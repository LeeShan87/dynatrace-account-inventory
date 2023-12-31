/*
Dynatrace Environment API

Testing ClusterTimeAPIService

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

func Test_environmentv1_ClusterTimeAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClusterTimeAPIService GetCurrentClusterTime", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ClusterTimeAPI.GetCurrentClusterTime(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
