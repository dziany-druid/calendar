package main

import (
	"fmt"
	"strings"
	"time"
	"crypto/md5"
)

func (calendar *Calendar) ICal() string {
	var content strings.Builder

	content.WriteString("BEGIN:VCALENDAR\n")
	content.WriteString("VERSION:2.0\n")
	content.WriteString("PRODID:polish-holidays\n")
	content.WriteString("NAME:" + calendar.Name + "\n")
	content.WriteString("X-WR-CALNAME:" + calendar.Name + "\n")

	dtStamp := time.Now().Format("20060102T150405Z")
	previousYear := time.Now().Year() - 1

	for _, event := range calendar.Events {
		content.WriteString("BEGIN:VEVENT\n")
		content.WriteString("UID:" + generateUid(event) + "\n")
		content.WriteString("SUMMARY:" + event.Summary + "\n")

		if event.Description != "" {
			content.WriteString("DESCRIPTION:" + event.Description + "\n")
		}

		content.WriteString("TRANSP:TRANSPARENT\n")
		content.WriteString("DTSTAMP:" + dtStamp + "\n")

		startYear := event.Year
		
		if event.Year <= 0 {
			startYear = previousYear
			content.WriteString("RRULE:FREQ=YEARLY;COUNT=3\n")
		}

		content.WriteString("DTSTART;VALUE=DATE:" + fmt.Sprintf("%d%02d%02d", startYear, event.Month, event.Day) + "\n")
		content.WriteString("END:VEVENT\n")
	}

	content.WriteString("END:VCALENDAR\n")

	return content.String()
}

func generateUid(event Event) string {
	input := fmt.Sprintf("%s%d%d%d", event.Year, event.Summary, event.Month, event.Day)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(input)))

	return hash
}
