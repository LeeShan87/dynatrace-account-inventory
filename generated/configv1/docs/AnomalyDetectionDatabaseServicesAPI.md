# \AnomalyDetectionDatabaseServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatabaseServiceAnomalyDetectionConfig**](AnomalyDetectionDatabaseServicesAPI.md#GetDatabaseServiceAnomalyDetectionConfig) | **Get** /anomalyDetection/databaseServices | Gets the configuration of anomaly detection for database services
[**UpdateDatabaseServiceAnomalyDetectionConfig**](AnomalyDetectionDatabaseServicesAPI.md#UpdateDatabaseServiceAnomalyDetectionConfig) | **Put** /anomalyDetection/databaseServices | Updates the configuration of anomaly detection for database services
[**ValidateDatabaseServiceAnomalyDetectionConfig**](AnomalyDetectionDatabaseServicesAPI.md#ValidateDatabaseServiceAnomalyDetectionConfig) | **Post** /anomalyDetection/databaseServices/validator | Validates the payload for the &#x60;PUT /anomalyDetection/databaseServices&#x60; request



## GetDatabaseServiceAnomalyDetectionConfig

> DatabaseAnomalyDetectionConfig GetDatabaseServiceAnomalyDetectionConfig(ctx).Execute()

Gets the configuration of anomaly detection for database services

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
    resp, r, err := apiClient.AnomalyDetectionDatabaseServicesAPI.GetDatabaseServiceAnomalyDetectionConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDatabaseServicesAPI.GetDatabaseServiceAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDatabaseServiceAnomalyDetectionConfig`: DatabaseAnomalyDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionDatabaseServicesAPI.GetDatabaseServiceAnomalyDetectionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabaseServiceAnomalyDetectionConfigRequest struct via the builder pattern


### Return type

[**DatabaseAnomalyDetectionConfig**](DatabaseAnomalyDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDatabaseServiceAnomalyDetectionConfig

> UpdateDatabaseServiceAnomalyDetectionConfig(ctx).DatabaseAnomalyDetectionConfig(databaseAnomalyDetectionConfig).Execute()

Updates the configuration of anomaly detection for database services

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
    databaseAnomalyDetectionConfig := *openapiclient.NewDatabaseAnomalyDetectionConfig(*openapiclient.NewDatabaseConnectionFailureDetectionConfig(false), *openapiclient.NewFailureRateIncreaseDetectionConfig("DetectionMode_example"), *openapiclient.NewResponseTimeDegradationDetectionConfig("DetectionMode_example")) // DatabaseAnomalyDetectionConfig | The JSON body of the request. Contains parameters of the database service anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionDatabaseServicesAPI.UpdateDatabaseServiceAnomalyDetectionConfig(context.Background()).DatabaseAnomalyDetectionConfig(databaseAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDatabaseServicesAPI.UpdateDatabaseServiceAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDatabaseServiceAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseAnomalyDetectionConfig** | [**DatabaseAnomalyDetectionConfig**](DatabaseAnomalyDetectionConfig.md) | The JSON body of the request. Contains parameters of the database service anomaly detection configuration. | 

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


## ValidateDatabaseServiceAnomalyDetectionConfig

> ValidateDatabaseServiceAnomalyDetectionConfig(ctx).DatabaseAnomalyDetectionConfig(databaseAnomalyDetectionConfig).Execute()

Validates the payload for the `PUT /anomalyDetection/databaseServices` request

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
    databaseAnomalyDetectionConfig := *openapiclient.NewDatabaseAnomalyDetectionConfig(*openapiclient.NewDatabaseConnectionFailureDetectionConfig(false), *openapiclient.NewFailureRateIncreaseDetectionConfig("DetectionMode_example"), *openapiclient.NewResponseTimeDegradationDetectionConfig("DetectionMode_example")) // DatabaseAnomalyDetectionConfig | The JSON body of the request. Contains parameters of the database service anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionDatabaseServicesAPI.ValidateDatabaseServiceAnomalyDetectionConfig(context.Background()).DatabaseAnomalyDetectionConfig(databaseAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionDatabaseServicesAPI.ValidateDatabaseServiceAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDatabaseServiceAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **databaseAnomalyDetectionConfig** | [**DatabaseAnomalyDetectionConfig**](DatabaseAnomalyDetectionConfig.md) | The JSON body of the request. Contains parameters of the database service anomaly detection configuration. | 

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

