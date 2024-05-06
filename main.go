package main

import (
	"restAPI/db"
	"restAPI/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server) // Setted up server routes
	server.Run()
}
