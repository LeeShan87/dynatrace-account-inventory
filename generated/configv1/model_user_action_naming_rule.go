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

// checks if the UserActionNamingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserActionNamingRule{}

// UserActionNamingRule The settings of naming rule.
type UserActionNamingRule struct {
	// Defines the conditions when the naming rule should apply.
	Conditions []UserActionNamingRuleCondition `json:"conditions,omitempty"`
	// Naming pattern. Use Curly brackets `{}` to select placeholders.
	Template string `json:"template"`
	// If set to `true` the conditions will be connected by logical OR instead of logical AND.
	UseOrConditions *bool `json:"useOrConditions,omitempty"`
}

// NewUserActionNamingRule instantiates a new UserActionNamingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActionNamingRule(template string) *UserActionNamingRule {
	this := UserActionNamingRule{}
	this.Template = template
	return &this
}

// NewUserActionNamingRuleWithDefaults instantiates a new UserActionNamingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActionNamingRuleWithDefaults() *UserActionNamingRule {
	this := UserActionNamingRule{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *UserActionNamingRule) GetConditions() []UserActionNamingRuleCondition {
	if o == nil || IsNil(o.Conditions) {
		var ret []UserActionNamingRuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionNamingRule) GetConditionsOk() ([]UserActionNamingRuleCondition, bool) {
	if o == nil || IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *UserActionNamingRule) HasConditions() bool {
	if o != nil && !IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []UserActionNamingRuleCondition and assigns it to the Conditions field.
func (o *UserActionNamingRule) SetConditions(v []UserActionNamingRuleCondition) {
	o.Conditions = v
}

// GetTemplate returns the Template field value
func (o *UserActionNamingRule) GetTemplate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *UserActionNamingRule) GetTemplateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *UserActionNamingRule) SetTemplate(v string) {
	o.Template = v
}

// GetUseOrConditions returns the UseOrConditions field value if set, zero value otherwise.
func (o *UserActionNamingRule) GetUseOrConditions() bool {
	if o == nil || IsNil(o.UseOrConditions) {
		var ret bool
		return ret
	}
	return *o.UseOrConditions
}

// GetUseOrConditionsOk returns a tuple with the UseOrConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionNamingRule) GetUseOrConditionsOk() (*bool, bool) {
	if o == nil || IsNil(o.UseOrConditions) {
		return nil, false
	}
	return o.UseOrConditions, true
}

// HasUseOrConditions returns a boolean if a field has been set.
func (o *UserActionNamingRule) HasUseOrConditions() bool {
	if o != nil && !IsNil(o.UseOrConditions) {
		return true
	}

	return false
}

// SetUseOrConditions gets a reference to the given bool and assigns it to the UseOrConditions field.
func (o *UserActionNamingRule) SetUseOrConditions(v bool) {
	o.UseOrConditions = &v
}

func (o UserActionNamingRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserActionNamingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Conditions) {
		toSerialize["conditions"] = o.Conditions
	}
	toSerialize["template"] = o.Template
	if !IsNil(o.UseOrConditions) {
		toSerialize["useOrConditions"] = o.UseOrConditions
	}
	return toSerialize, nil
}

type NullableUserActionNamingRule struct {
	value *UserActionNamingRule
	isSet bool
}

func (v NullableUserActionNamingRule) Get() *UserActionNamingRule {
	return v.value
}

func (v *NullableUserActionNamingRule) Set(val *UserActionNamingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActionNamingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActionNamingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActionNamingRule(val *UserActionNamingRule) *NullableUserActionNamingRule {
	return &NullableUserActionNamingRule{value: val, isSet: true}
}

func (v NullableUserActionNamingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActionNamingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


