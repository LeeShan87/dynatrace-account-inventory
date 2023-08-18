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

// checks if the FdcPredicateIntegerEquals type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdcPredicateIntegerEquals{}

// FdcPredicateIntegerEquals The predicate of the `INTEGER_EQUALS` type. It checks whether the attribute (which is an integer) equals one of the reference values.
type FdcPredicateIntegerEquals struct {
	// The reference value. The condition is fulfilled when the attribute (which is an integer) equals *any* of these.
	Values []int32 `json:"values"`
}

// NewFdcPredicateIntegerEquals instantiates a new FdcPredicateIntegerEquals object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdcPredicateIntegerEquals(values []int32, type_ string) *FdcPredicateIntegerEquals {
	this := FdcPredicateIntegerEquals{}
	this.Type = type_
	this.Values = values
	return &this
}

// NewFdcPredicateIntegerEqualsWithDefaults instantiates a new FdcPredicateIntegerEquals object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdcPredicateIntegerEqualsWithDefaults() *FdcPredicateIntegerEquals {
	this := FdcPredicateIntegerEquals{}
	return &this
}

// GetValues returns the Values field value
func (o *FdcPredicateIntegerEquals) GetValues() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *FdcPredicateIntegerEquals) GetValuesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *FdcPredicateIntegerEquals) SetValues(v []int32) {
	o.Values = v
}

func (o FdcPredicateIntegerEquals) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdcPredicateIntegerEquals) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableFdcPredicateIntegerEquals struct {
	value *FdcPredicateIntegerEquals
	isSet bool
}

func (v NullableFdcPredicateIntegerEquals) Get() *FdcPredicateIntegerEquals {
	return v.value
}

func (v *NullableFdcPredicateIntegerEquals) Set(val *FdcPredicateIntegerEquals) {
	v.value = val
	v.isSet = true
}

func (v NullableFdcPredicateIntegerEquals) IsSet() bool {
	return v.isSet
}

func (v *NullableFdcPredicateIntegerEquals) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdcPredicateIntegerEquals(val *FdcPredicateIntegerEquals) *NullableFdcPredicateIntegerEquals {
	return &NullableFdcPredicateIntegerEquals{value: val, isSet: true}
}

func (v NullableFdcPredicateIntegerEquals) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdcPredicateIntegerEquals) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

