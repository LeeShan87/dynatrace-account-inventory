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

// checks if the ConditionsFullWebRequestAttributeTypeDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConditionsFullWebRequestAttributeTypeDto{}

// ConditionsFullWebRequestAttributeTypeDto A condition of the service detection rule.
type ConditionsFullWebRequestAttributeTypeDto struct {
	// The type of the attribute to be checked.
	AttributeType string `json:"attributeType"`
	// A list of conditions for the rule.   If several conditions are specified, the AND logic applies.
	CompareOperations []CompareOperation `json:"compareOperations,omitempty"`
}

// NewConditionsFullWebRequestAttributeTypeDto instantiates a new ConditionsFullWebRequestAttributeTypeDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConditionsFullWebRequestAttributeTypeDto(attributeType string) *ConditionsFullWebRequestAttributeTypeDto {
	this := ConditionsFullWebRequestAttributeTypeDto{}
	this.AttributeType = attributeType
	return &this
}

// NewConditionsFullWebRequestAttributeTypeDtoWithDefaults instantiates a new ConditionsFullWebRequestAttributeTypeDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConditionsFullWebRequestAttributeTypeDtoWithDefaults() *ConditionsFullWebRequestAttributeTypeDto {
	this := ConditionsFullWebRequestAttributeTypeDto{}
	return &this
}

// GetAttributeType returns the AttributeType field value
func (o *ConditionsFullWebRequestAttributeTypeDto) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *ConditionsFullWebRequestAttributeTypeDto) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *ConditionsFullWebRequestAttributeTypeDto) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetCompareOperations returns the CompareOperations field value if set, zero value otherwise.
func (o *ConditionsFullWebRequestAttributeTypeDto) GetCompareOperations() []CompareOperation {
	if o == nil || IsNil(o.CompareOperations) {
		var ret []CompareOperation
		return ret
	}
	return o.CompareOperations
}

// GetCompareOperationsOk returns a tuple with the CompareOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConditionsFullWebRequestAttributeTypeDto) GetCompareOperationsOk() ([]CompareOperation, bool) {
	if o == nil || IsNil(o.CompareOperations) {
		return nil, false
	}
	return o.CompareOperations, true
}

// HasCompareOperations returns a boolean if a field has been set.
func (o *ConditionsFullWebRequestAttributeTypeDto) HasCompareOperations() bool {
	if o != nil && !IsNil(o.CompareOperations) {
		return true
	}

	return false
}

// SetCompareOperations gets a reference to the given []CompareOperation and assigns it to the CompareOperations field.
func (o *ConditionsFullWebRequestAttributeTypeDto) SetCompareOperations(v []CompareOperation) {
	o.CompareOperations = v
}

func (o ConditionsFullWebRequestAttributeTypeDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConditionsFullWebRequestAttributeTypeDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attributeType"] = o.AttributeType
	if !IsNil(o.CompareOperations) {
		toSerialize["compareOperations"] = o.CompareOperations
	}
	return toSerialize, nil
}

type NullableConditionsFullWebRequestAttributeTypeDto struct {
	value *ConditionsFullWebRequestAttributeTypeDto
	isSet bool
}

func (v NullableConditionsFullWebRequestAttributeTypeDto) Get() *ConditionsFullWebRequestAttributeTypeDto {
	return v.value
}

func (v *NullableConditionsFullWebRequestAttributeTypeDto) Set(val *ConditionsFullWebRequestAttributeTypeDto) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionsFullWebRequestAttributeTypeDto) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionsFullWebRequestAttributeTypeDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionsFullWebRequestAttributeTypeDto(val *ConditionsFullWebRequestAttributeTypeDto) *NullableConditionsFullWebRequestAttributeTypeDto {
	return &NullableConditionsFullWebRequestAttributeTypeDto{value: val, isSet: true}
}

func (v NullableConditionsFullWebRequestAttributeTypeDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionsFullWebRequestAttributeTypeDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


