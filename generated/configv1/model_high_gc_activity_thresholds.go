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

// checks if the HighGcActivityThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HighGcActivityThresholds{}

// HighGcActivityThresholds Custom thresholds for high GC activity. If not set, automatic mode is used.    Meeting **any** of these conditions triggers an alert.
type HighGcActivityThresholds struct {
	// GC suspension is higher than *X*% in 3 out of 5 samples.
	GcSuspensionPercentage int32 `json:"gcSuspensionPercentage"`
	// GC time is higher than *X*% in 3 out of 5 samples.
	GcTimePercentage int32 `json:"gcTimePercentage"`
}

// NewHighGcActivityThresholds instantiates a new HighGcActivityThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHighGcActivityThresholds(gcSuspensionPercentage int32, gcTimePercentage int32) *HighGcActivityThresholds {
	this := HighGcActivityThresholds{}
	this.GcSuspensionPercentage = gcSuspensionPercentage
	this.GcTimePercentage = gcTimePercentage
	return &this
}

// NewHighGcActivityThresholdsWithDefaults instantiates a new HighGcActivityThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHighGcActivityThresholdsWithDefaults() *HighGcActivityThresholds {
	this := HighGcActivityThresholds{}
	return &this
}

// GetGcSuspensionPercentage returns the GcSuspensionPercentage field value
func (o *HighGcActivityThresholds) GetGcSuspensionPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GcSuspensionPercentage
}

// GetGcSuspensionPercentageOk returns a tuple with the GcSuspensionPercentage field value
// and a boolean to check if the value has been set.
func (o *HighGcActivityThresholds) GetGcSuspensionPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcSuspensionPercentage, true
}

// SetGcSuspensionPercentage sets field value
func (o *HighGcActivityThresholds) SetGcSuspensionPercentage(v int32) {
	o.GcSuspensionPercentage = v
}

// GetGcTimePercentage returns the GcTimePercentage field value
func (o *HighGcActivityThresholds) GetGcTimePercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.GcTimePercentage
}

// GetGcTimePercentageOk returns a tuple with the GcTimePercentage field value
// and a boolean to check if the value has been set.
func (o *HighGcActivityThresholds) GetGcTimePercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GcTimePercentage, true
}

// SetGcTimePercentage sets field value
func (o *HighGcActivityThresholds) SetGcTimePercentage(v int32) {
	o.GcTimePercentage = v
}

func (o HighGcActivityThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HighGcActivityThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gcSuspensionPercentage"] = o.GcSuspensionPercentage
	toSerialize["gcTimePercentage"] = o.GcTimePercentage
	return toSerialize, nil
}

type NullableHighGcActivityThresholds struct {
	value *HighGcActivityThresholds
	isSet bool
}

func (v NullableHighGcActivityThresholds) Get() *HighGcActivityThresholds {
	return v.value
}

func (v *NullableHighGcActivityThresholds) Set(val *HighGcActivityThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableHighGcActivityThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableHighGcActivityThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighGcActivityThresholds(val *HighGcActivityThresholds) *NullableHighGcActivityThresholds {
	return &NullableHighGcActivityThresholds{value: val, isSet: true}
}

func (v NullableHighGcActivityThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHighGcActivityThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


