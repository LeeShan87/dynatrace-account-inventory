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

// checks if the AdvancedJavaScriptTagSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdvancedJavaScriptTagSettings{}

// AdvancedJavaScriptTagSettings Advanced JavaScript tag settings.
type AdvancedJavaScriptTagSettings struct {
	AdditionalEventHandlers AdditionalEventHandlers `json:"additionalEventHandlers"`
	EventWrapperSettings EventWrapperSettings `json:"eventWrapperSettings"`
	GlobalEventCaptureSettings GlobalEventCaptureSettings `json:"globalEventCaptureSettings"`
	// Instrumentation of unsupported Ajax frameworks enabled/disabled.
	InstrumentUnsupportedAjaxFrameworks bool `json:"instrumentUnsupportedAjaxFrameworks"`
	// Maximum character length for action names. Valid values range from 5 to 10000.
	MaxActionNameLength int32 `json:"maxActionNameLength"`
	// Maximum number of errors to be captured per page. Valid values range from 0 to 50.
	MaxErrorsToCapture int32 `json:"maxErrorsToCapture"`
	// Proxy wrapper enabled/disabled.
	ProxyWrapperEnabled *bool `json:"proxyWrapperEnabled,omitempty"`
	// Additional special characters that are to be escaped using non-alphanumeric characters in HTML escape format.
	SpecialCharactersToEscape string `json:"specialCharactersToEscape"`
	// Send the beacon signal as a synchronous XMLHttpRequest using Firefox enabled/disabled.
	SyncBeaconFirefox *bool `json:"syncBeaconFirefox,omitempty"`
	// Send the beacon signal as a synchronous XMLHttpRequest using Internet Explorer enabled/disabled.
	SyncBeaconInternetExplorer *bool `json:"syncBeaconInternetExplorer,omitempty"`
	// User action name attribute.
	UserActionNameAttribute *string `json:"userActionNameAttribute,omitempty"`
}

// NewAdvancedJavaScriptTagSettings instantiates a new AdvancedJavaScriptTagSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdvancedJavaScriptTagSettings(additionalEventHandlers AdditionalEventHandlers, eventWrapperSettings EventWrapperSettings, globalEventCaptureSettings GlobalEventCaptureSettings, instrumentUnsupportedAjaxFrameworks bool, maxActionNameLength int32, maxErrorsToCapture int32, specialCharactersToEscape string) *AdvancedJavaScriptTagSettings {
	this := AdvancedJavaScriptTagSettings{}
	this.AdditionalEventHandlers = additionalEventHandlers
	this.EventWrapperSettings = eventWrapperSettings
	this.GlobalEventCaptureSettings = globalEventCaptureSettings
	this.InstrumentUnsupportedAjaxFrameworks = instrumentUnsupportedAjaxFrameworks
	this.MaxActionNameLength = maxActionNameLength
	this.MaxErrorsToCapture = maxErrorsToCapture
	this.SpecialCharactersToEscape = specialCharactersToEscape
	return &this
}

// NewAdvancedJavaScriptTagSettingsWithDefaults instantiates a new AdvancedJavaScriptTagSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdvancedJavaScriptTagSettingsWithDefaults() *AdvancedJavaScriptTagSettings {
	this := AdvancedJavaScriptTagSettings{}
	return &this
}

// GetAdditionalEventHandlers returns the AdditionalEventHandlers field value
func (o *AdvancedJavaScriptTagSettings) GetAdditionalEventHandlers() AdditionalEventHandlers {
	if o == nil {
		var ret AdditionalEventHandlers
		return ret
	}

	return o.AdditionalEventHandlers
}

// GetAdditionalEventHandlersOk returns a tuple with the AdditionalEventHandlers field value
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetAdditionalEventHandlersOk() (*AdditionalEventHandlers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalEventHandlers, true
}

// SetAdditionalEventHandlers sets field value
func (o *AdvancedJavaScriptTagSettings) SetAdditionalEventHandlers(v AdditionalEventHandlers) {
	o.AdditionalEventHandlers = v
}

// GetEventWrapperSettings returns the EventWrapperSettings field value
func (o *AdvancedJavaScriptTagSettings) GetEventWrapperSettings() EventWrapperSettings {
	if o == nil {
		var ret EventWrapperSettings
		return ret
	}

	return o.EventWrapperSettings
}

// GetEventWrapperSettingsOk returns a tuple with the EventWrapperSettings field value
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetEventWrapperSettingsOk() (*EventWrapperSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventWrapperSettings, true
}

// SetEventWrapperSettings sets field value
func (o *AdvancedJavaScriptTagSettings) SetEventWrapperSettings(v EventWrapperSettings) {
	o.EventWrapperSettings = v
}

// GetGlobalEventCaptureSettings returns the GlobalEventCaptureSettings field value
func (o *AdvancedJavaScriptTagSettings) GetGlobalEventCaptureSettings() GlobalEventCaptureSettings {
	if o == nil {
		var ret GlobalEventCaptureSettings
		return ret
	}

	return o.GlobalEventCaptureSettings
}

