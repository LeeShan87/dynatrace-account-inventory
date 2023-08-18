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

// checks if the HighMemoryDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HighMemoryDetectionConfig{}

// HighMemoryDetectionConfig Configuration of high memory usage detection.
type HighMemoryDetectionConfig struct {
	CustomThresholds *HighMemoryThresholds `json:"customThresholds,omitempty"`
	// The detection is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
}

// NewHighMemoryDetectionConfig instantiates a new HighMemoryDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHighMemoryDetectionConfig(enabled bool) *HighMemoryDetectionConfig {
	this := HighMemoryDetectionConfig{}
	this.Enabled = enabled
	return &this
}

// NewHighMemoryDetectionConfigWithDefaults instantiates a new HighMemoryDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHighMemoryDetectionConfigWithDefaults() *HighMemoryDetectionConfig {
	this := HighMemoryDetectionConfig{}
	return &this
}

// GetCustomThresholds returns the CustomThresholds field value if set, zero value otherwise.
func (o *HighMemoryDetectionConfig) GetCustomThresholds() HighMemoryThresholds {
	if o == nil || IsNil(o.CustomThresholds) {
		var ret HighMemoryThresholds
		return ret
	}
	return *o.CustomThresholds
}

// GetCustomThresholdsOk returns a tuple with the CustomThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HighMemoryDetectionConfig) GetCustomThresholdsOk() (*HighMemoryThresholds, bool) {
	if o == nil || IsNil(o.CustomThresholds) {
		return nil, false
	}
	return o.CustomThresholds, true
}

// HasCustomThresholds returns a boolean if a field has been set.
func (o *HighMemoryDetectionConfig) HasCustomThresholds() bool {
	if o != nil && !IsNil(o.CustomThresholds) {
		return true
	}

	return false
}

// SetCustomThresholds gets a reference to the given HighMemoryThresholds and assigns it to the CustomThresholds field.
func (o *HighMemoryDetectionConfig) SetCustomThresholds(v HighMemoryThresholds) {
	o.CustomThresholds = &v
}

// GetEnabled returns the Enabled field value
func (o *HighMemoryDetectionConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *HighMemoryDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *HighMemoryDetectionConfig) SetEnabled(v bool) {
	o.Enabled = v
}

func (o HighMemoryDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HighMemoryDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomThresholds) {
		toSerialize["customThresholds"] = o.CustomThresholds
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableHighMemoryDetectionConfig struct {
	value *HighMemoryDetectionConfig
	isSet bool
}

func (v NullableHighMemoryDetectionConfig) Get() *HighMemoryDetectionConfig {
	return v.value
}

func (v *NullableHighMemoryDetectionConfig) Set(val *HighMemoryDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHighMemoryDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHighMemoryDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHighMemoryDetectionConfig(val *HighMemoryDetectionConfig) *NullableHighMemoryDetectionConfig {
	return &NullableHighMemoryDetectionConfig{value: val, isSet: true}
}

func (v NullableHighMemoryDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHighMemoryDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

