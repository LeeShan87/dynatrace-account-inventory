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

// checks if the AttackSecurityProblemAssessmentDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttackSecurityProblemAssessmentDto{}

// AttackSecurityProblemAssessmentDto The assessment of a security problem related to an attack.
type AttackSecurityProblemAssessmentDto struct {
	// The reachability of data assets by the attacked target.
	DataAssets *string `json:"dataAssets,omitempty"`
	// The level of exposure of the attacked target
	Exposure *string `json:"exposure,omitempty"`
	// The number of data assets reachable by the attacked target.
	NumberOfReachableDataAssets *int32 `json:"numberOfReachableDataAssets,omitempty"`
}

// NewAttackSecurityProblemAssessmentDto instantiates a new AttackSecurityProblemAssessmentDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackSecurityProblemAssessmentDto() *AttackSecurityProblemAssessmentDto {
	this := AttackSecurityProblemAssessmentDto{}
	return &this
}

// NewAttackSecurityProblemAssessmentDtoWithDefaults instantiates a new AttackSecurityProblemAssessmentDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackSecurityProblemAssessmentDtoWithDefaults() *AttackSecurityProblemAssessmentDto {
	this := AttackSecurityProblemAssessmentDto{}
	return &this
}

// GetDataAssets returns the DataAssets field value if set, zero value otherwise.
func (o *AttackSecurityProblemAssessmentDto) GetDataAssets() string {
	if o == nil || IsNil(o.DataAssets) {
		var ret string
		return ret
	}
	return *o.DataAssets
}

// GetDataAssetsOk returns a tuple with the DataAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackSecurityProblemAssessmentDto) GetDataAssetsOk() (*string, bool) {
	if o == nil || IsNil(o.DataAssets) {
		return nil, false
	}
	return o.DataAssets, true
}

// HasDataAssets returns a boolean if a field has been set.
func (o *AttackSecurityProblemAssessmentDto) HasDataAssets() bool {
	if o != nil && !IsNil(o.DataAssets) {
		return true
	}

	return false
}

// SetDataAssets gets a reference to the given string and assigns it to the DataAssets field.
func (o *AttackSecurityProblemAssessmentDto) SetDataAssets(v string) {
	o.DataAssets = &v
}

// GetExposure returns the Exposure field value if set, zero value otherwise.
func (o *AttackSecurityProblemAssessmentDto) GetExposure() string {
	if o == nil || IsNil(o.Exposure) {
		var ret string
		return ret
	}
	return *o.Exposure
}

// GetExposureOk returns a tuple with the Exposure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackSecurityProblemAssessmentDto) GetExposureOk() (*string, bool) {
	if o == nil || IsNil(o.Exposure) {
		return nil, false
	}
	return o.Exposure, true
}

// HasExposure returns a boolean if a field has been set.
func (o *AttackSecurityProblemAssessmentDto) HasExposure() bool {
	if o != nil && !IsNil(o.Exposure) {
		return true
	}

	return false
}

// SetExposure gets a reference to the given string and assigns it to the Exposure field.
func (o *AttackSecurityProblemAssessmentDto) SetExposure(v string) {
	o.Exposure = &v
}

// GetNumberOfReachableDataAssets returns the NumberOfReachableDataAssets field value if set, zero value otherwise.
func (o *AttackSecurityProblemAssessmentDto) GetNumberOfReachableDataAssets() int32 {
	if o == nil || IsNil(o.NumberOfReachableDataAssets) {
		var ret int32
		return ret
	}
	return *o.NumberOfReachableDataAssets
}

// GetNumberOfReachableDataAssetsOk returns a tuple with the NumberOfReachableDataAssets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttackSecurityProblemAssessmentDto) GetNumberOfReachableDataAssetsOk() (*int32, bool) {
	if o == nil || IsNil(o.NumberOfReachableDataAssets) {
		return nil, false
	}
	return o.NumberOfReachableDataAssets, true
}

// HasNumberOfReachableDataAssets returns a boolean if a field has been set.
func (o *AttackSecurityProblemAssessmentDto) HasNumberOfReachableDataAssets() bool {
	if o != nil && !IsNil(o.NumberOfReachableDataAssets) {
		return true
	}

	return false
}

// SetNumberOfReachableDataAssets gets a reference to the given int32 and assigns it to the NumberOfReachableDataAssets field.
func (o *AttackSecurityProblemAssessmentDto) SetNumberOfReachableDataAssets(v int32) {
	o.NumberOfReachableDataAssets = &v
}

func (o AttackSecurityProblemAssessmentDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttackSecurityProblemAssessmentDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataAssets) {
		toSerialize["dataAssets"] = o.DataAssets
	}
	if !IsNil(o.Exposure) {
		toSerialize["exposure"] = o.Exposure
	}
	if !IsNil(o.NumberOfReachableDataAssets) {
		toSerialize["numberOfReachableDataAssets"] = o.NumberOfReachableDataAssets
	}
	return toSerialize, nil
}

type NullableAttackSecurityProblemAssessmentDto struct {
	value *AttackSecurityProblemAssessmentDto
	isSet bool
}

func (v NullableAttackSecurityProblemAssessmentDto) Get() *AttackSecurityProblemAssessmentDto {
	return v.value
}

func (v *NullableAttackSecurityProblemAssessmentDto) Set(val *AttackSecurityProblemAssessmentDto) {
	v.value = val
	v.isSet = true
}

func (v NullableAttackSecurityProblemAssessmentDto) IsSet() bool {
	return v.isSet
}

func (v *NullableAttackSecurityProblemAssessmentDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttackSecurityProblemAssessmentDto(val *AttackSecurityProblemAssessmentDto) *NullableAttackSecurityProblemAssessmentDto {
	return &NullableAttackSecurityProblemAssessmentDto{value: val, isSet: true}
}

func (v NullableAttackSecurityProblemAssessmentDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttackSecurityProblemAssessmentDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


