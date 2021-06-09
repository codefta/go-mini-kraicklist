package controllers

import (
	"net/http"

	storage "github.com/fathisiddiqi/go-mini-kraicklist/storage"
	"github.com/gorilla/mux"
)

type API struct {
	storage *storage.Storage
}

type APIConfigs struct {
	Storage *storage.Storage
}

func NewAPI(configs APIConfigs) (*API, error) {
	// TODO: configs validation
	return &API{storage: configs.Storage}, nil
}

func (a *API) GetHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", a.Home).Methods(http.MethodGet)
	router.HandleFunc("/ads", a.GetAds).Methods(http.MethodGet)
	router.HandleFunc("/ads", a.PostAds).Methods(http.MethodPost)
	router.HandleFunc("/ads/{id}", a.UpdateAds).Methods(http.MethodPut)
	router.HandleFunc("/ads/{id}", a.DeleteAds).Methods(http.MethodDelete)
	router.HandleFunc("/stats", a.Getstats).Methods(http.MethodGet)

	return router
}