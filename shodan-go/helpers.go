package shodan

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// Facet is either a facet name or a (name, count) pair.
type Facet struct {
	Name  string
	Count int // 0 means no count suffix
}

// CreateFacetString converts facets into a comma-separated string understood by the Shodan API.
func CreateFacetString(facets []any) string {
	parts := make([]string, 0, len(facets))
	for _, facet := range facets {
		switch v := facet.(type) {
		case string:
			parts = append(parts, v)
		case Facet:
			if v.Count > 0 {
				parts = append(parts, fmt.Sprintf("%s:%d", v.Name, v.Count))
			} else {
				parts = append(parts, v.Name)
			}
		case [2]any:
			parts = append(parts, fmt.Sprintf("%v:%v", v[0], v[1]))
		default:
			parts = append(parts, fmt.Sprint(v))
		}
	}
	return strings.Join(parts, ",")
}

// IterateFiles loops over all records of the provided Shodan output file(s).
func IterateFiles(files []string) (<-chan map[string]any, <-chan error) {
	out := make(chan map[string]any)
	errc := make(chan error, 1)
	go func() {
		defer close(out)
		defer close(errc)
		for _, filename := range files {
			if err := iterateOneFile(filename, out); err != nil {
				errc <- err
				return
			}
		}
	}()
	return out, errc
}

func iterateOneFile(filename string, out chan<- map[string]any) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	var r io.Reader = f
	if strings.HasSuffix(filename, ".gz") {
		gz, err := gzip.NewReader(f)
		if err != nil {
			return err
		}
		defer gz.Close()
		r = gz
	}

	sc := bufio.NewScanner(r)
	buf := make([]byte, 0, 1024*1024)
	sc.Buffer(buf, 16*1024*1024)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		var banner map[string]any
		if err := json.Unmarshal([]byte(line), &banner); err != nil {
			return err
		}
		out <- banner
	}
	return sc.Err()
}

// GetScreenshot extracts a screenshot object from a banner if present.
func GetScreenshot(banner map[string]any) any {
	if s, ok := banner["screenshot"]; ok && s != nil {
		return s
	}
	if opts, ok := banner["opts"].(map[string]any); ok {
		if s, ok := opts["screenshot"]; ok {
			return s
		}
	}
	return nil
}

// GetIP returns the IPv6 or IPv4 address from a banner.
func GetIP(banner map[string]any) string {
	if v, ok := banner["ipv6"].(string); ok && v != "" {
		return v
	}
	if v, ok := banner["ip_str"].(string); ok {
		return v
	}
	return ""
}

// OpenFile opens a gzip file for writing banners.
func OpenFile(filename string) (*gzip.Writer, *os.File, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, nil, err
	}
	gw, err := gzip.NewWriterLevel(f, gzip.BestCompression)
	if err != nil {
		f.Close()
		return nil, nil, err
	}
	return gw, f, nil
}

// WriteBanner writes a banner as a JSON line.
func WriteBanner(w io.Writer, banner any) error {
	b, err := json.Marshal(banner)
	if err != nil {
		return err
	}
	_, err = w.Write(append(b, '\n'))
	return err
}

// HumanizeBytes returns a humanized string representation of a number of bytes.
func HumanizeBytes(byteCount float64, precision int) string {
	if byteCount == 1 {
		return "1 byte"
	}
	if byteCount < 1024 {
		return fmt.Sprintf("%.0f bytes", byteCount)
	}
	suffixes := []string{"KB", "MB", "GB", "TB", "PB"}
	multiple := 1024.0
	for _, suffix := range suffixes {
		byteCount /= multiple
		if byteCount < multiple {
			return fmt.Sprintf("%.*f %s", precision, byteCount, suffix)
		}
	}
	return fmt.Sprintf("%.*f %s", precision, byteCount, suffixes[len(suffixes)-1])
}
