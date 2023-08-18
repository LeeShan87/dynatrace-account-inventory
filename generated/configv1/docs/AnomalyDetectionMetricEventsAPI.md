# \AnomalyDetectionMetricEventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetricEvent**](AnomalyDetectionMetricEventsAPI.md#CreateMetricEvent) | **Post** /anomalyDetection/metricEvents | Creates a new metric event
[**DeleteMetricEvent**](AnomalyDetectionMetricEventsAPI.md#DeleteMetricEvent) | **Delete** /anomalyDetection/metricEvents/{id} | Deletes the specified metric event
[**GetMetricEventConfig**](AnomalyDetectionMetricEventsAPI.md#GetMetricEventConfig) | **Get** /anomalyDetection/metricEvents/{id} | Gets the properties of the specified metric event
[**ListMetricEventConfigs**](AnomalyDetectionMetricEventsAPI.md#ListMetricEventConfigs) | **Get** /anomalyDetection/metricEvents | Lists all configured metric events
[**UpdateMetricEvent**](AnomalyDetectionMetricEventsAPI.md#UpdateMetricEvent) | **Put** /anomalyDetection/metricEvents/{id} | Updates an existing metric event or creates a new one
[**ValidateCreateMetricEvent**](AnomalyDetectionMetricEventsAPI.md#ValidateCreateMetricEvent) | **Post** /anomalyDetection/metricEvents/validator | Validates the payload for the &#x60;POST /anomalyDetection/metricEvents&#x60; request
[**ValidateUpdateMetricEvent**](AnomalyDetectionMetricEventsAPI.md#ValidateUpdateMetricEvent) | **Post** /anomalyDetection/metricEvents/{id}/validator | Validates the payload for the &#x60;PUT /anomalyDetection/metricEvents/{id}&#x60; request



## CreateMetricEvent

> EntityShortRepresentation CreateMetricEvent(ctx).MetricEvent(metricEvent).Execute()

Creates a new metric event



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
    metricEvent := *openapiclient.NewMetricEvent("Description_example", false, *openapiclient.NewMetricEventMonitoringStrategy("Type_example"), "Name_example") // MetricEvent | The JSON body of the request. Contains parameters of the new metric event. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnomalyDetectionMetricEventsAPI.CreateMetricEvent(context.Background()).MetricEvent(metricEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionMetricEventsAPI.CreateMetricEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMetricEvent`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionMetricEventsAPI.CreateMetricEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetricEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricEvent** | [**MetricEvent**](MetricEvent.md) | The JSON body of the request. Contains parameters of the new metric event. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetricEvent

> DeleteMetricEvent(ctx, id).Execute()

Deletes the specified metric event



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
    id := "id_example" // string | The ID of the metric event to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionMetricEventsAPI.DeleteMetricEvent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionMetricEventsAPI.DeleteMetricEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the metric event to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetricEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricEventConfig

> MetricEvent GetMetricEventConfig(ctx, id).Execute()

Gets the properties of the specified metric event

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
    id := "id_example" // string | The ID of the required metric event.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnomalyDetectionMetricEventsAPI.GetMetricEventConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionMetricEventsAPI.GetMetricEventConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricEventConfig`: MetricEvent
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionMetricEventsAPI.GetMetricEventConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required metric event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricEventConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetricEvent**](MetricEvent.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMetricEventConfigs

> StubList ListMetricEventConfigs(ctx).IncludeEntityFilterMetricEvents(includeEntityFilterMetricEvents).Execute()

Lists all configured metric events

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
    includeEntityFilterMetricEvents := true // bool | Flag to include metric events with associated entities to the response.    Metric events with **entity** filters aren't compatible across environments. If set to `false`, metric events with entity filters will be removed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnomalyDetectionMetricEventsAPI.ListMetricEventConfigs(context.Background()).IncludeEntityFilterMetricEvents(includeEntityFilterMetricEvents).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionMetricEventsAPI.ListMetricEventConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMetricEventConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionMetricEventsAPI.ListMetricEventConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMetricEventConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeEntityFilterMetricEvents** | **bool** | Flag to include metric events with associated entities to the response.    Metric events with **entity** filters aren&#39;t compatible across environments. If set to &#x60;false&#x60;, metric events with entity filters will be removed. | [default to false]

### Return type

[**StubList**](StubList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricEvent

> EntityShortRepresentation UpdateMetricEvent(ctx, id).MetricEvent(metricEvent).Execute()

Updates an existing metric event or creates a new one



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
    id := "id_example" // string | The ID of the metric event to be updated.   If you also set the ID in the body, it must match this ID.
    metricEvent := *openapiclient.NewMetricEvent("Description_example", false, *openapiclient.NewMetricEventMonitoringStrategy("Type_example"), "Name_example") // MetricEvent | The JSON body of the request. Contains updated parameters of the metric event. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnomalyDetectionMetricEventsAPI.UpdateMetricEvent(context.Background(), id).MetricEvent(metricEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionMetricEventsAPI.UpdateMetricEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetricEvent`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionMetricEventsAPI.UpdateMetricEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the metric event to be updated.   If you also set the ID in the body, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricEvent** | [**MetricEvent**](MetricEvent.md) | The JSON body of the request. Contains updated parameters of the metric event. | 

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCreateMetricEvent

> ValidateCreateMetricEvent(ctx).MetricEvent(metricEvent).Execute()

Validates the payload for the `POST /anomalyDetection/metricEvents` request

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
    metricEvent := *openapiclient.NewMetricEvent("Description_example", false, *openapiclient.NewMetricEventMonitoringStrategy("Type_example"), "Name_example") // MetricEvent | The JSON body of the request. Contains the metric event configuration to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionMetricEventsAPI.ValidateCreateMetricEvent(context.Background()).MetricEvent(metricEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionMetricEventsAPI.ValidateCreateMetricEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateMetricEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricEvent** | [**MetricEvent**](MetricEvent.md) | The JSON body of the request. Contains the metric event configuration to validate. | 

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


## ValidateUpdateMetricEvent

> ValidateUpdateMetricEvent(ctx, id).MetricEvent(metricEvent).Execute()

Validates the payload for the `PUT /anomalyDetection/metricEvents/{id}` request

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
    id := "id_example" // string | The ID of the metric event to be validated.
    metricEvent := *openapiclient.NewMetricEvent("Description_example", false, *openapiclient.NewMetricEventMonitoringStrategy("Type_example"), "Name_example") // MetricEvent | The JSON body of the request. Contains the metric event configuration to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionMetricEventsAPI.ValidateUpdateMetricEvent(context.Background(), id).MetricEvent(metricEvent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionMetricEventsAPI.ValidateUpdateMetricEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the metric event to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateMetricEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricEvent** | [**MetricEvent**](MetricEvent.md) | The JSON body of the request. Contains the metric event configuration to validate. | 

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

