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

// checks if the FlawStateComparisonInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlawStateComparisonInfo{}

// FlawStateComparisonInfo Comparison for `FLAW_STATE` attributes.
type FlawStateComparisonInfo struct {
	// Operator of the comparision. You can reverse it by setting **negate** to `true`.
	Comparison string `json:"comparison"`
	// The value to compare to.
	Value *string `json:"value,omitempty"`
	// The values to compare to.
	Values *[]string `json:"values,omitempty"`
}

// NewFlawStateComparisonInfo instantiates a new FlawStateComparisonInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlawStateComparisonInfo(comparison string, negate bool, type_ string) *FlawStateComparisonInfo {
	this := FlawStateComparisonInfo{}
	this.Comparison = comparison
	this.Negate = negate
	this.Type = type_
	return &this
}

// NewFlawStateComparisonInfoWithDefaults instantiates a new FlawStateComparisonInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlawStateComparisonInfoWithDefaults() *FlawStateComparisonInfo {
	this := FlawStateComparisonInfo{}
	return &this
}

// GetComparison returns the Comparison field value
func (o *FlawStateComparisonInfo) GetComparison() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Comparison
}

// GetComparisonOk returns a tuple with the Comparison field value
// and a boolean to check if the value has been set.
func (o *FlawStateComparisonInfo) GetComparisonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Comparison, true
}

// SetComparison sets field value
func (o *FlawStateComparisonInfo) SetComparison(v string) {
	o.Comparison = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *FlawStateComparisonInfo) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlawStateComparisonInfo) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *FlawStateComparisonInfo) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *FlawStateComparisonInfo) SetValue(v string) {
	o.Value = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *FlawStateComparisonInfo) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlawStateComparisonInfo) GetValuesOk() (*[]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *FlawStateComparisonInfo) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *FlawStateComparisonInfo) SetValues(v []string) {
	o.Values = &v
}

func (o FlawStateComparisonInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlawStateComparisonInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["comparison"] = o.Comparison
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableFlawStateComparisonInfo struct {
	value *FlawStateComparisonInfo
	isSet bool
}

func (v NullableFlawStateComparisonInfo) Get() *FlawStateComparisonInfo {
	return v.value
}

func (v *NullableFlawStateComparisonInfo) Set(val *FlawStateComparisonInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFlawStateComparisonInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFlawStateComparisonInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlawStateComparisonInfo(val *FlawStateComparisonInfo) *NullableFlawStateComparisonInfo {
	return &NullableFlawStateComparisonInfo{value: val, isSet: true}
}

func (v NullableFlawStateComparisonInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlawStateComparisonInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


