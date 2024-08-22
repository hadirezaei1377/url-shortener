package storage

type Storage interface {
	Save(shortenedURL, originalURL string)
	Load(shortenedURL string) (string, bool)
}
