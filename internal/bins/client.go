package bins

import (
	"fmt"

	"github.com/gocolly/colly"
)

var formURL = "https://www.birmingham.gov.uk/xfp/form/619"

const (
	postcodeField = "q1f8ccce1d1e2f58649b4069712be6879a839233f_0_0"
	uprnField     = "q1f8ccce1d1e2f58649b4069712be6879a839233f_1_0"
)

func Fetch(postcode, uprn string, _ bool) ([]Collection, error) {
	c := colly.NewCollector(colly.MaxDepth(1))
	c.UserAgent = "Mozilla/5.0"

	var result []Collection

	c.OnHTML(".data-table tbody tr", func(e *colly.HTMLElement) {
		if row, ok := ParseRow(e); ok {
			result = append(result, row)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		if r.URL.String() != formURL {
			r.Headers.Set("Referer", formURL)
		}
	})

	if err := c.Visit(formURL); err != nil {
		return nil, fmt.Errorf("visit: %w", err)
	}

	fields := map[string][]byte{
		"page":        []byte("491"),
		postcodeField: []byte(postcode),
		uprnField:     []byte(uprn),
		"next":        []byte("Next"),
	}
	if err := c.PostMultipart(formURL, fields); err != nil {
		return nil, fmt.Errorf("post: %w", err)
	}

	return result, nil
}
