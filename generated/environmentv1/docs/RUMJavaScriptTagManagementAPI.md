# \RUMJavaScriptTagManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppRevision**](RUMJavaScriptTagManagementAPI.md#GetAppRevision) | **Get** /rum/appRevision/{entity} | Gets the version of the RUM JavaScript code injected into specified application
[**GetAsyncCodeSnippet**](RUMJavaScriptTagManagementAPI.md#GetAsyncCodeSnippet) | **Get** /rum/asyncCS/{entity} | Downloads the asynchronous code snippet
[**GetJsInlineScript**](RUMJavaScriptTagManagementAPI.md#GetJsInlineScript) | **Get** /rum/jsInlineScript/{entity} | Downloads inline code
[**GetJsLatestVersion**](RUMJavaScriptTagManagementAPI.md#GetJsLatestVersion) | **Get** /rum/jsLatestVersion | Gets the latest version of OneAgent JavaScript library
[**GetJsScript**](RUMJavaScriptTagManagementAPI.md#GetJsScript) | **Get** /rum/jsTag/{entity} | Downloads OneAgent JavaScript tag
[**GetJsTagComplete**](RUMJavaScriptTagManagementAPI.md#GetJsTagComplete) | **Get** /rum/jsTagComplete/{entity} | Downloads JavaScript tag
[**GetManualApps**](RUMJavaScriptTagManagementAPI.md#GetManualApps) | **Get** /rum/manualApps | Lists all manually injected applications
[**GetSyncCodeSnippet**](RUMJavaScriptTagManagementAPI.md#GetSyncCodeSnippet) | **Get** /rum/syncCS/{entity} | Downloads the synchronous code snippet



## GetAppRevision

> string GetAppRevision(ctx, entity).Execute()

Gets the version of the RUM JavaScript code injected into specified application

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
    entity := "entity_example" // string | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMJavaScriptTagManagementAPI.GetAppRevision(context.Background(), entity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMJavaScriptTagManagementAPI.GetAppRevision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppRevision`: string
    fmt.Fprintf(os.Stdout, "Response from `RUMJavaScriptTagManagementAPI.GetAppRevision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRevisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncCodeSnippet

> string GetAsyncCodeSnippet(ctx, entity).Execute()

Downloads the asynchronous code snippet



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
    entity := "entity_example" // string | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMJavaScriptTagManagementAPI.GetAsyncCodeSnippet(context.Background(), entity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMJavaScriptTagManagementAPI.GetAsyncCodeSnippet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsyncCodeSnippet`: string
    fmt.Fprintf(os.Stdout, "Response from `RUMJavaScriptTagManagementAPI.GetAsyncCodeSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncCodeSnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJsInlineScript

> string GetJsInlineScript(ctx, entity).Execute()

Downloads inline code



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
    entity := "entity_example" // string | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMJavaScriptTagManagementAPI.GetJsInlineScript(context.Background(), entity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMJavaScriptTagManagementAPI.GetJsInlineScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsInlineScript`: string
    fmt.Fprintf(os.Stdout, "Response from `RUMJavaScriptTagManagementAPI.GetJsInlineScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsInlineScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJsLatestVersion

> string GetJsLatestVersion(ctx).Execute()

Gets the latest version of OneAgent JavaScript library

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
    resp, r, err := apiClient.RUMJavaScriptTagManagementAPI.GetJsLatestVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMJavaScriptTagManagementAPI.GetJsLatestVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsLatestVersion`: string
    fmt.Fprintf(os.Stdout, "Response from `RUMJavaScriptTagManagementAPI.GetJsLatestVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsLatestVersionRequest struct via the builder pattern


### Return type

**string**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJsScript

> string GetJsScript(ctx, entity).Execute()

Downloads OneAgent JavaScript tag



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
    entity := "entity_example" // string | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMJavaScriptTagManagementAPI.GetJsScript(context.Background(), entity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMJavaScriptTagManagementAPI.GetJsScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsScript`: string
    fmt.Fprintf(os.Stdout, "Response from `RUMJavaScriptTagManagementAPI.GetJsScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJsTagComplete

> string GetJsTagComplete(ctx, entity).Execute()

Downloads JavaScript tag



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
    entity := "entity_example" // string | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMJavaScriptTagManagementAPI.GetJsTagComplete(context.Background(), entity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMJavaScriptTagManagementAPI.GetJsTagComplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJsTagComplete`: string
    fmt.Fprintf(os.Stdout, "Response from `RUMJavaScriptTagManagementAPI.GetJsTagComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetJsTagCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManualApps

> []ManualApplication GetManualApps(ctx).Execute()

Lists all manually injected applications

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
    resp, r, err := apiClient.RUMJavaScriptTagManagementAPI.GetManualApps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMJavaScriptTagManagementAPI.GetManualApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetManualApps`: []ManualApplication
    fmt.Fprintf(os.Stdout, "Response from `RUMJavaScriptTagManagementAPI.GetManualApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetManualAppsRequest struct via the builder pattern


### Return type

[**[]ManualApplication**](ManualApplication.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncCodeSnippet

> string GetSyncCodeSnippet(ctx, entity).Execute()

Downloads the synchronous code snippet



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
    entity := "entity_example" // string | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMJavaScriptTagManagementAPI.GetSyncCodeSnippet(context.Background(), entity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMJavaScriptTagManagementAPI.GetSyncCodeSnippet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSyncCodeSnippet`: string
    fmt.Fprintf(os.Stdout, "Response from `RUMJavaScriptTagManagementAPI.GetSyncCodeSnippet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entity** | **string** | The Dynatrace entity ID of the application.   You can obtain it from the response of the [GET the list of manually injected applications](https://dt-url.net/dl03sgo) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncCodeSnippetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

