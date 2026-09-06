package cli

import (
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	shodan "github.com/salarrbl/DOT/shodan-go"
)

type Converter interface {
	Process(files []string, fileSize int64) error
	Dirname() string
}

type csvConverter struct {
	fout   io.Writer
	fields []string
}

func NewCsvConverter(fout io.Writer, fields []string) *csvConverter {
	if len(fields) == 0 {
		fields = []string{
			"data", "hostnames", "ip", "ip_str", "ipv6", "org", "isp",
			"location.country_code", "location.city", "location.country_name",
			"location.latitude", "location.longitude", "os", "asn", "port",
			"tags", "timestamp", "transport", "product", "version", "vulns",
			"ssl.cipher.version", "ssl.cipher.bits", "ssl.cipher.name",
			"ssl.alpn", "ssl.versions", "ssl.cert.serial",
			"ssl.cert.fingerprint.sha1", "ssl.cert.fingerprint.sha256",
			"html", "title",
		}
	}
	return &csvConverter{fout: fout, fields: fields}
}

func (c *csvConverter) Dirname() string { return "" }

func (c *csvConverter) Process(files []string, fileSize int64) error {
	w := csv.NewWriter(c.fout)
	if err := w.Write(c.fields); err != nil {
		return err
	}
	ch, errc := shodan.IterateFiles(files)
	for banner := range ch {
		if vulns, ok := banner["vulns"].(map[string]any); ok {
			keys := make([]any, 0, len(vulns))
			for k := range vulns {
				keys = append(keys, k)
			}
			banner["vulns"] = keys
		}
		row := make([]string, len(c.fields))
		for i, field := range c.fields {
			row[i] = bannerField(banner, field)
		}
		_ = w.Write(row)
	}
	w.Flush()
	return <-errc
}

func bannerField(banner map[string]any, flatField string) string {
	v := GetBannerField(banner, flatField)
	if v == nil {
		return ""
	}
	if list, ok := v.([]any); ok {
		parts := make([]string, 0, len(list))
		for _, i := range list {
			parts = append(parts, fmt.Sprint(i))
		}
		return strings.Join(parts, ",")
	}
	return fmt.Sprint(v)
}

type geoJSONConverter struct{ fout io.Writer }

func NewGeoJSONConverter(fout io.Writer) *geoJSONConverter { return &geoJSONConverter{fout} }
func (c *geoJSONConverter) Dirname() string                 { return "" }

func (c *geoJSONConverter) Process(files []string, fileSize int64) error {
	_, _ = io.WriteString(c.fout, "{\n            \"type\": \"FeatureCollection\",\n            \"features\": [\n        ")
	unique := map[string]bool{}
	ch, errc := shodan.IterateFiles(files)
	for banner := range ch {
		ip := shodan.GetIP(banner)
		if ip == "" || unique[ip] {
			continue
		}
		unique[ip] = true
		loc := asMap(banner["location"])
		if loc == nil {
			continue
		}
		lat, lon := loc["latitude"], loc["longitude"]
		feature := map[string]any{
			"type": "Feature",
			"id":   ip,
			"properties": map[string]any{"name": ip, "lat": lat, "lon": lon},
			"geometry": map[string]any{
				"type":        "Point",
				"coordinates": []any{lon, lat},
			},
		}
		b, _ := json.Marshal(feature)
		_, _ = c.fout.Write(b)
		_, _ = io.WriteString(c.fout, ",")
	}
	_, _ = io.WriteString(c.fout, "{ }]}")
	return <-errc
}

type kmlConverter struct{ fout io.Writer }

func NewKmlConverter(fout io.Writer) *kmlConverter { return &kmlConverter{fout} }
func (c *kmlConverter) Dirname() string            { return "" }

func (c *kmlConverter) Process(files []string, fileSize int64) error {
	_, _ = io.WriteString(c.fout, `<?xml version="1.0" encoding="UTF-8"?>
<kml xmlns="http://www.opengis.net/kml/2.2">
  <Document>`)
	hosts := map[string]map[string]any{}
	ch, errc := shodan.IterateFiles(files)
	for banner := range ch {
		ip := shodan.GetIP(banner)
		if ip == "" {
			continue
		}
		if _, ok := hosts[ip]; !ok {
			hosts[ip] = banner
			hosts[ip]["ports"] = []any{}
		}
		ports, _ := hosts[ip]["ports"].([]any)
		hosts[ip]["ports"] = append(ports, banner["port"])
	}
	for _, host := range hosts {
		writeKML(c.fout, host)
	}
	_, _ = io.WriteString(c.fout, `</Document></kml>`)
	return <-errc
}

func writeKML(w io.Writer, host map[string]any) {
	defer func() { recover() }()
	ip := shodan.GetIP(host)
	loc := asMap(host["location"])
	lat, lon := loc["latitude"], loc["longitude"]
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`<Placemark><name><![CDATA[<h1>%s</h1>]]></name><description><![CDATA[`, ip))
	b.WriteString("<h2>Ports</h2><ul>")
	for _, port := range asStringSlice(host["ports"]) {
		b.WriteString(fmt.Sprintf("<li>%s</li>", port))
	}
	b.WriteString(fmt.Sprintf(`</ul><div><a href="https://www.shodan.io/host/%s">View Details</a></div>]]></description>`, ip))
	b.WriteString(fmt.Sprintf(`<Point><coordinates>%v,%v</coordinates></Point></Placemark>`, lon, lat))
	_, _ = io.WriteString(w, b.String())
}

type imagesConverter struct {
	fout    *os.File
	dirname string
}

func NewImagesConverter(fout *os.File) *imagesConverter {
	return &imagesConverter{fout: fout}
}
func (c *imagesConverter) Dirname() string { return c.dirname }

func (c *imagesConverter) Process(files []string, fileSize int64) error {
	name := c.fout.Name()
	if len(name) > 7 {
		c.dirname = name[:len(name)-7] + "-images"
	} else {
		c.dirname = name + "-images"
	}
	c.fout.Close()
	_ = os.Remove(name)
	_ = os.MkdirAll(c.dirname, 0o755)
	ch, errc := shodan.IterateFiles(files)
	for banner := range ch {
		shot := shodan.GetScreenshot(banner)
		sm, ok := shot.(map[string]any)
		if !ok || sm == nil {
			continue
		}
		data, _ := sm["data"].(string)
		if data == "" {
			continue
		}
		filename := filepath.Join(c.dirname, fmt.Sprintf("%s-%v", shodan.GetIP(banner), banner["port"]))
		tmp := filename
		counter := 0
		for {
			if _, err := os.Stat(tmp + ".jpg"); os.IsNotExist(err) {
				break
			}
			tmp = fmt.Sprintf("%s-%d", filename, counter)
			counter++
		}
		raw, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			continue
		}
		_ = os.WriteFile(tmp+".jpg", raw, 0o644)
	}
	return <-errc
}
