# Go API client for environmentv1

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).

Notes about compatibility:
* Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this.
* We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import environmentv1 "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), environmentv1.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), environmentv1.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), environmentv1.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), environmentv1.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AnonymizationAPI* | [**Anonymize**](docs/AnonymizationAPI.md#anonymize) | **Put** /anonymize/anonymizationJobs | Creates user session anonymization job
*AnonymizationAPI* | [**GetStatus**](docs/AnonymizationAPI.md#getstatus) | **Get** /anonymize/anonymizationJobs/{requestId} | Shows the progress of the specified anonymization job
*ClusterTimeAPI* | [**GetCurrentClusterTime**](docs/ClusterTimeAPI.md#getcurrentclustertime) | **Get** /time | Gets the current time of the Dynatrace server in UTC milliseconds
*ClusterVersionAPI* | [**GetVersion**](docs/ClusterVersionAPI.md#getversion) | **Get** /config/clusterversion | Gets the current version of the Dynatrace server
*DeploymentAPI* | [**DownloadAgentInstallerWithVersion**](docs/DeploymentAPI.md#downloadagentinstallerwithversion) | **Get** /deployment/installer/agent/{osType}/{installerType}/version/{version} | Downloads OneAgent installer of the specified version
*DeploymentAPI* | [**DownloadAgentOrchestrationSignatureWithVersion**](docs/DeploymentAPI.md#downloadagentorchestrationsignaturewithversion) | **Get** /deployment/orchestration/agent/{orchestrationType}/version/{version}/signature | Downloads the requested version matching OneAgent deployment orchestration tarball&#39;s signature
*DeploymentAPI* | [**DownloadAgentOrchestrationWithVersion**](docs/DeploymentAPI.md#downloadagentorchestrationwithversion) | **Get** /deployment/orchestration/agent/{orchestrationType}/version/{version} | Downloads the requested version matching OneAgent deployment orchestration tarball
*DeploymentAPI* | [**DownloadBoshReleaseWithVersion**](docs/DeploymentAPI.md#downloadboshreleasewithversion) | **Get** /deployment/boshrelease/agent/{osType}/version/{version} | Downloads BOSH release tarballs of the specified version, OneAgent included
*DeploymentAPI* | [**DownloadGatewayInstallerWithVersion**](docs/DeploymentAPI.md#downloadgatewayinstallerwithversion) | **Get** /deployment/installer/gateway/{osType}/version/{version} | Downloads the ActiveGate installer of the specified version
*DeploymentAPI* | [**DownloadLatestAgentInstaller**](docs/DeploymentAPI.md#downloadlatestagentinstaller) | **Get** /deployment/installer/agent/{osType}/{installerType}/latest | Downloads the latest OneAgent installer
*DeploymentAPI* | [**DownloadLatestAgentOrchestration**](docs/DeploymentAPI.md#downloadlatestagentorchestration) | **Get** /deployment/orchestration/agent/{orchestrationType}/latest | Downloads the latest OneAgent deployment orchestration tarball
*DeploymentAPI* | [**DownloadLatestAgentOrchestrationSignature**](docs/DeploymentAPI.md#downloadlatestagentorchestrationsignature) | **Get** /deployment/orchestration/agent/{orchestrationType}/latest/signature | Downloads the latest OneAgent deployment orchestration tarball&#39;s signature
*DeploymentAPI* | [**DownloadLatestGatewayInstaller**](docs/DeploymentAPI.md#downloadlatestgatewayinstaller) | **Get** /deployment/installer/gateway/{osType}/latest | Downloads the configured standard ActiveGate installer of the latest version for the specified OS
*DeploymentAPI* | [**GetActiveGateInstallerAvailableVersions**](docs/DeploymentAPI.md#getactivegateinstalleravailableversions) | **Get** /deployment/installer/gateway/versions/{osType} | Lists all available versions of ActiveGate installer
*DeploymentAPI* | [**GetActiveGateInstallerConnectionInfo**](docs/DeploymentAPI.md#getactivegateinstallerconnectioninfo) | **Get** /deployment/installer/gateway/connectioninfo | Gets the connectivity information for Environment ActiveGate
*DeploymentAPI* | [**GetAgentInstallerAvailableVersions**](docs/DeploymentAPI.md#getagentinstalleravailableversions) | **Get** /deployment/installer/agent/versions/{osType}/{installerType} | Lists all available versions of OneAgent installer
*DeploymentAPI* | [**GetAgentInstallerConnectionInfo**](docs/DeploymentAPI.md#getagentinstallerconnectioninfo) | **Get** /deployment/installer/agent/connectioninfo | Gets the connectivity information for OneAgent
*DeploymentAPI* | [**GetAgentInstallerConnectionInfoEndpoints**](docs/DeploymentAPI.md#getagentinstallerconnectioninfoendpoints) | **Get** /deployment/installer/agent/connectioninfo/endpoints | Gets the list of the ActiveGate-Endpoints to be used for Agents ordered by networkzone-priorities.
*DeploymentAPI* | [**GetAgentInstallerMetaInfo**](docs/DeploymentAPI.md#getagentinstallermetainfo) | **Get** /deployment/installer/agent/{osType}/{installerType}/latest/metainfo | Gets the latest available version of a OneAgent installer
*DeploymentAPI* | [**GetAgentInstallerWithVersionChecksum**](docs/DeploymentAPI.md#getagentinstallerwithversionchecksum) | **Get** /deployment/installer/agent/{osType}/{installerType}/version/{version}/checksum | Gets the checksum of a non-customized OneAgent installer
*DeploymentAPI* | [**GetAgentProcessModuleConfig**](docs/DeploymentAPI.md#getagentprocessmoduleconfig) | **Get** /deployment/installer/agent/processmoduleconfig | Gets the latest process module config | maturity&#x3D;EARLY_ADOPTER
*DeploymentAPI* | [**GetBoshReleaseAvailableVersions**](docs/DeploymentAPI.md#getboshreleaseavailableversions) | **Get** /deployment/boshrelease/versions/{osType} | Gets the list of available OneAgent versions for BOSH release tarballs
*DeploymentAPI* | [**GetBoshReleaseChecksum**](docs/DeploymentAPI.md#getboshreleasechecksum) | **Get** /deployment/boshrelease/agent/{osType}/version/{version}/checksum | Gets the checksum of the specified BOSH release tarball
*DeploymentAPI* | [**GetLatestLambdaBuildUnits**](docs/DeploymentAPI.md#getlatestlambdabuildunits) | **Get** /deployment/lambda/agent/latest | Get the latest version names of the OneAgent for AWS Lambda
*EventsAPI* | [**GetEventById**](docs/EventsAPI.md#geteventbyid) | **Get** /events/{eventId} | Gets the parameters of the specified event
*EventsAPI* | [**PostEvent**](docs/EventsAPI.md#postevent) | **Post** /events | Pushes custom events to one or more monitored entities
*EventsAPI* | [**QueryEvents**](docs/EventsAPI.md#queryevents) | **Get** /events | Lists all the events of your environment, along with their parameters
*LogMonitoringCustomDevicesAPI* | [**CustomDeviceLogJobDelete**](docs/LogMonitoringCustomDevicesAPI.md#customdevicelogjobdelete) | **Delete** /entity/infrastructure/custom-devices/{customDeviceId}/logs/jobs/{jobId} | Deletes or cancels the specified log analysis job
*LogMonitoringCustomDevicesAPI* | [**CustomDeviceLogJobRecords**](docs/LogMonitoringCustomDevicesAPI.md#customdevicelogjobrecords) | **Get** /entity/infrastructure/custom-devices/{customDeviceId}/logs/jobs/{jobId}/records | Gets the content of the analyzed log
*LogMonitoringCustomDevicesAPI* | [**CustomDeviceLogJobRecordsFiltered**](docs/LogMonitoringCustomDevicesAPI.md#customdevicelogjobrecordsfiltered) | **Post** /entity/infrastructure/custom-devices/{customDeviceId}/logs/jobs/{jobId}/records | Gets the filtered content of the analyzed log
*LogMonitoringCustomDevicesAPI* | [**CustomDeviceLogJobRecordsTop**](docs/LogMonitoringCustomDevicesAPI.md#customdevicelogjobrecordstop) | **Post** /entity/infrastructure/custom-devices/{customDeviceId}/logs/jobs/{jobId}/records/top | Gets the top values of fields present in the content of the analyzed log
*LogMonitoringCustomDevicesAPI* | [**CustomDeviceLogJobStart**](docs/LogMonitoringCustomDevicesAPI.md#customdevicelogjobstart) | **Post** /entity/infrastructure/custom-devices/{customDeviceId}/logs/{logPath} | Starts the analysis job for the specified custom device log
*LogMonitoringCustomDevicesAPI* | [**CustomDeviceLogJobStatus**](docs/LogMonitoringCustomDevicesAPI.md#customdevicelogjobstatus) | **Get** /entity/infrastructure/custom-devices/{customDeviceId}/logs/jobs/{jobId} | Gets status of the specified log analysis job
*LogMonitoringCustomDevicesAPI* | [**CustomDeviceLogList**](docs/LogMonitoringCustomDevicesAPI.md#customdeviceloglist) | **Get** /entity/infrastructure/custom-devices/{customDeviceId}/logs | Lists all the available logs on the specified custom device
*LogMonitoringHostsAPI* | [**HostLogJobDelete**](docs/LogMonitoringHostsAPI.md#hostlogjobdelete) | **Delete** /entity/infrastructure/hosts/{hostId}/logs/jobs/{jobId} | Deletes or cancels the specified log analysis job
*LogMonitoringHostsAPI* | [**HostLogJobRecords**](docs/LogMonitoringHostsAPI.md#hostlogjobrecords) | **Get** /entity/infrastructure/hosts/{hostId}/logs/jobs/{jobId}/records | Gets the full content of the analyzed log
*LogMonitoringHostsAPI* | [**HostLogJobRecordsFiltered**](docs/LogMonitoringHostsAPI.md#hostlogjobrecordsfiltered) | **Post** /entity/infrastructure/hosts/{hostId}/logs/jobs/{jobId}/records | Gets the filtered content of the analyzed log
*LogMonitoringHostsAPI* | [**HostLogJobRecordsTop**](docs/LogMonitoringHostsAPI.md#hostlogjobrecordstop) | **Post** /entity/infrastructure/hosts/{hostId}/logs/jobs/{jobId}/records/top | Gets the top values of fields present in the content of the analyzed log
*LogMonitoringHostsAPI* | [**HostLogJobStart**](docs/LogMonitoringHostsAPI.md#hostlogjobstart) | **Post** /entity/infrastructure/hosts/{hostId}/logs/{logPath} | Starts the analysis job for the specified OS log
*LogMonitoringHostsAPI* | [**HostLogJobStatus**](docs/LogMonitoringHostsAPI.md#hostlogjobstatus) | **Get** /entity/infrastructure/hosts/{hostId}/logs/jobs/{jobId} | Gets status of the specified log analysis job
*LogMonitoringHostsAPI* | [**HostLogList**](docs/LogMonitoringHostsAPI.md#hostloglist) | **Get** /entity/infrastructure/hosts/{hostId}/logs | Lists all the available OS logs on the specified host
*LogMonitoringProcessGroupsAPI* | [**ProcessGroupLogJobDelete**](docs/LogMonitoringProcessGroupsAPI.md#processgrouplogjobdelete) | **Delete** /entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId} | Deletes or cancels the specified log analysis job
*LogMonitoringProcessGroupsAPI* | [**ProcessGroupLogJobRecords**](docs/LogMonitoringProcessGroupsAPI.md#processgrouplogjobrecords) | **Get** /entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}/records | Gets the content of the analyzed log
*LogMonitoringProcessGroupsAPI* | [**ProcessGroupLogJobRecordsFiltered**](docs/LogMonitoringProcessGroupsAPI.md#processgrouplogjobrecordsfiltered) | **Post** /entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}/records | Gets the content of the analyzed log
*LogMonitoringProcessGroupsAPI* | [**ProcessGroupLogJobRecordsTop**](docs/LogMonitoringProcessGroupsAPI.md#processgrouplogjobrecordstop) | **Post** /entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId}/records/top | Gets the top values of fields present in the content of the analyzed log
*LogMonitoringProcessGroupsAPI* | [**ProcessGroupLogJobStart**](docs/LogMonitoringProcessGroupsAPI.md#processgrouplogjobstart) | **Post** /entity/infrastructure/process-groups/{pgId}/logs/{logPath} | Starts analysis job for the specified process group log
*LogMonitoringProcessGroupsAPI* | [**ProcessGroupLogJobStatus**](docs/LogMonitoringProcessGroupsAPI.md#processgrouplogjobstatus) | **Get** /entity/infrastructure/process-groups/{pgId}/logs/jobs/{jobId} | Gets status of the specified log analysis job
*LogMonitoringProcessGroupsAPI* | [**ProcessGroupLogList**](docs/LogMonitoringProcessGroupsAPI.md#processgrouploglist) | **Get** /entity/infrastructure/process-groups/{pgId}/logs | Lists all the available logs of the specified process group
*MaintenanceWindowAPI* | [**GetAllMaintenanceWindowConfigs**](docs/MaintenanceWindowAPI.md#getallmaintenancewindowconfigs) | **Get** /maintenance | Lists all parameters of all maintenance windows available in your Dynatrace environment.
*MaintenanceWindowAPI* | [**GetMaintenanceWindowConfig**](docs/MaintenanceWindowAPI.md#getmaintenancewindowconfig) | **Get** /maintenance/{uid} | Lists all parameters of the specified maintenance window.
*MaintenanceWindowAPI* | [**RemoveMaintenanceWindowConfig**](docs/MaintenanceWindowAPI.md#removemaintenancewindowconfig) | **Delete** /maintenance/{uid} | Deletes the specified maintenance window
*MaintenanceWindowAPI* | [**StoreMaintenanceWindowConfig**](docs/MaintenanceWindowAPI.md#storemaintenancewindowconfig) | **Post** /maintenance | Creates a new or updates an existing maintenance window
*OneAgentOnAHostAPI* | [**GetHostsWithSpecificAgents**](docs/OneAgentOnAHostAPI.md#gethostswithspecificagents) | **Get** /oneagents | Gets the list of hosts with OneAgent deployment information for each host
*ProblemAPI* | [**CloseProblem**](docs/ProblemAPI.md#closeproblem) | **Post** /problem/details/{problemId}/close | Closes the specified problem and adds a closing comment to it
*ProblemAPI* | [**DeleteComment**](docs/ProblemAPI.md#deletecomment) | **Delete** /problem/details/{problemId}/comments/{commentId} | Deletes an existing comment to the specified problem.
*ProblemAPI* | [**GetComment**](docs/ProblemAPI.md#getcomment) | **Get** /problem/details/{problemId}/comments | Gets all the comments to the specified problem
*ProblemAPI* | [**GetDetails**](docs/ProblemAPI.md#getdetails) | **Get** /problem/details/{problemId} | Gets the properties of the specified problem
*ProblemAPI* | [**GetFeed**](docs/ProblemAPI.md#getfeed) | **Get** /problem/feed | Gets the information about problems within the specified timeframe
*ProblemAPI* | [**GetProblemStatus**](docs/ProblemAPI.md#getproblemstatus) | **Get** /problem/status | Lists the number of open problems, split by impact level
*ProblemAPI* | [**PushComment**](docs/ProblemAPI.md#pushcomment) | **Post** /problem/details/{problemId}/comments | Adds a new comment to the specified problem
*ProblemAPI* | [**UpdateComment**](docs/ProblemAPI.md#updatecomment) | **Put** /problem/details/{problemId}/comments/{commentId} | Updates an existing comment to the specified problem
*RUMJavaScriptTagManagementAPI* | [**GetAppRevision**](docs/RUMJavaScriptTagManagementAPI.md#getapprevision) | **Get** /rum/appRevision/{entity} | Gets the version of the RUM JavaScript code injected into specified application
*RUMJavaScriptTagManagementAPI* | [**GetAsyncCodeSnippet**](docs/RUMJavaScriptTagManagementAPI.md#getasynccodesnippet) | **Get** /rum/asyncCS/{entity} | Downloads the asynchronous code snippet
*RUMJavaScriptTagManagementAPI* | [**GetJsInlineScript**](docs/RUMJavaScriptTagManagementAPI.md#getjsinlinescript) | **Get** /rum/jsInlineScript/{entity} | Downloads inline code
*RUMJavaScriptTagManagementAPI* | [**GetJsLatestVersion**](docs/RUMJavaScriptTagManagementAPI.md#getjslatestversion) | **Get** /rum/jsLatestVersion | Gets the latest version of OneAgent JavaScript library
*RUMJavaScriptTagManagementAPI* | [**GetJsScript**](docs/RUMJavaScriptTagManagementAPI.md#getjsscript) | **Get** /rum/jsTag/{entity} | Downloads OneAgent JavaScript tag
*RUMJavaScriptTagManagementAPI* | [**GetJsTagComplete**](docs/RUMJavaScriptTagManagementAPI.md#getjstagcomplete) | **Get** /rum/jsTagComplete/{entity} | Downloads JavaScript tag
*RUMJavaScriptTagManagementAPI* | [**GetManualApps**](docs/RUMJavaScriptTagManagementAPI.md#getmanualapps) | **Get** /rum/manualApps | Lists all manually injected applications
*RUMJavaScriptTagManagementAPI* | [**GetSyncCodeSnippet**](docs/RUMJavaScriptTagManagementAPI.md#getsynccodesnippet) | **Get** /rum/syncCS/{entity} | Downloads the synchronous code snippet
*RUMUserSessionsAPI* | [**GetUsqlResultAsTable**](docs/RUMUserSessionsAPI.md#getusqlresultastable) | **Get** /userSessionQueryLanguage/table | Returns the result of the query as a table structure
*RUMUserSessionsAPI* | [**GetUsqlResultAsTree**](docs/RUMUserSessionsAPI.md#getusqlresultastree) | **Get** /userSessionQueryLanguage/tree | Returns the result of the query as a tree structure
*SyntheticLocationsNodesAndConfigurationAPI* | [**AddLocation**](docs/SyntheticLocationsNodesAndConfigurationAPI.md#addlocation) | **Post** /synthetic/locations | Creates a new private synthetic location
*SyntheticLocationsNodesAndConfigurationAPI* | [**GetLocation**](docs/SyntheticLocationsNodesAndConfigurationAPI.md#getlocation) | **Get** /synthetic/locations/{locationId} | Gets properties of the specified location
*SyntheticLocationsNodesAndConfigurationAPI* | [**GetLocations**](docs/SyntheticLocationsNodesAndConfigurationAPI.md#getlocations) | **Get** /synthetic/locations | Lists all synthetic locations (both public and private) available for your environment
*SyntheticLocationsNodesAndConfigurationAPI* | [**GetLocationsStatus**](docs/SyntheticLocationsNodesAndConfigurationAPI.md#getlocationsstatus) | **Get** /synthetic/locations/status | Returns whether public locations are enabled or not | maturity&#x3D;EARLY_ADOPTER
*SyntheticLocationsNodesAndConfigurationAPI* | [**GetNode**](docs/SyntheticLocationsNodesAndConfigurationAPI.md#getnode) | **Get** /synthetic/nodes/{nodeId} | Lists properties of the specified synthetic node
*SyntheticLocationsNodesAndConfigurationAPI* | [**GetNodes**](docs/SyntheticLocationsNodesAndConfigurationAPI.md#getnodes) | **Get** /synthetic/nodes | Lists all synthetic nodes available in your environment
*SyntheticLocationsNodesAndConfigurationAPI* | [**RemoveLocation**](docs/SyntheticLocationsNodesAndConfigurationAPI.md#removelocation) | **Delete** /synthetic/locations/{locationId} | Deletes the specified private synthetic location
*SyntheticLocationsNodesAndConfigurationAPI* | [**UpdateLocation**](docs/SyntheticLocationsNodesAndConfigurationAPI.md#updatelocation) | **Put** /synthetic/locations/{locationId} | Updates the specified synthetic location
*SyntheticLocationsNodesAndConfigurationAPI* | [**UpdateLocationsStatus**](docs/SyntheticLocationsNodesAndConfigurationAPI.md#updatelocationsstatus) | **Put** /synthetic/locations/status | Enable/disable public synthetic locations | maturity&#x3D;EARLY_ADOPTER
*SyntheticMonitorsAPI* | [**AddMonitor**](docs/SyntheticMonitorsAPI.md#addmonitor) | **Post** /synthetic/monitors | Creates a new synthetic monitor
*SyntheticMonitorsAPI* | [**DeleteMonitor**](docs/SyntheticMonitorsAPI.md#deletemonitor) | **Delete** /synthetic/monitors/{monitorId} | Deletes the specified synthetic monitor
*SyntheticMonitorsAPI* | [**GetMonitor**](docs/SyntheticMonitorsAPI.md#getmonitor) | **Get** /synthetic/monitors/{monitorId} | Gets parameters of the specified synthetic monitor
*SyntheticMonitorsAPI* | [**GetMonitorsCollection**](docs/SyntheticMonitorsAPI.md#getmonitorscollection) | **Get** /synthetic/monitors | Lists all synthetic monitors in your Dynatrace environment
*SyntheticMonitorsAPI* | [**ReplaceMonitor**](docs/SyntheticMonitorsAPI.md#replacemonitor) | **Put** /synthetic/monitors/{monitorId} | Updates parameters of the specified synthetic monitor
*SyntheticThirdPartyAPI* | [**PushEvents**](docs/SyntheticThirdPartyAPI.md#pushevents) | **Post** /synthetic/ext/events | Pushes third-party synthetic events to Dynatrace
*SyntheticThirdPartyAPI* | [**PushStateModification**](docs/SyntheticThirdPartyAPI.md#pushstatemodification) | **Post** /synthetic/ext/stateModifications | Modifies the operation state of all third-party monitors
*SyntheticThirdPartyAPI* | [**TestResults**](docs/SyntheticThirdPartyAPI.md#testresults) | **Post** /synthetic/ext/tests | Pushes third-party synthetic monitors, locations, and results to Dynatrace
*ThresholdAPI* | [**CreateCustomThreshold**](docs/ThresholdAPI.md#createcustomthreshold) | **Put** /thresholds/{thresholdId} | Creates or updates an existing threshold for a plugin or a custom event
*ThresholdAPI* | [**DeleteCustomThreshold**](docs/ThresholdAPI.md#deletecustomthreshold) | **Delete** /thresholds/{thresholdId} | Deletes the specified threshold
*ThresholdAPI* | [**ReadCustomThresholds**](docs/ThresholdAPI.md#readcustomthresholds) | **Get** /thresholds | Gets all configured thresholds for plugins and custom events in your environment
*TimeseriesAPI* | [**CreateCustomTimeseries**](docs/TimeseriesAPI.md#createcustomtimeseries) | **Put** /timeseries/{timeseriesIdentifier} | Creates a new custom metric
*TimeseriesAPI* | [**DeleteCustomTimeseries**](docs/TimeseriesAPI.md#deletecustomtimeseries) | **Delete** /timeseries/{timeseriesIdentifier} | Deletes the specified custom metric
*TimeseriesAPI* | [**GetAllTimeseriesDefinitions**](docs/TimeseriesAPI.md#getalltimeseriesdefinitions) | **Get** /timeseries | Lists all metric definitions, with the parameters of each metric
*TimeseriesAPI* | [**ReadTimeseriesComplex**](docs/TimeseriesAPI.md#readtimeseriescomplex) | **Post** /timeseries/{timeseriesIdentifier} | Lists all available metric data points, matching the specified parameters
*TimeseriesAPI* | [**ReadTimeseriesData**](docs/TimeseriesAPI.md#readtimeseriesdata) | **Get** /timeseries/{timeseriesIdentifier} | Gets the parameters of the specified metric and, optionally, its data points
*TokensAPI* | [**CreateToken**](docs/TokensAPI.md#createtoken) | **Post** /tokens | Creates a new token
*TokensAPI* | [**DeleteToken**](docs/TokensAPI.md#deletetoken) | **Delete** /tokens/{id} | Deletes the specified token
*TokensAPI* | [**GetTokenMetadata**](docs/TokensAPI.md#gettokenmetadata) | **Get** /tokens/{id} | Lists token metadata by token ID
*TokensAPI* | [**GetTokenMetadataBySecret**](docs/TokensAPI.md#gettokenmetadatabysecret) | **Post** /tokens/lookup | Lists token metadata by token itself
*TokensAPI* | [**ListTokens**](docs/TokensAPI.md#listtokens) | **Get** /tokens | Lists available tokens in your environment
*TokensAPI* | [**UpdateToken**](docs/TokensAPI.md#updatetoken) | **Put** /tokens/{id} | Updates the specified token
*TopologySmartscapeApplicationAPI* | [**GetApplications**](docs/TopologySmartscapeApplicationAPI.md#getapplications) | **Get** /entity/applications | Gets the list of all applications in your environment along with their parameters
*TopologySmartscapeApplicationAPI* | [**GetBaselineValuesForSingleApplication**](docs/TopologySmartscapeApplicationAPI.md#getbaselinevaluesforsingleapplication) | **Get** /entity/applications/{meIdentifier}/baseline | Gets baseline data for the specified application | maturity&#x3D;EARLY_ADOPTER
*TopologySmartscapeApplicationAPI* | [**GetSingleApplication**](docs/TopologySmartscapeApplicationAPI.md#getsingleapplication) | **Get** /entity/applications/{meIdentifier} | Gets parameters of the specified application
*TopologySmartscapeApplicationAPI* | [**UpdateApplication**](docs/TopologySmartscapeApplicationAPI.md#updateapplication) | **Post** /entity/applications/{meIdentifier} | Updates parameters of the specified application
*TopologySmartscapeCustomDeviceAPI* | [**CreateCustomDataPoints**](docs/TopologySmartscapeCustomDeviceAPI.md#createcustomdatapoints) | **Post** /entity/infrastructure/custom/{customDeviceId} | Creates or updates a custom device, or reports metric data points to the custom device.
*TopologySmartscapeHostAPI* | [**GetHosts**](docs/TopologySmartscapeHostAPI.md#gethosts) | **Get** /entity/infrastructure/hosts | Lists all available hosts in your environment
*TopologySmartscapeHostAPI* | [**GetSingleHost**](docs/TopologySmartscapeHostAPI.md#getsinglehost) | **Get** /entity/infrastructure/hosts/{meIdentifier} | Gets parameters of the specified host
*TopologySmartscapeHostAPI* | [**RemoveTags**](docs/TopologySmartscapeHostAPI.md#removetags) | **Delete** /entity/infrastructure/hosts/{meIdentifier}/tags/{tag} | Remove tag of the specified host
*TopologySmartscapeHostAPI* | [**UpdateHost**](docs/TopologySmartscapeHostAPI.md#updatehost) | **Post** /entity/infrastructure/hosts/{meIdentifier} | Updates properties of the specified host
*TopologySmartscapeProcessAPI* | [**GetProcesses**](docs/TopologySmartscapeProcessAPI.md#getprocesses) | **Get** /entity/infrastructure/processes | Lists all monitored processes along with their parameters
*TopologySmartscapeProcessAPI* | [**GetSingleProcess**](docs/TopologySmartscapeProcessAPI.md#getsingleprocess) | **Get** /entity/infrastructure/processes/{meIdentifier} | List properties of the specified process
*TopologySmartscapeProcessGroupAPI* | [**GetProcessGroups**](docs/TopologySmartscapeProcessGroupAPI.md#getprocessgroups) | **Get** /entity/infrastructure/process-groups | Lists all process groups of your environment, along with their parameters
*TopologySmartscapeProcessGroupAPI* | [**GetSingleProcessGroup**](docs/TopologySmartscapeProcessGroupAPI.md#getsingleprocessgroup) | **Get** /entity/infrastructure/process-groups/{meIdentifier} | List properties of the specified process group
*TopologySmartscapeProcessGroupAPI* | [**UpdateProcessGroup**](docs/TopologySmartscapeProcessGroupAPI.md#updateprocessgroup) | **Post** /entity/infrastructure/process-groups/{meIdentifier} | Updates properties of the specified process group
*TopologySmartscapeServiceAPI* | [**GetBaselineValuesForSingleService**](docs/TopologySmartscapeServiceAPI.md#getbaselinevaluesforsingleservice) | **Get** /entity/services/{meIdentifier}/baseline | Gets baseline data for the specified service | maturity&#x3D;EARLY_ADOPTER
*TopologySmartscapeServiceAPI* | [**GetServices**](docs/TopologySmartscapeServiceAPI.md#getservices) | **Get** /entity/services | Lists all available services in your environment
*TopologySmartscapeServiceAPI* | [**GetSingleService**](docs/TopologySmartscapeServiceAPI.md#getsingleservice) | **Get** /entity/services/{meIdentifier} | Gets parameters of the specified service
*TopologySmartscapeServiceAPI* | [**UpdateService**](docs/TopologySmartscapeServiceAPI.md#updateservice) | **Post** /entity/services/{meIdentifier} | Updates parameters of the specified service


## Documentation For Models

 - [ActiveGateConnectionInfo](docs/ActiveGateConnectionInfo.md)
 - [ActiveGateInstallerVersions](docs/ActiveGateInstallerVersions.md)
 - [AgentInstallerVersions](docs/AgentInstallerVersions.md)
 - [AgentProcessModuleConfigResponse](docs/AgentProcessModuleConfigResponse.md)
 - [AgentVersion](docs/AgentVersion.md)
 - [AnomalyDetection](docs/AnomalyDetection.md)
 - [AnonymizationClusterRequestID](docs/AnonymizationClusterRequestID.md)
 - [AnonymizationIdResult](docs/AnonymizationIdResult.md)
 - [AnonymizationProgressResult](docs/AnonymizationProgressResult.md)
 - [Application](docs/Application.md)
 - [ApplicationBaselineValues](docs/ApplicationBaselineValues.md)
 - [ApplicationFromRelationships](docs/ApplicationFromRelationships.md)
 - [ApplicationToRelationships](docs/ApplicationToRelationships.md)
 - [BoshReleaseAvailableVersions](docs/BoshReleaseAvailableVersions.md)
 - [BoshReleaseChecksum](docs/BoshReleaseChecksum.md)
 - [BrowserSyntheticMonitor](docs/BrowserSyntheticMonitor.md)
 - [BrowserSyntheticMonitorUpdate](docs/BrowserSyntheticMonitorUpdate.md)
 - [ClusterVersion](docs/ClusterVersion.md)
 - [ConnectionInfo](docs/ConnectionInfo.md)
 - [ConstraintViolation](docs/ConstraintViolation.md)
 - [CreateToken](docs/CreateToken.md)
 - [CustomDevicePushMessage](docs/CustomDevicePushMessage.md)
 - [CustomDevicePushResult](docs/CustomDevicePushResult.md)
 - [DateProperty](docs/DateProperty.md)
 - [DoubleProperty](docs/DoubleProperty.md)
 - [Duration](docs/Duration.md)
 - [EntityBaselineData](docs/EntityBaselineData.md)
 - [EntityIdDto](docs/EntityIdDto.md)
 - [EntityShortRepresentation](docs/EntityShortRepresentation.md)
 - [EntityTimeseriesData](docs/EntityTimeseriesData.md)
 - [Error](docs/Error.md)
 - [ErrorEnvelope](docs/ErrorEnvelope.md)
 - [Event](docs/Event.md)
 - [EventCreation](docs/EventCreation.md)
 - [EventDto](docs/EventDto.md)
 - [EventQueryResult](docs/EventQueryResult.md)
 - [EventRestEntry](docs/EventRestEntry.md)
 - [EventRestImpact](docs/EventRestImpact.md)
 - [EventSeverity](docs/EventSeverity.md)
 - [EventStoreResult](docs/EventStoreResult.md)
 - [ExtractFields](docs/ExtractFields.md)
 - [FilterLogContent](docs/FilterLogContent.md)
 - [FilterTopLogRecords](docs/FilterTopLogRecords.md)
 - [GlobalOutagePolicy](docs/GlobalOutagePolicy.md)
 - [GlobalProblemStatus](docs/GlobalProblemStatus.md)
 - [GlobalProblemStatusOpenProblemCounts](docs/GlobalProblemStatusOpenProblemCounts.md)
 - [Host](docs/Host.md)
 - [Host4pg](docs/Host4pg.md)
 - [HostAgentInfo](docs/HostAgentInfo.md)
 - [HostFromRelationships](docs/HostFromRelationships.md)
 - [HostGroup](docs/HostGroup.md)
 - [HostToRelationships](docs/HostToRelationships.md)
 - [HostsListPage](docs/HostsListPage.md)
 - [HttpSyntheticMonitor](docs/HttpSyntheticMonitor.md)
 - [HttpSyntheticMonitorUpdate](docs/HttpSyntheticMonitorUpdate.md)
 - [InstallerMetaInfoDto](docs/InstallerMetaInfoDto.md)
 - [KeyPerformanceMetrics](docs/KeyPerformanceMetrics.md)
 - [LatestLambdaLayerNames](docs/LatestLambdaLayerNames.md)
 - [LoadingTimeThreshold](docs/LoadingTimeThreshold.md)
 - [LoadingTimeThresholdsPolicyDto](docs/LoadingTimeThresholdsPolicyDto.md)
 - [LocalOutagePolicy](docs/LocalOutagePolicy.md)
 - [LocationCollectionElement](docs/LocationCollectionElement.md)
 - [Log4host](docs/Log4host.md)
 - [LogFile4pg](docs/LogFile4pg.md)
 - [LogForCustomDevice](docs/LogForCustomDevice.md)
 - [LogJobDeleteResult](docs/LogJobDeleteResult.md)
 - [LogJobRecordsResult](docs/LogJobRecordsResult.md)
 - [LogJobRecordsTopValuesRestResult](docs/LogJobRecordsTopValuesRestResult.md)
 - [LogJobStatusResult](docs/LogJobStatusResult.md)
 - [LogList4hostResult](docs/LogList4hostResult.md)
 - [LogList4pgResult](docs/LogList4pgResult.md)
 - [LogListForCustomDeviceResult](docs/LogListForCustomDeviceResult.md)
 - [LogRecord](docs/LogRecord.md)
 - [LongProperty](docs/LongProperty.md)
 - [MaintenanceWindow](docs/MaintenanceWindow.md)
 - [MaintenanceWindowRecurrence](docs/MaintenanceWindowRecurrence.md)
 - [MaintenanceWindowSchedule](docs/MaintenanceWindowSchedule.md)
 - [MaintenanceWindowScope](docs/MaintenanceWindowScope.md)
 - [ManagementZone](docs/ManagementZone.md)
 - [ManualApplication](docs/ManualApplication.md)
 - [Model3rdPartyEventOpenNotification](docs/Model3rdPartyEventOpenNotification.md)
 - [Model3rdPartyEventResolvedNotification](docs/Model3rdPartyEventResolvedNotification.md)
 - [Model3rdPartySyntheticEvents](docs/Model3rdPartySyntheticEvents.md)
 - [Model3rdPartySyntheticLocation](docs/Model3rdPartySyntheticLocation.md)
 - [Model3rdPartySyntheticLocationTestResult](docs/Model3rdPartySyntheticLocationTestResult.md)
 - [Model3rdPartySyntheticMonitor](docs/Model3rdPartySyntheticMonitor.md)
 - [Model3rdPartySyntheticTestResult](docs/Model3rdPartySyntheticTestResult.md)
 - [Model3rdPartySyntheticTests](docs/Model3rdPartySyntheticTests.md)
 - [ModuleInfo](docs/ModuleInfo.md)
 - [ModuleInstance](docs/ModuleInstance.md)
 - [MonitorCollectionElement](docs/MonitorCollectionElement.md)
 - [MonitoredEntityFilter](docs/MonitoredEntityFilter.md)
 - [MonitoringState](docs/MonitoringState.md)
 - [Monitors](docs/Monitors.md)
 - [Node](docs/Node.md)
 - [NodeCollectionElement](docs/NodeCollectionElement.md)
 - [Nodes](docs/Nodes.md)
 - [Occurrence](docs/Occurrence.md)
 - [OneAgentInstallerChecksum](docs/OneAgentInstallerChecksum.md)
 - [OutageHandlingPolicy](docs/OutageHandlingPolicy.md)
 - [ParsingFieldTopValue](docs/ParsingFieldTopValue.md)
 - [PluginInfo](docs/PluginInfo.md)
 - [PluginInstance](docs/PluginInstance.md)
 - [PrivateSyntheticLocation](docs/PrivateSyntheticLocation.md)
 - [PrivateSyntheticLocationUpdate](docs/PrivateSyntheticLocationUpdate.md)
 - [Problem](docs/Problem.md)
 - [ProblemAffectedCounts](docs/ProblemAffectedCounts.md)
 - [ProblemCloseResult](docs/ProblemCloseResult.md)
 - [ProblemComment](docs/ProblemComment.md)
 - [ProblemCommentList](docs/ProblemCommentList.md)
 - [ProblemDetailsResultWrapper](docs/ProblemDetailsResultWrapper.md)
 - [ProblemFeedQueryResult](docs/ProblemFeedQueryResult.md)
 - [ProblemFeedQueryResultMonitored](docs/ProblemFeedQueryResultMonitored.md)
 - [ProblemFeedResultWrapper](docs/ProblemFeedResultWrapper.md)
 - [ProblemRecoveredCounts](docs/ProblemRecoveredCounts.md)
 - [ProblemStatusResultWrapper](docs/ProblemStatusResultWrapper.md)
 - [ProcessGroup](docs/ProcessGroup.md)
 - [ProcessGroupFromRelationships](docs/ProcessGroupFromRelationships.md)
 - [ProcessGroupInstance](docs/ProcessGroupInstance.md)
 - [ProcessGroupInstanceFromRelationships](docs/ProcessGroupInstanceFromRelationships.md)
 - [ProcessGroupInstanceModule](docs/ProcessGroupInstanceModule.md)
 - [ProcessGroupInstanceToRelationships](docs/ProcessGroupInstanceToRelationships.md)
 - [ProcessGroupMetadata](docs/ProcessGroupMetadata.md)
 - [ProcessGroupToRelationships](docs/ProcessGroupToRelationships.md)
 - [PublicSyntheticLocation](docs/PublicSyntheticLocation.md)
 - [PushEventAttachRules](docs/PushEventAttachRules.md)
 - [PushProblemComment](docs/PushProblemComment.md)
 - [RequestDto](docs/RequestDto.md)
 - [SectionProperty](docs/SectionProperty.md)
 - [Service](docs/Service.md)
 - [ServiceBaselineValues](docs/ServiceBaselineValues.md)
 - [ServiceFromRelationships](docs/ServiceFromRelationships.md)
 - [ServiceToRelationships](docs/ServiceToRelationships.md)
 - [SortingAttributes](docs/SortingAttributes.md)
 - [StateModification](docs/StateModification.md)
 - [StringProperty](docs/StringProperty.md)
 - [StubList](docs/StubList.md)
 - [SyntheticLocation](docs/SyntheticLocation.md)
 - [SyntheticLocationUpdate](docs/SyntheticLocationUpdate.md)
 - [SyntheticLocations](docs/SyntheticLocations.md)
 - [SyntheticMonitor](docs/SyntheticMonitor.md)
 - [SyntheticMonitorError](docs/SyntheticMonitorError.md)
 - [SyntheticMonitorStepResult](docs/SyntheticMonitorStepResult.md)
 - [SyntheticMonitorUpdate](docs/SyntheticMonitorUpdate.md)
 - [SyntheticPublicLocationUpdate](docs/SyntheticPublicLocationUpdate.md)
 - [SyntheticPublicLocationsStatus](docs/SyntheticPublicLocationsStatus.md)
 - [SyntheticTestLocation](docs/SyntheticTestLocation.md)
 - [SyntheticTestStep](docs/SyntheticTestStep.md)
 - [TagInfo](docs/TagInfo.md)
 - [TagMatchRule](docs/TagMatchRule.md)
 - [TagWithSourceInfo](docs/TagWithSourceInfo.md)
 - [TechnologyInfo](docs/TechnologyInfo.md)
 - [Threshold](docs/Threshold.md)
 - [ThresholdRegistrationMessage](docs/ThresholdRegistrationMessage.md)
 - [TimeseriesDataPointQueryResult](docs/TimeseriesDataPointQueryResult.md)
 - [TimeseriesDefinition](docs/TimeseriesDefinition.md)
 - [TimeseriesQueryMessage](docs/TimeseriesQueryMessage.md)
 - [TimeseriesQueryResult](docs/TimeseriesQueryResult.md)
 - [TimeseriesQueryResultWrapper](docs/TimeseriesQueryResultWrapper.md)
 - [TimeseriesRegistrationMessage](docs/TimeseriesRegistrationMessage.md)
 - [Token](docs/Token.md)
 - [TokenMetadata](docs/TokenMetadata.md)
 - [UniversalTag](docs/UniversalTag.md)
 - [UniversalTagKey](docs/UniversalTagKey.md)
 - [UpdateEntity](docs/UpdateEntity.md)
 - [UpdateToken](docs/UpdateToken.md)
 - [UserSession](docs/UserSession.md)
 - [UserSessionErrors](docs/UserSessionErrors.md)
 - [UserSessionEvents](docs/UserSessionEvents.md)
 - [UserSessionSyntheticEvent](docs/UserSessionSyntheticEvent.md)
 - [UserSessionUserAction](docs/UserSessionUserAction.md)
 - [UsqlResultAsTable](docs/UsqlResultAsTable.md)
 - [UsqlResultAsTree](docs/UsqlResultAsTree.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### Api-Token

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


