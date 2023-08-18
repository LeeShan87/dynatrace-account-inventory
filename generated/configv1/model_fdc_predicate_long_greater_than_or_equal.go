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

// checks if the FdcPredicateLongGreaterThanOrEqual type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdcPredicateLongGreaterThanOrEqual{}

// FdcPredicateLongGreaterThanOrEqual The predicate of the `LONG_GREATER_THAN_OR_EQUAL` type. It checks whether the attribute (which is a long) is greater than or equals the reference value.
type FdcPredicateLongGreaterThanOrEqual struct {
	// The reference value. The condition is fulfilled when the attribute (which is a long) is greater than or equals this value.
	Value *string `json:"value,omitempty"`
}

// NewFdcPredicateLongGreaterThanOrEqual instantiates a new FdcPredicateLongGreaterThanOrEqual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdcPredicateLongGreaterThanOrEqual(type_ string) *FdcPredicateLongGreaterThanOrEqual {
	this := FdcPredicateLongGreaterThanOrEqual{}
	this.Type = type_
	return &this
}

// NewFdcPredicateLongGreaterThanOrEqualWithDefaults instantiates a new FdcPredicateLongGreaterThanOrEqual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdcPredicateLongGreaterThanOrEqualWithDefaults() *FdcPredicateLongGreaterThanOrEqual {
	this := FdcPredicateLongGreaterThanOrEqual{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FdcPredicateLongGreaterThanOrEqual) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FdcPredicateLongGreaterThanOrEqual) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FdcPredicateLongGreaterThanOrEqual) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FdcPredicateLongGreaterThanOrEqual) SetValue(v string) {
	o.Value = &v
}

func (o FdcPredicateLongGreaterThanOrEqual) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdcPredicateLongGreaterThanOrEqual) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFdcPredicateLongGreaterThanOrEqual struct {
	value *FdcPredicateLongGreaterThanOrEqual
	isSet bool
}

func (v NullableFdcPredicateLongGreaterThanOrEqual) Get() *FdcPredicateLongGreaterThanOrEqual {
	return v.value
}

func (v *NullableFdcPredicateLongGreaterThanOrEqual) Set(val *FdcPredicateLongGreaterThanOrEqual) {
	v.value = val
	v.isSet = true
}

func (v NullableFdcPredicateLongGreaterThanOrEqual) IsSet() bool {
	return v.isSet
}

func (v *NullableFdcPredicateLongGreaterThanOrEqual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdcPredicateLongGreaterThanOrEqual(val *FdcPredicateLongGreaterThanOrEqual) *NullableFdcPredicateLongGreaterThanOrEqual {
	return &NullableFdcPredicateLongGreaterThanOrEqual{value: val, isSet: true}
}

func (v NullableFdcPredicateLongGreaterThanOrEqual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdcPredicateLongGreaterThanOrEqual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


