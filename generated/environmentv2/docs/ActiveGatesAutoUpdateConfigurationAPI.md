# \ActiveGatesAutoUpdateConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoUpdateConfigById**](ActiveGatesAutoUpdateConfigurationAPI.md#GetAutoUpdateConfigById) | **Get** /activeGates/{agId}/autoUpdate | Gets the configuration of auto-update for the specified ActiveGate
[**GetGlobalAutoUpdateConfigForTenant**](ActiveGatesAutoUpdateConfigurationAPI.md#GetGlobalAutoUpdateConfigForTenant) | **Get** /activeGates/autoUpdate | Gets the global auto-update configuration of environment ActiveGates.
[**PutAutoUpdateConfigById**](ActiveGatesAutoUpdateConfigurationAPI.md#PutAutoUpdateConfigById) | **Put** /activeGates/{agId}/autoUpdate | Updates the configuration of auto-update for the specified ActiveGate
[**PutGlobalAutoUpdateConfigForTenant**](ActiveGatesAutoUpdateConfigurationAPI.md#PutGlobalAutoUpdateConfigForTenant) | **Put** /activeGates/autoUpdate | Puts the global auto-update configuration of environment ActiveGates.
[**ValidateAutoUpdateConfigById**](ActiveGatesAutoUpdateConfigurationAPI.md#ValidateAutoUpdateConfigById) | **Post** /activeGates/{agId}/autoUpdate/validator | Validates the payload for the &#x60;POST /activeGates/{agId}/autoUpdate&#x60; request.
[**ValidateGlobalAutoUpdateConfigForTenant**](ActiveGatesAutoUpdateConfigurationAPI.md#ValidateGlobalAutoUpdateConfigForTenant) | **Post** /activeGates/autoUpdate/validator | Validates the payload for the &#x60;POST /activeGates/autoUpdate&#x60; request.



## GetAutoUpdateConfigById

> ActiveGateAutoUpdateConfig GetAutoUpdateConfigById(ctx, agId).Execute()

Gets the configuration of auto-update for the specified ActiveGate

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
    agId := "agId_example" // string | The ID of the required ActiveGate.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActiveGatesAutoUpdateConfigurationAPI.GetAutoUpdateConfigById(context.Background(), agId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationAPI.GetAutoUpdateConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoUpdateConfigById`: ActiveGateAutoUpdateConfig
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesAutoUpdateConfigurationAPI.GetAutoUpdateConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoUpdateConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActiveGateAutoUpdateConfig**](ActiveGateAutoUpdateConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalAutoUpdateConfigForTenant

> ActiveGateGlobalAutoUpdateConfig GetGlobalAutoUpdateConfigForTenant(ctx).Execute()

Gets the global auto-update configuration of environment ActiveGates.

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
    resp, r, err := apiClient.ActiveGatesAutoUpdateConfigurationAPI.GetGlobalAutoUpdateConfigForTenant(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationAPI.GetGlobalAutoUpdateConfigForTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalAutoUpdateConfigForTenant`: ActiveGateGlobalAutoUpdateConfig
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesAutoUpdateConfigurationAPI.GetGlobalAutoUpdateConfigForTenant`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalAutoUpdateConfigForTenantRequest struct via the builder pattern


### Return type

[**ActiveGateGlobalAutoUpdateConfig**](ActiveGateGlobalAutoUpdateConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAutoUpdateConfigById

> PutAutoUpdateConfigById(ctx, agId).ActiveGateAutoUpdateConfig(activeGateAutoUpdateConfig).Execute()

Updates the configuration of auto-update for the specified ActiveGate

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
    agId := "agId_example" // string | The ID of the required ActiveGate.
    activeGateAutoUpdateConfig := *openapiclient.NewActiveGateAutoUpdateConfig("INHERITED") // ActiveGateAutoUpdateConfig | JSON body of the request, containing auto update parameters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActiveGatesAutoUpdateConfigurationAPI.PutAutoUpdateConfigById(context.Background(), agId).ActiveGateAutoUpdateConfig(activeGateAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationAPI.PutAutoUpdateConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAutoUpdateConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeGateAutoUpdateConfig** | [**ActiveGateAutoUpdateConfig**](ActiveGateAutoUpdateConfig.md) | JSON body of the request, containing auto update parameters. | 

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


## PutGlobalAutoUpdateConfigForTenant

> PutGlobalAutoUpdateConfigForTenant(ctx).ActiveGateGlobalAutoUpdateConfig(activeGateGlobalAutoUpdateConfig).Execute()

Puts the global auto-update configuration of environment ActiveGates.

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
    activeGateGlobalAutoUpdateConfig := *openapiclient.NewActiveGateGlobalAutoUpdateConfig("GlobalSetting_example") // ActiveGateGlobalAutoUpdateConfig | JSON body of the request, containing global auto update parameters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActiveGatesAutoUpdateConfigurationAPI.PutGlobalAutoUpdateConfigForTenant(context.Background()).ActiveGateGlobalAutoUpdateConfig(activeGateGlobalAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationAPI.PutGlobalAutoUpdateConfigForTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutGlobalAutoUpdateConfigForTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeGateGlobalAutoUpdateConfig** | [**ActiveGateGlobalAutoUpdateConfig**](ActiveGateGlobalAutoUpdateConfig.md) | JSON body of the request, containing global auto update parameters. | 

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


## ValidateAutoUpdateConfigById

> ValidateAutoUpdateConfigById(ctx, agId).ActiveGateAutoUpdateConfig(activeGateAutoUpdateConfig).Execute()

Validates the payload for the `POST /activeGates/{agId}/autoUpdate` request.

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
    agId := "agId_example" // string | The ID of the required ActiveGate.
    activeGateAutoUpdateConfig := *openapiclient.NewActiveGateAutoUpdateConfig("INHERITED") // ActiveGateAutoUpdateConfig | JSON body of the request, containing auto update parameters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActiveGatesAutoUpdateConfigurationAPI.ValidateAutoUpdateConfigById(context.Background(), agId).ActiveGateAutoUpdateConfig(activeGateAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationAPI.ValidateAutoUpdateConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**agId** | **string** | The ID of the required ActiveGate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAutoUpdateConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activeGateAutoUpdateConfig** | [**ActiveGateAutoUpdateConfig**](ActiveGateAutoUpdateConfig.md) | JSON body of the request, containing auto update parameters. | 

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


## ValidateGlobalAutoUpdateConfigForTenant

> ValidateGlobalAutoUpdateConfigForTenant(ctx).ActiveGateGlobalAutoUpdateConfig(activeGateGlobalAutoUpdateConfig).Execute()

Validates the payload for the `POST /activeGates/autoUpdate` request.

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
    activeGateGlobalAutoUpdateConfig := *openapiclient.NewActiveGateGlobalAutoUpdateConfig("GlobalSetting_example") // ActiveGateGlobalAutoUpdateConfig | JSON body of the request, containing global auto update parameters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ActiveGatesAutoUpdateConfigurationAPI.ValidateGlobalAutoUpdateConfigForTenant(context.Background()).ActiveGateGlobalAutoUpdateConfig(activeGateGlobalAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesAutoUpdateConfigurationAPI.ValidateGlobalAutoUpdateConfigForTenant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateGlobalAutoUpdateConfigForTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeGateGlobalAutoUpdateConfig** | [**ActiveGateGlobalAutoUpdateConfig**](ActiveGateGlobalAutoUpdateConfig.md) | JSON body of the request, containing global auto update parameters. | 

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

