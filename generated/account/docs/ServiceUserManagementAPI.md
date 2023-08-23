# \ServiceUserManagementAPI

All URIs are relative to _http://localhost_

| Method                                                                                                                                 | HTTP request                                                       | Description                              |
| -------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------ | ---------------------------------------- |
| [**ServiceUsersControllerCreateServiceUserForAccount**](ServiceUserManagementAPI.md#ServiceUsersControllerCreateServiceUserForAccount) | **Post** /iam/v1/accounts/{accountUuid}/service-users              | Creates a new service user in an account |
| [**ServiceUsersControllerDeleteUser**](ServiceUserManagementAPI.md#ServiceUsersControllerDeleteUser)                                   | **Delete** /iam/v1/accounts/{accountUuid}/service-users/{userUuid} | Removes service user                     |
| [**ServiceUsersControllerGetServiceUsersFromAccount**](ServiceUserManagementAPI.md#ServiceUsersControllerGetServiceUsersFromAccount)   | **Get** /iam/v1/accounts/{accountUuid}/service-users               | Get service users assigned to account    |

## ServiceUsersControllerCreateServiceUserForAccount

> ServiceUserUuidDto ServiceUsersControllerCreateServiceUserForAccount(ctx, accountUuid).ServiceUserNameDto(serviceUserNameDto).Execute()

Creates a new service user in an account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/LeeShan87/dynatrace-account-inventory/generated/account"
)

func main() {
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
    serviceUserNameDto := *openapiclient.NewServiceUserNameDto("Name_example") // ServiceUserNameDto | The JSON body of the request. Contains the name of the new service user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceUserManagementAPI.ServiceUsersControllerCreateServiceUserForAccount(context.Background(), accountUuid).ServiceUserNameDto(serviceUserNameDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceUserManagementAPI.ServiceUsersControllerCreateServiceUserForAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceUsersControllerCreateServiceUserForAccount`: ServiceUserUuidDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceUserManagementAPI.ServiceUsersControllerCreateServiceUserForAccount`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                                                                                    | Notes |
| --------------- | ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                    |
| **accountUuid** | **string**          | The ID of the required account. You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. |

### Other Parameters

Other parameters are passed through a pointer to a apiServiceUsersControllerCreateServiceUserForAccountRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**serviceUserNameDto** | [**ServiceUserNameDto**](ServiceUserNameDto.md) | The JSON body of the request. Contains the name of the new service user. |

### Return type

[**ServiceUserUuidDto**](ServiceUserUuidDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ServiceUsersControllerDeleteUser

> ServiceUsersControllerDeleteUser(ctx, accountUuid, userUuid).Execute()

Removes service user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/LeeShan87/dynatrace-account-inventory/generated/account"
)

func main() {
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.
    userUuid := "userUuid_example" // string | The UUID of the required user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ServiceUserManagementAPI.ServiceUsersControllerDeleteUser(context.Background(), accountUuid, userUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceUserManagementAPI.ServiceUsersControllerDeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name            | Type                | Description                                                                                                                                    | Notes |
| --------------- | ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                    |
| **accountUuid** | **string**          | The ID of the required account. You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. |
| **userUuid**    | **string**          | The UUID of the required user.                                                                                                                 |

### Other Parameters

Other parameters are passed through a pointer to a apiServiceUsersControllerDeleteUserRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ServiceUsersControllerGetServiceUsersFromAccount

> ExternalServiceUsersPageDto ServiceUsersControllerGetServiceUsersFromAccount(ctx, accountUuid).Execute()

Get service users assigned to account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/LeeShan87/dynatrace-account-inventory/generated/account"
)

func main() {
    accountUuid := "accountUuid_example" // string | The ID of the required account.    You can find the UUID on the **Account > Account management API** page, during creation of an OAuth client.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServiceUserManagementAPI.ServiceUsersControllerGetServiceUsersFromAccount(context.Background(), accountUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceUserManagementAPI.ServiceUsersControllerGetServiceUsersFromAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceUsersControllerGetServiceUsersFromAccount`: ExternalServiceUsersPageDto
    fmt.Fprintf(os.Stdout, "Response from `ServiceUserManagementAPI.ServiceUsersControllerGetServiceUsersFromAccount`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                                                                                    | Notes |
| --------------- | ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                    |
| **accountUuid** | **string**          | The ID of the required account. You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. |

### Other Parameters

Other parameters are passed through a pointer to a apiServiceUsersControllerGetServiceUsersFromAccountRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**ExternalServiceUsersPageDto**](ExternalServiceUsersPageDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
