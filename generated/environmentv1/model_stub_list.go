/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv1

import (
	"encoding/json"
)

// checks if the StubList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StubList{}

// StubList An ordered list of short representations of Dynatrace entities.
type StubList struct {
	// An ordered list of short representations of Dynatrace entities.
	Values []EntityShortRepresentation `json:"values"`
}

// NewStubList instantiates a new StubList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStubList(values []EntityShortRepresentation) *StubList {
	this := StubList{}
	this.Values = values
	return &this
}

// NewStubListWithDefaults instantiates a new StubList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStubListWithDefaults() *StubList {
	this := StubList{}
	return &this
}

// GetValues returns the Values field value
func (o *StubList) GetValues() []EntityShortRepresentation {
	if o == nil {
		var ret []EntityShortRepresentation
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *StubList) GetValuesOk() ([]EntityShortRepresentation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *StubList) SetValues(v []EntityShortRepresentation) {
	o.Values = v
}

func (o StubList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StubList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableStubList struct {
	value *StubList
	isSet bool
}

func (v NullableStubList) Get() *StubList {
	return v.value
}

func (v *NullableStubList) Set(val *StubList) {
	v.value = val
	v.isSet = true
}

func (v NullableStubList) IsSet() bool {
	return v.isSet
}

func (v *NullableStubList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStubList(val *StubList) *NullableStubList {
	return &NullableStubList{value: val, isSet: true}
}

func (v NullableStubList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStubList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


