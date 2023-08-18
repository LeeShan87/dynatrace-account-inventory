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

// checks if the PathDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PathDefinition{}

// PathDefinition A filtering criterion for log path.
type PathDefinition struct {
	// The path to the required log path.    If the **type** is set to `WILDCARD`, it may contain wildcard characters (`*`).
	Definition string `json:"definition"`
	// The type of the log path **definition**: fixed or an expression with wildcards.
	Type string `json:"type"`
}

// NewPathDefinition instantiates a new PathDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathDefinition(definition string, type_ string) *PathDefinition {
	this := PathDefinition{}
	this.Definition = definition
	this.Type = type_
	return &this
}

// NewPathDefinitionWithDefaults instantiates a new PathDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathDefinitionWithDefaults() *PathDefinition {
	this := PathDefinition{}
	return &this
}

// GetDefinition returns the Definition field value
func (o *PathDefinition) GetDefinition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value
// and a boolean to check if the value has been set.
func (o *PathDefinition) GetDefinitionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Definition, true
}

// SetDefinition sets field value
func (o *PathDefinition) SetDefinition(v string) {
	o.Definition = v
}

// GetType returns the Type field value
func (o *PathDefinition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PathDefinition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PathDefinition) SetType(v string) {
	o.Type = v
}

func (o PathDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["definition"] = o.Definition
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullablePathDefinition struct {
	value *PathDefinition
	isSet bool
}

func (v NullablePathDefinition) Get() *PathDefinition {
	return v.value
}

func (v *NullablePathDefinition) Set(val *PathDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullablePathDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullablePathDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathDefinition(val *PathDefinition) *NullablePathDefinition {
	return &NullablePathDefinition{value: val, isSet: true}
}

func (v NullablePathDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


