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

// checks if the ResponseTimeDegradationAutodetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseTimeDegradationAutodetectionConfig{}

// ResponseTimeDegradationAutodetectionConfig Parameters of the response time degradation auto-detection. Required if the **detectionMode** is `DETECT_AUTOMATICALLY`. Not applicable otherwise.   Violation of **any** criterion triggers an alert.
type ResponseTimeDegradationAutodetectionConfig struct {
	// Minimal service load to detect response time degradation.    Response time degradation of services with smaller load won't trigger alerts.
	LoadThreshold string `json:"loadThreshold"`
	// Alert if the response time degrades beyond *X* milliseconds.
	ResponseTimeDegradationMilliseconds int32 `json:"responseTimeDegradationMilliseconds"`
	// Alert if the response time degrades beyond *X* %.
	ResponseTimeDegradationPercent int32 `json:"responseTimeDegradationPercent"`
	// Alert if the response time of the slowest 10% degrades beyond *X* milliseconds.
	SlowestResponseTimeDegradationMilliseconds int32 `json:"slowestResponseTimeDegradationMilliseconds"`
	// Alert if the response time of the slowest 10% degrades beyond *X* %.
	SlowestResponseTimeDegradationPercent int32 `json:"slowestResponseTimeDegradationPercent"`
}

// NewResponseTimeDegradationAutodetectionConfig instantiates a new ResponseTimeDegradationAutodetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseTimeDegradationAutodetectionConfig(loadThreshold string, responseTimeDegradationMilliseconds int32, responseTimeDegradationPercent int32, slowestResponseTimeDegradationMilliseconds int32, slowestResponseTimeDegradationPercent int32) *ResponseTimeDegradationAutodetectionConfig {
	this := ResponseTimeDegradationAutodetectionConfig{}
	this.LoadThreshold = loadThreshold
	this.ResponseTimeDegradationMilliseconds = responseTimeDegradationMilliseconds
	this.ResponseTimeDegradationPercent = responseTimeDegradationPercent
	this.SlowestResponseTimeDegradationMilliseconds = slowestResponseTimeDegradationMilliseconds
	this.SlowestResponseTimeDegradationPercent = slowestResponseTimeDegradationPercent
	return &this
}

// NewResponseTimeDegradationAutodetectionConfigWithDefaults instantiates a new ResponseTimeDegradationAutodetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseTimeDegradationAutodetectionConfigWithDefaults() *ResponseTimeDegradationAutodetectionConfig {
	this := ResponseTimeDegradationAutodetectionConfig{}
	return &this
}

// GetLoadThreshold returns the LoadThreshold field value
func (o *ResponseTimeDegradationAutodetectionConfig) GetLoadThreshold() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LoadThreshold
}

// GetLoadThresholdOk returns a tuple with the LoadThreshold field value
// and a boolean to check if the value has been set.
func (o *ResponseTimeDegradationAutodetectionConfig) GetLoadThresholdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LoadThreshold, true
}

// SetLoadThreshold sets field value
func (o *ResponseTimeDegradationAutodetectionConfig) SetLoadThreshold(v string) {
	o.LoadThreshold = v
}

// GetResponseTimeDegradationMilliseconds returns the ResponseTimeDegradationMilliseconds field value
func (o *ResponseTimeDegradationAutodetectionConfig) GetResponseTimeDegradationMilliseconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseTimeDegradationMilliseconds
}

// GetResponseTimeDegradationMillisecondsOk returns a tuple with the ResponseTimeDegradationMilliseconds field value
// and a boolean to check if the value has been set.
func (o *ResponseTimeDegradationAutodetectionConfig) GetResponseTimeDegradationMillisecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseTimeDegradationMilliseconds, true
}

// SetResponseTimeDegradationMilliseconds sets field value
func (o *ResponseTimeDegradationAutodetectionConfig) SetResponseTimeDegradationMilliseconds(v int32) {
	o.ResponseTimeDegradationMilliseconds = v
}

