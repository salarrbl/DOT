// Shodan CLI — Go copy of shodan/__main__.py and shodan/cli/*
//
// Always run "shodan init <api key>" before other commands.
package main

import (
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	shodan "github.com/salarrbl/DOT/shodan-go"
	"github.com/salarrbl/DOT/shodan-go/cli"
)

const version = "1.31.0"

func usage() {
	fmt.Fprintf(os.Stderr, `Shodan CLI (Go port of shodan-python)

Commands:
  init <api key>
  info
  count <query>
  search [--fields f] [--limit n] <query>
  download [--limit n] [--fields f] <filename> <query>
  host [--history] [--format pretty|tsv] [--save] <ip>
  parse [--fields f] [--filters k:v] <files...>
  convert [--fields f] <input> <kml|csv|geo.json|images>
  domain [--details] [--save] [--history] [--type T] <domain>
  myip [--ipv6]
  stats [--facets f] [--limit n] [-O file] <query>
  stream [filters...]
  trends [--facets f] <query>
  honeyscore <ip>
  version
  alert  (clear|create|list|info|remove|enable|disable|triggers|stats|download|export|import|domain)
  data   (list|download)
  org    (info|add|remove)
  scan   (list|submit|status|protocols|internet)

`)
	os.Exit(2)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	switch os.Args[1] {
	case "init":
		cmdInit(os.Args[2:])
	case "info":
		cmdInfo()
	case "count":
		cmdCount(os.Args[2:])
	case "search":
		cmdSearch(os.Args[2:])
	case "download":
		cmdDownload(os.Args[2:])
	case "host":
		cmdHost(os.Args[2:])
	case "parse":
		cmdParse(os.Args[2:])
	case "convert":
		cmdConvert(os.Args[2:])
	case "domain":
		cmdDomain(os.Args[2:])
	case "myip":
		cmdMyIP(os.Args[2:])
	case "stats":
		cmdStats(os.Args[2:])
	case "stream":
		cmdStream(os.Args[2:])
	case "trends":
		cmdTrends(os.Args[2:])
	case "honeyscore":
		cmdHoneyscore(os.Args[2:])
	case "version":
		fmt.Println(version)
	case "alert":
		cmdAlert(os.Args[2:])
	case "data":
		cmdData(os.Args[2:])
	case "org":
		cmdOrg(os.Args[2:])
	case "scan":
		cmdScan(os.Args[2:])
	case "-h", "--help", "help":
		usage()
	default:
		usage()
	}
}

func cmdInit(args []string) {
	if len(args) < 1 {
		usage()
	}
	key := strings.TrimSpace(args[0])
	dir := cli.ConfigDir()
	if err := os.MkdirAll(dir, 0o700); err != nil {
		cli.Die(fmt.Errorf("Unable to create directory to store the Shodan API key (%s)", dir))
	}
	api := shodan.New(key)
	if _, err := api.Info(); err != nil {
		cli.Die(err)
	}
	keyfile := filepath.Join(dir, "api_key")
	cli.Die(os.WriteFile(keyfile, []byte(key), 0o600))
	fmt.Println(cli.Color("Successfully initialized", "green"))
}

func cmdInfo() {
	api := cli.MustAPI()
	results, err := api.Info()
	cli.Die(err)
	fmt.Printf("Query credits available: %v\nScan credits available: %v\n", results["query_credits"], results["scan_credits"])
}

func cmdCount(args []string) {
	query := strings.Join(args, " ")
	if strings.TrimSpace(query) == "" {
		cli.Die(fmt.Errorf("Empty search query"))
	}
	api := cli.MustAPI()
	results, err := api.Count(query, nil)
	cli.Die(err)
	fmt.Println(results["total"])
}

