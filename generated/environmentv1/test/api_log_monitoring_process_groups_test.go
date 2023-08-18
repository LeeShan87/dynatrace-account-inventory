/*
Dynatrace Environment API

Testing LogMonitoringProcessGroupsAPIService

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

func Test_environmentv1_LogMonitoringProcessGroupsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LogMonitoringProcessGroupsAPIService ProcessGroupLogJobDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pgId string
		var jobId string

		resp, httpRes, err := apiClient.LogMonitoringProcessGroupsAPI.ProcessGroupLogJobDelete(context.Background(), pgId, jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogMonitoringProcessGroupsAPIService ProcessGroupLogJobRecords", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pgId string
		var jobId string

		resp, httpRes, err := apiClient.LogMonitoringProcessGroupsAPI.ProcessGroupLogJobRecords(context.Background(), pgId, jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogMonitoringProcessGroupsAPIService ProcessGroupLogJobRecordsFiltered", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pgId string
		var jobId string

		resp, httpRes, err := apiClient.LogMonitoringProcessGroupsAPI.ProcessGroupLogJobRecordsFiltered(context.Background(), pgId, jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogMonitoringProcessGroupsAPIService ProcessGroupLogJobRecordsTop", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pgId string
		var jobId string

		resp, httpRes, err := apiClient.LogMonitoringProcessGroupsAPI.ProcessGroupLogJobRecordsTop(context.Background(), pgId, jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogMonitoringProcessGroupsAPIService ProcessGroupLogJobStart", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pgId string
		var logPath string

		resp, httpRes, err := apiClient.LogMonitoringProcessGroupsAPI.ProcessGroupLogJobStart(context.Background(), pgId, logPath).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogMonitoringProcessGroupsAPIService ProcessGroupLogJobStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pgId string
		var jobId string

		resp, httpRes, err := apiClient.LogMonitoringProcessGroupsAPI.ProcessGroupLogJobStatus(context.Background(), pgId, jobId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LogMonitoringProcessGroupsAPIService ProcessGroupLogList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pgId string

		resp, httpRes, err := apiClient.LogMonitoringProcessGroupsAPI.ProcessGroupLogList(context.Background(), pgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
