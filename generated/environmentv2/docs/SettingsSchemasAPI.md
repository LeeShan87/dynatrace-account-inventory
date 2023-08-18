# \SettingsSchemasAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableSchemaDefinitions**](SettingsSchemasAPI.md#GetAvailableSchemaDefinitions) | **Get** /settings/schemas | Lists available settings schemas
[**GetSchemaDefinition**](SettingsSchemasAPI.md#GetSchemaDefinition) | **Get** /settings/schemas/{schemaId} | Gets parameters of the specified settings schema



## GetAvailableSchemaDefinitions

> SchemaList GetAvailableSchemaDefinitions(ctx).Execute()

Lists available settings schemas

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
    resp, r, err := apiClient.SettingsSchemasAPI.GetAvailableSchemaDefinitions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsSchemasAPI.GetAvailableSchemaDefinitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAvailableSchemaDefinitions`: SchemaList
    fmt.Fprintf(os.Stdout, "Response from `SettingsSchemasAPI.GetAvailableSchemaDefinitions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableSchemaDefinitionsRequest struct via the builder pattern


### Return type

[**SchemaList**](SchemaList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaDefinition

> SchemaDefinitionRestDto GetSchemaDefinition(ctx, schemaId).SchemaVersion(schemaVersion).Execute()

Gets parameters of the specified settings schema

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
    schemaId := "schemaId_example" // string | The ID of the required schema.
    schemaVersion := "schemaVersion_example" // string | The version of the required schema.    If not set, the most recent version is returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsSchemasAPI.GetSchemaDefinition(context.Background(), schemaId).SchemaVersion(schemaVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsSchemasAPI.GetSchemaDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaDefinition`: SchemaDefinitionRestDto
    fmt.Fprintf(os.Stdout, "Response from `SettingsSchemasAPI.GetSchemaDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemaId** | **string** | The ID of the required schema. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **schemaVersion** | **string** | The version of the required schema.    If not set, the most recent version is returned. | 

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

