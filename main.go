package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id`
	Title  string  `json:"title`
	Artist string  `json:"artist`
	Price  float64 `json:"price`
}

// albums slice to seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Jerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and ", Artist: "Jerry Mulligan", Price: 17.99},
}

func main() {
	router := gin.Default()          // Initialize a Gin router
	router.GET("/albums", getAlbums) // Associate GET HTTP method and album path with a handler function
	router.GET("albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080") // Attach router to http.Server and start server
}

// getAlbums responds with the list of all albums as JSON
// gin.Context carries request details, validates and serializes JSON, and more
// IndentedJSON serializes the struct into JSON and add it to the response
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum) // Add new album to the slice
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
