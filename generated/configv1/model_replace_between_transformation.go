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

// checks if the ReplaceBetweenTransformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReplaceBetweenTransformation{}

// ReplaceBetweenTransformation The transformation of the `REPLACE_BETWEEN` type.   The transformation replaces the content in between specified delimiters with the specified value. The rest of the string remains intact.
type ReplaceBetweenTransformation struct {
	// The starting delimiter. The transformation replaces everything from here until ending delimiter. The delimiter itself remain intact.
	After string `json:"after"`
	// The ending delimiter. The transformation replaces everything from starting delimiter until here. The delimiter itself remain intact.
	Before string `json:"before"`
	// The value to be used instead of the content between delimiters.
	Replacement string `json:"replacement"`
}

// NewReplaceBetweenTransformation instantiates a new ReplaceBetweenTransformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReplaceBetweenTransformation(after string, before string, replacement string, type_ string) *ReplaceBetweenTransformation {
	this := ReplaceBetweenTransformation{}
	this.Type = type_
	this.After = after
	this.Before = before
	this.Replacement = replacement
	return &this
}

// NewReplaceBetweenTransformationWithDefaults instantiates a new ReplaceBetweenTransformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReplaceBetweenTransformationWithDefaults() *ReplaceBetweenTransformation {
	this := ReplaceBetweenTransformation{}
	return &this
}

// GetAfter returns the After field value
func (o *ReplaceBetweenTransformation) GetAfter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.After
}

// GetAfterOk returns a tuple with the After field value
// and a boolean to check if the value has been set.
func (o *ReplaceBetweenTransformation) GetAfterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.After, true
}

// SetAfter sets field value
func (o *ReplaceBetweenTransformation) SetAfter(v string) {
	o.After = v
}

// GetBefore returns the Before field value
func (o *ReplaceBetweenTransformation) GetBefore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Before
}

// GetBeforeOk returns a tuple with the Before field value
// and a boolean to check if the value has been set.
func (o *ReplaceBetweenTransformation) GetBeforeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Before, true
}

// SetBefore sets field value
func (o *ReplaceBetweenTransformation) SetBefore(v string) {
	o.Before = v
}

// GetReplacement returns the Replacement field value
func (o *ReplaceBetweenTransformation) GetReplacement() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Replacement
}

// GetReplacementOk returns a tuple with the Replacement field value
// and a boolean to check if the value has been set.
func (o *ReplaceBetweenTransformation) GetReplacementOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Replacement, true
}

// SetReplacement sets field value
func (o *ReplaceBetweenTransformation) SetReplacement(v string) {
	o.Replacement = v
}

func (o ReplaceBetweenTransformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReplaceBetweenTransformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["after"] = o.After
	toSerialize["before"] = o.Before
	toSerialize["replacement"] = o.Replacement
	return toSerialize, nil
}

type NullableReplaceBetweenTransformation struct {
	value *ReplaceBetweenTransformation
	isSet bool
}

func (v NullableReplaceBetweenTransformation) Get() *ReplaceBetweenTransformation {
	return v.value
}

func (v *NullableReplaceBetweenTransformation) Set(val *ReplaceBetweenTransformation) {
	v.value = val
	v.isSet = true
}

func (v NullableReplaceBetweenTransformation) IsSet() bool {
	return v.isSet
}

func (v *NullableReplaceBetweenTransformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReplaceBetweenTransformation(val *ReplaceBetweenTransformation) *NullableReplaceBetweenTransformation {
	return &NullableReplaceBetweenTransformation{value: val, isSet: true}
}

func (v NullableReplaceBetweenTransformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReplaceBetweenTransformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

