package main

func main() {

	// TODO set up database

	// set up router
	router := setupRouter()
	router.Run("localhost:8080")
}
