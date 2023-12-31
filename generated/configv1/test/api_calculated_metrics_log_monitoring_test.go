/*
Dynatrace Configuration API

Testing CalculatedMetricsLogMonitoringAPIService

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

func Test_configv1_CalculatedMetricsLogMonitoringAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CalculatedMetricsLogMonitoringAPIService DeleteLogMetricConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		httpRes, err := apiClient.CalculatedMetricsLogMonitoringAPI.DeleteLogMetricConfig(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsLogMonitoringAPIService GetLogMetricConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		resp, httpRes, err := apiClient.CalculatedMetricsLogMonitoringAPI.GetLogMetricConfig(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsLogMonitoringAPIService ListLogMetricConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CalculatedMetricsLogMonitoringAPI.ListLogMetricConfigs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsLogMonitoringAPIService UpdateOrCreateLogMetricConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		resp, httpRes, err := apiClient.CalculatedMetricsLogMonitoringAPI.UpdateOrCreateLogMetricConfig(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CalculatedMetricsLogMonitoringAPIService ValidateMetric", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var metricKey string

		httpRes, err := apiClient.CalculatedMetricsLogMonitoringAPI.ValidateMetric(context.Background(), metricKey).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