func parseFlags(args []string) (map[string]string, []string) {
	flags := map[string]string{}
	var rest []string
	for i := 0; i < len(args); i++ {
		a := args[i]
		if a == "--" {
			rest = append(rest, args[i+1:]...)
			break
		}
		if strings.HasPrefix(a, "--") {
			name := strings.TrimPrefix(a, "--")
			if strings.Contains(name, "=") {
				kv := strings.SplitN(name, "=", 2)
				flags[kv[0]] = kv[1]
				continue
			}
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				flags[name] = args[i+1]
				i++
			} else {
				flags[name] = "true"
			}
			continue
		}
		if strings.HasPrefix(a, "-") && len(a) == 2 {
			if i+1 < len(args) && !strings.HasPrefix(args[i+1], "-") {
				flags[a] = args[i+1]
				i++
			} else {
				flags[a] = "true"
			}
			continue
		}
		rest = append(rest, a)
	}
	return flags, rest
}

func cmdSearch(args []string) {
	flags, rest := parseFlags(args)
	query := strings.Join(rest, " ")
	if strings.TrimSpace(query) == "" {
		cli.Die(fmt.Errorf("Empty search query"))
	}
	fields := "ip_str,port,hostnames,data"
	if flags["fields"] != "" {
		fields = flags["fields"]
	}
	limit := 100
	if flags["limit"] != "" {
		limit, _ = strconv.Atoi(flags["limit"])
	}
	if limit > 1000 {
		cli.Die(fmt.Errorf("Too many results requested, maximum is 1,000"))
	}
	sep := "\t"
	if flags["separator"] != "" {
		sep = flags["separator"]
	}
	color := flags["no-color"] != "true"
	flist := splitCSV(fields)
	api := cli.MustAPI()
	results, err := api.Search(query, 1, limit, 0, nil, false, flist)
	cli.Die(err)
	if asFloat(results["total"]) == 0 {
		cli.Die(fmt.Errorf("No search results found"))
	}
	matches, _ := results["matches"].([]any)
	for _, m := range matches {
		banner, _ := m.(map[string]any)
		fmt.Println(cli.FormatBannerRow(banner, flist, sep, color))
	}
}

func splitCSV(s string) []string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func asFloat(v any) float64 {
	f, _ := v.(float64)
	return f
}

func cmdDownload(args []string) {
	flags, rest := parseFlags(args)
	if len(rest) < 2 {
		usage()
	}
	filename := rest[0]
	query := strings.Join(rest[1:], " ")
	if query == "" || filename == "" {
		cli.Die(fmt.Errorf("Empty filename or query"))
	}
	filename = cli.EnsureJSONGz(filename)
	limit := 1000
	if flags["limit"] != "" {
		limit, _ = strconv.Atoi(flags["limit"])
	}
	var fields []string
	if flags["fields"] != "" {
		fields = splitCSV(flags["fields"])
	}
	api := cli.MustAPI()
	totalWrap, err := api.Count(query, nil)
	cli.Die(err)
	info, err := api.Info()
	cli.Die(err)
	total := int(asFloat(totalWrap["total"]))
	fmt.Printf("Search query:\t\t\t%s\n", query)
	fmt.Printf("Total number of results:\t%d\n", total)
	fmt.Printf("Query credits left:\t\t%v\n", info["unlocked_left"])
	fmt.Printf("Output file:\t\t\t%s\n", filename)
	if limit > total || limit <= 0 {
		limit = total
	}
	gw, f, err := shodan.OpenFile(filename)
	cli.Die(err)
	defer f.Close()
	defer gw.Close()
	ch, errc := api.SearchCursor(query, false, 5, fields)
	count := 0
	for banner := range ch {
		cli.Die(shodan.WriteBanner(gw, banner))
		count++
		if count >= limit {
			break
		}
	}
	select {
	case err := <-errc:
		if err != nil && count < limit {
			fmt.Println(cli.Color("Notice: fewer results were saved than requested", "yellow"))
		}
	default:
	}
	fmt.Println(cli.Color(fmt.Sprintf("Saved %d results into file %s", count, filename), "green"))
}

