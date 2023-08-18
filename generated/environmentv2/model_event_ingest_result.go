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

// checks if the EventIngestResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventIngestResult{}

// EventIngestResult The result of a created event report.
type EventIngestResult struct {
	// The correlation ID of the created event.
	CorrelationId *string `json:"correlationId,omitempty"`
	// The status of the ingestion.
	Status *string `json:"status,omitempty"`
}

// NewEventIngestResult instantiates a new EventIngestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventIngestResult() *EventIngestResult {
	this := EventIngestResult{}
	return &this
}

// NewEventIngestResultWithDefaults instantiates a new EventIngestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventIngestResultWithDefaults() *EventIngestResult {
	this := EventIngestResult{}
	return &this
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *EventIngestResult) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventIngestResult) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *EventIngestResult) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *EventIngestResult) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EventIngestResult) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventIngestResult) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EventIngestResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EventIngestResult) SetStatus(v string) {
	o.Status = &v
}

func (o EventIngestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventIngestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableEventIngestResult struct {
	value *EventIngestResult
	isSet bool
}

func (v NullableEventIngestResult) Get() *EventIngestResult {
	return v.value
}

func (v *NullableEventIngestResult) Set(val *EventIngestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEventIngestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEventIngestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventIngestResult(val *EventIngestResult) *NullableEventIngestResult {
	return &NullableEventIngestResult{value: val, isSet: true}
}

func (v NullableEventIngestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventIngestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


