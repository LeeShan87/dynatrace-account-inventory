# SessionReplaySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CostControlPercentage** | **int32** | Session replay sampling rating in percentage. | 
**CssResourceCapturingExclusionRules** | Pointer to **[]string** | A list of URLs to be excluded from CSS resource capturing. | [optional] 
**EnableCssResourceCapturing** | Pointer to **bool** | Capture (&#x60;true&#x60;) or don&#39;t capture (&#x60;false&#x60;) CSS resources from the session. | [optional] 
**Enabled** | **bool** | SessionReplay Enabled. | 

## Methods

### NewSessionReplaySetting

`func NewSessionReplaySetting(costControlPercentage int32, enabled bool, ) *SessionReplaySetting`

NewSessionReplaySetting instantiates a new SessionReplaySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionReplaySettingWithDefaults

`func NewSessionReplaySettingWithDefaults() *SessionReplaySetting`

NewSessionReplaySettingWithDefaults instantiates a new SessionReplaySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCostControlPercentage

`func (o *SessionReplaySetting) GetCostControlPercentage() int32`

GetCostControlPercentage returns the CostControlPercentage field if non-nil, zero value otherwise.

### GetCostControlPercentageOk

`func (o *SessionReplaySetting) GetCostControlPercentageOk() (*int32, bool)`

GetCostControlPercentageOk returns a tuple with the CostControlPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostControlPercentage

`func (o *SessionReplaySetting) SetCostControlPercentage(v int32)`

SetCostControlPercentage sets CostControlPercentage field to given value.


### GetCssResourceCapturingExclusionRules

`func (o *SessionReplaySetting) GetCssResourceCapturingExclusionRules() []string`

GetCssResourceCapturingExclusionRules returns the CssResourceCapturingExclusionRules field if non-nil, zero value otherwise.

### GetCssResourceCapturingExclusionRulesOk

`func (o *SessionReplaySetting) GetCssResourceCapturingExclusionRulesOk() (*[]string, bool)`

GetCssResourceCapturingExclusionRulesOk returns a tuple with the CssResourceCapturingExclusionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCssResourceCapturingExclusionRules

`func (o *SessionReplaySetting) SetCssResourceCapturingExclusionRules(v []string)`

SetCssResourceCapturingExclusionRules sets CssResourceCapturingExclusionRules field to given value.

### HasCssResourceCapturingExclusionRules

`func (o *SessionReplaySetting) HasCssResourceCapturingExclusionRules() bool`

HasCssResourceCapturingExclusionRules returns a boolean if a field has been set.

### GetEnableCssResourceCapturing

`func (o *SessionReplaySetting) GetEnableCssResourceCapturing() bool`

GetEnableCssResourceCapturing returns the EnableCssResourceCapturing field if non-nil, zero value otherwise.

### GetEnableCssResourceCapturingOk

`func (o *SessionReplaySetting) GetEnableCssResourceCapturingOk() (*bool, bool)`

GetEnableCssResourceCapturingOk returns a tuple with the EnableCssResourceCapturing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCssResourceCapturing

`func (o *SessionReplaySetting) SetEnableCssResourceCapturing(v bool)`

SetEnableCssResourceCapturing sets EnableCssResourceCapturing field to given value.

### HasEnableCssResourceCapturing

`func (o *SessionReplaySetting) HasEnableCssResourceCapturing() bool`

HasEnableCssResourceCapturing returns a boolean if a field has been set.

### GetEnabled

`func (o *SessionReplaySetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SessionReplaySetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SessionReplaySetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


