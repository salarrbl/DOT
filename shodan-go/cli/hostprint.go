package cli

import (
	"fmt"
	"sort"
	"strings"

	shodan "github.com/salarrbl/DOT/shodan-go"
)

func asStringSlice(v any) []string {
	switch t := v.(type) {
	case []any:
		out := make([]string, 0, len(t))
		for _, x := range t {
			out = append(out, fmt.Sprint(x))
		}
		return out
	case []string:
		return t
	default:
		return nil
	}
}

func asMap(v any) map[string]any {
	m, _ := v.(map[string]any)
	return m
}

func asFloat(v any) float64 {
	f, _ := v.(float64)
	return f
}

func HostPrintPretty(host map[string]any, history bool) {
	fmt.Println(Color(shodan.GetIP(host), "green"))
	hostnames := asStringSlice(host["hostnames"])
	if len(hostnames) > 0 {
		fmt.Printf("%-25s%s\n", "Hostnames:", strings.Join(hostnames, ";"))
	}
	if city, _ := host["city"].(string); city != "" {
		fmt.Printf("%-25s%s\n", "City:", city)
	}
	if cn, _ := host["country_name"].(string); cn != "" {
		fmt.Printf("%-25s%s\n", "Country:", cn)
	}
	if osn, _ := host["os"].(string); osn != "" {
		fmt.Printf("%-25s%s\n", "Operating System:", osn)
	}
	if org, _ := host["org"].(string); org != "" {
		fmt.Printf("%-25s%s\n", "Organization:", org)
	}
	if lu, _ := host["last_update"].(string); lu != "" {
		fmt.Printf("%-25s%s\n", "Updated:", lu)
	}
	ports := asStringSlice(host["ports"])
	if ports == nil {
		if ps, ok := host["ports"].([]any); ok {
			ports = make([]string, len(ps))
			for i, p := range ps {
				ports[i] = fmt.Sprint(p)
			}
		}
	}
	fmt.Printf("%-25s%d\n", "Number of open ports:", len(ports))

	if vulnsRaw := host["vulns"]; vulnsRaw != nil {
		var vulns []string
		switch t := vulnsRaw.(type) {
		case []any:
			for _, vuln := range t {
				s := fmt.Sprint(vuln)
				if strings.HasPrefix(s, "!") {
					continue
				}
				if strings.ToUpper(s) == "CVE-2014-0160" {
					vulns = append(vulns, Color("Heartbleed", "red"))
				} else {
					vulns = append(vulns, Color(s, "red"))
				}
			}
		case map[string]any:
			for s := range t {
				if strings.HasPrefix(s, "!") {
					continue
				}
				if strings.ToUpper(s) == "CVE-2014-0160" {
					vulns = append(vulns, Color("Heartbleed", "red"))
				} else {
					vulns = append(vulns, Color(s, "red"))
				}
			}
		}
		if len(vulns) > 0 {
			fmt.Printf("%-25s", "Vulnerabilities:")
			for _, v := range vulns {
				fmt.Print(v + "\t")
			}
			fmt.Println()
		}
	}
	fmt.Println()

	data, _ := host["data"].([]any)
	if len(ports) != len(data) && len(data) > 0 {
		seen := map[int]bool{}
		for _, b := range data {
			bm := asMap(b)
			seen[int(asFloat(bm["port"]))] = true
		}
		last := asMap(data[len(data)-1])
		for _, p := range host["ports"].([]any) {
			port := int(asFloat(p))
			if !seen[port] {
				data = append(data, map[string]any{
					"port":        float64(port),
					"transport":   "tcp",
					"timestamp":   last["timestamp"],
					"placeholder": true,
				})
			}
		}
		host["data"] = data
	}

	fmt.Println("Ports:")
	sort.Slice(data, func(i, j int) bool {
		return asFloat(asMap(data[i])["port"]) < asFloat(asMap(data[j])["port"])
	})
	for _, raw := range data {
		banner := asMap(raw)
		port := int(asFloat(banner["port"]))
		product, _ := banner["product"].(string)
		version := ""
		if v, _ := banner["version"].(string); v != "" {
			version = "(" + v + ")"
		}
		fmt.Print(Color(fmt.Sprintf("%7d", port), "cyan"))
		if tr, _ := banner["transport"].(string); tr != "" {
			fmt.Print("/")
			fmt.Print(Color(tr+" ", "yellow"))
		}
		fmt.Printf("%s %s", product, version)
		if history {
			if ts, _ := banner["timestamp"].(string); len(ts) >= 10 {
				fmt.Print(Color("\t\t("+ts[:10]+")", "white"))
			}
		}
		fmt.Println()
		if http := asMap(banner["http"]); http != nil {
			if title, _ := http["title"].(string); title != "" {
				fmt.Printf("\t|-- HTTP title: %s\n", title)
			}
		}
		if ssl := asMap(banner["ssl"]); ssl != nil {
			if cert := asMap(ssl["cert"]); cert != nil {
				if issuer := asMap(cert["issuer"]); issuer != nil {
					var parts []string
					for k, v := range issuer {
						parts = append(parts, fmt.Sprintf("%s=%v", k, v))
					}
					fmt.Printf("\t|-- Cert Issuer: %s\n", strings.Join(parts, ", "))
				}
				if subject := asMap(cert["subject"]); subject != nil {
					var parts []string
					for k, v := range subject {
						parts = append(parts, fmt.Sprintf("%s=%v", k, v))
					}
					fmt.Printf("\t|-- Cert Subject: %s\n", strings.Join(parts, ", "))
				}
			}
			if versions, ok := ssl["versions"].([]any); ok {
				var vs []string
				for _, item := range versions {
					s := fmt.Sprint(item)
					if !strings.HasPrefix(s, "-") {
						vs = append(vs, s)
					}
				}
				sort.Strings(vs)
				fmt.Printf("\t|-- SSL Versions: %s\n", strings.Join(vs, ", "))
			}
			if dh := asMap(ssl["dhparams"]); dh != nil {
				fmt.Println("\t|-- Diffie-Hellman Parameters:")
				fmt.Printf("\t\t%-15s%v\n\t\t%-15s%v\n", "Bits:", dh["bits"], "Generator:", dh["generator"])
				if fp, ok := dh["fingerprint"]; ok {
					fmt.Printf("\t\t%-15s%v\n", "Fingerprint:", fp)
				}
			}
		}
	}
}

func HostPrintTSV(host map[string]any, history bool) {
	data, _ := host["data"].([]any)
	sort.Slice(data, func(i, j int) bool {
		return asFloat(asMap(data[i])["port"]) < asFloat(asMap(data[j])["port"])
	})
	for _, raw := range data {
		banner := asMap(raw)
		fmt.Print(Color(fmt.Sprintf("%7d", int(asFloat(banner["port"]))), "cyan"))
		fmt.Print("\t")
		fmt.Print(Color(fmt.Sprintf("%v ", banner["transport"]), "yellow"))
		if history {
			if ts, _ := banner["timestamp"].(string); len(ts) >= 10 {
				fmt.Print(Color("\t("+ts[:10]+")", "white"))
			}
		}
		fmt.Println()
	}
}

var HostPrint = map[string]func(map[string]any, bool){
	"pretty": HostPrintPretty,
	"tsv":    HostPrintTSV,
}
