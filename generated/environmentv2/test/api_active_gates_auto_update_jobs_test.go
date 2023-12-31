/*
Dynatrace Environment API

Testing ActiveGatesAutoUpdateJobsAPIService

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

func Test_environmentv2_ActiveGatesAutoUpdateJobsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ActiveGatesAutoUpdateJobsAPIService CreateUpdateJobForAg", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var agId string

		resp, httpRes, err := apiClient.ActiveGatesAutoUpdateJobsAPI.CreateUpdateJobForAg(context.Background(), agId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveGatesAutoUpdateJobsAPIService DeleteUpdateJobByJobIdForAg", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var agId string
		var jobId string

		httpRes, err := apiClient.ActiveGatesAutoUpdateJobsAPI.DeleteUpdateJobByJobIdForAg(context.Background(), agId, jobId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveGatesAutoUpdateJobsAPIService GetAllUpdateJobList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ActiveGatesAutoUpdateJobsAPI.GetAllUpdateJobList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveGatesAutoUpdateJobsAPIService GetUpdateJobByJobIdForAg", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var agId string
		var jobId string

		resp, httpRes, err := apiClient.ActiveGatesAutoUpdateJobsAPI.GetUpdateJobByJobIdForAg(context.Background(), agId, jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveGatesAutoUpdateJobsAPIService GetUpdateJobListByAgId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var agId string

		resp, httpRes, err := apiClient.ActiveGatesAutoUpdateJobsAPI.GetUpdateJobListByAgId(context.Background(), agId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveGatesAutoUpdateJobsAPIService ValidateUpdateJobForAg", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var agId string

		httpRes, err := apiClient.ActiveGatesAutoUpdateJobsAPI.ValidateUpdateJobForAg(context.Background(), agId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
