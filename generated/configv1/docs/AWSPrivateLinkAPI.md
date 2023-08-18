# \AWSPrivateLinkAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllAllowlistedAccounts**](AWSPrivateLinkAPI.md#GetAllAllowlistedAccounts) | **Get** /aws/privateLink/allowlistedAccounts | Gets the information about all allowlisted accounts in AWS PrivateLink
[**GetPrivateLinkConfig**](AWSPrivateLinkAPI.md#GetPrivateLinkConfig) | **Get** /aws/privateLink | Gets the configuration information about AWS PrivateLink
[**PutAllowlistedAccount**](AWSPrivateLinkAPI.md#PutAllowlistedAccount) | **Put** /aws/privateLink/allowlistedAccounts/{id} | Updates the given AWS account id in the allowlist in AWS PrivateLink
[**PutPrivateLinkConfig**](AWSPrivateLinkAPI.md#PutPrivateLinkConfig) | **Put** /aws/privateLink | Updates the configuration information about AWS PrivateLink
[**RemoveAllowlistedAccount**](AWSPrivateLinkAPI.md#RemoveAllowlistedAccount) | **Delete** /aws/privateLink/allowlistedAccounts/{id} | Removes one AWS account from the allowlist in AWS PrivateLink



## GetAllAllowlistedAccounts

> AllowlistedAwsAccountList GetAllAllowlistedAccounts(ctx).Execute()

Gets the information about all allowlisted accounts in AWS PrivateLink



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
    resp, r, err := apiClient.AWSPrivateLinkAPI.GetAllAllowlistedAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSPrivateLinkAPI.GetAllAllowlistedAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllAllowlistedAccounts`: AllowlistedAwsAccountList
    fmt.Fprintf(os.Stdout, "Response from `AWSPrivateLinkAPI.GetAllAllowlistedAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAllowlistedAccountsRequest struct via the builder pattern


### Return type

[**AllowlistedAwsAccountList**](AllowlistedAwsAccountList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrivateLinkConfig

> AwsPrivateLinkConfig GetPrivateLinkConfig(ctx).Execute()

Gets the configuration information about AWS PrivateLink



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
    resp, r, err := apiClient.AWSPrivateLinkAPI.GetPrivateLinkConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSPrivateLinkAPI.GetPrivateLinkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPrivateLinkConfig`: AwsPrivateLinkConfig
    fmt.Fprintf(os.Stdout, "Response from `AWSPrivateLinkAPI.GetPrivateLinkConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrivateLinkConfigRequest struct via the builder pattern


### Return type

[**AwsPrivateLinkConfig**](AwsPrivateLinkConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAllowlistedAccount

> AllowlistedAwsAccount PutAllowlistedAccount(ctx, id).AllowlistedAwsAccount(allowlistedAwsAccount).Execute()

Updates the given AWS account id in the allowlist in AWS PrivateLink



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
    id := "id_example" // string | The AWS account id to be updated in the AWS PrivateLink allowlist, has to match the id in the provided payload.
    allowlistedAwsAccount := *openapiclient.NewAllowlistedAwsAccount("Id_example") // AllowlistedAwsAccount | The AWS account id to be updated in the AWS PrivateLink allowlist, has to match the id in the path.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSPrivateLinkAPI.PutAllowlistedAccount(context.Background(), id).AllowlistedAwsAccount(allowlistedAwsAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSPrivateLinkAPI.PutAllowlistedAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutAllowlistedAccount`: AllowlistedAwsAccount
    fmt.Fprintf(os.Stdout, "Response from `AWSPrivateLinkAPI.PutAllowlistedAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The AWS account id to be updated in the AWS PrivateLink allowlist, has to match the id in the provided payload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutAllowlistedAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **allowlistedAwsAccount** | [**AllowlistedAwsAccount**](AllowlistedAwsAccount.md) | The AWS account id to be updated in the AWS PrivateLink allowlist, has to match the id in the path. | 

### Return type

[**AllowlistedAwsAccount**](AllowlistedAwsAccount.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPrivateLinkConfig

> AwsPrivateLinkConfig PutPrivateLinkConfig(ctx).AwsPrivateLinkConfig(awsPrivateLinkConfig).Execute()

Updates the configuration information about AWS PrivateLink



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
    awsPrivateLinkConfig := *openapiclient.NewAwsPrivateLinkConfig(false) // AwsPrivateLinkConfig | The AWS PrivateLink configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSPrivateLinkAPI.PutPrivateLinkConfig(context.Background()).AwsPrivateLinkConfig(awsPrivateLinkConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSPrivateLinkAPI.PutPrivateLinkConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutPrivateLinkConfig`: AwsPrivateLinkConfig
    fmt.Fprintf(os.Stdout, "Response from `AWSPrivateLinkAPI.PutPrivateLinkConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutPrivateLinkConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsPrivateLinkConfig** | [**AwsPrivateLinkConfig**](AwsPrivateLinkConfig.md) | The AWS PrivateLink configuration. | 

### Return type

[**AwsPrivateLinkConfig**](AwsPrivateLinkConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAllowlistedAccount

> RemoveAllowlistedAccount(ctx, id).Execute()

Removes one AWS account from the allowlist in AWS PrivateLink



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
    id := "id_example" // string | The AWS account id to be deleted from the AWS PrivateLink allowlist

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AWSPrivateLinkAPI.RemoveAllowlistedAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSPrivateLinkAPI.RemoveAllowlistedAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The AWS account id to be deleted from the AWS PrivateLink allowlist | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAllowlistedAccountRequest struct via the builder pattern


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

