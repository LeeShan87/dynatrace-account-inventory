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

// checks if the FdpTagIntegerGreaterThanOrEqual type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdpTagIntegerGreaterThanOrEqual{}

// FdpTagIntegerGreaterThanOrEqual The predicate of the `INTEGER_GREATER_THAN_OR_EQUAL` type. It checks whether the tag (which is an integer) is greater than or equals the reference value.
type FdpTagIntegerGreaterThanOrEqual struct {
	// The reference value.
	Value *string `json:"value,omitempty"`
}

// NewFdpTagIntegerGreaterThanOrEqual instantiates a new FdpTagIntegerGreaterThanOrEqual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdpTagIntegerGreaterThanOrEqual(type_ string) *FdpTagIntegerGreaterThanOrEqual {
	this := FdpTagIntegerGreaterThanOrEqual{}
	this.Type = type_
	return &this
}

// NewFdpTagIntegerGreaterThanOrEqualWithDefaults instantiates a new FdpTagIntegerGreaterThanOrEqual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdpTagIntegerGreaterThanOrEqualWithDefaults() *FdpTagIntegerGreaterThanOrEqual {
	this := FdpTagIntegerGreaterThanOrEqual{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FdpTagIntegerGreaterThanOrEqual) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FdpTagIntegerGreaterThanOrEqual) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FdpTagIntegerGreaterThanOrEqual) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FdpTagIntegerGreaterThanOrEqual) SetValue(v string) {
	o.Value = &v
}

func (o FdpTagIntegerGreaterThanOrEqual) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdpTagIntegerGreaterThanOrEqual) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFdpTagIntegerGreaterThanOrEqual struct {
	value *FdpTagIntegerGreaterThanOrEqual
	isSet bool
}

func (v NullableFdpTagIntegerGreaterThanOrEqual) Get() *FdpTagIntegerGreaterThanOrEqual {
	return v.value
}

func (v *NullableFdpTagIntegerGreaterThanOrEqual) Set(val *FdpTagIntegerGreaterThanOrEqual) {
	v.value = val
	v.isSet = true
}

func (v NullableFdpTagIntegerGreaterThanOrEqual) IsSet() bool {
	return v.isSet
}

func (v *NullableFdpTagIntegerGreaterThanOrEqual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdpTagIntegerGreaterThanOrEqual(val *FdpTagIntegerGreaterThanOrEqual) *NullableFdpTagIntegerGreaterThanOrEqual {
	return &NullableFdpTagIntegerGreaterThanOrEqual{value: val, isSet: true}
}

func (v NullableFdpTagIntegerGreaterThanOrEqual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdpTagIntegerGreaterThanOrEqual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


