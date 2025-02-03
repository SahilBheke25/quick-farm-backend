package main

import (
	"fmt"
	"log"
	"net/http"

	Equipment "github.com/SahilBheke25/ResourceSharingApplication/internal/app/equipment"
	"github.com/SahilBheke25/ResourceSharingApplication/internal/app/login"

	repository "github.com/SahilBheke25/ResourceSharingApplication/internal/repository"

	_ "github.com/lib/pq"
)

func main() {

	// Creating DB connection
	repository.InitializeDatabase()
	defer repository.DB.Close()

	mux := http.DefaultServeMux
	mux.HandleFunc("POST /login", login.Verify)
	mux.HandleFunc("POST /register", login.Register)
	mux.HandleFunc("POST /equipment", Equipment.PostLendEquipmentHandler)
	mux.HandleFunc("GET /equipments", Equipment.GetAllEquipmentHandler)

	fmt.Println("listning to port 3000")
	log.Fatal(http.ListenAndServe(":3000", mux))

}
