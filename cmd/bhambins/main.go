package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/shmink/BhamBins/internal/bins"
	flag "github.com/spf13/pflag"
)

var (
	postcode string
	uprn     string
	verbose  bool
)

func init() {
	flag.StringVarP(&postcode, "postcode", "p", "", "Postcode of the property")
	flag.StringVarP(&uprn, "uprn", "u", "", "Unique Property Reference Number (see https://www.findmyaddress.co.uk/search)")
	flag.BoolVarP(&verbose, "verbose", "v", false, "Enable verbose debug logging")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `usage: bhambins [options]

Fetch Birmingham City Council bin collection dates.

options:
`)
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, `
example:
  bhambins --postcode "B17 0LY" --uprn "100070285236"
`)
	}
}

func main() {
	flag.Parse()

	if postcode == "" || uprn == "" {
		flag.Usage()
		os.Exit(1)
	}

	collections, err := bins.Fetch(postcode, uprn, verbose)
	if err != nil {
		log.Fatal(err)
	}

	result, _ := json.MarshalIndent(bins.Collections{Bins: collections}, "", "  ")
	fmt.Println(string(result))
}
