# SessionReplayMaskingSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaskingPreset** | **string** | The type of the masking:   * &#x60;MASK_ALL&#x60;: Mask all texts, user input, and images.  * &#x60;MASK_USER_INPUT&#x60;: Mask all data that is provided through user input  * &#x60;ALLOW_LIST&#x60;: Only elements, specified in **maskingRules** are shown, everything else is masked.  * &#x60;BLOCK_LIST&#x60;: Elements, specified in **maskingRules** are masked, everything else is shown. | 
**MaskingRules** | Pointer to [**[]MaskingRule**](MaskingRule.md) | A list of masking rules. | [optional] 

## Methods

### NewSessionReplayMaskingSetting

`func NewSessionReplayMaskingSetting(maskingPreset string, ) *SessionReplayMaskingSetting`

NewSessionReplayMaskingSetting instantiates a new SessionReplayMaskingSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionReplayMaskingSettingWithDefaults

`func NewSessionReplayMaskingSettingWithDefaults() *SessionReplayMaskingSetting`

NewSessionReplayMaskingSettingWithDefaults instantiates a new SessionReplayMaskingSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaskingPreset

`func (o *SessionReplayMaskingSetting) GetMaskingPreset() string`

GetMaskingPreset returns the MaskingPreset field if non-nil, zero value otherwise.

### GetMaskingPresetOk

`func (o *SessionReplayMaskingSetting) GetMaskingPresetOk() (*string, bool)`

GetMaskingPresetOk returns a tuple with the MaskingPreset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingPreset

`func (o *SessionReplayMaskingSetting) SetMaskingPreset(v string)`

SetMaskingPreset sets MaskingPreset field to given value.


### GetMaskingRules

`func (o *SessionReplayMaskingSetting) GetMaskingRules() []MaskingRule`

GetMaskingRules returns the MaskingRules field if non-nil, zero value otherwise.

### GetMaskingRulesOk

`func (o *SessionReplayMaskingSetting) GetMaskingRulesOk() (*[]MaskingRule, bool)`

GetMaskingRulesOk returns a tuple with the MaskingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingRules

`func (o *SessionReplayMaskingSetting) SetMaskingRules(v []MaskingRule)`

SetMaskingRules sets MaskingRules field to given value.

### HasMaskingRules

`func (o *SessionReplayMaskingSetting) HasMaskingRules() bool`

HasMaskingRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


