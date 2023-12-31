/*
Dynatrace Configuration API

Documentation of the Dynatrace Configuration API. To read about use-cases and examples, see [Dynatrace Documentation](https://dt-url.net/4u43kxw).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configv1

import (
	"encoding/json"
)

// checks if the HostConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostConfig{}

// HostConfig OneAgent configuration on a host.
type HostConfig struct {
	AutoUpdateConfig *HostAutoUpdateConfig `json:"autoUpdateConfig,omitempty"`
	// The Dynatrace entity ID of the host where OneAgent is deployed.
	Id *string `json:"id,omitempty"`
	MonitoringConfig *MonitoringConfig `json:"monitoringConfig,omitempty"`
	TechMonitoringConfigList *TechMonitoringList `json:"techMonitoringConfigList,omitempty"`
}

// NewHostConfig instantiates a new HostConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostConfig() *HostConfig {
	this := HostConfig{}
	return &this
}

// NewHostConfigWithDefaults instantiates a new HostConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostConfigWithDefaults() *HostConfig {
	this := HostConfig{}
	return &this
}

// GetAutoUpdateConfig returns the AutoUpdateConfig field value if set, zero value otherwise.
func (o *HostConfig) GetAutoUpdateConfig() HostAutoUpdateConfig {
	if o == nil || IsNil(o.AutoUpdateConfig) {
		var ret HostAutoUpdateConfig
		return ret
	}
	return *o.AutoUpdateConfig
}

// GetAutoUpdateConfigOk returns a tuple with the AutoUpdateConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostConfig) GetAutoUpdateConfigOk() (*HostAutoUpdateConfig, bool) {
	if o == nil || IsNil(o.AutoUpdateConfig) {
		return nil, false
	}
	return o.AutoUpdateConfig, true
}

// HasAutoUpdateConfig returns a boolean if a field has been set.
func (o *HostConfig) HasAutoUpdateConfig() bool {
	if o != nil && !IsNil(o.AutoUpdateConfig) {
		return true
	}

	return false
}

// SetAutoUpdateConfig gets a reference to the given HostAutoUpdateConfig and assigns it to the AutoUpdateConfig field.
func (o *HostConfig) SetAutoUpdateConfig(v HostAutoUpdateConfig) {
	o.AutoUpdateConfig = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HostConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HostConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *HostConfig) SetId(v string) {
	o.Id = &v
}

// GetMonitoringConfig returns the MonitoringConfig field value if set, zero value otherwise.
func (o *HostConfig) GetMonitoringConfig() MonitoringConfig {
	if o == nil || IsNil(o.MonitoringConfig) {
		var ret MonitoringConfig
		return ret
	}
	return *o.MonitoringConfig
}

// GetMonitoringConfigOk returns a tuple with the MonitoringConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostConfig) GetMonitoringConfigOk() (*MonitoringConfig, bool) {
	if o == nil || IsNil(o.MonitoringConfig) {
		return nil, false
	}
	return o.MonitoringConfig, true
}

// HasMonitoringConfig returns a boolean if a field has been set.
func (o *HostConfig) HasMonitoringConfig() bool {
	if o != nil && !IsNil(o.MonitoringConfig) {
		return true
	}

	return false
}

// SetMonitoringConfig gets a reference to the given MonitoringConfig and assigns it to the MonitoringConfig field.
func (o *HostConfig) SetMonitoringConfig(v MonitoringConfig) {
	o.MonitoringConfig = &v
}

// GetTechMonitoringConfigList returns the TechMonitoringConfigList field value if set, zero value otherwise.
func (o *HostConfig) GetTechMonitoringConfigList() TechMonitoringList {
	if o == nil || IsNil(o.TechMonitoringConfigList) {
		var ret TechMonitoringList
		return ret
	}
	return *o.TechMonitoringConfigList
}

// GetTechMonitoringConfigListOk returns a tuple with the TechMonitoringConfigList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostConfig) GetTechMonitoringConfigListOk() (*TechMonitoringList, bool) {
	if o == nil || IsNil(o.TechMonitoringConfigList) {
		return nil, false
	}
	return o.TechMonitoringConfigList, true
}

// HasTechMonitoringConfigList returns a boolean if a field has been set.
func (o *HostConfig) HasTechMonitoringConfigList() bool {
	if o != nil && !IsNil(o.TechMonitoringConfigList) {
		return true
	}

	return false
}

// SetTechMonitoringConfigList gets a reference to the given TechMonitoringList and assigns it to the TechMonitoringConfigList field.
func (o *HostConfig) SetTechMonitoringConfigList(v TechMonitoringList) {
	o.TechMonitoringConfigList = &v
}

func (o HostConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoUpdateConfig) {
		toSerialize["autoUpdateConfig"] = o.AutoUpdateConfig
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MonitoringConfig) {
		toSerialize["monitoringConfig"] = o.MonitoringConfig
	}
	if !IsNil(o.TechMonitoringConfigList) {
		toSerialize["techMonitoringConfigList"] = o.TechMonitoringConfigList
	}
	return toSerialize, nil
}

type NullableHostConfig struct {
	value *HostConfig
	isSet bool
}

func (v NullableHostConfig) Get() *HostConfig {
	return v.value
}

func (v *NullableHostConfig) Set(val *HostConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHostConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHostConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostConfig(val *HostConfig) *NullableHostConfig {
	return &NullableHostConfig{value: val, isSet: true}
}

func (v NullableHostConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


