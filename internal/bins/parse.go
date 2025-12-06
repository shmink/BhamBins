package bins

import (
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func ParseRow(e *colly.HTMLElement) (Collection, bool) {
	fullRow := strings.TrimSpace(e.Text)

	reg := regexp.MustCompile(`^(.*?)(Mon|Tue|Wed|Thu|Fri|Sat|Sun)\s+(.+)$`)
	match := reg.FindStringSubmatch(fullRow)

	name := match[1]
	date := match[3]

	parsedDate, err := time.Parse("02/01/2006", date)
	if err != nil {
		return Collection{}, false
	}

	return Collection{
		Name: name,
		Date: parsedDate.Format("2006-01-02"),
	}, true
}