func cmdHost(args []string) {
	flags, rest := parseFlags(args)
	if len(rest) < 1 {
		usage()
	}
	ip := rest[0]
	format := flags["format"]
	if format == "" {
		format = "pretty"
	}
	history := flags["history"] == "true"
	save := flags["save"] == "true" || flags["-S"] == "true"
	filename := flags["filename"]
	if filename == "" {
		filename = flags["-O"]
	}
	api := cli.MustAPI()
	raw, err := api.Host([]string{ip}, history, false)
	cli.Die(err)
	host, ok := raw.(map[string]any)
	if !ok {
		cli.Die(fmt.Errorf("unexpected host response"))
	}
	fn := cli.HostPrint[format]
	if fn == nil {
		cli.Die(fmt.Errorf("unknown format"))
	}
	fn(host, history)
	if filename != "" || save {
		if save {
			filename = ip + ".json.gz"
		}
		filename = cli.EnsureJSONGz(filename)
		gw, f, err := shodan.OpenFile(filename)
		cli.Die(err)
		defer f.Close()
		defer gw.Close()
		data, _ := host["data"].([]any)
		for _, b := range data {
			banner, _ := b.(map[string]any)
			if _, ok := banner["placeholder"]; ok {
				continue
			}
			cli.Die(shodan.WriteBanner(gw, banner))
		}
	}
}

func cmdParse(args []string) {
	flags, rest := parseFlags(args)
	fields := splitCSV(orDefault(flags["fields"], "ip_str,port,hostnames,data"))
	if len(fields) == 0 {
		cli.Die(fmt.Errorf("Please define at least one property to show"))
	}
	var filters []string
	if flags["filters"] != "" {
		filters = append(filters, flags["filters"])
	}
	if flags["-f"] != "" {
		filters = append(filters, flags["-f"])
	}
	sep := orDefault(flags["separator"], "\t")
	color := flags["no-color"] != "true"
	filename := flags["filename"]
	if filename == "" {
		filename = flags["-O"]
	}
	var fout *gzip.Writer
	var ff *os.File
	if filename != "" {
		if len(filters) == 0 {
			cli.Die(fmt.Errorf("Output file specified without any filters. Need to use filters with this option."))
		}
		filename = cli.EnsureJSONGz(filename)
		var err error
		fout, ff, err = shodan.OpenFile(filename)
		cli.Die(err)
		defer ff.Close()
		defer fout.Close()
	}
	ch, errc := shodan.IterateFiles(rest)
	for banner := range ch {
		if len(filters) > 0 && !cli.MatchFilters(banner, filters) {
			continue
		}
		if fout != nil {
			_ = shodan.WriteBanner(fout, banner)
		}
		fmt.Println(cli.FormatBannerRow(banner, fields, sep, color))
	}
	if err := <-errc; err != nil {
		cli.Die(err)
	}
}

func orDefault(s, d string) string {
	if s == "" {
		return d
	}
	return s
}

func cmdConvert(args []string) {
	flags, rest := parseFlags(args)
	if len(rest) < 2 {
		usage()
	}
	input, format := rest[0], rest[1]
	st, err := os.Stat(input)
	cli.Die(err)
	basename := strings.TrimSuffix(strings.TrimSuffix(input, ".json.gz"), ".json")
	outname := basename + "." + format
	fout, err := os.Create(outname)
	cli.Die(err)
	var conv cli.Converter
	fields := splitCSV(flags["fields"])
	switch format {
	case "csv":
		conv = cli.NewCsvConverter(fout, fields)
	case "kml":
		conv = cli.NewKmlConverter(fout)
	case "geo.json":
		conv = cli.NewGeoJSONConverter(fout)
	case "images":
		conv = cli.NewImagesConverter(fout)
	default:
		cli.Die(fmt.Errorf("unknown format %s", format))
	}
	cli.Die(conv.Process([]string{input}, st.Size()))
	if format == "images" {
		fmt.Println(cli.Color("Successfully extracted images to directory: "+conv.Dirname(), "green"))
	} else {
		fout.Close()
		fmt.Println(cli.Color("Successfully created new file: "+outname, "green"))
	}
}

