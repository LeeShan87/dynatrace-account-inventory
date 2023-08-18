# \JavaScriptMappingFilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteJavaScriptMappingFile**](JavaScriptMappingFilesAPI.md#DeleteJavaScriptMappingFile) | **Delete** /jsMappingFiles/{minifiedJsFileUrl}/{fileType} | Deletes the specified JavaScript mapping file | maturity&#x3D;EARLY_ADOPTER
[**GetJavaScriptMappingFilesMetadata**](JavaScriptMappingFilesAPI.md#GetJavaScriptMappingFilesMetadata) | **Get** /jsMappingFiles | Lists metadata of JavaScript mapping files | maturity&#x3D;EARLY_ADOPTER
[**UpdateJavaScriptMappingFileMetadata**](JavaScriptMappingFilesAPI.md#UpdateJavaScriptMappingFileMetadata) | **Put** /jsMappingFiles/{minifiedJsFileUrl}/{fileType}/metadata | Updates metadata of the specified JavaScript mapping file | maturity&#x3D;EARLY_ADOPTER
[**UploadJavaScriptMappingFile**](JavaScriptMappingFilesAPI.md#UploadJavaScriptMappingFile) | **Put** /jsMappingFiles/{minifiedJsFileUrl}/{fileType} | Uploads new or updates existing JavaScript mapping file | maturity&#x3D;EARLY_ADOPTER



## DeleteJavaScriptMappingFile

> DeleteJavaScriptMappingFile(ctx, minifiedJsFileUrl, fileType).Execute()

Deletes the specified JavaScript mapping file | maturity=EARLY_ADOPTER

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
    minifiedJsFileUrl := "minifiedJsFileUrl_example" // string | The URL of the minified JavaScript file.
    fileType := "fileType_example" // string | The type of the JavaScript mapping file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.JavaScriptMappingFilesAPI.DeleteJavaScriptMappingFile(context.Background(), minifiedJsFileUrl, fileType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JavaScriptMappingFilesAPI.DeleteJavaScriptMappingFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**minifiedJsFileUrl** | **string** | The URL of the minified JavaScript file. | 
**fileType** | **string** | The type of the JavaScript mapping file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteJavaScriptMappingFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetJavaScriptMappingFilesMetadata

> JavaScriptMappingFileListDto GetJavaScriptMappingFilesMetadata(ctx).NextPageKey(nextPageKey).PageSize(pageSize).MinifiedJsFileUrl(minifiedJsFileUrl).FileType(fileType).Execute()

Lists metadata of JavaScript mapping files | maturity=EARLY_ADOPTER

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
    pageSize := int64(789) // int64 | The amount of JavaScript mapping files in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. (optional)
    minifiedJsFileUrl := "minifiedJsFileUrl_example" // string | Filters the resulting set of JavaScript mapping files by the minified JavaScript file URL. Only equals are taken into account. (optional)
    fileType := "fileType_example" // string | Filters the resulting set of JavaScript mapping files by file type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JavaScriptMappingFilesAPI.GetJavaScriptMappingFilesMetadata(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).MinifiedJsFileUrl(minifiedJsFileUrl).FileType(fileType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JavaScriptMappingFilesAPI.GetJavaScriptMappingFilesMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetJavaScriptMappingFilesMetadata`: JavaScriptMappingFileListDto
    fmt.Fprintf(os.Stdout, "Response from `JavaScriptMappingFilesAPI.GetJavaScriptMappingFilesMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetJavaScriptMappingFilesMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of JavaScript mapping files in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. | 
 **minifiedJsFileUrl** | **string** | Filters the resulting set of JavaScript mapping files by the minified JavaScript file URL. Only equals are taken into account. | 
 **fileType** | **string** | Filters the resulting set of JavaScript mapping files by file type. | 

### Return type

[**JavaScriptMappingFileListDto**](JavaScriptMappingFileListDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateJavaScriptMappingFileMetadata

> JavaScriptMappingFileDto UpdateJavaScriptMappingFileMetadata(ctx, minifiedJsFileUrl, fileType).JavaScriptMappingFileMetadataDto(javaScriptMappingFileMetadataDto).Execute()

Updates metadata of the specified JavaScript mapping file | maturity=EARLY_ADOPTER

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
    minifiedJsFileUrl := "minifiedJsFileUrl_example" // string | The URL of the minified JavaScript file.
    fileType := "fileType_example" // string | The type of the JavaScript mapping file.
    javaScriptMappingFileMetadataDto := *openapiclient.NewJavaScriptMappingFileMetadataDto() // JavaScriptMappingFileMetadataDto | The JSON body of the request. Contains updated metadata of the file. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JavaScriptMappingFilesAPI.UpdateJavaScriptMappingFileMetadata(context.Background(), minifiedJsFileUrl, fileType).JavaScriptMappingFileMetadataDto(javaScriptMappingFileMetadataDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JavaScriptMappingFilesAPI.UpdateJavaScriptMappingFileMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateJavaScriptMappingFileMetadata`: JavaScriptMappingFileDto
    fmt.Fprintf(os.Stdout, "Response from `JavaScriptMappingFilesAPI.UpdateJavaScriptMappingFileMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**minifiedJsFileUrl** | **string** | The URL of the minified JavaScript file. | 
**fileType** | **string** | The type of the JavaScript mapping file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateJavaScriptMappingFileMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **javaScriptMappingFileMetadataDto** | [**JavaScriptMappingFileMetadataDto**](JavaScriptMappingFileMetadataDto.md) | The JSON body of the request. Contains updated metadata of the file. | 

### Return type

[**JavaScriptMappingFileDto**](JavaScriptMappingFileDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadJavaScriptMappingFile

> UploadJavaScriptMappingFile(ctx, minifiedJsFileUrl, fileType).File(file).Execute()

Uploads new or updates existing JavaScript mapping file | maturity=EARLY_ADOPTER

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
    minifiedJsFileUrl := "minifiedJsFileUrl_example" // string | The URL of the minified JavaScript file.
    fileType := "fileType_example" // string | The type of the JavaScript mapping file.
    file := os.NewFile(1234, "some_file") // *os.File | JavaScript mapping file to upload.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.JavaScriptMappingFilesAPI.UploadJavaScriptMappingFile(context.Background(), minifiedJsFileUrl, fileType).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JavaScriptMappingFilesAPI.UploadJavaScriptMappingFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**minifiedJsFileUrl** | **string** | The URL of the minified JavaScript file. | 
**fileType** | **string** | The type of the JavaScript mapping file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadJavaScriptMappingFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **file** | ***os.File** | JavaScript mapping file to upload. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

