/*
Dynatrace Configuration API

Testing AnomalyDetectionVMwareAPIService

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

func Test_configv1_AnomalyDetectionVMwareAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AnomalyDetectionVMwareAPIService GetVMwareAnomalyDetectionConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AnomalyDetectionVMwareAPI.GetVMwareAnomalyDetectionConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnomalyDetectionVMwareAPIService UpdateVMwareAnomalyDetectionConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AnomalyDetectionVMwareAPI.UpdateVMwareAnomalyDetectionConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AnomalyDetectionVMwareAPIService ValidateVMwareAnomalyDetectionConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AnomalyDetectionVMwareAPI.ValidateVMwareAnomalyDetectionConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}