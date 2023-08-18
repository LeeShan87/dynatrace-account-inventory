# VisualizationConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Axes** | Pointer to [**Axes**](Axes.md) |  | [optional] 
**Global** | Pointer to [**VisualizationGlobalProperties**](VisualizationGlobalProperties.md) |  | [optional] 
**GraphChartSettings** | Pointer to [**GraphChartSettings**](GraphChartSettings.md) |  | [optional] 
**HeatmapSettings** | Pointer to [**HeatmapSettings**](HeatmapSettings.md) |  | [optional] 
**HoneycombSettings** | Pointer to [**HoneycombSettings**](HoneycombSettings.md) |  | [optional] 
**Rules** | [**[]VisualizationRule**](VisualizationRule.md) | Rules for Visualization | 
**SingleValueSettings** | Pointer to [**SingleValueSettings**](SingleValueSettings.md) |  | [optional] 
**TableSettings** | Pointer to [**TableSettings**](TableSettings.md) |  | [optional] 
**Thresholds** | Pointer to [**[]VisualizationThreshold**](VisualizationThreshold.md) | Thresholds for Visualization | [optional] 
**Type** | **string** | The id of the query | 

## Methods

### NewVisualizationConfiguration

`func NewVisualizationConfiguration(rules []VisualizationRule, type_ string, ) *VisualizationConfiguration`

NewVisualizationConfiguration instantiates a new VisualizationConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizationConfigurationWithDefaults

`func NewVisualizationConfigurationWithDefaults() *VisualizationConfiguration`

NewVisualizationConfigurationWithDefaults instantiates a new VisualizationConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxes

`func (o *VisualizationConfiguration) GetAxes() Axes`

GetAxes returns the Axes field if non-nil, zero value otherwise.

### GetAxesOk

`func (o *VisualizationConfiguration) GetAxesOk() (*Axes, bool)`

GetAxesOk returns a tuple with the Axes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxes

`func (o *VisualizationConfiguration) SetAxes(v Axes)`

SetAxes sets Axes field to given value.

### HasAxes

`func (o *VisualizationConfiguration) HasAxes() bool`

HasAxes returns a boolean if a field has been set.

### GetGlobal

`func (o *VisualizationConfiguration) GetGlobal() VisualizationGlobalProperties`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *VisualizationConfiguration) GetGlobalOk() (*VisualizationGlobalProperties, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *VisualizationConfiguration) SetGlobal(v VisualizationGlobalProperties)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *VisualizationConfiguration) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetGraphChartSettings

`func (o *VisualizationConfiguration) GetGraphChartSettings() GraphChartSettings`

GetGraphChartSettings returns the GraphChartSettings field if non-nil, zero value otherwise.

### GetGraphChartSettingsOk

`func (o *VisualizationConfiguration) GetGraphChartSettingsOk() (*GraphChartSettings, bool)`

GetGraphChartSettingsOk returns a tuple with the GraphChartSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphChartSettings

`func (o *VisualizationConfiguration) SetGraphChartSettings(v GraphChartSettings)`

SetGraphChartSettings sets GraphChartSettings field to given value.

### HasGraphChartSettings

`func (o *VisualizationConfiguration) HasGraphChartSettings() bool`

HasGraphChartSettings returns a boolean if a field has been set.

### GetHeatmapSettings

`func (o *VisualizationConfiguration) GetHeatmapSettings() HeatmapSettings`

GetHeatmapSettings returns the HeatmapSettings field if non-nil, zero value otherwise.

### GetHeatmapSettingsOk

`func (o *VisualizationConfiguration) GetHeatmapSettingsOk() (*HeatmapSettings, bool)`

GetHeatmapSettingsOk returns a tuple with the HeatmapSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmapSettings

`func (o *VisualizationConfiguration) SetHeatmapSettings(v HeatmapSettings)`

SetHeatmapSettings sets HeatmapSettings field to given value.

### HasHeatmapSettings

`func (o *VisualizationConfiguration) HasHeatmapSettings() bool`

HasHeatmapSettings returns a boolean if a field has been set.

### GetHoneycombSettings

`func (o *VisualizationConfiguration) GetHoneycombSettings() HoneycombSettings`

GetHoneycombSettings returns the HoneycombSettings field if non-nil, zero value otherwise.

### GetHoneycombSettingsOk

`func (o *VisualizationConfiguration) GetHoneycombSettingsOk() (*HoneycombSettings, bool)`

GetHoneycombSettingsOk returns a tuple with the HoneycombSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoneycombSettings

`func (o *VisualizationConfiguration) SetHoneycombSettings(v HoneycombSettings)`

SetHoneycombSettings sets HoneycombSettings field to given value.

### HasHoneycombSettings

`func (o *VisualizationConfiguration) HasHoneycombSettings() bool`

HasHoneycombSettings returns a boolean if a field has been set.

### GetRules

`func (o *VisualizationConfiguration) GetRules() []VisualizationRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *VisualizationConfiguration) GetRulesOk() (*[]VisualizationRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *VisualizationConfiguration) SetRules(v []VisualizationRule)`

SetRules sets Rules field to given value.


### GetSingleValueSettings

`func (o *VisualizationConfiguration) GetSingleValueSettings() SingleValueSettings`

GetSingleValueSettings returns the SingleValueSettings field if non-nil, zero value otherwise.

### GetSingleValueSettingsOk

`func (o *VisualizationConfiguration) GetSingleValueSettingsOk() (*SingleValueSettings, bool)`

GetSingleValueSettingsOk returns a tuple with the SingleValueSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleValueSettings

`func (o *VisualizationConfiguration) SetSingleValueSettings(v SingleValueSettings)`

SetSingleValueSettings sets SingleValueSettings field to given value.

### HasSingleValueSettings

`func (o *VisualizationConfiguration) HasSingleValueSettings() bool`

HasSingleValueSettings returns a boolean if a field has been set.

### GetTableSettings

`func (o *VisualizationConfiguration) GetTableSettings() TableSettings`

GetTableSettings returns the TableSettings field if non-nil, zero value otherwise.

### GetTableSettingsOk

`func (o *VisualizationConfiguration) GetTableSettingsOk() (*TableSettings, bool)`

GetTableSettingsOk returns a tuple with the TableSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableSettings

`func (o *VisualizationConfiguration) SetTableSettings(v TableSettings)`

SetTableSettings sets TableSettings field to given value.

### HasTableSettings

`func (o *VisualizationConfiguration) HasTableSettings() bool`

HasTableSettings returns a boolean if a field has been set.

### GetThresholds

`func (o *VisualizationConfiguration) GetThresholds() []VisualizationThreshold`

GetThresholds returns the Thresholds field if non-nil, zero value otherwise.

### GetThresholdsOk

`func (o *VisualizationConfiguration) GetThresholdsOk() (*[]VisualizationThreshold, bool)`

GetThresholdsOk returns a tuple with the Thresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholds

`func (o *VisualizationConfiguration) SetThresholds(v []VisualizationThreshold)`

SetThresholds sets Thresholds field to given value.

### HasThresholds

`func (o *VisualizationConfiguration) HasThresholds() bool`

HasThresholds returns a boolean if a field has been set.

### GetType

`func (o *VisualizationConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VisualizationConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VisualizationConfiguration) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


