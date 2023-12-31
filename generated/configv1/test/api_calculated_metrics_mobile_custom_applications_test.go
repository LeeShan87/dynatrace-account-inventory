/*
Dynatrace Configuration API

Testing CalculatedMetricsMobileCustomApplicationsAPIService

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

func Test_configv1_CalculatedMetricsMobileCustomApplicationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CalculatedMetricsMobileCustomApplicationsAPIService CreateMobileMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.CreateMobileMetric(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsMobileCustomApplicationsAPIService DeleteMobileMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		httpRes, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.DeleteMobileMetric(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsMobileCustomApplicationsAPIService GetMobileMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		resp, httpRes, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.GetMobileMetric(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsMobileCustomApplicationsAPIService ListMobileMetrics", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.ListMobileMetrics(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsMobileCustomApplicationsAPIService UpdateMobileMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		httpRes, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.UpdateMobileMetric(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsMobileCustomApplicationsAPIService ValidateCreateMobileMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.ValidateCreateMobileMetric(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsMobileCustomApplicationsAPIService ValidateUpdateMobileMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		httpRes, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.ValidateUpdateMobileMetric(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
