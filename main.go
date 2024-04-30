package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run()
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"Message": "True"})
}
