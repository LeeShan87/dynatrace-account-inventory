# \EnvironmentManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentResourcesControllerGetEnvironmentResources**](EnvironmentManagementAPI.md#EnvironmentResourcesControllerGetEnvironmentResources) | **Get** /env/v1/accounts/{accountUuid}/environments | Lists all environments and management zones of an account



## EnvironmentResourcesControllerGetEnvironmentResources

> EnvironmentResourceDto EnvironmentResourcesControllerGetEnvironmentResources(ctx, accountUuid).Execute()

Lists all environments and management zones of an account

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
    resp, r, err := apiClient.EnvironmentManagementAPI.EnvironmentResourcesControllerGetEnvironmentResources(context.Background(), accountUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentManagementAPI.EnvironmentResourcesControllerGetEnvironmentResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentResourcesControllerGetEnvironmentResources`: EnvironmentResourceDto
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentManagementAPI.EnvironmentResourcesControllerGetEnvironmentResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentResourcesControllerGetEnvironmentResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnvironmentResourceDto**](EnvironmentResourceDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

