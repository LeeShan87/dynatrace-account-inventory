# \SyntheticThirdPartyAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PushEvents**](SyntheticThirdPartyAPI.md#PushEvents) | **Post** /synthetic/ext/events | Pushes third-party synthetic events to Dynatrace
[**PushStateModification**](SyntheticThirdPartyAPI.md#PushStateModification) | **Post** /synthetic/ext/stateModifications | Modifies the operation state of all third-party monitors
[**TestResults**](SyntheticThirdPartyAPI.md#TestResults) | **Post** /synthetic/ext/tests | Pushes third-party synthetic monitors, locations, and results to Dynatrace



## PushEvents

> PushEvents(ctx).Model3rdPartySyntheticEvents(model3rdPartySyntheticEvents).Execute()

Pushes third-party synthetic events to Dynatrace

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
    model3rdPartySyntheticEvents := *openapiclient.NewModel3rdPartySyntheticEvents("SyntheticEngineName_example") // Model3rdPartySyntheticEvents | The JSON body of the request. Contains third-party synthetic events.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SyntheticThirdPartyAPI.PushEvents(context.Background()).Model3rdPartySyntheticEvents(model3rdPartySyntheticEvents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticThirdPartyAPI.PushEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model3rdPartySyntheticEvents** | [**Model3rdPartySyntheticEvents**](Model3rdPartySyntheticEvents.md) | The JSON body of the request. Contains third-party synthetic events. | 

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


## PushStateModification

> PushStateModification(ctx).StateModification(stateModification).Execute()

Modifies the operation state of all third-party monitors

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
    stateModification := *openapiclient.NewStateModification("State_example") // StateModification | The JSON body of the request. Contains new operational status of third-party synthetic monitors.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SyntheticThirdPartyAPI.PushStateModification(context.Background()).StateModification(stateModification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticThirdPartyAPI.PushStateModification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPushStateModificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stateModification** | [**StateModification**](StateModification.md) | The JSON body of the request. Contains new operational status of third-party synthetic monitors. | 

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


## TestResults

> TestResults(ctx).Model3rdPartySyntheticTests(model3rdPartySyntheticTests).Execute()

Pushes third-party synthetic monitors, locations, and results to Dynatrace

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
    model3rdPartySyntheticTests := *openapiclient.NewModel3rdPartySyntheticTests([]openapiclient.Model3rdPartySyntheticLocation{*openapiclient.NewModel3rdPartySyntheticLocation("Id_example", "Name_example")}, int64(123), "SyntheticEngineName_example", []openapiclient.Model3rdPartySyntheticMonitor{*openapiclient.NewModel3rdPartySyntheticMonitor("Id_example", []openapiclient.SyntheticTestLocation{*openapiclient.NewSyntheticTestLocation("Id_example")}, int32(123), "Title_example")}) // Model3rdPartySyntheticTests | The JSON body of the request. Contains third-party synthetic monitors, locations, and results.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SyntheticThirdPartyAPI.TestResults(context.Background()).Model3rdPartySyntheticTests(model3rdPartySyntheticTests).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticThirdPartyAPI.TestResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **model3rdPartySyntheticTests** | [**Model3rdPartySyntheticTests**](Model3rdPartySyntheticTests.md) | The JSON body of the request. Contains third-party synthetic monitors, locations, and results. | 

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

