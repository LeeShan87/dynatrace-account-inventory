/*
Dynatrace Configuration API

Documentation of the Dynatrace Configuration API. To read about use-cases and examples, see [Dynatrace Documentation](https://dt-url.net/4u43kxw).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configv1

import (
	"encoding/json"
)

// checks if the SessionReplayContentMaskingSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionReplayContentMaskingSettings{}

// SessionReplayContentMaskingSettings Content masking settings for Session Replay.   For more details, see [Configure Session Replay](https://dt-url.net/0m03slq) in Dynatrace Documentation.
type SessionReplayContentMaskingSettings struct {
	PlaybackMaskingSettings *SessionReplayMaskingSetting `json:"playbackMaskingSettings,omitempty"`
	RecordingMaskingSettings *SessionReplayMaskingSetting `json:"recordingMaskingSettings,omitempty"`
	// The version of the content masking.   You can use this API only with the version 2.   If you're using version 1, set this field to `2` in the PUT request to switch to version 2.
	RecordingMaskingSettingsVersion int32 `json:"recordingMaskingSettingsVersion"`
}

// NewSessionReplayContentMaskingSettings instantiates a new SessionReplayContentMaskingSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionReplayContentMaskingSettings(recordingMaskingSettingsVersion int32) *SessionReplayContentMaskingSettings {
	this := SessionReplayContentMaskingSettings{}
	this.RecordingMaskingSettingsVersion = recordingMaskingSettingsVersion
	return &this
}

// NewSessionReplayContentMaskingSettingsWithDefaults instantiates a new SessionReplayContentMaskingSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionReplayContentMaskingSettingsWithDefaults() *SessionReplayContentMaskingSettings {
	this := SessionReplayContentMaskingSettings{}
	return &this
}

// GetPlaybackMaskingSettings returns the PlaybackMaskingSettings field value if set, zero value otherwise.
func (o *SessionReplayContentMaskingSettings) GetPlaybackMaskingSettings() SessionReplayMaskingSetting {
	if o == nil || IsNil(o.PlaybackMaskingSettings) {
		var ret SessionReplayMaskingSetting
		return ret
	}
	return *o.PlaybackMaskingSettings
}

// GetPlaybackMaskingSettingsOk returns a tuple with the PlaybackMaskingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionReplayContentMaskingSettings) GetPlaybackMaskingSettingsOk() (*SessionReplayMaskingSetting, bool) {
	if o == nil || IsNil(o.PlaybackMaskingSettings) {
		return nil, false
	}
	return o.PlaybackMaskingSettings, true
}

// HasPlaybackMaskingSettings returns a boolean if a field has been set.
func (o *SessionReplayContentMaskingSettings) HasPlaybackMaskingSettings() bool {
	if o != nil && !IsNil(o.PlaybackMaskingSettings) {
		return true
	}

	return false
}

// SetPlaybackMaskingSettings gets a reference to the given SessionReplayMaskingSetting and assigns it to the PlaybackMaskingSettings field.
func (o *SessionReplayContentMaskingSettings) SetPlaybackMaskingSettings(v SessionReplayMaskingSetting) {
	o.PlaybackMaskingSettings = &v
}

// GetRecordingMaskingSettings returns the RecordingMaskingSettings field value if set, zero value otherwise.
func (o *SessionReplayContentMaskingSettings) GetRecordingMaskingSettings() SessionReplayMaskingSetting {
	if o == nil || IsNil(o.RecordingMaskingSettings) {
		var ret SessionReplayMaskingSetting
		return ret
	}
	return *o.RecordingMaskingSettings
}

// GetRecordingMaskingSettingsOk returns a tuple with the RecordingMaskingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionReplayContentMaskingSettings) GetRecordingMaskingSettingsOk() (*SessionReplayMaskingSetting, bool) {
	if o == nil || IsNil(o.RecordingMaskingSettings) {
		return nil, false
	}
	return o.RecordingMaskingSettings, true
}

// HasRecordingMaskingSettings returns a boolean if a field has been set.
func (o *SessionReplayContentMaskingSettings) HasRecordingMaskingSettings() bool {
	if o != nil && !IsNil(o.RecordingMaskingSettings) {
		return true
	}

	return false
}

// SetRecordingMaskingSettings gets a reference to the given SessionReplayMaskingSetting and assigns it to the RecordingMaskingSettings field.
func (o *SessionReplayContentMaskingSettings) SetRecordingMaskingSettings(v SessionReplayMaskingSetting) {
	o.RecordingMaskingSettings = &v
}

// GetRecordingMaskingSettingsVersion returns the RecordingMaskingSettingsVersion field value
func (o *SessionReplayContentMaskingSettings) GetRecordingMaskingSettingsVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RecordingMaskingSettingsVersion
}

// GetRecordingMaskingSettingsVersionOk returns a tuple with the RecordingMaskingSettingsVersion field value
// and a boolean to check if the value has been set.
func (o *SessionReplayContentMaskingSettings) GetRecordingMaskingSettingsVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordingMaskingSettingsVersion, true
}

// SetRecordingMaskingSettingsVersion sets field value
func (o *SessionReplayContentMaskingSettings) SetRecordingMaskingSettingsVersion(v int32) {
	o.RecordingMaskingSettingsVersion = v
}

func (o SessionReplayContentMaskingSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionReplayContentMaskingSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlaybackMaskingSettings) {
		toSerialize["playbackMaskingSettings"] = o.PlaybackMaskingSettings
	}
	if !IsNil(o.RecordingMaskingSettings) {
		toSerialize["recordingMaskingSettings"] = o.RecordingMaskingSettings
	}
	toSerialize["recordingMaskingSettingsVersion"] = o.RecordingMaskingSettingsVersion
	return toSerialize, nil
}

type NullableSessionReplayContentMaskingSettings struct {
	value *SessionReplayContentMaskingSettings
	isSet bool
}

func (v NullableSessionReplayContentMaskingSettings) Get() *SessionReplayContentMaskingSettings {
	return v.value
}

func (v *NullableSessionReplayContentMaskingSettings) Set(val *SessionReplayContentMaskingSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionReplayContentMaskingSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionReplayContentMaskingSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionReplayContentMaskingSettings(val *SessionReplayContentMaskingSettings) *NullableSessionReplayContentMaskingSettings {
	return &NullableSessionReplayContentMaskingSettings{value: val, isSet: true}
}

func (v NullableSessionReplayContentMaskingSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionReplayContentMaskingSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


