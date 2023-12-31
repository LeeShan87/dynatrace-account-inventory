# \RemoteEnvironmentsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRemoteEnvironmentConfig**](RemoteEnvironmentsAPI.md#CreateRemoteEnvironmentConfig) | **Post** /remoteEnvironments | Creates a new remote environment configuration | maturity&#x3D;EARLY_ADOPTER
[**DeleteRemoteEnvironmentConfig**](RemoteEnvironmentsAPI.md#DeleteRemoteEnvironmentConfig) | **Delete** /remoteEnvironments/{id} | Deletes the specified remote environment configuration | maturity&#x3D;EARLY_ADOPTER
[**GetRemoteEnvironmentConfig**](RemoteEnvironmentsAPI.md#GetRemoteEnvironmentConfig) | **Get** /remoteEnvironments/{id} | Gets the properties of the specified remote environment configuration | maturity&#x3D;EARLY_ADOPTER
[**ListRemoteEnvironmentConfigs**](RemoteEnvironmentsAPI.md#ListRemoteEnvironmentConfigs) | **Get** /remoteEnvironments | Lists all remote environment configurations | maturity&#x3D;EARLY_ADOPTER
[**UpdateRemoteEnvironmentConfig**](RemoteEnvironmentsAPI.md#UpdateRemoteEnvironmentConfig) | **Put** /remoteEnvironments/{id} | Updates an existing remote environment configuration or creates a new one | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateRemoteEnvironmentConfig**](RemoteEnvironmentsAPI.md#ValidateCreateRemoteEnvironmentConfig) | **Post** /remoteEnvironments/validator | Validates the payload for the &#x60;POST /remoteEnvironments&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateRemoteEnvironmentConfig**](RemoteEnvironmentsAPI.md#ValidateUpdateRemoteEnvironmentConfig) | **Post** /remoteEnvironments/{id}/validator | Validates the payload for the &#x60;PUT /remoteEnvironments/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateRemoteEnvironmentConfig

> RemoteEnvironmentConfigStub CreateRemoteEnvironmentConfig(ctx).RemoteEnvironmentConfigDto(remoteEnvironmentConfigDto).Execute()

Creates a new remote environment configuration | maturity=EARLY_ADOPTER



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
    remoteEnvironmentConfigDto := *openapiclient.NewRemoteEnvironmentConfigDto("DisplayName_example", "Uri_example") // RemoteEnvironmentConfigDto | The JSON body of the request. Contains parameters of the new remote environment configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteEnvironmentsAPI.CreateRemoteEnvironmentConfig(context.Background()).RemoteEnvironmentConfigDto(remoteEnvironmentConfigDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteEnvironmentsAPI.CreateRemoteEnvironmentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRemoteEnvironmentConfig`: RemoteEnvironmentConfigStub
    fmt.Fprintf(os.Stdout, "Response from `RemoteEnvironmentsAPI.CreateRemoteEnvironmentConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRemoteEnvironmentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteEnvironmentConfigDto** | [**RemoteEnvironmentConfigDto**](RemoteEnvironmentConfigDto.md) | The JSON body of the request. Contains parameters of the new remote environment configuration. | 

### Return type

[**RemoteEnvironmentConfigStub**](RemoteEnvironmentConfigStub.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRemoteEnvironmentConfig

> DeleteRemoteEnvironmentConfig(ctx, id).Execute()

Deletes the specified remote environment configuration | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the configuration to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RemoteEnvironmentsAPI.DeleteRemoteEnvironmentConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteEnvironmentsAPI.DeleteRemoteEnvironmentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the configuration to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteEnvironmentConfigRequest struct via the builder pattern


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


## GetRemoteEnvironmentConfig

> RemoteEnvironmentConfigDto GetRemoteEnvironmentConfig(ctx, id).Execute()

Gets the properties of the specified remote environment configuration | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteEnvironmentsAPI.GetRemoteEnvironmentConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteEnvironmentsAPI.GetRemoteEnvironmentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteEnvironmentConfig`: RemoteEnvironmentConfigDto
    fmt.Fprintf(os.Stdout, "Response from `RemoteEnvironmentsAPI.GetRemoteEnvironmentConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteEnvironmentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteEnvironmentConfigDto**](RemoteEnvironmentConfigDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRemoteEnvironmentConfigs

> RemoteEnvironmentConfigListDto ListRemoteEnvironmentConfigs(ctx).Execute()

Lists all remote environment configurations | maturity=EARLY_ADOPTER

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
    resp, r, err := apiClient.RemoteEnvironmentsAPI.ListRemoteEnvironmentConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteEnvironmentsAPI.ListRemoteEnvironmentConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRemoteEnvironmentConfigs`: RemoteEnvironmentConfigListDto
    fmt.Fprintf(os.Stdout, "Response from `RemoteEnvironmentsAPI.ListRemoteEnvironmentConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRemoteEnvironmentConfigsRequest struct via the builder pattern


### Return type

[**RemoteEnvironmentConfigListDto**](RemoteEnvironmentConfigListDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteEnvironmentConfig

> RemoteEnvironmentConfigStub UpdateRemoteEnvironmentConfig(ctx, id).RemoteEnvironmentConfigDto(remoteEnvironmentConfigDto).Execute()

Updates an existing remote environment configuration or creates a new one | maturity=EARLY_ADOPTER



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
    id := "id_example" // string | The ID of the configuration to be updated.    If you set the ID in the body as well, it must match this ID.
    remoteEnvironmentConfigDto := *openapiclient.NewRemoteEnvironmentConfigDto("DisplayName_example", "Uri_example") // RemoteEnvironmentConfigDto | The JSON body of the request. Contains updated parameters of the remote environment configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteEnvironmentsAPI.UpdateRemoteEnvironmentConfig(context.Background(), id).RemoteEnvironmentConfigDto(remoteEnvironmentConfigDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteEnvironmentsAPI.UpdateRemoteEnvironmentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRemoteEnvironmentConfig`: RemoteEnvironmentConfigStub
    fmt.Fprintf(os.Stdout, "Response from `RemoteEnvironmentsAPI.UpdateRemoteEnvironmentConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the configuration to be updated.    If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteEnvironmentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteEnvironmentConfigDto** | [**RemoteEnvironmentConfigDto**](RemoteEnvironmentConfigDto.md) | The JSON body of the request. Contains updated parameters of the remote environment configuration. | 

### Return type

[**RemoteEnvironmentConfigStub**](RemoteEnvironmentConfigStub.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCreateRemoteEnvironmentConfig

> ValidateCreateRemoteEnvironmentConfig(ctx).RemoteEnvironmentConfigDto(remoteEnvironmentConfigDto).Execute()

Validates the payload for the `POST /remoteEnvironments` request | maturity=EARLY_ADOPTER

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
    remoteEnvironmentConfigDto := *openapiclient.NewRemoteEnvironmentConfigDto("DisplayName_example", "Uri_example") // RemoteEnvironmentConfigDto | The JSON body of the request. Contains the remote environment configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RemoteEnvironmentsAPI.ValidateCreateRemoteEnvironmentConfig(context.Background()).RemoteEnvironmentConfigDto(remoteEnvironmentConfigDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteEnvironmentsAPI.ValidateCreateRemoteEnvironmentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateRemoteEnvironmentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **remoteEnvironmentConfigDto** | [**RemoteEnvironmentConfigDto**](RemoteEnvironmentConfigDto.md) | The JSON body of the request. Contains the remote environment configuration to be validated. | 

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


## ValidateUpdateRemoteEnvironmentConfig

> ValidateUpdateRemoteEnvironmentConfig(ctx, id).RemoteEnvironmentConfigDto(remoteEnvironmentConfigDto).Execute()

Validates the payload for the `PUT /remoteEnvironments/{id}` request | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the remote environment configuration to be validated.
    remoteEnvironmentConfigDto := *openapiclient.NewRemoteEnvironmentConfigDto("DisplayName_example", "Uri_example") // RemoteEnvironmentConfigDto | The JSON body of the request. Contains the remote environment configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RemoteEnvironmentsAPI.ValidateUpdateRemoteEnvironmentConfig(context.Background(), id).RemoteEnvironmentConfigDto(remoteEnvironmentConfigDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteEnvironmentsAPI.ValidateUpdateRemoteEnvironmentConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the remote environment configuration to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateRemoteEnvironmentConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteEnvironmentConfigDto** | [**RemoteEnvironmentConfigDto**](RemoteEnvironmentConfigDto.md) | The JSON body of the request. Contains the remote environment configuration to be validated. | 

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

