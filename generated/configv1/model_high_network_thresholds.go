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

// checks if the HighNetworkThresholds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HighNetworkThresholds{}

// HighNetworkThresholds Custom thresholds for high network utilization. If not set, automatic mode is used.
type HighNetworkThresholds struct {
	// Alert if sent/received traffic utilization is higher than *X*% in 3 out of 5 samples.
	UtilizationPercentage int32 `json:"utilizationPercentage"`
}

// NewHighNetworkThresholds instantiates a new HighNetworkThresholds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHighNetworkThresholds(utilizationPercentage int32) *HighNetworkThresholds {
	this := HighNetworkThresholds{}
	this.UtilizationPercentage = utilizationPercentage
	return &this
}

// NewHighNetworkThresholdsWithDefaults instantiates a new HighNetworkThresholds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHighNetworkThresholdsWithDefaults() *HighNetworkThresholds {
	this := HighNetworkThresholds{}
	return &this
}

// GetUtilizationPercentage returns the UtilizationPercentage field value
func (o *HighNetworkThresholds) GetUtilizationPercentage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UtilizationPercentage
}

// GetUtilizationPercentageOk returns a tuple with the UtilizationPercentage field value
// and a boolean to check if the value has been set.
func (o *HighNetworkThresholds) GetUtilizationPercentageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UtilizationPercentage, true
}

// SetUtilizationPercentage sets field value
func (o *HighNetworkThresholds) SetUtilizationPercentage(v int32) {
	o.UtilizationPercentage = v
}

func (o HighNetworkThresholds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HighNetworkThresholds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["utilizationPercentage"] = o.UtilizationPercentage
	return toSerialize, nil
}

type NullableHighNetworkThresholds struct {
	value *HighNetworkThresholds
	isSet bool
}

func (v NullableHighNetworkThresholds) Get() *HighNetworkThresholds {
	return v.value
}

func (v *NullableHighNetworkThresholds) Set(val *HighNetworkThresholds) {
	v.value = val
	v.isSet = true
}

func (v NullableHighNetworkThresholds) IsSet() bool {
	return v.isSet
}

func (v *NullableHighNetworkThresholds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighNetworkThresholds(val *HighNetworkThresholds) *NullableHighNetworkThresholds {
	return &NullableHighNetworkThresholds{value: val, isSet: true}
}

func (v NullableHighNetworkThresholds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHighNetworkThresholds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


