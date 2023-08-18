# \OneAgentOnAHostAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHostsWithSpecificAgents**](OneAgentOnAHostAPI.md#GetHostsWithSpecificAgents) | **Get** /oneagents | Gets the list of hosts with OneAgent deployment information for each host



## GetHostsWithSpecificAgents

> HostsListPage GetHostsWithSpecificAgents(ctx).IncludeDetails(includeDetails).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).ManagementZoneId(managementZoneId).ManagementZone(managementZone).NetworkZoneId(networkZoneId).HostGroupId(hostGroupId).HostGroupName(hostGroupName).OsType(osType).CloudType(cloudType).AutoInjection(autoInjection).AvailabilityState(availabilityState).DetailedAvailabilityState(detailedAvailabilityState).MonitoringType(monitoringType).AgentVersionIs(agentVersionIs).AgentVersionNumber(agentVersionNumber).AutoUpdateSetting(autoUpdateSetting).UpdateStatus(updateStatus).FaultyVersion(faultyVersion).ActiveGateId(activeGateId).TechnologyModuleType(technologyModuleType).TechnologyModuleVersionIs(technologyModuleVersionIs).TechnologyModuleVersionNumber(technologyModuleVersionNumber).TechnologyModuleFaultyVersion(technologyModuleFaultyVersion).PluginName(pluginName).PluginVersionIs(pluginVersionIs).PluginVersionNumber(pluginVersionNumber).PluginState(pluginState).NextPageKey(nextPageKey).Execute()

