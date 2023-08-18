# PrivateSyntheticLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpdateChromium** | Pointer to **bool** | Auto upgrade of Chromium is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**AvailabilityLocationOutage** | Pointer to **bool** | Alerting for location outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). Supported only for private Synthetic locations. | [optional] 
**AvailabilityNodeOutage** | Pointer to **bool** | Alerting for node outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). \\n\\n If enabled, the outage of *any* node in the location triggers an alert. Supported only for private Synthetic locations. | [optional] 
**AvailabilityNotificationsEnabled** | Pointer to **bool** | Notifications for location and node outage are enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). Supported only for private Synthetic locations. | [optional] 
**DeploymentType** | Pointer to **string** | The deployment type of the location:   * &#x60;STANDARD&#x60;: The location is deployed on Windows or Linux. * &#x60;KUBERNETES&#x60;: The location is deployed on Kubernetes. | [optional] 
**LocationNodeOutageDelayInMinutes** | Pointer to **int32** | Alert if location or node outage lasts longer than *X* minutes. \\n\\n Only applicable when &#x60;availabilityLocationOutage&#x60; or &#x60;availabilityNodeOutage&#x60; is set to &#x60;true&#x60;. Supported only for private Synthetic locations. | [optional] 
**Nodes** | **[]string** | A list of synthetic nodes belonging to the location.    You can retrieve the list of available nodes with the [GET all nodes](https://dt-url.net/miy3rpl) call. | 
**UseNewKubernetesVersion** | Pointer to **bool** | Boolean value describes which kubernetes version will be used:  * &#x60;false&#x60;: Version 1.23+ that is older than 1.26 * &#x60;true&#x60;: Version 1.26+. | [optional] 

## Methods

### NewPrivateSyntheticLocation

`func NewPrivateSyntheticLocation(nodes []string, ) *PrivateSyntheticLocation`

NewPrivateSyntheticLocation instantiates a new PrivateSyntheticLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateSyntheticLocationWithDefaults

`func NewPrivateSyntheticLocationWithDefaults() *PrivateSyntheticLocation`

NewPrivateSyntheticLocationWithDefaults instantiates a new PrivateSyntheticLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpdateChromium

`func (o *PrivateSyntheticLocation) GetAutoUpdateChromium() bool`

GetAutoUpdateChromium returns the AutoUpdateChromium field if non-nil, zero value otherwise.

### GetAutoUpdateChromiumOk

`func (o *PrivateSyntheticLocation) GetAutoUpdateChromiumOk() (*bool, bool)`

GetAutoUpdateChromiumOk returns a tuple with the AutoUpdateChromium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateChromium

`func (o *PrivateSyntheticLocation) SetAutoUpdateChromium(v bool)`

SetAutoUpdateChromium sets AutoUpdateChromium field to given value.

### HasAutoUpdateChromium

`func (o *PrivateSyntheticLocation) HasAutoUpdateChromium() bool`

HasAutoUpdateChromium returns a boolean if a field has been set.

### GetAvailabilityLocationOutage

`func (o *PrivateSyntheticLocation) GetAvailabilityLocationOutage() bool`

GetAvailabilityLocationOutage returns the AvailabilityLocationOutage field if non-nil, zero value otherwise.

### GetAvailabilityLocationOutageOk

`func (o *PrivateSyntheticLocation) GetAvailabilityLocationOutageOk() (*bool, bool)`

GetAvailabilityLocationOutageOk returns a tuple with the AvailabilityLocationOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityLocationOutage

`func (o *PrivateSyntheticLocation) SetAvailabilityLocationOutage(v bool)`

SetAvailabilityLocationOutage sets AvailabilityLocationOutage field to given value.

### HasAvailabilityLocationOutage

`func (o *PrivateSyntheticLocation) HasAvailabilityLocationOutage() bool`

HasAvailabilityLocationOutage returns a boolean if a field has been set.

### GetAvailabilityNodeOutage

`func (o *PrivateSyntheticLocation) GetAvailabilityNodeOutage() bool`

GetAvailabilityNodeOutage returns the AvailabilityNodeOutage field if non-nil, zero value otherwise.

### GetAvailabilityNodeOutageOk

`func (o *PrivateSyntheticLocation) GetAvailabilityNodeOutageOk() (*bool, bool)`

GetAvailabilityNodeOutageOk returns a tuple with the AvailabilityNodeOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityNodeOutage

`func (o *PrivateSyntheticLocation) SetAvailabilityNodeOutage(v bool)`

SetAvailabilityNodeOutage sets AvailabilityNodeOutage field to given value.

### HasAvailabilityNodeOutage

`func (o *PrivateSyntheticLocation) HasAvailabilityNodeOutage() bool`

HasAvailabilityNodeOutage returns a boolean if a field has been set.

### GetAvailabilityNotificationsEnabled

`func (o *PrivateSyntheticLocation) GetAvailabilityNotificationsEnabled() bool`

GetAvailabilityNotificationsEnabled returns the AvailabilityNotificationsEnabled field if non-nil, zero value otherwise.

### GetAvailabilityNotificationsEnabledOk

`func (o *PrivateSyntheticLocation) GetAvailabilityNotificationsEnabledOk() (*bool, bool)`

GetAvailabilityNotificationsEnabledOk returns a tuple with the AvailabilityNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityNotificationsEnabled

`func (o *PrivateSyntheticLocation) SetAvailabilityNotificationsEnabled(v bool)`

SetAvailabilityNotificationsEnabled sets AvailabilityNotificationsEnabled field to given value.

### HasAvailabilityNotificationsEnabled

`func (o *PrivateSyntheticLocation) HasAvailabilityNotificationsEnabled() bool`

HasAvailabilityNotificationsEnabled returns a boolean if a field has been set.

### GetDeploymentType

`func (o *PrivateSyntheticLocation) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *PrivateSyntheticLocation) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *PrivateSyntheticLocation) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *PrivateSyntheticLocation) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetLocationNodeOutageDelayInMinutes

`func (o *PrivateSyntheticLocation) GetLocationNodeOutageDelayInMinutes() int32`

GetLocationNodeOutageDelayInMinutes returns the LocationNodeOutageDelayInMinutes field if non-nil, zero value otherwise.

### GetLocationNodeOutageDelayInMinutesOk

`func (o *PrivateSyntheticLocation) GetLocationNodeOutageDelayInMinutesOk() (*int32, bool)`

GetLocationNodeOutageDelayInMinutesOk returns a tuple with the LocationNodeOutageDelayInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationNodeOutageDelayInMinutes

`func (o *PrivateSyntheticLocation) SetLocationNodeOutageDelayInMinutes(v int32)`

SetLocationNodeOutageDelayInMinutes sets LocationNodeOutageDelayInMinutes field to given value.

### HasLocationNodeOutageDelayInMinutes

`func (o *PrivateSyntheticLocation) HasLocationNodeOutageDelayInMinutes() bool`

HasLocationNodeOutageDelayInMinutes returns a boolean if a field has been set.

### GetNodes

`func (o *PrivateSyntheticLocation) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *PrivateSyntheticLocation) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *PrivateSyntheticLocation) SetNodes(v []string)`

SetNodes sets Nodes field to given value.


### GetUseNewKubernetesVersion

`func (o *PrivateSyntheticLocation) GetUseNewKubernetesVersion() bool`

GetUseNewKubernetesVersion returns the UseNewKubernetesVersion field if non-nil, zero value otherwise.

### GetUseNewKubernetesVersionOk

`func (o *PrivateSyntheticLocation) GetUseNewKubernetesVersionOk() (*bool, bool)`

GetUseNewKubernetesVersionOk returns a tuple with the UseNewKubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNewKubernetesVersion

`func (o *PrivateSyntheticLocation) SetUseNewKubernetesVersion(v bool)`

SetUseNewKubernetesVersion sets UseNewKubernetesVersion field to given value.

### HasUseNewKubernetesVersion

`func (o *PrivateSyntheticLocation) HasUseNewKubernetesVersion() bool`

HasUseNewKubernetesVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


