# \CredentialVaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentials**](CredentialVaultAPI.md#CreateCredentials) | **Post** /credentials | Creates a new credentials set.
[**GetCredentials**](CredentialVaultAPI.md#GetCredentials) | **Get** /credentials/{id} | Gets the metadata of the specified credentials set.
[**GetCredentialsDetails**](CredentialVaultAPI.md#GetCredentialsDetails) | **Get** /credentials/{id}/details | Gets the details of the specified credentials set.
[**ListCredentials**](CredentialVaultAPI.md#ListCredentials) | **Get** /credentials | Lists all sets of credentials for synthetic monitors stored in your environment.
[**RemoveCredentials**](CredentialVaultAPI.md#RemoveCredentials) | **Delete** /credentials/{id} | Deletes the specified credentials set
[**UpdateCredentials**](CredentialVaultAPI.md#UpdateCredentials) | **Put** /credentials/{id} | Updates the specified credentials set.



## CreateCredentials

> CredentialsId CreateCredentials(ctx).Credentials(credentials).Execute()

Creates a new credentials set.



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

Gets the metadata of the specified credentials set.



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


## GetCredentialsDetails

> AbstractCredentialsResponseElement GetCredentialsDetails(ctx, id).Execute()

Gets the details of the specified credentials set.



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
    resp, r, err := apiClient.CredentialVaultAPI.GetCredentialsDetails(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialVaultAPI.GetCredentialsDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCredentialsDetails`: AbstractCredentialsResponseElement
    fmt.Fprintf(os.Stdout, "Response from `CredentialVaultAPI.GetCredentialsDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The Dynatrace entity ID of the required credentials set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCredentialsDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AbstractCredentialsResponseElement**](AbstractCredentialsResponseElement.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCredentials

> CredentialsList ListCredentials(ctx).Type_(type_).Name(name).User(user).Scope(scope).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Lists all sets of credentials for synthetic monitors stored in your environment.



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
    name := "name_example" // string | Filters the result by the name. When in quotation marks, whole phrase is taken. Case insensitive. (optional)
    user := "user_example" // string | Filters credentials accessible to the user (owned by the user or the ones that are accessible for all). (optional)
    scope := "scope_example" // string | Filters credentials with specified scope. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of credentials in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialVaultAPI.ListCredentials(context.Background()).Type_(type_).Name(name).User(user).Scope(scope).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
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
 **name** | **string** | Filters the result by the name. When in quotation marks, whole phrase is taken. Case insensitive. | 
 **user** | **string** | Filters credentials accessible to the user (owned by the user or the ones that are accessible for all). | 
 **scope** | **string** | Filters credentials with specified scope. | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of credentials in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. | 

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
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentials

> CredentialsId UpdateCredentials(ctx, id).Credentials(credentials).Execute()

Updates the specified credentials set.



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