func cmdDomain(args []string) {
	flags, rest := parseFlags(args)
	if len(rest) < 1 {
		usage()
	}
	domain := rest[0]
	api := cli.MustAPI()
	info, err := api.DNS.DomainInfo(domain, flags["history"] == "true" || flags["-H"] == "true", flags["type"], 1)
	if flags["-T"] != "" {
		info, err = api.DNS.DomainInfo(domain, flags["history"] == "true", flags["-T"], 1)
	}
	cli.Die(err)
	fmt.Println(cli.Color(strings.ToUpper(fmt.Sprint(info["domain"])), "green"))
	fmt.Println()
	data, _ := info["data"].([]any)
	hosts := map[string]map[string]any{}
	if flags["details"] == "true" || flags["-D"] == "true" {
		ips := map[string]bool{}
		for _, rec := range data {
			m, _ := rec.(map[string]any)
			if t, _ := m["type"].(string); t == "A" || t == "AAAA" {
				ips[fmt.Sprint(m["value"])] = true
			}
		}
		var fout *gzip.Writer
		var ff *os.File
		if flags["save"] == "true" || flags["-S"] == "true" {
			var err error
			fout, ff, err = shodan.OpenFile(domain + "-hosts.json.gz")
			cli.Die(err)
			defer ff.Close()
			defer fout.Close()
		}
		for ip := range ips {
			h, err := api.Host([]string{ip}, false, false)
			if err != nil {
				continue
			}
			hm, _ := h.(map[string]any)
			hosts[ip] = hm
			if fout != nil {
				for _, b := range hm["data"].([]any) {
					banner, _ := b.(map[string]any)
					if _, ok := banner["placeholder"]; !ok {
						_ = shodan.WriteBanner(fout, banner)
					}
				}
			}
		}
	}
	if flags["save"] == "true" || flags["-S"] == "true" {
		fout, ff, err := shodan.OpenFile(domain + ".json.gz")
		cli.Die(err)
		defer ff.Close()
		defer fout.Close()
		for _, rec := range data {
			_ = shodan.WriteBanner(fout, rec)
		}
	}
	for _, rec := range data {
		m, _ := rec.(map[string]any)
		line := fmt.Sprintf("%-32s  %-14s  %v", cli.Color(fmt.Sprint(m["subdomain"]), "cyan"), cli.Color(fmt.Sprint(m["type"]), "yellow"), m["value"])
		fmt.Print(line)
		if h, ok := hosts[fmt.Sprint(m["value"])]; ok {
			ports := []string{}
			if ps, ok := h["ports"].([]any); ok {
				for _, p := range ps {
					ports = append(ports, fmt.Sprint(int(asFloat(p))))
				}
			}
			fmt.Print(cli.Color(" Ports: "+strings.Join(ports, ", "), "blue"))
		}
		fmt.Println()
	}
}

func cmdMyIP(args []string) {
	flags, _ := parseFlags(args)
	api := cli.MustAPI()
	if flags["ipv6"] == "true" || flags["-6"] == "true" {
		api.BaseURL = "https://apiv6.shodan.io"
	}
	v, err := api.Tools.MyIP()
	cli.Die(err)
	fmt.Println(v)
}

func cmdStats(args []string) {
	flags, rest := parseFlags(args)
	query := strings.Join(rest, " ")
	if strings.TrimSpace(query) == "" {
		cli.Die(fmt.Errorf("Empty search query"))
	}
	limit := 10
	if flags["limit"] != "" {
		limit, _ = strconv.Atoi(flags["limit"])
	}
	facetNames := splitCSV(orDefault(flags["facets"], "country,org"))
	var facets []any
	for _, f := range facetNames {
		facets = append(facets, shodan.Facet{Name: f, Count: limit})
	}
	api := cli.MustAPI()
	results, err := api.Count(query, facets)
	cli.Die(err)
	printFacets(results)
	filename := flags["filename"]
	if filename == "" {
		filename = flags["-O"]
	}
	if filename != "" {
		if !strings.HasSuffix(filename, ".csv") {
			filename += ".csv"
		}
		writeFacetCSV(filename, results, query, true)
	}
}

