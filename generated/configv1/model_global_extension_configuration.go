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

// checks if the GlobalExtensionConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalExtensionConfiguration{}

// GlobalExtensionConfiguration Global Configuration of OneAgent and JMX extension
type GlobalExtensionConfiguration struct {
	// The extension is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// The ID of the extension.
	ExtensionId *string `json:"extensionId,omitempty"`
	// The plugin is enabled (`true`) or disabled (`false`) globally for hosts in infrastructure-only monitoring mode
	InfraOnlyEnabled *bool `json:"infraOnlyEnabled,omitempty"`
	// The list of configuration parameters.    Each parameter is a key-value pair.
	Properties *map[string]string `json:"properties,omitempty"`
}

// NewGlobalExtensionConfiguration instantiates a new GlobalExtensionConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalExtensionConfiguration(enabled bool) *GlobalExtensionConfiguration {
	this := GlobalExtensionConfiguration{}
	this.Enabled = enabled
	return &this
}

// NewGlobalExtensionConfigurationWithDefaults instantiates a new GlobalExtensionConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalExtensionConfigurationWithDefaults() *GlobalExtensionConfiguration {
	this := GlobalExtensionConfiguration{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *GlobalExtensionConfiguration) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GlobalExtensionConfiguration) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GlobalExtensionConfiguration) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExtensionId returns the ExtensionId field value if set, zero value otherwise.
func (o *GlobalExtensionConfiguration) GetExtensionId() string {
	if o == nil || IsNil(o.ExtensionId) {
		var ret string
		return ret
	}
	return *o.ExtensionId
}

// GetExtensionIdOk returns a tuple with the ExtensionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalExtensionConfiguration) GetExtensionIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExtensionId) {
		return nil, false
	}
	return o.ExtensionId, true
}

// HasExtensionId returns a boolean if a field has been set.
func (o *GlobalExtensionConfiguration) HasExtensionId() bool {
	if o != nil && !IsNil(o.ExtensionId) {
		return true
	}

	return false
}

// SetExtensionId gets a reference to the given string and assigns it to the ExtensionId field.
func (o *GlobalExtensionConfiguration) SetExtensionId(v string) {
	o.ExtensionId = &v
}

// GetInfraOnlyEnabled returns the InfraOnlyEnabled field value if set, zero value otherwise.
func (o *GlobalExtensionConfiguration) GetInfraOnlyEnabled() bool {
	if o == nil || IsNil(o.InfraOnlyEnabled) {
		var ret bool
		return ret
	}
	return *o.InfraOnlyEnabled
}

// GetInfraOnlyEnabledOk returns a tuple with the InfraOnlyEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalExtensionConfiguration) GetInfraOnlyEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.InfraOnlyEnabled) {
		return nil, false
	}
	return o.InfraOnlyEnabled, true
}

// HasInfraOnlyEnabled returns a boolean if a field has been set.
func (o *GlobalExtensionConfiguration) HasInfraOnlyEnabled() bool {
	if o != nil && !IsNil(o.InfraOnlyEnabled) {
		return true
	}

	return false
}

// SetInfraOnlyEnabled gets a reference to the given bool and assigns it to the InfraOnlyEnabled field.
func (o *GlobalExtensionConfiguration) SetInfraOnlyEnabled(v bool) {
	o.InfraOnlyEnabled = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *GlobalExtensionConfiguration) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalExtensionConfiguration) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *GlobalExtensionConfiguration) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *GlobalExtensionConfiguration) SetProperties(v map[string]string) {
	o.Properties = &v
}

func (o GlobalExtensionConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalExtensionConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.ExtensionId) {
		toSerialize["extensionId"] = o.ExtensionId
	}
	if !IsNil(o.InfraOnlyEnabled) {
		toSerialize["infraOnlyEnabled"] = o.InfraOnlyEnabled
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableGlobalExtensionConfiguration struct {
	value *GlobalExtensionConfiguration
	isSet bool
}

func (v NullableGlobalExtensionConfiguration) Get() *GlobalExtensionConfiguration {
	return v.value
}

func (v *NullableGlobalExtensionConfiguration) Set(val *GlobalExtensionConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalExtensionConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalExtensionConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalExtensionConfiguration(val *GlobalExtensionConfiguration) *NullableGlobalExtensionConfiguration {
	return &NullableGlobalExtensionConfiguration{value: val, isSet: true}
}

func (v NullableGlobalExtensionConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalExtensionConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


