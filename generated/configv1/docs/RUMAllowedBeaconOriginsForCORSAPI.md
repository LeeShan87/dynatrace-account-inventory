# \RUMAllowedBeaconOriginsForCORSAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllowedBeaconOriginsConfig**](RUMAllowedBeaconOriginsForCORSAPI.md#GetAllowedBeaconOriginsConfig) | **Get** /allowedBeaconOriginsForCors | Gets the configuration of the allowed beacon origins for CORS requests
[**PutAllowedBeaconOriginsConfig**](RUMAllowedBeaconOriginsForCORSAPI.md#PutAllowedBeaconOriginsConfig) | **Put** /allowedBeaconOriginsForCors | Updates the configuration of the allowed beacon origins for CORS requests
[**ValidateAllowedBeaconOriginsConfig**](RUMAllowedBeaconOriginsForCORSAPI.md#ValidateAllowedBeaconOriginsConfig) | **Post** /allowedBeaconOriginsForCors/validator | Validates the payload for the &#x60;PUT /allowedBeaconOriginsForCors&#x60; request



## GetAllowedBeaconOriginsConfig

> AllowedBeaconOrigins GetAllowedBeaconOriginsConfig(ctx).Execute()

Gets the configuration of the allowed beacon origins for CORS requests

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
    resp, r, err := apiClient.RUMAllowedBeaconOriginsForCORSAPI.GetAllowedBeaconOriginsConfig(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMAllowedBeaconOriginsForCORSAPI.GetAllowedBeaconOriginsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllowedBeaconOriginsConfig`: AllowedBeaconOrigins
    fmt.Fprintf(os.Stdout, "Response from `RUMAllowedBeaconOriginsForCORSAPI.GetAllowedBeaconOriginsConfig`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowedBeaconOriginsConfigRequest struct via the builder pattern


### Return type

[**AllowedBeaconOrigins**](AllowedBeaconOrigins.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutAllowedBeaconOriginsConfig

> PutAllowedBeaconOriginsConfig(ctx).AllowedBeaconOrigins(allowedBeaconOrigins).Execute()

Updates the configuration of the allowed beacon origins for CORS requests



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
    allowedBeaconOrigins := *openapiclient.NewAllowedBeaconOrigins() // AllowedBeaconOrigins | The JSON body of the request. Contains the configuration of the allowed beacon origins for CORS requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMAllowedBeaconOriginsForCORSAPI.PutAllowedBeaconOriginsConfig(context.Background()).AllowedBeaconOrigins(allowedBeaconOrigins).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMAllowedBeaconOriginsForCORSAPI.PutAllowedBeaconOriginsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutAllowedBeaconOriginsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowedBeaconOrigins** | [**AllowedBeaconOrigins**](AllowedBeaconOrigins.md) | The JSON body of the request. Contains the configuration of the allowed beacon origins for CORS requests. | 

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


## ValidateAllowedBeaconOriginsConfig

> ValidateAllowedBeaconOriginsConfig(ctx).AllowedBeaconOrigins(allowedBeaconOrigins).Execute()

Validates the payload for the `PUT /allowedBeaconOriginsForCors` request

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
    allowedBeaconOrigins := *openapiclient.NewAllowedBeaconOrigins() // AllowedBeaconOrigins | The JSON body of the request. Contains the configuration of the allowed beacon origins for CORS requests. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RUMAllowedBeaconOriginsForCORSAPI.ValidateAllowedBeaconOriginsConfig(context.Background()).AllowedBeaconOrigins(allowedBeaconOrigins).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RUMAllowedBeaconOriginsForCORSAPI.ValidateAllowedBeaconOriginsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAllowedBeaconOriginsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowedBeaconOrigins** | [**AllowedBeaconOrigins**](AllowedBeaconOrigins.md) | The JSON body of the request. Contains the configuration of the allowed beacon origins for CORS requests. | 

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

