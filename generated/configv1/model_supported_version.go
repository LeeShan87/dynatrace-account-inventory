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

// checks if the SupportedVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedVersion{}

// SupportedVersion struct for SupportedVersion
type SupportedVersion struct {
	// The supported iOS file format version.
	Version *string `json:"version,omitempty"`
}

// NewSupportedVersion instantiates a new SupportedVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedVersion() *SupportedVersion {
	this := SupportedVersion{}
	return &this
}

// NewSupportedVersionWithDefaults instantiates a new SupportedVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedVersionWithDefaults() *SupportedVersion {
	this := SupportedVersion{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *SupportedVersion) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedVersion) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *SupportedVersion) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *SupportedVersion) SetVersion(v string) {
	o.Version = &v
}

func (o SupportedVersion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableSupportedVersion struct {
	value *SupportedVersion
	isSet bool
}

func (v NullableSupportedVersion) Get() *SupportedVersion {
	return v.value
}

func (v *NullableSupportedVersion) Set(val *SupportedVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedVersion(val *SupportedVersion) *NullableSupportedVersion {
	return &NullableSupportedVersion{value: val, isSet: true}
}

func (v NullableSupportedVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

