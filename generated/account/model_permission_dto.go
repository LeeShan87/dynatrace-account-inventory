/*
Dynatrace Account Management API

The enterprise management API for Dynatrace SaaS enables automation of operational tasks related to user access and environment lifecycle management.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package account

import (
	"encoding/json"
)

// checks if the PermissionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionDto{}

// PermissionDto struct for PermissionDto
type PermissionDto struct {
	// The ID of the permission.
	Id string `json:"id"`
	// The display name of the permission.
	Description string `json:"description"`
}

// NewPermissionDto instantiates a new PermissionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionDto(id string, description string) *PermissionDto {
	this := PermissionDto{}
	this.Id = id
	this.Description = description
	return &this
}

// NewPermissionDtoWithDefaults instantiates a new PermissionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionDtoWithDefaults() *PermissionDto {
	this := PermissionDto{}
	return &this
}

// GetId returns the Id field value
func (o *PermissionDto) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PermissionDto) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PermissionDto) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *PermissionDto) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PermissionDto) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PermissionDto) SetDescription(v string) {
	o.Description = v
}

func (o PermissionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["description"] = o.Description
	return toSerialize, nil
}

type NullablePermissionDto struct {
	value *PermissionDto
	isSet bool
}

func (v NullablePermissionDto) Get() *PermissionDto {
	return v.value
}

func (v *NullablePermissionDto) Set(val *PermissionDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionDto(val *PermissionDto) *NullablePermissionDto {
	return &NullablePermissionDto{value: val, isSet: true}
}

func (v NullablePermissionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


