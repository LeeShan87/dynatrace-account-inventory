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

// checks if the UiCustomization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UiCustomization{}

// UiCustomization Customization for UI elements
type UiCustomization struct {
	Callback *UiCallbackCustomization `json:"callback,omitempty"`
	Expandable *UiExpandableCustomization `json:"expandable,omitempty"`
	Table *UiTableCustomization `json:"table,omitempty"`
	Tabs *UiTabsCustomization `json:"tabs,omitempty"`
}

// NewUiCustomization instantiates a new UiCustomization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUiCustomization() *UiCustomization {
	this := UiCustomization{}
	return &this
}

// NewUiCustomizationWithDefaults instantiates a new UiCustomization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUiCustomizationWithDefaults() *UiCustomization {
	this := UiCustomization{}
	return &this
}

// GetCallback returns the Callback field value if set, zero value otherwise.
func (o *UiCustomization) GetCallback() UiCallbackCustomization {
	if o == nil || IsNil(o.Callback) {
		var ret UiCallbackCustomization
		return ret
	}
	return *o.Callback
}

// GetCallbackOk returns a tuple with the Callback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiCustomization) GetCallbackOk() (*UiCallbackCustomization, bool) {
	if o == nil || IsNil(o.Callback) {
		return nil, false
	}
	return o.Callback, true
}

// HasCallback returns a boolean if a field has been set.
func (o *UiCustomization) HasCallback() bool {
	if o != nil && !IsNil(o.Callback) {
		return true
	}

	return false
}

// SetCallback gets a reference to the given UiCallbackCustomization and assigns it to the Callback field.
func (o *UiCustomization) SetCallback(v UiCallbackCustomization) {
	o.Callback = &v
}

// GetExpandable returns the Expandable field value if set, zero value otherwise.
func (o *UiCustomization) GetExpandable() UiExpandableCustomization {
	if o == nil || IsNil(o.Expandable) {
		var ret UiExpandableCustomization
		return ret
	}
	return *o.Expandable
}

// GetExpandableOk returns a tuple with the Expandable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiCustomization) GetExpandableOk() (*UiExpandableCustomization, bool) {
	if o == nil || IsNil(o.Expandable) {
		return nil, false
	}
	return o.Expandable, true
}

// HasExpandable returns a boolean if a field has been set.
func (o *UiCustomization) HasExpandable() bool {
	if o != nil && !IsNil(o.Expandable) {
		return true
	}

	return false
}

// SetExpandable gets a reference to the given UiExpandableCustomization and assigns it to the Expandable field.
func (o *UiCustomization) SetExpandable(v UiExpandableCustomization) {
	o.Expandable = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *UiCustomization) GetTable() UiTableCustomization {
	if o == nil || IsNil(o.Table) {
		var ret UiTableCustomization
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiCustomization) GetTableOk() (*UiTableCustomization, bool) {
	if o == nil || IsNil(o.Table) {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *UiCustomization) HasTable() bool {
	if o != nil && !IsNil(o.Table) {
		return true
	}

	return false
}

// SetTable gets a reference to the given UiTableCustomization and assigns it to the Table field.
func (o *UiCustomization) SetTable(v UiTableCustomization) {
	o.Table = &v
}

// GetTabs returns the Tabs field value if set, zero value otherwise.
func (o *UiCustomization) GetTabs() UiTabsCustomization {
	if o == nil || IsNil(o.Tabs) {
		var ret UiTabsCustomization
		return ret
	}
	return *o.Tabs
}

// GetTabsOk returns a tuple with the Tabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UiCustomization) GetTabsOk() (*UiTabsCustomization, bool) {
	if o == nil || IsNil(o.Tabs) {
		return nil, false
	}
	return o.Tabs, true
}

// HasTabs returns a boolean if a field has been set.
func (o *UiCustomization) HasTabs() bool {
	if o != nil && !IsNil(o.Tabs) {
		return true
	}

	return false
}

// SetTabs gets a reference to the given UiTabsCustomization and assigns it to the Tabs field.
func (o *UiCustomization) SetTabs(v UiTabsCustomization) {
	o.Tabs = &v
}

func (o UiCustomization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UiCustomization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Callback) {
		toSerialize["callback"] = o.Callback
	}
	if !IsNil(o.Expandable) {
		toSerialize["expandable"] = o.Expandable
	}
	if !IsNil(o.Table) {
		toSerialize["table"] = o.Table
	}
	if !IsNil(o.Tabs) {
		toSerialize["tabs"] = o.Tabs
	}
	return toSerialize, nil
}

type NullableUiCustomization struct {
	value *UiCustomization
	isSet bool
}

func (v NullableUiCustomization) Get() *UiCustomization {
	return v.value
}

func (v *NullableUiCustomization) Set(val *UiCustomization) {
	v.value = val
	v.isSet = true
}

func (v NullableUiCustomization) IsSet() bool {
	return v.isSet
}

func (v *NullableUiCustomization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUiCustomization(val *UiCustomization) *NullableUiCustomization {
	return &NullableUiCustomization{value: val, isSet: true}
}

func (v NullableUiCustomization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUiCustomization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


