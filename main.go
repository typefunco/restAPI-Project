package main

import (
	"fmt"
	"net/http"
	"restAPI/db"
	"restAPI/events"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", saveEvent)

	server.Run()
}

func getEvents(context *gin.Context) {
	events, err := events.GetEvents()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "could not fetch data"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func saveEvent(context *gin.Context) {
	var event events.Events
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create"})
		return
	}

	err = event.Save() // Unique ID

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "could not fetch data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event created", "event": event})

}
