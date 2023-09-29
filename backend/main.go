package main

import (
	"database/sql"
	"gin-backend/database"

	"log"
)

var fileRepository *database.SQLiteRepository

// sets up a database of the provided fileName
func SetupDatabase(fileName string) *database.SQLiteRepository {
	db, err := sql.Open("sqlite3", fileName)
	if err != nil {
		log.Fatal(err)
	}
	fileRepo := database.NewSQLiteRepository(db)
	if err := fileRepo.Migrate(); err != nil {
		log.Fatal(err)
	}
	return fileRepo
}

func main() {

	// TODO set up database

	// set up router
	router := setupRouter()
	router.Run("localhost:8080")
}
