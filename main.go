package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/api", api)
	fs := http.FileServer(http.Dir("etc"))
	http.Handle("/etc/", http.StripPrefix("/etc/", fs))
	fmt.Println("Server is listening...", "\n", "localhost:8080")
	http.ListenAndServe(":8080", nil)

}
