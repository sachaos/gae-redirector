package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var targetHost = ""

func init() {
	targetHost = os.Getenv("TARGET_HOST")
	if targetHost == "" {
		panic("TARGET_HOST is required")
	}
}

func redirector(w http.ResponseWriter, r *http.Request) {
	target := targetHost + r.URL.EscapedPath()

	w.Header().Set("Location", target)
	w.WriteHeader(http.StatusFound)
}

func main() {
	http.HandleFunc("/", redirector)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
