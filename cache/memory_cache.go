package cache

type memoryCache struct {
	data map[string]string
}

func NewMemoryCache() Cache {
	return &memoryCache{data: make(map[string]string)}
}

func (m *memoryCache) Set(key, value string) {
	m.data[key] = value
}

func (m *memoryCache) Get(key string) (string, bool) {
	value, found := m.data[key]
	return value, found
}
