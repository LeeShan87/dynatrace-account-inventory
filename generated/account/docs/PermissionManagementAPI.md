# \PermissionManagementAPI

All URIs are relative to _http://localhost_

| Method                                                                                                                          | HTTP request                                                             | Description                                                                  |
| ------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------ | ---------------------------------------------------------------------------- |
| [**PermissionsControllerAddGroupPermissions**](PermissionManagementAPI.md#PermissionsControllerAddGroupPermissions)             | **Post** /iam/v1/accounts/{accountUuid}/groups/{groupUuid}/permissions   | Assigns permissions to a user group. Existing permissions remain unaffected. |
| [**PermissionsControllerGetGroupPermissions**](PermissionManagementAPI.md#PermissionsControllerGetGroupPermissions)             | **Get** /iam/v1/accounts/{accountUuid}/groups/{groupUuid}/permissions    | Lists all permissions of a user group                                        |
| [**PermissionsControllerOverwriteGroupPermissions**](PermissionManagementAPI.md#PermissionsControllerOverwriteGroupPermissions) | **Put** /iam/v1/accounts/{accountUuid}/groups/{groupUuid}/permissions    | Sets permissions of a user group. Existing permissions are overwritten.      |
| [**PermissionsControllerRemoveGroupPermissions**](PermissionManagementAPI.md#PermissionsControllerRemoveGroupPermissions)       | **Delete** /iam/v1/accounts/{accountUuid}/groups/{groupUuid}/permissions | Removes a permission from a user group                                       |

## PermissionsControllerAddGroupPermissions

> PermissionsControllerAddGroupPermissions(ctx, accountUuid, groupUuid).PermissionsDto(permissionsDto).Execute()

Assigns permissions to a user group. Existing permissions remain unaffected.

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
    groupUuid := "groupUuid_example" // string | The UUID of the required user group.
    permissionsDto := []openapiclient.PermissionsDto{*openapiclient.NewPermissionsDto("PermissionName_example", "Scope_example", "ScopeType_example")} // []PermissionsDto | The body of the request. Contains a list of permissions to be assigned to the group.   Existing permissions remain unaffected.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PermissionManagementAPI.PermissionsControllerAddGroupPermissions(context.Background(), accountUuid, groupUuid).PermissionsDto(permissionsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.PermissionsControllerAddGroupPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name            | Type                | Description                                                                                                                                    | Notes |
| --------------- | ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                    |
| **accountUuid** | **string**          | The ID of the required account. You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. |
| **groupUuid**   | **string**          | The UUID of the required user group.                                                                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsControllerAddGroupPermissionsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**permissionsDto** | [**[]PermissionsDto**](PermissionsDto.md) | The body of the request. Contains a list of permissions to be assigned to the group. Existing permissions remain unaffected. |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PermissionsControllerGetGroupPermissions

> PermissionsGroupDto PermissionsControllerGetGroupPermissions(ctx, accountUuid, groupUuid).Execute()

Lists all permissions of a user group

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
    groupUuid := "groupUuid_example" // string | The UUID of the required user group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PermissionManagementAPI.PermissionsControllerGetGroupPermissions(context.Background(), accountUuid, groupUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.PermissionsControllerGetGroupPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PermissionsControllerGetGroupPermissions`: PermissionsGroupDto
    fmt.Fprintf(os.Stdout, "Response from `PermissionManagementAPI.PermissionsControllerGetGroupPermissions`: %v\n", resp)
}
```

### Path Parameters

| Name            | Type                | Description                                                                                                                                    | Notes |
| --------------- | ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                    |
| **accountUuid** | **string**          | The ID of the required account. You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. |
| **groupUuid**   | **string**          | The UUID of the required user group.                                                                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsControllerGetGroupPermissionsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**PermissionsGroupDto**](PermissionsGroupDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PermissionsControllerOverwriteGroupPermissions

> PermissionsControllerOverwriteGroupPermissions(ctx, accountUuid, groupUuid).PermissionsDto(permissionsDto).Execute()

Sets permissions of a user group. Existing permissions are overwritten.

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
    groupUuid := "groupUuid_example" // string | The UUID of the required user group.
    permissionsDto := []openapiclient.PermissionsDto{*openapiclient.NewPermissionsDto("PermissionName_example", "Scope_example", "ScopeType_example")} // []PermissionsDto | The body of the request. Contains a list of permissions to be assigned to the group.    Existing permissions are overwritten.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PermissionManagementAPI.PermissionsControllerOverwriteGroupPermissions(context.Background(), accountUuid, groupUuid).PermissionsDto(permissionsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.PermissionsControllerOverwriteGroupPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name            | Type                | Description                                                                                                                                    | Notes |
| --------------- | ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                    |
| **accountUuid** | **string**          | The ID of the required account. You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. |
| **groupUuid**   | **string**          | The UUID of the required user group.                                                                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsControllerOverwriteGroupPermissionsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**permissionsDto** | [**[]PermissionsDto**](PermissionsDto.md) | The body of the request. Contains a list of permissions to be assigned to the group. Existing permissions are overwritten. |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PermissionsControllerRemoveGroupPermissions

> PermissionsControllerRemoveGroupPermissions(ctx, accountUuid, groupUuid).Scope(scope).PermissionName(permissionName).ScopeType(scopeType).Execute()

Removes a permission from a user group

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
    groupUuid := "groupUuid_example" // string | The UUID of the required user group.
    scope := "scope_example" // string | The scope of the permission to be deleted. Depending on the type of the scope, specify one of the following:    * `account`: The UUID of the account.  * `tenant`: The ID of the environment.  * `management-zone`: The ID of the management zone from an environment in `{environment-id}:{management-zone-id}` format.
    permissionName := "permissionName_example" // string | The name of the permission to be deleted.
    scopeType := "scopeType_example" // string | The scope type of the permission to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PermissionManagementAPI.PermissionsControllerRemoveGroupPermissions(context.Background(), accountUuid, groupUuid).Scope(scope).PermissionName(permissionName).ScopeType(scopeType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PermissionManagementAPI.PermissionsControllerRemoveGroupPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name            | Type                | Description                                                                                                                                    | Notes |
| --------------- | ------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**         | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                    |
| **accountUuid** | **string**          | The ID of the required account. You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. |
| **groupUuid**   | **string**          | The UUID of the required user group.                                                                                                           |

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsControllerRemoveGroupPermissionsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**scope** | **string** | The scope of the permission to be deleted. Depending on the type of the scope, specify one of the following: _ &#x60;account&#x60;: The UUID of the account. _ &#x60;tenant&#x60;: The ID of the environment. \* &#x60;management-zone&#x60;: The ID of the management zone from an environment in &#x60;{environment-id}:{management-zone-id}&#x60; format. |
**permissionName** | **string** | The name of the permission to be deleted. |
**scopeType** | **string** | The scope type of the permission to be deleted. |

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
