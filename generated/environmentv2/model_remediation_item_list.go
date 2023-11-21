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

// checks if the RemediationItemList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemediationItemList{}

// RemediationItemList A list of remediation items.
type RemediationItemList struct {
	// A list of remediation items.
	RemediationItems []RemediationItem `json:"remediationItems,omitempty"`
}

// NewRemediationItemList instantiates a new RemediationItemList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemediationItemList() *RemediationItemList {
	this := RemediationItemList{}
	return &this
}

// NewRemediationItemListWithDefaults instantiates a new RemediationItemList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemediationItemListWithDefaults() *RemediationItemList {
	this := RemediationItemList{}
	return &this
}

// GetRemediationItems returns the RemediationItems field value if set, zero value otherwise.
func (o *RemediationItemList) GetRemediationItems() []RemediationItem {
	if o == nil || IsNil(o.RemediationItems) {
		var ret []RemediationItem
		return ret
	}
	return o.RemediationItems
}

// GetRemediationItemsOk returns a tuple with the RemediationItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemList) GetRemediationItemsOk() ([]RemediationItem, bool) {
	if o == nil || IsNil(o.RemediationItems) {
		return nil, false
	}
	return o.RemediationItems, true
}

// HasRemediationItems returns a boolean if a field has been set.
func (o *RemediationItemList) HasRemediationItems() bool {
	if o != nil && !IsNil(o.RemediationItems) {
		return true
	}

	return false
}

// SetRemediationItems gets a reference to the given []RemediationItem and assigns it to the RemediationItems field.
func (o *RemediationItemList) SetRemediationItems(v []RemediationItem) {
	o.RemediationItems = v
}

func (o RemediationItemList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemediationItemList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RemediationItems) {
		toSerialize["remediationItems"] = o.RemediationItems
	}
	return toSerialize, nil
}

type NullableRemediationItemList struct {
	value *RemediationItemList
	isSet bool
}

func (v NullableRemediationItemList) Get() *RemediationItemList {
	return v.value
}

func (v *NullableRemediationItemList) Set(val *RemediationItemList) {
	v.value = val
	v.isSet = true
}

func (v NullableRemediationItemList) IsSet() bool {
	return v.isSet
}

func (v *NullableRemediationItemList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemediationItemList(val *RemediationItemList) *NullableRemediationItemList {
	return &NullableRemediationItemList{value: val, isSet: true}
}

func (v NullableRemediationItemList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemediationItemList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


