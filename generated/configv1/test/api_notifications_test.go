/*
Dynatrace Configuration API

Testing NotificationsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package configv1

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_configv1_NotificationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NotificationsAPIService CreateNotificationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.NotificationsAPI.CreateNotificationConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationsAPIService DeleteNotificationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.NotificationsAPI.DeleteNotificationConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationsAPIService GetNotificationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NotificationsAPI.GetNotificationConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationsAPIService ListNotificationConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NotificationsAPI.ListNotificationConfigs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationsAPIService UpdateNotificationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.NotificationsAPI.UpdateNotificationConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationsAPIService ValidateCreateNotificationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.NotificationsAPI.ValidateCreateNotificationConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationsAPIService ValidateUpdateNotificationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.NotificationsAPI.ValidateUpdateNotificationConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}