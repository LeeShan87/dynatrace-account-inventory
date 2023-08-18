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

// checks if the ConditionKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionKey{}

// ConditionKey The key to identify the data we're matching.   The actual set of fields and possible values depend on the type of the key. Find the list of actual objects in the description of the **type** field or see [JSON models](https://dt-url.net/0b83s6z).
type ConditionKey struct {
	// The attribute to be used for comparison.
	Attribute string `json:"attribute"`
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `PROCESS_CUSTOM_METADATA_KEY` -> CustomProcessMetadataConditionKey  * `HOST_CUSTOM_METADATA_KEY` -> CustomHostMetadataConditionKey  * `PROCESS_PREDEFINED_METADATA_KEY` -> ProcessMetadataConditionKey  * `STRING` -> StringConditionKey  * `STATIC` -> StaticConditionKey  
	Type *string `json:"type,omitempty"`
}

// NewConditionKey instantiates a new ConditionKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionKey(attribute string) *ConditionKey {
	this := ConditionKey{}
	this.Attribute = attribute
	return &this
}

// NewConditionKeyWithDefaults instantiates a new ConditionKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionKeyWithDefaults() *ConditionKey {
	this := ConditionKey{}
	return &this
}

// GetAttribute returns the Attribute field value
func (o *ConditionKey) GetAttribute() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Attribute
}

// GetAttributeOk returns a tuple with the Attribute field value
// and a boolean to check if the value has been set.
func (o *ConditionKey) GetAttributeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attribute, true
}

// SetAttribute sets field value
func (o *ConditionKey) SetAttribute(v string) {
	o.Attribute = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConditionKey) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionKey) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConditionKey) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConditionKey) SetType(v string) {
	o.Type = &v
}

func (o ConditionKey) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attribute"] = o.Attribute
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableConditionKey struct {
	value *ConditionKey
	isSet bool
}

func (v NullableConditionKey) Get() *ConditionKey {
	return v.value
}

func (v *NullableConditionKey) Set(val *ConditionKey) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionKey) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionKey(val *ConditionKey) *NullableConditionKey {
	return &NullableConditionKey{value: val, isSet: true}
}

func (v NullableConditionKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

