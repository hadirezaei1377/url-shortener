package storage

type memoryStorage struct {
	data map[string]string
}

func NewMemoryStorage() Storage {
	return &memoryStorage{data: make(map[string]string)}
}

func (m *memoryStorage) Save(shortenedURL, originalURL string) {
	m.data[shortenedURL] = originalURL
}

func (m *memoryStorage) Load(shortenedURL string) (string, bool) {
	originalURL, found := m.data[shortenedURL]
	return originalURL, found
}
