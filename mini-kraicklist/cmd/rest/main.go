package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", home).Methods(http.MethodGet)
	router.HandleFunc("/ads", getAds).Methods(http.MethodGet)
	router.HandleFunc("/ads", postAds).Methods(http.MethodPost)

	log.Println("Server running on port :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}