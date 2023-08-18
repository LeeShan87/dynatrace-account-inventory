# \CalculatedMetricsWebApplicationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRumMetric**](CalculatedMetricsWebApplicationsAPI.md#CreateRumMetric) | **Post** /calculatedMetrics/rum | Creates a new calculated web application metric
[**DeleteRumMetric**](CalculatedMetricsWebApplicationsAPI.md#DeleteRumMetric) | **Delete** /calculatedMetrics/rum/{metricKey} | Deletes the specified calculated web application metric
[**GetRumMetric**](CalculatedMetricsWebApplicationsAPI.md#GetRumMetric) | **Get** /calculatedMetrics/rum/{metricKey} | Gets the descriptor of the specified calculated web application metric
[**ListRumMetrics**](CalculatedMetricsWebApplicationsAPI.md#ListRumMetrics) | **Get** /calculatedMetrics/rum | Lists all calculated web application metrics
[**UpdateRumMetric**](CalculatedMetricsWebApplicationsAPI.md#UpdateRumMetric) | **Put** /calculatedMetrics/rum/{metricKey} | Updates the specified calculated web application metric
[**ValidateCreateRumMetric**](CalculatedMetricsWebApplicationsAPI.md#ValidateCreateRumMetric) | **Post** /calculatedMetrics/rum/validator | Validates the payload for the &#x60;POST /calculatedMetrics/rum&#x60; request
[**ValidateUpdateRumMetric**](CalculatedMetricsWebApplicationsAPI.md#ValidateUpdateRumMetric) | **Post** /calculatedMetrics/rum/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetrics/rum/{metricKey}&#x60; request



## CreateRumMetric

> EntityShortRepresentation CreateRumMetric(ctx).WebApplicationMetric(webApplicationMetric).Execute()

Creates a new calculated web application metric

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
    webApplicationMetric := *openapiclient.NewWebApplicationMetric("ApplicationIdentifier_example", false, *openapiclient.NewWebApplicationMetricDefinition("Metric_example"), "MetricKey_example", "Name_example") // WebApplicationMetric | The JSON body of the request. Contains the descriptor of the new calculated web application metric.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalculatedMetricsWebApplicationsAPI.CreateRumMetric(context.Background()).WebApplicationMetric(webApplicationMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsAPI.CreateRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRumMetric`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsWebApplicationsAPI.CreateRumMetric`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationMetric** | [**WebApplicationMetric**](WebApplicationMetric.md) | The JSON body of the request. Contains the descriptor of the new calculated web application metric. | 

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


## DeleteRumMetric

> DeleteRumMetric(ctx, metricKey).Execute()

Deletes the specified calculated web application metric

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
    metricKey := "metricKey_example" // string | The key of the metric to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CalculatedMetricsWebApplicationsAPI.DeleteRumMetric(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsAPI.DeleteRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the metric to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRumMetricRequest struct via the builder pattern


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


## GetRumMetric

> WebApplicationMetric GetRumMetric(ctx, metricKey).Execute()

Gets the descriptor of the specified calculated web application metric

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
    metricKey := "metricKey_example" // string | The key of the required metric.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalculatedMetricsWebApplicationsAPI.GetRumMetric(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsAPI.GetRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRumMetric`: WebApplicationMetric
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsWebApplicationsAPI.GetRumMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the required metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebApplicationMetric**](WebApplicationMetric.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRumMetrics

> StubList ListRumMetrics(ctx).Execute()

Lists all calculated web application metrics

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
    resp, r, err := apiClient.CalculatedMetricsWebApplicationsAPI.ListRumMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsAPI.ListRumMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRumMetrics`: StubList
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsWebApplicationsAPI.ListRumMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRumMetricsRequest struct via the builder pattern


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


## UpdateRumMetric

> UpdateRumMetric(ctx, metricKey).WebApplicationMetricUpdate(webApplicationMetricUpdate).Execute()

Updates the specified calculated web application metric

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
    metricKey := "metricKey_example" // string | The key of the calculated web application metric to be updated.
    webApplicationMetricUpdate := *openapiclient.NewWebApplicationMetricUpdate() // WebApplicationMetricUpdate | The JSON body of the request. Contains the updated parameters of the calculated web application metric.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CalculatedMetricsWebApplicationsAPI.UpdateRumMetric(context.Background(), metricKey).WebApplicationMetricUpdate(webApplicationMetricUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsAPI.UpdateRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the calculated web application metric to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webApplicationMetricUpdate** | [**WebApplicationMetricUpdate**](WebApplicationMetricUpdate.md) | The JSON body of the request. Contains the updated parameters of the calculated web application metric. | 

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


## ValidateCreateRumMetric

> ValidateCreateRumMetric(ctx).WebApplicationMetric(webApplicationMetric).Execute()

Validates the payload for the `POST /calculatedMetrics/rum` request



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
    webApplicationMetric := *openapiclient.NewWebApplicationMetric("ApplicationIdentifier_example", false, *openapiclient.NewWebApplicationMetricDefinition("Metric_example"), "MetricKey_example", "Name_example") // WebApplicationMetric | The JSON body of the request. Contains the descriptor of the metric to be validated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CalculatedMetricsWebApplicationsAPI.ValidateCreateRumMetric(context.Background()).WebApplicationMetric(webApplicationMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsAPI.ValidateCreateRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationMetric** | [**WebApplicationMetric**](WebApplicationMetric.md) | The JSON body of the request. Contains the descriptor of the metric to be validated. | 

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


## ValidateUpdateRumMetric

> ValidateUpdateRumMetric(ctx, metricKey).WebApplicationMetricUpdate(webApplicationMetricUpdate).Execute()

Validates the payload for the `PUT /calculatedMetrics/rum/{metricKey}` request

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
    metricKey := "metricKey_example" // string | The key of the calculated web application metric to be validated.
    webApplicationMetricUpdate := *openapiclient.NewWebApplicationMetricUpdate() // WebApplicationMetricUpdate | The JSON body of the request. Contains the descriptor of the metric to be validated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CalculatedMetricsWebApplicationsAPI.ValidateUpdateRumMetric(context.Background(), metricKey).WebApplicationMetricUpdate(webApplicationMetricUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsWebApplicationsAPI.ValidateUpdateRumMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the calculated web application metric to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateRumMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webApplicationMetricUpdate** | [**WebApplicationMetricUpdate**](WebApplicationMetricUpdate.md) | The JSON body of the request. Contains the descriptor of the metric to be validated. | 

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

