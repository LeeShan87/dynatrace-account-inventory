/*
Dynatrace Configuration API

Testing RUMMobileAndCustomApplicationConfigurationAPIService

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

func Test_configv1_RUMMobileAndCustomApplicationConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService CreateMobileApplicationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.CreateMobileApplicationConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService CreateMobileKeyUserAction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var actionName string

		resp, httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.CreateMobileKeyUserAction(context.Background(), applicationId, actionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService CreateSessionProperty", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.CreateSessionProperty(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService DeleteMobileApplicationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.DeleteMobileApplicationConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService DeleteMobileKeyUserAction", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var actionName string

		httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.DeleteMobileKeyUserAction(context.Background(), applicationId, actionName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService DeleteSessionProperty", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var key string

		httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.DeleteSessionProperty(context.Background(), applicationId, key).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService GetMobileApplicationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.GetMobileApplicationConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService GetSessionProperty", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var key string

		resp, httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.GetSessionProperty(context.Background(), applicationId, key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService ListMobileApplicationConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ListMobileApplicationConfigs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService ListMobileKeyUserActions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ListMobileKeyUserActions(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService ListSessionProperties", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		resp, httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ListSessionProperties(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService UpdateMobileApplicationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.UpdateMobileApplicationConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService UpdateSessionProperty", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var key string

		resp, httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.UpdateSessionProperty(context.Background(), applicationId, key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService ValidateCreateMobileApplicationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ValidateCreateMobileApplicationConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService ValidateCreateSessionProperty", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string

		httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ValidateCreateSessionProperty(context.Background(), applicationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService ValidateUpdateMobileApplicationConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ValidateUpdateMobileApplicationConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMMobileAndCustomApplicationConfigurationAPIService ValidateUpdateSessionProperty", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var key string

		httpRes, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ValidateUpdateSessionProperty(context.Background(), applicationId, key).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
