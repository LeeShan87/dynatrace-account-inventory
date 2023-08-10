# \UserManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersControllerAddUserToGroups**](UserManagementAPI.md#UsersControllerAddUserToGroups) | **Post** /iam/v1/accounts/{accountUuid}/users/{email} | Adds a user to groups. Any existing group membership remains unaffected.
[**UsersControllerCreateUserForAccount**](UserManagementAPI.md#UsersControllerCreateUserForAccount) | **Post** /iam/v1/accounts/{accountUuid}/users | Creates a new user in an account
[**UsersControllerGetUserGroups**](UserManagementAPI.md#UsersControllerGetUserGroups) | **Get** /iam/v1/accounts/{accountUuid}/users/{email} | Lists all groups of a user
[**UsersControllerGetUsers**](UserManagementAPI.md#UsersControllerGetUsers) | **Get** /iam/v1/accounts/{accountUuid}/users | Lists all users of an account
[**UsersControllerRemoveUserFromAccount**](UserManagementAPI.md#UsersControllerRemoveUserFromAccount) | **Delete** /iam/v1/accounts/{accountUuid}/users/{email} | Removes a user from an account
[**UsersControllerRemoveUserFromGroups**](UserManagementAPI.md#UsersControllerRemoveUserFromGroups) | **Delete** /iam/v1/accounts/{accountUuid}/users/{email}/groups | Removes a user from groups
[**UsersControllerReplaceUserGroups**](UserManagementAPI.md#UsersControllerReplaceUserGroups) | **Put** /iam/v1/accounts/{accountUuid}/users/{email}/groups | Sets group membership of a user. Any existing membership is overwritten.



## UsersControllerAddUserToGroups

> UsersControllerAddUserToGroups(ctx, accountUuid, email).RequestBody(requestBody).Execute()

Adds a user to groups. Any existing group membership remains unaffected.

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
    email := "email_example" // string | The email address of the required user.
    requestBody := []string{"Property_example"} // []string | The body of the request. Contains a list of groups (specified by UUIDs) to which the user is to be added.    Any existing group membership remains unaffected.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserManagementAPI.UsersControllerAddUserToGroups(context.Background(), accountUuid, email).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UsersControllerAddUserToGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**email** | **string** | The email address of the required user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersControllerAddUserToGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]string** | The body of the request. Contains a list of groups (specified by UUIDs) to which the user is to be added.    Any existing group membership remains unaffected. | 

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


## UsersControllerCreateUserForAccount

> UsersControllerCreateUserForAccount(ctx, accountUuid).UserEmailDto(userEmailDto).Execute()

Creates a new user in an account

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
    userEmailDto := *openapiclient.NewUserEmailDto("Email_example") // UserEmailDto | The JSON body of the request. Contains the email address of the new user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserManagementAPI.UsersControllerCreateUserForAccount(context.Background(), accountUuid).UserEmailDto(userEmailDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UsersControllerCreateUserForAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersControllerCreateUserForAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userEmailDto** | [**UserEmailDto**](UserEmailDto.md) | The JSON body of the request. Contains the email address of the new user. | 

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


## UsersControllerGetUserGroups

> GroupUserDto UsersControllerGetUserGroups(ctx, accountUuid, email).Execute()

Lists all groups of a user

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
    email := "email_example" // string | The email address of the required user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementAPI.UsersControllerGetUserGroups(context.Background(), accountUuid, email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UsersControllerGetUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersControllerGetUserGroups`: GroupUserDto
    fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.UsersControllerGetUserGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**email** | **string** | The email address of the required user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersControllerGetUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupUserDto**](GroupUserDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersControllerGetUsers

> UserListDto UsersControllerGetUsers(ctx, accountUuid).ServiceUsers(serviceUsers).Execute()

Lists all users of an account

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
    serviceUsers := true // bool | Specifies whether service users are included in results. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserManagementAPI.UsersControllerGetUsers(context.Background(), accountUuid).ServiceUsers(serviceUsers).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UsersControllerGetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersControllerGetUsers`: UserListDto
    fmt.Fprintf(os.Stdout, "Response from `UserManagementAPI.UsersControllerGetUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersControllerGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serviceUsers** | **bool** | Specifies whether service users are included in results. | 

### Return type

[**UserListDto**](UserListDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersControllerRemoveUserFromAccount

> UsersControllerRemoveUserFromAccount(ctx, accountUuid, email).Execute()

Removes a user from an account

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
    email := "email_example" // string | The email address of the required user.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserManagementAPI.UsersControllerRemoveUserFromAccount(context.Background(), accountUuid, email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UsersControllerRemoveUserFromAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**email** | **string** | The email address of the required user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersControllerRemoveUserFromAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## UsersControllerRemoveUserFromGroups

> UsersControllerRemoveUserFromGroups(ctx, accountUuid, email).GroupUuid(groupUuid).Execute()

Removes a user from groups

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
    email := "email_example" // string | The email address of the required user.
    groupUuid := []string{"Inner_example"} // []string | A list of groups the user is no longer a member of.    To specify several groups, use the following format: `group-uuid=aaaaaa&group-uuid=bbbb`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserManagementAPI.UsersControllerRemoveUserFromGroups(context.Background(), accountUuid, email).GroupUuid(groupUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UsersControllerRemoveUserFromGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**email** | **string** | The email address of the required user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersControllerRemoveUserFromGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **groupUuid** | **[]string** | A list of groups the user is no longer a member of.    To specify several groups, use the following format: &#x60;group-uuid&#x3D;aaaaaa&amp;group-uuid&#x3D;bbbb&#x60;. | 

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


## UsersControllerReplaceUserGroups

> UsersControllerReplaceUserGroups(ctx, accountUuid, email).RequestBody(requestBody).Execute()

Sets group membership of a user. Any existing membership is overwritten.

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
    email := "email_example" // string | The email address of the required user.
    requestBody := []string{"Property_example"} // []string | The body of the request. Contains a list of groups (specified by UUIDs) where the user is to be a member.    The user will be removed from any group that is not specified here.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserManagementAPI.UsersControllerReplaceUserGroups(context.Background(), accountUuid, email).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserManagementAPI.UsersControllerReplaceUserGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**email** | **string** | The email address of the required user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersControllerReplaceUserGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **[]string** | The body of the request. Contains a list of groups (specified by UUIDs) where the user is to be a member.    The user will be removed from any group that is not specified here. | 

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

