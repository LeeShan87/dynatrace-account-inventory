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

// checks if the MetricDimensionCardinality type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricDimensionCardinality{}

// MetricDimensionCardinality The dimension cardinalities of a metric.
type MetricDimensionCardinality struct {
	// The cardinality estimate of the dimension.
	Estimate int64 `json:"estimate"`
	// The key of the dimension.    It must be unique within the metric.
	Key string `json:"key"`
	// The relative cardinality of the dimension expressed as percentage
	Relative float64 `json:"relative"`
}

// NewMetricDimensionCardinality instantiates a new MetricDimensionCardinality object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDimensionCardinality(estimate int64, key string, relative float64) *MetricDimensionCardinality {
	this := MetricDimensionCardinality{}
	this.Estimate = estimate
	this.Key = key
	this.Relative = relative
	return &this
}

// NewMetricDimensionCardinalityWithDefaults instantiates a new MetricDimensionCardinality object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDimensionCardinalityWithDefaults() *MetricDimensionCardinality {
	this := MetricDimensionCardinality{}
	return &this
}

// GetEstimate returns the Estimate field value
func (o *MetricDimensionCardinality) GetEstimate() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Estimate
}

// GetEstimateOk returns a tuple with the Estimate field value
// and a boolean to check if the value has been set.
func (o *MetricDimensionCardinality) GetEstimateOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Estimate, true
}

// SetEstimate sets field value
func (o *MetricDimensionCardinality) SetEstimate(v int64) {
	o.Estimate = v
}

// GetKey returns the Key field value
func (o *MetricDimensionCardinality) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MetricDimensionCardinality) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MetricDimensionCardinality) SetKey(v string) {
	o.Key = v
}

// GetRelative returns the Relative field value
func (o *MetricDimensionCardinality) GetRelative() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Relative
}

// GetRelativeOk returns a tuple with the Relative field value
// and a boolean to check if the value has been set.
func (o *MetricDimensionCardinality) GetRelativeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relative, true
}

// SetRelative sets field value
func (o *MetricDimensionCardinality) SetRelative(v float64) {
	o.Relative = v
}

func (o MetricDimensionCardinality) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricDimensionCardinality) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["estimate"] = o.Estimate
	toSerialize["key"] = o.Key
	toSerialize["relative"] = o.Relative
	return toSerialize, nil
}

type NullableMetricDimensionCardinality struct {
	value *MetricDimensionCardinality
	isSet bool
}

func (v NullableMetricDimensionCardinality) Get() *MetricDimensionCardinality {
	return v.value
}

func (v *NullableMetricDimensionCardinality) Set(val *MetricDimensionCardinality) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricDimensionCardinality) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricDimensionCardinality) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricDimensionCardinality(val *MetricDimensionCardinality) *NullableMetricDimensionCardinality {
	return &NullableMetricDimensionCardinality{value: val, isSet: true}
}

func (v NullableMetricDimensionCardinality) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricDimensionCardinality) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


