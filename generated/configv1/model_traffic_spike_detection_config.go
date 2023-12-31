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

// checks if the TrafficSpikeDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficSpikeDetectionConfig{}

// TrafficSpikeDetectionConfig The configuration of traffic spikes detection.
type TrafficSpikeDetectionConfig struct {
	// The detection is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// Alert if the observed traffic is more than *X* % of the expected value.
	TrafficSpikePercent *int32 `json:"trafficSpikePercent,omitempty"`
}

// NewTrafficSpikeDetectionConfig instantiates a new TrafficSpikeDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficSpikeDetectionConfig(enabled bool) *TrafficSpikeDetectionConfig {
	this := TrafficSpikeDetectionConfig{}
	this.Enabled = enabled
	return &this
}

// NewTrafficSpikeDetectionConfigWithDefaults instantiates a new TrafficSpikeDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficSpikeDetectionConfigWithDefaults() *TrafficSpikeDetectionConfig {
	this := TrafficSpikeDetectionConfig{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *TrafficSpikeDetectionConfig) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *TrafficSpikeDetectionConfig) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *TrafficSpikeDetectionConfig) SetEnabled(v bool) {
	o.Enabled = v
}

// GetTrafficSpikePercent returns the TrafficSpikePercent field value if set, zero value otherwise.
func (o *TrafficSpikeDetectionConfig) GetTrafficSpikePercent() int32 {
	if o == nil || IsNil(o.TrafficSpikePercent) {
		var ret int32
		return ret
	}
	return *o.TrafficSpikePercent
}

// GetTrafficSpikePercentOk returns a tuple with the TrafficSpikePercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficSpikeDetectionConfig) GetTrafficSpikePercentOk() (*int32, bool) {
	if o == nil || IsNil(o.TrafficSpikePercent) {
		return nil, false
	}
	return o.TrafficSpikePercent, true
}

// HasTrafficSpikePercent returns a boolean if a field has been set.
func (o *TrafficSpikeDetectionConfig) HasTrafficSpikePercent() bool {
	if o != nil && !IsNil(o.TrafficSpikePercent) {
		return true
	}

	return false
}

// SetTrafficSpikePercent gets a reference to the given int32 and assigns it to the TrafficSpikePercent field.
func (o *TrafficSpikeDetectionConfig) SetTrafficSpikePercent(v int32) {
	o.TrafficSpikePercent = &v
}

func (o TrafficSpikeDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficSpikeDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.TrafficSpikePercent) {
		toSerialize["trafficSpikePercent"] = o.TrafficSpikePercent
	}
	return toSerialize, nil
}

type NullableTrafficSpikeDetectionConfig struct {
	value *TrafficSpikeDetectionConfig
	isSet bool
}

func (v NullableTrafficSpikeDetectionConfig) Get() *TrafficSpikeDetectionConfig {
	return v.value
}

func (v *NullableTrafficSpikeDetectionConfig) Set(val *TrafficSpikeDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficSpikeDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficSpikeDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficSpikeDetectionConfig(val *TrafficSpikeDetectionConfig) *NullableTrafficSpikeDetectionConfig {
	return &NullableTrafficSpikeDetectionConfig{value: val, isSet: true}
}

func (v NullableTrafficSpikeDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficSpikeDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


