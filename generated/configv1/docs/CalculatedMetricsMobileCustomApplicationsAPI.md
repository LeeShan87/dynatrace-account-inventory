# \CalculatedMetricsMobileCustomApplicationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMobileMetric**](CalculatedMetricsMobileCustomApplicationsAPI.md#CreateMobileMetric) | **Post** /calculatedMetrics/mobile | Creates a new calculated metric for a mobile or custom app
[**DeleteMobileMetric**](CalculatedMetricsMobileCustomApplicationsAPI.md#DeleteMobileMetric) | **Delete** /calculatedMetrics/mobile/{metricKey} | Deletes the specified calculated metric for mobile or custom app
[**GetMobileMetric**](CalculatedMetricsMobileCustomApplicationsAPI.md#GetMobileMetric) | **Get** /calculatedMetrics/mobile/{metricKey} | Gets the descriptor of the specified calculated metric for mobile or custom app
[**ListMobileMetrics**](CalculatedMetricsMobileCustomApplicationsAPI.md#ListMobileMetrics) | **Get** /calculatedMetrics/mobile | Lists all calculated metrics for mobile and custom apps configured in your environment
[**UpdateMobileMetric**](CalculatedMetricsMobileCustomApplicationsAPI.md#UpdateMobileMetric) | **Put** /calculatedMetrics/mobile/{metricKey} | Updates the specified calculated metric for a mobile or custom app
[**ValidateCreateMobileMetric**](CalculatedMetricsMobileCustomApplicationsAPI.md#ValidateCreateMobileMetric) | **Post** /calculatedMetrics/mobile/validator | Validates the payload for the &#x60;POST /calculatedMetrics/mobile&#x60; request
[**ValidateUpdateMobileMetric**](CalculatedMetricsMobileCustomApplicationsAPI.md#ValidateUpdateMobileMetric) | **Post** /calculatedMetrics/mobile/{metricKey}/validator | Validates the payload for the &#x60;PUT /calculatedMetrics/mobile/{metricKey}&#x60; request



## CreateMobileMetric

> EntityShortRepresentation CreateMobileMetric(ctx).CalculatedMobileMetric(calculatedMobileMetric).Execute()

Creates a new calculated metric for a mobile or custom app

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
    calculatedMobileMetric := *openapiclient.NewCalculatedMobileMetric("ApplicationIdentifier_example", false, "MetricKey_example", "MetricType_example", "Name_example") // CalculatedMobileMetric | The JSON body of the request. Contains the definition of the new calculated metric for mobile or custom app. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.CreateMobileMetric(context.Background()).CalculatedMobileMetric(calculatedMobileMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsMobileCustomApplicationsAPI.CreateMobileMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMobileMetric`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsMobileCustomApplicationsAPI.CreateMobileMetric`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMobileMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculatedMobileMetric** | [**CalculatedMobileMetric**](CalculatedMobileMetric.md) | The JSON body of the request. Contains the definition of the new calculated metric for mobile or custom app. | 

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


## DeleteMobileMetric

> DeleteMobileMetric(ctx, metricKey).Execute()

Deletes the specified calculated metric for mobile or custom app

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
    r, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.DeleteMobileMetric(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsMobileCustomApplicationsAPI.DeleteMobileMetric``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteMobileMetricRequest struct via the builder pattern


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


## GetMobileMetric

> CalculatedMobileMetric GetMobileMetric(ctx, metricKey).Execute()

Gets the descriptor of the specified calculated metric for mobile or custom app

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
    resp, r, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.GetMobileMetric(context.Background(), metricKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsMobileCustomApplicationsAPI.GetMobileMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileMetric`: CalculatedMobileMetric
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsMobileCustomApplicationsAPI.GetMobileMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the required metric. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CalculatedMobileMetric**](CalculatedMobileMetric.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMobileMetrics

> StubList ListMobileMetrics(ctx).Execute()

Lists all calculated metrics for mobile and custom apps configured in your environment

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
    resp, r, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.ListMobileMetrics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsMobileCustomApplicationsAPI.ListMobileMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMobileMetrics`: StubList
    fmt.Fprintf(os.Stdout, "Response from `CalculatedMetricsMobileCustomApplicationsAPI.ListMobileMetrics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMobileMetricsRequest struct via the builder pattern


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


## UpdateMobileMetric

> UpdateMobileMetric(ctx, metricKey).CalculatedMobileMetricUpdate(calculatedMobileMetricUpdate).Execute()

Updates the specified calculated metric for a mobile or custom app

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
    metricKey := "metricKey_example" // string | The key of the metric to be updated.
    calculatedMobileMetricUpdate := *openapiclient.NewCalculatedMobileMetricUpdate() // CalculatedMobileMetricUpdate | The JSON body of the request. Contains the updated definition of the calculated mobile metric. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.UpdateMobileMetric(context.Background(), metricKey).CalculatedMobileMetricUpdate(calculatedMobileMetricUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsMobileCustomApplicationsAPI.UpdateMobileMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the metric to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMobileMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **calculatedMobileMetricUpdate** | [**CalculatedMobileMetricUpdate**](CalculatedMobileMetricUpdate.md) | The JSON body of the request. Contains the updated definition of the calculated mobile metric. | 

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


## ValidateCreateMobileMetric

> ValidateCreateMobileMetric(ctx).CalculatedMobileMetric(calculatedMobileMetric).Execute()

Validates the payload for the `POST /calculatedMetrics/mobile` request

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
    calculatedMobileMetric := *openapiclient.NewCalculatedMobileMetric("ApplicationIdentifier_example", false, "MetricKey_example", "MetricType_example", "Name_example") // CalculatedMobileMetric | The JSON body of the request. Contains the definition of the metric to be validated.   The key of the metric must be omitted. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.ValidateCreateMobileMetric(context.Background()).CalculatedMobileMetric(calculatedMobileMetric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsMobileCustomApplicationsAPI.ValidateCreateMobileMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateMobileMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **calculatedMobileMetric** | [**CalculatedMobileMetric**](CalculatedMobileMetric.md) | The JSON body of the request. Contains the definition of the metric to be validated.   The key of the metric must be omitted. | 

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


## ValidateUpdateMobileMetric

> ValidateUpdateMobileMetric(ctx, metricKey).CalculatedMobileMetricUpdate(calculatedMobileMetricUpdate).Execute()

Validates the payload for the `PUT /calculatedMetrics/mobile/{metricKey}` request

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
    metricKey := "metricKey_example" // string | The key of the metric to be validated.
    calculatedMobileMetricUpdate := *openapiclient.NewCalculatedMobileMetricUpdate() // CalculatedMobileMetricUpdate | The JSON body of the request. Contains the definition of the metric to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CalculatedMetricsMobileCustomApplicationsAPI.ValidateUpdateMobileMetric(context.Background(), metricKey).CalculatedMobileMetricUpdate(calculatedMobileMetricUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CalculatedMetricsMobileCustomApplicationsAPI.ValidateUpdateMobileMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**metricKey** | **string** | The key of the metric to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateMobileMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **calculatedMobileMetricUpdate** | [**CalculatedMobileMetricUpdate**](CalculatedMobileMetricUpdate.md) | The JSON body of the request. Contains the definition of the metric to be validated. | 

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

