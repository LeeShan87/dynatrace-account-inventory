package dynatrace

import (
	"os"

	"github.com/LeeShan87/dynatrace-account-inventory/generated/account"
	"github.com/LeeShan87/dynatrace-account-inventory/utils"
)

func GetApiClient(authClient DTAuthClient, apiURL string) (*account.APIClient, error) {
	accessToken, err := authClient.GetAccessToken()
	if err != nil {
		return nil, err
	}
	config := account.NewConfiguration()
	config.Servers[0].URL = apiURL
	useCache := os.Getenv("USE_HTTP_CACHE")
	if useCache == "true" {
		config.HTTPClient = utils.GetFSCacheClient()
	}
	config.AddDefaultHeader("Authorization", "Bearer "+accessToken)
	client := account.NewAPIClient(config)
	return client, nil
}
