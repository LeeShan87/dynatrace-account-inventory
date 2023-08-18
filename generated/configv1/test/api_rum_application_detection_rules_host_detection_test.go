/*
Dynatrace Configuration API

Testing RUMApplicationDetectionRulesHostDetectionAPIService

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

func Test_configv1_RUMApplicationDetectionRulesHostDetectionAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RUMApplicationDetectionRulesHostDetectionAPIService GetHostDetectionConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RUMApplicationDetectionRulesHostDetectionAPI.GetHostDetectionConfig(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMApplicationDetectionRulesHostDetectionAPIService UpdateHostDetectionConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RUMApplicationDetectionRulesHostDetectionAPI.UpdateHostDetectionConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RUMApplicationDetectionRulesHostDetectionAPIService ValidateHostDetectionConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.RUMApplicationDetectionRulesHostDetectionAPI.ValidateHostDetectionConfig(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}