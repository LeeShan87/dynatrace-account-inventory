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

// checks if the ErrorEnvelope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorEnvelope{}

// ErrorEnvelope struct for ErrorEnvelope
type ErrorEnvelope struct {
	Error *Error `json:"error,omitempty"`
}

// NewErrorEnvelope instantiates a new ErrorEnvelope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorEnvelope() *ErrorEnvelope {
	this := ErrorEnvelope{}
	return &this
}

// NewErrorEnvelopeWithDefaults instantiates a new ErrorEnvelope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorEnvelopeWithDefaults() *ErrorEnvelope {
	this := ErrorEnvelope{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ErrorEnvelope) GetError() Error {
	if o == nil || IsNil(o.Error) {
		var ret Error
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorEnvelope) GetErrorOk() (*Error, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ErrorEnvelope) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given Error and assigns it to the Error field.
func (o *ErrorEnvelope) SetError(v Error) {
	o.Error = &v
}

func (o ErrorEnvelope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorEnvelope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableErrorEnvelope struct {
	value *ErrorEnvelope
	isSet bool
}

func (v NullableErrorEnvelope) Get() *ErrorEnvelope {
	return v.value
}

func (v *NullableErrorEnvelope) Set(val *ErrorEnvelope) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorEnvelope) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorEnvelope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorEnvelope(val *ErrorEnvelope) *NullableErrorEnvelope {
	return &NullableErrorEnvelope{value: val, isSet: true}
}

func (v NullableErrorEnvelope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorEnvelope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


