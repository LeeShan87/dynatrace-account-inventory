# SessionReplayDataPrivacySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentMaskingSettings** | Pointer to [**SessionReplayContentMaskingSettings**](SessionReplayContentMaskingSettings.md) |  | [optional] 
**OptInModeEnabled** | Pointer to **bool** | If &#x60;true&#x60;, session recording is disabled until JavaScriptAPI &#x60;dtrum.enableSessionReplay()&#x60; is called. | [optional] 
**UrlExclusionRules** | Pointer to **[]string** | A list of URLs to be excluded from recording. | [optional] 

## Methods

### NewSessionReplayDataPrivacySettings

`func NewSessionReplayDataPrivacySettings() *SessionReplayDataPrivacySettings`

NewSessionReplayDataPrivacySettings instantiates a new SessionReplayDataPrivacySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionReplayDataPrivacySettingsWithDefaults

`func NewSessionReplayDataPrivacySettingsWithDefaults() *SessionReplayDataPrivacySettings`

NewSessionReplayDataPrivacySettingsWithDefaults instantiates a new SessionReplayDataPrivacySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentMaskingSettings

`func (o *SessionReplayDataPrivacySettings) GetContentMaskingSettings() SessionReplayContentMaskingSettings`

GetContentMaskingSettings returns the ContentMaskingSettings field if non-nil, zero value otherwise.

### GetContentMaskingSettingsOk

`func (o *SessionReplayDataPrivacySettings) GetContentMaskingSettingsOk() (*SessionReplayContentMaskingSettings, bool)`

GetContentMaskingSettingsOk returns a tuple with the ContentMaskingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentMaskingSettings

`func (o *SessionReplayDataPrivacySettings) SetContentMaskingSettings(v SessionReplayContentMaskingSettings)`

SetContentMaskingSettings sets ContentMaskingSettings field to given value.

### HasContentMaskingSettings

`func (o *SessionReplayDataPrivacySettings) HasContentMaskingSettings() bool`

HasContentMaskingSettings returns a boolean if a field has been set.

### GetOptInModeEnabled

`func (o *SessionReplayDataPrivacySettings) GetOptInModeEnabled() bool`

GetOptInModeEnabled returns the OptInModeEnabled field if non-nil, zero value otherwise.

### GetOptInModeEnabledOk

`func (o *SessionReplayDataPrivacySettings) GetOptInModeEnabledOk() (*bool, bool)`

GetOptInModeEnabledOk returns a tuple with the OptInModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInModeEnabled

`func (o *SessionReplayDataPrivacySettings) SetOptInModeEnabled(v bool)`

SetOptInModeEnabled sets OptInModeEnabled field to given value.

### HasOptInModeEnabled

`func (o *SessionReplayDataPrivacySettings) HasOptInModeEnabled() bool`

HasOptInModeEnabled returns a boolean if a field has been set.

### GetUrlExclusionRules

`func (o *SessionReplayDataPrivacySettings) GetUrlExclusionRules() []string`

GetUrlExclusionRules returns the UrlExclusionRules field if non-nil, zero value otherwise.

### GetUrlExclusionRulesOk

`func (o *SessionReplayDataPrivacySettings) GetUrlExclusionRulesOk() (*[]string, bool)`

GetUrlExclusionRulesOk returns a tuple with the UrlExclusionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlExclusionRules

`func (o *SessionReplayDataPrivacySettings) SetUrlExclusionRules(v []string)`

SetUrlExclusionRules sets UrlExclusionRules field to given value.

### HasUrlExclusionRules

`func (o *SessionReplayDataPrivacySettings) HasUrlExclusionRules() bool`

HasUrlExclusionRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


