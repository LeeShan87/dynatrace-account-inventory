# \NotificationsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationConfig**](NotificationsAPI.md#CreateNotificationConfig) | **Post** /notifications | Creates a new notification configuration
[**DeleteNotificationConfig**](NotificationsAPI.md#DeleteNotificationConfig) | **Delete** /notifications/{id} | Deletes the specified notification configuration
[**GetNotificationConfig**](NotificationsAPI.md#GetNotificationConfig) | **Get** /notifications/{id} | Gets the properties of the specified notification configuration
[**ListNotificationConfigs**](NotificationsAPI.md#ListNotificationConfigs) | **Get** /notifications | Lists all notification configurations available in your environment
[**UpdateNotificationConfig**](NotificationsAPI.md#UpdateNotificationConfig) | **Put** /notifications/{id} | Updates an existing notification configuration or creates a new one
[**ValidateCreateNotificationConfig**](NotificationsAPI.md#ValidateCreateNotificationConfig) | **Post** /notifications/validator | Validates the payload for the &#x60;POST /notifications&#x60; request
[**ValidateUpdateNotificationConfig**](NotificationsAPI.md#ValidateUpdateNotificationConfig) | **Post** /notifications/{id}/validator | Validates the payload the &#x60;PUT /notifications/{id}&#x60; request



## CreateNotificationConfig

> CreateNotificationConfig(ctx).NotificationConfig(notificationConfig).Execute()

Creates a new notification configuration

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
    notificationConfig := *openapiclient.NewNotificationConfig(false, "AlertingProfile_example", "Name_example", "Type_example") // NotificationConfig | The JSON body of the request. Contains parameters of the new notification configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsAPI.CreateNotificationConfig(context.Background()).NotificationConfig(notificationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.CreateNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationConfig** | [**NotificationConfig**](NotificationConfig.md) | The JSON body of the request. Contains parameters of the new notification configuration. | 

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


## DeleteNotificationConfig

> DeleteNotificationConfig(ctx, id).Execute()

Deletes the specified notification configuration

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the notification configuration to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsAPI.DeleteNotificationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.DeleteNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the notification configuration to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationConfigRequest struct via the builder pattern


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


## GetNotificationConfig

> NotificationConfig GetNotificationConfig(ctx, id).Execute()

Gets the properties of the specified notification configuration

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required notification configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.GetNotificationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.GetNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNotificationConfig`: NotificationConfig
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.GetNotificationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required notification configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationConfig**](NotificationConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNotificationConfigs

> NotificationConfigStubListDto ListNotificationConfigs(ctx).Execute()

Lists all notification configurations available in your environment

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
    resp, r, err := apiClient.NotificationsAPI.ListNotificationConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ListNotificationConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNotificationConfigs`: NotificationConfigStubListDto
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.ListNotificationConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListNotificationConfigsRequest struct via the builder pattern


### Return type

[**NotificationConfigStubListDto**](NotificationConfigStubListDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationConfig

> NotificationConfigStub UpdateNotificationConfig(ctx, id).NotificationConfig(notificationConfig).Execute()

Updates an existing notification configuration or creates a new one



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the notification configuration to be updated.    If you set the ID in the body as well, it must match this ID.
    notificationConfig := *openapiclient.NewNotificationConfig(false, "AlertingProfile_example", "Name_example", "Type_example") // NotificationConfig | The JSON body of the request. Contains updated parameters of the notification configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NotificationsAPI.UpdateNotificationConfig(context.Background(), id).NotificationConfig(notificationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.UpdateNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNotificationConfig`: NotificationConfigStub
    fmt.Fprintf(os.Stdout, "Response from `NotificationsAPI.UpdateNotificationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the notification configuration to be updated.    If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationConfig** | [**NotificationConfig**](NotificationConfig.md) | The JSON body of the request. Contains updated parameters of the notification configuration. | 

### Return type

[**NotificationConfigStub**](NotificationConfigStub.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCreateNotificationConfig

> ValidateCreateNotificationConfig(ctx).NotificationConfig(notificationConfig).Execute()

Validates the payload for the `POST /notifications` request

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
    notificationConfig := *openapiclient.NewNotificationConfig(false, "AlertingProfile_example", "Name_example", "Type_example") // NotificationConfig | The JSON body of the request. Contains the notification configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsAPI.ValidateCreateNotificationConfig(context.Background()).NotificationConfig(notificationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ValidateCreateNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationConfig** | [**NotificationConfig**](NotificationConfig.md) | The JSON body of the request. Contains the notification configuration to be validated. | 

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


## ValidateUpdateNotificationConfig

> ValidateUpdateNotificationConfig(ctx, id).NotificationConfig(notificationConfig).Execute()

Validates the payload the `PUT /notifications/{id}` request

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the notification configuration to be validated.
    notificationConfig := *openapiclient.NewNotificationConfig(false, "AlertingProfile_example", "Name_example", "Type_example") // NotificationConfig | The JSON body of the request. Contains the notification configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NotificationsAPI.ValidateUpdateNotificationConfig(context.Background(), id).NotificationConfig(notificationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationsAPI.ValidateUpdateNotificationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the notification configuration to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateNotificationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationConfig** | [**NotificationConfig**](NotificationConfig.md) | The JSON body of the request. Contains the notification configuration to be validated. | 

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

