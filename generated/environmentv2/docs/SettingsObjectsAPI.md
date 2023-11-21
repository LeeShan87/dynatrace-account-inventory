# \SettingsObjectsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSettingsObjectByObjectId**](SettingsObjectsAPI.md#DeleteSettingsObjectByObjectId) | **Delete** /settings/objects/{objectId} | Deletes the specified settings object
[**GetEffectiveSettingsValues**](SettingsObjectsAPI.md#GetEffectiveSettingsValues) | **Get** /settings/effectiveValues | Lists effective settings values
[**GetSettingsObjectByObjectId**](SettingsObjectsAPI.md#GetSettingsObjectByObjectId) | **Get** /settings/objects/{objectId} | Gets the specified settings object
[**GetSettingsObjects**](SettingsObjectsAPI.md#GetSettingsObjects) | **Get** /settings/objects | Lists persisted settings objects
[**PostSettingsObjects**](SettingsObjectsAPI.md#PostSettingsObjects) | **Post** /settings/objects | Creates a new settings object
[**PutSettingsObjectByObjectId**](SettingsObjectsAPI.md#PutSettingsObjectByObjectId) | **Put** /settings/objects/{objectId} | Updates an existing settings object



## DeleteSettingsObjectByObjectId

> DeleteSettingsObjectByObjectId(ctx, objectId).UpdateToken(updateToken).Execute()

Deletes the specified settings object

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
    objectId := "objectId_example" // string | The ID of the required settings object.
    updateToken := "updateToken_example" // string | The update token of the object. You can use it to detect simultaneous modifications by different users.   It is generated upon retrieval (GET requests). If set on update (PUT request) or deletion, the update/deletion will be allowed only if there wasn't any change between the retrieval and the update.   If omitted on update/deletion, the operation overrides the current value or deletes it without any checks. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SettingsObjectsAPI.DeleteSettingsObjectByObjectId(context.Background(), objectId).UpdateToken(updateToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsAPI.DeleteSettingsObjectByObjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the required settings object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSettingsObjectByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateToken** | **string** | The update token of the object. You can use it to detect simultaneous modifications by different users.   It is generated upon retrieval (GET requests). If set on update (PUT request) or deletion, the update/deletion will be allowed only if there wasn&#39;t any change between the retrieval and the update.   If omitted on update/deletion, the operation overrides the current value or deletes it without any checks. | 

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


## GetEffectiveSettingsValues

> EffectiveSettingsValuesList GetEffectiveSettingsValues(ctx).SchemaIds(schemaIds).Scope(scope).Fields(fields).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Lists effective settings values



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
    schemaIds := "schemaIds_example" // string | A list of comma-separated schema IDs to which the requested objects belong.   Only considered on load of the first page, when the **nextPageKey** is not set. (optional)
    scope := "scope_example" // string | The scope that the requested objects target.   The selection only matches objects directly targeting the specified scope. For example, `environment` will not match objects that target a host within environment. For more details, please see [Dynatrace Documentation](https://dt-url.net/ky03459).   To load the first page, when the **nextPageKey** is not set, this parameter is required.  (optional)
    fields := "fields_example" // string | A list of fields to be included to the response. The provided set of fields replaces the default set.    Specify the required top-level fields, separated by commas (for example, `origin,value`).  Supported fields: `summary`, `searchSummary`, `created`, `modified`, `createdBy`, `modifiedBy`, `author`, `origin`, `schemaId`, `schemaVersion`, `value`, `externalId`. (optional) (default to "origin,value")
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of settings objects in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsObjectsAPI.GetEffectiveSettingsValues(context.Background()).SchemaIds(schemaIds).Scope(scope).Fields(fields).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsAPI.GetEffectiveSettingsValues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEffectiveSettingsValues`: EffectiveSettingsValuesList
    fmt.Fprintf(os.Stdout, "Response from `SettingsObjectsAPI.GetEffectiveSettingsValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEffectiveSettingsValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaIds** | **string** | A list of comma-separated schema IDs to which the requested objects belong.   Only considered on load of the first page, when the **nextPageKey** is not set. | 
 **scope** | **string** | The scope that the requested objects target.   The selection only matches objects directly targeting the specified scope. For example, &#x60;environment&#x60; will not match objects that target a host within environment. For more details, please see [Dynatrace Documentation](https://dt-url.net/ky03459).   To load the first page, when the **nextPageKey** is not set, this parameter is required.  | 
 **fields** | **string** | A list of fields to be included to the response. The provided set of fields replaces the default set.    Specify the required top-level fields, separated by commas (for example, &#x60;origin,value&#x60;).  Supported fields: &#x60;summary&#x60;, &#x60;searchSummary&#x60;, &#x60;created&#x60;, &#x60;modified&#x60;, &#x60;createdBy&#x60;, &#x60;modifiedBy&#x60;, &#x60;author&#x60;, &#x60;origin&#x60;, &#x60;schemaId&#x60;, &#x60;schemaVersion&#x60;, &#x60;value&#x60;, &#x60;externalId&#x60;. | [default to &quot;origin,value&quot;]
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of settings objects in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. | 

### Return type

[**EffectiveSettingsValuesList**](EffectiveSettingsValuesList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsObjectByObjectId

> SettingsObject GetSettingsObjectByObjectId(ctx, objectId).Execute()

Gets the specified settings object

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
    objectId := "objectId_example" // string | The ID of the required settings object.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsObjectsAPI.GetSettingsObjectByObjectId(context.Background(), objectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsAPI.GetSettingsObjectByObjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettingsObjectByObjectId`: SettingsObject
    fmt.Fprintf(os.Stdout, "Response from `SettingsObjectsAPI.GetSettingsObjectByObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the required settings object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsObjectByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SettingsObject**](SettingsObject.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSettingsObjects

> ObjectsList GetSettingsObjects(ctx).SchemaIds(schemaIds).Scopes(scopes).ExternalIds(externalIds).Fields(fields).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Lists persisted settings objects



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
    schemaIds := "schemaIds_example" // string | A list of comma-separated schema IDs to which the requested objects belong.   To load the first page, when the **nextPageKey** is not set, either this parameter or **scopes** is required.   To load all objects belonging to the given schema IDs leave the **scopes** parameter empty. (optional)
    scopes := "scopes_example" // string | A list of comma-separated scopes, that the requested objects target.   The selection only matches objects directly targeting the specified scopes. For example, `environment` will not match objects that target a host within environment. For more details, please see [Dynatrace Documentation](https://dt-url.net/ky03459).   To load the first page, when the **nextPageKey** is not set, either this parameter or **schemaIds** is required.   To load all objects belonging to the given scopes leave the **schemaIds** parameter empty. (optional)
    externalIds := "externalIds_example" // string | A list of comma-separated external IDs that the requested objects have.   Each external ID has a maximum length of 500 characters.  Only considered on load of the first page, when the **nextPageKey** is not set. (optional)
    fields := "fields_example" // string | A list of fields to be included to the response. The provided set of fields replaces the default set.    Specify the required top-level fields, separated by commas (for example, `objectId,value`).  Supported fields: `objectId`, `summary`, `searchSummary`, `created`, `modified`, `createdBy`, `modifiedBy`, `author`, `updateToken`, `scope`, `modificationInfo`, `schemaId`, `schemaVersion`, `value`, `externalId`. (optional) (default to "objectId,value")
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of settings objects in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsObjectsAPI.GetSettingsObjects(context.Background()).SchemaIds(schemaIds).Scopes(scopes).ExternalIds(externalIds).Fields(fields).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsAPI.GetSettingsObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSettingsObjects`: ObjectsList
    fmt.Fprintf(os.Stdout, "Response from `SettingsObjectsAPI.GetSettingsObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemaIds** | **string** | A list of comma-separated schema IDs to which the requested objects belong.   To load the first page, when the **nextPageKey** is not set, either this parameter or **scopes** is required.   To load all objects belonging to the given schema IDs leave the **scopes** parameter empty. | 
 **scopes** | **string** | A list of comma-separated scopes, that the requested objects target.   The selection only matches objects directly targeting the specified scopes. For example, &#x60;environment&#x60; will not match objects that target a host within environment. For more details, please see [Dynatrace Documentation](https://dt-url.net/ky03459).   To load the first page, when the **nextPageKey** is not set, either this parameter or **schemaIds** is required.   To load all objects belonging to the given scopes leave the **schemaIds** parameter empty. | 
 **externalIds** | **string** | A list of comma-separated external IDs that the requested objects have.   Each external ID has a maximum length of 500 characters.  Only considered on load of the first page, when the **nextPageKey** is not set. | 
 **fields** | **string** | A list of fields to be included to the response. The provided set of fields replaces the default set.    Specify the required top-level fields, separated by commas (for example, &#x60;objectId,value&#x60;).  Supported fields: &#x60;objectId&#x60;, &#x60;summary&#x60;, &#x60;searchSummary&#x60;, &#x60;created&#x60;, &#x60;modified&#x60;, &#x60;createdBy&#x60;, &#x60;modifiedBy&#x60;, &#x60;author&#x60;, &#x60;updateToken&#x60;, &#x60;scope&#x60;, &#x60;modificationInfo&#x60;, &#x60;schemaId&#x60;, &#x60;schemaVersion&#x60;, &#x60;value&#x60;, &#x60;externalId&#x60;. | [default to &quot;objectId,value&quot;]
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of settings objects in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. | 

### Return type

[**ObjectsList**](ObjectsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSettingsObjects

> []SettingsObjectResponse PostSettingsObjects(ctx).ValidateOnly(validateOnly).SettingsObjectCreate(settingsObjectCreate).Execute()

Creates a new settings object



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
    validateOnly := true // bool | If `true`, the request runs only validation of the submitted settings objects, without saving them. (optional) (default to false)
    settingsObjectCreate := []openapiclient.SettingsObjectCreate{*openapiclient.NewSettingsObjectCreate("builtin:container.built-in-monitoring-rule", map[string]interface{}({"autoMonitoring":true}))} // []SettingsObjectCreate | The JSON body of the request. Contains the settings objects. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsObjectsAPI.PostSettingsObjects(context.Background()).ValidateOnly(validateOnly).SettingsObjectCreate(settingsObjectCreate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsAPI.PostSettingsObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSettingsObjects`: []SettingsObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsObjectsAPI.PostSettingsObjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSettingsObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateOnly** | **bool** | If &#x60;true&#x60;, the request runs only validation of the submitted settings objects, without saving them. | [default to false]
 **settingsObjectCreate** | [**[]SettingsObjectCreate**](SettingsObjectCreate.md) | The JSON body of the request. Contains the settings objects. | 

### Return type

[**[]SettingsObjectResponse**](SettingsObjectResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSettingsObjectByObjectId

> SettingsObjectResponse PutSettingsObjectByObjectId(ctx, objectId).SettingsObjectUpdate(settingsObjectUpdate).Execute()

Updates an existing settings object



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
    objectId := "objectId_example" // string | The ID of the required settings object.
    settingsObjectUpdate := *openapiclient.NewSettingsObjectUpdate(map[string]interface{}({"autoMonitoring":true})) // SettingsObjectUpdate | The JSON body of the request. Contains updated parameters of the settings object. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsObjectsAPI.PutSettingsObjectByObjectId(context.Background(), objectId).SettingsObjectUpdate(settingsObjectUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsObjectsAPI.PutSettingsObjectByObjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSettingsObjectByObjectId`: SettingsObjectResponse
    fmt.Fprintf(os.Stdout, "Response from `SettingsObjectsAPI.PutSettingsObjectByObjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** | The ID of the required settings object. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSettingsObjectByObjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **settingsObjectUpdate** | [**SettingsObjectUpdate**](SettingsObjectUpdate.md) | The JSON body of the request. Contains updated parameters of the settings object. | 

### Return type

[**SettingsObjectResponse**](SettingsObjectResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

