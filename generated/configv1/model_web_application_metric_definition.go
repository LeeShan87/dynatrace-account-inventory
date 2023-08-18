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

// checks if the WebApplicationMetricDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebApplicationMetricDefinition{}

// WebApplicationMetricDefinition Definition of the web application metric.
type WebApplicationMetricDefinition struct {
	// The type of the web application metric.
	Metric string `json:"metric"`
	// The key of the user action property.    Only applicable for `DoubleProperty` and `LongProperty` metrics.
	PropertyKey *string `json:"propertyKey,omitempty"`
}

// NewWebApplicationMetricDefinition instantiates a new WebApplicationMetricDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebApplicationMetricDefinition(metric string) *WebApplicationMetricDefinition {
	this := WebApplicationMetricDefinition{}
	this.Metric = metric
	return &this
}

// NewWebApplicationMetricDefinitionWithDefaults instantiates a new WebApplicationMetricDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebApplicationMetricDefinitionWithDefaults() *WebApplicationMetricDefinition {
	this := WebApplicationMetricDefinition{}
	return &this
}

// GetMetric returns the Metric field value
func (o *WebApplicationMetricDefinition) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *WebApplicationMetricDefinition) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *WebApplicationMetricDefinition) SetMetric(v string) {
	o.Metric = v
}

// GetPropertyKey returns the PropertyKey field value if set, zero value otherwise.
func (o *WebApplicationMetricDefinition) GetPropertyKey() string {
	if o == nil || IsNil(o.PropertyKey) {
		var ret string
		return ret
	}
	return *o.PropertyKey
}

// GetPropertyKeyOk returns a tuple with the PropertyKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebApplicationMetricDefinition) GetPropertyKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PropertyKey) {
		return nil, false
	}
	return o.PropertyKey, true
}

// HasPropertyKey returns a boolean if a field has been set.
func (o *WebApplicationMetricDefinition) HasPropertyKey() bool {
	if o != nil && !IsNil(o.PropertyKey) {
		return true
	}

	return false
}

// SetPropertyKey gets a reference to the given string and assigns it to the PropertyKey field.
func (o *WebApplicationMetricDefinition) SetPropertyKey(v string) {
	o.PropertyKey = &v
}

func (o WebApplicationMetricDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebApplicationMetricDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metric"] = o.Metric
	if !IsNil(o.PropertyKey) {
		toSerialize["propertyKey"] = o.PropertyKey
	}
	return toSerialize, nil
}

type NullableWebApplicationMetricDefinition struct {
	value *WebApplicationMetricDefinition
	isSet bool
}

func (v NullableWebApplicationMetricDefinition) Get() *WebApplicationMetricDefinition {
	return v.value
}

func (v *NullableWebApplicationMetricDefinition) Set(val *WebApplicationMetricDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableWebApplicationMetricDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableWebApplicationMetricDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebApplicationMetricDefinition(val *WebApplicationMetricDefinition) *NullableWebApplicationMetricDefinition {
	return &NullableWebApplicationMetricDefinition{value: val, isSet: true}
}

func (v NullableWebApplicationMetricDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebApplicationMetricDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

