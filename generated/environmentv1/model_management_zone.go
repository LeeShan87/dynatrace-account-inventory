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

// checks if the ManagementZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementZone{}

// ManagementZone The configuration of a management zone.
type ManagementZone struct {
	// The ID of the management zone.
	Id string `json:"id"`
	// The name of the management zone.
	Name string `json:"name"`
}

// NewManagementZone instantiates a new ManagementZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementZone(id string, name string) *ManagementZone {
	this := ManagementZone{}
	this.Id = id
	this.Name = name
	return &this
}

// NewManagementZoneWithDefaults instantiates a new ManagementZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementZoneWithDefaults() *ManagementZone {
	this := ManagementZone{}
	return &this
}

// GetId returns the Id field value
func (o *ManagementZone) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ManagementZone) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ManagementZone) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ManagementZone) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ManagementZone) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ManagementZone) SetName(v string) {
	o.Name = v
}

func (o ManagementZone) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableManagementZone struct {
	value *ManagementZone
	isSet bool
}

func (v NullableManagementZone) Get() *ManagementZone {
	return v.value
}

func (v *NullableManagementZone) Set(val *ManagementZone) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementZone) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementZone(val *ManagementZone) *NullableManagementZone {
	return &NullableManagementZone{value: val, isSet: true}
}

func (v NullableManagementZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