Gets the list of hosts with OneAgent deployment information for each host



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
    includeDetails := true // bool | Includes (`true`) or excludes (`false`) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then `true` is used. (optional) (default to true)
    startTimestamp := int64(789) // int64 | The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used. (optional)
    endTimestamp := int64(789) // int64 | The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 7 months (214 days). (optional)
    relativeTime := "relativeTime_example" // string | The relative timeframe, back from now.   If you need to specify relative timeframe that is not presented in the list of possible values, specify the **startTimestamp** (up to 214 days back from now) and leave **endTimestamp** and **relativeTime** empty. (optional)
    tag := []string{"Inner_example"} // []string | Filters the resulting set of hosts by the specified tag. You can specify several tags in the following format: `tag=tag1&tag=tag2`. The host has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: `tag=[context]key:value`. For custom key-value tags, omit the context: `tag=key:value`. (optional)
    entity := []string{"Inner_example"} // []string | Filters result to the specified hosts only.    To specify several hosts use the following format: `entity=ID1&entity=ID2`. (optional)
    managementZoneId := int64(789) // int64 | Only return hosts that are part of the specified management zone.   Specify the management zone ID here. (optional)
    managementZone := "managementZone_example" // string | Only return hosts that are part of the specified management zone.   Specify the management zone name here.   If the **managementZoneId** parameter is set, this parameter is ignored. (optional)
    networkZoneId := "networkZoneId_example" // string | Filters the resulting set of hosts by the specified network zone.    Specify the Dynatrace entity ID of the required network zone. You can fetch the list of available network zones with the [GET all network zones](https://dt-url.net/u4o3r7z) call. (optional)
    hostGroupId := "hostGroupId_example" // string | Filters the resulting set of hosts by the specified host group.    Specify the Dynatrace entity ID of the required host group. (optional)
    hostGroupName := "hostGroupName_example" // string | Filters the resulting set of hosts by the specified host group.    Specify the name of the required host group. (optional)
    osType := "osType_example" // string | Filters the resulting set of hosts by the OS type. (optional)
    cloudType := "cloudType_example" // string | Filters the resulting set of hosts by the cloud type. (optional)
    autoInjection := "autoInjection_example" // string | Filters the resulting set of hosts by the auto-injection status. (optional)
    availabilityState := "availabilityState_example" // string | Filters the resulting set of hosts by the availability state of OneAgent.   * `MONITORED`: Hosts where OneAgent is enabled and active. * `UNMONITORED`: Hosts where OneAgent is disabled and inactive. * `CRASHED`: Hosts where OneAgent has returned a crash status code. * `LOST`: Hosts where it is impossible to establish connection with OneAgent. * `PRE_MONITORED`: Hosts where OneAgent is being initialized for monitoring. * `SHUTDOWN`: Hosts where OneAgent is shutting down in a controlled process. * `UNEXPECTED_SHUTDOWN`: Hosts where OneAgent is shutting down in an uncontrolled process. * `UNKNOWN`: Hosts where the state of OneAgent is unknown. (optional)
    detailedAvailabilityState := "detailedAvailabilityState_example" // string | Filters the resulting set of hosts by the detailed availability state of OneAgent.   * `UNKNOWN`: Hosts where the state of OneAgent is unknown. * `PRE_MONITORED`: Hosts where OneAgent is being initialized for monitoring. * `CRASHED_UNKNOWN`: Hosts where OneAgent has crashed for unknown reason. * `CRASHED_FAILURE`: Hosts where OneAgent has returned a crash status code. * `LOST_UNKNOWN`: Hosts where it is impossible to establish connection with OneAgent for unknown reason. * `LOST_CONNECTION`: Hosts where OneAgent has been recognized to be inactive. * `LOST_AGENT_UPGRADE_FAILED`: Hosts where OneAgent has a potential update problem due to inactivity after update. * `SHUTDOWN_UNKNOWN_UNEXPECTED`: Hosts where OneAgent is shutting down in an uncontrolled process. * `SHUTDOWN_UNKNOWN`: Hosts where OneAgent has shutdown for unknown reason. * `SHUTDOWN_GRACEFUL`: Hosts where OneAgent has shutdown because of host shutdown. * `SHUTDOWN_STOPPED`: Hosts where OneAgent has shutdown because the host has stopped. * `SHUTDOWN_AGENT_LOST`: Hosts where PaaS module has been recognized to be inactive. * `SHUTDOWN_SPOT_INSTANCE`: Hosts where OneAgent shutdown was triggered by the AWS Spot Instance interruption. * `UNMONITORED_UNKNOWN`: Hosts where OneAgent is disabled and inactive for unknown reason. * `UNMONITORED_TERMINATED`: Hosts where OneAgent has terminated. * `UNMONITORED_DISABLED`: Hosts where OneAgent has been disabled in configuration. * `UNMONITORED_AGENT_STOPPED`: Hosts where OneAgent is stopped. * `UNMONITORED_AGENT_RESTART_TRIGGERED`: Hosts where OneAgent is being restarted. * `UNMONITORED_AGENT_UNINSTALLED`: Hosts where OneAgent is uninstalled. * `UNMONITORED_AGENT_DISABLED`: Hosts where OneAgent reported that it was disabled. * `UNMONITORED_AGENT_UPGRADE_FAILED`: Hosts where OneAgent has a potential update problem. * `UNMONITORED_ID_CHANGED`: Hosts where OneAgent has potentially changed ID during update. * `UNMONITORED_AGENT_LOST`: Hosts where OneAgent has been recognized to be unavailable due to server communication issues. * `UNMONITORED_AGENT_UNREGISTERED`: Hosts where a code module has been recognized to be unavailable because of shutdown. * `UNMONITORED_AGENT_VERSION_REJECTED`: Hosts where OneAgent was rejected because the version does not meet the minimum agent version requirement. * `MONITORED`: Hosts where OneAgent is enabled and active. * `MONITORED_ENABLED`: Hosts where OneAgent has been enabled in configuration. * `MONITORED_AGENT_REGISTERED`: Hosts where the new OneAgent has been recognized. * `MONITORED_AGENT_UPGRADE_STARTED`: Hosts where OneAgent has shutdown due to an update. * `MONITORED_AGENT_ENABLED`: Hosts where OneAgent reported that it was enabled. * `MONITORED_AGENT_VERSION_ACCEPTED`: Hosts where OneAgent was accepted because the version meets the minimum agent version requirement. (optional)
    monitoringType := "monitoringType_example" // string | Filters the resulting set of hosts by monitoring mode of OneAgent deployed on the host. (optional)
    agentVersionIs := "agentVersionIs_example" // string | Filters the resulting set of hosts to those that have a certain OneAgent version deployed on the host.   Specify the comparison operator here. (optional)
    agentVersionNumber := "agentVersionNumber_example" // string | Filters the resulting set of hosts to those that have a certain OneAgent version deployed on the host.   Specify the version in the `<major>.<minor>.<revision>` format, for example `1.182.0`. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call. (optional)
    autoUpdateSetting := "autoUpdateSetting_example" // string | Filters the resulting set of hosts by the actual state of the auto-update setting of deployed OneAgents. (optional)
    updateStatus := "updateStatus_example" // string | Filters the resulting set of hosts by the update status of OneAgent deployed on the host. (optional)
    faultyVersion := true // bool | Filters the resulting set of hosts to those that run OneAgent version that is marked as faulty. (optional)
    activeGateId := "activeGateId_example" // string | Filters the resulting set of hosts to those that are currently connected to the ActiveGate with the specified ID.   Use **DIRECT_COMMUNICATION** keyword to find the hosts not connected to any ActiveGate. (optional)
    technologyModuleType := "technologyModuleType_example" // string | Filters the resulting set of hosts to those that run the specified OneAgent code module.   If several code module filters are specified, the code module has to match **all** the filters. (optional)
    technologyModuleVersionIs := "technologyModuleVersionIs_example" // string | Filters the resulting set of hosts to those that have a certain code module version deployed on the host.   Specify the comparison operator here.   If several code module filters are specified, the code module has to match **all** the filters. (optional)
    technologyModuleVersionNumber := "technologyModuleVersionNumber_example" // string | Filters the resulting set of hosts to those that have a certain code module version deployed on the host.   Specify the version in the `<major>.<minor>.<revision>` format, for example `1.182.0`. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call.   If several code module filters are specified, the code module has to match **all** the filters. (optional)
    technologyModuleFaultyVersion := true // bool | Filters the resulting set of hosts to those that run the code module version that is marked as faulty.   If several code module filters are specified, the code module has to match **all** the filters. (optional)
    pluginName := "pluginName_example" // string | Filters the resulting set of hosts to those that run the plugin with the specified name.   The **CONTAINS** operator is applied to the specified value.   If several plugin filters are specified, the plugin has to match **all** the filters. (optional)
    pluginVersionIs := "pluginVersionIs_example" // string | Filters the resulting set of hosts to those that have a certain plugin version deployed on the host.   Specify the comparison operator here.   If several plugin filters are specified, the plugin has to match **all** the filters. (optional)
    pluginVersionNumber := "pluginVersionNumber_example" // string | Filters the resulting set of hosts to those that have a certain plugin version deployed on the host.   Specify the version in the `<major>.<minor>.<revision>` format, for example `1.182.0`. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call.   `<minor>` and `<revision>` parts of the version number are optional.   If several plugin filters are specified, the plugin has to match **all** the filters. (optional)
    pluginState := "pluginState_example" // string | Filters the resulting set of hosts to those that run the plugin with the specified state. (optional)
    nextPageKey := "nextPageKey_example" // string | The cursor for the next page of results, if results do not fit on one page. You can find the cursor value on the current page of the response, in the **nextPageKey** field.   To obtain subsequent pages, you must specify this cursor value in your query, and keep all other query parameters as they were in the original request.   If you don't specify the cursor, the first page will always be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OneAgentOnAHostAPI.GetHostsWithSpecificAgents(context.Background()).IncludeDetails(includeDetails).StartTimestamp(startTimestamp).EndTimestamp(endTimestamp).RelativeTime(relativeTime).Tag(tag).Entity(entity).ManagementZoneId(managementZoneId).ManagementZone(managementZone).NetworkZoneId(networkZoneId).HostGroupId(hostGroupId).HostGroupName(hostGroupName).OsType(osType).CloudType(cloudType).AutoInjection(autoInjection).AvailabilityState(availabilityState).DetailedAvailabilityState(detailedAvailabilityState).MonitoringType(monitoringType).AgentVersionIs(agentVersionIs).AgentVersionNumber(agentVersionNumber).AutoUpdateSetting(autoUpdateSetting).UpdateStatus(updateStatus).FaultyVersion(faultyVersion).ActiveGateId(activeGateId).TechnologyModuleType(technologyModuleType).TechnologyModuleVersionIs(technologyModuleVersionIs).TechnologyModuleVersionNumber(technologyModuleVersionNumber).TechnologyModuleFaultyVersion(technologyModuleFaultyVersion).PluginName(pluginName).PluginVersionIs(pluginVersionIs).PluginVersionNumber(pluginVersionNumber).PluginState(pluginState).NextPageKey(nextPageKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OneAgentOnAHostAPI.GetHostsWithSpecificAgents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHostsWithSpecificAgents`: HostsListPage
    fmt.Fprintf(os.Stdout, "Response from `OneAgentOnAHostAPI.GetHostsWithSpecificAgents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetHostsWithSpecificAgentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeDetails** | **bool** | Includes (&#x60;true&#x60;) or excludes (&#x60;false&#x60;) details which are queried from related entities.  Excluding details may make queries faster.   If not set, then &#x60;true&#x60; is used. | [default to true]
 **startTimestamp** | **int64** | The start timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then 72 hours behind from now is used. | 
 **endTimestamp** | **int64** | The end timestamp of the requested timeframe, in milliseconds (UTC).   If not set, then the current timestamp is used.   The timeframe must not exceed 7 months (214 days). | 
 **relativeTime** | **string** | The relative timeframe, back from now.   If you need to specify relative timeframe that is not presented in the list of possible values, specify the **startTimestamp** (up to 214 days back from now) and leave **endTimestamp** and **relativeTime** empty. | 
 **tag** | **[]string** | Filters the resulting set of hosts by the specified tag. You can specify several tags in the following format: &#x60;tag&#x3D;tag1&amp;tag&#x3D;tag2&#x60;. The host has to match **all** the specified tags.   In case of key-value tags, such as imported AWS or CloudFoundry tags, use the following format: &#x60;tag&#x3D;[context]key:value&#x60;. For custom key-value tags, omit the context: &#x60;tag&#x3D;key:value&#x60;. | 
 **entity** | **[]string** | Filters result to the specified hosts only.    To specify several hosts use the following format: &#x60;entity&#x3D;ID1&amp;entity&#x3D;ID2&#x60;. | 
 **managementZoneId** | **int64** | Only return hosts that are part of the specified management zone.   Specify the management zone ID here. | 
 **managementZone** | **string** | Only return hosts that are part of the specified management zone.   Specify the management zone name here.   If the **managementZoneId** parameter is set, this parameter is ignored. | 
 **networkZoneId** | **string** | Filters the resulting set of hosts by the specified network zone.    Specify the Dynatrace entity ID of the required network zone. You can fetch the list of available network zones with the [GET all network zones](https://dt-url.net/u4o3r7z) call. | 
 **hostGroupId** | **string** | Filters the resulting set of hosts by the specified host group.    Specify the Dynatrace entity ID of the required host group. | 
 **hostGroupName** | **string** | Filters the resulting set of hosts by the specified host group.    Specify the name of the required host group. | 
 **osType** | **string** | Filters the resulting set of hosts by the OS type. | 
 **cloudType** | **string** | Filters the resulting set of hosts by the cloud type. | 
 **autoInjection** | **string** | Filters the resulting set of hosts by the auto-injection status. | 
 **availabilityState** | **string** | Filters the resulting set of hosts by the availability state of OneAgent.   * &#x60;MONITORED&#x60;: Hosts where OneAgent is enabled and active. * &#x60;UNMONITORED&#x60;: Hosts where OneAgent is disabled and inactive. * &#x60;CRASHED&#x60;: Hosts where OneAgent has returned a crash status code. * &#x60;LOST&#x60;: Hosts where it is impossible to establish connection with OneAgent. * &#x60;PRE_MONITORED&#x60;: Hosts where OneAgent is being initialized for monitoring. * &#x60;SHUTDOWN&#x60;: Hosts where OneAgent is shutting down in a controlled process. * &#x60;UNEXPECTED_SHUTDOWN&#x60;: Hosts where OneAgent is shutting down in an uncontrolled process. * &#x60;UNKNOWN&#x60;: Hosts where the state of OneAgent is unknown. | 
 **detailedAvailabilityState** | **string** | Filters the resulting set of hosts by the detailed availability state of OneAgent.   * &#x60;UNKNOWN&#x60;: Hosts where the state of OneAgent is unknown. * &#x60;PRE_MONITORED&#x60;: Hosts where OneAgent is being initialized for monitoring. * &#x60;CRASHED_UNKNOWN&#x60;: Hosts where OneAgent has crashed for unknown reason. * &#x60;CRASHED_FAILURE&#x60;: Hosts where OneAgent has returned a crash status code. * &#x60;LOST_UNKNOWN&#x60;: Hosts where it is impossible to establish connection with OneAgent for unknown reason. * &#x60;LOST_CONNECTION&#x60;: Hosts where OneAgent has been recognized to be inactive. * &#x60;LOST_AGENT_UPGRADE_FAILED&#x60;: Hosts where OneAgent has a potential update problem due to inactivity after update. * &#x60;SHUTDOWN_UNKNOWN_UNEXPECTED&#x60;: Hosts where OneAgent is shutting down in an uncontrolled process. * &#x60;SHUTDOWN_UNKNOWN&#x60;: Hosts where OneAgent has shutdown for unknown reason. * &#x60;SHUTDOWN_GRACEFUL&#x60;: Hosts where OneAgent has shutdown because of host shutdown. * &#x60;SHUTDOWN_STOPPED&#x60;: Hosts where OneAgent has shutdown because the host has stopped. * &#x60;SHUTDOWN_AGENT_LOST&#x60;: Hosts where PaaS module has been recognized to be inactive. * &#x60;SHUTDOWN_SPOT_INSTANCE&#x60;: Hosts where OneAgent shutdown was triggered by the AWS Spot Instance interruption. * &#x60;UNMONITORED_UNKNOWN&#x60;: Hosts where OneAgent is disabled and inactive for unknown reason. * &#x60;UNMONITORED_TERMINATED&#x60;: Hosts where OneAgent has terminated. * &#x60;UNMONITORED_DISABLED&#x60;: Hosts where OneAgent has been disabled in configuration. * &#x60;UNMONITORED_AGENT_STOPPED&#x60;: Hosts where OneAgent is stopped. * &#x60;UNMONITORED_AGENT_RESTART_TRIGGERED&#x60;: Hosts where OneAgent is being restarted. * &#x60;UNMONITORED_AGENT_UNINSTALLED&#x60;: Hosts where OneAgent is uninstalled. * &#x60;UNMONITORED_AGENT_DISABLED&#x60;: Hosts where OneAgent reported that it was disabled. * &#x60;UNMONITORED_AGENT_UPGRADE_FAILED&#x60;: Hosts where OneAgent has a potential update problem. * &#x60;UNMONITORED_ID_CHANGED&#x60;: Hosts where OneAgent has potentially changed ID during update. * &#x60;UNMONITORED_AGENT_LOST&#x60;: Hosts where OneAgent has been recognized to be unavailable due to server communication issues. * &#x60;UNMONITORED_AGENT_UNREGISTERED&#x60;: Hosts where a code module has been recognized to be unavailable because of shutdown. * &#x60;UNMONITORED_AGENT_VERSION_REJECTED&#x60;: Hosts where OneAgent was rejected because the version does not meet the minimum agent version requirement. * &#x60;MONITORED&#x60;: Hosts where OneAgent is enabled and active. * &#x60;MONITORED_ENABLED&#x60;: Hosts where OneAgent has been enabled in configuration. * &#x60;MONITORED_AGENT_REGISTERED&#x60;: Hosts where the new OneAgent has been recognized. * &#x60;MONITORED_AGENT_UPGRADE_STARTED&#x60;: Hosts where OneAgent has shutdown due to an update. * &#x60;MONITORED_AGENT_ENABLED&#x60;: Hosts where OneAgent reported that it was enabled. * &#x60;MONITORED_AGENT_VERSION_ACCEPTED&#x60;: Hosts where OneAgent was accepted because the version meets the minimum agent version requirement. | 
 **monitoringType** | **string** | Filters the resulting set of hosts by monitoring mode of OneAgent deployed on the host. | 
 **agentVersionIs** | **string** | Filters the resulting set of hosts to those that have a certain OneAgent version deployed on the host.   Specify the comparison operator here. | 
 **agentVersionNumber** | **string** | Filters the resulting set of hosts to those that have a certain OneAgent version deployed on the host.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60; format, for example &#x60;1.182.0&#x60;. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call. | 
 **autoUpdateSetting** | **string** | Filters the resulting set of hosts by the actual state of the auto-update setting of deployed OneAgents. | 
 **updateStatus** | **string** | Filters the resulting set of hosts by the update status of OneAgent deployed on the host. | 
 **faultyVersion** | **bool** | Filters the resulting set of hosts to those that run OneAgent version that is marked as faulty. | 
 **activeGateId** | **string** | Filters the resulting set of hosts to those that are currently connected to the ActiveGate with the specified ID.   Use **DIRECT_COMMUNICATION** keyword to find the hosts not connected to any ActiveGate. | 
 **technologyModuleType** | **string** | Filters the resulting set of hosts to those that run the specified OneAgent code module.   If several code module filters are specified, the code module has to match **all** the filters. | 
 **technologyModuleVersionIs** | **string** | Filters the resulting set of hosts to those that have a certain code module version deployed on the host.   Specify the comparison operator here.   If several code module filters are specified, the code module has to match **all** the filters. | 
 **technologyModuleVersionNumber** | **string** | Filters the resulting set of hosts to those that have a certain code module version deployed on the host.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60; format, for example &#x60;1.182.0&#x60;. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call.   If several code module filters are specified, the code module has to match **all** the filters. | 
 **technologyModuleFaultyVersion** | **bool** | Filters the resulting set of hosts to those that run the code module version that is marked as faulty.   If several code module filters are specified, the code module has to match **all** the filters. | 
 **pluginName** | **string** | Filters the resulting set of hosts to those that run the plugin with the specified name.   The **CONTAINS** operator is applied to the specified value.   If several plugin filters are specified, the plugin has to match **all** the filters. | 
 **pluginVersionIs** | **string** | Filters the resulting set of hosts to those that have a certain plugin version deployed on the host.   Specify the comparison operator here.   If several plugin filters are specified, the plugin has to match **all** the filters. | 
 **pluginVersionNumber** | **string** | Filters the resulting set of hosts to those that have a certain plugin version deployed on the host.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60; format, for example &#x60;1.182.0&#x60;. You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call.   &#x60;&lt;minor&gt;&#x60; and &#x60;&lt;revision&gt;&#x60; parts of the version number are optional.   If several plugin filters are specified, the plugin has to match **all** the filters. | 
 **pluginState** | **string** | Filters the resulting set of hosts to those that run the plugin with the specified state. | 
 **nextPageKey** | **string** | The cursor for the next page of results, if results do not fit on one page. You can find the cursor value on the current page of the response, in the **nextPageKey** field.   To obtain subsequent pages, you must specify this cursor value in your query, and keep all other query parameters as they were in the original request.   If you don&#39;t specify the cursor, the first page will always be returned. | 

### Return type

[**HostsListPage**](HostsListPage.md)

### Authorization

[Api-Token](../README.md#Api-Token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

