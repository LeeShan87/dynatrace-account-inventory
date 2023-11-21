# \RUMGeographicRegionsCustomClientIPHeadersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGeolocationRegionsIpHeaders**](RUMGeographicRegionsCustomClientIPHeadersAPI.md#GetGeolocationRegionsIpHeaders) | **Get** /geographicRegions/ipDetectionHeaders | Gets the configuration of custom client IP headers
[**PutGeolocationRegionsIpHeaders**](RUMGeographicRegionsCustomClientIPHeadersAPI.md#PutGeolocationRegionsIpHeaders) | **Put** /geographicRegions/ipDetectionHeaders | Updates the configuration of custom client IP headers
[**ValidateGeolocationRegionsIpHeaders**](RUMGeographicRegionsCustomClientIPHeadersAPI.md#ValidateGeolocationRegionsIpHeaders) | **Post** /geographicRegions/ipDetectionHeaders/validator | Validates the payload for the &#x60;PUT /geographicRegions/ipDetectionHeaders&#x60; request



## GetGeolocationRegionsIpHeaders

> IpDetectionHeaders GetGeolocationRegionsIpHeaders(ctx).Execute()

Gets the configuration of custom client IP headers

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
    resp, r, err := apiClient.RUMGeographicRegionsCustomClientIPHeadersAPI.GetGeolocationRegionsIpHeaders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMGeographicRegionsCustomClientIPHeadersAPI.GetGeolocationRegionsIpHeaders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGeolocationRegionsIpHeaders`: IpDetectionHeaders
    fmt.Fprintf(os.Stdout, "Response from `RUMGeographicRegionsCustomClientIPHeadersAPI.GetGeolocationRegionsIpHeaders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGeolocationRegionsIpHeadersRequest struct via the builder pattern


### Return type

[**IpDetectionHeaders**](IpDetectionHeaders.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutGeolocationRegionsIpHeaders

> PutGeolocationRegionsIpHeaders(ctx).IpDetectionHeaders(ipDetectionHeaders).Execute()

Updates the configuration of custom client IP headers



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
    ipDetectionHeaders := *openapiclient.NewIpDetectionHeaders() // IpDetectionHeaders | The JSON body of the request. Contains the configuration of the custom client IP headers. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMGeographicRegionsCustomClientIPHeadersAPI.PutGeolocationRegionsIpHeaders(context.Background()).IpDetectionHeaders(ipDetectionHeaders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMGeographicRegionsCustomClientIPHeadersAPI.PutGeolocationRegionsIpHeaders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutGeolocationRegionsIpHeadersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipDetectionHeaders** | [**IpDetectionHeaders**](IpDetectionHeaders.md) | The JSON body of the request. Contains the configuration of the custom client IP headers. | 

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


## ValidateGeolocationRegionsIpHeaders

> ValidateGeolocationRegionsIpHeaders(ctx).IpDetectionHeaders(ipDetectionHeaders).Execute()

Validates the payload for the `PUT /geographicRegions/ipDetectionHeaders` request

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
    ipDetectionHeaders := *openapiclient.NewIpDetectionHeaders() // IpDetectionHeaders | The JSON body of the request. Contains the configuration of the custom client IP headers to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMGeographicRegionsCustomClientIPHeadersAPI.ValidateGeolocationRegionsIpHeaders(context.Background()).IpDetectionHeaders(ipDetectionHeaders).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMGeographicRegionsCustomClientIPHeadersAPI.ValidateGeolocationRegionsIpHeaders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateGeolocationRegionsIpHeadersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipDetectionHeaders** | [**IpDetectionHeaders**](IpDetectionHeaders.md) | The JSON body of the request. Contains the configuration of the custom client IP headers to be validated. | 

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

