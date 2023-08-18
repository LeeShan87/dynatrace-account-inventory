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

// checks if the CalculatedMetricDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CalculatedMetricDefinition{}

// CalculatedMetricDefinition The definition of a calculated service metric.
type CalculatedMetricDefinition struct {
	// The metric to be captured.
	Metric string `json:"metric"`
	// The request attribute to be captured.    Only applicable when the **metric** parameter is set to `REQUEST_ATTRIBUTE`.
	RequestAttribute *string `json:"requestAttribute,omitempty"`
}

// NewCalculatedMetricDefinition instantiates a new CalculatedMetricDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalculatedMetricDefinition(metric string) *CalculatedMetricDefinition {
	this := CalculatedMetricDefinition{}
	this.Metric = metric
	return &this
}

// NewCalculatedMetricDefinitionWithDefaults instantiates a new CalculatedMetricDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalculatedMetricDefinitionWithDefaults() *CalculatedMetricDefinition {
	this := CalculatedMetricDefinition{}
	return &this
}

// GetMetric returns the Metric field value
func (o *CalculatedMetricDefinition) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *CalculatedMetricDefinition) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *CalculatedMetricDefinition) SetMetric(v string) {
	o.Metric = v
}

// GetRequestAttribute returns the RequestAttribute field value if set, zero value otherwise.
func (o *CalculatedMetricDefinition) GetRequestAttribute() string {
	if o == nil || IsNil(o.RequestAttribute) {
		var ret string
		return ret
	}
	return *o.RequestAttribute
}

// GetRequestAttributeOk returns a tuple with the RequestAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalculatedMetricDefinition) GetRequestAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.RequestAttribute) {
		return nil, false
	}
	return o.RequestAttribute, true
}

// HasRequestAttribute returns a boolean if a field has been set.
func (o *CalculatedMetricDefinition) HasRequestAttribute() bool {
	if o != nil && !IsNil(o.RequestAttribute) {
		return true
	}

	return false
}

// SetRequestAttribute gets a reference to the given string and assigns it to the RequestAttribute field.
func (o *CalculatedMetricDefinition) SetRequestAttribute(v string) {
	o.RequestAttribute = &v
}

func (o CalculatedMetricDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CalculatedMetricDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metric"] = o.Metric
	if !IsNil(o.RequestAttribute) {
		toSerialize["requestAttribute"] = o.RequestAttribute
	}
	return toSerialize, nil
}

type NullableCalculatedMetricDefinition struct {
	value *CalculatedMetricDefinition
	isSet bool
}

func (v NullableCalculatedMetricDefinition) Get() *CalculatedMetricDefinition {
	return v.value
}

func (v *NullableCalculatedMetricDefinition) Set(val *CalculatedMetricDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableCalculatedMetricDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableCalculatedMetricDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalculatedMetricDefinition(val *CalculatedMetricDefinition) *NullableCalculatedMetricDefinition {
	return &NullableCalculatedMetricDefinition{value: val, isSet: true}
}

func (v NullableCalculatedMetricDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalculatedMetricDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


