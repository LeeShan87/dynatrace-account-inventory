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

// checks if the VisitNumActionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisitNumActionDetails{}

// VisitNumActionDetails Configuration of a number of user actions-based conversion goal.
type VisitNumActionDetails struct {
	// The number of user actions to hit the conversion goal.
	NumUserActions *int32 `json:"numUserActions,omitempty"`
}

// NewVisitNumActionDetails instantiates a new VisitNumActionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisitNumActionDetails() *VisitNumActionDetails {
	this := VisitNumActionDetails{}
	return &this
}

// NewVisitNumActionDetailsWithDefaults instantiates a new VisitNumActionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisitNumActionDetailsWithDefaults() *VisitNumActionDetails {
	this := VisitNumActionDetails{}
	return &this
}

// GetNumUserActions returns the NumUserActions field value if set, zero value otherwise.
func (o *VisitNumActionDetails) GetNumUserActions() int32 {
	if o == nil || IsNil(o.NumUserActions) {
		var ret int32
		return ret
	}
	return *o.NumUserActions
}

// GetNumUserActionsOk returns a tuple with the NumUserActions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisitNumActionDetails) GetNumUserActionsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumUserActions) {
		return nil, false
	}
	return o.NumUserActions, true
}

// HasNumUserActions returns a boolean if a field has been set.
func (o *VisitNumActionDetails) HasNumUserActions() bool {
	if o != nil && !IsNil(o.NumUserActions) {
		return true
	}

	return false
}

// SetNumUserActions gets a reference to the given int32 and assigns it to the NumUserActions field.
func (o *VisitNumActionDetails) SetNumUserActions(v int32) {
	o.NumUserActions = &v
}

func (o VisitNumActionDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisitNumActionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumUserActions) {
		toSerialize["numUserActions"] = o.NumUserActions
	}
	return toSerialize, nil
}

type NullableVisitNumActionDetails struct {
	value *VisitNumActionDetails
	isSet bool
}

func (v NullableVisitNumActionDetails) Get() *VisitNumActionDetails {
	return v.value
}

func (v *NullableVisitNumActionDetails) Set(val *VisitNumActionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableVisitNumActionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableVisitNumActionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisitNumActionDetails(val *VisitNumActionDetails) *NullableVisitNumActionDetails {
	return &NullableVisitNumActionDetails{value: val, isSet: true}
}

func (v NullableVisitNumActionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisitNumActionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

