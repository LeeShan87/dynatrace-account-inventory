# \RUMGeographicRegionsIPAddressMappingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGeolocationRegionsIpAddress**](RUMGeographicRegionsIPAddressMappingAPI.md#GetGeolocationRegionsIpAddress) | **Get** /geographicRegions/ipAddressMappings | Gets the configuration of mapping between IP address and geographic regions
[**PutGeolocationRegionsIpAddress**](RUMGeographicRegionsIPAddressMappingAPI.md#PutGeolocationRegionsIpAddress) | **Put** /geographicRegions/ipAddressMappings | Updates the configuration of mapping between IP address and geographic regions
[**ValidateGeolocationRegionsIpAddress**](RUMGeographicRegionsIPAddressMappingAPI.md#ValidateGeolocationRegionsIpAddress) | **Post** /geographicRegions/ipAddressMappings/validator | Validates the payload for the &#x60;PUT /geographicRegions/ipAddressMappings&#x60; request



## GetGeolocationRegionsIpAddress

> IpAddressMappings GetGeolocationRegionsIpAddress(ctx).Execute()

Gets the configuration of mapping between IP address and geographic regions

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
    resp, r, err := apiClient.RUMGeographicRegionsIPAddressMappingAPI.GetGeolocationRegionsIpAddress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMGeographicRegionsIPAddressMappingAPI.GetGeolocationRegionsIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGeolocationRegionsIpAddress`: IpAddressMappings
    fmt.Fprintf(os.Stdout, "Response from `RUMGeographicRegionsIPAddressMappingAPI.GetGeolocationRegionsIpAddress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGeolocationRegionsIpAddressRequest struct via the builder pattern


### Return type

[**IpAddressMappings**](IpAddressMappings.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutGeolocationRegionsIpAddress

> PutGeolocationRegionsIpAddress(ctx).IpAddressMappings(ipAddressMappings).Execute()

Updates the configuration of mapping between IP address and geographic regions



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
    ipAddressMappings := *openapiclient.NewIpAddressMappings() // IpAddressMappings | The JSON body of the request. Contains the configuration of the IP address mapping. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMGeographicRegionsIPAddressMappingAPI.PutGeolocationRegionsIpAddress(context.Background()).IpAddressMappings(ipAddressMappings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMGeographicRegionsIPAddressMappingAPI.PutGeolocationRegionsIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutGeolocationRegionsIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipAddressMappings** | [**IpAddressMappings**](IpAddressMappings.md) | The JSON body of the request. Contains the configuration of the IP address mapping. | 

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


## ValidateGeolocationRegionsIpAddress

> ValidateGeolocationRegionsIpAddress(ctx).IpAddressMappings(ipAddressMappings).Execute()

Validates the payload for the `PUT /geographicRegions/ipAddressMappings` request

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
    ipAddressMappings := *openapiclient.NewIpAddressMappings() // IpAddressMappings | The JSON body of the request. Contains the configuration of the IP address. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMGeographicRegionsIPAddressMappingAPI.ValidateGeolocationRegionsIpAddress(context.Background()).IpAddressMappings(ipAddressMappings).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMGeographicRegionsIPAddressMappingAPI.ValidateGeolocationRegionsIpAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateGeolocationRegionsIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipAddressMappings** | [**IpAddressMappings**](IpAddressMappings.md) | The JSON body of the request. Contains the configuration of the IP address. | 

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

