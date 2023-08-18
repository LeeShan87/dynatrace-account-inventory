# \ServiceCustomServicesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomService**](ServiceCustomServicesAPI.md#CreateCustomService) | **Post** /service/customServices/{technology} | Creates a custom service
[**DeleteCustomService**](ServiceCustomServicesAPI.md#DeleteCustomService) | **Delete** /service/customServices/{technology}/{id} | Deletes the specified custom service
[**GetCustomService**](ServiceCustomServicesAPI.md#GetCustomService) | **Get** /service/customServices/{technology}/{id} | Gets the definition of the specified custom service
[**ListCustomServices**](ServiceCustomServicesAPI.md#ListCustomServices) | **Get** /service/customServices/{technology} | Lists all custom services of the specified technology
[**OrderCustomServices**](ServiceCustomServicesAPI.md#OrderCustomServices) | **Put** /service/customServices/{technology}/order | Reorders the custom services of the specified technology
[**UpdateCustomService**](ServiceCustomServicesAPI.md#UpdateCustomService) | **Put** /service/customServices/{technology}/{id} | Updates the specified custom service or create a new one.
[**ValidateCreateCustomService**](ServiceCustomServicesAPI.md#ValidateCreateCustomService) | **Post** /service/customServices/{technology}/validator | Validate the new custom service for the &#x60;POST /customServices/{technology}&#x60; request
[**ValidateUpdateCustomService**](ServiceCustomServicesAPI.md#ValidateUpdateCustomService) | **Post** /service/customServices/{technology}/{id}/validator | Validate the new custom service for the &#x60;PUT /customServices/{technology}/{id}&#x60; request



## CreateCustomService

> EntityShortRepresentation CreateCustomService(ctx, technology).Position(position).CustomService(customService).Execute()

Creates a custom service



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
    technology := "technology_example" // string | Technology of the new custom service.
    position := "position_example" // string | Order of the new custom service. Set to `PREPEND` to prepend it to the list, `APPEND` to append it. Defaults to `APPEND`. (optional) (default to "APPEND")
    customService := *openapiclient.NewCustomService(false, "Name_example", false, []openapiclient.DetectionRule{*openapiclient.NewDetectionRule(false, []openapiclient.MethodRule{*openapiclient.NewMethodRule("MethodName_example", "ReturnType_example")})}) // CustomService | JSON body of the request containing definition of the new custom service.   You must not specify the IDs for the custom service or any of its rules. The *order* field is not allowed either. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceCustomServicesAPI.CreateCustomService(context.Background(), technology).Position(position).CustomService(customService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCustomServicesAPI.CreateCustomService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomService`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceCustomServicesAPI.CreateCustomService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string** | Technology of the new custom service. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **position** | **string** | Order of the new custom service. Set to &#x60;PREPEND&#x60; to prepend it to the list, &#x60;APPEND&#x60; to append it. Defaults to &#x60;APPEND&#x60;. | [default to &quot;APPEND&quot;]
 **customService** | [**CustomService**](CustomService.md) | JSON body of the request containing definition of the new custom service.   You must not specify the IDs for the custom service or any of its rules. The *order* field is not allowed either. | 

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


## DeleteCustomService

> DeleteCustomService(ctx, technology, id).Execute()

Deletes the specified custom service

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
    technology := "technology_example" // string | Technology of the custom service to delete.
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the custom service to delete.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceCustomServicesAPI.DeleteCustomService(context.Background(), technology, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCustomServicesAPI.DeleteCustomService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string** | Technology of the custom service to delete. | 
**id** | **string** | The ID of the custom service to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomServiceRequest struct via the builder pattern


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


## GetCustomService

> CustomService GetCustomService(ctx, technology, id).IncludeProcessGroupReferences(includeProcessGroupReferences).Execute()

Gets the definition of the specified custom service

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
    technology := "technology_example" // string | Technology of the custom service you're inquiring.
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the custom service you're inquiring.
    includeProcessGroupReferences := true // bool | Flag to include process group references to the response.    Process group references aren't compatible across environments.   `false` is used if the value is not set. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceCustomServicesAPI.GetCustomService(context.Background(), technology, id).IncludeProcessGroupReferences(includeProcessGroupReferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCustomServicesAPI.GetCustomService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomService`: CustomService
    fmt.Fprintf(os.Stdout, "Response from `ServiceCustomServicesAPI.GetCustomService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string** | Technology of the custom service you&#39;re inquiring. | 
**id** | **string** | The ID of the custom service you&#39;re inquiring. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeProcessGroupReferences** | **bool** | Flag to include process group references to the response.    Process group references aren&#39;t compatible across environments.   &#x60;false&#x60; is used if the value is not set. | [default to false]

### Return type

[**CustomService**](CustomService.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomServices

> StubList ListCustomServices(ctx, technology).Execute()

Lists all custom services of the specified technology

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
    technology := "technology_example" // string | Technology of the required custom services.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceCustomServicesAPI.ListCustomServices(context.Background(), technology).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCustomServicesAPI.ListCustomServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomServices`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceCustomServicesAPI.ListCustomServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string** | Technology of the required custom services. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## OrderCustomServices

> OrderCustomServices(ctx, technology).StubList(stubList).Execute()

Reorders the custom services of the specified technology



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
    technology := "technology_example" // string | Technology of custom services to update.
    stubList := *openapiclient.NewStubList([]openapiclient.EntityShortRepresentation{*openapiclient.NewEntityShortRepresentation("Id_example")}) // StubList | JSON body of the request containing the IDs of the custom services in the desired order. Any further properties (*name*, *description*) will be ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceCustomServicesAPI.OrderCustomServices(context.Background(), technology).StubList(stubList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCustomServicesAPI.OrderCustomServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string** | Technology of custom services to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrderCustomServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stubList** | [**StubList**](StubList.md) | JSON body of the request containing the IDs of the custom services in the desired order. Any further properties (*name*, *description*) will be ignored. | 

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


## UpdateCustomService

> EntityShortRepresentation UpdateCustomService(ctx, technology, id).CustomService(customService).Execute()

Updates the specified custom service or create a new one.



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
    technology := "technology_example" // string | Technology of the custom service to update.
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the custom service to update.   The ID of the custom service in the body of the request must match this ID.
    customService := *openapiclient.NewCustomService(false, "Name_example", false, []openapiclient.DetectionRule{*openapiclient.NewDetectionRule(false, []openapiclient.MethodRule{*openapiclient.NewMethodRule("MethodName_example", "ReturnType_example")})}) // CustomService | JSON body of the request containing updated definition of the custom service. If *order* is present, it will be used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceCustomServicesAPI.UpdateCustomService(context.Background(), technology, id).CustomService(customService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCustomServicesAPI.UpdateCustomService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomService`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceCustomServicesAPI.UpdateCustomService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string** | Technology of the custom service to update. | 
**id** | **string** | The ID of the custom service to update.   The ID of the custom service in the body of the request must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customService** | [**CustomService**](CustomService.md) | JSON body of the request containing updated definition of the custom service. If *order* is present, it will be used. | 

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


## ValidateCreateCustomService

> ValidateCreateCustomService(ctx, technology).CustomService(customService).Execute()

Validate the new custom service for the `POST /customServices/{technology}` request

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
    technology := "technology_example" // string | Technology of the custom service to validate.
    customService := *openapiclient.NewCustomService(false, "Name_example", false, []openapiclient.DetectionRule{*openapiclient.NewDetectionRule(false, []openapiclient.MethodRule{*openapiclient.NewMethodRule("MethodName_example", "ReturnType_example")})}) // CustomService | JSON body of the request containing definition of the custom service to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceCustomServicesAPI.ValidateCreateCustomService(context.Background(), technology).CustomService(customService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCustomServicesAPI.ValidateCreateCustomService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string** | Technology of the custom service to validate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateCustomServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customService** | [**CustomService**](CustomService.md) | JSON body of the request containing definition of the custom service to validate. | 

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


## ValidateUpdateCustomService

> ValidateUpdateCustomService(ctx, technology, id).CustomService(customService).Execute()

Validate the new custom service for the `PUT /customServices/{technology}/{id}` request

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
    technology := "technology_example" // string | Technology of the custom service to validate.
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the custom service to validate.
    customService := *openapiclient.NewCustomService(false, "Name_example", false, []openapiclient.DetectionRule{*openapiclient.NewDetectionRule(false, []openapiclient.MethodRule{*openapiclient.NewMethodRule("MethodName_example", "ReturnType_example")})}) // CustomService | JSON body of the request containing definition of the custom service to validate. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceCustomServicesAPI.ValidateUpdateCustomService(context.Background(), technology, id).CustomService(customService).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceCustomServicesAPI.ValidateUpdateCustomService``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**technology** | **string** | Technology of the custom service to validate. | 
**id** | **string** | The ID of the custom service to validate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateCustomServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customService** | [**CustomService**](CustomService.md) | JSON body of the request containing definition of the custom service to validate. | 

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

