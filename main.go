package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type Album struct {
	id     string  `json:"id"`
	title  string  `json:"title"`
	artist string  `json:"artist"`
	price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{id: "1", title: "Blue Train", artist: "John Coltrane", price: 56.99},
	{id: "2", title: "Jeru", artist: "Gerry Mulligan", price: 17.99},
	{id: "3", title: "Sarah Vaughan and Clifford Brown", artist: "Sarah Vaughan", price: 39.99},
}

// getAlbums responds with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum Album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then retuns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter
	for _, a := range albums {
		if a.id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
