package controllers

import (
	"net/http"

	"github.com/fathisiddiqi/go-mini-kraicklist/cmd/rest/helpers"
)

func (a *API) Getstats(w http.ResponseWriter, r *http.Request) {
	stat, err := a.storage.GetStatistics()
	if err != nil {
		helpers.WriteResp(w, helpers.NewInternalErrorResp(err))
	}

	helpers.WriteResp(w, helpers.NewSuccessResp(stat))
}