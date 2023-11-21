# \RUMMobileDeobfuscationAndSymbolicationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdate**](RUMMobileDeobfuscationAndSymbolicationAPI.md#CreateOrUpdate) | **Put** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName} | Upload a symbol file
[**CreateOrUpdatePinning**](RUMMobileDeobfuscationAndSymbolicationAPI.md#CreateOrUpdatePinning) | **Put** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName}/pinning | Pins/unpins all symbol files of an app version
[**DeleteApp**](RUMMobileDeobfuscationAndSymbolicationAPI.md#DeleteApp) | **Delete** /symfiles/{applicationId} | Deletes all symbol file belonging for a mobile app
[**DeleteSingleFile**](RUMMobileDeobfuscationAndSymbolicationAPI.md#DeleteSingleFile) | **Delete** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName} | Deletes the symbol file belonging to the specified app, OS, and an app version
[**GetAll**](RUMMobileDeobfuscationAndSymbolicationAPI.md#GetAll) | **Get** /symfiles | Lists metadata of all symbol files from the Symbol File Store
[**GetAllPerApplication**](RUMMobileDeobfuscationAndSymbolicationAPI.md#GetAllPerApplication) | **Get** /symfiles/{applicationId} | Lists metadata of all symbol files for a mobile app from the Symbol File Store
[**GetDssClientUrl**](RUMMobileDeobfuscationAndSymbolicationAPI.md#GetDssClientUrl) | **Get** /symfiles/dtxdss-download | Gets a download link for the latest version of the DTXDssClient
[**GetInfo**](RUMMobileDeobfuscationAndSymbolicationAPI.md#GetInfo) | **Get** /symfiles/info | Gets information about symbol files storage
[**GetSingle**](RUMMobileDeobfuscationAndSymbolicationAPI.md#GetSingle) | **Get** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName} | Gets metadata of the symbol file for an app version
[**GetSupportedVersion**](RUMMobileDeobfuscationAndSymbolicationAPI.md#GetSupportedVersion) | **Get** /symfiles/ios/supportedversion | Gets the supported file format version for iOS symbol files
[**ValidatePinning**](RUMMobileDeobfuscationAndSymbolicationAPI.md#ValidatePinning) | **Put** /symfiles/{applicationId}/{packageName}/{os}/{versionCode}/{versionName}/pinning/validator | Validates the payload for the &#x60;PUT /{applicationId}/{packageName}/{os}/{versionName}/pinning&#x60; request



## CreateOrUpdate

> CreateOrUpdate(ctx, applicationId, packageName, os, versionCode, versionName).Body(body).ContentType(contentType).Execute()

Upload a symbol file



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required mobile app.
    packageName := "packageName_example" // string | The CFBundleIdentifier (iOS) or the package name (Android) of the required mobile app.
    os := "os_example" // string | The operating system of the required app.
    versionCode := "versionCode_example" // string | The version code (Android) / CFBundleVersion (iOS) of the required app.
    versionName := "versionName_example" // string | The version name (Android) / CFBundleShortVersionString (iOS) of the required app.
    body := os.NewFile(1234, "some_file") // *os.File | The file to be uploaded: a proguard file (*.txt) for Android or the zip file produced by the DTXDssClient provided with the OneAgent for iOS. 
    contentType := "contentType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.CreateOrUpdate(context.Background(), applicationId, packageName, os, versionCode, versionName).Body(body).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.CreateOrUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required mobile app. | 
**packageName** | **string** | The CFBundleIdentifier (iOS) or the package name (Android) of the required mobile app. | 
**os** | **string** | The operating system of the required app. | 
**versionCode** | **string** | The version code (Android) / CFBundleVersion (iOS) of the required app. | 
**versionName** | **string** | The version name (Android) / CFBundleShortVersionString (iOS) of the required app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | ***os.File** | The file to be uploaded: a proguard file (*.txt) for Android or the zip file produced by the DTXDssClient provided with the OneAgent for iOS.  | 
 **contentType** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/x-compressed, application/x-zip-compressed, application/zip, text/plain
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrUpdatePinning

> CreateOrUpdatePinning(ctx, applicationId, packageName, os, versionCode, versionName).SymbolFilePinning(symbolFilePinning).Execute()

Pins/unpins all symbol files of an app version



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The operating system of the required app.
    packageName := "packageName_example" // string | The CFBundleIdentifier (iOS) or the package name (Android) of the required mobile app.
    os := "os_example" // string | The operating system of the required app.
    versionCode := "versionCode_example" // string | The version code (Android) / CFBundleVersion (iOS) of the required app.
    versionName := "versionName_example" // string | The version name (Android) / CFBundleShortVersionString (iOS) of the required app.
    symbolFilePinning := *openapiclient.NewSymbolFilePinning(false) // SymbolFilePinning | The JSON body of the request. Contains the pinning status to set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.CreateOrUpdatePinning(context.Background(), applicationId, packageName, os, versionCode, versionName).SymbolFilePinning(symbolFilePinning).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.CreateOrUpdatePinning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The operating system of the required app. | 
**packageName** | **string** | The CFBundleIdentifier (iOS) or the package name (Android) of the required mobile app. | 
**os** | **string** | The operating system of the required app. | 
**versionCode** | **string** | The version code (Android) / CFBundleVersion (iOS) of the required app. | 
**versionName** | **string** | The version name (Android) / CFBundleShortVersionString (iOS) of the required app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdatePinningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **symbolFilePinning** | [**SymbolFilePinning**](SymbolFilePinning.md) | The JSON body of the request. Contains the pinning status to set. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApp

> DeleteApp(ctx, applicationId).Execute()

Deletes all symbol file belonging for a mobile app

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required mobile app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.DeleteApp(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.DeleteApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required mobile app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSingleFile

> DeleteSingleFile(ctx, applicationId, packageName, os, versionCode, versionName).Execute()

Deletes the symbol file belonging to the specified app, OS, and an app version

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required mobile app.
    packageName := "packageName_example" // string | The CFBundleIdentifier (iOS) or the package name (Android) of the required mobile app.
    os := "os_example" // string | The operating system of the required app.
    versionCode := "versionCode_example" // string | The version code (Android) / CFBundleVersion (iOS) of the required app.
    versionName := "versionName_example" // string | The version name (Android) / CFBundleShortVersionString (iOS) of the required app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.DeleteSingleFile(context.Background(), applicationId, packageName, os, versionCode, versionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.DeleteSingleFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required mobile app. | 
**packageName** | **string** | The CFBundleIdentifier (iOS) or the package name (Android) of the required mobile app. | 
**os** | **string** | The operating system of the required app. | 
**versionCode** | **string** | The version code (Android) / CFBundleVersion (iOS) of the required app. | 
**versionName** | **string** | The version name (Android) / CFBundleShortVersionString (iOS) of the required app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSingleFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAll

> SymbolFileList GetAll(ctx).Execute()

Lists metadata of all symbol files from the Symbol File Store



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
    resp, r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.GetAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.GetAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAll`: SymbolFileList
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileDeobfuscationAndSymbolicationAPI.GetAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRequest struct via the builder pattern


### Return type

[**SymbolFileList**](SymbolFileList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPerApplication

> SymbolFileList GetAllPerApplication(ctx, applicationId).Execute()

Lists metadata of all symbol files for a mobile app from the Symbol File Store



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The Dynatrace entity ID of the required mobile app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.GetAllPerApplication(context.Background(), applicationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.GetAllPerApplication``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPerApplication`: SymbolFileList
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileDeobfuscationAndSymbolicationAPI.GetAllPerApplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The Dynatrace entity ID of the required mobile app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPerApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SymbolFileList**](SymbolFileList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDssClientUrl

> SymbolFileClientLinkDto GetDssClientUrl(ctx).Execute()

Gets a download link for the latest version of the DTXDssClient

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
    resp, r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.GetDssClientUrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.GetDssClientUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDssClientUrl`: SymbolFileClientLinkDto
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileDeobfuscationAndSymbolicationAPI.GetDssClientUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDssClientUrlRequest struct via the builder pattern


### Return type

[**SymbolFileClientLinkDto**](SymbolFileClientLinkDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInfo

> SymbolFileStorageInfo GetInfo(ctx).Execute()

Gets information about symbol files storage

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
    resp, r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.GetInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.GetInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInfo`: SymbolFileStorageInfo
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileDeobfuscationAndSymbolicationAPI.GetInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInfoRequest struct via the builder pattern


### Return type

[**SymbolFileStorageInfo**](SymbolFileStorageInfo.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingle

> SymbolFile GetSingle(ctx, applicationId, packageName, os, versionCode, versionName).Execute()

Gets metadata of the symbol file for an app version



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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required mobile app.
    packageName := "packageName_example" // string | The CFBundleIdentifier (iOS) or the package name (Android) of the required mobile app.
    os := "os_example" // string | The operating system of the required app.
    versionCode := "versionCode_example" // string | The version code (Android) / CFBundleVersion (iOS) of the required app.
    versionName := "versionName_example" // string | The version name (Android) / CFBundleShortVersionString (iOS) of the required app.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.GetSingle(context.Background(), applicationId, packageName, os, versionCode, versionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.GetSingle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingle`: SymbolFile
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileDeobfuscationAndSymbolicationAPI.GetSingle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required mobile app. | 
**packageName** | **string** | The CFBundleIdentifier (iOS) or the package name (Android) of the required mobile app. | 
**os** | **string** | The operating system of the required app. | 
**versionCode** | **string** | The version code (Android) / CFBundleVersion (iOS) of the required app. | 
**versionName** | **string** | The version name (Android) / CFBundleShortVersionString (iOS) of the required app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**SymbolFile**](SymbolFile.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedVersion

> SupportedVersion GetSupportedVersion(ctx).Execute()

Gets the supported file format version for iOS symbol files

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
    resp, r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.GetSupportedVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.GetSupportedVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedVersion`: SupportedVersion
    fmt.Fprintf(os.Stdout, "Response from `RUMMobileDeobfuscationAndSymbolicationAPI.GetSupportedVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedVersionRequest struct via the builder pattern


### Return type

[**SupportedVersion**](SupportedVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidatePinning

> ValidatePinning(ctx, applicationId, packageName, os, versionCode, versionName).SymbolFilePinning(symbolFilePinning).Execute()

Validates the payload for the `PUT /{applicationId}/{packageName}/{os}/{versionName}/pinning` request

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
    applicationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required mobile app.
    packageName := "packageName_example" // string | The CFBundleIdentifier (iOS) or the package name (Android) of the required mobile app.
    os := "os_example" // string | The operating system of the required app.
    versionCode := "versionCode_example" // string | The version code (Android) / CFBundleVersion (iOS) of the required app.
    versionName := "versionName_example" // string | The version name (Android) / CFBundleShortVersionString (iOS) of the required app.
    symbolFilePinning := *openapiclient.NewSymbolFilePinning(false) // SymbolFilePinning | The JSON body of the request. Contains the pinning status to set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMMobileDeobfuscationAndSymbolicationAPI.ValidatePinning(context.Background(), applicationId, packageName, os, versionCode, versionName).SymbolFilePinning(symbolFilePinning).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMMobileDeobfuscationAndSymbolicationAPI.ValidatePinning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** | The ID of the required mobile app. | 
**packageName** | **string** | The CFBundleIdentifier (iOS) or the package name (Android) of the required mobile app. | 
**os** | **string** | The operating system of the required app. | 
**versionCode** | **string** | The version code (Android) / CFBundleVersion (iOS) of the required app. | 
**versionName** | **string** | The version name (Android) / CFBundleShortVersionString (iOS) of the required app. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidatePinningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **symbolFilePinning** | [**SymbolFilePinning**](SymbolFilePinning.md) | The JSON body of the request. Contains the pinning status to set. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

