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

// checks if the CustomHostMetadataConditionKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomHostMetadataConditionKey{}

// CustomHostMetadataConditionKey The key for dynamic attributes of the `HOST_CUSTOM_METADATA_KEY` type.
type CustomHostMetadataConditionKey struct {
	DynamicKey CustomHostMetadataKey `json:"dynamicKey"`
}

// NewCustomHostMetadataConditionKey instantiates a new CustomHostMetadataConditionKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomHostMetadataConditionKey(dynamicKey CustomHostMetadataKey, attribute string) *CustomHostMetadataConditionKey {
	this := CustomHostMetadataConditionKey{}
	this.Attribute = attribute
	this.DynamicKey = dynamicKey
	return &this
}

// NewCustomHostMetadataConditionKeyWithDefaults instantiates a new CustomHostMetadataConditionKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomHostMetadataConditionKeyWithDefaults() *CustomHostMetadataConditionKey {
	this := CustomHostMetadataConditionKey{}
	return &this
}

// GetDynamicKey returns the DynamicKey field value
func (o *CustomHostMetadataConditionKey) GetDynamicKey() CustomHostMetadataKey {
	if o == nil {
		var ret CustomHostMetadataKey
		return ret
	}

	return o.DynamicKey
}

// GetDynamicKeyOk returns a tuple with the DynamicKey field value
// and a boolean to check if the value has been set.
func (o *CustomHostMetadataConditionKey) GetDynamicKeyOk() (*CustomHostMetadataKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DynamicKey, true
}

// SetDynamicKey sets field value
func (o *CustomHostMetadataConditionKey) SetDynamicKey(v CustomHostMetadataKey) {
	o.DynamicKey = v
}

func (o CustomHostMetadataConditionKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomHostMetadataConditionKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dynamicKey"] = o.DynamicKey
	return toSerialize, nil
}

type NullableCustomHostMetadataConditionKey struct {
	value *CustomHostMetadataConditionKey
	isSet bool
}

func (v NullableCustomHostMetadataConditionKey) Get() *CustomHostMetadataConditionKey {
	return v.value
}

func (v *NullableCustomHostMetadataConditionKey) Set(val *CustomHostMetadataConditionKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomHostMetadataConditionKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomHostMetadataConditionKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomHostMetadataConditionKey(val *CustomHostMetadataConditionKey) *NullableCustomHostMetadataConditionKey {
	return &NullableCustomHostMetadataConditionKey{value: val, isSet: true}
}

func (v NullableCustomHostMetadataConditionKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomHostMetadataConditionKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


