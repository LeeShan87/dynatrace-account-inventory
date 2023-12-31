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

// checks if the HighCpuSaturationThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HighCpuSaturationThresholds{}

// HighCpuSaturationThresholds Custom thresholds for high CPU saturation. If not set then the automatic mode is used.
type HighCpuSaturationThresholds struct {
	// Alert if CPU usage is higher than *X*% in 3 out of 5 samples.
	CpuSaturation int32 `json:"cpuSaturation"`
}

// NewHighCpuSaturationThresholds instantiates a new HighCpuSaturationThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHighCpuSaturationThresholds(cpuSaturation int32) *HighCpuSaturationThresholds {
	this := HighCpuSaturationThresholds{}
	this.CpuSaturation = cpuSaturation
	return &this
}

// NewHighCpuSaturationThresholdsWithDefaults instantiates a new HighCpuSaturationThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHighCpuSaturationThresholdsWithDefaults() *HighCpuSaturationThresholds {
	this := HighCpuSaturationThresholds{}
	return &this
}

// GetCpuSaturation returns the CpuSaturation field value
func (o *HighCpuSaturationThresholds) GetCpuSaturation() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CpuSaturation
}

// GetCpuSaturationOk returns a tuple with the CpuSaturation field value
// and a boolean to check if the value has been set.
func (o *HighCpuSaturationThresholds) GetCpuSaturationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuSaturation, true
}

// SetCpuSaturation sets field value
func (o *HighCpuSaturationThresholds) SetCpuSaturation(v int32) {
	o.CpuSaturation = v
}

func (o HighCpuSaturationThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HighCpuSaturationThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpuSaturation"] = o.CpuSaturation
	return toSerialize, nil
}

type NullableHighCpuSaturationThresholds struct {
	value *HighCpuSaturationThresholds
	isSet bool
}

func (v NullableHighCpuSaturationThresholds) Get() *HighCpuSaturationThresholds {
	return v.value
}

func (v *NullableHighCpuSaturationThresholds) Set(val *HighCpuSaturationThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableHighCpuSaturationThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableHighCpuSaturationThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighCpuSaturationThresholds(val *HighCpuSaturationThresholds) *NullableHighCpuSaturationThresholds {
	return &NullableHighCpuSaturationThresholds{value: val, isSet: true}
}

func (v NullableHighCpuSaturationThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHighCpuSaturationThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


