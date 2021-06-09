package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fathisiddiqi/go-mini-kraicklist/cmd/rest/helpers"
	model "github.com/fathisiddiqi/go-mini-kraicklist/cmd/rest/models"
	"github.com/gorilla/mux"
)

func (a *API) Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to mini Kraicklist"))
}

func (a *API) GetAds(w http.ResponseWriter, r *http.Request) {
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	
	if limit < 1 {
		limit = 5
	}

	lists, err := a.storage.GetList(limit)

	if err != nil {
		helpers.WriteResp(w, helpers.NewInternalErrorResp(err))
		return
	}
	helpers.WriteResp(w, helpers.NewSuccessResp(map[string]interface{}{
		"ads": lists,
	}))
}

func (a *API) PostAds(w http.ResponseWriter, r *http.Request) {
	// parse request reqBody
	var reqBody postNewAdBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		helpers.WriteResp(w, helpers.NewBadRequestErrorResp(err.Error()))
		return
	}
	// validate request body
	err = reqBody.Validate()
	if err != nil {
		helpers.WriteResp(w, helpers.NewBadRequestErrorResp(err.Error()))
		return
	}
	// save list to db
	list := model.List{
		Title: reqBody.Title,
		Body:  reqBody.Body,
		Tags:  reqBody.Tags,
	}
	listSaved, err := a.storage.AddList(list)
	if err != nil {
		helpers.WriteResp(w, helpers.NewInternalErrorResp(err))
		return
	}
	// output success response
	helpers.WriteResp(w, helpers.NewSuccessResp(listSaved))
}

func (a *API) UpdateAds(w http.ResponseWriter, r *http.Request) {
	var reqBody updateAdBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		helpers.WriteResp(w, helpers.NewBadRequestErrorResp(err.Error()))
		return
	}

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil || id < 1 {
		helpers.WriteResp(w, helpers.NewBadRequestErrorResp("at least one parameter must be specified"))
		return
	}

	_, err = a.storage.FindListById(id)
	if err != nil {
		helpers.WriteResp(w, helpers.NewNotFoundErrorResp(err.Error()))
		return
	}

	list := model.List{
		ID: id,
		Title: reqBody.Title,
		Body: reqBody.Body,
		Tags: reqBody.Tags,
	}

	listUpdated, err := a.storage.UpdateList(list)
	if err != nil {
		helpers.WriteResp(w, helpers.NewInternalErrorResp(err))
		return
	}

	helpers.WriteResp(w, helpers.NewSuccessResp(listUpdated))
}

func (a *API) DeleteAds(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		helpers.WriteResp(w, helpers.NewBadRequestErrorResp("at least one parameter must be specified"))
		return
	}

	_, err = a.storage.FindListById(id)
	if err != nil {
		helpers.WriteResp(w, helpers.NewNotFoundErrorResp(err.Error()))
		return
	}

	listDeleted, err := a.storage.DeleteList(id)
	if err != nil {
		helpers.WriteResp(w, helpers.NewInternalErrorResp(err))
		return
	}

	helpers.WriteResp(w, helpers.NewSuccessResp(listDeleted))
}

type updateAdBody struct {
	Title string `json:"title"`
	Body string `json:"body"`
	Tags []string `json:"tags"`
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
