# \RUMWebApplicationConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeyUserAction**](RUMWebApplicationConfigurationAPI.md#CreateKeyUserAction) | **Post** /applications/web/{id}/keyUserActions | Marks the user action as a key user action in the specified web application
[**CreateOrUpdateDefaultConfiguration**](RUMWebApplicationConfigurationAPI.md#CreateOrUpdateDefaultConfiguration) | **Put** /applications/web/default | Updates the configuration of the default web application
[**CreateWebApplicationConfig**](RUMWebApplicationConfigurationAPI.md#CreateWebApplicationConfig) | **Post** /applications/web | Creates a new web application
[**DeleteKeyUserAction**](RUMWebApplicationConfigurationAPI.md#DeleteKeyUserAction) | **Delete** /applications/web/{id}/keyUserActions/{keyUserActionId} | Removes the specified user action from the list of key user actions in the specified web application
[**DeleteWebApplicationConfig**](RUMWebApplicationConfigurationAPI.md#DeleteWebApplicationConfig) | **Delete** /applications/web/{id} | Deletes the specified web application
[**GetApplicationErrorConfig**](RUMWebApplicationConfigurationAPI.md#GetApplicationErrorConfig) | **Get** /applications/web/{id}/errorRules | Gets the configuration of error rules in the specified application
[**GetDataPrivacySettings**](RUMWebApplicationConfigurationAPI.md#GetDataPrivacySettings) | **Get** /applications/web/{id}/dataPrivacy | Gets the data privacy settings of the specified web application
[**GetDefaultApplication**](RUMWebApplicationConfigurationAPI.md#GetDefaultApplication) | **Get** /applications/web/default | Gets the configuration of the default web application
[**GetDefaultApplicationDataPrivacySettings**](RUMWebApplicationConfigurationAPI.md#GetDefaultApplicationDataPrivacySettings) | **Get** /applications/web/default/dataPrivacy | Gets the data privacy settings of the default web application
[**GetWebApplicationConfig**](RUMWebApplicationConfigurationAPI.md#GetWebApplicationConfig) | **Get** /applications/web/{id} | Gets the configuration of the specified web application
[**ListDataPrivacySettings**](RUMWebApplicationConfigurationAPI.md#ListDataPrivacySettings) | **Get** /applications/web/dataPrivacy | Lists data privacy settings of all web applications
[**ListKeyUserActions**](RUMWebApplicationConfigurationAPI.md#ListKeyUserActions) | **Get** /applications/web/{id}/keyUserActions | Gets the list of key user actions in the specified web application
[**ListWebApplicationConfigs**](RUMWebApplicationConfigurationAPI.md#ListWebApplicationConfigs) | **Get** /applications/web | Lists all existing web applications
[**UpdateApplicationErrorConfig**](RUMWebApplicationConfigurationAPI.md#UpdateApplicationErrorConfig) | **Put** /applications/web/{id}/errorRules | Updates the configuration of error rules in the specified application
[**UpdateDataPrivacySettings**](RUMWebApplicationConfigurationAPI.md#UpdateDataPrivacySettings) | **Put** /applications/web/{id}/dataPrivacy | Updates the data privacy settings of the specified web application
[**UpdateDefaultApplicationDataPrivacySettings**](RUMWebApplicationConfigurationAPI.md#UpdateDefaultApplicationDataPrivacySettings) | **Put** /applications/web/default/dataPrivacy | Updates the data privacy settings of the default web application
[**UpdateWebApplicationConfig**](RUMWebApplicationConfigurationAPI.md#UpdateWebApplicationConfig) | **Put** /applications/web/{id} | Updates the configuration of the specified web application or creates a new one
[**ValidateCreateWebApplicationConfig**](RUMWebApplicationConfigurationAPI.md#ValidateCreateWebApplicationConfig) | **Post** /applications/web/validator | Validates the configuration of the web application for the &#x60;POST /applications/web&#x60; request
[**ValidateDataPrivacySettings**](RUMWebApplicationConfigurationAPI.md#ValidateDataPrivacySettings) | **Post** /applications/web/{id}/dataPrivacy/validator | Validates data privacy settings for the &#x60;PUT /applications/web/{id}/dataPrivacy&#x60; request
[**ValidateDefaultApplicationDataPrivacySettings**](RUMWebApplicationConfigurationAPI.md#ValidateDefaultApplicationDataPrivacySettings) | **Post** /applications/web/default/dataPrivacy/validator | Validates data privacy settings of the default web application for the &#x60;PUT /applications/web/default/dataPrivacy&#x60; request
[**ValidateDefaultConfiguration**](RUMWebApplicationConfigurationAPI.md#ValidateDefaultConfiguration) | **Post** /applications/web/default/validator | Validates the configuration of the default web application for the &#x60;PUT /applications/web/default&#x60; request
[**ValidateUpdateWebApplicationConfig**](RUMWebApplicationConfigurationAPI.md#ValidateUpdateWebApplicationConfig) | **Post** /applications/web/{id}/validator | Validates the configuration of the web application for the &#x60;PUT /applications/web/{id}&#x60; request.



## CreateKeyUserAction

> EntityShortRepresentation CreateKeyUserAction(ctx, id).KeyUserAction(keyUserAction).Execute()

Marks the user action as a key user action in the specified web application

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
    id := "id_example" // string | The ID of the required web application.
    keyUserAction := *openapiclient.NewKeyUserAction("ActionType_example", "Name_example") // KeyUserAction | The JSON body of the request. Contains the action to be marked as a key user action. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.CreateKeyUserAction(context.Background(), id).KeyUserAction(keyUserAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.CreateKeyUserAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKeyUserAction`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.CreateKeyUserAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required web application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyUserActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keyUserAction** | [**KeyUserAction**](KeyUserAction.md) | The JSON body of the request. Contains the action to be marked as a key user action. | 

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


## CreateOrUpdateDefaultConfiguration

> CreateOrUpdateDefaultConfiguration(ctx).WebApplicationConfig(webApplicationConfig).Execute()

Updates the configuration of the default web application



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
    webApplicationConfig := *openapiclient.NewWebApplicationConfig(float32(123), *openapiclient.NewApdex(), *openapiclient.NewApdex(), "LoadActionKeyPerformanceMetric_example", *openapiclient.NewMonitoringSettings(*openapiclient.NewAdvancedJavaScriptTagSettings(*openapiclient.NewAdditionalEventHandlers(false, false, false, int32(123), false, false, false), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings("AdditionalEventCapturedAsUserInput_example", false, false, false, false, false, false, false), false, int32(123), int32(123), "SpecialCharactersToEscape_example"), false, *openapiclient.NewContentCapture(false, *openapiclient.NewResourceTimingSettings(false, int32(123), false), *openapiclient.NewTimeoutSettings(int32(123), int32(123), false), false), "CookiePlacementDomain_example", "CustomConfigurationProperties_example", "ExcludeXhrRegex_example", false, "InjectionMode_example", *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), false, "ServerRequestPathId_example", false), "Name_example", false, *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewApdex(), "XhrActionKeyPerformanceMetric_example") // WebApplicationConfig | JSON body of the request, containing the new parameters of the default web application. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.CreateOrUpdateDefaultConfiguration(context.Background()).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.CreateOrUpdateDefaultConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateDefaultConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing the new parameters of the default web application. | 

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


## CreateWebApplicationConfig

> EntityShortRepresentation CreateWebApplicationConfig(ctx).WebApplicationConfig(webApplicationConfig).Execute()

Creates a new web application



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
    webApplicationConfig := *openapiclient.NewWebApplicationConfig(float32(123), *openapiclient.NewApdex(), *openapiclient.NewApdex(), "LoadActionKeyPerformanceMetric_example", *openapiclient.NewMonitoringSettings(*openapiclient.NewAdvancedJavaScriptTagSettings(*openapiclient.NewAdditionalEventHandlers(false, false, false, int32(123), false, false, false), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings("AdditionalEventCapturedAsUserInput_example", false, false, false, false, false, false, false), false, int32(123), int32(123), "SpecialCharactersToEscape_example"), false, *openapiclient.NewContentCapture(false, *openapiclient.NewResourceTimingSettings(false, int32(123), false), *openapiclient.NewTimeoutSettings(int32(123), int32(123), false), false), "CookiePlacementDomain_example", "CustomConfigurationProperties_example", "ExcludeXhrRegex_example", false, "InjectionMode_example", *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), false, "ServerRequestPathId_example", false), "Name_example", false, *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewApdex(), "XhrActionKeyPerformanceMetric_example") // WebApplicationConfig | JSON body of the request, containing parameters of the new web application configuraiton. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.CreateWebApplicationConfig(context.Background()).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.CreateWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebApplicationConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.CreateWebApplicationConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing parameters of the new web application configuraiton. | 

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


## DeleteKeyUserAction

> DeleteKeyUserAction(ctx, id, keyUserActionId).Execute()

Removes the specified user action from the list of key user actions in the specified web application

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
    id := "id_example" // string | The ID of the required web application.
    keyUserActionId := "keyUserActionId_example" // string | The ID of the user action to be removed from the key user actions list.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.DeleteKeyUserAction(context.Background(), id, keyUserActionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.DeleteKeyUserAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required web application. | 
**keyUserActionId** | **string** | The ID of the user action to be removed from the key user actions list. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeyUserActionRequest struct via the builder pattern


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


## DeleteWebApplicationConfig

> DeleteWebApplicationConfig(ctx, id).Execute()

Deletes the specified web application

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
    id := "id_example" // string | The ID of the web application to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.DeleteWebApplicationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.DeleteWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebApplicationConfigRequest struct via the builder pattern


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


## GetApplicationErrorConfig

> ApplicationErrorRules GetApplicationErrorConfig(ctx, id).Execute()

Gets the configuration of error rules in the specified application

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
    id := "id_example" // string | The ID of the required web application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.GetApplicationErrorConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.GetApplicationErrorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationErrorConfig`: ApplicationErrorRules
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.GetApplicationErrorConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required web application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationErrorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationErrorRules**](ApplicationErrorRules.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataPrivacySettings

> ApplicationDataPrivacy GetDataPrivacySettings(ctx, id).Execute()

Gets the data privacy settings of the specified web application

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
    id := "id_example" // string | The ID of the web application where you want to check data privacy settings.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.GetDataPrivacySettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.GetDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataPrivacySettings`: ApplicationDataPrivacy
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.GetDataPrivacySettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application where you want to check data privacy settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataPrivacySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationDataPrivacy**](ApplicationDataPrivacy.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultApplication

> WebApplicationConfig GetDefaultApplication(ctx).Execute()

Gets the configuration of the default web application



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
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.GetDefaultApplication(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.GetDefaultApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultApplication`: WebApplicationConfig
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.GetDefaultApplication`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultApplicationRequest struct via the builder pattern


### Return type

[**WebApplicationConfig**](WebApplicationConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultApplicationDataPrivacySettings

> ApplicationDataPrivacy GetDefaultApplicationDataPrivacySettings(ctx).Execute()

Gets the data privacy settings of the default web application



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
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.GetDefaultApplicationDataPrivacySettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.GetDefaultApplicationDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultApplicationDataPrivacySettings`: ApplicationDataPrivacy
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.GetDefaultApplicationDataPrivacySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultApplicationDataPrivacySettingsRequest struct via the builder pattern


### Return type

[**ApplicationDataPrivacy**](ApplicationDataPrivacy.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWebApplicationConfig

> WebApplicationConfig GetWebApplicationConfig(ctx, id).Execute()

Gets the configuration of the specified web application

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
    id := "id_example" // string | The ID of the requested web application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.GetWebApplicationConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.GetWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebApplicationConfig`: WebApplicationConfig
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.GetWebApplicationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the requested web application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WebApplicationConfig**](WebApplicationConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDataPrivacySettings

> ApplicationDataPrivacyList ListDataPrivacySettings(ctx).Execute()

Lists data privacy settings of all web applications

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
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.ListDataPrivacySettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.ListDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDataPrivacySettings`: ApplicationDataPrivacyList
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.ListDataPrivacySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDataPrivacySettingsRequest struct via the builder pattern


### Return type

[**ApplicationDataPrivacyList**](ApplicationDataPrivacyList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeyUserActions

> KeyUserActionList ListKeyUserActions(ctx, id).Execute()

Gets the list of key user actions in the specified web application

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
    id := "id_example" // string | The ID of the required web application.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.ListKeyUserActions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.ListKeyUserActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListKeyUserActions`: KeyUserActionList
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.ListKeyUserActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required web application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKeyUserActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KeyUserActionList**](KeyUserActionList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWebApplicationConfigs

> StubList ListWebApplicationConfigs(ctx).Execute()

Lists all existing web applications

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
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.ListWebApplicationConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.ListWebApplicationConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebApplicationConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.ListWebApplicationConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebApplicationConfigsRequest struct via the builder pattern


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


## UpdateApplicationErrorConfig

> UpdateApplicationErrorConfig(ctx, id).ApplicationErrorRules(applicationErrorRules).Execute()

Updates the configuration of error rules in the specified application

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
    id := "id_example" // string | The ID of the required web application.
    applicationErrorRules := *openapiclient.NewApplicationErrorRules([]openapiclient.CustomErrorRule{*openapiclient.NewCustomErrorRule(false, false, false)}, []openapiclient.HttpErrorRule{*openapiclient.NewHttpErrorRule(false, false, false, false, false)}, false, false, false) // ApplicationErrorRules | The JSON body of the request. Contains the updated configuration of error rules. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.UpdateApplicationErrorConfig(context.Background(), id).ApplicationErrorRules(applicationErrorRules).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.UpdateApplicationErrorConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required web application. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationErrorConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationErrorRules** | [**ApplicationErrorRules**](ApplicationErrorRules.md) | The JSON body of the request. Contains the updated configuration of error rules. | 

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


## UpdateDataPrivacySettings

> UpdateDataPrivacySettings(ctx, id).ApplicationDataPrivacy(applicationDataPrivacy).Execute()

Updates the data privacy settings of the specified web application

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
    id := "id_example" // string | The ID of the web application, where you want to update data privacy settings.
    applicationDataPrivacy := *openapiclient.NewApplicationDataPrivacy(false, "DoNotTrackBehaviour_example", false) // ApplicationDataPrivacy | JSON body of the request, containing new data privacy settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.UpdateDataPrivacySettings(context.Background(), id).ApplicationDataPrivacy(applicationDataPrivacy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.UpdateDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application, where you want to update data privacy settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDataPrivacySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDataPrivacy** | [**ApplicationDataPrivacy**](ApplicationDataPrivacy.md) | JSON body of the request, containing new data privacy settings. | 

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


## UpdateDefaultApplicationDataPrivacySettings

> UpdateDefaultApplicationDataPrivacySettings(ctx).ApplicationDataPrivacy(applicationDataPrivacy).Execute()

Updates the data privacy settings of the default web application



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
    applicationDataPrivacy := *openapiclient.NewApplicationDataPrivacy(false, "DoNotTrackBehaviour_example", false) // ApplicationDataPrivacy | JSON body of the request, containing new data privacy settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.UpdateDefaultApplicationDataPrivacySettings(context.Background()).ApplicationDataPrivacy(applicationDataPrivacy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.UpdateDefaultApplicationDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultApplicationDataPrivacySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDataPrivacy** | [**ApplicationDataPrivacy**](ApplicationDataPrivacy.md) | JSON body of the request, containing new data privacy settings. | 

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


## UpdateWebApplicationConfig

> EntityShortRepresentation UpdateWebApplicationConfig(ctx, id).WebApplicationConfig(webApplicationConfig).Execute()

Updates the configuration of the specified web application or creates a new one



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
    id := "id_example" // string | The ID of the web application to update.   If you set the ID in the body as well, it must match this ID.
    webApplicationConfig := *openapiclient.NewWebApplicationConfig(float32(123), *openapiclient.NewApdex(), *openapiclient.NewApdex(), "LoadActionKeyPerformanceMetric_example", *openapiclient.NewMonitoringSettings(*openapiclient.NewAdvancedJavaScriptTagSettings(*openapiclient.NewAdditionalEventHandlers(false, false, false, int32(123), false, false, false), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings("AdditionalEventCapturedAsUserInput_example", false, false, false, false, false, false, false), false, int32(123), int32(123), "SpecialCharactersToEscape_example"), false, *openapiclient.NewContentCapture(false, *openapiclient.NewResourceTimingSettings(false, int32(123), false), *openapiclient.NewTimeoutSettings(int32(123), int32(123), false), false), "CookiePlacementDomain_example", "CustomConfigurationProperties_example", "ExcludeXhrRegex_example", false, "InjectionMode_example", *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), false, "ServerRequestPathId_example", false), "Name_example", false, *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewApdex(), "XhrActionKeyPerformanceMetric_example") // WebApplicationConfig | JSON body of the request, containing updated configuration of the web application. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMWebApplicationConfigurationAPI.UpdateWebApplicationConfig(context.Background(), id).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.UpdateWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWebApplicationConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMWebApplicationConfigurationAPI.UpdateWebApplicationConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application to update.   If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWebApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing updated configuration of the web application. | 

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


## ValidateCreateWebApplicationConfig

> ValidateCreateWebApplicationConfig(ctx).WebApplicationConfig(webApplicationConfig).Execute()

Validates the configuration of the web application for the `POST /applications/web` request

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
    webApplicationConfig := *openapiclient.NewWebApplicationConfig(float32(123), *openapiclient.NewApdex(), *openapiclient.NewApdex(), "LoadActionKeyPerformanceMetric_example", *openapiclient.NewMonitoringSettings(*openapiclient.NewAdvancedJavaScriptTagSettings(*openapiclient.NewAdditionalEventHandlers(false, false, false, int32(123), false, false, false), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings("AdditionalEventCapturedAsUserInput_example", false, false, false, false, false, false, false), false, int32(123), int32(123), "SpecialCharactersToEscape_example"), false, *openapiclient.NewContentCapture(false, *openapiclient.NewResourceTimingSettings(false, int32(123), false), *openapiclient.NewTimeoutSettings(int32(123), int32(123), false), false), "CookiePlacementDomain_example", "CustomConfigurationProperties_example", "ExcludeXhrRegex_example", false, "InjectionMode_example", *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), false, "ServerRequestPathId_example", false), "Name_example", false, *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewApdex(), "XhrActionKeyPerformanceMetric_example") // WebApplicationConfig | JSON body of the request, containing the web application configuration to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.ValidateCreateWebApplicationConfig(context.Background()).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.ValidateCreateWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateWebApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing the web application configuration to validate. | 

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


## ValidateDataPrivacySettings

> ValidateDataPrivacySettings(ctx, id).ApplicationDataPrivacy(applicationDataPrivacy).Execute()

Validates data privacy settings for the `PUT /applications/web/{id}/dataPrivacy` request

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
    id := "id_example" // string | The ID of the web application, where you want to validate data privacy settings.
    applicationDataPrivacy := *openapiclient.NewApplicationDataPrivacy(false, "DoNotTrackBehaviour_example", false) // ApplicationDataPrivacy | JSON body of the request, containing new data privacy settings. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.ValidateDataPrivacySettings(context.Background(), id).ApplicationDataPrivacy(applicationDataPrivacy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.ValidateDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application, where you want to validate data privacy settings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateDataPrivacySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDataPrivacy** | [**ApplicationDataPrivacy**](ApplicationDataPrivacy.md) | JSON body of the request, containing new data privacy settings. | 

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


## ValidateDefaultApplicationDataPrivacySettings

> ValidateDefaultApplicationDataPrivacySettings(ctx).ApplicationDataPrivacy(applicationDataPrivacy).Execute()

Validates data privacy settings of the default web application for the `PUT /applications/web/default/dataPrivacy` request

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
    applicationDataPrivacy := *openapiclient.NewApplicationDataPrivacy(false, "DoNotTrackBehaviour_example", false) // ApplicationDataPrivacy | JSON body of the request, containing the data privacy settings to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.ValidateDefaultApplicationDataPrivacySettings(context.Background()).ApplicationDataPrivacy(applicationDataPrivacy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.ValidateDefaultApplicationDataPrivacySettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDefaultApplicationDataPrivacySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDataPrivacy** | [**ApplicationDataPrivacy**](ApplicationDataPrivacy.md) | JSON body of the request, containing the data privacy settings to validate. | 

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


## ValidateDefaultConfiguration

> ValidateDefaultConfiguration(ctx).WebApplicationConfig(webApplicationConfig).Execute()

Validates the configuration of the default web application for the `PUT /applications/web/default` request

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
    webApplicationConfig := *openapiclient.NewWebApplicationConfig(float32(123), *openapiclient.NewApdex(), *openapiclient.NewApdex(), "LoadActionKeyPerformanceMetric_example", *openapiclient.NewMonitoringSettings(*openapiclient.NewAdvancedJavaScriptTagSettings(*openapiclient.NewAdditionalEventHandlers(false, false, false, int32(123), false, false, false), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings("AdditionalEventCapturedAsUserInput_example", false, false, false, false, false, false, false), false, int32(123), int32(123), "SpecialCharactersToEscape_example"), false, *openapiclient.NewContentCapture(false, *openapiclient.NewResourceTimingSettings(false, int32(123), false), *openapiclient.NewTimeoutSettings(int32(123), int32(123), false), false), "CookiePlacementDomain_example", "CustomConfigurationProperties_example", "ExcludeXhrRegex_example", false, "InjectionMode_example", *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), false, "ServerRequestPathId_example", false), "Name_example", false, *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewApdex(), "XhrActionKeyPerformanceMetric_example") // WebApplicationConfig | JSON body of the request, containing web application configuration to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.ValidateDefaultConfiguration(context.Background()).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.ValidateDefaultConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDefaultConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing web application configuration to validate. | 

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


## ValidateUpdateWebApplicationConfig

> ValidateUpdateWebApplicationConfig(ctx, id).WebApplicationConfig(webApplicationConfig).Execute()

Validates the configuration of the web application for the `PUT /applications/web/{id}` request.

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
    id := "id_example" // string | The ID of the web application to validate.
    webApplicationConfig := *openapiclient.NewWebApplicationConfig(float32(123), *openapiclient.NewApdex(), *openapiclient.NewApdex(), "LoadActionKeyPerformanceMetric_example", *openapiclient.NewMonitoringSettings(*openapiclient.NewAdvancedJavaScriptTagSettings(*openapiclient.NewAdditionalEventHandlers(false, false, false, int32(123), false, false, false), *openapiclient.NewEventWrapperSettings(false, false, false, false, false, false), *openapiclient.NewGlobalEventCaptureSettings("AdditionalEventCapturedAsUserInput_example", false, false, false, false, false, false, false), false, int32(123), int32(123), "SpecialCharactersToEscape_example"), false, *openapiclient.NewContentCapture(false, *openapiclient.NewResourceTimingSettings(false, int32(123), false), *openapiclient.NewTimeoutSettings(int32(123), int32(123), false), false), "CookiePlacementDomain_example", "CustomConfigurationProperties_example", "ExcludeXhrRegex_example", false, "InjectionMode_example", *openapiclient.NewJavaScriptFrameworkSupport(false, false, false, false, false, false, false, false), false, "ServerRequestPathId_example", false), "Name_example", false, *openapiclient.NewWaterfallSettings(int32(123), int32(123), int32(123), int32(123), int32(123), int32(123), int32(123)), *openapiclient.NewApdex(), "XhrActionKeyPerformanceMetric_example") // WebApplicationConfig | JSON body of the request, containing the web application configuration to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMWebApplicationConfigurationAPI.ValidateUpdateWebApplicationConfig(context.Background(), id).WebApplicationConfig(webApplicationConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMWebApplicationConfigurationAPI.ValidateUpdateWebApplicationConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the web application to validate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateWebApplicationConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **webApplicationConfig** | [**WebApplicationConfig**](WebApplicationConfig.md) | JSON body of the request, containing the web application configuration to validate. | 

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

