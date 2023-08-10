# \DynatracePlatformSubscriptionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionsControllerGetEnvironmentCost**](DynatracePlatformSubscriptionAPI.md#SubscriptionsControllerGetEnvironmentCost) | **Get** /sub/v2/accounts/{accountUuid}/subscriptions/{subscriptionUuid}/environments/cost | Get cost summary grouped by environment for a given subscription
[**SubscriptionsControllerGetEnvironmentUsage**](DynatracePlatformSubscriptionAPI.md#SubscriptionsControllerGetEnvironmentUsage) | **Get** /sub/v2/accounts/{accountUuid}/subscriptions/{subscriptionUuid}/environments/usage | Get usage summary grouped by environment for a given subscription
[**SubscriptionsControllerGetEvents**](DynatracePlatformSubscriptionAPI.md#SubscriptionsControllerGetEvents) | **Get** /sub/v2/accounts/{accountUuid}/subscriptions/events | Get all notification events for a given subscription
[**SubscriptionsControllerGetForecast**](DynatracePlatformSubscriptionAPI.md#SubscriptionsControllerGetForecast) | **Get** /sub/v2/accounts/{accountUuid}/subscriptions/forecast | Get forecast snapshot for a given account
[**SubscriptionsControllerGetSubscription**](DynatracePlatformSubscriptionAPI.md#SubscriptionsControllerGetSubscription) | **Get** /sub/v2/accounts/{accountUuid}/subscriptions/{subscriptionUuid} | Get specific subscription by uuid
[**SubscriptionsControllerGetTotalSubscriptionCost**](DynatracePlatformSubscriptionAPI.md#SubscriptionsControllerGetTotalSubscriptionCost) | **Get** /sub/v2/accounts/{accountUuid}/subscriptions/{subscriptionUuid}/cost | Get aggregated cost data grouped by date for a given subscription
[**SubscriptionsControllerGetTotalSubscriptionUsage**](DynatracePlatformSubscriptionAPI.md#SubscriptionsControllerGetTotalSubscriptionUsage) | **Get** /sub/v2/accounts/{accountUuid}/subscriptions/{subscriptionUuid}/usage | Get aggregated usage data grouped by date for a given subscription
[**SubscriptionsControllerListSubscriptions**](DynatracePlatformSubscriptionAPI.md#SubscriptionsControllerListSubscriptions) | **Get** /sub/v2/accounts/{accountUuid}/subscriptions | Get list of subscriptions by account uuid



## SubscriptionsControllerGetEnvironmentCost

> SubscriptionEnvironmentCostListDto SubscriptionsControllerGetEnvironmentCost(ctx, accountUuid, subscriptionUuid).StartTime(startTime).EndTime(endTime).EnvironmentIds(environmentIds).CapabilityKeys(capabilityKeys).Execute()

Get cost summary grouped by environment for a given subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
    subscriptionUuid := "subscriptionUuid_example" // string | The UUID of the requested subscription
    startTime := time.Now() // time.Time | Date as ISO string
    endTime := time.Now() // time.Time | Date as ISO string
    environmentIds := []string{"Inner_example"} // []string |  (optional)
    capabilityKeys := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetEnvironmentCost(context.Background(), accountUuid, subscriptionUuid).StartTime(startTime).EndTime(endTime).EnvironmentIds(environmentIds).CapabilityKeys(capabilityKeys).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetEnvironmentCost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsControllerGetEnvironmentCost`: SubscriptionEnvironmentCostListDto
    fmt.Fprintf(os.Stdout, "Response from `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetEnvironmentCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**subscriptionUuid** | **string** | The UUID of the requested subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsControllerGetEnvironmentCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **time.Time** | Date as ISO string | 
 **endTime** | **time.Time** | Date as ISO string | 
 **environmentIds** | **[]string** |  | 
 **capabilityKeys** | **[]string** |  | 

### Return type

[**SubscriptionEnvironmentCostListDto**](SubscriptionEnvironmentCostListDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsControllerGetEnvironmentUsage

> SubscriptionEnvironmentUsageListDto SubscriptionsControllerGetEnvironmentUsage(ctx, accountUuid, subscriptionUuid).StartTime(startTime).EndTime(endTime).EnvironmentIds(environmentIds).CapabilityKeys(capabilityKeys).Execute()

Get usage summary grouped by environment for a given subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
    subscriptionUuid := "subscriptionUuid_example" // string | The UUID of the requested subscription
    startTime := time.Now() // time.Time | Date as ISO string
    endTime := time.Now() // time.Time | Date as ISO string
    environmentIds := []string{"Inner_example"} // []string |  (optional)
    capabilityKeys := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetEnvironmentUsage(context.Background(), accountUuid, subscriptionUuid).StartTime(startTime).EndTime(endTime).EnvironmentIds(environmentIds).CapabilityKeys(capabilityKeys).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetEnvironmentUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsControllerGetEnvironmentUsage`: SubscriptionEnvironmentUsageListDto
    fmt.Fprintf(os.Stdout, "Response from `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetEnvironmentUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**subscriptionUuid** | **string** | The UUID of the requested subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsControllerGetEnvironmentUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startTime** | **time.Time** | Date as ISO string | 
 **endTime** | **time.Time** | Date as ISO string | 
 **environmentIds** | **[]string** |  | 
 **capabilityKeys** | **[]string** |  | 

### Return type

[**SubscriptionEnvironmentUsageListDto**](SubscriptionEnvironmentUsageListDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsControllerGetEvents

> []Event SubscriptionsControllerGetEvents(ctx, accountUuid).StartTime(startTime).EndTime(endTime).EventType(eventType).Execute()

Get all notification events for a given subscription

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
    startTime := time.Now() // time.Time | Date as ISO string (optional)
    endTime := time.Now() // time.Time | Date as ISO string (optional)
    eventType := "eventType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetEvents(context.Background(), accountUuid).StartTime(startTime).EndTime(endTime).EventType(eventType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsControllerGetEvents`: []Event
    fmt.Fprintf(os.Stdout, "Response from `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsControllerGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTime** | **time.Time** | Date as ISO string | 
 **endTime** | **time.Time** | Date as ISO string | 
 **eventType** | **string** |  | 

### Return type

[**[]Event**](Event.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsControllerGetForecast

> Forecast SubscriptionsControllerGetForecast(ctx, accountUuid).Execute()

Get forecast snapshot for a given account

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
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetForecast(context.Background(), accountUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetForecast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsControllerGetForecast`: Forecast
    fmt.Fprintf(os.Stdout, "Response from `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetForecast`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsControllerGetForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Forecast**](Forecast.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsControllerGetSubscription

> SubscriptionDto SubscriptionsControllerGetSubscription(ctx, accountUuid, subscriptionUuid).Execute()

Get specific subscription by uuid

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
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
    subscriptionUuid := "subscriptionUuid_example" // string | The UUID of the requested subscription

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetSubscription(context.Background(), accountUuid, subscriptionUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsControllerGetSubscription`: SubscriptionDto
    fmt.Fprintf(os.Stdout, "Response from `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**subscriptionUuid** | **string** | The UUID of the requested subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsControllerGetSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubscriptionDto**](SubscriptionDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsControllerGetTotalSubscriptionCost

> SubscriptionCostListDto SubscriptionsControllerGetTotalSubscriptionCost(ctx, accountUuid, subscriptionUuid).EnvironmentIds(environmentIds).CapabilityKeys(capabilityKeys).Execute()

Get aggregated cost data grouped by date for a given subscription

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
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
    subscriptionUuid := "subscriptionUuid_example" // string | The UUID of the requested subscription
    environmentIds := []string{"Inner_example"} // []string |  (optional)
    capabilityKeys := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetTotalSubscriptionCost(context.Background(), accountUuid, subscriptionUuid).EnvironmentIds(environmentIds).CapabilityKeys(capabilityKeys).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetTotalSubscriptionCost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsControllerGetTotalSubscriptionCost`: SubscriptionCostListDto
    fmt.Fprintf(os.Stdout, "Response from `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetTotalSubscriptionCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**subscriptionUuid** | **string** | The UUID of the requested subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsControllerGetTotalSubscriptionCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentIds** | **[]string** |  | 
 **capabilityKeys** | **[]string** |  | 

### Return type

[**SubscriptionCostListDto**](SubscriptionCostListDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsControllerGetTotalSubscriptionUsage

> SubscriptionUsageListDto SubscriptionsControllerGetTotalSubscriptionUsage(ctx, accountUuid, subscriptionUuid).EnvironmentIds(environmentIds).CapabilityKeys(capabilityKeys).Execute()

Get aggregated usage data grouped by date for a given subscription

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
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
    subscriptionUuid := "subscriptionUuid_example" // string | The UUID of the requested subscription
    environmentIds := []string{"Inner_example"} // []string |  (optional)
    capabilityKeys := []string{"Inner_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetTotalSubscriptionUsage(context.Background(), accountUuid, subscriptionUuid).EnvironmentIds(environmentIds).CapabilityKeys(capabilityKeys).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetTotalSubscriptionUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsControllerGetTotalSubscriptionUsage`: SubscriptionUsageListDto
    fmt.Fprintf(os.Stdout, "Response from `DynatracePlatformSubscriptionAPI.SubscriptionsControllerGetTotalSubscriptionUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**subscriptionUuid** | **string** | The UUID of the requested subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsControllerGetTotalSubscriptionUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentIds** | **[]string** |  | 
 **capabilityKeys** | **[]string** |  | 

### Return type

[**SubscriptionUsageListDto**](SubscriptionUsageListDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionsControllerListSubscriptions

> SubscriptionListDto SubscriptionsControllerListSubscriptions(ctx, accountUuid).Execute()

Get list of subscriptions by account uuid

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
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DynatracePlatformSubscriptionAPI.SubscriptionsControllerListSubscriptions(context.Background(), accountUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DynatracePlatformSubscriptionAPI.SubscriptionsControllerListSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriptionsControllerListSubscriptions`: SubscriptionListDto
    fmt.Fprintf(os.Stdout, "Response from `DynatracePlatformSubscriptionAPI.SubscriptionsControllerListSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionsControllerListSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionListDto**](SubscriptionListDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

