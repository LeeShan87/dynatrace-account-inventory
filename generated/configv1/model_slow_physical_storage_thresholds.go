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

// checks if the SlowPhysicalStorageThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SlowPhysicalStorageThresholds{}

// SlowPhysicalStorageThresholds Custom thresholds for slow running physical storage device. If not set then the automatic mode is used.    Fulfillment of **any** condition triggers an alert.
type SlowPhysicalStorageThresholds struct {
	// Read/write latency is higher than *X* milliseconds in 4 out of 5 samples.
	AvgReadWriteLatency int32 `json:"avgReadWriteLatency"`
	// Peak value for read/write latency is higher than *X* milliseconds in 4 out of 5 samples.
	PeakReadWriteLatency int32 `json:"peakReadWriteLatency"`
}

// NewSlowPhysicalStorageThresholds instantiates a new SlowPhysicalStorageThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSlowPhysicalStorageThresholds(avgReadWriteLatency int32, peakReadWriteLatency int32) *SlowPhysicalStorageThresholds {
	this := SlowPhysicalStorageThresholds{}
	this.AvgReadWriteLatency = avgReadWriteLatency
	this.PeakReadWriteLatency = peakReadWriteLatency
	return &this
}

// NewSlowPhysicalStorageThresholdsWithDefaults instantiates a new SlowPhysicalStorageThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSlowPhysicalStorageThresholdsWithDefaults() *SlowPhysicalStorageThresholds {
	this := SlowPhysicalStorageThresholds{}
	return &this
}

// GetAvgReadWriteLatency returns the AvgReadWriteLatency field value
func (o *SlowPhysicalStorageThresholds) GetAvgReadWriteLatency() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AvgReadWriteLatency
}

// GetAvgReadWriteLatencyOk returns a tuple with the AvgReadWriteLatency field value
// and a boolean to check if the value has been set.
func (o *SlowPhysicalStorageThresholds) GetAvgReadWriteLatencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgReadWriteLatency, true
}

// SetAvgReadWriteLatency sets field value
func (o *SlowPhysicalStorageThresholds) SetAvgReadWriteLatency(v int32) {
	o.AvgReadWriteLatency = v
}

// GetPeakReadWriteLatency returns the PeakReadWriteLatency field value
func (o *SlowPhysicalStorageThresholds) GetPeakReadWriteLatency() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PeakReadWriteLatency
}

// GetPeakReadWriteLatencyOk returns a tuple with the PeakReadWriteLatency field value
// and a boolean to check if the value has been set.
func (o *SlowPhysicalStorageThresholds) GetPeakReadWriteLatencyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeakReadWriteLatency, true
}

// SetPeakReadWriteLatency sets field value
func (o *SlowPhysicalStorageThresholds) SetPeakReadWriteLatency(v int32) {
	o.PeakReadWriteLatency = v
}

func (o SlowPhysicalStorageThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SlowPhysicalStorageThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["avgReadWriteLatency"] = o.AvgReadWriteLatency
	toSerialize["peakReadWriteLatency"] = o.PeakReadWriteLatency
	return toSerialize, nil
}

type NullableSlowPhysicalStorageThresholds struct {
	value *SlowPhysicalStorageThresholds
	isSet bool
}

func (v NullableSlowPhysicalStorageThresholds) Get() *SlowPhysicalStorageThresholds {
	return v.value
}

func (v *NullableSlowPhysicalStorageThresholds) Set(val *SlowPhysicalStorageThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableSlowPhysicalStorageThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableSlowPhysicalStorageThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSlowPhysicalStorageThresholds(val *SlowPhysicalStorageThresholds) *NullableSlowPhysicalStorageThresholds {
	return &NullableSlowPhysicalStorageThresholds{value: val, isSet: true}
}

func (v NullableSlowPhysicalStorageThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSlowPhysicalStorageThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


