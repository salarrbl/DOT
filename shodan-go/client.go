package shodan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultBaseURL         = "https://api.shodan.io"
	defaultExploitsBaseURL = "https://exploits.shodan.io"
	defaultTrendsBaseURL   = "https://trends.shodan.io"
)

// Client is a wrapper around the Shodan REST and Streaming APIs.
type Client struct {
	APIKey           string
	BaseURL          string
	BaseExploitsURL  string
	BaseTrendsURL    string
	HTTP             *http.Client
	APIRateLimit     float64 // requests per second; 0 disables
	Data             *DataAPI
	DNS              *DNSAPI
	Exploits         *ExploitsAPI
	Trends           *TrendsAPI
	Labs             *LabsAPI
	Notifier         *NotifierAPI
	Org              *OrganizationAPI
	Tools            *ToolsAPI
	Stream           *Stream

	mu           sync.Mutex
	lastQueryAt  time.Time
}

// New creates a Shodan API client.
func New(apiKey string) *Client {
	c := &Client{
		APIKey:          apiKey,
		BaseURL:         defaultBaseURL,
		BaseExploitsURL: defaultExploitsBaseURL,
		BaseTrendsURL:   defaultTrendsBaseURL,
		HTTP:            &http.Client{Timeout: 60 * time.Second},
		APIRateLimit:    1,
	}
	if u := os.Getenv("SHODAN_API_URL"); u != "" {
		c.BaseURL = u
	}
	c.Data = &DataAPI{parent: c}
	c.DNS = &DNSAPI{parent: c}
	c.Exploits = &ExploitsAPI{parent: c}
	c.Trends = &TrendsAPI{parent: c}
	c.Labs = &LabsAPI{parent: c}
	c.Notifier = &NotifierAPI{parent: c}
	c.Org = &OrganizationAPI{parent: c}
	c.Tools = &ToolsAPI{parent: c}
	c.Stream = newStream(apiKey, c.HTTP)
	return c
}

// WithHTTPClient sets a custom HTTP client (proxies, TLS, etc.).
func (c *Client) WithHTTPClient(h *http.Client) *Client {
	c.HTTP = h
	c.Stream.HTTP = h
	return c
}

func (c *Client) waitRate() {
	if c.APIRateLimit <= 0 {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	minInterval := time.Duration(float64(time.Second) / c.APIRateLimit)
	if !c.lastQueryAt.IsZero() {
		for time.Since(c.lastQueryAt) < minInterval {
			time.Sleep(minInterval / 10)
		}
	}
	c.lastQueryAt = time.Now()
}

func (c *Client) baseFor(service string) string {
	switch service {
	case "exploits":
		return c.BaseExploitsURL
	case "trends":
		return c.BaseTrendsURL
	default:
		return c.BaseURL
	}
}

func queryEscape(s string) string {
	return url.QueryEscape(s)
}

func (c *Client) request(function string, params url.Values, service, method string, jsonData any) (any, error) {
	if params == nil {
		params = url.Values{}
	}
	params.Set("key", c.APIKey)

	c.waitRate()

	base := c.baseFor(service)
	u := base + function
	var body io.Reader
	reqURL := u
	m := strings.ToUpper(method)
	if m == "" {
		m = http.MethodGet
	}

	if m == http.MethodPost && jsonData != nil {
		b, err := json.Marshal(jsonData)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(b)
		reqURL = u + "?" + params.Encode()
	} else if m == http.MethodPost && jsonData == nil {
		// Python posts params as form body
		body = strings.NewReader(params.Encode())
	} else {
		reqURL = u + "?" + params.Encode()
	}

	req, err := http.NewRequest(m, reqURL, body)
	if err != nil {
		return nil, newAPIError("Unable to connect to Shodan")
	}
	if m == http.MethodPost && jsonData != nil {
		req.Header.Set("Content-Type", "application/json")
	} else if m == http.MethodPost {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, newAPIError("Unable to connect to Shodan")
	}
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, newAPIError("Unable to connect to Shodan")
	}

	if resp.StatusCode == http.StatusUnauthorized {
		var obj map[string]any
		if json.Unmarshal(raw, &obj) == nil {
			if e, ok := obj["error"].(string); ok {
				return nil, newAPIError(e)
			}
		}
		if len(raw) > 0 && raw[0] == '<' {
			return nil, newAPIError("Invalid API key")
		}
		return nil, newAPIError("Invalid API key")
	}
	if resp.StatusCode == http.StatusForbidden {
		return nil, newAPIError("Access denied (403 Forbidden)")
	}
	if resp.StatusCode == http.StatusBadGateway {
		return nil, newAPIError("Bad Gateway (502)")
	}

	var data any
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, newAPIError("Unable to parse JSON response")
	}
	if m, ok := data.(map[string]any); ok {
		if e, ok := m["error"].(string); ok && e != "" {
			return nil, newAPIError(e)
		}
	}
	return data, nil
}

