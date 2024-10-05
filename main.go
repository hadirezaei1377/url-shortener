package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", shortenURLHandler)
	http.HandleFunc("/expand/", expandURLHandler)

	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
