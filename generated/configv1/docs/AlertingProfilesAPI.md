# \AlertingProfilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlertingProfile**](AlertingProfilesAPI.md#CreateAlertingProfile) | **Post** /alertingProfiles | Creates a new alerting profile
[**CreateOrUpdateAlertingProfile**](AlertingProfilesAPI.md#CreateOrUpdateAlertingProfile) | **Put** /alertingProfiles/{id} | Updates an existing alerting profile
[**DeleteAlertingProfile**](AlertingProfilesAPI.md#DeleteAlertingProfile) | **Delete** /alertingProfiles/{id} | Deletes the specified alerting profile
[**GetAlertingProfile**](AlertingProfilesAPI.md#GetAlertingProfile) | **Get** /alertingProfiles/{id} | Gets the properties of the specified alerting profile
[**GetAlertingProfiles**](AlertingProfilesAPI.md#GetAlertingProfiles) | **Get** /alertingProfiles | Lists all alerting profiles available in your environment
[**ValidateCreateAlertingProfile**](AlertingProfilesAPI.md#ValidateCreateAlertingProfile) | **Post** /alertingProfiles/validator | Validates the payload the &#x60;POST /alertingProfiles&#x60; request
[**ValidateCreateOrUpdateAlertingProfile**](AlertingProfilesAPI.md#ValidateCreateOrUpdateAlertingProfile) | **Post** /alertingProfiles/{id}/validator | Validates the payload the &#x60;PUT /alertingProfiles/{id}&#x60; request



## CreateAlertingProfile

> EntityShortRepresentation CreateAlertingProfile(ctx).AlertingProfile(alertingProfile).Execute()

Creates a new alerting profile



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
    alertingProfile := *openapiclient.NewAlertingProfile("DisplayName_example") // AlertingProfile | The JSON body of the request. Contains parameters of the new alerting profile. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingProfilesAPI.CreateAlertingProfile(context.Background()).AlertingProfile(alertingProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.CreateAlertingProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAlertingProfile`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesAPI.CreateAlertingProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAlertingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertingProfile** | [**AlertingProfile**](AlertingProfile.md) | The JSON body of the request. Contains parameters of the new alerting profile. | 

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


## CreateOrUpdateAlertingProfile

> EntityShortRepresentation CreateOrUpdateAlertingProfile(ctx, id).AlertingProfile(alertingProfile).Execute()

Updates an existing alerting profile



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the alerting profile to be updated.
    alertingProfile := *openapiclient.NewAlertingProfile("DisplayName_example") // AlertingProfile | The JSON body of the request. Contains updated parameters of the alerting profile. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingProfilesAPI.CreateOrUpdateAlertingProfile(context.Background(), id).AlertingProfile(alertingProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.CreateOrUpdateAlertingProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrUpdateAlertingProfile`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesAPI.CreateOrUpdateAlertingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the alerting profile to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateAlertingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertingProfile** | [**AlertingProfile**](AlertingProfile.md) | The JSON body of the request. Contains updated parameters of the alerting profile. | 

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


## DeleteAlertingProfile

> DeleteAlertingProfile(ctx, id).Execute()

Deletes the specified alerting profile



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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the alerting profile to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertingProfilesAPI.DeleteAlertingProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.DeleteAlertingProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the alerting profile to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAlertingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingProfile

> AlertingProfile GetAlertingProfile(ctx, id).Execute()

Gets the properties of the specified alerting profile

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the required alerting profile.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AlertingProfilesAPI.GetAlertingProfile(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.GetAlertingProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertingProfile`: AlertingProfile
    fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesAPI.GetAlertingProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the required alerting profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AlertingProfile**](AlertingProfile.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAlertingProfiles

> StubList GetAlertingProfiles(ctx).Execute()

Lists all alerting profiles available in your environment

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
    resp, r, err := apiClient.AlertingProfilesAPI.GetAlertingProfiles(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.GetAlertingProfiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlertingProfiles`: StubList
    fmt.Fprintf(os.Stdout, "Response from `AlertingProfilesAPI.GetAlertingProfiles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlertingProfilesRequest struct via the builder pattern


### Return type

[**StubList**](StubList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCreateAlertingProfile

> ValidateCreateAlertingProfile(ctx).AlertingProfile(alertingProfile).Execute()

Validates the payload the `POST /alertingProfiles` request

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
    alertingProfile := *openapiclient.NewAlertingProfile("DisplayName_example") // AlertingProfile | The JSON body of the request. Contains parameters of the alerting profile to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertingProfilesAPI.ValidateCreateAlertingProfile(context.Background()).AlertingProfile(alertingProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.ValidateCreateAlertingProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateAlertingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alertingProfile** | [**AlertingProfile**](AlertingProfile.md) | The JSON body of the request. Contains parameters of the alerting profile to be validated. | 

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


## ValidateCreateOrUpdateAlertingProfile

> ValidateCreateOrUpdateAlertingProfile(ctx, id).AlertingProfile(alertingProfile).Execute()

Validates the payload the `PUT /alertingProfiles/{id}` request

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the alerting profile to be validated.
    alertingProfile := *openapiclient.NewAlertingProfile("DisplayName_example") // AlertingProfile | The JSON body of the request. Contains parameters of the alerting profile to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AlertingProfilesAPI.ValidateCreateOrUpdateAlertingProfile(context.Background(), id).AlertingProfile(alertingProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AlertingProfilesAPI.ValidateCreateOrUpdateAlertingProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the alerting profile to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateOrUpdateAlertingProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alertingProfile** | [**AlertingProfile**](AlertingProfile.md) | The JSON body of the request. Contains parameters of the alerting profile to be validated. | 

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