// GetGlobalEventCaptureSettingsOk returns a tuple with the GlobalEventCaptureSettings field value
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetGlobalEventCaptureSettingsOk() (*GlobalEventCaptureSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalEventCaptureSettings, true
}

// SetGlobalEventCaptureSettings sets field value
func (o *AdvancedJavaScriptTagSettings) SetGlobalEventCaptureSettings(v GlobalEventCaptureSettings) {
	o.GlobalEventCaptureSettings = v
}

// GetInstrumentUnsupportedAjaxFrameworks returns the InstrumentUnsupportedAjaxFrameworks field value
func (o *AdvancedJavaScriptTagSettings) GetInstrumentUnsupportedAjaxFrameworks() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.InstrumentUnsupportedAjaxFrameworks
}

// GetInstrumentUnsupportedAjaxFrameworksOk returns a tuple with the InstrumentUnsupportedAjaxFrameworks field value
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetInstrumentUnsupportedAjaxFrameworksOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstrumentUnsupportedAjaxFrameworks, true
}

// SetInstrumentUnsupportedAjaxFrameworks sets field value
func (o *AdvancedJavaScriptTagSettings) SetInstrumentUnsupportedAjaxFrameworks(v bool) {
	o.InstrumentUnsupportedAjaxFrameworks = v
}

// GetMaxActionNameLength returns the MaxActionNameLength field value
func (o *AdvancedJavaScriptTagSettings) GetMaxActionNameLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxActionNameLength
}

// GetMaxActionNameLengthOk returns a tuple with the MaxActionNameLength field value
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetMaxActionNameLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxActionNameLength, true
}

// SetMaxActionNameLength sets field value
func (o *AdvancedJavaScriptTagSettings) SetMaxActionNameLength(v int32) {
	o.MaxActionNameLength = v
}

// GetMaxErrorsToCapture returns the MaxErrorsToCapture field value
func (o *AdvancedJavaScriptTagSettings) GetMaxErrorsToCapture() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxErrorsToCapture
}

// GetMaxErrorsToCaptureOk returns a tuple with the MaxErrorsToCapture field value
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetMaxErrorsToCaptureOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxErrorsToCapture, true
}

// SetMaxErrorsToCapture sets field value
func (o *AdvancedJavaScriptTagSettings) SetMaxErrorsToCapture(v int32) {
	o.MaxErrorsToCapture = v
}

// GetProxyWrapperEnabled returns the ProxyWrapperEnabled field value if set, zero value otherwise.
func (o *AdvancedJavaScriptTagSettings) GetProxyWrapperEnabled() bool {
	if o == nil || IsNil(o.ProxyWrapperEnabled) {
		var ret bool
		return ret
	}
	return *o.ProxyWrapperEnabled
}

// GetProxyWrapperEnabledOk returns a tuple with the ProxyWrapperEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetProxyWrapperEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ProxyWrapperEnabled) {
		return nil, false
	}
	return o.ProxyWrapperEnabled, true
}

// HasProxyWrapperEnabled returns a boolean if a field has been set.
func (o *AdvancedJavaScriptTagSettings) HasProxyWrapperEnabled() bool {
	if o != nil && !IsNil(o.ProxyWrapperEnabled) {
		return true
	}

	return false
}

// SetProxyWrapperEnabled gets a reference to the given bool and assigns it to the ProxyWrapperEnabled field.
func (o *AdvancedJavaScriptTagSettings) SetProxyWrapperEnabled(v bool) {
	o.ProxyWrapperEnabled = &v
}

// GetSpecialCharactersToEscape returns the SpecialCharactersToEscape field value
func (o *AdvancedJavaScriptTagSettings) GetSpecialCharactersToEscape() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpecialCharactersToEscape
}

// GetSpecialCharactersToEscapeOk returns a tuple with the SpecialCharactersToEscape field value
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetSpecialCharactersToEscapeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecialCharactersToEscape, true
}

// SetSpecialCharactersToEscape sets field value
func (o *AdvancedJavaScriptTagSettings) SetSpecialCharactersToEscape(v string) {
	o.SpecialCharactersToEscape = v
}

// GetSyncBeaconFirefox returns the SyncBeaconFirefox field value if set, zero value otherwise.
func (o *AdvancedJavaScriptTagSettings) GetSyncBeaconFirefox() bool {
	if o == nil || IsNil(o.SyncBeaconFirefox) {
		var ret bool
		return ret
	}
	return *o.SyncBeaconFirefox
}

// GetSyncBeaconFirefoxOk returns a tuple with the SyncBeaconFirefox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetSyncBeaconFirefoxOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncBeaconFirefox) {
		return nil, false
	}
	return o.SyncBeaconFirefox, true
}

// HasSyncBeaconFirefox returns a boolean if a field has been set.
func (o *AdvancedJavaScriptTagSettings) HasSyncBeaconFirefox() bool {
	if o != nil && !IsNil(o.SyncBeaconFirefox) {
		return true
	}

	return false
}

