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

// checks if the FdcPredicateSetOfIntegersContainsAny type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdcPredicateSetOfIntegersContainsAny{}

// FdcPredicateSetOfIntegersContainsAny The predicate of the `SET_OF_INTEGERS_CONTAINS_ANY` type. It checks whether the attribute (which is a set of integers) contains one of the reference values.
type FdcPredicateSetOfIntegersContainsAny struct {
	// A list of reference values. The condition is fulfilled when the attribute (which is a set of integers) contains *any* of these.
	Values []int32 `json:"values"`
}

// NewFdcPredicateSetOfIntegersContainsAny instantiates a new FdcPredicateSetOfIntegersContainsAny object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdcPredicateSetOfIntegersContainsAny(values []int32, type_ string) *FdcPredicateSetOfIntegersContainsAny {
	this := FdcPredicateSetOfIntegersContainsAny{}
	this.Type = type_
	this.Values = values
	return &this
}

// NewFdcPredicateSetOfIntegersContainsAnyWithDefaults instantiates a new FdcPredicateSetOfIntegersContainsAny object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdcPredicateSetOfIntegersContainsAnyWithDefaults() *FdcPredicateSetOfIntegersContainsAny {
	this := FdcPredicateSetOfIntegersContainsAny{}
	return &this
}

// GetValues returns the Values field value
func (o *FdcPredicateSetOfIntegersContainsAny) GetValues() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *FdcPredicateSetOfIntegersContainsAny) GetValuesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *FdcPredicateSetOfIntegersContainsAny) SetValues(v []int32) {
	o.Values = v
}

func (o FdcPredicateSetOfIntegersContainsAny) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdcPredicateSetOfIntegersContainsAny) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableFdcPredicateSetOfIntegersContainsAny struct {
	value *FdcPredicateSetOfIntegersContainsAny
	isSet bool
}

func (v NullableFdcPredicateSetOfIntegersContainsAny) Get() *FdcPredicateSetOfIntegersContainsAny {
	return v.value
}

func (v *NullableFdcPredicateSetOfIntegersContainsAny) Set(val *FdcPredicateSetOfIntegersContainsAny) {
	v.value = val
	v.isSet = true
}

func (v NullableFdcPredicateSetOfIntegersContainsAny) IsSet() bool {
	return v.isSet
}

func (v *NullableFdcPredicateSetOfIntegersContainsAny) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdcPredicateSetOfIntegersContainsAny(val *FdcPredicateSetOfIntegersContainsAny) *NullableFdcPredicateSetOfIntegersContainsAny {
	return &NullableFdcPredicateSetOfIntegersContainsAny{value: val, isSet: true}
}

func (v NullableFdcPredicateSetOfIntegersContainsAny) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdcPredicateSetOfIntegersContainsAny) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


