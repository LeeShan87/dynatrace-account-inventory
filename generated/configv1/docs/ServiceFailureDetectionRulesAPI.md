# \ServiceFailureDetectionRulesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeRuleId**](ServiceFailureDetectionRulesAPI.md#ChangeRuleId) | **Put** /service/failureDetection/parameterSelection/rules/{id}/changeId | Changes the ID of the specified rule | maturity&#x3D;EARLY_ADOPTER
[**CreateOrUpdateRule**](ServiceFailureDetectionRulesAPI.md#CreateOrUpdateRule) | **Put** /service/failureDetection/parameterSelection/rules/{id} | Updates the specified failure detection rule rule | maturity&#x3D;EARLY_ADOPTER
[**CreateRule**](ServiceFailureDetectionRulesAPI.md#CreateRule) | **Post** /service/failureDetection/parameterSelection/rules | Creates a new failure detection rule | maturity&#x3D;EARLY_ADOPTER
[**DeleteRule**](ServiceFailureDetectionRulesAPI.md#DeleteRule) | **Delete** /service/failureDetection/parameterSelection/rules/{id} | Deletes the specified failure detection rule | maturity&#x3D;EARLY_ADOPTER
[**GetAllRules**](ServiceFailureDetectionRulesAPI.md#GetAllRules) | **Get** /service/failureDetection/parameterSelection/rules | Lists all available failure detection rules | maturity&#x3D;EARLY_ADOPTER
[**GetRule**](ServiceFailureDetectionRulesAPI.md#GetRule) | **Get** /service/failureDetection/parameterSelection/rules/{id} | Gets the properties of the specified rule | maturity&#x3D;EARLY_ADOPTER
[**ReorderRules**](ServiceFailureDetectionRulesAPI.md#ReorderRules) | **Put** /service/failureDetection/parameterSelection/rules/reorderRules | Reorders failure detection rules | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateRule**](ServiceFailureDetectionRulesAPI.md#ValidateCreateRule) | **Post** /service/failureDetection/parameterSelection/rules/validator | Validates the payload for the &#x60;POST /service/failureDetection/parameterSelection/rules&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateRule**](ServiceFailureDetectionRulesAPI.md#ValidateUpdateRule) | **Post** /service/failureDetection/parameterSelection/rules/{id}/validator | Validates the payload for the &#x60;PUT /service/failureDetection/parameterSelection/rules/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## ChangeRuleId

> ChangeRuleId(ctx, id).EntityShortRepresentation(entityShortRepresentation).Execute()

Changes the ID of the specified rule | maturity=EARLY_ADOPTER

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
    id := "123e4567-e89b-12d3-a456-426614174000" // string | The ID of the required failure detection rule. Needs to be a valid RFC 4122 UUID.
    entityShortRepresentation := *openapiclient.NewEntityShortRepresentation("Id_example") // EntityShortRepresentation | The JSON body of the request. Contains the new ID of the rule. All other fields are ignored.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceFailureDetectionRulesAPI.ChangeRuleId(context.Background(), id).EntityShortRepresentation(entityShortRepresentation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionRulesAPI.ChangeRuleId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required failure detection rule. Needs to be a valid RFC 4122 UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityShortRepresentation** | [**EntityShortRepresentation**](EntityShortRepresentation.md) | The JSON body of the request. Contains the new ID of the rule. All other fields are ignored. | 

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


## CreateOrUpdateRule

> EntityShortRepresentation CreateOrUpdateRule(ctx, id).FailureDetectionRule(failureDetectionRule).Execute()

Updates the specified failure detection rule rule | maturity=EARLY_ADOPTER



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
    id := "123e4567-e89b-12d3-a456-426614174000" // string | The ID of the required failure detection rule. Needs to be a valid RFC 4122 UUID.
    failureDetectionRule := *openapiclient.NewFailureDetectionRule([]openapiclient.FailureDetectionCondition{*openapiclient.NewFailureDetectionCondition()}, false, "FdpId_example") // FailureDetectionRule | The JSON body of the request. Contains the updated configuration of the failure detection rule.   You can't update the ID with this request. Use the change ID request instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceFailureDetectionRulesAPI.CreateOrUpdateRule(context.Background(), id).FailureDetectionRule(failureDetectionRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionRulesAPI.CreateOrUpdateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceFailureDetectionRulesAPI.CreateOrUpdateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required failure detection rule. Needs to be a valid RFC 4122 UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **failureDetectionRule** | [**FailureDetectionRule**](FailureDetectionRule.md) | The JSON body of the request. Contains the updated configuration of the failure detection rule.   You can&#39;t update the ID with this request. Use the change ID request instead. | 

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


## CreateRule

> EntityShortRepresentation CreateRule(ctx).FailureDetectionRule(failureDetectionRule).Execute()

Creates a new failure detection rule | maturity=EARLY_ADOPTER



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
    failureDetectionRule := *openapiclient.NewFailureDetectionRule([]openapiclient.FailureDetectionCondition{*openapiclient.NewFailureDetectionCondition()}, false, "FdpId_example") // FailureDetectionRule | The JSON body of the request. Contains the configuration of the new failure detection rule.   Dynatrace will generate a random UUID for you, if you don't specify an ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceFailureDetectionRulesAPI.CreateRule(context.Background()).FailureDetectionRule(failureDetectionRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionRulesAPI.CreateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRule`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceFailureDetectionRulesAPI.CreateRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failureDetectionRule** | [**FailureDetectionRule**](FailureDetectionRule.md) | The JSON body of the request. Contains the configuration of the new failure detection rule.   Dynatrace will generate a random UUID for you, if you don&#39;t specify an ID. | 

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


## DeleteRule

> DeleteRule(ctx, id).Execute()

Deletes the specified failure detection rule | maturity=EARLY_ADOPTER

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
    id := "123e4567-e89b-12d3-a456-426614174000" // string | The ID of the required failure detection rule. Needs to be a valid RFC 4122 UUID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceFailureDetectionRulesAPI.DeleteRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionRulesAPI.DeleteRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required failure detection rule. Needs to be a valid RFC 4122 UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


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


## GetAllRules

> StubList GetAllRules(ctx).Execute()

Lists all available failure detection rules | maturity=EARLY_ADOPTER

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
    resp, r, err := apiClient.ServiceFailureDetectionRulesAPI.GetAllRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionRulesAPI.GetAllRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRules`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceFailureDetectionRulesAPI.GetAllRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRulesRequest struct via the builder pattern


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


## GetRule

> FailureDetectionRule GetRule(ctx, id).Execute()

Gets the properties of the specified rule | maturity=EARLY_ADOPTER

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
    id := "123e4567-e89b-12d3-a456-426614174000" // string | The ID of the required failure detection rule. Needs to be a valid RFC 4122 UUID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceFailureDetectionRulesAPI.GetRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionRulesAPI.GetRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRule`: FailureDetectionRule
    fmt.Fprintf(os.Stdout, "Response from `ServiceFailureDetectionRulesAPI.GetRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required failure detection rule. Needs to be a valid RFC 4122 UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FailureDetectionRule**](FailureDetectionRule.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReorderRules

> ReorderRules(ctx).FdpSelectionRuleOrder(fdpSelectionRuleOrder).Execute()

Reorders failure detection rules | maturity=EARLY_ADOPTER



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
    fdpSelectionRuleOrder := *openapiclient.NewFdpSelectionRuleOrder([]string{"RuleIds_example"}) // FdpSelectionRuleOrder | The JSON body of the request. Contains the failure detection rules in the required order.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceFailureDetectionRulesAPI.ReorderRules(context.Background()).FdpSelectionRuleOrder(fdpSelectionRuleOrder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionRulesAPI.ReorderRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReorderRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fdpSelectionRuleOrder** | [**FdpSelectionRuleOrder**](FdpSelectionRuleOrder.md) | The JSON body of the request. Contains the failure detection rules in the required order. | 

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


## ValidateCreateRule

> ValidateCreateRule(ctx).FailureDetectionRule(failureDetectionRule).Execute()

Validates the payload for the `POST /service/failureDetection/parameterSelection/rules` request | maturity=EARLY_ADOPTER

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
    failureDetectionRule := *openapiclient.NewFailureDetectionRule([]openapiclient.FailureDetectionCondition{*openapiclient.NewFailureDetectionCondition()}, false, "FdpId_example") // FailureDetectionRule | The JSON body of the request. Contains the parameters of the failure detection rule to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceFailureDetectionRulesAPI.ValidateCreateRule(context.Background()).FailureDetectionRule(failureDetectionRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionRulesAPI.ValidateCreateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failureDetectionRule** | [**FailureDetectionRule**](FailureDetectionRule.md) | The JSON body of the request. Contains the parameters of the failure detection rule to be validated. | 

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


## ValidateUpdateRule

> ValidateUpdateRule(ctx, id).FailureDetectionRule(failureDetectionRule).Execute()

Validates the payload for the `PUT /service/failureDetection/parameterSelection/rules/{id}` request | maturity=EARLY_ADOPTER

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
    id := "123e4567-e89b-12d3-a456-426614174000" // string | The ID of the required failure detection rule. Needs to be a valid RFC 4122 UUID.
    failureDetectionRule := *openapiclient.NewFailureDetectionRule([]openapiclient.FailureDetectionCondition{*openapiclient.NewFailureDetectionCondition()}, false, "FdpId_example") // FailureDetectionRule | The JSON body of the request. Contains the configuration of the failure detection rule to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceFailureDetectionRulesAPI.ValidateUpdateRule(context.Background(), id).FailureDetectionRule(failureDetectionRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionRulesAPI.ValidateUpdateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required failure detection rule. Needs to be a valid RFC 4122 UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **failureDetectionRule** | [**FailureDetectionRule**](FailureDetectionRule.md) | The JSON body of the request. Contains the configuration of the failure detection rule to be validated. | 

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

