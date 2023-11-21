# \RUMApplicationDetectionRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationDetectionConfig**](RUMApplicationDetectionRulesAPI.md#CreateApplicationDetectionConfig) | **Post** /applicationDetectionRules | Creates a new application detection rule
[**DeleteApplicationDetectionConfig**](RUMApplicationDetectionRulesAPI.md#DeleteApplicationDetectionConfig) | **Delete** /applicationDetectionRules/{id} | Deletes the specified application detection rule
[**GetApplicationDetectionConfig**](RUMApplicationDetectionRulesAPI.md#GetApplicationDetectionConfig) | **Get** /applicationDetectionRules/{id} | Gets the parameters of the specified application detection rule
[**ListApplicationDetectionConfigs**](RUMApplicationDetectionRulesAPI.md#ListApplicationDetectionConfigs) | **Get** /applicationDetectionRules | Lists all available application detection rules
[**OrderApplicationDetectionConfigs**](RUMApplicationDetectionRulesAPI.md#OrderApplicationDetectionConfigs) | **Put** /applicationDetectionRules/order | Reorders the application detection rules
[**UpdateApplicationDetectionConfig**](RUMApplicationDetectionRulesAPI.md#UpdateApplicationDetectionConfig) | **Put** /applicationDetectionRules/{id} | Updates the specified application detection rule
[**ValidateCreateApplicationDetectionConfig**](RUMApplicationDetectionRulesAPI.md#ValidateCreateApplicationDetectionConfig) | **Post** /applicationDetectionRules/validator | Validates the payload for the &#x60;POST /applicationDetection&#x60; request
[**ValidateUpdateApplicationDetectionConfig**](RUMApplicationDetectionRulesAPI.md#ValidateUpdateApplicationDetectionConfig) | **Post** /applicationDetectionRules/{id}/validator | Validate the payload for the &#x60;PUT /applicationDetection/{id}&#x60; request



## CreateApplicationDetectionConfig

> EntityShortRepresentation CreateApplicationDetectionConfig(ctx).Position(position).ApplicationDetectionRuleConfig(applicationDetectionRuleConfig).Execute()

Creates a new application detection rule



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
    position := "position_example" // string | The position of the new rule:    * `APPEND`: at the bottom of the rule list.   * `PREPEND`: at the top of the rule list.    If not set, the `APPEND` is used. (optional)
    applicationDetectionRuleConfig := *openapiclient.NewApplicationDetectionRuleConfig("APPLICATION-123456", *openapiclient.NewApplicationFilter("DOMAIN", "EQUALS", "myapp.example.com")) // ApplicationDetectionRuleConfig | The JSON body of the request. Contains configuration of the new application detection rule.    You must not specify the ID of the rule.   The **order** field is ignored in this request. To enforce a particular order use the `PUT /applicationDetectionRules/order` request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMApplicationDetectionRulesAPI.CreateApplicationDetectionConfig(context.Background()).Position(position).ApplicationDetectionRuleConfig(applicationDetectionRuleConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesAPI.CreateApplicationDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApplicationDetectionConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMApplicationDetectionRulesAPI.CreateApplicationDetectionConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **string** | The position of the new rule:    * &#x60;APPEND&#x60;: at the bottom of the rule list.   * &#x60;PREPEND&#x60;: at the top of the rule list.    If not set, the &#x60;APPEND&#x60; is used. | 
 **applicationDetectionRuleConfig** | [**ApplicationDetectionRuleConfig**](ApplicationDetectionRuleConfig.md) | The JSON body of the request. Contains configuration of the new application detection rule.    You must not specify the ID of the rule.   The **order** field is ignored in this request. To enforce a particular order use the &#x60;PUT /applicationDetectionRules/order&#x60; request. | 

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


## DeleteApplicationDetectionConfig

> DeleteApplicationDetectionConfig(ctx, id).Execute()

Deletes the specified application detection rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the application detection rule to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMApplicationDetectionRulesAPI.DeleteApplicationDetectionConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesAPI.DeleteApplicationDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the application detection rule to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationDetectionConfigRequest struct via the builder pattern


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


## GetApplicationDetectionConfig

> ApplicationDetectionRuleConfig GetApplicationDetectionConfig(ctx, id).Execute()

Gets the parameters of the specified application detection rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required application detection rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMApplicationDetectionRulesAPI.GetApplicationDetectionConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesAPI.GetApplicationDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationDetectionConfig`: ApplicationDetectionRuleConfig
    fmt.Fprintf(os.Stdout, "Response from `RUMApplicationDetectionRulesAPI.GetApplicationDetectionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required application detection rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationDetectionRuleConfig**](ApplicationDetectionRuleConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationDetectionConfigs

> StubList ListApplicationDetectionConfigs(ctx).Execute()

Lists all available application detection rules

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
    resp, r, err := apiClient.RUMApplicationDetectionRulesAPI.ListApplicationDetectionConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesAPI.ListApplicationDetectionConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationDetectionConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `RUMApplicationDetectionRulesAPI.ListApplicationDetectionConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationDetectionConfigsRequest struct via the builder pattern


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


## OrderApplicationDetectionConfigs

> OrderApplicationDetectionConfigs(ctx).StubList(stubList).Execute()

Reorders the application detection rules



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
    stubList := *openapiclient.NewStubList([]openapiclient.EntityShortRepresentation{*openapiclient.NewEntityShortRepresentation("Id_example")}) // StubList | The JSON body of the request. Contains the IDs of the application detection rules in the desired order. Any further properties (**name**, **description**) are ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMApplicationDetectionRulesAPI.OrderApplicationDetectionConfigs(context.Background()).StubList(stubList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesAPI.OrderApplicationDetectionConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderApplicationDetectionConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stubList** | [**StubList**](StubList.md) | The JSON body of the request. Contains the IDs of the application detection rules in the desired order. Any further properties (**name**, **description**) are ignored. | 

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


## UpdateApplicationDetectionConfig

> EntityShortRepresentation UpdateApplicationDetectionConfig(ctx, id).ApplicationDetectionRuleConfig(applicationDetectionRuleConfig).Execute()

Updates the specified application detection rule



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the application detection rule to be updated.    If you set the ID in the body as well, it must match this ID.
    applicationDetectionRuleConfig := *openapiclient.NewApplicationDetectionRuleConfig("APPLICATION-123456", *openapiclient.NewApplicationFilter("DOMAIN", "EQUALS", "myapp.example.com")) // ApplicationDetectionRuleConfig | The JSON body of the request. Contains updated parameters of the application detection rule.    If the **order** parameter is set, the rule is placed to this position. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMApplicationDetectionRulesAPI.UpdateApplicationDetectionConfig(context.Background(), id).ApplicationDetectionRuleConfig(applicationDetectionRuleConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesAPI.UpdateApplicationDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApplicationDetectionConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `RUMApplicationDetectionRulesAPI.UpdateApplicationDetectionConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the application detection rule to be updated.    If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDetectionRuleConfig** | [**ApplicationDetectionRuleConfig**](ApplicationDetectionRuleConfig.md) | The JSON body of the request. Contains updated parameters of the application detection rule.    If the **order** parameter is set, the rule is placed to this position. | 

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


## ValidateCreateApplicationDetectionConfig

> ValidateCreateApplicationDetectionConfig(ctx).ApplicationDetectionRuleConfig(applicationDetectionRuleConfig).Execute()

Validates the payload for the `POST /applicationDetection` request

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
    applicationDetectionRuleConfig := *openapiclient.NewApplicationDetectionRuleConfig("APPLICATION-123456", *openapiclient.NewApplicationFilter("DOMAIN", "EQUALS", "myapp.example.com")) // ApplicationDetectionRuleConfig | The JSON body of the request. Contains the application detection rule to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMApplicationDetectionRulesAPI.ValidateCreateApplicationDetectionConfig(context.Background()).ApplicationDetectionRuleConfig(applicationDetectionRuleConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesAPI.ValidateCreateApplicationDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateApplicationDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationDetectionRuleConfig** | [**ApplicationDetectionRuleConfig**](ApplicationDetectionRuleConfig.md) | The JSON body of the request. Contains the application detection rule to be validated. | 

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


## ValidateUpdateApplicationDetectionConfig

> ValidateUpdateApplicationDetectionConfig(ctx, id).ApplicationDetectionRuleConfig(applicationDetectionRuleConfig).Execute()

Validate the payload for the `PUT /applicationDetection/{id}` request

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the application detection rule to be validated.    If you set the ID in the body as well, it must match this ID.
    applicationDetectionRuleConfig := *openapiclient.NewApplicationDetectionRuleConfig("APPLICATION-123456", *openapiclient.NewApplicationFilter("DOMAIN", "EQUALS", "myapp.example.com")) // ApplicationDetectionRuleConfig | The JSON body of the request. Contains the application detection rule to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMApplicationDetectionRulesAPI.ValidateUpdateApplicationDetectionConfig(context.Background(), id).ApplicationDetectionRuleConfig(applicationDetectionRuleConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMApplicationDetectionRulesAPI.ValidateUpdateApplicationDetectionConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the application detection rule to be validated.    If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateApplicationDetectionConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **applicationDetectionRuleConfig** | [**ApplicationDetectionRuleConfig**](ApplicationDetectionRuleConfig.md) | The JSON body of the request. Contains the application detection rule to be validated. | 

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

