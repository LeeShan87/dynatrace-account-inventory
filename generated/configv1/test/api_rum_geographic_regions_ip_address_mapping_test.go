/*
Dynatrace Configuration API

Testing RUMGeographicRegionsIPAddressMappingAPIService

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

func Test_configv1_RUMGeographicRegionsIPAddressMappingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RUMGeographicRegionsIPAddressMappingAPIService GetGeolocationRegionsIpAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RUMGeographicRegionsIPAddressMappingAPI.GetGeolocationRegionsIpAddress(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMGeographicRegionsIPAddressMappingAPIService PutGeolocationRegionsIpAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RUMGeographicRegionsIPAddressMappingAPI.PutGeolocationRegionsIpAddress(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMGeographicRegionsIPAddressMappingAPIService ValidateGeolocationRegionsIpAddress", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RUMGeographicRegionsIPAddressMappingAPI.ValidateGeolocationRegionsIpAddress(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}