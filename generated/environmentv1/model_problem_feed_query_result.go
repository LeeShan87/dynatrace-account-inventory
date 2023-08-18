/*
Dynatrace Environment API

Documentation of the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/xc03k3c).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv1

import (
	"encoding/json"
)

// checks if the ProblemFeedQueryResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProblemFeedQueryResult{}

// ProblemFeedQueryResult Details on open problems in your environment.
type ProblemFeedQueryResult struct {
	Monitored *ProblemFeedQueryResultMonitored `json:"monitored,omitempty"`
	// The list of problems and their details.   Contains all problems within specified timeframe, open and closed.
	Problems []Problem `json:"problems,omitempty"`
}

// NewProblemFeedQueryResult instantiates a new ProblemFeedQueryResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProblemFeedQueryResult() *ProblemFeedQueryResult {
	this := ProblemFeedQueryResult{}
	return &this
}

// NewProblemFeedQueryResultWithDefaults instantiates a new ProblemFeedQueryResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProblemFeedQueryResultWithDefaults() *ProblemFeedQueryResult {
	this := ProblemFeedQueryResult{}
	return &this
}

// GetMonitored returns the Monitored field value if set, zero value otherwise.
func (o *ProblemFeedQueryResult) GetMonitored() ProblemFeedQueryResultMonitored {
	if o == nil || IsNil(o.Monitored) {
		var ret ProblemFeedQueryResultMonitored
		return ret
	}
	return *o.Monitored
}

// GetMonitoredOk returns a tuple with the Monitored field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemFeedQueryResult) GetMonitoredOk() (*ProblemFeedQueryResultMonitored, bool) {
	if o == nil || IsNil(o.Monitored) {
		return nil, false
	}
	return o.Monitored, true
}

// HasMonitored returns a boolean if a field has been set.
func (o *ProblemFeedQueryResult) HasMonitored() bool {
	if o != nil && !IsNil(o.Monitored) {
		return true
	}

	return false
}

// SetMonitored gets a reference to the given ProblemFeedQueryResultMonitored and assigns it to the Monitored field.
func (o *ProblemFeedQueryResult) SetMonitored(v ProblemFeedQueryResultMonitored) {
	o.Monitored = &v
}

// GetProblems returns the Problems field value if set, zero value otherwise.
func (o *ProblemFeedQueryResult) GetProblems() []Problem {
	if o == nil || IsNil(o.Problems) {
		var ret []Problem
		return ret
	}
	return o.Problems
}

// GetProblemsOk returns a tuple with the Problems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProblemFeedQueryResult) GetProblemsOk() ([]Problem, bool) {
	if o == nil || IsNil(o.Problems) {
		return nil, false
	}
	return o.Problems, true
}

// HasProblems returns a boolean if a field has been set.
func (o *ProblemFeedQueryResult) HasProblems() bool {
	if o != nil && !IsNil(o.Problems) {
		return true
	}

	return false
}

// SetProblems gets a reference to the given []Problem and assigns it to the Problems field.
func (o *ProblemFeedQueryResult) SetProblems(v []Problem) {
	o.Problems = v
}

func (o ProblemFeedQueryResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProblemFeedQueryResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Monitored) {
		toSerialize["monitored"] = o.Monitored
	}
	if !IsNil(o.Problems) {
		toSerialize["problems"] = o.Problems
	}
	return toSerialize, nil
}

type NullableProblemFeedQueryResult struct {
	value *ProblemFeedQueryResult
	isSet bool
}

func (v NullableProblemFeedQueryResult) Get() *ProblemFeedQueryResult {
	return v.value
}

func (v *NullableProblemFeedQueryResult) Set(val *ProblemFeedQueryResult) {
	v.value = val
	v.isSet = true
}

func (v NullableProblemFeedQueryResult) IsSet() bool {
	return v.isSet
}

func (v *NullableProblemFeedQueryResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProblemFeedQueryResult(val *ProblemFeedQueryResult) *NullableProblemFeedQueryResult {
	return &NullableProblemFeedQueryResult{value: val, isSet: true}
}

func (v NullableProblemFeedQueryResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProblemFeedQueryResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


