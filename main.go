// doc: https://go.dev/doc/tutorial/web-service-gin
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) { // gin.Context is the most important part of Gin. It carries request details, validates and serializes JSON, and more. (Despite the similar name, this is different from Go’s built-in context package.)
	c.IndentedJSON(http.StatusOK, albums) // Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
}

func createAlbum(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id") // Use Context.Param to retrieve the id path parameter from the URL. When you map this handler to a path, you’ll include a placeholder for the parameter in the path.

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusCreated, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default() // Initialize a Gin router using Default.
	router.GET("/albums", getAlbums)
	router.POST("/album", createAlbum)
	router.GET("/album/:id", getAlbum)

	router.Run("localhost:8080") // Use the Run function to attach the router to an http.Server and start the server.
}
