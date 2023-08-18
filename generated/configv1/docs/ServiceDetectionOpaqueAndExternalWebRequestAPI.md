# \ServiceDetectionOpaqueAndExternalWebRequestAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOpaqueAndExternalWebRequestDetectionRule**](ServiceDetectionOpaqueAndExternalWebRequestAPI.md#CreateOpaqueAndExternalWebRequestDetectionRule) | **Post** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST | Creates a new service detection rule
[**DeleteOpaqueAndExternalWebRequestDetectionRule**](ServiceDetectionOpaqueAndExternalWebRequestAPI.md#DeleteOpaqueAndExternalWebRequestDetectionRule) | **Delete** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id} | Deletes the specified service detection rule
[**GetOpaqueAndExternalWebRequestDetectionRule**](ServiceDetectionOpaqueAndExternalWebRequestAPI.md#GetOpaqueAndExternalWebRequestDetectionRule) | **Get** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id} | Shows the properties of the specified service detection rule
[**ListOpaqueAndExternalWebRequestDetectionRules**](ServiceDetectionOpaqueAndExternalWebRequestAPI.md#ListOpaqueAndExternalWebRequestDetectionRules) | **Get** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST | Lists all opaque and external web request service detection rules
[**OrderOpaqueAndExternalWebRequestDetectionRules**](ServiceDetectionOpaqueAndExternalWebRequestAPI.md#OrderOpaqueAndExternalWebRequestDetectionRules) | **Put** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/order | Reorders the service detection rules of the specified type
[**UpdateOpaqueAndExternalWebRequestDetectionRule**](ServiceDetectionOpaqueAndExternalWebRequestAPI.md#UpdateOpaqueAndExternalWebRequestDetectionRule) | **Put** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id} | Updates an existing service detection rule
[**ValidateCreateOpaqueAndExternalWebRequestDetectionRule**](ServiceDetectionOpaqueAndExternalWebRequestAPI.md#ValidateCreateOpaqueAndExternalWebRequestDetectionRule) | **Post** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/validator | Validates the payload for the &#x60;POST /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST&#x60; request
[**ValidateUpdateOpaqueAndExternalWebRequestDetectionRule**](ServiceDetectionOpaqueAndExternalWebRequestAPI.md#ValidateUpdateOpaqueAndExternalWebRequestDetectionRule) | **Post** /service/detectionRules/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}/validator | Validate the payload for the &#x60;PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}&#x60; request



## CreateOpaqueAndExternalWebRequestDetectionRule

> EntityShortRepresentation CreateOpaqueAndExternalWebRequestDetectionRule(ctx).Position(position).OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule).Execute()

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
    opaqueAndExternalWebRequestRule := *openapiclient.NewOpaqueAndExternalWebRequestRule(false, "Name_example", "Type_example") // OpaqueAndExternalWebRequestRule | The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the `PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST/reorder` request. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDetectionOpaqueAndExternalWebRequestAPI.CreateOpaqueAndExternalWebRequestDetectionRule(context.Background()).Position(position).OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebRequestAPI.CreateOpaqueAndExternalWebRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOpaqueAndExternalWebRequestDetectionRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionOpaqueAndExternalWebRequestAPI.CreateOpaqueAndExternalWebRequestDetectionRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOpaqueAndExternalWebRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **string** | The position of the new rule:    * &#x60;APPEND&#x60;: at the end of the rule list.   * &#x60;PREPEND&#x60;: on top of the rule list.    | [default to &quot;APPEND&quot;]
 **opaqueAndExternalWebRequestRule** | [**OpaqueAndExternalWebRequestRule**](OpaqueAndExternalWebRequestRule.md) | The JSON body of the request containing parameters of the new service detection rule.    You must not specify the ID of the rule!   The **order** field is ignored in this request. To enforce a particular order use the &#x60;PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST/reorder&#x60; request. | 

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


## DeleteOpaqueAndExternalWebRequestDetectionRule

> DeleteOpaqueAndExternalWebRequestDetectionRule(ctx, id).Execute()

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
    r, err := apiClient.ServiceDetectionOpaqueAndExternalWebRequestAPI.DeleteOpaqueAndExternalWebRequestDetectionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebRequestAPI.DeleteOpaqueAndExternalWebRequestDetectionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteOpaqueAndExternalWebRequestDetectionRuleRequest struct via the builder pattern


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


## GetOpaqueAndExternalWebRequestDetectionRule

> OpaqueAndExternalWebRequestRule GetOpaqueAndExternalWebRequestDetectionRule(ctx, id).Execute()

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
    resp, r, err := apiClient.ServiceDetectionOpaqueAndExternalWebRequestAPI.GetOpaqueAndExternalWebRequestDetectionRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebRequestAPI.GetOpaqueAndExternalWebRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOpaqueAndExternalWebRequestDetectionRule`: OpaqueAndExternalWebRequestRule
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionOpaqueAndExternalWebRequestAPI.GetOpaqueAndExternalWebRequestDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required service detection rule. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpaqueAndExternalWebRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**OpaqueAndExternalWebRequestRule**](OpaqueAndExternalWebRequestRule.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOpaqueAndExternalWebRequestDetectionRules

> StubList ListOpaqueAndExternalWebRequestDetectionRules(ctx).Execute()

Lists all opaque and external web request service detection rules

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
    resp, r, err := apiClient.ServiceDetectionOpaqueAndExternalWebRequestAPI.ListOpaqueAndExternalWebRequestDetectionRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebRequestAPI.ListOpaqueAndExternalWebRequestDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOpaqueAndExternalWebRequestDetectionRules`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionOpaqueAndExternalWebRequestAPI.ListOpaqueAndExternalWebRequestDetectionRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOpaqueAndExternalWebRequestDetectionRulesRequest struct via the builder pattern


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


## OrderOpaqueAndExternalWebRequestDetectionRules

> OrderOpaqueAndExternalWebRequestDetectionRules(ctx).StubList(stubList).Execute()

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
    r, err := apiClient.ServiceDetectionOpaqueAndExternalWebRequestAPI.OrderOpaqueAndExternalWebRequestDetectionRules(context.Background()).StubList(stubList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebRequestAPI.OrderOpaqueAndExternalWebRequestDetectionRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderOpaqueAndExternalWebRequestDetectionRulesRequest struct via the builder pattern


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


## UpdateOpaqueAndExternalWebRequestDetectionRule

> EntityShortRepresentation UpdateOpaqueAndExternalWebRequestDetectionRule(ctx, id).OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule).Execute()

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
    opaqueAndExternalWebRequestRule := *openapiclient.NewOpaqueAndExternalWebRequestRule(false, "Name_example", "Type_example") // OpaqueAndExternalWebRequestRule | The JSON body of the request containing updated parameters of the service detection rule. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceDetectionOpaqueAndExternalWebRequestAPI.UpdateOpaqueAndExternalWebRequestDetectionRule(context.Background(), id).OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebRequestAPI.UpdateOpaqueAndExternalWebRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOpaqueAndExternalWebRequestDetectionRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceDetectionOpaqueAndExternalWebRequestAPI.UpdateOpaqueAndExternalWebRequestDetectionRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the rule to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOpaqueAndExternalWebRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaqueAndExternalWebRequestRule** | [**OpaqueAndExternalWebRequestRule**](OpaqueAndExternalWebRequestRule.md) | The JSON body of the request containing updated parameters of the service detection rule. | 

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


## ValidateCreateOpaqueAndExternalWebRequestDetectionRule

> ValidateCreateOpaqueAndExternalWebRequestDetectionRule(ctx).OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule).Execute()

Validates the payload for the `POST /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST` request

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
    opaqueAndExternalWebRequestRule := *openapiclient.NewOpaqueAndExternalWebRequestRule(false, "Name_example", "Type_example") // OpaqueAndExternalWebRequestRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceDetectionOpaqueAndExternalWebRequestAPI.ValidateCreateOpaqueAndExternalWebRequestDetectionRule(context.Background()).OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebRequestAPI.ValidateCreateOpaqueAndExternalWebRequestDetectionRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateOpaqueAndExternalWebRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **opaqueAndExternalWebRequestRule** | [**OpaqueAndExternalWebRequestRule**](OpaqueAndExternalWebRequestRule.md) |  | 

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


## ValidateUpdateOpaqueAndExternalWebRequestDetectionRule

> ValidateUpdateOpaqueAndExternalWebRequestDetectionRule(ctx, id).OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule).Execute()

Validate the payload for the `PUT /ruleBasedServiceDetection/OPAQUE_AND_EXTERNAL_WEB_REQUEST/{id}` request

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
    opaqueAndExternalWebRequestRule := *openapiclient.NewOpaqueAndExternalWebRequestRule(false, "Name_example", "Type_example") // OpaqueAndExternalWebRequestRule | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceDetectionOpaqueAndExternalWebRequestAPI.ValidateUpdateOpaqueAndExternalWebRequestDetectionRule(context.Background(), id).OpaqueAndExternalWebRequestRule(opaqueAndExternalWebRequestRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceDetectionOpaqueAndExternalWebRequestAPI.ValidateUpdateOpaqueAndExternalWebRequestDetectionRule``: %v\n", err)
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

Other parameters are passed through a pointer to a apiValidateUpdateOpaqueAndExternalWebRequestDetectionRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **opaqueAndExternalWebRequestRule** | [**OpaqueAndExternalWebRequestRule**](OpaqueAndExternalWebRequestRule.md) |  | 

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

