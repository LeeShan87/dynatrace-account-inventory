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

// checks if the AwsMonitoredServicesDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsMonitoredServicesDto{}

// AwsMonitoredServicesDto struct for AwsMonitoredServicesDto
type AwsMonitoredServicesDto struct {
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// A list of AWS services to be monitored. Available services are listed by [/aws/supportedServices](https://dt-url.net/me02sh2) operation.  For each service, a list of metrics and dimensions can be specified. A list of supported metrics and dimensions for a given service can be checked in [documentation](https://dt-url.net/r12v0pkl).  List of metrics can be skipped (set to null), resulting in recommended (default) set of metrics and dimensions being chosen for monitoring. For built-in services, adjusting the list of metrics is not supported, therefore it needs to be null.
	Services []AwsSupportingServiceConfig `json:"services"`
}

// NewAwsMonitoredServicesDto instantiates a new AwsMonitoredServicesDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsMonitoredServicesDto(services []AwsSupportingServiceConfig) *AwsMonitoredServicesDto {
	this := AwsMonitoredServicesDto{}
	this.Services = services
	return &this
}

// NewAwsMonitoredServicesDtoWithDefaults instantiates a new AwsMonitoredServicesDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsMonitoredServicesDtoWithDefaults() *AwsMonitoredServicesDto {
	this := AwsMonitoredServicesDto{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *AwsMonitoredServicesDto) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsMonitoredServicesDto) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *AwsMonitoredServicesDto) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *AwsMonitoredServicesDto) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetServices returns the Services field value
func (o *AwsMonitoredServicesDto) GetServices() []AwsSupportingServiceConfig {
	if o == nil {
		var ret []AwsSupportingServiceConfig
		return ret
	}

	return o.Services
}

// GetServicesOk returns a tuple with the Services field value
// and a boolean to check if the value has been set.
func (o *AwsMonitoredServicesDto) GetServicesOk() ([]AwsSupportingServiceConfig, bool) {
	if o == nil {
		return nil, false
	}
	return o.Services, true
}

// SetServices sets field value
func (o *AwsMonitoredServicesDto) SetServices(v []AwsSupportingServiceConfig) {
	o.Services = v
}

func (o AwsMonitoredServicesDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsMonitoredServicesDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["services"] = o.Services
	return toSerialize, nil
}

type NullableAwsMonitoredServicesDto struct {
	value *AwsMonitoredServicesDto
	isSet bool
}

func (v NullableAwsMonitoredServicesDto) Get() *AwsMonitoredServicesDto {
	return v.value
}

func (v *NullableAwsMonitoredServicesDto) Set(val *AwsMonitoredServicesDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsMonitoredServicesDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsMonitoredServicesDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsMonitoredServicesDto(val *AwsMonitoredServicesDto) *NullableAwsMonitoredServicesDto {
	return &NullableAwsMonitoredServicesDto{value: val, isSet: true}
}

func (v NullableAwsMonitoredServicesDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsMonitoredServicesDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

