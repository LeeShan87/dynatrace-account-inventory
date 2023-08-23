# \PolicyManagementAPI

All URIs are relative to _http://localhost_

| Method                                                                                                                                              | HTTP request                                                                    | Description                                                                                                    |
| --------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | --------------------------- |
| [**PolicyControllerAppendLevelPolicyBindings**](PolicyManagementAPI.md#PolicyControllerAppendLevelPolicyBindings)                                   | **Post** /iam/v1/repo/{levelType}/{levelId}/bindings/{policyUuid}               | Adds policy bindings to a level                                                                                |
| [**PolicyControllerAppendParticularGroupBinding**](PolicyManagementAPI.md#PolicyControllerAppendParticularGroupBinding)                             | **Post** /iam/v1/repo/{levelType}/{levelId}/bindings/{policyUuid}/{groupUuid}   | Append policy bindings within a level for a user group                                                         |
| [**PolicyControllerCreateLevelPolicy**](PolicyManagementAPI.md#PolicyControllerCreateLevelPolicy)                                                   | **Post** /iam/v1/repo/{levelType}/{levelId}/policies                            | Creates a new policy                                                                                           |
| [**PolicyControllerDeleteLevelPolicy**](PolicyManagementAPI.md#PolicyControllerDeleteLevelPolicy)                                                   | **Delete** /iam/v1/repo/{levelType}/{levelId}/policies/{policyUuid}             | Deletes a policy                                                                                               |
| [**PolicyControllerDeleteLevelPolicyBindings**](PolicyManagementAPI.md#PolicyControllerDeleteLevelPolicyBindings)                                   | **Delete** /iam/v1/repo/{levelType}/{levelId}/bindings                          | Deletes all policy bindings from a level                                                                       |
| [**PolicyControllerDeleteLevelPolicyBindingsForPolicy**](PolicyManagementAPI.md#PolicyControllerDeleteLevelPolicyBindingsForPolicy)                 | **Delete** /iam/v1/repo/{levelType}/{levelId}/bindings/{policyUuid}             | Deletes all bindings of a policy                                                                               |
| [**PolicyControllerDeleteLevelPolicyBindingsForPolicyAndGroup**](PolicyManagementAPI.md#PolicyControllerDeleteLevelPolicyBindingsForPolicyAndGroup) | **Delete** /iam/v1/repo/{levelType}/{levelId}/bindings/{policyUuid}/{groupUuid} | Deletes a policy binding from a user group                                                                     |
| [**PolicyControllerGetAllLevelPoliciesBindings**](PolicyManagementAPI.md#PolicyControllerGetAllLevelPoliciesBindings)                               | **Get** /iam/v1/repo/{levelType}/{levelId}/bindings                             | Lists all policy bindings of a level                                                                           |
| [**PolicyControllerGetEffectivePermissions**](PolicyManagementAPI.md#PolicyControllerGetEffectivePermissions)                                       | **Get** /iam/v1/resolution/{levelType}/{levelId}/effectivepermissions           | Gets effective permissions for a user or group                                                                 | maturity&#x3D;EARLY_ADOPTER |
| [**PolicyControllerGetLevelDescendantsPolicyBindings**](PolicyManagementAPI.md#PolicyControllerGetLevelDescendantsPolicyBindings)                   | **Get** /iam/v1/repo/{levelType}/{levelId}/bindings/descendants/{policyUuid}    | Get policy bindings within descendants of a level                                                              |
| [**PolicyControllerGetLevelPolicies**](PolicyManagementAPI.md#PolicyControllerGetLevelPolicies)                                                     | **Get** /iam/v1/repo/{levelType}/{levelId}/policies                             | Lists all native policies of a level                                                                           |
| [**PolicyControllerGetLevelPolicy**](PolicyManagementAPI.md#PolicyControllerGetLevelPolicy)                                                         | **Get** /iam/v1/repo/{levelType}/{levelId}/policies/{policyUuid}                | Gets a policy                                                                                                  |
| [**PolicyControllerGetLevelPolicyBindings**](PolicyManagementAPI.md#PolicyControllerGetLevelPolicyBindings)                                         | **Get** /iam/v1/repo/{levelType}/{levelId}/bindings/{policyUuid}                | Get policy bindings within a level                                                                             |
| [**PolicyControllerGetLevelPolicyBindingsForGroup**](PolicyManagementAPI.md#PolicyControllerGetLevelPolicyBindingsForGroup)                         | **Get** /iam/v1/repo/{levelType}/{levelId}/bindings/{policyUuid}/{groupUuid}    | Get policy bindings within a level                                                                             |
| [**PolicyControllerGetPolicyOverviewList**](PolicyManagementAPI.md#PolicyControllerGetPolicyOverviewList)                                           | **Get** /iam/v1/repo/{levelType}/{levelId}/policies/aggregate                   | Lists all policies for a level, including inherited from higher levels                                         |
| [**PolicyControllerGetPolicyUuidsBindings**](PolicyManagementAPI.md#PolicyControllerGetPolicyUuidsBindings)                                         | **Get** /iam/v1/repo/{levelType}/{levelId}/bindings/groups/{groupUuid}          | Lists all policies for a user group                                                                            |
| [**PolicyControllerUpdateLevelPolicy**](PolicyManagementAPI.md#PolicyControllerUpdateLevelPolicy)                                                   | **Put** /iam/v1/repo/{levelType}/{levelId}/policies/{policyUuid}                | Updates a policy                                                                                               |
| [**PolicyControllerUpdateLevelPolicyBindings**](PolicyManagementAPI.md#PolicyControllerUpdateLevelPolicyBindings)                                   | **Put** /iam/v1/repo/{levelType}/{levelId}/bindings                             | Updates policy bindings of a level                                                                             |
| [**PolicyControllerUpdatePolicyBindingsToGroup**](PolicyManagementAPI.md#PolicyControllerUpdatePolicyBindingsToGroup)                               | **Put** /iam/v1/repo/{levelType}/{levelId}/bindings/groups/{groupUuid}          | Updates policies for a user group                                                                              |
| [**PolicyControllerValidateLevelPolicy**](PolicyManagementAPI.md#PolicyControllerValidateLevelPolicy)                                               | **Post** /iam/v1/repo/{levelType}/{levelId}/policies/validation/{policyUuid}    | Validates the payload for the &#x60;PUT /iam/v1/repo/{levelType}/{levelId}/policies/{policyUuid}&#x60; request |
| [**PolicyControllerValidateNewLevelPolicy**](PolicyManagementAPI.md#PolicyControllerValidateNewLevelPolicy)                                         | **Post** /iam/v1/repo/{levelType}/{levelId}/policies/validation                 | Validates the payload for the &#x60;POST /iam/v1/repo/{levelType}/{levelId}/policies&#x60; request             |

## PolicyControllerAppendLevelPolicyBindings

> PolicyControllerAppendLevelPolicyBindings(ctx, levelType, levelId, policyUuid).AppendLevelPolicyBindingsRequestDto(appendLevelPolicyBindingsRequestDto).Execute()

Adds policy bindings to a level

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the required policy.
    appendLevelPolicyBindingsRequestDto := *openapiclient.NewAppendLevelPolicyBindingsRequestDto([]string{"Groups_example"}) // AppendLevelPolicyBindingsRequestDto | The JSON body of the request. Contains user groups that must use the policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerAppendLevelPolicyBindings(context.Background(), levelType, levelId, policyUuid).AppendLevelPolicyBindingsRequestDto(appendLevelPolicyBindingsRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerAppendLevelPolicyBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name           | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType**  | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**    | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |
| **policyUuid** | **string**          | The ID of the required policy.                                                                                                                                                                                                                                                                                                                                                                                                                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerAppendLevelPolicyBindingsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**appendLevelPolicyBindingsRequestDto** | [**AppendLevelPolicyBindingsRequestDto**](AppendLevelPolicyBindingsRequestDto.md) | The JSON body of the request. Contains user groups that must use the policy. |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerAppendParticularGroupBinding

> PolicyControllerAppendParticularGroupBinding(ctx, levelType, levelId, policyUuid, groupUuid).AppendLevelPolicyBindingForGroupDto(appendLevelPolicyBindingForGroupDto).Execute()

Append policy bindings within a level for a user group

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the required policy.
    groupUuid := "groupUuid_example" // string | The ID of the required user group.
    appendLevelPolicyBindingForGroupDto := *openapiclient.NewAppendLevelPolicyBindingForGroupDto() // AppendLevelPolicyBindingForGroupDto | The JSON body of the request. Contains parameters and metadata

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerAppendParticularGroupBinding(context.Background(), levelType, levelId, policyUuid, groupUuid).AppendLevelPolicyBindingForGroupDto(appendLevelPolicyBindingForGroupDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerAppendParticularGroupBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name           | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType**  | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**    | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |
| **policyUuid** | **string**          | The ID of the required policy.                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| **groupUuid**  | **string**          | The ID of the required user group.                                                                                                                                                                                                                                                                                                                                                                                                                                                            |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerAppendParticularGroupBindingRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**appendLevelPolicyBindingForGroupDto** | [**AppendLevelPolicyBindingForGroupDto**](AppendLevelPolicyBindingForGroupDto.md) | The JSON body of the request. Contains parameters and metadata |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerCreateLevelPolicy

> LevelPolicyDto PolicyControllerCreateLevelPolicy(ctx, levelType, levelId).CreateOrUpdateLevelPolicyRequestDto(createOrUpdateLevelPolicyRequestDto).Execute()

Creates a new policy

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * account: use the UUID of the account.  * environment: use the ID of the environment.
    createOrUpdateLevelPolicyRequestDto := *openapiclient.NewCreateOrUpdateLevelPolicyRequestDto("Name_example", "Description_example", []string{"Tags_example"}, "StatementQuery_example") // CreateOrUpdateLevelPolicyRequestDto | The JSON body of the request. Contains the configuration of a new policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyManagementAPI.PolicyControllerCreateLevelPolicy(context.Background(), levelType, levelId).CreateOrUpdateLevelPolicyRequestDto(createOrUpdateLevelPolicyRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerCreateLevelPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyControllerCreateLevelPolicy`: LevelPolicyDto
    fmt.Fprintf(os.Stdout, "Response from `PolicyManagementAPI.PolicyControllerCreateLevelPolicy`: %v\n", resp)
}
```

### Path Parameters

| Name          | Type                | Description                                                                                                                                                                                                                                                                                                                                                          | Notes |
| ------------- | ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                          |
| **levelType** | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;account&#x60;: An account policy applies to all environments of an account. _ &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**   | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ account: use the UUID of the account. _ environment: use the ID of the environment.                                                                                                                                                                                      |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerCreateLevelPolicyRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**createOrUpdateLevelPolicyRequestDto** | [**CreateOrUpdateLevelPolicyRequestDto**](CreateOrUpdateLevelPolicyRequestDto.md) | The JSON body of the request. Contains the configuration of a new policy. |

### Return type

[**LevelPolicyDto**](LevelPolicyDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerDeleteLevelPolicy

> PolicyControllerDeleteLevelPolicy(ctx, levelType, levelId, policyUuid, force).Execute()

Deletes a policy

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the required policy.
    force := true // bool | Set to `true` to delete a policy that is still in use.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerDeleteLevelPolicy(context.Background(), levelType, levelId, policyUuid, force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerDeleteLevelPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name           | Type                | Description                                                                                                                                                                                                                                                                                                                                                          | Notes |
| -------------- | ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                          |
| **levelType**  | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;account&#x60;: An account policy applies to all environments of an account. _ &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**    | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ account: use the UUID of the account. _ environment: use the ID of the environment.                                                                                                                                                                                      |
| **policyUuid** | **string**          | The ID of the required policy.                                                                                                                                                                                                                                                                                                                                       |
| **force**      | **bool**            | Set to &#x60;true&#x60; to delete a policy that is still in use.                                                                                                                                                                                                                                                                                                     |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerDeleteLevelPolicyRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerDeleteLevelPolicyBindings

> PolicyControllerDeleteLevelPolicyBindings(ctx, levelType, levelId).Execute()

Deletes all policy bindings from a level

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerDeleteLevelPolicyBindings(context.Background(), levelType, levelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerDeleteLevelPolicyBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name          | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType** | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**   | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerDeleteLevelPolicyBindingsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerDeleteLevelPolicyBindingsForPolicy

> PolicyControllerDeleteLevelPolicyBindingsForPolicy(ctx, levelType, levelId, policyUuid).ForceMultiple(forceMultiple).Execute()

Deletes all bindings of a policy

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the required policy.
    forceMultiple := true // bool | Forces multiple in case delete by parameters and metadata query

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerDeleteLevelPolicyBindingsForPolicy(context.Background(), levelType, levelId, policyUuid).ForceMultiple(forceMultiple).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerDeleteLevelPolicyBindingsForPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name           | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType**  | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**    | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |
| **policyUuid** | **string**          | The ID of the required policy.                                                                                                                                                                                                                                                                                                                                                                                                                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerDeleteLevelPolicyBindingsForPolicyRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**forceMultiple** | **bool** | Forces multiple in case delete by parameters and metadata query |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerDeleteLevelPolicyBindingsForPolicyAndGroup

> PolicyControllerDeleteLevelPolicyBindingsForPolicyAndGroup(ctx, levelType, levelId, policyUuid, groupUuid, forceMultiple).Execute()

Deletes a policy binding from a user group

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the required policy.
    groupUuid := "groupUuid_example" // string | The ID of the required user group.
    forceMultiple := true // bool | Forces multiple in case delete by parameters and metadata query

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerDeleteLevelPolicyBindingsForPolicyAndGroup(context.Background(), levelType, levelId, policyUuid, groupUuid, forceMultiple).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerDeleteLevelPolicyBindingsForPolicyAndGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name              | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| ----------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**           | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType**     | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**       | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |
| **policyUuid**    | **string**          | The ID of the required policy.                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| **groupUuid**     | **string**          | The ID of the required user group.                                                                                                                                                                                                                                                                                                                                                                                                                                                            |
| **forceMultiple** | **bool**            | Forces multiple in case delete by parameters and metadata query                                                                                                                                                                                                                                                                                                                                                                                                                               |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerDeleteLevelPolicyBindingsForPolicyAndGroupRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerGetAllLevelPoliciesBindings

> LevelPolicyBindingDto PolicyControllerGetAllLevelPoliciesBindings(ctx, levelType, levelId).Execute()

Lists all policy bindings of a level

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyManagementAPI.PolicyControllerGetAllLevelPoliciesBindings(context.Background(), levelType, levelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerGetAllLevelPoliciesBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyControllerGetAllLevelPoliciesBindings`: LevelPolicyBindingDto
    fmt.Fprintf(os.Stdout, "Response from `PolicyManagementAPI.PolicyControllerGetAllLevelPoliciesBindings`: %v\n", resp)
}
```

### Path Parameters

| Name          | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType** | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**   | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerGetAllLevelPoliciesBindingsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**LevelPolicyBindingDto**](LevelPolicyBindingDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerGetEffectivePermissions

> EffectivePermissions PolicyControllerGetEffectivePermissions(ctx, levelType, levelId).EntityType(entityType).EntityId(entityId).Explain(explain).Services(services).Page(page).Size(size).Execute()

Gets effective permissions for a user or group | maturity=EARLY_ADOPTER

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * account: use the UUID of the account.  * environment: use the ID of the environment.
    entityType := "entityType_example" // string | Required entity type. The following values are available:   * user  * group
    entityId := "entityId_example" // string | Required entity id.
    explain := true // bool |
    services := []string{"Inner_example"} // []string | Optional services list. Policies for given services will be returned (optional)
    page := float32(8.14) // float32 |  (optional)
    size := float32(8.14) // float32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyManagementAPI.PolicyControllerGetEffectivePermissions(context.Background(), levelType, levelId).EntityType(entityType).EntityId(entityId).Explain(explain).Services(services).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerGetEffectivePermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyControllerGetEffectivePermissions`: EffectivePermissions
    fmt.Fprintf(os.Stdout, "Response from `PolicyManagementAPI.PolicyControllerGetEffectivePermissions`: %v\n", resp)
}
```

### Path Parameters

| Name          | Type                | Description                                                                                                                                                                                                                                                                                                                                                          | Notes |
| ------------- | ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                          |
| **levelType** | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;account&#x60;: An account policy applies to all environments of an account. _ &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**   | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ account: use the UUID of the account. _ environment: use the ID of the environment.                                                                                                                                                                                      |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerGetEffectivePermissionsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**entityType** | **string** | Required entity type. The following values are available: _ user _ group |
**entityId** | **string** | Required entity id. |
**explain** | **bool** | |
**services** | **[]string** | Optional services list. Policies for given services will be returned |
**page** | **float32** | |
**size** | **float32** | |

### Return type

[**EffectivePermissions**](EffectivePermissions.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerGetLevelDescendantsPolicyBindings

> PolicyControllerGetLevelDescendantsPolicyBindings(ctx, levelType, levelId, policyUuid).Page(page).Size(size).Execute()

Get policy bindings within descendants of a level

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `account`: An account policy applies to all environments of an account.   Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the required policy.
    page := float32(8.14) // float32 |  (optional)
    size := float32(8.14) // float32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerGetLevelDescendantsPolicyBindings(context.Background(), levelType, levelId, policyUuid).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerGetLevelDescendantsPolicyBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name           | Type                | Description                                                                                                                                                                                                                                                                       | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                       |
| **levelType**  | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: \* &#x60;account&#x60;: An account policy applies to all environments of an account. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**    | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                      |
| **policyUuid** | **string**          | The ID of the required policy.                                                                                                                                                                                                                                                    |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerGetLevelDescendantsPolicyBindingsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**page** | **float32** | |
**size** | **float32** | |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerGetLevelPolicies

> PolicyDtoList PolicyControllerGetLevelPolicies(ctx, levelType, levelId).Name(name).Execute()

Lists all native policies of a level

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    name := "name_example" // string | Optional policy name. Only policies that are of equal name will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyManagementAPI.PolicyControllerGetLevelPolicies(context.Background(), levelType, levelId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerGetLevelPolicies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyControllerGetLevelPolicies`: PolicyDtoList
    fmt.Fprintf(os.Stdout, "Response from `PolicyManagementAPI.PolicyControllerGetLevelPolicies`: %v\n", resp)
}
```

### Path Parameters

| Name          | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType** | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**   | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerGetLevelPoliciesRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**name** | **string** | Optional policy name. Only policies that are of equal name will be returned. |

### Return type

[**PolicyDtoList**](PolicyDtoList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerGetLevelPolicy

> LevelPolicyDto PolicyControllerGetLevelPolicy(ctx, levelType, levelId, policyUuid).Execute()

Gets a policy

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the required policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyManagementAPI.PolicyControllerGetLevelPolicy(context.Background(), levelType, levelId, policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerGetLevelPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyControllerGetLevelPolicy`: LevelPolicyDto
    fmt.Fprintf(os.Stdout, "Response from `PolicyManagementAPI.PolicyControllerGetLevelPolicy`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType**  | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**    | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |
| **policyUuid** | **string**          | The ID of the required policy.                                                                                                                                                                                                                                                                                                                                                                                                                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerGetLevelPolicyRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**LevelPolicyDto**](LevelPolicyDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerGetLevelPolicyBindings

> PolicyControllerGetLevelPolicyBindings(ctx, levelType, levelId, policyUuid).Execute()

Get policy bindings within a level

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the required policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerGetLevelPolicyBindings(context.Background(), levelType, levelId, policyUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerGetLevelPolicyBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name           | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType**  | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**    | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |
| **policyUuid** | **string**          | The ID of the required policy.                                                                                                                                                                                                                                                                                                                                                                                                                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerGetLevelPolicyBindingsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerGetLevelPolicyBindingsForGroup

> PolicyControllerGetLevelPolicyBindingsForGroup(ctx, levelType, levelId, policyUuid, groupUuid).Execute()

Get policy bindings within a level

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the required policy.
    groupUuid := "groupUuid_example" // string | The ID of the required user group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerGetLevelPolicyBindingsForGroup(context.Background(), levelType, levelId, policyUuid, groupUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerGetLevelPolicyBindingsForGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name           | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| -------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType**  | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**    | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |
| **policyUuid** | **string**          | The ID of the required policy.                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| **groupUuid**  | **string**          | The ID of the required user group.                                                                                                                                                                                                                                                                                                                                                                                                                                                            |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerGetLevelPolicyBindingsForGroupRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerGetPolicyOverviewList

> PolicyOverviewDtoList PolicyControllerGetPolicyOverviewList(ctx, levelType, levelId).Execute()

Lists all policies for a level, including inherited from higher levels

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyManagementAPI.PolicyControllerGetPolicyOverviewList(context.Background(), levelType, levelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerGetPolicyOverviewList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyControllerGetPolicyOverviewList`: PolicyOverviewDtoList
    fmt.Fprintf(os.Stdout, "Response from `PolicyManagementAPI.PolicyControllerGetPolicyOverviewList`: %v\n", resp)
}
```

### Path Parameters

| Name          | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType** | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**   | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerGetPolicyOverviewListRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**PolicyOverviewDtoList**](PolicyOverviewDtoList.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerGetPolicyUuidsBindings

> PolicyUuidsWithoutMetadataDto PolicyControllerGetPolicyUuidsBindings(ctx, levelType, levelId, groupUuid).Execute()

Lists all policies for a user group

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    groupUuid := "groupUuid_example" // string | The ID of the required user group.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyManagementAPI.PolicyControllerGetPolicyUuidsBindings(context.Background(), levelType, levelId, groupUuid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerGetPolicyUuidsBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyControllerGetPolicyUuidsBindings`: PolicyUuidsWithoutMetadataDto
    fmt.Fprintf(os.Stdout, "Response from `PolicyManagementAPI.PolicyControllerGetPolicyUuidsBindings`: %v\n", resp)
}
```

### Path Parameters

| Name          | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType** | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**   | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |
| **groupUuid** | **string**          | The ID of the required user group.                                                                                                                                                                                                                                                                                                                                                                                                                                                            |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerGetPolicyUuidsBindingsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

### Return type

[**PolicyUuidsWithoutMetadataDto**](PolicyUuidsWithoutMetadataDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerUpdateLevelPolicy

> LevelPolicyDto PolicyControllerUpdateLevelPolicy(ctx, levelType, levelId, policyUuid).CreateOrUpdateLevelPolicyRequestDto(createOrUpdateLevelPolicyRequestDto).Execute()

Updates a policy

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the required policy.
    createOrUpdateLevelPolicyRequestDto := *openapiclient.NewCreateOrUpdateLevelPolicyRequestDto("Name_example", "Description_example", []string{"Tags_example"}, "StatementQuery_example") // CreateOrUpdateLevelPolicyRequestDto | The JSON body of the request. Contains the updated configuration of a policy.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyManagementAPI.PolicyControllerUpdateLevelPolicy(context.Background(), levelType, levelId, policyUuid).CreateOrUpdateLevelPolicyRequestDto(createOrUpdateLevelPolicyRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerUpdateLevelPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyControllerUpdateLevelPolicy`: LevelPolicyDto
    fmt.Fprintf(os.Stdout, "Response from `PolicyManagementAPI.PolicyControllerUpdateLevelPolicy`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                                                                                                                                                                                                                                                                                                          | Notes |
| -------------- | ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                          |
| **levelType**  | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;account&#x60;: An account policy applies to all environments of an account. _ &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**    | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ account: use the UUID of the account. _ environment: use the ID of the environment.                                                                                                                                                                                      |
| **policyUuid** | **string**          | The ID of the required policy.                                                                                                                                                                                                                                                                                                                                       |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerUpdateLevelPolicyRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**createOrUpdateLevelPolicyRequestDto** | [**CreateOrUpdateLevelPolicyRequestDto**](CreateOrUpdateLevelPolicyRequestDto.md) | The JSON body of the request. Contains the updated configuration of a policy. |

### Return type

[**LevelPolicyDto**](LevelPolicyDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerUpdateLevelPolicyBindings

> PolicyControllerUpdateLevelPolicyBindings(ctx, levelType, levelId).CreateLevelPolicyBindingsRequestDto(createLevelPolicyBindingsRequestDto).Execute()

Updates policy bindings of a level

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    createLevelPolicyBindingsRequestDto := *openapiclient.NewCreateLevelPolicyBindingsRequestDto([]openapiclient.Binding{*openapiclient.NewBinding("PolicyUuid_example", []string{"Groups_example"})}) // CreateLevelPolicyBindingsRequestDto | The JSON body of the request. Contains new policy bindings of a level.    Any existing binding not presented in the request is discarded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerUpdateLevelPolicyBindings(context.Background(), levelType, levelId).CreateLevelPolicyBindingsRequestDto(createLevelPolicyBindingsRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerUpdateLevelPolicyBindings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name          | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType** | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**   | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerUpdateLevelPolicyBindingsRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**createLevelPolicyBindingsRequestDto** | [**CreateLevelPolicyBindingsRequestDto**](CreateLevelPolicyBindingsRequestDto.md) | The JSON body of the request. Contains new policy bindings of a level. Any existing binding not presented in the request is discarded. |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerUpdatePolicyBindingsToGroup

> PolicyControllerUpdatePolicyBindingsToGroup(ctx, levelType, levelId, groupUuid).PolicyUuidsDto(policyUuidsDto).Execute()

Updates policies for a user group

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `global`: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace.  * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * global: use the `global` value.  * account: use the UUID of the account.  * environment: use the ID of the environment.
    groupUuid := "groupUuid_example" // string | The ID of the required user group.
    policyUuidsDto := *openapiclient.NewPolicyUuidsDto([]string{"PolicyUuids_example"}) // PolicyUuidsDto | The JSON body of the request. Contains new policies for the group.    Any policy not presented in the request is discarded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyManagementAPI.PolicyControllerUpdatePolicyBindingsToGroup(context.Background(), levelType, levelId, groupUuid).PolicyUuidsDto(policyUuidsDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerUpdatePolicyBindingsToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

| Name          | Type                | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Notes |
| ------------- | ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                                                                                                                                                   |
| **levelType** | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;global&#x60;: A global policy applies to all accounts and environments. It is defined and managed by Dynatrace. _ &#x60;account&#x60;: An account policy applies to all environments of an account. \* &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**   | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ global: use the &#x60;global&#x60; value. _ account: use the UUID of the account. \* environment: use the ID of the environment.                                                                                                                                                                                                                                                                  |
| **groupUuid** | **string**          | The ID of the required user group.                                                                                                                                                                                                                                                                                                                                                                                                                                                            |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerUpdatePolicyBindingsToGroupRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**policyUuidsDto** | [**PolicyUuidsDto**](PolicyUuidsDto.md) | The JSON body of the request. Contains new policies for the group. Any policy not presented in the request is discarded. |

### Return type

(empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerValidateLevelPolicy

> ValidationDto PolicyControllerValidateLevelPolicy(ctx, levelType, levelId, policyUuid).CreateOrUpdateLevelPolicyRequestDto(createOrUpdateLevelPolicyRequestDto).Execute()

Validates the payload for the `PUT /iam/v1/repo/{levelType}/{levelId}/policies/{policyUuid}` request

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * account: use the UUID of the account.  * environment: use the ID of the environment.
    policyUuid := "policyUuid_example" // string | The ID of the policy to be validated.
    createOrUpdateLevelPolicyRequestDto := *openapiclient.NewCreateOrUpdateLevelPolicyRequestDto("Name_example", "Description_example", []string{"Tags_example"}, "StatementQuery_example") // CreateOrUpdateLevelPolicyRequestDto | The JSON body of the request. Contains the configuration of a policy to be validated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyManagementAPI.PolicyControllerValidateLevelPolicy(context.Background(), levelType, levelId, policyUuid).CreateOrUpdateLevelPolicyRequestDto(createOrUpdateLevelPolicyRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerValidateLevelPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyControllerValidateLevelPolicy`: ValidationDto
    fmt.Fprintf(os.Stdout, "Response from `PolicyManagementAPI.PolicyControllerValidateLevelPolicy`: %v\n", resp)
}
```

### Path Parameters

| Name           | Type                | Description                                                                                                                                                                                                                                                                                                                                                          | Notes |
| -------------- | ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**        | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                          |
| **levelType**  | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;account&#x60;: An account policy applies to all environments of an account. _ &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**    | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ account: use the UUID of the account. _ environment: use the ID of the environment.                                                                                                                                                                                      |
| **policyUuid** | **string**          | The ID of the policy to be validated.                                                                                                                                                                                                                                                                                                                                |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerValidateLevelPolicyRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**createOrUpdateLevelPolicyRequestDto** | [**CreateOrUpdateLevelPolicyRequestDto**](CreateOrUpdateLevelPolicyRequestDto.md) | The JSON body of the request. Contains the configuration of a policy to be validated. |

### Return type

[**ValidationDto**](ValidationDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PolicyControllerValidateNewLevelPolicy

> ValidationDto PolicyControllerValidateNewLevelPolicy(ctx, levelType, levelId).CreateOrUpdateLevelPolicyRequestDto(createOrUpdateLevelPolicyRequestDto).Execute()

Validates the payload for the `POST /iam/v1/repo/{levelType}/{levelId}/policies` request

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
    levelType := "levelType_example" // string | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available:   * `account`: An account policy applies to all environments of an account.  * `environment`: An environment policy applies to a specific environment.    Each level inherits the policies of the higher level and extends them with its own policies.
    levelId := "levelId_example" // string | The ID of the policy level. Use one of the following values, depending on the level type:   * account: use the UUID of the account.  * environment: use the ID of the environment.
    createOrUpdateLevelPolicyRequestDto := *openapiclient.NewCreateOrUpdateLevelPolicyRequestDto("Name_example", "Description_example", []string{"Tags_example"}, "StatementQuery_example") // CreateOrUpdateLevelPolicyRequestDto | The JSON body of the request. Contains the configuration of a policy to be validated.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyManagementAPI.PolicyControllerValidateNewLevelPolicy(context.Background(), levelType, levelId).CreateOrUpdateLevelPolicyRequestDto(createOrUpdateLevelPolicyRequestDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyManagementAPI.PolicyControllerValidateNewLevelPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PolicyControllerValidateNewLevelPolicy`: ValidationDto
    fmt.Fprintf(os.Stdout, "Response from `PolicyManagementAPI.PolicyControllerValidateNewLevelPolicy`: %v\n", resp)
}
```

### Path Parameters

| Name          | Type                | Description                                                                                                                                                                                                                                                                                                                                                          | Notes |
| ------------- | ------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----- |
| **ctx**       | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.                                                                                                                                                                                                                                                                                          |
| **levelType** | **string**          | The type of the [policy](https://dt-url.net/eu03uap) level. The following values are available: _ &#x60;account&#x60;: An account policy applies to all environments of an account. _ &#x60;environment&#x60;: An environment policy applies to a specific environment. Each level inherits the policies of the higher level and extends them with its own policies. |
| **levelId**   | **string**          | The ID of the policy level. Use one of the following values, depending on the level type: _ account: use the UUID of the account. _ environment: use the ID of the environment.                                                                                                                                                                                      |

### Other Parameters

Other parameters are passed through a pointer to a apiPolicyControllerValidateNewLevelPolicyRequest struct via the builder pattern

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |

**createOrUpdateLevelPolicyRequestDto** | [**CreateOrUpdateLevelPolicyRequestDto**](CreateOrUpdateLevelPolicyRequestDto.md) | The JSON body of the request. Contains the configuration of a policy to be validated. |

### Return type

[**ValidationDto**](ValidationDto.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
