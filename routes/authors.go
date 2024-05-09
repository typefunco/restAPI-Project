package routes

import (
	"net/http"
	"restAPI/author"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAuthor(context *gin.Context) {
	AuthorId, err := strconv.ParseInt(context.Param("id"), 10, 32)

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Can't parse id"})
	}

	Author, err := author.GetAuthorById(int(AuthorId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No author with this id"})
	}

	context.JSON(http.StatusOK, gin.H{"Status": "Author collected", "event": Author})

}

func GetAuthors(context *gin.Context) {
	authors, err := author.GetAuthors()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "could not fetch data"})
		return
	}

	context.JSON(http.StatusOK, authors)
}

func SaveAuthor(context *gin.Context) {
	var author author.Author
	err := context.ShouldBindJSON(&author)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create"})
		return
	}

	err = author.PostAuthors() // Unique ID

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "could not fetch data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Author": "Event created", "event": author})

}

func UpdateAuthor(context *gin.Context) {
	// some code
}

func DeleteAuthors(context *gin.Context) {
	// some code
}
