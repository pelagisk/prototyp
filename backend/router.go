package main

import (
	"fmt"
	"gin-backend/database"
	"log"
	"mime/multipart"
	"net/http"
	"path"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

// binds to POST request JSON
type BindStruct struct {
	Description string                `form:"description" binding:"required"`
	Uploader    string                `form:"uploader" binding:"required"`
	FileHeader  *multipart.FileHeader `form:"file" binding:"required"`
}

// converts from binding struct to the format of the database
func bindStructToMetadata(bindStruct BindStruct) database.Metadata {
	return database.Metadata{
		Filename:      bindStruct.FileHeader.Filename,
		Mime:          bindStruct.FileHeader.Header["Content-Type"][0],
		Description:   bindStruct.Description,
		Uploader:      bindStruct.Uploader,
		UnixTimestamp: time.Now().Unix(),
	}
}

// get a list of all files
func getAllFiles(c *gin.Context) {
	all, err := fileRepository.All()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("database error: %s", err.Error()))
		return
	}
	if all == nil {
		c.IndentedJSON(http.StatusOK, []string{})
	} else {
		c.IndentedJSON(http.StatusOK, all)
	}
}

// upload a single file
func uploadFile(c *gin.Context) {

	var bindStruct BindStruct

	// bind file
	if err := c.ShouldBind(&bindStruct); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("err: %s", err.Error()))
		return
	}

	fileHeader := bindStruct.FileHeader

	// check if extension is allowed
	contentType := fileHeader.Header["Content-Type"][0]
	allowedContentTypes := []string{"application/pdf", "image/jpeg", "application/xml", "text/xml"}
	if !slices.Contains(allowedContentTypes, contentType) {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Content type: %s is not allowed!", contentType))
		return
	}

	// save metadata in database
	metadata := bindStructToMetadata(bindStruct)
	createdFile, err := fileRepository.Create(metadata)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("database error: %s", err.Error()))
		return
	}

	// save uploaded file in file store
	dst := path.Join(fileStorePath, createdFile.Filename)
	if err := c.SaveUploadedFile(fileHeader, dst); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	c.IndentedJSON(http.StatusCreated, createdFile)
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
