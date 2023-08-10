package dynatrace

import (
	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
)

func GetApiClient(authClient DTAuthClient, apiURL string) (*account.APIClient, error) {
	accessToken, err := authClient.GetAccessToken()
	if err != nil {
		return nil, err
	}
	config := account.NewConfiguration()
	config.Servers[0].URL = apiURL
	config.AddDefaultHeader("Authorization", "Bearer "+accessToken)
	client := account.NewAPIClient(config)
	return client, nil
}
