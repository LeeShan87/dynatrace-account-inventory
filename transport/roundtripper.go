package transport

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LeeShan87/dynatrace-account-inventory/cache"
)

type CacheRoundTripper struct {
	Proxied http.RoundTripper
	Cache   cache.Cache
}

// CachedResponse represents the structure of the cached HTTP response.
type CachedResponse struct {
	StatusCode int                 `json:"status_code"`
	Headers    map[string][]string `json:"headers"`
	Body       string              `json:"body"`
}

func (c *CacheRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	cacheKey := MD5Hash(req.URL.String())
	if value, found := c.Cache.Get(cacheKey); found {
		var cachedResp CachedResponse
		json.Unmarshal([]byte(value), &cachedResp)
		return &http.Response{
			StatusCode: cachedResp.StatusCode,
			Header:     cachedResp.Headers,
			Body:       ioutil.NopCloser(bytes.NewBufferString(cachedResp.Body)),
			Request:    req,
		}, nil
	}

	resp, err := c.Proxied.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Cache the response
	cachedResp := CachedResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       string(bodyBytes),
	}
	data, _ := json.Marshal(cachedResp)
	c.Cache.Set(cacheKey, string(data))
	resp.Body = ioutil.NopCloser(bytes.NewReader(bodyBytes))

	return resp, nil
}

func MD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
