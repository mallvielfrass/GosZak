package main

import (
	"log"
	"flag"
	"os"
	"path/filepath"
	"net/http"
)

func main() {
	var ip string
	var hostDir string
	var api Api

	flag.StringVar(&ip, "ip", ":8081", "ip address and port to listen on")
	flag.StringVar(&hostDir, "host-dir", "../client", "directory to host")
	flag.Parse()

	if !filepath.IsAbs(hostDir) {
		ex, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}

		exPath := filepath.Dir(ex)
		hostDir = filepath.Join(exPath, hostDir)
	}

	fs := http.FileServer(http.Dir(hostDir))
	http.Handle("/", http.StripPrefix("/", fs))
	http.HandleFunc("/api", api.Root)
	http.HandleFunc("/api/search", api.Search)

	log.Printf("Server is listening on %s...\n", ip)
	log.Fatal(http.ListenAndServe(ip, nil))
}
