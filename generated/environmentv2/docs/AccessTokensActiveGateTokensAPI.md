# \AccessTokensActiveGateTokensAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](AccessTokensActiveGateTokensAPI.md#CreateToken) | **Post** /activeGateTokens | Creates a new ActiveGate token
[**GetToken**](AccessTokensActiveGateTokensAPI.md#GetToken) | **Get** /activeGateTokens/{activeGateTokenIdentifier} | Gets metadata of an ActiveGate token
[**ListTokens**](AccessTokensActiveGateTokensAPI.md#ListTokens) | **Get** /activeGateTokens | Lists all available ActiveGate tokens
[**RevokeToken**](AccessTokensActiveGateTokensAPI.md#RevokeToken) | **Delete** /activeGateTokens/{activeGateTokenIdentifier} | Deletes an ActiveGate token



## CreateToken

> ActiveGateTokenCreated CreateToken(ctx).ActiveGateTokenCreate(activeGateTokenCreate).Execute()

Creates a new ActiveGate token



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
    activeGateTokenCreate := *openapiclient.NewActiveGateTokenCreate("ActiveGateType_example", "myToken") // ActiveGateTokenCreate | The JSON body of the request. Contains parameters of the new ActiveGate token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokensActiveGateTokensAPI.CreateToken(context.Background()).ActiveGateTokenCreate(activeGateTokenCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensActiveGateTokensAPI.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: ActiveGateTokenCreated
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensActiveGateTokensAPI.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **activeGateTokenCreate** | [**ActiveGateTokenCreate**](ActiveGateTokenCreate.md) | The JSON body of the request. Contains parameters of the new ActiveGate token. | 

### Return type

[**ActiveGateTokenCreated**](ActiveGateTokenCreated.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetToken

> ActiveGateToken GetToken(ctx, activeGateTokenIdentifier).Execute()

Gets metadata of an ActiveGate token



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
    activeGateTokenIdentifier := "activeGateTokenIdentifier_example" // string | The ActiveGate token identifier, consisting of [prefix and public part](https://dt-url.net/rn00tjg) of the token.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokensActiveGateTokensAPI.GetToken(context.Background(), activeGateTokenIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensActiveGateTokensAPI.GetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetToken`: ActiveGateToken
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensActiveGateTokensAPI.GetToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activeGateTokenIdentifier** | **string** | The ActiveGate token identifier, consisting of [prefix and public part](https://dt-url.net/rn00tjg) of the token. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ActiveGateToken**](ActiveGateToken.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokens

> ActiveGateTokenList ListTokens(ctx).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Lists all available ActiveGate tokens



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
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of ActiveGate tokens in a single response payload.   The maximal allowed page size is 3000 and the minimal size is 100.   If not set, 100 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokensActiveGateTokensAPI.ListTokens(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensActiveGateTokensAPI.ListTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokens`: ActiveGateTokenList
    fmt.Fprintf(os.Stdout, "Response from `AccessTokensActiveGateTokensAPI.ListTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of ActiveGate tokens in a single response payload.   The maximal allowed page size is 3000 and the minimal size is 100.   If not set, 100 is used. | 

### Return type

[**ActiveGateTokenList**](ActiveGateTokenList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeToken

> RevokeToken(ctx, activeGateTokenIdentifier).Execute()

Deletes an ActiveGate token

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
    activeGateTokenIdentifier := "activeGateTokenIdentifier_example" // string | The ActiveGate token identifier, consisting of [prefix and public part](https://dt-url.net/rn00tjg) of the token to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AccessTokensActiveGateTokensAPI.RevokeToken(context.Background(), activeGateTokenIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokensActiveGateTokensAPI.RevokeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activeGateTokenIdentifier** | **string** | The ActiveGate token identifier, consisting of [prefix and public part](https://dt-url.net/rn00tjg) of the token to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenRequest struct via the builder pattern


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

