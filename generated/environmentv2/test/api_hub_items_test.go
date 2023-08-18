/*
Dynatrace Environment API

Testing HubItemsAPIService

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

func Test_environmentv2_HubItemsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HubItemsAPIService AddToEnvironment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extensionName string

		resp, httpRes, err := apiClient.HubItemsAPI.AddToEnvironment(context.Background(), extensionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HubItemsAPIService GetArtifact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extension1FQN string
		var version string

		httpRes, err := apiClient.HubItemsAPI.GetArtifact(context.Background(), extension1FQN, version).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HubItemsAPIService GetDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extension1FQN string

		resp, httpRes, err := apiClient.HubItemsAPI.GetDetails(context.Background(), extension1FQN).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HubItemsAPIService GetExtensionDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extensionName string

		resp, httpRes, err := apiClient.HubItemsAPI.GetExtensionDetails(context.Background(), extensionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HubItemsAPIService GetImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var encodedUrl string

		httpRes, err := apiClient.HubItemsAPI.GetImage(context.Background(), encodedUrl).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HubItemsAPIService GetTechnologyDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var slug string

		resp, httpRes, err := apiClient.HubItemsAPI.GetTechnologyDetails(context.Background(), slug).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HubItemsAPIService ListCategories", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HubItemsAPI.ListCategories(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HubItemsAPIService ListItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HubItemsAPI.ListItems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HubItemsAPIService Update", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extensionName string

		resp, httpRes, err := apiClient.HubItemsAPI.Update(context.Background(), extensionName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HubItemsAPIService UpdateReleaseNotes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extensionName string
		var version string

		httpRes, err := apiClient.HubItemsAPI.UpdateReleaseNotes(context.Background(), extensionName, version).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HubItemsAPIService UploadMetadataForExtension", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var extensionName string

		httpRes, err := apiClient.HubItemsAPI.UploadMetadataForExtension(context.Background(), extensionName).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}