package main

import (
	"encoding/json"
	"io/ioutil"
)

type Event struct {
	Summary string `json:"summary"`
	Description string `json:"description"`
	Year int `json:"-"`
	Month int `json:"month"`
	Day int `json:"day"`
}

type Calendar struct {
	Name string
	Events []Event
}

func (calendar *Calendar) AddEvent(event Event) {
	calendar.Events = append(calendar.Events, event)
}

func LoadEventsFromJson(fileName string) ([]Event, error) {
	fileContent, error := ioutil.ReadFile(fileName)

	if error != nil {
		return nil, error
	}

	var events []Event

	error = json.Unmarshal(fileContent, &events)

	if error != nil {
		return nil, error
	}

	return events, nil
}
