package main

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "main.html")

}
func about(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "about.html")
}
