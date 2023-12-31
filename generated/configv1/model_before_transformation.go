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

// checks if the BeforeTransformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeforeTransformation{}

// BeforeTransformation The transformation of the `BEFORE` type.   The transformation keeps the value before the specified delimiter and removes everything after it.
type BeforeTransformation struct {
	// The delimiter of the transformation. The transformation keeps everything before this delimiter and removes everything after it.    The delimiter itself is not kept.   If several delimiters appear in the initial value, only the first one is used.
	Delimiter string `json:"delimiter"`
}

// NewBeforeTransformation instantiates a new BeforeTransformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeforeTransformation(delimiter string, type_ string) *BeforeTransformation {
	this := BeforeTransformation{}
	this.Type = type_
	this.Delimiter = delimiter
	return &this
}

// NewBeforeTransformationWithDefaults instantiates a new BeforeTransformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeforeTransformationWithDefaults() *BeforeTransformation {
	this := BeforeTransformation{}
	return &this
}

// GetDelimiter returns the Delimiter field value
func (o *BeforeTransformation) GetDelimiter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value
// and a boolean to check if the value has been set.
func (o *BeforeTransformation) GetDelimiterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delimiter, true
}

// SetDelimiter sets field value
func (o *BeforeTransformation) SetDelimiter(v string) {
	o.Delimiter = v
}

func (o BeforeTransformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeforeTransformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["delimiter"] = o.Delimiter
	return toSerialize, nil
}

type NullableBeforeTransformation struct {
	value *BeforeTransformation
	isSet bool
}

func (v NullableBeforeTransformation) Get() *BeforeTransformation {
	return v.value
}

func (v *NullableBeforeTransformation) Set(val *BeforeTransformation) {
	v.value = val
	v.isSet = true
}

func (v NullableBeforeTransformation) IsSet() bool {
	return v.isSet
}

func (v *NullableBeforeTransformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeforeTransformation(val *BeforeTransformation) *NullableBeforeTransformation {
	return &NullableBeforeTransformation{value: val, isSet: true}
}

func (v NullableBeforeTransformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeforeTransformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


