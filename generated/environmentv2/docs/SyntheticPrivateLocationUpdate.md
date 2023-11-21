# SyntheticPrivateLocationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoUpdateChromium** | Pointer to **bool** | Auto upgrade of Chromium is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**AvailabilityLocationOutage** | Pointer to **bool** | Alerting for location outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). Supported only for private Synthetic locations. | [optional] 
**AvailabilityNodeOutage** | Pointer to **bool** | Alerting for node outage is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). \\n\\n If enabled, the outage of *any* node in the location triggers an alert. Supported only for private Synthetic locations. | [optional] 
**AvailabilityNotificationsEnabled** | Pointer to **bool** | Notifications for location and node outage are enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). Supported only for private Synthetic locations. | [optional] 
**City** | Pointer to **string** | The city of the location. | [optional] 
**CountryCode** | Pointer to **string** | The country code of the location.    To fetch the list of available country codes, use the [GET all countries](https://dt-url.net/37030go) request. | [optional] 
**DeploymentType** | Pointer to **string** | The deployment type of the location:   * &#x60;STANDARD&#x60;: The location is deployed on Windows or Linux. * &#x60;KUBERNETES&#x60;: The location is deployed on Kubernetes. | [optional] 
**Latitude** | **float64** | The latitude of the location in &#x60;DDD.dddd&#x60; format. | 
**LocationNodeOutageDelayInMinutes** | Pointer to **int32** | Alert if location or node outage lasts longer than *X* minutes. \\n\\n Only applicable when &#x60;availabilityLocationOutage&#x60; or &#x60;availabilityNodeOutage&#x60; is set to &#x60;true&#x60;. Supported only for private Synthetic locations. | [optional] 
**Longitude** | **float64** | The longitude of the location in &#x60;DDD.dddd&#x60; format. | 
**Name** | **string** | The name of the location. | 
**Nodes** | **[]string** | A list of synthetic nodes belonging to the location.    You can retrieve the list of available nodes with the [GET all nodes](https://dt-url.net/miy3rpl) call. | 
**RegionCode** | Pointer to **string** | The region code of the location.    To fetch the list of available region codes, use the [GET regions of the country](https://dt-url.net/az230x0) request. | [optional] 
**Status** | Pointer to **string** | The status of the location:   * &#x60;ENABLED&#x60;: The location is displayed as active in the UI. You can assign monitors to the location.  * &#x60;DISABLED&#x60;: The location is displayed as inactive in the UI. You can&#39;t assign monitors to the location. Monitors already assigned to the location will stay there and will be executed from the location.  * &#x60;HIDDEN&#x60;: The location is not displayed in the UI. You can&#39;t assign monitors to the location. You can only set location as &#x60;HIDDEN&#x60; when no monitor is assigned to it. | [optional] 
**UseNewKubernetesVersion** | Pointer to **bool** | Boolean value describes which kubernetes version will be used:  * &#x60;false&#x60;: Version 1.23+ that is older than 1.26 * &#x60;true&#x60;: Version 1.26+. | [optional] 

## Methods

### NewSyntheticPrivateLocationUpdate

`func NewSyntheticPrivateLocationUpdate(latitude float64, longitude float64, name string, nodes []string, ) *SyntheticPrivateLocationUpdate`

NewSyntheticPrivateLocationUpdate instantiates a new SyntheticPrivateLocationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticPrivateLocationUpdateWithDefaults

`func NewSyntheticPrivateLocationUpdateWithDefaults() *SyntheticPrivateLocationUpdate`

NewSyntheticPrivateLocationUpdateWithDefaults instantiates a new SyntheticPrivateLocationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoUpdateChromium

`func (o *SyntheticPrivateLocationUpdate) GetAutoUpdateChromium() bool`

GetAutoUpdateChromium returns the AutoUpdateChromium field if non-nil, zero value otherwise.

### GetAutoUpdateChromiumOk

`func (o *SyntheticPrivateLocationUpdate) GetAutoUpdateChromiumOk() (*bool, bool)`

GetAutoUpdateChromiumOk returns a tuple with the AutoUpdateChromium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdateChromium

`func (o *SyntheticPrivateLocationUpdate) SetAutoUpdateChromium(v bool)`

SetAutoUpdateChromium sets AutoUpdateChromium field to given value.

### HasAutoUpdateChromium

`func (o *SyntheticPrivateLocationUpdate) HasAutoUpdateChromium() bool`

HasAutoUpdateChromium returns a boolean if a field has been set.

### GetAvailabilityLocationOutage

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityLocationOutage() bool`

GetAvailabilityLocationOutage returns the AvailabilityLocationOutage field if non-nil, zero value otherwise.

### GetAvailabilityLocationOutageOk

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityLocationOutageOk() (*bool, bool)`

GetAvailabilityLocationOutageOk returns a tuple with the AvailabilityLocationOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityLocationOutage

`func (o *SyntheticPrivateLocationUpdate) SetAvailabilityLocationOutage(v bool)`

SetAvailabilityLocationOutage sets AvailabilityLocationOutage field to given value.

### HasAvailabilityLocationOutage

`func (o *SyntheticPrivateLocationUpdate) HasAvailabilityLocationOutage() bool`

HasAvailabilityLocationOutage returns a boolean if a field has been set.

### GetAvailabilityNodeOutage

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNodeOutage() bool`

GetAvailabilityNodeOutage returns the AvailabilityNodeOutage field if non-nil, zero value otherwise.

### GetAvailabilityNodeOutageOk

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNodeOutageOk() (*bool, bool)`

GetAvailabilityNodeOutageOk returns a tuple with the AvailabilityNodeOutage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityNodeOutage

`func (o *SyntheticPrivateLocationUpdate) SetAvailabilityNodeOutage(v bool)`

SetAvailabilityNodeOutage sets AvailabilityNodeOutage field to given value.

### HasAvailabilityNodeOutage

`func (o *SyntheticPrivateLocationUpdate) HasAvailabilityNodeOutage() bool`

HasAvailabilityNodeOutage returns a boolean if a field has been set.

### GetAvailabilityNotificationsEnabled

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNotificationsEnabled() bool`

GetAvailabilityNotificationsEnabled returns the AvailabilityNotificationsEnabled field if non-nil, zero value otherwise.

### GetAvailabilityNotificationsEnabledOk

`func (o *SyntheticPrivateLocationUpdate) GetAvailabilityNotificationsEnabledOk() (*bool, bool)`

GetAvailabilityNotificationsEnabledOk returns a tuple with the AvailabilityNotificationsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityNotificationsEnabled

`func (o *SyntheticPrivateLocationUpdate) SetAvailabilityNotificationsEnabled(v bool)`

SetAvailabilityNotificationsEnabled sets AvailabilityNotificationsEnabled field to given value.

### HasAvailabilityNotificationsEnabled

`func (o *SyntheticPrivateLocationUpdate) HasAvailabilityNotificationsEnabled() bool`

HasAvailabilityNotificationsEnabled returns a boolean if a field has been set.

### GetCity

`func (o *SyntheticPrivateLocationUpdate) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SyntheticPrivateLocationUpdate) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SyntheticPrivateLocationUpdate) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SyntheticPrivateLocationUpdate) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountryCode

