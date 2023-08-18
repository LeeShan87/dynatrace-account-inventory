# \SyntheticOnDemandMonitorExecutionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Execute**](SyntheticOnDemandMonitorExecutionsAPI.md#Execute) | **Post** /synthetic/executions/batch | Triggers on-demand executions for synthetic monitors
[**GetBatch**](SyntheticOnDemandMonitorExecutionsAPI.md#GetBatch) | **Get** /synthetic/executions/batch/{batchId} | Gets summary information and the list of failed executions for the given batch ID
[**GetExecution**](SyntheticOnDemandMonitorExecutionsAPI.md#GetExecution) | **Get** /synthetic/executions/{executionId} | Gets basic results of the specified on-demand execution
[**GetExecutionFullReport**](SyntheticOnDemandMonitorExecutionsAPI.md#GetExecutionFullReport) | **Get** /synthetic/executions/{executionId}/fullReport | Gets detailed results of the specified on-demand execution
[**GetExecutions**](SyntheticOnDemandMonitorExecutionsAPI.md#GetExecutions) | **Get** /synthetic/executions | Gets the list of all on-demand executions of synthetic monitors



## Execute

> SyntheticOnDemandExecutionResult Execute(ctx).SyntheticOnDemandExecutionRequest(syntheticOnDemandExecutionRequest).Execute()

Triggers on-demand executions for synthetic monitors

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
    syntheticOnDemandExecutionRequest := *openapiclient.NewSyntheticOnDemandExecutionRequest() // SyntheticOnDemandExecutionRequest | The JSON body of the request. Contains the parameters of the triggered on-demand execution. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticOnDemandMonitorExecutionsAPI.Execute(context.Background()).SyntheticOnDemandExecutionRequest(syntheticOnDemandExecutionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticOnDemandMonitorExecutionsAPI.Execute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Execute`: SyntheticOnDemandExecutionResult
    fmt.Fprintf(os.Stdout, "Response from `SyntheticOnDemandMonitorExecutionsAPI.Execute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syntheticOnDemandExecutionRequest** | [**SyntheticOnDemandExecutionRequest**](SyntheticOnDemandExecutionRequest.md) | The JSON body of the request. Contains the parameters of the triggered on-demand execution. | 

### Return type

[**SyntheticOnDemandExecutionResult**](SyntheticOnDemandExecutionResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatch

> SyntheticOnDemandBatchStatus GetBatch(ctx, batchId).Execute()

Gets summary information and the list of failed executions for the given batch ID

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
    batchId := int64(789) // int64 | The batch identifier of the executions.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticOnDemandMonitorExecutionsAPI.GetBatch(context.Background(), batchId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticOnDemandMonitorExecutionsAPI.GetBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBatch`: SyntheticOnDemandBatchStatus
    fmt.Fprintf(os.Stdout, "Response from `SyntheticOnDemandMonitorExecutionsAPI.GetBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **int64** | The batch identifier of the executions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticOnDemandBatchStatus**](SyntheticOnDemandBatchStatus.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecution

> SyntheticOnDemandExecution GetExecution(ctx, executionId).Execute()

Gets basic results of the specified on-demand execution

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
    executionId := int64(789) // int64 | The identifier of the on-demand execution.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticOnDemandMonitorExecutionsAPI.GetExecution(context.Background(), executionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticOnDemandMonitorExecutionsAPI.GetExecution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecution`: SyntheticOnDemandExecution
    fmt.Fprintf(os.Stdout, "Response from `SyntheticOnDemandMonitorExecutionsAPI.GetExecution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **int64** | The identifier of the on-demand execution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticOnDemandExecution**](SyntheticOnDemandExecution.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutionFullReport

> SyntheticOnDemandExecution GetExecutionFullReport(ctx, executionId).Execute()

Gets detailed results of the specified on-demand execution

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
    executionId := int64(789) // int64 | The identifier of the on-demand execution.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticOnDemandMonitorExecutionsAPI.GetExecutionFullReport(context.Background(), executionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticOnDemandMonitorExecutionsAPI.GetExecutionFullReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecutionFullReport`: SyntheticOnDemandExecution
    fmt.Fprintf(os.Stdout, "Response from `SyntheticOnDemandMonitorExecutionsAPI.GetExecutionFullReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**executionId** | **int64** | The identifier of the on-demand execution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionFullReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticOnDemandExecution**](SyntheticOnDemandExecution.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExecutions

> SyntheticOnDemandExecutions GetExecutions(ctx).ExecutionStage(executionStage).SchedulingFrom(schedulingFrom).SchedulingTo(schedulingTo).ExecutionFrom(executionFrom).ExecutionTo(executionTo).DataDeliveryFrom(dataDeliveryFrom).DataDeliveryTo(dataDeliveryTo).BatchId(batchId).MonitorId(monitorId).LocationId(locationId).UserId(userId).Source(source).Execute()

Gets the list of all on-demand executions of synthetic monitors

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
    executionStage := "executionStage_example" // string | Filters the resulting set of executions by their stage. (optional)
    schedulingFrom := "schedulingFrom_example" // string | The start of the requested timeframe for scheduling timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the maximum relative timeframe of six hours is used (`now-6h`). (optional)
    schedulingTo := "schedulingTo_example" // string | The end of the requested timeframe for scheduling timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    executionFrom := "executionFrom_example" // string | The start of the requested timeframe for execution timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the maximum relative timeframe of six hours is used (`now-6h`). (optional)
    executionTo := "executionTo_example" // string | The end of the requested timeframe for execution timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    dataDeliveryFrom := "dataDeliveryFrom_example" // string | The start of the requested timeframe for data delivering timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the maximum relative timeframe of six hours is used (`now-6h`). (optional)
    dataDeliveryTo := "dataDeliveryTo_example" // string | The end of the requested timeframe for data delivering timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    batchId := int64(789) // int64 | Filters the resulting set of the executions by batch. Specify the ID of the batch. (optional)
    monitorId := "monitorId_example" // string | Filters the resulting set of the executions by monitor synthetic monitor. Specify the ID of the monitor. (optional)
    locationId := "locationId_example" // string | Filters the resulting set of the executions by Synthetic location. Specify the ID of the location. (optional)
    userId := "userId_example" // string | Filters the resulting set of executions by scheduled user. (optional)
    source := "source_example" // string | Filters the resulting set of the executions by the source of the triggering request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticOnDemandMonitorExecutionsAPI.GetExecutions(context.Background()).ExecutionStage(executionStage).SchedulingFrom(schedulingFrom).SchedulingTo(schedulingTo).ExecutionFrom(executionFrom).ExecutionTo(executionTo).DataDeliveryFrom(dataDeliveryFrom).DataDeliveryTo(dataDeliveryTo).BatchId(batchId).MonitorId(monitorId).LocationId(locationId).UserId(userId).Source(source).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticOnDemandMonitorExecutionsAPI.GetExecutions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExecutions`: SyntheticOnDemandExecutions
    fmt.Fprintf(os.Stdout, "Response from `SyntheticOnDemandMonitorExecutionsAPI.GetExecutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExecutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executionStage** | **string** | Filters the resulting set of executions by their stage. | 
 **schedulingFrom** | **string** | The start of the requested timeframe for scheduling timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the maximum relative timeframe of six hours is used (&#x60;now-6h&#x60;). | 
 **schedulingTo** | **string** | The end of the requested timeframe for scheduling timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **executionFrom** | **string** | The start of the requested timeframe for execution timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the maximum relative timeframe of six hours is used (&#x60;now-6h&#x60;). | 
 **executionTo** | **string** | The end of the requested timeframe for execution timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **dataDeliveryFrom** | **string** | The start of the requested timeframe for data delivering timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the maximum relative timeframe of six hours is used (&#x60;now-6h&#x60;). | 
 **dataDeliveryTo** | **string** | The end of the requested timeframe for data delivering timestamps.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **batchId** | **int64** | Filters the resulting set of the executions by batch. Specify the ID of the batch. | 
 **monitorId** | **string** | Filters the resulting set of the executions by monitor synthetic monitor. Specify the ID of the monitor. | 
 **locationId** | **string** | Filters the resulting set of the executions by Synthetic location. Specify the ID of the location. | 
 **userId** | **string** | Filters the resulting set of executions by scheduled user. | 
 **source** | **string** | Filters the resulting set of the executions by the source of the triggering request. | 

### Return type

[**SyntheticOnDemandExecutions**](SyntheticOnDemandExecutions.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

