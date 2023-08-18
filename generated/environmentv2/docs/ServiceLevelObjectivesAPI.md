# \ServiceLevelObjectivesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlert**](ServiceLevelObjectivesAPI.md#CreateAlert) | **Post** /slo/{id}/alert | Creates an alert of the provided alert type for an SLO
[**CreateSlo**](ServiceLevelObjectivesAPI.md#CreateSlo) | **Post** /slo | Creates a new SLO
[**DeleteSlo**](ServiceLevelObjectivesAPI.md#DeleteSlo) | **Delete** /slo/{id} | Deletes an SLO
[**GetSlo**](ServiceLevelObjectivesAPI.md#GetSlo) | **Get** /slo | Lists all available SLOs along with calculated values
[**GetSloById**](ServiceLevelObjectivesAPI.md#GetSloById) | **Get** /slo/{id} | Gets parameters and calculated values of a specific SLO
[**UpdateSloById**](ServiceLevelObjectivesAPI.md#UpdateSloById) | **Put** /slo/{id} | Updates an existing SLO



## CreateAlert

> EntityShortRepresentation CreateAlert(ctx, id).AbstractSloAlertDto(abstractSloAlertDto).From(from).To(to).TimeFrame(timeFrame).Execute()

Creates an alert of the provided alert type for an SLO

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required SLO.
    abstractSloAlertDto := *openapiclient.NewAbstractSloAlertDto("AlertName_example", float64(123), "AlertType_example") // AbstractSloAlertDto | The JSON body of the request. Contains the parameters of the new SLO alert.
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    timeFrame := "timeFrame_example" // string | The timeframe to calculate the SLO values:   * `CURRENT`: SLO's own timeframe.  * `GTF`: timeframe specified by **from** and **to** parameters.   If not set, the `CURRENT` value is used. (optional) (default to "CURRENT")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesAPI.CreateAlert(context.Background(), id).AbstractSloAlertDto(abstractSloAlertDto).From(from).To(to).TimeFrame(timeFrame).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesAPI.CreateAlert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlert`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesAPI.CreateAlert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required SLO. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **abstractSloAlertDto** | [**AbstractSloAlertDto**](AbstractSloAlertDto.md) | The JSON body of the request. Contains the parameters of the new SLO alert. | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **timeFrame** | **string** | The timeframe to calculate the SLO values:   * &#x60;CURRENT&#x60;: SLO&#39;s own timeframe.  * &#x60;GTF&#x60;: timeframe specified by **from** and **to** parameters.   If not set, the &#x60;CURRENT&#x60; value is used. | [default to &quot;CURRENT&quot;]

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


## CreateSlo

> CreateSlo(ctx).SloConfigItemDtoImpl(sloConfigItemDtoImpl).Execute()

Creates a new SLO

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
    sloConfigItemDtoImpl := *openapiclient.NewSloConfigItemDtoImpl("AGGREGATE", "Payment service availability", float64(95.0), "-1d", float64(97.5)) // SloConfigItemDtoImpl | The JSON body of the request. Contains the parameters of the new SLO.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceLevelObjectivesAPI.CreateSlo(context.Background()).SloConfigItemDtoImpl(sloConfigItemDtoImpl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesAPI.CreateSlo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sloConfigItemDtoImpl** | [**SloConfigItemDtoImpl**](SloConfigItemDtoImpl.md) | The JSON body of the request. Contains the parameters of the new SLO. | 

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


## DeleteSlo

> DeleteSlo(ctx, id).Execute()

Deletes an SLO

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required SLO.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceLevelObjectivesAPI.DeleteSlo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesAPI.DeleteSlo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required SLO. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSloRequest struct via the builder pattern


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


## GetSlo

> SLOs GetSlo(ctx).NextPageKey(nextPageKey).PageSize(pageSize).From(from).To(to).SloSelector(sloSelector).Sort(sort).TimeFrame(timeFrame).Demo(demo).Evaluate(evaluate).EnabledSlos(enabledSlos).ShowGlobalSlos(showGlobalSlos).Execute()

Lists all available SLOs along with calculated values



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
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The amount of SLOs in a single response payload.   The maximal allowed page size is 10000.   If not set, 10 is used. (optional)
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    sloSelector := "sloSelector_example" // string | The scope of the query. Only SLOs matching the provided criteria are included in the response.   You can add one or several of the criteria listed below. * SLO ID: `id(\"id-1\",\"id-2\")`. * Name: `name(\"name\")`. Filters for an SLO with the given name. The filter is case-sensitive. * Health state: `healthState(\"HEALTHY\")` or `healthState(\"UNHEALTHY\")`. Filters for SLOs that have no related open problems (`HEALTHY`) or SLOs that have related open problems (`UNHEALTHY`). You can specify only one health state. * Text: `text(\"value\")`. Filters for all SLOs that contain the given value in their name or description. The filter is case-insensitive. * Problem: `problemDisplayName(\"value\")`. Filters for all SLOs that are related to a given problem display name (e.g. P-12345). * Management zone name: `managementZone(\"MZ-A\")`. Filters for all SLOs that are related to the given management zone name. * Management zone ID: `managementZoneID(\"123\")`. Filters for all SLOs that are related to the given management zone ID.  To set several criteria, separate them with comma (`,`). Only SLOs matching all criteria are included in the response. Examples: * .../api/v2/slo?sloSelector=name(\"Service Availability\") * .../api/v2/slo?sloSelector=id(\"id\") * .../api/v2/slo?sloSelector=text(\"Description\"),healthState(\"HEALTHY\").  The special characters `~` and `\"` need to be escaped using a `~` (e.g. double quote search `text(\"~\"\")`). (optional) (default to "")
    sort := "sort_example" // string | The sorting of SLO entries:  * `name`: Names in ascending order.  * `-name`: Names in descending order.   If not set, the ascending order is used. (optional) (default to "name")
    timeFrame := "timeFrame_example" // string | The timeframe to calculate the SLO values:   * `CURRENT`: SLO's own timeframe.  * `GTF`: timeframe specified by **from** and **to** parameters.   If not set, the `CURRENT` value is used. (optional) (default to "CURRENT")
    demo := true // bool | Get your SLOs (`false`) or a set of demo SLOs (`true`). (optional) (default to false)
    evaluate := "evaluate_example" // string | Get your SLOs without them being evaluated (`false`) or with evaluations (`true`) with maximum `pageSize` of 25. (optional) (default to "false")
    enabledSlos := "enabledSlos_example" // string | Get your enabled SLOs (`true`), disabled ones (`false`) or both enabled and disabled ones (`all`). (optional) (default to "true")
    showGlobalSlos := true // bool | Get your global SLOs (`true`) regardless of the selected filter or filter them out (`false`). (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesAPI.GetSlo(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).From(from).To(to).SloSelector(sloSelector).Sort(sort).TimeFrame(timeFrame).Demo(demo).Evaluate(evaluate).EnabledSlos(enabledSlos).ShowGlobalSlos(showGlobalSlos).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesAPI.GetSlo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSlo`: SLOs
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesAPI.GetSlo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSloRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of SLOs in a single response payload.   The maximal allowed page size is 10000.   If not set, 10 is used. | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **sloSelector** | **string** | The scope of the query. Only SLOs matching the provided criteria are included in the response.   You can add one or several of the criteria listed below. * SLO ID: &#x60;id(\&quot;id-1\&quot;,\&quot;id-2\&quot;)&#x60;. * Name: &#x60;name(\&quot;name\&quot;)&#x60;. Filters for an SLO with the given name. The filter is case-sensitive. * Health state: &#x60;healthState(\&quot;HEALTHY\&quot;)&#x60; or &#x60;healthState(\&quot;UNHEALTHY\&quot;)&#x60;. Filters for SLOs that have no related open problems (&#x60;HEALTHY&#x60;) or SLOs that have related open problems (&#x60;UNHEALTHY&#x60;). You can specify only one health state. * Text: &#x60;text(\&quot;value\&quot;)&#x60;. Filters for all SLOs that contain the given value in their name or description. The filter is case-insensitive. * Problem: &#x60;problemDisplayName(\&quot;value\&quot;)&#x60;. Filters for all SLOs that are related to a given problem display name (e.g. P-12345). * Management zone name: &#x60;managementZone(\&quot;MZ-A\&quot;)&#x60;. Filters for all SLOs that are related to the given management zone name. * Management zone ID: &#x60;managementZoneID(\&quot;123\&quot;)&#x60;. Filters for all SLOs that are related to the given management zone ID.  To set several criteria, separate them with comma (&#x60;,&#x60;). Only SLOs matching all criteria are included in the response. Examples: * .../api/v2/slo?sloSelector&#x3D;name(\&quot;Service Availability\&quot;) * .../api/v2/slo?sloSelector&#x3D;id(\&quot;id\&quot;) * .../api/v2/slo?sloSelector&#x3D;text(\&quot;Description\&quot;),healthState(\&quot;HEALTHY\&quot;).  The special characters &#x60;~&#x60; and &#x60;\&quot;&#x60; need to be escaped using a &#x60;~&#x60; (e.g. double quote search &#x60;text(\&quot;~\&quot;\&quot;)&#x60;). | [default to &quot;&quot;]
 **sort** | **string** | The sorting of SLO entries:  * &#x60;name&#x60;: Names in ascending order.  * &#x60;-name&#x60;: Names in descending order.   If not set, the ascending order is used. | [default to &quot;name&quot;]
 **timeFrame** | **string** | The timeframe to calculate the SLO values:   * &#x60;CURRENT&#x60;: SLO&#39;s own timeframe.  * &#x60;GTF&#x60;: timeframe specified by **from** and **to** parameters.   If not set, the &#x60;CURRENT&#x60; value is used. | [default to &quot;CURRENT&quot;]
 **demo** | **bool** | Get your SLOs (&#x60;false&#x60;) or a set of demo SLOs (&#x60;true&#x60;). | [default to false]
 **evaluate** | **string** | Get your SLOs without them being evaluated (&#x60;false&#x60;) or with evaluations (&#x60;true&#x60;) with maximum &#x60;pageSize&#x60; of 25. | [default to &quot;false&quot;]
 **enabledSlos** | **string** | Get your enabled SLOs (&#x60;true&#x60;), disabled ones (&#x60;false&#x60;) or both enabled and disabled ones (&#x60;all&#x60;). | [default to &quot;true&quot;]
 **showGlobalSlos** | **bool** | Get your global SLOs (&#x60;true&#x60;) regardless of the selected filter or filter them out (&#x60;false&#x60;). | [default to true]

### Return type

[**SLOs**](SLOs.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSloById

> SLO GetSloById(ctx, id).From(from).To(to).TimeFrame(timeFrame).Execute()

Gets parameters and calculated values of a specific SLO



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required SLO.
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    timeFrame := "timeFrame_example" // string | The timeframe to calculate the SLO values:   * `CURRENT`: SLO's own timeframe.  * `GTF`: timeframe specified by **from** and **to** parameters.   If not set, the `CURRENT` value is used. (optional) (default to "CURRENT")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceLevelObjectivesAPI.GetSloById(context.Background(), id).From(from).To(to).TimeFrame(timeFrame).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesAPI.GetSloById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSloById`: SLO
    fmt.Fprintf(os.Stdout, "Response from `ServiceLevelObjectivesAPI.GetSloById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required SLO. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSloByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **timeFrame** | **string** | The timeframe to calculate the SLO values:   * &#x60;CURRENT&#x60;: SLO&#39;s own timeframe.  * &#x60;GTF&#x60;: timeframe specified by **from** and **to** parameters.   If not set, the &#x60;CURRENT&#x60; value is used. | [default to &quot;CURRENT&quot;]

### Return type

[**SLO**](SLO.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSloById

> UpdateSloById(ctx, id).SloConfigItemDtoImpl(sloConfigItemDtoImpl).Execute()

Updates an existing SLO

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required SLO.
    sloConfigItemDtoImpl := *openapiclient.NewSloConfigItemDtoImpl("AGGREGATE", "Payment service availability", float64(95.0), "-1d", float64(97.5)) // SloConfigItemDtoImpl | The JSON body of the request. Contains the updated parameters of the SLO.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceLevelObjectivesAPI.UpdateSloById(context.Background(), id).SloConfigItemDtoImpl(sloConfigItemDtoImpl).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceLevelObjectivesAPI.UpdateSloById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required SLO. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSloByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sloConfigItemDtoImpl** | [**SloConfigItemDtoImpl**](SloConfigItemDtoImpl.md) | The JSON body of the request. Contains the updated parameters of the SLO. | 

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

