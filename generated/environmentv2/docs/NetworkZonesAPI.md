# \NetworkZonesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateNetworkZone**](NetworkZonesAPI.md#CreateOrUpdateNetworkZone) | **Put** /networkZones/{id} | Updates an existing network zone or creates a new one
[**DeleteNetworkZone**](NetworkZonesAPI.md#DeleteNetworkZone) | **Delete** /networkZones/{id} | Deletes the specified network zone
[**GetAllNetworkZones**](NetworkZonesAPI.md#GetAllNetworkZones) | **Get** /networkZones | Lists all existing network zones
[**GetHostStats**](NetworkZonesAPI.md#GetHostStats) | **Get** /networkZones/{id}/hostConnectionStatistics | Gets the statistics about hosts using the network zone
[**GetNetworkZoneSettings**](NetworkZonesAPI.md#GetNetworkZoneSettings) | **Get** /networkZoneSettings | Gets the global configuration of network zones
[**GetSingleNetworkZone**](NetworkZonesAPI.md#GetSingleNetworkZone) | **Get** /networkZones/{id} | Gets parameters of the specified network zone
[**UpdateNetworkZoneSettings**](NetworkZonesAPI.md#UpdateNetworkZoneSettings) | **Put** /networkZoneSettings | Updates the global configuration of network zones



## CreateOrUpdateNetworkZone

> EntityShortRepresentation CreateOrUpdateNetworkZone(ctx, id).NetworkZone(networkZone).Execute()

Updates an existing network zone or creates a new one



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
    id := "id_example" // string | The ID of the network zone to be updated.    If you set the ID in the body as well, it must match this ID.    The ID is not case sensitive. Dynatrace stores the ID in lowercase.
    networkZone := *openapiclient.NewNetworkZone() // NetworkZone | The JSON body of the request. Contains parameters of the network zone.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZonesAPI.CreateOrUpdateNetworkZone(context.Background(), id).NetworkZone(networkZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesAPI.CreateOrUpdateNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateNetworkZone`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `NetworkZonesAPI.CreateOrUpdateNetworkZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the network zone to be updated.    If you set the ID in the body as well, it must match this ID.    The ID is not case sensitive. Dynatrace stores the ID in lowercase. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateNetworkZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkZone** | [**NetworkZone**](NetworkZone.md) | The JSON body of the request. Contains parameters of the network zone. | 

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


## DeleteNetworkZone

> DeleteNetworkZone(ctx, id).Execute()

Deletes the specified network zone



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
    id := "id_example" // string | The ID of the network zone to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkZonesAPI.DeleteNetworkZone(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesAPI.DeleteNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the network zone to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkZoneRequest struct via the builder pattern


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


## GetAllNetworkZones

> NetworkZoneList GetAllNetworkZones(ctx).Execute()

Lists all existing network zones

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
    resp, r, err := apiClient.NetworkZonesAPI.GetAllNetworkZones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesAPI.GetAllNetworkZones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllNetworkZones`: NetworkZoneList
    fmt.Fprintf(os.Stdout, "Response from `NetworkZonesAPI.GetAllNetworkZones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNetworkZonesRequest struct via the builder pattern


### Return type

[**NetworkZoneList**](NetworkZoneList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHostStats

> NetworkZoneConnectionStatistics GetHostStats(ctx, id).Filter(filter).Execute()

Gets the statistics about hosts using the network zone

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
    id := "id_example" // string | The ID of the required network zone.
    filter := "filter_example" // string | Filters the resulting set of hosts:   * `all`: All hosts using the zone.  * `configuredButNotConnectedOnly`: Hosts from the network zone that use other zones.  * `connectedAsAlternativeOnly`: Hosts that use the network zone as an alternative.  * `connectedAsFailoverOnly`: Hosts from other zones that use the zone (not configured as an alternative) even though ActiveGates of higher priority are available.  * `connectedAsFailoverWithoutOwnActiveGatesOnly`: Hosts from other zones that use the zone (not configured as an alternative) and **no** ActiveGates of higher priority are available.   If not set, `all` is used. (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZonesAPI.GetHostStats(context.Background(), id).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesAPI.GetHostStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostStats`: NetworkZoneConnectionStatistics
    fmt.Fprintf(os.Stdout, "Response from `NetworkZonesAPI.GetHostStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required network zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHostStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filters the resulting set of hosts:   * &#x60;all&#x60;: All hosts using the zone.  * &#x60;configuredButNotConnectedOnly&#x60;: Hosts from the network zone that use other zones.  * &#x60;connectedAsAlternativeOnly&#x60;: Hosts that use the network zone as an alternative.  * &#x60;connectedAsFailoverOnly&#x60;: Hosts from other zones that use the zone (not configured as an alternative) even though ActiveGates of higher priority are available.  * &#x60;connectedAsFailoverWithoutOwnActiveGatesOnly&#x60;: Hosts from other zones that use the zone (not configured as an alternative) and **no** ActiveGates of higher priority are available.   If not set, &#x60;all&#x60; is used. | [default to &quot;all&quot;]

### Return type

[**NetworkZoneConnectionStatistics**](NetworkZoneConnectionStatistics.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkZoneSettings

> NetworkZoneSettings GetNetworkZoneSettings(ctx).Execute()

Gets the global configuration of network zones

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
    resp, r, err := apiClient.NetworkZonesAPI.GetNetworkZoneSettings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesAPI.GetNetworkZoneSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkZoneSettings`: NetworkZoneSettings
    fmt.Fprintf(os.Stdout, "Response from `NetworkZonesAPI.GetNetworkZoneSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkZoneSettingsRequest struct via the builder pattern


### Return type

[**NetworkZoneSettings**](NetworkZoneSettings.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSingleNetworkZone

> NetworkZone GetSingleNetworkZone(ctx, id).Execute()

Gets parameters of the specified network zone

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
    id := "id_example" // string | The ID of the required network zone.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkZonesAPI.GetSingleNetworkZone(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesAPI.GetSingleNetworkZone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSingleNetworkZone`: NetworkZone
    fmt.Fprintf(os.Stdout, "Response from `NetworkZonesAPI.GetSingleNetworkZone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required network zone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSingleNetworkZoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NetworkZone**](NetworkZone.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkZoneSettings

> UpdateNetworkZoneSettings(ctx).NetworkZoneSettings(networkZoneSettings).Execute()

Updates the global configuration of network zones

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
    networkZoneSettings := *openapiclient.NewNetworkZoneSettings() // NetworkZoneSettings | The JSON body of the request. Contains global configuration of network zones.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.NetworkZonesAPI.UpdateNetworkZoneSettings(context.Background()).NetworkZoneSettings(networkZoneSettings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkZonesAPI.UpdateNetworkZoneSettings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkZoneSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkZoneSettings** | [**NetworkZoneSettings**](NetworkZoneSettings.md) | The JSON body of the request. Contains global configuration of network zones. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

