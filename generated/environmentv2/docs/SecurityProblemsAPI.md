# \SecurityProblemsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkMuteRemediationItems**](SecurityProblemsAPI.md#BulkMuteRemediationItems) | **Post** /securityProblems/{id}/remediationItems/mute | Mutes several remediation items
[**BulkMuteSecurityProblems**](SecurityProblemsAPI.md#BulkMuteSecurityProblems) | **Post** /securityProblems/mute | Mutes several security problems
[**BulkUnmuteRemediationItems**](SecurityProblemsAPI.md#BulkUnmuteRemediationItems) | **Post** /securityProblems/{id}/remediationItems/unmute | Un-mutes several remediation items
[**BulkUnmuteSecurityProblems**](SecurityProblemsAPI.md#BulkUnmuteSecurityProblems) | **Post** /securityProblems/unmute | Un-mutes several security problems
[**GetEventsForSecurityProblem**](SecurityProblemsAPI.md#GetEventsForSecurityProblem) | **Get** /securityProblems/{id}/events | Lists all events of a security problem
[**GetRemediationItem**](SecurityProblemsAPI.md#GetRemediationItem) | **Get** /securityProblems/{id}/remediationItems/{remediationItemId} | Gets parameters of a remediation item of a security problem
[**GetRemediationItems**](SecurityProblemsAPI.md#GetRemediationItems) | **Get** /securityProblems/{id}/remediationItems | Lists remediation items of a third-party security problem
[**GetRemediationProgressEntities**](SecurityProblemsAPI.md#GetRemediationProgressEntities) | **Get** /securityProblems/{id}/remediationItems/{remediationItemId}/remediationProgressEntities | Lists remediation progress entities
[**GetSecurityProblem**](SecurityProblemsAPI.md#GetSecurityProblem) | **Get** /securityProblems/{id} | Gets parameters of a security problem
[**GetSecurityProblems**](SecurityProblemsAPI.md#GetSecurityProblems) | **Get** /securityProblems | Lists all security problems
[**GetVulnerableFunctions**](SecurityProblemsAPI.md#GetVulnerableFunctions) | **Get** /securityProblems/{id}/vulnerableFunctions | Lists all vulnerable functions and their usage for a third-party security problem
[**MuteSecurityProblem**](SecurityProblemsAPI.md#MuteSecurityProblem) | **Post** /securityProblems/{id}/mute | Mutes a security problem
[**SetRemediationItemMuteState**](SecurityProblemsAPI.md#SetRemediationItemMuteState) | **Put** /securityProblems/{id}/remediationItems/{remediationItemId}/muteState | Sets the mute state of a remediation item
[**UnmuteSecurityProblem**](SecurityProblemsAPI.md#UnmuteSecurityProblem) | **Post** /securityProblems/{id}/unmute | Un-mutes a security problem



## BulkMuteRemediationItems

> RemediationItemsBulkMuteResponse BulkMuteRemediationItems(ctx, id).RemediationItemsBulkMute(remediationItemsBulkMute).Execute()

Mutes several remediation items

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
    id := "id_example" // string | The ID of the required third-party security problem.
    remediationItemsBulkMute := *openapiclient.NewRemediationItemsBulkMute("Reason_example", []string{"RemediationItemIds_example"}) // RemediationItemsBulkMute | The JSON body of the request. Contains the muting information. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.BulkMuteRemediationItems(context.Background(), id).RemediationItemsBulkMute(remediationItemsBulkMute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.BulkMuteRemediationItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkMuteRemediationItems`: RemediationItemsBulkMuteResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.BulkMuteRemediationItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required third-party security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkMuteRemediationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remediationItemsBulkMute** | [**RemediationItemsBulkMute**](RemediationItemsBulkMute.md) | The JSON body of the request. Contains the muting information. | 

### Return type

[**RemediationItemsBulkMuteResponse**](RemediationItemsBulkMuteResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkMuteSecurityProblems

> SecurityProblemsBulkMuteResponse BulkMuteSecurityProblems(ctx).SecurityProblemsBulkMute(securityProblemsBulkMute).Execute()

Mutes several security problems

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
    securityProblemsBulkMute := *openapiclient.NewSecurityProblemsBulkMute("Reason_example", []string{"SecurityProblemIds_example"}) // SecurityProblemsBulkMute | The JSON body of the request. Contains the muting information. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.BulkMuteSecurityProblems(context.Background()).SecurityProblemsBulkMute(securityProblemsBulkMute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.BulkMuteSecurityProblems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkMuteSecurityProblems`: SecurityProblemsBulkMuteResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.BulkMuteSecurityProblems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkMuteSecurityProblemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityProblemsBulkMute** | [**SecurityProblemsBulkMute**](SecurityProblemsBulkMute.md) | The JSON body of the request. Contains the muting information. | 

### Return type

[**SecurityProblemsBulkMuteResponse**](SecurityProblemsBulkMuteResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUnmuteRemediationItems

> RemediationItemsBulkUnmuteResponse BulkUnmuteRemediationItems(ctx, id).RemediationItemsBulkUnmute(remediationItemsBulkUnmute).Execute()

Un-mutes several remediation items

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
    id := "id_example" // string | The ID of the required third-party security problem.
    remediationItemsBulkUnmute := *openapiclient.NewRemediationItemsBulkUnmute("Reason_example", []string{"RemediationItemIds_example"}) // RemediationItemsBulkUnmute | The JSON body of the request. Contains the un-muting information. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.BulkUnmuteRemediationItems(context.Background(), id).RemediationItemsBulkUnmute(remediationItemsBulkUnmute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.BulkUnmuteRemediationItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUnmuteRemediationItems`: RemediationItemsBulkUnmuteResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.BulkUnmuteRemediationItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required third-party security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkUnmuteRemediationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remediationItemsBulkUnmute** | [**RemediationItemsBulkUnmute**](RemediationItemsBulkUnmute.md) | The JSON body of the request. Contains the un-muting information. | 

### Return type

[**RemediationItemsBulkUnmuteResponse**](RemediationItemsBulkUnmuteResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUnmuteSecurityProblems

> SecurityProblemsBulkUnmuteResponse BulkUnmuteSecurityProblems(ctx).SecurityProblemsBulkUnmute(securityProblemsBulkUnmute).Execute()

Un-mutes several security problems

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
    securityProblemsBulkUnmute := *openapiclient.NewSecurityProblemsBulkUnmute("Reason_example", []string{"SecurityProblemIds_example"}) // SecurityProblemsBulkUnmute | The JSON body of the request. Contains the un-muting information. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.BulkUnmuteSecurityProblems(context.Background()).SecurityProblemsBulkUnmute(securityProblemsBulkUnmute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.BulkUnmuteSecurityProblems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUnmuteSecurityProblems`: SecurityProblemsBulkUnmuteResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.BulkUnmuteSecurityProblems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUnmuteSecurityProblemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityProblemsBulkUnmute** | [**SecurityProblemsBulkUnmute**](SecurityProblemsBulkUnmute.md) | The JSON body of the request. Contains the un-muting information. | 

### Return type

[**SecurityProblemsBulkUnmuteResponse**](SecurityProblemsBulkUnmuteResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: application/json; charset=utf-8
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsForSecurityProblem

> SecurityProblemEventsList GetEventsForSecurityProblem(ctx, id).From(from).To(to).Execute()

Lists all events of a security problem

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
    id := "id_example" // string | The ID of the required security problem.
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of thirty days is used (`now-30d`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.GetEventsForSecurityProblem(context.Background(), id).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.GetEventsForSecurityProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventsForSecurityProblem`: SecurityProblemEventsList
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.GetEventsForSecurityProblem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsForSecurityProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of thirty days is used (&#x60;now-30d&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 

### Return type

[**SecurityProblemEventsList**](SecurityProblemEventsList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemediationItem

> RemediationDetailsItem GetRemediationItem(ctx, id, remediationItemId).Execute()

Gets parameters of a remediation item of a security problem

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
    id := "id_example" // string | The ID of the required third-party security problem.
    remediationItemId := "remediationItemId_example" // string | The ID of the remediation item.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.GetRemediationItem(context.Background(), id, remediationItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.GetRemediationItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemediationItem`: RemediationDetailsItem
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.GetRemediationItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required third-party security problem. | 
**remediationItemId** | **string** | The ID of the remediation item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemediationItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemediationDetailsItem**](RemediationDetailsItem.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemediationItems

> RemediationItemList GetRemediationItems(ctx, id).RemediationItemSelector(remediationItemSelector).Execute()

Lists remediation items of a third-party security problem

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
    id := "id_example" // string | The ID of the required third-party security problem.
    remediationItemSelector := "remediationItemSelector_example" // string | Defines the scope of the query. Only remediable entities matching the specified criteria are included in the response.  You can add one or more of the following criteria. Values are *not* case-sensitive and the `EQUALS` operator is used unless otherwise specified.  * Vulnerability state: `vulnerabilityState(\"value\")`. Find the possible values in the description of the **vulnerabilityState** field of the response. If not set, all entities are returned. * Muted: `muted(\"value\")`. Possible values are `TRUE` or `FALSE`. * Reachable data asset assessment: `assessment.dataAssets(\"value\")` Possible values are `REACHABLE`, and `NOT_DETECTED`. * Network exposure assessment: `assessment.exposure(\"value\")` Possible values are `PUBLIC_NETWORK`, and `NOT_DETECTED`. * Vulnerable function usage assessment: `assessment.vulnerableFunctionUsage(\"value\")` Possible values are `IN_USE`, and `NOT_IN_USE`. * Vulnerable function in use contains: `assessment.vulnerableFunctionInUseContains(\"value\")`. Possible values are `class::function`, `class::` and `function`. The `CONTAINS` operator is used. Only vulnerable functions in use are considered. * Assessment accuracy: `assessment.accuracy(\"value\")` Possible values are `FULL` and `REDUCED`. * Entity name contains: `entityNameContains(\"value-1\")`. The `CONTAINS` operator is used.  To set several criteria, separate them with a comma (`,`). Only results matching **all** criteria are included in the response.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (`~`) inside quotes: * Tilde `~`  * Quote `\"` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.GetRemediationItems(context.Background(), id).RemediationItemSelector(remediationItemSelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.GetRemediationItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemediationItems`: RemediationItemList
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.GetRemediationItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required third-party security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemediationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remediationItemSelector** | **string** | Defines the scope of the query. Only remediable entities matching the specified criteria are included in the response.  You can add one or more of the following criteria. Values are *not* case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.  * Vulnerability state: &#x60;vulnerabilityState(\&quot;value\&quot;)&#x60;. Find the possible values in the description of the **vulnerabilityState** field of the response. If not set, all entities are returned. * Muted: &#x60;muted(\&quot;value\&quot;)&#x60;. Possible values are &#x60;TRUE&#x60; or &#x60;FALSE&#x60;. * Reachable data asset assessment: &#x60;assessment.dataAssets(\&quot;value\&quot;)&#x60; Possible values are &#x60;REACHABLE&#x60;, and &#x60;NOT_DETECTED&#x60;. * Network exposure assessment: &#x60;assessment.exposure(\&quot;value\&quot;)&#x60; Possible values are &#x60;PUBLIC_NETWORK&#x60;, and &#x60;NOT_DETECTED&#x60;. * Vulnerable function usage assessment: &#x60;assessment.vulnerableFunctionUsage(\&quot;value\&quot;)&#x60; Possible values are &#x60;IN_USE&#x60;, and &#x60;NOT_IN_USE&#x60;. * Vulnerable function in use contains: &#x60;assessment.vulnerableFunctionInUseContains(\&quot;value\&quot;)&#x60;. Possible values are &#x60;class::function&#x60;, &#x60;class::&#x60; and &#x60;function&#x60;. The &#x60;CONTAINS&#x60; operator is used. Only vulnerable functions in use are considered. * Assessment accuracy: &#x60;assessment.accuracy(\&quot;value\&quot;)&#x60; Possible values are &#x60;FULL&#x60; and &#x60;REDUCED&#x60;. * Entity name contains: &#x60;entityNameContains(\&quot;value-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used.  To set several criteria, separate them with a comma (&#x60;,&#x60;). Only results matching **all** criteria are included in the response.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (&#x60;~&#x60;) inside quotes: * Tilde &#x60;~&#x60;  * Quote &#x60;\&quot;&#x60; | 

### Return type

[**RemediationItemList**](RemediationItemList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemediationProgressEntities

> RemediationProgressEntityList GetRemediationProgressEntities(ctx, id, remediationItemId).RemediationProgressEntitySelector(remediationProgressEntitySelector).Execute()

Lists remediation progress entities

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
    id := "id_example" // string | The ID of the required third-party security problem.
    remediationItemId := "remediationItemId_example" // string | The ID of the remediation item.
    remediationProgressEntitySelector := "remediationProgressEntitySelector_example" // string | Defines the scope of the query. Only remediation progress entities matching the specified criteria are included in the response.  You can add one or more of the following criteria. Values are *not* case-sensitive and the `EQUALS` operator is used unless otherwise specified.  * State: `state(\"value\")`. Possible values the **state** field are `AFFECTED` and `UNAFFECTED`. If not set, all entities are returned. * Vulnerable function usage assessment: `assessment.vulnerableFunctionUsage(\"value\")` Possible values are `IN_USE`, and `NOT_IN_USE`. * Vulnerable function in use contains: `assessment.vulnerableFunctionInUseContains(\"value\")`. Possible values are `class::function`, `class::` and `function`. The `CONTAINS` operator is used. Only vulnerable functions in use are considered. * Entity name contains: `entityNameContains(\"value-1\")`. The `CONTAINS` operator is used.  To set several criteria, separate them with a comma (`,`). Only results matching **all** criteria are included in the response.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (`~`) inside quotes: * Tilde `~`  * Quote `\"` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.GetRemediationProgressEntities(context.Background(), id, remediationItemId).RemediationProgressEntitySelector(remediationProgressEntitySelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.GetRemediationProgressEntities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemediationProgressEntities`: RemediationProgressEntityList
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.GetRemediationProgressEntities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required third-party security problem. | 
**remediationItemId** | **string** | The ID of the remediation item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemediationProgressEntitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **remediationProgressEntitySelector** | **string** | Defines the scope of the query. Only remediation progress entities matching the specified criteria are included in the response.  You can add one or more of the following criteria. Values are *not* case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.  * State: &#x60;state(\&quot;value\&quot;)&#x60;. Possible values the **state** field are &#x60;AFFECTED&#x60; and &#x60;UNAFFECTED&#x60;. If not set, all entities are returned. * Vulnerable function usage assessment: &#x60;assessment.vulnerableFunctionUsage(\&quot;value\&quot;)&#x60; Possible values are &#x60;IN_USE&#x60;, and &#x60;NOT_IN_USE&#x60;. * Vulnerable function in use contains: &#x60;assessment.vulnerableFunctionInUseContains(\&quot;value\&quot;)&#x60;. Possible values are &#x60;class::function&#x60;, &#x60;class::&#x60; and &#x60;function&#x60;. The &#x60;CONTAINS&#x60; operator is used. Only vulnerable functions in use are considered. * Entity name contains: &#x60;entityNameContains(\&quot;value-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used.  To set several criteria, separate them with a comma (&#x60;,&#x60;). Only results matching **all** criteria are included in the response.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (&#x60;~&#x60;) inside quotes: * Tilde &#x60;~&#x60;  * Quote &#x60;\&quot;&#x60; | 

### Return type

[**RemediationProgressEntityList**](RemediationProgressEntityList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityProblem

> SecurityProblemDetails GetSecurityProblem(ctx, id).Fields(fields).From(from).Execute()

Gets parameters of a security problem

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
    id := "id_example" // string | The ID of the required security problem.
    fields := "fields_example" // string | A list of additional security problem properties you can add to the response.  The following properties are available (all other properties are always included and you can't remove them from the response):  * `riskAssessment`: A risk assessment of the security problem. * `managementZones`: The management zone where the security problem occurred. * `codeLevelVulnerabilityDetails`: Details of the code-level vulnerability. * `globalCounts`: Globally calculated statistics about the security problem. No management zone information is taken into account. * `filteredCounts`: Statistics about the security problem, filtered by the management zone and timeframe start ('from') query parameters.  * `description`: The description of the vulnerability. * `events`: The security problem's last 10 events within the last 365 days, sorted from newest to oldest. * `vulnerableComponents`: A list of vulnerable components of the security problem within the provided filter range.  * `affectedEntities`: A list of affected entities of the security problem within the provided filter range.  * `exposedEntities`: A list of exposed entities of the security problem within the provided filter range.  * `reachableDataAssets`: A list of data assets reachable by affected entities of the security problem within the provided filter range.  * `relatedEntities`: A list of related entities of the security problem within the provided filter range.  * `relatedContainerImages`: A list of related container images of the security problem within the provided filter range.  * `relatedAttacks`: A list of attacks detected on the exposed security problem.  * `entryPoints`: A list of entry points and a flag which indicates whether this list was truncated or not.  To add properties, specify them in a comma-separated list and prefix each property with a plus (for example, `+riskAssessment,+managementZones`).  (optional)
    from := "from_example" // string | Based on the timeframe start the affected-, related- and vulnerable entities are being calculated. You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the default timeframe start of 24 hours in the past is used (`now-24h`).  The timeframe start must not be older than 365 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.GetSecurityProblem(context.Background(), id).Fields(fields).From(from).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.GetSecurityProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityProblem`: SecurityProblemDetails
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.GetSecurityProblem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of additional security problem properties you can add to the response.  The following properties are available (all other properties are always included and you can&#39;t remove them from the response):  * &#x60;riskAssessment&#x60;: A risk assessment of the security problem. * &#x60;managementZones&#x60;: The management zone where the security problem occurred. * &#x60;codeLevelVulnerabilityDetails&#x60;: Details of the code-level vulnerability. * &#x60;globalCounts&#x60;: Globally calculated statistics about the security problem. No management zone information is taken into account. * &#x60;filteredCounts&#x60;: Statistics about the security problem, filtered by the management zone and timeframe start (&#39;from&#39;) query parameters.  * &#x60;description&#x60;: The description of the vulnerability. * &#x60;events&#x60;: The security problem&#39;s last 10 events within the last 365 days, sorted from newest to oldest. * &#x60;vulnerableComponents&#x60;: A list of vulnerable components of the security problem within the provided filter range.  * &#x60;affectedEntities&#x60;: A list of affected entities of the security problem within the provided filter range.  * &#x60;exposedEntities&#x60;: A list of exposed entities of the security problem within the provided filter range.  * &#x60;reachableDataAssets&#x60;: A list of data assets reachable by affected entities of the security problem within the provided filter range.  * &#x60;relatedEntities&#x60;: A list of related entities of the security problem within the provided filter range.  * &#x60;relatedContainerImages&#x60;: A list of related container images of the security problem within the provided filter range.  * &#x60;relatedAttacks&#x60;: A list of attacks detected on the exposed security problem.  * &#x60;entryPoints&#x60;: A list of entry points and a flag which indicates whether this list was truncated or not.  To add properties, specify them in a comma-separated list and prefix each property with a plus (for example, &#x60;+riskAssessment,+managementZones&#x60;).  | 
 **from** | **string** | Based on the timeframe start the affected-, related- and vulnerable entities are being calculated. You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the default timeframe start of 24 hours in the past is used (&#x60;now-24h&#x60;).  The timeframe start must not be older than 365 days. | 

### Return type

[**SecurityProblemDetails**](SecurityProblemDetails.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityProblems

> SecurityProblemList GetSecurityProblems(ctx).NextPageKey(nextPageKey).PageSize(pageSize).SecurityProblemSelector(securityProblemSelector).Sort(sort).Fields(fields).From(from).To(to).Execute()

Lists all security problems

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
    pageSize := int64(789) // int64 | The amount of security problems in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. (optional)
    securityProblemSelector := "securityProblemSelector_example" // string | Defines the scope of the query. Only security problems matching the specified criteria are included in the response.  You can add one or more of the following criteria. Values are *not* case-sensitive and the `EQUALS` operator is used unless otherwise specified.  * Status: `status(\"value\")`. Find the possible values in the description of the **status** field of the response. If not set, all security problems are returned. * Muted: `muted(\"value\")`. Possible values are `TRUE` or `FALSE`. * Risk level: `riskLevel(\"value\")`. The Davis risk level. Find the possible values in the description of the **riskLevel** field of the response. * Minimum risk score: `minRiskScore(\"5.5\")`. The Davis minimum risk score. The `GREATER THAN OR EQUAL TO` operator is used. Specify a number between `1.0` and `10.0`. * Maximum risk score: `maxRiskScore(\"5.5\")`. The Davis maximum risk score. The `LESS THAN` operator is used. Specify a number between `1.0` and `10.0`. * Base risk level: `baseRiskLevel(\"value\")`. The Base risk level from the CVSS. Find the possible values in the description of the **riskLevel** field of the response. * Minimum base risk score: `minBaseRiskScore(\"5.5\")`. The minimum base risk score from the CVSS. The `GREATER THAN OR EQUAL TO` operator is used. Specify a number between `1.0` and `10.0`. * Maximum base risk score: `maxBaseRiskScore(\"5.5\")`. The maximum base risk score from the CVSS. The `LESS THAN` operator is used. Specify a number between `1.0` and `10.0`. * External vulnerability ID contains: `externalVulnerabilityIdContains(\"id-1\")`. The `CONTAINS` operator is used. Maximum value length is 48 characters. * External vulnerability ID: `externalVulnerabilityId(\"id-1\", \"id-2\")`.  * CVE ID: `cveId(\"id\")`. * Risk assessment `riskAssessment(\"value-1\", \"value-2\")` Possible values are `EXPOSED`, `SENSITIVE`, `EXPLOIT`, `VULNERABLE_FUNCTION_IN_USE` and `ACCURACY_REDUCED`. * Related host ID: `relatedHostIds(\"value-1\", \"value-2\")`. Specify Dynatrace entity IDs here. * Related host name: `relatedHostNames(\"value-1\", \"value-2\")`. Values are case-sensitive. * Related host name contains: `relatedHostNameContains(\"value-1\")`. The `CONTAINS` operator is used. * Related Kubernetes cluster ID: `relatedKubernetesClusterIds(\"value-1\", \"value-2\")`. Specify Dynatrace entity IDs here. * Related Kubernetes cluster name: `relatedKubernetesClusterNames(\"value-1\", \"value-2\")`. Values are case-sensitive. * Related Kubernetes cluster name contains: `relatedKubernetesClusterNameContains(\"value-1\")`. The `CONTAINS` operator is used. * Related Kubernetes workload ID: `relatedKubernetesWorkloadIds(\"value-1\", \"value-2\")`. Specify Dynatrace entity IDs here. * Related Kubernetes workload name: `relatedKubernetesWorkloadNames(\"value-1\", \"value-2\")`. Values are case-sensitive. * Related Kubernetes workload name contains: `relatedKubernetesWorkloadNameContains(\"value-1\")`. The `CONTAINS` operator is used. * Management zone ID: `managementZoneIds(\"mzId-1\", \"mzId-2\")`. * Management zone name: `managementZones(\"name-1\", \"name-2\")`. Values are case-sensitive. * Affected process group instance ID: `affectedPgiIds(\"pgiId-1\", \"pgiId-2\")`. Specify Dynatrace entity IDs here. * Affected process group ID: `affectedPgIds(\"pgId-1\", \"pgId-2\")`. Specify Dynatrace entity IDs here. * Affected process group name: `affectedPgNames(\"name-1\", \"name-2\")`. Values are case-sensitive. * Affected process group name contains: `affectedPgNameContains(\"name-1\")`. The `CONTAINS` operator is used. * Vulnerable component ID: `vulnerableComponentIds(\"componentId-1\", \"componentId-2\")`. Specify component IDs here. * Vulnerable component name: `vulnerableComponentNames(\"name-1\", \"name-2\")`. Values are case-sensitive. * Vulnerable component name contains: `vulnerableComponentNameContains(\"name-1\")`. The `CONTAINS` operator is used. * Host tags: `hostTags(\"hostTag-1\")`. The `CONTAINS` operator is used. Maximum value length is 48 characters. * Process group tags: `pgTags(\"pgTag-1\")`. The `CONTAINS` operator is used. Maximum value length is 48 characters. * Process group instance tags: `pgiTags(\"pgiTag-1\")`. The `CONTAINS` operator is used. Maximum value length is 48 characters. * Tags: `tags(\"tag-1\")`. The `CONTAINS` operator is used. This selector picks hosts, process groups, and process group instances at the same time. Maximum value length is 48 characters. * Display ID: `displayIds(\"S-1234\", \"S-5678\")`. The `EQUALS` operator is used. * Security problem ID: `securityProblemIds(\"12544152654387159360\", \"5904857564184044850\")`. The `EQUALS` operator is used. * Technology: `technology(\"technology-1\", \"technology-2\")`. Find the possible values in the description of the **technology** field of the response. The `EQUALS` operator is used. * Vulnerability type: `vulnerabilityType(\"type-1\", \"type-2\")`. Possible values are `THIRD_PARTY`, `CODE_LEVEL`, `RUNTIME`.  Risk score and risk category are mutually exclusive (cannot be used at the same time).  To set several criteria, separate them with a comma (`,`). Only results matching **all** criteria are included in the response.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (`~`) inside quotes: * Tilde `~`  * Quote `\"` (optional)
    sort := "sort_example" // string | Specifies one or more fields for sorting the security problem list. Multiple fields can be concatenated using a comma (`,`) as a separator (e.g. `+status,-timestamp`).  You can sort by the following properties with a sign prefix for the sorting order.   * `status`: The security problem status (`+` open first or `-` resolved first)  * `muted`: The security problem mute state (`+` unmuted first or `-` muted first)  * `technology`: The security problem technology  * `firstSeenTimestamp`: The timestamp of the first occurrence of the security problem (`+` new problems first or `-` old problems first)  * `lastUpdatedTimestamp`: The timestamp of the last update of the security problem (`+` recently updated problems first or `-`  earlier updated problems first) * `securityProblemId`: The auto-generated ID of the security problem (`+` lower number first or `-` higher number first)  * `externalVulnerabilityId`: The ID of the external vulnerability (`+` lower number first or `-` higher number first)  * `displayId`: The display ID (`+` lower number first or `-` higher number first)  * `riskAssessment.riskScore`: The Davis security score (`+` lower score first or `-` higher score first)  * `riskAssessment.riskLevel`: The Davis security level (`+` lower level first or `-` higher level first)  * `riskAssessment.exposure`: Whether the problem is exposed to the internet  * `riskAssessment.dataAssets`: Whether data assets are affected  * `riskAssessment.vulnerableFunctionUsage`: Whether vulnerable functions are used  * `riskAssessment.assessmentAccuracy`: The assessments accuracy (`+` less accuracy first or `-` more accuracy first)  * `globalCounts.affectedNodes`: Number of affected nodes (`+` lower number first or `-` higher number first)  * `globalCounts.affectedProcessGroupInstances`: Number of affected process group instances (`+` lower number first or `-` higher number first)  * `globalCounts.affectedProcessGroups`: Number of affected process groups (`+` lower number first or `-` higher number first)  * `globalCounts.exposedProcessGroups`: Number of exposed process groups (`+` lower number first or `-` higher number first)  * `globalCounts.reachableDataAssets`: Number of reachable data assets (`+` lower number first or `-` higher number first)  * `globalCounts.relatedApplications`: Number of related applications (`+` lower number first or `-` higher number first)  * `globalCounts.relatedAttacks`: Number of attacks on the security problem (`+` lower number first or `-` higher number first)  * `globalCounts.relatedHosts`: Number of related hosts (`+` lower number first or `-` higher number first)  * `globalCounts.relatedKubernetesClusters`: Number of related Kubernetes cluster (`+` lower number first or `-` higher number first)  * `globalCounts.relatedKubernetesWorkloads`: Number of related Kubernetes workloads (`+` lower number first or `-` higher number first)  * `globalCounts.relatedServices`: Number of related services (`+` lower number first or `-` higher number first)  * `globalCounts.vulnerableComponents`: Number of vulnerable components (`+` lower number first or `-` higher number first)   If no prefix is set, `+` is used. (optional)
    fields := "fields_example" // string | A list of additional security problem properties you can add to the response.  The following properties are available (all other properties are always included and you can't remove them from the response):  * `riskAssessment`: A risk assessment of the security problem. * `managementZones`: The management zone where the security problem occurred. * `codeLevelVulnerabilityDetails`: Details of the code-level vulnerability. * `globalCounts`: Globally calculated statistics about the security problem. No management zone information is taken into account.  To add properties, specify them in a comma-separated list and prefix each property with a plus (for example, `+riskAssessment,+managementZones`).  (optional)
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of thirty days is used (`now-30d`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used.  The end of the timeframe must not be older than 365 days. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.GetSecurityProblems(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).SecurityProblemSelector(securityProblemSelector).Sort(sort).Fields(fields).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.GetSecurityProblems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSecurityProblems`: SecurityProblemList
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.GetSecurityProblems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityProblemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of security problems in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. | 
 **securityProblemSelector** | **string** | Defines the scope of the query. Only security problems matching the specified criteria are included in the response.  You can add one or more of the following criteria. Values are *not* case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.  * Status: &#x60;status(\&quot;value\&quot;)&#x60;. Find the possible values in the description of the **status** field of the response. If not set, all security problems are returned. * Muted: &#x60;muted(\&quot;value\&quot;)&#x60;. Possible values are &#x60;TRUE&#x60; or &#x60;FALSE&#x60;. * Risk level: &#x60;riskLevel(\&quot;value\&quot;)&#x60;. The Davis risk level. Find the possible values in the description of the **riskLevel** field of the response. * Minimum risk score: &#x60;minRiskScore(\&quot;5.5\&quot;)&#x60;. The Davis minimum risk score. The &#x60;GREATER THAN OR EQUAL TO&#x60; operator is used. Specify a number between &#x60;1.0&#x60; and &#x60;10.0&#x60;. * Maximum risk score: &#x60;maxRiskScore(\&quot;5.5\&quot;)&#x60;. The Davis maximum risk score. The &#x60;LESS THAN&#x60; operator is used. Specify a number between &#x60;1.0&#x60; and &#x60;10.0&#x60;. * Base risk level: &#x60;baseRiskLevel(\&quot;value\&quot;)&#x60;. The Base risk level from the CVSS. Find the possible values in the description of the **riskLevel** field of the response. * Minimum base risk score: &#x60;minBaseRiskScore(\&quot;5.5\&quot;)&#x60;. The minimum base risk score from the CVSS. The &#x60;GREATER THAN OR EQUAL TO&#x60; operator is used. Specify a number between &#x60;1.0&#x60; and &#x60;10.0&#x60;. * Maximum base risk score: &#x60;maxBaseRiskScore(\&quot;5.5\&quot;)&#x60;. The maximum base risk score from the CVSS. The &#x60;LESS THAN&#x60; operator is used. Specify a number between &#x60;1.0&#x60; and &#x60;10.0&#x60;. * External vulnerability ID contains: &#x60;externalVulnerabilityIdContains(\&quot;id-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. Maximum value length is 48 characters. * External vulnerability ID: &#x60;externalVulnerabilityId(\&quot;id-1\&quot;, \&quot;id-2\&quot;)&#x60;.  * CVE ID: &#x60;cveId(\&quot;id\&quot;)&#x60;. * Risk assessment &#x60;riskAssessment(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60; Possible values are &#x60;EXPOSED&#x60;, &#x60;SENSITIVE&#x60;, &#x60;EXPLOIT&#x60;, &#x60;VULNERABLE_FUNCTION_IN_USE&#x60; and &#x60;ACCURACY_REDUCED&#x60;. * Related host ID: &#x60;relatedHostIds(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Specify Dynatrace entity IDs here. * Related host name: &#x60;relatedHostNames(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Values are case-sensitive. * Related host name contains: &#x60;relatedHostNameContains(\&quot;value-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Related Kubernetes cluster ID: &#x60;relatedKubernetesClusterIds(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Specify Dynatrace entity IDs here. * Related Kubernetes cluster name: &#x60;relatedKubernetesClusterNames(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Values are case-sensitive. * Related Kubernetes cluster name contains: &#x60;relatedKubernetesClusterNameContains(\&quot;value-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Related Kubernetes workload ID: &#x60;relatedKubernetesWorkloadIds(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Specify Dynatrace entity IDs here. * Related Kubernetes workload name: &#x60;relatedKubernetesWorkloadNames(\&quot;value-1\&quot;, \&quot;value-2\&quot;)&#x60;. Values are case-sensitive. * Related Kubernetes workload name contains: &#x60;relatedKubernetesWorkloadNameContains(\&quot;value-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Management zone ID: &#x60;managementZoneIds(\&quot;mzId-1\&quot;, \&quot;mzId-2\&quot;)&#x60;. * Management zone name: &#x60;managementZones(\&quot;name-1\&quot;, \&quot;name-2\&quot;)&#x60;. Values are case-sensitive. * Affected process group instance ID: &#x60;affectedPgiIds(\&quot;pgiId-1\&quot;, \&quot;pgiId-2\&quot;)&#x60;. Specify Dynatrace entity IDs here. * Affected process group ID: &#x60;affectedPgIds(\&quot;pgId-1\&quot;, \&quot;pgId-2\&quot;)&#x60;. Specify Dynatrace entity IDs here. * Affected process group name: &#x60;affectedPgNames(\&quot;name-1\&quot;, \&quot;name-2\&quot;)&#x60;. Values are case-sensitive. * Affected process group name contains: &#x60;affectedPgNameContains(\&quot;name-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Vulnerable component ID: &#x60;vulnerableComponentIds(\&quot;componentId-1\&quot;, \&quot;componentId-2\&quot;)&#x60;. Specify component IDs here. * Vulnerable component name: &#x60;vulnerableComponentNames(\&quot;name-1\&quot;, \&quot;name-2\&quot;)&#x60;. Values are case-sensitive. * Vulnerable component name contains: &#x60;vulnerableComponentNameContains(\&quot;name-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. * Host tags: &#x60;hostTags(\&quot;hostTag-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. Maximum value length is 48 characters. * Process group tags: &#x60;pgTags(\&quot;pgTag-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. Maximum value length is 48 characters. * Process group instance tags: &#x60;pgiTags(\&quot;pgiTag-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. Maximum value length is 48 characters. * Tags: &#x60;tags(\&quot;tag-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used. This selector picks hosts, process groups, and process group instances at the same time. Maximum value length is 48 characters. * Display ID: &#x60;displayIds(\&quot;S-1234\&quot;, \&quot;S-5678\&quot;)&#x60;. The &#x60;EQUALS&#x60; operator is used. * Security problem ID: &#x60;securityProblemIds(\&quot;12544152654387159360\&quot;, \&quot;5904857564184044850\&quot;)&#x60;. The &#x60;EQUALS&#x60; operator is used. * Technology: &#x60;technology(\&quot;technology-1\&quot;, \&quot;technology-2\&quot;)&#x60;. Find the possible values in the description of the **technology** field of the response. The &#x60;EQUALS&#x60; operator is used. * Vulnerability type: &#x60;vulnerabilityType(\&quot;type-1\&quot;, \&quot;type-2\&quot;)&#x60;. Possible values are &#x60;THIRD_PARTY&#x60;, &#x60;CODE_LEVEL&#x60;, &#x60;RUNTIME&#x60;.  Risk score and risk category are mutually exclusive (cannot be used at the same time).  To set several criteria, separate them with a comma (&#x60;,&#x60;). Only results matching **all** criteria are included in the response.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (&#x60;~&#x60;) inside quotes: * Tilde &#x60;~&#x60;  * Quote &#x60;\&quot;&#x60; | 
 **sort** | **string** | Specifies one or more fields for sorting the security problem list. Multiple fields can be concatenated using a comma (&#x60;,&#x60;) as a separator (e.g. &#x60;+status,-timestamp&#x60;).  You can sort by the following properties with a sign prefix for the sorting order.   * &#x60;status&#x60;: The security problem status (&#x60;+&#x60; open first or &#x60;-&#x60; resolved first)  * &#x60;muted&#x60;: The security problem mute state (&#x60;+&#x60; unmuted first or &#x60;-&#x60; muted first)  * &#x60;technology&#x60;: The security problem technology  * &#x60;firstSeenTimestamp&#x60;: The timestamp of the first occurrence of the security problem (&#x60;+&#x60; new problems first or &#x60;-&#x60; old problems first)  * &#x60;lastUpdatedTimestamp&#x60;: The timestamp of the last update of the security problem (&#x60;+&#x60; recently updated problems first or &#x60;-&#x60;  earlier updated problems first) * &#x60;securityProblemId&#x60;: The auto-generated ID of the security problem (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;externalVulnerabilityId&#x60;: The ID of the external vulnerability (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;displayId&#x60;: The display ID (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;riskAssessment.riskScore&#x60;: The Davis security score (&#x60;+&#x60; lower score first or &#x60;-&#x60; higher score first)  * &#x60;riskAssessment.riskLevel&#x60;: The Davis security level (&#x60;+&#x60; lower level first or &#x60;-&#x60; higher level first)  * &#x60;riskAssessment.exposure&#x60;: Whether the problem is exposed to the internet  * &#x60;riskAssessment.dataAssets&#x60;: Whether data assets are affected  * &#x60;riskAssessment.vulnerableFunctionUsage&#x60;: Whether vulnerable functions are used  * &#x60;riskAssessment.assessmentAccuracy&#x60;: The assessments accuracy (&#x60;+&#x60; less accuracy first or &#x60;-&#x60; more accuracy first)  * &#x60;globalCounts.affectedNodes&#x60;: Number of affected nodes (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.affectedProcessGroupInstances&#x60;: Number of affected process group instances (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.affectedProcessGroups&#x60;: Number of affected process groups (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.exposedProcessGroups&#x60;: Number of exposed process groups (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.reachableDataAssets&#x60;: Number of reachable data assets (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.relatedApplications&#x60;: Number of related applications (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.relatedAttacks&#x60;: Number of attacks on the security problem (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.relatedHosts&#x60;: Number of related hosts (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.relatedKubernetesClusters&#x60;: Number of related Kubernetes cluster (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.relatedKubernetesWorkloads&#x60;: Number of related Kubernetes workloads (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.relatedServices&#x60;: Number of related services (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)  * &#x60;globalCounts.vulnerableComponents&#x60;: Number of vulnerable components (&#x60;+&#x60; lower number first or &#x60;-&#x60; higher number first)   If no prefix is set, &#x60;+&#x60; is used. | 
 **fields** | **string** | A list of additional security problem properties you can add to the response.  The following properties are available (all other properties are always included and you can&#39;t remove them from the response):  * &#x60;riskAssessment&#x60;: A risk assessment of the security problem. * &#x60;managementZones&#x60;: The management zone where the security problem occurred. * &#x60;codeLevelVulnerabilityDetails&#x60;: Details of the code-level vulnerability. * &#x60;globalCounts&#x60;: Globally calculated statistics about the security problem. No management zone information is taken into account.  To add properties, specify them in a comma-separated list and prefix each property with a plus (for example, &#x60;+riskAssessment,+managementZones&#x60;).  | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of thirty days is used (&#x60;now-30d&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used.  The end of the timeframe must not be older than 365 days. | 

### Return type

[**SecurityProblemList**](SecurityProblemList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVulnerableFunctions

> VulnerableFunctionsContainer GetVulnerableFunctions(ctx, id).VulnerableFunctionsSelector(vulnerableFunctionsSelector).GroupBy(groupBy).Execute()

Lists all vulnerable functions and their usage for a third-party security problem

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
    id := "id_example" // string | The ID of the required third-party security problem.
    vulnerableFunctionsSelector := "vulnerableFunctionsSelector_example" // string | Defines the scope of the query. Only vulnerable functions matching the specified criteria are included in the response.  You can add the following criteria. Values are *not* case sensitive and the `EQUALS` operator is used unless otherwise specified.  * Management zone ID: `managementZoneIds(\"mzId-1\", \"mzId-2\")`. * Management zone name: `managementZones(\"name-1\", \"name-2\")`. Values are case sensitive. * Process group ID: `processGroupIds(\"pgId-1\", \"pgId-2\")`. Specify Dynatrace entity IDs here. * Process group name: `processGroupNames(\"name-1\", \"name-2\")`. Values are case sensitive. * Process group name contains: `processGroupNameContains(\"name-1\")`. The `CONTAINS` operator is used.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (`~`) inside quotes: * Tilde `~`  * Quote `\"` (optional)
    groupBy := "groupBy_example" // string | Defines additional grouping types in which vulnerable functions should be displayed.  You can add one of the following grouping types.  * Process group: `PROCESS_GROUP` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityProblemsAPI.GetVulnerableFunctions(context.Background(), id).VulnerableFunctionsSelector(vulnerableFunctionsSelector).GroupBy(groupBy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.GetVulnerableFunctions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVulnerableFunctions`: VulnerableFunctionsContainer
    fmt.Fprintf(os.Stdout, "Response from `SecurityProblemsAPI.GetVulnerableFunctions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required third-party security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVulnerableFunctionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **vulnerableFunctionsSelector** | **string** | Defines the scope of the query. Only vulnerable functions matching the specified criteria are included in the response.  You can add the following criteria. Values are *not* case sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.  * Management zone ID: &#x60;managementZoneIds(\&quot;mzId-1\&quot;, \&quot;mzId-2\&quot;)&#x60;. * Management zone name: &#x60;managementZones(\&quot;name-1\&quot;, \&quot;name-2\&quot;)&#x60;. Values are case sensitive. * Process group ID: &#x60;processGroupIds(\&quot;pgId-1\&quot;, \&quot;pgId-2\&quot;)&#x60;. Specify Dynatrace entity IDs here. * Process group name: &#x60;processGroupNames(\&quot;name-1\&quot;, \&quot;name-2\&quot;)&#x60;. Values are case sensitive. * Process group name contains: &#x60;processGroupNameContains(\&quot;name-1\&quot;)&#x60;. The &#x60;CONTAINS&#x60; operator is used.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (&#x60;~&#x60;) inside quotes: * Tilde &#x60;~&#x60;  * Quote &#x60;\&quot;&#x60; | 
 **groupBy** | **string** | Defines additional grouping types in which vulnerable functions should be displayed.  You can add one of the following grouping types.  * Process group: &#x60;PROCESS_GROUP&#x60; | 

### Return type

[**VulnerableFunctionsContainer**](VulnerableFunctionsContainer.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteSecurityProblem

> MuteSecurityProblem(ctx, id).SecurityProblemMute(securityProblemMute).Execute()

Mutes a security problem

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
    id := "id_example" // string | The ID of the required security problem.
    securityProblemMute := *openapiclient.NewSecurityProblemMute("Reason_example") // SecurityProblemMute | The JSON body of the request. Contains the muting information. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecurityProblemsAPI.MuteSecurityProblem(context.Background(), id).SecurityProblemMute(securityProblemMute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.MuteSecurityProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMuteSecurityProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityProblemMute** | [**SecurityProblemMute**](SecurityProblemMute.md) | The JSON body of the request. Contains the muting information. | 

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


## SetRemediationItemMuteState

> SetRemediationItemMuteState(ctx, id, remediationItemId).RemediationItemMuteStateChange(remediationItemMuteStateChange).Execute()

Sets the mute state of a remediation item

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
    id := "id_example" // string | The ID of the required third-party security problem.
    remediationItemId := "remediationItemId_example" // string | The ID of the remediation item.
    remediationItemMuteStateChange := *openapiclient.NewRemediationItemMuteStateChange("Comment_example", false, "IGNORE") // RemediationItemMuteStateChange | The JSON body of the request. Contains the mute state information to be applied. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecurityProblemsAPI.SetRemediationItemMuteState(context.Background(), id, remediationItemId).RemediationItemMuteStateChange(remediationItemMuteStateChange).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.SetRemediationItemMuteState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required third-party security problem. | 
**remediationItemId** | **string** | The ID of the remediation item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRemediationItemMuteStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **remediationItemMuteStateChange** | [**RemediationItemMuteStateChange**](RemediationItemMuteStateChange.md) | The JSON body of the request. Contains the mute state information to be applied. | 

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


## UnmuteSecurityProblem

> UnmuteSecurityProblem(ctx, id).SecurityProblemUnmute(securityProblemUnmute).Execute()

Un-mutes a security problem

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
    id := "id_example" // string | The ID of the required security problem.
    securityProblemUnmute := *openapiclient.NewSecurityProblemUnmute("Reason_example") // SecurityProblemUnmute | The JSON body of the request. Contains the un-muting information. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SecurityProblemsAPI.UnmuteSecurityProblem(context.Background(), id).SecurityProblemUnmute(securityProblemUnmute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityProblemsAPI.UnmuteSecurityProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required security problem. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteSecurityProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityProblemUnmute** | [**SecurityProblemUnmute**](SecurityProblemUnmute.md) | The JSON body of the request. Contains the un-muting information. | 

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