func asMap(v any, err error) (map[string]any, error) {
	if err != nil {
		return nil, err
	}
	m, ok := v.(map[string]any)
	if !ok {
		return nil, newAPIError("unexpected response type")
	}
	return m, nil
}

func asAny(v any, err error) (any, error) {
	return v, err
}

// Count returns the total number of search results for the query.
func (c *Client) Count(query string, facets []any) (map[string]any, error) {
	p := url.Values{"query": {query}}
	if len(facets) > 0 {
		p.Set("facets", CreateFacetString(facets))
	}
	return asMap(c.request("/shodan/host/count", p, "shodan", http.MethodGet, nil))
}

// Host returns all available information on one or more IPs.
func (c *Client) Host(ips []string, history, minify bool) (any, error) {
	p := url.Values{}
	if history {
		p.Set("history", "true")
	}
	if minify {
		p.Set("minify", "true")
	}
	return asAny(c.request("/shodan/host/"+strings.Join(ips, ","), p, "shodan", http.MethodGet, nil))
}

// Info returns information about the current API key.
func (c *Client) Info() (map[string]any, error) {
	return asMap(c.request("/api-info", url.Values{}, "shodan", http.MethodGet, nil))
}

// Ports returns the list of ports that Shodan crawls.
func (c *Client) Ports() (any, error) {
	return asAny(c.request("/shodan/ports", url.Values{}, "shodan", http.MethodGet, nil))
}

// Protocols returns protocols the on-demand scanning API supports.
func (c *Client) Protocols() (any, error) {
	return asAny(c.request("/shodan/protocols", url.Values{}, "shodan", http.MethodGet, nil))
}

// Scan submits IPs or a structured scan request.
func (c *Client) Scan(ips any, force bool) (map[string]any, error) {
	var networks string
	switch v := ips.(type) {
	case string:
		networks = v
	case []string:
		networks = strings.Join(v, ",")
	default:
		b, err := json.Marshal(ips)
		if err != nil {
			return nil, err
		}
		networks = string(b)
	}
	p := url.Values{
		"ips":   {networks},
		"force": {strconv.FormatBool(force)},
	}
	return asMap(c.request("/shodan/scan", p, "shodan", http.MethodPost, nil))
}

// Scans lists submitted scans.
func (c *Client) Scans(page int) (map[string]any, error) {
	if page <= 0 {
		page = 1
	}
	return asMap(c.request("/shodan/scans", url.Values{"page": {strconv.Itoa(page)}}, "shodan", http.MethodGet, nil))
}

// ScanInternet starts an internet-wide scan for a port/protocol.
func (c *Client) ScanInternet(port int, protocol string) (map[string]any, error) {
	p := url.Values{
		"port":     {strconv.Itoa(port)},
		"protocol": {protocol},
	}
	return asMap(c.request("/shodan/scan/internet", p, "shodan", http.MethodPost, nil))
}

// ScanStatus returns status of a previously submitted scan.
func (c *Client) ScanStatus(scanID string) (map[string]any, error) {
	return asMap(c.request("/shodan/scan/"+scanID, url.Values{}, "shodan", http.MethodGet, nil))
}

// Search searches the Shodan database.
func (c *Client) Search(query string, page, limit, offset int, facets []any, minify bool, fields []string) (map[string]any, error) {
	p := url.Values{
		"query":  {query},
		"minify": {strconv.FormatBool(minify)},
	}
	if limit > 0 {
		p.Set("limit", strconv.Itoa(limit))
		if offset > 0 {
			p.Set("offset", strconv.Itoa(offset))
		}
	} else {
		if page <= 0 {
			page = 1
		}
		p.Set("page", strconv.Itoa(page))
	}
	if len(facets) > 0 {
		p.Set("facets", CreateFacetString(facets))
	}
	if len(fields) > 0 {
		p.Set("fields", strings.Join(fields, ","))
	}
	return asMap(c.request("/shodan/host/search", p, "shodan", http.MethodGet, nil))
}

