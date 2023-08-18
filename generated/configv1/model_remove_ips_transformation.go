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

// checks if the RemoveIPsTransformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveIPsTransformation{}

// RemoveIPsTransformation The transformation of the `REMOVE_IPS` type.   The transformation automatically detects and removes IP addresses. No additional parameters needed.
type RemoveIPsTransformation struct {
}

// NewRemoveIPsTransformation instantiates a new RemoveIPsTransformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveIPsTransformation(type_ string) *RemoveIPsTransformation {
	this := RemoveIPsTransformation{}
	this.Type = type_
	return &this
}

// NewRemoveIPsTransformationWithDefaults instantiates a new RemoveIPsTransformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveIPsTransformationWithDefaults() *RemoveIPsTransformation {
	this := RemoveIPsTransformation{}
	return &this
}

func (o RemoveIPsTransformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveIPsTransformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableRemoveIPsTransformation struct {
	value *RemoveIPsTransformation
	isSet bool
}

func (v NullableRemoveIPsTransformation) Get() *RemoveIPsTransformation {
	return v.value
}

func (v *NullableRemoveIPsTransformation) Set(val *RemoveIPsTransformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveIPsTransformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveIPsTransformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveIPsTransformation(val *RemoveIPsTransformation) *NullableRemoveIPsTransformation {
	return &NullableRemoveIPsTransformation{value: val, isSet: true}
}

func (v NullableRemoveIPsTransformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveIPsTransformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


