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

// checks if the AllowlistedAwsAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllowlistedAwsAccount{}

// AllowlistedAwsAccount struct for AllowlistedAwsAccount
type AllowlistedAwsAccount struct {
	// The AWS account id to allowlist
	Id string `json:"id"`
}

// NewAllowlistedAwsAccount instantiates a new AllowlistedAwsAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowlistedAwsAccount(id string) *AllowlistedAwsAccount {
	this := AllowlistedAwsAccount{}
	this.Id = id
	return &this
}

// NewAllowlistedAwsAccountWithDefaults instantiates a new AllowlistedAwsAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowlistedAwsAccountWithDefaults() *AllowlistedAwsAccount {
	this := AllowlistedAwsAccount{}
	return &this
}

// GetId returns the Id field value
func (o *AllowlistedAwsAccount) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AllowlistedAwsAccount) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AllowlistedAwsAccount) SetId(v string) {
	o.Id = v
}

func (o AllowlistedAwsAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllowlistedAwsAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAllowlistedAwsAccount struct {
	value *AllowlistedAwsAccount
	isSet bool
}

func (v NullableAllowlistedAwsAccount) Get() *AllowlistedAwsAccount {
	return v.value
}

func (v *NullableAllowlistedAwsAccount) Set(val *AllowlistedAwsAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowlistedAwsAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowlistedAwsAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowlistedAwsAccount(val *AllowlistedAwsAccount) *NullableAllowlistedAwsAccount {
	return &NullableAllowlistedAwsAccount{value: val, isSet: true}
}

func (v NullableAllowlistedAwsAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowlistedAwsAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

