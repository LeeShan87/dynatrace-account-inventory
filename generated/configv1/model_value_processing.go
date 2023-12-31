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

// checks if the ValueProcessing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueProcessing{}

// ValueProcessing Process values as specified.
type ValueProcessing struct {
	ExtractSubstring *ExtractSubstring `json:"extractSubstring,omitempty"`
	// Split (preprocessed) string values at this separator.
	SplitAt *string `json:"splitAt,omitempty"`
	// Prune Whitespaces. Defaults to false.
	Trim bool `json:"trim"`
	ValueCondition *ValueCondition `json:"valueCondition,omitempty"`
	// Extract value from captured data per regex.
	ValueExtractorRegex *string `json:"valueExtractorRegex,omitempty"`
}

// NewValueProcessing instantiates a new ValueProcessing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueProcessing(trim bool) *ValueProcessing {
	this := ValueProcessing{}
	this.Trim = trim
	return &this
}

// NewValueProcessingWithDefaults instantiates a new ValueProcessing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueProcessingWithDefaults() *ValueProcessing {
	this := ValueProcessing{}
	return &this
}

// GetExtractSubstring returns the ExtractSubstring field value if set, zero value otherwise.
func (o *ValueProcessing) GetExtractSubstring() ExtractSubstring {
	if o == nil || IsNil(o.ExtractSubstring) {
		var ret ExtractSubstring
		return ret
	}
	return *o.ExtractSubstring
}

// GetExtractSubstringOk returns a tuple with the ExtractSubstring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueProcessing) GetExtractSubstringOk() (*ExtractSubstring, bool) {
	if o == nil || IsNil(o.ExtractSubstring) {
		return nil, false
	}
	return o.ExtractSubstring, true
}

// HasExtractSubstring returns a boolean if a field has been set.
func (o *ValueProcessing) HasExtractSubstring() bool {
	if o != nil && !IsNil(o.ExtractSubstring) {
		return true
	}

	return false
}

// SetExtractSubstring gets a reference to the given ExtractSubstring and assigns it to the ExtractSubstring field.
func (o *ValueProcessing) SetExtractSubstring(v ExtractSubstring) {
	o.ExtractSubstring = &v
}

// GetSplitAt returns the SplitAt field value if set, zero value otherwise.
func (o *ValueProcessing) GetSplitAt() string {
	if o == nil || IsNil(o.SplitAt) {
		var ret string
		return ret
	}
	return *o.SplitAt
}

// GetSplitAtOk returns a tuple with the SplitAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueProcessing) GetSplitAtOk() (*string, bool) {
	if o == nil || IsNil(o.SplitAt) {
		return nil, false
	}
	return o.SplitAt, true
}

// HasSplitAt returns a boolean if a field has been set.
func (o *ValueProcessing) HasSplitAt() bool {
	if o != nil && !IsNil(o.SplitAt) {
		return true
	}

	return false
}

// SetSplitAt gets a reference to the given string and assigns it to the SplitAt field.
func (o *ValueProcessing) SetSplitAt(v string) {
	o.SplitAt = &v
}

// GetTrim returns the Trim field value
func (o *ValueProcessing) GetTrim() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Trim
}

// GetTrimOk returns a tuple with the Trim field value
// and a boolean to check if the value has been set.
func (o *ValueProcessing) GetTrimOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trim, true
}

// SetTrim sets field value
func (o *ValueProcessing) SetTrim(v bool) {
	o.Trim = v
}

// GetValueCondition returns the ValueCondition field value if set, zero value otherwise.
func (o *ValueProcessing) GetValueCondition() ValueCondition {
	if o == nil || IsNil(o.ValueCondition) {
		var ret ValueCondition
		return ret
	}
	return *o.ValueCondition
}

// GetValueConditionOk returns a tuple with the ValueCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueProcessing) GetValueConditionOk() (*ValueCondition, bool) {
	if o == nil || IsNil(o.ValueCondition) {
		return nil, false
	}
	return o.ValueCondition, true
}

// HasValueCondition returns a boolean if a field has been set.
func (o *ValueProcessing) HasValueCondition() bool {
	if o != nil && !IsNil(o.ValueCondition) {
		return true
	}

	return false
}

// SetValueCondition gets a reference to the given ValueCondition and assigns it to the ValueCondition field.
func (o *ValueProcessing) SetValueCondition(v ValueCondition) {
	o.ValueCondition = &v
}

// GetValueExtractorRegex returns the ValueExtractorRegex field value if set, zero value otherwise.
func (o *ValueProcessing) GetValueExtractorRegex() string {
	if o == nil || IsNil(o.ValueExtractorRegex) {
		var ret string
		return ret
	}
	return *o.ValueExtractorRegex
}

// GetValueExtractorRegexOk returns a tuple with the ValueExtractorRegex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueProcessing) GetValueExtractorRegexOk() (*string, bool) {
	if o == nil || IsNil(o.ValueExtractorRegex) {
		return nil, false
	}
	return o.ValueExtractorRegex, true
}

// HasValueExtractorRegex returns a boolean if a field has been set.
func (o *ValueProcessing) HasValueExtractorRegex() bool {
	if o != nil && !IsNil(o.ValueExtractorRegex) {
		return true
	}

	return false
}

// SetValueExtractorRegex gets a reference to the given string and assigns it to the ValueExtractorRegex field.
func (o *ValueProcessing) SetValueExtractorRegex(v string) {
	o.ValueExtractorRegex = &v
}

func (o ValueProcessing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValueProcessing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtractSubstring) {
		toSerialize["extractSubstring"] = o.ExtractSubstring
	}
	if !IsNil(o.SplitAt) {
		toSerialize["splitAt"] = o.SplitAt
	}
	toSerialize["trim"] = o.Trim
	if !IsNil(o.ValueCondition) {
		toSerialize["valueCondition"] = o.ValueCondition
	}
	if !IsNil(o.ValueExtractorRegex) {
		toSerialize["valueExtractorRegex"] = o.ValueExtractorRegex
	}
	return toSerialize, nil
}

type NullableValueProcessing struct {
	value *ValueProcessing
	isSet bool
}

func (v NullableValueProcessing) Get() *ValueProcessing {
	return v.value
}

func (v *NullableValueProcessing) Set(val *ValueProcessing) {
	v.value = val
	v.isSet = true
}

func (v NullableValueProcessing) IsSet() bool {
	return v.isSet
}

func (v *NullableValueProcessing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueProcessing(val *ValueProcessing) *NullableValueProcessing {
	return &NullableValueProcessing{value: val, isSet: true}
}

func (v NullableValueProcessing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValueProcessing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


