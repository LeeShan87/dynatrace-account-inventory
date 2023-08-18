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

// checks if the ApplicationAnomalyDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationAnomalyDetectionConfig{}

// ApplicationAnomalyDetectionConfig The configuration of anomaly detection for applications.
type ApplicationAnomalyDetectionConfig struct {
	FailureRateIncrease FailureRateIncreaseDetectionConfig `json:"failureRateIncrease"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	ResponseTimeDegradation ResponseTimeDegradationDetectionConfig `json:"responseTimeDegradation"`
	TrafficDrop TrafficDropDetectionConfig `json:"trafficDrop"`
	TrafficSpike TrafficSpikeDetectionConfig `json:"trafficSpike"`
}

// NewApplicationAnomalyDetectionConfig instantiates a new ApplicationAnomalyDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationAnomalyDetectionConfig(failureRateIncrease FailureRateIncreaseDetectionConfig, responseTimeDegradation ResponseTimeDegradationDetectionConfig, trafficDrop TrafficDropDetectionConfig, trafficSpike TrafficSpikeDetectionConfig) *ApplicationAnomalyDetectionConfig {
	this := ApplicationAnomalyDetectionConfig{}
	this.FailureRateIncrease = failureRateIncrease
	this.ResponseTimeDegradation = responseTimeDegradation
	this.TrafficDrop = trafficDrop
	this.TrafficSpike = trafficSpike
	return &this
}

// NewApplicationAnomalyDetectionConfigWithDefaults instantiates a new ApplicationAnomalyDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationAnomalyDetectionConfigWithDefaults() *ApplicationAnomalyDetectionConfig {
	this := ApplicationAnomalyDetectionConfig{}
	return &this
}

// GetFailureRateIncrease returns the FailureRateIncrease field value
func (o *ApplicationAnomalyDetectionConfig) GetFailureRateIncrease() FailureRateIncreaseDetectionConfig {
	if o == nil {
		var ret FailureRateIncreaseDetectionConfig
		return ret
	}

	return o.FailureRateIncrease
}

// GetFailureRateIncreaseOk returns a tuple with the FailureRateIncrease field value
// and a boolean to check if the value has been set.
func (o *ApplicationAnomalyDetectionConfig) GetFailureRateIncreaseOk() (*FailureRateIncreaseDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureRateIncrease, true
}

// SetFailureRateIncrease sets field value
func (o *ApplicationAnomalyDetectionConfig) SetFailureRateIncrease(v FailureRateIncreaseDetectionConfig) {
	o.FailureRateIncrease = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApplicationAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApplicationAnomalyDetectionConfig) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *ApplicationAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetResponseTimeDegradation returns the ResponseTimeDegradation field value
func (o *ApplicationAnomalyDetectionConfig) GetResponseTimeDegradation() ResponseTimeDegradationDetectionConfig {
	if o == nil {
		var ret ResponseTimeDegradationDetectionConfig
		return ret
	}

	return o.ResponseTimeDegradation
}

// GetResponseTimeDegradationOk returns a tuple with the ResponseTimeDegradation field value
// and a boolean to check if the value has been set.
func (o *ApplicationAnomalyDetectionConfig) GetResponseTimeDegradationOk() (*ResponseTimeDegradationDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseTimeDegradation, true
}

// SetResponseTimeDegradation sets field value
func (o *ApplicationAnomalyDetectionConfig) SetResponseTimeDegradation(v ResponseTimeDegradationDetectionConfig) {
	o.ResponseTimeDegradation = v
}

// GetTrafficDrop returns the TrafficDrop field value
func (o *ApplicationAnomalyDetectionConfig) GetTrafficDrop() TrafficDropDetectionConfig {
	if o == nil {
		var ret TrafficDropDetectionConfig
		return ret
	}

	return o.TrafficDrop
}

// GetTrafficDropOk returns a tuple with the TrafficDrop field value
// and a boolean to check if the value has been set.
func (o *ApplicationAnomalyDetectionConfig) GetTrafficDropOk() (*TrafficDropDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrafficDrop, true
}

// SetTrafficDrop sets field value
func (o *ApplicationAnomalyDetectionConfig) SetTrafficDrop(v TrafficDropDetectionConfig) {
	o.TrafficDrop = v
}

// GetTrafficSpike returns the TrafficSpike field value
func (o *ApplicationAnomalyDetectionConfig) GetTrafficSpike() TrafficSpikeDetectionConfig {
	if o == nil {
		var ret TrafficSpikeDetectionConfig
		return ret
	}

	return o.TrafficSpike
}

// GetTrafficSpikeOk returns a tuple with the TrafficSpike field value
// and a boolean to check if the value has been set.
func (o *ApplicationAnomalyDetectionConfig) GetTrafficSpikeOk() (*TrafficSpikeDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrafficSpike, true
}

// SetTrafficSpike sets field value
func (o *ApplicationAnomalyDetectionConfig) SetTrafficSpike(v TrafficSpikeDetectionConfig) {
	o.TrafficSpike = v
}

func (o ApplicationAnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationAnomalyDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["failureRateIncrease"] = o.FailureRateIncrease
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["responseTimeDegradation"] = o.ResponseTimeDegradation
	toSerialize["trafficDrop"] = o.TrafficDrop
	toSerialize["trafficSpike"] = o.TrafficSpike
	return toSerialize, nil
}

type NullableApplicationAnomalyDetectionConfig struct {
	value *ApplicationAnomalyDetectionConfig
	isSet bool
}

func (v NullableApplicationAnomalyDetectionConfig) Get() *ApplicationAnomalyDetectionConfig {
	return v.value
}

func (v *NullableApplicationAnomalyDetectionConfig) Set(val *ApplicationAnomalyDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationAnomalyDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationAnomalyDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationAnomalyDetectionConfig(val *ApplicationAnomalyDetectionConfig) *NullableApplicationAnomalyDetectionConfig {
	return &NullableApplicationAnomalyDetectionConfig{value: val, isSet: true}
}

func (v NullableApplicationAnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationAnomalyDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

