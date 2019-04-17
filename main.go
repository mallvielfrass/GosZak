package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/index", index)
	http.HandleFunc("/api", api)
	fs := http.FileServer(http.Dir("etc"))
	http.Handle("/etc/", http.StripPrefix("/etc/", fs))
	fmt.Println("Server is listening...", "\n", "localhost:8080")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
