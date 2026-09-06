package main

import (
	"encoding/json"
	"fmt"
	"os"

	shodan "github.com/salarrbl/DOT/shodan-go"
)

func usage() {
	fmt.Fprintf(os.Stderr, `shodan-go — Go client for the Shodan API

Usage:
  shodan host <ip>
  shodan search <query>
  shodan count <query>
  shodan info
  shodan myip

API key is read from SHODAN_API_KEY.
`)
	os.Exit(2)
}

func mustJSON(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(v)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	key := os.Getenv("SHODAN_API_KEY")
	if key == "" {
		fmt.Fprintln(os.Stderr, "SHODAN_API_KEY is not set")
		os.Exit(1)
	}
	api := shodan.New(key)
	switch os.Args[1] {
	case "host":
		if len(os.Args) < 3 {
			usage()
		}
		v, err := api.Host([]string{os.Args[2]}, false, false)
		die(err)
		mustJSON(v)
	case "search":
		if len(os.Args) < 3 {
			usage()
		}
		v, err := api.Search(os.Args[2], 1, 0, 0, nil, true, nil)
		die(err)
		mustJSON(v)
	case "count":
		if len(os.Args) < 3 {
			usage()
		}
		v, err := api.Count(os.Args[2], nil)
		die(err)
		mustJSON(v)
	case "info":
		v, err := api.Info()
		die(err)
		mustJSON(v)
	case "myip":
		v, err := api.Tools.MyIP()
		die(err)
		mustJSON(v)
	default:
		usage()
	}
}

func die(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
