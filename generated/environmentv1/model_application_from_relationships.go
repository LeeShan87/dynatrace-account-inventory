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

// checks if the ApplicationFromRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationFromRelationships{}

// ApplicationFromRelationships The list of outgoing calls from the application.
type ApplicationFromRelationships struct {
	Calls []string `json:"calls,omitempty"`
}

// NewApplicationFromRelationships instantiates a new ApplicationFromRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFromRelationships() *ApplicationFromRelationships {
	this := ApplicationFromRelationships{}
	return &this
}

// NewApplicationFromRelationshipsWithDefaults instantiates a new ApplicationFromRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFromRelationshipsWithDefaults() *ApplicationFromRelationships {
	this := ApplicationFromRelationships{}
	return &this
}

// GetCalls returns the Calls field value if set, zero value otherwise.
func (o *ApplicationFromRelationships) GetCalls() []string {
	if o == nil || IsNil(o.Calls) {
		var ret []string
		return ret
	}
	return o.Calls
}

// GetCallsOk returns a tuple with the Calls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFromRelationships) GetCallsOk() ([]string, bool) {
	if o == nil || IsNil(o.Calls) {
		return nil, false
	}
	return o.Calls, true
}

// HasCalls returns a boolean if a field has been set.
func (o *ApplicationFromRelationships) HasCalls() bool {
	if o != nil && !IsNil(o.Calls) {
		return true
	}

	return false
}

// SetCalls gets a reference to the given []string and assigns it to the Calls field.
func (o *ApplicationFromRelationships) SetCalls(v []string) {
	o.Calls = v
}

func (o ApplicationFromRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationFromRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Calls) {
		toSerialize["calls"] = o.Calls
	}
	return toSerialize, nil
}

type NullableApplicationFromRelationships struct {
	value *ApplicationFromRelationships
	isSet bool
}

func (v NullableApplicationFromRelationships) Get() *ApplicationFromRelationships {
	return v.value
}

func (v *NullableApplicationFromRelationships) Set(val *ApplicationFromRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFromRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFromRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFromRelationships(val *ApplicationFromRelationships) *NullableApplicationFromRelationships {
	return &NullableApplicationFromRelationships{value: val, isSet: true}
}

func (v NullableApplicationFromRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFromRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


