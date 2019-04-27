package main

import (
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type Api struct {}

type SearchQuery struct {
	LawNumber       []string
	ProcedureStatus []string
	SortDirection   string
	SortBy          string
	CityName        string
	PublishDateFrom int64
	PublishDateTo   int64
	PageNumber      int64
	SearchString    string
}

type ServerError struct {
	Error string
}

func (_ Api) JsonifyError(errorToJsonify error) string {
	serverError := ServerError{
		Error: errorToJsonify.Error(),
	}

	errorJson, err := json.Marshal(serverError)
	if err != nil {
		log.Print(err)
		return `{"Error":"undefined error"}`
	}

	return string(errorJson)
}

func (_ Api) Root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, "{}")

}

func (m Api) Search(w http.ResponseWriter, r *http.Request) {
	var query SearchQuery
	var getInfo GetInfo

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		http.Error(w, m.JsonifyError(err), 500)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &query)
	if err != nil {
		log.Print(err)
		http.Error(w, m.JsonifyError(err), 500)
		return
	}

	searchResult, err := getInfo.Search(query)
	if err != nil {
		log.Print(err)
		http.Error(w, m.JsonifyError(err), 500)
		return
	}

	jsonResult, err := json.Marshal(searchResult)
	if err != nil {
		log.Print(err)
		http.Error(w, m.JsonifyError(err), 500)
		return
	}

	fmt.Fprint(w, string(jsonResult))

}
