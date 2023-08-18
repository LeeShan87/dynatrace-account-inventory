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

// checks if the DatabaseAnomalyDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatabaseAnomalyDetectionConfig{}

// DatabaseAnomalyDetectionConfig The configuration of the anomaly detection for database services.
type DatabaseAnomalyDetectionConfig struct {
	DatabaseConnectionFailureCount DatabaseConnectionFailureDetectionConfig `json:"databaseConnectionFailureCount"`
	FailureRateIncrease FailureRateIncreaseDetectionConfig `json:"failureRateIncrease"`
	LoadDrop *LoadDropDetectionConfig `json:"loadDrop,omitempty"`
	LoadSpike *LoadSpikeDetectionConfig `json:"loadSpike,omitempty"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	ResponseTimeDegradation ResponseTimeDegradationDetectionConfig `json:"responseTimeDegradation"`
}

// NewDatabaseAnomalyDetectionConfig instantiates a new DatabaseAnomalyDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatabaseAnomalyDetectionConfig(databaseConnectionFailureCount DatabaseConnectionFailureDetectionConfig, failureRateIncrease FailureRateIncreaseDetectionConfig, responseTimeDegradation ResponseTimeDegradationDetectionConfig) *DatabaseAnomalyDetectionConfig {
	this := DatabaseAnomalyDetectionConfig{}
	this.DatabaseConnectionFailureCount = databaseConnectionFailureCount
	this.FailureRateIncrease = failureRateIncrease
	this.ResponseTimeDegradation = responseTimeDegradation
	return &this
}

// NewDatabaseAnomalyDetectionConfigWithDefaults instantiates a new DatabaseAnomalyDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatabaseAnomalyDetectionConfigWithDefaults() *DatabaseAnomalyDetectionConfig {
	this := DatabaseAnomalyDetectionConfig{}
	return &this
}

// GetDatabaseConnectionFailureCount returns the DatabaseConnectionFailureCount field value
func (o *DatabaseAnomalyDetectionConfig) GetDatabaseConnectionFailureCount() DatabaseConnectionFailureDetectionConfig {
	if o == nil {
		var ret DatabaseConnectionFailureDetectionConfig
		return ret
	}

	return o.DatabaseConnectionFailureCount
}

// GetDatabaseConnectionFailureCountOk returns a tuple with the DatabaseConnectionFailureCount field value
// and a boolean to check if the value has been set.
func (o *DatabaseAnomalyDetectionConfig) GetDatabaseConnectionFailureCountOk() (*DatabaseConnectionFailureDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseConnectionFailureCount, true
}

// SetDatabaseConnectionFailureCount sets field value
func (o *DatabaseAnomalyDetectionConfig) SetDatabaseConnectionFailureCount(v DatabaseConnectionFailureDetectionConfig) {
	o.DatabaseConnectionFailureCount = v
}

// GetFailureRateIncrease returns the FailureRateIncrease field value
func (o *DatabaseAnomalyDetectionConfig) GetFailureRateIncrease() FailureRateIncreaseDetectionConfig {
	if o == nil {
		var ret FailureRateIncreaseDetectionConfig
		return ret
	}

	return o.FailureRateIncrease
}

// GetFailureRateIncreaseOk returns a tuple with the FailureRateIncrease field value
// and a boolean to check if the value has been set.
func (o *DatabaseAnomalyDetectionConfig) GetFailureRateIncreaseOk() (*FailureRateIncreaseDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureRateIncrease, true
}

// SetFailureRateIncrease sets field value
func (o *DatabaseAnomalyDetectionConfig) SetFailureRateIncrease(v FailureRateIncreaseDetectionConfig) {
	o.FailureRateIncrease = v
}

// GetLoadDrop returns the LoadDrop field value if set, zero value otherwise.
func (o *DatabaseAnomalyDetectionConfig) GetLoadDrop() LoadDropDetectionConfig {
	if o == nil || IsNil(o.LoadDrop) {
		var ret LoadDropDetectionConfig
		return ret
	}
	return *o.LoadDrop
}

// GetLoadDropOk returns a tuple with the LoadDrop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseAnomalyDetectionConfig) GetLoadDropOk() (*LoadDropDetectionConfig, bool) {
	if o == nil || IsNil(o.LoadDrop) {
		return nil, false
	}
	return o.LoadDrop, true
}

// HasLoadDrop returns a boolean if a field has been set.
func (o *DatabaseAnomalyDetectionConfig) HasLoadDrop() bool {
	if o != nil && !IsNil(o.LoadDrop) {
		return true
	}

	return false
}

// SetLoadDrop gets a reference to the given LoadDropDetectionConfig and assigns it to the LoadDrop field.
func (o *DatabaseAnomalyDetectionConfig) SetLoadDrop(v LoadDropDetectionConfig) {
	o.LoadDrop = &v
}

// GetLoadSpike returns the LoadSpike field value if set, zero value otherwise.
func (o *DatabaseAnomalyDetectionConfig) GetLoadSpike() LoadSpikeDetectionConfig {
	if o == nil || IsNil(o.LoadSpike) {
		var ret LoadSpikeDetectionConfig
		return ret
	}
	return *o.LoadSpike
}

// GetLoadSpikeOk returns a tuple with the LoadSpike field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseAnomalyDetectionConfig) GetLoadSpikeOk() (*LoadSpikeDetectionConfig, bool) {
	if o == nil || IsNil(o.LoadSpike) {
		return nil, false
	}
	return o.LoadSpike, true
}

// HasLoadSpike returns a boolean if a field has been set.
func (o *DatabaseAnomalyDetectionConfig) HasLoadSpike() bool {
	if o != nil && !IsNil(o.LoadSpike) {
		return true
	}

	return false
}

// SetLoadSpike gets a reference to the given LoadSpikeDetectionConfig and assigns it to the LoadSpike field.
func (o *DatabaseAnomalyDetectionConfig) SetLoadSpike(v LoadSpikeDetectionConfig) {
	o.LoadSpike = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DatabaseAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatabaseAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DatabaseAnomalyDetectionConfig) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *DatabaseAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetResponseTimeDegradation returns the ResponseTimeDegradation field value
func (o *DatabaseAnomalyDetectionConfig) GetResponseTimeDegradation() ResponseTimeDegradationDetectionConfig {
	if o == nil {
		var ret ResponseTimeDegradationDetectionConfig
		return ret
	}

	return o.ResponseTimeDegradation
}

// GetResponseTimeDegradationOk returns a tuple with the ResponseTimeDegradation field value
// and a boolean to check if the value has been set.
func (o *DatabaseAnomalyDetectionConfig) GetResponseTimeDegradationOk() (*ResponseTimeDegradationDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseTimeDegradation, true
}

// SetResponseTimeDegradation sets field value
func (o *DatabaseAnomalyDetectionConfig) SetResponseTimeDegradation(v ResponseTimeDegradationDetectionConfig) {
	o.ResponseTimeDegradation = v
}

func (o DatabaseAnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatabaseAnomalyDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["databaseConnectionFailureCount"] = o.DatabaseConnectionFailureCount
	toSerialize["failureRateIncrease"] = o.FailureRateIncrease
	if !IsNil(o.LoadDrop) {
		toSerialize["loadDrop"] = o.LoadDrop
	}
	if !IsNil(o.LoadSpike) {
		toSerialize["loadSpike"] = o.LoadSpike
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["responseTimeDegradation"] = o.ResponseTimeDegradation
	return toSerialize, nil
}

type NullableDatabaseAnomalyDetectionConfig struct {
	value *DatabaseAnomalyDetectionConfig
	isSet bool
}

func (v NullableDatabaseAnomalyDetectionConfig) Get() *DatabaseAnomalyDetectionConfig {
	return v.value
}

func (v *NullableDatabaseAnomalyDetectionConfig) Set(val *DatabaseAnomalyDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDatabaseAnomalyDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDatabaseAnomalyDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatabaseAnomalyDetectionConfig(val *DatabaseAnomalyDetectionConfig) *NullableDatabaseAnomalyDetectionConfig {
	return &NullableDatabaseAnomalyDetectionConfig{value: val, isSet: true}
}

func (v NullableDatabaseAnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatabaseAnomalyDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

