package handler

import (
	"encoding/json"
	"net/http"
	"url-shortener/storage"
	"url-shortener/utils"
)

var store storage.Storage = storage.NewMemoryStorage()

type shortenRequest struct {
	URL string `json:"url"`
}

type shortenResponse struct {
	ShortenedURL string `json:"shortened_url"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req shortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	shortened := utils.Shorten(req.URL)
	store.Save(shortened, req.URL)

	resp := shortenResponse{ShortenedURL: shortened}
	json.NewEncoder(w).Encode(resp)
}

type resolveRequest struct {
	ShortenedURL string `json:"shortened_url"`
}

type resolveResponse struct {
	URL string `json:"url"`
}

func ResolveURL(w http.ResponseWriter, r *http.Request) {
	var req resolveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	originalURL, found := store.Load(req.ShortenedURL)
	if !found {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	resp := resolveResponse{URL: originalURL}
	json.NewEncoder(w).Encode(resp)
}
