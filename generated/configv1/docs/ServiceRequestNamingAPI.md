# \ServiceRequestNamingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequestNaming**](ServiceRequestNamingAPI.md#CreateRequestNaming) | **Post** /service/requestNaming | Creates a new request naming rule
[**DeleteRequestNaming**](ServiceRequestNamingAPI.md#DeleteRequestNaming) | **Delete** /service/requestNaming/{id} | Deletes the specified request naming rule
[**Get**](ServiceRequestNamingAPI.md#Get) | **Get** /service/resourceNaming | Lists the global service resource requests.
[**GetRequestNaming**](ServiceRequestNamingAPI.md#GetRequestNaming) | **Get** /service/requestNaming/{id} | Gets the parameters of the specified request naming rule
[**ListRequestNaming**](ServiceRequestNamingAPI.md#ListRequestNaming) | **Get** /service/requestNaming | Lists all request naming rules along with their parameters
[**OrderRequestNaming**](ServiceRequestNamingAPI.md#OrderRequestNaming) | **Put** /service/requestNaming/order | Reorders the request namings
[**Put**](ServiceRequestNamingAPI.md#Put) | **Put** /service/resourceNaming | Updates the global service resource requests.
[**UpdateRequestNaming**](ServiceRequestNamingAPI.md#UpdateRequestNaming) | **Put** /service/requestNaming/{id} | Updates the specified request naming rule
[**Validate**](ServiceRequestNamingAPI.md#Validate) | **Post** /service/resourceNaming/validator | Validates new resource requests settings for the &#x60;PUT /service/resourceRequest&#x60; request.
[**ValidateCreateRequestNaming**](ServiceRequestNamingAPI.md#ValidateCreateRequestNaming) | **Post** /service/requestNaming/validator | Validates the new request naming rule for the &#x60;POST /requestNaming&#x60; request
[**ValidateUpdateRequestNaming**](ServiceRequestNamingAPI.md#ValidateUpdateRequestNaming) | **Post** /service/requestNaming/{id}/validator | Validates the new request naming for the &#x60;PUT /requestNaming/{id}&#x60; request



## CreateRequestNaming

> EntityShortRepresentation CreateRequestNaming(ctx).Position(position).RequestNaming(requestNaming).Execute()

Creates a new request naming rule



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
    position := "position_example" // string | Order of the new request naming rule. Set to `PREPEND` to prepend it to the list, `APPEND` to append it. Defaults to `APPEND`. (optional) (default to "APPEND")
    requestNaming := *openapiclient.NewRequestNaming([]openapiclient.Condition{*openapiclient.NewCondition("Attribute_example", *openapiclient.NewComparisonInfo("Comparison_example", false, "Type_example"))}, false, "NamingPattern_example") // RequestNaming | The JSON body of the request containing parameters of the new request naming rule.    You must not specify the ID of the rule! (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRequestNamingAPI.CreateRequestNaming(context.Background()).Position(position).RequestNaming(requestNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.CreateRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequestNaming`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestNamingAPI.CreateRequestNaming`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **position** | **string** | Order of the new request naming rule. Set to &#x60;PREPEND&#x60; to prepend it to the list, &#x60;APPEND&#x60; to append it. Defaults to &#x60;APPEND&#x60;. | [default to &quot;APPEND&quot;]
 **requestNaming** | [**RequestNaming**](RequestNaming.md) | The JSON body of the request containing parameters of the new request naming rule.    You must not specify the ID of the rule! | 

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


## DeleteRequestNaming

> DeleteRequestNaming(ctx, id).Execute()

Deletes the specified request naming rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the request naming rule to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRequestNamingAPI.DeleteRequestNaming(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.DeleteRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the request naming rule to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestNamingRequest struct via the builder pattern


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


## Get

> ResourceNaming Get(ctx).Execute()

Lists the global service resource requests.



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
    resp, r, err := apiClient.ServiceRequestNamingAPI.Get(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: ResourceNaming
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestNamingAPI.Get`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


### Return type

[**ResourceNaming**](ResourceNaming.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestNaming

> RequestNaming GetRequestNaming(ctx, id).Execute()

Gets the parameters of the specified request naming rule

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the request naming rule you're inquiring.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRequestNamingAPI.GetRequestNaming(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.GetRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestNaming`: RequestNaming
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestNamingAPI.GetRequestNaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the request naming rule you&#39;re inquiring. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestNaming**](RequestNaming.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestNaming

> StubList ListRequestNaming(ctx).Execute()

Lists all request naming rules along with their parameters

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
    resp, r, err := apiClient.ServiceRequestNamingAPI.ListRequestNaming(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.ListRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestNaming`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestNamingAPI.ListRequestNaming`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequestNamingRequest struct via the builder pattern


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


## OrderRequestNaming

> OrderRequestNaming(ctx).StubList(stubList).Execute()

Reorders the request namings



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
    stubList := *openapiclient.NewStubList([]openapiclient.EntityShortRepresentation{*openapiclient.NewEntityShortRepresentation("Id_example")}) // StubList | JSON body of the request containing the IDs of the request naming rules in the desired order. Any further properties (*name*, *description*) will be ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRequestNamingAPI.OrderRequestNaming(context.Background()).StubList(stubList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.OrderRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrderRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **stubList** | [**StubList**](StubList.md) | JSON body of the request containing the IDs of the request naming rules in the desired order. Any further properties (*name*, *description*) will be ignored. | 

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


## Put

> Put(ctx).ResourceNaming(resourceNaming).Execute()

Updates the global service resource requests.



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
    resourceNaming := *openapiclient.NewResourceNaming([]string{"Binary_example"}, []string{"Image_example"}) // ResourceNaming | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRequestNamingAPI.Put(context.Background()).ResourceNaming(resourceNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.Put``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceNaming** | [**ResourceNaming**](ResourceNaming.md) |  | 

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


## UpdateRequestNaming

> EntityShortRepresentation UpdateRequestNaming(ctx, id).RequestNaming(requestNaming).Execute()

Updates the specified request naming rule



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the request naming to be updated.   The ID of the request naming in the body of the request must match this ID.
    requestNaming := *openapiclient.NewRequestNaming([]openapiclient.Condition{*openapiclient.NewCondition("Attribute_example", *openapiclient.NewComparisonInfo("Comparison_example", false, "Type_example"))}, false, "NamingPattern_example") // RequestNaming | The JSON body of the request containing updated parameters of the request naming. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRequestNamingAPI.UpdateRequestNaming(context.Background(), id).RequestNaming(requestNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.UpdateRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRequestNaming`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestNamingAPI.UpdateRequestNaming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the request naming to be updated.   The ID of the request naming in the body of the request must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestNaming** | [**RequestNaming**](RequestNaming.md) | The JSON body of the request containing updated parameters of the request naming. | 

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


## Validate

> Validate(ctx).ResourceNaming(resourceNaming).Execute()

Validates new resource requests settings for the `PUT /service/resourceRequest` request.

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
    resourceNaming := *openapiclient.NewResourceNaming([]string{"Binary_example"}, []string{"Image_example"}) // ResourceNaming | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRequestNamingAPI.Validate(context.Background()).ResourceNaming(resourceNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.Validate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceNaming** | [**ResourceNaming**](ResourceNaming.md) |  | 

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


## ValidateCreateRequestNaming

> ValidateCreateRequestNaming(ctx).RequestNaming(requestNaming).Execute()

Validates the new request naming rule for the `POST /requestNaming` request

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
    requestNaming := *openapiclient.NewRequestNaming([]openapiclient.Condition{*openapiclient.NewCondition("Attribute_example", *openapiclient.NewComparisonInfo("Comparison_example", false, "Type_example"))}, false, "NamingPattern_example") // RequestNaming | The JSON body of the request containing parameters of the new request naming rule.    You must not specify the ID of the rule! (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRequestNamingAPI.ValidateCreateRequestNaming(context.Background()).RequestNaming(requestNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.ValidateCreateRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestNaming** | [**RequestNaming**](RequestNaming.md) | The JSON body of the request containing parameters of the new request naming rule.    You must not specify the ID of the rule! | 

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


## ValidateUpdateRequestNaming

> ValidateUpdateRequestNaming(ctx, id).RequestNaming(requestNaming).Execute()

Validates the new request naming for the `PUT /requestNaming/{id}` request



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the request naming to be updated.   The ID of the request naming in the body of the request must match this ID.
    requestNaming := *openapiclient.NewRequestNaming([]openapiclient.Condition{*openapiclient.NewCondition("Attribute_example", *openapiclient.NewComparisonInfo("Comparison_example", false, "Type_example"))}, false, "NamingPattern_example") // RequestNaming | The JSON body of the request containing updated parameters of the request naming. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRequestNamingAPI.ValidateUpdateRequestNaming(context.Background(), id).RequestNaming(requestNaming).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestNamingAPI.ValidateUpdateRequestNaming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the request naming to be updated.   The ID of the request naming in the body of the request must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateRequestNamingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestNaming** | [**RequestNaming**](RequestNaming.md) | The JSON body of the request containing updated parameters of the request naming. | 

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

