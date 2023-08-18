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

// checks if the NotificationConfigStub type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationConfigStub{}

// NotificationConfigStub The short representation of a notification.
type NotificationConfigStub struct {
	// A short description of the Dynatrace entity.
	Description *string `json:"description,omitempty"`
	// The ID of the Dynatrace entity.
	Id string `json:"id"`
	// The name of the Dynatrace entity.
	Name *string `json:"name,omitempty"`
	// The type of the notification.
	Type *string `json:"type,omitempty"`
}

// NewNotificationConfigStub instantiates a new NotificationConfigStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationConfigStub(id string) *NotificationConfigStub {
	this := NotificationConfigStub{}
	this.Id = id
	return &this
}

// NewNotificationConfigStubWithDefaults instantiates a new NotificationConfigStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationConfigStubWithDefaults() *NotificationConfigStub {
	this := NotificationConfigStub{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NotificationConfigStub) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationConfigStub) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *NotificationConfigStub) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NotificationConfigStub) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value
func (o *NotificationConfigStub) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NotificationConfigStub) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NotificationConfigStub) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NotificationConfigStub) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationConfigStub) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NotificationConfigStub) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NotificationConfigStub) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *NotificationConfigStub) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationConfigStub) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *NotificationConfigStub) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *NotificationConfigStub) SetType(v string) {
	o.Type = &v
}

func (o NotificationConfigStub) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationConfigStub) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableNotificationConfigStub struct {
	value *NotificationConfigStub
	isSet bool
}

func (v NullableNotificationConfigStub) Get() *NotificationConfigStub {
	return v.value
}

func (v *NullableNotificationConfigStub) Set(val *NotificationConfigStub) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationConfigStub) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationConfigStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationConfigStub(val *NotificationConfigStub) *NullableNotificationConfigStub {
	return &NullableNotificationConfigStub{value: val, isSet: true}
}

func (v NullableNotificationConfigStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationConfigStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

