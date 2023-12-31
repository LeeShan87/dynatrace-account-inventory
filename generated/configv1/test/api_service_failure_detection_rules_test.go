/*
Dynatrace Configuration API

Testing ServiceFailureDetectionRulesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package configv1

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_configv1_ServiceFailureDetectionRulesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceFailureDetectionRulesAPIService ChangeRuleId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ServiceFailureDetectionRulesAPI.ChangeRuleId(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceFailureDetectionRulesAPIService CreateOrUpdateRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ServiceFailureDetectionRulesAPI.CreateOrUpdateRule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceFailureDetectionRulesAPIService CreateRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServiceFailureDetectionRulesAPI.CreateRule(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceFailureDetectionRulesAPIService DeleteRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ServiceFailureDetectionRulesAPI.DeleteRule(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceFailureDetectionRulesAPIService GetAllRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServiceFailureDetectionRulesAPI.GetAllRules(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceFailureDetectionRulesAPIService GetRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.ServiceFailureDetectionRulesAPI.GetRule(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceFailureDetectionRulesAPIService ReorderRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ServiceFailureDetectionRulesAPI.ReorderRules(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceFailureDetectionRulesAPIService ValidateCreateRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ServiceFailureDetectionRulesAPI.ValidateCreateRule(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceFailureDetectionRulesAPIService ValidateUpdateRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		httpRes, err := apiClient.ServiceFailureDetectionRulesAPI.ValidateUpdateRule(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
