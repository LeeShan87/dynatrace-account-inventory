package transport

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LeeShan87/dynatrace-account-inventory/cache"
)

func TestCacheRoundTripper(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))
	defer testServer.Close()

	mockCache := cache.NewInMemoryCache()
	client := &http.Client{
		Transport: &CacheRoundTripper{
			Proxied: http.DefaultTransport,
			Cache:   mockCache,
		},
	}

	resp, _ := client.Get(testServer.URL)
	body, _ := ioutil.ReadAll(resp.Body)

	if string(body) != "Hello, World!" {
		t.Error("Expected response to be 'Hello, World!'")
	}
	key := MD5Hash(testServer.URL)
	var cachedResponse CachedResponse
	cachedVal, _ := mockCache.Get(key)
	json.Unmarshal([]byte(cachedVal), &cachedResponse)
	if cachedResponse.Body != "Hello, World!" {
		t.Error("Expected cache to contain 'Hello, World!'", cachedVal)
	}
}
