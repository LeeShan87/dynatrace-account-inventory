# \EventsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEvent**](EventsAPI.md#CreateEvent) | **Post** /events/ingest | Ingests a custom event
[**GetEvent**](EventsAPI.md#GetEvent) | **Get** /events/{eventId} | Gets the properties of an event
[**GetEventProperties**](EventsAPI.md#GetEventProperties) | **Get** /eventProperties | Lists all event properties
[**GetEventProperty**](EventsAPI.md#GetEventProperty) | **Get** /eventProperties/{propertyKey} | Gets the details of an event property
[**GetEventType**](EventsAPI.md#GetEventType) | **Get** /eventTypes/{eventType} | Gets the properties of an event type
[**GetEventTypes**](EventsAPI.md#GetEventTypes) | **Get** /eventTypes | Lists all event types
[**GetEvents**](EventsAPI.md#GetEvents) | **Get** /events | Lists events within the specified timeframe



## CreateEvent

> EventIngestResults CreateEvent(ctx).EventIngest(eventIngest).Execute()

Ingests a custom event



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
    eventIngest := *openapiclient.NewEventIngest("EventType_example", "Title_example") // EventIngest | The JSON body of the request. Contains properties of the new event. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.CreateEvent(context.Background()).EventIngest(eventIngest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.CreateEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEvent`: EventIngestResults
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.CreateEvent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **eventIngest** | [**EventIngest**](EventIngest.md) | The JSON body of the request. Contains properties of the new event. | 

### Return type

[**EventIngestResults**](EventIngestResults.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvent

> Event GetEvent(ctx, eventId).Execute()

Gets the properties of an event

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
    eventId := "eventId_example" // string | The ID of the required event.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEvent(context.Background(), eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: Event
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventId** | **string** | The ID of the required event. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Event**](Event.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventProperties

> EventPropertiesList GetEventProperties(ctx).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Lists all event properties

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
    pageSize := int64(789) // int64 | The amount of event properties in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEventProperties(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventProperties``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventProperties`: EventPropertiesList
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of event properties in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. | 

### Return type

[**EventPropertiesList**](EventPropertiesList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventProperty

> EventPropertyDetails GetEventProperty(ctx, propertyKey).Execute()

Gets the details of an event property

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
    propertyKey := "propertyKey_example" // string | The event property key you're inquiring.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEventProperty(context.Background(), propertyKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventProperty``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventProperty`: EventPropertyDetails
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyKey** | **string** | The event property key you&#39;re inquiring. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventPropertyDetails**](EventPropertyDetails.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventType

> EventType GetEventType(ctx, eventType).Execute()

Gets the properties of an event type

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
    eventType := "eventType_example" // string | The event type you're inquiring.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEventType(context.Background(), eventType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventType`: EventType
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**eventType** | **string** | The event type you&#39;re inquiring. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventType**](EventType.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventTypes

> EventTypeList GetEventTypes(ctx).NextPageKey(nextPageKey).PageSize(pageSize).Execute()

Lists all event types

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
    pageSize := int64(789) // int64 | The amount of event types in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEventTypes(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEventTypes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventTypes`: EventTypeList
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEventTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of event types in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. | 

### Return type

[**EventTypeList**](EventTypeList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> EventList GetEvents(ctx).NextPageKey(nextPageKey).PageSize(pageSize).From(from).To(to).EventSelector(eventSelector).EntitySelector(entitySelector).Execute()

Lists events within the specified timeframe

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
    pageSize := int64(789) // int64 | The amount of events in a single response payload.   The maximal allowed page size is 1000.   If not set, 100 is used. (optional)
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of two hours is used (`now-2h`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)
    eventSelector := "eventSelector_example" // string | Defines the scope of the query. Only events matching the specified criteria are included in the response.  You can add one or several of the criteria listed below. For each criterion you can specify multiple comma-separated values, unless stated otherwise. If several values are specified, the **OR** logic applies. * Event ID: `eventId(\"id-1\", \"id-2\")`. * ID of related entity: `entityId(\"id-1\", \"id-2\")`. * Event status: `status(\"OPEN\")` or `status(\"CLOSED\")`. You can specify only one status. * Management zone ID: `managementZoneId(\"123\", \"321\")`. * Event type: `eventType(\"event-type\")`. You can specify only one event type. You can fetch the list of possible event types with the [GET event types](https://dt-url.net/qc03u6w) call.  * Correlation ID: `correlationId(\"id-1\", \"id-2\")`. * Happened during maintenance (true, false): `underMaintenance(true)`.  * Notifications are suppressed (true, false): `suppressAlert(true)`.  * Problem creation is suppressed (true, false): `suppressProblem(true)`.  * Frequent event (true, false): `frequentEvent(true)`.  * Event property: `property.<key>(\"value-1\", \"value-2\")`.  To set several criteria, separate them with commas (`,`). Only results matching **all** criteria are included in the response.    (optional)
    entitySelector := "entitySelector_example" // string | The entity scope of the query. You must set one of these criteria:   * Entity type: `type(\"TYPE\")`  * Dynatrace entity ID: `entityId(\"id\")`. You can specify several IDs, separated by a comma (`entityId(\"id-1\",\"id-2\")`). All requested entities must be of the same type.   You can add one or more of the following criteria. Values are case-sensitive and the `EQUALS` operator is used unless otherwise specified.   * Tag: `tag(\"value\")`. Tags in `[context]key:value`, `key:value`, and `value` formats are detected and parsed automatically. Any colons (`:`) that are part of the key or value must be escaped with a backslash(`\\`). Otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: `mzId(123)`  * Management zone name: `mzName(\"value\")` * Entity name:   * `entityName.equals`: performs a non-casesensitive `EQUALS` query.   * `entityName.startsWith`: changes the operator to `BEGINS WITH`.   * `entityName.in`: enables you to provide multiple values. The `EQUALS` operator applies.   * `caseSensitive(entityName.equals(\"value\"))`: takes any entity name criterion as an argument and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): `healthState(\"HEALTHY\")` * First seen timestamp: `firstSeenTms.<operator>(now-3h)`. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * `lte`: earlier than or at the specified time  * `lt`: earlier than the specified time  * `gte`: later than or at the specified time  * `gt`: later than the specified time * Entity attribute: `<attribute>(\"value1\",\"value2\")` and `<attribute>.exists()`. To fetch the list of available attributes, execute the [GET entity type](https://dt-url.net/2ka3ivt) request and check the **properties** field of the response.  * Relationships: `fromRelationships.<relationshipName>()` and `toRelationships.<relationshipName>()`. This criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET entity type](https://dt-url.net/2ka3ivt) request and check the **fromRelationships** and **toRelationships** fields. * Negation: `not(<criterion>)`. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (`,`). For example, `type(\"HOST\"),healthState(\"HEALTHY\")`. Only results matching **all** criteria are included in the response.   The maximum string length is 2,000 characters.   The number of entities that can be selected is limited to 10000. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsAPI.GetEvents(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).From(from).To(to).EventSelector(eventSelector).EntitySelector(entitySelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsAPI.GetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvents`: EventList
    fmt.Fprintf(os.Stdout, "Response from `EventsAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of events in a single response payload.   The maximal allowed page size is 1000.   If not set, 100 is used. | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of two hours is used (&#x60;now-2h&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 
 **eventSelector** | **string** | Defines the scope of the query. Only events matching the specified criteria are included in the response.  You can add one or several of the criteria listed below. For each criterion you can specify multiple comma-separated values, unless stated otherwise. If several values are specified, the **OR** logic applies. * Event ID: &#x60;eventId(\&quot;id-1\&quot;, \&quot;id-2\&quot;)&#x60;. * ID of related entity: &#x60;entityId(\&quot;id-1\&quot;, \&quot;id-2\&quot;)&#x60;. * Event status: &#x60;status(\&quot;OPEN\&quot;)&#x60; or &#x60;status(\&quot;CLOSED\&quot;)&#x60;. You can specify only one status. * Management zone ID: &#x60;managementZoneId(\&quot;123\&quot;, \&quot;321\&quot;)&#x60;. * Event type: &#x60;eventType(\&quot;event-type\&quot;)&#x60;. You can specify only one event type. You can fetch the list of possible event types with the [GET event types](https://dt-url.net/qc03u6w) call.  * Correlation ID: &#x60;correlationId(\&quot;id-1\&quot;, \&quot;id-2\&quot;)&#x60;. * Happened during maintenance (true, false): &#x60;underMaintenance(true)&#x60;.  * Notifications are suppressed (true, false): &#x60;suppressAlert(true)&#x60;.  * Problem creation is suppressed (true, false): &#x60;suppressProblem(true)&#x60;.  * Frequent event (true, false): &#x60;frequentEvent(true)&#x60;.  * Event property: &#x60;property.&lt;key&gt;(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;.  To set several criteria, separate them with commas (&#x60;,&#x60;). Only results matching **all** criteria are included in the response.    | 
 **entitySelector** | **string** | The entity scope of the query. You must set one of these criteria:   * Entity type: &#x60;type(\&quot;TYPE\&quot;)&#x60;  * Dynatrace entity ID: &#x60;entityId(\&quot;id\&quot;)&#x60;. You can specify several IDs, separated by a comma (&#x60;entityId(\&quot;id-1\&quot;,\&quot;id-2\&quot;)&#x60;). All requested entities must be of the same type.   You can add one or more of the following criteria. Values are case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.   * Tag: &#x60;tag(\&quot;value\&quot;)&#x60;. Tags in &#x60;[context]key:value&#x60;, &#x60;key:value&#x60;, and &#x60;value&#x60; formats are detected and parsed automatically. Any colons (&#x60;:&#x60;) that are part of the key or value must be escaped with a backslash(&#x60;\\&#x60;). Otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: &#x60;mzId(123)&#x60;  * Management zone name: &#x60;mzName(\&quot;value\&quot;)&#x60; * Entity name:   * &#x60;entityName.equals&#x60;: performs a non-casesensitive &#x60;EQUALS&#x60; query.   * &#x60;entityName.startsWith&#x60;: changes the operator to &#x60;BEGINS WITH&#x60;.   * &#x60;entityName.in&#x60;: enables you to provide multiple values. The &#x60;EQUALS&#x60; operator applies.   * &#x60;caseSensitive(entityName.equals(\&quot;value\&quot;))&#x60;: takes any entity name criterion as an argument and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): &#x60;healthState(\&quot;HEALTHY\&quot;)&#x60; * First seen timestamp: &#x60;firstSeenTms.&lt;operator&gt;(now-3h)&#x60;. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * &#x60;lte&#x60;: earlier than or at the specified time  * &#x60;lt&#x60;: earlier than the specified time  * &#x60;gte&#x60;: later than or at the specified time  * &#x60;gt&#x60;: later than the specified time * Entity attribute: &#x60;&lt;attribute&gt;(\&quot;value1\&quot;,\&quot;value2\&quot;)&#x60; and &#x60;&lt;attribute&gt;.exists()&#x60;. To fetch the list of available attributes, execute the [GET entity type](https://dt-url.net/2ka3ivt) request and check the **properties** field of the response.  * Relationships: &#x60;fromRelationships.&lt;relationshipName&gt;()&#x60; and &#x60;toRelationships.&lt;relationshipName&gt;()&#x60;. This criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET entity type](https://dt-url.net/2ka3ivt) request and check the **fromRelationships** and **toRelationships** fields. * Negation: &#x60;not(&lt;criterion&gt;)&#x60;. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (&#x60;,&#x60;). For example, &#x60;type(\&quot;HOST\&quot;),healthState(\&quot;HEALTHY\&quot;)&#x60;. Only results matching **all** criteria are included in the response.   The maximum string length is 2,000 characters.   The number of entities that can be selected is limited to 10000. | 

### Return type

[**EventList**](EventList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

