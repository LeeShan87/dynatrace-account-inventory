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

// checks if the DexpFilterRelationship type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DexpFilterRelationship{}

// DexpFilterRelationship struct for DexpFilterRelationship
type DexpFilterRelationship struct {
	// The id of the relationship. e.g runsOn, isStepOf, etc
	Id *string `json:"id,omitempty"`
	// The target entity of the relationship. e.g HOST, VCENTER, SERVICE etc
	TargetEntity *string `json:"targetEntity,omitempty"`
	// The type of the relationship
	Type *string `json:"type,omitempty"`
}

// NewDexpFilterRelationship instantiates a new DexpFilterRelationship object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDexpFilterRelationship() *DexpFilterRelationship {
	this := DexpFilterRelationship{}
	return &this
}

// NewDexpFilterRelationshipWithDefaults instantiates a new DexpFilterRelationship object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDexpFilterRelationshipWithDefaults() *DexpFilterRelationship {
	this := DexpFilterRelationship{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DexpFilterRelationship) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DexpFilterRelationship) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DexpFilterRelationship) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DexpFilterRelationship) SetId(v string) {
	o.Id = &v
}

// GetTargetEntity returns the TargetEntity field value if set, zero value otherwise.
func (o *DexpFilterRelationship) GetTargetEntity() string {
	if o == nil || IsNil(o.TargetEntity) {
		var ret string
		return ret
	}
	return *o.TargetEntity
}

// GetTargetEntityOk returns a tuple with the TargetEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DexpFilterRelationship) GetTargetEntityOk() (*string, bool) {
	if o == nil || IsNil(o.TargetEntity) {
		return nil, false
	}
	return o.TargetEntity, true
}

// HasTargetEntity returns a boolean if a field has been set.
func (o *DexpFilterRelationship) HasTargetEntity() bool {
	if o != nil && !IsNil(o.TargetEntity) {
		return true
	}

	return false
}

// SetTargetEntity gets a reference to the given string and assigns it to the TargetEntity field.
func (o *DexpFilterRelationship) SetTargetEntity(v string) {
	o.TargetEntity = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DexpFilterRelationship) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DexpFilterRelationship) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DexpFilterRelationship) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DexpFilterRelationship) SetType(v string) {
	o.Type = &v
}

func (o DexpFilterRelationship) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DexpFilterRelationship) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TargetEntity) {
		toSerialize["targetEntity"] = o.TargetEntity
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDexpFilterRelationship struct {
	value *DexpFilterRelationship
	isSet bool
}

func (v NullableDexpFilterRelationship) Get() *DexpFilterRelationship {
	return v.value
}

func (v *NullableDexpFilterRelationship) Set(val *DexpFilterRelationship) {
	v.value = val
	v.isSet = true
}

func (v NullableDexpFilterRelationship) IsSet() bool {
	return v.isSet
}

func (v *NullableDexpFilterRelationship) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDexpFilterRelationship(val *DexpFilterRelationship) *NullableDexpFilterRelationship {
	return &NullableDexpFilterRelationship{value: val, isSet: true}
}

func (v NullableDexpFilterRelationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDexpFilterRelationship) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


