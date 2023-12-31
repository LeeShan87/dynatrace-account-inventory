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

// checks if the MonitoringConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringConfig{}

// MonitoringConfig Monitoring configuration of OneAgent.
type MonitoringConfig struct {
	// Code modules will be injected automatically into monitored applications if this setting is enabled. This setting won't apply if auto-injection is disabled via oneagentctl (see https://dt-url.net/oneagentctl).
	AutoInjectionEnabled *bool `json:"autoInjectionEnabled,omitempty"`
	// The Dynatrace entity ID of the host where OneAgent is deployed.
	Id *string `json:"id,omitempty"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// The monitoring is enabled (`true`) or disabled (`false`).
	MonitoringEnabled bool `json:"monitoringEnabled"`
	// The monitoring mode for the host: full stack or infrastructure only.
	MonitoringMode string `json:"monitoringMode"`
}

// NewMonitoringConfig instantiates a new MonitoringConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringConfig(monitoringEnabled bool, monitoringMode string) *MonitoringConfig {
	this := MonitoringConfig{}
	this.MonitoringEnabled = monitoringEnabled
	this.MonitoringMode = monitoringMode
	return &this
}

// NewMonitoringConfigWithDefaults instantiates a new MonitoringConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringConfigWithDefaults() *MonitoringConfig {
	this := MonitoringConfig{}
	return &this
}

// GetAutoInjectionEnabled returns the AutoInjectionEnabled field value if set, zero value otherwise.
func (o *MonitoringConfig) GetAutoInjectionEnabled() bool {
	if o == nil || IsNil(o.AutoInjectionEnabled) {
		var ret bool
		return ret
	}
	return *o.AutoInjectionEnabled
}

// GetAutoInjectionEnabledOk returns a tuple with the AutoInjectionEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfig) GetAutoInjectionEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoInjectionEnabled) {
		return nil, false
	}
	return o.AutoInjectionEnabled, true
}

// HasAutoInjectionEnabled returns a boolean if a field has been set.
func (o *MonitoringConfig) HasAutoInjectionEnabled() bool {
	if o != nil && !IsNil(o.AutoInjectionEnabled) {
		return true
	}

	return false
}

// SetAutoInjectionEnabled gets a reference to the given bool and assigns it to the AutoInjectionEnabled field.
func (o *MonitoringConfig) SetAutoInjectionEnabled(v bool) {
	o.AutoInjectionEnabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MonitoringConfig) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfig) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MonitoringConfig) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MonitoringConfig) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MonitoringConfig) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringConfig) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MonitoringConfig) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *MonitoringConfig) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetMonitoringEnabled returns the MonitoringEnabled field value
func (o *MonitoringConfig) GetMonitoringEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MonitoringEnabled
}

// GetMonitoringEnabledOk returns a tuple with the MonitoringEnabled field value
// and a boolean to check if the value has been set.
func (o *MonitoringConfig) GetMonitoringEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitoringEnabled, true
}

// SetMonitoringEnabled sets field value
func (o *MonitoringConfig) SetMonitoringEnabled(v bool) {
	o.MonitoringEnabled = v
}

// GetMonitoringMode returns the MonitoringMode field value
func (o *MonitoringConfig) GetMonitoringMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MonitoringMode
}

// GetMonitoringModeOk returns a tuple with the MonitoringMode field value
// and a boolean to check if the value has been set.
func (o *MonitoringConfig) GetMonitoringModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitoringMode, true
}

// SetMonitoringMode sets field value
func (o *MonitoringConfig) SetMonitoringMode(v string) {
	o.MonitoringMode = v
}

func (o MonitoringConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoInjectionEnabled) {
		toSerialize["autoInjectionEnabled"] = o.AutoInjectionEnabled
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["monitoringEnabled"] = o.MonitoringEnabled
	toSerialize["monitoringMode"] = o.MonitoringMode
	return toSerialize, nil
}

type NullableMonitoringConfig struct {
	value *MonitoringConfig
	isSet bool
}

func (v NullableMonitoringConfig) Get() *MonitoringConfig {
	return v.value
}

func (v *NullableMonitoringConfig) Set(val *MonitoringConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringConfig(val *MonitoringConfig) *NullableMonitoringConfig {
	return &NullableMonitoringConfig{value: val, isSet: true}
}

func (v NullableMonitoringConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


