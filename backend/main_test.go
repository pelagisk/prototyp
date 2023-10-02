package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-backend/database"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestDb(t *testing.T) {

	dbName := "sqlite_test.db"
	fileRepo := setupDatabase(dbName)

	file := database.Metadata{
		Filename:      "test.jpg",
		Mime:          "image/jpeg",
		Description:   "A PDF document",
		Uploader:      "Axel Gagge",
		UnixTimestamp: time.Now().Unix(),
	}

	if _, err := fileRepo.Create(file); err != nil {
		t.Errorf("Failed to create file with error: %s", err)
	}

	if _, err := fileRepo.GetById(1); err != nil {
		t.Errorf("Failed to get file with error: %s", err)
	}

	os.Remove(dbName)
}

func TestApiGet(t *testing.T) {

	dbName := "sqlite_test.db"
	fileRepository = setupDatabase(dbName)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/files", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "[]", w.Body.String())

	os.Remove(dbName)
}

func TestApiUpload(t *testing.T) {

	// setup test
	dbName := "sqlite_test.db"
	fileRepository = setupDatabase(dbName)
	router := setupRouter()

	// run test

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	var data = map[string]string{
		"description": "An image",
		"uploader":    "Axel Gagge",
	}
	for key, value := range data {
		p, _ := writer.CreateFormField(key)
		io.WriteString(p, value)
	}

	// file with correct header, a little trickier
	// https://stackoverflow.com/questions/74832003/i-cant-add-a-header-to-a-specific-multipart-part-in-golang
	filePath := "testfiles/PLC.jpg"
	file, _ := os.Open(filePath)
	defer file.Close()
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filePath))
	header.Set("Content-Type", "image/jpeg")
	part, err := writer.CreatePart(header)
	if err != nil {
		t.Errorf("Failed to create request with error: %s", err)
	}
	// part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)

	writer.Close()

	wPost := httptest.NewRecorder()
	reqPost, _ := http.NewRequest("POST", "/files", body)
	reqPost.Header.Add("Content-Type", writer.FormDataContentType())
	router.ServeHTTP(wPost, reqPost)

	assert.Equal(t, 201, wPost.Code)

	wGet := httptest.NewRecorder()
	reqGet, _ := http.NewRequest("GET", "/files", nil)
	router.ServeHTTP(wGet, reqGet)

	// match JSON with the metadata struct
	metadata := new(database.Metadata)
	metadatas := []*database.Metadata{metadata}
	err3 := json.Unmarshal(wGet.Body.Bytes(), &metadatas)
	if err3 != nil {
		t.Errorf("Failed to convert metadata JSON to struct with error: %s", err)
	}
	println("Filename: ", metadata.Filename, "Description: ", metadata.Description, "Unix: ", metadata.UnixTimestamp)

	assert.Equal(t, 200, wGet.Code)
	assert.Equal(t, "PLC.jpg", metadata.Filename)

	// tear down test

	os.Remove(dbName)
	// os.Remove(path.Join(fileStorePath, metadata.Filename))
}
