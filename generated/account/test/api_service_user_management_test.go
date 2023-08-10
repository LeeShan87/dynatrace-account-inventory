/*
Dynatrace Account Management API

Testing ServiceUserManagementAPIService

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

func Test_account_ServiceUserManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceUserManagementAPIService ServiceUsersControllerCreateServiceUserForAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string

		resp, httpRes, err := apiClient.ServiceUserManagementAPI.ServiceUsersControllerCreateServiceUserForAccount(context.Background(), accountUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceUserManagementAPIService ServiceUsersControllerDeleteUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string
		var userUuid string

		httpRes, err := apiClient.ServiceUserManagementAPI.ServiceUsersControllerDeleteUser(context.Background(), accountUuid, userUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceUserManagementAPIService ServiceUsersControllerGetServiceUsersFromAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string

		resp, httpRes, err := apiClient.ServiceUserManagementAPI.ServiceUsersControllerGetServiceUsersFromAccount(context.Background(), accountUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
