package bins

import (
	"strings"
	"time"
)

func calculateNextDate(webDate string) string {
	parts := strings.Fields(webDate)
	if len(parts) < 2 {
		return ""
	}

	dayOfMonthStr := parts[1] // e.g. "(16th)"
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
