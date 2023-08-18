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

// checks if the ServiceTypeComparison type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTypeComparison{}

// ServiceTypeComparison Comparison for `SERVICE_TYPE` attributes.
type ServiceTypeComparison struct {
	// Operator of the comparison. You can reverse it by setting **negate** to `true`.   Possible values depend on the **type** of the comparison. Find the list of actual models in the description of the **type** field and check the description of the model you need.
	Operator string `json:"operator"`
	// The value to compare to.
	Value *string `json:"value,omitempty"`
}

// NewServiceTypeComparison instantiates a new ServiceTypeComparison object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTypeComparison(operator string, negate bool, type_ string) *ServiceTypeComparison {
	this := ServiceTypeComparison{}
	this.Negate = negate
	this.Operator = operator
	this.Type = type_
	return &this
}

// NewServiceTypeComparisonWithDefaults instantiates a new ServiceTypeComparison object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTypeComparisonWithDefaults() *ServiceTypeComparison {
	this := ServiceTypeComparison{}
	return &this
}

// GetOperator returns the Operator field value
func (o *ServiceTypeComparison) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *ServiceTypeComparison) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *ServiceTypeComparison) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ServiceTypeComparison) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTypeComparison) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ServiceTypeComparison) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ServiceTypeComparison) SetValue(v string) {
	o.Value = &v
}

func (o ServiceTypeComparison) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTypeComparison) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operator"] = o.Operator
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableServiceTypeComparison struct {
	value *ServiceTypeComparison
	isSet bool
}

func (v NullableServiceTypeComparison) Get() *ServiceTypeComparison {
	return v.value
}

func (v *NullableServiceTypeComparison) Set(val *ServiceTypeComparison) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTypeComparison) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTypeComparison) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTypeComparison(val *ServiceTypeComparison) *NullableServiceTypeComparison {
	return &NullableServiceTypeComparison{value: val, isSet: true}
}

func (v NullableServiceTypeComparison) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTypeComparison) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


