# \AWSCredentialsConfigurationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAwsCredentialsConfig**](AWSCredentialsConfigurationAPI.md#CreateAwsCredentialsConfig) | **Post** /aws/credentials | Creates a new AWS credentials configuration
[**DeleteAwsCredentialsConfig**](AWSCredentialsConfigurationAPI.md#DeleteAwsCredentialsConfig) | **Delete** /aws/credentials/{id} | Deletes the specified AWS credentials configuration
[**GetAwsCredentialsConfig**](AWSCredentialsConfigurationAPI.md#GetAwsCredentialsConfig) | **Get** /aws/credentials/{id} | Gets the configuration of the specified AWS credentials
[**GetAwsServicesConfig**](AWSCredentialsConfigurationAPI.md#GetAwsServicesConfig) | **Get** /aws/credentials/{id}/services | Gets the monitored services configuration of the specified AWS credentials
[**GetSupportedServices**](AWSCredentialsConfigurationAPI.md#GetSupportedServices) | **Get** /aws/supportedServices | Gets the list of AWS supported services
[**ListAwsCredentialConfigs**](AWSCredentialsConfigurationAPI.md#ListAwsCredentialConfigs) | **Get** /aws/credentials | Lists all available AWS credentials configurations
[**ReadIamExternalIdToken**](AWSCredentialsConfigurationAPI.md#ReadIamExternalIdToken) | **Get** /aws/iamExternalId | Gets the external ID token for setting an IAM role
[**UpdateAwsCredentialsConfig**](AWSCredentialsConfigurationAPI.md#UpdateAwsCredentialsConfig) | **Put** /aws/credentials/{id} | Updates an existing AWS credentials configuration
[**UpdateAwsServicesConfig**](AWSCredentialsConfigurationAPI.md#UpdateAwsServicesConfig) | **Put** /aws/credentials/{id}/services | Replace an existing monitored services configuration of the specified AWS credentials
[**ValidateCreateAwsCredentialsConfig**](AWSCredentialsConfigurationAPI.md#ValidateCreateAwsCredentialsConfig) | **Post** /aws/credentials/validator | Validates the payload for the &#x60;POST /aws/credentials&#x60; request
[**ValidateUpdateAwsCredentialsConfig**](AWSCredentialsConfigurationAPI.md#ValidateUpdateAwsCredentialsConfig) | **Post** /aws/credentials/{id}/validator | Validates the payload for the &#x60;PUT /aws/credentials/{id}&#x60; request
[**ValidateUpdateAwsServicesConfig**](AWSCredentialsConfigurationAPI.md#ValidateUpdateAwsServicesConfig) | **Post** /aws/credentials/{id}/services/validator | Validates the payload for the &#x60;PUT /aws/credentials/{id}/services&#x60; request



## CreateAwsCredentialsConfig

> EntityShortRepresentation CreateAwsCredentialsConfig(ctx).AwsCredentialsConfig(awsCredentialsConfig).Execute()

Creates a new AWS credentials configuration



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
    awsCredentialsConfig := *openapiclient.NewAwsCredentialsConfig(*openapiclient.NewAwsAuthenticationData("Type_example"), "Label_example", "PartitionType_example", false, []openapiclient.AwsConfigTag{*openapiclient.NewAwsConfigTag("Name_example", "Value_example")}) // AwsCredentialsConfig | The JSON body of the request. Contains parameters of the new AWS credentials configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCredentialsConfigurationAPI.CreateAwsCredentialsConfig(context.Background()).AwsCredentialsConfig(awsCredentialsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.CreateAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAwsCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationAPI.CreateAwsCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCredentialsConfig** | [**AwsCredentialsConfig**](AwsCredentialsConfig.md) | The JSON body of the request. Contains parameters of the new AWS credentials configuration. | 

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


## DeleteAwsCredentialsConfig

> DeleteAwsCredentialsConfig(ctx, id).Execute()

Deletes the specified AWS credentials configuration

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
    id := "id_example" // string | The ID of AWS credentials configuration to be deleted.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AWSCredentialsConfigurationAPI.DeleteAwsCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.DeleteAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of AWS credentials configuration to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAwsCredentialsConfig

> AwsCredentialsConfig GetAwsCredentialsConfig(ctx, id).Execute()

Gets the configuration of the specified AWS credentials

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
    id := "id_example" // string | The ID of the specified AWS credentials configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCredentialsConfigurationAPI.GetAwsCredentialsConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.GetAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAwsCredentialsConfig`: AwsCredentialsConfig
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationAPI.GetAwsCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the specified AWS credentials configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsCredentialsConfig**](AwsCredentialsConfig.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAwsServicesConfig

> AwsMonitoredServicesDto GetAwsServicesConfig(ctx, id).Execute()

Gets the monitored services configuration of the specified AWS credentials

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
    id := "id_example" // string | The ID of the specified AWS credentials configuration.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCredentialsConfigurationAPI.GetAwsServicesConfig(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.GetAwsServicesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAwsServicesConfig`: AwsMonitoredServicesDto
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationAPI.GetAwsServicesConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the specified AWS credentials configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAwsServicesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AwsMonitoredServicesDto**](AwsMonitoredServicesDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSupportedServices

> CloudSupportedServicesList GetSupportedServices(ctx).Execute()

Gets the list of AWS supported services



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
    resp, r, err := apiClient.AWSCredentialsConfigurationAPI.GetSupportedServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.GetSupportedServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSupportedServices`: CloudSupportedServicesList
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationAPI.GetSupportedServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSupportedServicesRequest struct via the builder pattern


### Return type

[**CloudSupportedServicesList**](CloudSupportedServicesList.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAwsCredentialConfigs

> []EntityShortRepresentation ListAwsCredentialConfigs(ctx).Execute()

Lists all available AWS credentials configurations

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
    resp, r, err := apiClient.AWSCredentialsConfigurationAPI.ListAwsCredentialConfigs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.ListAwsCredentialConfigs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAwsCredentialConfigs`: []EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationAPI.ListAwsCredentialConfigs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAwsCredentialConfigsRequest struct via the builder pattern


### Return type

[**[]EntityShortRepresentation**](EntityShortRepresentation.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadIamExternalIdToken

> AwsIamToken ReadIamExternalIdToken(ctx).Execute()

Gets the external ID token for setting an IAM role



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
    resp, r, err := apiClient.AWSCredentialsConfigurationAPI.ReadIamExternalIdToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.ReadIamExternalIdToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadIamExternalIdToken`: AwsIamToken
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationAPI.ReadIamExternalIdToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReadIamExternalIdTokenRequest struct via the builder pattern


### Return type

[**AwsIamToken**](AwsIamToken.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAwsCredentialsConfig

> EntityShortRepresentation UpdateAwsCredentialsConfig(ctx, id).AwsCredentialsConfig(awsCredentialsConfig).Execute()

Updates an existing AWS credentials configuration

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
    id := "id_example" // string | The ID of the AWS credentials configuration to be updated.
    awsCredentialsConfig := *openapiclient.NewAwsCredentialsConfig(*openapiclient.NewAwsAuthenticationData("Type_example"), "Label_example", "PartitionType_example", false, []openapiclient.AwsConfigTag{*openapiclient.NewAwsConfigTag("Name_example", "Value_example")}) // AwsCredentialsConfig | The JSON body of the request. Contains updated parameters of the AWS credentials configuration. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AWSCredentialsConfigurationAPI.UpdateAwsCredentialsConfig(context.Background(), id).AwsCredentialsConfig(awsCredentialsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.UpdateAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAwsCredentialsConfig`: EntityShortRepresentation
    fmt.Fprintf(os.Stdout, "Response from `AWSCredentialsConfigurationAPI.UpdateAwsCredentialsConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the AWS credentials configuration to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsCredentialsConfig** | [**AwsCredentialsConfig**](AwsCredentialsConfig.md) | The JSON body of the request. Contains updated parameters of the AWS credentials configuration. | 

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


## UpdateAwsServicesConfig

> UpdateAwsServicesConfig(ctx, id).AwsMonitoredServicesDto(awsMonitoredServicesDto).Execute()

Replace an existing monitored services configuration of the specified AWS credentials

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
    id := "id_example" // string | The ID of the AWS credentials to be updated with new monitored services configuration.
    awsMonitoredServicesDto := *openapiclient.NewAwsMonitoredServicesDto([]openapiclient.AwsSupportingServiceConfig{*openapiclient.NewAwsSupportingServiceConfig("Name_example")}) // AwsMonitoredServicesDto | The JSON body of the request. Contains updated monitored services configuration for AWS credentials. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AWSCredentialsConfigurationAPI.UpdateAwsServicesConfig(context.Background(), id).AwsMonitoredServicesDto(awsMonitoredServicesDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.UpdateAwsServicesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the AWS credentials to be updated with new monitored services configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAwsServicesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsMonitoredServicesDto** | [**AwsMonitoredServicesDto**](AwsMonitoredServicesDto.md) | The JSON body of the request. Contains updated monitored services configuration for AWS credentials. | 

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


## ValidateCreateAwsCredentialsConfig

> ValidateCreateAwsCredentialsConfig(ctx).AwsCredentialsConfig(awsCredentialsConfig).Execute()

Validates the payload for the `POST /aws/credentials` request

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
    awsCredentialsConfig := *openapiclient.NewAwsCredentialsConfig(*openapiclient.NewAwsAuthenticationData("Type_example"), "Label_example", "PartitionType_example", false, []openapiclient.AwsConfigTag{*openapiclient.NewAwsConfigTag("Name_example", "Value_example")}) // AwsCredentialsConfig | The JSON body of the request. Contains the AWS credentials configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AWSCredentialsConfigurationAPI.ValidateCreateAwsCredentialsConfig(context.Background()).AwsCredentialsConfig(awsCredentialsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.ValidateCreateAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **awsCredentialsConfig** | [**AwsCredentialsConfig**](AwsCredentialsConfig.md) | The JSON body of the request. Contains the AWS credentials configuration to be validated. | 

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


## ValidateUpdateAwsCredentialsConfig

> ValidateUpdateAwsCredentialsConfig(ctx, id).AwsCredentialsConfig(awsCredentialsConfig).Execute()

Validates the payload for the `PUT /aws/credentials/{id}` request

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
    id := "id_example" // string | The ID of the AWS credentials configuration to be validated.
    awsCredentialsConfig := *openapiclient.NewAwsCredentialsConfig(*openapiclient.NewAwsAuthenticationData("Type_example"), "Label_example", "PartitionType_example", false, []openapiclient.AwsConfigTag{*openapiclient.NewAwsConfigTag("Name_example", "Value_example")}) // AwsCredentialsConfig | The JSON body of the request. Contains the AWS credentials configuration to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AWSCredentialsConfigurationAPI.ValidateUpdateAwsCredentialsConfig(context.Background(), id).AwsCredentialsConfig(awsCredentialsConfig).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.ValidateUpdateAwsCredentialsConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the AWS credentials configuration to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateAwsCredentialsConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsCredentialsConfig** | [**AwsCredentialsConfig**](AwsCredentialsConfig.md) | The JSON body of the request. Contains the AWS credentials configuration to be validated. | 

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


## ValidateUpdateAwsServicesConfig

> ValidateUpdateAwsServicesConfig(ctx, id).AwsMonitoredServicesDto(awsMonitoredServicesDto).Execute()

Validates the payload for the `PUT /aws/credentials/{id}/services` request

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
    id := "id_example" // string | The ID of the AWS credentials for which the monitored services configuration is going to be validated.
    awsMonitoredServicesDto := *openapiclient.NewAwsMonitoredServicesDto([]openapiclient.AwsSupportingServiceConfig{*openapiclient.NewAwsSupportingServiceConfig("Name_example")}) // AwsMonitoredServicesDto | The JSON body of the request. Contains a monitored services configuration for AWS credentials to be validated. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AWSCredentialsConfigurationAPI.ValidateUpdateAwsServicesConfig(context.Background(), id).AwsMonitoredServicesDto(awsMonitoredServicesDto).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AWSCredentialsConfigurationAPI.ValidateUpdateAwsServicesConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the AWS credentials for which the monitored services configuration is going to be validated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateAwsServicesConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **awsMonitoredServicesDto** | [**AwsMonitoredServicesDto**](AwsMonitoredServicesDto.md) | The JSON body of the request. Contains a monitored services configuration for AWS credentials to be validated. | 

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

