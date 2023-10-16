package main

import (
	"fmt"
	"gin-backend/database"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slices"
)

// binds to POST request JSON
type BindStruct struct {
	Description string                `form:"description" binding:"required"`
	Uploader    string                `form:"uploader" binding:"required"`
	Filename    string                `form:"filename"`
	FileHeader  *multipart.FileHeader `form:"file" binding:"required"`
}

// converts from binding struct to the format of the database
func bindStructToMetadata(filename string, bindStruct BindStruct) database.Metadata {

	return database.Metadata{
		Filename:      filename,
		Mime:          bindStruct.FileHeader.Header["Content-Type"][0],
		Description:   bindStruct.Description,
		Uploader:      bindStruct.Uploader,
		UnixTimestamp: time.Now().Unix(),
	}
}

func validateFilename(filename string, header *multipart.FileHeader) (string, bool) {
	if filename == "" {
		return header.Filename, true
	} else {
		// check that the file has correct suffix or else append it to the name
		contentType := header.Header["Content-Type"][0]
		suffixes := map[string]string{
			"application/pdf": "pdf",
			"image/jpeg":      "jpg",
			"application/xml": "xml",
			"text/xml":        "xml",
		}
		suffix, ok := suffixes[contentType]
		if !ok {
			return "", false // in this case, the bool error is false
		}
		if !strings.HasSuffix(filename, suffix) {
			return filename + "." + suffix, true
		} else {
			return filename, true
		}
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

	// validate the filename
	filename, ok := validateFilename(bindStruct.Filename, fileHeader)
	if !ok {
		c.String(http.StatusBadRequest, fmt.Sprintf("was not able to validate filename: %s and header filename: %s", bindStruct.Filename, fileHeader.Filename))
		return
	}

	// save metadata in database
	metadata := bindStructToMetadata(filename, bindStruct)
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

	// get metadata from database
	gotFile, err := fileRepository.GetById(id)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("database error: %s", err.Error()))
		return
	}

	// get file from filestore
	c.File(path.Join(fileStorePath, gotFile.Filename))
}

// delete a file of the provided ID
func deleteFileById(c *gin.Context) {

	// convert id to integer
	idString := c.Param("id")
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("id to int error: %s", err.Error()))
	}

	// get metadata from database
	gotFile, err := fileRepository.GetById(id)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("database error: %s", err.Error()))
		return
	}

	// delete metadata in database
	os.Remove(path.Join(fileStorePath, gotFile.Filename))

	// delete metadata from database
	if err := fileRepository.Delete(id); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("database error: %s", err.Error()))
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("File with id = %d successfully deleted", id))
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
	router.GET("/v1/files", getAllFiles)
	router.POST("/v1/files", uploadFile)
	router.GET("/v1/files/:id", downloadFileById)
	router.DELETE("/v1/files/:id", deleteFileById)

	return router
}
