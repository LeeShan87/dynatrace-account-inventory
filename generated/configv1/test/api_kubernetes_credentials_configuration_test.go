/*
Dynatrace Configuration API

Testing KubernetesCredentialsConfigurationAPIService

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

func Test_configv1_KubernetesCredentialsConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test KubernetesCredentialsConfigurationAPIService CreateKubernetesCredentialsConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesCredentialsConfigurationAPI.CreateKubernetesCredentialsConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesCredentialsConfigurationAPIService DeleteKubernetesCredentialsConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.KubernetesCredentialsConfigurationAPI.DeleteKubernetesCredentialsConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesCredentialsConfigurationAPIService GetKubernetesCredentialsConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.KubernetesCredentialsConfigurationAPI.GetKubernetesCredentialsConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesCredentialsConfigurationAPIService ListKubernetesCredentialsConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.KubernetesCredentialsConfigurationAPI.ListKubernetesCredentialsConfigs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesCredentialsConfigurationAPIService UpdateKubernetesCredentialsConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.KubernetesCredentialsConfigurationAPI.UpdateKubernetesCredentialsConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesCredentialsConfigurationAPIService ValidateCreateKubernetesCredentialsConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.KubernetesCredentialsConfigurationAPI.ValidateCreateKubernetesCredentialsConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test KubernetesCredentialsConfigurationAPIService ValidateUpdateKubernetesCredentialsConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.KubernetesCredentialsConfigurationAPI.ValidateUpdateKubernetesCredentialsConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}