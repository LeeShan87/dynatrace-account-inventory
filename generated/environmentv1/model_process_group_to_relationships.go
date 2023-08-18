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

// checks if the ProcessGroupToRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessGroupToRelationships{}

// ProcessGroupToRelationships struct for ProcessGroupToRelationships
type ProcessGroupToRelationships struct {
	IsInstanceOf []string `json:"isInstanceOf,omitempty"`
	IsNetworkClientOfProcessGroup []string `json:"isNetworkClientOfProcessGroup,omitempty"`
	RunsOn []string `json:"runsOn,omitempty"`
}

// NewProcessGroupToRelationships instantiates a new ProcessGroupToRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessGroupToRelationships() *ProcessGroupToRelationships {
	this := ProcessGroupToRelationships{}
	return &this
}

// NewProcessGroupToRelationshipsWithDefaults instantiates a new ProcessGroupToRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessGroupToRelationshipsWithDefaults() *ProcessGroupToRelationships {
	this := ProcessGroupToRelationships{}
	return &this
}

// GetIsInstanceOf returns the IsInstanceOf field value if set, zero value otherwise.
func (o *ProcessGroupToRelationships) GetIsInstanceOf() []string {
	if o == nil || IsNil(o.IsInstanceOf) {
		var ret []string
		return ret
	}
	return o.IsInstanceOf
}

// GetIsInstanceOfOk returns a tuple with the IsInstanceOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupToRelationships) GetIsInstanceOfOk() ([]string, bool) {
	if o == nil || IsNil(o.IsInstanceOf) {
		return nil, false
	}
	return o.IsInstanceOf, true
}

// HasIsInstanceOf returns a boolean if a field has been set.
func (o *ProcessGroupToRelationships) HasIsInstanceOf() bool {
	if o != nil && !IsNil(o.IsInstanceOf) {
		return true
	}

	return false
}

// SetIsInstanceOf gets a reference to the given []string and assigns it to the IsInstanceOf field.
func (o *ProcessGroupToRelationships) SetIsInstanceOf(v []string) {
	o.IsInstanceOf = v
}

// GetIsNetworkClientOfProcessGroup returns the IsNetworkClientOfProcessGroup field value if set, zero value otherwise.
func (o *ProcessGroupToRelationships) GetIsNetworkClientOfProcessGroup() []string {
	if o == nil || IsNil(o.IsNetworkClientOfProcessGroup) {
		var ret []string
		return ret
	}
	return o.IsNetworkClientOfProcessGroup
}

// GetIsNetworkClientOfProcessGroupOk returns a tuple with the IsNetworkClientOfProcessGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupToRelationships) GetIsNetworkClientOfProcessGroupOk() ([]string, bool) {
	if o == nil || IsNil(o.IsNetworkClientOfProcessGroup) {
		return nil, false
	}
	return o.IsNetworkClientOfProcessGroup, true
}

// HasIsNetworkClientOfProcessGroup returns a boolean if a field has been set.
func (o *ProcessGroupToRelationships) HasIsNetworkClientOfProcessGroup() bool {
	if o != nil && !IsNil(o.IsNetworkClientOfProcessGroup) {
		return true
	}

	return false
}

// SetIsNetworkClientOfProcessGroup gets a reference to the given []string and assigns it to the IsNetworkClientOfProcessGroup field.
func (o *ProcessGroupToRelationships) SetIsNetworkClientOfProcessGroup(v []string) {
	o.IsNetworkClientOfProcessGroup = v
}

// GetRunsOn returns the RunsOn field value if set, zero value otherwise.
func (o *ProcessGroupToRelationships) GetRunsOn() []string {
	if o == nil || IsNil(o.RunsOn) {
		var ret []string
		return ret
	}
	return o.RunsOn
}

// GetRunsOnOk returns a tuple with the RunsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupToRelationships) GetRunsOnOk() ([]string, bool) {
	if o == nil || IsNil(o.RunsOn) {
		return nil, false
	}
	return o.RunsOn, true
}

// HasRunsOn returns a boolean if a field has been set.
func (o *ProcessGroupToRelationships) HasRunsOn() bool {
	if o != nil && !IsNil(o.RunsOn) {
		return true
	}

	return false
}

// SetRunsOn gets a reference to the given []string and assigns it to the RunsOn field.
func (o *ProcessGroupToRelationships) SetRunsOn(v []string) {
	o.RunsOn = v
}

func (o ProcessGroupToRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessGroupToRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsInstanceOf) {
		toSerialize["isInstanceOf"] = o.IsInstanceOf
	}
	if !IsNil(o.IsNetworkClientOfProcessGroup) {
		toSerialize["isNetworkClientOfProcessGroup"] = o.IsNetworkClientOfProcessGroup
	}
	if !IsNil(o.RunsOn) {
		toSerialize["runsOn"] = o.RunsOn
	}
	return toSerialize, nil
}

type NullableProcessGroupToRelationships struct {
	value *ProcessGroupToRelationships
	isSet bool
}

func (v NullableProcessGroupToRelationships) Get() *ProcessGroupToRelationships {
	return v.value
}

func (v *NullableProcessGroupToRelationships) Set(val *ProcessGroupToRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessGroupToRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessGroupToRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessGroupToRelationships(val *ProcessGroupToRelationships) *NullableProcessGroupToRelationships {
	return &NullableProcessGroupToRelationships{value: val, isSet: true}
}

func (v NullableProcessGroupToRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessGroupToRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


