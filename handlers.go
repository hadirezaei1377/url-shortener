package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var requestBody struct {
		LongURL string `json:"long_url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortURL := generateShortURL(requestBody.LongURL)
	saveURL(shortURL, requestBody.LongURL)
	cacheURL(shortURL, requestBody.LongURL)

	response := map[string]string{"short_url": fmt.Sprintf("http://localhost:8080/%s", shortURL)}
	json.NewEncoder(w).Encode(response)
}

func expandURLHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := strings.TrimPrefix(r.URL.Path, "/expand/")

	longURL, found := getURLFromCache(shortURL)
	if !found {

		longURL, found = getOriginalURL(shortURL)
		if !found {
			http.Error(w, "URL not found", http.StatusNotFound)
			return
		}

		cacheURL(shortURL, longURL)
	}

	http.Redirect(w, r, longURL, http.StatusFound)
}
