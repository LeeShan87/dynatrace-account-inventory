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

// checks if the LessThanCompareOperation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LessThanCompareOperation{}

// LessThanCompareOperation The condition of the `LESS_THAN` type.   The condition checks whether the integer value is less than the specified value.
type LessThanCompareOperation struct {
	// The value to compare to.
	Value int32 `json:"value"`
}

// NewLessThanCompareOperation instantiates a new LessThanCompareOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLessThanCompareOperation(value int32, type_ string) *LessThanCompareOperation {
	this := LessThanCompareOperation{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewLessThanCompareOperationWithDefaults instantiates a new LessThanCompareOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLessThanCompareOperationWithDefaults() *LessThanCompareOperation {
	this := LessThanCompareOperation{}
	return &this
}

// GetValue returns the Value field value
func (o *LessThanCompareOperation) GetValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LessThanCompareOperation) GetValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *LessThanCompareOperation) SetValue(v int32) {
	o.Value = v
}

func (o LessThanCompareOperation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LessThanCompareOperation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableLessThanCompareOperation struct {
	value *LessThanCompareOperation
	isSet bool
}

func (v NullableLessThanCompareOperation) Get() *LessThanCompareOperation {
	return v.value
}

func (v *NullableLessThanCompareOperation) Set(val *LessThanCompareOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableLessThanCompareOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableLessThanCompareOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLessThanCompareOperation(val *LessThanCompareOperation) *NullableLessThanCompareOperation {
	return &NullableLessThanCompareOperation{value: val, isSet: true}
}

func (v NullableLessThanCompareOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLessThanCompareOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

