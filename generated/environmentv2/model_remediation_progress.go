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

// checks if the RemediationProgress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemediationProgress{}

// RemediationProgress The progress of this remediation item. It contains affected and unaffected entities.
type RemediationProgress struct {
	// A list of related entities that are affected by the security problem.
	AffectedEntities []string `json:"affectedEntities,omitempty"`
	// A list of related entities that are affected by the security problem.
	UnaffectedEntities []string `json:"unaffectedEntities,omitempty"`
}

// NewRemediationProgress instantiates a new RemediationProgress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemediationProgress() *RemediationProgress {
	this := RemediationProgress{}
	return &this
}

// NewRemediationProgressWithDefaults instantiates a new RemediationProgress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemediationProgressWithDefaults() *RemediationProgress {
	this := RemediationProgress{}
	return &this
}

// GetAffectedEntities returns the AffectedEntities field value if set, zero value otherwise.
func (o *RemediationProgress) GetAffectedEntities() []string {
	if o == nil || IsNil(o.AffectedEntities) {
		var ret []string
		return ret
	}
	return o.AffectedEntities
}

// GetAffectedEntitiesOk returns a tuple with the AffectedEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProgress) GetAffectedEntitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.AffectedEntities) {
		return nil, false
	}
	return o.AffectedEntities, true
}

// HasAffectedEntities returns a boolean if a field has been set.
func (o *RemediationProgress) HasAffectedEntities() bool {
	if o != nil && !IsNil(o.AffectedEntities) {
		return true
	}

	return false
}

// SetAffectedEntities gets a reference to the given []string and assigns it to the AffectedEntities field.
func (o *RemediationProgress) SetAffectedEntities(v []string) {
	o.AffectedEntities = v
}

// GetUnaffectedEntities returns the UnaffectedEntities field value if set, zero value otherwise.
func (o *RemediationProgress) GetUnaffectedEntities() []string {
	if o == nil || IsNil(o.UnaffectedEntities) {
		var ret []string
		return ret
	}
	return o.UnaffectedEntities
}

// GetUnaffectedEntitiesOk returns a tuple with the UnaffectedEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationProgress) GetUnaffectedEntitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.UnaffectedEntities) {
		return nil, false
	}
	return o.UnaffectedEntities, true
}

// HasUnaffectedEntities returns a boolean if a field has been set.
func (o *RemediationProgress) HasUnaffectedEntities() bool {
	if o != nil && !IsNil(o.UnaffectedEntities) {
		return true
	}

	return false
}

// SetUnaffectedEntities gets a reference to the given []string and assigns it to the UnaffectedEntities field.
func (o *RemediationProgress) SetUnaffectedEntities(v []string) {
	o.UnaffectedEntities = v
}

func (o RemediationProgress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemediationProgress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffectedEntities) {
		toSerialize["affectedEntities"] = o.AffectedEntities
	}
	if !IsNil(o.UnaffectedEntities) {
		toSerialize["unaffectedEntities"] = o.UnaffectedEntities
	}
	return toSerialize, nil
}

type NullableRemediationProgress struct {
	value *RemediationProgress
	isSet bool
}

func (v NullableRemediationProgress) Get() *RemediationProgress {
	return v.value
}

func (v *NullableRemediationProgress) Set(val *RemediationProgress) {
	v.value = val
	v.isSet = true
}

func (v NullableRemediationProgress) IsSet() bool {
	return v.isSet
}

func (v *NullableRemediationProgress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemediationProgress(val *RemediationProgress) *NullableRemediationProgress {
	return &NullableRemediationProgress{value: val, isSet: true}
}

func (v NullableRemediationProgress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemediationProgress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


