package main

import (
	"log"
	"net/http"
	"url-shortener/handler"
)

func main() {

	http.HandleFunc("/shorten", handler.ShortenURL)
	http.HandleFunc("/resolve", handler.ResolveURL)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
