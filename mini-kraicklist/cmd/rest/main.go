package main

import (
	"log"
	"net/http"
)

const addr = ":8080"

func main() {
	// initialize storage
	storage, err := NewStorage(StorageConfigs{
		DBUser: GoDotEnv("MYSQL_USER"),
		DBPass: GoDotEnv("MYSQL_PASSWORD"),
		DBName: GoDotEnv("MYSQL_DATABASE"),
	})
	if err != nil {
		log.Fatalf("unable to initialize storage due: %v", err)
	}
	// initialize api
	api, err := NewAPI(APIConfigs{
		Storage: storage,
	})
	if err != nil {
		log.Fatalf("unable to initialize api due: %v", err)
	}
	// execute http server
	log.Printf("Server running on port %v", addr)
	err = http.ListenAndServe(addr, api.GetHandler())
	if err != nil {
		log.Fatalf("unable to execute http server due: %v", err)
	}
}
