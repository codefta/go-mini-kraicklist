package main

import (
	"log"
	"net/http"
	"os"

	"github.com/fathisiddiqi/go-mini-kraicklist/cmd/rest/controllers"
	"github.com/fathisiddiqi/go-mini-kraicklist/storage"
)

const addr = ":8080"

const (
	envKeyDBHost = "MYSQL_HOST"
	envKeyDBPort = "MYSQL_PORT"
	envKeyDBUser = "MYSQL_USER"
	envKeyDBPass = "MYSQL_PASSWORD"
	envKeyDBName = "MYSQL_DATABASE"
)

func main() {
	// initialize storage
	s, err := storage.NewStorage(storage.StorageConfigs{
		DBHost: os.Getenv(envKeyDBHost),
		DBPort: os.Getenv(envKeyDBPort),
		DBUser: os.Getenv(envKeyDBUser),
		DBPass: os.Getenv(envKeyDBPass),
		DBName: os.Getenv(envKeyDBName),
	})
	if err != nil {
		log.Fatalf("unable to initialize storage due: %v", err)
	}
	// initialize api
	api, err := controllers.NewAPI(controllers.APIConfigs{
		Storage: s,
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
