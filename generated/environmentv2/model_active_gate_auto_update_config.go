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

// checks if the ActiveGateAutoUpdateConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveGateAutoUpdateConfig{}

// ActiveGateAutoUpdateConfig Configuration of the ActiveGate auto-updates.
type ActiveGateAutoUpdateConfig struct {
	// The actual state of the ActiveGate auto-update.   Applicable only if the **setting** parameter is set to `INHERITED`. In that case, the value is taken from the parent setting. Otherwise, it's just a duplicate of the **setting** value.
	EffectiveSetting *string `json:"effectiveSetting,omitempty"`
	// The state of the ActiveGate auto-update: enabled, disabled, or inherited.   If set to `INHERITED`, the setting is inherited from the global configuration set on the environment or Managed cluster level.
	Setting string `json:"setting"`
}

// NewActiveGateAutoUpdateConfig instantiates a new ActiveGateAutoUpdateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveGateAutoUpdateConfig(setting string) *ActiveGateAutoUpdateConfig {
	this := ActiveGateAutoUpdateConfig{}
	this.Setting = setting
	return &this
}

// NewActiveGateAutoUpdateConfigWithDefaults instantiates a new ActiveGateAutoUpdateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveGateAutoUpdateConfigWithDefaults() *ActiveGateAutoUpdateConfig {
	this := ActiveGateAutoUpdateConfig{}
	return &this
}

// GetEffectiveSetting returns the EffectiveSetting field value if set, zero value otherwise.
func (o *ActiveGateAutoUpdateConfig) GetEffectiveSetting() string {
	if o == nil || IsNil(o.EffectiveSetting) {
		var ret string
		return ret
	}
	return *o.EffectiveSetting
}

// GetEffectiveSettingOk returns a tuple with the EffectiveSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateAutoUpdateConfig) GetEffectiveSettingOk() (*string, bool) {
	if o == nil || IsNil(o.EffectiveSetting) {
		return nil, false
	}
	return o.EffectiveSetting, true
}

// HasEffectiveSetting returns a boolean if a field has been set.
func (o *ActiveGateAutoUpdateConfig) HasEffectiveSetting() bool {
	if o != nil && !IsNil(o.EffectiveSetting) {
		return true
	}

	return false
}

// SetEffectiveSetting gets a reference to the given string and assigns it to the EffectiveSetting field.
func (o *ActiveGateAutoUpdateConfig) SetEffectiveSetting(v string) {
	o.EffectiveSetting = &v
}

// GetSetting returns the Setting field value
func (o *ActiveGateAutoUpdateConfig) GetSetting() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Setting
}

// GetSettingOk returns a tuple with the Setting field value
// and a boolean to check if the value has been set.
func (o *ActiveGateAutoUpdateConfig) GetSettingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Setting, true
}

// SetSetting sets field value
func (o *ActiveGateAutoUpdateConfig) SetSetting(v string) {
	o.Setting = v
}

func (o ActiveGateAutoUpdateConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveGateAutoUpdateConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EffectiveSetting) {
		toSerialize["effectiveSetting"] = o.EffectiveSetting
	}
	toSerialize["setting"] = o.Setting
	return toSerialize, nil
}

type NullableActiveGateAutoUpdateConfig struct {
	value *ActiveGateAutoUpdateConfig
	isSet bool
}

func (v NullableActiveGateAutoUpdateConfig) Get() *ActiveGateAutoUpdateConfig {
	return v.value
}

func (v *NullableActiveGateAutoUpdateConfig) Set(val *ActiveGateAutoUpdateConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveGateAutoUpdateConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveGateAutoUpdateConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveGateAutoUpdateConfig(val *ActiveGateAutoUpdateConfig) *NullableActiveGateAutoUpdateConfig {
	return &NullableActiveGateAutoUpdateConfig{value: val, isSet: true}
}

func (v NullableActiveGateAutoUpdateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveGateAutoUpdateConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


