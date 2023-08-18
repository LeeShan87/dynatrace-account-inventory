# \LogsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportLogRecords**](LogsAPI.md#ExportLogRecords) | **Get** /logs/export | Exports log records | maturity&#x3D;EARLY_ADOPTER
[**GetLogHistogramData**](LogsAPI.md#GetLogHistogramData) | **Get** /logs/aggregate | Gets aggregated log records | maturity&#x3D;EARLY_ADOPTER
[**GetLogRecords**](LogsAPI.md#GetLogRecords) | **Get** /logs/search | Reads log records | maturity&#x3D;EARLY_ADOPTER
[**StoreLog**](LogsAPI.md#StoreLog) | **Post** /logs/ingest | Pushes log records to Dynatrace



## ExportLogRecords

> ExportedLogRecordList ExportLogRecords(ctx).From(from).To(to).NextPageKey(nextPageKey).PageSize(pageSize).Query(query).Sort(sort).Execute()

Exports log records | maturity=EARLY_ADOPTER



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
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don't specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  (optional)
    pageSize := int64(789) // int64 | The number of results per result page. (optional) (default to 1000)
    query := "query_example" // string | The log search query.   The query must use the [Dynatrace search query language](https://dt-url.net/pe03s6f). (optional) (default to "")
    sort := "sort_example" // string | Defines the ordering of the log records.  Each field has a sign prefix (+/-) for sorting order. If no sign prefix is set, then the `+` option will be applied.   Currently, ordering is available only for the timestamp (+timestamp for the oldest records first, or -timestamp for the newest records first).   When millisecond resolution provided by the timestamp is not enough, log records are sorted based on the order in which they appear in the log source (remote process writing to REST API endpoint or remote process from which logs are collected). (optional) (default to "-timestamp")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsAPI.ExportLogRecords(context.Background()).From(from).To(to).NextPageKey(nextPageKey).PageSize(pageSize).Query(query).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.ExportLogRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportLogRecords`: ExportedLogRecordList
    fmt.Fprintf(os.Stdout, "Response from `LogsAPI.ExportLogRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportLogRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The number of results per result page. | [default to 1000]
 **query** | **string** | The log search query.   The query must use the [Dynatrace search query language](https://dt-url.net/pe03s6f). | [default to &quot;&quot;]
 **sort** | **string** | Defines the ordering of the log records.  Each field has a sign prefix (+/-) for sorting order. If no sign prefix is set, then the &#x60;+&#x60; option will be applied.   Currently, ordering is available only for the timestamp (+timestamp for the oldest records first, or -timestamp for the newest records first).   When millisecond resolution provided by the timestamp is not enough, log records are sorted based on the order in which they appear in the log source (remote process writing to REST API endpoint or remote process from which logs are collected). | [default to &quot;-timestamp&quot;]

### Return type

[**ExportedLogRecordList**](ExportedLogRecordList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogHistogramData

> AggregatedLog GetLogHistogramData(ctx).From(from).To(to).Query(query).TimeBuckets(timeBuckets).MaxGroupValues(maxGroupValues).GroupBy(groupBy).Execute()

Gets aggregated log records | maturity=EARLY_ADOPTER



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
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    query := "query_example" // string | The log search query.   The query must use the [Dynatrace search query language](https://dt-url.net/pe03s6f). (optional) (default to "")
    timeBuckets := int32(56) // int32 | The number of time slots in the result.   The query timeframe is divided equally into the specified number of slots.   The minimum length of a slot is 1 ms.   If not set, 1 is used. (optional) (default to 1)
    maxGroupValues := int32(56) // int32 | The maximum number of values in each group.   You can get up to 100 values per group.   If not set, 10 is used. (optional) (default to 10)
    groupBy := []string{"Inner_example"} // []string | The groupings to be included in the response.   You can specify several groups in the following format: `groupBy=status&groupBy=log.source`.   If not set, all possible groups are returned. You can use this option to check for possible grouping values.   Unique log data attributes (high-cardinality attributes) for example, `span_id` or `trace_id` cannot be used for grouping. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsAPI.GetLogHistogramData(context.Background()).From(from).To(to).Query(query).TimeBuckets(timeBuckets).MaxGroupValues(maxGroupValues).GroupBy(groupBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetLogHistogramData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogHistogramData`: AggregatedLog
    fmt.Fprintf(os.Stdout, "Response from `LogsAPI.GetLogHistogramData`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogHistogramDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **query** | **string** | The log search query.   The query must use the [Dynatrace search query language](https://dt-url.net/pe03s6f). | [default to &quot;&quot;]
 **timeBuckets** | **int32** | The number of time slots in the result.   The query timeframe is divided equally into the specified number of slots.   The minimum length of a slot is 1 ms.   If not set, 1 is used. | [default to 1]
 **maxGroupValues** | **int32** | The maximum number of values in each group.   You can get up to 100 values per group.   If not set, 10 is used. | [default to 10]
 **groupBy** | **[]string** | The groupings to be included in the response.   You can specify several groups in the following format: &#x60;groupBy&#x3D;status&amp;groupBy&#x3D;log.source&#x60;.   If not set, all possible groups are returned. You can use this option to check for possible grouping values.   Unique log data attributes (high-cardinality attributes) for example, &#x60;span_id&#x60; or &#x60;trace_id&#x60; cannot be used for grouping. | 

### Return type

[**AggregatedLog**](AggregatedLog.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogRecords

> LogRecordsList GetLogRecords(ctx).From(from).To(to).Limit(limit).Query(query).Sort(sort).NextSliceKey(nextSliceKey).Execute()

Reads log records | maturity=EARLY_ADOPTER



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
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two weeks is used (`now-2w`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    limit := int32(56) // int32 | The desired amount of log records.   The maximal allowed limit is 1000.   If not set, 1000 is used. (optional) (default to 1000)
    query := "query_example" // string | The log search query.   The query must use the [Dynatrace search query language](https://dt-url.net/pe03s6f). (optional) (default to "")
    sort := "sort_example" // string | Defines the ordering of the log records.  Each field has a sign prefix (+/-) for sorting order. If no sign prefix is set, then the `+` option will be applied.   Currently, ordering is available only for the timestamp (+timestamp for the oldest records first, or -timestamp for the newest records first).   When millisecond resolution provided by the timestamp is not enough, log records are sorted based on the order in which they appear in the log source (remote process writing to REST API endpoint or remote process from which logs are collected). (optional) (default to "-timestamp")
    nextSliceKey := "nextSliceKey_example" // string | The cursor for the next slice of results. You can find it in the **nextSliceKey** field of the previous response.   The first slice is always returned if you don't specify this parameter.   If this parameter is set, all other query parameters are ignored. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsAPI.GetLogRecords(context.Background()).From(from).To(to).Limit(limit).Query(query).Sort(sort).NextSliceKey(nextSliceKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.GetLogRecords``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogRecords`: LogRecordsList
    fmt.Fprintf(os.Stdout, "Response from `LogsAPI.GetLogRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two weeks is used (&#x60;now-2w&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **limit** | **int32** | The desired amount of log records.   The maximal allowed limit is 1000.   If not set, 1000 is used. | [default to 1000]
 **query** | **string** | The log search query.   The query must use the [Dynatrace search query language](https://dt-url.net/pe03s6f). | [default to &quot;&quot;]
 **sort** | **string** | Defines the ordering of the log records.  Each field has a sign prefix (+/-) for sorting order. If no sign prefix is set, then the &#x60;+&#x60; option will be applied.   Currently, ordering is available only for the timestamp (+timestamp for the oldest records first, or -timestamp for the newest records first).   When millisecond resolution provided by the timestamp is not enough, log records are sorted based on the order in which they appear in the log source (remote process writing to REST API endpoint or remote process from which logs are collected). | [default to &quot;-timestamp&quot;]
 **nextSliceKey** | **string** | The cursor for the next slice of results. You can find it in the **nextSliceKey** field of the previous response.   The first slice is always returned if you don&#39;t specify this parameter.   If this parameter is set, all other query parameters are ignored. | 

### Return type

[**LogRecordsList**](LogRecordsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreLog

> SuccessEnvelope StoreLog(ctx).Body(body).Execute()

Pushes log records to Dynatrace



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
    body := map[string]interface{}{ ... } // map[string]interface{} | The body of the request. Contains one or more log events to be ingested.   The endpoint accepts one of the following payload types, defined by the **Accept** header:   * `text/plain`: supports only one log event.  * `application/json`: supports multiple log events in a single payload. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogsAPI.StoreLog(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogsAPI.StoreLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StoreLog`: SuccessEnvelope
    fmt.Fprintf(os.Stdout, "Response from `LogsAPI.StoreLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | The body of the request. Contains one or more log events to be ingested.   The endpoint accepts one of the following payload types, defined by the **Accept** header:   * &#x60;text/plain&#x60;: supports only one log event.  * &#x60;application/json&#x60;: supports multiple log events in a single payload. | 

### Return type

[**SuccessEnvelope**](SuccessEnvelope.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8, text/plain; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

