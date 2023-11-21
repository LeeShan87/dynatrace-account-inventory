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

// checks if the TruncatableListString type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TruncatableListString{}

// TruncatableListString A list of values that has possibly been truncated.
type TruncatableListString struct {
	TruncationInfo *TruncationInfo `json:"truncationInfo,omitempty"`
	// Values of the list.
	Values []string `json:"values,omitempty"`
}

// NewTruncatableListString instantiates a new TruncatableListString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTruncatableListString() *TruncatableListString {
	this := TruncatableListString{}
	return &this
}

// NewTruncatableListStringWithDefaults instantiates a new TruncatableListString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTruncatableListStringWithDefaults() *TruncatableListString {
	this := TruncatableListString{}
	return &this
}

// GetTruncationInfo returns the TruncationInfo field value if set, zero value otherwise.
func (o *TruncatableListString) GetTruncationInfo() TruncationInfo {
	if o == nil || IsNil(o.TruncationInfo) {
		var ret TruncationInfo
		return ret
	}
	return *o.TruncationInfo
}

// GetTruncationInfoOk returns a tuple with the TruncationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TruncatableListString) GetTruncationInfoOk() (*TruncationInfo, bool) {
	if o == nil || IsNil(o.TruncationInfo) {
		return nil, false
	}
	return o.TruncationInfo, true
}

// HasTruncationInfo returns a boolean if a field has been set.
func (o *TruncatableListString) HasTruncationInfo() bool {
	if o != nil && !IsNil(o.TruncationInfo) {
		return true
	}

	return false
}

// SetTruncationInfo gets a reference to the given TruncationInfo and assigns it to the TruncationInfo field.
func (o *TruncatableListString) SetTruncationInfo(v TruncationInfo) {
	o.TruncationInfo = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *TruncatableListString) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TruncatableListString) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *TruncatableListString) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *TruncatableListString) SetValues(v []string) {
	o.Values = v
}

func (o TruncatableListString) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TruncatableListString) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TruncationInfo) {
		toSerialize["truncationInfo"] = o.TruncationInfo
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableTruncatableListString struct {
	value *TruncatableListString
	isSet bool
}

func (v NullableTruncatableListString) Get() *TruncatableListString {
	return v.value
}

func (v *NullableTruncatableListString) Set(val *TruncatableListString) {
	v.value = val
	v.isSet = true
}

func (v NullableTruncatableListString) IsSet() bool {
	return v.isSet
}

func (v *NullableTruncatableListString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTruncatableListString(val *TruncatableListString) *NullableTruncatableListString {
	return &NullableTruncatableListString{value: val, isSet: true}
}

func (v NullableTruncatableListString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTruncatableListString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


