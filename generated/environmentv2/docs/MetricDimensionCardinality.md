# MetricDimensionCardinality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Estimate** | **int64** | The cardinality estimate of the dimension. | 
**Key** | **string** | The key of the dimension.    It must be unique within the metric. | 
**Relative** | **float64** | The relative cardinality of the dimension expressed as percentage | 

## Methods

### NewMetricDimensionCardinality

`func NewMetricDimensionCardinality(estimate int64, key string, relative float64, ) *MetricDimensionCardinality`

NewMetricDimensionCardinality instantiates a new MetricDimensionCardinality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDimensionCardinalityWithDefaults

`func NewMetricDimensionCardinalityWithDefaults() *MetricDimensionCardinality`

NewMetricDimensionCardinalityWithDefaults instantiates a new MetricDimensionCardinality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimate

`func (o *MetricDimensionCardinality) GetEstimate() int64`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *MetricDimensionCardinality) GetEstimateOk() (*int64, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *MetricDimensionCardinality) SetEstimate(v int64)`

SetEstimate sets Estimate field to given value.


### GetKey

`func (o *MetricDimensionCardinality) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricDimensionCardinality) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricDimensionCardinality) SetKey(v string)`

SetKey sets Key field to given value.


### GetRelative

`func (o *MetricDimensionCardinality) GetRelative() float64`

GetRelative returns the Relative field if non-nil, zero value otherwise.

### GetRelativeOk

`func (o *MetricDimensionCardinality) GetRelativeOk() (*float64, bool)`

GetRelativeOk returns a tuple with the Relative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelative

`func (o *MetricDimensionCardinality) SetRelative(v float64)`

SetRelative sets Relative field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


