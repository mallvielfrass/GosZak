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
	LawNumber []string
	ProcedureType []string
	SortDirection string
	CityName string
	PublishDateFrom int
	PublishDateTo int
	SearchQuery string
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
	fmt.Fprintf(w, "{}")

}

func (m Api) Search(w http.ResponseWriter, r *http.Request) {
	var query SearchQuery

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Print(err)
		http.Error(w, m.JsonifyError(err), 500)
		return
	}

	err = json.Unmarshal(body, &query)
	if err != nil {
		log.Print(err)
		http.Error(w, m.JsonifyError(err), 500)
		return
	}

	result := get(query)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, result)

}
