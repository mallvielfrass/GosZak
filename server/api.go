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
	SearchString string
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
	var getInfo GetInfo

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, m.JsonifyError(err), 500)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &query)
	if err != nil {
		log.Print(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, m.JsonifyError(err), 500)
		return
	}

	result, err := getInfo.Search(query)
	if err != nil {
		log.Print(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, m.JsonifyError(err), 500)
		return
	}

	fmt.Fprintf(w, result)

}
