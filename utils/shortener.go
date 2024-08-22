package utils

import (
	"crypto/sha256"
	"encoding/base64"
)

func Shorten(url string) string {
	hash := sha256.Sum256([]byte(url))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:8]
	return shortURL
}
