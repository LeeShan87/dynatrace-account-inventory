# \GeographicRegionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCities**](GeographicRegionsAPI.md#GetCities) | **Get** /rum/cities/{countryCode}/{regionCode} | Lists all cities of a country&#39;s region
[**GetCountries**](GeographicRegionsAPI.md#GetCountries) | **Get** /rum/countries | Lists all countries
[**GetRegions**](GeographicRegionsAPI.md#GetRegions) | **Get** /rum/regions | Lists all countries along with their regions
[**GetRegionsAndCities**](GeographicRegionsAPI.md#GetRegionsAndCities) | **Get** /rum/cities/{countryCode} | Lists all cities of a country
[**GetRegionsOfCountry**](GeographicRegionsAPI.md#GetRegionsOfCountry) | **Get** /rum/regions/{countryCode} | Lists all regions of a country



## GetCities

> CountryWithRegionsWithCities GetCities(ctx, countryCode, regionCode).Execute()

Lists all cities of a country's region

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
    countryCode := "countryCode_example" // string | The ISO code of the required country.    To fetch the list of available country codes, use the [GET all countries](https://dt-url.net/37030go) request.
    regionCode := "regionCode_example" // string | The code of the required region.    To fetch the list of available region codes, use the [GET regions of the country](https://dt-url.net/az230x0) request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeographicRegionsAPI.GetCities(context.Background(), countryCode, regionCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographicRegionsAPI.GetCities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCities`: CountryWithRegionsWithCities
    fmt.Fprintf(os.Stdout, "Response from `GeographicRegionsAPI.GetCities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | The ISO code of the required country.    To fetch the list of available country codes, use the [GET all countries](https://dt-url.net/37030go) request. | 
**regionCode** | **string** | The code of the required region.    To fetch the list of available region codes, use the [GET regions of the country](https://dt-url.net/az230x0) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CountryWithRegionsWithCities**](CountryWithRegionsWithCities.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCountries

> CountryList GetCountries(ctx).Execute()

Lists all countries

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
    resp, r, err := apiClient.GeographicRegionsAPI.GetCountries(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographicRegionsAPI.GetCountries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCountries`: CountryList
    fmt.Fprintf(os.Stdout, "Response from `GeographicRegionsAPI.GetCountries`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCountriesRequest struct via the builder pattern


### Return type

[**CountryList**](CountryList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegions

> CountryListWithRegions GetRegions(ctx).Execute()

Lists all countries along with their regions

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
    resp, r, err := apiClient.GeographicRegionsAPI.GetRegions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographicRegionsAPI.GetRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegions`: CountryListWithRegions
    fmt.Fprintf(os.Stdout, "Response from `GeographicRegionsAPI.GetRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsRequest struct via the builder pattern


### Return type

[**CountryListWithRegions**](CountryListWithRegions.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegionsAndCities

> CountryWithRegionsWithCities GetRegionsAndCities(ctx, countryCode).Execute()

Lists all cities of a country

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
    countryCode := "countryCode_example" // string | The ISO code of the required country.    To fetch the list of available country codes, use the [GET all countries](https://dt-url.net/37030go) request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeographicRegionsAPI.GetRegionsAndCities(context.Background(), countryCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographicRegionsAPI.GetRegionsAndCities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegionsAndCities`: CountryWithRegionsWithCities
    fmt.Fprintf(os.Stdout, "Response from `GeographicRegionsAPI.GetRegionsAndCities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | The ISO code of the required country.    To fetch the list of available country codes, use the [GET all countries](https://dt-url.net/37030go) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsAndCitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CountryWithRegionsWithCities**](CountryWithRegionsWithCities.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegionsOfCountry

> CountryWithRegions GetRegionsOfCountry(ctx, countryCode).Execute()

Lists all regions of a country

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
    countryCode := "countryCode_example" // string | The ISO code of the required country.    To fetch the list of available country codes, use the [GET all countries](https://dt-url.net/37030go) request.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GeographicRegionsAPI.GetRegionsOfCountry(context.Background(), countryCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeographicRegionsAPI.GetRegionsOfCountry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegionsOfCountry`: CountryWithRegions
    fmt.Fprintf(os.Stdout, "Response from `GeographicRegionsAPI.GetRegionsOfCountry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | The ISO code of the required country.    To fetch the list of available country codes, use the [GET all countries](https://dt-url.net/37030go) request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionsOfCountryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CountryWithRegions**](CountryWithRegions.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

