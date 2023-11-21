# \RUMMobileAndCustomApplicationConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationAPI.md#CreateMobileApplicationConfig) | **Post** /applications/mobile | Creates a new mobile or custom application
[**CreateMobileKeyUserAction**](RUMMobileAndCustomApplicationConfigurationAPI.md#CreateMobileKeyUserAction) | **Post** /applications/mobile/{applicationId}/keyUserActions/{actionName} | Marks the user action as a key user action in the specified application
[**CreateSessionProperty**](RUMMobileAndCustomApplicationConfigurationAPI.md#CreateSessionProperty) | **Post** /applications/mobile/{applicationId}/userActionAndSessionProperties | Creates a new mobile session or user action property for the specified application
[**DeleteMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationAPI.md#DeleteMobileApplicationConfig) | **Delete** /applications/mobile/{id} | Deletes the specified mobile or custom application
[**DeleteMobileKeyUserAction**](RUMMobileAndCustomApplicationConfigurationAPI.md#DeleteMobileKeyUserAction) | **Delete** /applications/mobile/{applicationId}/keyUserActions/{actionName} | Removes the specified user action from the list of key user actions in the specified application
[**DeleteSessionProperty**](RUMMobileAndCustomApplicationConfigurationAPI.md#DeleteSessionProperty) | **Delete** /applications/mobile/{applicationId}/userActionAndSessionProperties/{key} | Deletes the specified mobile session or user action property for an application
[**GetMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationAPI.md#GetMobileApplicationConfig) | **Get** /applications/mobile/{id} | Gets the configuration of the specified mobile or custom application
[**GetSessionProperty**](RUMMobileAndCustomApplicationConfigurationAPI.md#GetSessionProperty) | **Get** /applications/mobile/{applicationId}/userActionAndSessionProperties/{key} | Gets the specified mobile session or user action property of an application
[**ListMobileApplicationConfigs**](RUMMobileAndCustomApplicationConfigurationAPI.md#ListMobileApplicationConfigs) | **Get** /applications/mobile | Lists all existing mobile and custom applications
[**ListMobileKeyUserActions**](RUMMobileAndCustomApplicationConfigurationAPI.md#ListMobileKeyUserActions) | **Get** /applications/mobile/{applicationId}/keyUserActions | Gets the list of key user actions in the specified application
[**ListSessionProperties**](RUMMobileAndCustomApplicationConfigurationAPI.md#ListSessionProperties) | **Get** /applications/mobile/{applicationId}/userActionAndSessionProperties | Lists all mobile session and user action properties for the specified application
[**UpdateMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationAPI.md#UpdateMobileApplicationConfig) | **Put** /applications/mobile/{id} | Updates the configuration of the specified mobile or custom application
[**UpdateSessionProperty**](RUMMobileAndCustomApplicationConfigurationAPI.md#UpdateSessionProperty) | **Put** /applications/mobile/{applicationId}/userActionAndSessionProperties/{key} | Updates the specified mobile session or user action property for an application
[**ValidateCreateMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationAPI.md#ValidateCreateMobileApplicationConfig) | **Post** /applications/mobile/validator | Validates the payload for the &#x60;POST /applications/mobile&#x60; request
[**ValidateCreateSessionProperty**](RUMMobileAndCustomApplicationConfigurationAPI.md#ValidateCreateSessionProperty) | **Post** /applications/mobile/{applicationId}/userActionAndSessionProperties/validator | Validates the payload for the &#x60;POST /applications/mobile/{applicationId}/userActionAndSessionProperties&#x60; request
[**ValidateUpdateMobileApplicationConfig**](RUMMobileAndCustomApplicationConfigurationAPI.md#ValidateUpdateMobileApplicationConfig) | **Post** /applications/mobile/{id}/validator | Validates the payload for the &#x60;PUT /applications/mobile/{id}&#x60; request.
[**ValidateUpdateSessionProperty**](RUMMobileAndCustomApplicationConfigurationAPI.md#ValidateUpdateSessionProperty) | **Post** /applications/mobile/{applicationId}/userActionAndSessionProperties/{key}/validator | Validates the payload for the &#x60;PUT /applications/mobile/{applicationId}/userActionAndSessionProperties/{key}&#x60; request



## CreateMobileApplicationConfig

> EntityShortRepresentation CreateMobileApplicationConfig(ctx).NewMobileCustomAppConfig(newMobileCustomAppConfig).Execute()

Creates a new mobile or custom application



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
    newMobileCustomAppConfig := *openapiclient.NewNewMobileCustomAppConfig("ApplicationType_example", "Name_example") // NewMobileCustomAppConfig | The JSON body of the request. Contains the parameters of the new mobile or custom application. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.CreateMobileApplicationConfig(context.Background()).NewMobileCustomAppConfig(newMobileCustomAppConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.CreateMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMobileApplicationConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationAPI.CreateMobileApplicationConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMobileApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newMobileCustomAppConfig** | [**NewMobileCustomAppConfig**](NewMobileCustomAppConfig.md) | The JSON body of the request. Contains the parameters of the new mobile or custom application. | 

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


## CreateMobileKeyUserAction

> KeyUserActionMobile CreateMobileKeyUserAction(ctx, applicationId, actionName).Execute()

Marks the user action as a key user action in the specified application

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    actionName := "actionName_example" // string | The name of the user action to be marked as a key user action.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.CreateMobileKeyUserAction(context.Background(), applicationId, actionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.CreateMobileKeyUserAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMobileKeyUserAction`: KeyUserActionMobile
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationAPI.CreateMobileKeyUserAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**actionName** | **string** | The name of the user action to be marked as a key user action. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMobileKeyUserActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KeyUserActionMobile**](KeyUserActionMobile.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSessionProperty

> MobileSessionUserActionPropertyShort CreateSessionProperty(ctx, applicationId).MobileSessionUserActionProperty(mobileSessionUserActionProperty).Execute()

Creates a new mobile session or user action property for the specified application

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    mobileSessionUserActionProperty := *openapiclient.NewMobileSessionUserActionProperty("Key_example", "Origin_example", "Type_example") // MobileSessionUserActionProperty | The JSON body of the request. Contains the configuration of the property. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.CreateSessionProperty(context.Background(), applicationId).MobileSessionUserActionProperty(mobileSessionUserActionProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.CreateSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSessionProperty`: MobileSessionUserActionPropertyShort
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationAPI.CreateSessionProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSessionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mobileSessionUserActionProperty** | [**MobileSessionUserActionProperty**](MobileSessionUserActionProperty.md) | The JSON body of the request. Contains the configuration of the property. | 

### Return type

[**MobileSessionUserActionPropertyShort**](MobileSessionUserActionPropertyShort.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMobileApplicationConfig

> DeleteMobileApplicationConfig(ctx, id).Execute()

Deletes the specified mobile or custom application

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
    id := "id_example" // string | The ID of the application to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.DeleteMobileApplicationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.DeleteMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the application to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMobileApplicationConfigRequest struct via the builder pattern


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


## DeleteMobileKeyUserAction

> DeleteMobileKeyUserAction(ctx, applicationId, actionName).Execute()

Removes the specified user action from the list of key user actions in the specified application

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    actionName := "actionName_example" // string | The ID of the user action to be removed from the key user actions list.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.DeleteMobileKeyUserAction(context.Background(), applicationId, actionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.DeleteMobileKeyUserAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**actionName** | **string** | The ID of the user action to be removed from the key user actions list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMobileKeyUserActionRequest struct via the builder pattern


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


## DeleteSessionProperty

> DeleteSessionProperty(ctx, applicationId, key).Execute()

Deletes the specified mobile session or user action property for an application

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    key := "key_example" // string | The key of the required mobile session or user action property.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.DeleteSessionProperty(context.Background(), applicationId, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.DeleteSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**key** | **string** | The key of the required mobile session or user action property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSessionPropertyRequest struct via the builder pattern


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


## GetMobileApplicationConfig

> MobileCustomAppConfig GetMobileApplicationConfig(ctx, id).Execute()

Gets the configuration of the specified mobile or custom application

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
    id := "id_example" // string | The ID of the required mobile or custom application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.GetMobileApplicationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.GetMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMobileApplicationConfig`: MobileCustomAppConfig
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationAPI.GetMobileApplicationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required mobile or custom application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMobileApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileCustomAppConfig**](MobileCustomAppConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionProperty

> MobileSessionUserActionProperty GetSessionProperty(ctx, applicationId, key).Execute()

Gets the specified mobile session or user action property of an application

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    key := "key_example" // string | The key of the required mobile session or user action property.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.GetSessionProperty(context.Background(), applicationId, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.GetSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionProperty`: MobileSessionUserActionProperty
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationAPI.GetSessionProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**key** | **string** | The key of the required mobile session or user action property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MobileSessionUserActionProperty**](MobileSessionUserActionProperty.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMobileApplicationConfigs

> StubList ListMobileApplicationConfigs(ctx).Execute()

Lists all existing mobile and custom applications

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
    resp, r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ListMobileApplicationConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.ListMobileApplicationConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMobileApplicationConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationAPI.ListMobileApplicationConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMobileApplicationConfigsRequest struct via the builder pattern


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


## ListMobileKeyUserActions

> KeyUserActionMobileList ListMobileKeyUserActions(ctx, applicationId).Execute()

Gets the list of key user actions in the specified application

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
    applicationId := "applicationId_example" // string | The ID of the required application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ListMobileKeyUserActions(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.ListMobileKeyUserActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMobileKeyUserActions`: KeyUserActionMobileList
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationAPI.ListMobileKeyUserActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMobileKeyUserActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyUserActionMobileList**](KeyUserActionMobileList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionProperties

> MobileSessionUserActionPropertyList ListSessionProperties(ctx, applicationId).Execute()

Lists all mobile session and user action properties for the specified application

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
    applicationId := "applicationId_example" // string | The ID of the required application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ListSessionProperties(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.ListSessionProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionProperties`: MobileSessionUserActionPropertyList
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationAPI.ListSessionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MobileSessionUserActionPropertyList**](MobileSessionUserActionPropertyList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMobileApplicationConfig

> UpdateMobileApplicationConfig(ctx, id).MobileCustomAppConfig(mobileCustomAppConfig).Execute()

Updates the configuration of the specified mobile or custom application



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
    id := "id_example" // string | The ID of the application to be updated.
    mobileCustomAppConfig := *openapiclient.NewMobileCustomAppConfig(*openapiclient.NewMobileCustomApdex(false, int32(123), int32(123)), "BeaconEndpointType_example", int32(123), "Name_example") // MobileCustomAppConfig | The JSON body of the request. Contains updated configuration of the mobile or custom application.   Do not set the identifier in the body. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.UpdateMobileApplicationConfig(context.Background(), id).MobileCustomAppConfig(mobileCustomAppConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.UpdateMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the application to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMobileApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mobileCustomAppConfig** | [**MobileCustomAppConfig**](MobileCustomAppConfig.md) | The JSON body of the request. Contains updated configuration of the mobile or custom application.   Do not set the identifier in the body. | 

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


## UpdateSessionProperty

> MobileSessionUserActionPropertyShort UpdateSessionProperty(ctx, applicationId, key).MobileSessionUserActionPropertyUpd(mobileSessionUserActionPropertyUpd).Execute()

Updates the specified mobile session or user action property for an application



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
    applicationId := "applicationId_example" // string | The ID of the required application.
    key := "key_example" // string | The key of the required mobile session or user action property.
    mobileSessionUserActionPropertyUpd := *openapiclient.NewMobileSessionUserActionPropertyUpd("Origin_example", "Type_example") // MobileSessionUserActionPropertyUpd | The JSON body of the request. Contains the configuration of the property. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.UpdateSessionProperty(context.Background(), applicationId, key).MobileSessionUserActionPropertyUpd(mobileSessionUserActionPropertyUpd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.UpdateSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSessionProperty`: MobileSessionUserActionPropertyShort
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileAndCustomApplicationConfigurationAPI.UpdateSessionProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**key** | **string** | The key of the required mobile session or user action property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSessionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mobileSessionUserActionPropertyUpd** | [**MobileSessionUserActionPropertyUpd**](MobileSessionUserActionPropertyUpd.md) | The JSON body of the request. Contains the configuration of the property. | 

### Return type

[**MobileSessionUserActionPropertyShort**](MobileSessionUserActionPropertyShort.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCreateMobileApplicationConfig

> ValidateCreateMobileApplicationConfig(ctx).NewMobileCustomAppConfig(newMobileCustomAppConfig).Execute()

Validates the payload for the `POST /applications/mobile` request

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
    newMobileCustomAppConfig := *openapiclient.NewNewMobileCustomAppConfig("ApplicationType_example", "Name_example") // NewMobileCustomAppConfig | The JSON body of the request. Contains the mobile or custom application configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ValidateCreateMobileApplicationConfig(context.Background()).NewMobileCustomAppConfig(newMobileCustomAppConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.ValidateCreateMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateMobileApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newMobileCustomAppConfig** | [**NewMobileCustomAppConfig**](NewMobileCustomAppConfig.md) | The JSON body of the request. Contains the mobile or custom application configuration to be validated. | 

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


## ValidateCreateSessionProperty

> ValidateCreateSessionProperty(ctx, applicationId).MobileSessionUserActionProperty(mobileSessionUserActionProperty).Execute()

Validates the payload for the `POST /applications/mobile/{applicationId}/userActionAndSessionProperties` request

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    mobileSessionUserActionProperty := *openapiclient.NewMobileSessionUserActionProperty("Key_example", "Origin_example", "Type_example") // MobileSessionUserActionProperty | The JSON body of the request. Contains the configuration of the property to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ValidateCreateSessionProperty(context.Background(), applicationId).MobileSessionUserActionProperty(mobileSessionUserActionProperty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.ValidateCreateSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateSessionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mobileSessionUserActionProperty** | [**MobileSessionUserActionProperty**](MobileSessionUserActionProperty.md) | The JSON body of the request. Contains the configuration of the property to be validated. | 

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


## ValidateUpdateMobileApplicationConfig

> ValidateUpdateMobileApplicationConfig(ctx, id).MobileCustomAppConfig(mobileCustomAppConfig).Execute()

Validates the payload for the `PUT /applications/mobile/{id}` request.

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
    id := "id_example" // string | The ID of the mobile or custom application to be validated.
    mobileCustomAppConfig := *openapiclient.NewMobileCustomAppConfig(*openapiclient.NewMobileCustomApdex(false, int32(123), int32(123)), "BeaconEndpointType_example", int32(123), "Name_example") // MobileCustomAppConfig | The JSON body of the request. Contains the mobile or custom application configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ValidateUpdateMobileApplicationConfig(context.Background(), id).MobileCustomAppConfig(mobileCustomAppConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.ValidateUpdateMobileApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the mobile or custom application to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateMobileApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mobileCustomAppConfig** | [**MobileCustomAppConfig**](MobileCustomAppConfig.md) | The JSON body of the request. Contains the mobile or custom application configuration to be validated. | 

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


## ValidateUpdateSessionProperty

> ValidateUpdateSessionProperty(ctx, applicationId, key).MobileSessionUserActionPropertyUpd(mobileSessionUserActionPropertyUpd).Execute()

Validates the payload for the `PUT /applications/mobile/{applicationId}/userActionAndSessionProperties/{key}` request

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
    applicationId := "applicationId_example" // string | The ID of the required application.
    key := "key_example" // string | The key of the required mobile session or user action property.
    mobileSessionUserActionPropertyUpd := *openapiclient.NewMobileSessionUserActionPropertyUpd("Origin_example", "Type_example") // MobileSessionUserActionPropertyUpd | The JSON body of the request. Contains the configuration of the property to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileAndCustomApplicationConfigurationAPI.ValidateUpdateSessionProperty(context.Background(), applicationId, key).MobileSessionUserActionPropertyUpd(mobileSessionUserActionPropertyUpd).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileAndCustomApplicationConfigurationAPI.ValidateUpdateSessionProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required application. | 
**key** | **string** | The key of the required mobile session or user action property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateSessionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mobileSessionUserActionPropertyUpd** | [**MobileSessionUserActionPropertyUpd**](MobileSessionUserActionPropertyUpd.md) | The JSON body of the request. Contains the configuration of the property to be validated. | 

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

