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

// checks if the StringEqualsCompareOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StringEqualsCompareOperation{}

// StringEqualsCompareOperation The condition of the `STRING_EQUALS` type.   The condition checks whether the string value equals the specified text.
type StringEqualsCompareOperation struct {
	// The condition is case sensitive (`false`) or case insensitive (`true`).   If not set, then `false` is used, making the condition case sensitive.
	IgnoreCase *bool `json:"ignoreCase,omitempty"`
	// Inverts the operation of the condition. Set to `true` to turn **equals** into **does not equal**.    If not set, then `false` is used.
	Negate *bool `json:"negate,omitempty"`
	// The value to compare to.   If several values are specified, the OR logic applies.
	Values []string `json:"values"`
}

// NewStringEqualsCompareOperation instantiates a new StringEqualsCompareOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStringEqualsCompareOperation(values []string, type_ string) *StringEqualsCompareOperation {
	this := StringEqualsCompareOperation{}
	this.Type = type_
	this.Values = values
	return &this
}

// NewStringEqualsCompareOperationWithDefaults instantiates a new StringEqualsCompareOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStringEqualsCompareOperationWithDefaults() *StringEqualsCompareOperation {
	this := StringEqualsCompareOperation{}
	return &this
}

// GetIgnoreCase returns the IgnoreCase field value if set, zero value otherwise.
func (o *StringEqualsCompareOperation) GetIgnoreCase() bool {
	if o == nil || IsNil(o.IgnoreCase) {
		var ret bool
		return ret
	}
	return *o.IgnoreCase
}

// GetIgnoreCaseOk returns a tuple with the IgnoreCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringEqualsCompareOperation) GetIgnoreCaseOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreCase) {
		return nil, false
	}
	return o.IgnoreCase, true
}

// HasIgnoreCase returns a boolean if a field has been set.
func (o *StringEqualsCompareOperation) HasIgnoreCase() bool {
	if o != nil && !IsNil(o.IgnoreCase) {
		return true
	}

	return false
}

// SetIgnoreCase gets a reference to the given bool and assigns it to the IgnoreCase field.
func (o *StringEqualsCompareOperation) SetIgnoreCase(v bool) {
	o.IgnoreCase = &v
}

// GetNegate returns the Negate field value if set, zero value otherwise.
func (o *StringEqualsCompareOperation) GetNegate() bool {
	if o == nil || IsNil(o.Negate) {
		var ret bool
		return ret
	}
	return *o.Negate
}

// GetNegateOk returns a tuple with the Negate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StringEqualsCompareOperation) GetNegateOk() (*bool, bool) {
	if o == nil || IsNil(o.Negate) {
		return nil, false
	}
	return o.Negate, true
}

// HasNegate returns a boolean if a field has been set.
func (o *StringEqualsCompareOperation) HasNegate() bool {
	if o != nil && !IsNil(o.Negate) {
		return true
	}

	return false
}

// SetNegate gets a reference to the given bool and assigns it to the Negate field.
func (o *StringEqualsCompareOperation) SetNegate(v bool) {
	o.Negate = &v
}

// GetValues returns the Values field value
func (o *StringEqualsCompareOperation) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *StringEqualsCompareOperation) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *StringEqualsCompareOperation) SetValues(v []string) {
	o.Values = v
}

func (o StringEqualsCompareOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StringEqualsCompareOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IgnoreCase) {
		toSerialize["ignoreCase"] = o.IgnoreCase
	}
	if !IsNil(o.Negate) {
		toSerialize["negate"] = o.Negate
	}
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

type NullableStringEqualsCompareOperation struct {
	value *StringEqualsCompareOperation
	isSet bool
}

func (v NullableStringEqualsCompareOperation) Get() *StringEqualsCompareOperation {
	return v.value
}

func (v *NullableStringEqualsCompareOperation) Set(val *StringEqualsCompareOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableStringEqualsCompareOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableStringEqualsCompareOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStringEqualsCompareOperation(val *StringEqualsCompareOperation) *NullableStringEqualsCompareOperation {
	return &NullableStringEqualsCompareOperation{value: val, isSet: true}
}

func (v NullableStringEqualsCompareOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStringEqualsCompareOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

