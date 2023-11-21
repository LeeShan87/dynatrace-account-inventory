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

// checks if the MobileSessionUserActionPropertyShort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MobileSessionUserActionPropertyShort{}

// MobileSessionUserActionPropertyShort A short representation of mobile session or user action property.
type MobileSessionUserActionPropertyShort struct {
	// The display name of the session or user action property.
	DisplayName *string `json:"displayName,omitempty"`
	// The key of the session or user action property.
	Key string `json:"key"`
}

// NewMobileSessionUserActionPropertyShort instantiates a new MobileSessionUserActionPropertyShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMobileSessionUserActionPropertyShort(key string) *MobileSessionUserActionPropertyShort {
	this := MobileSessionUserActionPropertyShort{}
	this.Key = key
	return &this
}

// NewMobileSessionUserActionPropertyShortWithDefaults instantiates a new MobileSessionUserActionPropertyShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMobileSessionUserActionPropertyShortWithDefaults() *MobileSessionUserActionPropertyShort {
	this := MobileSessionUserActionPropertyShort{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MobileSessionUserActionPropertyShort) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MobileSessionUserActionPropertyShort) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MobileSessionUserActionPropertyShort) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MobileSessionUserActionPropertyShort) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetKey returns the Key field value
func (o *MobileSessionUserActionPropertyShort) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MobileSessionUserActionPropertyShort) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MobileSessionUserActionPropertyShort) SetKey(v string) {
	o.Key = v
}

func (o MobileSessionUserActionPropertyShort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MobileSessionUserActionPropertyShort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	toSerialize["key"] = o.Key
	return toSerialize, nil
}

type NullableMobileSessionUserActionPropertyShort struct {
	value *MobileSessionUserActionPropertyShort
	isSet bool
}

func (v NullableMobileSessionUserActionPropertyShort) Get() *MobileSessionUserActionPropertyShort {
	return v.value
}

func (v *NullableMobileSessionUserActionPropertyShort) Set(val *MobileSessionUserActionPropertyShort) {
	v.value = val
	v.isSet = true
}

func (v NullableMobileSessionUserActionPropertyShort) IsSet() bool {
	return v.isSet
}

func (v *NullableMobileSessionUserActionPropertyShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMobileSessionUserActionPropertyShort(val *MobileSessionUserActionPropertyShort) *NullableMobileSessionUserActionPropertyShort {
	return &NullableMobileSessionUserActionPropertyShort{value: val, isSet: true}
}

func (v NullableMobileSessionUserActionPropertyShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMobileSessionUserActionPropertyShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