// SearchCursor iterates over all search results.
func (c *Client) SearchCursor(query string, minify bool, retries int, fields []string) (<-chan map[string]any, <-chan error) {
	out := make(chan map[string]any)
	errc := make(chan error, 1)
	if retries <= 0 {
		retries = 5
	}
	go func() {
		defer close(out)
		defer close(errc)
		page := 1
		results, err := c.Search(query, page, 0, 0, nil, minify, fields)
		if err != nil {
			errc <- err
			return
		}
		total := 0.0
		if t, ok := results["total"].(float64); ok {
			total = t
		}
		totalPages := 0
		if total > 0 {
			totalPages = int(math.Ceil(total / 100))
		}
		if matches, ok := results["matches"].([]any); ok {
			for _, m := range matches {
				if banner, ok := m.(map[string]any); ok {
					out <- banner
				}
			}
		}
		page++
		tries := 0
		for page <= totalPages {
			results, err = c.Search(query, page, 0, 0, nil, minify, fields)
			if err != nil {
				if tries >= retries {
					errc <- newAPIErrorf("Retry limit reached (%d)", retries)
					return
				}
				tries++
				time.Sleep(time.Duration(tries) * time.Second)
				continue
			}
			if matches, ok := results["matches"].([]any); ok {
				for _, m := range matches {
					if banner, ok := m.(map[string]any); ok {
						out <- banner
					}
				}
			}
			page++
			tries = 0
		}
	}()
	return out, errc
}

// SearchFacets returns available search facets.
func (c *Client) SearchFacets() (any, error) {
	return asAny(c.request("/shodan/host/search/facets", url.Values{}, "shodan", http.MethodGet, nil))
}

// SearchFilters returns available search filters.
func (c *Client) SearchFilters() (any, error) {
	return asAny(c.request("/shodan/host/search/filters", url.Values{}, "shodan", http.MethodGet, nil))
}

// SearchTokens returns information about a search query.
func (c *Client) SearchTokens(query string) (map[string]any, error) {
	return asMap(c.request("/shodan/host/search/tokens", url.Values{"query": {query}}, "shodan", http.MethodGet, nil))
}

// Services returns ports/services that Shodan crawls.
func (c *Client) Services() (any, error) {
	return asAny(c.request("/shodan/services", url.Values{}, "shodan", http.MethodGet, nil))
}

// Queries lists shared search queries.
func (c *Client) Queries(page int, sort, order string) (map[string]any, error) {
	if page <= 0 {
		page = 1
	}
	if sort == "" {
		sort = "timestamp"
	}
	if order == "" {
		order = "desc"
	}
	p := url.Values{
		"page":  {strconv.Itoa(page)},
		"sort":  {sort},
		"order": {order},
	}
	return asMap(c.request("/shodan/query", p, "shodan", http.MethodGet, nil))
}

// QueriesSearch searches the directory of saved queries.
func (c *Client) QueriesSearch(query string, page int) (map[string]any, error) {
	if page <= 0 {
		page = 1
	}
	p := url.Values{"page": {strconv.Itoa(page)}, "query": {query}}
	return asMap(c.request("/shodan/query/search", p, "shodan", http.MethodGet, nil))
}

// QueriesTags returns popular query tags.
func (c *Client) QueriesTags(size int) (map[string]any, error) {
	if size <= 0 {
		size = 10
	}
	return asMap(c.request("/shodan/query/tags", url.Values{"size": {strconv.Itoa(size)}}, "shodan", http.MethodGet, nil))
}

// CreateAlert creates a network alert for IP range(s).
func (c *Client) CreateAlert(name string, ip any, expires int) (map[string]any, error) {
	data := map[string]any{
		"name": name,
		"filters": map[string]any{
			"ip": ip,
		},
		"expires": expires,
	}
	return asMap(c.request("/shodan/alert", url.Values{}, "shodan", http.MethodPost, data))
}

// EditAlert edits IPs monitored by an alert.
func (c *Client) EditAlert(aid string, ip any) (map[string]any, error) {
	data := map[string]any{
		"filters": map[string]any{"ip": ip},
	}
	return asMap(c.request("/shodan/alert/"+aid, url.Values{}, "shodan", http.MethodPost, data))
}

// Alerts lists alerts, or a single alert if aid is set.
func (c *Client) Alerts(aid string, includeExpired bool) (any, error) {
	fn := "/shodan/alert/info"
	if aid != "" {
		fn = "/shodan/alert/" + aid + "/info"
	}
	p := url.Values{"include_expired": {strconv.FormatBool(includeExpired)}}
	return asAny(c.request(fn, p, "shodan", http.MethodGet, nil))
}

