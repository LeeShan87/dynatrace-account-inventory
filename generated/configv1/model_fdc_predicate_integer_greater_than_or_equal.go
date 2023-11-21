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

// checks if the FdcPredicateIntegerGreaterThanOrEqual type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdcPredicateIntegerGreaterThanOrEqual{}

// FdcPredicateIntegerGreaterThanOrEqual The predicate of the `INTEGER_GREATER_THAN_OR_EQUAL` type. It checks whether the attribute (which is an integer) is greater than or equals the reference value.
type FdcPredicateIntegerGreaterThanOrEqual struct {
	// The reference value. The condition is fulfilled when the attribute (which is an integer) is greater than or equals this value.
	Value *int32 `json:"value,omitempty"`
}

// NewFdcPredicateIntegerGreaterThanOrEqual instantiates a new FdcPredicateIntegerGreaterThanOrEqual object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdcPredicateIntegerGreaterThanOrEqual(type_ string) *FdcPredicateIntegerGreaterThanOrEqual {
	this := FdcPredicateIntegerGreaterThanOrEqual{}
	this.Type = type_
	return &this
}

// NewFdcPredicateIntegerGreaterThanOrEqualWithDefaults instantiates a new FdcPredicateIntegerGreaterThanOrEqual object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdcPredicateIntegerGreaterThanOrEqualWithDefaults() *FdcPredicateIntegerGreaterThanOrEqual {
	this := FdcPredicateIntegerGreaterThanOrEqual{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FdcPredicateIntegerGreaterThanOrEqual) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FdcPredicateIntegerGreaterThanOrEqual) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FdcPredicateIntegerGreaterThanOrEqual) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *FdcPredicateIntegerGreaterThanOrEqual) SetValue(v int32) {
	o.Value = &v
}

func (o FdcPredicateIntegerGreaterThanOrEqual) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdcPredicateIntegerGreaterThanOrEqual) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFdcPredicateIntegerGreaterThanOrEqual struct {
	value *FdcPredicateIntegerGreaterThanOrEqual
	isSet bool
}

func (v NullableFdcPredicateIntegerGreaterThanOrEqual) Get() *FdcPredicateIntegerGreaterThanOrEqual {
	return v.value
}

func (v *NullableFdcPredicateIntegerGreaterThanOrEqual) Set(val *FdcPredicateIntegerGreaterThanOrEqual) {
	v.value = val
	v.isSet = true
}

func (v NullableFdcPredicateIntegerGreaterThanOrEqual) IsSet() bool {
	return v.isSet
}

func (v *NullableFdcPredicateIntegerGreaterThanOrEqual) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdcPredicateIntegerGreaterThanOrEqual(val *FdcPredicateIntegerGreaterThanOrEqual) *NullableFdcPredicateIntegerGreaterThanOrEqual {
	return &NullableFdcPredicateIntegerGreaterThanOrEqual{value: val, isSet: true}
}

func (v NullableFdcPredicateIntegerGreaterThanOrEqual) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdcPredicateIntegerGreaterThanOrEqual) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


