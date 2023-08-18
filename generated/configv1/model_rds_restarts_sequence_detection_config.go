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

// checks if the RdsRestartsSequenceDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RdsRestartsSequenceDetectionConfig{}

// RdsRestartsSequenceDetectionConfig The configuration of the restarts sequence on RDS detection.
type RdsRestartsSequenceDetectionConfig struct {
	CustomThresholds *RdsRestartsThresholds `json:"customThresholds,omitempty"`
	// The detection is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
}

// NewRdsRestartsSequenceDetectionConfig instantiates a new RdsRestartsSequenceDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRdsRestartsSequenceDetectionConfig(enabled bool) *RdsRestartsSequenceDetectionConfig {
	this := RdsRestartsSequenceDetectionConfig{}
	this.Enabled = enabled
	return &this
}

// NewRdsRestartsSequenceDetectionConfigWithDefaults instantiates a new RdsRestartsSequenceDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRdsRestartsSequenceDetectionConfigWithDefaults() *RdsRestartsSequenceDetectionConfig {
	this := RdsRestartsSequenceDetectionConfig{}
	return &this
}

// GetCustomThresholds returns the CustomThresholds field value if set, zero value otherwise.
func (o *RdsRestartsSequenceDetectionConfig) GetCustomThresholds() RdsRestartsThresholds {
	if o == nil || IsNil(o.CustomThresholds) {
		var ret RdsRestartsThresholds
		return ret
	}
	return *o.CustomThresholds
}

// GetCustomThresholdsOk returns a tuple with the CustomThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RdsRestartsSequenceDetectionConfig) GetCustomThresholdsOk() (*RdsRestartsThresholds, bool) {
	if o == nil || IsNil(o.CustomThresholds) {
		return nil, false
	}
	return o.CustomThresholds, true
}

// HasCustomThresholds returns a boolean if a field has been set.
func (o *RdsRestartsSequenceDetectionConfig) HasCustomThresholds() bool {
	if o != nil && !IsNil(o.CustomThresholds) {
		return true
	}

	return false
}

// SetCustomThresholds gets a reference to the given RdsRestartsThresholds and assigns it to the CustomThresholds field.
func (o *RdsRestartsSequenceDetectionConfig) SetCustomThresholds(v RdsRestartsThresholds) {
	o.CustomThresholds = &v
}

// GetEnabled returns the Enabled field value
func (o *RdsRestartsSequenceDetectionConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RdsRestartsSequenceDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *RdsRestartsSequenceDetectionConfig) SetEnabled(v bool) {
	o.Enabled = v
}

func (o RdsRestartsSequenceDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RdsRestartsSequenceDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomThresholds) {
		toSerialize["customThresholds"] = o.CustomThresholds
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableRdsRestartsSequenceDetectionConfig struct {
	value *RdsRestartsSequenceDetectionConfig
	isSet bool
}

func (v NullableRdsRestartsSequenceDetectionConfig) Get() *RdsRestartsSequenceDetectionConfig {
	return v.value
}

func (v *NullableRdsRestartsSequenceDetectionConfig) Set(val *RdsRestartsSequenceDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableRdsRestartsSequenceDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableRdsRestartsSequenceDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRdsRestartsSequenceDetectionConfig(val *RdsRestartsSequenceDetectionConfig) *NullableRdsRestartsSequenceDetectionConfig {
	return &NullableRdsRestartsSequenceDetectionConfig{value: val, isSet: true}
}

func (v NullableRdsRestartsSequenceDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRdsRestartsSequenceDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

