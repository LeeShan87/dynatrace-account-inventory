/*
Dynatrace Environment API

Testing GeographicRegionsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package environmentv2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_environmentv2_GeographicRegionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GeographicRegionsAPIService GetCities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var countryCode string
		var regionCode string

		resp, httpRes, err := apiClient.GeographicRegionsAPI.GetCities(context.Background(), countryCode, regionCode).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GeographicRegionsAPIService GetCountries", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GeographicRegionsAPI.GetCountries(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GeographicRegionsAPIService GetRegions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GeographicRegionsAPI.GetRegions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GeographicRegionsAPIService GetRegionsAndCities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var countryCode string

		resp, httpRes, err := apiClient.GeographicRegionsAPI.GetRegionsAndCities(context.Background(), countryCode).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GeographicRegionsAPIService GetRegionsOfCountry", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var countryCode string

		resp, httpRes, err := apiClient.GeographicRegionsAPI.GetRegionsOfCountry(context.Background(), countryCode).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
