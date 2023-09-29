package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// get a list of all files
func getAllFiles(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "all files",
	})
}

// upload a single file
func uploadFile(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "uploaded",
	})
}

// download a file of the provided ID
func downloadFileById(c *gin.Context) {

	// convert id to integer
	idString := c.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("id to int error: %s", err.Error()))
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("downloaded file of id: %d", id),
	})
}

// delete a file of the provided ID
func deleteFileById(c *gin.Context) {

	// convert id to integer
	idString := c.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("id to int error: %s", err.Error()))
	}

	c.JSON(200, gin.H{
		"message": fmt.Sprintf("deleted file of id: %d", id),
	})
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// only allow CORS in debug mode
	if gin.DebugMode == "debug" {
		log.Print("WARNING: running debug mode. Allowing all origin for CORS. Dangerous in production.")
		router.Use(cors.Default())
	}

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	// define API routes
	router.GET("/files", getAllFiles)
	router.POST("/files", uploadFile)
	router.GET("/files/:id", downloadFileById)
	router.DELETE("/files/:id", deleteFileById)

	return router
}
