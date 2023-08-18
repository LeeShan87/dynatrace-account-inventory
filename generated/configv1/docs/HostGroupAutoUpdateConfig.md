# HostGroupAutoUpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveSetting** | Pointer to **string** | The actual state of the auto-update on the hosts in a host group.   Applicable only if the **setting** parameter is set to &#x60;INHERITED&#x60;. In that case the value is taken from the environment-wide setting. | [optional] [readonly] 
**EffectiveVersion** | Pointer to **string** | The actual version to which the OneAgent must be updated.   Applicable only if the **setting** parameter is set to &#x60;INHERITED&#x60;. In that case the value is taken from the environment-wide setting. | [optional] [readonly] 
**Id** | Pointer to **string** | The Dynatrace entity ID of the host group. | [optional] [readonly] 
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Setting** | **string** | The auto-update state of OneAgents in a host group:   * &#x60;ENABLED&#x60;: OneAgents automatically update to the most recent version.  * &#x60;DISABLED&#x60;: OneAgents update to the version specified in the **version** field. * &#x60;INHERITED&#x60;: The setting from the environment-wide configuration is used.  OneAgents installed on hosts of the host group use this configuration only when their **setting** parameter is set to &#x60;INHERITED&#x60;. | 
**TargetVersion** | Pointer to **string** | Version to update a OneAgent to when automatic updates are enabled.  Supports relative versions &#x60;latest&#x60;, &#x60;previous&#x60; and &#x60;older&#x60; as well as specific version in &#x60;&lt;major&gt;.&lt;minor&gt;&#x60; format (for example &#x60;1.261&#x60;) or &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;.&lt;timestamp&gt;&#x60; format (for example &#x60;1.261.178.20230313-090930&#x60;).  Only applicable when the **setting** parameter is set to &#x60;ENABLED&#x60;. | [optional] 
**UpdateWindows** | Pointer to [**UpdateWindowsConfig**](UpdateWindowsConfig.md) |  | [optional] 
**Version** | Pointer to **string** | The version to which the OneAgent must be updated.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60; format (for example &#x60;1.181.0&#x60;) or &#x60;&lt;major&gt;.&lt;minor&gt;&#x60; format (for example &#x60;1.181&#x60;). You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call. If no suitable installer is found for the provided version or the value is set to &#x60;null&#x60;, OneAgent won&#39;t be updated.   Only applicable when the **setting** parameter is set to &#x60;DISABLED&#x60;. | [optional] 

## Methods

### NewHostGroupAutoUpdateConfig

`func NewHostGroupAutoUpdateConfig(setting string, ) *HostGroupAutoUpdateConfig`

NewHostGroupAutoUpdateConfig instantiates a new HostGroupAutoUpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostGroupAutoUpdateConfigWithDefaults

`func NewHostGroupAutoUpdateConfigWithDefaults() *HostGroupAutoUpdateConfig`

NewHostGroupAutoUpdateConfigWithDefaults instantiates a new HostGroupAutoUpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveSetting

`func (o *HostGroupAutoUpdateConfig) GetEffectiveSetting() string`

GetEffectiveSetting returns the EffectiveSetting field if non-nil, zero value otherwise.

### GetEffectiveSettingOk

`func (o *HostGroupAutoUpdateConfig) GetEffectiveSettingOk() (*string, bool)`

GetEffectiveSettingOk returns a tuple with the EffectiveSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveSetting

`func (o *HostGroupAutoUpdateConfig) SetEffectiveSetting(v string)`

SetEffectiveSetting sets EffectiveSetting field to given value.

### HasEffectiveSetting

`func (o *HostGroupAutoUpdateConfig) HasEffectiveSetting() bool`

HasEffectiveSetting returns a boolean if a field has been set.

### GetEffectiveVersion

`func (o *HostGroupAutoUpdateConfig) GetEffectiveVersion() string`

GetEffectiveVersion returns the EffectiveVersion field if non-nil, zero value otherwise.

### GetEffectiveVersionOk

`func (o *HostGroupAutoUpdateConfig) GetEffectiveVersionOk() (*string, bool)`

GetEffectiveVersionOk returns a tuple with the EffectiveVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveVersion

`func (o *HostGroupAutoUpdateConfig) SetEffectiveVersion(v string)`

SetEffectiveVersion sets EffectiveVersion field to given value.

### HasEffectiveVersion

`func (o *HostGroupAutoUpdateConfig) HasEffectiveVersion() bool`

HasEffectiveVersion returns a boolean if a field has been set.

### GetId

`func (o *HostGroupAutoUpdateConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostGroupAutoUpdateConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostGroupAutoUpdateConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HostGroupAutoUpdateConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetadata

`func (o *HostGroupAutoUpdateConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *HostGroupAutoUpdateConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *HostGroupAutoUpdateConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *HostGroupAutoUpdateConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSetting

`func (o *HostGroupAutoUpdateConfig) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *HostGroupAutoUpdateConfig) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *HostGroupAutoUpdateConfig) SetSetting(v string)`

SetSetting sets Setting field to given value.


### GetTargetVersion

`func (o *HostGroupAutoUpdateConfig) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *HostGroupAutoUpdateConfig) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *HostGroupAutoUpdateConfig) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *HostGroupAutoUpdateConfig) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetUpdateWindows

`func (o *HostGroupAutoUpdateConfig) GetUpdateWindows() UpdateWindowsConfig`

GetUpdateWindows returns the UpdateWindows field if non-nil, zero value otherwise.

### GetUpdateWindowsOk

`func (o *HostGroupAutoUpdateConfig) GetUpdateWindowsOk() (*UpdateWindowsConfig, bool)`

GetUpdateWindowsOk returns a tuple with the UpdateWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateWindows

`func (o *HostGroupAutoUpdateConfig) SetUpdateWindows(v UpdateWindowsConfig)`

SetUpdateWindows sets UpdateWindows field to given value.

### HasUpdateWindows

`func (o *HostGroupAutoUpdateConfig) HasUpdateWindows() bool`

HasUpdateWindows returns a boolean if a field has been set.

### GetVersion

`func (o *HostGroupAutoUpdateConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HostGroupAutoUpdateConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HostGroupAutoUpdateConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HostGroupAutoUpdateConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


