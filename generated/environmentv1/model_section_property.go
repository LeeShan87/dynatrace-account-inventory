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

// checks if the SectionProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SectionProperty{}

// SectionProperty A single agent property with it's associated section.
type SectionProperty struct {
	// The property key.
	Key *string `json:"key,omitempty"`
	// The section this property belongs to.
	Section *string `json:"section,omitempty"`
	// The property value.
	Value *string `json:"value,omitempty"`
}

// NewSectionProperty instantiates a new SectionProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSectionProperty() *SectionProperty {
	this := SectionProperty{}
	return &this
}

// NewSectionPropertyWithDefaults instantiates a new SectionProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSectionPropertyWithDefaults() *SectionProperty {
	this := SectionProperty{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *SectionProperty) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectionProperty) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *SectionProperty) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *SectionProperty) SetKey(v string) {
	o.Key = &v
}

// GetSection returns the Section field value if set, zero value otherwise.
func (o *SectionProperty) GetSection() string {
	if o == nil || IsNil(o.Section) {
		var ret string
		return ret
	}
	return *o.Section
}

// GetSectionOk returns a tuple with the Section field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectionProperty) GetSectionOk() (*string, bool) {
	if o == nil || IsNil(o.Section) {
		return nil, false
	}
	return o.Section, true
}

// HasSection returns a boolean if a field has been set.
func (o *SectionProperty) HasSection() bool {
	if o != nil && !IsNil(o.Section) {
		return true
	}

	return false
}

// SetSection gets a reference to the given string and assigns it to the Section field.
func (o *SectionProperty) SetSection(v string) {
	o.Section = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SectionProperty) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SectionProperty) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SectionProperty) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SectionProperty) SetValue(v string) {
	o.Value = &v
}

func (o SectionProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SectionProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Section) {
		toSerialize["section"] = o.Section
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSectionProperty struct {
	value *SectionProperty
	isSet bool
}

func (v NullableSectionProperty) Get() *SectionProperty {
	return v.value
}

func (v *NullableSectionProperty) Set(val *SectionProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableSectionProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableSectionProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSectionProperty(val *SectionProperty) *NullableSectionProperty {
	return &NullableSectionProperty{value: val, isSet: true}
}

func (v NullableSectionProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSectionProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

