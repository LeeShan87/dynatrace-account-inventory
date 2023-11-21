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

// checks if the EsxiHighCpuThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EsxiHighCpuThresholds{}

// EsxiHighCpuThresholds Custom thresholds for CPU saturation detection on ESXi. If not set then the automatic mode is used.    **All** conditions must be fulfilled to trigger an alert.
type EsxiHighCpuThresholds struct {
	// At least one peak higher than *X*% occurred in 3 out of 5 samples.
	CpuPeakPercentage int32 `json:"cpuPeakPercentage"`
	// CPU usage is higher than *X*% in 3 out of 5 samples.
	CpuUsagePercentage int32 `json:"cpuUsagePercentage"`
	// VM CPU ready is higher than *X*% in 3 out of 5 samples.
	VmCpuReadyPercentage int32 `json:"vmCpuReadyPercentage"`
}

// NewEsxiHighCpuThresholds instantiates a new EsxiHighCpuThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEsxiHighCpuThresholds(cpuPeakPercentage int32, cpuUsagePercentage int32, vmCpuReadyPercentage int32) *EsxiHighCpuThresholds {
	this := EsxiHighCpuThresholds{}
	this.CpuPeakPercentage = cpuPeakPercentage
	this.CpuUsagePercentage = cpuUsagePercentage
	this.VmCpuReadyPercentage = vmCpuReadyPercentage
	return &this
}

// NewEsxiHighCpuThresholdsWithDefaults instantiates a new EsxiHighCpuThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEsxiHighCpuThresholdsWithDefaults() *EsxiHighCpuThresholds {
	this := EsxiHighCpuThresholds{}
	return &this
}

// GetCpuPeakPercentage returns the CpuPeakPercentage field value
func (o *EsxiHighCpuThresholds) GetCpuPeakPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CpuPeakPercentage
}

// GetCpuPeakPercentageOk returns a tuple with the CpuPeakPercentage field value
// and a boolean to check if the value has been set.
func (o *EsxiHighCpuThresholds) GetCpuPeakPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuPeakPercentage, true
}

// SetCpuPeakPercentage sets field value
func (o *EsxiHighCpuThresholds) SetCpuPeakPercentage(v int32) {
	o.CpuPeakPercentage = v
}

// GetCpuUsagePercentage returns the CpuUsagePercentage field value
func (o *EsxiHighCpuThresholds) GetCpuUsagePercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CpuUsagePercentage
}

// GetCpuUsagePercentageOk returns a tuple with the CpuUsagePercentage field value
// and a boolean to check if the value has been set.
func (o *EsxiHighCpuThresholds) GetCpuUsagePercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CpuUsagePercentage, true
}

// SetCpuUsagePercentage sets field value
func (o *EsxiHighCpuThresholds) SetCpuUsagePercentage(v int32) {
	o.CpuUsagePercentage = v
}

// GetVmCpuReadyPercentage returns the VmCpuReadyPercentage field value
func (o *EsxiHighCpuThresholds) GetVmCpuReadyPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VmCpuReadyPercentage
}

// GetVmCpuReadyPercentageOk returns a tuple with the VmCpuReadyPercentage field value
// and a boolean to check if the value has been set.
func (o *EsxiHighCpuThresholds) GetVmCpuReadyPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VmCpuReadyPercentage, true
}

// SetVmCpuReadyPercentage sets field value
func (o *EsxiHighCpuThresholds) SetVmCpuReadyPercentage(v int32) {
	o.VmCpuReadyPercentage = v
}

func (o EsxiHighCpuThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EsxiHighCpuThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpuPeakPercentage"] = o.CpuPeakPercentage
	toSerialize["cpuUsagePercentage"] = o.CpuUsagePercentage
	toSerialize["vmCpuReadyPercentage"] = o.VmCpuReadyPercentage
	return toSerialize, nil
}

type NullableEsxiHighCpuThresholds struct {
	value *EsxiHighCpuThresholds
	isSet bool
}

func (v NullableEsxiHighCpuThresholds) Get() *EsxiHighCpuThresholds {
	return v.value
}

func (v *NullableEsxiHighCpuThresholds) Set(val *EsxiHighCpuThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableEsxiHighCpuThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableEsxiHighCpuThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEsxiHighCpuThresholds(val *EsxiHighCpuThresholds) *NullableEsxiHighCpuThresholds {
	return &NullableEsxiHighCpuThresholds{value: val, isSet: true}
}

func (v NullableEsxiHighCpuThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEsxiHighCpuThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


