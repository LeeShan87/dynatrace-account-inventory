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

// checks if the FdpTagIntegerGreaterThan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdpTagIntegerGreaterThan{}

// FdpTagIntegerGreaterThan The predicate of the `INTEGER_GREATER_THAN` type. It checks whether the tag (which is an integer) is greater than the reference value.
type FdpTagIntegerGreaterThan struct {
	// The reference value.
	Value *string `json:"value,omitempty"`
}

// NewFdpTagIntegerGreaterThan instantiates a new FdpTagIntegerGreaterThan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdpTagIntegerGreaterThan(type_ string) *FdpTagIntegerGreaterThan {
	this := FdpTagIntegerGreaterThan{}
	this.Type = type_
	return &this
}

// NewFdpTagIntegerGreaterThanWithDefaults instantiates a new FdpTagIntegerGreaterThan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdpTagIntegerGreaterThanWithDefaults() *FdpTagIntegerGreaterThan {
	this := FdpTagIntegerGreaterThan{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FdpTagIntegerGreaterThan) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FdpTagIntegerGreaterThan) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FdpTagIntegerGreaterThan) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FdpTagIntegerGreaterThan) SetValue(v string) {
	o.Value = &v
}

func (o FdpTagIntegerGreaterThan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdpTagIntegerGreaterThan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFdpTagIntegerGreaterThan struct {
	value *FdpTagIntegerGreaterThan
	isSet bool
}

func (v NullableFdpTagIntegerGreaterThan) Get() *FdpTagIntegerGreaterThan {
	return v.value
}

func (v *NullableFdpTagIntegerGreaterThan) Set(val *FdpTagIntegerGreaterThan) {
	v.value = val
	v.isSet = true
}

func (v NullableFdpTagIntegerGreaterThan) IsSet() bool {
	return v.isSet
}

func (v *NullableFdpTagIntegerGreaterThan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdpTagIntegerGreaterThan(val *FdpTagIntegerGreaterThan) *NullableFdpTagIntegerGreaterThan {
	return &NullableFdpTagIntegerGreaterThan{value: val, isSet: true}
}

func (v NullableFdpTagIntegerGreaterThan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdpTagIntegerGreaterThan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


