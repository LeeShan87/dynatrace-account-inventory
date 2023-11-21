# \AutomaticallyAppliedTagsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutoTag**](AutomaticallyAppliedTagsAPI.md#CreateAutoTag) | **Post** /autoTags | Creates a new auto-tag
[**DeleteAutoTag**](AutomaticallyAppliedTagsAPI.md#DeleteAutoTag) | **Delete** /autoTags/{id} | Deletes the specified auto-tag
[**GetAutoTag**](AutomaticallyAppliedTagsAPI.md#GetAutoTag) | **Get** /autoTags/{id} | Gets the properties of the specified auto-tag
[**ListAutoTags**](AutomaticallyAppliedTagsAPI.md#ListAutoTags) | **Get** /autoTags | Lists all configured auto-tags
[**UpdateAutoTag**](AutomaticallyAppliedTagsAPI.md#UpdateAutoTag) | **Put** /autoTags/{id} | Updates an existing auto-tag
[**ValidateCreateAutoTag**](AutomaticallyAppliedTagsAPI.md#ValidateCreateAutoTag) | **Post** /autoTags/validator | Validates the payload for the &#x60;POST /autoTags&#x60; request
[**ValidateUpdateAutoTag**](AutomaticallyAppliedTagsAPI.md#ValidateUpdateAutoTag) | **Post** /autoTags/{id}/validator | Validates the payload for the &#x60;PUT /autoTags/{id}&#x60; request



## CreateAutoTag

> EntityShortRepresentation CreateAutoTag(ctx).AutoTag(autoTag).Execute()

Creates a new auto-tag



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
    autoTag := *openapiclient.NewAutoTag("Name_example") // AutoTag | The JSON body of the request. Contains parameters of the new auto-tag. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomaticallyAppliedTagsAPI.CreateAutoTag(context.Background()).AutoTag(autoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticallyAppliedTagsAPI.CreateAutoTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutoTag`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AutomaticallyAppliedTagsAPI.CreateAutoTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutoTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTag** | [**AutoTag**](AutoTag.md) | The JSON body of the request. Contains parameters of the new auto-tag. | 

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


## DeleteAutoTag

> DeleteAutoTag(ctx, id).Execute()

Deletes the specified auto-tag

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the auto-tag to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutomaticallyAppliedTagsAPI.DeleteAutoTag(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticallyAppliedTagsAPI.DeleteAutoTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the auto-tag to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutoTagRequest struct via the builder pattern


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


## GetAutoTag

> AutoTag GetAutoTag(ctx, id).IncludeProcessGroupReferences(includeProcessGroupReferences).Execute()

Gets the properties of the specified auto-tag

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required auto-tag.
    includeProcessGroupReferences := true // bool | Flag to include process group references to the response.    Process group references aren't compatible across environments.   When this flag is set to `false`, conditions with process group references are removed. If that results in a rule having no conditions, the entire rule is removed. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomaticallyAppliedTagsAPI.GetAutoTag(context.Background(), id).IncludeProcessGroupReferences(includeProcessGroupReferences).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticallyAppliedTagsAPI.GetAutoTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAutoTag`: AutoTag
    fmt.Fprintf(os.Stdout, "Response from `AutomaticallyAppliedTagsAPI.GetAutoTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required auto-tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProcessGroupReferences** | **bool** | Flag to include process group references to the response.    Process group references aren&#39;t compatible across environments.   When this flag is set to &#x60;false&#x60;, conditions with process group references are removed. If that results in a rule having no conditions, the entire rule is removed. | [default to false]

### Return type

[**AutoTag**](AutoTag.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAutoTags

> StubList ListAutoTags(ctx).Execute()

Lists all configured auto-tags

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
    resp, r, err := apiClient.AutomaticallyAppliedTagsAPI.ListAutoTags(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticallyAppliedTagsAPI.ListAutoTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAutoTags`: StubList
    fmt.Fprintf(os.Stdout, "Response from `AutomaticallyAppliedTagsAPI.ListAutoTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAutoTagsRequest struct via the builder pattern


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


## UpdateAutoTag

> EntityShortRepresentation UpdateAutoTag(ctx, id).AutoTag(autoTag).Execute()

Updates an existing auto-tag



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the auto-tag to be updated.   If you set the ID in the body as well, it must match this ID.
    autoTag := *openapiclient.NewAutoTag("Name_example") // AutoTag | The JSON body of the request. Contains updated parameters of the auto-tag. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AutomaticallyAppliedTagsAPI.UpdateAutoTag(context.Background(), id).AutoTag(autoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticallyAppliedTagsAPI.UpdateAutoTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAutoTag`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AutomaticallyAppliedTagsAPI.UpdateAutoTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the auto-tag to be updated.   If you set the ID in the body as well, it must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutoTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoTag** | [**AutoTag**](AutoTag.md) | The JSON body of the request. Contains updated parameters of the auto-tag. | 

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


## ValidateCreateAutoTag

> ValidateCreateAutoTag(ctx).AutoTag(autoTag).Execute()

Validates the payload for the `POST /autoTags` request

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
    autoTag := *openapiclient.NewAutoTag("Name_example") // AutoTag | The JSON body of the request. Contains auto-tag configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutomaticallyAppliedTagsAPI.ValidateCreateAutoTag(context.Background()).AutoTag(autoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticallyAppliedTagsAPI.ValidateCreateAutoTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateAutoTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **autoTag** | [**AutoTag**](AutoTag.md) | The JSON body of the request. Contains auto-tag configuration to be validated. | 

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


## ValidateUpdateAutoTag

> ValidateUpdateAutoTag(ctx, id).AutoTag(autoTag).Execute()

Validates the payload for the `PUT /autoTags/{id}` request

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the auto-tag to be validated.
    autoTag := *openapiclient.NewAutoTag("Name_example") // AutoTag | The JSON body of the request. Contains auto-tag configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AutomaticallyAppliedTagsAPI.ValidateUpdateAutoTag(context.Background(), id).AutoTag(autoTag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AutomaticallyAppliedTagsAPI.ValidateUpdateAutoTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the auto-tag to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateAutoTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **autoTag** | [**AutoTag**](AutoTag.md) | The JSON body of the request. Contains auto-tag configuration to be validated. | 

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

