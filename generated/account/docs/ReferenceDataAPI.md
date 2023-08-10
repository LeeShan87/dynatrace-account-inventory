# \ReferenceDataAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReferenceDataControllerGetPermissions**](ReferenceDataAPI.md#ReferenceDataControllerGetPermissions) | **Get** /ref/v1/account/permissions | Lists all available permissions
[**ReferenceDataControllerGetRegions**](ReferenceDataAPI.md#ReferenceDataControllerGetRegions) | **Get** /ref/v1/regions | Lists all available regions
[**ReferenceDataControllerGetTimezones**](ReferenceDataAPI.md#ReferenceDataControllerGetTimezones) | **Get** /ref/v1/time-zones | Lists all available time zones



## ReferenceDataControllerGetPermissions

> []PermissionDto ReferenceDataControllerGetPermissions(ctx).Execute()

Lists all available permissions

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
    resp, r, err := apiClient.ReferenceDataAPI.ReferenceDataControllerGetPermissions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.ReferenceDataControllerGetPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferenceDataControllerGetPermissions`: []PermissionDto
    fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.ReferenceDataControllerGetPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReferenceDataControllerGetPermissionsRequest struct via the builder pattern


### Return type

[**[]PermissionDto**](PermissionDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceDataControllerGetRegions

> []RegionDto ReferenceDataControllerGetRegions(ctx).Execute()

Lists all available regions

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
    resp, r, err := apiClient.ReferenceDataAPI.ReferenceDataControllerGetRegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.ReferenceDataControllerGetRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferenceDataControllerGetRegions`: []RegionDto
    fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.ReferenceDataControllerGetRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReferenceDataControllerGetRegionsRequest struct via the builder pattern


### Return type

[**[]RegionDto**](RegionDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceDataControllerGetTimezones

> []TimeZoneDto ReferenceDataControllerGetTimezones(ctx).Execute()

Lists all available time zones

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
    resp, r, err := apiClient.ReferenceDataAPI.ReferenceDataControllerGetTimezones(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceDataAPI.ReferenceDataControllerGetTimezones``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReferenceDataControllerGetTimezones`: []TimeZoneDto
    fmt.Fprintf(os.Stdout, "Response from `ReferenceDataAPI.ReferenceDataControllerGetTimezones`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReferenceDataControllerGetTimezonesRequest struct via the builder pattern


### Return type

[**[]TimeZoneDto**](TimeZoneDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

