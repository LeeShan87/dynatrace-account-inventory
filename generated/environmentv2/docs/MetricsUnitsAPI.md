# \MetricsUnitsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AllUnits**](MetricsUnitsAPI.md#AllUnits) | **Get** /units | Lists all available units
[**Convert**](MetricsUnitsAPI.md#Convert) | **Get** /units/{unitId}/convert | Converts a value from a source unit into a target unit
[**Unit**](MetricsUnitsAPI.md#Unit) | **Get** /units/{unitId} | Gets the properties of the specified unit



## AllUnits

> UnitList AllUnits(ctx).UnitSelector(unitSelector).Fields(fields).Execute()

Lists all available units



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
    unitSelector := "unitSelector_example" // string | Selects units to be included to the response. Available criteria:  * Compatibility: `compatibleTo(\"unit\",\"display-format\")`. Returns units that can be converted to the specified unit. The optional display format (`binary` or `decimal`) argument is supported by bit- and byte-based units and returns only units for the specified format. (optional)
    fields := "fields_example" // string | Defines the list of properties to be included in the response. The ID of the unit is **always** included. The following additional properties are available:   * `displayName`: The display name of the unit.  * `symbol`: The symbol of the unit.  * `description`: A short description of the unit.  By default, the ID, the display name, and the symbol are included.   To add properties, list them with leading plus `+`. To exclude default properties, list them with leading minus `-`.  To specify several properties, join them with a comma (for example `fields=+description,-symbol`).  If you specify just one property, the response contains the unitId and the specified property. To return unit IDs only, specify `unitId` here. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsUnitsAPI.AllUnits(context.Background()).UnitSelector(unitSelector).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsUnitsAPI.AllUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AllUnits`: UnitList
    fmt.Fprintf(os.Stdout, "Response from `MetricsUnitsAPI.AllUnits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAllUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **unitSelector** | **string** | Selects units to be included to the response. Available criteria:  * Compatibility: &#x60;compatibleTo(\&quot;unit\&quot;,\&quot;display-format\&quot;)&#x60;. Returns units that can be converted to the specified unit. The optional display format (&#x60;binary&#x60; or &#x60;decimal&#x60;) argument is supported by bit- and byte-based units and returns only units for the specified format. | 
 **fields** | **string** | Defines the list of properties to be included in the response. The ID of the unit is **always** included. The following additional properties are available:   * &#x60;displayName&#x60;: The display name of the unit.  * &#x60;symbol&#x60;: The symbol of the unit.  * &#x60;description&#x60;: A short description of the unit.  By default, the ID, the display name, and the symbol are included.   To add properties, list them with leading plus &#x60;+&#x60;. To exclude default properties, list them with leading minus &#x60;-&#x60;.  To specify several properties, join them with a comma (for example &#x60;fields&#x3D;+description,-symbol&#x60;).  If you specify just one property, the response contains the unitId and the specified property. To return unit IDs only, specify &#x60;unitId&#x60; here. | 

### Return type

[**UnitList**](UnitList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Convert

> UnitConversionResult Convert(ctx, unitId).Value(value).TargetUnit(targetUnit).NumberFormat(numberFormat).Execute()

Converts a value from a source unit into a target unit



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
    unitId := "unitId_example" // string | The ID of the source unit.
    value := float64(1.2) // float64 | The value to be converted.
    targetUnit := "targetUnit_example" // string | The ID of the target unit.   If not set, the request finds an appropriate target unit automatically, based on the specified number format. (optional)
    numberFormat := "numberFormat_example" // string | The preferred number format of the target value. You can specify the following formats:   * `binary`  * `decimal`   `Only used if the target unit if not set. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsUnitsAPI.Convert(context.Background(), unitId).Value(value).TargetUnit(targetUnit).NumberFormat(numberFormat).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsUnitsAPI.Convert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Convert`: UnitConversionResult
    fmt.Fprintf(os.Stdout, "Response from `MetricsUnitsAPI.Convert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unitId** | **string** | The ID of the source unit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConvertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **value** | **float64** | The value to be converted. | 
 **targetUnit** | **string** | The ID of the target unit.   If not set, the request finds an appropriate target unit automatically, based on the specified number format. | 
 **numberFormat** | **string** | The preferred number format of the target value. You can specify the following formats:   * &#x60;binary&#x60;  * &#x60;decimal&#x60;   &#x60;Only used if the target unit if not set. | 

### Return type

[**UnitConversionResult**](UnitConversionResult.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Unit

> Unit Unit(ctx, unitId).Execute()

Gets the properties of the specified unit

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
    unitId := "unitId_example" // string | The ID of the required unit.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricsUnitsAPI.Unit(context.Background(), unitId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsUnitsAPI.Unit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Unit`: Unit
    fmt.Fprintf(os.Stdout, "Response from `MetricsUnitsAPI.Unit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**unitId** | **string** | The ID of the required unit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Unit**](Unit.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

