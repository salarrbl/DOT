package shodan

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const defaultStreamURL = "https://stream.shodan.io"

// Stream provides access to the Shodan Streaming API.
type Stream struct {
	APIKey  string
	BaseURL string
	HTTP    *http.Client
}

func newStream(apiKey string, httpClient *http.Client) *Stream {
	return &Stream{
		APIKey:  apiKey,
		BaseURL: defaultStreamURL,
		HTTP:    httpClient,
	}
}

func (s *Stream) createStream(name string, query string, timeout time.Duration) (io.ReadCloser, error) {
	params := url.Values{}
	params.Set("key", s.APIKey)
	if timeout > 0 {
		params.Set("heartbeat", "false")
	}
	if query != "" {
		params.Set("query", query)
	}

	u := s.BaseURL + name + "?" + params.Encode()
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, newAPIError("Unable to contact the Shodan Streaming API")
	}

	client := s.HTTP
	if timeout > 0 {
		c := *s.HTTP
		c.Timeout = timeout
		client = &c
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, newAPIError("Unable to contact the Shodan Streaming API")
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		var data map[string]any
		_ = json.NewDecoder(resp.Body).Decode(&data)
		if errMsg, ok := data["error"].(string); ok && errMsg != "" {
			return nil, newAPIError(errMsg)
		}
		return nil, newAPIError("Invalid API key or you do not have access to the Streaming API")
	}
	return resp.Body, nil
}

// BannerEvent is a decoded stream line, or raw bytes when Raw is true.
type BannerEvent struct {
	Raw    []byte
	Object map[string]any
}

func iterStream(body io.Reader, raw bool) (<-chan BannerEvent, <-chan error) {
	out := make(chan BannerEvent)
	errc := make(chan error, 1)
	go func() {
		defer close(out)
		defer close(errc)
		sc := bufio.NewScanner(body)
		buf := make([]byte, 0, 1024*1024)
		sc.Buffer(buf, 16*1024*1024)
		for sc.Scan() {
			line := sc.Bytes()
			if len(line) == 0 {
				continue
			}
			if raw {
				cp := make([]byte, len(line))
				copy(cp, line)
				out <- BannerEvent{Raw: cp}
				continue
			}
			var obj map[string]any
			if err := json.Unmarshal(line, &obj); err != nil {
				errc <- err
				return
			}
			out <- BannerEvent{Object: obj}
		}
		if err := sc.Err(); err != nil {
			errc <- newAPIError("Stream timed out")
		}
	}()
	return out, errc
}

func (s *Stream) consume(path, query string, timeout time.Duration, raw bool) (<-chan BannerEvent, <-chan error, io.Closer, error) {
	body, err := s.createStream(path, query, timeout)
	if err != nil {
		return nil, nil, nil, err
	}
	ch, errc := iterStream(body, raw)
	return ch, errc, body, nil
}

// Alert streams network alerts. If aid is empty, all alerts are streamed.
func (s *Stream) Alert(aid string, timeout time.Duration, raw bool) (<-chan BannerEvent, <-chan error, io.Closer, error) {
	path := "/shodan/alert"
	if aid != "" {
		path = "/shodan/alert/" + aid
	}
	return s.consume(path, "", timeout, raw)
}

// ASN streams banners matching the given ASNs.
func (s *Stream) ASN(asn []string, timeout time.Duration, raw bool) (<-chan BannerEvent, <-chan error, io.Closer, error) {
	return s.consume("/shodan/asn/"+strings.Join(asn, ","), "", timeout, raw)
}

// Banners is a real-time feed of data Shodan is currently collecting.
func (s *Stream) Banners(timeout time.Duration, raw bool) (<-chan BannerEvent, <-chan error, io.Closer, error) {
	return s.consume("/shodan/banners", "", timeout, raw)
}

// Countries streams banners matching the given country codes.
func (s *Stream) Countries(countries []string, timeout time.Duration, raw bool) (<-chan BannerEvent, <-chan error, io.Closer, error) {
	return s.consume("/shodan/countries/"+strings.Join(countries, ","), "", timeout, raw)
}

// Custom streams banners matching a custom query.
func (s *Stream) Custom(query string, timeout time.Duration, raw bool) (<-chan BannerEvent, <-chan error, io.Closer, error) {
	return s.consume("/shodan/custom", query, timeout, raw)
}

// Ports streams banners matching the given ports.
func (s *Stream) Ports(ports []int, timeout time.Duration, raw bool) (<-chan BannerEvent, <-chan error, io.Closer, error) {
	ss := make([]string, len(ports))
	for i, p := range ports {
		ss[i] = strconv.Itoa(p)
	}
	return s.consume("/shodan/ports/"+strings.Join(ss, ","), "", timeout, raw)
}

// Tags streams banners matching the given tags.
func (s *Stream) Tags(tags []string, timeout time.Duration, raw bool) (<-chan BannerEvent, <-chan error, io.Closer, error) {
	return s.consume("/shodan/tags/"+strings.Join(tags, ","), "", timeout, raw)
}

// Vulns streams banners matching the given vulnerabilities.
func (s *Stream) Vulns(vulns []string, timeout time.Duration, raw bool) (<-chan BannerEvent, <-chan error, io.Closer, error) {
	return s.consume("/shodan/vulns/"+strings.Join(vulns, ","), "", timeout, raw)
}

// StreamPath is exported for tests.
func (s *Stream) StreamPath(name string) string {
	return fmt.Sprintf("%s%s", s.BaseURL, name)
}
