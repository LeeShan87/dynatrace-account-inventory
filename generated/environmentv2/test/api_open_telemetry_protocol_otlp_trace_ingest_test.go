/*
Dynatrace Environment API

Testing OpenTelemetryProtocolOTLPTraceIngestAPIService

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

func Test_environmentv2_OpenTelemetryProtocolOTLPTraceIngestAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OpenTelemetryProtocolOTLPTraceIngestAPIService IngestTrace", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.OpenTelemetryProtocolOTLPTraceIngestAPI.IngestTrace(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}