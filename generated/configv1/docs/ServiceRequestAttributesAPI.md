# \ServiceRequestAttributesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequestAttributesConfig**](ServiceRequestAttributesAPI.md#CreateRequestAttributesConfig) | **Post** /service/requestAttributes | Creates a new request attribute
[**DeleteRequestAttributesConfig**](ServiceRequestAttributesAPI.md#DeleteRequestAttributesConfig) | **Delete** /service/requestAttributes/{id} | Deletes the specified request attribute
[**GetRequestAttributesConfig**](ServiceRequestAttributesAPI.md#GetRequestAttributesConfig) | **Get** /service/requestAttributes/{id} | Shows the properties of the specified request attribute
[**ListRequestAttributesConfigs**](ServiceRequestAttributesAPI.md#ListRequestAttributesConfigs) | **Get** /service/requestAttributes | Lists all available request attributes
[**UpdateRequestAttributesConfig**](ServiceRequestAttributesAPI.md#UpdateRequestAttributesConfig) | **Put** /service/requestAttributes/{id} | Updates an existing request attribute or creates a new one
[**ValidateCreateRequestAttributesConfig**](ServiceRequestAttributesAPI.md#ValidateCreateRequestAttributesConfig) | **Post** /service/requestAttributes/validator | Validates new request attributes for the &#x60;POST /requestAttributes&#x60; request
[**ValidateUpdateRequestAttributesConfig**](ServiceRequestAttributesAPI.md#ValidateUpdateRequestAttributesConfig) | **Post** /service/requestAttributes/{id}/validator | Validate updates of existing request attribute for the &#x60;PUT /requestAttributes/{id}&#x60; request



## CreateRequestAttributesConfig

> EntityShortRepresentation CreateRequestAttributesConfig(ctx).RequestAttribute(requestAttribute).Execute()

Creates a new request attribute



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
    requestAttribute := *openapiclient.NewRequestAttribute("Aggregation_example", false, []openapiclient.DataSource{*openapiclient.NewDataSource(false, "Source_example")}, "DataType_example", false, "Name_example", "Normalization_example", false) // RequestAttribute | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRequestAttributesAPI.CreateRequestAttributesConfig(context.Background()).RequestAttribute(requestAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesAPI.CreateRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRequestAttributesConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestAttributesAPI.CreateRequestAttributesConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestAttributesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestAttribute** | [**RequestAttribute**](RequestAttribute.md) |  | 

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


## DeleteRequestAttributesConfig

> DeleteRequestAttributesConfig(ctx, id).Execute()

Deletes the specified request attribute

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the request attribute to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRequestAttributesAPI.DeleteRequestAttributesConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesAPI.DeleteRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the request attribute to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestAttributesConfigRequest struct via the builder pattern


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


## GetRequestAttributesConfig

> RequestAttribute GetRequestAttributesConfig(ctx, id).IncludeProcessGroupReferences(includeProcessGroupReferences).Execute()

Shows the properties of the specified request attribute

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required request attribute.
    includeProcessGroupReferences := true // bool | Flag to include process group references to the response.    Process Group group references aren't compatible across environments. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRequestAttributesAPI.GetRequestAttributesConfig(context.Background(), id).IncludeProcessGroupReferences(includeProcessGroupReferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesAPI.GetRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRequestAttributesConfig`: RequestAttribute
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestAttributesAPI.GetRequestAttributesConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required request attribute. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestAttributesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **bool** | Flag to include process group references to the response.    Process Group group references aren&#39;t compatible across environments. | [default to false]

### Return type

[**RequestAttribute**](RequestAttribute.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRequestAttributesConfigs

> StubList ListRequestAttributesConfigs(ctx).Execute()

Lists all available request attributes

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
    resp, r, err := apiClient.ServiceRequestAttributesAPI.ListRequestAttributesConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesAPI.ListRequestAttributesConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRequestAttributesConfigs`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestAttributesAPI.ListRequestAttributesConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRequestAttributesConfigsRequest struct via the builder pattern


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


## UpdateRequestAttributesConfig

> EntityShortRepresentation UpdateRequestAttributesConfig(ctx, id).RequestAttribute(requestAttribute).Execute()

Updates an existing request attribute or creates a new one

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the request attribute to be updated.   If you set the ID in the body as well, it must match this ID.
    requestAttribute := *openapiclient.NewRequestAttribute("Aggregation_example", false, []openapiclient.DataSource{*openapiclient.NewDataSource(false, "Source_example")}, "DataType_example", false, "Name_example", "Normalization_example", false) // RequestAttribute | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceRequestAttributesAPI.UpdateRequestAttributesConfig(context.Background(), id).RequestAttribute(requestAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesAPI.UpdateRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRequestAttributesConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceRequestAttributesAPI.UpdateRequestAttributesConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the request attribute to be updated.   If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequestAttributesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestAttribute** | [**RequestAttribute**](RequestAttribute.md) |  | 

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


## ValidateCreateRequestAttributesConfig

> ValidateCreateRequestAttributesConfig(ctx).RequestAttribute(requestAttribute).Execute()

Validates new request attributes for the `POST /requestAttributes` request

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
    requestAttribute := *openapiclient.NewRequestAttribute("Aggregation_example", false, []openapiclient.DataSource{*openapiclient.NewDataSource(false, "Source_example")}, "DataType_example", false, "Name_example", "Normalization_example", false) // RequestAttribute | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRequestAttributesAPI.ValidateCreateRequestAttributesConfig(context.Background()).RequestAttribute(requestAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesAPI.ValidateCreateRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateRequestAttributesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestAttribute** | [**RequestAttribute**](RequestAttribute.md) |  | 

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


## ValidateUpdateRequestAttributesConfig

> ValidateUpdateRequestAttributesConfig(ctx, id).RequestAttribute(requestAttribute).Execute()

Validate updates of existing request attribute for the `PUT /requestAttributes/{id}` request

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the request attribute to be validated.
    requestAttribute := *openapiclient.NewRequestAttribute("Aggregation_example", false, []openapiclient.DataSource{*openapiclient.NewDataSource(false, "Source_example")}, "DataType_example", false, "Name_example", "Normalization_example", false) // RequestAttribute | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceRequestAttributesAPI.ValidateUpdateRequestAttributesConfig(context.Background(), id).RequestAttribute(requestAttribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceRequestAttributesAPI.ValidateUpdateRequestAttributesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the request attribute to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateRequestAttributesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestAttribute** | [**RequestAttribute**](RequestAttribute.md) |  | 

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

