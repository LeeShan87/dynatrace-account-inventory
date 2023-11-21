# \OneAgentOnAHostAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoUpdateConfig1**](OneAgentOnAHostAPI.md#GetAutoUpdateConfig1) | **Get** /hosts/{id}/autoupdate | Gets the configuration of OneAgent auto-update on the specified host
[**GetHostConfig**](OneAgentOnAHostAPI.md#GetHostConfig) | **Get** /hosts/{id} | Gets the OneAgent configuration on the specified host
[**GetMonitoringConfig**](OneAgentOnAHostAPI.md#GetMonitoringConfig) | **Get** /hosts/{id}/monitoring | Deprecated. Manage host monitoring settings via the environment API v2 settings endpoint with schemaIds &#39;builtin:host.monitoring&#39; and &#39;builtin:host.monitoring.mode&#39; instead.
[**GetTechHostConfigs**](OneAgentOnAHostAPI.md#GetTechHostConfigs) | **Get** /hosts/{id}/technologies | Gets the configuration of monitored technologies on the specified host
[**UpdateAutoUpdateConfig**](OneAgentOnAHostAPI.md#UpdateAutoUpdateConfig) | **Put** /hosts/{id}/autoupdate | Updates the configuration of OneAgent auto-update on the specified host
[**UpdateMonitoringConfig**](OneAgentOnAHostAPI.md#UpdateMonitoringConfig) | **Put** /hosts/{id}/monitoring | Deprecated. Manage host monitoring settings via the environment API v2 settings endpoint with schemaIds &#39;builtin:host.monitoring&#39; and &#39;builtin:host.monitoring.mode&#39; instead.
[**ValidateAutoUpdateConfig1**](OneAgentOnAHostAPI.md#ValidateAutoUpdateConfig1) | **Post** /hosts/{id}/autoupdate/validator | Validates the payload for the &#x60;PUT /hosts/{id}/autoupdate&#x60; request
[**ValidateMonitoringConfig**](OneAgentOnAHostAPI.md#ValidateMonitoringConfig) | **Post** /hosts/{id}/monitoring/validator | Deprecated. Manage host monitoring settings via the environment API v2 settings endpoint with schemaIds &#39;builtin:host.monitoring&#39; and &#39;builtin:host.monitoring.mode&#39; instead.



## GetAutoUpdateConfig1

> HostAutoUpdateConfig GetAutoUpdateConfig1(ctx, id).Execute()

Gets the configuration of OneAgent auto-update on the specified host

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
    id := "id_example" // string | The Dynatrace entity ID of the required host.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OneAgentOnAHostAPI.GetAutoUpdateConfig1(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostAPI.GetAutoUpdateConfig1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoUpdateConfig1`: HostAutoUpdateConfig
    fmt.Fprintf(os.Stdout, "Response from `OneAgentOnAHostAPI.GetAutoUpdateConfig1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoUpdateConfig1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostAutoUpdateConfig**](HostAutoUpdateConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostConfig

> HostConfig GetHostConfig(ctx, id).Execute()

Gets the OneAgent configuration on the specified host

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
    id := "id_example" // string | The Dynatrace entity ID of the required host.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OneAgentOnAHostAPI.GetHostConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostAPI.GetHostConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostConfig`: HostConfig
    fmt.Fprintf(os.Stdout, "Response from `OneAgentOnAHostAPI.GetHostConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HostConfig**](HostConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonitoringConfig

> MonitoringConfig GetMonitoringConfig(ctx, id).Execute()

Deprecated. Manage host monitoring settings via the environment API v2 settings endpoint with schemaIds 'builtin:host.monitoring' and 'builtin:host.monitoring.mode' instead.



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
    id := "id_example" // string | The Dynatrace entity ID of the required host.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OneAgentOnAHostAPI.GetMonitoringConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostAPI.GetMonitoringConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonitoringConfig`: MonitoringConfig
    fmt.Fprintf(os.Stdout, "Response from `OneAgentOnAHostAPI.GetMonitoringConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMonitoringConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonitoringConfig**](MonitoringConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTechHostConfigs

> TechMonitoringList GetTechHostConfigs(ctx, id).Execute()

Gets the configuration of monitored technologies on the specified host

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
    id := "id_example" // string | The Dynatrace entity ID of the required host.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OneAgentOnAHostAPI.GetTechHostConfigs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostAPI.GetTechHostConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTechHostConfigs`: TechMonitoringList
    fmt.Fprintf(os.Stdout, "Response from `OneAgentOnAHostAPI.GetTechHostConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTechHostConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TechMonitoringList**](TechMonitoringList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAutoUpdateConfig

> UpdateAutoUpdateConfig(ctx, id).HostAutoUpdateConfig(hostAutoUpdateConfig).Execute()

Updates the configuration of OneAgent auto-update on the specified host



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
    id := "id_example" // string | The Dynatrace entity ID of the required host.
    hostAutoUpdateConfig := *openapiclient.NewHostAutoUpdateConfig("DISABLED") // HostAutoUpdateConfig | The JSON body of the request. Contains OneAgent auto-update parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OneAgentOnAHostAPI.UpdateAutoUpdateConfig(context.Background(), id).HostAutoUpdateConfig(hostAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostAPI.UpdateAutoUpdateConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoUpdateConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostAutoUpdateConfig** | [**HostAutoUpdateConfig**](HostAutoUpdateConfig.md) | The JSON body of the request. Contains OneAgent auto-update parameters. | 

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


## UpdateMonitoringConfig

> UpdateMonitoringConfig(ctx, id).MonitoringConfig(monitoringConfig).Execute()

Deprecated. Manage host monitoring settings via the environment API v2 settings endpoint with schemaIds 'builtin:host.monitoring' and 'builtin:host.monitoring.mode' instead.



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
    id := "id_example" // string | The Dynatrace entity ID of the required host.
    monitoringConfig := *openapiclient.NewMonitoringConfig(true, "FULL_STACK") // MonitoringConfig | The JSON body of the request. Contains OneAgent monitoring parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OneAgentOnAHostAPI.UpdateMonitoringConfig(context.Background(), id).MonitoringConfig(monitoringConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostAPI.UpdateMonitoringConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitoringConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitoringConfig** | [**MonitoringConfig**](MonitoringConfig.md) | The JSON body of the request. Contains OneAgent monitoring parameters. | 

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


## ValidateAutoUpdateConfig1

> ValidateAutoUpdateConfig1(ctx, id).HostAutoUpdateConfig(hostAutoUpdateConfig).Execute()

Validates the payload for the `PUT /hosts/{id}/autoupdate` request

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
    id := "id_example" // string | The Dynatrace entity ID of the required host.
    hostAutoUpdateConfig := *openapiclient.NewHostAutoUpdateConfig("DISABLED") // HostAutoUpdateConfig | The JSON body of the request. Contains OneAgent auto-update parameters to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OneAgentOnAHostAPI.ValidateAutoUpdateConfig1(context.Background(), id).HostAutoUpdateConfig(hostAutoUpdateConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostAPI.ValidateAutoUpdateConfig1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAutoUpdateConfig1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **hostAutoUpdateConfig** | [**HostAutoUpdateConfig**](HostAutoUpdateConfig.md) | The JSON body of the request. Contains OneAgent auto-update parameters to be validated. | 

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


## ValidateMonitoringConfig

> ValidateMonitoringConfig(ctx, id).MonitoringConfig(monitoringConfig).Execute()

Deprecated. Manage host monitoring settings via the environment API v2 settings endpoint with schemaIds 'builtin:host.monitoring' and 'builtin:host.monitoring.mode' instead.



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
    id := "id_example" // string | The Dynatrace entity ID of the required host.
    monitoringConfig := *openapiclient.NewMonitoringConfig(true, "FULL_STACK") // MonitoringConfig | The JSON body of the request. Contains OneAgent monitoring parameters. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.OneAgentOnAHostAPI.ValidateMonitoringConfig(context.Background(), id).MonitoringConfig(monitoringConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostAPI.ValidateMonitoringConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required host. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateMonitoringConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitoringConfig** | [**MonitoringConfig**](MonitoringConfig.md) | The JSON body of the request. Contains OneAgent monitoring parameters. | 

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

