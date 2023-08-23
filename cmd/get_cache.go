package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/LeeShan87/dynatrace-account-inventory/cache"
	"github.com/LeeShan87/dynatrace-account-inventory/transport"
)

func main() {
	// get the first command line argument
	// and store it in the variable "name"
	name := os.Args[1]
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cacheDir := filepath.Join(cwd, ".apiResponses")

	fsCache := cache.NewFSCache(cacheDir)
	value, found := fsCache.Get(name)
	if !found {
		return
	}
	var cachedResp transport.CachedResponse
	json.Unmarshal([]byte(value), &cachedResp)
	resp := &http.Response{
		StatusCode: cachedResp.StatusCode,
		Header:     cachedResp.Headers,
		Body:       ioutil.NopCloser(bytes.NewBufferString(cachedResp.Body)),
		Request:    nil,
	}
	body, _ := ioutil.ReadAll(resp.Body)

	var prettyBuffer bytes.Buffer
	json.Indent(&prettyBuffer, body, "", "  ")
	prettyJSON := prettyBuffer.Bytes()
	fmt.Println(string(prettyJSON))
}
