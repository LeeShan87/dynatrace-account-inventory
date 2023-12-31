/*
Dynatrace Environment API

Testing JavaScriptMappingFilesAPIService

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

func Test_environmentv2_JavaScriptMappingFilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test JavaScriptMappingFilesAPIService DeleteJavaScriptMappingFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var minifiedJsFileUrl string
		var fileType string

		httpRes, err := apiClient.JavaScriptMappingFilesAPI.DeleteJavaScriptMappingFile(context.Background(), minifiedJsFileUrl, fileType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JavaScriptMappingFilesAPIService GetJavaScriptMappingFilesMetadata", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.JavaScriptMappingFilesAPI.GetJavaScriptMappingFilesMetadata(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JavaScriptMappingFilesAPIService UpdateJavaScriptMappingFileMetadata", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var minifiedJsFileUrl string
		var fileType string

		resp, httpRes, err := apiClient.JavaScriptMappingFilesAPI.UpdateJavaScriptMappingFileMetadata(context.Background(), minifiedJsFileUrl, fileType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test JavaScriptMappingFilesAPIService UploadJavaScriptMappingFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var minifiedJsFileUrl string
		var fileType string

		httpRes, err := apiClient.JavaScriptMappingFilesAPI.UploadJavaScriptMappingFile(context.Background(), minifiedJsFileUrl, fileType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
