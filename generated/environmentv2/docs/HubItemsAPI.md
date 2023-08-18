# \HubItemsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddToEnvironment**](HubItemsAPI.md#AddToEnvironment) | **Post** /hub/extensions2/{extensionName}/actions/addToEnvironment | Registers the given extension version in the environment (using the recommended version if no other release is registered)
[**GetArtifact**](HubItemsAPI.md#GetArtifact) | **Get** /hub/extensions1/{extension1FQN}/releases/{version} | Artifact of the given version of an item
[**GetDetails**](HubItemsAPI.md#GetDetails) | **Get** /hub/extensions1/{extension1FQN} | The details of the given extension version 1
[**GetExtensionDetails**](HubItemsAPI.md#GetExtensionDetails) | **Get** /hub/extensions2/{extensionName} | The details of the given extension
[**GetImage**](HubItemsAPI.md#GetImage) | **Get** /hub/assets/images/{encodedUrl} | Retrieves the image associated with the given url. Can only obtain images used in hub items.
[**GetTechnologyDetails**](HubItemsAPI.md#GetTechnologyDetails) | **Get** /hub/technologies/{slug} | The details of the given technology
[**ListCategories**](HubItemsAPI.md#ListCategories) | **Get** /hub/categories | Lists all the available category groupings of items
[**ListItems**](HubItemsAPI.md#ListItems) | **Get** /hub/items | Lists all the available items in your environment
[**Update**](HubItemsAPI.md#Update) | **Post** /hub/extensions2/{extensionName}/actions/update | Updates the given extension to the latest compatible published version
[**UpdateReleaseNotes**](HubItemsAPI.md#UpdateReleaseNotes) | **Put** /hub/extensions2/{extensionName}/releases/{version}/releaseNotes | Sets or updates the release notes of the specified extension release
[**UploadMetadataForExtension**](HubItemsAPI.md#UploadMetadataForExtension) | **Put** /hub/extensions2/{extensionName}/metadata | Create/Replace metadata for an extension that does not have Dynatrace distributed metadata.



## AddToEnvironment

> RegisteredExtensionResultDto AddToEnvironment(ctx, extensionName).ExtensionVersion(extensionVersion).Execute()

Registers the given extension version in the environment (using the recommended version if no other release is registered)

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
    extensionName := "extensionName_example" // string | Fully qualified name of the extension
    extensionVersion := "extensionVersion_example" // string | Version of the extension. Fallback to the evaluated recommended version when the version is not provided (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HubItemsAPI.AddToEnvironment(context.Background(), extensionName).ExtensionVersion(extensionVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.AddToEnvironment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddToEnvironment`: RegisteredExtensionResultDto
    fmt.Fprintf(os.Stdout, "Response from `HubItemsAPI.AddToEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | Fully qualified name of the extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddToEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionVersion** | **string** | Version of the extension. Fallback to the evaluated recommended version when the version is not provided | 

### Return type

[**RegisteredExtensionResultDto**](RegisteredExtensionResultDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifact

> GetArtifact(ctx, extension1FQN, version).Execute()

Artifact of the given version of an item

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
    extension1FQN := "extension1FQN_example" // string | Fully qualified name of the extension1/plugin
    version := "version_example" // string | Version of the release of the extension1/plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HubItemsAPI.GetArtifact(context.Background(), extension1FQN, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.GetArtifact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extension1FQN** | **string** | Fully qualified name of the extension1/plugin | 
**version** | **string** | Version of the release of the extension1/plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetails

> ItemDetails GetDetails(ctx, extension1FQN).Execute()

The details of the given extension version 1

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
    extension1FQN := "extension1FQN_example" // string | Fully qualified name of the extension1/plugin

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HubItemsAPI.GetDetails(context.Background(), extension1FQN).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.GetDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDetails`: ItemDetails
    fmt.Fprintf(os.Stdout, "Response from `HubItemsAPI.GetDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extension1FQN** | **string** | Fully qualified name of the extension1/plugin | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ItemDetails**](ItemDetails.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionDetails

> ItemDetails GetExtensionDetails(ctx, extensionName).Execute()

The details of the given extension

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
    extensionName := "extensionName_example" // string | Fully qualified name of the extension

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HubItemsAPI.GetExtensionDetails(context.Background(), extensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.GetExtensionDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtensionDetails`: ItemDetails
    fmt.Fprintf(os.Stdout, "Response from `HubItemsAPI.GetExtensionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | Fully qualified name of the extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ItemDetails**](ItemDetails.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetImage

> GetImage(ctx, encodedUrl).Execute()

Retrieves the image associated with the given url. Can only obtain images used in hub items.

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
    encodedUrl := "encodedUrl_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HubItemsAPI.GetImage(context.Background(), encodedUrl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.GetImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedUrl** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRequest struct via the builder pattern


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


## GetTechnologyDetails

> ItemDetails GetTechnologyDetails(ctx, slug).Execute()

The details of the given technology

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
    slug := "slug_example" // string | Slug of the technology to be fetched

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HubItemsAPI.GetTechnologyDetails(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.GetTechnologyDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTechnologyDetails`: ItemDetails
    fmt.Fprintf(os.Stdout, "Response from `HubItemsAPI.GetTechnologyDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | Slug of the technology to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTechnologyDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ItemDetails**](ItemDetails.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCategories

> CategoryList ListCategories(ctx).Execute()

Lists all the available category groupings of items

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
    resp, r, err := apiClient.HubItemsAPI.ListCategories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.ListCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategories`: CategoryList
    fmt.Fprintf(os.Stdout, "Response from `HubItemsAPI.ListCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCategoriesRequest struct via the builder pattern


### Return type

[**CategoryList**](CategoryList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListItems

> ItemList ListItems(ctx).NextPageKey(nextPageKey).PageSize(pageSize).ItemType(itemType).Query(query).Installed(installed).CategoryId(categoryId).Offset(offset).Execute()

Lists all the available items in your environment

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
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of hub items in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. (optional)
    itemType := "itemType_example" // string | If provided, will filter out results based on the given item type. (optional)
    query := "query_example" // string | Filter the results for items matching the query string within id, name, author, description or any tag. - Case insensitive - Spaces in the query string will lead to an intersection of the results of each term (optional)
    installed := true // bool | If provided, will restrict the result to Extensions 2.0 that have the respective installed state. - Only applies if itemType filter is not set or EXTENSION2 (optional)
    categoryId := "categoryId_example" // string | If provided, will filter items that belong to the given category. - This overrides the itemType or installed filters - For a list of category ids refer to /categories - Will return the items in the order of the category (optional)
    offset := "offset_example" // string | If provided, will skip the desired number of results, allowing for pagination in combination with page size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HubItemsAPI.ListItems(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).ItemType(itemType).Query(query).Installed(installed).CategoryId(categoryId).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.ListItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListItems`: ItemList
    fmt.Fprintf(os.Stdout, "Response from `HubItemsAPI.ListItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of hub items in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. | 
 **itemType** | **string** | If provided, will filter out results based on the given item type. | 
 **query** | **string** | Filter the results for items matching the query string within id, name, author, description or any tag. - Case insensitive - Spaces in the query string will lead to an intersection of the results of each term | 
 **installed** | **bool** | If provided, will restrict the result to Extensions 2.0 that have the respective installed state. - Only applies if itemType filter is not set or EXTENSION2 | 
 **categoryId** | **string** | If provided, will filter items that belong to the given category. - This overrides the itemType or installed filters - For a list of category ids refer to /categories - Will return the items in the order of the category | 
 **offset** | **string** | If provided, will skip the desired number of results, allowing for pagination in combination with page size | 

### Return type

[**ItemList**](ItemList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> RegisteredExtensionResultDto Update(ctx, extensionName).ExtensionVersion(extensionVersion).Execute()

Updates the given extension to the latest compatible published version

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
    extensionName := "extensionName_example" // string | Fully qualified name of the extension
    extensionVersion := "extensionVersion_example" // string | Version of the extension. Fallback to the evaluated recommended version when the version is not provided (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HubItemsAPI.Update(context.Background(), extensionName).ExtensionVersion(extensionVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: RegisteredExtensionResultDto
    fmt.Fprintf(os.Stdout, "Response from `HubItemsAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | Fully qualified name of the extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionVersion** | **string** | Version of the extension. Fallback to the evaluated recommended version when the version is not provided | 

### Return type

[**RegisteredExtensionResultDto**](RegisteredExtensionResultDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReleaseNotes

> UpdateReleaseNotes(ctx, extensionName, version).ExtensionReleaseNotes(extensionReleaseNotes).Execute()

Sets or updates the release notes of the specified extension release

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
    extensionName := "extensionName_example" // string | Fully qualified name of the extension
    version := "version_example" // string | Version of the extension release
    extensionReleaseNotes := *openapiclient.NewExtensionReleaseNotes() // ExtensionReleaseNotes | The JSON body of the request. Contains the markdown for release notes (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HubItemsAPI.UpdateReleaseNotes(context.Background(), extensionName, version).ExtensionReleaseNotes(extensionReleaseNotes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.UpdateReleaseNotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | Fully qualified name of the extension | 
**version** | **string** | Version of the extension release | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReleaseNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extensionReleaseNotes** | [**ExtensionReleaseNotes**](ExtensionReleaseNotes.md) | The JSON body of the request. Contains the markdown for release notes | 

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


## UploadMetadataForExtension

> UploadMetadataForExtension(ctx, extensionName).Description(description).Logo(logo).Name(name).Execute()

Create/Replace metadata for an extension that does not have Dynatrace distributed metadata.

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
    extensionName := "extensionName_example" // string | Fully qualified name of the extension
    description := "description_example" // string |  (optional)
    logo := os.NewFile(1234, "some_file") // *os.File | Logo of the extension (optional)
    name := "name_example" // string | If left empty or blank, the extension name will be used as name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HubItemsAPI.UploadMetadataForExtension(context.Background(), extensionName).Description(description).Logo(logo).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HubItemsAPI.UploadMetadataForExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | Fully qualified name of the extension | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadMetadataForExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **description** | **string** |  | 
 **logo** | ***os.File** | Logo of the extension | 
 **name** | **string** | If left empty or blank, the extension name will be used as name | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

