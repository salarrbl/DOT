package shodan

import (
	"bufio"
	"encoding/json"
	"io"
	"net/http"
)

// Threatnet wraps the Threatnet REST and Streaming APIs.
type Threatnet struct {
	APIKey  string
	BaseURL string
	Stream  *ThreatnetStream
	HTTP    *http.Client
}

// ThreatnetStream provides Threatnet streaming endpoints.
type ThreatnetStream struct {
	parent  *Threatnet
	BaseURL string
}

// NewThreatnet creates a Threatnet client.
func NewThreatnet(apiKey string) *Threatnet {
	t := &Threatnet{
		APIKey:  apiKey,
		BaseURL: defaultBaseURL,
		HTTP:    &http.Client{},
	}
	t.Stream = &ThreatnetStream{parent: t, BaseURL: defaultStreamURL}
	return t
}

func (s *ThreatnetStream) createStream(name string) (io.ReadCloser, error) {
	u := s.BaseURL + name + "?key=" + urlQueryEscape(s.parent.APIKey)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, newAPIError("Unable to contact the Shodan Streaming API")
	}
	resp, err := s.parent.HTTP.Do(req)
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

func urlQueryEscape(s string) string {
	return queryEscape(s)
}

func iterJSONLines(body io.Reader) (<-chan map[string]any, <-chan error) {
	out := make(chan map[string]any)
	errc := make(chan error, 1)
	go func() {
		defer close(out)
		defer close(errc)
		sc := bufio.NewScanner(body)
		for sc.Scan() {
			line := sc.Bytes()
			if len(line) == 0 {
				continue
			}
			var banner map[string]any
			if err := json.Unmarshal(line, &banner); err != nil {
				errc <- err
				return
			}
			out <- banner
		}
		if err := sc.Err(); err != nil {
			errc <- err
		}
	}()
	return out, errc
}

// Events streams Threatnet events.
func (s *ThreatnetStream) Events() (<-chan map[string]any, <-chan error, io.Closer, error) {
	body, err := s.createStream("/threatnet/events")
	if err != nil {
		return nil, nil, nil, err
	}
	ch, errc := iterJSONLines(body)
	return ch, errc, body, nil
}

// Backscatter streams Threatnet backscatter.
func (s *ThreatnetStream) Backscatter() (<-chan map[string]any, <-chan error, io.Closer, error) {
	body, err := s.createStream("/threatnet/backscatter")
	if err != nil {
		return nil, nil, nil, err
	}
	ch, errc := iterJSONLines(body)
	return ch, errc, body, nil
}

// Activity streams Threatnet SSH activity.
func (s *ThreatnetStream) Activity() (<-chan map[string]any, <-chan error, io.Closer, error) {
	body, err := s.createStream("/threatnet/ssh")
	if err != nil {
		return nil, nil, nil, err
	}
	ch, errc := iterJSONLines(body)
	return ch, errc, body, nil
}
