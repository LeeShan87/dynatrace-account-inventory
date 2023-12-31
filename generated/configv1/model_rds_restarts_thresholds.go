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

// checks if the RdsRestartsThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RdsRestartsThresholds{}

// RdsRestartsThresholds Custom thresholds for restarts sequence on RDS. If not set, automatic mode is used.
type RdsRestartsThresholds struct {
	// Alert if number of restarts is *X* per minute or higher in 3 out of 20 samples.
	RestartsPerMinute int32 `json:"restartsPerMinute"`
}

// NewRdsRestartsThresholds instantiates a new RdsRestartsThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRdsRestartsThresholds(restartsPerMinute int32) *RdsRestartsThresholds {
	this := RdsRestartsThresholds{}
	this.RestartsPerMinute = restartsPerMinute
	return &this
}

// NewRdsRestartsThresholdsWithDefaults instantiates a new RdsRestartsThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRdsRestartsThresholdsWithDefaults() *RdsRestartsThresholds {
	this := RdsRestartsThresholds{}
	return &this
}

// GetRestartsPerMinute returns the RestartsPerMinute field value
func (o *RdsRestartsThresholds) GetRestartsPerMinute() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RestartsPerMinute
}

// GetRestartsPerMinuteOk returns a tuple with the RestartsPerMinute field value
// and a boolean to check if the value has been set.
func (o *RdsRestartsThresholds) GetRestartsPerMinuteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RestartsPerMinute, true
}

// SetRestartsPerMinute sets field value
func (o *RdsRestartsThresholds) SetRestartsPerMinute(v int32) {
	o.RestartsPerMinute = v
}

func (o RdsRestartsThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RdsRestartsThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["restartsPerMinute"] = o.RestartsPerMinute
	return toSerialize, nil
}

type NullableRdsRestartsThresholds struct {
	value *RdsRestartsThresholds
	isSet bool
}

func (v NullableRdsRestartsThresholds) Get() *RdsRestartsThresholds {
	return v.value
}

func (v *NullableRdsRestartsThresholds) Set(val *RdsRestartsThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableRdsRestartsThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableRdsRestartsThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRdsRestartsThresholds(val *RdsRestartsThresholds) *NullableRdsRestartsThresholds {
	return &NullableRdsRestartsThresholds{value: val, isSet: true}
}

func (v NullableRdsRestartsThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRdsRestartsThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


