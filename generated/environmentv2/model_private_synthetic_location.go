/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
)

// checks if the PrivateSyntheticLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrivateSyntheticLocation{}

// PrivateSyntheticLocation Configuration of a private synthetic location.    Some fields are inherited from the base *SyntheticLocation* object.
type PrivateSyntheticLocation struct {
	// Auto upgrade of Chromium is enabled (`true`) or disabled (`false`).
	AutoUpdateChromium *bool `json:"autoUpdateChromium,omitempty"`
	// Alerting for location outage is enabled (`true`) or disabled (`false`). Supported only for private Synthetic locations.
	AvailabilityLocationOutage *bool `json:"availabilityLocationOutage,omitempty"`
	// Alerting for node outage is enabled (`true`) or disabled (`false`). \\n\\n If enabled, the outage of *any* node in the location triggers an alert. Supported only for private Synthetic locations.
	AvailabilityNodeOutage *bool `json:"availabilityNodeOutage,omitempty"`
	// Notifications for location and node outage are enabled (`true`) or disabled (`false`). Supported only for private Synthetic locations.
	AvailabilityNotificationsEnabled *bool `json:"availabilityNotificationsEnabled,omitempty"`
	// The deployment type of the location:   * `STANDARD`: The location is deployed on Windows or Linux. * `KUBERNETES`: The location is deployed on Kubernetes.
	DeploymentType *string `json:"deploymentType,omitempty"`
	// Alert if location or node outage lasts longer than *X* minutes. \\n\\n Only applicable when `availabilityLocationOutage` or `availabilityNodeOutage` is set to `true`. Supported only for private Synthetic locations.
	LocationNodeOutageDelayInMinutes *int32 `json:"locationNodeOutageDelayInMinutes,omitempty"`
	// A list of synthetic nodes belonging to the location.    You can retrieve the list of available nodes with the [GET all nodes](https://dt-url.net/miy3rpl) call.
	Nodes []string `json:"nodes"`
	// Boolean value describes which kubernetes version will be used:  * `false`: Version 1.23+ that is older than 1.26 * `true`: Version 1.26+.
	UseNewKubernetesVersion *bool `json:"useNewKubernetesVersion,omitempty"`
}

// NewPrivateSyntheticLocation instantiates a new PrivateSyntheticLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrivateSyntheticLocation(nodes []string, entityId string, latitude float64, longitude float64, name string, type_ string) *PrivateSyntheticLocation {
	this := PrivateSyntheticLocation{}
	this.EntityId = entityId
	this.Latitude = latitude
	this.Longitude = longitude
	this.Name = name
	this.Type = type_
	this.Nodes = nodes
	return &this
}

// NewPrivateSyntheticLocationWithDefaults instantiates a new PrivateSyntheticLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrivateSyntheticLocationWithDefaults() *PrivateSyntheticLocation {
	this := PrivateSyntheticLocation{}
	return &this
}

// GetAutoUpdateChromium returns the AutoUpdateChromium field value if set, zero value otherwise.
func (o *PrivateSyntheticLocation) GetAutoUpdateChromium() bool {
	if o == nil || IsNil(o.AutoUpdateChromium) {
		var ret bool
		return ret
	}
	return *o.AutoUpdateChromium
}

// GetAutoUpdateChromiumOk returns a tuple with the AutoUpdateChromium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocation) GetAutoUpdateChromiumOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoUpdateChromium) {
		return nil, false
	}
	return o.AutoUpdateChromium, true
}

// HasAutoUpdateChromium returns a boolean if a field has been set.
func (o *PrivateSyntheticLocation) HasAutoUpdateChromium() bool {
	if o != nil && !IsNil(o.AutoUpdateChromium) {
		return true
	}

	return false
}

// SetAutoUpdateChromium gets a reference to the given bool and assigns it to the AutoUpdateChromium field.
func (o *PrivateSyntheticLocation) SetAutoUpdateChromium(v bool) {
	o.AutoUpdateChromium = &v
}

// GetAvailabilityLocationOutage returns the AvailabilityLocationOutage field value if set, zero value otherwise.
func (o *PrivateSyntheticLocation) GetAvailabilityLocationOutage() bool {
	if o == nil || IsNil(o.AvailabilityLocationOutage) {
		var ret bool
		return ret
	}
	return *o.AvailabilityLocationOutage
}

