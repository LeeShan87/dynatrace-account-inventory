# \ExtensionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLocalExtensionConfiguration**](ExtensionsAPI.md#CreateLocalExtensionConfiguration) | **Post** /extensions/{id}/instances | Creates instance of local configuration for given extension | maturity&#x3D;EARLY_ADOPTER
[**DeleteExtension**](ExtensionsAPI.md#DeleteExtension) | **Delete** /extensions/{id} | Deletes the ZIP file of the specified extension | maturity&#x3D;EARLY_ADOPTER
[**DeleteLocalExtensionConfiguration**](ExtensionsAPI.md#DeleteLocalExtensionConfiguration) | **Delete** /extensions/{id}/instances/{configurationId} | Deletes an existing configuration of the extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtension**](ExtensionsAPI.md#GetExtension) | **Get** /extensions/{id} | Lists the properties of the specified extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtensionBinary**](ExtensionsAPI.md#GetExtensionBinary) | **Get** /extensions/{id}/binary | Downloads the ZIP file of the specified extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtensionConfigurations**](ExtensionsAPI.md#GetExtensionConfigurations) | **Get** /extensions/{id}/instances | Returns list of all local configuration instances for given extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtensionGlobalConfiguration**](ExtensionsAPI.md#GetExtensionGlobalConfiguration) | **Get** /extensions/{id}/global | Get the global configuration of the specified OneAgent or JMX extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtensionStates**](ExtensionsAPI.md#GetExtensionStates) | **Get** /extensions/{id}/states | Lists the states of the specified extension | maturity&#x3D;EARLY_ADOPTER
[**GetExtensions**](ExtensionsAPI.md#GetExtensions) | **Get** /extensions | Lists all uploaded extensions | maturity&#x3D;EARLY_ADOPTER
[**GetHostsForTechnology**](ExtensionsAPI.md#GetHostsForTechnology) | **Get** /extensions/{technology}/availableHosts | Lists all available hosts that have specified technology running | maturity&#x3D;EARLY_ADOPTER
[**GetLocalExtensionConfiguration**](ExtensionsAPI.md#GetLocalExtensionConfiguration) | **Get** /extensions/{id}/instances/{configurationId} | Returns instance of local configuration for given extension | maturity&#x3D;EARLY_ADOPTER
[**GetRemoteExtensionModules**](ExtensionsAPI.md#GetRemoteExtensionModules) | **Get** /extensions/activeGateExtensionModules | List available ActiveGate extension modules | maturity&#x3D;EARLY_ADOPTER
[**UpdateGlobalExtensionConfiguration**](ExtensionsAPI.md#UpdateGlobalExtensionConfiguration) | **Put** /extensions/{id}/global | Updates the configuration of the specified OneAgent or JMX extension | maturity&#x3D;EARLY_ADOPTER
[**UpdateLocalExtensionConfiguration**](ExtensionsAPI.md#UpdateLocalExtensionConfiguration) | **Put** /extensions/{id}/instances/{configurationId} | Updates instance of local configuration for given extension | maturity&#x3D;EARLY_ADOPTER
[**UploadExtension**](ExtensionsAPI.md#UploadExtension) | **Post** /extensions | Uploads a ZIP extension file | maturity&#x3D;EARLY_ADOPTER
[**ValidateExtension**](ExtensionsAPI.md#ValidateExtension) | **Post** /extensions/validator | Validates a ZIP extension file for &#x60;POST/extensions&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateLocalExtensionConfiguration**](ExtensionsAPI.md#ValidateLocalExtensionConfiguration) | **Post** /extensions/{id}/instances/validator | Validates the payload for the &#x60;POST /extensions/{id}/instances&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateLocalExtensionConfiguration

> EntityShortRepresentation CreateLocalExtensionConfiguration(ctx, id).ExtensionConfigurationDto(extensionConfigurationDto).Execute()

Creates instance of local configuration for given extension | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the extension
    extensionConfigurationDto := *openapiclient.NewExtensionConfigurationDto(false, false) // ExtensionConfigurationDto | The JSON body of the request. Contains new configuration of the extension. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsAPI.CreateLocalExtensionConfiguration(context.Background(), id).ExtensionConfigurationDto(extensionConfigurationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.CreateLocalExtensionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLocalExtensionConfiguration`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.CreateLocalExtensionConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLocalExtensionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionConfigurationDto** | [**ExtensionConfigurationDto**](ExtensionConfigurationDto.md) | The JSON body of the request. Contains new configuration of the extension. | 

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


## DeleteExtension

> DeleteExtension(ctx, id).Execute()

Deletes the ZIP file of the specified extension | maturity=EARLY_ADOPTER



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
    id := "id_example" // string | The ID of the extension to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExtensionsAPI.DeleteExtension(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.DeleteExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the extension to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExtensionRequest struct via the builder pattern


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


## DeleteLocalExtensionConfiguration

> DeleteLocalExtensionConfiguration(ctx, id, configurationId).Execute()

Deletes an existing configuration of the extension | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the extension where you want to delete the configuration.
    configurationId := "configurationId_example" // string | The ID of the configuration to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExtensionsAPI.DeleteLocalExtensionConfiguration(context.Background(), id, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.DeleteLocalExtensionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the extension where you want to delete the configuration. | 
**configurationId** | **string** | The ID of the configuration to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLocalExtensionConfigurationRequest struct via the builder pattern


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


## GetExtension

> Extension GetExtension(ctx, id).Execute()

Lists the properties of the specified extension | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required extension.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsAPI.GetExtension(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtension`: Extension
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.GetExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required extension. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Extension**](Extension.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionBinary

> GetExtensionBinary(ctx, id).Execute()

Downloads the ZIP file of the specified extension | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the extension you want to download.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExtensionsAPI.GetExtensionBinary(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetExtensionBinary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the extension you want to download. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionBinaryRequest struct via the builder pattern


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


## GetExtensionConfigurations

> ExtensionConfigurationList GetExtensionConfigurations(ctx, id).PageSize(pageSize).NextPageKey(nextPageKey).Execute()

Returns list of all local configuration instances for given extension | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required extension.
    pageSize := int32(56) // int32 | The number of results per result page. Must be between 1 and 500 (optional) (default to 200)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsAPI.GetExtensionConfigurations(context.Background(), id).PageSize(pageSize).NextPageKey(nextPageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetExtensionConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtensionConfigurations`: ExtensionConfigurationList
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.GetExtensionConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required extension. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of results per result page. Must be between 1 and 500 | [default to 200]
 **nextPageKey** | **string** | The cursor for the next page of results. | 

### Return type

[**ExtensionConfigurationList**](ExtensionConfigurationList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionGlobalConfiguration

> GlobalExtensionConfiguration GetExtensionGlobalConfiguration(ctx, id).Execute()

Get the global configuration of the specified OneAgent or JMX extension | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the extension to be updated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsAPI.GetExtensionGlobalConfiguration(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetExtensionGlobalConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtensionGlobalConfiguration`: GlobalExtensionConfiguration
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.GetExtensionGlobalConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the extension to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionGlobalConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalExtensionConfiguration**](GlobalExtensionConfiguration.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionStates

> ExtensionStateList GetExtensionStates(ctx, id).PageSize(pageSize).NextPageKey(nextPageKey).State(state).Execute()

Lists the states of the specified extension | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required extension.
    pageSize := int32(56) // int32 | The number of results per result page. Must be between 1 and 500 (optional) (default to 200)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. (optional)
    state := "state_example" // string | Extension state to filter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsAPI.GetExtensionStates(context.Background(), id).PageSize(pageSize).NextPageKey(nextPageKey).State(state).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetExtensionStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtensionStates`: ExtensionStateList
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.GetExtensionStates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required extension. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of results per result page. Must be between 1 and 500 | [default to 200]
 **nextPageKey** | **string** | The cursor for the next page of results. | 
 **state** | **string** | Extension state to filter. | 

### Return type

[**ExtensionStateList**](ExtensionStateList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensions

> ExtensionListDto GetExtensions(ctx).PageSize(pageSize).NextPageKey(nextPageKey).Execute()

Lists all uploaded extensions | maturity=EARLY_ADOPTER

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
    pageSize := int32(56) // int32 | The number of results per result page. Must be between 1 and 500 (optional) (default to 200)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsAPI.GetExtensions(context.Background()).PageSize(pageSize).NextPageKey(nextPageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtensions`: ExtensionListDto
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.GetExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | The number of results per result page. Must be between 1 and 500 | [default to 200]
 **nextPageKey** | **string** | The cursor for the next page of results. | 

### Return type

[**ExtensionListDto**](ExtensionListDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostsForTechnology

> HostList GetHostsForTechnology(ctx, technology).Tag(tag).ManagementZone(managementZone).HostGroupId(hostGroupId).HostGroupName(hostGroupName).PageSize(pageSize).NextPageKey(nextPageKey).Execute()

Lists all available hosts that have specified technology running | maturity=EARLY_ADOPTER

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
    technology := "technology_example" // string | Name of requested technology
    tag := []string{"Inner_example"} // []string | Filters the resulting set of hosts by the specified tag.    You can specify several tags in the following format: `tag=tag1&tag=tag2`. The host has to match **all** the specified tags. (optional)
    managementZone := int64(789) // int64 | Only return hosts that are part of the specified management zone. (optional)
    hostGroupId := "hostGroupId_example" // string | Filters the resulting set of hosts by the specified host group.    Specify the Dynatrace IDs of the host group you're interested in. (optional)
    hostGroupName := "hostGroupName_example" // string | Filters the resulting set of hosts by the specified host group.    Specify the name of the host group you're interested in. (optional)
    pageSize := int32(56) // int32 | The number of results per result page. Must be between 1 and 500 (optional) (default to 200)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsAPI.GetHostsForTechnology(context.Background(), technology).Tag(tag).ManagementZone(managementZone).HostGroupId(hostGroupId).HostGroupName(hostGroupName).PageSize(pageSize).NextPageKey(nextPageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetHostsForTechnology``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostsForTechnology`: HostList
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.GetHostsForTechnology`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string** | Name of requested technology | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsForTechnologyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tag** | **[]string** | Filters the resulting set of hosts by the specified tag.    You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The host has to match **all** the specified tags. | 
 **managementZone** | **int64** | Only return hosts that are part of the specified management zone. | 
 **hostGroupId** | **string** | Filters the resulting set of hosts by the specified host group.    Specify the Dynatrace IDs of the host group you&#39;re interested in. | 
 **hostGroupName** | **string** | Filters the resulting set of hosts by the specified host group.    Specify the name of the host group you&#39;re interested in. | 
 **pageSize** | **int32** | The number of results per result page. Must be between 1 and 500 | [default to 200]
 **nextPageKey** | **string** | The cursor for the next page of results. | 

### Return type

[**HostList**](HostList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocalExtensionConfiguration

> ExtensionConfigurationDto GetLocalExtensionConfiguration(ctx, id, configurationId).Execute()

Returns instance of local configuration for given extension | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the extension.
    configurationId := "configurationId_example" // string | The ID of the configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsAPI.GetLocalExtensionConfiguration(context.Background(), id, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetLocalExtensionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocalExtensionConfiguration`: ExtensionConfigurationDto
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.GetLocalExtensionConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the extension. | 
**configurationId** | **string** | The ID of the configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocalExtensionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExtensionConfigurationDto**](ExtensionConfigurationDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteExtensionModules

> StubList GetRemoteExtensionModules(ctx).Execute()

List available ActiveGate extension modules | maturity=EARLY_ADOPTER

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
    resp, r, err := apiClient.ExtensionsAPI.GetRemoteExtensionModules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetRemoteExtensionModules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteExtensionModules`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.GetRemoteExtensionModules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteExtensionModulesRequest struct via the builder pattern


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


## UpdateGlobalExtensionConfiguration

> UpdateGlobalExtensionConfiguration(ctx, id).GlobalExtensionConfiguration(globalExtensionConfiguration).Execute()

Updates the configuration of the specified OneAgent or JMX extension | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the extension to be updated.
    globalExtensionConfiguration := *openapiclient.NewGlobalExtensionConfiguration(false) // GlobalExtensionConfiguration | The JSON body of the request. Contains updated configuration of the extension. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExtensionsAPI.UpdateGlobalExtensionConfiguration(context.Background(), id).GlobalExtensionConfiguration(globalExtensionConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.UpdateGlobalExtensionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the extension to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalExtensionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalExtensionConfiguration** | [**GlobalExtensionConfiguration**](GlobalExtensionConfiguration.md) | The JSON body of the request. Contains updated configuration of the extension. | 

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


## UpdateLocalExtensionConfiguration

> UpdateLocalExtensionConfiguration(ctx, id, configurationId).ExtensionConfigurationDto(extensionConfigurationDto).Execute()

Updates instance of local configuration for given extension | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the extension where you want to update the configuration.   If you set the extension ID in the body as well, it must match this ID.
    configurationId := "configurationId_example" // string | The ID of the configuration to be updated.
    extensionConfigurationDto := *openapiclient.NewExtensionConfigurationDto(false, false) // ExtensionConfigurationDto | The JSON body of the request. Contains updated parameters of the extension configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExtensionsAPI.UpdateLocalExtensionConfiguration(context.Background(), id, configurationId).ExtensionConfigurationDto(extensionConfigurationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.UpdateLocalExtensionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the extension where you want to update the configuration.   If you set the extension ID in the body as well, it must match this ID. | 
**configurationId** | **string** | The ID of the configuration to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocalExtensionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extensionConfigurationDto** | [**ExtensionConfigurationDto**](ExtensionConfigurationDto.md) | The JSON body of the request. Contains updated parameters of the extension configuration. | 

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


## UploadExtension

> EntityShortRepresentation UploadExtension(ctx).File(file).OverrideAlerts(overrideAlerts).Execute()

Uploads a ZIP extension file | maturity=EARLY_ADOPTER

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
    file := os.NewFile(1234, "some_file") // *os.File | Extension ZIP file to be uploaded.    File name must match the **name** field in the `plugin.json` file.   For example, for the extension whose **name** is `custom.remote.python.demo`, the name of the extension file must be `custom.remote.python.demo.zip`.
    overrideAlerts := true // bool | Use extension-defined thresholds for alerts (`true`) or user-defined thresholds (`false`).    Extension-defined thresholds are stored in the `plugin.json` file.   If not set, user-defined thresholds are used. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExtensionsAPI.UploadExtension(context.Background()).File(file).OverrideAlerts(overrideAlerts).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.UploadExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadExtension`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.UploadExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | Extension ZIP file to be uploaded.    File name must match the **name** field in the &#x60;plugin.json&#x60; file.   For example, for the extension whose **name** is &#x60;custom.remote.python.demo&#x60;, the name of the extension file must be &#x60;custom.remote.python.demo.zip&#x60;. | 
 **overrideAlerts** | **bool** | Use extension-defined thresholds for alerts (&#x60;true&#x60;) or user-defined thresholds (&#x60;false&#x60;).    Extension-defined thresholds are stored in the &#x60;plugin.json&#x60; file.   If not set, user-defined thresholds are used. | [default to false]

### Return type

[**EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateExtension

> ValidateExtension(ctx).File(file).JsonOnly(jsonOnly).Execute()

Validates a ZIP extension file for `POST/extensions` request | maturity=EARLY_ADOPTER

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
    file := os.NewFile(1234, "some_file") // *os.File | The ZIP extension file to be uploaded.    The file name must match the ID of the extension. Example: `custom.remote.python.demo.zip`
    jsonOnly := true // bool | Validate only the `plugin.json` file (`true`) or the entire extension structure (`false`).    If not set, the entire extension is validated.  (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExtensionsAPI.ValidateExtension(context.Background()).File(file).JsonOnly(jsonOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.ValidateExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** | The ZIP extension file to be uploaded.    The file name must match the ID of the extension. Example: &#x60;custom.remote.python.demo.zip&#x60; | 
 **jsonOnly** | **bool** | Validate only the &#x60;plugin.json&#x60; file (&#x60;true&#x60;) or the entire extension structure (&#x60;false&#x60;).    If not set, the entire extension is validated.  | [default to false]

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateLocalExtensionConfiguration

> ValidateLocalExtensionConfiguration(ctx, id).ExtensionConfigurationDto(extensionConfigurationDto).Execute()

Validates the payload for the `POST /extensions/{id}/instances` request | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the extension.
    extensionConfigurationDto := *openapiclient.NewExtensionConfigurationDto(false, false) // ExtensionConfigurationDto | The JSON body of the request. Contains new configuration of the extension to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ExtensionsAPI.ValidateLocalExtensionConfiguration(context.Background(), id).ExtensionConfigurationDto(extensionConfigurationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.ValidateLocalExtensionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the extension. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateLocalExtensionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionConfigurationDto** | [**ExtensionConfigurationDto**](ExtensionConfigurationDto.md) | The JSON body of the request. Contains new configuration of the extension to be validated. | 

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

