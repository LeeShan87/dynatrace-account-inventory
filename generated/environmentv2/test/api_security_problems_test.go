/*
Dynatrace Environment API

Testing SecurityProblemsAPIService

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

func Test_environmentv2_SecurityProblemsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SecurityProblemsAPIService BulkMuteRemediationItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SecurityProblemsAPI.BulkMuteRemediationItems(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService BulkMuteSecurityProblems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SecurityProblemsAPI.BulkMuteSecurityProblems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService BulkUnmuteRemediationItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SecurityProblemsAPI.BulkUnmuteRemediationItems(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService BulkUnmuteSecurityProblems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SecurityProblemsAPI.BulkUnmuteSecurityProblems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService GetEventsForSecurityProblem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SecurityProblemsAPI.GetEventsForSecurityProblem(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService GetRemediationItem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var remediationItemId string

		resp, httpRes, err := apiClient.SecurityProblemsAPI.GetRemediationItem(context.Background(), id, remediationItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService GetRemediationItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SecurityProblemsAPI.GetRemediationItems(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService GetRemediationProgressEntities", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var remediationItemId string

		resp, httpRes, err := apiClient.SecurityProblemsAPI.GetRemediationProgressEntities(context.Background(), id, remediationItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService GetSecurityProblem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SecurityProblemsAPI.GetSecurityProblem(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService GetSecurityProblems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SecurityProblemsAPI.GetSecurityProblems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService GetVulnerableFunctions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.SecurityProblemsAPI.GetVulnerableFunctions(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService MuteSecurityProblem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.SecurityProblemsAPI.MuteSecurityProblem(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService SetRemediationItemMuteState", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var remediationItemId string

		httpRes, err := apiClient.SecurityProblemsAPI.SetRemediationItemMuteState(context.Background(), id, remediationItemId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SecurityProblemsAPIService UnmuteSecurityProblem", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.SecurityProblemsAPI.UnmuteSecurityProblem(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
