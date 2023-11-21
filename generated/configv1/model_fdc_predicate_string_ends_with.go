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

// checks if the FdcPredicateStringEndsWith type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FdcPredicateStringEndsWith{}

// FdcPredicateStringEndsWith The predicate of the `STRING_ENDS_WITH ` type. It checks whether the attribute (which is a string) ends with one of the reference values.
type FdcPredicateStringEndsWith struct {
	// The condition is case sensitive (`false`) or case insensitive (`true`).   If not set, then `false` is used, making the condition case sensitive.
	IgnoreCase bool `json:"ignoreCase"`
	// A list of reference values. The condition is fulfilled when the attribute (which is a string) ends with *any* of these.
	Values []string `json:"values"`
}

// NewFdcPredicateStringEndsWith instantiates a new FdcPredicateStringEndsWith object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFdcPredicateStringEndsWith(ignoreCase bool, values []string, type_ string) *FdcPredicateStringEndsWith {
	this := FdcPredicateStringEndsWith{}
	this.Type = type_
	this.IgnoreCase = ignoreCase
	this.Values = values
	return &this
}

// NewFdcPredicateStringEndsWithWithDefaults instantiates a new FdcPredicateStringEndsWith object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFdcPredicateStringEndsWithWithDefaults() *FdcPredicateStringEndsWith {
	this := FdcPredicateStringEndsWith{}
	return &this
}

// GetIgnoreCase returns the IgnoreCase field value
func (o *FdcPredicateStringEndsWith) GetIgnoreCase() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgnoreCase
}

// GetIgnoreCaseOk returns a tuple with the IgnoreCase field value
// and a boolean to check if the value has been set.
func (o *FdcPredicateStringEndsWith) GetIgnoreCaseOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnoreCase, true
}

// SetIgnoreCase sets field value
func (o *FdcPredicateStringEndsWith) SetIgnoreCase(v bool) {
	o.IgnoreCase = v
}

// GetValues returns the Values field value
func (o *FdcPredicateStringEndsWith) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *FdcPredicateStringEndsWith) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *FdcPredicateStringEndsWith) SetValues(v []string) {
	o.Values = v
}

func (o FdcPredicateStringEndsWith) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FdcPredicateStringEndsWith) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ignoreCase"] = o.IgnoreCase
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableFdcPredicateStringEndsWith struct {
	value *FdcPredicateStringEndsWith
	isSet bool
}

func (v NullableFdcPredicateStringEndsWith) Get() *FdcPredicateStringEndsWith {
	return v.value
}

func (v *NullableFdcPredicateStringEndsWith) Set(val *FdcPredicateStringEndsWith) {
	v.value = val
	v.isSet = true
}

func (v NullableFdcPredicateStringEndsWith) IsSet() bool {
	return v.isSet
}

func (v *NullableFdcPredicateStringEndsWith) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFdcPredicateStringEndsWith(val *FdcPredicateStringEndsWith) *NullableFdcPredicateStringEndsWith {
	return &NullableFdcPredicateStringEndsWith{value: val, isSet: true}
}

func (v NullableFdcPredicateStringEndsWith) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFdcPredicateStringEndsWith) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


