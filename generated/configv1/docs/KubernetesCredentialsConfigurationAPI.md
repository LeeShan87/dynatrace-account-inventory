# \KubernetesCredentialsConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKubernetesCredentialsConfig**](KubernetesCredentialsConfigurationAPI.md#CreateKubernetesCredentialsConfig) | **Post** /kubernetes/credentials | Creates a new Kubernetes credentials configuration | maturity&#x3D;EARLY_ADOPTER
[**DeleteKubernetesCredentialsConfig**](KubernetesCredentialsConfigurationAPI.md#DeleteKubernetesCredentialsConfig) | **Delete** /kubernetes/credentials/{id} | Deletes the specified Kubernetes credentials configuration | maturity&#x3D;EARLY_ADOPTER
[**GetKubernetesCredentialsConfig**](KubernetesCredentialsConfigurationAPI.md#GetKubernetesCredentialsConfig) | **Get** /kubernetes/credentials/{id} | Gets the configuration of the specified Kubernetes credentials | maturity&#x3D;EARLY_ADOPTER
[**ListKubernetesCredentialsConfigs**](KubernetesCredentialsConfigurationAPI.md#ListKubernetesCredentialsConfigs) | **Get** /kubernetes/credentials | Lists all available Kubernetes credentials configurations | maturity&#x3D;EARLY_ADOPTER
[**UpdateKubernetesCredentialsConfig**](KubernetesCredentialsConfigurationAPI.md#UpdateKubernetesCredentialsConfig) | **Put** /kubernetes/credentials/{id} | Updates an existing Kubernetes credentials configuration | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateKubernetesCredentialsConfig**](KubernetesCredentialsConfigurationAPI.md#ValidateCreateKubernetesCredentialsConfig) | **Post** /kubernetes/credentials/validator | Validates the payload for the &#x60;POST /kubernetes/credentials&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateKubernetesCredentialsConfig**](KubernetesCredentialsConfigurationAPI.md#ValidateUpdateKubernetesCredentialsConfig) | **Post** /kubernetes/credentials/{id}/validator | Validates the payload for the &#x60;PUT /kubernetes/credentials/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## CreateKubernetesCredentialsConfig

> EntityShortRepresentation CreateKubernetesCredentialsConfig(ctx).KubernetesCredentials(kubernetesCredentials).Execute()

Creates a new Kubernetes credentials configuration | maturity=EARLY_ADOPTER



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
    kubernetesCredentials := *openapiclient.NewKubernetesCredentials("EndpointUrl_example", "Label_example") // KubernetesCredentials | The JSON body of the request. Contains parameters of the new Kubernetes credentials configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesCredentialsConfigurationAPI.CreateKubernetesCredentialsConfig(context.Background()).KubernetesCredentials(kubernetesCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesCredentialsConfigurationAPI.CreateKubernetesCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKubernetesCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `KubernetesCredentialsConfigurationAPI.CreateKubernetesCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKubernetesCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesCredentials** | [**KubernetesCredentials**](KubernetesCredentials.md) | The JSON body of the request. Contains parameters of the new Kubernetes credentials configuration. | 

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


## DeleteKubernetesCredentialsConfig

> DeleteKubernetesCredentialsConfig(ctx, id).Execute()

Deletes the specified Kubernetes credentials configuration | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the Kubernetes credentials configuration be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesCredentialsConfigurationAPI.DeleteKubernetesCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesCredentialsConfigurationAPI.DeleteKubernetesCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Kubernetes credentials configuration be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKubernetesCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKubernetesCredentialsConfig

> KubernetesCredentials GetKubernetesCredentialsConfig(ctx, id).Execute()

Gets the configuration of the specified Kubernetes credentials | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required Kubernetes credentials configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesCredentialsConfigurationAPI.GetKubernetesCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesCredentialsConfigurationAPI.GetKubernetesCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKubernetesCredentialsConfig`: KubernetesCredentials
    fmt.Fprintf(os.Stdout, "Response from `KubernetesCredentialsConfigurationAPI.GetKubernetesCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required Kubernetes credentials configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKubernetesCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KubernetesCredentials**](KubernetesCredentials.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKubernetesCredentialsConfigs

> KubernetesConfigStubListDto ListKubernetesCredentialsConfigs(ctx).Execute()

Lists all available Kubernetes credentials configurations | maturity=EARLY_ADOPTER

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
    resp, r, err := apiClient.KubernetesCredentialsConfigurationAPI.ListKubernetesCredentialsConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesCredentialsConfigurationAPI.ListKubernetesCredentialsConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKubernetesCredentialsConfigs`: KubernetesConfigStubListDto
    fmt.Fprintf(os.Stdout, "Response from `KubernetesCredentialsConfigurationAPI.ListKubernetesCredentialsConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListKubernetesCredentialsConfigsRequest struct via the builder pattern


### Return type

[**KubernetesConfigStubListDto**](KubernetesConfigStubListDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateKubernetesCredentialsConfig

> EntityShortRepresentation UpdateKubernetesCredentialsConfig(ctx, id).KubernetesCredentials(kubernetesCredentials).Execute()

Updates an existing Kubernetes credentials configuration | maturity=EARLY_ADOPTER



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
    id := "id_example" // string | The ID of the Kubernetes credentials configuration to be updated.
    kubernetesCredentials := *openapiclient.NewKubernetesCredentials("EndpointUrl_example", "Label_example") // KubernetesCredentials | The JSON body of the request. Contains updated parameters of the Kubernetes credentials configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KubernetesCredentialsConfigurationAPI.UpdateKubernetesCredentialsConfig(context.Background(), id).KubernetesCredentials(kubernetesCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesCredentialsConfigurationAPI.UpdateKubernetesCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateKubernetesCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `KubernetesCredentialsConfigurationAPI.UpdateKubernetesCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Kubernetes credentials configuration to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateKubernetesCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesCredentials** | [**KubernetesCredentials**](KubernetesCredentials.md) | The JSON body of the request. Contains updated parameters of the Kubernetes credentials configuration. | 

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


## ValidateCreateKubernetesCredentialsConfig

> ValidateCreateKubernetesCredentialsConfig(ctx).KubernetesCredentials(kubernetesCredentials).Execute()

Validates the payload for the `POST /kubernetes/credentials` request | maturity=EARLY_ADOPTER

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
    kubernetesCredentials := *openapiclient.NewKubernetesCredentials("EndpointUrl_example", "Label_example") // KubernetesCredentials | The JSON body of the request. Contains the Kubernetes credentials configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesCredentialsConfigurationAPI.ValidateCreateKubernetesCredentialsConfig(context.Background()).KubernetesCredentials(kubernetesCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesCredentialsConfigurationAPI.ValidateCreateKubernetesCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateKubernetesCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kubernetesCredentials** | [**KubernetesCredentials**](KubernetesCredentials.md) | The JSON body of the request. Contains the Kubernetes credentials configuration to be validated. | 

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


## ValidateUpdateKubernetesCredentialsConfig

> ValidateUpdateKubernetesCredentialsConfig(ctx, id).KubernetesCredentials(kubernetesCredentials).Execute()

Validates the payload for the `PUT /kubernetes/credentials/{id}` request | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the Kubernetes credentials configuration to be validated.
    kubernetesCredentials := *openapiclient.NewKubernetesCredentials("EndpointUrl_example", "Label_example") // KubernetesCredentials | The JSON body of the request. Contains the Kubernetes credentials configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.KubernetesCredentialsConfigurationAPI.ValidateUpdateKubernetesCredentialsConfig(context.Background(), id).KubernetesCredentials(kubernetesCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KubernetesCredentialsConfigurationAPI.ValidateUpdateKubernetesCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Kubernetes credentials configuration to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateKubernetesCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kubernetesCredentials** | [**KubernetesCredentials**](KubernetesCredentials.md) | The JSON body of the request. Contains the Kubernetes credentials configuration to be validated. | 

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

