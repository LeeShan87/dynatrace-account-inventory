# EventEvidence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Event**](Event.md) |  | [optional] 
**EndTime** | **int64** | The end timestamp of the event, in UTC milliseconds.    Has &#x60;-1&#x60; value, if the event is still active. | 
**EventId** | **string** | The ID of the event. | 
**EventType** | **string** | The type of the event. | 

## Methods

### NewEventEvidence

`func NewEventEvidence(endTime int64, eventId string, eventType string, ) *EventEvidence`

NewEventEvidence instantiates a new EventEvidence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventEvidenceWithDefaults

`func NewEventEvidenceWithDefaults() *EventEvidence`

NewEventEvidenceWithDefaults instantiates a new EventEvidence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *EventEvidence) GetData() Event`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EventEvidence) GetDataOk() (*Event, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EventEvidence) SetData(v Event)`

SetData sets Data field to given value.

### HasData

`func (o *EventEvidence) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEndTime

`func (o *EventEvidence) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EventEvidence) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EventEvidence) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetEventId

`func (o *EventEvidence) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *EventEvidence) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *EventEvidence) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventType

`func (o *EventEvidence) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventEvidence) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventEvidence) SetEventType(v string)`

SetEventType sets EventType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


