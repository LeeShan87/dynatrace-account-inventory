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

// checks if the DiskLowSpaceThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiskLowSpaceThresholds{}

// DiskLowSpaceThresholds Custom thresholds for low disk space. If not set, automatic mode is used.
type DiskLowSpaceThresholds struct {
	// Alert if free disk space is lower than *X*% in 3 out of 5 samples.
	FreeSpacePercentage int32 `json:"freeSpacePercentage"`
}

// NewDiskLowSpaceThresholds instantiates a new DiskLowSpaceThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiskLowSpaceThresholds(freeSpacePercentage int32) *DiskLowSpaceThresholds {
	this := DiskLowSpaceThresholds{}
	this.FreeSpacePercentage = freeSpacePercentage
	return &this
}

// NewDiskLowSpaceThresholdsWithDefaults instantiates a new DiskLowSpaceThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiskLowSpaceThresholdsWithDefaults() *DiskLowSpaceThresholds {
	this := DiskLowSpaceThresholds{}
	return &this
}

// GetFreeSpacePercentage returns the FreeSpacePercentage field value
func (o *DiskLowSpaceThresholds) GetFreeSpacePercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FreeSpacePercentage
}

// GetFreeSpacePercentageOk returns a tuple with the FreeSpacePercentage field value
// and a boolean to check if the value has been set.
func (o *DiskLowSpaceThresholds) GetFreeSpacePercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FreeSpacePercentage, true
}

// SetFreeSpacePercentage sets field value
func (o *DiskLowSpaceThresholds) SetFreeSpacePercentage(v int32) {
	o.FreeSpacePercentage = v
}

func (o DiskLowSpaceThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiskLowSpaceThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["freeSpacePercentage"] = o.FreeSpacePercentage
	return toSerialize, nil
}

type NullableDiskLowSpaceThresholds struct {
	value *DiskLowSpaceThresholds
	isSet bool
}

func (v NullableDiskLowSpaceThresholds) Get() *DiskLowSpaceThresholds {
	return v.value
}

func (v *NullableDiskLowSpaceThresholds) Set(val *DiskLowSpaceThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableDiskLowSpaceThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableDiskLowSpaceThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiskLowSpaceThresholds(val *DiskLowSpaceThresholds) *NullableDiskLowSpaceThresholds {
	return &NullableDiskLowSpaceThresholds{value: val, isSet: true}
}

func (v NullableDiskLowSpaceThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiskLowSpaceThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


