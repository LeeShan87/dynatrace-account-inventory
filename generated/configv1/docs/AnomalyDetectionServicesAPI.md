# \AnomalyDetectionServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetServiceAnomalyDetectionConfig**](AnomalyDetectionServicesAPI.md#GetServiceAnomalyDetectionConfig) | **Get** /anomalyDetection/services | Gets the service anomaly detection configuration
[**UpdateServiceAnomalyDetectionConfig**](AnomalyDetectionServicesAPI.md#UpdateServiceAnomalyDetectionConfig) | **Put** /anomalyDetection/services | Updates the service anomaly detection configuration
[**ValidateServiceAnomalyDetectionConfig**](AnomalyDetectionServicesAPI.md#ValidateServiceAnomalyDetectionConfig) | **Post** /anomalyDetection/services/validator | Validates the payload for the &#x60;PUT /anomalyDetection/services&#x60; request



## GetServiceAnomalyDetectionConfig

> ServiceAnomalyDetectionConfig GetServiceAnomalyDetectionConfig(ctx).Execute()

Gets the service anomaly detection configuration

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
    resp, r, err := apiClient.AnomalyDetectionServicesAPI.GetServiceAnomalyDetectionConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionServicesAPI.GetServiceAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAnomalyDetectionConfig`: ServiceAnomalyDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionServicesAPI.GetServiceAnomalyDetectionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAnomalyDetectionConfigRequest struct via the builder pattern


### Return type

[**ServiceAnomalyDetectionConfig**](ServiceAnomalyDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServiceAnomalyDetectionConfig

> UpdateServiceAnomalyDetectionConfig(ctx).ServiceAnomalyDetectionConfig(serviceAnomalyDetectionConfig).Execute()

Updates the service anomaly detection configuration

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
    serviceAnomalyDetectionConfig := *openapiclient.NewServiceAnomalyDetectionConfig(*openapiclient.NewFailureRateIncreaseDetectionConfig("DetectionMode_example"), *openapiclient.NewResponseTimeDegradationDetectionConfig("DetectionMode_example")) // ServiceAnomalyDetectionConfig | The JSON body of the request. Contains parameters of the service anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionServicesAPI.UpdateServiceAnomalyDetectionConfig(context.Background()).ServiceAnomalyDetectionConfig(serviceAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionServicesAPI.UpdateServiceAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAnomalyDetectionConfig** | [**ServiceAnomalyDetectionConfig**](ServiceAnomalyDetectionConfig.md) | The JSON body of the request. Contains parameters of the service anomaly detection configuration. | 

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


## ValidateServiceAnomalyDetectionConfig

> ValidateServiceAnomalyDetectionConfig(ctx).ServiceAnomalyDetectionConfig(serviceAnomalyDetectionConfig).Execute()

Validates the payload for the `PUT /anomalyDetection/services` request

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
    serviceAnomalyDetectionConfig := *openapiclient.NewServiceAnomalyDetectionConfig(*openapiclient.NewFailureRateIncreaseDetectionConfig("DetectionMode_example"), *openapiclient.NewResponseTimeDegradationDetectionConfig("DetectionMode_example")) // ServiceAnomalyDetectionConfig | The JSON body of the request. Contains parameters of the service anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionServicesAPI.ValidateServiceAnomalyDetectionConfig(context.Background()).ServiceAnomalyDetectionConfig(serviceAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionServicesAPI.ValidateServiceAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateServiceAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAnomalyDetectionConfig** | [**ServiceAnomalyDetectionConfig**](ServiceAnomalyDetectionConfig.md) | The JSON body of the request. Contains parameters of the service anomaly detection configuration. | 

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

