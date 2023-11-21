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

// checks if the MetricSeriesCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricSeriesCollection{}

// MetricSeriesCollection Data points of a metric.
type MetricSeriesCollection struct {
	// A list of filtered metric keys along with filters that have been applied to these keys, from the `optionalFilter` parameter.
	AppliedOptionalFilters []AppliedFilter `json:"appliedOptionalFilters,omitempty"`
	// Data points of the metric.
	Data []MetricSeries `json:"data"`
	// The ratio of queried data points divided by the maximum number of data points per metric that are allowed in a single query.
	DataPointCountRatio float32 `json:"dataPointCountRatio"`
	// The ratio of queried dimension tuples divided by the maximum number of dimension tuples allowed in a single query.
	DimensionCountRatio float32 `json:"dimensionCountRatio"`
	// The key of the metric.   If any transformation is applied, it is included here.
	MetricId string `json:"metricId"`
	// A list of potential warnings that affect this ID. For example deprecated feature usage etc.
	Warnings []string `json:"warnings,omitempty"`
}

// NewMetricSeriesCollection instantiates a new MetricSeriesCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricSeriesCollection(data []MetricSeries, dataPointCountRatio float32, dimensionCountRatio float32, metricId string) *MetricSeriesCollection {
	this := MetricSeriesCollection{}
	this.Data = data
	this.DataPointCountRatio = dataPointCountRatio
	this.DimensionCountRatio = dimensionCountRatio
	this.MetricId = metricId
	return &this
}

// NewMetricSeriesCollectionWithDefaults instantiates a new MetricSeriesCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricSeriesCollectionWithDefaults() *MetricSeriesCollection {
	this := MetricSeriesCollection{}
	return &this
}

// GetAppliedOptionalFilters returns the AppliedOptionalFilters field value if set, zero value otherwise.
func (o *MetricSeriesCollection) GetAppliedOptionalFilters() []AppliedFilter {
	if o == nil || IsNil(o.AppliedOptionalFilters) {
		var ret []AppliedFilter
		return ret
	}
	return o.AppliedOptionalFilters
}

// GetAppliedOptionalFiltersOk returns a tuple with the AppliedOptionalFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSeriesCollection) GetAppliedOptionalFiltersOk() ([]AppliedFilter, bool) {
	if o == nil || IsNil(o.AppliedOptionalFilters) {
		return nil, false
	}
	return o.AppliedOptionalFilters, true
}

// HasAppliedOptionalFilters returns a boolean if a field has been set.
func (o *MetricSeriesCollection) HasAppliedOptionalFilters() bool {
	if o != nil && !IsNil(o.AppliedOptionalFilters) {
		return true
	}

	return false
}

// SetAppliedOptionalFilters gets a reference to the given []AppliedFilter and assigns it to the AppliedOptionalFilters field.
func (o *MetricSeriesCollection) SetAppliedOptionalFilters(v []AppliedFilter) {
	o.AppliedOptionalFilters = v
}

// GetData returns the Data field value
func (o *MetricSeriesCollection) GetData() []MetricSeries {
	if o == nil {
		var ret []MetricSeries
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MetricSeriesCollection) GetDataOk() ([]MetricSeries, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *MetricSeriesCollection) SetData(v []MetricSeries) {
	o.Data = v
}

// GetDataPointCountRatio returns the DataPointCountRatio field value
func (o *MetricSeriesCollection) GetDataPointCountRatio() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DataPointCountRatio
}

// GetDataPointCountRatioOk returns a tuple with the DataPointCountRatio field value
// and a boolean to check if the value has been set.
func (o *MetricSeriesCollection) GetDataPointCountRatioOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataPointCountRatio, true
}

// SetDataPointCountRatio sets field value
func (o *MetricSeriesCollection) SetDataPointCountRatio(v float32) {
	o.DataPointCountRatio = v
}

// GetDimensionCountRatio returns the DimensionCountRatio field value
func (o *MetricSeriesCollection) GetDimensionCountRatio() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DimensionCountRatio
}

// GetDimensionCountRatioOk returns a tuple with the DimensionCountRatio field value
// and a boolean to check if the value has been set.
func (o *MetricSeriesCollection) GetDimensionCountRatioOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DimensionCountRatio, true
}

// SetDimensionCountRatio sets field value
func (o *MetricSeriesCollection) SetDimensionCountRatio(v float32) {
	o.DimensionCountRatio = v
}

// GetMetricId returns the MetricId field value
func (o *MetricSeriesCollection) GetMetricId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *MetricSeriesCollection) GetMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value
func (o *MetricSeriesCollection) SetMetricId(v string) {
	o.MetricId = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *MetricSeriesCollection) GetWarnings() []string {
	if o == nil || IsNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricSeriesCollection) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *MetricSeriesCollection) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *MetricSeriesCollection) SetWarnings(v []string) {
	o.Warnings = v
}

func (o MetricSeriesCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricSeriesCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppliedOptionalFilters) {
		toSerialize["appliedOptionalFilters"] = o.AppliedOptionalFilters
	}
	toSerialize["data"] = o.Data
	toSerialize["dataPointCountRatio"] = o.DataPointCountRatio
	toSerialize["dimensionCountRatio"] = o.DimensionCountRatio
	toSerialize["metricId"] = o.MetricId
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableMetricSeriesCollection struct {
	value *MetricSeriesCollection
	isSet bool
}

func (v NullableMetricSeriesCollection) Get() *MetricSeriesCollection {
	return v.value
}

func (v *NullableMetricSeriesCollection) Set(val *MetricSeriesCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricSeriesCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricSeriesCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricSeriesCollection(val *MetricSeriesCollection) *NullableMetricSeriesCollection {
	return &NullableMetricSeriesCollection{value: val, isSet: true}
}

func (v NullableMetricSeriesCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricSeriesCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


