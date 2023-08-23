package cache

import (
	"os"
	"testing"
)

func TestFSCache(t *testing.T) {
	dir := "./test_cache"
	os.Mkdir(dir, os.ModePerm)
	defer os.RemoveAll(dir)

	cache := NewFSCache(dir)

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
