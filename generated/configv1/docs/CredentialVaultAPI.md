# \CredentialVaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentials**](CredentialVaultAPI.md#CreateCredentials) | **Post** /credentials | Creates a new credentials set
[**GetCredentials**](CredentialVaultAPI.md#GetCredentials) | **Get** /credentials/{id} | Gets the metadata of the specified credentials set
[**ListCredentials**](CredentialVaultAPI.md#ListCredentials) | **Get** /credentials | Lists all sets of credentials for synthetic monitors stored in your environment
[**RemoveCredentials**](CredentialVaultAPI.md#RemoveCredentials) | **Delete** /credentials/{id} | Deletes the specified credentials set
[**UpdateCredentials**](CredentialVaultAPI.md#UpdateCredentials) | **Put** /credentials/{id} | Updates the specified credentials set



## CreateCredentials

> CredentialsId CreateCredentials(ctx).Credentials(credentials).Execute()

Creates a new credentials set



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
    credentials := *openapiclient.NewCredentials("Name_example", "Scope_example", []string{"Scopes_example"}) // Credentials | The JSON body of the request. Contains parameters of the new credentials set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialVaultAPI.CreateCredentials(context.Background()).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultAPI.CreateCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredentials`: CredentialsId
    fmt.Fprintf(os.Stdout, "Response from `CredentialVaultAPI.CreateCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **credentials** | [**Credentials**](Credentials.md) | The JSON body of the request. Contains parameters of the new credentials set. | 

### Return type

[**CredentialsId**](CredentialsId.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentials

> CredentialsResponseElement GetCredentials(ctx, id).Execute()

Gets the metadata of the specified credentials set



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
    id := "id_example" // string | The Dynatrace entity ID of the required credentials set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialVaultAPI.GetCredentials(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultAPI.GetCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCredentials`: CredentialsResponseElement
    fmt.Fprintf(os.Stdout, "Response from `CredentialVaultAPI.GetCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required credentials set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CredentialsResponseElement**](CredentialsResponseElement.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentials

> CredentialsList ListCredentials(ctx).Type_(type_).Execute()

Lists all sets of credentials for synthetic monitors stored in your environment

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
    type_ := "type__example" // string | Filters the result by the specified credentials type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialVaultAPI.ListCredentials(context.Background()).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultAPI.ListCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCredentials`: CredentialsList
    fmt.Fprintf(os.Stdout, "Response from `CredentialVaultAPI.ListCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Filters the result by the specified credentials type. | 

### Return type

[**CredentialsList**](CredentialsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCredentials

> RemoveCredentials(ctx, id).Execute()

Deletes the specified credentials set

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
    id := "id_example" // string | The ID of the credentials set to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CredentialVaultAPI.RemoveCredentials(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultAPI.RemoveCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the credentials set to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCredentialsRequest struct via the builder pattern


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


## UpdateCredentials

> CredentialsId UpdateCredentials(ctx, id).Credentials(credentials).Execute()

Updates the specified credentials set

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
    id := "id_example" // string | The Dynatrace entity ID of the credentials set to be updated.
    credentials := *openapiclient.NewCredentials("Name_example", "Scope_example", []string{"Scopes_example"}) // Credentials | The JSON body of the request. Contains updated parameters of the credentials set.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialVaultAPI.UpdateCredentials(context.Background(), id).Credentials(credentials).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultAPI.UpdateCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCredentials`: CredentialsId
    fmt.Fprintf(os.Stdout, "Response from `CredentialVaultAPI.UpdateCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the credentials set to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentials** | [**Credentials**](Credentials.md) | The JSON body of the request. Contains updated parameters of the credentials set. | 

### Return type

[**CredentialsId**](CredentialsId.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

