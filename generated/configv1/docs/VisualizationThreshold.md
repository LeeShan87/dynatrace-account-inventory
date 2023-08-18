# VisualizationThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AxisTarget** | Pointer to **string** |  | [optional] 
**ColumnId** | Pointer to **string** |  | [optional] 
**QueryId** | Pointer to **string** |  | [optional] 
**Rules** | [**[]VisualizationThresholdRule**](VisualizationThresholdRule.md) |  | 
**Visible** | Pointer to **bool** |  | [optional] 

## Methods

### NewVisualizationThreshold

`func NewVisualizationThreshold(rules []VisualizationThresholdRule, ) *VisualizationThreshold`

NewVisualizationThreshold instantiates a new VisualizationThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizationThresholdWithDefaults

`func NewVisualizationThresholdWithDefaults() *VisualizationThreshold`

NewVisualizationThresholdWithDefaults instantiates a new VisualizationThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxisTarget

`func (o *VisualizationThreshold) GetAxisTarget() string`

GetAxisTarget returns the AxisTarget field if non-nil, zero value otherwise.

### GetAxisTargetOk

`func (o *VisualizationThreshold) GetAxisTargetOk() (*string, bool)`

GetAxisTargetOk returns a tuple with the AxisTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxisTarget

`func (o *VisualizationThreshold) SetAxisTarget(v string)`

SetAxisTarget sets AxisTarget field to given value.

### HasAxisTarget

`func (o *VisualizationThreshold) HasAxisTarget() bool`

HasAxisTarget returns a boolean if a field has been set.

### GetColumnId

`func (o *VisualizationThreshold) GetColumnId() string`

GetColumnId returns the ColumnId field if non-nil, zero value otherwise.

### GetColumnIdOk

`func (o *VisualizationThreshold) GetColumnIdOk() (*string, bool)`

GetColumnIdOk returns a tuple with the ColumnId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnId

`func (o *VisualizationThreshold) SetColumnId(v string)`

SetColumnId sets ColumnId field to given value.

### HasColumnId

`func (o *VisualizationThreshold) HasColumnId() bool`

HasColumnId returns a boolean if a field has been set.

### GetQueryId

`func (o *VisualizationThreshold) GetQueryId() string`

GetQueryId returns the QueryId field if non-nil, zero value otherwise.

### GetQueryIdOk

`func (o *VisualizationThreshold) GetQueryIdOk() (*string, bool)`

GetQueryIdOk returns a tuple with the QueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryId

`func (o *VisualizationThreshold) SetQueryId(v string)`

SetQueryId sets QueryId field to given value.

### HasQueryId

`func (o *VisualizationThreshold) HasQueryId() bool`

HasQueryId returns a boolean if a field has been set.

### GetRules

`func (o *VisualizationThreshold) GetRules() []VisualizationThresholdRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *VisualizationThreshold) GetRulesOk() (*[]VisualizationThresholdRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *VisualizationThreshold) SetRules(v []VisualizationThresholdRule)`

SetRules sets Rules field to given value.


### GetVisible

`func (o *VisualizationThreshold) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *VisualizationThreshold) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *VisualizationThreshold) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *VisualizationThreshold) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