// SetSyncBeaconFirefox gets a reference to the given bool and assigns it to the SyncBeaconFirefox field.
func (o *AdvancedJavaScriptTagSettings) SetSyncBeaconFirefox(v bool) {
	o.SyncBeaconFirefox = &v
}

// GetSyncBeaconInternetExplorer returns the SyncBeaconInternetExplorer field value if set, zero value otherwise.
func (o *AdvancedJavaScriptTagSettings) GetSyncBeaconInternetExplorer() bool {
	if o == nil || IsNil(o.SyncBeaconInternetExplorer) {
		var ret bool
		return ret
	}
	return *o.SyncBeaconInternetExplorer
}

// GetSyncBeaconInternetExplorerOk returns a tuple with the SyncBeaconInternetExplorer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetSyncBeaconInternetExplorerOk() (*bool, bool) {
	if o == nil || IsNil(o.SyncBeaconInternetExplorer) {
		return nil, false
	}
	return o.SyncBeaconInternetExplorer, true
}

// HasSyncBeaconInternetExplorer returns a boolean if a field has been set.
func (o *AdvancedJavaScriptTagSettings) HasSyncBeaconInternetExplorer() bool {
	if o != nil && !IsNil(o.SyncBeaconInternetExplorer) {
		return true
	}

	return false
}

// SetSyncBeaconInternetExplorer gets a reference to the given bool and assigns it to the SyncBeaconInternetExplorer field.
func (o *AdvancedJavaScriptTagSettings) SetSyncBeaconInternetExplorer(v bool) {
	o.SyncBeaconInternetExplorer = &v
}

// GetUserActionNameAttribute returns the UserActionNameAttribute field value if set, zero value otherwise.
func (o *AdvancedJavaScriptTagSettings) GetUserActionNameAttribute() string {
	if o == nil || IsNil(o.UserActionNameAttribute) {
		var ret string
		return ret
	}
	return *o.UserActionNameAttribute
}

// GetUserActionNameAttributeOk returns a tuple with the UserActionNameAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdvancedJavaScriptTagSettings) GetUserActionNameAttributeOk() (*string, bool) {
	if o == nil || IsNil(o.UserActionNameAttribute) {
		return nil, false
	}
	return o.UserActionNameAttribute, true
}

// HasUserActionNameAttribute returns a boolean if a field has been set.
func (o *AdvancedJavaScriptTagSettings) HasUserActionNameAttribute() bool {
	if o != nil && !IsNil(o.UserActionNameAttribute) {
		return true
	}

	return false
}

// SetUserActionNameAttribute gets a reference to the given string and assigns it to the UserActionNameAttribute field.
func (o *AdvancedJavaScriptTagSettings) SetUserActionNameAttribute(v string) {
	o.UserActionNameAttribute = &v
}

func (o AdvancedJavaScriptTagSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdvancedJavaScriptTagSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["additionalEventHandlers"] = o.AdditionalEventHandlers
	toSerialize["eventWrapperSettings"] = o.EventWrapperSettings
	toSerialize["globalEventCaptureSettings"] = o.GlobalEventCaptureSettings
	toSerialize["instrumentUnsupportedAjaxFrameworks"] = o.InstrumentUnsupportedAjaxFrameworks
	toSerialize["maxActionNameLength"] = o.MaxActionNameLength
	toSerialize["maxErrorsToCapture"] = o.MaxErrorsToCapture
	if !IsNil(o.ProxyWrapperEnabled) {
		toSerialize["proxyWrapperEnabled"] = o.ProxyWrapperEnabled
	}
	toSerialize["specialCharactersToEscape"] = o.SpecialCharactersToEscape
	if !IsNil(o.SyncBeaconFirefox) {
		toSerialize["syncBeaconFirefox"] = o.SyncBeaconFirefox
	}
	if !IsNil(o.SyncBeaconInternetExplorer) {
		toSerialize["syncBeaconInternetExplorer"] = o.SyncBeaconInternetExplorer
	}
	if !IsNil(o.UserActionNameAttribute) {
		toSerialize["userActionNameAttribute"] = o.UserActionNameAttribute
	}
	return toSerialize, nil
}

type NullableAdvancedJavaScriptTagSettings struct {
	value *AdvancedJavaScriptTagSettings
	isSet bool
}

func (v NullableAdvancedJavaScriptTagSettings) Get() *AdvancedJavaScriptTagSettings {
	return v.value
}

func (v *NullableAdvancedJavaScriptTagSettings) Set(val *AdvancedJavaScriptTagSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAdvancedJavaScriptTagSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAdvancedJavaScriptTagSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdvancedJavaScriptTagSettings(val *AdvancedJavaScriptTagSettings) *NullableAdvancedJavaScriptTagSettings {
	return &NullableAdvancedJavaScriptTagSettings{value: val, isSet: true}
}

func (v NullableAdvancedJavaScriptTagSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdvancedJavaScriptTagSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


