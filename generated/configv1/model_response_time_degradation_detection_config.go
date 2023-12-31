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

// checks if the ResponseTimeDegradationDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTimeDegradationDetectionConfig{}

// ResponseTimeDegradationDetectionConfig Configuration of response time degradation detection.
type ResponseTimeDegradationDetectionConfig struct {
	AutomaticDetection *ResponseTimeDegradationAutodetectionConfig `json:"automaticDetection,omitempty"`
	// How to detect response time degradation: automatically, or based on fixed thresholds, or do not detect.
	DetectionMode string `json:"detectionMode"`
	Thresholds *ResponseTimeDegradationThresholdConfig `json:"thresholds,omitempty"`
}

// NewResponseTimeDegradationDetectionConfig instantiates a new ResponseTimeDegradationDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTimeDegradationDetectionConfig(detectionMode string) *ResponseTimeDegradationDetectionConfig {
	this := ResponseTimeDegradationDetectionConfig{}
	this.DetectionMode = detectionMode
	return &this
}

// NewResponseTimeDegradationDetectionConfigWithDefaults instantiates a new ResponseTimeDegradationDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTimeDegradationDetectionConfigWithDefaults() *ResponseTimeDegradationDetectionConfig {
	this := ResponseTimeDegradationDetectionConfig{}
	return &this
}

// GetAutomaticDetection returns the AutomaticDetection field value if set, zero value otherwise.
func (o *ResponseTimeDegradationDetectionConfig) GetAutomaticDetection() ResponseTimeDegradationAutodetectionConfig {
	if o == nil || IsNil(o.AutomaticDetection) {
		var ret ResponseTimeDegradationAutodetectionConfig
		return ret
	}
	return *o.AutomaticDetection
}

// GetAutomaticDetectionOk returns a tuple with the AutomaticDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTimeDegradationDetectionConfig) GetAutomaticDetectionOk() (*ResponseTimeDegradationAutodetectionConfig, bool) {
	if o == nil || IsNil(o.AutomaticDetection) {
		return nil, false
	}
	return o.AutomaticDetection, true
}

// HasAutomaticDetection returns a boolean if a field has been set.
func (o *ResponseTimeDegradationDetectionConfig) HasAutomaticDetection() bool {
	if o != nil && !IsNil(o.AutomaticDetection) {
		return true
	}

	return false
}

// SetAutomaticDetection gets a reference to the given ResponseTimeDegradationAutodetectionConfig and assigns it to the AutomaticDetection field.
func (o *ResponseTimeDegradationDetectionConfig) SetAutomaticDetection(v ResponseTimeDegradationAutodetectionConfig) {
	o.AutomaticDetection = &v
}

// GetDetectionMode returns the DetectionMode field value
func (o *ResponseTimeDegradationDetectionConfig) GetDetectionMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DetectionMode
}

// GetDetectionModeOk returns a tuple with the DetectionMode field value
// and a boolean to check if the value has been set.
func (o *ResponseTimeDegradationDetectionConfig) GetDetectionModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DetectionMode, true
}

// SetDetectionMode sets field value
func (o *ResponseTimeDegradationDetectionConfig) SetDetectionMode(v string) {
	o.DetectionMode = v
}

// GetThresholds returns the Thresholds field value if set, zero value otherwise.
func (o *ResponseTimeDegradationDetectionConfig) GetThresholds() ResponseTimeDegradationThresholdConfig {
	if o == nil || IsNil(o.Thresholds) {
		var ret ResponseTimeDegradationThresholdConfig
		return ret
	}
	return *o.Thresholds
}

// GetThresholdsOk returns a tuple with the Thresholds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseTimeDegradationDetectionConfig) GetThresholdsOk() (*ResponseTimeDegradationThresholdConfig, bool) {
	if o == nil || IsNil(o.Thresholds) {
		return nil, false
	}
	return o.Thresholds, true
}

// HasThresholds returns a boolean if a field has been set.
func (o *ResponseTimeDegradationDetectionConfig) HasThresholds() bool {
	if o != nil && !IsNil(o.Thresholds) {
		return true
	}

	return false
}

// SetThresholds gets a reference to the given ResponseTimeDegradationThresholdConfig and assigns it to the Thresholds field.
func (o *ResponseTimeDegradationDetectionConfig) SetThresholds(v ResponseTimeDegradationThresholdConfig) {
	o.Thresholds = &v
}

func (o ResponseTimeDegradationDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTimeDegradationDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutomaticDetection) {
		toSerialize["automaticDetection"] = o.AutomaticDetection
	}
	toSerialize["detectionMode"] = o.DetectionMode
	if !IsNil(o.Thresholds) {
		toSerialize["thresholds"] = o.Thresholds
	}
	return toSerialize, nil
}

type NullableResponseTimeDegradationDetectionConfig struct {
	value *ResponseTimeDegradationDetectionConfig
	isSet bool
}

func (v NullableResponseTimeDegradationDetectionConfig) Get() *ResponseTimeDegradationDetectionConfig {
	return v.value
}

func (v *NullableResponseTimeDegradationDetectionConfig) Set(val *ResponseTimeDegradationDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTimeDegradationDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTimeDegradationDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTimeDegradationDetectionConfig(val *ResponseTimeDegradationDetectionConfig) *NullableResponseTimeDegradationDetectionConfig {
	return &NullableResponseTimeDegradationDetectionConfig{value: val, isSet: true}
}

func (v NullableResponseTimeDegradationDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTimeDegradationDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


