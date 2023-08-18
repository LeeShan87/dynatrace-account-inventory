# \ServiceFailureDetectionParameterSetsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeFdpId**](ServiceFailureDetectionParameterSetsAPI.md#ChangeFdpId) | **Put** /service/failureDetection/parameterSelection/parameterSets/{id}/changeId | Changes the ID of the specified failure detection parameter set | maturity&#x3D;EARLY_ADOPTER
[**CreateFdp**](ServiceFailureDetectionParameterSetsAPI.md#CreateFdp) | **Post** /service/failureDetection/parameterSelection/parameterSets | Creates a new failure detection parameter set | maturity&#x3D;EARLY_ADOPTER
[**CreateOrUpdateFdp**](ServiceFailureDetectionParameterSetsAPI.md#CreateOrUpdateFdp) | **Put** /service/failureDetection/parameterSelection/parameterSets/{id} | Updates the specified failure detection parameter set | maturity&#x3D;EARLY_ADOPTER
[**DeleteFdp**](ServiceFailureDetectionParameterSetsAPI.md#DeleteFdp) | **Delete** /service/failureDetection/parameterSelection/parameterSets/{id} | Deletes the specified failure detection parameter set | maturity&#x3D;EARLY_ADOPTER
[**GetAllFdps**](ServiceFailureDetectionParameterSetsAPI.md#GetAllFdps) | **Get** /service/failureDetection/parameterSelection/parameterSets | Lists all available failure detection parameter sets | maturity&#x3D;EARLY_ADOPTER
[**GetFdp**](ServiceFailureDetectionParameterSetsAPI.md#GetFdp) | **Get** /service/failureDetection/parameterSelection/parameterSets/{id} | Gets the specified failure detection parameter set | maturity&#x3D;EARLY_ADOPTER
[**ValidateCreateFdp**](ServiceFailureDetectionParameterSetsAPI.md#ValidateCreateFdp) | **Post** /service/failureDetection/parameterSelection/parameterSets/validator | Validates the payload for the &#x60;POST /service/failureDetection/parameterSelection/parameterSets&#x60; request | maturity&#x3D;EARLY_ADOPTER
[**ValidateUpdateFdp**](ServiceFailureDetectionParameterSetsAPI.md#ValidateUpdateFdp) | **Post** /service/failureDetection/parameterSelection/parameterSets/{id}/validator | Validates the payload for the &#x60;PUT /service/failureDetection/parameterSelection/parameterSets/{id}&#x60; request | maturity&#x3D;EARLY_ADOPTER



## ChangeFdpId

> ChangeFdpId(ctx, id).EntityShortRepresentation(entityShortRepresentation).Execute()

Changes the ID of the specified failure detection parameter set | maturity=EARLY_ADOPTER



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
    id := "123e4567-e89b-12d3-a456-426614174000" // string | The ID of the required failure detection parameter set. Needs to be a valid RFC 4122 UUID.
    entityShortRepresentation := *openapiclient.NewEntityShortRepresentation("Id_example") // EntityShortRepresentation | The JSON body of the request. Contains the new ID of the set. All other fields are ignored.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceFailureDetectionParameterSetsAPI.ChangeFdpId(context.Background(), id).EntityShortRepresentation(entityShortRepresentation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionParameterSetsAPI.ChangeFdpId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required failure detection parameter set. Needs to be a valid RFC 4122 UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeFdpIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityShortRepresentation** | [**EntityShortRepresentation**](EntityShortRepresentation.md) | The JSON body of the request. Contains the new ID of the set. All other fields are ignored. | 

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


## CreateFdp

> EntityShortRepresentation CreateFdp(ctx).FailureDetectionParameterSet(failureDetectionParameterSet).Execute()

Creates a new failure detection parameter set | maturity=EARLY_ADOPTER

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
    failureDetectionParameterSet := *openapiclient.NewFailureDetectionParameterSet() // FailureDetectionParameterSet | The JSON body of the request. Contains the new failure detection parameter set.   Dynatrace will generate a random UUID for you, if you don't specify an ID. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceFailureDetectionParameterSetsAPI.CreateFdp(context.Background()).FailureDetectionParameterSet(failureDetectionParameterSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionParameterSetsAPI.CreateFdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFdp`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceFailureDetectionParameterSetsAPI.CreateFdp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failureDetectionParameterSet** | [**FailureDetectionParameterSet**](FailureDetectionParameterSet.md) | The JSON body of the request. Contains the new failure detection parameter set.   Dynatrace will generate a random UUID for you, if you don&#39;t specify an ID. | 

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


## CreateOrUpdateFdp

> EntityShortRepresentation CreateOrUpdateFdp(ctx, id).FailureDetectionParameterSet(failureDetectionParameterSet).Execute()

Updates the specified failure detection parameter set | maturity=EARLY_ADOPTER



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
    id := "123e4567-e89b-12d3-a456-426614174000" // string | The ID of the required failure detection parameter set. Needs to be a valid RFC 4122 UUID.
    failureDetectionParameterSet := *openapiclient.NewFailureDetectionParameterSet() // FailureDetectionParameterSet | The JSON body of the request. Contains the updated failure detection parameter set.   You can't update the ID with this request. Use the change ID request instead. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceFailureDetectionParameterSetsAPI.CreateOrUpdateFdp(context.Background(), id).FailureDetectionParameterSet(failureDetectionParameterSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionParameterSetsAPI.CreateOrUpdateFdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateFdp`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceFailureDetectionParameterSetsAPI.CreateOrUpdateFdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required failure detection parameter set. Needs to be a valid RFC 4122 UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateFdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **failureDetectionParameterSet** | [**FailureDetectionParameterSet**](FailureDetectionParameterSet.md) | The JSON body of the request. Contains the updated failure detection parameter set.   You can&#39;t update the ID with this request. Use the change ID request instead. | 

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


## DeleteFdp

> DeleteFdp(ctx, id).Execute()

Deletes the specified failure detection parameter set | maturity=EARLY_ADOPTER



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
    id := "123e4567-e89b-12d3-a456-426614174000" // string | The ID of the required failure detection parameter set. Needs to be a valid RFC 4122 UUID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceFailureDetectionParameterSetsAPI.DeleteFdp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionParameterSetsAPI.DeleteFdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required failure detection parameter set. Needs to be a valid RFC 4122 UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFdpRequest struct via the builder pattern


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


## GetAllFdps

> StubList GetAllFdps(ctx).Execute()

Lists all available failure detection parameter sets | maturity=EARLY_ADOPTER

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
    resp, r, err := apiClient.ServiceFailureDetectionParameterSetsAPI.GetAllFdps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionParameterSetsAPI.GetAllFdps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllFdps`: StubList
    fmt.Fprintf(os.Stdout, "Response from `ServiceFailureDetectionParameterSetsAPI.GetAllFdps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFdpsRequest struct via the builder pattern


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


## GetFdp

> FailureDetectionParameterSet GetFdp(ctx, id).Execute()

Gets the specified failure detection parameter set | maturity=EARLY_ADOPTER

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
    id := "123e4567-e89b-12d3-a456-426614174000" // string | The ID of the required failure detection parameter set. Needs to be a valid RFC 4122 UUID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceFailureDetectionParameterSetsAPI.GetFdp(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionParameterSetsAPI.GetFdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFdp`: FailureDetectionParameterSet
    fmt.Fprintf(os.Stdout, "Response from `ServiceFailureDetectionParameterSetsAPI.GetFdp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required failure detection parameter set. Needs to be a valid RFC 4122 UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FailureDetectionParameterSet**](FailureDetectionParameterSet.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCreateFdp

> ValidateCreateFdp(ctx).FailureDetectionParameterSet(failureDetectionParameterSet).Execute()

Validates the payload for the `POST /service/failureDetection/parameterSelection/parameterSets` request | maturity=EARLY_ADOPTER

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
    failureDetectionParameterSet := *openapiclient.NewFailureDetectionParameterSet() // FailureDetectionParameterSet | The JSON body of the request. Contains the failure detection parameter set to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceFailureDetectionParameterSetsAPI.ValidateCreateFdp(context.Background()).FailureDetectionParameterSet(failureDetectionParameterSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionParameterSetsAPI.ValidateCreateFdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateFdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **failureDetectionParameterSet** | [**FailureDetectionParameterSet**](FailureDetectionParameterSet.md) | The JSON body of the request. Contains the failure detection parameter set to be validated. | 

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


## ValidateUpdateFdp

> ValidateUpdateFdp(ctx, id).FailureDetectionParameterSet(failureDetectionParameterSet).Execute()

Validates the payload for the `PUT /service/failureDetection/parameterSelection/parameterSets/{id}` request | maturity=EARLY_ADOPTER

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
    id := "123e4567-e89b-12d3-a456-426614174000" // string | The ID of the required failure detection parameter set. Needs to be a valid RFC 4122 UUID.
    failureDetectionParameterSet := *openapiclient.NewFailureDetectionParameterSet() // FailureDetectionParameterSet | The JSON body of the request. Contains the failure detection parameter set to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceFailureDetectionParameterSetsAPI.ValidateUpdateFdp(context.Background(), id).FailureDetectionParameterSet(failureDetectionParameterSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceFailureDetectionParameterSetsAPI.ValidateUpdateFdp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required failure detection parameter set. Needs to be a valid RFC 4122 UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateFdpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **failureDetectionParameterSet** | [**FailureDetectionParameterSet**](FailureDetectionParameterSet.md) | The JSON body of the request. Contains the failure detection parameter set to be validated. | 

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

