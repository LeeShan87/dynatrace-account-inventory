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

// checks if the SchemaList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaList{}

// SchemaList The list of available settings schemas.
type SchemaList struct {
	// A list of settings schemas.
	Items []SchemaStub `json:"items"`
	// The number of schemas in the list.
	TotalCount int64 `json:"totalCount"`
}

// NewSchemaList instantiates a new SchemaList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaList(items []SchemaStub, totalCount int64) *SchemaList {
	this := SchemaList{}
	this.Items = items
	this.TotalCount = totalCount
	return &this
}

// NewSchemaListWithDefaults instantiates a new SchemaList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaListWithDefaults() *SchemaList {
	this := SchemaList{}
	return &this
}

// GetItems returns the Items field value
func (o *SchemaList) GetItems() []SchemaStub {
	if o == nil {
		var ret []SchemaStub
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *SchemaList) GetItemsOk() ([]SchemaStub, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *SchemaList) SetItems(v []SchemaStub) {
	o.Items = v
}

// GetTotalCount returns the TotalCount field value
func (o *SchemaList) GetTotalCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *SchemaList) GetTotalCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *SchemaList) SetTotalCount(v int64) {
	o.TotalCount = v
}

func (o SchemaList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["items"] = o.Items
	toSerialize["totalCount"] = o.TotalCount
	return toSerialize, nil
}

type NullableSchemaList struct {
	value *SchemaList
	isSet bool
}

func (v NullableSchemaList) Get() *SchemaList {
	return v.value
}

func (v *NullableSchemaList) Set(val *SchemaList) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaList) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaList(val *SchemaList) *NullableSchemaList {
	return &NullableSchemaList{value: val, isSet: true}
}

func (v NullableSchemaList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


