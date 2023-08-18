# EnvironmentAutoUpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Setting** | **string** | The auto-update state of OneAgents connecting to the environment:   * &#x60;ENABLED&#x60;: OneAgents automatically update to the most recent version.  * &#x60;DISABLED&#x60;: OneAgents update to the version specified in the **version** field.  OneAgents that connect to the environment use this configuration only when their **setting** parameter is set to &#x60;INHERITED&#x60;. | 
**TargetVersion** | Pointer to **string** | Version to update a OneAgent to when automatic updates are enabled.  Supports relative versions &#x60;latest&#x60;, &#x60;previous&#x60; and &#x60;older&#x60; as well as specific version in &#x60;&lt;major&gt;.&lt;minor&gt;&#x60; format (for example &#x60;1.261&#x60;) or &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;.&lt;timestamp&gt;&#x60; format (for example &#x60;1.261.178.20230313-090930&#x60;).  Only applicable when the **setting** parameter is set to &#x60;ENABLED&#x60;. | [optional] 
**UpdateWindows** | Pointer to [**UpdateWindowsConfig**](UpdateWindowsConfig.md) |  | [optional] 
**Version** | Pointer to **string** | The version to which the OneAgent must be updated.   Specify the version in the &#x60;&lt;major&gt;.&lt;minor&gt;.&lt;revision&gt;&#x60; format (for example &#x60;1.181.0&#x60;) or &#x60;&lt;major&gt;.&lt;minor&gt;&#x60; format (for example &#x60;1.181&#x60;). You can fetch the list of available versions with the [GET available versions](https://dt-url.net/fo23rb5) call. If no suitable installer is found for the provided version or the value is set to &#x60;null&#x60;, OneAgent won&#39;t be updated.   Only applicable when the **setting** parameter is set to &#x60;DISABLED&#x60;. | [optional] 

## Methods

### NewEnvironmentAutoUpdateConfig

`func NewEnvironmentAutoUpdateConfig(setting string, ) *EnvironmentAutoUpdateConfig`

NewEnvironmentAutoUpdateConfig instantiates a new EnvironmentAutoUpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentAutoUpdateConfigWithDefaults

`func NewEnvironmentAutoUpdateConfigWithDefaults() *EnvironmentAutoUpdateConfig`

NewEnvironmentAutoUpdateConfigWithDefaults instantiates a new EnvironmentAutoUpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *EnvironmentAutoUpdateConfig) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnvironmentAutoUpdateConfig) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnvironmentAutoUpdateConfig) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnvironmentAutoUpdateConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSetting

`func (o *EnvironmentAutoUpdateConfig) GetSetting() string`

GetSetting returns the Setting field if non-nil, zero value otherwise.

### GetSettingOk

`func (o *EnvironmentAutoUpdateConfig) GetSettingOk() (*string, bool)`

GetSettingOk returns a tuple with the Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetting

`func (o *EnvironmentAutoUpdateConfig) SetSetting(v string)`

SetSetting sets Setting field to given value.


### GetTargetVersion

`func (o *EnvironmentAutoUpdateConfig) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *EnvironmentAutoUpdateConfig) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *EnvironmentAutoUpdateConfig) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *EnvironmentAutoUpdateConfig) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetUpdateWindows

`func (o *EnvironmentAutoUpdateConfig) GetUpdateWindows() UpdateWindowsConfig`

GetUpdateWindows returns the UpdateWindows field if non-nil, zero value otherwise.

### GetUpdateWindowsOk

`func (o *EnvironmentAutoUpdateConfig) GetUpdateWindowsOk() (*UpdateWindowsConfig, bool)`

GetUpdateWindowsOk returns a tuple with the UpdateWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateWindows

`func (o *EnvironmentAutoUpdateConfig) SetUpdateWindows(v UpdateWindowsConfig)`

SetUpdateWindows sets UpdateWindows field to given value.

### HasUpdateWindows

`func (o *EnvironmentAutoUpdateConfig) HasUpdateWindows() bool`

HasUpdateWindows returns a boolean if a field has been set.

### GetVersion

`func (o *EnvironmentAutoUpdateConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *EnvironmentAutoUpdateConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *EnvironmentAutoUpdateConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *EnvironmentAutoUpdateConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


