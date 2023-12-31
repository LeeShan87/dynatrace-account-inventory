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

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Event{}

// Event Configuration of an event.
type Event struct {
	// The correlation ID of the event. 
	CorrelationId *string `json:"correlationId,omitempty"`
	// The timestamp when the event was closed, in UTC milliseconds.    Has the value of `null` if the event is still active.
	EndTime *int64 `json:"endTime,omitempty"`
	EntityId *EntityStub `json:"entityId,omitempty"`
	// A list of tags of the related entity.
	EntityTags []METag `json:"entityTags,omitempty"`
	// The ID of the event.
	EventId *string `json:"eventId,omitempty"`
	// The type of the event.
	EventType *string `json:"eventType,omitempty"`
	// If `true`, the event happens [frequently](https://dt-url.net/4da3kdg).    A frequent event doesn't raise a problem.
	FrequentEvent *bool `json:"frequentEvent,omitempty"`
	// A list of all management zones that the event belongs to.
	ManagementZones []ManagementZone `json:"managementZones,omitempty"`
	// A list of event properties.
	Properties []EventProperty `json:"properties,omitempty"`
	// The timestamp when the event was raised, in UTC milliseconds.
	StartTime *int64 `json:"startTime,omitempty"`
	// The status of the event.
	Status *string `json:"status,omitempty"`
	// The alerting status during a [maintenance](https://dt-url.net/b2123rg0):    * `false`: Alerting works as usual.  * `true`: Alerting is disabled.
	SuppressAlert *bool `json:"suppressAlert,omitempty"`
	// The problem detection status during a [maintenance](https://dt-url.net/b2123rg0):    * `false`: Problem detection works as usual.  * `true`: Problem detection is disabled.
	SuppressProblem *bool `json:"suppressProblem,omitempty"`
	// The title of the event.
	Title *string `json:"title,omitempty"`
	// If `true`, the event happened while the monitored system was under maintenance.
	UnderMaintenance *bool `json:"underMaintenance,omitempty"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent() *Event {
	this := Event{}
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *Event) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *Event) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *Event) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *Event) GetEndTime() int64 {
	if o == nil || IsNil(o.EndTime) {
		var ret int64
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEndTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *Event) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int64 and assigns it to the EndTime field.
func (o *Event) SetEndTime(v int64) {
	o.EndTime = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *Event) GetEntityId() EntityStub {
	if o == nil || IsNil(o.EntityId) {
		var ret EntityStub
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEntityIdOk() (*EntityStub, bool) {
	if o == nil || IsNil(o.EntityId) {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *Event) HasEntityId() bool {
	if o != nil && !IsNil(o.EntityId) {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given EntityStub and assigns it to the EntityId field.
func (o *Event) SetEntityId(v EntityStub) {
	o.EntityId = &v
}

// GetEntityTags returns the EntityTags field value if set, zero value otherwise.
func (o *Event) GetEntityTags() []METag {
	if o == nil || IsNil(o.EntityTags) {
		var ret []METag
		return ret
	}
	return o.EntityTags
}

// GetEntityTagsOk returns a tuple with the EntityTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEntityTagsOk() ([]METag, bool) {
	if o == nil || IsNil(o.EntityTags) {
		return nil, false
	}
	return o.EntityTags, true
}

// HasEntityTags returns a boolean if a field has been set.
func (o *Event) HasEntityTags() bool {
	if o != nil && !IsNil(o.EntityTags) {
		return true
	}

	return false
}

// SetEntityTags gets a reference to the given []METag and assigns it to the EntityTags field.
func (o *Event) SetEntityTags(v []METag) {
	o.EntityTags = v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *Event) GetEventId() string {
	if o == nil || IsNil(o.EventId) {
		var ret string
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEventIdOk() (*string, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *Event) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given string and assigns it to the EventId field.
func (o *Event) SetEventId(v string) {
	o.EventId = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *Event) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *Event) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *Event) SetEventType(v string) {
	o.EventType = &v
}

// GetFrequentEvent returns the FrequentEvent field value if set, zero value otherwise.
func (o *Event) GetFrequentEvent() bool {
	if o == nil || IsNil(o.FrequentEvent) {
		var ret bool
		return ret
	}
	return *o.FrequentEvent
}

// GetFrequentEventOk returns a tuple with the FrequentEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetFrequentEventOk() (*bool, bool) {
	if o == nil || IsNil(o.FrequentEvent) {
		return nil, false
	}
	return o.FrequentEvent, true
}

// HasFrequentEvent returns a boolean if a field has been set.
func (o *Event) HasFrequentEvent() bool {
	if o != nil && !IsNil(o.FrequentEvent) {
		return true
	}

	return false
}

// SetFrequentEvent gets a reference to the given bool and assigns it to the FrequentEvent field.
func (o *Event) SetFrequentEvent(v bool) {
	o.FrequentEvent = &v
}

// GetManagementZones returns the ManagementZones field value if set, zero value otherwise.
func (o *Event) GetManagementZones() []ManagementZone {
	if o == nil || IsNil(o.ManagementZones) {
		var ret []ManagementZone
		return ret
	}
	return o.ManagementZones
}

// GetManagementZonesOk returns a tuple with the ManagementZones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetManagementZonesOk() ([]ManagementZone, bool) {
	if o == nil || IsNil(o.ManagementZones) {
		return nil, false
	}
	return o.ManagementZones, true
}

// HasManagementZones returns a boolean if a field has been set.
func (o *Event) HasManagementZones() bool {
	if o != nil && !IsNil(o.ManagementZones) {
		return true
	}

	return false
}

// SetManagementZones gets a reference to the given []ManagementZone and assigns it to the ManagementZones field.
func (o *Event) SetManagementZones(v []ManagementZone) {
	o.ManagementZones = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *Event) GetProperties() []EventProperty {
	if o == nil || IsNil(o.Properties) {
		var ret []EventProperty
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetPropertiesOk() ([]EventProperty, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *Event) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []EventProperty and assigns it to the Properties field.
func (o *Event) SetProperties(v []EventProperty) {
	o.Properties = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Event) GetStartTime() int64 {
	if o == nil || IsNil(o.StartTime) {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetStartTimeOk() (*int64, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Event) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *Event) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Event) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Event) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Event) SetStatus(v string) {
	o.Status = &v
}

// GetSuppressAlert returns the SuppressAlert field value if set, zero value otherwise.
func (o *Event) GetSuppressAlert() bool {
	if o == nil || IsNil(o.SuppressAlert) {
		var ret bool
		return ret
	}
	return *o.SuppressAlert
}

// GetSuppressAlertOk returns a tuple with the SuppressAlert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSuppressAlertOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressAlert) {
		return nil, false
	}
	return o.SuppressAlert, true
}

// HasSuppressAlert returns a boolean if a field has been set.
func (o *Event) HasSuppressAlert() bool {
	if o != nil && !IsNil(o.SuppressAlert) {
		return true
	}

	return false
}

// SetSuppressAlert gets a reference to the given bool and assigns it to the SuppressAlert field.
func (o *Event) SetSuppressAlert(v bool) {
	o.SuppressAlert = &v
}

// GetSuppressProblem returns the SuppressProblem field value if set, zero value otherwise.
func (o *Event) GetSuppressProblem() bool {
	if o == nil || IsNil(o.SuppressProblem) {
		var ret bool
		return ret
	}
	return *o.SuppressProblem
}

// GetSuppressProblemOk returns a tuple with the SuppressProblem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSuppressProblemOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressProblem) {
		return nil, false
	}
	return o.SuppressProblem, true
}

// HasSuppressProblem returns a boolean if a field has been set.
func (o *Event) HasSuppressProblem() bool {
	if o != nil && !IsNil(o.SuppressProblem) {
		return true
	}

	return false
}

// SetSuppressProblem gets a reference to the given bool and assigns it to the SuppressProblem field.
func (o *Event) SetSuppressProblem(v bool) {
	o.SuppressProblem = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Event) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Event) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Event) SetTitle(v string) {
	o.Title = &v
}

// GetUnderMaintenance returns the UnderMaintenance field value if set, zero value otherwise.
func (o *Event) GetUnderMaintenance() bool {
	if o == nil || IsNil(o.UnderMaintenance) {
		var ret bool
		return ret
	}
	return *o.UnderMaintenance
}

// GetUnderMaintenanceOk returns a tuple with the UnderMaintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetUnderMaintenanceOk() (*bool, bool) {
	if o == nil || IsNil(o.UnderMaintenance) {
		return nil, false
	}
	return o.UnderMaintenance, true
}

// HasUnderMaintenance returns a boolean if a field has been set.
func (o *Event) HasUnderMaintenance() bool {
	if o != nil && !IsNil(o.UnderMaintenance) {
		return true
	}

	return false
}

// SetUnderMaintenance gets a reference to the given bool and assigns it to the UnderMaintenance field.
func (o *Event) SetUnderMaintenance(v bool) {
	o.UnderMaintenance = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlationId"] = o.CorrelationId
	}
	if !IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	if !IsNil(o.EntityId) {
		toSerialize["entityId"] = o.EntityId
	}
	if !IsNil(o.EntityTags) {
		toSerialize["entityTags"] = o.EntityTags
	}
	if !IsNil(o.EventId) {
		toSerialize["eventId"] = o.EventId
	}
	if !IsNil(o.EventType) {
		toSerialize["eventType"] = o.EventType
	}
	if !IsNil(o.FrequentEvent) {
		toSerialize["frequentEvent"] = o.FrequentEvent
	}
	if !IsNil(o.ManagementZones) {
		toSerialize["managementZones"] = o.ManagementZones
	}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.SuppressAlert) {
		toSerialize["suppressAlert"] = o.SuppressAlert
	}
	if !IsNil(o.SuppressProblem) {
		toSerialize["suppressProblem"] = o.SuppressProblem
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UnderMaintenance) {
		toSerialize["underMaintenance"] = o.UnderMaintenance
	}
	return toSerialize, nil
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


