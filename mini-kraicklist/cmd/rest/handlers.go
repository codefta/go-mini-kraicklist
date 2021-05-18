package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to mini Kraicklist"))
}

func getAds(w http.ResponseWriter, r *http.Request) {
	lists := getList()
	
	res := map[string]interface{} {
		"success": true,
		"data": map[string]interface{} {
			"ads": lists,
		},
	}

	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}

func postAds(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var list *List
	json.Unmarshal(reqBody, &list)

	err := listValidation(list)

	if err != nil {
		w.Header().Set("Content-Type","application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		res, _ := json.Marshal(ResponseError{
			false,
			http.StatusText(http.StatusBadRequest),
			string(err.Error()),
		})
		fmt.Fprintf(w, string(res))
		return
	}

	listSaved := addList(list)

	res := map[string]interface{} {
		"success": true,
		"data": listSaved,
	}

	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}
