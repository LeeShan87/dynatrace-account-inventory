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

// checks if the FrequentIssueDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrequentIssueDetectionConfig{}

// FrequentIssueDetectionConfig Parameters of the frequent issue detection. To learn more about it, see [Detection of frequent issues](https://dt-url.net/4da3kdg) in Dynatrace Documentation. 
type FrequentIssueDetectionConfig struct {
	// The detection for applications is enabled (`true`) or disabled (`false`).
	FrequentIssueDetectionApplicationEnabled bool `json:"frequentIssueDetectionApplicationEnabled"`
	// The detection for infrastructure is enabled (`true`) or disabled (`false`).
	FrequentIssueDetectionInfrastructureEnabled bool `json:"frequentIssueDetectionInfrastructureEnabled"`
	// The detection for services is enabled (`true`) or disabled (`false`).
	FrequentIssueDetectionServiceEnabled bool `json:"frequentIssueDetectionServiceEnabled"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
}

// NewFrequentIssueDetectionConfig instantiates a new FrequentIssueDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrequentIssueDetectionConfig(frequentIssueDetectionApplicationEnabled bool, frequentIssueDetectionInfrastructureEnabled bool, frequentIssueDetectionServiceEnabled bool) *FrequentIssueDetectionConfig {
	this := FrequentIssueDetectionConfig{}
	this.FrequentIssueDetectionApplicationEnabled = frequentIssueDetectionApplicationEnabled
	this.FrequentIssueDetectionInfrastructureEnabled = frequentIssueDetectionInfrastructureEnabled
	this.FrequentIssueDetectionServiceEnabled = frequentIssueDetectionServiceEnabled
	return &this
}

// NewFrequentIssueDetectionConfigWithDefaults instantiates a new FrequentIssueDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrequentIssueDetectionConfigWithDefaults() *FrequentIssueDetectionConfig {
	this := FrequentIssueDetectionConfig{}
	return &this
}

// GetFrequentIssueDetectionApplicationEnabled returns the FrequentIssueDetectionApplicationEnabled field value
func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionApplicationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FrequentIssueDetectionApplicationEnabled
}

// GetFrequentIssueDetectionApplicationEnabledOk returns a tuple with the FrequentIssueDetectionApplicationEnabled field value
// and a boolean to check if the value has been set.
func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionApplicationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrequentIssueDetectionApplicationEnabled, true
}

// SetFrequentIssueDetectionApplicationEnabled sets field value
func (o *FrequentIssueDetectionConfig) SetFrequentIssueDetectionApplicationEnabled(v bool) {
	o.FrequentIssueDetectionApplicationEnabled = v
}

// GetFrequentIssueDetectionInfrastructureEnabled returns the FrequentIssueDetectionInfrastructureEnabled field value
func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionInfrastructureEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FrequentIssueDetectionInfrastructureEnabled
}

// GetFrequentIssueDetectionInfrastructureEnabledOk returns a tuple with the FrequentIssueDetectionInfrastructureEnabled field value
// and a boolean to check if the value has been set.
func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionInfrastructureEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrequentIssueDetectionInfrastructureEnabled, true
}

// SetFrequentIssueDetectionInfrastructureEnabled sets field value
func (o *FrequentIssueDetectionConfig) SetFrequentIssueDetectionInfrastructureEnabled(v bool) {
	o.FrequentIssueDetectionInfrastructureEnabled = v
}

// GetFrequentIssueDetectionServiceEnabled returns the FrequentIssueDetectionServiceEnabled field value
func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionServiceEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FrequentIssueDetectionServiceEnabled
}

// GetFrequentIssueDetectionServiceEnabledOk returns a tuple with the FrequentIssueDetectionServiceEnabled field value
// and a boolean to check if the value has been set.
func (o *FrequentIssueDetectionConfig) GetFrequentIssueDetectionServiceEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FrequentIssueDetectionServiceEnabled, true
}

// SetFrequentIssueDetectionServiceEnabled sets field value
func (o *FrequentIssueDetectionConfig) SetFrequentIssueDetectionServiceEnabled(v bool) {
	o.FrequentIssueDetectionServiceEnabled = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *FrequentIssueDetectionConfig) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrequentIssueDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *FrequentIssueDetectionConfig) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *FrequentIssueDetectionConfig) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

func (o FrequentIssueDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrequentIssueDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["frequentIssueDetectionApplicationEnabled"] = o.FrequentIssueDetectionApplicationEnabled
	toSerialize["frequentIssueDetectionInfrastructureEnabled"] = o.FrequentIssueDetectionInfrastructureEnabled
	toSerialize["frequentIssueDetectionServiceEnabled"] = o.FrequentIssueDetectionServiceEnabled
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableFrequentIssueDetectionConfig struct {
	value *FrequentIssueDetectionConfig
	isSet bool
}

func (v NullableFrequentIssueDetectionConfig) Get() *FrequentIssueDetectionConfig {
	return v.value
}

func (v *NullableFrequentIssueDetectionConfig) Set(val *FrequentIssueDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableFrequentIssueDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableFrequentIssueDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrequentIssueDetectionConfig(val *FrequentIssueDetectionConfig) *NullableFrequentIssueDetectionConfig {
	return &NullableFrequentIssueDetectionConfig{value: val, isSet: true}
}

func (v NullableFrequentIssueDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrequentIssueDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

