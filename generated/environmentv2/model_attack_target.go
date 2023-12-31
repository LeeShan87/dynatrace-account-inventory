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

// checks if the AttackTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackTarget{}

// AttackTarget Information about the targeted host/database of an attack.
type AttackTarget struct {
	// The monitored entity ID of the targeted host/database.
	EntityId *string `json:"entityId,omitempty"`
	// The name of the targeted host/database.
	Name *string `json:"name,omitempty"`
}

// NewAttackTarget instantiates a new AttackTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackTarget() *AttackTarget {
	this := AttackTarget{}
	return &this
}

// NewAttackTargetWithDefaults instantiates a new AttackTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackTargetWithDefaults() *AttackTarget {
	this := AttackTarget{}
	return &this
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *AttackTarget) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackTarget) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *AttackTarget) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *AttackTarget) SetEntityId(v string) {
	o.EntityId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttackTarget) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackTarget) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttackTarget) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttackTarget) SetName(v string) {
	o.Name = &v
}

func (o AttackTarget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAttackTarget struct {
	value *AttackTarget
	isSet bool
}

func (v NullableAttackTarget) Get() *AttackTarget {
	return v.value
}

func (v *NullableAttackTarget) Set(val *AttackTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackTarget(val *AttackTarget) *NullableAttackTarget {
	return &NullableAttackTarget{value: val, isSet: true}
}

func (v NullableAttackTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


