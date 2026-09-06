package shodan

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	c := New("test-key")
	c.BaseURL = srv.URL
	c.BaseExploitsURL = srv.URL
	c.BaseTrendsURL = srv.URL
	c.APIRateLimit = 0
	c.HTTP = srv.Client()
	return c
}

func TestHost(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/shodan/host/8.8.8.8" {
			t.Errorf("path = %s", r.URL.Path)
		}
		if r.URL.Query().Get("key") != "test-key" {
			t.Error("missing api key")
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"ip_str": "8.8.8.8", "ports": []int{53}})
	})
	v, err := c.Host([]string{"8.8.8.8"}, false, false)
	if err != nil {
		t.Fatal(err)
	}
	m := v.(map[string]any)
	if m["ip_str"] != "8.8.8.8" {
		t.Fatalf("got %#v", m)
	}
}

func TestSearchAndCount(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/shodan/host/search":
			if r.URL.Query().Get("query") != "apache" {
				t.Errorf("query=%s", r.URL.Query().Get("query"))
			}
			_ = json.NewEncoder(w).Encode(map[string]any{
				"total":   2,
				"matches": []any{map[string]any{"ip_str": "1.1.1.1"}},
			})
		case "/shodan/host/count":
			_ = json.NewEncoder(w).Encode(map[string]any{"total": 42})
		default:
			http.NotFound(w, r)
		}
	})
	s, err := c.Search("apache", 1, 0, 0, nil, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s["total"].(float64) != 2 {
		t.Fatalf("total=%v", s["total"])
	}
	cnt, err := c.Count("apache", []any{"country"})
	if err != nil {
		t.Fatal(err)
	}
	if cnt["total"].(float64) != 42 {
		t.Fatalf("count=%v", cnt["total"])
	}
}

func TestAPIErrorJSON(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(map[string]any{"error": "Invalid API key"})
	})
	_, err := c.Info()
	if err == nil || !strings.Contains(err.Error(), "Invalid API key") {
		t.Fatalf("err=%v", err)
	}
}

func TestUnauthorizedHTML(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("<html>denied</html>"))
	})
	_, err := c.Info()
	if err == nil || err.Error() != "Invalid API key" {
		t.Fatalf("err=%v", err)
	}
}

func TestForbiddenAndBadGateway(t *testing.T) {
	n := 0
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		n++
		if n == 1 {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		w.WriteHeader(http.StatusBadGateway)
	})
	_, err := c.Info()
	if err == nil || !strings.Contains(err.Error(), "403") {
		t.Fatalf("err=%v", err)
	}
	_, err = c.Info()
	if err == nil || !strings.Contains(err.Error(), "502") {
		t.Fatalf("err=%v", err)
	}
}

func TestCreateAlertJSONBody(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method=%s", r.Method)
		}
		if r.URL.Path != "/shodan/alert" {
			t.Errorf("path=%s", r.URL.Path)
		}
		b, _ := io.ReadAll(r.Body)
		var body map[string]any
		if err := json.Unmarshal(b, &body); err != nil {
			t.Fatal(err)
		}
		if body["name"] != "net" {
			t.Errorf("body=%s", b)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"id": "A1", "name": "net"})
	})
	v, err := c.CreateAlert("net", []string{"1.2.3.0/24"}, 0)
	if err != nil {
		t.Fatal(err)
	}
	if v["id"] != "A1" {
		t.Fatalf("%v", v)
	}
}

func TestExploitsAndTrends(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/api/search":
			_ = json.NewEncoder(w).Encode(map[string]any{"total": 1, "matches": []any{}})
		case "/api/v1/search":
			_ = json.NewEncoder(w).Encode(map[string]any{"total": 9})
		default:
			http.NotFound(w, r)
		}
	})
	e, err := c.Exploits.Search("apache", 1, nil)
	if err != nil {
		t.Fatal(err)
	}
	if e["total"].(float64) != 1 {
		t.Fatal(e)
	}
	tr, err := c.Trends.Search("apache", []any{"country"})
	if err != nil {
		t.Fatal(err)
	}
	if tr["total"].(float64) != 9 {
		t.Fatal(tr)
	}
}

func TestOrgMember(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			t.Errorf("method=%s", r.Method)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"success": true})
	})
	ok, err := c.Org.AddMember("bob", true)
	if err != nil || !ok {
		t.Fatalf("ok=%v err=%v", ok, err)
	}
}

func TestCreateFacetString(t *testing.T) {
	got := CreateFacetString([]any{"country", Facet{Name: "port", Count: 5}})
	if got != "country,port:5" {
		t.Fatalf("got %q", got)
	}
}

func TestHumanizeBytes(t *testing.T) {
	if HumanizeBytes(1, 1) != "1 byte" {
		t.Fatal(HumanizeBytes(1, 1))
	}
	if HumanizeBytes(1024, 1) != "1.0 KB" {
		t.Fatal(HumanizeBytes(1024, 1))
	}
}

func TestGetIPAndScreenshot(t *testing.T) {
	b := map[string]any{"ip_str": "1.2.3.4"}
	if GetIP(b) != "1.2.3.4" {
		t.Fatal(GetIP(b))
	}
	b["ipv6"] = "fe80::1"
	if GetIP(b) != "fe80::1" {
		t.Fatal(GetIP(b))
	}
	if GetScreenshot(b) != nil {
		t.Fatal("expected nil screenshot")
	}
	b["opts"] = map[string]any{"screenshot": "img"}
	if GetScreenshot(b) != "img" {
		t.Fatal(GetScreenshot(b))
	}
}

func TestIterateFiles(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "banners.json")
	if err := os.WriteFile(p, []byte("{\"ip_str\":\"9.9.9.9\"}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ch, errc := IterateFiles([]string{p})
	n := 0
	for b := range ch {
		n++
		if b["ip_str"] != "9.9.9.9" {
			t.Fatalf("%v", b)
		}
	}
	if err := <-errc; err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("n=%d", n)
	}
}

func TestIgnoreVulnerableTrigger(t *testing.T) {
	c := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		want := "/shodan/alert/A/trigger/vulnerable/ignore/1.1.1.1:443/CVE-1,CVE-2"
		if r.URL.Path != want {
			t.Errorf("path=%s", r.URL.Path)
		}
		if r.Method != http.MethodPut {
			t.Errorf("method=%s", r.Method)
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"success": true})
	})
	_, err := c.IgnoreAlertTriggerNotification("A", "vulnerable", "1.1.1.1", 443, []string{"CVE-1", "CVE-2"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestStreamBanners(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/shodan/banners" {
			t.Errorf("path=%s", r.URL.Path)
		}
		_, _ = w.Write([]byte("{\"ip_str\":\"8.8.8.8\"}\n\n{\"ip_str\":\"1.1.1.1\"}\n"))
	}))
	t.Cleanup(srv.Close)
	s := newStream("k", srv.Client())
	s.BaseURL = srv.URL
	ch, errc, closer, err := s.Banners(0, false)
	if err != nil {
		t.Fatal(err)
	}
	defer closer.Close()
	var ips []string
	for ev := range ch {
		ips = append(ips, ev.Object["ip_str"].(string))
	}
	if err := <-errc; err != nil {
		t.Fatal(err)
	}
	if len(ips) != 2 || ips[0] != "8.8.8.8" {
		t.Fatalf("%v", ips)
	}
}

func TestAPIErrorString(t *testing.T) {
	e := &APIError{Value: "boom"}
	if e.Error() != "boom" {
		t.Fatal(e.Error())
	}
}
