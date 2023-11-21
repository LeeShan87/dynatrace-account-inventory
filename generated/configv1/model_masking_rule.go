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

// checks if the MaskingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskingRule{}

// MaskingRule The masking rule defining how data is hidden.
type MaskingRule struct {
	// The type of the masking rule.
	MaskingRuleType string `json:"maskingRuleType"`
	// The selector for the element or the attribute to be masked.   Specify a CSS expression for an element or a [regular expression](https://dt-url.net/k9e0iaq) for an attribute.
	Selector string `json:"selector"`
	// Interactions with the element are (`true`) or are not (`false) masked.
	UserInteractionHidden bool `json:"userInteractionHidden"`
}

// NewMaskingRule instantiates a new MaskingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskingRule(maskingRuleType string, selector string, userInteractionHidden bool) *MaskingRule {
	this := MaskingRule{}
	this.MaskingRuleType = maskingRuleType
	this.Selector = selector
	this.UserInteractionHidden = userInteractionHidden
	return &this
}

// NewMaskingRuleWithDefaults instantiates a new MaskingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskingRuleWithDefaults() *MaskingRule {
	this := MaskingRule{}
	return &this
}

// GetMaskingRuleType returns the MaskingRuleType field value
func (o *MaskingRule) GetMaskingRuleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaskingRuleType
}

// GetMaskingRuleTypeOk returns a tuple with the MaskingRuleType field value
// and a boolean to check if the value has been set.
func (o *MaskingRule) GetMaskingRuleTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskingRuleType, true
}

// SetMaskingRuleType sets field value
func (o *MaskingRule) SetMaskingRuleType(v string) {
	o.MaskingRuleType = v
}

// GetSelector returns the Selector field value
func (o *MaskingRule) GetSelector() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value
// and a boolean to check if the value has been set.
func (o *MaskingRule) GetSelectorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Selector, true
}

// SetSelector sets field value
func (o *MaskingRule) SetSelector(v string) {
	o.Selector = v
}

// GetUserInteractionHidden returns the UserInteractionHidden field value
func (o *MaskingRule) GetUserInteractionHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UserInteractionHidden
}

// GetUserInteractionHiddenOk returns a tuple with the UserInteractionHidden field value
// and a boolean to check if the value has been set.
func (o *MaskingRule) GetUserInteractionHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserInteractionHidden, true
}

// SetUserInteractionHidden sets field value
func (o *MaskingRule) SetUserInteractionHidden(v bool) {
	o.UserInteractionHidden = v
}

func (o MaskingRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maskingRuleType"] = o.MaskingRuleType
	toSerialize["selector"] = o.Selector
	toSerialize["userInteractionHidden"] = o.UserInteractionHidden
	return toSerialize, nil
}

type NullableMaskingRule struct {
	value *MaskingRule
	isSet bool
}

func (v NullableMaskingRule) Get() *MaskingRule {
	return v.value
}

func (v *NullableMaskingRule) Set(val *MaskingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskingRule(val *MaskingRule) *NullableMaskingRule {
	return &NullableMaskingRule{value: val, isSet: true}
}

func (v NullableMaskingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


