# CustomFilterChartConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AxisLimits** | Pointer to **map[string]float64** | The optional custom y-axis limits. | [optional] 
**LeftAxisCustomUnit** | Pointer to **string** | The custom unit for the left Y-axis. | [optional] 
**LegendShown** | Pointer to **bool** | Defines if a legend should be shown. | [optional] 
**ResultMetadata** | [**map[string]CustomChartingItemMetadataConfig**](CustomChartingItemMetadataConfig.md) | Additional information about charted metric. | 
**RightAxisCustomUnit** | Pointer to **string** | The custom unit for the right Y-axis. | [optional] 
**Series** | [**[]CustomFilterChartSeriesConfig**](CustomFilterChartSeriesConfig.md) | A list of charted metrics. | 
**Type** | **string** | The type of the chart. | 

## Methods

### NewCustomFilterChartConfig

`func NewCustomFilterChartConfig(resultMetadata map[string]CustomChartingItemMetadataConfig, series []CustomFilterChartSeriesConfig, type_ string, ) *CustomFilterChartConfig`

NewCustomFilterChartConfig instantiates a new CustomFilterChartConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFilterChartConfigWithDefaults

`func NewCustomFilterChartConfigWithDefaults() *CustomFilterChartConfig`

NewCustomFilterChartConfigWithDefaults instantiates a new CustomFilterChartConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxisLimits

`func (o *CustomFilterChartConfig) GetAxisLimits() map[string]float64`

GetAxisLimits returns the AxisLimits field if non-nil, zero value otherwise.

### GetAxisLimitsOk

`func (o *CustomFilterChartConfig) GetAxisLimitsOk() (*map[string]float64, bool)`

GetAxisLimitsOk returns a tuple with the AxisLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxisLimits

`func (o *CustomFilterChartConfig) SetAxisLimits(v map[string]float64)`

SetAxisLimits sets AxisLimits field to given value.

### HasAxisLimits

`func (o *CustomFilterChartConfig) HasAxisLimits() bool`

HasAxisLimits returns a boolean if a field has been set.

### GetLeftAxisCustomUnit

`func (o *CustomFilterChartConfig) GetLeftAxisCustomUnit() string`

GetLeftAxisCustomUnit returns the LeftAxisCustomUnit field if non-nil, zero value otherwise.

### GetLeftAxisCustomUnitOk

`func (o *CustomFilterChartConfig) GetLeftAxisCustomUnitOk() (*string, bool)`

GetLeftAxisCustomUnitOk returns a tuple with the LeftAxisCustomUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftAxisCustomUnit

`func (o *CustomFilterChartConfig) SetLeftAxisCustomUnit(v string)`

SetLeftAxisCustomUnit sets LeftAxisCustomUnit field to given value.

### HasLeftAxisCustomUnit

`func (o *CustomFilterChartConfig) HasLeftAxisCustomUnit() bool`

HasLeftAxisCustomUnit returns a boolean if a field has been set.

### GetLegendShown

`func (o *CustomFilterChartConfig) GetLegendShown() bool`

GetLegendShown returns the LegendShown field if non-nil, zero value otherwise.

### GetLegendShownOk

`func (o *CustomFilterChartConfig) GetLegendShownOk() (*bool, bool)`

GetLegendShownOk returns a tuple with the LegendShown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegendShown

`func (o *CustomFilterChartConfig) SetLegendShown(v bool)`

SetLegendShown sets LegendShown field to given value.

### HasLegendShown

`func (o *CustomFilterChartConfig) HasLegendShown() bool`

HasLegendShown returns a boolean if a field has been set.

### GetResultMetadata

`func (o *CustomFilterChartConfig) GetResultMetadata() map[string]CustomChartingItemMetadataConfig`

GetResultMetadata returns the ResultMetadata field if non-nil, zero value otherwise.

### GetResultMetadataOk

`func (o *CustomFilterChartConfig) GetResultMetadataOk() (*map[string]CustomChartingItemMetadataConfig, bool)`

GetResultMetadataOk returns a tuple with the ResultMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultMetadata

`func (o *CustomFilterChartConfig) SetResultMetadata(v map[string]CustomChartingItemMetadataConfig)`

SetResultMetadata sets ResultMetadata field to given value.


### GetRightAxisCustomUnit

`func (o *CustomFilterChartConfig) GetRightAxisCustomUnit() string`

GetRightAxisCustomUnit returns the RightAxisCustomUnit field if non-nil, zero value otherwise.

### GetRightAxisCustomUnitOk

`func (o *CustomFilterChartConfig) GetRightAxisCustomUnitOk() (*string, bool)`

GetRightAxisCustomUnitOk returns a tuple with the RightAxisCustomUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRightAxisCustomUnit

`func (o *CustomFilterChartConfig) SetRightAxisCustomUnit(v string)`

SetRightAxisCustomUnit sets RightAxisCustomUnit field to given value.

### HasRightAxisCustomUnit

`func (o *CustomFilterChartConfig) HasRightAxisCustomUnit() bool`

HasRightAxisCustomUnit returns a boolean if a field has been set.

### GetSeries

`func (o *CustomFilterChartConfig) GetSeries() []CustomFilterChartSeriesConfig`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *CustomFilterChartConfig) GetSeriesOk() (*[]CustomFilterChartSeriesConfig, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *CustomFilterChartConfig) SetSeries(v []CustomFilterChartSeriesConfig)`

SetSeries sets Series field to given value.


### GetType

`func (o *CustomFilterChartConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomFilterChartConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomFilterChartConfig) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


