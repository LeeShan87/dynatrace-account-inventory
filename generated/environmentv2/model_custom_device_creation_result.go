/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
)

// checks if the CustomDeviceCreationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomDeviceCreationResult{}

// CustomDeviceCreationResult The short representation of a newly created custom device.
type CustomDeviceCreationResult struct {
	// The Dynatrace entity ID of the custom device.
	EntityId *string `json:"entityId,omitempty"`
	// The Dynatrace entity ID of the custom device group.
	GroupId *string `json:"groupId,omitempty"`
}

// NewCustomDeviceCreationResult instantiates a new CustomDeviceCreationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomDeviceCreationResult() *CustomDeviceCreationResult {
	this := CustomDeviceCreationResult{}
	return &this
}

// NewCustomDeviceCreationResultWithDefaults instantiates a new CustomDeviceCreationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomDeviceCreationResultWithDefaults() *CustomDeviceCreationResult {
	this := CustomDeviceCreationResult{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *CustomDeviceCreationResult) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreationResult) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *CustomDeviceCreationResult) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *CustomDeviceCreationResult) SetEntityId(v string) {
	o.EntityId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *CustomDeviceCreationResult) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomDeviceCreationResult) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *CustomDeviceCreationResult) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *CustomDeviceCreationResult) SetGroupId(v string) {
	o.GroupId = &v
}

func (o CustomDeviceCreationResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomDeviceCreationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	return toSerialize, nil
}

type NullableCustomDeviceCreationResult struct {
	value *CustomDeviceCreationResult
	isSet bool
}

func (v NullableCustomDeviceCreationResult) Get() *CustomDeviceCreationResult {
	return v.value
}

func (v *NullableCustomDeviceCreationResult) Set(val *CustomDeviceCreationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomDeviceCreationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomDeviceCreationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomDeviceCreationResult(val *CustomDeviceCreationResult) *NullableCustomDeviceCreationResult {
	return &NullableCustomDeviceCreationResult{value: val, isSet: true}
}

func (v NullableCustomDeviceCreationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomDeviceCreationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


