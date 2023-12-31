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

// checks if the AvailabilityEvidence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvailabilityEvidence{}

// AvailabilityEvidence The availability evidence of the problem.   Indicates an entity that has been unavailable during the problem lifespan and that might be related to the root cause.
type AvailabilityEvidence struct {
	// The end time of the evidence, in UTC milliseconds.
	EndTime int64 `json:"endTime"`
}

// NewAvailabilityEvidence instantiates a new AvailabilityEvidence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailabilityEvidence(endTime int64, displayName string, entity EntityStub, evidenceType string, rootCauseRelevant bool, startTime int64) *AvailabilityEvidence {
	this := AvailabilityEvidence{}
	this.DisplayName = displayName
	this.Entity = entity
	this.EvidenceType = evidenceType
	this.RootCauseRelevant = rootCauseRelevant
	this.StartTime = startTime
	this.EndTime = endTime
	return &this
}

// NewAvailabilityEvidenceWithDefaults instantiates a new AvailabilityEvidence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityEvidenceWithDefaults() *AvailabilityEvidence {
	this := AvailabilityEvidence{}
	return &this
}

// GetEndTime returns the EndTime field value
func (o *AvailabilityEvidence) GetEndTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *AvailabilityEvidence) GetEndTimeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *AvailabilityEvidence) SetEndTime(v int64) {
	o.EndTime = v
}

func (o AvailabilityEvidence) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvailabilityEvidence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["endTime"] = o.EndTime
	return toSerialize, nil
}

type NullableAvailabilityEvidence struct {
	value *AvailabilityEvidence
	isSet bool
}

func (v NullableAvailabilityEvidence) Get() *AvailabilityEvidence {
	return v.value
}

func (v *NullableAvailabilityEvidence) Set(val *AvailabilityEvidence) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailabilityEvidence) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailabilityEvidence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailabilityEvidence(val *AvailabilityEvidence) *NullableAvailabilityEvidence {
	return &NullableAvailabilityEvidence{value: val, isSet: true}
}

func (v NullableAvailabilityEvidence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailabilityEvidence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


