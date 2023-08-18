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

// checks if the TransformationBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransformationBase{}

// TransformationBase Configuration of transformation of the detected value.   If several transformations are specified, they are handled sequentially from top to bottom. Each transformation is applied to the result of the preceding transformation. For example, the second transformation is applied to the result of the first transformation.   The actual set of fields depends on the type of the transformation. Find the list of actual objects in the description of the **type** field or see [Service detection API - JSON models](https://dt-url.net/2ie3slq).
type TransformationBase struct {
	// Defines the actual set of fields depending on the value. See one of the following objects:   * `BEFORE` -> BeforeTransformation  * `AFTER` -> AfterTransformation  * `BETWEEN` -> BetweenTransformation  * `REPLACE_BETWEEN` -> ReplaceBetweenTransformation  * `REMOVE_NUMBERS` -> RemoveNumbersTransformation  * `REMOVE_CREDIT_CARDS` -> RemoveCreditCardNumbersTransformation  * `REMOVE_IBANS` -> RemoveIBANsTransformation  * `REMOVE_IPS` -> RemoveIPsTransformation  * `SPLIT_SELECT` -> SplitSelectTransformation  * `TAKE_SEGMENTS` -> TakeSegmentsTransformation  
	Type string `json:"type"`
}

// NewTransformationBase instantiates a new TransformationBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransformationBase(type_ string) *TransformationBase {
	this := TransformationBase{}
	this.Type = type_
	return &this
}

// NewTransformationBaseWithDefaults instantiates a new TransformationBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransformationBaseWithDefaults() *TransformationBase {
	this := TransformationBase{}
	return &this
}

// GetType returns the Type field value
func (o *TransformationBase) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransformationBase) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransformationBase) SetType(v string) {
	o.Type = v
}

func (o TransformationBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransformationBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTransformationBase struct {
	value *TransformationBase
	isSet bool
}

func (v NullableTransformationBase) Get() *TransformationBase {
	return v.value
}

func (v *NullableTransformationBase) Set(val *TransformationBase) {
	v.value = val
	v.isSet = true
}

func (v NullableTransformationBase) IsSet() bool {
	return v.isSet
}

func (v *NullableTransformationBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransformationBase(val *TransformationBase) *NullableTransformationBase {
	return &NullableTransformationBase{value: val, isSet: true}
}

func (v NullableTransformationBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransformationBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

