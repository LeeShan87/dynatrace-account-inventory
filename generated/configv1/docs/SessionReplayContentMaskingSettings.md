# SessionReplayContentMaskingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlaybackMaskingSettings** | Pointer to [**SessionReplayMaskingSetting**](SessionReplayMaskingSetting.md) |  | [optional] 
**RecordingMaskingSettings** | Pointer to [**SessionReplayMaskingSetting**](SessionReplayMaskingSetting.md) |  | [optional] 
**RecordingMaskingSettingsVersion** | **int32** | The version of the content masking.   You can use this API only with the version 2.   If you&#39;re using version 1, set this field to &#x60;2&#x60; in the PUT request to switch to version 2. | 

## Methods

### NewSessionReplayContentMaskingSettings

`func NewSessionReplayContentMaskingSettings(recordingMaskingSettingsVersion int32, ) *SessionReplayContentMaskingSettings`

NewSessionReplayContentMaskingSettings instantiates a new SessionReplayContentMaskingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionReplayContentMaskingSettingsWithDefaults

`func NewSessionReplayContentMaskingSettingsWithDefaults() *SessionReplayContentMaskingSettings`

NewSessionReplayContentMaskingSettingsWithDefaults instantiates a new SessionReplayContentMaskingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlaybackMaskingSettings

`func (o *SessionReplayContentMaskingSettings) GetPlaybackMaskingSettings() SessionReplayMaskingSetting`

GetPlaybackMaskingSettings returns the PlaybackMaskingSettings field if non-nil, zero value otherwise.

### GetPlaybackMaskingSettingsOk

`func (o *SessionReplayContentMaskingSettings) GetPlaybackMaskingSettingsOk() (*SessionReplayMaskingSetting, bool)`

GetPlaybackMaskingSettingsOk returns a tuple with the PlaybackMaskingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybackMaskingSettings

`func (o *SessionReplayContentMaskingSettings) SetPlaybackMaskingSettings(v SessionReplayMaskingSetting)`

SetPlaybackMaskingSettings sets PlaybackMaskingSettings field to given value.

### HasPlaybackMaskingSettings

`func (o *SessionReplayContentMaskingSettings) HasPlaybackMaskingSettings() bool`

HasPlaybackMaskingSettings returns a boolean if a field has been set.

### GetRecordingMaskingSettings

`func (o *SessionReplayContentMaskingSettings) GetRecordingMaskingSettings() SessionReplayMaskingSetting`

GetRecordingMaskingSettings returns the RecordingMaskingSettings field if non-nil, zero value otherwise.

### GetRecordingMaskingSettingsOk

`func (o *SessionReplayContentMaskingSettings) GetRecordingMaskingSettingsOk() (*SessionReplayMaskingSetting, bool)`

GetRecordingMaskingSettingsOk returns a tuple with the RecordingMaskingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingMaskingSettings

`func (o *SessionReplayContentMaskingSettings) SetRecordingMaskingSettings(v SessionReplayMaskingSetting)`

SetRecordingMaskingSettings sets RecordingMaskingSettings field to given value.

### HasRecordingMaskingSettings

`func (o *SessionReplayContentMaskingSettings) HasRecordingMaskingSettings() bool`

HasRecordingMaskingSettings returns a boolean if a field has been set.

### GetRecordingMaskingSettingsVersion

`func (o *SessionReplayContentMaskingSettings) GetRecordingMaskingSettingsVersion() int32`

GetRecordingMaskingSettingsVersion returns the RecordingMaskingSettingsVersion field if non-nil, zero value otherwise.

### GetRecordingMaskingSettingsVersionOk

`func (o *SessionReplayContentMaskingSettings) GetRecordingMaskingSettingsVersionOk() (*int32, bool)`

GetRecordingMaskingSettingsVersionOk returns a tuple with the RecordingMaskingSettingsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingMaskingSettingsVersion

`func (o *SessionReplayContentMaskingSettings) SetRecordingMaskingSettingsVersion(v int32)`

SetRecordingMaskingSettingsVersion sets RecordingMaskingSettingsVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


