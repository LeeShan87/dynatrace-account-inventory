/*
Dynatrace Configuration API

Testing CalculatedMetricsServicesAPIService

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

func Test_configv1_CalculatedMetricsServicesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CalculatedMetricsServicesAPIService CreateServiceMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CalculatedMetricsServicesAPI.CreateServiceMetric(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsServicesAPIService DeleteServiceMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		httpRes, err := apiClient.CalculatedMetricsServicesAPI.DeleteServiceMetric(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsServicesAPIService GetServiceMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		resp, httpRes, err := apiClient.CalculatedMetricsServicesAPI.GetServiceMetric(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsServicesAPIService ListServiceMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CalculatedMetricsServicesAPI.ListServiceMetrics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsServicesAPIService UpdateServiceMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		resp, httpRes, err := apiClient.CalculatedMetricsServicesAPI.UpdateServiceMetric(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsServicesAPIService ValidateCreateServiceMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CalculatedMetricsServicesAPI.ValidateCreateServiceMetric(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsServicesAPIService ValidateUpdateServiceMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		httpRes, err := apiClient.CalculatedMetricsServicesAPI.ValidateUpdateServiceMetric(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
