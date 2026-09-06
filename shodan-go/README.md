# shodan-go

Go rewrite of the official [shodan-python](https://github.com/achillean/shodan-python) library.

Shodan is a search engine for Internet-connected devices. This package wraps the REST, Streaming, Exploits, Trends, and Threatnet APIs.

## Install

```bash
go get github.com/salarrbl/DOT/shodan-go
```

## Quick start

```go
package main

import (
	"fmt"
	"log"

	"github.com/salarrbl/DOT/shodan-go"
)

func main() {
	api := shodan.New("MY API KEY")

	ipinfo, err := api.Host([]string{"8.8.8.8"}, false, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipinfo)

	ch, errc := api.SearchCursor(`http.title:"hacked by"`, true, 5, nil)
	for banner := range ch {
		fmt.Println(banner)
	}
	if err := <-errc; err != nil {
		log.Fatal(err)
	}

	ics, err := api.Count("tag:ics", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Industrial Control Systems:", ics["total"])
}
```

Get an API key from https://account.shodan.io

Override the REST base URL with `SHODAN_API_URL` (same as the Python client).

## CLI

```bash
export SHODAN_API_KEY=...
go run ./cmd/shodan host 8.8.8.8
go run ./cmd/shodan search apache
go run ./cmd/shodan count tag:ics
go run ./cmd/shodan info
go run ./cmd/shodan myip
```

## Coverage vs shodan-python

| Area | Methods |
| --- | --- |
| Search | `Search`, `SearchCursor`, `Count`, `SearchFacets`, `SearchFilters`, `SearchTokens` |
| Host | `Host` |
| Scan | `Scan`, `Scans`, `ScanInternet`, `ScanStatus` |
| Alerts | `CreateAlert`, `EditAlert`, `Alerts`, `DeleteAlert`, triggers, notifiers |
| Data / DNS / Labs / Org / Tools | nested APIs matching Python |
| Exploits / Trends | nested APIs |
| Stream | banners, ports, ASN, countries, tags, vulns, custom, alert |
| Threatnet | events, backscatter, SSH activity |
| Helpers | facets, iterate files, screenshots, humanize bytes |

The Python CLI (`shodan` command with converters, world map, etc.) is not fully ported; a small CLI covers host/search/count/info/myip.

## License

MIT, following the original shodan-python library.
