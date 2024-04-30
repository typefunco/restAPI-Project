package main

import (
	"net/http"
	"restAPI/events"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", saveEvent)

	server.Run()
}

func getEvents(context *gin.Context) {
	events := events.GetEvents()
	context.JSON(http.StatusOK, events)
}

func saveEvent(context *gin.Context) {
	var event events.Events
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create"})
		return
	}

	event.ID = 1

	context.JSON(http.StatusOK, gin.H{"Message": "Event created", "event": event})
	event.Save()
}
