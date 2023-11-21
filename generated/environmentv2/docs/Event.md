# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | The correlation ID of the event.  | [optional] 
**EndTime** | Pointer to **int64** | The timestamp when the event was closed, in UTC milliseconds.    Has the value of &#x60;null&#x60; if the event is still active. | [optional] 
**EntityId** | Pointer to [**EntityStub**](EntityStub.md) |  | [optional] 
**EntityTags** | Pointer to [**[]METag**](METag.md) | A list of tags of the related entity. | [optional] 
**EventId** | Pointer to **string** | The ID of the event. | [optional] 
**EventType** | Pointer to **string** | The type of the event. | [optional] 
**FrequentEvent** | Pointer to **bool** | If &#x60;true&#x60;, the event happens [frequently](https://dt-url.net/4da3kdg).    A frequent event doesn&#39;t raise a problem. | [optional] 
**ManagementZones** | Pointer to [**[]ManagementZone**](ManagementZone.md) | A list of all management zones that the event belongs to. | [optional] 
**Properties** | Pointer to [**[]EventProperty**](EventProperty.md) | A list of event properties. | [optional] 
**StartTime** | Pointer to **int64** | The timestamp when the event was raised, in UTC milliseconds. | [optional] 
**Status** | Pointer to **string** | The status of the event. | [optional] 
**SuppressAlert** | Pointer to **bool** | The alerting status during a [maintenance](https://dt-url.net/b2123rg0):    * &#x60;false&#x60;: Alerting works as usual.  * &#x60;true&#x60;: Alerting is disabled. | [optional] 
**SuppressProblem** | Pointer to **bool** | The problem detection status during a [maintenance](https://dt-url.net/b2123rg0):    * &#x60;false&#x60;: Problem detection works as usual.  * &#x60;true&#x60;: Problem detection is disabled. | [optional] 
**Title** | Pointer to **string** | The title of the event. | [optional] 
**UnderMaintenance** | Pointer to **bool** | If &#x60;true&#x60;, the event happened while the monitored system was under maintenance. | [optional] 

## Methods

### NewEvent

`func NewEvent() *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *Event) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *Event) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *Event) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *Event) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetEndTime

`func (o *Event) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Event) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Event) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Event) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntityId

`func (o *Event) GetEntityId() EntityStub`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *Event) GetEntityIdOk() (*EntityStub, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *Event) SetEntityId(v EntityStub)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *Event) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetEntityTags

`func (o *Event) GetEntityTags() []METag`

GetEntityTags returns the EntityTags field if non-nil, zero value otherwise.

### GetEntityTagsOk

`func (o *Event) GetEntityTagsOk() (*[]METag, bool)`

GetEntityTagsOk returns a tuple with the EntityTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTags

`func (o *Event) SetEntityTags(v []METag)`

SetEntityTags sets EntityTags field to given value.

### HasEntityTags

`func (o *Event) HasEntityTags() bool`

HasEntityTags returns a boolean if a field has been set.

### GetEventId

`func (o *Event) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *Event) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *Event) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *Event) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetEventType

`func (o *Event) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *Event) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *Event) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *Event) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetFrequentEvent

`func (o *Event) GetFrequentEvent() bool`

GetFrequentEvent returns the FrequentEvent field if non-nil, zero value otherwise.

### GetFrequentEventOk

`func (o *Event) GetFrequentEventOk() (*bool, bool)`

GetFrequentEventOk returns a tuple with the FrequentEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequentEvent

`func (o *Event) SetFrequentEvent(v bool)`

SetFrequentEvent sets FrequentEvent field to given value.

### HasFrequentEvent

`func (o *Event) HasFrequentEvent() bool`

HasFrequentEvent returns a boolean if a field has been set.

### GetManagementZones

`func (o *Event) GetManagementZones() []ManagementZone`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *Event) GetManagementZonesOk() (*[]ManagementZone, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *Event) SetManagementZones(v []ManagementZone)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *Event) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetProperties

`func (o *Event) GetProperties() []EventProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Event) GetPropertiesOk() (*[]EventProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Event) SetProperties(v []EventProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Event) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStartTime

`func (o *Event) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Event) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Event) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Event) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *Event) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Event) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Event) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Event) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuppressAlert

`func (o *Event) GetSuppressAlert() bool`

GetSuppressAlert returns the SuppressAlert field if non-nil, zero value otherwise.

### GetSuppressAlertOk

`func (o *Event) GetSuppressAlertOk() (*bool, bool)`

GetSuppressAlertOk returns a tuple with the SuppressAlert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressAlert

`func (o *Event) SetSuppressAlert(v bool)`

SetSuppressAlert sets SuppressAlert field to given value.

### HasSuppressAlert

`func (o *Event) HasSuppressAlert() bool`

HasSuppressAlert returns a boolean if a field has been set.

### GetSuppressProblem

`func (o *Event) GetSuppressProblem() bool`

GetSuppressProblem returns the SuppressProblem field if non-nil, zero value otherwise.

### GetSuppressProblemOk

`func (o *Event) GetSuppressProblemOk() (*bool, bool)`

GetSuppressProblemOk returns a tuple with the SuppressProblem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuppressProblem

`func (o *Event) SetSuppressProblem(v bool)`

SetSuppressProblem sets SuppressProblem field to given value.

### HasSuppressProblem

`func (o *Event) HasSuppressProblem() bool`

HasSuppressProblem returns a boolean if a field has been set.

### GetTitle

`func (o *Event) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Event) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Event) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Event) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUnderMaintenance

`func (o *Event) GetUnderMaintenance() bool`

GetUnderMaintenance returns the UnderMaintenance field if non-nil, zero value otherwise.

### GetUnderMaintenanceOk

`func (o *Event) GetUnderMaintenanceOk() (*bool, bool)`

GetUnderMaintenanceOk returns a tuple with the UnderMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderMaintenance

`func (o *Event) SetUnderMaintenance(v bool)`

SetUnderMaintenance sets UnderMaintenance field to given value.

### HasUnderMaintenance

`func (o *Event) HasUnderMaintenance() bool`

HasUnderMaintenance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


