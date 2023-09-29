package main

import (
	"gin-backend/database"
	"os"
	"testing"
	"time"
)

func TestDb(t *testing.T) {

	dbName := "sqlite_test.db"
	fileRepo := SetupDatabase(dbName)

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
