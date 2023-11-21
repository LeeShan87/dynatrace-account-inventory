# DataExplorerQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultValue** | Pointer to **float64** | Replaces null data points with the provided value | [optional] 
**Enabled** | **bool** | Status of the query | 
**FilterBy** | [**DataExplorerFilter**](DataExplorerFilter.md) |  | 
**FoldTransformation** | Pointer to **string** | The fold transformation | [optional] 
**GeneratedMetricSelector** | Pointer to **string** | Generated metric selector | [optional] 
**Id** | **string** | The id of the query | 
**Limit** | Pointer to **int32** | Limit the results of the query | [optional] 
**Metric** | **string** | The metric id | 
**MetricSelector** | Pointer to **string** | The metric selector | [optional] 
**Rate** | Pointer to **string** | Converts a count-based metric (for example, bytes) into a rate-based metric (bytes per second) | [optional] 
**SortBy** | Pointer to **string** | The order of the sorting applied to the query | [optional] 
**SortByDimension** | Pointer to **string** | The dimension where sorting is applied. Sorting by value if null | [optional] 
**SpaceAggregation** | **string** | Space aggregation applied to the query | 
**SplitBy** | **[]string** | The splittings applied to the query | 
**TimeAggregation** | **string** | Time roll up applied to the query | 
**TimeShift** | Pointer to [**DataExplorerTimeShift**](DataExplorerTimeShift.md) |  | [optional] 

## Methods

### NewDataExplorerQuery

`func NewDataExplorerQuery(enabled bool, filterBy DataExplorerFilter, id string, metric string, spaceAggregation string, splitBy []string, timeAggregation string, ) *DataExplorerQuery`

NewDataExplorerQuery instantiates a new DataExplorerQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataExplorerQueryWithDefaults

`func NewDataExplorerQueryWithDefaults() *DataExplorerQuery`

NewDataExplorerQueryWithDefaults instantiates a new DataExplorerQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultValue

`func (o *DataExplorerQuery) GetDefaultValue() float64`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *DataExplorerQuery) GetDefaultValueOk() (*float64, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *DataExplorerQuery) SetDefaultValue(v float64)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *DataExplorerQuery) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetEnabled

`func (o *DataExplorerQuery) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DataExplorerQuery) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DataExplorerQuery) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFilterBy

`func (o *DataExplorerQuery) GetFilterBy() DataExplorerFilter`

GetFilterBy returns the FilterBy field if non-nil, zero value otherwise.

### GetFilterByOk

`func (o *DataExplorerQuery) GetFilterByOk() (*DataExplorerFilter, bool)`

GetFilterByOk returns a tuple with the FilterBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterBy

`func (o *DataExplorerQuery) SetFilterBy(v DataExplorerFilter)`

SetFilterBy sets FilterBy field to given value.


### GetFoldTransformation

`func (o *DataExplorerQuery) GetFoldTransformation() string`

GetFoldTransformation returns the FoldTransformation field if non-nil, zero value otherwise.

### GetFoldTransformationOk

`func (o *DataExplorerQuery) GetFoldTransformationOk() (*string, bool)`

GetFoldTransformationOk returns a tuple with the FoldTransformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoldTransformation

`func (o *DataExplorerQuery) SetFoldTransformation(v string)`

SetFoldTransformation sets FoldTransformation field to given value.

### HasFoldTransformation

`func (o *DataExplorerQuery) HasFoldTransformation() bool`

HasFoldTransformation returns a boolean if a field has been set.

### GetGeneratedMetricSelector

`func (o *DataExplorerQuery) GetGeneratedMetricSelector() string`

GetGeneratedMetricSelector returns the GeneratedMetricSelector field if non-nil, zero value otherwise.

### GetGeneratedMetricSelectorOk

`func (o *DataExplorerQuery) GetGeneratedMetricSelectorOk() (*string, bool)`

GetGeneratedMetricSelectorOk returns a tuple with the GeneratedMetricSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedMetricSelector

`func (o *DataExplorerQuery) SetGeneratedMetricSelector(v string)`

SetGeneratedMetricSelector sets GeneratedMetricSelector field to given value.

### HasGeneratedMetricSelector

`func (o *DataExplorerQuery) HasGeneratedMetricSelector() bool`

HasGeneratedMetricSelector returns a boolean if a field has been set.

### GetId

`func (o *DataExplorerQuery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataExplorerQuery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataExplorerQuery) SetId(v string)`

SetId sets Id field to given value.


### GetLimit

`func (o *DataExplorerQuery) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *DataExplorerQuery) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *DataExplorerQuery) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *DataExplorerQuery) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetMetric

