/*
Dynatrace Environment API

Testing MetricsUnitsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package environmentv2

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_environmentv2_MetricsUnitsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MetricsUnitsAPIService AllUnits", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MetricsUnitsAPI.AllUnits(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsUnitsAPIService Convert", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var unitId string

		resp, httpRes, err := apiClient.MetricsUnitsAPI.Convert(context.Background(), unitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MetricsUnitsAPIService Unit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var unitId string

		resp, httpRes, err := apiClient.MetricsUnitsAPI.Unit(context.Background(), unitId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}