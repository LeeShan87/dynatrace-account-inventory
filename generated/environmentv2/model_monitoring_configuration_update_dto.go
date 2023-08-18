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

// checks if the MonitoringConfigurationUpdateDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringConfigurationUpdateDto{}

// MonitoringConfigurationUpdateDto struct for MonitoringConfigurationUpdateDto
type MonitoringConfigurationUpdateDto struct {
	// The monitoring configuration
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewMonitoringConfigurationUpdateDto instantiates a new MonitoringConfigurationUpdateDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringConfigurationUpdateDto() *MonitoringConfigurationUpdateDto {
	this := MonitoringConfigurationUpdateDto{}
	return &this
}

// NewMonitoringConfigurationUpdateDtoWithDefaults instantiates a new MonitoringConfigurationUpdateDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringConfigurationUpdateDtoWithDefaults() *MonitoringConfigurationUpdateDto {
	this := MonitoringConfigurationUpdateDto{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MonitoringConfigurationUpdateDto) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfigurationUpdateDto) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MonitoringConfigurationUpdateDto) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *MonitoringConfigurationUpdateDto) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o MonitoringConfigurationUpdateDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringConfigurationUpdateDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableMonitoringConfigurationUpdateDto struct {
	value *MonitoringConfigurationUpdateDto
	isSet bool
}

func (v NullableMonitoringConfigurationUpdateDto) Get() *MonitoringConfigurationUpdateDto {
	return v.value
}

func (v *NullableMonitoringConfigurationUpdateDto) Set(val *MonitoringConfigurationUpdateDto) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringConfigurationUpdateDto) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringConfigurationUpdateDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringConfigurationUpdateDto(val *MonitoringConfigurationUpdateDto) *NullableMonitoringConfigurationUpdateDto {
	return &NullableMonitoringConfigurationUpdateDto{value: val, isSet: true}
}

func (v NullableMonitoringConfigurationUpdateDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringConfigurationUpdateDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