// GetResponseTimeDegradationPercent returns the ResponseTimeDegradationPercent field value
func (o *ResponseTimeDegradationAutodetectionConfig) GetResponseTimeDegradationPercent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ResponseTimeDegradationPercent
}

// GetResponseTimeDegradationPercentOk returns a tuple with the ResponseTimeDegradationPercent field value
// and a boolean to check if the value has been set.
func (o *ResponseTimeDegradationAutodetectionConfig) GetResponseTimeDegradationPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseTimeDegradationPercent, true
}

// SetResponseTimeDegradationPercent sets field value
func (o *ResponseTimeDegradationAutodetectionConfig) SetResponseTimeDegradationPercent(v int32) {
	o.ResponseTimeDegradationPercent = v
}

// GetSlowestResponseTimeDegradationMilliseconds returns the SlowestResponseTimeDegradationMilliseconds field value
func (o *ResponseTimeDegradationAutodetectionConfig) GetSlowestResponseTimeDegradationMilliseconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SlowestResponseTimeDegradationMilliseconds
}

// GetSlowestResponseTimeDegradationMillisecondsOk returns a tuple with the SlowestResponseTimeDegradationMilliseconds field value
// and a boolean to check if the value has been set.
func (o *ResponseTimeDegradationAutodetectionConfig) GetSlowestResponseTimeDegradationMillisecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowestResponseTimeDegradationMilliseconds, true
}

// SetSlowestResponseTimeDegradationMilliseconds sets field value
func (o *ResponseTimeDegradationAutodetectionConfig) SetSlowestResponseTimeDegradationMilliseconds(v int32) {
	o.SlowestResponseTimeDegradationMilliseconds = v
}

// GetSlowestResponseTimeDegradationPercent returns the SlowestResponseTimeDegradationPercent field value
func (o *ResponseTimeDegradationAutodetectionConfig) GetSlowestResponseTimeDegradationPercent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SlowestResponseTimeDegradationPercent
}

// GetSlowestResponseTimeDegradationPercentOk returns a tuple with the SlowestResponseTimeDegradationPercent field value
// and a boolean to check if the value has been set.
func (o *ResponseTimeDegradationAutodetectionConfig) GetSlowestResponseTimeDegradationPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SlowestResponseTimeDegradationPercent, true
}

// SetSlowestResponseTimeDegradationPercent sets field value
func (o *ResponseTimeDegradationAutodetectionConfig) SetSlowestResponseTimeDegradationPercent(v int32) {
	o.SlowestResponseTimeDegradationPercent = v
}

func (o ResponseTimeDegradationAutodetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseTimeDegradationAutodetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["loadThreshold"] = o.LoadThreshold
	toSerialize["responseTimeDegradationMilliseconds"] = o.ResponseTimeDegradationMilliseconds
	toSerialize["responseTimeDegradationPercent"] = o.ResponseTimeDegradationPercent
	toSerialize["slowestResponseTimeDegradationMilliseconds"] = o.SlowestResponseTimeDegradationMilliseconds
	toSerialize["slowestResponseTimeDegradationPercent"] = o.SlowestResponseTimeDegradationPercent
	return toSerialize, nil
}

type NullableResponseTimeDegradationAutodetectionConfig struct {
	value *ResponseTimeDegradationAutodetectionConfig
	isSet bool
}

func (v NullableResponseTimeDegradationAutodetectionConfig) Get() *ResponseTimeDegradationAutodetectionConfig {
	return v.value
}

func (v *NullableResponseTimeDegradationAutodetectionConfig) Set(val *ResponseTimeDegradationAutodetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseTimeDegradationAutodetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseTimeDegradationAutodetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseTimeDegradationAutodetectionConfig(val *ResponseTimeDegradationAutodetectionConfig) *NullableResponseTimeDegradationAutodetectionConfig {
	return &NullableResponseTimeDegradationAutodetectionConfig{value: val, isSet: true}
}

func (v NullableResponseTimeDegradationAutodetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseTimeDegradationAutodetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

