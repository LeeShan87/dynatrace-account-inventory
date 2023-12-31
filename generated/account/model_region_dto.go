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

// checks if the RegionDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegionDto{}

// RegionDto struct for RegionDto
type RegionDto struct {
	// The name of the region.
	Name string `json:"name"`
}

// NewRegionDto instantiates a new RegionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegionDto(name string) *RegionDto {
	this := RegionDto{}
	this.Name = name
	return &this
}

// NewRegionDtoWithDefaults instantiates a new RegionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegionDtoWithDefaults() *RegionDto {
	this := RegionDto{}
	return &this
}

// GetName returns the Name field value
func (o *RegionDto) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RegionDto) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RegionDto) SetName(v string) {
	o.Name = v
}

func (o RegionDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegionDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableRegionDto struct {
	value *RegionDto
	isSet bool
}

func (v NullableRegionDto) Get() *RegionDto {
	return v.value
}

func (v *NullableRegionDto) Set(val *RegionDto) {
	v.value = val
	v.isSet = true
}

func (v NullableRegionDto) IsSet() bool {
	return v.isSet
}

func (v *NullableRegionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegionDto(val *RegionDto) *NullableRegionDto {
	return &NullableRegionDto{value: val, isSet: true}
}

func (v NullableRegionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


