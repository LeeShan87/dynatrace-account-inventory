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

// checks if the FdpTagDoubleEquals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdpTagDoubleEquals{}

// FdpTagDoubleEquals The predicate of the `DOUBLE_EQUALS` type. It checks whether the tag (which is a double) equals the reference value.
type FdpTagDoubleEquals struct {
	Negated *bool `json:"negated,omitempty"`
	// The reference value.
	Value *float64 `json:"value,omitempty"`
}

// NewFdpTagDoubleEquals instantiates a new FdpTagDoubleEquals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdpTagDoubleEquals(type_ string) *FdpTagDoubleEquals {
	this := FdpTagDoubleEquals{}
	this.Type = type_
	return &this
}

// NewFdpTagDoubleEqualsWithDefaults instantiates a new FdpTagDoubleEquals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdpTagDoubleEqualsWithDefaults() *FdpTagDoubleEquals {
	this := FdpTagDoubleEquals{}
	return &this
}

// GetNegated returns the Negated field value if set, zero value otherwise.
func (o *FdpTagDoubleEquals) GetNegated() bool {
	if o == nil || IsNil(o.Negated) {
		var ret bool
		return ret
	}
	return *o.Negated
}

// GetNegatedOk returns a tuple with the Negated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FdpTagDoubleEquals) GetNegatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Negated) {
		return nil, false
	}
	return o.Negated, true
}

// HasNegated returns a boolean if a field has been set.
func (o *FdpTagDoubleEquals) HasNegated() bool {
	if o != nil && !IsNil(o.Negated) {
		return true
	}

	return false
}

// SetNegated gets a reference to the given bool and assigns it to the Negated field.
func (o *FdpTagDoubleEquals) SetNegated(v bool) {
	o.Negated = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FdpTagDoubleEquals) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FdpTagDoubleEquals) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FdpTagDoubleEquals) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *FdpTagDoubleEquals) SetValue(v float64) {
	o.Value = &v
}

func (o FdpTagDoubleEquals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdpTagDoubleEquals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Negated) {
		toSerialize["negated"] = o.Negated
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFdpTagDoubleEquals struct {
	value *FdpTagDoubleEquals
	isSet bool
}

func (v NullableFdpTagDoubleEquals) Get() *FdpTagDoubleEquals {
	return v.value
}

func (v *NullableFdpTagDoubleEquals) Set(val *FdpTagDoubleEquals) {
	v.value = val
	v.isSet = true
}

func (v NullableFdpTagDoubleEquals) IsSet() bool {
	return v.isSet
}

func (v *NullableFdpTagDoubleEquals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdpTagDoubleEquals(val *FdpTagDoubleEquals) *NullableFdpTagDoubleEquals {
	return &NullableFdpTagDoubleEquals{value: val, isSet: true}
}

func (v NullableFdpTagDoubleEquals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdpTagDoubleEquals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

