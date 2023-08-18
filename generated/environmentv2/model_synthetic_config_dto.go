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

// checks if the SyntheticConfigDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyntheticConfigDto{}

// SyntheticConfigDto A DTO for synthetic configuration.
type SyntheticConfigDto struct {
	// bmMonitorTimeout - browser monitor execution timeout (ms)
	BmMonitorTimeout int64 `json:"bmMonitorTimeout"`
	// bmStepTimeout - browser monitor single step execution timeout (ms)
	BmStepTimeout int64 `json:"bmStepTimeout"`
}

// NewSyntheticConfigDto instantiates a new SyntheticConfigDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticConfigDto(bmMonitorTimeout int64, bmStepTimeout int64) *SyntheticConfigDto {
	this := SyntheticConfigDto{}
	this.BmMonitorTimeout = bmMonitorTimeout
	this.BmStepTimeout = bmStepTimeout
	return &this
}

// NewSyntheticConfigDtoWithDefaults instantiates a new SyntheticConfigDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticConfigDtoWithDefaults() *SyntheticConfigDto {
	this := SyntheticConfigDto{}
	return &this
}

// GetBmMonitorTimeout returns the BmMonitorTimeout field value
func (o *SyntheticConfigDto) GetBmMonitorTimeout() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BmMonitorTimeout
}

// GetBmMonitorTimeoutOk returns a tuple with the BmMonitorTimeout field value
// and a boolean to check if the value has been set.
func (o *SyntheticConfigDto) GetBmMonitorTimeoutOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BmMonitorTimeout, true
}

// SetBmMonitorTimeout sets field value
func (o *SyntheticConfigDto) SetBmMonitorTimeout(v int64) {
	o.BmMonitorTimeout = v
}

// GetBmStepTimeout returns the BmStepTimeout field value
func (o *SyntheticConfigDto) GetBmStepTimeout() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.BmStepTimeout
}

// GetBmStepTimeoutOk returns a tuple with the BmStepTimeout field value
// and a boolean to check if the value has been set.
func (o *SyntheticConfigDto) GetBmStepTimeoutOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BmStepTimeout, true
}

// SetBmStepTimeout sets field value
func (o *SyntheticConfigDto) SetBmStepTimeout(v int64) {
	o.BmStepTimeout = v
}

func (o SyntheticConfigDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyntheticConfigDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bmMonitorTimeout"] = o.BmMonitorTimeout
	toSerialize["bmStepTimeout"] = o.BmStepTimeout
	return toSerialize, nil
}

type NullableSyntheticConfigDto struct {
	value *SyntheticConfigDto
	isSet bool
}

func (v NullableSyntheticConfigDto) Get() *SyntheticConfigDto {
	return v.value
}

func (v *NullableSyntheticConfigDto) Set(val *SyntheticConfigDto) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticConfigDto) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticConfigDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticConfigDto(val *SyntheticConfigDto) *NullableSyntheticConfigDto {
	return &NullableSyntheticConfigDto{value: val, isSet: true}
}

func (v NullableSyntheticConfigDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticConfigDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

