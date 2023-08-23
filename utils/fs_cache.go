package utils

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/LeeShan87/dynatrace-account-inventory/cache"
	"github.com/LeeShan87/dynatrace-account-inventory/transport"
)

func GetFSCacheClient() *http.Client {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cacheDir := filepath.Join(cwd, ".apiResponses")
	//fmt.Println("Cache directory:", cacheDir)

	fsCache := cache.NewFSCache(cacheDir)
	client := &http.Client{
		Transport: &transport.CacheRoundTripper{
			Proxied: http.DefaultTransport,
			Cache:   fsCache,
		},
	}
	return client
}
