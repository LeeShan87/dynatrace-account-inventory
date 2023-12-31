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

// checks if the AwsAnomalyDetectionConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsAnomalyDetectionConfig{}

// AwsAnomalyDetectionConfig The configuration of anomaly detection for AWS.
type AwsAnomalyDetectionConfig struct {
	Ec2CandidateCpuSaturationDetection *Ec2CandidateCpuSaturationDetectionConfig `json:"ec2CandidateCpuSaturationDetection,omitempty"`
	ElbHighConnectionErrorsDetection ElbHighConnectionErrorsDetectionConfig `json:"elbHighConnectionErrorsDetection"`
	LambdaHighErrorRateDetection LambdaHighErrorRateDetectionConfig `json:"lambdaHighErrorRateDetection"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	RdsHighCpuDetection RdsHighCpuDetectionConfig `json:"rdsHighCpuDetection"`
	RdsHighMemoryDetection RdsHighMemoryDetectionConfig `json:"rdsHighMemoryDetection"`
	RdsHighWriteReadLatencyDetection RdsHighWriteReadLatencyDetectionConfig `json:"rdsHighWriteReadLatencyDetection"`
	RdsLowStorageDetection RdsLowStorageDetectionConfig `json:"rdsLowStorageDetection"`
	RdsRestartsSequenceDetection RdsRestartsSequenceDetectionConfig `json:"rdsRestartsSequenceDetection"`
}

// NewAwsAnomalyDetectionConfig instantiates a new AwsAnomalyDetectionConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsAnomalyDetectionConfig(elbHighConnectionErrorsDetection ElbHighConnectionErrorsDetectionConfig, lambdaHighErrorRateDetection LambdaHighErrorRateDetectionConfig, rdsHighCpuDetection RdsHighCpuDetectionConfig, rdsHighMemoryDetection RdsHighMemoryDetectionConfig, rdsHighWriteReadLatencyDetection RdsHighWriteReadLatencyDetectionConfig, rdsLowStorageDetection RdsLowStorageDetectionConfig, rdsRestartsSequenceDetection RdsRestartsSequenceDetectionConfig) *AwsAnomalyDetectionConfig {
	this := AwsAnomalyDetectionConfig{}
	this.ElbHighConnectionErrorsDetection = elbHighConnectionErrorsDetection
	this.LambdaHighErrorRateDetection = lambdaHighErrorRateDetection
	this.RdsHighCpuDetection = rdsHighCpuDetection
	this.RdsHighMemoryDetection = rdsHighMemoryDetection
	this.RdsHighWriteReadLatencyDetection = rdsHighWriteReadLatencyDetection
	this.RdsLowStorageDetection = rdsLowStorageDetection
	this.RdsRestartsSequenceDetection = rdsRestartsSequenceDetection
	return &this
}

// NewAwsAnomalyDetectionConfigWithDefaults instantiates a new AwsAnomalyDetectionConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsAnomalyDetectionConfigWithDefaults() *AwsAnomalyDetectionConfig {
	this := AwsAnomalyDetectionConfig{}
	return &this
}

// GetEc2CandidateCpuSaturationDetection returns the Ec2CandidateCpuSaturationDetection field value if set, zero value otherwise.
func (o *AwsAnomalyDetectionConfig) GetEc2CandidateCpuSaturationDetection() Ec2CandidateCpuSaturationDetectionConfig {
	if o == nil || IsNil(o.Ec2CandidateCpuSaturationDetection) {
		var ret Ec2CandidateCpuSaturationDetectionConfig
		return ret
	}
	return *o.Ec2CandidateCpuSaturationDetection
}

// GetEc2CandidateCpuSaturationDetectionOk returns a tuple with the Ec2CandidateCpuSaturationDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAnomalyDetectionConfig) GetEc2CandidateCpuSaturationDetectionOk() (*Ec2CandidateCpuSaturationDetectionConfig, bool) {
	if o == nil || IsNil(o.Ec2CandidateCpuSaturationDetection) {
		return nil, false
	}
	return o.Ec2CandidateCpuSaturationDetection, true
}

// HasEc2CandidateCpuSaturationDetection returns a boolean if a field has been set.
func (o *AwsAnomalyDetectionConfig) HasEc2CandidateCpuSaturationDetection() bool {
	if o != nil && !IsNil(o.Ec2CandidateCpuSaturationDetection) {
		return true
	}

	return false
}

// SetEc2CandidateCpuSaturationDetection gets a reference to the given Ec2CandidateCpuSaturationDetectionConfig and assigns it to the Ec2CandidateCpuSaturationDetection field.
func (o *AwsAnomalyDetectionConfig) SetEc2CandidateCpuSaturationDetection(v Ec2CandidateCpuSaturationDetectionConfig) {
	o.Ec2CandidateCpuSaturationDetection = &v
}

// GetElbHighConnectionErrorsDetection returns the ElbHighConnectionErrorsDetection field value
func (o *AwsAnomalyDetectionConfig) GetElbHighConnectionErrorsDetection() ElbHighConnectionErrorsDetectionConfig {
	if o == nil {
		var ret ElbHighConnectionErrorsDetectionConfig
		return ret
	}

	return o.ElbHighConnectionErrorsDetection
}

// GetElbHighConnectionErrorsDetectionOk returns a tuple with the ElbHighConnectionErrorsDetection field value
// and a boolean to check if the value has been set.
func (o *AwsAnomalyDetectionConfig) GetElbHighConnectionErrorsDetectionOk() (*ElbHighConnectionErrorsDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ElbHighConnectionErrorsDetection, true
}

// SetElbHighConnectionErrorsDetection sets field value
func (o *AwsAnomalyDetectionConfig) SetElbHighConnectionErrorsDetection(v ElbHighConnectionErrorsDetectionConfig) {
	o.ElbHighConnectionErrorsDetection = v
}

// GetLambdaHighErrorRateDetection returns the LambdaHighErrorRateDetection field value
func (o *AwsAnomalyDetectionConfig) GetLambdaHighErrorRateDetection() LambdaHighErrorRateDetectionConfig {
	if o == nil {
		var ret LambdaHighErrorRateDetectionConfig
		return ret
	}

	return o.LambdaHighErrorRateDetection
}

// GetLambdaHighErrorRateDetectionOk returns a tuple with the LambdaHighErrorRateDetection field value
// and a boolean to check if the value has been set.
func (o *AwsAnomalyDetectionConfig) GetLambdaHighErrorRateDetectionOk() (*LambdaHighErrorRateDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LambdaHighErrorRateDetection, true
}

// SetLambdaHighErrorRateDetection sets field value
func (o *AwsAnomalyDetectionConfig) SetLambdaHighErrorRateDetection(v LambdaHighErrorRateDetectionConfig) {
	o.LambdaHighErrorRateDetection = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AwsAnomalyDetectionConfig) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAnomalyDetectionConfig) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AwsAnomalyDetectionConfig) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *AwsAnomalyDetectionConfig) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetRdsHighCpuDetection returns the RdsHighCpuDetection field value
func (o *AwsAnomalyDetectionConfig) GetRdsHighCpuDetection() RdsHighCpuDetectionConfig {
	if o == nil {
		var ret RdsHighCpuDetectionConfig
		return ret
	}

	return o.RdsHighCpuDetection
}

// GetRdsHighCpuDetectionOk returns a tuple with the RdsHighCpuDetection field value
// and a boolean to check if the value has been set.
func (o *AwsAnomalyDetectionConfig) GetRdsHighCpuDetectionOk() (*RdsHighCpuDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RdsHighCpuDetection, true
}

// SetRdsHighCpuDetection sets field value
func (o *AwsAnomalyDetectionConfig) SetRdsHighCpuDetection(v RdsHighCpuDetectionConfig) {
	o.RdsHighCpuDetection = v
}

// GetRdsHighMemoryDetection returns the RdsHighMemoryDetection field value
func (o *AwsAnomalyDetectionConfig) GetRdsHighMemoryDetection() RdsHighMemoryDetectionConfig {
	if o == nil {
		var ret RdsHighMemoryDetectionConfig
		return ret
	}

	return o.RdsHighMemoryDetection
}

// GetRdsHighMemoryDetectionOk returns a tuple with the RdsHighMemoryDetection field value
// and a boolean to check if the value has been set.
func (o *AwsAnomalyDetectionConfig) GetRdsHighMemoryDetectionOk() (*RdsHighMemoryDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RdsHighMemoryDetection, true
}

// SetRdsHighMemoryDetection sets field value
func (o *AwsAnomalyDetectionConfig) SetRdsHighMemoryDetection(v RdsHighMemoryDetectionConfig) {
	o.RdsHighMemoryDetection = v
}

// GetRdsHighWriteReadLatencyDetection returns the RdsHighWriteReadLatencyDetection field value
func (o *AwsAnomalyDetectionConfig) GetRdsHighWriteReadLatencyDetection() RdsHighWriteReadLatencyDetectionConfig {
	if o == nil {
		var ret RdsHighWriteReadLatencyDetectionConfig
		return ret
	}

	return o.RdsHighWriteReadLatencyDetection
}

// GetRdsHighWriteReadLatencyDetectionOk returns a tuple with the RdsHighWriteReadLatencyDetection field value
// and a boolean to check if the value has been set.
func (o *AwsAnomalyDetectionConfig) GetRdsHighWriteReadLatencyDetectionOk() (*RdsHighWriteReadLatencyDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RdsHighWriteReadLatencyDetection, true
}

// SetRdsHighWriteReadLatencyDetection sets field value
func (o *AwsAnomalyDetectionConfig) SetRdsHighWriteReadLatencyDetection(v RdsHighWriteReadLatencyDetectionConfig) {
	o.RdsHighWriteReadLatencyDetection = v
}

// GetRdsLowStorageDetection returns the RdsLowStorageDetection field value
func (o *AwsAnomalyDetectionConfig) GetRdsLowStorageDetection() RdsLowStorageDetectionConfig {
	if o == nil {
		var ret RdsLowStorageDetectionConfig
		return ret
	}

	return o.RdsLowStorageDetection
}

// GetRdsLowStorageDetectionOk returns a tuple with the RdsLowStorageDetection field value
// and a boolean to check if the value has been set.
func (o *AwsAnomalyDetectionConfig) GetRdsLowStorageDetectionOk() (*RdsLowStorageDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RdsLowStorageDetection, true
}

// SetRdsLowStorageDetection sets field value
func (o *AwsAnomalyDetectionConfig) SetRdsLowStorageDetection(v RdsLowStorageDetectionConfig) {
	o.RdsLowStorageDetection = v
}

// GetRdsRestartsSequenceDetection returns the RdsRestartsSequenceDetection field value
func (o *AwsAnomalyDetectionConfig) GetRdsRestartsSequenceDetection() RdsRestartsSequenceDetectionConfig {
	if o == nil {
		var ret RdsRestartsSequenceDetectionConfig
		return ret
	}

	return o.RdsRestartsSequenceDetection
}

// GetRdsRestartsSequenceDetectionOk returns a tuple with the RdsRestartsSequenceDetection field value
// and a boolean to check if the value has been set.
func (o *AwsAnomalyDetectionConfig) GetRdsRestartsSequenceDetectionOk() (*RdsRestartsSequenceDetectionConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RdsRestartsSequenceDetection, true
}

// SetRdsRestartsSequenceDetection sets field value
func (o *AwsAnomalyDetectionConfig) SetRdsRestartsSequenceDetection(v RdsRestartsSequenceDetectionConfig) {
	o.RdsRestartsSequenceDetection = v
}

func (o AwsAnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsAnomalyDetectionConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ec2CandidateCpuSaturationDetection) {
		toSerialize["ec2CandidateCpuSaturationDetection"] = o.Ec2CandidateCpuSaturationDetection
	}
	toSerialize["elbHighConnectionErrorsDetection"] = o.ElbHighConnectionErrorsDetection
	toSerialize["lambdaHighErrorRateDetection"] = o.LambdaHighErrorRateDetection
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["rdsHighCpuDetection"] = o.RdsHighCpuDetection
	toSerialize["rdsHighMemoryDetection"] = o.RdsHighMemoryDetection
	toSerialize["rdsHighWriteReadLatencyDetection"] = o.RdsHighWriteReadLatencyDetection
	toSerialize["rdsLowStorageDetection"] = o.RdsLowStorageDetection
	toSerialize["rdsRestartsSequenceDetection"] = o.RdsRestartsSequenceDetection
	return toSerialize, nil
}

type NullableAwsAnomalyDetectionConfig struct {
	value *AwsAnomalyDetectionConfig
	isSet bool
}

func (v NullableAwsAnomalyDetectionConfig) Get() *AwsAnomalyDetectionConfig {
	return v.value
}

func (v *NullableAwsAnomalyDetectionConfig) Set(val *AwsAnomalyDetectionConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsAnomalyDetectionConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsAnomalyDetectionConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsAnomalyDetectionConfig(val *AwsAnomalyDetectionConfig) *NullableAwsAnomalyDetectionConfig {
	return &NullableAwsAnomalyDetectionConfig{value: val, isSet: true}
}

func (v NullableAwsAnomalyDetectionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsAnomalyDetectionConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


