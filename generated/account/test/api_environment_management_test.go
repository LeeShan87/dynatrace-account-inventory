/*
Dynatrace Account Management API

Testing EnvironmentManagementAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package account

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_account_EnvironmentManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EnvironmentManagementAPIService EnvironmentResourcesControllerGetEnvironmentResources", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string

		resp, httpRes, err := apiClient.EnvironmentManagementAPI.EnvironmentResourcesControllerGetEnvironmentResources(context.Background(), accountUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
