# \Extensions20API

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateExtensionEnvironmentConfiguration**](Extensions20API.md#ActivateExtensionEnvironmentConfiguration) | **Post** /extensions/{extensionName}/environmentConfiguration | Activates the environment configuration from the specified version of the extension 2.0
[**CreateMonitoringConfiguration**](Extensions20API.md#CreateMonitoringConfiguration) | **Post** /extensions/{extensionName}/monitoringConfigurations | Creates new monitoring configuration for the specified extension 2.0
[**DeleteEnvironmentConfiguration**](Extensions20API.md#DeleteEnvironmentConfiguration) | **Delete** /extensions/{extensionName}/environmentConfiguration | Deactivates the environment configuration of the specified extension 2.0
[**ExtensionConfigurationSchema**](Extensions20API.md#ExtensionConfigurationSchema) | **Get** /extensions/{extensionName}/{extensionVersion}/schema | Gets the configuration schema of the specified version of the extension 2.0
[**ExtensionDetails**](Extensions20API.md#ExtensionDetails) | **Get** /extensions/{extensionName}/{extensionVersion} | Gets details of the specified version of the extension 2.0
[**ExtensionMonitoringConfigurations**](Extensions20API.md#ExtensionMonitoringConfigurations) | **Get** /extensions/{extensionName}/monitoringConfigurations | Lists all the monitoring configurations of the specified extension 2.0
[**GetActiveEnvironmentConfiguration**](Extensions20API.md#GetActiveEnvironmentConfiguration) | **Get** /extensions/{extensionName}/environmentConfiguration | Gets the active environment configuration version of the specified extension 2.0
[**GetEnvironmentConfigurationAssetsInfo**](Extensions20API.md#GetEnvironmentConfigurationAssetsInfo) | **Get** /extensions/{extensionName}/environmentConfiguration/assets | Gets the information about assets in an active extension 2.0
[**GetEnvironmentConfigurationEvents**](Extensions20API.md#GetEnvironmentConfigurationEvents) | **Get** /extensions/{extensionName}/environmentConfiguration/events | List of the latest extension environment configuration events
[**GetExtensionMonitoringConfigurationEvents**](Extensions20API.md#GetExtensionMonitoringConfigurationEvents) | **Get** /extensions/{extensionName}/monitoringConfigurations/{configurationId}/events | Gets the list of the events linked to specific monitoring configuration
[**GetExtensionMonitoringConfigurationStatus**](Extensions20API.md#GetExtensionMonitoringConfigurationStatus) | **Get** /extensions/{extensionName}/monitoringConfigurations/{configurationId}/status | Gets the most recent status of the execution of given monitoring configuration
[**GetSchemaFile**](Extensions20API.md#GetSchemaFile) | **Get** /extensions/schemas/{schemaVersion}/{fileName} | Gets the extension 2.0 schema file in the specified version
[**ListExtensionVersions**](Extensions20API.md#ListExtensionVersions) | **Get** /extensions/{extensionName} | Lists all versions of the extension 2.0
[**ListExtensions**](Extensions20API.md#ListExtensions) | **Get** /extensions | Lists all the extensions 2.0 available in your environment
[**ListSchemaFiles**](Extensions20API.md#ListSchemaFiles) | **Get** /extensions/schemas/{schemaVersion} | Lists all the files available for the specified extension 2.0 schema version
[**ListSchemas**](Extensions20API.md#ListSchemas) | **Get** /extensions/schemas | Lists all the extension 2.0 schemas versions available in your environment
[**MonitoringConfigurationDetails**](Extensions20API.md#MonitoringConfigurationDetails) | **Get** /extensions/{extensionName}/monitoringConfigurations/{configurationId} | Gets the details of the specified monitoring configuration
[**RemoveExtension**](Extensions20API.md#RemoveExtension) | **Delete** /extensions/{extensionName}/{extensionVersion} | Deletes the specified version of the extension 2.0
[**RemoveMonitoringConfiguration**](Extensions20API.md#RemoveMonitoringConfiguration) | **Delete** /extensions/{extensionName}/monitoringConfigurations/{configurationId} | Deletes the specified monitoring configuration
[**UpdateExtensionEnvironmentConfiguration**](Extensions20API.md#UpdateExtensionEnvironmentConfiguration) | **Put** /extensions/{extensionName}/environmentConfiguration | Updates the active environment configuration version of the extension 2.0
[**UpdateMonitoringConfiguration**](Extensions20API.md#UpdateMonitoringConfiguration) | **Put** /extensions/{extensionName}/monitoringConfigurations/{configurationId} | Updates the specified monitoring configuration
[**UploadExtension**](Extensions20API.md#UploadExtension) | **Post** /extensions | Uploads or verifies a new extension 2.0



## ActivateExtensionEnvironmentConfiguration

> ExtensionEnvironmentConfigurationVersion ActivateExtensionEnvironmentConfiguration(ctx, extensionName).ExtensionEnvironmentConfigurationVersion(extensionEnvironmentConfigurationVersion).Execute()

Activates the environment configuration from the specified version of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    extensionEnvironmentConfigurationVersion := *openapiclient.NewExtensionEnvironmentConfigurationVersion("1.2.3") // ExtensionEnvironmentConfigurationVersion | The version of the requested environment configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.ActivateExtensionEnvironmentConfiguration(context.Background(), extensionName).ExtensionEnvironmentConfigurationVersion(extensionEnvironmentConfigurationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.ActivateExtensionEnvironmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateExtensionEnvironmentConfiguration`: ExtensionEnvironmentConfigurationVersion
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.ActivateExtensionEnvironmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateExtensionEnvironmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionEnvironmentConfigurationVersion** | [**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md) | The version of the requested environment configuration. | 

### Return type

[**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMonitoringConfiguration

> []MonitoringConfigurationResponse CreateMonitoringConfiguration(ctx, extensionName).MonitoringConfigurationDto(monitoringConfigurationDto).Execute()

Creates new monitoring configuration for the specified extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    monitoringConfigurationDto := []openapiclient.MonitoringConfigurationDto{*openapiclient.NewMonitoringConfigurationDto("HOST-D3A3C5A146830A79")} // []MonitoringConfigurationDto | JSON body of the request, containing monitoring configuration parameters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.CreateMonitoringConfiguration(context.Background(), extensionName).MonitoringConfigurationDto(monitoringConfigurationDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.CreateMonitoringConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMonitoringConfiguration`: []MonitoringConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.CreateMonitoringConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMonitoringConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **monitoringConfigurationDto** | [**[]MonitoringConfigurationDto**](MonitoringConfigurationDto.md) | JSON body of the request, containing monitoring configuration parameters. | 

### Return type

[**[]MonitoringConfigurationResponse**](MonitoringConfigurationResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentConfiguration

> ExtensionEnvironmentConfigurationVersion DeleteEnvironmentConfiguration(ctx, extensionName).Execute()

Deactivates the environment configuration of the specified extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.DeleteEnvironmentConfiguration(context.Background(), extensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.DeleteEnvironmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEnvironmentConfiguration`: ExtensionEnvironmentConfigurationVersion
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.DeleteEnvironmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtensionConfigurationSchema

> SchemaDefinitionRestDto ExtensionConfigurationSchema(ctx, extensionName, extensionVersion).Execute()

Gets the configuration schema of the specified version of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    extensionVersion := "extensionVersion_example" // string | The version of the requested extension 2.0

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.ExtensionConfigurationSchema(context.Background(), extensionName, extensionVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.ExtensionConfigurationSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtensionConfigurationSchema`: SchemaDefinitionRestDto
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.ExtensionConfigurationSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**extensionVersion** | **string** | The version of the requested extension 2.0 | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtensionConfigurationSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SchemaDefinitionRestDto**](SchemaDefinitionRestDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtensionDetails

> Extension ExtensionDetails(ctx, extensionName, extensionVersion).Accept(accept).Execute()

Gets details of the specified version of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    extensionVersion := "extensionVersion_example" // string | The version of the requested extension 2.0
    accept := "accept_example" // string | Accept header. Specifies part of the extension 2.0 that will be returned:  * application/json; charset=utf-8 - returns extension 2.0 metadata in JSON  * application/octet-stream - returns extension 2.0 zip package stored on the server. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.ExtensionDetails(context.Background(), extensionName, extensionVersion).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.ExtensionDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtensionDetails`: Extension
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.ExtensionDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**extensionVersion** | **string** | The version of the requested extension 2.0 | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtensionDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accept** | **string** | Accept header. Specifies part of the extension 2.0 that will be returned:  * application/json; charset&#x3D;utf-8 - returns extension 2.0 metadata in JSON  * application/octet-stream - returns extension 2.0 zip package stored on the server. | 

### Return type

[**Extension**](Extension.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExtensionMonitoringConfigurations

> ExtensionMonitoringConfigurationsList ExtensionMonitoringConfigurations(ctx, extensionName).NextPageKey(nextPageKey).PageSize(pageSize).Version(version).Active(active).Execute()

Lists all the monitoring configurations of the specified extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. (optional)
    version := "version_example" // string | Filters the resulting set of configurations by extension 2.0 version. (optional)
    active := true // bool | Filters the resulting set of configurations by the active state. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.ExtensionMonitoringConfigurations(context.Background(), extensionName).NextPageKey(nextPageKey).PageSize(pageSize).Version(version).Active(active).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.ExtensionMonitoringConfigurations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExtensionMonitoringConfigurations`: ExtensionMonitoringConfigurationsList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.ExtensionMonitoringConfigurations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtensionMonitoringConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. | 
 **version** | **string** | Filters the resulting set of configurations by extension 2.0 version. | 
 **active** | **bool** | Filters the resulting set of configurations by the active state. | 

### Return type

[**ExtensionMonitoringConfigurationsList**](ExtensionMonitoringConfigurationsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveEnvironmentConfiguration

> ExtensionEnvironmentConfigurationVersion GetActiveEnvironmentConfiguration(ctx, extensionName).Execute()

Gets the active environment configuration version of the specified extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.GetActiveEnvironmentConfiguration(context.Background(), extensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.GetActiveEnvironmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveEnvironmentConfiguration`: ExtensionEnvironmentConfigurationVersion
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.GetActiveEnvironmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveEnvironmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentConfigurationAssetsInfo

> ExtensionAssetsDto GetEnvironmentConfigurationAssetsInfo(ctx, extensionName).Execute()

Gets the information about assets in an active extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.GetEnvironmentConfigurationAssetsInfo(context.Background(), extensionName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.GetEnvironmentConfigurationAssetsInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentConfigurationAssetsInfo`: ExtensionAssetsDto
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.GetEnvironmentConfigurationAssetsInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentConfigurationAssetsInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExtensionAssetsDto**](ExtensionAssetsDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentConfigurationEvents

> ExtensionEventsList GetEnvironmentConfigurationEvents(ctx, extensionName).From(from).To(to).Content(content).Status(status).Execute()

List of the latest extension environment configuration events

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two hours is used (`now-2h`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    content := "content_example" // string | Content of the event (optional)
    status := "status_example" // string | Status of the event (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.GetEnvironmentConfigurationEvents(context.Background(), extensionName).From(from).To(to).Content(content).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.GetEnvironmentConfigurationEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEnvironmentConfigurationEvents`: ExtensionEventsList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.GetEnvironmentConfigurationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentConfigurationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two hours is used (&#x60;now-2h&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **content** | **string** | Content of the event | 
 **status** | **string** | Status of the event | 

### Return type

[**ExtensionEventsList**](ExtensionEventsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionMonitoringConfigurationEvents

> ExtensionEventsList GetExtensionMonitoringConfigurationEvents(ctx, extensionName, configurationId).From(from).To(to).DtEntityHost(dtEntityHost).DtActiveGateId(dtActiveGateId).DtExtensionDs(dtExtensionDs).Content(content).Status(status).Execute()

Gets the list of the events linked to specific monitoring configuration

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    configurationId := "configurationId_example" // string | The ID of the requested monitoring configuration.
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two hours is used (`now-2h`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    dtEntityHost := "dtEntityHost_example" // string | Host that uses this monitoring configuration.  Example: `HOST-ABCDEF0123456789` (optional)
    dtActiveGateId := "dtActiveGateId_example" // string | Hexadecimal ID of Active Gate that uses this monitoring configuration.  Example: `0x1a2b3c4d` (optional)
    dtExtensionDs := "dtExtensionDs_example" // string | Data source that uses this monitoring configuration.  Example: `snmp` (optional)
    content := "content_example" // string | Content of the event (optional)
    status := "status_example" // string | Status of the event (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.GetExtensionMonitoringConfigurationEvents(context.Background(), extensionName, configurationId).From(from).To(to).DtEntityHost(dtEntityHost).DtActiveGateId(dtActiveGateId).DtExtensionDs(dtExtensionDs).Content(content).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.GetExtensionMonitoringConfigurationEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtensionMonitoringConfigurationEvents`: ExtensionEventsList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.GetExtensionMonitoringConfigurationEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**configurationId** | **string** | The ID of the requested monitoring configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionMonitoringConfigurationEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two hours is used (&#x60;now-2h&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **dtEntityHost** | **string** | Host that uses this monitoring configuration.  Example: &#x60;HOST-ABCDEF0123456789&#x60; | 
 **dtActiveGateId** | **string** | Hexadecimal ID of Active Gate that uses this monitoring configuration.  Example: &#x60;0x1a2b3c4d&#x60; | 
 **dtExtensionDs** | **string** | Data source that uses this monitoring configuration.  Example: &#x60;snmp&#x60; | 
 **content** | **string** | Content of the event | 
 **status** | **string** | Status of the event | 

### Return type

[**ExtensionEventsList**](ExtensionEventsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExtensionMonitoringConfigurationStatus

> ExtensionStatusDto GetExtensionMonitoringConfigurationStatus(ctx, extensionName, configurationId).Execute()

Gets the most recent status of the execution of given monitoring configuration

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    configurationId := "configurationId_example" // string | The ID of the requested monitoring configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.GetExtensionMonitoringConfigurationStatus(context.Background(), extensionName, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.GetExtensionMonitoringConfigurationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExtensionMonitoringConfigurationStatus`: ExtensionStatusDto
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.GetExtensionMonitoringConfigurationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**configurationId** | **string** | The ID of the requested monitoring configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionMonitoringConfigurationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExtensionStatusDto**](ExtensionStatusDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaFile

> map[string]interface{} GetSchemaFile(ctx, schemaVersion, fileName).Execute()

Gets the extension 2.0 schema file in the specified version

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
    schemaVersion := "schemaVersion_example" // string | The version of the schema.
    fileName := "fileName_example" // string | The name of the schema file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.GetSchemaFile(context.Background(), schemaVersion, fileName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.GetSchemaFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaFile`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.GetSchemaFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaVersion** | **string** | The version of the schema. | 
**fileName** | **string** | The name of the schema file. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtensionVersions

> ExtensionList ListExtensionVersions(ctx, extensionName).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Lists all versions of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.ListExtensionVersions(context.Background(), extensionName).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.ListExtensionVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExtensionVersions`: ExtensionList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.ListExtensionVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExtensionVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. | 

### Return type

[**ExtensionList**](ExtensionList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExtensions

> ExtensionList ListExtensions(ctx).NextPageKey(nextPageKey).PageSize(pageSize).Name(name).Execute()

Lists all the extensions 2.0 available in your environment

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
    pageSize := int64(789) // int64 | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. (optional)
    name := "name_example" // string | Filters the resulting set of extensions 2.0 by name. You can specify a partial name. In that case, the CONTAINS operator is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.ListExtensions(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.ListExtensions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListExtensions`: ExtensionList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.ListExtensions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListExtensionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of extensions in a single response payload.   The maximal allowed page size is 100.   If not set, 20 is used. | 
 **name** | **string** | Filters the resulting set of extensions 2.0 by name. You can specify a partial name. In that case, the CONTAINS operator is used. | 

### Return type

[**ExtensionList**](ExtensionList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchemaFiles

> SchemaFiles ListSchemaFiles(ctx, schemaVersion).Accept(accept).Execute()

Lists all the files available for the specified extension 2.0 schema version

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
    schemaVersion := "schemaVersion_example" // string | The version of the schema.
    accept := "accept_example" // string | Accept header. Specifies part of the extension 2.0 that will be returned:  * application/json; charset=utf-8 - returns extension 2.0 metadata in JSON  * application/octet-stream - returns extension 2.0 zip package stored on the server. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.ListSchemaFiles(context.Background(), schemaVersion).Accept(accept).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.ListSchemaFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchemaFiles`: SchemaFiles
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.ListSchemaFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaVersion** | **string** | The version of the schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSchemaFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accept** | **string** | Accept header. Specifies part of the extension 2.0 that will be returned:  * application/json; charset&#x3D;utf-8 - returns extension 2.0 metadata in JSON  * application/octet-stream - returns extension 2.0 zip package stored on the server. | 

### Return type

[**SchemaFiles**](SchemaFiles.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8, application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSchemas

> SchemasList ListSchemas(ctx).Execute()

Lists all the extension 2.0 schemas versions available in your environment

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
    resp, r, err := apiClient.Extensions20API.ListSchemas(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.ListSchemas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSchemas`: SchemasList
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.ListSchemas`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSchemasRequest struct via the builder pattern


### Return type

[**SchemasList**](SchemasList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MonitoringConfigurationDetails

> ExtensionMonitoringConfiguration MonitoringConfigurationDetails(ctx, extensionName, configurationId).Execute()

Gets the details of the specified monitoring configuration

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    configurationId := "configurationId_example" // string | The ID of the requested monitoring configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.MonitoringConfigurationDetails(context.Background(), extensionName, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.MonitoringConfigurationDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MonitoringConfigurationDetails`: ExtensionMonitoringConfiguration
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.MonitoringConfigurationDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**configurationId** | **string** | The ID of the requested monitoring configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMonitoringConfigurationDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ExtensionMonitoringConfiguration**](ExtensionMonitoringConfiguration.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveExtension

> Extension RemoveExtension(ctx, extensionName, extensionVersion).Execute()

Deletes the specified version of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    extensionVersion := "extensionVersion_example" // string | The version of the requested extension 2.0

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.RemoveExtension(context.Background(), extensionName, extensionVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.RemoveExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveExtension`: Extension
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.RemoveExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**extensionVersion** | **string** | The version of the requested extension 2.0 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Extension**](Extension.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMonitoringConfiguration

> RemoveMonitoringConfiguration(ctx, extensionName, configurationId).Execute()

Deletes the specified monitoring configuration

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    configurationId := "configurationId_example" // string | The ID of the requested monitoring configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.Extensions20API.RemoveMonitoringConfiguration(context.Background(), extensionName, configurationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.RemoveMonitoringConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**configurationId** | **string** | The ID of the requested monitoring configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMonitoringConfigurationRequest struct via the builder pattern


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


## UpdateExtensionEnvironmentConfiguration

> ExtensionEnvironmentConfigurationVersion UpdateExtensionEnvironmentConfiguration(ctx, extensionName).ExtensionEnvironmentConfigurationVersion(extensionEnvironmentConfigurationVersion).Execute()

Updates the active environment configuration version of the extension 2.0

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    extensionEnvironmentConfigurationVersion := *openapiclient.NewExtensionEnvironmentConfigurationVersion("1.2.3") // ExtensionEnvironmentConfigurationVersion | The version of the requested environment configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.UpdateExtensionEnvironmentConfiguration(context.Background(), extensionName).ExtensionEnvironmentConfigurationVersion(extensionEnvironmentConfigurationVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.UpdateExtensionEnvironmentConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateExtensionEnvironmentConfiguration`: ExtensionEnvironmentConfigurationVersion
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.UpdateExtensionEnvironmentConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtensionEnvironmentConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extensionEnvironmentConfigurationVersion** | [**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md) | The version of the requested environment configuration. | 

### Return type

[**ExtensionEnvironmentConfigurationVersion**](ExtensionEnvironmentConfigurationVersion.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMonitoringConfiguration

> MonitoringConfigurationResponse UpdateMonitoringConfiguration(ctx, extensionName, configurationId).MonitoringConfigurationUpdateDto(monitoringConfigurationUpdateDto).Execute()

Updates the specified monitoring configuration

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
    extensionName := "extensionName_example" // string | The name of the requested extension 2.0.
    configurationId := "configurationId_example" // string | The ID of the requested monitoring configuration.
    monitoringConfigurationUpdateDto := *openapiclient.NewMonitoringConfigurationUpdateDto() // MonitoringConfigurationUpdateDto | JSON body of the request, containing monitoring configuration parameters.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.UpdateMonitoringConfiguration(context.Background(), extensionName, configurationId).MonitoringConfigurationUpdateDto(monitoringConfigurationUpdateDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.UpdateMonitoringConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMonitoringConfiguration`: MonitoringConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.UpdateMonitoringConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionName** | **string** | The name of the requested extension 2.0. | 
**configurationId** | **string** | The ID of the requested monitoring configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMonitoringConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **monitoringConfigurationUpdateDto** | [**MonitoringConfigurationUpdateDto**](MonitoringConfigurationUpdateDto.md) | JSON body of the request, containing monitoring configuration parameters. | 

### Return type

[**MonitoringConfigurationResponse**](MonitoringConfigurationResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadExtension

> ExtensionUploadResponseDto UploadExtension(ctx).Body(body).ValidateOnly(validateOnly).Execute()

Uploads or verifies a new extension 2.0

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
    body := os.NewFile(1234, "some_file") // *os.File | Extension 2.0 zip file to upload.
    validateOnly := true // bool | Only run validation but do not persist the extension even if validation was successful. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Extensions20API.UploadExtension(context.Background()).Body(body).ValidateOnly(validateOnly).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Extensions20API.UploadExtension``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadExtension`: ExtensionUploadResponseDto
    fmt.Fprintf(os.Stdout, "Response from `Extensions20API.UploadExtension`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** | Extension 2.0 zip file to upload. | 
 **validateOnly** | **bool** | Only run validation but do not persist the extension even if validation was successful. | [default to false]

### Return type

[**ExtensionUploadResponseDto**](ExtensionUploadResponseDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/octet-stream, multipart/form-data
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

