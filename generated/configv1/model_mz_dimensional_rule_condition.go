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

// checks if the MzDimensionalRuleCondition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MzDimensionalRuleCondition{}

// MzDimensionalRuleCondition A condition of the management zone dimensional rule.
type MzDimensionalRuleCondition struct {
	// The type of the condition.
	ConditionType string `json:"conditionType"`
	// The reference value for comparison.   For conditions of the `DIMENSION` type, specify the key here.
	Key string `json:"key"`
	// How we compare the values.
	RuleMatcher string `json:"ruleMatcher"`
	// The value of the dimension.   Only applicable when the **conditionType** is set to `DIMENSION`.
	Value *string `json:"value,omitempty"`
}

// NewMzDimensionalRuleCondition instantiates a new MzDimensionalRuleCondition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMzDimensionalRuleCondition(conditionType string, key string, ruleMatcher string) *MzDimensionalRuleCondition {
	this := MzDimensionalRuleCondition{}
	this.ConditionType = conditionType
	this.Key = key
	this.RuleMatcher = ruleMatcher
	return &this
}

// NewMzDimensionalRuleConditionWithDefaults instantiates a new MzDimensionalRuleCondition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMzDimensionalRuleConditionWithDefaults() *MzDimensionalRuleCondition {
	this := MzDimensionalRuleCondition{}
	return &this
}

// GetConditionType returns the ConditionType field value
func (o *MzDimensionalRuleCondition) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *MzDimensionalRuleCondition) GetConditionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *MzDimensionalRuleCondition) SetConditionType(v string) {
	o.ConditionType = v
}

// GetKey returns the Key field value
func (o *MzDimensionalRuleCondition) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MzDimensionalRuleCondition) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MzDimensionalRuleCondition) SetKey(v string) {
	o.Key = v
}

// GetRuleMatcher returns the RuleMatcher field value
func (o *MzDimensionalRuleCondition) GetRuleMatcher() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleMatcher
}

// GetRuleMatcherOk returns a tuple with the RuleMatcher field value
// and a boolean to check if the value has been set.
func (o *MzDimensionalRuleCondition) GetRuleMatcherOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleMatcher, true
}

// SetRuleMatcher sets field value
func (o *MzDimensionalRuleCondition) SetRuleMatcher(v string) {
	o.RuleMatcher = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MzDimensionalRuleCondition) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MzDimensionalRuleCondition) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MzDimensionalRuleCondition) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MzDimensionalRuleCondition) SetValue(v string) {
	o.Value = &v
}

func (o MzDimensionalRuleCondition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MzDimensionalRuleCondition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conditionType"] = o.ConditionType
	toSerialize["key"] = o.Key
	toSerialize["ruleMatcher"] = o.RuleMatcher
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableMzDimensionalRuleCondition struct {
	value *MzDimensionalRuleCondition
	isSet bool
}

func (v NullableMzDimensionalRuleCondition) Get() *MzDimensionalRuleCondition {
	return v.value
}

func (v *NullableMzDimensionalRuleCondition) Set(val *MzDimensionalRuleCondition) {
	v.value = val
	v.isSet = true
}

func (v NullableMzDimensionalRuleCondition) IsSet() bool {
	return v.isSet
}

func (v *NullableMzDimensionalRuleCondition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMzDimensionalRuleCondition(val *MzDimensionalRuleCondition) *NullableMzDimensionalRuleCondition {
	return &NullableMzDimensionalRuleCondition{value: val, isSet: true}
}

func (v NullableMzDimensionalRuleCondition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMzDimensionalRuleCondition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


