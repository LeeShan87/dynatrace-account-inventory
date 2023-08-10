/*
Dynatrace Account Management API

Testing UserManagementAPIService

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

func Test_account_UserManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserManagementAPIService UsersControllerAddUserToGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string
		var email string

		httpRes, err := apiClient.UserManagementAPI.UsersControllerAddUserToGroups(context.Background(), accountUuid, email).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserManagementAPIService UsersControllerCreateUserForAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string

		httpRes, err := apiClient.UserManagementAPI.UsersControllerCreateUserForAccount(context.Background(), accountUuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserManagementAPIService UsersControllerGetUserGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string
		var email string

		resp, httpRes, err := apiClient.UserManagementAPI.UsersControllerGetUserGroups(context.Background(), accountUuid, email).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserManagementAPIService UsersControllerGetUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string

		resp, httpRes, err := apiClient.UserManagementAPI.UsersControllerGetUsers(context.Background(), accountUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserManagementAPIService UsersControllerRemoveUserFromAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string
		var email string

		httpRes, err := apiClient.UserManagementAPI.UsersControllerRemoveUserFromAccount(context.Background(), accountUuid, email).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserManagementAPIService UsersControllerRemoveUserFromGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string
		var email string

		httpRes, err := apiClient.UserManagementAPI.UsersControllerRemoveUserFromGroups(context.Background(), accountUuid, email).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserManagementAPIService UsersControllerReplaceUserGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var accountUuid string
		var email string

		httpRes, err := apiClient.UserManagementAPI.UsersControllerReplaceUserGroups(context.Background(), accountUuid, email).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