func printFacets(results map[string]any) {
	fm, _ := results["facets"].(map[string]any)
	for facet, raw := range fm {
		items, _ := raw.([]any)
		fmt.Printf("Top %d Results for Facet: %s\n", len(items), facet)
		for _, item := range items {
			m, _ := item.(map[string]any)
			fmt.Print(cli.Color(fmt.Sprintf("%-28s", fmt.Sprint(m["value"])), "cyan"))
			fmt.Println(cli.Color(fmt.Sprintf("%12.0f", asFloat(m["count"])), "green"))
		}
		fmt.Println()
	}
}

func writeFacetCSV(filename string, results map[string]any, query string, withQuery bool) {
	f, err := os.Create(filename)
	cli.Die(err)
	defer f.Close()
	w := csv.NewWriter(f)
	if withQuery {
		_ = w.Write([]string{"Query", query})
		_ = w.Write([]string{})
	}
	fm, _ := results["facets"].(map[string]any)
	header := []string{}
	keys := []string{}
	for facet := range fm {
		keys = append(keys, facet)
		header = append(header, facet, "")
	}
	_ = w.Write(header)
	counter := 0
	for {
		row := make([]string, len(keys)*2)
		has := false
		pos := 0
		for _, facet := range keys {
			values, _ := fm[facet].([]any)
			if len(values) > counter {
				has = true
				m, _ := values[counter].(map[string]any)
				row[pos] = fmt.Sprint(m["value"])
				row[pos+1] = fmt.Sprint(m["count"])
			}
			pos += 2
		}
		if !has {
			break
		}
		_ = w.Write(row)
		counter++
	}
	w.Flush()
}

func cmdStream(args []string) {
	flags, _ := parseFlags(args)
	api := cli.MustAPI()
	if flags["streamer"] != "" {
		api.Stream.BaseURL = flags["streamer"]
	}
	fields := splitCSV(orDefault(flags["fields"], "ip_str,port,hostnames,data"))
	sep := orDefault(flags["separator"], "\t")
	color := flags["no-color"] != "true"
	quiet := flags["quiet"] == "true"
	limit := -1
	if flags["limit"] != "" {
		limit, _ = strconv.Atoi(flags["limit"])
	}
	timeout := time.Duration(0)
	if flags["timeout"] != "" {
		sec, _ := strconv.Atoi(flags["timeout"])
		timeout = time.Duration(sec) * time.Second
	}
	var ch <-chan shodan.BannerEvent
	var errc <-chan error
	var closer io.Closer
	var err error
	nfilters := 0
	pick := func() { nfilters++ }
	if flags["ports"] != "" {
		pick()
		var ports []int
		for _, p := range splitCSV(flags["ports"]) {
			n, _ := strconv.Atoi(p)
			ports = append(ports, n)
		}
		ch, errc, closer, err = api.Stream.Ports(ports, timeout, false)
	} else if flags["countries"] != "" {
		pick()
		ch, errc, closer, err = api.Stream.Countries(splitCSV(flags["countries"]), timeout, false)
	} else if flags["asn"] != "" {
		pick()
		ch, errc, closer, err = api.Stream.ASN(splitCSV(flags["asn"]), timeout, false)
	} else if flags["alert"] != "" {
		pick()
		aid := flags["alert"]
		if strings.ToLower(aid) == "all" {
			aid = ""
		}
		ch, errc, closer, err = api.Stream.Alert(aid, timeout, false)
	} else if flags["tags"] != "" {
		pick()
		ch, errc, closer, err = api.Stream.Tags(splitCSV(flags["tags"]), timeout, false)
	} else if flags["vulns"] != "" {
		pick()
		ch, errc, closer, err = api.Stream.Vulns(splitCSV(flags["vulns"]), timeout, false)
	} else if flags["custom-filters"] != "" {
		pick()
		ch, errc, closer, err = api.Stream.Custom(flags["custom-filters"], timeout, false)
	} else {
		ch, errc, closer, err = api.Stream.Banners(timeout, false)
	}
	cli.Die(err)
	defer closer.Close()
	counter := 0
	for ev := range ch {
		if limit > 0 {
			counter++
			if counter > limit {
				break
			}
		}
		if !quiet {
			fmt.Println(cli.FormatBannerRow(ev.Object, fields, sep, color))
		}
	}
	if err := <-errc; err != nil {
		cli.Die(err)
	}
}

