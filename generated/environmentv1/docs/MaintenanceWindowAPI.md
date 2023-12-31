# \MaintenanceWindowAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllMaintenanceWindowConfigs**](MaintenanceWindowAPI.md#GetAllMaintenanceWindowConfigs) | **Get** /maintenance | Lists all parameters of all maintenance windows available in your Dynatrace environment.
[**GetMaintenanceWindowConfig**](MaintenanceWindowAPI.md#GetMaintenanceWindowConfig) | **Get** /maintenance/{uid} | Lists all parameters of the specified maintenance window.
[**RemoveMaintenanceWindowConfig**](MaintenanceWindowAPI.md#RemoveMaintenanceWindowConfig) | **Delete** /maintenance/{uid} | Deletes the specified maintenance window
[**StoreMaintenanceWindowConfig**](MaintenanceWindowAPI.md#StoreMaintenanceWindowConfig) | **Post** /maintenance | Creates a new or updates an existing maintenance window



## GetAllMaintenanceWindowConfigs

> []MaintenanceWindow GetAllMaintenanceWindowConfigs(ctx).From(from).To(to).Type_(type_).Execute()

Lists all parameters of all maintenance windows available in your Dynatrace environment.

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
    from := int64(789) // int64 | The start timestamp of the inquiry timeframe, in UTC milliseconds.    If not set, the current time is used. (optional) (default to 0)
    to := int64(789) // int64 | The end timestamp of the inquiry timeframe, in UTC milliseconds.    If not set, all maintenance windows beginning after the `from` timestamp will be returned. (optional) (default to 0)
    type_ := "type__example" // string | The type of the maintenance window to return.    If `Unknown` or not set, all maintenance windows are returned. (optional) (default to "Unknown")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceWindowAPI.GetAllMaintenanceWindowConfigs(context.Background()).From(from).To(to).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowAPI.GetAllMaintenanceWindowConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllMaintenanceWindowConfigs`: []MaintenanceWindow
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowAPI.GetAllMaintenanceWindowConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllMaintenanceWindowConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | The start timestamp of the inquiry timeframe, in UTC milliseconds.    If not set, the current time is used. | [default to 0]
 **to** | **int64** | The end timestamp of the inquiry timeframe, in UTC milliseconds.    If not set, all maintenance windows beginning after the &#x60;from&#x60; timestamp will be returned. | [default to 0]
 **type_** | **string** | The type of the maintenance window to return.    If &#x60;Unknown&#x60; or not set, all maintenance windows are returned. | [default to &quot;Unknown&quot;]

### Return type

[**[]MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMaintenanceWindowConfig

> MaintenanceWindow GetMaintenanceWindowConfig(ctx, uid).Execute()

Lists all parameters of the specified maintenance window.

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
    uid := "uid_example" // string | The ID of the required maintenance window.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MaintenanceWindowAPI.GetMaintenanceWindowConfig(context.Background(), uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowAPI.GetMaintenanceWindowConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMaintenanceWindowConfig`: MaintenanceWindow
    fmt.Fprintf(os.Stdout, "Response from `MaintenanceWindowAPI.GetMaintenanceWindowConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | The ID of the required maintenance window. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMaintenanceWindowConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MaintenanceWindow**](MaintenanceWindow.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMaintenanceWindowConfig

> RemoveMaintenanceWindowConfig(ctx, uid).Execute()

Deletes the specified maintenance window



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
    uid := "uid_example" // string | The ID of the maintenance window to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MaintenanceWindowAPI.RemoveMaintenanceWindowConfig(context.Background(), uid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowAPI.RemoveMaintenanceWindowConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uid** | **string** | The ID of the maintenance window to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMaintenanceWindowConfigRequest struct via the builder pattern


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


## StoreMaintenanceWindowConfig

> StoreMaintenanceWindowConfig(ctx).MaintenanceWindow(maintenanceWindow).Execute()

Creates a new or updates an existing maintenance window

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
    maintenanceWindow := *openapiclient.NewMaintenanceWindow(*openapiclient.NewMaintenanceWindowSchedule("MaintenanceEnd_example", "MaintenanceStart_example", "Type_example"), "Type_example") // MaintenanceWindow |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MaintenanceWindowAPI.StoreMaintenanceWindowConfig(context.Background()).MaintenanceWindow(maintenanceWindow).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MaintenanceWindowAPI.StoreMaintenanceWindowConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreMaintenanceWindowConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maintenanceWindow** | [**MaintenanceWindow**](MaintenanceWindow.md) |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

