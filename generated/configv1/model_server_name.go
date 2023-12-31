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

// checks if the ServerName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerName{}

// ServerName The contribution to the service ID calculation from the detected server name.    You have two mutually exclusive options:  * Override the detected value with a specified static value. Specify the new value in the **valueOverride** field.  * Dynamically transform the detected value. Specify the transformation parameters in the **transformations** field.
type ServerName struct {
	// Transformations to be applied to the detected value.
	Transformations []TransformationBase `json:"transformations,omitempty"`
	// The value to be used instead of the detected value.
	ValueOverride *string `json:"valueOverride,omitempty"`
}

// NewServerName instantiates a new ServerName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerName() *ServerName {
	this := ServerName{}
	return &this
}

// NewServerNameWithDefaults instantiates a new ServerName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerNameWithDefaults() *ServerName {
	this := ServerName{}
	return &this
}

// GetTransformations returns the Transformations field value if set, zero value otherwise.
func (o *ServerName) GetTransformations() []TransformationBase {
	if o == nil || IsNil(o.Transformations) {
		var ret []TransformationBase
		return ret
	}
	return o.Transformations
}

// GetTransformationsOk returns a tuple with the Transformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerName) GetTransformationsOk() ([]TransformationBase, bool) {
	if o == nil || IsNil(o.Transformations) {
		return nil, false
	}
	return o.Transformations, true
}

// HasTransformations returns a boolean if a field has been set.
func (o *ServerName) HasTransformations() bool {
	if o != nil && !IsNil(o.Transformations) {
		return true
	}

	return false
}

// SetTransformations gets a reference to the given []TransformationBase and assigns it to the Transformations field.
func (o *ServerName) SetTransformations(v []TransformationBase) {
	o.Transformations = v
}

// GetValueOverride returns the ValueOverride field value if set, zero value otherwise.
func (o *ServerName) GetValueOverride() string {
	if o == nil || IsNil(o.ValueOverride) {
		var ret string
		return ret
	}
	return *o.ValueOverride
}

// GetValueOverrideOk returns a tuple with the ValueOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerName) GetValueOverrideOk() (*string, bool) {
	if o == nil || IsNil(o.ValueOverride) {
		return nil, false
	}
	return o.ValueOverride, true
}

// HasValueOverride returns a boolean if a field has been set.
func (o *ServerName) HasValueOverride() bool {
	if o != nil && !IsNil(o.ValueOverride) {
		return true
	}

	return false
}

// SetValueOverride gets a reference to the given string and assigns it to the ValueOverride field.
func (o *ServerName) SetValueOverride(v string) {
	o.ValueOverride = &v
}

func (o ServerName) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Transformations) {
		toSerialize["transformations"] = o.Transformations
	}
	if !IsNil(o.ValueOverride) {
		toSerialize["valueOverride"] = o.ValueOverride
	}
	return toSerialize, nil
}

type NullableServerName struct {
	value *ServerName
	isSet bool
}

func (v NullableServerName) Get() *ServerName {
	return v.value
}

func (v *NullableServerName) Set(val *ServerName) {
	v.value = val
	v.isSet = true
}

func (v NullableServerName) IsSet() bool {
	return v.isSet
}

func (v *NullableServerName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerName(val *ServerName) *NullableServerName {
	return &NullableServerName{value: val, isSet: true}
}

func (v NullableServerName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