func cmdTrends(args []string) {
	flags, rest := parseFlags(args)
	query := strings.Join(rest, " ")
	if strings.TrimSpace(query) == "" {
		cli.Die(fmt.Errorf("Empty search query"))
	}
	var parsed []any
	for _, facet := range splitCSV(flags["facets"]) {
		parts := strings.SplitN(facet, ":", 2)
		if len(parts) > 1 {
			n, _ := strconv.Atoi(parts[1])
			parsed = append(parsed, shodan.Facet{Name: parts[0], Count: n})
		} else if parts[0] != "" {
			parsed = append(parsed, parts[0])
		}
	}
	api := cli.MustAPI()
	results, err := api.Trends.Search(query, parsed)
	cli.Die(err)
	if asFloat(results["total"]) == 0 {
		cli.Die(fmt.Errorf("No search results found"))
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(results)
}

func cmdHoneyscore(args []string) {
	if len(args) < 1 {
		usage()
	}
	api := cli.MustAPI()
	score, err := api.Labs.HoneyScore(args[0])
	cli.Die(err)
	f := asFloat(score)
	switch {
	case f == 1.0:
		fmt.Println(cli.Color("Honeypot detected", "red"))
	case f > 0.5:
		fmt.Println(cli.Color("Probably a honeypot", "yellow"))
	default:
		fmt.Println(cli.Color("Not a honeypot", "green"))
	}
	fmt.Printf("Score: %v\n", score)
}

func cmdAlert(args []string) {
	if len(args) < 1 {
		usage()
	}
	api := cli.MustAPI()
	switch args[0] {
	case "clear":
		alerts, err := api.Alerts("", true)
		cli.Die(err)
		list, _ := alerts.([]any)
		for _, a := range list {
			m, _ := a.(map[string]any)
			fmt.Printf("Removing %v (%v)\n", m["name"], m["id"])
			_, err := api.DeleteAlert(fmt.Sprint(m["id"]))
			cli.Die(err)
		}
		fmt.Println("Alerts deleted")
	case "create":
		if len(args) < 3 {
			usage()
		}
		alert, err := api.CreateAlert(args[1], args[2:], 0)
		cli.Die(err)
		fmt.Println(cli.Color("Successfully created network alert!", "green"))
		fmt.Println(cli.Color("Alert ID: "+fmt.Sprint(alert["id"]), "cyan"))
	case "list":
		results, err := api.Alerts("", true)
		cli.Die(err)
		list, _ := results.([]any)
		if len(list) == 0 {
			fmt.Println("You haven't created any alerts yet.")
			return
		}
		fmt.Printf("# %-14s %-21s %-15s\n", "Alert ID", "Name", "IP/ Network")
		for _, a := range list {
			m, _ := a.(map[string]any)
			filters := asMap(m["filters"])
			ips := []string{}
			if ip, ok := filters["ip"].([]any); ok {
				for _, x := range ip {
					ips = append(ips, fmt.Sprint(x))
				}
			}
			fmt.Printf("%-16s %-30s %-35s\n", cli.Color(fmt.Sprint(m["id"]), "yellow"), cli.Color(fmt.Sprint(m["name"]), "cyan"), strings.Join(ips, ", "))
		}
	case "info":
		if len(args) < 2 {
			usage()
		}
		info, err := api.Alerts(args[1], true)
		cli.Die(err)
		m, _ := info.(map[string]any)
		fmt.Println(cli.Color(fmt.Sprint(m["name"]), "cyan"))
		fmt.Println("Created:", m["created"])
		filters := asMap(m["filters"])
		fmt.Println("Network Range(s):")
		if ip, ok := filters["ip"].([]any); ok {
			for _, n := range ip {
				fmt.Println(" >", cli.Color(fmt.Sprint(n), "yellow"))
			}
		}
	case "remove":
		if len(args) < 2 {
			usage()
		}
		_, err := api.DeleteAlert(args[1])
		cli.Die(err)
		fmt.Println("Alert deleted")
	case "triggers":
		results, err := api.AlertTriggers()
		cli.Die(err)
		list, _ := results.([]any)
		for _, t := range list {
			m, _ := t.(map[string]any)
			fmt.Println("Name        ", cli.Color(fmt.Sprint(m["name"]), "yellow"))
			fmt.Println("Description ", cli.Color(fmt.Sprint(m["description"]), "cyan"))
			fmt.Println("Rule        ", m["rule"])
			fmt.Println()
		}
	case "enable":
		if len(args) < 3 {
			usage()
		}
		_, err := api.EnableAlertTrigger(args[1], args[2])
		cli.Die(err)
		fmt.Println(cli.Color("Successfully enabled the trigger: "+args[2], "green"))
	case "disable":
		if len(args) < 3 {
			usage()
		}
		_, err := api.DisableAlertTrigger(args[1], args[2])
		cli.Die(err)
		fmt.Println(cli.Color("Successfully disabled the trigger: "+args[2], "green"))
	case "stats":
		flags, rest := parseFlags(args[1:])
		if len(rest) == 0 {
			cli.Die(fmt.Errorf("No facets provided"))
		}
		limit := 10
		if flags["limit"] != "" {
			limit, _ = strconv.Atoi(flags["limit"])
		}
		alerts, err := api.Alerts("", true)
		cli.Die(err)
		nets := []string{}
		if list, ok := alerts.([]any); ok {
			for _, a := range list {
				m := asMap(a)
				filters := asMap(m["filters"])
				if ip, ok := filters["ip"].([]any); ok {
					for _, n := range ip {
						nets = append(nets, fmt.Sprint(n))
					}
				}
			}
		}
		var facets []any
		for _, f := range rest {
			facets = append(facets, shodan.Facet{Name: f, Count: limit})
		}
		query := "net:" + strings.Join(nets, ",")
		results, err := api.Count(query, facets)
		cli.Die(err)
		printFacets(results)
	default:
		usage()
	}
}

func asMap(v any) map[string]any {
	m, _ := v.(map[string]any)
	return m
}

func cmdData(args []string) {
	if len(args) < 1 {
		usage()
	}
	api := cli.MustAPI()
	switch args[0] {
	case "list":
		flags, _ := parseFlags(args[1:])
		if flags["dataset"] != "" {
			files, err := api.Data.ListFiles(flags["dataset"])
			cli.Die(err)
			list, _ := files.([]any)
			for _, f := range list {
				m := asMap(f)
				fmt.Printf("%-20s%-10s", cli.Color(fmt.Sprint(m["name"]), "cyan"), cli.Color(shodan.HumanizeBytes(asFloat(m["size"]), 1), "yellow"))
				if sha, ok := m["sha1"].(string); ok && sha != "" {
					fmt.Printf("%-42s", cli.Color(sha, "green"))
				}
				fmt.Println(m["url"])
			}
		} else {
			ds, err := api.Data.ListDatasets()
			cli.Die(err)
			list, _ := ds.([]any)
			for _, d := range list {
				m := asMap(d)
				fmt.Printf("%-15s%s\n", cli.Color(fmt.Sprint(m["name"]), "cyan"), m["description"])
			}
		}
	case "download":
		flags, rest := parseFlags(args[1:])
		if len(rest) < 2 {
			usage()
		}
		dataset, name := rest[0], rest[1]
		files, err := api.Data.ListFiles(dataset)
		cli.Die(err)
		var file map[string]any
		for _, tmp := range files.([]any) {
			m := asMap(tmp)
			if fmt.Sprint(m["name"]) == name {
				file = m
				break
			}
		}
		if file == nil {
			cli.Die(fmt.Errorf("File not found"))
		}
		filename := flags["filename"]
		if filename == "" {
			filename = flags["-O"]
		}
		if filename == "" {
			filename = dataset + "-" + name
		}
		resp, err := http.Get(fmt.Sprint(file["url"]))
		cli.Die(err)
		defer resp.Body.Close()
		f, err := os.Create(filename)
		cli.Die(err)
		defer f.Close()
		_, err = io.Copy(f, resp.Body)
		cli.Die(err)
		fmt.Println(cli.Color("Download completed: "+filename, "green"))
	default:
		usage()
	}
}

func cmdOrg(args []string) {
	if len(args) < 1 {
		usage()
	}
	api := cli.MustAPI()
	switch args[0] {
	case "add":
		if len(args) < 2 {
			usage()
		}
		_, err := api.Org.AddMember(args[len(args)-1], true)
		cli.Die(err)
		fmt.Println(cli.Color("Successfully added the new member", "green"))
	case "remove":
		if len(args) < 2 {
			usage()
		}
		_, err := api.Org.RemoveMember(args[1])
		cli.Die(err)
		fmt.Println(cli.Color("Successfully removed the member", "green"))
	case "info":
		org, err := api.Org.Info()
		cli.Die(err)
		fmt.Println(cli.Color(fmt.Sprint(org["name"]), "cyan"))
		fmt.Print(cli.Dim("Access Level: "))
		fmt.Println(cli.Color(cli.HumanizeAPIPlan(fmt.Sprint(org["upgrade_type"])), "magenta"))
		fmt.Println()
		fmt.Println(cli.Dim("Administrators:"))
		if admins, ok := org["admins"].([]any); ok {
			for _, a := range admins {
				m := asMap(a)
				fmt.Printf(" > %-30s\t%-30v\n", cli.Color(fmt.Sprint(m["username"]), "yellow"), m["email"])
			}
		}
	default:
		usage()
	}
}

func cmdScan(args []string) {
	if len(args) < 1 {
		usage()
	}
	api := cli.MustAPI()
	switch args[0] {
	case "list":
		scans, err := api.Scans(1)
		cli.Die(err)
		fmt.Printf("# %v Scans Total - Showing 10 most recent scans:\n", scans["total"])
		fmt.Printf("# %-20s %-15s %-10s %-15s\n", "Scan ID", "Status", "Size", "Timestamp")
		matches, _ := scans["matches"].([]any)
		n := len(matches)
		if n > 10 {
			n = 10
		}
		for _, s := range matches[:n] {
			m := asMap(s)
			fmt.Printf("%-31s %-24s %-10v %-15v\n", cli.Color(fmt.Sprint(m["id"]), "yellow"), cli.Color(fmt.Sprint(m["status"]), "cyan"), m["size"], m["created"])
		}
	case "protocols":
		p, err := api.Protocols()
		cli.Die(err)
		m := asMap(p)
		for name, desc := range m {
			fmt.Printf("%s%s\n", cli.Color(fmt.Sprintf("%-30s", name), "cyan"), desc)
		}
	case "status":
		if len(args) < 2 {
			usage()
		}
		s, err := api.ScanStatus(args[1])
		cli.Die(err)
		fmt.Println(s["status"])
	case "submit":
		flags, rest := parseFlags(args[1:])
		if len(rest) < 1 {
			usage()
		}
		force := flags["force"] == "true"
		scan, err := api.Scan(rest, force)
		cli.Die(err)
		now := time.Now().Format("2006-01-02 15:04")
		fmt.Printf("\nStarting Shodan scan at %s - %v scan credits left\n", now, scan["credits_left"])
		wait := 20
		if flags["wait"] != "" {
			wait, _ = strconv.Atoi(flags["wait"])
		}
		if wait <= 0 {
			fmt.Println("Scan ID:", scan["id"])
			return
		}
		fmt.Println("Scan ID:", scan["id"])
	case "internet":
		_, rest := parseFlags(args[1:])
		if len(rest) < 2 {
			usage()
		}
		port, _ := strconv.Atoi(rest[0])
		scan, err := api.ScanInternet(port, rest[1])
		cli.Die(err)
		fmt.Println("Submitting Internet scan to Shodan...Done")
		fmt.Println("Scan ID:", scan["id"])
	default:
		usage()
	}
}

