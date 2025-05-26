package viewmodels

import "time"

type Event struct { //todo add a sort structure to make the timeline in unorreder events
	Id      int
	Start   CustomTime
	End     CustomTime
	Title   string
	Summary string
	Image   string
}

type Career struct {
	TimeLine [][]Event
	Earliest time.Time
}
