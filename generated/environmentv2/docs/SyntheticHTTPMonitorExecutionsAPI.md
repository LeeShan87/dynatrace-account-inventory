# \SyntheticHTTPMonitorExecutionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExecutionResult**](SyntheticHTTPMonitorExecutionsAPI.md#GetExecutionResult) | **Get** /synthetic/execution/{monitorId}/{resultType} | Gets detailed information about the last execution of the specified HTTP monitor



## GetExecutionResult

> MonitorExecutionResults GetExecutionResult(ctx, monitorId, resultType).LocationId(locationId).Execute()

Gets detailed information about the last execution of the specified HTTP monitor

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    monitorId := "monitorId_example" // string | Identifier of the HTTP monitor for which last execution result is returned.
    resultType := "resultType_example" // string | Defines the result type of the last HTTP monitor's execution.
    locationId := "locationId_example" // string | Filters the results to those executed by specified Synthetic location. Specify the ID of the location. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticHTTPMonitorExecutionsAPI.GetExecutionResult(context.Background(), monitorId, resultType).LocationId(locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticHTTPMonitorExecutionsAPI.GetExecutionResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecutionResult`: MonitorExecutionResults
    fmt.Fprintf(os.Stdout, "Response from `SyntheticHTTPMonitorExecutionsAPI.GetExecutionResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**monitorId** | **string** | Identifier of the HTTP monitor for which last execution result is returned. | 
**resultType** | **string** | Defines the result type of the last HTTP monitor&#39;s execution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **locationId** | **string** | Filters the results to those executed by specified Synthetic location. Specify the ID of the location. | 

### Return type

[**MonitorExecutionResults**](MonitorExecutionResults.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

