package main

import (
	"log"
	"net/http"
	"os"
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
		ical := calendar.ICal()
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
