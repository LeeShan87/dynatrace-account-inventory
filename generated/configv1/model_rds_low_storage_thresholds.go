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

// checks if the RdsLowStorageThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RdsLowStorageThresholds{}

// RdsLowStorageThresholds Custom thresholds for low free storage space on RDS. If not set, automatic mode is used.
type RdsLowStorageThresholds struct {
	// Alert if free storage space divided by allocated storage is lower than *X*% in 3 out of 5 samples.
	FreeStoragePercentage int32 `json:"freeStoragePercentage"`
}

// NewRdsLowStorageThresholds instantiates a new RdsLowStorageThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRdsLowStorageThresholds(freeStoragePercentage int32) *RdsLowStorageThresholds {
	this := RdsLowStorageThresholds{}
	this.FreeStoragePercentage = freeStoragePercentage
	return &this
}

// NewRdsLowStorageThresholdsWithDefaults instantiates a new RdsLowStorageThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRdsLowStorageThresholdsWithDefaults() *RdsLowStorageThresholds {
	this := RdsLowStorageThresholds{}
	return &this
}

// GetFreeStoragePercentage returns the FreeStoragePercentage field value
func (o *RdsLowStorageThresholds) GetFreeStoragePercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FreeStoragePercentage
}

// GetFreeStoragePercentageOk returns a tuple with the FreeStoragePercentage field value
// and a boolean to check if the value has been set.
func (o *RdsLowStorageThresholds) GetFreeStoragePercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeStoragePercentage, true
}

// SetFreeStoragePercentage sets field value
func (o *RdsLowStorageThresholds) SetFreeStoragePercentage(v int32) {
	o.FreeStoragePercentage = v
}

func (o RdsLowStorageThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RdsLowStorageThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["freeStoragePercentage"] = o.FreeStoragePercentage
	return toSerialize, nil
}

type NullableRdsLowStorageThresholds struct {
	value *RdsLowStorageThresholds
	isSet bool
}

func (v NullableRdsLowStorageThresholds) Get() *RdsLowStorageThresholds {
	return v.value
}

func (v *NullableRdsLowStorageThresholds) Set(val *RdsLowStorageThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableRdsLowStorageThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableRdsLowStorageThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRdsLowStorageThresholds(val *RdsLowStorageThresholds) *NullableRdsLowStorageThresholds {
	return &NullableRdsLowStorageThresholds{value: val, isSet: true}
}

func (v NullableRdsLowStorageThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRdsLowStorageThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