`func (o *DataExplorerQuery) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *DataExplorerQuery) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *DataExplorerQuery) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetMetricSelector

`func (o *DataExplorerQuery) GetMetricSelector() string`

GetMetricSelector returns the MetricSelector field if non-nil, zero value otherwise.

### GetMetricSelectorOk

`func (o *DataExplorerQuery) GetMetricSelectorOk() (*string, bool)`

GetMetricSelectorOk returns a tuple with the MetricSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelector

`func (o *DataExplorerQuery) SetMetricSelector(v string)`

SetMetricSelector sets MetricSelector field to given value.

### HasMetricSelector

`func (o *DataExplorerQuery) HasMetricSelector() bool`

HasMetricSelector returns a boolean if a field has been set.

### GetRate

`func (o *DataExplorerQuery) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *DataExplorerQuery) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *DataExplorerQuery) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *DataExplorerQuery) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetSortBy

`func (o *DataExplorerQuery) GetSortBy() string`

GetSortBy returns the SortBy field if non-nil, zero value otherwise.

### GetSortByOk

`func (o *DataExplorerQuery) GetSortByOk() (*string, bool)`

GetSortByOk returns a tuple with the SortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortBy

`func (o *DataExplorerQuery) SetSortBy(v string)`

SetSortBy sets SortBy field to given value.

### HasSortBy

`func (o *DataExplorerQuery) HasSortBy() bool`

HasSortBy returns a boolean if a field has been set.

### GetSortByDimension

`func (o *DataExplorerQuery) GetSortByDimension() string`

GetSortByDimension returns the SortByDimension field if non-nil, zero value otherwise.

### GetSortByDimensionOk

`func (o *DataExplorerQuery) GetSortByDimensionOk() (*string, bool)`

GetSortByDimensionOk returns a tuple with the SortByDimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortByDimension

`func (o *DataExplorerQuery) SetSortByDimension(v string)`

SetSortByDimension sets SortByDimension field to given value.

### HasSortByDimension

`func (o *DataExplorerQuery) HasSortByDimension() bool`

HasSortByDimension returns a boolean if a field has been set.

### GetSpaceAggregation

`func (o *DataExplorerQuery) GetSpaceAggregation() string`

GetSpaceAggregation returns the SpaceAggregation field if non-nil, zero value otherwise.

### GetSpaceAggregationOk

`func (o *DataExplorerQuery) GetSpaceAggregationOk() (*string, bool)`

GetSpaceAggregationOk returns a tuple with the SpaceAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceAggregation

`func (o *DataExplorerQuery) SetSpaceAggregation(v string)`

SetSpaceAggregation sets SpaceAggregation field to given value.


### GetSplitBy

`func (o *DataExplorerQuery) GetSplitBy() []string`

GetSplitBy returns the SplitBy field if non-nil, zero value otherwise.

### GetSplitByOk

`func (o *DataExplorerQuery) GetSplitByOk() (*[]string, bool)`

GetSplitByOk returns a tuple with the SplitBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitBy

`func (o *DataExplorerQuery) SetSplitBy(v []string)`

SetSplitBy sets SplitBy field to given value.


### GetTimeAggregation

`func (o *DataExplorerQuery) GetTimeAggregation() string`

GetTimeAggregation returns the TimeAggregation field if non-nil, zero value otherwise.

### GetTimeAggregationOk

`func (o *DataExplorerQuery) GetTimeAggregationOk() (*string, bool)`

GetTimeAggregationOk returns a tuple with the TimeAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeAggregation

`func (o *DataExplorerQuery) SetTimeAggregation(v string)`

SetTimeAggregation sets TimeAggregation field to given value.


### GetTimeShift

`func (o *DataExplorerQuery) GetTimeShift() DataExplorerTimeShift`

GetTimeShift returns the TimeShift field if non-nil, zero value otherwise.

### GetTimeShiftOk

`func (o *DataExplorerQuery) GetTimeShiftOk() (*DataExplorerTimeShift, bool)`

GetTimeShiftOk returns a tuple with the TimeShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeShift

`func (o *DataExplorerQuery) SetTimeShift(v DataExplorerTimeShift)`

SetTimeShift sets TimeShift field to given value.

### HasTimeShift

`func (o *DataExplorerQuery) HasTimeShift() bool`

HasTimeShift returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


