/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv1

import (
	"encoding/json"
)

// checks if the LogListForCustomDeviceResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LogListForCustomDeviceResult{}

// LogListForCustomDeviceResult Logs available on the Custom Device.
type LogListForCustomDeviceResult struct {
	// The list of available logs.
	Logs []LogForCustomDevice `json:"logs,omitempty"`
}

// NewLogListForCustomDeviceResult instantiates a new LogListForCustomDeviceResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogListForCustomDeviceResult() *LogListForCustomDeviceResult {
	this := LogListForCustomDeviceResult{}
	return &this
}

// NewLogListForCustomDeviceResultWithDefaults instantiates a new LogListForCustomDeviceResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogListForCustomDeviceResultWithDefaults() *LogListForCustomDeviceResult {
	this := LogListForCustomDeviceResult{}
	return &this
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *LogListForCustomDeviceResult) GetLogs() []LogForCustomDevice {
	if o == nil || IsNil(o.Logs) {
		var ret []LogForCustomDevice
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogListForCustomDeviceResult) GetLogsOk() ([]LogForCustomDevice, bool) {
	if o == nil || IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *LogListForCustomDeviceResult) HasLogs() bool {
	if o != nil && !IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []LogForCustomDevice and assigns it to the Logs field.
func (o *LogListForCustomDeviceResult) SetLogs(v []LogForCustomDevice) {
	o.Logs = v
}

func (o LogListForCustomDeviceResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LogListForCustomDeviceResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Logs) {
		toSerialize["logs"] = o.Logs
	}
	return toSerialize, nil
}

type NullableLogListForCustomDeviceResult struct {
	value *LogListForCustomDeviceResult
	isSet bool
}

func (v NullableLogListForCustomDeviceResult) Get() *LogListForCustomDeviceResult {
	return v.value
}

func (v *NullableLogListForCustomDeviceResult) Set(val *LogListForCustomDeviceResult) {
	v.value = val
	v.isSet = true
}

func (v NullableLogListForCustomDeviceResult) IsSet() bool {
	return v.isSet
}

func (v *NullableLogListForCustomDeviceResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogListForCustomDeviceResult(val *LogListForCustomDeviceResult) *NullableLogListForCustomDeviceResult {
	return &NullableLogListForCustomDeviceResult{value: val, isSet: true}
}

func (v NullableLogListForCustomDeviceResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogListForCustomDeviceResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


