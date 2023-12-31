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

// checks if the AttackEntrypointPayloadInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackEntrypointPayloadInner{}

// AttackEntrypointPayloadInner A list of values that has possibly been truncated.
type AttackEntrypointPayloadInner struct {
	TruncationInfo *TruncationInfo `json:"truncationInfo,omitempty"`
	// Values of the list.
	Values []EntrypointPayload `json:"values,omitempty"`
}

// NewAttackEntrypointPayloadInner instantiates a new AttackEntrypointPayloadInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackEntrypointPayloadInner() *AttackEntrypointPayloadInner {
	this := AttackEntrypointPayloadInner{}
	return &this
}

// NewAttackEntrypointPayloadInnerWithDefaults instantiates a new AttackEntrypointPayloadInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackEntrypointPayloadInnerWithDefaults() *AttackEntrypointPayloadInner {
	this := AttackEntrypointPayloadInner{}
	return &this
}

// GetTruncationInfo returns the TruncationInfo field value if set, zero value otherwise.
func (o *AttackEntrypointPayloadInner) GetTruncationInfo() TruncationInfo {
	if o == nil || IsNil(o.TruncationInfo) {
		var ret TruncationInfo
		return ret
	}
	return *o.TruncationInfo
}

// GetTruncationInfoOk returns a tuple with the TruncationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackEntrypointPayloadInner) GetTruncationInfoOk() (*TruncationInfo, bool) {
	if o == nil || IsNil(o.TruncationInfo) {
		return nil, false
	}
	return o.TruncationInfo, true
}

// HasTruncationInfo returns a boolean if a field has been set.
func (o *AttackEntrypointPayloadInner) HasTruncationInfo() bool {
	if o != nil && !IsNil(o.TruncationInfo) {
		return true
	}

	return false
}

// SetTruncationInfo gets a reference to the given TruncationInfo and assigns it to the TruncationInfo field.
func (o *AttackEntrypointPayloadInner) SetTruncationInfo(v TruncationInfo) {
	o.TruncationInfo = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *AttackEntrypointPayloadInner) GetValues() []EntrypointPayload {
	if o == nil || IsNil(o.Values) {
		var ret []EntrypointPayload
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackEntrypointPayloadInner) GetValuesOk() ([]EntrypointPayload, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *AttackEntrypointPayloadInner) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []EntrypointPayload and assigns it to the Values field.
func (o *AttackEntrypointPayloadInner) SetValues(v []EntrypointPayload) {
	o.Values = v
}

func (o AttackEntrypointPayloadInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackEntrypointPayloadInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TruncationInfo) {
		toSerialize["truncationInfo"] = o.TruncationInfo
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableAttackEntrypointPayloadInner struct {
	value *AttackEntrypointPayloadInner
	isSet bool
}

func (v NullableAttackEntrypointPayloadInner) Get() *AttackEntrypointPayloadInner {
	return v.value
}

func (v *NullableAttackEntrypointPayloadInner) Set(val *AttackEntrypointPayloadInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackEntrypointPayloadInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackEntrypointPayloadInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackEntrypointPayloadInner(val *AttackEntrypointPayloadInner) *NullableAttackEntrypointPayloadInner {
	return &NullableAttackEntrypointPayloadInner{value: val, isSet: true}
}

func (v NullableAttackEntrypointPayloadInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackEntrypointPayloadInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


