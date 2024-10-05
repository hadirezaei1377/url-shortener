package main

import (
	"crypto/sha1"
	"encoding/base64"
)

func generateShortURL(longURL string) string {
	h := sha1.New()
	h.Write([]byte(longURL))
	hash := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return hash[:8] // just advanced 8 characters
}