`func (o *SyntheticPrivateLocationUpdate) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *SyntheticPrivateLocationUpdate) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *SyntheticPrivateLocationUpdate) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *SyntheticPrivateLocationUpdate) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetDeploymentType

`func (o *SyntheticPrivateLocationUpdate) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *SyntheticPrivateLocationUpdate) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *SyntheticPrivateLocationUpdate) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.

### HasDeploymentType

`func (o *SyntheticPrivateLocationUpdate) HasDeploymentType() bool`

HasDeploymentType returns a boolean if a field has been set.

### GetLatitude

`func (o *SyntheticPrivateLocationUpdate) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SyntheticPrivateLocationUpdate) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SyntheticPrivateLocationUpdate) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.


### GetLocationNodeOutageDelayInMinutes

`func (o *SyntheticPrivateLocationUpdate) GetLocationNodeOutageDelayInMinutes() int32`

GetLocationNodeOutageDelayInMinutes returns the LocationNodeOutageDelayInMinutes field if non-nil, zero value otherwise.

### GetLocationNodeOutageDelayInMinutesOk

`func (o *SyntheticPrivateLocationUpdate) GetLocationNodeOutageDelayInMinutesOk() (*int32, bool)`

GetLocationNodeOutageDelayInMinutesOk returns a tuple with the LocationNodeOutageDelayInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationNodeOutageDelayInMinutes

`func (o *SyntheticPrivateLocationUpdate) SetLocationNodeOutageDelayInMinutes(v int32)`

SetLocationNodeOutageDelayInMinutes sets LocationNodeOutageDelayInMinutes field to given value.

### HasLocationNodeOutageDelayInMinutes

`func (o *SyntheticPrivateLocationUpdate) HasLocationNodeOutageDelayInMinutes() bool`

HasLocationNodeOutageDelayInMinutes returns a boolean if a field has been set.

### GetLongitude

`func (o *SyntheticPrivateLocationUpdate) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SyntheticPrivateLocationUpdate) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SyntheticPrivateLocationUpdate) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.


### GetName

`func (o *SyntheticPrivateLocationUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SyntheticPrivateLocationUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SyntheticPrivateLocationUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetNodes

`func (o *SyntheticPrivateLocationUpdate) GetNodes() []string`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *SyntheticPrivateLocationUpdate) GetNodesOk() (*[]string, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *SyntheticPrivateLocationUpdate) SetNodes(v []string)`

SetNodes sets Nodes field to given value.


### GetRegionCode

`func (o *SyntheticPrivateLocationUpdate) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *SyntheticPrivateLocationUpdate) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *SyntheticPrivateLocationUpdate) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *SyntheticPrivateLocationUpdate) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetStatus

`func (o *SyntheticPrivateLocationUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SyntheticPrivateLocationUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SyntheticPrivateLocationUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SyntheticPrivateLocationUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUseNewKubernetesVersion

`func (o *SyntheticPrivateLocationUpdate) GetUseNewKubernetesVersion() bool`

GetUseNewKubernetesVersion returns the UseNewKubernetesVersion field if non-nil, zero value otherwise.

### GetUseNewKubernetesVersionOk

`func (o *SyntheticPrivateLocationUpdate) GetUseNewKubernetesVersionOk() (*bool, bool)`

GetUseNewKubernetesVersionOk returns a tuple with the UseNewKubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNewKubernetesVersion

`func (o *SyntheticPrivateLocationUpdate) SetUseNewKubernetesVersion(v bool)`

SetUseNewKubernetesVersion sets UseNewKubernetesVersion field to given value.

### HasUseNewKubernetesVersion

`func (o *SyntheticPrivateLocationUpdate) HasUseNewKubernetesVersion() bool`

HasUseNewKubernetesVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


