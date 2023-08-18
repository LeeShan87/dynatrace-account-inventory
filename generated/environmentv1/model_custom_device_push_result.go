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

// checks if the CustomDevicePushResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDevicePushResult{}

// CustomDevicePushResult The result of a custom device push request. The entity ID is calculated automatically.
type CustomDevicePushResult struct {
	// The Dynatrace entity ID of the custom device.
	EntityId *string `json:"entityId,omitempty"`
	// The Dynatrace entity ID of the custom device group.
	GroupId *string `json:"groupId,omitempty"`
}

// NewCustomDevicePushResult instantiates a new CustomDevicePushResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDevicePushResult() *CustomDevicePushResult {
	this := CustomDevicePushResult{}
	return &this
}

// NewCustomDevicePushResultWithDefaults instantiates a new CustomDevicePushResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDevicePushResultWithDefaults() *CustomDevicePushResult {
	this := CustomDevicePushResult{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *CustomDevicePushResult) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushResult) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *CustomDevicePushResult) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *CustomDevicePushResult) SetEntityId(v string) {
	o.EntityId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CustomDevicePushResult) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDevicePushResult) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CustomDevicePushResult) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CustomDevicePushResult) SetGroupId(v string) {
	o.GroupId = &v
}

func (o CustomDevicePushResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDevicePushResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	return toSerialize, nil
}

type NullableCustomDevicePushResult struct {
	value *CustomDevicePushResult
	isSet bool
}

func (v NullableCustomDevicePushResult) Get() *CustomDevicePushResult {
	return v.value
}

func (v *NullableCustomDevicePushResult) Set(val *CustomDevicePushResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDevicePushResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDevicePushResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDevicePushResult(val *CustomDevicePushResult) *NullableCustomDevicePushResult {
	return &NullableCustomDevicePushResult{value: val, isSet: true}
}

func (v NullableCustomDevicePushResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDevicePushResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

