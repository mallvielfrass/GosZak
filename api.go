package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func api(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Body)
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	word := string(b)
	fmt.Println(word)
	result := get(word)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, result)

}
