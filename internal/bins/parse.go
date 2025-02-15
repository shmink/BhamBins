package bins

import (
	"strings"

	"github.com/gocolly/colly"
)

func ParseRow(e *colly.HTMLElement) (Collection, bool) {
	name := strings.TrimSpace(e.ChildText("td:nth-child(1)"))
	webDate := strings.TrimSpace(e.ChildText("td:nth-child(2)"))

	if name == "" || webDate == "" {
		return Collection{}, false
	}

	return Collection{
		Name:       name,
		WebDate:    webDate,
		ActualDate: calculateNextDate(webDate),
	}, true
}
