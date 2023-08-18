# \ServiceDetectionFullWebRequestAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequestDetectionRule**](ServiceDetectionFullWebRequestAPI.md#CreateRequestDetectionRule) | **Post** /service/detectionRules/FULL_WEB_REQUEST | Creates a new service detection rule
[**DeleteRequestDetectionRule**](ServiceDetectionFullWebRequestAPI.md#DeleteRequestDetectionRule) | **Delete** /service/detectionRules/FULL_WEB_REQUEST/{id} | Deletes the specified service detection rule
[**GetRequestDetectionRule**](ServiceDetectionFullWebRequestAPI.md#GetRequestDetectionRule) | **Get** /service/detectionRules/FULL_WEB_REQUEST/{id} | Gets the properties of the specified service detection rule
[**ListRequestDetectionRules**](ServiceDetectionFullWebRequestAPI.md#ListRequestDetectionRules) | **Get** /service/detectionRules/FULL_WEB_REQUEST | Lists all full web request service detection rules
[**OrderRequestDetectionRules**](ServiceDetectionFullWebRequestAPI.md#OrderRequestDetectionRules) | **Put** /service/detectionRules/FULL_WEB_REQUEST/order | Reorders the service detection rules of the specified type
[**UpdateRequestDetectionRule**](ServiceDetectionFullWebRequestAPI.md#UpdateRequestDetectionRule) | **Put** /service/detectionRules/FULL_WEB_REQUEST/{id} | Updates an existing service detection rule
[**ValidateCreateRequestDetectionRule**](ServiceDetectionFullWebRequestAPI.md#ValidateCreateRequestDetectionRule) | **Post** /service/detectionRules/FULL_WEB_REQUEST/validator | Validates the payload for the &#x60;POST /ruleBasedServiceDetection/FULL_WEB_REQUEST&#x60; request
[**ValidateUpdateRequestDetectionRule**](ServiceDetectionFullWebRequestAPI.md#ValidateUpdateRequestDetectionRule) | **Post** /service/detectionRules/FULL_WEB_REQUEST/{id}/validator | Validates the payload for the &#x60;PUT /service/detectionRules/FULL_WEB_REQUEST/{id}&#x60; request



## CreateRequestDetectionRule

> EntityShortRepresentation CreateRequestDetectionRule(ctx).Position(position).FullWebRequestRule(fullWebRequestRule).Execute()

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
    position := "position_example" // string | The position of the new rule:    * `APPEND`: at the bottom of the rule list.   * `PREPEND`: at the top of the rule list.    If not set, the `APPEND` is used. (optional) (default to "APPEND")
    fullWebRequestRule := *openapiclient.NewFullWebRequestRule(false, "Name_example", "Type_example") // FullWebRequestRule | The JSON body of the request. Contains parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order, use the `PUT /service/detectionRules/FULL_WEB_REQUEST/reorder` request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDetectionFullWebRequestAPI.CreateRequestDetectionRule(context.Background()).Position(position).FullWebRequestRule(fullWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestAPI.CreateRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequestDetectionRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebRequestAPI.CreateRequestDetectionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **string** | The position of the new rule:    * &#x60;APPEND&#x60;: at the bottom of the rule list.   * &#x60;PREPEND&#x60;: at the top of the rule list.    If not set, the &#x60;APPEND&#x60; is used. | [default to &quot;APPEND&quot;]
 **fullWebRequestRule** | [**FullWebRequestRule**](FullWebRequestRule.md) | The JSON body of the request. Contains parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order, use the &#x60;PUT /service/detectionRules/FULL_WEB_REQUEST/reorder&#x60; request. | 

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


## DeleteRequestDetectionRule

> DeleteRequestDetectionRule(ctx, id).Execute()

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
    r, err := apiClient.ServiceDetectionFullWebRequestAPI.DeleteRequestDetectionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestAPI.DeleteRequestDetectionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRequestDetectionRuleRequest struct via the builder pattern


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


## GetRequestDetectionRule

> FullWebRequestRule GetRequestDetectionRule(ctx, id).Execute()

Gets the properties of the specified service detection rule

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
    resp, r, err := apiClient.ServiceDetectionFullWebRequestAPI.GetRequestDetectionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestAPI.GetRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestDetectionRule`: FullWebRequestRule
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebRequestAPI.GetRequestDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required service detection rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FullWebRequestRule**](FullWebRequestRule.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestDetectionRules

> StubList ListRequestDetectionRules(ctx).Execute()

Lists all full web request service detection rules

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
    resp, r, err := apiClient.ServiceDetectionFullWebRequestAPI.ListRequestDetectionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestAPI.ListRequestDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestDetectionRules`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebRequestAPI.ListRequestDetectionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequestDetectionRulesRequest struct via the builder pattern


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


## OrderRequestDetectionRules

> OrderRequestDetectionRules(ctx).StubList(stubList).Execute()

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
    r, err := apiClient.ServiceDetectionFullWebRequestAPI.OrderRequestDetectionRules(context.Background()).StubList(stubList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestAPI.OrderRequestDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderRequestDetectionRulesRequest struct via the builder pattern


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


## UpdateRequestDetectionRule

> EntityShortRepresentation UpdateRequestDetectionRule(ctx, id).FullWebRequestRule(fullWebRequestRule).Execute()

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
    fullWebRequestRule := *openapiclient.NewFullWebRequestRule(false, "Name_example", "Type_example") // FullWebRequestRule | The JSON body of the request. Contains updated parameters of the service detection rule. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDetectionFullWebRequestAPI.UpdateRequestDetectionRule(context.Background(), id).FullWebRequestRule(fullWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestAPI.UpdateRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRequestDetectionRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionFullWebRequestAPI.UpdateRequestDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebRequestRule** | [**FullWebRequestRule**](FullWebRequestRule.md) | The JSON body of the request. Contains updated parameters of the service detection rule. | 

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


## ValidateCreateRequestDetectionRule

> ValidateCreateRequestDetectionRule(ctx).FullWebRequestRule(fullWebRequestRule).Execute()

Validates the payload for the `POST /ruleBasedServiceDetection/FULL_WEB_REQUEST` request

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
    fullWebRequestRule := *openapiclient.NewFullWebRequestRule(false, "Name_example", "Type_example") // FullWebRequestRule | The JSON body of the request. Contains parameters of the service detection rule to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceDetectionFullWebRequestAPI.ValidateCreateRequestDetectionRule(context.Background()).FullWebRequestRule(fullWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestAPI.ValidateCreateRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fullWebRequestRule** | [**FullWebRequestRule**](FullWebRequestRule.md) | The JSON body of the request. Contains parameters of the service detection rule to be validated. | 

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


## ValidateUpdateRequestDetectionRule

> ValidateUpdateRequestDetectionRule(ctx, id).FullWebRequestRule(fullWebRequestRule).Execute()

Validates the payload for the `PUT /service/detectionRules/FULL_WEB_REQUEST/{id}` request

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
    fullWebRequestRule := *openapiclient.NewFullWebRequestRule(false, "Name_example", "Type_example") // FullWebRequestRule | The JSON body of the request. Contains parameters of the service detection rule to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceDetectionFullWebRequestAPI.ValidateUpdateRequestDetectionRule(context.Background(), id).FullWebRequestRule(fullWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionFullWebRequestAPI.ValidateUpdateRequestDetectionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiValidateUpdateRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fullWebRequestRule** | [**FullWebRequestRule**](FullWebRequestRule.md) | The JSON body of the request. Contains parameters of the service detection rule to be validated. | 

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

