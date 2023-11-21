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

// checks if the NotificationConfigStubListDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationConfigStubListDto{}

// NotificationConfigStubListDto struct for NotificationConfigStubListDto
type NotificationConfigStubListDto struct {
	Values []NotificationConfigStub `json:"values,omitempty"`
}

// NewNotificationConfigStubListDto instantiates a new NotificationConfigStubListDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationConfigStubListDto() *NotificationConfigStubListDto {
	this := NotificationConfigStubListDto{}
	return &this
}

// NewNotificationConfigStubListDtoWithDefaults instantiates a new NotificationConfigStubListDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationConfigStubListDtoWithDefaults() *NotificationConfigStubListDto {
	this := NotificationConfigStubListDto{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *NotificationConfigStubListDto) GetValues() []NotificationConfigStub {
	if o == nil || IsNil(o.Values) {
		var ret []NotificationConfigStub
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationConfigStubListDto) GetValuesOk() ([]NotificationConfigStub, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *NotificationConfigStubListDto) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []NotificationConfigStub and assigns it to the Values field.
func (o *NotificationConfigStubListDto) SetValues(v []NotificationConfigStub) {
	o.Values = v
}

func (o NotificationConfigStubListDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationConfigStubListDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableNotificationConfigStubListDto struct {
	value *NotificationConfigStubListDto
	isSet bool
}

func (v NullableNotificationConfigStubListDto) Get() *NotificationConfigStubListDto {
	return v.value
}

func (v *NullableNotificationConfigStubListDto) Set(val *NotificationConfigStubListDto) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationConfigStubListDto) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationConfigStubListDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationConfigStubListDto(val *NotificationConfigStubListDto) *NullableNotificationConfigStubListDto {
	return &NullableNotificationConfigStubListDto{value: val, isSet: true}
}

func (v NullableNotificationConfigStubListDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationConfigStubListDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


