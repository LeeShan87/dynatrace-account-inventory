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

// checks if the ApplicationErrorRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationErrorRules{}

// ApplicationErrorRules Configuration of error rules in the web application.
type ApplicationErrorRules struct {
	// An ordered list of custom errors.   Rules are evaluated from top to bottom; the first matching rule applies.
	CustomErrorRules []CustomErrorRule `json:"customErrorRules"`
	// An ordered list of HTTP errors.   Rules are evaluated from top to bottom; the first matching rule applies.
	HttpErrorRules []HttpErrorRule `json:"httpErrorRules"`
	// Exclude (`true`) or include (`false`) custom errors listed in **customErrorRules** in Apdex calculation.
	IgnoreCustomErrorsInApdexCalculation bool `json:"ignoreCustomErrorsInApdexCalculation"`
	// Exclude (`true`) or include (`false`) HTTP errors listed in **httpErrorRules** in Apdex calculation.
	IgnoreHttpErrorsInApdexCalculation bool `json:"ignoreHttpErrorsInApdexCalculation"`
	// Exclude (`true`) or include (`false`) JavaScript errors in Apdex calculation.
	IgnoreJavaScriptErrorsInApdexCalculation bool `json:"ignoreJavaScriptErrorsInApdexCalculation"`
}

// NewApplicationErrorRules instantiates a new ApplicationErrorRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationErrorRules(customErrorRules []CustomErrorRule, httpErrorRules []HttpErrorRule, ignoreCustomErrorsInApdexCalculation bool, ignoreHttpErrorsInApdexCalculation bool, ignoreJavaScriptErrorsInApdexCalculation bool) *ApplicationErrorRules {
	this := ApplicationErrorRules{}
	this.CustomErrorRules = customErrorRules
	this.HttpErrorRules = httpErrorRules
	this.IgnoreCustomErrorsInApdexCalculation = ignoreCustomErrorsInApdexCalculation
	this.IgnoreHttpErrorsInApdexCalculation = ignoreHttpErrorsInApdexCalculation
	this.IgnoreJavaScriptErrorsInApdexCalculation = ignoreJavaScriptErrorsInApdexCalculation
	return &this
}

// NewApplicationErrorRulesWithDefaults instantiates a new ApplicationErrorRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationErrorRulesWithDefaults() *ApplicationErrorRules {
	this := ApplicationErrorRules{}
	return &this
}

// GetCustomErrorRules returns the CustomErrorRules field value
func (o *ApplicationErrorRules) GetCustomErrorRules() []CustomErrorRule {
	if o == nil {
		var ret []CustomErrorRule
		return ret
	}

	return o.CustomErrorRules
}

// GetCustomErrorRulesOk returns a tuple with the CustomErrorRules field value
// and a boolean to check if the value has been set.
func (o *ApplicationErrorRules) GetCustomErrorRulesOk() ([]CustomErrorRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomErrorRules, true
}

// SetCustomErrorRules sets field value
func (o *ApplicationErrorRules) SetCustomErrorRules(v []CustomErrorRule) {
	o.CustomErrorRules = v
}

// GetHttpErrorRules returns the HttpErrorRules field value
func (o *ApplicationErrorRules) GetHttpErrorRules() []HttpErrorRule {
	if o == nil {
		var ret []HttpErrorRule
		return ret
	}

	return o.HttpErrorRules
}

// GetHttpErrorRulesOk returns a tuple with the HttpErrorRules field value
// and a boolean to check if the value has been set.
func (o *ApplicationErrorRules) GetHttpErrorRulesOk() ([]HttpErrorRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.HttpErrorRules, true
}

// SetHttpErrorRules sets field value
func (o *ApplicationErrorRules) SetHttpErrorRules(v []HttpErrorRule) {
	o.HttpErrorRules = v
}

// GetIgnoreCustomErrorsInApdexCalculation returns the IgnoreCustomErrorsInApdexCalculation field value
func (o *ApplicationErrorRules) GetIgnoreCustomErrorsInApdexCalculation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgnoreCustomErrorsInApdexCalculation
}

// GetIgnoreCustomErrorsInApdexCalculationOk returns a tuple with the IgnoreCustomErrorsInApdexCalculation field value
// and a boolean to check if the value has been set.
func (o *ApplicationErrorRules) GetIgnoreCustomErrorsInApdexCalculationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnoreCustomErrorsInApdexCalculation, true
}

// SetIgnoreCustomErrorsInApdexCalculation sets field value
func (o *ApplicationErrorRules) SetIgnoreCustomErrorsInApdexCalculation(v bool) {
	o.IgnoreCustomErrorsInApdexCalculation = v
}

// GetIgnoreHttpErrorsInApdexCalculation returns the IgnoreHttpErrorsInApdexCalculation field value
func (o *ApplicationErrorRules) GetIgnoreHttpErrorsInApdexCalculation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgnoreHttpErrorsInApdexCalculation
}

// GetIgnoreHttpErrorsInApdexCalculationOk returns a tuple with the IgnoreHttpErrorsInApdexCalculation field value
// and a boolean to check if the value has been set.
func (o *ApplicationErrorRules) GetIgnoreHttpErrorsInApdexCalculationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnoreHttpErrorsInApdexCalculation, true
}

// SetIgnoreHttpErrorsInApdexCalculation sets field value
func (o *ApplicationErrorRules) SetIgnoreHttpErrorsInApdexCalculation(v bool) {
	o.IgnoreHttpErrorsInApdexCalculation = v
}

// GetIgnoreJavaScriptErrorsInApdexCalculation returns the IgnoreJavaScriptErrorsInApdexCalculation field value
func (o *ApplicationErrorRules) GetIgnoreJavaScriptErrorsInApdexCalculation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IgnoreJavaScriptErrorsInApdexCalculation
}

// GetIgnoreJavaScriptErrorsInApdexCalculationOk returns a tuple with the IgnoreJavaScriptErrorsInApdexCalculation field value
// and a boolean to check if the value has been set.
func (o *ApplicationErrorRules) GetIgnoreJavaScriptErrorsInApdexCalculationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IgnoreJavaScriptErrorsInApdexCalculation, true
}

// SetIgnoreJavaScriptErrorsInApdexCalculation sets field value
func (o *ApplicationErrorRules) SetIgnoreJavaScriptErrorsInApdexCalculation(v bool) {
	o.IgnoreJavaScriptErrorsInApdexCalculation = v
}

func (o ApplicationErrorRules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationErrorRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["customErrorRules"] = o.CustomErrorRules
	toSerialize["httpErrorRules"] = o.HttpErrorRules
	toSerialize["ignoreCustomErrorsInApdexCalculation"] = o.IgnoreCustomErrorsInApdexCalculation
	toSerialize["ignoreHttpErrorsInApdexCalculation"] = o.IgnoreHttpErrorsInApdexCalculation
	toSerialize["ignoreJavaScriptErrorsInApdexCalculation"] = o.IgnoreJavaScriptErrorsInApdexCalculation
	return toSerialize, nil
}

type NullableApplicationErrorRules struct {
	value *ApplicationErrorRules
	isSet bool
}

func (v NullableApplicationErrorRules) Get() *ApplicationErrorRules {
	return v.value
}

func (v *NullableApplicationErrorRules) Set(val *ApplicationErrorRules) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationErrorRules) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationErrorRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationErrorRules(val *ApplicationErrorRules) *NullableApplicationErrorRules {
	return &NullableApplicationErrorRules{value: val, isSet: true}
}

func (v NullableApplicationErrorRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationErrorRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


