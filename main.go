package main

//package main

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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, word)

}
func index(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "main.html")

}
func about(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "about.html")
}
func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api", api)
	http.HandleFunc("/about", about)
	fs := http.FileServer(http.Dir("etc"))
	http.Handle("/etc/", http.StripPrefix("/etc/", fs))
	fmt.Println("Server is listening...", "\n", "localhost:8080")
	http.ListenAndServe(":8080", nil)

}
