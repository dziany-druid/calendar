package main

import "time"

func (calendar *Calendar) AddMovableFeast(year int) {
	easter := easterDate(year)
	ashWednesday := easter.AddDate(0, 0, -46)
	fatThursday := ashWednesday.AddDate(0, 0, -6)

	calendar.AddEvent(
		Event{
			Summary: "Tłusty Czwartek",
			Year: fatThursday.Year(),
			Month: int(fatThursday.Month()),
			Day: fatThursday.Day(),
		},
	)

	shroveTuesday := ashWednesday.AddDate(0, 0, -1)

	calendar.AddEvent(
		Event{
			Summary: "Ostatki",
			Year: shroveTuesday.Year(),
			Month: int(shroveTuesday.Month()),
			Day: shroveTuesday.Day(),
		},
	)

	calendar.AddEvent(
		Event{
			Summary: "Środa Popielcowa",
			Year: ashWednesday.Year(),
			Month: int(ashWednesday.Month()),
			Day: ashWednesday.Day(),
		},
	)

	maundyThursday := easter.AddDate(0, 0, -3)

	calendar.AddEvent(
		Event{
			Summary: "Wielki Czwartek",
			Year: maundyThursday.Year(),
			Month: int(maundyThursday.Month()),
			Day: maundyThursday.Day(),
		},
	)

	goodFriday := easter.AddDate(0, 0, -2)

	calendar.AddEvent(
		Event{
			Summary: "Wielki Piątek",
			Year: goodFriday.Year(),
			Month: int(goodFriday.Month()),
			Day: goodFriday.Day(),
		},
	)

	holySaturday := easter.AddDate(0, 0, -1)

	calendar.AddEvent(
		Event{
			Summary: "Wielka Sobota",
			Year: holySaturday.Year(),
			Month: int(holySaturday.Month()),
			Day: holySaturday.Day(),
		},
	)

	calendar.AddEvent(
		Event{
			Summary: "Wielkanoc",
			Description: "Dzień wolny od pracy.",
			Year: easter.Year(),
			Month: int(easter.Month()),
			Day: easter.Day(),
		},
	)

	easterMonday := easter.AddDate(0, 0, 1)

	calendar.AddEvent(
		Event{
			Summary: "Poniedziałek Wielkanocny",
			Description: "Dzień wolny od pracy.",
			Year: easterMonday.Year(),
			Month: int(easterMonday.Month()),
			Day: easterMonday.Day(),
		},
	)

	pentecost := easter.AddDate(0, 0, 49)

	calendar.AddEvent(
		Event{
			Summary: "Zesłanie Ducha Świętego (Zielone Świątki)",
			Description: "Dzień wolny od pracy.",
			Year: pentecost.Year(),
			Month: int(pentecost.Month()),
			Day: pentecost.Day(),
		},
	)

	corpusChristi := easter.AddDate(0, 0, 60)

	calendar.AddEvent(
		Event{
			Summary: "Boże Ciało",
			Description: "Dzień wolny od pracy.",
			Year: corpusChristi.Year(),
			Month: int(corpusChristi.Month()),
			Day: corpusChristi.Day(),
		},
	)
}

func (calendar *Calendar) AddDaylightSavingStart(year int) {
	lastSundayOfMarch := lastSundayOfMonth(year, time.March)

	calendar.AddEvent(
		Event{
			Summary: "Zmiana czasu z zimowego na letni",
			Description: "Zmiana czasu z godziny 2.00 na 3.00 w nocy, czyli o jedną pełną godzinę do przodu.",
			Year: lastSundayOfMarch.Year(),
			Month: int(lastSundayOfMarch.Month()),
			Day: lastSundayOfMarch.Day(),
		},
	)
}

func (calendar *Calendar) AddDaylightSavingEnd(year int) {
	lastSundayOfOctober := lastSundayOfMonth(year, time.October)

	calendar.AddEvent(
		Event{
			Summary: "Zmiana czasu z letniego na zimowy",
			Description: "Zmiana czasu z godziny 3.00 na 2.00 w nocy, czyli o jedną pełną godzinę do tyłu.",
			Year: lastSundayOfOctober.Year(),
			Month: int(lastSundayOfOctober.Month()),
			Day: lastSundayOfOctober.Day(),
		},
	)
}

func easterDate(year int) time.Time {
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19 * a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2 * e + 2 * i - h - k) % 7
	m := (a + 11 * h + 22 * l) / 451
	month := (h + l - 7 * m + 114) / 31
	day := ((h + l - 7 * m + 114) % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func lastSundayOfMonth(year int, month time.Month) time.Time {
	firstDayOfMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)
	daysToSubtract := int(lastDayOfMonth.Weekday())
	
	return lastDayOfMonth.AddDate(0, 0, -daysToSubtract)
}
