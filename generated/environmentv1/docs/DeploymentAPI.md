# \DeploymentAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadAgentInstallerWithVersion**](DeploymentAPI.md#DownloadAgentInstallerWithVersion) | **Get** /deployment/installer/agent/{osType}/{installerType}/version/{version} | Downloads OneAgent installer of the specified version
[**DownloadAgentOrchestrationSignatureWithVersion**](DeploymentAPI.md#DownloadAgentOrchestrationSignatureWithVersion) | **Get** /deployment/orchestration/agent/{orchestrationType}/version/{version}/signature | Downloads the requested version matching OneAgent deployment orchestration tarball&#39;s signature
[**DownloadAgentOrchestrationWithVersion**](DeploymentAPI.md#DownloadAgentOrchestrationWithVersion) | **Get** /deployment/orchestration/agent/{orchestrationType}/version/{version} | Downloads the requested version matching OneAgent deployment orchestration tarball
[**DownloadBoshReleaseWithVersion**](DeploymentAPI.md#DownloadBoshReleaseWithVersion) | **Get** /deployment/boshrelease/agent/{osType}/version/{version} | Downloads BOSH release tarballs of the specified version, OneAgent included
[**DownloadGatewayInstallerWithVersion**](DeploymentAPI.md#DownloadGatewayInstallerWithVersion) | **Get** /deployment/installer/gateway/{osType}/version/{version} | Downloads the ActiveGate installer of the specified version
[**DownloadLatestAgentInstaller**](DeploymentAPI.md#DownloadLatestAgentInstaller) | **Get** /deployment/installer/agent/{osType}/{installerType}/latest | Downloads the latest OneAgent installer
[**DownloadLatestAgentOrchestration**](DeploymentAPI.md#DownloadLatestAgentOrchestration) | **Get** /deployment/orchestration/agent/{orchestrationType}/latest | Downloads the latest OneAgent deployment orchestration tarball
[**DownloadLatestAgentOrchestrationSignature**](DeploymentAPI.md#DownloadLatestAgentOrchestrationSignature) | **Get** /deployment/orchestration/agent/{orchestrationType}/latest/signature | Downloads the latest OneAgent deployment orchestration tarball&#39;s signature
[**DownloadLatestGatewayInstaller**](DeploymentAPI.md#DownloadLatestGatewayInstaller) | **Get** /deployment/installer/gateway/{osType}/latest | Downloads the configured standard ActiveGate installer of the latest version for the specified OS
[**GetActiveGateInstallerAvailableVersions**](DeploymentAPI.md#GetActiveGateInstallerAvailableVersions) | **Get** /deployment/installer/gateway/versions/{osType} | Lists all available versions of ActiveGate installer
[**GetActiveGateInstallerConnectionInfo**](DeploymentAPI.md#GetActiveGateInstallerConnectionInfo) | **Get** /deployment/installer/gateway/connectioninfo | Gets the connectivity information for Environment ActiveGate
[**GetAgentInstallerAvailableVersions**](DeploymentAPI.md#GetAgentInstallerAvailableVersions) | **Get** /deployment/installer/agent/versions/{osType}/{installerType} | Lists all available versions of OneAgent installer
[**GetAgentInstallerConnectionInfo**](DeploymentAPI.md#GetAgentInstallerConnectionInfo) | **Get** /deployment/installer/agent/connectioninfo | Gets the connectivity information for OneAgent
[**GetAgentInstallerConnectionInfoEndpoints**](DeploymentAPI.md#GetAgentInstallerConnectionInfoEndpoints) | **Get** /deployment/installer/agent/connectioninfo/endpoints | Gets the list of the ActiveGate-Endpoints to be used for Agents ordered by networkzone-priorities.
[**GetAgentInstallerMetaInfo**](DeploymentAPI.md#GetAgentInstallerMetaInfo) | **Get** /deployment/installer/agent/{osType}/{installerType}/latest/metainfo | Gets the latest available version of a OneAgent installer
[**GetAgentInstallerWithVersionChecksum**](DeploymentAPI.md#GetAgentInstallerWithVersionChecksum) | **Get** /deployment/installer/agent/{osType}/{installerType}/version/{version}/checksum | Gets the checksum of a non-customized OneAgent installer
[**GetAgentProcessModuleConfig**](DeploymentAPI.md#GetAgentProcessModuleConfig) | **Get** /deployment/installer/agent/processmoduleconfig | Gets the latest process module config | maturity&#x3D;EARLY_ADOPTER
[**GetBoshReleaseAvailableVersions**](DeploymentAPI.md#GetBoshReleaseAvailableVersions) | **Get** /deployment/boshrelease/versions/{osType} | Gets the list of available OneAgent versions for BOSH release tarballs
[**GetBoshReleaseChecksum**](DeploymentAPI.md#GetBoshReleaseChecksum) | **Get** /deployment/boshrelease/agent/{osType}/version/{version}/checksum | Gets the checksum of the specified BOSH release tarball
[**GetLatestLambdaBuildUnits**](DeploymentAPI.md#GetLatestLambdaBuildUnits) | **Get** /deployment/lambda/agent/latest | Get the latest version names of the OneAgent for AWS Lambda



## DownloadAgentInstallerWithVersion

> DownloadAgentInstallerWithVersion(ctx, osType, installerType, version).IfNoneMatch(ifNoneMatch).Flavor(flavor).Arch(arch).Bitness(bitness).Include(include).SkipMetadata(skipMetadata).NetworkZone(networkZone).Execute()

Downloads OneAgent installer of the specified version



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
    osType := "osType_example" // string | The operating system of the installer.
    installerType := "installerType_example" // string | The type of the installer:   * `default`: Self-extracting installer for manual installation. Downloads an `.exe` file for Windows or an `.sh` file for Unix.  * `default-unattended`: Self-extracting installer for unattended installation. Windows only. Downloads a `.zip` archive, containing the `.msi` installer and the batch file. This option is deprecated with OneAgent version 1.173  * `mainframe`: Downloads all code modules for z/OS combined in a single `*.pax` archive.  * `paas`: Code modules installer. Downloads a `*.zip` archive, containing the `manifest.json` file with meta information or a `.jar` file for z/OS.  * `paas-sh`: Code modules installer. Downloads a self-extracting shell script with the embedded `tar.gz` archive.
    version := "version_example" // string | The required version of the OneAgent in `1.155.275.20181112-084458` format.   You can retrieve the list of available versions with the [**GET available versions of OneAgent**](https://dt-url.net/fo23rb5) call.
    ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. Do not download if it matches the ETag of the installer. (optional)
    flavor := "flavor_example" // string | The flavor of your Linux distribution:   * `musl` for Linux distributions, which are using the musl C standard library, for example Alpine Linux.  * 'multidistro` for all Linux distributions which are using musl C and glibc standard library.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "default")
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Use this value for AIX and z/OS. Defaults to `x86` for other OS types.  * `x86`: x86 architecture. * `ppc`: PowerPC architecture, only supported for AIX and Linux. * `ppcle`: PowerPC Little Endian architecture, only supported for Linux. * `sparc`: Sparc architecture, only supported for Solaris.   * `arm`: ARM architecture, only supported for Linux.   * `s390`: S/390 architecture, only supported for Linux.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    bitness := "bitness_example" // string | The bitness of your OS. Must be supported by the OS.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    include := []string{"Include_example"} // []string | The code modules to be included to the installer. You can specify several modules in the following format: `include=java&include=dotnet`.   Only applicable to the `paas` and `paas-sh` installer types. (optional)
    skipMetadata := true // bool | Set `true` to omit the OneAgent connectivity information from the installer.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to false)
    networkZone := "networkZone_example" // string | The network zone you want the result to be configured with. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.DownloadAgentInstallerWithVersion(context.Background(), osType, installerType, version).IfNoneMatch(ifNoneMatch).Flavor(flavor).Arch(arch).Bitness(bitness).Include(include).SkipMetadata(skipMetadata).NetworkZone(networkZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DownloadAgentInstallerWithVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**installerType** | **string** | The type of the installer:   * &#x60;default&#x60;: Self-extracting installer for manual installation. Downloads an &#x60;.exe&#x60; file for Windows or an &#x60;.sh&#x60; file for Unix.  * &#x60;default-unattended&#x60;: Self-extracting installer for unattended installation. Windows only. Downloads a &#x60;.zip&#x60; archive, containing the &#x60;.msi&#x60; installer and the batch file. This option is deprecated with OneAgent version 1.173  * &#x60;mainframe&#x60;: Downloads all code modules for z/OS combined in a single &#x60;*.pax&#x60; archive.  * &#x60;paas&#x60;: Code modules installer. Downloads a &#x60;*.zip&#x60; archive, containing the &#x60;manifest.json&#x60; file with meta information or a &#x60;.jar&#x60; file for z/OS.  * &#x60;paas-sh&#x60;: Code modules installer. Downloads a self-extracting shell script with the embedded &#x60;tar.gz&#x60; archive. | 
**version** | **string** | The required version of the OneAgent in &#x60;1.155.275.20181112-084458&#x60; format.   You can retrieve the list of available versions with the [**GET available versions of OneAgent**](https://dt-url.net/fo23rb5) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAgentInstallerWithVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifNoneMatch** | **string** | The ETag of the previous request. Do not download if it matches the ETag of the installer. | 
 **flavor** | **string** | The flavor of your Linux distribution:   * &#x60;musl&#x60; for Linux distributions, which are using the musl C standard library, for example Alpine Linux.  * &#39;multidistro&#x60; for all Linux distributions which are using musl C and glibc standard library.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;default&quot;]
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Use this value for AIX and z/OS. Defaults to &#x60;x86&#x60; for other OS types.  * &#x60;x86&#x60;: x86 architecture. * &#x60;ppc&#x60;: PowerPC architecture, only supported for AIX and Linux. * &#x60;ppcle&#x60;: PowerPC Little Endian architecture, only supported for Linux. * &#x60;sparc&#x60;: Sparc architecture, only supported for Solaris.   * &#x60;arm&#x60;: ARM architecture, only supported for Linux.   * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **bitness** | **string** | The bitness of your OS. Must be supported by the OS.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **include** | **[]string** | The code modules to be included to the installer. You can specify several modules in the following format: &#x60;include&#x3D;java&amp;include&#x3D;dotnet&#x60;.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | 
 **skipMetadata** | **bool** | Set &#x60;true&#x60; to omit the OneAgent connectivity information from the installer.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to false]
 **networkZone** | **string** | The network zone you want the result to be configured with. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAgentOrchestrationSignatureWithVersion

> DownloadAgentOrchestrationSignatureWithVersion(ctx, orchestrationType, version).Execute()

Downloads the requested version matching OneAgent deployment orchestration tarball's signature



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
    orchestrationType := "orchestrationType_example" // string | The Orchestration Type of the orchestration deployment script.
    version := "version_example" // string | The requested version of the OneAgent deployment orchestration tarball in `0.1.0.20200925-120822` format.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.DownloadAgentOrchestrationSignatureWithVersion(context.Background(), orchestrationType, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DownloadAgentOrchestrationSignatureWithVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orchestrationType** | **string** | The Orchestration Type of the orchestration deployment script. | 
**version** | **string** | The requested version of the OneAgent deployment orchestration tarball in &#x60;0.1.0.20200925-120822&#x60; format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAgentOrchestrationSignatureWithVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadAgentOrchestrationWithVersion

> DownloadAgentOrchestrationWithVersion(ctx, orchestrationType, version).Execute()

Downloads the requested version matching OneAgent deployment orchestration tarball



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
    orchestrationType := "orchestrationType_example" // string | The Orchestration Type of the orchestration deployment script.
    version := "version_example" // string | The requested version of the OneAgent orchestration deployment tarball in `0.1.0.20200925-120822` format.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.DownloadAgentOrchestrationWithVersion(context.Background(), orchestrationType, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DownloadAgentOrchestrationWithVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orchestrationType** | **string** | The Orchestration Type of the orchestration deployment script. | 
**version** | **string** | The requested version of the OneAgent orchestration deployment tarball in &#x60;0.1.0.20200925-120822&#x60; format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadAgentOrchestrationWithVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadBoshReleaseWithVersion

> DownloadBoshReleaseWithVersion(ctx, osType, version).SkipMetadata(skipMetadata).NetworkZone(networkZone).Execute()

Downloads BOSH release tarballs of the specified version, OneAgent included



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
    osType := "osType_example" // string | The operating system of the installer.
    version := "version_example" // string | The required version of the OneAgent in the `1.155.275.20181112-084458` format.   You can retrieve the list of available versions with the [**GET available versions of BOSH tarballs**](https://dt-url.net/j703kdn) call.
    skipMetadata := true // bool | Set `true` to omit the OneAgent connectivity information from the installer.    If not set, `false` is used. (optional) (default to false)
    networkZone := "networkZone_example" // string | The network zone you want the result to be configured with. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.DownloadBoshReleaseWithVersion(context.Background(), osType, version).SkipMetadata(skipMetadata).NetworkZone(networkZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DownloadBoshReleaseWithVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**version** | **string** | The required version of the OneAgent in the &#x60;1.155.275.20181112-084458&#x60; format.   You can retrieve the list of available versions with the [**GET available versions of BOSH tarballs**](https://dt-url.net/j703kdn) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadBoshReleaseWithVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipMetadata** | **bool** | Set &#x60;true&#x60; to omit the OneAgent connectivity information from the installer.    If not set, &#x60;false&#x60; is used. | [default to false]
 **networkZone** | **string** | The network zone you want the result to be configured with. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadGatewayInstallerWithVersion

> DownloadGatewayInstallerWithVersion(ctx, osType, version).IfNoneMatch(ifNoneMatch).NetworkZone(networkZone).Arch(arch).Execute()

Downloads the ActiveGate installer of the specified version

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
    osType := "osType_example" // string | The operating system of the installer.
    version := "version_example" // string | The required version of the ActiveGate installer, in `1.155.275.20181112-084458` format.   You can retrieve the list of available versions with the [**GET available versions of ActiveGate**](https://dt-url.net/kh43rha) call.
    ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. Do not download if it matches the ETag of the installer. (optional)
    networkZone := "networkZone_example" // string | The network zone you want the result to be configured with. Requires at least ActiveGate version 1.247. (optional)
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Defaults to `amd64`.  * `amd64`: amd64 architecture. * `s390`: S/390 architecture, only supported for Linux.    (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.DownloadGatewayInstallerWithVersion(context.Background(), osType, version).IfNoneMatch(ifNoneMatch).NetworkZone(networkZone).Arch(arch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DownloadGatewayInstallerWithVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**version** | **string** | The required version of the ActiveGate installer, in &#x60;1.155.275.20181112-084458&#x60; format.   You can retrieve the list of available versions with the [**GET available versions of ActiveGate**](https://dt-url.net/kh43rha) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadGatewayInstallerWithVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** | The ETag of the previous request. Do not download if it matches the ETag of the installer. | 
 **networkZone** | **string** | The network zone you want the result to be configured with. Requires at least ActiveGate version 1.247. | 
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Defaults to &#x60;amd64&#x60;.  * &#x60;amd64&#x60;: amd64 architecture. * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.    | [default to &quot;all&quot;]

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLatestAgentInstaller

> DownloadLatestAgentInstaller(ctx, osType, installerType).IfNoneMatch(ifNoneMatch).Flavor(flavor).Arch(arch).Bitness(bitness).Include(include).SkipMetadata(skipMetadata).NetworkZone(networkZone).Execute()

Downloads the latest OneAgent installer



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
    osType := "osType_example" // string | The operating system of the installer.
    installerType := "installerType_example" // string | The type of the installer:   * `default`: Self-extracting installer for manual installation. Downloads an `.exe` file for Windows or an `.sh` file for Unix.  * `default-unattended`: Self-extracting installer for unattended installation. Windows only. Downloads a `.zip` archive, containing the `.msi` installer and the batch file. This option is deprecated with OneAgent version 1.173  * `mainframe`: Downloads all code modules for z/OS combined in a single `*.pax` archive.  * `paas`: Code modules installer. Downloads a `*.zip` archive, containing the `manifest.json` file with meta information or a `.jar` file for z/OS.  * `paas-sh`: Code modules installer. Downloads a self-extracting shell script with the embedded `tar.gz` archive.
    ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. Do not download if it matches the ETag of the installer. (optional)
    flavor := "flavor_example" // string | The flavor of your Linux distribution:   * `musl` for Linux distributions, which are using the musl C standard library, for example Alpine Linux.  * 'multidistro` for all Linux distributions which are using musl C and glibc standard library.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "default")
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Use this value for AIX and z/OS. Defaults to `x86` for other OS types.  * `x86`: x86 architecture. * `ppc`: PowerPC architecture, only supported for AIX and Linux. * `ppcle`: PowerPC Little Endian architecture, only supported for Linux. * `sparc`: Sparc architecture, only supported for Solaris.   * `arm`: ARM architecture, only supported for Linux.   * `s390`: S/390 architecture, only supported for Linux.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    bitness := "bitness_example" // string | The bitness of your OS. Must be supported by the OS.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    include := []string{"Include_example"} // []string | The code modules to be included to the installer. You can specify several modules in the following format: `include=java&include=dotnet`.   Only applicable to the `paas` and `paas-sh` installer types. (optional)
    skipMetadata := true // bool | Set `true` to omit the OneAgent connectivity information from the installer.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to false)
    networkZone := "networkZone_example" // string | The network zone you want the result to be configured with. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.DownloadLatestAgentInstaller(context.Background(), osType, installerType).IfNoneMatch(ifNoneMatch).Flavor(flavor).Arch(arch).Bitness(bitness).Include(include).SkipMetadata(skipMetadata).NetworkZone(networkZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DownloadLatestAgentInstaller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**installerType** | **string** | The type of the installer:   * &#x60;default&#x60;: Self-extracting installer for manual installation. Downloads an &#x60;.exe&#x60; file for Windows or an &#x60;.sh&#x60; file for Unix.  * &#x60;default-unattended&#x60;: Self-extracting installer for unattended installation. Windows only. Downloads a &#x60;.zip&#x60; archive, containing the &#x60;.msi&#x60; installer and the batch file. This option is deprecated with OneAgent version 1.173  * &#x60;mainframe&#x60;: Downloads all code modules for z/OS combined in a single &#x60;*.pax&#x60; archive.  * &#x60;paas&#x60;: Code modules installer. Downloads a &#x60;*.zip&#x60; archive, containing the &#x60;manifest.json&#x60; file with meta information or a &#x60;.jar&#x60; file for z/OS.  * &#x60;paas-sh&#x60;: Code modules installer. Downloads a self-extracting shell script with the embedded &#x60;tar.gz&#x60; archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLatestAgentInstallerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** | The ETag of the previous request. Do not download if it matches the ETag of the installer. | 
 **flavor** | **string** | The flavor of your Linux distribution:   * &#x60;musl&#x60; for Linux distributions, which are using the musl C standard library, for example Alpine Linux.  * &#39;multidistro&#x60; for all Linux distributions which are using musl C and glibc standard library.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;default&quot;]
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Use this value for AIX and z/OS. Defaults to &#x60;x86&#x60; for other OS types.  * &#x60;x86&#x60;: x86 architecture. * &#x60;ppc&#x60;: PowerPC architecture, only supported for AIX and Linux. * &#x60;ppcle&#x60;: PowerPC Little Endian architecture, only supported for Linux. * &#x60;sparc&#x60;: Sparc architecture, only supported for Solaris.   * &#x60;arm&#x60;: ARM architecture, only supported for Linux.   * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **bitness** | **string** | The bitness of your OS. Must be supported by the OS.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **include** | **[]string** | The code modules to be included to the installer. You can specify several modules in the following format: &#x60;include&#x3D;java&amp;include&#x3D;dotnet&#x60;.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | 
 **skipMetadata** | **bool** | Set &#x60;true&#x60; to omit the OneAgent connectivity information from the installer.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to false]
 **networkZone** | **string** | The network zone you want the result to be configured with. | 

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLatestAgentOrchestration

> DownloadLatestAgentOrchestration(ctx, orchestrationType).Execute()

Downloads the latest OneAgent deployment orchestration tarball



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
    orchestrationType := "orchestrationType_example" // string | The Orchestration Type of the orchestration deployment script.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.DownloadLatestAgentOrchestration(context.Background(), orchestrationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DownloadLatestAgentOrchestration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orchestrationType** | **string** | The Orchestration Type of the orchestration deployment script. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLatestAgentOrchestrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLatestAgentOrchestrationSignature

> DownloadLatestAgentOrchestrationSignature(ctx, orchestrationType).Execute()

Downloads the latest OneAgent deployment orchestration tarball's signature



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
    orchestrationType := "orchestrationType_example" // string | The Orchestration Type of the orchestration deployment script.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.DownloadLatestAgentOrchestrationSignature(context.Background(), orchestrationType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DownloadLatestAgentOrchestrationSignature``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orchestrationType** | **string** | The Orchestration Type of the orchestration deployment script. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLatestAgentOrchestrationSignatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadLatestGatewayInstaller

> DownloadLatestGatewayInstaller(ctx, osType).IfNoneMatch(ifNoneMatch).NetworkZone(networkZone).Arch(arch).Execute()

Downloads the configured standard ActiveGate installer of the latest version for the specified OS

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
    osType := "osType_example" // string | The operating system of the installer.
    ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. Do not download if it matches the ETag of the installer. (optional)
    networkZone := "networkZone_example" // string | The network zone you want the result to be configured with. Requires at least ActiveGate version 1.247. (optional)
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Defaults to `amd64`.  * `amd64`: amd64 architecture. * `s390`: S/390 architecture, only supported for Linux.    (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.DownloadLatestGatewayInstaller(context.Background(), osType).IfNoneMatch(ifNoneMatch).NetworkZone(networkZone).Arch(arch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.DownloadLatestGatewayInstaller``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadLatestGatewayInstallerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifNoneMatch** | **string** | The ETag of the previous request. Do not download if it matches the ETag of the installer. | 
 **networkZone** | **string** | The network zone you want the result to be configured with. Requires at least ActiveGate version 1.247. | 
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Defaults to &#x60;amd64&#x60;.  * &#x60;amd64&#x60;: amd64 architecture. * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.    | [default to &quot;all&quot;]

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveGateInstallerAvailableVersions

> ActiveGateInstallerVersions GetActiveGateInstallerAvailableVersions(ctx, osType).Arch(arch).Execute()

Lists all available versions of ActiveGate installer

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
    osType := "osType_example" // string | The operating system of the installer.
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Defaults to `amd64`.  * `amd64`: amd64 architecture. * `s390`: S/390 architecture, only supported for Linux.    (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetActiveGateInstallerAvailableVersions(context.Background(), osType).Arch(arch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetActiveGateInstallerAvailableVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveGateInstallerAvailableVersions`: ActiveGateInstallerVersions
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetActiveGateInstallerAvailableVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveGateInstallerAvailableVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Defaults to &#x60;amd64&#x60;.  * &#x60;amd64&#x60;: amd64 architecture. * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.    | [default to &quot;all&quot;]

### Return type

[**ActiveGateInstallerVersions**](ActiveGateInstallerVersions.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetActiveGateInstallerConnectionInfo

> ActiveGateConnectionInfo GetActiveGateInstallerConnectionInfo(ctx).NetworkZone(networkZone).Execute()

Gets the connectivity information for Environment ActiveGate

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
    networkZone := "networkZone_example" // string | The network zone you want the result to be configured with. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetActiveGateInstallerConnectionInfo(context.Background()).NetworkZone(networkZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetActiveGateInstallerConnectionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActiveGateInstallerConnectionInfo`: ActiveGateConnectionInfo
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetActiveGateInstallerConnectionInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActiveGateInstallerConnectionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkZone** | **string** | The network zone you want the result to be configured with. | 

### Return type

[**ActiveGateConnectionInfo**](ActiveGateConnectionInfo.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentInstallerAvailableVersions

> AgentInstallerVersions GetAgentInstallerAvailableVersions(ctx, osType, installerType).Flavor(flavor).Arch(arch).Execute()

Lists all available versions of OneAgent installer

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
    osType := "osType_example" // string | The operating system of the installer.
    installerType := "installerType_example" // string | The type of the installer:   * `default`: Self-extracting installer for manual installation. Downloads an `.exe` file for Windows or an `.sh` file for Unix.  * `default-unattended`: Self-extracting installer for unattended installation. Windows only. Downloads a `.zip` archive, containing the `.msi` installer and the batch file. This option is deprecated with OneAgent version 1.173  * `mainframe`: Downloads all code modules for z/OS combined in a single `*.pax` archive.  * `paas`: Code modules installer. Downloads a `*.zip` archive, containing the `manifest.json` file with meta information or a `.jar` file for z/OS.  * `paas-sh`: Code modules installer. Downloads a self-extracting shell script with the embedded `tar.gz` archive.
    flavor := "flavor_example" // string | The flavor of your Linux distribution:   * `musl` for Linux distributions, which are using the musl C standard library, for example Alpine Linux.  * 'multidistro` for all Linux distributions which are using musl C and glibc standard library.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "default")
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Use this value for AIX and z/OS. Defaults to `x86` for other OS types.  * `x86`: x86 architecture. * `ppc`: PowerPC architecture, only supported for AIX and Linux. * `ppcle`: PowerPC Little Endian architecture, only supported for Linux. * `sparc`: Sparc architecture, only supported for Solaris.   * `arm`: ARM architecture, only supported for Linux.   * `s390`: S/390 architecture, only supported for Linux.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetAgentInstallerAvailableVersions(context.Background(), osType, installerType).Flavor(flavor).Arch(arch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetAgentInstallerAvailableVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentInstallerAvailableVersions`: AgentInstallerVersions
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetAgentInstallerAvailableVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**installerType** | **string** | The type of the installer:   * &#x60;default&#x60;: Self-extracting installer for manual installation. Downloads an &#x60;.exe&#x60; file for Windows or an &#x60;.sh&#x60; file for Unix.  * &#x60;default-unattended&#x60;: Self-extracting installer for unattended installation. Windows only. Downloads a &#x60;.zip&#x60; archive, containing the &#x60;.msi&#x60; installer and the batch file. This option is deprecated with OneAgent version 1.173  * &#x60;mainframe&#x60;: Downloads all code modules for z/OS combined in a single &#x60;*.pax&#x60; archive.  * &#x60;paas&#x60;: Code modules installer. Downloads a &#x60;*.zip&#x60; archive, containing the &#x60;manifest.json&#x60; file with meta information or a &#x60;.jar&#x60; file for z/OS.  * &#x60;paas-sh&#x60;: Code modules installer. Downloads a self-extracting shell script with the embedded &#x60;tar.gz&#x60; archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInstallerAvailableVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flavor** | **string** | The flavor of your Linux distribution:   * &#x60;musl&#x60; for Linux distributions, which are using the musl C standard library, for example Alpine Linux.  * &#39;multidistro&#x60; for all Linux distributions which are using musl C and glibc standard library.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;default&quot;]
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Use this value for AIX and z/OS. Defaults to &#x60;x86&#x60; for other OS types.  * &#x60;x86&#x60;: x86 architecture. * &#x60;ppc&#x60;: PowerPC architecture, only supported for AIX and Linux. * &#x60;ppcle&#x60;: PowerPC Little Endian architecture, only supported for Linux. * &#x60;sparc&#x60;: Sparc architecture, only supported for Solaris.   * &#x60;arm&#x60;: ARM architecture, only supported for Linux.   * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]

### Return type

[**AgentInstallerVersions**](AgentInstallerVersions.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentInstallerConnectionInfo

> ConnectionInfo GetAgentInstallerConnectionInfo(ctx).NetworkZone(networkZone).DefaultZoneFallback(defaultZoneFallback).Version(version).Execute()

Gets the connectivity information for OneAgent

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
    networkZone := "networkZone_example" // string | The network zone you want the result to be configured with. (optional)
    defaultZoneFallback := true // bool | Set `true` to perform a fallback to the default network zone if the provided network zone does not exist. (optional) (default to false)
    version := "version_example" // string | The version of the OneAgent for which you're requesting connectivity information, in the `1.221` format.   Set this parameter to get the best format of endpoint list for optimal performance. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetAgentInstallerConnectionInfo(context.Background()).NetworkZone(networkZone).DefaultZoneFallback(defaultZoneFallback).Version(version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetAgentInstallerConnectionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentInstallerConnectionInfo`: ConnectionInfo
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetAgentInstallerConnectionInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInstallerConnectionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkZone** | **string** | The network zone you want the result to be configured with. | 
 **defaultZoneFallback** | **bool** | Set &#x60;true&#x60; to perform a fallback to the default network zone if the provided network zone does not exist. | [default to false]
 **version** | **string** | The version of the OneAgent for which you&#39;re requesting connectivity information, in the &#x60;1.221&#x60; format.   Set this parameter to get the best format of endpoint list for optimal performance. | 

### Return type

[**ConnectionInfo**](ConnectionInfo.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentInstallerConnectionInfoEndpoints

> GetAgentInstallerConnectionInfoEndpoints(ctx).NetworkZone(networkZone).DefaultZoneFallback(defaultZoneFallback).Execute()

Gets the list of the ActiveGate-Endpoints to be used for Agents ordered by networkzone-priorities.



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
    networkZone := "networkZone_example" // string |  (optional)
    defaultZoneFallback := true // bool | Set `true` to perform a fallback to the default network zone if the provided network zone does not exist. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeploymentAPI.GetAgentInstallerConnectionInfoEndpoints(context.Background()).NetworkZone(networkZone).DefaultZoneFallback(defaultZoneFallback).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetAgentInstallerConnectionInfoEndpoints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInstallerConnectionInfoEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkZone** | **string** |  | 
 **defaultZoneFallback** | **bool** | Set &#x60;true&#x60; to perform a fallback to the default network zone if the provided network zone does not exist. | [default to false]

### Return type

 (empty response body)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentInstallerMetaInfo

> InstallerMetaInfoDto GetAgentInstallerMetaInfo(ctx, osType, installerType).Flavor(flavor).Arch(arch).Bitness(bitness).Execute()

Gets the latest available version of a OneAgent installer



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
    osType := "osType_example" // string | The operating system of the installer.
    installerType := "installerType_example" // string | The type of the installer:   * `default`: Self-extracting installer for manual installation. Downloads an `.exe` file for Windows or an `.sh` file for Unix.  * `default-unattended`: Self-extracting installer for unattended installation. Windows only. Downloads a `.zip` archive, containing the `.msi` installer and the batch file. This option is deprecated with OneAgent version 1.173  * `mainframe`: Downloads all code modules for z/OS combined in a single `*.pax` archive.  * `paas`: Code modules installer. Downloads a `*.zip` archive, containing the `manifest.json` file with meta information or a `.jar` file for z/OS.  * `paas-sh`: Code modules installer. Downloads a self-extracting shell script with the embedded `tar.gz` archive.
    flavor := "flavor_example" // string | The flavor of your Linux distribution:   * `musl` for Linux distributions, which are using the musl C standard library, for example Alpine Linux.  * 'multidistro` for all Linux distributions which are using musl C and glibc standard library.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "default")
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Use this value for AIX and z/OS. Defaults to `x86` for other OS types.  * `x86`: x86 architecture. * `ppc`: PowerPC architecture, only supported for AIX and Linux. * `ppcle`: PowerPC Little Endian architecture, only supported for Linux. * `sparc`: Sparc architecture, only supported for Solaris.   * `arm`: ARM architecture, only supported for Linux.   * `s390`: S/390 architecture, only supported for Linux.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    bitness := "bitness_example" // string | The bitness of your OS. Must be supported by the OS.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetAgentInstallerMetaInfo(context.Background(), osType, installerType).Flavor(flavor).Arch(arch).Bitness(bitness).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetAgentInstallerMetaInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentInstallerMetaInfo`: InstallerMetaInfoDto
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetAgentInstallerMetaInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**installerType** | **string** | The type of the installer:   * &#x60;default&#x60;: Self-extracting installer for manual installation. Downloads an &#x60;.exe&#x60; file for Windows or an &#x60;.sh&#x60; file for Unix.  * &#x60;default-unattended&#x60;: Self-extracting installer for unattended installation. Windows only. Downloads a &#x60;.zip&#x60; archive, containing the &#x60;.msi&#x60; installer and the batch file. This option is deprecated with OneAgent version 1.173  * &#x60;mainframe&#x60;: Downloads all code modules for z/OS combined in a single &#x60;*.pax&#x60; archive.  * &#x60;paas&#x60;: Code modules installer. Downloads a &#x60;*.zip&#x60; archive, containing the &#x60;manifest.json&#x60; file with meta information or a &#x60;.jar&#x60; file for z/OS.  * &#x60;paas-sh&#x60;: Code modules installer. Downloads a self-extracting shell script with the embedded &#x60;tar.gz&#x60; archive. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInstallerMetaInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **flavor** | **string** | The flavor of your Linux distribution:   * &#x60;musl&#x60; for Linux distributions, which are using the musl C standard library, for example Alpine Linux.  * &#39;multidistro&#x60; for all Linux distributions which are using musl C and glibc standard library.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;default&quot;]
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Use this value for AIX and z/OS. Defaults to &#x60;x86&#x60; for other OS types.  * &#x60;x86&#x60;: x86 architecture. * &#x60;ppc&#x60;: PowerPC architecture, only supported for AIX and Linux. * &#x60;ppcle&#x60;: PowerPC Little Endian architecture, only supported for Linux. * &#x60;sparc&#x60;: Sparc architecture, only supported for Solaris.   * &#x60;arm&#x60;: ARM architecture, only supported for Linux.   * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **bitness** | **string** | The bitness of your OS. Must be supported by the OS.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]

### Return type

[**InstallerMetaInfoDto**](InstallerMetaInfoDto.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentInstallerWithVersionChecksum

> OneAgentInstallerChecksum GetAgentInstallerWithVersionChecksum(ctx, osType, installerType, version).IfNoneMatch(ifNoneMatch).Flavor(flavor).Arch(arch).Bitness(bitness).Include(include).NetworkZone(networkZone).Execute()

Gets the checksum of a non-customized OneAgent installer



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
    osType := "osType_example" // string | The operating system of the installer.
    installerType := "installerType_example" // string | The type of the installer.
    version := "version_example" // string | The required version of the OneAgent in `1.155.275.20181112-084458` format.   You can retrieve the list of available versions with the [**GET available versions of OneAgent**](https://dt-url.net/fo23rb5) call.
    ifNoneMatch := "ifNoneMatch_example" // string | The ETag of the previous request. Do not download if it matches the ETag of the installer. (optional)
    flavor := "flavor_example" // string | The flavor of your Linux distribution:   * `musl` for Linux distributions, which are using the musl C standard library, for example Alpine Linux.  * 'multidistro` for all Linux distributions which are using musl C and glibc standard library.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "default")
    arch := "arch_example" // string | The architecture of your OS:   * `all`: Use this value for AIX and z/OS. Defaults to `x86` for other OS types.  * `x86`: x86 architecture. * `ppc`: PowerPC architecture, only supported for AIX and Linux. * `ppcle`: PowerPC Little Endian architecture, only supported for Linux. * `sparc`: Sparc architecture, only supported for Solaris.   * `arm`: ARM architecture, only supported for Linux.   * `s390`: S/390 architecture, only supported for Linux.   Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    bitness := "bitness_example" // string | The bitness of your OS. Must be supported by the OS.    Only applicable to the `paas` and `paas-sh` installer types. (optional) (default to "all")
    include := []string{"Include_example"} // []string | The code modules to be included to the installer. You can specify several modules in the following format: `include=java&include=dotnet`.   Only applicable to the `paas` and `paas-sh` installer types. (optional)
    networkZone := "networkZone_example" // string | The network zone you want the result to be configured with. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetAgentInstallerWithVersionChecksum(context.Background(), osType, installerType, version).IfNoneMatch(ifNoneMatch).Flavor(flavor).Arch(arch).Bitness(bitness).Include(include).NetworkZone(networkZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetAgentInstallerWithVersionChecksum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentInstallerWithVersionChecksum`: OneAgentInstallerChecksum
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetAgentInstallerWithVersionChecksum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**installerType** | **string** | The type of the installer. | 
**version** | **string** | The required version of the OneAgent in &#x60;1.155.275.20181112-084458&#x60; format.   You can retrieve the list of available versions with the [**GET available versions of OneAgent**](https://dt-url.net/fo23rb5) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentInstallerWithVersionChecksumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **ifNoneMatch** | **string** | The ETag of the previous request. Do not download if it matches the ETag of the installer. | 
 **flavor** | **string** | The flavor of your Linux distribution:   * &#x60;musl&#x60; for Linux distributions, which are using the musl C standard library, for example Alpine Linux.  * &#39;multidistro&#x60; for all Linux distributions which are using musl C and glibc standard library.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;default&quot;]
 **arch** | **string** | The architecture of your OS:   * &#x60;all&#x60;: Use this value for AIX and z/OS. Defaults to &#x60;x86&#x60; for other OS types.  * &#x60;x86&#x60;: x86 architecture. * &#x60;ppc&#x60;: PowerPC architecture, only supported for AIX and Linux. * &#x60;ppcle&#x60;: PowerPC Little Endian architecture, only supported for Linux. * &#x60;sparc&#x60;: Sparc architecture, only supported for Solaris.   * &#x60;arm&#x60;: ARM architecture, only supported for Linux.   * &#x60;s390&#x60;: S/390 architecture, only supported for Linux.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **bitness** | **string** | The bitness of your OS. Must be supported by the OS.    Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | [default to &quot;all&quot;]
 **include** | **[]string** | The code modules to be included to the installer. You can specify several modules in the following format: &#x60;include&#x3D;java&amp;include&#x3D;dotnet&#x60;.   Only applicable to the &#x60;paas&#x60; and &#x60;paas-sh&#x60; installer types. | 
 **networkZone** | **string** | The network zone you want the result to be configured with. | 

### Return type

[**OneAgentInstallerChecksum**](OneAgentInstallerChecksum.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAgentProcessModuleConfig

> AgentProcessModuleConfigResponse GetAgentProcessModuleConfig(ctx).Revision(revision).Execute()

Gets the latest process module config | maturity=EARLY_ADOPTER



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
    revision := int64(789) // int64 | The previously received revision to compare against. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetAgentProcessModuleConfig(context.Background()).Revision(revision).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetAgentProcessModuleConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAgentProcessModuleConfig`: AgentProcessModuleConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetAgentProcessModuleConfig`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAgentProcessModuleConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revision** | **int64** | The previously received revision to compare against. | 

### Return type

[**AgentProcessModuleConfigResponse**](AgentProcessModuleConfigResponse.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoshReleaseAvailableVersions

> BoshReleaseAvailableVersions GetBoshReleaseAvailableVersions(ctx, osType).Execute()

Gets the list of available OneAgent versions for BOSH release tarballs

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
    osType := "osType_example" // string | The operating system of the installer.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetBoshReleaseAvailableVersions(context.Background(), osType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetBoshReleaseAvailableVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBoshReleaseAvailableVersions`: BoshReleaseAvailableVersions
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetBoshReleaseAvailableVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoshReleaseAvailableVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BoshReleaseAvailableVersions**](BoshReleaseAvailableVersions.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoshReleaseChecksum

> BoshReleaseChecksum GetBoshReleaseChecksum(ctx, osType, version).SkipMetadata(skipMetadata).NetworkZone(networkZone).Execute()

Gets the checksum of the specified BOSH release tarball



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
    osType := "osType_example" // string | The operating system of the installer.
    version := "version_example" // string | The required version of the OneAgent in the `1.155.275.20181112-084458` format.   You can retrieve the list of available versions with the [**GET available versions of BOSH tarballs**](https://dt-url.net/j703kdn) call.
    skipMetadata := true // bool | Set `true` to omit the OneAgent connectivity information from the installer.    If not set, `false` is used. (optional) (default to false)
    networkZone := "networkZone_example" // string | The network zone you want the result to be configured with. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeploymentAPI.GetBoshReleaseChecksum(context.Background(), osType, version).SkipMetadata(skipMetadata).NetworkZone(networkZone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetBoshReleaseChecksum``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBoshReleaseChecksum`: BoshReleaseChecksum
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetBoshReleaseChecksum`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**osType** | **string** | The operating system of the installer. | 
**version** | **string** | The required version of the OneAgent in the &#x60;1.155.275.20181112-084458&#x60; format.   You can retrieve the list of available versions with the [**GET available versions of BOSH tarballs**](https://dt-url.net/j703kdn) call. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoshReleaseChecksumRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **skipMetadata** | **bool** | Set &#x60;true&#x60; to omit the OneAgent connectivity information from the installer.    If not set, &#x60;false&#x60; is used. | [default to false]
 **networkZone** | **string** | The network zone you want the result to be configured with. | 

### Return type

[**BoshReleaseChecksum**](BoshReleaseChecksum.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestLambdaBuildUnits

> LatestLambdaLayerNames GetLatestLambdaBuildUnits(ctx).Execute()

Get the latest version names of the OneAgent for AWS Lambda



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
    resp, r, err := apiClient.DeploymentAPI.GetLatestLambdaBuildUnits(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeploymentAPI.GetLatestLambdaBuildUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestLambdaBuildUnits`: LatestLambdaLayerNames
    fmt.Fprintf(os.Stdout, "Response from `DeploymentAPI.GetLatestLambdaBuildUnits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestLambdaBuildUnitsRequest struct via the builder pattern


### Return type

[**LatestLambdaLayerNames**](LatestLambdaLayerNames.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

