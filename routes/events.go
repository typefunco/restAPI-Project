package routes

import (
	"net/http"
	"restAPI/events"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := events.GetEvents()

	if err != nil {
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

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 32)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error message": "Can't reach path"})
		return
	}

	event, err := events.GetEventById(int(eventId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "Could't get data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Status": "Event collected", "event": event})

}
