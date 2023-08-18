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

// checks if the ExceptionPattern type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExceptionPattern{}

// ExceptionPattern A list of faulty exceptions.   If an exception on *any* node of the service matches *any* of these patterns it is considered a failure.
type ExceptionPattern struct {
	ClassPattern *string `json:"classPattern,omitempty"`
	MessagePattern *string `json:"messagePattern,omitempty"`
}

// NewExceptionPattern instantiates a new ExceptionPattern object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionPattern() *ExceptionPattern {
	this := ExceptionPattern{}
	return &this
}

// NewExceptionPatternWithDefaults instantiates a new ExceptionPattern object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionPatternWithDefaults() *ExceptionPattern {
	this := ExceptionPattern{}
	return &this
}

// GetClassPattern returns the ClassPattern field value if set, zero value otherwise.
func (o *ExceptionPattern) GetClassPattern() string {
	if o == nil || IsNil(o.ClassPattern) {
		var ret string
		return ret
	}
	return *o.ClassPattern
}

// GetClassPatternOk returns a tuple with the ClassPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionPattern) GetClassPatternOk() (*string, bool) {
	if o == nil || IsNil(o.ClassPattern) {
		return nil, false
	}
	return o.ClassPattern, true
}

// HasClassPattern returns a boolean if a field has been set.
func (o *ExceptionPattern) HasClassPattern() bool {
	if o != nil && !IsNil(o.ClassPattern) {
		return true
	}

	return false
}

// SetClassPattern gets a reference to the given string and assigns it to the ClassPattern field.
func (o *ExceptionPattern) SetClassPattern(v string) {
	o.ClassPattern = &v
}

// GetMessagePattern returns the MessagePattern field value if set, zero value otherwise.
func (o *ExceptionPattern) GetMessagePattern() string {
	if o == nil || IsNil(o.MessagePattern) {
		var ret string
		return ret
	}
	return *o.MessagePattern
}

// GetMessagePatternOk returns a tuple with the MessagePattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionPattern) GetMessagePatternOk() (*string, bool) {
	if o == nil || IsNil(o.MessagePattern) {
		return nil, false
	}
	return o.MessagePattern, true
}

// HasMessagePattern returns a boolean if a field has been set.
func (o *ExceptionPattern) HasMessagePattern() bool {
	if o != nil && !IsNil(o.MessagePattern) {
		return true
	}

	return false
}

// SetMessagePattern gets a reference to the given string and assigns it to the MessagePattern field.
func (o *ExceptionPattern) SetMessagePattern(v string) {
	o.MessagePattern = &v
}

func (o ExceptionPattern) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExceptionPattern) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClassPattern) {
		toSerialize["classPattern"] = o.ClassPattern
	}
	if !IsNil(o.MessagePattern) {
		toSerialize["messagePattern"] = o.MessagePattern
	}
	return toSerialize, nil
}

type NullableExceptionPattern struct {
	value *ExceptionPattern
	isSet bool
}

func (v NullableExceptionPattern) Get() *ExceptionPattern {
	return v.value
}

func (v *NullableExceptionPattern) Set(val *ExceptionPattern) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionPattern) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionPattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionPattern(val *ExceptionPattern) *NullableExceptionPattern {
	return &NullableExceptionPattern{value: val, isSet: true}
}

func (v NullableExceptionPattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionPattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


