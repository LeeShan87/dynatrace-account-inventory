package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type DTAuthClient struct {
	url      string
	clientID string
	secret   string
	scopes   string
	resource string
}
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func NewDTAuthClient(url, clientID, secret, scopes, resource string) *DTAuthClient {
	return &DTAuthClient{
		url:      url,
		clientID: clientID,
		secret:   secret,
		scopes:   scopes,
		resource: resource,
	}
}

func (dt *DTAuthClient) GetAccessToken() (string, error) {
	payload := strings.NewReader(fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&scope=%s&resource=%s", dt.clientID, dt.secret, dt.scopes, dt.resource))

	req, err := http.NewRequest("POST", dt.url, payload)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Add("User-Agent", "vscode-restclient")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	var accessTokenResponse AccessTokenResponse
	err = json.Unmarshal(body, &accessTokenResponse)
	if err != nil {
		return "", fmt.Errorf("error decoding response JSON: %w", err)
	}

	return accessTokenResponse.AccessToken, nil
}
