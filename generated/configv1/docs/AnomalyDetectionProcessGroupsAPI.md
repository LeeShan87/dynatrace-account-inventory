# \AnomalyDetectionProcessGroupsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLocalAvailabilityConfig**](AnomalyDetectionProcessGroupsAPI.md#DeleteLocalAvailabilityConfig) | **Delete** /anomalyDetection/processGroups/{id} | Switches off anomaly detection for the specified process group | maturity&#x3D;EARLY_ADOPTER
[**GetLocalAvailabilityConfig**](AnomalyDetectionProcessGroupsAPI.md#GetLocalAvailabilityConfig) | **Get** /anomalyDetection/processGroups/{id} | Gets the configuration of anomaly detection for the specified process group | maturity&#x3D;EARLY_ADOPTER
[**SetLocalAvailabilityConfig**](AnomalyDetectionProcessGroupsAPI.md#SetLocalAvailabilityConfig) | **Put** /anomalyDetection/processGroups/{id} | Updates the configuration of anomaly detection for the specified process group | maturity&#x3D;EARLY_ADOPTER
[**ValidateLocalPgAvailabilityAlertingEvent**](AnomalyDetectionProcessGroupsAPI.md#ValidateLocalPgAvailabilityAlertingEvent) | **Post** /anomalyDetection/processGroups/{id}/validator | Validates the payload for the &#x60;PUT /anomalyDetection/processGroups/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## DeleteLocalAvailabilityConfig

> DeleteLocalAvailabilityConfig(ctx, id).Execute()

Switches off anomaly detection for the specified process group | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The Dynatrace entity ID of the required process group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionProcessGroupsAPI.DeleteLocalAvailabilityConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionProcessGroupsAPI.DeleteLocalAvailabilityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required process group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalAvailabilityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalAvailabilityConfig

> AnomalyDetectionPG GetLocalAvailabilityConfig(ctx, id).Execute()

Gets the configuration of anomaly detection for the specified process group | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The Dynatrace entity ID of the required process group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnomalyDetectionProcessGroupsAPI.GetLocalAvailabilityConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionProcessGroupsAPI.GetLocalAvailabilityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocalAvailabilityConfig`: AnomalyDetectionPG
    fmt.Fprintf(os.Stdout, "Response from `AnomalyDetectionProcessGroupsAPI.GetLocalAvailabilityConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required process group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalAvailabilityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AnomalyDetectionPG**](AnomalyDetectionPG.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLocalAvailabilityConfig

> SetLocalAvailabilityConfig(ctx, id).AnomalyDetectionPG(anomalyDetectionPG).Execute()

Updates the configuration of anomaly detection for the specified process group | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The Dynatrace entity ID of the required process group.
    anomalyDetectionPG := *openapiclient.NewAnomalyDetectionPG() // AnomalyDetectionPG | The JSON body of the request. Contains parameters of anomaly detection for a process group. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionProcessGroupsAPI.SetLocalAvailabilityConfig(context.Background(), id).AnomalyDetectionPG(anomalyDetectionPG).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionProcessGroupsAPI.SetLocalAvailabilityConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required process group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLocalAvailabilityConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **anomalyDetectionPG** | [**AnomalyDetectionPG**](AnomalyDetectionPG.md) | The JSON body of the request. Contains parameters of anomaly detection for a process group. | 

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


## ValidateLocalPgAvailabilityAlertingEvent

> ValidateLocalPgAvailabilityAlertingEvent(ctx, id).AnomalyDetectionPG(anomalyDetectionPG).Execute()

Validates the payload for the `PUT /anomalyDetection/processGroups/{id}` request | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The Dynatrace entity ID of the required process group.
    anomalyDetectionPG := *openapiclient.NewAnomalyDetectionPG() // AnomalyDetectionPG | The JSON body of the request. Contains anomaly detection configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnomalyDetectionProcessGroupsAPI.ValidateLocalPgAvailabilityAlertingEvent(context.Background(), id).AnomalyDetectionPG(anomalyDetectionPG).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnomalyDetectionProcessGroupsAPI.ValidateLocalPgAvailabilityAlertingEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required process group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateLocalPgAvailabilityAlertingEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **anomalyDetectionPG** | [**AnomalyDetectionPG**](AnomalyDetectionPG.md) | The JSON body of the request. Contains anomaly detection configuration to be validated. | 

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