// GetAvailabilityLocationOutageOk returns a tuple with the AvailabilityLocationOutage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocation) GetAvailabilityLocationOutageOk() (*bool, bool) {
	if o == nil || IsNil(o.AvailabilityLocationOutage) {
		return nil, false
	}
	return o.AvailabilityLocationOutage, true
}

// HasAvailabilityLocationOutage returns a boolean if a field has been set.
func (o *PrivateSyntheticLocation) HasAvailabilityLocationOutage() bool {
	if o != nil && !IsNil(o.AvailabilityLocationOutage) {
		return true
	}

	return false
}

// SetAvailabilityLocationOutage gets a reference to the given bool and assigns it to the AvailabilityLocationOutage field.
func (o *PrivateSyntheticLocation) SetAvailabilityLocationOutage(v bool) {
	o.AvailabilityLocationOutage = &v
}

// GetAvailabilityNodeOutage returns the AvailabilityNodeOutage field value if set, zero value otherwise.
func (o *PrivateSyntheticLocation) GetAvailabilityNodeOutage() bool {
	if o == nil || IsNil(o.AvailabilityNodeOutage) {
		var ret bool
		return ret
	}
	return *o.AvailabilityNodeOutage
}

// GetAvailabilityNodeOutageOk returns a tuple with the AvailabilityNodeOutage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocation) GetAvailabilityNodeOutageOk() (*bool, bool) {
	if o == nil || IsNil(o.AvailabilityNodeOutage) {
		return nil, false
	}
	return o.AvailabilityNodeOutage, true
}

// HasAvailabilityNodeOutage returns a boolean if a field has been set.
func (o *PrivateSyntheticLocation) HasAvailabilityNodeOutage() bool {
	if o != nil && !IsNil(o.AvailabilityNodeOutage) {
		return true
	}

	return false
}

// SetAvailabilityNodeOutage gets a reference to the given bool and assigns it to the AvailabilityNodeOutage field.
func (o *PrivateSyntheticLocation) SetAvailabilityNodeOutage(v bool) {
	o.AvailabilityNodeOutage = &v
}

// GetAvailabilityNotificationsEnabled returns the AvailabilityNotificationsEnabled field value if set, zero value otherwise.
func (o *PrivateSyntheticLocation) GetAvailabilityNotificationsEnabled() bool {
	if o == nil || IsNil(o.AvailabilityNotificationsEnabled) {
		var ret bool
		return ret
	}
	return *o.AvailabilityNotificationsEnabled
}

// GetAvailabilityNotificationsEnabledOk returns a tuple with the AvailabilityNotificationsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocation) GetAvailabilityNotificationsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AvailabilityNotificationsEnabled) {
		return nil, false
	}
	return o.AvailabilityNotificationsEnabled, true
}

// HasAvailabilityNotificationsEnabled returns a boolean if a field has been set.
func (o *PrivateSyntheticLocation) HasAvailabilityNotificationsEnabled() bool {
	if o != nil && !IsNil(o.AvailabilityNotificationsEnabled) {
		return true
	}

	return false
}

// SetAvailabilityNotificationsEnabled gets a reference to the given bool and assigns it to the AvailabilityNotificationsEnabled field.
func (o *PrivateSyntheticLocation) SetAvailabilityNotificationsEnabled(v bool) {
	o.AvailabilityNotificationsEnabled = &v
}

// GetDeploymentType returns the DeploymentType field value if set, zero value otherwise.
func (o *PrivateSyntheticLocation) GetDeploymentType() string {
	if o == nil || IsNil(o.DeploymentType) {
		var ret string
		return ret
	}
	return *o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocation) GetDeploymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DeploymentType) {
		return nil, false
	}
	return o.DeploymentType, true
}

// HasDeploymentType returns a boolean if a field has been set.
func (o *PrivateSyntheticLocation) HasDeploymentType() bool {
	if o != nil && !IsNil(o.DeploymentType) {
		return true
	}

	return false
}

// SetDeploymentType gets a reference to the given string and assigns it to the DeploymentType field.
func (o *PrivateSyntheticLocation) SetDeploymentType(v string) {
	o.DeploymentType = &v
}

// GetLocationNodeOutageDelayInMinutes returns the LocationNodeOutageDelayInMinutes field value if set, zero value otherwise.
func (o *PrivateSyntheticLocation) GetLocationNodeOutageDelayInMinutes() int32 {
	if o == nil || IsNil(o.LocationNodeOutageDelayInMinutes) {
		var ret int32
		return ret
	}
	return *o.LocationNodeOutageDelayInMinutes
}

