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

// checks if the UiEmptyStateCustomization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiEmptyStateCustomization{}

// UiEmptyStateCustomization UI customization for empty state in a table
type UiEmptyStateCustomization struct {
	// The text to be shown in the empty state
	Text *string `json:"text,omitempty"`
}

// NewUiEmptyStateCustomization instantiates a new UiEmptyStateCustomization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiEmptyStateCustomization() *UiEmptyStateCustomization {
	this := UiEmptyStateCustomization{}
	return &this
}

// NewUiEmptyStateCustomizationWithDefaults instantiates a new UiEmptyStateCustomization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiEmptyStateCustomizationWithDefaults() *UiEmptyStateCustomization {
	this := UiEmptyStateCustomization{}
	return &this
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *UiEmptyStateCustomization) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiEmptyStateCustomization) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *UiEmptyStateCustomization) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *UiEmptyStateCustomization) SetText(v string) {
	o.Text = &v
}

func (o UiEmptyStateCustomization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiEmptyStateCustomization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableUiEmptyStateCustomization struct {
	value *UiEmptyStateCustomization
	isSet bool
}

func (v NullableUiEmptyStateCustomization) Get() *UiEmptyStateCustomization {
	return v.value
}

func (v *NullableUiEmptyStateCustomization) Set(val *UiEmptyStateCustomization) {
	v.value = val
	v.isSet = true
}

func (v NullableUiEmptyStateCustomization) IsSet() bool {
	return v.isSet
}

func (v *NullableUiEmptyStateCustomization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiEmptyStateCustomization(val *UiEmptyStateCustomization) *NullableUiEmptyStateCustomization {
	return &NullableUiEmptyStateCustomization{value: val, isSet: true}
}

func (v NullableUiEmptyStateCustomization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiEmptyStateCustomization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


