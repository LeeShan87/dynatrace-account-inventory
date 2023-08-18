# \AnomalyDetectionVMwareAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetVMwareAnomalyDetectionConfig**](AnomalyDetectionVMwareAPI.md#GetVMwareAnomalyDetectionConfig) | **Get** /anomalyDetection/vmware | Gets the configuration of anomaly detection for VMware
[**UpdateVMwareAnomalyDetectionConfig**](AnomalyDetectionVMwareAPI.md#UpdateVMwareAnomalyDetectionConfig) | **Put** /anomalyDetection/vmware | Updates the configuration of anomaly detection for VMware
[**ValidateVMwareAnomalyDetectionConfig**](AnomalyDetectionVMwareAPI.md#ValidateVMwareAnomalyDetectionConfig) | **Post** /anomalyDetection/vmware/validator | Validates the configuration of anomaly detection for VMware for the &#x60;PUT /anomalyDetection/vmware&#x60; request



## GetVMwareAnomalyDetectionConfig

> VMwareAnomalyDetectionConfig GetVMwareAnomalyDetectionConfig(ctx).Execute()

Gets the configuration of anomaly detection for VMware

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
    resp, r, err := apiClient.AnomalyDetectionVMwareAPI.GetVMwareAnomalyDetectionConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionVMwareAPI.GetVMwareAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMwareAnomalyDetectionConfig`: VMwareAnomalyDetectionConfig
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionVMwareAPI.GetVMwareAnomalyDetectionConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMwareAnomalyDetectionConfigRequest struct via the builder pattern


### Return type

[**VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMwareAnomalyDetectionConfig

> UpdateVMwareAnomalyDetectionConfig(ctx).VMwareAnomalyDetectionConfig(vMwareAnomalyDetectionConfig).Execute()

Updates the configuration of anomaly detection for VMware

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
    vMwareAnomalyDetectionConfig := *openapiclient.NewVMwareAnomalyDetectionConfig(*openapiclient.NewDroppedPacketsDetectionConfig(false), *openapiclient.NewEsxiHighCpuSaturationConfig(false), *openapiclient.NewEsxiHighMemoryDetectionConfig(false), *openapiclient.NewLowDatastoreSpaceDetectionConfig(false), *openapiclient.NewOverloadedStorageDetectionConfig(false), *openapiclient.NewSlowPhysicalStorageDetectionConfig(false), *openapiclient.NewUndersizedStorageDetectionConfig(false)) // VMwareAnomalyDetectionConfig | JSON body of the request, containing parameters of the VMware anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionVMwareAPI.UpdateVMwareAnomalyDetectionConfig(context.Background()).VMwareAnomalyDetectionConfig(vMwareAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionVMwareAPI.UpdateVMwareAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMwareAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vMwareAnomalyDetectionConfig** | [**VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md) | JSON body of the request, containing parameters of the VMware anomaly detection configuration. | 

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


## ValidateVMwareAnomalyDetectionConfig

> ValidateVMwareAnomalyDetectionConfig(ctx).VMwareAnomalyDetectionConfig(vMwareAnomalyDetectionConfig).Execute()

Validates the configuration of anomaly detection for VMware for the `PUT /anomalyDetection/vmware` request

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
    vMwareAnomalyDetectionConfig := *openapiclient.NewVMwareAnomalyDetectionConfig(*openapiclient.NewDroppedPacketsDetectionConfig(false), *openapiclient.NewEsxiHighCpuSaturationConfig(false), *openapiclient.NewEsxiHighMemoryDetectionConfig(false), *openapiclient.NewLowDatastoreSpaceDetectionConfig(false), *openapiclient.NewOverloadedStorageDetectionConfig(false), *openapiclient.NewSlowPhysicalStorageDetectionConfig(false), *openapiclient.NewUndersizedStorageDetectionConfig(false)) // VMwareAnomalyDetectionConfig | JSON body of the request, containing parameters of the VMware anomaly detection configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionVMwareAPI.ValidateVMwareAnomalyDetectionConfig(context.Background()).VMwareAnomalyDetectionConfig(vMwareAnomalyDetectionConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionVMwareAPI.ValidateVMwareAnomalyDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateVMwareAnomalyDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vMwareAnomalyDetectionConfig** | [**VMwareAnomalyDetectionConfig**](VMwareAnomalyDetectionConfig.md) | JSON body of the request, containing parameters of the VMware anomaly detection configuration. | 

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