// DeleteAlert deletes an alert.
func (c *Client) DeleteAlert(aid string) (any, error) {
	return asAny(c.request("/shodan/alert/"+aid, url.Values{}, "shodan", http.MethodDelete, nil))
}

// AlertTriggers returns available alert triggers.
func (c *Client) AlertTriggers() (any, error) {
	return asAny(c.request("/shodan/alert/triggers", url.Values{}, "shodan", http.MethodGet, nil))
}

// EnableAlertTrigger enables a trigger on an alert.
func (c *Client) EnableAlertTrigger(aid, trigger string) (any, error) {
	return asAny(c.request(fmt.Sprintf("/shodan/alert/%s/trigger/%s", aid, trigger), url.Values{}, "shodan", http.MethodPut, nil))
}

// DisableAlertTrigger disables a trigger on an alert.
func (c *Client) DisableAlertTrigger(aid, trigger string) (any, error) {
	return asAny(c.request(fmt.Sprintf("/shodan/alert/%s/trigger/%s", aid, trigger), url.Values{}, "shodan", http.MethodDelete, nil))
}

// IgnoreAlertTriggerNotification ignores trigger notifications for IP:port.
func (c *Client) IgnoreAlertTriggerNotification(aid, trigger, ip string, port int, vulns []string) (any, error) {
	if (trigger == "vulnerable" || trigger == "vulnerable_unverified") && len(vulns) > 0 {
		path := fmt.Sprintf("/shodan/alert/%s/trigger/%s/ignore/%s:%d/%s", aid, trigger, ip, port, strings.Join(vulns, ","))
		return asAny(c.request(path, url.Values{}, "shodan", http.MethodPut, nil))
	}
	path := fmt.Sprintf("/shodan/alert/%s/trigger/%s/ignore/%s:%d", aid, trigger, ip, port)
	return asAny(c.request(path, url.Values{}, "shodan", http.MethodPut, nil))
}

// UnignoreAlertTriggerNotification re-enables trigger notifications.
func (c *Client) UnignoreAlertTriggerNotification(aid, trigger, ip string, port int) (any, error) {
	path := fmt.Sprintf("/shodan/alert/%s/trigger/%s/ignore/%s:%d", aid, trigger, ip, port)
	return asAny(c.request(path, url.Values{}, "shodan", http.MethodDelete, nil))
}

// AddAlertNotifier enables a notifier for an alert.
func (c *Client) AddAlertNotifier(aid, nid string) (any, error) {
	return asAny(c.request(fmt.Sprintf("/shodan/alert/%s/notifier/%s", aid, nid), url.Values{}, "shodan", http.MethodPut, nil))
}

// RemoveAlertNotifier removes a notifier from an alert.
func (c *Client) RemoveAlertNotifier(aid, nid string) (any, error) {
	return asAny(c.request(fmt.Sprintf("/shodan/alert/%s/notifier/%s", aid, nid), url.Values{}, "shodan", http.MethodDelete, nil))
}

// DataAPI wraps dataset download endpoints.
type DataAPI struct{ parent *Client }

func (d *DataAPI) ListDatasets() (any, error) {
	return asAny(d.parent.request("/shodan/data", url.Values{}, "shodan", http.MethodGet, nil))
}

func (d *DataAPI) ListFiles(dataset string) (any, error) {
	return asAny(d.parent.request("/shodan/data/"+dataset, url.Values{}, "shodan", http.MethodGet, nil))
}

// DNSAPI wraps DNS endpoints.
type DNSAPI struct{ parent *Client }

func (d *DNSAPI) DomainInfo(domain string, history bool, recType string, page int) (map[string]any, error) {
	if page <= 0 {
		page = 1
	}
	p := url.Values{"page": {strconv.Itoa(page)}}
	if history {
		p.Set("history", "true")
	}
	if recType != "" {
		p.Set("type", recType)
	}
	return asMap(d.parent.request("/dns/domain/"+domain, p, "shodan", http.MethodGet, nil))
}

// NotifierAPI wraps notifier endpoints.
type NotifierAPI struct{ parent *Client }

func (n *NotifierAPI) Create(provider string, args map[string]any, description string) (map[string]any, error) {
	if args == nil {
		args = map[string]any{}
	}
	p := url.Values{}
	for k, v := range args {
		p.Set(k, fmt.Sprint(v))
	}
	p.Set("provider", provider)
	if description != "" {
		p.Set("description", description)
	}
	return asMap(n.parent.request("/notifier", p, "shodan", http.MethodPost, nil))
}

