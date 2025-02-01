package main

import (
	"eventbackendproject/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// CREATING SERVER INSTANCE
	server := gin.Default()
	// IMPLEMENTING GET METHOD TO GET THE EVENTS
	server.GET("/events", getEvents)
	// IMPLEMENTING POST METHOD TO CREATE A NEW EVENT
	server.POST("/events", createEvent)
	// RUNNING THE SERVER ON PORT 8080 ON LOCAL HOST
	server.Run(":8080")
}

// IMPLEMENTING getEvents FUNCTION
func getEvents(cntxt *gin.Context) {
	events := model.GetAllEvents()
	cntxt.JSON(http.StatusOK, gin.H{"message": "THis is a GET request", "events": events})
}

// IMPLEMENATING createEvent FUNCTION
func createEvent(cntxt *gin.Context) {
	var event model.Event
	err := cntxt.ShouldBindJSON(&event)
	if err != nil {
		cntxt.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data properly"})
		return
	}
	event.ID = 1
	event.UserID = 1
	event.Save()
	cntxt.JSON(http.StatusCreated, gin.H{"message": "This is a POST Request, event successfully createed", "event": event})
}
