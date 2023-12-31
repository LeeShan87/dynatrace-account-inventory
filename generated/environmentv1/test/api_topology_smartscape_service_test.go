/*
Dynatrace Environment API

Testing TopologySmartscapeServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package environmentv1

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_environmentv1_TopologySmartscapeServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TopologySmartscapeServiceAPIService GetBaselineValuesForSingleService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var meIdentifier string

		resp, httpRes, err := apiClient.TopologySmartscapeServiceAPI.GetBaselineValuesForSingleService(context.Background(), meIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TopologySmartscapeServiceAPIService GetServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TopologySmartscapeServiceAPI.GetServices(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TopologySmartscapeServiceAPIService GetSingleService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var meIdentifier string

		resp, httpRes, err := apiClient.TopologySmartscapeServiceAPI.GetSingleService(context.Background(), meIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TopologySmartscapeServiceAPIService UpdateService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var meIdentifier string

		httpRes, err := apiClient.TopologySmartscapeServiceAPI.UpdateService(context.Background(), meIdentifier).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
