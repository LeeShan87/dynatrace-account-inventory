# HeatmapSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShowLabels** | Pointer to **bool** |  | [optional] 
**XAxisBuckets** | Pointer to **int32** | Number of buckets in the X axis | [optional] 
**YAxis** | **string** | Y axis aggregation criteria | 
**YAxisBuckets** | Pointer to **int32** | Number of buckets in the Y axis | [optional] 

## Methods

### NewHeatmapSettings

`func NewHeatmapSettings(yAxis string, ) *HeatmapSettings`

NewHeatmapSettings instantiates a new HeatmapSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHeatmapSettingsWithDefaults

`func NewHeatmapSettingsWithDefaults() *HeatmapSettings`

NewHeatmapSettingsWithDefaults instantiates a new HeatmapSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShowLabels

`func (o *HeatmapSettings) GetShowLabels() bool`

GetShowLabels returns the ShowLabels field if non-nil, zero value otherwise.

### GetShowLabelsOk

`func (o *HeatmapSettings) GetShowLabelsOk() (*bool, bool)`

GetShowLabelsOk returns a tuple with the ShowLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLabels

`func (o *HeatmapSettings) SetShowLabels(v bool)`

SetShowLabels sets ShowLabels field to given value.

### HasShowLabels

`func (o *HeatmapSettings) HasShowLabels() bool`

HasShowLabels returns a boolean if a field has been set.

### GetXAxisBuckets

`func (o *HeatmapSettings) GetXAxisBuckets() int32`

GetXAxisBuckets returns the XAxisBuckets field if non-nil, zero value otherwise.

### GetXAxisBucketsOk

`func (o *HeatmapSettings) GetXAxisBucketsOk() (*int32, bool)`

GetXAxisBucketsOk returns a tuple with the XAxisBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXAxisBuckets

`func (o *HeatmapSettings) SetXAxisBuckets(v int32)`

SetXAxisBuckets sets XAxisBuckets field to given value.

### HasXAxisBuckets

`func (o *HeatmapSettings) HasXAxisBuckets() bool`

HasXAxisBuckets returns a boolean if a field has been set.

### GetYAxis

`func (o *HeatmapSettings) GetYAxis() string`

GetYAxis returns the YAxis field if non-nil, zero value otherwise.

### GetYAxisOk

`func (o *HeatmapSettings) GetYAxisOk() (*string, bool)`

GetYAxisOk returns a tuple with the YAxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxis

`func (o *HeatmapSettings) SetYAxis(v string)`

SetYAxis sets YAxis field to given value.


### GetYAxisBuckets

`func (o *HeatmapSettings) GetYAxisBuckets() int32`

GetYAxisBuckets returns the YAxisBuckets field if non-nil, zero value otherwise.

### GetYAxisBucketsOk

`func (o *HeatmapSettings) GetYAxisBucketsOk() (*int32, bool)`

GetYAxisBucketsOk returns a tuple with the YAxisBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYAxisBuckets

`func (o *HeatmapSettings) SetYAxisBuckets(v int32)`

SetYAxisBuckets sets YAxisBuckets field to given value.

### HasYAxisBuckets

`func (o *HeatmapSettings) HasYAxisBuckets() bool`

HasYAxisBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


