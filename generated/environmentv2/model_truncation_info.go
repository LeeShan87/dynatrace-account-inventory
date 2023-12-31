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

// checks if the TruncationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TruncationInfo{}

// TruncationInfo Information on a possible truncation.
type TruncationInfo struct {
	// If the list/value has been truncated.
	Truncated *bool `json:"truncated,omitempty"`
}

// NewTruncationInfo instantiates a new TruncationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTruncationInfo() *TruncationInfo {
	this := TruncationInfo{}
	return &this
}

// NewTruncationInfoWithDefaults instantiates a new TruncationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTruncationInfoWithDefaults() *TruncationInfo {
	this := TruncationInfo{}
	return &this
}

// GetTruncated returns the Truncated field value if set, zero value otherwise.
func (o *TruncationInfo) GetTruncated() bool {
	if o == nil || IsNil(o.Truncated) {
		var ret bool
		return ret
	}
	return *o.Truncated
}

// GetTruncatedOk returns a tuple with the Truncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TruncationInfo) GetTruncatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Truncated) {
		return nil, false
	}
	return o.Truncated, true
}

// HasTruncated returns a boolean if a field has been set.
func (o *TruncationInfo) HasTruncated() bool {
	if o != nil && !IsNil(o.Truncated) {
		return true
	}

	return false
}

// SetTruncated gets a reference to the given bool and assigns it to the Truncated field.
func (o *TruncationInfo) SetTruncated(v bool) {
	o.Truncated = &v
}

func (o TruncationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TruncationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Truncated) {
		toSerialize["truncated"] = o.Truncated
	}
	return toSerialize, nil
}

type NullableTruncationInfo struct {
	value *TruncationInfo
	isSet bool
}

func (v NullableTruncationInfo) Get() *TruncationInfo {
	return v.value
}

func (v *NullableTruncationInfo) Set(val *TruncationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTruncationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTruncationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTruncationInfo(val *TruncationInfo) *NullableTruncationInfo {
	return &NullableTruncationInfo{value: val, isSet: true}
}

func (v NullableTruncationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTruncationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


