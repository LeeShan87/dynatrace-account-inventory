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

// checks if the EventRestEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventRestEntry{}

// EventRestEntry Set of parameters of the event.    Apart from the general properties mentioned here, which each event has, an actual event has a set of metadata that varies depending on the event type.
type EventRestEntry struct {
	// The timestamp of the event closure, in UTC milliseconds
	EndTime *int64 `json:"endTime,omitempty"`
	// The ID of the affected Dynatrace entity.
	EntityId *string `json:"entityId,omitempty"`
	// The name of the affected Dynatrace entity.
	EntityName *string `json:"entityName,omitempty"`
	// The state of the event: open or closed.
	EventStatus *string `json:"eventStatus,omitempty"`
	// The type of the event.
	EventType *string `json:"eventType,omitempty"`
	// The encoded ID of the event. The format is *eventID_startTime*.    You should use the value from this field when you need an event ID. 
	Id *string `json:"id,omitempty"`
	// The impact level of the event. It shows what is affected by the problem: infrastructure, service, or application.
	ImpactLevel *string `json:"impactLevel,omitempty"`
	// The id of the resource the event occurred on.
	ResourceId *string `json:"resourceId,omitempty"`
	// The name of the resource the event occurred on.
	ResourceName *string `json:"resourceName,omitempty"`
	// The severity of the event.
	SeverityLevel *string `json:"severityLevel,omitempty"`
	// The timestamp of the event detection, in UTC milliseconds.
	StartTime *int64 `json:"startTime,omitempty"`
	// Tags of the Dynatrace entity that raised the event.
	Tags []TagInfo `json:"tags,omitempty"`
}

// NewEventRestEntry instantiates a new EventRestEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventRestEntry() *EventRestEntry {
	this := EventRestEntry{}
	return &this
}

// NewEventRestEntryWithDefaults instantiates a new EventRestEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventRestEntryWithDefaults() *EventRestEntry {
	this := EventRestEntry{}
	return &this
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *EventRestEntry) GetEndTime() int64 {
	if o == nil || IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetEndTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *EventRestEntry) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *EventRestEntry) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *EventRestEntry) GetEntityId() string {
	if o == nil || IsNil(o.EntityId) {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetEntityIdOk() (*string, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *EventRestEntry) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *EventRestEntry) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityName returns the EntityName field value if set, zero value otherwise.
func (o *EventRestEntry) GetEntityName() string {
	if o == nil || IsNil(o.EntityName) {
		var ret string
		return ret
	}
	return *o.EntityName
}

// GetEntityNameOk returns a tuple with the EntityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetEntityNameOk() (*string, bool) {
	if o == nil || IsNil(o.EntityName) {
		return nil, false
	}
	return o.EntityName, true
}

// HasEntityName returns a boolean if a field has been set.
func (o *EventRestEntry) HasEntityName() bool {
	if o != nil && !IsNil(o.EntityName) {
		return true
	}

	return false
}

// SetEntityName gets a reference to the given string and assigns it to the EntityName field.
func (o *EventRestEntry) SetEntityName(v string) {
	o.EntityName = &v
}

// GetEventStatus returns the EventStatus field value if set, zero value otherwise.
func (o *EventRestEntry) GetEventStatus() string {
	if o == nil || IsNil(o.EventStatus) {
		var ret string
		return ret
	}
	return *o.EventStatus
}

// GetEventStatusOk returns a tuple with the EventStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetEventStatusOk() (*string, bool) {
	if o == nil || IsNil(o.EventStatus) {
		return nil, false
	}
	return o.EventStatus, true
}

// HasEventStatus returns a boolean if a field has been set.
func (o *EventRestEntry) HasEventStatus() bool {
	if o != nil && !IsNil(o.EventStatus) {
		return true
	}

	return false
}

// SetEventStatus gets a reference to the given string and assigns it to the EventStatus field.
func (o *EventRestEntry) SetEventStatus(v string) {
	o.EventStatus = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *EventRestEntry) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *EventRestEntry) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *EventRestEntry) SetEventType(v string) {
	o.EventType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EventRestEntry) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EventRestEntry) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EventRestEntry) SetId(v string) {
	o.Id = &v
}

// GetImpactLevel returns the ImpactLevel field value if set, zero value otherwise.
func (o *EventRestEntry) GetImpactLevel() string {
	if o == nil || IsNil(o.ImpactLevel) {
		var ret string
		return ret
	}
	return *o.ImpactLevel
}

// GetImpactLevelOk returns a tuple with the ImpactLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetImpactLevelOk() (*string, bool) {
	if o == nil || IsNil(o.ImpactLevel) {
		return nil, false
	}
	return o.ImpactLevel, true
}

// HasImpactLevel returns a boolean if a field has been set.
func (o *EventRestEntry) HasImpactLevel() bool {
	if o != nil && !IsNil(o.ImpactLevel) {
		return true
	}

	return false
}

// SetImpactLevel gets a reference to the given string and assigns it to the ImpactLevel field.
func (o *EventRestEntry) SetImpactLevel(v string) {
	o.ImpactLevel = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *EventRestEntry) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *EventRestEntry) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *EventRestEntry) SetResourceId(v string) {
	o.ResourceId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *EventRestEntry) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *EventRestEntry) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *EventRestEntry) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetSeverityLevel returns the SeverityLevel field value if set, zero value otherwise.
func (o *EventRestEntry) GetSeverityLevel() string {
	if o == nil || IsNil(o.SeverityLevel) {
		var ret string
		return ret
	}
	return *o.SeverityLevel
}

// GetSeverityLevelOk returns a tuple with the SeverityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetSeverityLevelOk() (*string, bool) {
	if o == nil || IsNil(o.SeverityLevel) {
		return nil, false
	}
	return o.SeverityLevel, true
}

// HasSeverityLevel returns a boolean if a field has been set.
func (o *EventRestEntry) HasSeverityLevel() bool {
	if o != nil && !IsNil(o.SeverityLevel) {
		return true
	}

	return false
}

// SetSeverityLevel gets a reference to the given string and assigns it to the SeverityLevel field.
func (o *EventRestEntry) SetSeverityLevel(v string) {
	o.SeverityLevel = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *EventRestEntry) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *EventRestEntry) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *EventRestEntry) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *EventRestEntry) GetTags() []TagInfo {
	if o == nil || IsNil(o.Tags) {
		var ret []TagInfo
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventRestEntry) GetTagsOk() ([]TagInfo, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *EventRestEntry) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []TagInfo and assigns it to the Tags field.
func (o *EventRestEntry) SetTags(v []TagInfo) {
	o.Tags = v
}

func (o EventRestEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventRestEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.EntityName) {
		toSerialize["entityName"] = o.EntityName
	}
	if !IsNil(o.EventStatus) {
		toSerialize["eventStatus"] = o.EventStatus
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ImpactLevel) {
		toSerialize["impactLevel"] = o.ImpactLevel
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}
	if !IsNil(o.SeverityLevel) {
		toSerialize["severityLevel"] = o.SeverityLevel
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableEventRestEntry struct {
	value *EventRestEntry
	isSet bool
}

func (v NullableEventRestEntry) Get() *EventRestEntry {
	return v.value
}

func (v *NullableEventRestEntry) Set(val *EventRestEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableEventRestEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableEventRestEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventRestEntry(val *EventRestEntry) *NullableEventRestEntry {
	return &NullableEventRestEntry{value: val, isSet: true}
}

func (v NullableEventRestEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventRestEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


