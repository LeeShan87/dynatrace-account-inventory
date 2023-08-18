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

// checks if the UiButtonCustomization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiButtonCustomization{}

// UiButtonCustomization UI customization for defining a button that calls a function when pressed
type UiButtonCustomization struct {
	// The description to be shown in a tooltip when hovering over the button
	Description *string `json:"description,omitempty"`
	// The label of the button
	DisplayName string `json:"displayName"`
	// The identifier of the function to be called when the button is pressed
	Identifier string `json:"identifier"`
	Insert UiButtonCustomizationInsert `json:"insert"`
}

// NewUiButtonCustomization instantiates a new UiButtonCustomization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiButtonCustomization(displayName string, identifier string, insert UiButtonCustomizationInsert) *UiButtonCustomization {
	this := UiButtonCustomization{}
	this.DisplayName = displayName
	this.Identifier = identifier
	this.Insert = insert
	return &this
}

// NewUiButtonCustomizationWithDefaults instantiates a new UiButtonCustomization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiButtonCustomizationWithDefaults() *UiButtonCustomization {
	this := UiButtonCustomization{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UiButtonCustomization) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiButtonCustomization) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UiButtonCustomization) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UiButtonCustomization) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value
func (o *UiButtonCustomization) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *UiButtonCustomization) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *UiButtonCustomization) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetIdentifier returns the Identifier field value
func (o *UiButtonCustomization) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *UiButtonCustomization) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *UiButtonCustomization) SetIdentifier(v string) {
	o.Identifier = v
}

// GetInsert returns the Insert field value
func (o *UiButtonCustomization) GetInsert() UiButtonCustomizationInsert {
	if o == nil {
		var ret UiButtonCustomizationInsert
		return ret
	}

	return o.Insert
}

// GetInsertOk returns a tuple with the Insert field value
// and a boolean to check if the value has been set.
func (o *UiButtonCustomization) GetInsertOk() (*UiButtonCustomizationInsert, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Insert, true
}

// SetInsert sets field value
func (o *UiButtonCustomization) SetInsert(v UiButtonCustomizationInsert) {
	o.Insert = v
}

func (o UiButtonCustomization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiButtonCustomization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["identifier"] = o.Identifier
	toSerialize["insert"] = o.Insert
	return toSerialize, nil
}

type NullableUiButtonCustomization struct {
	value *UiButtonCustomization
	isSet bool
}

func (v NullableUiButtonCustomization) Get() *UiButtonCustomization {
	return v.value
}

func (v *NullableUiButtonCustomization) Set(val *UiButtonCustomization) {
	v.value = val
	v.isSet = true
}

func (v NullableUiButtonCustomization) IsSet() bool {
	return v.isSet
}

func (v *NullableUiButtonCustomization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiButtonCustomization(val *UiButtonCustomization) *NullableUiButtonCustomization {
	return &NullableUiButtonCustomization{value: val, isSet: true}
}

func (v NullableUiButtonCustomization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiButtonCustomization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


