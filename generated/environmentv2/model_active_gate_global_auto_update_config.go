/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
)

// checks if the ActiveGateGlobalAutoUpdateConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveGateGlobalAutoUpdateConfig{}

// ActiveGateGlobalAutoUpdateConfig Global configuration of ActiveGates auto-update.
type ActiveGateGlobalAutoUpdateConfig struct {
	// The state of auto-updates for all ActiveGates connected to the environment or Managed cluster.   This setting is inherited by all ActiveGates that have the `INHERITED` setting.
	GlobalSetting string `json:"globalSetting"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
}

// NewActiveGateGlobalAutoUpdateConfig instantiates a new ActiveGateGlobalAutoUpdateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveGateGlobalAutoUpdateConfig(globalSetting string) *ActiveGateGlobalAutoUpdateConfig {
	this := ActiveGateGlobalAutoUpdateConfig{}
	this.GlobalSetting = globalSetting
	return &this
}

// NewActiveGateGlobalAutoUpdateConfigWithDefaults instantiates a new ActiveGateGlobalAutoUpdateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveGateGlobalAutoUpdateConfigWithDefaults() *ActiveGateGlobalAutoUpdateConfig {
	this := ActiveGateGlobalAutoUpdateConfig{}
	return &this
}

// GetGlobalSetting returns the GlobalSetting field value
func (o *ActiveGateGlobalAutoUpdateConfig) GetGlobalSetting() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalSetting
}

// GetGlobalSettingOk returns a tuple with the GlobalSetting field value
// and a boolean to check if the value has been set.
func (o *ActiveGateGlobalAutoUpdateConfig) GetGlobalSettingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalSetting, true
}

// SetGlobalSetting sets field value
func (o *ActiveGateGlobalAutoUpdateConfig) SetGlobalSetting(v string) {
	o.GlobalSetting = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ActiveGateGlobalAutoUpdateConfig) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateGlobalAutoUpdateConfig) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ActiveGateGlobalAutoUpdateConfig) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *ActiveGateGlobalAutoUpdateConfig) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

func (o ActiveGateGlobalAutoUpdateConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveGateGlobalAutoUpdateConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["globalSetting"] = o.GlobalSetting
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableActiveGateGlobalAutoUpdateConfig struct {
	value *ActiveGateGlobalAutoUpdateConfig
	isSet bool
}

func (v NullableActiveGateGlobalAutoUpdateConfig) Get() *ActiveGateGlobalAutoUpdateConfig {
	return v.value
}

func (v *NullableActiveGateGlobalAutoUpdateConfig) Set(val *ActiveGateGlobalAutoUpdateConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveGateGlobalAutoUpdateConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveGateGlobalAutoUpdateConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveGateGlobalAutoUpdateConfig(val *ActiveGateGlobalAutoUpdateConfig) *NullableActiveGateGlobalAutoUpdateConfig {
	return &NullableActiveGateGlobalAutoUpdateConfig{value: val, isSet: true}
}

func (v NullableActiveGateGlobalAutoUpdateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveGateGlobalAutoUpdateConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


