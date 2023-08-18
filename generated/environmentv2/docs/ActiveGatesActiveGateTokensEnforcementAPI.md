# \ActiveGatesActiveGateTokensEnforcementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokenEnforcement**](ActiveGatesActiveGateTokensEnforcementAPI.md#GetTokenEnforcement) | **Get** /activeGates/tokenEnforcement | Gets the status of ActiveGate tokens enforcement



## GetTokenEnforcement

> ActiveGateTokenEnforcement GetTokenEnforcement(ctx).Execute()

Gets the status of ActiveGate tokens enforcement

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
    resp, r, err := apiClient.ActiveGatesActiveGateTokensEnforcementAPI.GetTokenEnforcement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActiveGatesActiveGateTokensEnforcementAPI.GetTokenEnforcement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenEnforcement`: ActiveGateTokenEnforcement
    fmt.Fprintf(os.Stdout, "Response from `ActiveGatesActiveGateTokensEnforcementAPI.GetTokenEnforcement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenEnforcementRequest struct via the builder pattern


### Return type

[**ActiveGateTokenEnforcement**](ActiveGateTokenEnforcement.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

