# \AttacksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAttack**](AttacksAPI.md#GetAttack) | **Get** /attacks/{id} | Gets the specified attack | maturity&#x3D;EARLY_ADOPTER
[**GetAttacks**](AttacksAPI.md#GetAttacks) | **Get** /attacks | Lists all attacks | maturity&#x3D;EARLY_ADOPTER



## GetAttack

> Attack GetAttack(ctx, id).Fields(fields).Execute()

Gets the specified attack | maturity=EARLY_ADOPTER

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
    id := "id_example" // string | The ID of the attack.
    fields := "fields_example" // string | A list of additional attack properties you can add to the response.  The following properties are available (all other properties are always included and you can't remove them from the response):  * `attackTarget`: The targeted host/database of an attack. * `request`: The request that was sent from the attacker. * `entrypoint`: The entry point used by an attacker to start a specific attack. * `vulnerability`: The vulnerability utilized by the attack. * `securityProblem`: The related security problem. * `attacker`: The attacker of an attack. * `managementZones`: The related management zones.  To add properties, specify them in a comma-separated list and prefix each property with a plus (for example, `+attackTarget,+securityProblem`).  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttacksAPI.GetAttack(context.Background(), id).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttacksAPI.GetAttack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttack`: Attack
    fmt.Fprintf(os.Stdout, "Response from `AttacksAPI.GetAttack`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attack. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | A list of additional attack properties you can add to the response.  The following properties are available (all other properties are always included and you can&#39;t remove them from the response):  * &#x60;attackTarget&#x60;: The targeted host/database of an attack. * &#x60;request&#x60;: The request that was sent from the attacker. * &#x60;entrypoint&#x60;: The entry point used by an attacker to start a specific attack. * &#x60;vulnerability&#x60;: The vulnerability utilized by the attack. * &#x60;securityProblem&#x60;: The related security problem. * &#x60;attacker&#x60;: The attacker of an attack. * &#x60;managementZones&#x60;: The related management zones.  To add properties, specify them in a comma-separated list and prefix each property with a plus (for example, &#x60;+attackTarget,+securityProblem&#x60;).  | 

### Return type

[**Attack**](Attack.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttacks

> AttackList GetAttacks(ctx).NextPageKey(nextPageKey).PageSize(pageSize).AttackSelector(attackSelector).Sort(sort).Fields(fields).From(from).To(to).Execute()

Lists all attacks | maturity=EARLY_ADOPTER

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
    pageSize := int64(789) // int64 | The amount of attacks in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. (optional)
    attackSelector := "attackSelector_example" // string | Defines the scope of the query. Only attacks matching the specified criteria are included in the response. You can add one or more of the following criteria. Values are *not* case-sensitive and the `EQUALS` operator is used unless otherwise specified.  * State: `state(\"value\")`. The state of the attack. Possible values are `EXPLOITED`, `BLOCKED`, and `ALLOWLISTED`. * Attack Type: `attackType(\"value\")`. The type of the attack. Find the possible values in the description of the **attackType** field of the response. * Country Code: `countryCode(\"value\")`. The country code of the attacker. Supported values include all ISO-3166-1 alpha-2 country codes (2-letter). Supplying empty filter value `countryCode()` will return attacks, where location is not available. * Request path contains: `requestPathContains(\"value\")`. Filters for a substring in the request path. The `CONTAINS` operator is used. A maximum of 48 characters are allowed. * Process group name contains: `processGroupNameContains(\"value\")`. Filters for a substring in the targeted process group's name. The `CONTAINS` operator is used. * Vulnerability ID: `vulnerabilityId(\"123456789\")`. The exact ID of the vulnerability. * Source IPs: `sourceIps(\"93.184.216.34\", \"63.124.6.12\")`. The exact IPv4/IPv6 addresses of the attacker. * Management zone ID: `managementZoneIds(\"mzId-1\", \"mzId-2\")`. * Management zone name: `managementZones(\"name-1\", \"name-2\")`. Values are case sensitive. * Technology: `technology(\"technology-1\", \"technology-2\")`. Find the possible values in the description of the **technology** field of the response. The `EQUALS` operator is used.  To set several criteria, separate them with a comma (`,`). Only results matching (**all** criteria are included in the response.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (`~`) inside quotes: * Tilde `~`  * Quote `\"` (optional)
    sort := "sort_example" // string | Specifies one or more fields for sorting the attack list. Multiple fields can be concatenated using a comma (`,`) as a separator (e.g. `+state,-timestamp`).  You can sort by the following properties with a sign prefix for the sorting order.   * `displayId`: The attack's display ID. * `displayName`: The attack's display name. * `attackType`: The type of the attack (e.g. SQL_INJECTION, JNDI_INJECTION, etc.). * `state`: The state of the attack. (`+` low severity state first `-` high severity state first) * `sourceIp`: The IP address of the attacker. Sorts by the numerical IP value. * `requestPath`: The request path where the attack was started. * `timestamp`: When the attack was executed. (`+` old attacks first or `-` new attacks first) If no prefix is set, `+` is used. (optional)
    fields := "fields_example" // string | A list of additional attack properties you can add to the response.  The following properties are available (all other properties are always included and you can't remove them from the response):  * `attackTarget`: The targeted host/database of an attack. * `request`: The request that was sent from the attacker. * `entrypoint`: The entry point used by an attacker to start a specific attack. * `vulnerability`: The vulnerability utilized by the attack. * `securityProblem`: The related security problem. * `attacker`: The attacker of an attack. * `managementZones`: The related management zones. * `affectedEntities`: The affected entities of an attack.  To add properties, specify them in a comma-separated list and prefix each property with a plus (for example, `+attackTarget,+securityProblem`).  (optional)
    from := "from_example" // string | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the relative timeframe of thirty days is used (`now-30d`). (optional)
    to := "to_example" // string | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of `2021-01-25T05:57:01.123+01:00`. If no time zone is specified, UTC is used. You can use a space character instead of the `T`. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is `now-NU/A`, where `N` is the amount of time, `U` is the unit of time, and `A` is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, `now-1y/w` is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: `now-NU`.  Supported time units for the relative timeframe are:     * `m`: minutes     * `h`: hours     * `d`: days     * `w`: weeks     * `M`: months     * `y`: years   If not set, the current timestamp is used. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AttacksAPI.GetAttacks(context.Background()).NextPageKey(nextPageKey).PageSize(pageSize).AttackSelector(attackSelector).Sort(sort).Fields(fields).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AttacksAPI.GetAttacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAttacks`: AttackList
    fmt.Fprintf(os.Stdout, "Response from `AttacksAPI.GetAttacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAttacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nextPageKey** | **string** | The cursor for the next page of results. You can find it in the **nextPageKey** field of the previous response.   The first page is always returned if you don&#39;t specify the **nextPageKey** query parameter.   When the **nextPageKey** is set to obtain subsequent pages, you must omit all other query parameters.  | 
 **pageSize** | **int64** | The amount of attacks in a single response payload.   The maximal allowed page size is 500.   If not set, 100 is used. | 
 **attackSelector** | **string** | Defines the scope of the query. Only attacks matching the specified criteria are included in the response. You can add one or more of the following criteria. Values are *not* case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.  * State: &#x60;state(\&quot;value\&quot;)&#x60;. The state of the attack. Possible values are &#x60;EXPLOITED&#x60;, &#x60;BLOCKED&#x60;, and &#x60;ALLOWLISTED&#x60;. * Attack Type: &#x60;attackType(\&quot;value\&quot;)&#x60;. The type of the attack. Find the possible values in the description of the **attackType** field of the response. * Country Code: &#x60;countryCode(\&quot;value\&quot;)&#x60;. The country code of the attacker. Supported values include all ISO-3166-1 alpha-2 country codes (2-letter). Supplying empty filter value &#x60;countryCode()&#x60; will return attacks, where location is not available. * Request path contains: &#x60;requestPathContains(\&quot;value\&quot;)&#x60;. Filters for a substring in the request path. The &#x60;CONTAINS&#x60; operator is used. A maximum of 48 characters are allowed. * Process group name contains: &#x60;processGroupNameContains(\&quot;value\&quot;)&#x60;. Filters for a substring in the targeted process group&#39;s name. The &#x60;CONTAINS&#x60; operator is used. * Vulnerability ID: &#x60;vulnerabilityId(\&quot;123456789\&quot;)&#x60;. The exact ID of the vulnerability. * Source IPs: &#x60;sourceIps(\&quot;93.184.216.34\&quot;, \&quot;63.124.6.12\&quot;)&#x60;. The exact IPv4/IPv6 addresses of the attacker. * Management zone ID: &#x60;managementZoneIds(\&quot;mzId-1\&quot;, \&quot;mzId-2\&quot;)&#x60;. * Management zone name: &#x60;managementZones(\&quot;name-1\&quot;, \&quot;name-2\&quot;)&#x60;. Values are case sensitive. * Technology: &#x60;technology(\&quot;technology-1\&quot;, \&quot;technology-2\&quot;)&#x60;. Find the possible values in the description of the **technology** field of the response. The &#x60;EQUALS&#x60; operator is used.  To set several criteria, separate them with a comma (&#x60;,&#x60;). Only results matching (**all** criteria are included in the response.  Specify the value of a criterion as a quoted string. The following special characters must be escaped with a tilde (&#x60;~&#x60;) inside quotes: * Tilde &#x60;~&#x60;  * Quote &#x60;\&quot;&#x60; | 
 **sort** | **string** | Specifies one or more fields for sorting the attack list. Multiple fields can be concatenated using a comma (&#x60;,&#x60;) as a separator (e.g. &#x60;+state,-timestamp&#x60;).  You can sort by the following properties with a sign prefix for the sorting order.   * &#x60;displayId&#x60;: The attack&#39;s display ID. * &#x60;displayName&#x60;: The attack&#39;s display name. * &#x60;attackType&#x60;: The type of the attack (e.g. SQL_INJECTION, JNDI_INJECTION, etc.). * &#x60;state&#x60;: The state of the attack. (&#x60;+&#x60; low severity state first &#x60;-&#x60; high severity state first) * &#x60;sourceIp&#x60;: The IP address of the attacker. Sorts by the numerical IP value. * &#x60;requestPath&#x60;: The request path where the attack was started. * &#x60;timestamp&#x60;: When the attack was executed. (&#x60;+&#x60; old attacks first or &#x60;-&#x60; new attacks first) If no prefix is set, &#x60;+&#x60; is used. | 
 **fields** | **string** | A list of additional attack properties you can add to the response.  The following properties are available (all other properties are always included and you can&#39;t remove them from the response):  * &#x60;attackTarget&#x60;: The targeted host/database of an attack. * &#x60;request&#x60;: The request that was sent from the attacker. * &#x60;entrypoint&#x60;: The entry point used by an attacker to start a specific attack. * &#x60;vulnerability&#x60;: The vulnerability utilized by the attack. * &#x60;securityProblem&#x60;: The related security problem. * &#x60;attacker&#x60;: The attacker of an attack. * &#x60;managementZones&#x60;: The related management zones. * &#x60;affectedEntities&#x60;: The affected entities of an attack.  To add properties, specify them in a comma-separated list and prefix each property with a plus (for example, &#x60;+attackTarget,+securityProblem&#x60;).  | 
 **from** | **string** | The start of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the relative timeframe of thirty days is used (&#x60;now-30d&#x60;). | 
 **to** | **string** | The end of the requested timeframe.   You can use one of the following formats:  * Timestamp in UTC milliseconds.  * Human-readable format of &#x60;2021-01-25T05:57:01.123+01:00&#x60;. If no time zone is specified, UTC is used. You can use a space character instead of the &#x60;T&#x60;. Seconds and fractions of a second are optional.  * Relative timeframe, back from now. The format is &#x60;now-NU/A&#x60;, where &#x60;N&#x60; is the amount of time, &#x60;U&#x60; is the unit of time, and &#x60;A&#x60; is an alignment. The alignment rounds all the smaller values to the nearest zero in the past. For example, &#x60;now-1y/w&#x60; is one year back, aligned by a week.  You can also specify relative timeframe without an alignment: &#x60;now-NU&#x60;.  Supported time units for the relative timeframe are:     * &#x60;m&#x60;: minutes     * &#x60;h&#x60;: hours     * &#x60;d&#x60;: days     * &#x60;w&#x60;: weeks     * &#x60;M&#x60;: months     * &#x60;y&#x60;: years   If not set, the current timestamp is used. | 

### Return type

[**AttackList**](AttackList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

