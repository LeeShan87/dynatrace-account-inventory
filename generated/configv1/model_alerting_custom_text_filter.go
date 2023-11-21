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

// checks if the AlertingCustomTextFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertingCustomTextFilter{}

// AlertingCustomTextFilter Configuration of a matching filter.
type AlertingCustomTextFilter struct {
	// The condition is case sensitive (`false`) or case insensitive (`true`).    If not set, then `false` is used, making the condition case sensitive.
	CaseInsensitive bool `json:"caseInsensitive"`
	// The filter is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// Reverses the comparison **operator**. For example it turns the **begins with** into **does not begin with**.
	Negate bool `json:"negate"`
	// Operator of the comparison.    You can reverse it by setting **negate** to `true`.
	Operator string `json:"operator"`
	// The value to compare to.
	Value string `json:"value"`
}

// NewAlertingCustomTextFilter instantiates a new AlertingCustomTextFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertingCustomTextFilter(caseInsensitive bool, enabled bool, negate bool, operator string, value string) *AlertingCustomTextFilter {
	this := AlertingCustomTextFilter{}
	this.CaseInsensitive = caseInsensitive
	this.Enabled = enabled
	this.Negate = negate
	this.Operator = operator
	this.Value = value
	return &this
}

// NewAlertingCustomTextFilterWithDefaults instantiates a new AlertingCustomTextFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertingCustomTextFilterWithDefaults() *AlertingCustomTextFilter {
	this := AlertingCustomTextFilter{}
	return &this
}

// GetCaseInsensitive returns the CaseInsensitive field value
func (o *AlertingCustomTextFilter) GetCaseInsensitive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CaseInsensitive
}

// GetCaseInsensitiveOk returns a tuple with the CaseInsensitive field value
// and a boolean to check if the value has been set.
func (o *AlertingCustomTextFilter) GetCaseInsensitiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CaseInsensitive, true
}

// SetCaseInsensitive sets field value
func (o *AlertingCustomTextFilter) SetCaseInsensitive(v bool) {
	o.CaseInsensitive = v
}

// GetEnabled returns the Enabled field value
func (o *AlertingCustomTextFilter) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AlertingCustomTextFilter) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AlertingCustomTextFilter) SetEnabled(v bool) {
	o.Enabled = v
}

// GetNegate returns the Negate field value
func (o *AlertingCustomTextFilter) GetNegate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Negate
}

// GetNegateOk returns a tuple with the Negate field value
// and a boolean to check if the value has been set.
func (o *AlertingCustomTextFilter) GetNegateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Negate, true
}

// SetNegate sets field value
func (o *AlertingCustomTextFilter) SetNegate(v bool) {
	o.Negate = v
}

// GetOperator returns the Operator field value
func (o *AlertingCustomTextFilter) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *AlertingCustomTextFilter) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *AlertingCustomTextFilter) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value
func (o *AlertingCustomTextFilter) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AlertingCustomTextFilter) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AlertingCustomTextFilter) SetValue(v string) {
	o.Value = v
}

func (o AlertingCustomTextFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertingCustomTextFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["caseInsensitive"] = o.CaseInsensitive
	toSerialize["enabled"] = o.Enabled
	toSerialize["negate"] = o.Negate
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableAlertingCustomTextFilter struct {
	value *AlertingCustomTextFilter
	isSet bool
}

func (v NullableAlertingCustomTextFilter) Get() *AlertingCustomTextFilter {
	return v.value
}

func (v *NullableAlertingCustomTextFilter) Set(val *AlertingCustomTextFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertingCustomTextFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertingCustomTextFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertingCustomTextFilter(val *AlertingCustomTextFilter) *NullableAlertingCustomTextFilter {
	return &NullableAlertingCustomTextFilter{value: val, isSet: true}
}

func (v NullableAlertingCustomTextFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertingCustomTextFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


