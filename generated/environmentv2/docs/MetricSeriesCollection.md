# MetricSeriesCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliedOptionalFilters** | Pointer to [**[]AppliedFilter**](AppliedFilter.md) | A list of filtered metric keys along with filters that have been applied to these keys, from the &#x60;optionalFilter&#x60; parameter. | [optional] 
**Data** | [**[]MetricSeries**](MetricSeries.md) | Data points of the metric. | 
**DataPointCountRatio** | **float32** | The ratio of queried data points divided by the maximum number of data points per metric that are allowed in a single query. | 
**DimensionCountRatio** | **float32** | The ratio of queried dimension tuples divided by the maximum number of dimension tuples allowed in a single query. | 
**MetricId** | **string** | The key of the metric.   If any transformation is applied, it is included here. | 
**Warnings** | Pointer to **[]string** | A list of potential warnings that affect this ID. For example deprecated feature usage etc. | [optional] 

## Methods

### NewMetricSeriesCollection

`func NewMetricSeriesCollection(data []MetricSeries, dataPointCountRatio float32, dimensionCountRatio float32, metricId string, ) *MetricSeriesCollection`

NewMetricSeriesCollection instantiates a new MetricSeriesCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricSeriesCollectionWithDefaults

`func NewMetricSeriesCollectionWithDefaults() *MetricSeriesCollection`

NewMetricSeriesCollectionWithDefaults instantiates a new MetricSeriesCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliedOptionalFilters

`func (o *MetricSeriesCollection) GetAppliedOptionalFilters() []AppliedFilter`

GetAppliedOptionalFilters returns the AppliedOptionalFilters field if non-nil, zero value otherwise.

### GetAppliedOptionalFiltersOk

`func (o *MetricSeriesCollection) GetAppliedOptionalFiltersOk() (*[]AppliedFilter, bool)`

GetAppliedOptionalFiltersOk returns a tuple with the AppliedOptionalFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliedOptionalFilters

`func (o *MetricSeriesCollection) SetAppliedOptionalFilters(v []AppliedFilter)`

SetAppliedOptionalFilters sets AppliedOptionalFilters field to given value.

### HasAppliedOptionalFilters

`func (o *MetricSeriesCollection) HasAppliedOptionalFilters() bool`

HasAppliedOptionalFilters returns a boolean if a field has been set.

### GetData

`func (o *MetricSeriesCollection) GetData() []MetricSeries`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MetricSeriesCollection) GetDataOk() (*[]MetricSeries, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MetricSeriesCollection) SetData(v []MetricSeries)`

SetData sets Data field to given value.


### GetDataPointCountRatio

`func (o *MetricSeriesCollection) GetDataPointCountRatio() float32`

GetDataPointCountRatio returns the DataPointCountRatio field if non-nil, zero value otherwise.

### GetDataPointCountRatioOk

`func (o *MetricSeriesCollection) GetDataPointCountRatioOk() (*float32, bool)`

GetDataPointCountRatioOk returns a tuple with the DataPointCountRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPointCountRatio

`func (o *MetricSeriesCollection) SetDataPointCountRatio(v float32)`

SetDataPointCountRatio sets DataPointCountRatio field to given value.


### GetDimensionCountRatio

`func (o *MetricSeriesCollection) GetDimensionCountRatio() float32`

GetDimensionCountRatio returns the DimensionCountRatio field if non-nil, zero value otherwise.

### GetDimensionCountRatioOk

`func (o *MetricSeriesCollection) GetDimensionCountRatioOk() (*float32, bool)`

GetDimensionCountRatioOk returns a tuple with the DimensionCountRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionCountRatio

`func (o *MetricSeriesCollection) SetDimensionCountRatio(v float32)`

SetDimensionCountRatio sets DimensionCountRatio field to given value.


### GetMetricId

`func (o *MetricSeriesCollection) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricSeriesCollection) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricSeriesCollection) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetWarnings

`func (o *MetricSeriesCollection) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MetricSeriesCollection) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MetricSeriesCollection) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *MetricSeriesCollection) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


