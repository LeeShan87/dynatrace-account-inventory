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

// checks if the KeyUserActionMobileList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyUserActionMobileList{}

// KeyUserActionMobileList A list of key actions in an application.
type KeyUserActionMobileList struct {
	// A list of key actions in an application.
	KeyUserActions []KeyUserActionMobile `json:"keyUserActions,omitempty"`
}

// NewKeyUserActionMobileList instantiates a new KeyUserActionMobileList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyUserActionMobileList() *KeyUserActionMobileList {
	this := KeyUserActionMobileList{}
	return &this
}

// NewKeyUserActionMobileListWithDefaults instantiates a new KeyUserActionMobileList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyUserActionMobileListWithDefaults() *KeyUserActionMobileList {
	this := KeyUserActionMobileList{}
	return &this
}

// GetKeyUserActions returns the KeyUserActions field value if set, zero value otherwise.
func (o *KeyUserActionMobileList) GetKeyUserActions() []KeyUserActionMobile {
	if o == nil || IsNil(o.KeyUserActions) {
		var ret []KeyUserActionMobile
		return ret
	}
	return o.KeyUserActions
}

// GetKeyUserActionsOk returns a tuple with the KeyUserActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyUserActionMobileList) GetKeyUserActionsOk() ([]KeyUserActionMobile, bool) {
	if o == nil || IsNil(o.KeyUserActions) {
		return nil, false
	}
	return o.KeyUserActions, true
}

// HasKeyUserActions returns a boolean if a field has been set.
func (o *KeyUserActionMobileList) HasKeyUserActions() bool {
	if o != nil && !IsNil(o.KeyUserActions) {
		return true
	}

	return false
}

// SetKeyUserActions gets a reference to the given []KeyUserActionMobile and assigns it to the KeyUserActions field.
func (o *KeyUserActionMobileList) SetKeyUserActions(v []KeyUserActionMobile) {
	o.KeyUserActions = v
}

func (o KeyUserActionMobileList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyUserActionMobileList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeyUserActions) {
		toSerialize["keyUserActions"] = o.KeyUserActions
	}
	return toSerialize, nil
}

type NullableKeyUserActionMobileList struct {
	value *KeyUserActionMobileList
	isSet bool
}

func (v NullableKeyUserActionMobileList) Get() *KeyUserActionMobileList {
	return v.value
}

func (v *NullableKeyUserActionMobileList) Set(val *KeyUserActionMobileList) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyUserActionMobileList) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyUserActionMobileList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyUserActionMobileList(val *KeyUserActionMobileList) *NullableKeyUserActionMobileList {
	return &NullableKeyUserActionMobileList{value: val, isSet: true}
}

func (v NullableKeyUserActionMobileList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyUserActionMobileList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

