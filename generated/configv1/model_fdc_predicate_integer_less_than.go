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

// checks if the FdcPredicateIntegerLessThan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdcPredicateIntegerLessThan{}

// FdcPredicateIntegerLessThan The predicate of the `INTEGER_LESS_THAN` type. It checks whether the attribute (which is an integer) is less than the reference value.
type FdcPredicateIntegerLessThan struct {
	// The reference value. The condition is fulfilled when the attribute (which is an integer) is less than this value.
	Value *int32 `json:"value,omitempty"`
}

// NewFdcPredicateIntegerLessThan instantiates a new FdcPredicateIntegerLessThan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdcPredicateIntegerLessThan(type_ string) *FdcPredicateIntegerLessThan {
	this := FdcPredicateIntegerLessThan{}
	this.Type = type_
	return &this
}

// NewFdcPredicateIntegerLessThanWithDefaults instantiates a new FdcPredicateIntegerLessThan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdcPredicateIntegerLessThanWithDefaults() *FdcPredicateIntegerLessThan {
	this := FdcPredicateIntegerLessThan{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FdcPredicateIntegerLessThan) GetValue() int32 {
	if o == nil || IsNil(o.Value) {
		var ret int32
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FdcPredicateIntegerLessThan) GetValueOk() (*int32, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FdcPredicateIntegerLessThan) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given int32 and assigns it to the Value field.
func (o *FdcPredicateIntegerLessThan) SetValue(v int32) {
	o.Value = &v
}

func (o FdcPredicateIntegerLessThan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdcPredicateIntegerLessThan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableFdcPredicateIntegerLessThan struct {
	value *FdcPredicateIntegerLessThan
	isSet bool
}

func (v NullableFdcPredicateIntegerLessThan) Get() *FdcPredicateIntegerLessThan {
	return v.value
}

func (v *NullableFdcPredicateIntegerLessThan) Set(val *FdcPredicateIntegerLessThan) {
	v.value = val
	v.isSet = true
}

func (v NullableFdcPredicateIntegerLessThan) IsSet() bool {
	return v.isSet
}

func (v *NullableFdcPredicateIntegerLessThan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdcPredicateIntegerLessThan(val *FdcPredicateIntegerLessThan) *NullableFdcPredicateIntegerLessThan {
	return &NullableFdcPredicateIntegerLessThan{value: val, isSet: true}
}

func (v NullableFdcPredicateIntegerLessThan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdcPredicateIntegerLessThan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


