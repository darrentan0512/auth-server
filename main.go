package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"myauthserver/handlers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Register).METHOD("POST")
	r.HandleFunc("/login", handlers.Login).METHOD("POST")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}