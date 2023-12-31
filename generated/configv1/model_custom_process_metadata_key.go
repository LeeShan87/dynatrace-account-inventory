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

// checks if the CustomProcessMetadataKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomProcessMetadataKey{}

// CustomProcessMetadataKey The key of the attribute, which need dynamic keys.   Not applicable otherwise, as the attibute itself acts as a key.
type CustomProcessMetadataKey struct {
	// The actual key of the custom metadata.
	Key string `json:"key"`
	// The source of the custom metadata.
	Source string `json:"source"`
}

// NewCustomProcessMetadataKey instantiates a new CustomProcessMetadataKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomProcessMetadataKey(key string, source string) *CustomProcessMetadataKey {
	this := CustomProcessMetadataKey{}
	this.Key = key
	this.Source = source
	return &this
}

// NewCustomProcessMetadataKeyWithDefaults instantiates a new CustomProcessMetadataKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomProcessMetadataKeyWithDefaults() *CustomProcessMetadataKey {
	this := CustomProcessMetadataKey{}
	return &this
}

// GetKey returns the Key field value
func (o *CustomProcessMetadataKey) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *CustomProcessMetadataKey) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *CustomProcessMetadataKey) SetKey(v string) {
	o.Key = v
}

// GetSource returns the Source field value
func (o *CustomProcessMetadataKey) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CustomProcessMetadataKey) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *CustomProcessMetadataKey) SetSource(v string) {
	o.Source = v
}

func (o CustomProcessMetadataKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomProcessMetadataKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullableCustomProcessMetadataKey struct {
	value *CustomProcessMetadataKey
	isSet bool
}

func (v NullableCustomProcessMetadataKey) Get() *CustomProcessMetadataKey {
	return v.value
}

func (v *NullableCustomProcessMetadataKey) Set(val *CustomProcessMetadataKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomProcessMetadataKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomProcessMetadataKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomProcessMetadataKey(val *CustomProcessMetadataKey) *NullableCustomProcessMetadataKey {
	return &NullableCustomProcessMetadataKey{value: val, isSet: true}
}

func (v NullableCustomProcessMetadataKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomProcessMetadataKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


