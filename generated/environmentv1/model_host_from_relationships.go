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

// checks if the HostFromRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostFromRelationships{}

// HostFromRelationships struct for HostFromRelationships
type HostFromRelationships struct {
	IsNetworkClientOfHost []string `json:"isNetworkClientOfHost,omitempty"`
}

// NewHostFromRelationships instantiates a new HostFromRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostFromRelationships() *HostFromRelationships {
	this := HostFromRelationships{}
	return &this
}

// NewHostFromRelationshipsWithDefaults instantiates a new HostFromRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostFromRelationshipsWithDefaults() *HostFromRelationships {
	this := HostFromRelationships{}
	return &this
}

// GetIsNetworkClientOfHost returns the IsNetworkClientOfHost field value if set, zero value otherwise.
func (o *HostFromRelationships) GetIsNetworkClientOfHost() []string {
	if o == nil || IsNil(o.IsNetworkClientOfHost) {
		var ret []string
		return ret
	}
	return o.IsNetworkClientOfHost
}

// GetIsNetworkClientOfHostOk returns a tuple with the IsNetworkClientOfHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostFromRelationships) GetIsNetworkClientOfHostOk() ([]string, bool) {
	if o == nil || IsNil(o.IsNetworkClientOfHost) {
		return nil, false
	}
	return o.IsNetworkClientOfHost, true
}

// HasIsNetworkClientOfHost returns a boolean if a field has been set.
func (o *HostFromRelationships) HasIsNetworkClientOfHost() bool {
	if o != nil && !IsNil(o.IsNetworkClientOfHost) {
		return true
	}

	return false
}

// SetIsNetworkClientOfHost gets a reference to the given []string and assigns it to the IsNetworkClientOfHost field.
func (o *HostFromRelationships) SetIsNetworkClientOfHost(v []string) {
	o.IsNetworkClientOfHost = v
}

func (o HostFromRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostFromRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsNetworkClientOfHost) {
		toSerialize["isNetworkClientOfHost"] = o.IsNetworkClientOfHost
	}
	return toSerialize, nil
}

type NullableHostFromRelationships struct {
	value *HostFromRelationships
	isSet bool
}

func (v NullableHostFromRelationships) Get() *HostFromRelationships {
	return v.value
}

func (v *NullableHostFromRelationships) Set(val *HostFromRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableHostFromRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableHostFromRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostFromRelationships(val *HostFromRelationships) *NullableHostFromRelationships {
	return &NullableHostFromRelationships{value: val, isSet: true}
}

func (v NullableHostFromRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostFromRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

