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

// checks if the UiExpandableCustomization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiExpandableCustomization{}

// UiExpandableCustomization UI customization for expandable section
type UiExpandableCustomization struct {
	// The display name
	DisplayName *string `json:"displayName,omitempty"`
	// Defines if the item should be expanded by default
	Expanded *bool `json:"expanded,omitempty"`
	// A list of sections
	Sections []UiExpandableSectionCustomization `json:"sections,omitempty"`
}

// NewUiExpandableCustomization instantiates a new UiExpandableCustomization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiExpandableCustomization() *UiExpandableCustomization {
	this := UiExpandableCustomization{}
	return &this
}

// NewUiExpandableCustomizationWithDefaults instantiates a new UiExpandableCustomization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiExpandableCustomizationWithDefaults() *UiExpandableCustomization {
	this := UiExpandableCustomization{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UiExpandableCustomization) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiExpandableCustomization) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UiExpandableCustomization) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UiExpandableCustomization) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExpanded returns the Expanded field value if set, zero value otherwise.
func (o *UiExpandableCustomization) GetExpanded() bool {
	if o == nil || IsNil(o.Expanded) {
		var ret bool
		return ret
	}
	return *o.Expanded
}

// GetExpandedOk returns a tuple with the Expanded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiExpandableCustomization) GetExpandedOk() (*bool, bool) {
	if o == nil || IsNil(o.Expanded) {
		return nil, false
	}
	return o.Expanded, true
}

// HasExpanded returns a boolean if a field has been set.
func (o *UiExpandableCustomization) HasExpanded() bool {
	if o != nil && !IsNil(o.Expanded) {
		return true
	}

	return false
}

// SetExpanded gets a reference to the given bool and assigns it to the Expanded field.
func (o *UiExpandableCustomization) SetExpanded(v bool) {
	o.Expanded = &v
}

// GetSections returns the Sections field value if set, zero value otherwise.
func (o *UiExpandableCustomization) GetSections() []UiExpandableSectionCustomization {
	if o == nil || IsNil(o.Sections) {
		var ret []UiExpandableSectionCustomization
		return ret
	}
	return o.Sections
}

// GetSectionsOk returns a tuple with the Sections field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiExpandableCustomization) GetSectionsOk() ([]UiExpandableSectionCustomization, bool) {
	if o == nil || IsNil(o.Sections) {
		return nil, false
	}
	return o.Sections, true
}

// HasSections returns a boolean if a field has been set.
func (o *UiExpandableCustomization) HasSections() bool {
	if o != nil && !IsNil(o.Sections) {
		return true
	}

	return false
}

// SetSections gets a reference to the given []UiExpandableSectionCustomization and assigns it to the Sections field.
func (o *UiExpandableCustomization) SetSections(v []UiExpandableSectionCustomization) {
	o.Sections = v
}

func (o UiExpandableCustomization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiExpandableCustomization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.Expanded) {
		toSerialize["expanded"] = o.Expanded
	}
	if !IsNil(o.Sections) {
		toSerialize["sections"] = o.Sections
	}
	return toSerialize, nil
}

type NullableUiExpandableCustomization struct {
	value *UiExpandableCustomization
	isSet bool
}

func (v NullableUiExpandableCustomization) Get() *UiExpandableCustomization {
	return v.value
}

func (v *NullableUiExpandableCustomization) Set(val *UiExpandableCustomization) {
	v.value = val
	v.isSet = true
}

func (v NullableUiExpandableCustomization) IsSet() bool {
	return v.isSet
}

func (v *NullableUiExpandableCustomization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiExpandableCustomization(val *UiExpandableCustomization) *NullableUiExpandableCustomization {
	return &NullableUiExpandableCustomization{value: val, isSet: true}
}

func (v NullableUiExpandableCustomization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiExpandableCustomization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


