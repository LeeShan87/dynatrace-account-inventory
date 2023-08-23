package cache

type InMemoryCache struct {
	data map[string]string
}

func NewInMemoryCache() *InMemoryCache {
	return &InMemoryCache{data: make(map[string]string)}
}

func (i *InMemoryCache) Get(key string) (string, bool) {
	val, ok := i.data[key]
	return val, ok
}

func (i *InMemoryCache) Set(key string, value string) {
	i.data[key] = value
}
