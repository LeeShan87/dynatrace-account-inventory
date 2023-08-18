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

// checks if the EntityIdAlertingScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntityIdAlertingScope{}

// EntityIdAlertingScope A scope filter for a monitored entity identifier.
type EntityIdAlertingScope struct {
	// The monitored entities id to match on.
	EntityId string `json:"entityId"`
}

// NewEntityIdAlertingScope instantiates a new EntityIdAlertingScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntityIdAlertingScope(entityId string, filterType string) *EntityIdAlertingScope {
	this := EntityIdAlertingScope{}
	this.FilterType = filterType
	this.EntityId = entityId
	return &this
}

// NewEntityIdAlertingScopeWithDefaults instantiates a new EntityIdAlertingScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntityIdAlertingScopeWithDefaults() *EntityIdAlertingScope {
	this := EntityIdAlertingScope{}
	return &this
}

// GetEntityId returns the EntityId field value
func (o *EntityIdAlertingScope) GetEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value
// and a boolean to check if the value has been set.
func (o *EntityIdAlertingScope) GetEntityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntityId, true
}

// SetEntityId sets field value
func (o *EntityIdAlertingScope) SetEntityId(v string) {
	o.EntityId = v
}

func (o EntityIdAlertingScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntityIdAlertingScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entityId"] = o.EntityId
	return toSerialize, nil
}

type NullableEntityIdAlertingScope struct {
	value *EntityIdAlertingScope
	isSet bool
}

func (v NullableEntityIdAlertingScope) Get() *EntityIdAlertingScope {
	return v.value
}

func (v *NullableEntityIdAlertingScope) Set(val *EntityIdAlertingScope) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityIdAlertingScope) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityIdAlertingScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityIdAlertingScope(val *EntityIdAlertingScope) *NullableEntityIdAlertingScope {
	return &NullableEntityIdAlertingScope{value: val, isSet: true}
}

func (v NullableEntityIdAlertingScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityIdAlertingScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


