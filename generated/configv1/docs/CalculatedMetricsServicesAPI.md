# \CalculatedMetricsServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceMetric**](CalculatedMetricsServicesAPI.md#CreateServiceMetric) | **Post** /calculatedMetrics/service | Creates a new calculated service metric
[**DeleteServiceMetric**](CalculatedMetricsServicesAPI.md#DeleteServiceMetric) | **Delete** /calculatedMetrics/service/{metricKey} | Deletes the specified calculated service metric
[**GetServiceMetric**](CalculatedMetricsServicesAPI.md#GetServiceMetric) | **Get** /calculatedMetrics/service/{metricKey} | Gets the descriptor of the specified calculated service metric
[**ListServiceMetrics**](CalculatedMetricsServicesAPI.md#ListServiceMetrics) | **Get** /calculatedMetrics/service | Lists all calculated service metrics configured in your environment
[**UpdateServiceMetric**](CalculatedMetricsServicesAPI.md#UpdateServiceMetric) | **Put** /calculatedMetrics/service/{metricKey} | Updates the specified calculated service metric
[**ValidateCreateServiceMetric**](CalculatedMetricsServicesAPI.md#ValidateCreateServiceMetric) | **Post** /calculatedMetrics/service/validator | Validates the payload for the &#x60;POST /calculatedMetric/service&#x60; request
[**ValidateUpdateServiceMetric**](CalculatedMetricsServicesAPI.md#ValidateUpdateServiceMetric) | **Post** /calculatedMetrics/service/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetric/service/{id}&#x60; request



## CreateServiceMetric

> EntityShortRepresentation CreateServiceMetric(ctx).CalculatedServiceMetric(calculatedServiceMetric).Execute()

Creates a new calculated service metric

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
    calculatedServiceMetric := *openapiclient.NewCalculatedServiceMetric(false, *openapiclient.NewCalculatedMetricDefinition("Metric_example"), "Name_example", "TsmMetricKey_example", "Unit_example") // CalculatedServiceMetric | The JSON body of the request. Contains parameters of the new calculated service metric.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalculatedMetricsServicesAPI.CreateServiceMetric(context.Background()).CalculatedServiceMetric(calculatedServiceMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsServicesAPI.CreateServiceMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceMetric`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsServicesAPI.CreateServiceMetric`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculatedServiceMetric** | [**CalculatedServiceMetric**](CalculatedServiceMetric.md) | The JSON body of the request. Contains parameters of the new calculated service metric. | 

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


## DeleteServiceMetric

> DeleteServiceMetric(ctx, metricKey).Execute()

Deletes the specified calculated service metric

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
    metricKey := "metricKey_example" // string | The key of the calculated service metric to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CalculatedMetricsServicesAPI.DeleteServiceMetric(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsServicesAPI.DeleteServiceMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the calculated service metric to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceMetricRequest struct via the builder pattern


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


## GetServiceMetric

> CalculatedServiceMetric GetServiceMetric(ctx, metricKey).Execute()

Gets the descriptor of the specified calculated service metric

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
    metricKey := "metricKey_example" // string | The key of the required calculated service metric.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalculatedMetricsServicesAPI.GetServiceMetric(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsServicesAPI.GetServiceMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceMetric`: CalculatedServiceMetric
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsServicesAPI.GetServiceMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the required calculated service metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CalculatedServiceMetric**](CalculatedServiceMetric.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceMetrics

> StubList ListServiceMetrics(ctx).Execute()

Lists all calculated service metrics configured in your environment

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
    resp, r, err := apiClient.CalculatedMetricsServicesAPI.ListServiceMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsServicesAPI.ListServiceMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceMetrics`: StubList
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsServicesAPI.ListServiceMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceMetricsRequest struct via the builder pattern


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


## UpdateServiceMetric

> EntityShortRepresentation UpdateServiceMetric(ctx, metricKey).CalculatedServiceMetric(calculatedServiceMetric).Execute()

Updates the specified calculated service metric



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
    metricKey := "metricKey_example" // string | The key of the calculated service metric to be updated.    The key of the calculated service metric in the body of the request must match this key.
    calculatedServiceMetric := *openapiclient.NewCalculatedServiceMetric(false, *openapiclient.NewCalculatedMetricDefinition("Metric_example"), "Name_example", "TsmMetricKey_example", "Unit_example") // CalculatedServiceMetric | The JSON body of the request. Contains updated parameters of the calculated service metric.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalculatedMetricsServicesAPI.UpdateServiceMetric(context.Background(), metricKey).CalculatedServiceMetric(calculatedServiceMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsServicesAPI.UpdateServiceMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceMetric`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsServicesAPI.UpdateServiceMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the calculated service metric to be updated.    The key of the calculated service metric in the body of the request must match this key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **calculatedServiceMetric** | [**CalculatedServiceMetric**](CalculatedServiceMetric.md) | The JSON body of the request. Contains updated parameters of the calculated service metric. | 

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


## ValidateCreateServiceMetric

> ValidateCreateServiceMetric(ctx).CalculatedServiceMetric(calculatedServiceMetric).Execute()

Validates the payload for the `POST /calculatedMetric/service` request

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
    calculatedServiceMetric := *openapiclient.NewCalculatedServiceMetric(false, *openapiclient.NewCalculatedMetricDefinition("Metric_example"), "Name_example", "TsmMetricKey_example", "Unit_example") // CalculatedServiceMetric | The JSON body of the request. Contains the calculated service metric to be validated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CalculatedMetricsServicesAPI.ValidateCreateServiceMetric(context.Background()).CalculatedServiceMetric(calculatedServiceMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsServicesAPI.ValidateCreateServiceMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateServiceMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculatedServiceMetric** | [**CalculatedServiceMetric**](CalculatedServiceMetric.md) | The JSON body of the request. Contains the calculated service metric to be validated. | 

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


## ValidateUpdateServiceMetric

> ValidateUpdateServiceMetric(ctx, metricKey).CalculatedServiceMetric(calculatedServiceMetric).Execute()

Validates the payload for the `PUT /calculatedMetric/service/{id}` request

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
    metricKey := "metricKey_example" // string | The key of the calculated service metric to be validated .   The key of the metric in the body of the request must match this key.
    calculatedServiceMetric := *openapiclient.NewCalculatedServiceMetric(false, *openapiclient.NewCalculatedMetricDefinition("Metric_example"), "Name_example", "TsmMetricKey_example", "Unit_example") // CalculatedServiceMetric | The JSON body of the request. Contains the calculated service metric to be validated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CalculatedMetricsServicesAPI.ValidateUpdateServiceMetric(context.Background(), metricKey).CalculatedServiceMetric(calculatedServiceMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsServicesAPI.ValidateUpdateServiceMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the calculated service metric to be validated .   The key of the metric in the body of the request must match this key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateServiceMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **calculatedServiceMetric** | [**CalculatedServiceMetric**](CalculatedServiceMetric.md) | The JSON body of the request. Contains the calculated service metric to be validated. | 

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

