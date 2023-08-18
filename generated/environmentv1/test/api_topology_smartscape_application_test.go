/*
Dynatrace Environment API

Testing TopologySmartscapeApplicationAPIService

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

func Test_environmentv1_TopologySmartscapeApplicationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TopologySmartscapeApplicationAPIService GetApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TopologySmartscapeApplicationAPI.GetApplications(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TopologySmartscapeApplicationAPIService GetBaselineValuesForSingleApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var meIdentifier string

		resp, httpRes, err := apiClient.TopologySmartscapeApplicationAPI.GetBaselineValuesForSingleApplication(context.Background(), meIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TopologySmartscapeApplicationAPIService GetSingleApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var meIdentifier string

		resp, httpRes, err := apiClient.TopologySmartscapeApplicationAPI.GetSingleApplication(context.Background(), meIdentifier).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TopologySmartscapeApplicationAPIService UpdateApplication", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var meIdentifier string

		httpRes, err := apiClient.TopologySmartscapeApplicationAPI.UpdateApplication(context.Background(), meIdentifier).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
