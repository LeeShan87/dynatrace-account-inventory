# \ActiveGatesRemoteConfigurationManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemoteIdentityOperationJob**](ActiveGatesRemoteConfigurationManagementAPI.md#CreateRemoteIdentityOperationJob) | **Post** /activeGates/remoteConfigurationManagement | Creates a new remote configuration management job
[**CreateRemoteIdentityOperationPreview**](ActiveGatesRemoteConfigurationManagementAPI.md#CreateRemoteIdentityOperationPreview) | **Post** /activeGates/remoteConfigurationManagement/preview | Creates a preview for remote configuration management job - applicable only to network zone, host group and ActiveGate group
[**GetCurrentRemoteIdentityOperationJob**](ActiveGatesRemoteConfigurationManagementAPI.md#GetCurrentRemoteIdentityOperationJob) | **Get** /activeGates/remoteConfigurationManagement/current | Gets remote configuration management job that is currently running
[**GetRemoteIdentityOperationJob**](ActiveGatesRemoteConfigurationManagementAPI.md#GetRemoteIdentityOperationJob) | **Get** /activeGates/remoteConfigurationManagement/{id} | Gets the specified remote configuration management job
[**GetRemoteIdentityOperations**](ActiveGatesRemoteConfigurationManagementAPI.md#GetRemoteIdentityOperations) | **Get** /activeGates/remoteConfigurationManagement | Lists finished ActiveGate remote configuration management jobs
[**ValidateRemoteIdentityOperation**](ActiveGatesRemoteConfigurationManagementAPI.md#ValidateRemoteIdentityOperation) | **Post** /activeGates/remoteConfigurationManagement/validator | Validates the payload for the &#x60;POST /activeGates/remoteConfigurationManagement&#x60; request.



## CreateRemoteIdentityOperationJob

> RemoteConfigurationManagementJob CreateRemoteIdentityOperationJob(ctx).RemoteConfigurationManagementOperationActiveGateRequest(remoteConfigurationManagementOperationActiveGateRequest).Execute()

Creates a new remote configuration management job

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
    remoteConfigurationManagementOperationActiveGateRequest := *openapiclient.NewRemoteConfigurationManagementOperationActiveGateRequest([]string{"Entities_example"}, []openapiclient.RemoteConfigurationManagementOperation{*openapiclient.NewRemoteConfigurationManagementOperation("networkZone", "set")}) // RemoteConfigurationManagementOperationActiveGateRequest | JSON body of the request, containing remote configuration management job definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveGatesRemoteConfigurationManagementAPI.CreateRemoteIdentityOperationJob(context.Background()).RemoteConfigurationManagementOperationActiveGateRequest(remoteConfigurationManagementOperationActiveGateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesRemoteConfigurationManagementAPI.CreateRemoteIdentityOperationJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemoteIdentityOperationJob`: RemoteConfigurationManagementJob
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesRemoteConfigurationManagementAPI.CreateRemoteIdentityOperationJob`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemoteIdentityOperationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteConfigurationManagementOperationActiveGateRequest** | [**RemoteConfigurationManagementOperationActiveGateRequest**](RemoteConfigurationManagementOperationActiveGateRequest.md) | JSON body of the request, containing remote configuration management job definition. | 

### Return type

[**RemoteConfigurationManagementJob**](RemoteConfigurationManagementJob.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRemoteIdentityOperationPreview

> RemoteConfigurationManagementPreviewList CreateRemoteIdentityOperationPreview(ctx).RemoteConfigurationManagementOperationActiveGateRequest(remoteConfigurationManagementOperationActiveGateRequest).Execute()

Creates a preview for remote configuration management job - applicable only to network zone, host group and ActiveGate group

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
    remoteConfigurationManagementOperationActiveGateRequest := *openapiclient.NewRemoteConfigurationManagementOperationActiveGateRequest([]string{"Entities_example"}, []openapiclient.RemoteConfigurationManagementOperation{*openapiclient.NewRemoteConfigurationManagementOperation("networkZone", "set")}) // RemoteConfigurationManagementOperationActiveGateRequest | JSON body of the request, containing remote configuration management job definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveGatesRemoteConfigurationManagementAPI.CreateRemoteIdentityOperationPreview(context.Background()).RemoteConfigurationManagementOperationActiveGateRequest(remoteConfigurationManagementOperationActiveGateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesRemoteConfigurationManagementAPI.CreateRemoteIdentityOperationPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemoteIdentityOperationPreview`: RemoteConfigurationManagementPreviewList
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesRemoteConfigurationManagementAPI.CreateRemoteIdentityOperationPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemoteIdentityOperationPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteConfigurationManagementOperationActiveGateRequest** | [**RemoteConfigurationManagementOperationActiveGateRequest**](RemoteConfigurationManagementOperationActiveGateRequest.md) | JSON body of the request, containing remote configuration management job definition. | 

### Return type

[**RemoteConfigurationManagementPreviewList**](RemoteConfigurationManagementPreviewList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentRemoteIdentityOperationJob

> RemoteConfigurationManagementJob GetCurrentRemoteIdentityOperationJob(ctx).Execute()

Gets remote configuration management job that is currently running



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveGatesRemoteConfigurationManagementAPI.GetCurrentRemoteIdentityOperationJob(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesRemoteConfigurationManagementAPI.GetCurrentRemoteIdentityOperationJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentRemoteIdentityOperationJob`: RemoteConfigurationManagementJob
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesRemoteConfigurationManagementAPI.GetCurrentRemoteIdentityOperationJob`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentRemoteIdentityOperationJobRequest struct via the builder pattern


### Return type

[**RemoteConfigurationManagementJob**](RemoteConfigurationManagementJob.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteIdentityOperationJob

> RemoteConfigurationManagementJob GetRemoteIdentityOperationJob(ctx, id).Execute()

Gets the specified remote configuration management job

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
    id := "id_example" // string | The ID of the required remote configuration management job.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveGatesRemoteConfigurationManagementAPI.GetRemoteIdentityOperationJob(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesRemoteConfigurationManagementAPI.GetRemoteIdentityOperationJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteIdentityOperationJob`: RemoteConfigurationManagementJob
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesRemoteConfigurationManagementAPI.GetRemoteIdentityOperationJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required remote configuration management job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteIdentityOperationJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteConfigurationManagementJob**](RemoteConfigurationManagementJob.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteIdentityOperations

> RemoteConfigurationManagementJobList GetRemoteIdentityOperations(ctx).From(from).To(to).Execute()

Lists finished ActiveGate remote configuration management jobs



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
    from := "from_example" // string | The start of the requested timeframe for a remote configuration management job.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years    (optional)
    to := "to_example" // string | The end of the requested timeframe for a remote configuration management job.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveGatesRemoteConfigurationManagementAPI.GetRemoteIdentityOperations(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesRemoteConfigurationManagementAPI.GetRemoteIdentityOperations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteIdentityOperations`: RemoteConfigurationManagementJobList
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesRemoteConfigurationManagementAPI.GetRemoteIdentityOperations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteIdentityOperationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The start of the requested timeframe for a remote configuration management job.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years    | 
 **to** | **string** | The end of the requested timeframe for a remote configuration management job.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 

### Return type

[**RemoteConfigurationManagementJobList**](RemoteConfigurationManagementJobList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateRemoteIdentityOperation

> ValidateRemoteIdentityOperation(ctx).RemoteConfigurationManagementOperationActiveGateRequest(remoteConfigurationManagementOperationActiveGateRequest).Execute()

Validates the payload for the `POST /activeGates/remoteConfigurationManagement` request.

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
    remoteConfigurationManagementOperationActiveGateRequest := *openapiclient.NewRemoteConfigurationManagementOperationActiveGateRequest([]string{"Entities_example"}, []openapiclient.RemoteConfigurationManagementOperation{*openapiclient.NewRemoteConfigurationManagementOperation("networkZone", "set")}) // RemoteConfigurationManagementOperationActiveGateRequest | JSON body of the request, containing remote configuration management job definition.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActiveGatesRemoteConfigurationManagementAPI.ValidateRemoteIdentityOperation(context.Background()).RemoteConfigurationManagementOperationActiveGateRequest(remoteConfigurationManagementOperationActiveGateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesRemoteConfigurationManagementAPI.ValidateRemoteIdentityOperation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRemoteIdentityOperationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteConfigurationManagementOperationActiveGateRequest** | [**RemoteConfigurationManagementOperationActiveGateRequest**](RemoteConfigurationManagementOperationActiveGateRequest.md) | JSON body of the request, containing remote configuration management job definition. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

