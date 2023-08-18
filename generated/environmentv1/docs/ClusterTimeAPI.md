# \ClusterTimeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentClusterTime**](ClusterTimeAPI.md#GetCurrentClusterTime) | **Get** /time | Gets the current time of the Dynatrace server in UTC milliseconds



## GetCurrentClusterTime

> int64 GetCurrentClusterTime(ctx).Execute()

Gets the current time of the Dynatrace server in UTC milliseconds

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
    resp, r, err := apiClient.ClusterTimeAPI.GetCurrentClusterTime(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterTimeAPI.GetCurrentClusterTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentClusterTime`: int64
    fmt.Fprintf(os.Stdout, "Response from `ClusterTimeAPI.GetCurrentClusterTime`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentClusterTimeRequest struct via the builder pattern


### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

