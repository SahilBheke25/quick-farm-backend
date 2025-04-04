package main

import (
	"log"
	"net/http"

	"github.com/SahilBheke25/quick-farm-backend/internal/app"
	"github.com/joho/godotenv"

	"github.com/SahilBheke25/quick-farm-backend/internal/repository"

	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found, err : ", err)
		return
	}

	db := repository.InitializeDatabase()
	defer db.Close()

	dependencies := app.InitializeDependencies(db)

	router := app.InitializeRoutes(dependencies)

	log.Println("Server up and running on port: 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
