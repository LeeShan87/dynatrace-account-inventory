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

// checks if the JavaScriptInjectionRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JavaScriptInjectionRules{}

// JavaScriptInjectionRules Rules for javascript injection
type JavaScriptInjectionRules struct {
	// The enable or disable rule of the java script injection.
	Enabled bool `json:"enabled"`
	// The html pattern of the java script injection.
	HtmlPattern *string `json:"htmlPattern,omitempty"`
	// The url rule of the java script injection.
	Rule string `json:"rule"`
	// The target against which the rule of the java script injection should be matched.
	Target *string `json:"target,omitempty"`
	// The url operator of the java script injection.
	UrlOperator string `json:"urlOperator"`
	// The url pattern of the java script injection.
	UrlPattern *string `json:"urlPattern,omitempty"`
}

// NewJavaScriptInjectionRules instantiates a new JavaScriptInjectionRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJavaScriptInjectionRules(enabled bool, rule string, urlOperator string) *JavaScriptInjectionRules {
	this := JavaScriptInjectionRules{}
	this.Enabled = enabled
	this.Rule = rule
	this.UrlOperator = urlOperator
	return &this
}

// NewJavaScriptInjectionRulesWithDefaults instantiates a new JavaScriptInjectionRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJavaScriptInjectionRulesWithDefaults() *JavaScriptInjectionRules {
	this := JavaScriptInjectionRules{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *JavaScriptInjectionRules) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *JavaScriptInjectionRules) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *JavaScriptInjectionRules) SetEnabled(v bool) {
	o.Enabled = v
}

// GetHtmlPattern returns the HtmlPattern field value if set, zero value otherwise.
func (o *JavaScriptInjectionRules) GetHtmlPattern() string {
	if o == nil || IsNil(o.HtmlPattern) {
		var ret string
		return ret
	}
	return *o.HtmlPattern
}

// GetHtmlPatternOk returns a tuple with the HtmlPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaScriptInjectionRules) GetHtmlPatternOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlPattern) {
		return nil, false
	}
	return o.HtmlPattern, true
}

// HasHtmlPattern returns a boolean if a field has been set.
func (o *JavaScriptInjectionRules) HasHtmlPattern() bool {
	if o != nil && !IsNil(o.HtmlPattern) {
		return true
	}

	return false
}

// SetHtmlPattern gets a reference to the given string and assigns it to the HtmlPattern field.
func (o *JavaScriptInjectionRules) SetHtmlPattern(v string) {
	o.HtmlPattern = &v
}

// GetRule returns the Rule field value
func (o *JavaScriptInjectionRules) GetRule() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rule
}

// GetRuleOk returns a tuple with the Rule field value
// and a boolean to check if the value has been set.
func (o *JavaScriptInjectionRules) GetRuleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rule, true
}

// SetRule sets field value
func (o *JavaScriptInjectionRules) SetRule(v string) {
	o.Rule = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *JavaScriptInjectionRules) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaScriptInjectionRules) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *JavaScriptInjectionRules) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *JavaScriptInjectionRules) SetTarget(v string) {
	o.Target = &v
}

// GetUrlOperator returns the UrlOperator field value
func (o *JavaScriptInjectionRules) GetUrlOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlOperator
}

// GetUrlOperatorOk returns a tuple with the UrlOperator field value
// and a boolean to check if the value has been set.
func (o *JavaScriptInjectionRules) GetUrlOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlOperator, true
}

// SetUrlOperator sets field value
func (o *JavaScriptInjectionRules) SetUrlOperator(v string) {
	o.UrlOperator = v
}

// GetUrlPattern returns the UrlPattern field value if set, zero value otherwise.
func (o *JavaScriptInjectionRules) GetUrlPattern() string {
	if o == nil || IsNil(o.UrlPattern) {
		var ret string
		return ret
	}
	return *o.UrlPattern
}

// GetUrlPatternOk returns a tuple with the UrlPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JavaScriptInjectionRules) GetUrlPatternOk() (*string, bool) {
	if o == nil || IsNil(o.UrlPattern) {
		return nil, false
	}
	return o.UrlPattern, true
}

// HasUrlPattern returns a boolean if a field has been set.
func (o *JavaScriptInjectionRules) HasUrlPattern() bool {
	if o != nil && !IsNil(o.UrlPattern) {
		return true
	}

	return false
}

// SetUrlPattern gets a reference to the given string and assigns it to the UrlPattern field.
func (o *JavaScriptInjectionRules) SetUrlPattern(v string) {
	o.UrlPattern = &v
}

func (o JavaScriptInjectionRules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JavaScriptInjectionRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.HtmlPattern) {
		toSerialize["htmlPattern"] = o.HtmlPattern
	}
	toSerialize["rule"] = o.Rule
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	toSerialize["urlOperator"] = o.UrlOperator
	if !IsNil(o.UrlPattern) {
		toSerialize["urlPattern"] = o.UrlPattern
	}
	return toSerialize, nil
}

type NullableJavaScriptInjectionRules struct {
	value *JavaScriptInjectionRules
	isSet bool
}

func (v NullableJavaScriptInjectionRules) Get() *JavaScriptInjectionRules {
	return v.value
}

func (v *NullableJavaScriptInjectionRules) Set(val *JavaScriptInjectionRules) {
	v.value = val
	v.isSet = true
}

func (v NullableJavaScriptInjectionRules) IsSet() bool {
	return v.isSet
}

func (v *NullableJavaScriptInjectionRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJavaScriptInjectionRules(val *JavaScriptInjectionRules) *NullableJavaScriptInjectionRules {
	return &NullableJavaScriptInjectionRules{value: val, isSet: true}
}

func (v NullableJavaScriptInjectionRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJavaScriptInjectionRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

