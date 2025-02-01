package model

import "time"

type Event struct {
	// EVENT ID
	ID int
	// EVENT NAME
	Name string
	// EVENT DESCRIPTION
	Description string
	// EVENT LOCATION
	Location string
	// EVENT DATE TIME
	DateTime time.Time
	// EVENT CREATOR USER ID
	UserID int
}

// METHOD TO SAVE THE EVENT
var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

// METHOD TO GET ALL EVENTS
func GetAllEvents() []Event {
	return events
}
