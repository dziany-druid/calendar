package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fixedHolidays, error := LoadEventsFromJson("fixed_holidays.json")

	if error != nil {
		log.Fatal(error)
	}

	calendar := Calendar{
		Name: "Polskie święta",
		Events: fixedHolidays,
	}

	currentYear := time.Now().Year()

	for _, year := range []int{currentYear - 1, currentYear, currentYear + 1} {
		calendar.AddMovableFeast(year)
		calendar.AddDaylightSavingStart(year)
		calendar.AddDaylightSavingEnd(year)
	}

	ical := calendar.ICal()

	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		if request.URL.Path != "/" {
			response.WriteHeader(http.StatusNotFound)

			return
		}

		response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		response.Header().Set("Pragma", "no-cache")
		response.Header().Set("Content-Type", "text/calendar; charset=UTF-8")
		response.Header().Set("Content-Disposition", "inline; filename=holidays.ics")
		response.WriteHeader(http.StatusOK)
		response.Write([]byte(ical))
	})

	port := os.Getenv("CALENDAR_SERVER_PORT")

	if port == "" {
		port = "8080"
	}

	if error := http.ListenAndServe(":" + port, nil); error != nil {
		log.Fatal(error)
	}
}
