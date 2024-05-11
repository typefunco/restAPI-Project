package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", saveEvent)
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	server.GET("/authors", GetAuthors)
	server.GET("/authors/:id", GetAuthor)
	server.POST("/authors", SaveAuthor)
	server.PUT("/authors/:id", UpdateAuthor)
	server.DELETE("/authors/:id", DeleteAuthors)

	server.POST("/signup", signup)
	server.POST("/login", login)

}