// GetLocationNodeOutageDelayInMinutesOk returns a tuple with the LocationNodeOutageDelayInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocation) GetLocationNodeOutageDelayInMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.LocationNodeOutageDelayInMinutes) {
		return nil, false
	}
	return o.LocationNodeOutageDelayInMinutes, true
}

// HasLocationNodeOutageDelayInMinutes returns a boolean if a field has been set.
func (o *PrivateSyntheticLocation) HasLocationNodeOutageDelayInMinutes() bool {
	if o != nil && !IsNil(o.LocationNodeOutageDelayInMinutes) {
		return true
	}

	return false
}

// SetLocationNodeOutageDelayInMinutes gets a reference to the given int32 and assigns it to the LocationNodeOutageDelayInMinutes field.
func (o *PrivateSyntheticLocation) SetLocationNodeOutageDelayInMinutes(v int32) {
	o.LocationNodeOutageDelayInMinutes = &v
}

// GetNodes returns the Nodes field value
func (o *PrivateSyntheticLocation) GetNodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocation) GetNodesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *PrivateSyntheticLocation) SetNodes(v []string) {
	o.Nodes = v
}

// GetUseNewKubernetesVersion returns the UseNewKubernetesVersion field value if set, zero value otherwise.
func (o *PrivateSyntheticLocation) GetUseNewKubernetesVersion() bool {
	if o == nil || IsNil(o.UseNewKubernetesVersion) {
		var ret bool
		return ret
	}
	return *o.UseNewKubernetesVersion
}

// GetUseNewKubernetesVersionOk returns a tuple with the UseNewKubernetesVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrivateSyntheticLocation) GetUseNewKubernetesVersionOk() (*bool, bool) {
	if o == nil || IsNil(o.UseNewKubernetesVersion) {
		return nil, false
	}
	return o.UseNewKubernetesVersion, true
}

// HasUseNewKubernetesVersion returns a boolean if a field has been set.
func (o *PrivateSyntheticLocation) HasUseNewKubernetesVersion() bool {
	if o != nil && !IsNil(o.UseNewKubernetesVersion) {
		return true
	}

	return false
}

// SetUseNewKubernetesVersion gets a reference to the given bool and assigns it to the UseNewKubernetesVersion field.
func (o *PrivateSyntheticLocation) SetUseNewKubernetesVersion(v bool) {
	o.UseNewKubernetesVersion = &v
}

func (o PrivateSyntheticLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrivateSyntheticLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoUpdateChromium) {
		toSerialize["autoUpdateChromium"] = o.AutoUpdateChromium
	}
	if !IsNil(o.AvailabilityLocationOutage) {
		toSerialize["availabilityLocationOutage"] = o.AvailabilityLocationOutage
	}
	if !IsNil(o.AvailabilityNodeOutage) {
		toSerialize["availabilityNodeOutage"] = o.AvailabilityNodeOutage
	}
	if !IsNil(o.AvailabilityNotificationsEnabled) {
		toSerialize["availabilityNotificationsEnabled"] = o.AvailabilityNotificationsEnabled
	}
	if !IsNil(o.DeploymentType) {
		toSerialize["deploymentType"] = o.DeploymentType
	}
	if !IsNil(o.LocationNodeOutageDelayInMinutes) {
		toSerialize["locationNodeOutageDelayInMinutes"] = o.LocationNodeOutageDelayInMinutes
	}
	toSerialize["nodes"] = o.Nodes
	if !IsNil(o.UseNewKubernetesVersion) {
		toSerialize["useNewKubernetesVersion"] = o.UseNewKubernetesVersion
	}
	return toSerialize, nil
}

type NullablePrivateSyntheticLocation struct {
	value *PrivateSyntheticLocation
	isSet bool
}

func (v NullablePrivateSyntheticLocation) Get() *PrivateSyntheticLocation {
	return v.value
}

func (v *NullablePrivateSyntheticLocation) Set(val *PrivateSyntheticLocation) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivateSyntheticLocation) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivateSyntheticLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivateSyntheticLocation(val *PrivateSyntheticLocation) *NullablePrivateSyntheticLocation {
	return &NullablePrivateSyntheticLocation{value: val, isSet: true}
}

func (v NullablePrivateSyntheticLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivateSyntheticLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


