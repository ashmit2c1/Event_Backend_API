package main

import (
	"eventbackendproject/db"
	"eventbackendproject/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.Run(":8080") // local host 8080

}

func getEvents(context *gin.Context) {
	events, err := model.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Coudl not fetch events"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Hello this is a GET request", "event": events})
}

func createEvent(context *gin.Context) {
	var event model.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data properly"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save the data"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "New event created", "event": event})

}
