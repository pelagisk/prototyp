package main

import (
	"bytes"
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

func TestGetFiles(t *testing.T) {

	dbName := "sqlite_test.db"
	fileRepository = setupDatabase(dbName)
	router := setupRouter()

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
	filePath := "/Users/axelgagge/PLC.jpg"
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

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/files", body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	router.ServeHTTP(w, req)
	println(w.Body.String())

	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/files", nil)
	router.ServeHTTP(w2, req2)

	println(w2.Body.String())
	assert.Equal(t, 200, w2.Code)
	assert.Equal(t, "[]", w2.Body.String())

	os.Remove(dbName)
}
