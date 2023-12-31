/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
)

// checks if the MetricIngestError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricIngestError{}

// MetricIngestError struct for MetricIngestError
type MetricIngestError struct {
	Code *int32 `json:"code,omitempty"`
	InvalidLines []InvalidLine `json:"invalidLines,omitempty"`
	Message *string `json:"message,omitempty"`
}

// NewMetricIngestError instantiates a new MetricIngestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricIngestError() *MetricIngestError {
	this := MetricIngestError{}
	return &this
}

// NewMetricIngestErrorWithDefaults instantiates a new MetricIngestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricIngestErrorWithDefaults() *MetricIngestError {
	this := MetricIngestError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MetricIngestError) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIngestError) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MetricIngestError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *MetricIngestError) SetCode(v int32) {
	o.Code = &v
}

// GetInvalidLines returns the InvalidLines field value if set, zero value otherwise.
func (o *MetricIngestError) GetInvalidLines() []InvalidLine {
	if o == nil || IsNil(o.InvalidLines) {
		var ret []InvalidLine
		return ret
	}
	return o.InvalidLines
}

// GetInvalidLinesOk returns a tuple with the InvalidLines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIngestError) GetInvalidLinesOk() ([]InvalidLine, bool) {
	if o == nil || IsNil(o.InvalidLines) {
		return nil, false
	}
	return o.InvalidLines, true
}

// HasInvalidLines returns a boolean if a field has been set.
func (o *MetricIngestError) HasInvalidLines() bool {
	if o != nil && !IsNil(o.InvalidLines) {
		return true
	}

	return false
}

// SetInvalidLines gets a reference to the given []InvalidLine and assigns it to the InvalidLines field.
func (o *MetricIngestError) SetInvalidLines(v []InvalidLine) {
	o.InvalidLines = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MetricIngestError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricIngestError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MetricIngestError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MetricIngestError) SetMessage(v string) {
	o.Message = &v
}

func (o MetricIngestError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricIngestError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.InvalidLines) {
		toSerialize["invalidLines"] = o.InvalidLines
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableMetricIngestError struct {
	value *MetricIngestError
	isSet bool
}

func (v NullableMetricIngestError) Get() *MetricIngestError {
	return v.value
}

func (v *NullableMetricIngestError) Set(val *MetricIngestError) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricIngestError) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricIngestError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricIngestError(val *MetricIngestError) *NullableMetricIngestError {
	return &NullableMetricIngestError{value: val, isSet: true}
}

func (v NullableMetricIngestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricIngestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


