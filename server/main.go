package main

import (
	"log"
	"flag"
	"net/http"
)

func main() {
	var ip string
	var api Api

	flag.StringVar(&ip, "ip", ":8081", "ip address and port to listen on")

	flag.Parse()

	fs := http.FileServer(http.Dir("../client"))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/api", api.Root)
	http.HandleFunc("/api/search", api.Search)

	log.Printf("Server is listening on %s...\n", ip)
	log.Fatal(http.ListenAndServe(ip, nil))
}
