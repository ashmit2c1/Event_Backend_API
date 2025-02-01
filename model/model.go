package model

import "time"

type Event struct {
	// EVENT ID
	ID int
	// EVENT NAME
	Name string `binding:"required"`
	// EVENT DESCRIPTION
	Description string `binding:"required"`
	// EVENT LOCATION
	Location string `binding:"required"`
	// EVENT DATE TIME
	DateTime time.Time `binding:"required"`
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
