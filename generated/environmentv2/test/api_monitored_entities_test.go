/*
Dynatrace Environment API

Testing MonitoredEntitiesAPIService

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

func Test_environmentv2_MonitoredEntitiesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MonitoredEntitiesAPIService GetEntities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MonitoredEntitiesAPI.GetEntities(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MonitoredEntitiesAPIService GetEntity", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var entityId string

		resp, httpRes, err := apiClient.MonitoredEntitiesAPI.GetEntity(context.Background(), entityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MonitoredEntitiesAPIService GetEntityType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string

		resp, httpRes, err := apiClient.MonitoredEntitiesAPI.GetEntityType(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MonitoredEntitiesAPIService GetEntityTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MonitoredEntitiesAPI.GetEntityTypes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MonitoredEntitiesAPIService PushCustomDevice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MonitoredEntitiesAPI.PushCustomDevice(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
