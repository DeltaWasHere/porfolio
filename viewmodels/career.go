package viewmodels

import "time"

type Event struct {
	id    int
	start time.Time
	end   time.Time
	title string
}

type EventDetails struct {
	id      int
	summary string
	image   string
}
