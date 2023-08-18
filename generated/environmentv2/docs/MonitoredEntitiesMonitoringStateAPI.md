# \MonitoredEntitiesMonitoringStateAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStates**](MonitoredEntitiesMonitoringStateAPI.md#GetStates) | **Get** /monitoringstate | Lists monitoring states of entities



## GetStates

> MonitoredStates GetStates(ctx).EntitySelector(entitySelector).Execute()

Lists monitoring states of entities



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
    entitySelector := "entitySelector_example" // string | Specifies the process group instances where you're querying the state. Use the `PROCESS_GROUP_INSTANCE` entity type.   You must set one of these criteria:   * Entity type: `type(\"TYPE\")`  * Dynatrace entity ID: `entityId(\"id\")`. You can specify several IDs, separated by a comma (`entityId(\"id-1\",\"id-2\")`). All requested entities must be of the same type.   You can add one or more of the following criteria. Values are case-sensitive and the `EQUALS` operator is used unless otherwise specified.   * Tag: `tag(\"value\")`. Tags in `[context]key:value`, `key:value`, and `value` formats are detected and parsed automatically. Any colons (`:`) that are part of the key or value must be escaped with a backslash(`\\`). Otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: `mzId(123)`  * Management zone name: `mzName(\"value\")` * Entity name:   * `entityName.equals`: performs a non-casesensitive `EQUALS` query.   * `entityName.startsWith`: changes the operator to `BEGINS WITH`.   * `entityName.in`: enables you to provide multiple values. The `EQUALS` operator applies.   * `caseSensitive(entityName.equals(\"value\"))`: takes any entity name criterion as an argument and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): `healthState(\"HEALTHY\")` * First seen timestamp: `firstSeenTms.<operator>(now-3h)`. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * `lte`: earlier than or at the specified time  * `lt`: earlier than the specified time  * `gte`: later than or at the specified time  * `gt`: later than the specified time * Entity attribute: `<attribute>(\"value1\",\"value2\")` and `<attribute>.exists()`. To fetch the list of available attributes, execute the [GET entity type](https://dt-url.net/2ka3ivt) request and check the **properties** field of the response.  * Relationships: `fromRelationships.<relationshipName>()` and `toRelationships.<relationshipName>()`. This criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET entity type](https://dt-url.net/2ka3ivt) request and check the **fromRelationships** and **toRelationships** fields. * Negation: `not(<criterion>)`. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (`,`). For example, `type(\"HOST\"),healthState(\"HEALTHY\")`. Only results matching **all** criteria are included in the response.   The maximum string length is 2,000 characters.   

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MonitoredEntitiesMonitoringStateAPI.GetStates(context.Background()).EntitySelector(entitySelector).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitoredEntitiesMonitoringStateAPI.GetStates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStates`: MonitoredStates
    fmt.Fprintf(os.Stdout, "Response from `MonitoredEntitiesMonitoringStateAPI.GetStates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitySelector** | **string** | Specifies the process group instances where you&#39;re querying the state. Use the &#x60;PROCESS_GROUP_INSTANCE&#x60; entity type.   You must set one of these criteria:   * Entity type: &#x60;type(\&quot;TYPE\&quot;)&#x60;  * Dynatrace entity ID: &#x60;entityId(\&quot;id\&quot;)&#x60;. You can specify several IDs, separated by a comma (&#x60;entityId(\&quot;id-1\&quot;,\&quot;id-2\&quot;)&#x60;). All requested entities must be of the same type.   You can add one or more of the following criteria. Values are case-sensitive and the &#x60;EQUALS&#x60; operator is used unless otherwise specified.   * Tag: &#x60;tag(\&quot;value\&quot;)&#x60;. Tags in &#x60;[context]key:value&#x60;, &#x60;key:value&#x60;, and &#x60;value&#x60; formats are detected and parsed automatically. Any colons (&#x60;:&#x60;) that are part of the key or value must be escaped with a backslash(&#x60;\\&#x60;). Otherwise, it will be interpreted as the separator between the key and the value. All tag values are case-sensitive.  * Management zone ID: &#x60;mzId(123)&#x60;  * Management zone name: &#x60;mzName(\&quot;value\&quot;)&#x60; * Entity name:   * &#x60;entityName.equals&#x60;: performs a non-casesensitive &#x60;EQUALS&#x60; query.   * &#x60;entityName.startsWith&#x60;: changes the operator to &#x60;BEGINS WITH&#x60;.   * &#x60;entityName.in&#x60;: enables you to provide multiple values. The &#x60;EQUALS&#x60; operator applies.   * &#x60;caseSensitive(entityName.equals(\&quot;value\&quot;))&#x60;: takes any entity name criterion as an argument and makes the value case-sensitive. * Health state (HEALTHY,UNHEALTHY): &#x60;healthState(\&quot;HEALTHY\&quot;)&#x60; * First seen timestamp: &#x60;firstSeenTms.&lt;operator&gt;(now-3h)&#x60;. Use any timestamp format from the **from**_/_**to** parameters.   The following operators are available:  * &#x60;lte&#x60;: earlier than or at the specified time  * &#x60;lt&#x60;: earlier than the specified time  * &#x60;gte&#x60;: later than or at the specified time  * &#x60;gt&#x60;: later than the specified time * Entity attribute: &#x60;&lt;attribute&gt;(\&quot;value1\&quot;,\&quot;value2\&quot;)&#x60; and &#x60;&lt;attribute&gt;.exists()&#x60;. To fetch the list of available attributes, execute the [GET entity type](https://dt-url.net/2ka3ivt) request and check the **properties** field of the response.  * Relationships: &#x60;fromRelationships.&lt;relationshipName&gt;()&#x60; and &#x60;toRelationships.&lt;relationshipName&gt;()&#x60;. This criterion takes an entity selector as an attribute. To fetch the list of available relationships, execute the [GET entity type](https://dt-url.net/2ka3ivt) request and check the **fromRelationships** and **toRelationships** fields. * Negation: &#x60;not(&lt;criterion&gt;)&#x60;. Inverts any criterion except for **type**.   For more information, see [Entity selector](https://dt-url.net/apientityselector) in Dynatrace Documentation.   To set several criteria, separate them with a comma (&#x60;,&#x60;). For example, &#x60;type(\&quot;HOST\&quot;),healthState(\&quot;HEALTHY\&quot;)&#x60;. Only results matching **all** criteria are included in the response.   The maximum string length is 2,000 characters.    | 

### Return type

[**MonitoredStates**](MonitoredStates.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

