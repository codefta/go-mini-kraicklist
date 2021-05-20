package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	storage *Storage
}

type APIConfigs struct {
	Storage *Storage
}

func NewAPI(configs APIConfigs) (*API, error) {
	// TODO: configs validation
	return &API{storage: configs.Storage}, nil
}

func (a *API) GetHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", a.home).Methods(http.MethodGet)
	router.HandleFunc("/ads", a.getAds).Methods(http.MethodGet)
	router.HandleFunc("/ads", a.postAds).Methods(http.MethodPost)

	return router
}

func (a *API) home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to mini Kraicklist"))
}

func (a *API) getAds(w http.ResponseWriter, r *http.Request) {
	lists, _ := a.storage.GetList()

	res := map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"ads": lists,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}

func (a *API) postAds(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var list *List
	json.Unmarshal(reqBody, &list)

	err := listValidation(list)

	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(ResponseError{
			false,
			http.StatusText(http.StatusBadRequest),
			string(err.Error()),
		})
		fmt.Fprintf(w, string(res))
		return
	}

	listSaved, _ := a.storage.AddList(*list)

	res := map[string]interface{}{
		"success": true,
		"data":    listSaved,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}
