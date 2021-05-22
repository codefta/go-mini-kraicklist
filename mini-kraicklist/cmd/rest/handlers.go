package main

import (
	"encoding/json"
	"fmt"
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
	lists, err := a.storage.GetList()
	if err != nil {
		writeResp(w, newInternalErrorResp(err))
		return
	}
	writeResp(w, newSuccessResp(map[string]interface{}{
		"ads": lists,
	}))
}

func (a *API) postAds(w http.ResponseWriter, r *http.Request) {
	// parse request reqBody
	var reqBody postNewAdBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		writeResp(w, newBadRequestErrorResp(err.Error()))
		return
	}
	// validate request body
	err = reqBody.Validate()
	if err != nil {
		writeResp(w, newBadRequestErrorResp(err.Error()))
		return
	}
	// save list to db
	list := List{
		Title: reqBody.Title,
		Body:  reqBody.Body,
		Tags:  reqBody.Tags,
	}
	listSaved, err := a.storage.AddList(list)
	if err != nil {
		writeResp(w, newInternalErrorResp(err))
		return
	}
	// output success response
	writeResp(w, newSuccessResp(listSaved))
}

type postNewAdBody struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

func (b postNewAdBody) Validate() error {
	if len(b.Title) == 0 {
		return fmt.Errorf("field `title` cannot be empty")
	}
	if len(b.Body) == 0 {
		return fmt.Errorf("field `body` cannot be empty")
	}
	return nil
}

type APIResp struct {
	StatusCode int         `json:"-"`
	Success    bool        `json:"success"`
	Data       interface{} `json:"data,omitempty"`
	Err        string      `json:"err,omitempty"`
	Message    string      `json:"message,omitempty"`
}

func newSuccessResp(data interface{}) APIResp {
	return APIResp{
		StatusCode: http.StatusOK,
		Success:    true,
		Data:       data,
	}
}

func newErrorResp(statusCode int, errCode string, message string) APIResp {
	return APIResp{
		StatusCode: statusCode,
		Err:        errCode,
		Message:    message,
	}
}

func newBadRequestErrorResp(message string) APIResp {
	return newErrorResp(http.StatusBadRequest, "ERR_BAD_REQUEST", message)
}

func newInternalErrorResp(err error) APIResp {
	return newErrorResp(http.StatusInternalServerError, "ERR_INTERNAL_ERROR", err.Error())
}

func writeResp(w http.ResponseWriter, resp APIResp) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(resp)
}
