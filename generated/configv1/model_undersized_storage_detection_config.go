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

// checks if the UndersizedStorageDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UndersizedStorageDetectionConfig{}

// UndersizedStorageDetectionConfig Undersized storage device detection cofing
type UndersizedStorageDetectionConfig struct {
	CustomThresholds *UndersizedStorageThresholds `json:"customThresholds,omitempty"`
	// The detection is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
}

// NewUndersizedStorageDetectionConfig instantiates a new UndersizedStorageDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUndersizedStorageDetectionConfig(enabled bool) *UndersizedStorageDetectionConfig {
	this := UndersizedStorageDetectionConfig{}
	this.Enabled = enabled
	return &this
}

// NewUndersizedStorageDetectionConfigWithDefaults instantiates a new UndersizedStorageDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUndersizedStorageDetectionConfigWithDefaults() *UndersizedStorageDetectionConfig {
	this := UndersizedStorageDetectionConfig{}
	return &this
}

// GetCustomThresholds returns the CustomThresholds field value if set, zero value otherwise.
func (o *UndersizedStorageDetectionConfig) GetCustomThresholds() UndersizedStorageThresholds {
	if o == nil || IsNil(o.CustomThresholds) {
		var ret UndersizedStorageThresholds
		return ret
	}
	return *o.CustomThresholds
}

// GetCustomThresholdsOk returns a tuple with the CustomThresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UndersizedStorageDetectionConfig) GetCustomThresholdsOk() (*UndersizedStorageThresholds, bool) {
	if o == nil || IsNil(o.CustomThresholds) {
		return nil, false
	}
	return o.CustomThresholds, true
}

// HasCustomThresholds returns a boolean if a field has been set.
func (o *UndersizedStorageDetectionConfig) HasCustomThresholds() bool {
	if o != nil && !IsNil(o.CustomThresholds) {
		return true
	}

	return false
}

// SetCustomThresholds gets a reference to the given UndersizedStorageThresholds and assigns it to the CustomThresholds field.
func (o *UndersizedStorageDetectionConfig) SetCustomThresholds(v UndersizedStorageThresholds) {
	o.CustomThresholds = &v
}

// GetEnabled returns the Enabled field value
func (o *UndersizedStorageDetectionConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *UndersizedStorageDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *UndersizedStorageDetectionConfig) SetEnabled(v bool) {
	o.Enabled = v
}

func (o UndersizedStorageDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UndersizedStorageDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomThresholds) {
		toSerialize["customThresholds"] = o.CustomThresholds
	}
	toSerialize["enabled"] = o.Enabled
	return toSerialize, nil
}

type NullableUndersizedStorageDetectionConfig struct {
	value *UndersizedStorageDetectionConfig
	isSet bool
}

func (v NullableUndersizedStorageDetectionConfig) Get() *UndersizedStorageDetectionConfig {
	return v.value
}

func (v *NullableUndersizedStorageDetectionConfig) Set(val *UndersizedStorageDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUndersizedStorageDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUndersizedStorageDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUndersizedStorageDetectionConfig(val *UndersizedStorageDetectionConfig) *NullableUndersizedStorageDetectionConfig {
	return &NullableUndersizedStorageDetectionConfig{value: val, isSet: true}
}

func (v NullableUndersizedStorageDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUndersizedStorageDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


