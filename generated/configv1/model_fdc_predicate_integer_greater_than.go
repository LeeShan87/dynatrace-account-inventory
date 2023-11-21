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

// checks if the FdcPredicateIntegerGreaterThan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdcPredicateIntegerGreaterThan{}

// FdcPredicateIntegerGreaterThan The predicate of the `INTEGER_GREATER_THAN` type. It checks whether the attribute (which is an integer) is greater than the reference value.
type FdcPredicateIntegerGreaterThan struct {
	// The reference value. The condition is fulfilled when the attribute (which is an integer) is greater than this value.
	Value *int32 `json:"value,omitempty"`
}

// NewFdcPredicateIntegerGreaterThan instantiates a new FdcPredicateIntegerGreaterThan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdcPredicateIntegerGreaterThan(type_ string) *FdcPredicateIntegerGreaterThan {
	this := FdcPredicateIntegerGreaterThan{}
	this.Type = type_
	return &this
}

// NewFdcPredicateIntegerGreaterThanWithDefaults instantiates a new FdcPredicateIntegerGreaterThan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdcPredicateIntegerGreaterThanWithDefaults() *FdcPredicateIntegerGreaterThan {
	this := FdcPredicateIntegerGreaterThan{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FdcPredicateIntegerGreaterThan) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FdcPredicateIntegerGreaterThan) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FdcPredicateIntegerGreaterThan) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *FdcPredicateIntegerGreaterThan) SetValue(v int32) {
	o.Value = &v
}

func (o FdcPredicateIntegerGreaterThan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdcPredicateIntegerGreaterThan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFdcPredicateIntegerGreaterThan struct {
	value *FdcPredicateIntegerGreaterThan
	isSet bool
}

func (v NullableFdcPredicateIntegerGreaterThan) Get() *FdcPredicateIntegerGreaterThan {
	return v.value
}

func (v *NullableFdcPredicateIntegerGreaterThan) Set(val *FdcPredicateIntegerGreaterThan) {
	v.value = val
	v.isSet = true
}

func (v NullableFdcPredicateIntegerGreaterThan) IsSet() bool {
	return v.isSet
}

func (v *NullableFdcPredicateIntegerGreaterThan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdcPredicateIntegerGreaterThan(val *FdcPredicateIntegerGreaterThan) *NullableFdcPredicateIntegerGreaterThan {
	return &NullableFdcPredicateIntegerGreaterThan{value: val, isSet: true}
}

func (v NullableFdcPredicateIntegerGreaterThan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdcPredicateIntegerGreaterThan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


