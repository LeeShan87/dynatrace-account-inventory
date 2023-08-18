# \ServiceDetectionFullWebServiceAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateServiceDetectionRule**](ServiceDetectionFullWebServiceAPI.md#CreateServiceDetectionRule) | **Post** /service/detectionRules/FULL_WEB_SERVICE | Creates a new service detection rule
[**DeleteServiceDetectionRule**](ServiceDetectionFullWebServiceAPI.md#DeleteServiceDetectionRule) | **Delete** /service/detectionRules/FULL_WEB_SERVICE/{id} | Deletes the specified service detection rule
[**GetServiceDetectionRule**](ServiceDetectionFullWebServiceAPI.md#GetServiceDetectionRule) | **Get** /service/detectionRules/FULL_WEB_SERVICE/{id} | Shows the properties of the specified service detection rule
[**ListServiceDetectionRules**](ServiceDetectionFullWebServiceAPI.md#ListServiceDetectionRules) | **Get** /service/detectionRules/FULL_WEB_SERVICE | Lists all full web service detection rules
[**OrderServiceDetectionRules**](ServiceDetectionFullWebServiceAPI.md#OrderServiceDetectionRules) | **Put** /service/detectionRules/FULL_WEB_SERVICE/order | Reorders the service detection rules of the specified type
[**UpdateServiceDetectionRule**](ServiceDetectionFullWebServiceAPI.md#UpdateServiceDetectionRule) | **Put** /service/detectionRules/FULL_WEB_SERVICE/{id} | Updates an existing service detection rule
[**ValidateCreateServiceDetectionRule**](ServiceDetectionFullWebServiceAPI.md#ValidateCreateServiceDetectionRule) | **Post** /service/detectionRules/FULL_WEB_SERVICE/validator | Validates the payload for the &#x60;POST /ruleBasedServiceDetection/FULL_WEB_SERVICE&#x60; request
[**ValidateUpdateServiceDetectionRule**](ServiceDetectionFullWebServiceAPI.md#ValidateUpdateServiceDetectionRule) | **Post** /service/detectionRules/FULL_WEB_SERVICE/{id}/validator | Validate the payload for the &#x60;PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/{id}&#x60; request



## CreateServiceDetectionRule

> EntityShortRepresentation CreateServiceDetectionRule(ctx).Position(position).FullWebServiceRule(fullWebServiceRule).Execute()

Creates a new service detection rule



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
    position := "position_example" // string | The position of the new rule:    * `APPEND`: at the end of the rule list.   * `PREPEND`: on top of the rule list.    (optional) (default to "APPEND")
    fullWebServiceRule := *openapiclient.NewFullWebServiceRule(false, "Name_example", "Type_example") // FullWebServiceRule | The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the `PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/reorder` request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDetectionFullWebServiceAPI.CreateServiceDetectionRule(context.Background()).Position(position).FullWebServiceRule(fullWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceAPI.CreateServiceDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceDetectionRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebServiceAPI.CreateServiceDetectionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **string** | The position of the new rule:    * &#x60;APPEND&#x60;: at the end of the rule list.   * &#x60;PREPEND&#x60;: on top of the rule list.    | [default to &quot;APPEND&quot;]
 **fullWebServiceRule** | [**FullWebServiceRule**](FullWebServiceRule.md) | The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the &#x60;PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/reorder&#x60; request. | 

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


## DeleteServiceDetectionRule

> DeleteServiceDetectionRule(ctx, id).Execute()

Deletes the specified service detection rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the service detection rule to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceDetectionFullWebServiceAPI.DeleteServiceDetectionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceAPI.DeleteServiceDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the service detection rule to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceDetectionRuleRequest struct via the builder pattern


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


## GetServiceDetectionRule

> FullWebServiceRule GetServiceDetectionRule(ctx, id).Execute()

Shows the properties of the specified service detection rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required service detection rule.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDetectionFullWebServiceAPI.GetServiceDetectionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceAPI.GetServiceDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceDetectionRule`: FullWebServiceRule
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebServiceAPI.GetServiceDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required service detection rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullWebServiceRule**](FullWebServiceRule.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceDetectionRules

> StubList ListServiceDetectionRules(ctx).Execute()

Lists all full web service detection rules

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
    resp, r, err := apiClient.ServiceDetectionFullWebServiceAPI.ListServiceDetectionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceAPI.ListServiceDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceDetectionRules`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebServiceAPI.ListServiceDetectionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceDetectionRulesRequest struct via the builder pattern


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


## OrderServiceDetectionRules

> OrderServiceDetectionRules(ctx).StubList(stubList).Execute()

Reorders the service detection rules of the specified type



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
    stubList := *openapiclient.NewStubList([]openapiclient.EntityShortRepresentation{*openapiclient.NewEntityShortRepresentation("Id_example")}) // StubList | The JSON body of the request containing the service detection rules in the required order. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceDetectionFullWebServiceAPI.OrderServiceDetectionRules(context.Background()).StubList(stubList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceAPI.OrderServiceDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderServiceDetectionRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stubList** | [**StubList**](StubList.md) | The JSON body of the request containing the service detection rules in the required order. | 

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


## UpdateServiceDetectionRule

> EntityShortRepresentation UpdateServiceDetectionRule(ctx, id).FullWebServiceRule(fullWebServiceRule).Execute()

Updates an existing service detection rule



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the rule to be updated.
    fullWebServiceRule := *openapiclient.NewFullWebServiceRule(false, "Name_example", "Type_example") // FullWebServiceRule | The JSON body of the request containing updated parameters of the service detection rule. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDetectionFullWebServiceAPI.UpdateServiceDetectionRule(context.Background(), id).FullWebServiceRule(fullWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceAPI.UpdateServiceDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateServiceDetectionRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebServiceAPI.UpdateServiceDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServiceDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebServiceRule** | [**FullWebServiceRule**](FullWebServiceRule.md) | The JSON body of the request containing updated parameters of the service detection rule. | 

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


## ValidateCreateServiceDetectionRule

> ValidateCreateServiceDetectionRule(ctx).FullWebServiceRule(fullWebServiceRule).Execute()

Validates the payload for the `POST /ruleBasedServiceDetection/FULL_WEB_SERVICE` request

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
    fullWebServiceRule := *openapiclient.NewFullWebServiceRule(false, "Name_example", "Type_example") // FullWebServiceRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceDetectionFullWebServiceAPI.ValidateCreateServiceDetectionRule(context.Background()).FullWebServiceRule(fullWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceAPI.ValidateCreateServiceDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateServiceDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullWebServiceRule** | [**FullWebServiceRule**](FullWebServiceRule.md) |  | 

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


## ValidateUpdateServiceDetectionRule

> ValidateUpdateServiceDetectionRule(ctx, id).FullWebServiceRule(fullWebServiceRule).Execute()

Validate the payload for the `PUT /ruleBasedServiceDetection/FULL_WEB_SERVICE/{id}` request

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the service detection rule to be validated.
    fullWebServiceRule := *openapiclient.NewFullWebServiceRule(false, "Name_example", "Type_example") // FullWebServiceRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceDetectionFullWebServiceAPI.ValidateUpdateServiceDetectionRule(context.Background(), id).FullWebServiceRule(fullWebServiceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebServiceAPI.ValidateUpdateServiceDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the service detection rule to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateServiceDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebServiceRule** | [**FullWebServiceRule**](FullWebServiceRule.md) |  | 

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

