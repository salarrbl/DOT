# shodan-go

Go **copy** of the official [shodan-python](https://github.com/achillean/shodan-python) library (v1.31.0).

Package layout mirrors Python:

| Python | Go |
| --- | --- |
| `shodan.client.Shodan` | `shodan.Client` |
| `shodan.exception.APIError` | `shodan.APIError` |
| `shodan.helpers` | `shodan` helpers |
| `shodan.stream.Stream` | `shodan.Stream` |
| `shodan.threatnet.Threatnet` | `shodan.Threatnet` |
| `shodan.__main__` + `shodan.cli.*` | `cmd/shodan` + `cli/` |

Shodan is a search engine for Internet-connected devices. This package wraps the REST, Streaming, Exploits, Trends, and Threatnet APIs, plus the CLI commands from the Python tool.

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

Same command set as `shodan` in Python (`init`, `info`, `count`, `search`, `download`, `host`, `parse`, `convert`, `domain`, `myip`, `stats`, `stream`, `trends`, `honeyscore`, `alert`, `data`, `org`, `scan`, `version`).

```bash
go run ./cmd/shodan init YOUR_API_KEY
go run ./cmd/shodan host 8.8.8.8
go run ./cmd/shodan search apache
go run ./cmd/shodan count tag:ics
```

API key is stored in `~/.shodan/api_key` or `~/.config/shodan/api_key`, matching Python.

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
