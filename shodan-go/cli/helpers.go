package cli

import (
	"compress/gzip"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	shodan "github.com/salarrbl/DOT/shodan-go"
)

func GetAPIKey() (string, error) {
	keyfile := filepath.Join(ConfigDir(), "api_key")
	if _, err := os.Stat(keyfile); err != nil {
		return "", fmt.Errorf(`Please run "shodan init <api key>" before using this command`)
	}
	_ = os.Chmod(keyfile, 0o600)
	b, err := os.ReadFile(keyfile)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

func MustAPI() *shodan.Client {
	key, err := GetAPIKey()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
	return shodan.New(key)
}

func EscapeData(args any) string {
	s := fmt.Sprint(args)
	s = strings.ReplaceAll(s, "\n", "\\n")
	s = strings.ReplaceAll(s, "\r", "\\r")
	s = strings.ReplaceAll(s, "\t", "\\t")
	return s
}

func Timestr() string {
	return time.Now().UTC().Format("2006-01-02")
}

func OpenStreamingFile(directory, timestr string, compresslevel int) (*gzip.Writer, *os.File, error) {
	if compresslevel < gzip.NoCompression || compresslevel > gzip.BestCompression {
		compresslevel = gzip.BestCompression
	}
	f, err := os.OpenFile(filepath.Join(directory, timestr+".json.gz"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return nil, nil, err
	}
	w, err := gzip.NewWriterLevel(f, compresslevel)
	if err != nil {
		f.Close()
		return nil, nil, err
	}
	return w, f, nil
}

func GetBannerField(banner map[string]any, flatField string) any {
	fields := strings.Split(flatField, ".")
	var current any = banner
	for _, field := range fields {
		m, ok := current.(map[string]any)
		if !ok {
			return nil
		}
		current, ok = m[field]
		if !ok {
			return nil
		}
	}
	return current
}

func filterWithNetmask(banner map[string]any, netmask string) bool {
	_, network, err := net.ParseCIDR(netmask)
	if err != nil {
		return false
	}
	ipField := GetBannerField(banner, "ip")
	if ipField == nil {
		return false
	}
	var ip net.IP
	switch v := ipField.(type) {
	case string:
		ip = net.ParseIP(v)
	case float64:
		// Shodan often stores IPv4 as integer
		n := uint32(v)
		ip = net.IPv4(byte(n>>24), byte(n>>16), byte(n>>8), byte(n))
	default:
		ip = net.ParseIP(fmt.Sprint(v))
	}
	if ip == nil {
		return false
	}
	return network.Contains(ip)
}

func MatchFilters(banner map[string]any, filters []string) bool {
	for _, args := range filters {
		parts := strings.SplitN(args, ":", 2)
		if len(parts) != 2 {
			continue
		}
		flatField, check := parts[0], parts[1]
		if flatField == "net" {
			return filterWithNetmask(banner, check)
		}
		value := GetBannerField(banner, flatField)
		if value == nil {
			return false
		}
		switch v := value.(type) {
		case []any:
			ok := false
			for _, item := range v {
				if strings.Contains(fmt.Sprint(item), check) || fmt.Sprint(item) == check {
					ok = true
					break
				}
			}
			if !ok {
				return false
			}
		case string:
			if !strings.Contains(v, check) {
				return false
			}
		case float64:
			n, err := strconv.ParseFloat(check, 64)
			if err != nil || n != v {
				i, err2 := strconv.Atoi(check)
				if err2 != nil || float64(i) != v {
					return false
				}
			}
		}
	}
	return true
}

func HumanizeAPIPlan(plan string) string {
	m := map[string]string{
		"oss":        "Free",
		"dev":        "Membership",
		"basic":      "Freelancer API",
		"plus":       "Small Business API",
		"corp":       "Corporate API",
		"stream-100": "Enterprise",
	}
	if s, ok := m[plan]; ok {
		return s
	}
	return plan
}

func Color(s, fg string) string {
	codes := map[string]string{
		"green":   "32",
		"yellow":  "33",
		"cyan":    "36",
		"magenta": "35",
		"red":     "31",
		"white":   "37",
		"blue":    "34",
	}
	c, ok := codes[fg]
	if !ok {
		return s
	}
	return "\x1b[" + c + "m" + s + "\x1b[0m"
}

func Dim(s string) string {
	return "\x1b[2m" + s + "\x1b[0m"
}

func Die(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func FormatBannerRow(banner map[string]any, fields []string, sep string, color bool) string {
	var row strings.Builder
	for i, field := range fields {
		tmp := ""
		value := GetBannerField(banner, field)
		if value != nil {
			switch v := value.(type) {
			case []any:
				parts := make([]string, 0, len(v))
				for _, x := range v {
					parts = append(parts, fmt.Sprint(x))
				}
				tmp = strings.Join(parts, ";")
			case float64:
				if v == float64(int64(v)) {
					tmp = strconv.FormatInt(int64(v), 10)
				} else {
					tmp = fmt.Sprint(v)
				}
			default:
				tmp = EscapeData(v)
			}
			if color {
				tmp = Color(tmp, ColorizeFields[field])
				if ColorizeFields[field] == "" {
					tmp = Color(tmp, "white")
				}
			}
		}
		if i > 0 {
			row.WriteString(sep)
		}
		row.WriteString(tmp)
	}
	return row.String()
}

func EnsureJSONGz(filename string) string {
	if !strings.HasSuffix(filename, ".json.gz") {
		return filename + ".json.gz"
	}
	return filename
}
