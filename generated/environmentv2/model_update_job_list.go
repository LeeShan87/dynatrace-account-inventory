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

// checks if the UpdateJobList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateJobList{}

// UpdateJobList A list of update jobs of the ActiveGate.
type UpdateJobList struct {
	// The ID of the ActiveGate.
	AgId *string `json:"agId,omitempty"`
	// A list of update jobs of the ActiveGate.
	UpdateJobs []UpdateJob `json:"updateJobs,omitempty"`
}

// NewUpdateJobList instantiates a new UpdateJobList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateJobList() *UpdateJobList {
	this := UpdateJobList{}
	return &this
}

// NewUpdateJobListWithDefaults instantiates a new UpdateJobList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateJobListWithDefaults() *UpdateJobList {
	this := UpdateJobList{}
	return &this
}

// GetAgId returns the AgId field value if set, zero value otherwise.
func (o *UpdateJobList) GetAgId() string {
	if o == nil || IsNil(o.AgId) {
		var ret string
		return ret
	}
	return *o.AgId
}

// GetAgIdOk returns a tuple with the AgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateJobList) GetAgIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgId) {
		return nil, false
	}
	return o.AgId, true
}

// HasAgId returns a boolean if a field has been set.
func (o *UpdateJobList) HasAgId() bool {
	if o != nil && !IsNil(o.AgId) {
		return true
	}

	return false
}

// SetAgId gets a reference to the given string and assigns it to the AgId field.
func (o *UpdateJobList) SetAgId(v string) {
	o.AgId = &v
}

// GetUpdateJobs returns the UpdateJobs field value if set, zero value otherwise.
func (o *UpdateJobList) GetUpdateJobs() []UpdateJob {
	if o == nil || IsNil(o.UpdateJobs) {
		var ret []UpdateJob
		return ret
	}
	return o.UpdateJobs
}

// GetUpdateJobsOk returns a tuple with the UpdateJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateJobList) GetUpdateJobsOk() ([]UpdateJob, bool) {
	if o == nil || IsNil(o.UpdateJobs) {
		return nil, false
	}
	return o.UpdateJobs, true
}

// HasUpdateJobs returns a boolean if a field has been set.
func (o *UpdateJobList) HasUpdateJobs() bool {
	if o != nil && !IsNil(o.UpdateJobs) {
		return true
	}

	return false
}

// SetUpdateJobs gets a reference to the given []UpdateJob and assigns it to the UpdateJobs field.
func (o *UpdateJobList) SetUpdateJobs(v []UpdateJob) {
	o.UpdateJobs = v
}

func (o UpdateJobList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateJobList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgId) {
		toSerialize["agId"] = o.AgId
	}
	if !IsNil(o.UpdateJobs) {
		toSerialize["updateJobs"] = o.UpdateJobs
	}
	return toSerialize, nil
}

type NullableUpdateJobList struct {
	value *UpdateJobList
	isSet bool
}

func (v NullableUpdateJobList) Get() *UpdateJobList {
	return v.value
}

func (v *NullableUpdateJobList) Set(val *UpdateJobList) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateJobList) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateJobList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateJobList(val *UpdateJobList) *NullableUpdateJobList {
	return &NullableUpdateJobList{value: val, isSet: true}
}

func (v NullableUpdateJobList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateJobList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


