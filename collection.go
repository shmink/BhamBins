package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type BinCollection struct {
	Name       string `json:"name"`
	WebDate    string `json:"webDate"`
	ActualDate string `json:"actualDate"`
}

type BinCollections struct {
	Bins []BinCollection `json:"bins"`
}

func main() {
	postcode := flag.String("p", "", "Postcode for the search")
	uprn := flag.String("u", "", "UPRN for the search")
	flag.Parse()

	if *postcode == "" || *uprn == "" {
		log.Fatal("Postcode and UPRN must be provided - find uprn on https://www.findmyaddress.co.uk/search")
	}

	fetchBins(*postcode, *uprn)
}

func fetchBins(postcode string, uprn string) {
	url := "https://www.birmingham.gov.uk/xfp/form/619"

	c := colly.NewCollector()

	var collections []BinCollection

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundarySNAllVg4gWBSl1jZ")
	})

	c.OnHTML(".data-table tbody tr", func(e *colly.HTMLElement) {
		name := strings.TrimSpace(e.ChildText("td:nth-child(1)"))
		webDate := strings.TrimSpace(e.ChildText("td:nth-child(2)"))

		if name != "" && webDate != "" {
			actualDate := calculateNextDate(webDate)

			collection := BinCollection{
				Name:       name,
				WebDate:    webDate,
				ActualDate: actualDate,
			}
			collections = append(collections, collection)
		}
	})

	c.OnScraped(func(r *colly.Response) {
		binCollections := BinCollections{Bins: collections}

		jsonData, err := json.MarshalIndent(binCollections, "", "    ")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(jsonData))
	})

	formData := fmt.Sprintf(`------WebKitFormBoundarySNAllVg4gWBSl1jZ
Content-Disposition: form-data; name="page"

491
------WebKitFormBoundarySNAllVg4gWBSl1jZ
Content-Disposition: form-data; name="q1f8ccce1d1e2f58649b4069712be6879a839233f_0_0"

%s
------WebKitFormBoundarySNAllVg4gWBSl1jZ
Content-Disposition: form-data; name="q1f8ccce1d1e2f58649b4069712be6879a839233f_1_0"

%s
------WebKitFormBoundarySNAllVg4gWBSl1jZ
Content-Disposition: form-data; name="next"

Next
------WebKitFormBoundarySNAllVg4gWBSl1jZ--`, postcode, uprn)

	formDataBytes := []byte(formData)

	err := c.PostRaw(url, formDataBytes)
	if err != nil {
		log.Fatal(err)
	}
}

func calculateNextDate(webDate string) string {
	parts := strings.Fields(webDate)
	if len(parts) < 2 {
		return ""
	}

	dayOfMonthStr := parts[1] // e.g. 16th

	dayOfMonthStr = strings.Trim(dayOfMonthStr, "()")
	dayOfMonthStr = strings.TrimSuffix(dayOfMonthStr, "th")
	dayOfMonthStr = strings.TrimSuffix(dayOfMonthStr, "st")
	dayOfMonthStr = strings.TrimSuffix(dayOfMonthStr, "nd")
	dayOfMonthStr = strings.TrimSuffix(dayOfMonthStr, "rd")

	dayOfMonth, err := time.Parse("2", dayOfMonthStr)
	if err != nil {
		return ""
	}

	// Take todays date. Loop forward one day at a time until the next occurence is of that actual day is found
	today := time.Now()
	for {
		if today.Day() == dayOfMonth.Day() {
			break
		}
		today = today.AddDate(0, 0, 1)
	}

	return today.Format("2006-01-02")
}
