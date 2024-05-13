package routes

import (
	"net/http"
	"restAPI/models"
	"restAPI/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "could not fetch data"})
		return
	}

	context.JSON(http.StatusOK, events)
}

func saveEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "No auth token"})
		return
	}

	UserId, err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "No auth token"})
		return
	}

	var event models.Events
	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create"})
		return
	}

	event.UserId = UserId
	err = event.Save() // Unique ID

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "could not fetch data"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"Message": "Event created", "event": event})

}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 32)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error message": "Can't reach path"})
		return
	}

	event, err := models.GetEventById(int(eventId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "Could't get data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Status": "Event collected", "event": event})

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 32)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error message": "Can't parse data from path"})
		return
	}

	_, err = models.GetEventById(int(eventId))

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error message": "Can't collect event from request"})
		return
	}

	var updatedEvent models.Events
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error message": "Can't parse data"})
		return
	}

	updatedEvent.ID = int(eventId)
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "Can't update data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error message": "Can't parse data from path"})
		return
	}

	event, err := models.GetEventById(int(eventId))

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error message": "Can't collect event from db"})
		return
	}

	err = event.Delete(int(eventId))

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"error message": "Can't delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": "Event deleted successfully"})
}
