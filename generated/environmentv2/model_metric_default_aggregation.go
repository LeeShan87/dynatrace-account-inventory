/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
)

// checks if the MetricDefaultAggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricDefaultAggregation{}

// MetricDefaultAggregation The default aggregation of a metric.
type MetricDefaultAggregation struct {
	// The percentile to be delivered. Valid values are between `0` and `100`.   Applicable only to the `percentile` aggregation type.
	Parameter *float64 `json:"parameter,omitempty"`
	// The type of default aggregation.
	Type string `json:"type"`
}

// NewMetricDefaultAggregation instantiates a new MetricDefaultAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDefaultAggregation(type_ string) *MetricDefaultAggregation {
	this := MetricDefaultAggregation{}
	this.Type = type_
	return &this
}

// NewMetricDefaultAggregationWithDefaults instantiates a new MetricDefaultAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDefaultAggregationWithDefaults() *MetricDefaultAggregation {
	this := MetricDefaultAggregation{}
	return &this
}

// GetParameter returns the Parameter field value if set, zero value otherwise.
func (o *MetricDefaultAggregation) GetParameter() float64 {
	if o == nil || IsNil(o.Parameter) {
		var ret float64
		return ret
	}
	return *o.Parameter
}

// GetParameterOk returns a tuple with the Parameter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDefaultAggregation) GetParameterOk() (*float64, bool) {
	if o == nil || IsNil(o.Parameter) {
		return nil, false
	}
	return o.Parameter, true
}

// HasParameter returns a boolean if a field has been set.
func (o *MetricDefaultAggregation) HasParameter() bool {
	if o != nil && !IsNil(o.Parameter) {
		return true
	}

	return false
}

// SetParameter gets a reference to the given float64 and assigns it to the Parameter field.
func (o *MetricDefaultAggregation) SetParameter(v float64) {
	o.Parameter = &v
}

// GetType returns the Type field value
func (o *MetricDefaultAggregation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MetricDefaultAggregation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MetricDefaultAggregation) SetType(v string) {
	o.Type = v
}

func (o MetricDefaultAggregation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricDefaultAggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Parameter) {
		toSerialize["parameter"] = o.Parameter
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableMetricDefaultAggregation struct {
	value *MetricDefaultAggregation
	isSet bool
}

func (v NullableMetricDefaultAggregation) Get() *MetricDefaultAggregation {
	return v.value
}

func (v *NullableMetricDefaultAggregation) Set(val *MetricDefaultAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricDefaultAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricDefaultAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricDefaultAggregation(val *MetricDefaultAggregation) *NullableMetricDefaultAggregation {
	return &NullableMetricDefaultAggregation{value: val, isSet: true}
}

func (v NullableMetricDefaultAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricDefaultAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


