# \AnomalyDetectionApplicationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApplicationAnomalyDetectionConfig**](AnomalyDetectionApplicationsAPI.md#GetApplicationAnomalyDetectionConfig) | **Get** /anomalyDetection/applications | Gets the configuration of anomaly detection for applications
[**UpdateApplicationAnomalyDetectionConfig**](AnomalyDetectionApplicationsAPI.md#UpdateApplicationAnomalyDetectionConfig) | **Put** /anomalyDetection/applications | Updates the configuration of anomaly detection for applications
[**ValidateApplicationAnomalyDetectionConfig**](AnomalyDetectionApplicationsAPI.md#ValidateApplicationAnomalyDetectionConfig) | **Post** /anomalyDetection/applications/validator | Validates the configuration of anomaly detection for applications for the &#x60;PUT /anomalyDetection/applications&#x60; request



## GetApplicationAnomalyDetectionConfig

> ApplicationAnomalyDetectionConfig GetApplicationAnomalyDetectionConfig(ctx).Execute()

Gets the configuration of anomaly detection for applications

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
    resp, r, err := apiClient.AnomalyDetectionApplicationsAPI.GetApplicationAnomalyDetectionConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionApplicationsAPI.GetApplicationAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationAnomalyDetectionConfig`: ApplicationAnomalyDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionApplicationsAPI.GetApplicationAnomalyDetectionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationAnomalyDetectionConfigRequest struct via the builder pattern


### Return type

[**ApplicationAnomalyDetectionConfig**](ApplicationAnomalyDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationAnomalyDetectionConfig

> UpdateApplicationAnomalyDetectionConfig(ctx).ApplicationAnomalyDetectionConfig(applicationAnomalyDetectionConfig).Execute()

Updates the configuration of anomaly detection for applications

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
    applicationAnomalyDetectionConfig := *openapiclient.NewApplicationAnomalyDetectionConfig(*openapiclient.NewFailureRateIncreaseDetectionConfig("DetectionMode_example"), *openapiclient.NewResponseTimeDegradationDetectionConfig("DetectionMode_example"), *openapiclient.NewTrafficDropDetectionConfig(false), *openapiclient.NewTrafficSpikeDetectionConfig(false)) // ApplicationAnomalyDetectionConfig | The JSON body of the request, containing parameters of the application anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionApplicationsAPI.UpdateApplicationAnomalyDetectionConfig(context.Background()).ApplicationAnomalyDetectionConfig(applicationAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionApplicationsAPI.UpdateApplicationAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationAnomalyDetectionConfig** | [**ApplicationAnomalyDetectionConfig**](ApplicationAnomalyDetectionConfig.md) | The JSON body of the request, containing parameters of the application anomaly detection configuration. | 

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


## ValidateApplicationAnomalyDetectionConfig

> ValidateApplicationAnomalyDetectionConfig(ctx).ApplicationAnomalyDetectionConfig(applicationAnomalyDetectionConfig).Execute()

Validates the configuration of anomaly detection for applications for the `PUT /anomalyDetection/applications` request

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
    applicationAnomalyDetectionConfig := *openapiclient.NewApplicationAnomalyDetectionConfig(*openapiclient.NewFailureRateIncreaseDetectionConfig("DetectionMode_example"), *openapiclient.NewResponseTimeDegradationDetectionConfig("DetectionMode_example"), *openapiclient.NewTrafficDropDetectionConfig(false), *openapiclient.NewTrafficSpikeDetectionConfig(false)) // ApplicationAnomalyDetectionConfig | The JSON body of the request, containing parameters of the application anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionApplicationsAPI.ValidateApplicationAnomalyDetectionConfig(context.Background()).ApplicationAnomalyDetectionConfig(applicationAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionApplicationsAPI.ValidateApplicationAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateApplicationAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationAnomalyDetectionConfig** | [**ApplicationAnomalyDetectionConfig**](ApplicationAnomalyDetectionConfig.md) | The JSON body of the request, containing parameters of the application anomaly detection configuration. | 

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

