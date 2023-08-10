package main

import (
	"fmt"
	"os"

	auth "github.com/LeeShan87/dynatrace-account-inventory/dynatrace"
	"github.com/joho/godotenv"
)

var (
	url       = "https://sso.dynatrace.com/sso/oauth2/token"
	accountID string
	clientID  string
	secret    string
	scopes    string
	resource  string
)

func main() {
	_ = godotenv.Load()
	clientID = os.Getenv("AOA_CLIENT")
	secret = os.Getenv("AOA_SECRET")
	scopes = os.Getenv("TOKEN_SCOPE")
	accountID = os.Getenv("ACOUNT_ID")
	resource = "urn:dtaccount:" + accountID
	fmt.Println("ClientID", clientID)
	fmt.Println("Secret", secret)
	fmt.Println("Scopes", scopes)
	fmt.Println("Resource", resource)
	authClient := auth.NewDTAuthClient(
		url,
		clientID,
		secret,
		scopes,
		resource,
	)
	accessToken, err := authClient.GetAccessToken()
	if err != nil {
		panic(err)
	}
	fmt.Println("AccessToken", accessToken)
}
