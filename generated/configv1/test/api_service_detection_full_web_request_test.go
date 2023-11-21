/*
Dynatrace Configuration API

Testing ServiceDetectionFullWebRequestAPIService

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

func Test_configv1_ServiceDetectionFullWebRequestAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceDetectionFullWebRequestAPIService CreateRequestDetectionRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServiceDetectionFullWebRequestAPI.CreateRequestDetectionRule(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDetectionFullWebRequestAPIService DeleteRequestDetectionRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ServiceDetectionFullWebRequestAPI.DeleteRequestDetectionRule(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDetectionFullWebRequestAPIService GetRequestDetectionRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ServiceDetectionFullWebRequestAPI.GetRequestDetectionRule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDetectionFullWebRequestAPIService ListRequestDetectionRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServiceDetectionFullWebRequestAPI.ListRequestDetectionRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDetectionFullWebRequestAPIService OrderRequestDetectionRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ServiceDetectionFullWebRequestAPI.OrderRequestDetectionRules(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDetectionFullWebRequestAPIService UpdateRequestDetectionRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ServiceDetectionFullWebRequestAPI.UpdateRequestDetectionRule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDetectionFullWebRequestAPIService ValidateCreateRequestDetectionRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ServiceDetectionFullWebRequestAPI.ValidateCreateRequestDetectionRule(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceDetectionFullWebRequestAPIService ValidateUpdateRequestDetectionRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ServiceDetectionFullWebRequestAPI.ValidateUpdateRequestDetectionRule(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
