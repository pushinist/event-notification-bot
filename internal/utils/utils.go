package utils

import (
	"fmt"
	"regexp"
	"time"
)

var (
	dateRegex = regexp.MustCompile(`Дата:\s*(\d{2})\.(\d{2})`)
	timeRegex = regexp.MustCompile(`Время:\s*(\d{2}):(\d{2})`)
)

func TryParseEventTime(text string) (time.Time, bool) {
	dateMatch := dateRegex.FindStringSubmatch(text)
	timeMatch := timeRegex.FindStringSubmatch(text)

	if len(dateMatch) != 3 || len(timeMatch) != 3 {
		return time.Time{}, false
	}

	day := dateMatch[1]
	month := dateMatch[2]
	hour := timeMatch[1]
	minute := timeMatch[2]

	currentYear := time.Now().Year()

	dateStr := fmt.Sprintf("%d-%s-%s %s:%s", currentYear, month, day, hour, minute)

	t, err := time.ParseInLocation("2006-01-02 15:04", dateStr, time.Local)
	if err != nil {
		return time.Time{}, false
	}

	if t.Before(time.Now()) {
		t.AddDate(1, 0, 0)
	}

	return t, true
}
