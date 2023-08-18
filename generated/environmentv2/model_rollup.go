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

// checks if the Rollup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Rollup{}

// Rollup A way of viewing a series as a single value for the purpose of sorting or series-based filters.
type Rollup struct {
	Parameter *float32 `json:"parameter,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewRollup instantiates a new Rollup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollup() *Rollup {
	this := Rollup{}
	return &this
}

// NewRollupWithDefaults instantiates a new Rollup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollupWithDefaults() *Rollup {
	this := Rollup{}
	return &this
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *Rollup) GetParameter() float32 {
	if o == nil || IsNil(o.Parameter) {
		var ret float32
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rollup) GetParameterOk() (*float32, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *Rollup) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given float32 and assigns it to the Parameter field.
func (o *Rollup) SetParameter(v float32) {
	o.Parameter = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Rollup) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rollup) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Rollup) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Rollup) SetType(v string) {
	o.Type = &v
}

func (o Rollup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Rollup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableRollup struct {
	value *Rollup
	isSet bool
}

func (v NullableRollup) Get() *Rollup {
	return v.value
}

func (v *NullableRollup) Set(val *Rollup) {
	v.value = val
	v.isSet = true
}

func (v NullableRollup) IsSet() bool {
	return v.isSet
}

func (v *NullableRollup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollup(val *Rollup) *NullableRollup {
	return &NullableRollup{value: val, isSet: true}
}

func (v NullableRollup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


