package main

import (
	"encoding/csv"
	"os"
	"sync"
)

var mu sync.Mutex

func saveURL(shortURL, longURL string) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.OpenFile("urls.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.Write([]string{shortURL, longURL}); err != nil {
		panic(err)
	}
}

func getOriginalURL(shortURL string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Open("urls.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		if record[0] == shortURL {
			return record[1], true
		}
	}

	return "", false
}
