# \GroupManagementAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsControllerCreateGroups**](GroupManagementAPI.md#GroupsControllerCreateGroups) | **Post** /iam/v1/accounts/{accountUuid}/groups | Creates new user groups
[**GroupsControllerDeleteGroup**](GroupManagementAPI.md#GroupsControllerDeleteGroup) | **Delete** /iam/v1/accounts/{accountUuid}/groups/{groupUuid} | Deletes a user group
[**GroupsControllerEditGroup**](GroupManagementAPI.md#GroupsControllerEditGroup) | **Put** /iam/v1/accounts/{accountUuid}/groups/{groupUuid} | Edits a user group
[**GroupsControllerGetGroups**](GroupManagementAPI.md#GroupsControllerGetGroups) | **Get** /iam/v1/accounts/{accountUuid}/groups | Lists all user groups of an account
[**GroupsControllerGetUsersForGroup**](GroupManagementAPI.md#GroupsControllerGetUsersForGroup) | **Get** /iam/v1/accounts/{accountUuid}/groups/{groupUuid}/users | Lists all members of a group



## GroupsControllerCreateGroups

> []GetGroupDto GroupsControllerCreateGroups(ctx, accountUuid).PutGroupDto(putGroupDto).Execute()

Creates new user groups

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
    putGroupDto := []openapiclient.PutGroupDto{*openapiclient.NewPutGroupDto("Name_example")} // []PutGroupDto | The body of the request. Contains a list of configurations for new groups.    Do not specify a UUID. A UUID is assigned automatically by Dynatrace. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementAPI.GroupsControllerCreateGroups(context.Background(), accountUuid).PutGroupDto(putGroupDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementAPI.GroupsControllerCreateGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsControllerCreateGroups`: []GetGroupDto
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementAPI.GroupsControllerCreateGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsControllerCreateGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putGroupDto** | [**[]PutGroupDto**](PutGroupDto.md) | The body of the request. Contains a list of configurations for new groups.    Do not specify a UUID. A UUID is assigned automatically by Dynatrace.  | 

### Return type

[**[]GetGroupDto**](GetGroupDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsControllerDeleteGroup

> GroupsControllerDeleteGroup(ctx, accountUuid, groupUuid).Execute()

Deletes a user group

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
    groupUuid := "groupUuid_example" // string | The UUID of the required user group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupManagementAPI.GroupsControllerDeleteGroup(context.Background(), accountUuid, groupUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementAPI.GroupsControllerDeleteGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**groupUuid** | **string** | The UUID of the required user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsControllerDeleteGroupRequest struct via the builder pattern


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


## GroupsControllerEditGroup

> GroupsControllerEditGroup(ctx, accountUuid, groupUuid).PutGroupDto(putGroupDto).Execute()

Edits a user group

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
    groupUuid := "groupUuid_example" // string | The UUID of the required user group.
    putGroupDto := *openapiclient.NewPutGroupDto("Name_example") // PutGroupDto | The body of the request. Contains the updated parameters of the group.    You can't change the UUID of the group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GroupManagementAPI.GroupsControllerEditGroup(context.Background(), accountUuid, groupUuid).PutGroupDto(putGroupDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementAPI.GroupsControllerEditGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**groupUuid** | **string** | The UUID of the required user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsControllerEditGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putGroupDto** | [**PutGroupDto**](PutGroupDto.md) | The body of the request. Contains the updated parameters of the group.    You can&#39;t change the UUID of the group. | 

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


## GroupsControllerGetGroups

> GroupListDto GroupsControllerGetGroups(ctx, accountUuid).Execute()

Lists all user groups of an account

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
    resp, r, err := apiClient.GroupManagementAPI.GroupsControllerGetGroups(context.Background(), accountUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementAPI.GroupsControllerGetGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsControllerGetGroups`: GroupListDto
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementAPI.GroupsControllerGetGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsControllerGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GroupListDto**](GroupListDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsControllerGetUsersForGroup

> GroupUserListDto GroupsControllerGetUsersForGroup(ctx, accountUuid, groupUuid).Execute()

Lists all members of a group

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
    groupUuid := "groupUuid_example" // string | The UUID of the required user group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupManagementAPI.GroupsControllerGetUsersForGroup(context.Background(), accountUuid, groupUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupManagementAPI.GroupsControllerGetUsersForGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GroupsControllerGetUsersForGroup`: GroupUserListDto
    fmt.Fprintf(os.Stdout, "Response from `GroupManagementAPI.GroupsControllerGetUsersForGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountUuid** | **string** | The ID of the required account.    You can find the UUID on the **Account &gt; Account management API** page, during creation of an OAuth client. | 
**groupUuid** | **string** | The UUID of the required user group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsControllerGetUsersForGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GroupUserListDto**](GroupUserListDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

