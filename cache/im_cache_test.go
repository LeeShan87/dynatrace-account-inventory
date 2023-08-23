package cache

import "testing"

func TestInMemoryCache(t *testing.T) {
	cache := NewInMemoryCache()

	_, found := cache.Get("key1")
	if found {
		t.Error("Expected key1 to not be found")
	}

	cache.Set("key1", "value1")

	val, found := cache.Get("key1")
	if !found || val != "value1" {
		t.Error("Expected key1 to return value1")
	}
}