func (n *NotifierAPI) Edit(nid string, args map[string]any) (map[string]any, error) {
	p := url.Values{}
	for k, v := range args {
		p.Set(k, fmt.Sprint(v))
	}
	return asMap(n.parent.request("/notifier/"+nid, p, "shodan", http.MethodPut, nil))
}

func (n *NotifierAPI) Get(nid string) (map[string]any, error) {
	return asMap(n.parent.request("/notifier/"+nid, url.Values{}, "shodan", http.MethodGet, nil))
}

func (n *NotifierAPI) ListNotifiers() (any, error) {
	return asAny(n.parent.request("/notifier", url.Values{}, "shodan", http.MethodGet, nil))
}

func (n *NotifierAPI) ListProviders() (any, error) {
	return asAny(n.parent.request("/notifier/provider", url.Values{}, "shodan", http.MethodGet, nil))
}

func (n *NotifierAPI) Remove(nid string) (any, error) {
	return asAny(n.parent.request("/notifier/"+nid, url.Values{}, "shodan", http.MethodDelete, nil))
}

// ToolsAPI wraps utility endpoints.
type ToolsAPI struct{ parent *Client }

func (t *ToolsAPI) MyIP() (any, error) {
	return asAny(t.parent.request("/tools/myip", url.Values{}, "shodan", http.MethodGet, nil))
}

// ExploitsAPI wraps the Exploits REST API.
type ExploitsAPI struct{ parent *Client }

func (e *ExploitsAPI) Search(query string, page int, facets []any) (map[string]any, error) {
	if page <= 0 {
		page = 1
	}
	p := url.Values{"query": {query}, "page": {strconv.Itoa(page)}}
	if len(facets) > 0 {
		p.Set("facets", CreateFacetString(facets))
	}
	return asMap(e.parent.request("/api/search", p, "exploits", http.MethodGet, nil))
}

func (e *ExploitsAPI) Count(query string, facets []any) (map[string]any, error) {
	p := url.Values{"query": {query}}
	if len(facets) > 0 {
		p.Set("facets", CreateFacetString(facets))
	}
	return asMap(e.parent.request("/api/count", p, "exploits", http.MethodGet, nil))
}

// LabsAPI wraps experimental endpoints.
type LabsAPI struct{ parent *Client }

func (l *LabsAPI) HoneyScore(ip string) (any, error) {
	return asAny(l.parent.request("/labs/honeyscore/"+ip, url.Values{}, "shodan", http.MethodGet, nil))
}

// OrganizationAPI wraps organization endpoints.
type OrganizationAPI struct{ parent *Client }

func (o *OrganizationAPI) AddMember(user string, notify bool) (bool, error) {
	p := url.Values{"notify": {strconv.FormatBool(notify)}}
	m, err := asMap(o.parent.request("/org/member/"+user, p, "shodan", http.MethodPut, nil))
	if err != nil {
		return false, err
	}
	ok, _ := m["success"].(bool)
	return ok, nil
}

func (o *OrganizationAPI) Info() (map[string]any, error) {
	return asMap(o.parent.request("/org", url.Values{}, "shodan", http.MethodGet, nil))
}

func (o *OrganizationAPI) RemoveMember(user string) (bool, error) {
	m, err := asMap(o.parent.request("/org/member/"+user, url.Values{}, "shodan", http.MethodDelete, nil))
	if err != nil {
		return false, err
	}
	ok, _ := m["success"].(bool)
	return ok, nil
}

// TrendsAPI wraps the Trends REST API.
type TrendsAPI struct{ parent *Client }

func (t *TrendsAPI) Search(query string, facets []any) (map[string]any, error) {
	p := url.Values{"query": {query}}
	if len(facets) > 0 {
		p.Set("facets", CreateFacetString(facets))
	}
	return asMap(t.parent.request("/api/v1/search", p, "trends", http.MethodGet, nil))
}

func (t *TrendsAPI) SearchFacets() (any, error) {
	return asAny(t.parent.request("/api/v1/search/facets", url.Values{}, "trends", http.MethodGet, nil))
}

func (t *TrendsAPI) SearchFilters() (any, error) {
	return asAny(t.parent.request("/api/v1/search/filters", url.Values{}, "trends", http.MethodGet, nil))
}
