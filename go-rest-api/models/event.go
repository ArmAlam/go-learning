package models

import "time"

var events = []Event{}

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

func (e Event) Save() {
	// TODO: SAVE TO DATABASE
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
