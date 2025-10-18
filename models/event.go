package models

import "time"

type Events struct {
	ID          int
	Name        string    `binding: "required" json:"name"`
	Description string    `binding: "required" json:"description"`
	Location    string    `binding: "required" json:"location"`
	Date        time.Time `binding: "required" json:"date"`
	UserID      int
}

var events = []Events{}

func (e Events) Save() {
	// later: add it to the database
	events = append(events, e)
}

func GetAllEvents() []Events {
	return events
}
