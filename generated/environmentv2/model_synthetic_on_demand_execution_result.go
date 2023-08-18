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

// checks if the SyntheticOnDemandExecutionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SyntheticOnDemandExecutionResult{}

// SyntheticOnDemandExecutionResult The result of on-demand synthetic monitor execution.
type SyntheticOnDemandExecutionResult struct {
	// The batch identifier of the triggered executions.
	BatchId string `json:"batchId"`
	// Monitors for which on-demand executions were triggered.
	Triggered []SyntheticOnDemandTriggeredMonitor `json:"triggered,omitempty"`
	// The total number of the triggered executions within the batch.
	TriggeredCount int32 `json:"triggeredCount"`
	// The total number of problems within the batch.
	TriggeringProblemsCount int32 `json:"triggeringProblemsCount"`
	// List with the entities for which triggering problems occurred.
	TriggeringProblemsDetails []SyntheticOnDemandTriggeringProblemDetails `json:"triggeringProblemsDetails,omitempty"`
}

// NewSyntheticOnDemandExecutionResult instantiates a new SyntheticOnDemandExecutionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticOnDemandExecutionResult(batchId string, triggeredCount int32, triggeringProblemsCount int32) *SyntheticOnDemandExecutionResult {
	this := SyntheticOnDemandExecutionResult{}
	this.BatchId = batchId
	this.TriggeredCount = triggeredCount
	this.TriggeringProblemsCount = triggeringProblemsCount
	return &this
}

// NewSyntheticOnDemandExecutionResultWithDefaults instantiates a new SyntheticOnDemandExecutionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticOnDemandExecutionResultWithDefaults() *SyntheticOnDemandExecutionResult {
	this := SyntheticOnDemandExecutionResult{}
	return &this
}

// GetBatchId returns the BatchId field value
func (o *SyntheticOnDemandExecutionResult) GetBatchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value
// and a boolean to check if the value has been set.
func (o *SyntheticOnDemandExecutionResult) GetBatchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchId, true
}

// SetBatchId sets field value
func (o *SyntheticOnDemandExecutionResult) SetBatchId(v string) {
	o.BatchId = v
}

// GetTriggered returns the Triggered field value if set, zero value otherwise.
func (o *SyntheticOnDemandExecutionResult) GetTriggered() []SyntheticOnDemandTriggeredMonitor {
	if o == nil || IsNil(o.Triggered) {
		var ret []SyntheticOnDemandTriggeredMonitor
		return ret
	}
	return o.Triggered
}

// GetTriggeredOk returns a tuple with the Triggered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticOnDemandExecutionResult) GetTriggeredOk() ([]SyntheticOnDemandTriggeredMonitor, bool) {
	if o == nil || IsNil(o.Triggered) {
		return nil, false
	}
	return o.Triggered, true
}

// HasTriggered returns a boolean if a field has been set.
func (o *SyntheticOnDemandExecutionResult) HasTriggered() bool {
	if o != nil && !IsNil(o.Triggered) {
		return true
	}

	return false
}

// SetTriggered gets a reference to the given []SyntheticOnDemandTriggeredMonitor and assigns it to the Triggered field.
func (o *SyntheticOnDemandExecutionResult) SetTriggered(v []SyntheticOnDemandTriggeredMonitor) {
	o.Triggered = v
}

// GetTriggeredCount returns the TriggeredCount field value
func (o *SyntheticOnDemandExecutionResult) GetTriggeredCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TriggeredCount
}

// GetTriggeredCountOk returns a tuple with the TriggeredCount field value
// and a boolean to check if the value has been set.
func (o *SyntheticOnDemandExecutionResult) GetTriggeredCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggeredCount, true
}

// SetTriggeredCount sets field value
func (o *SyntheticOnDemandExecutionResult) SetTriggeredCount(v int32) {
	o.TriggeredCount = v
}

// GetTriggeringProblemsCount returns the TriggeringProblemsCount field value
func (o *SyntheticOnDemandExecutionResult) GetTriggeringProblemsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TriggeringProblemsCount
}

// GetTriggeringProblemsCountOk returns a tuple with the TriggeringProblemsCount field value
// and a boolean to check if the value has been set.
func (o *SyntheticOnDemandExecutionResult) GetTriggeringProblemsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggeringProblemsCount, true
}

// SetTriggeringProblemsCount sets field value
func (o *SyntheticOnDemandExecutionResult) SetTriggeringProblemsCount(v int32) {
	o.TriggeringProblemsCount = v
}

// GetTriggeringProblemsDetails returns the TriggeringProblemsDetails field value if set, zero value otherwise.
func (o *SyntheticOnDemandExecutionResult) GetTriggeringProblemsDetails() []SyntheticOnDemandTriggeringProblemDetails {
	if o == nil || IsNil(o.TriggeringProblemsDetails) {
		var ret []SyntheticOnDemandTriggeringProblemDetails
		return ret
	}
	return o.TriggeringProblemsDetails
}

// GetTriggeringProblemsDetailsOk returns a tuple with the TriggeringProblemsDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticOnDemandExecutionResult) GetTriggeringProblemsDetailsOk() ([]SyntheticOnDemandTriggeringProblemDetails, bool) {
	if o == nil || IsNil(o.TriggeringProblemsDetails) {
		return nil, false
	}
	return o.TriggeringProblemsDetails, true
}

// HasTriggeringProblemsDetails returns a boolean if a field has been set.
func (o *SyntheticOnDemandExecutionResult) HasTriggeringProblemsDetails() bool {
	if o != nil && !IsNil(o.TriggeringProblemsDetails) {
		return true
	}

	return false
}

// SetTriggeringProblemsDetails gets a reference to the given []SyntheticOnDemandTriggeringProblemDetails and assigns it to the TriggeringProblemsDetails field.
func (o *SyntheticOnDemandExecutionResult) SetTriggeringProblemsDetails(v []SyntheticOnDemandTriggeringProblemDetails) {
	o.TriggeringProblemsDetails = v
}

func (o SyntheticOnDemandExecutionResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SyntheticOnDemandExecutionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["batchId"] = o.BatchId
	if !IsNil(o.Triggered) {
		toSerialize["triggered"] = o.Triggered
	}
	toSerialize["triggeredCount"] = o.TriggeredCount
	toSerialize["triggeringProblemsCount"] = o.TriggeringProblemsCount
	if !IsNil(o.TriggeringProblemsDetails) {
		toSerialize["triggeringProblemsDetails"] = o.TriggeringProblemsDetails
	}
	return toSerialize, nil
}

type NullableSyntheticOnDemandExecutionResult struct {
	value *SyntheticOnDemandExecutionResult
	isSet bool
}

func (v NullableSyntheticOnDemandExecutionResult) Get() *SyntheticOnDemandExecutionResult {
	return v.value
}

func (v *NullableSyntheticOnDemandExecutionResult) Set(val *SyntheticOnDemandExecutionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticOnDemandExecutionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticOnDemandExecutionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticOnDemandExecutionResult(val *SyntheticOnDemandExecutionResult) *NullableSyntheticOnDemandExecutionResult {
	return &NullableSyntheticOnDemandExecutionResult{value: val, isSet: true}
}

func (v NullableSyntheticOnDemandExecutionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticOnDemandExecutionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


