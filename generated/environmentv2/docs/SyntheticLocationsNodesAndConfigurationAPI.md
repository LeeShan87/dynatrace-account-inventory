# \SyntheticLocationsNodesAndConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLocation**](SyntheticLocationsNodesAndConfigurationAPI.md#AddLocation) | **Post** /synthetic/locations | Creates a new private synthetic location
[**GetConfiguration**](SyntheticLocationsNodesAndConfigurationAPI.md#GetConfiguration) | **Get** /synthetic/config | Gets set of synthetic related parameters defined for whole tenant (affects all monitors and all private locations). | maturity&#x3D;EARLY_ADOPTER
[**GetLocation**](SyntheticLocationsNodesAndConfigurationAPI.md#GetLocation) | **Get** /synthetic/locations/{locationId} | Gets properties of the specified location
[**GetLocationDemploymentApplyCommands**](SyntheticLocationsNodesAndConfigurationAPI.md#GetLocationDemploymentApplyCommands) | **Get** /synthetic/locations/commands/apply | Gets list of commands to deploy synthetic location in Kubernetes/Openshift cluster | maturity&#x3D;EARLY_ADOPTER
[**GetLocationDemploymentDeleteCommands**](SyntheticLocationsNodesAndConfigurationAPI.md#GetLocationDemploymentDeleteCommands) | **Get** /synthetic/locations/{locationId}/commands/delete | Gets list of commands to delete synthetic location in Kubernetes/Openshift cluster | maturity&#x3D;EARLY_ADOPTER
[**GetLocationDemploymentYaml**](SyntheticLocationsNodesAndConfigurationAPI.md#GetLocationDemploymentYaml) | **Get** /synthetic/locations/{locationId}/yaml | Gets yaml file content to deploy location in Kubernetes/Openshift cluster | maturity&#x3D;EARLY_ADOPTER
[**GetLocations**](SyntheticLocationsNodesAndConfigurationAPI.md#GetLocations) | **Get** /synthetic/locations | Lists all synthetic locations (both public and private) available for your environment
[**GetLocationsStatus**](SyntheticLocationsNodesAndConfigurationAPI.md#GetLocationsStatus) | **Get** /synthetic/locations/status | Checks the status of public synthetic locations
[**GetNode**](SyntheticLocationsNodesAndConfigurationAPI.md#GetNode) | **Get** /synthetic/nodes/{nodeId} | Lists properties of the specified synthetic node
[**GetNodes**](SyntheticLocationsNodesAndConfigurationAPI.md#GetNodes) | **Get** /synthetic/nodes | Lists all synthetic nodes available in your environment
[**RemoveLocation**](SyntheticLocationsNodesAndConfigurationAPI.md#RemoveLocation) | **Delete** /synthetic/locations/{locationId} | Deletes the specified private synthetic location
[**UpdateConfiguration**](SyntheticLocationsNodesAndConfigurationAPI.md#UpdateConfiguration) | **Put** /synthetic/config | Updates set of synthetic related parameters defined for whole tenant (affects all monitors and all private locations). | maturity&#x3D;EARLY_ADOPTER
[**UpdateLocation**](SyntheticLocationsNodesAndConfigurationAPI.md#UpdateLocation) | **Put** /synthetic/locations/{locationId} | Updates the specified synthetic location
[**UpdateLocationsStatus**](SyntheticLocationsNodesAndConfigurationAPI.md#UpdateLocationsStatus) | **Put** /synthetic/locations/status | Changes the status of public synthetic locations



## AddLocation

> SyntheticLocationIdsDto AddLocation(ctx).PrivateSyntheticLocation(privateSyntheticLocation).Execute()

Creates a new private synthetic location

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
    privateSyntheticLocation := *openapiclient.NewPrivateSyntheticLocation([]string{"Nodes_example"}, "EntityId_example", float64(123), float64(123), "Name_example", "Type_example") // PrivateSyntheticLocation | The JSON body of the request. Contains parameters of the new private synthetic location. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.AddLocation(context.Background()).PrivateSyntheticLocation(privateSyntheticLocation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.AddLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLocation`: SyntheticLocationIdsDto
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsNodesAndConfigurationAPI.AddLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **privateSyntheticLocation** | [**PrivateSyntheticLocation**](PrivateSyntheticLocation.md) | The JSON body of the request. Contains parameters of the new private synthetic location. | 

### Return type

[**SyntheticLocationIdsDto**](SyntheticLocationIdsDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> SyntheticConfigDto GetConfiguration(ctx).Execute()

Gets set of synthetic related parameters defined for whole tenant (affects all monitors and all private locations). | maturity=EARLY_ADOPTER

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
    resp, r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.GetConfiguration(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.GetConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConfiguration`: SyntheticConfigDto
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsNodesAndConfigurationAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


### Return type

[**SyntheticConfigDto**](SyntheticConfigDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocation

> SyntheticLocation GetLocation(ctx, locationId).Execute()

Gets properties of the specified location

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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the required location.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.GetLocation(context.Background(), locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.GetLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocation`: SyntheticLocation
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsNodesAndConfigurationAPI.GetLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The Dynatrace entity ID of the required location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SyntheticLocation**](SyntheticLocation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationDemploymentApplyCommands

> GetLocationDemploymentApplyCommands(ctx).Platform(platform).Namespace(namespace).Filename(filename).Execute()

Gets list of commands to deploy synthetic location in Kubernetes/Openshift cluster | maturity=EARLY_ADOPTER

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
    platform := "platform_example" // string | Container platform, currently supported are: KUBERNETES and OPENSHIFT. Default value is KUBERNETES. (optional)
    namespace := "namespace_example" // string | Namespace (optional)
    filename := "filename_example" // string | Filename (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.GetLocationDemploymentApplyCommands(context.Background()).Platform(platform).Namespace(namespace).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.GetLocationDemploymentApplyCommands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationDemploymentApplyCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** | Container platform, currently supported are: KUBERNETES and OPENSHIFT. Default value is KUBERNETES. | 
 **namespace** | **string** | Namespace | 
 **filename** | **string** | Filename | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationDemploymentDeleteCommands

> GetLocationDemploymentDeleteCommands(ctx, locationId).Platform(platform).Namespace(namespace).Filename(filename).Execute()

Gets list of commands to delete synthetic location in Kubernetes/Openshift cluster | maturity=EARLY_ADOPTER

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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the required location.
    platform := "platform_example" // string | Container platform, currently supported are: KUBERNETES and OPENSHIFT. Default value is KUBERNETES. (optional)
    namespace := "namespace_example" // string | Namespace (optional)
    filename := "filename_example" // string | Filename (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.GetLocationDemploymentDeleteCommands(context.Background(), locationId).Platform(platform).Namespace(namespace).Filename(filename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.GetLocationDemploymentDeleteCommands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The Dynatrace entity ID of the required location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationDemploymentDeleteCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platform** | **string** | Container platform, currently supported are: KUBERNETES and OPENSHIFT. Default value is KUBERNETES. | 
 **namespace** | **string** | Namespace | 
 **filename** | **string** | Filename | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationDemploymentYaml

> GetLocationDemploymentYaml(ctx, locationId).Platform(platform).ActiveGateName(activeGateName).Namespace(namespace).Execute()

Gets yaml file content to deploy location in Kubernetes/Openshift cluster | maturity=EARLY_ADOPTER

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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the required location.
    platform := "platform_example" // string | Container platform, currently supported are: KUBERNETES and OPENSHIFT. Default value is KUBERNETES. (optional)
    activeGateName := "activeGateName_example" // string | Active gate name (optional)
    namespace := "namespace_example" // string | Namespace (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.GetLocationDemploymentYaml(context.Background(), locationId).Platform(platform).ActiveGateName(activeGateName).Namespace(namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.GetLocationDemploymentYaml``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The Dynatrace entity ID of the required location. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationDemploymentYamlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platform** | **string** | Container platform, currently supported are: KUBERNETES and OPENSHIFT. Default value is KUBERNETES. | 
 **activeGateName** | **string** | Active gate name | 
 **namespace** | **string** | Namespace | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/yaml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocations

> SyntheticLocations GetLocations(ctx).CloudPlatform(cloudPlatform).Type_(type_).Execute()

Lists all synthetic locations (both public and private) available for your environment

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
    cloudPlatform := "cloudPlatform_example" // string | Filters the resulting set of locations to those which are hosted on a specific cloud platform. (optional)
    type_ := "type__example" // string | Filters the resulting set of locations to those of a specific type. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.GetLocations(context.Background()).CloudPlatform(cloudPlatform).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.GetLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocations`: SyntheticLocations
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsNodesAndConfigurationAPI.GetLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudPlatform** | **string** | Filters the resulting set of locations to those which are hosted on a specific cloud platform. | 
 **type_** | **string** | Filters the resulting set of locations to those of a specific type. | 

### Return type

[**SyntheticLocations**](SyntheticLocations.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLocationsStatus

> SyntheticPublicLocationsStatus GetLocationsStatus(ctx).Execute()

Checks the status of public synthetic locations

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
    resp, r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.GetLocationsStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.GetLocationsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLocationsStatus`: SyntheticPublicLocationsStatus
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsNodesAndConfigurationAPI.GetLocationsStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLocationsStatusRequest struct via the builder pattern


### Return type

[**SyntheticPublicLocationsStatus**](SyntheticPublicLocationsStatus.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNode

> Node GetNode(ctx, nodeId).Execute()

Lists properties of the specified synthetic node

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
    nodeId := "nodeId_example" // string | The ID of the required synthetic node.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.GetNode(context.Background(), nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.GetNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNode`: Node
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsNodesAndConfigurationAPI.GetNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nodeId** | **string** | The ID of the required synthetic node. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Node**](Node.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodes

> Nodes GetNodes(ctx).Execute()

Lists all synthetic nodes available in your environment

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
    resp, r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.GetNodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.GetNodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodes`: Nodes
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsNodesAndConfigurationAPI.GetNodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodesRequest struct via the builder pattern


### Return type

[**Nodes**](Nodes.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLocation

> RemoveLocation(ctx, locationId).Execute()

Deletes the specified private synthetic location

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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the private synthetic location to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.RemoveLocation(context.Background(), locationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.RemoveLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The Dynatrace entity ID of the private synthetic location to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLocationRequest struct via the builder pattern


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


## UpdateConfiguration

> SyntheticConfigDto UpdateConfiguration(ctx).SyntheticConfigDto(syntheticConfigDto).Execute()

Updates set of synthetic related parameters defined for whole tenant (affects all monitors and all private locations). | maturity=EARLY_ADOPTER

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
    syntheticConfigDto := *openapiclient.NewSyntheticConfigDto(int64(123), int64(123)) // SyntheticConfigDto |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.UpdateConfiguration(context.Background()).SyntheticConfigDto(syntheticConfigDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.UpdateConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConfiguration`: SyntheticConfigDto
    fmt.Fprintf(os.Stdout, "Response from `SyntheticLocationsNodesAndConfigurationAPI.UpdateConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syntheticConfigDto** | [**SyntheticConfigDto**](SyntheticConfigDto.md) |  | 

### Return type

[**SyntheticConfigDto**](SyntheticConfigDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLocation

> UpdateLocation(ctx, locationId).SyntheticLocationUpdate(syntheticLocationUpdate).Execute()

Updates the specified synthetic location



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
    locationId := "locationId_example" // string | The Dynatrace entity ID of the synthetic location to be updated.
    syntheticLocationUpdate := *openapiclient.NewSyntheticLocationUpdate("Type_example") // SyntheticLocationUpdate | The JSON body of the request. Contains updated parameters of the location. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.UpdateLocation(context.Background(), locationId).SyntheticLocationUpdate(syntheticLocationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.UpdateLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**locationId** | **string** | The Dynatrace entity ID of the synthetic location to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **syntheticLocationUpdate** | [**SyntheticLocationUpdate**](SyntheticLocationUpdate.md) | The JSON body of the request. Contains updated parameters of the location. | 

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


## UpdateLocationsStatus

> UpdateLocationsStatus(ctx).SyntheticPublicLocationsStatus(syntheticPublicLocationsStatus).Execute()

Changes the status of public synthetic locations

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
    syntheticPublicLocationsStatus := *openapiclient.NewSyntheticPublicLocationsStatus(false) // SyntheticPublicLocationsStatus | The new status of public synthetic locations. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SyntheticLocationsNodesAndConfigurationAPI.UpdateLocationsStatus(context.Background()).SyntheticPublicLocationsStatus(syntheticPublicLocationsStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SyntheticLocationsNodesAndConfigurationAPI.UpdateLocationsStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLocationsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syntheticPublicLocationsStatus** | [**SyntheticPublicLocationsStatus**](SyntheticPublicLocationsStatus.md) | The new status of public synthetic locations. | 

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

