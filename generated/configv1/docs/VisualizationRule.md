# VisualizationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matcher** | **string** |  | 
**Properties** | [**VisualizationProperties**](VisualizationProperties.md) |  | 
**SeriesOverrides** | Pointer to [**[]VisualizationSerieOverride**](VisualizationSerieOverride.md) |  | [optional] 
**UnitTransform** | Pointer to **string** |  | [optional] 
**ValueFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewVisualizationRule

`func NewVisualizationRule(matcher string, properties VisualizationProperties, ) *VisualizationRule`

NewVisualizationRule instantiates a new VisualizationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualizationRuleWithDefaults

`func NewVisualizationRuleWithDefaults() *VisualizationRule`

NewVisualizationRuleWithDefaults instantiates a new VisualizationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatcher

`func (o *VisualizationRule) GetMatcher() string`

GetMatcher returns the Matcher field if non-nil, zero value otherwise.

### GetMatcherOk

`func (o *VisualizationRule) GetMatcherOk() (*string, bool)`

GetMatcherOk returns a tuple with the Matcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatcher

`func (o *VisualizationRule) SetMatcher(v string)`

SetMatcher sets Matcher field to given value.


### GetProperties

`func (o *VisualizationRule) GetProperties() VisualizationProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *VisualizationRule) GetPropertiesOk() (*VisualizationProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *VisualizationRule) SetProperties(v VisualizationProperties)`

SetProperties sets Properties field to given value.


### GetSeriesOverrides

`func (o *VisualizationRule) GetSeriesOverrides() []VisualizationSerieOverride`

GetSeriesOverrides returns the SeriesOverrides field if non-nil, zero value otherwise.

### GetSeriesOverridesOk

`func (o *VisualizationRule) GetSeriesOverridesOk() (*[]VisualizationSerieOverride, bool)`

GetSeriesOverridesOk returns a tuple with the SeriesOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesOverrides

`func (o *VisualizationRule) SetSeriesOverrides(v []VisualizationSerieOverride)`

SetSeriesOverrides sets SeriesOverrides field to given value.

### HasSeriesOverrides

`func (o *VisualizationRule) HasSeriesOverrides() bool`

HasSeriesOverrides returns a boolean if a field has been set.

### GetUnitTransform

`func (o *VisualizationRule) GetUnitTransform() string`

GetUnitTransform returns the UnitTransform field if non-nil, zero value otherwise.

### GetUnitTransformOk

`func (o *VisualizationRule) GetUnitTransformOk() (*string, bool)`

GetUnitTransformOk returns a tuple with the UnitTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitTransform

`func (o *VisualizationRule) SetUnitTransform(v string)`

SetUnitTransform sets UnitTransform field to given value.

### HasUnitTransform

`func (o *VisualizationRule) HasUnitTransform() bool`

HasUnitTransform returns a boolean if a field has been set.

### GetValueFormat

`func (o *VisualizationRule) GetValueFormat() string`

GetValueFormat returns the ValueFormat field if non-nil, zero value otherwise.

### GetValueFormatOk

`func (o *VisualizationRule) GetValueFormatOk() (*string, bool)`

GetValueFormatOk returns a tuple with the ValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFormat

`func (o *VisualizationRule) SetValueFormat(v string)`

SetValueFormat sets ValueFormat field to given value.

### HasValueFormat

`func (o *VisualizationRule) HasValueFormat() bool`

HasValueFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


