# \AzureCredentialsConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAzureCredentialsConfig**](AzureCredentialsConfigurationAPI.md#CreateAzureCredentialsConfig) | **Post** /azure/credentials | Creates a new Azure credentials configuration
[**DeleteAzureCredentialsConfig**](AzureCredentialsConfigurationAPI.md#DeleteAzureCredentialsConfig) | **Delete** /azure/credentials/{id} | Deletes the specified Azure credentials configuration
[**GetAzureCredentialsConfig**](AzureCredentialsConfigurationAPI.md#GetAzureCredentialsConfig) | **Get** /azure/credentials/{id} | Gets the configuration of the specified Azure credentials
[**GetAzureServicesConfig**](AzureCredentialsConfigurationAPI.md#GetAzureServicesConfig) | **Get** /azure/credentials/{id}/services | Gets the monitored services configuration of the specified Azure credentials
[**GetSupportedServices1**](AzureCredentialsConfigurationAPI.md#GetSupportedServices1) | **Get** /azure/supportedServices | Gets the list of Azure supported services
[**ListAzureCredentialsConfigs**](AzureCredentialsConfigurationAPI.md#ListAzureCredentialsConfigs) | **Get** /azure/credentials | Lists all available Azure credentials configurations
[**UpdateAzureCredentialsConfig**](AzureCredentialsConfigurationAPI.md#UpdateAzureCredentialsConfig) | **Put** /azure/credentials/{id} | Updates an existing Azure credentials configuration
[**UpdateAzureServicesConfig**](AzureCredentialsConfigurationAPI.md#UpdateAzureServicesConfig) | **Put** /azure/credentials/{id}/services | Replace an existing monitored services configuration of the specified Azure credentials
[**ValidateAzureCredentialsConfig**](AzureCredentialsConfigurationAPI.md#ValidateAzureCredentialsConfig) | **Post** /azure/credentials/validator | Validates the payload for the &#x60;POST /azure/credentials&#x60; request
[**ValidateAzureServicesConfig**](AzureCredentialsConfigurationAPI.md#ValidateAzureServicesConfig) | **Post** /azure/credentials/{id}/services/validator | Validates the payload for the &#x60;PUT /azure/credentials/{id}/services&#x60; request
[**ValidateConfigurationUpdate**](AzureCredentialsConfigurationAPI.md#ValidateConfigurationUpdate) | **Post** /azure/credentials/{id}/validator | Validates the payload for the &#x60;PUT /azure/credentials/{id}&#x60; request



## CreateAzureCredentialsConfig

> EntityShortRepresentation CreateAzureCredentialsConfig(ctx).AzureCredentialsCreation(azureCredentialsCreation).Execute()

Creates a new Azure credentials configuration



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
    azureCredentialsCreation := *openapiclient.NewAzureCredentialsCreation("AppId_example", false, "DirectoryId_example", "Key_example", "Label_example", false) // AzureCredentialsCreation | The JSON body of the request. Contains parameters of the new Azure credentials configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCredentialsConfigurationAPI.CreateAzureCredentialsConfig(context.Background()).AzureCredentialsCreation(azureCredentialsCreation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.CreateAzureCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAzureCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AzureCredentialsConfigurationAPI.CreateAzureCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAzureCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureCredentialsCreation** | [**AzureCredentialsCreation**](AzureCredentialsCreation.md) | The JSON body of the request. Contains parameters of the new Azure credentials configuration. | 

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


## DeleteAzureCredentialsConfig

> DeleteAzureCredentialsConfig(ctx, id).Execute()

Deletes the specified Azure credentials configuration

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
    id := "id_example" // string | The ID of the Azure credentials configuration to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AzureCredentialsConfigurationAPI.DeleteAzureCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.DeleteAzureCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Azure credentials configuration to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAzureCredentialsConfigRequest struct via the builder pattern


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


## GetAzureCredentialsConfig

> AzureCredentials GetAzureCredentialsConfig(ctx, id).Execute()

Gets the configuration of the specified Azure credentials

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
    id := "id_example" // string | The ID of the required Azure credentials configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCredentialsConfigurationAPI.GetAzureCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.GetAzureCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAzureCredentialsConfig`: AzureCredentials
    fmt.Fprintf(os.Stdout, "Response from `AzureCredentialsConfigurationAPI.GetAzureCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required Azure credentials configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AzureCredentials**](AzureCredentials.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAzureServicesConfig

> AzureMonitoredServicesDto GetAzureServicesConfig(ctx, id).Execute()

Gets the monitored services configuration of the specified Azure credentials

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
    id := "id_example" // string | The ID of the required Azure credentials configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCredentialsConfigurationAPI.GetAzureServicesConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.GetAzureServicesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAzureServicesConfig`: AzureMonitoredServicesDto
    fmt.Fprintf(os.Stdout, "Response from `AzureCredentialsConfigurationAPI.GetAzureServicesConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required Azure credentials configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAzureServicesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AzureMonitoredServicesDto**](AzureMonitoredServicesDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedServices1

> CloudSupportedServicesList GetSupportedServices1(ctx).Execute()

Gets the list of Azure supported services



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
    resp, r, err := apiClient.AzureCredentialsConfigurationAPI.GetSupportedServices1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.GetSupportedServices1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedServices1`: CloudSupportedServicesList
    fmt.Fprintf(os.Stdout, "Response from `AzureCredentialsConfigurationAPI.GetSupportedServices1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedServices1Request struct via the builder pattern


### Return type

[**CloudSupportedServicesList**](CloudSupportedServicesList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAzureCredentialsConfigs

> StubList ListAzureCredentialsConfigs(ctx).Execute()

Lists all available Azure credentials configurations

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
    resp, r, err := apiClient.AzureCredentialsConfigurationAPI.ListAzureCredentialsConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.ListAzureCredentialsConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAzureCredentialsConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `AzureCredentialsConfigurationAPI.ListAzureCredentialsConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAzureCredentialsConfigsRequest struct via the builder pattern


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


## UpdateAzureCredentialsConfig

> EntityShortRepresentation UpdateAzureCredentialsConfig(ctx, id).AzureCredentials(azureCredentials).Execute()

Updates an existing Azure credentials configuration



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
    id := "id_example" // string | The ID of the Azure credentials configuration to be updated.
    azureCredentials := *openapiclient.NewAzureCredentials(false, "Label_example", false) // AzureCredentials | The JSON body of the request. Contains updated parameters of the Azure credentials configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AzureCredentialsConfigurationAPI.UpdateAzureCredentialsConfig(context.Background(), id).AzureCredentials(azureCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.UpdateAzureCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAzureCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AzureCredentialsConfigurationAPI.UpdateAzureCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Azure credentials configuration to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureCredentials** | [**AzureCredentials**](AzureCredentials.md) | The JSON body of the request. Contains updated parameters of the Azure credentials configuration. | 

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


## UpdateAzureServicesConfig

> UpdateAzureServicesConfig(ctx, id).AzureMonitoredServicesDto(azureMonitoredServicesDto).Execute()

Replace an existing monitored services configuration of the specified Azure credentials

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
    id := "id_example" // string | The ID of the Azure credentials to be updated with new monitored services configuration.
    azureMonitoredServicesDto := *openapiclient.NewAzureMonitoredServicesDto([]openapiclient.AzureSupportingService{*openapiclient.NewAzureSupportingService("Name_example")}) // AzureMonitoredServicesDto | The JSON body of the request. Contains updated monitored services configuration for Azure credentials. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AzureCredentialsConfigurationAPI.UpdateAzureServicesConfig(context.Background(), id).AzureMonitoredServicesDto(azureMonitoredServicesDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.UpdateAzureServicesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Azure credentials to be updated with new monitored services configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAzureServicesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureMonitoredServicesDto** | [**AzureMonitoredServicesDto**](AzureMonitoredServicesDto.md) | The JSON body of the request. Contains updated monitored services configuration for Azure credentials. | 

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


## ValidateAzureCredentialsConfig

> ValidateAzureCredentialsConfig(ctx).AzureCredentials(azureCredentials).Execute()

Validates the payload for the `POST /azure/credentials` request

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
    azureCredentials := *openapiclient.NewAzureCredentials(false, "Label_example", false) // AzureCredentials | The JSON body of the request. Contains the Azure credentials configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AzureCredentialsConfigurationAPI.ValidateAzureCredentialsConfig(context.Background()).AzureCredentials(azureCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.ValidateAzureCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAzureCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **azureCredentials** | [**AzureCredentials**](AzureCredentials.md) | The JSON body of the request. Contains the Azure credentials configuration to be validated. | 

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


## ValidateAzureServicesConfig

> ValidateAzureServicesConfig(ctx, id).AzureMonitoredServicesDto(azureMonitoredServicesDto).Execute()

Validates the payload for the `PUT /azure/credentials/{id}/services` request

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
    id := "id_example" // string | The ID of the Azure credentials for which the monitored services configuration is going to be validated.
    azureMonitoredServicesDto := *openapiclient.NewAzureMonitoredServicesDto([]openapiclient.AzureSupportingService{*openapiclient.NewAzureSupportingService("Name_example")}) // AzureMonitoredServicesDto | The JSON body of the request. Contains a monitored services configuration for Azure credentials to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AzureCredentialsConfigurationAPI.ValidateAzureServicesConfig(context.Background(), id).AzureMonitoredServicesDto(azureMonitoredServicesDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.ValidateAzureServicesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Azure credentials for which the monitored services configuration is going to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateAzureServicesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureMonitoredServicesDto** | [**AzureMonitoredServicesDto**](AzureMonitoredServicesDto.md) | The JSON body of the request. Contains a monitored services configuration for Azure credentials to be validated. | 

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


## ValidateConfigurationUpdate

> ValidateConfigurationUpdate(ctx, id).AzureCredentials(azureCredentials).Execute()

Validates the payload for the `PUT /azure/credentials/{id}` request

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
    id := "id_example" // string | The ID of the Azure credentials configuration to be validated.
    azureCredentials := *openapiclient.NewAzureCredentials(false, "Label_example", false) // AzureCredentials | The JSON body of the request. Contains the Azure credentials configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AzureCredentialsConfigurationAPI.ValidateConfigurationUpdate(context.Background(), id).AzureCredentials(azureCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AzureCredentialsConfigurationAPI.ValidateConfigurationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Azure credentials configuration to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateConfigurationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **azureCredentials** | [**AzureCredentials**](AzureCredentials.md) | The JSON body of the request. Contains the Azure credentials configuration to be validated. | 

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

