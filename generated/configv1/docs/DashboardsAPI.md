# \DashboardsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDashboard**](DashboardsAPI.md#CreateDashboard) | **Post** /dashboards | Creates a dashboard
[**DeleteDashboard**](DashboardsAPI.md#DeleteDashboard) | **Delete** /dashboards/{id} | Deletes the specified dashboard
[**GetDashboard**](DashboardsAPI.md#GetDashboard) | **Get** /dashboards/{id} | Gets the properties of the specified dashboard
[**GetDashboardSharingSettings**](DashboardsAPI.md#GetDashboardSharingSettings) | **Get** /dashboards/{id}/shareSettings | Gets the sharing configuration of the specified dashboard | maturity&#x3D;EARLY_ADOPTER
[**GetDashboardStubsList**](DashboardsAPI.md#GetDashboardStubsList) | **Get** /dashboards | Lists all dashboards of the environment
[**MigrateAllDashboard**](DashboardsAPI.md#MigrateAllDashboard) | **Get** /dashboards/migrateAll | Migrate all dashboards to replace Custom Charting tiles with DataExplorer ones | maturity&#x3D;EARLY_ADOPTER
[**MigrateDashboard**](DashboardsAPI.md#MigrateDashboard) | **Put** /dashboards/{dashboardId}/migrate | Migrate all the Custom Charting tiles to DataExplorer ones | maturity&#x3D;EARLY_ADOPTER
[**UpdateDashboard**](DashboardsAPI.md#UpdateDashboard) | **Put** /dashboards/{id} | Updates the specified dashboard
[**UpdateShareSettings**](DashboardsAPI.md#UpdateShareSettings) | **Put** /dashboards/{id}/shareSettings | Updates the sharing configuration of the specified dashboard | maturity&#x3D;EARLY_ADOPTER
[**ValidateDashboardCreation**](DashboardsAPI.md#ValidateDashboardCreation) | **Post** /dashboards/validator | Validates the payload for the &#39;POST /dashboards&#39; request
[**ValidateDashboardUpdate**](DashboardsAPI.md#ValidateDashboardUpdate) | **Post** /dashboards/{id}/validator | Validates the payload for the &#39;PUT /dashboards/{id}&#39; request
[**ValidateShareSettingsUpdate**](DashboardsAPI.md#ValidateShareSettingsUpdate) | **Post** /dashboards/{id}/shareSettings/validator | Validates the payload for the &#39;PUT /dashboards/{id}/shareSettings&#39; request | maturity&#x3D;EARLY_ADOPTER



## CreateDashboard

> EntityShortRepresentation CreateDashboard(ctx).Dashboard(dashboard).Execute()

Creates a dashboard



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
    dashboard := *openapiclient.NewDashboard(*openapiclient.NewDashboardMetadata("Name_example", "Owner_example"), []openapiclient.Tile{*openapiclient.NewTile(*openapiclient.NewTileBounds(), "Name_example", "TileType_example")}) // Dashboard | The JSON body of the request. Contains parameters of the new dashboard. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsAPI.CreateDashboard(context.Background()).Dashboard(dashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.CreateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDashboard`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | [**Dashboard**](Dashboard.md) | The JSON body of the request. Contains parameters of the new dashboard. | 

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


## DeleteDashboard

> DeleteDashboard(ctx, id).Execute()

Deletes the specified dashboard

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
    id := "id_example" // string | The ID of the dashboard to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DashboardsAPI.DeleteDashboard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DeleteDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the dashboard to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


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


## GetDashboard

> Dashboard GetDashboard(ctx, id).Execute()

Gets the properties of the specified dashboard

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
    id := "id_example" // string | The ID of the required dashboard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsAPI.GetDashboard(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboard`: Dashboard
    fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardSharingSettings

> DashboardSharing GetDashboardSharingSettings(ctx, id).Execute()

Gets the sharing configuration of the specified dashboard | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required dashboard.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsAPI.GetDashboardSharingSettings(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboardSharingSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardSharingSettings`: DashboardSharing
    fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboardSharingSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardSharingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DashboardSharing**](DashboardSharing.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardStubsList

> DashboardList GetDashboardStubsList(ctx).Owner(owner).Tags(tags).Execute()

Lists all dashboards of the environment

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
    owner := "owner_example" // string | The owner of the dashboard. (optional)
    tags := []string{"Inner_example"} // []string | A list of tags applied to the dashboard.    The dashboard must match **all** the specified tags. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsAPI.GetDashboardStubsList(context.Background()).Owner(owner).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboardStubsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDashboardStubsList`: DashboardList
    fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboardStubsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardStubsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **owner** | **string** | The owner of the dashboard. | 
 **tags** | **[]string** | A list of tags applied to the dashboard.    The dashboard must match **all** the specified tags. | 

### Return type

[**DashboardList**](DashboardList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateAllDashboard

> MigrateAllDashboard(ctx).Execute()

Migrate all dashboards to replace Custom Charting tiles with DataExplorer ones | maturity=EARLY_ADOPTER

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
    r, err := apiClient.DashboardsAPI.MigrateAllDashboard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.MigrateAllDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateAllDashboardRequest struct via the builder pattern


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


## MigrateDashboard

> MigrateDashboard(ctx, dashboardId).Execute()

Migrate all the Custom Charting tiles to DataExplorer ones | maturity=EARLY_ADOPTER

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
    dashboardId := "dashboardId_example" // string | The ID of the dashboard to be migrated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DashboardsAPI.MigrateDashboard(context.Background(), dashboardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.MigrateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard to be migrated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrateDashboardRequest struct via the builder pattern


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


## UpdateDashboard

> EntityShortRepresentation UpdateDashboard(ctx, id).Dashboard(dashboard).Execute()

Updates the specified dashboard



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
    id := "id_example" // string | The ID of the dashboard to be updated.    The ID in the request body must match this ID.
    dashboard := *openapiclient.NewDashboard(*openapiclient.NewDashboardMetadata("Name_example", "Owner_example"), []openapiclient.Tile{*openapiclient.NewTile(*openapiclient.NewTileBounds(), "Name_example", "TileType_example")}) // Dashboard | The JSON body of the request. Contains updated parameters of the dashboard. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardsAPI.UpdateDashboard(context.Background(), id).Dashboard(dashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.UpdateDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDashboard`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.UpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the dashboard to be updated.    The ID in the request body must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**Dashboard**](Dashboard.md) | The JSON body of the request. Contains updated parameters of the dashboard. | 

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


## UpdateShareSettings

> UpdateShareSettings(ctx, id).DashboardSharing(dashboardSharing).Execute()

Updates the sharing configuration of the specified dashboard | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the required dashboard.
    dashboardSharing := *openapiclient.NewDashboardSharing("Id_example", []openapiclient.DashboardSharePermissions{*openapiclient.NewDashboardSharePermissions("Permission_example", "Type_example")}, *openapiclient.NewDashboardAnonymousAccess([]string{"ManagementZoneIds_example"})) // DashboardSharing | The JSON body of the request. Contains updated parameters of the dashboard sharing. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DashboardsAPI.UpdateShareSettings(context.Background(), id).DashboardSharing(dashboardSharing).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.UpdateShareSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShareSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardSharing** | [**DashboardSharing**](DashboardSharing.md) | The JSON body of the request. Contains updated parameters of the dashboard sharing. | 

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


## ValidateDashboardCreation

> ValidateDashboardCreation(ctx).Dashboard(dashboard).Execute()

Validates the payload for the 'POST /dashboards' request



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
    dashboard := *openapiclient.NewDashboard(*openapiclient.NewDashboardMetadata("Name_example", "Owner_example"), []openapiclient.Tile{*openapiclient.NewTile(*openapiclient.NewTileBounds(), "Name_example", "TileType_example")}) // Dashboard | The JSON body of the request. Containing the dashboard to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DashboardsAPI.ValidateDashboardCreation(context.Background()).Dashboard(dashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.ValidateDashboardCreation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateDashboardCreationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboard** | [**Dashboard**](Dashboard.md) | The JSON body of the request. Containing the dashboard to be validated. | 

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


## ValidateDashboardUpdate

> ValidateDashboardUpdate(ctx, id).Dashboard(dashboard).Execute()

Validates the payload for the 'PUT /dashboards/{id}' request

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
    id := "id_example" // string | The ID of the dashboard to be validated.    The ID in the request body must match this ID.
    dashboard := *openapiclient.NewDashboard(*openapiclient.NewDashboardMetadata("Name_example", "Owner_example"), []openapiclient.Tile{*openapiclient.NewTile(*openapiclient.NewTileBounds(), "Name_example", "TileType_example")}) // Dashboard | The JSON body of the request. Contains the dashboard to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DashboardsAPI.ValidateDashboardUpdate(context.Background(), id).Dashboard(dashboard).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.ValidateDashboardUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the dashboard to be validated.    The ID in the request body must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateDashboardUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboard** | [**Dashboard**](Dashboard.md) | The JSON body of the request. Contains the dashboard to be validated. | 

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


## ValidateShareSettingsUpdate

> ValidateShareSettingsUpdate(ctx, id).DashboardSharing(dashboardSharing).Execute()

Validates the payload for the 'PUT /dashboards/{id}/shareSettings' request | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the dashboard to be validated.    The ID in the request body must match this ID.
    dashboardSharing := *openapiclient.NewDashboardSharing("Id_example", []openapiclient.DashboardSharePermissions{*openapiclient.NewDashboardSharePermissions("Permission_example", "Type_example")}, *openapiclient.NewDashboardAnonymousAccess([]string{"ManagementZoneIds_example"})) // DashboardSharing | The JSON body of the request. Contains parameters of the dashboard sharing to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DashboardsAPI.ValidateShareSettingsUpdate(context.Background(), id).DashboardSharing(dashboardSharing).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.ValidateShareSettingsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the dashboard to be validated.    The ID in the request body must match this ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateShareSettingsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardSharing** | [**DashboardSharing**](DashboardSharing.md) | The JSON body of the request. Contains parameters of the dashboard sharing to be validated. | 

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

