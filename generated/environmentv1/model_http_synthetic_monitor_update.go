/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv1

import (
	"encoding/json"
)

// checks if the HttpSyntheticMonitorUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpSyntheticMonitorUpdate{}

// HttpSyntheticMonitorUpdate HTTP synthetic monitor update. Some fields are inherited from base `SyntheticMonitorUpdate` model.
type HttpSyntheticMonitorUpdate struct {
}

// NewHttpSyntheticMonitorUpdate instantiates a new HttpSyntheticMonitorUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpSyntheticMonitorUpdate(enabled bool, frequencyMin int32, locations []string, manuallyAssignedApps []string, name string, script map[string]interface{}, tags []TagWithSourceInfo, type_ string) *HttpSyntheticMonitorUpdate {
	this := HttpSyntheticMonitorUpdate{}
	this.Enabled = enabled
	this.FrequencyMin = frequencyMin
	this.Locations = locations
	this.ManuallyAssignedApps = manuallyAssignedApps
	this.Name = name
	this.Script = script
	this.Tags = tags
	this.Type = type_
	return &this
}

// NewHttpSyntheticMonitorUpdateWithDefaults instantiates a new HttpSyntheticMonitorUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpSyntheticMonitorUpdateWithDefaults() *HttpSyntheticMonitorUpdate {
	this := HttpSyntheticMonitorUpdate{}
	return &this
}

func (o HttpSyntheticMonitorUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpSyntheticMonitorUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableHttpSyntheticMonitorUpdate struct {
	value *HttpSyntheticMonitorUpdate
	isSet bool
}

func (v NullableHttpSyntheticMonitorUpdate) Get() *HttpSyntheticMonitorUpdate {
	return v.value
}

func (v *NullableHttpSyntheticMonitorUpdate) Set(val *HttpSyntheticMonitorUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpSyntheticMonitorUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpSyntheticMonitorUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpSyntheticMonitorUpdate(val *HttpSyntheticMonitorUpdate) *NullableHttpSyntheticMonitorUpdate {
	return &NullableHttpSyntheticMonitorUpdate{value: val, isSet: true}
}

func (v NullableHttpSyntheticMonitorUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpSyntheticMonitorUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


