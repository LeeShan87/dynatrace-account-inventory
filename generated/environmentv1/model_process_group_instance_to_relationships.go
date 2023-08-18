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

// checks if the ProcessGroupInstanceToRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProcessGroupInstanceToRelationships{}

// ProcessGroupInstanceToRelationships struct for ProcessGroupInstanceToRelationships
type ProcessGroupInstanceToRelationships struct {
	IsNetworkClientOf []string `json:"isNetworkClientOf,omitempty"`
	RunsOnProcessGroupInstance []string `json:"runsOnProcessGroupInstance,omitempty"`
}

// NewProcessGroupInstanceToRelationships instantiates a new ProcessGroupInstanceToRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProcessGroupInstanceToRelationships() *ProcessGroupInstanceToRelationships {
	this := ProcessGroupInstanceToRelationships{}
	return &this
}

// NewProcessGroupInstanceToRelationshipsWithDefaults instantiates a new ProcessGroupInstanceToRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProcessGroupInstanceToRelationshipsWithDefaults() *ProcessGroupInstanceToRelationships {
	this := ProcessGroupInstanceToRelationships{}
	return &this
}

// GetIsNetworkClientOf returns the IsNetworkClientOf field value if set, zero value otherwise.
func (o *ProcessGroupInstanceToRelationships) GetIsNetworkClientOf() []string {
	if o == nil || IsNil(o.IsNetworkClientOf) {
		var ret []string
		return ret
	}
	return o.IsNetworkClientOf
}

// GetIsNetworkClientOfOk returns a tuple with the IsNetworkClientOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupInstanceToRelationships) GetIsNetworkClientOfOk() ([]string, bool) {
	if o == nil || IsNil(o.IsNetworkClientOf) {
		return nil, false
	}
	return o.IsNetworkClientOf, true
}

// HasIsNetworkClientOf returns a boolean if a field has been set.
func (o *ProcessGroupInstanceToRelationships) HasIsNetworkClientOf() bool {
	if o != nil && !IsNil(o.IsNetworkClientOf) {
		return true
	}

	return false
}

// SetIsNetworkClientOf gets a reference to the given []string and assigns it to the IsNetworkClientOf field.
func (o *ProcessGroupInstanceToRelationships) SetIsNetworkClientOf(v []string) {
	o.IsNetworkClientOf = v
}

// GetRunsOnProcessGroupInstance returns the RunsOnProcessGroupInstance field value if set, zero value otherwise.
func (o *ProcessGroupInstanceToRelationships) GetRunsOnProcessGroupInstance() []string {
	if o == nil || IsNil(o.RunsOnProcessGroupInstance) {
		var ret []string
		return ret
	}
	return o.RunsOnProcessGroupInstance
}

// GetRunsOnProcessGroupInstanceOk returns a tuple with the RunsOnProcessGroupInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProcessGroupInstanceToRelationships) GetRunsOnProcessGroupInstanceOk() ([]string, bool) {
	if o == nil || IsNil(o.RunsOnProcessGroupInstance) {
		return nil, false
	}
	return o.RunsOnProcessGroupInstance, true
}

// HasRunsOnProcessGroupInstance returns a boolean if a field has been set.
func (o *ProcessGroupInstanceToRelationships) HasRunsOnProcessGroupInstance() bool {
	if o != nil && !IsNil(o.RunsOnProcessGroupInstance) {
		return true
	}

	return false
}

// SetRunsOnProcessGroupInstance gets a reference to the given []string and assigns it to the RunsOnProcessGroupInstance field.
func (o *ProcessGroupInstanceToRelationships) SetRunsOnProcessGroupInstance(v []string) {
	o.RunsOnProcessGroupInstance = v
}

func (o ProcessGroupInstanceToRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProcessGroupInstanceToRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsNetworkClientOf) {
		toSerialize["isNetworkClientOf"] = o.IsNetworkClientOf
	}
	if !IsNil(o.RunsOnProcessGroupInstance) {
		toSerialize["runsOnProcessGroupInstance"] = o.RunsOnProcessGroupInstance
	}
	return toSerialize, nil
}

type NullableProcessGroupInstanceToRelationships struct {
	value *ProcessGroupInstanceToRelationships
	isSet bool
}

func (v NullableProcessGroupInstanceToRelationships) Get() *ProcessGroupInstanceToRelationships {
	return v.value
}

func (v *NullableProcessGroupInstanceToRelationships) Set(val *ProcessGroupInstanceToRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableProcessGroupInstanceToRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableProcessGroupInstanceToRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProcessGroupInstanceToRelationships(val *ProcessGroupInstanceToRelationships) *NullableProcessGroupInstanceToRelationships {
	return &NullableProcessGroupInstanceToRelationships{value: val, isSet: true}
}

func (v NullableProcessGroupInstanceToRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProcessGroupInstanceToRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


