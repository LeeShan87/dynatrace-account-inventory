/*
Dynatrace Configuration API

Testing ServiceCustomServicesAPIService

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

func Test_configv1_ServiceCustomServicesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceCustomServicesAPIService CreateCustomService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var technology string

		resp, httpRes, err := apiClient.ServiceCustomServicesAPI.CreateCustomService(context.Background(), technology).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceCustomServicesAPIService DeleteCustomService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var technology string
		var id string

		httpRes, err := apiClient.ServiceCustomServicesAPI.DeleteCustomService(context.Background(), technology, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceCustomServicesAPIService GetCustomService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var technology string
		var id string

		resp, httpRes, err := apiClient.ServiceCustomServicesAPI.GetCustomService(context.Background(), technology, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceCustomServicesAPIService ListCustomServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var technology string

		resp, httpRes, err := apiClient.ServiceCustomServicesAPI.ListCustomServices(context.Background(), technology).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceCustomServicesAPIService OrderCustomServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var technology string

		httpRes, err := apiClient.ServiceCustomServicesAPI.OrderCustomServices(context.Background(), technology).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceCustomServicesAPIService UpdateCustomService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var technology string
		var id string

		resp, httpRes, err := apiClient.ServiceCustomServicesAPI.UpdateCustomService(context.Background(), technology, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceCustomServicesAPIService ValidateCreateCustomService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var technology string

		httpRes, err := apiClient.ServiceCustomServicesAPI.ValidateCreateCustomService(context.Background(), technology).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceCustomServicesAPIService ValidateUpdateCustomService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var technology string
		var id string

		httpRes, err := apiClient.ServiceCustomServicesAPI.ValidateUpdateCustomService(context.Background(), technology, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
