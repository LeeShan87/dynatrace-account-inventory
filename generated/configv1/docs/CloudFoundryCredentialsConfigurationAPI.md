# \CloudFoundryCredentialsConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCloudFoundryCredentialsConfig**](CloudFoundryCredentialsConfigurationAPI.md#CreateCloudFoundryCredentialsConfig) | **Post** /cloudFoundry/credentials | Create new credentials for a single foundation. | maturity&#x3D;EARLY_ADOPTER
[**DeleteCloudFoundryCredentialsConfig**](CloudFoundryCredentialsConfigurationAPI.md#DeleteCloudFoundryCredentialsConfig) | **Delete** /cloudFoundry/credentials/{id} | Delete the specified Cloud Foundry foundation credentials. | maturity&#x3D;EARLY_ADOPTER
[**GetCloudFoundryCredentialsConfig**](CloudFoundryCredentialsConfigurationAPI.md#GetCloudFoundryCredentialsConfig) | **Get** /cloudFoundry/credentials/{id} | Show the properties for the specified Cloud Foundry foundation credentials. | maturity&#x3D;EARLY_ADOPTER
[**ListCloudFoundryCredentialsConfigs**](CloudFoundryCredentialsConfigurationAPI.md#ListCloudFoundryCredentialsConfigs) | **Get** /cloudFoundry/credentials | List all the Cloud Foundry foundations credentials. | maturity&#x3D;EARLY_ADOPTER
[**UpdateCloudFoundryCredentialsConfig**](CloudFoundryCredentialsConfigurationAPI.md#UpdateCloudFoundryCredentialsConfig) | **Put** /cloudFoundry/credentials/{id} | Create or update credentials for a single Cloud Foundry foundation. | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateCloudFoundryCredentialsConfig**](CloudFoundryCredentialsConfigurationAPI.md#ValidateCreateCloudFoundryCredentialsConfig) | **Post** /cloudFoundry/credentials/validator | Validate that creating credentials would be successful. | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateCloudFoundryCredentialsConfig**](CloudFoundryCredentialsConfigurationAPI.md#ValidateUpdateCloudFoundryCredentialsConfig) | **Post** /cloudFoundry/credentials/{id}/validator | Validate that updating or creating credentials would be successful. | maturity&#x3D;EARLY_ADOPTER



## CreateCloudFoundryCredentialsConfig

> EntityShortRepresentation CreateCloudFoundryCredentialsConfig(ctx).CloudFoundryCredentials(cloudFoundryCredentials).Execute()

Create new credentials for a single foundation. | maturity=EARLY_ADOPTER

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
    cloudFoundryCredentials := *openapiclient.NewCloudFoundryCredentials("ApiUrl_example", "LoginUrl_example", "Name_example", "Username_example") // CloudFoundryCredentials | `name`, `apiUrl` and `loginUrl` must be unique.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudFoundryCredentialsConfigurationAPI.CreateCloudFoundryCredentialsConfig(context.Background()).CloudFoundryCredentials(cloudFoundryCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudFoundryCredentialsConfigurationAPI.CreateCloudFoundryCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCloudFoundryCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `CloudFoundryCredentialsConfigurationAPI.CreateCloudFoundryCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCloudFoundryCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudFoundryCredentials** | [**CloudFoundryCredentials**](CloudFoundryCredentials.md) | &#x60;name&#x60;, &#x60;apiUrl&#x60; and &#x60;loginUrl&#x60; must be unique. | 

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


## DeleteCloudFoundryCredentialsConfig

> DeleteCloudFoundryCredentialsConfig(ctx, id).Execute()

Delete the specified Cloud Foundry foundation credentials. | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the Cloud Foundry foundation credentials to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CloudFoundryCredentialsConfigurationAPI.DeleteCloudFoundryCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudFoundryCredentialsConfigurationAPI.DeleteCloudFoundryCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Cloud Foundry foundation credentials to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCloudFoundryCredentialsConfigRequest struct via the builder pattern


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


## GetCloudFoundryCredentialsConfig

> CloudFoundryCredentials GetCloudFoundryCredentialsConfig(ctx, id).Execute()

Show the properties for the specified Cloud Foundry foundation credentials. | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required Cloud Foundry foundation credentials.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudFoundryCredentialsConfigurationAPI.GetCloudFoundryCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudFoundryCredentialsConfigurationAPI.GetCloudFoundryCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCloudFoundryCredentialsConfig`: CloudFoundryCredentials
    fmt.Fprintf(os.Stdout, "Response from `CloudFoundryCredentialsConfigurationAPI.GetCloudFoundryCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required Cloud Foundry foundation credentials. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCloudFoundryCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudFoundryCredentials**](CloudFoundryCredentials.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudFoundryCredentialsConfigs

> StubList ListCloudFoundryCredentialsConfigs(ctx).Execute()

List all the Cloud Foundry foundations credentials. | maturity=EARLY_ADOPTER

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
    resp, r, err := apiClient.CloudFoundryCredentialsConfigurationAPI.ListCloudFoundryCredentialsConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudFoundryCredentialsConfigurationAPI.ListCloudFoundryCredentialsConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudFoundryCredentialsConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `CloudFoundryCredentialsConfigurationAPI.ListCloudFoundryCredentialsConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudFoundryCredentialsConfigsRequest struct via the builder pattern


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


## UpdateCloudFoundryCredentialsConfig

> EntityShortRepresentation UpdateCloudFoundryCredentialsConfig(ctx, id).CloudFoundryCredentials(cloudFoundryCredentials).Execute()

Create or update credentials for a single Cloud Foundry foundation. | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the Cloud Foundry foundation credentials.
    cloudFoundryCredentials := *openapiclient.NewCloudFoundryCredentials("ApiUrl_example", "LoginUrl_example", "Name_example", "Username_example") // CloudFoundryCredentials | `name` must be unique. `password` can be omitted for updates, the existing one will be used. `apiUrl` and `loginUrl` must be set and may not differ from the existing config if it exists. Use this endpoint for copying credentials between environments while keeping their IDs and for updating existing credentials. You can *not* use this to create new credentials with an arbitrary ID, use POST instead.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudFoundryCredentialsConfigurationAPI.UpdateCloudFoundryCredentialsConfig(context.Background(), id).CloudFoundryCredentials(cloudFoundryCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudFoundryCredentialsConfigurationAPI.UpdateCloudFoundryCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCloudFoundryCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `CloudFoundryCredentialsConfigurationAPI.UpdateCloudFoundryCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Cloud Foundry foundation credentials. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCloudFoundryCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudFoundryCredentials** | [**CloudFoundryCredentials**](CloudFoundryCredentials.md) | &#x60;name&#x60; must be unique. &#x60;password&#x60; can be omitted for updates, the existing one will be used. &#x60;apiUrl&#x60; and &#x60;loginUrl&#x60; must be set and may not differ from the existing config if it exists. Use this endpoint for copying credentials between environments while keeping their IDs and for updating existing credentials. You can *not* use this to create new credentials with an arbitrary ID, use POST instead. | 

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


## ValidateCreateCloudFoundryCredentialsConfig

> ValidateCreateCloudFoundryCredentialsConfig(ctx).CloudFoundryCredentials(cloudFoundryCredentials).Execute()

Validate that creating credentials would be successful. | maturity=EARLY_ADOPTER

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
    cloudFoundryCredentials := *openapiclient.NewCloudFoundryCredentials("ApiUrl_example", "LoginUrl_example", "Name_example", "Username_example") // CloudFoundryCredentials | `name`, `apiUrl` and `loginUrl` must be unique.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CloudFoundryCredentialsConfigurationAPI.ValidateCreateCloudFoundryCredentialsConfig(context.Background()).CloudFoundryCredentials(cloudFoundryCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudFoundryCredentialsConfigurationAPI.ValidateCreateCloudFoundryCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateCloudFoundryCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudFoundryCredentials** | [**CloudFoundryCredentials**](CloudFoundryCredentials.md) | &#x60;name&#x60;, &#x60;apiUrl&#x60; and &#x60;loginUrl&#x60; must be unique. | 

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


## ValidateUpdateCloudFoundryCredentialsConfig

> ValidateUpdateCloudFoundryCredentialsConfig(ctx, id).CloudFoundryCredentials(cloudFoundryCredentials).Execute()

Validate that updating or creating credentials would be successful. | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the Cloud Foundry foundation credentials.
    cloudFoundryCredentials := *openapiclient.NewCloudFoundryCredentials("ApiUrl_example", "LoginUrl_example", "Name_example", "Username_example") // CloudFoundryCredentials | `name` must be unique. `password` can be omitted for updates. See the constraints for the PUT /cloudFoundry/credentials/{id} request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CloudFoundryCredentialsConfigurationAPI.ValidateUpdateCloudFoundryCredentialsConfig(context.Background(), id).CloudFoundryCredentials(cloudFoundryCredentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudFoundryCredentialsConfigurationAPI.ValidateUpdateCloudFoundryCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the Cloud Foundry foundation credentials. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateCloudFoundryCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudFoundryCredentials** | [**CloudFoundryCredentials**](CloudFoundryCredentials.md) | &#x60;name&#x60; must be unique. &#x60;password&#x60; can be omitted for updates. See the constraints for the PUT /cloudFoundry/credentials/{id} request. | 

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

