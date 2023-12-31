/*
Dynatrace Configuration API

Testing AnomalyDetectionMetricEventsAPIService

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

func Test_configv1_AnomalyDetectionMetricEventsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AnomalyDetectionMetricEventsAPIService CreateMetricEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AnomalyDetectionMetricEventsAPI.CreateMetricEvent(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnomalyDetectionMetricEventsAPIService DeleteMetricEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AnomalyDetectionMetricEventsAPI.DeleteMetricEvent(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnomalyDetectionMetricEventsAPIService GetMetricEventConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AnomalyDetectionMetricEventsAPI.GetMetricEventConfig(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnomalyDetectionMetricEventsAPIService ListMetricEventConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AnomalyDetectionMetricEventsAPI.ListMetricEventConfigs(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnomalyDetectionMetricEventsAPIService UpdateMetricEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AnomalyDetectionMetricEventsAPI.UpdateMetricEvent(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnomalyDetectionMetricEventsAPIService ValidateCreateMetricEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AnomalyDetectionMetricEventsAPI.ValidateCreateMetricEvent(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnomalyDetectionMetricEventsAPIService ValidateUpdateMetricEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.AnomalyDetectionMetricEventsAPI.ValidateUpdateMetricEvent(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
