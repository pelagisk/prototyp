package main

import (
	"database/sql"
	"gin-backend/database"

	"log"
)

const fileStorePath = "filestore"

var fileRepository *database.SQLiteRepository

// sets up a database of the provided fileName
func setupDatabase(fileName string) *database.SQLiteRepository {
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

	// set up database
	dbName := "sqlite.db"
	fileRepository = setupDatabase(dbName)

	// set up router
	router := setupRouter()
	router.Run("localhost:8080")
}
