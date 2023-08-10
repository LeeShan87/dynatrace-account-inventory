package dynatrace

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type DTAuthClient struct {
	url      string
	clientID string
	secret   string
	scopes   string
	resource string
}

var (
	account_api_url = "https://sso.dynatrace.com/sso/oauth2/token"
	accountID       string
	clientID        string
	secret          string
	scopes          string
	resource        string
)

func init() {
	_ = godotenv.Load()
	clientID = os.Getenv("AOA_CLIENT")
	secret = os.Getenv("AOA_SECRET")
	scopes = os.Getenv("TOKEN_SCOPE")
	accountID = os.Getenv("ACOUNT_ID")
	resource = "urn:dtaccount:" + accountID
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

func NewDTAuthClient() *DTAuthClient {
	return &DTAuthClient{
		url:      account_api_url,
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
