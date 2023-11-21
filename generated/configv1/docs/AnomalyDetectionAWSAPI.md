# \AnomalyDetectionAWSAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAwsAnomalyDetectionConfig**](AnomalyDetectionAWSAPI.md#GetAwsAnomalyDetectionConfig) | **Get** /anomalyDetection/aws | Gets the configuration of anomaly detection for AWS
[**UpdateAwsAnomalyDetectionConfig**](AnomalyDetectionAWSAPI.md#UpdateAwsAnomalyDetectionConfig) | **Put** /anomalyDetection/aws | Updates the configuration of anomaly detection for AWS
[**ValidateAwsAnomalyDetectionConfig**](AnomalyDetectionAWSAPI.md#ValidateAwsAnomalyDetectionConfig) | **Post** /anomalyDetection/aws/validator | Validates the configuration of anomaly detection for AWS for the &#x60;PUT /anomalyDetection/aws&#x60; request



## GetAwsAnomalyDetectionConfig

> AwsAnomalyDetectionConfig GetAwsAnomalyDetectionConfig(ctx).Execute()

Gets the configuration of anomaly detection for AWS

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
    resp, r, err := apiClient.AnomalyDetectionAWSAPI.GetAwsAnomalyDetectionConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionAWSAPI.GetAwsAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAwsAnomalyDetectionConfig`: AwsAnomalyDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionAWSAPI.GetAwsAnomalyDetectionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsAnomalyDetectionConfigRequest struct via the builder pattern


### Return type

[**AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAwsAnomalyDetectionConfig

> UpdateAwsAnomalyDetectionConfig(ctx).AwsAnomalyDetectionConfig(awsAnomalyDetectionConfig).Execute()

Updates the configuration of anomaly detection for AWS

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
    awsAnomalyDetectionConfig := *openapiclient.NewAwsAnomalyDetectionConfig(*openapiclient.NewElbHighConnectionErrorsDetectionConfig(false), *openapiclient.NewLambdaHighErrorRateDetectionConfig(false), *openapiclient.NewRdsHighCpuDetectionConfig(false), *openapiclient.NewRdsHighMemoryDetectionConfig(false), *openapiclient.NewRdsHighWriteReadLatencyDetectionConfig(false), *openapiclient.NewRdsLowStorageDetectionConfig(false), *openapiclient.NewRdsRestartsSequenceDetectionConfig(false)) // AwsAnomalyDetectionConfig | JSON body of the request, containing parameters of the AWS anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionAWSAPI.UpdateAwsAnomalyDetectionConfig(context.Background()).AwsAnomalyDetectionConfig(awsAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionAWSAPI.UpdateAwsAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAwsAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsAnomalyDetectionConfig** | [**AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md) | JSON body of the request, containing parameters of the AWS anomaly detection configuration. | 

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


## ValidateAwsAnomalyDetectionConfig

> ValidateAwsAnomalyDetectionConfig(ctx).AwsAnomalyDetectionConfig(awsAnomalyDetectionConfig).Execute()

Validates the configuration of anomaly detection for AWS for the `PUT /anomalyDetection/aws` request

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
    awsAnomalyDetectionConfig := *openapiclient.NewAwsAnomalyDetectionConfig(*openapiclient.NewElbHighConnectionErrorsDetectionConfig(false), *openapiclient.NewLambdaHighErrorRateDetectionConfig(false), *openapiclient.NewRdsHighCpuDetectionConfig(false), *openapiclient.NewRdsHighMemoryDetectionConfig(false), *openapiclient.NewRdsHighWriteReadLatencyDetectionConfig(false), *openapiclient.NewRdsLowStorageDetectionConfig(false), *openapiclient.NewRdsRestartsSequenceDetectionConfig(false)) // AwsAnomalyDetectionConfig | JSON body of the request, containing parameters of the AWS anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionAWSAPI.ValidateAwsAnomalyDetectionConfig(context.Background()).AwsAnomalyDetectionConfig(awsAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionAWSAPI.ValidateAwsAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAwsAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsAnomalyDetectionConfig** | [**AwsAnomalyDetectionConfig**](AwsAnomalyDetectionConfig.md) | JSON body of the request, containing parameters of the AWS anomaly detection configuration. | 

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

