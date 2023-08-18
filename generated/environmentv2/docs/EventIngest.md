# EventIngest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | Pointer to **int64** | The end time of the event, in UTC milliseconds.  If not set, the start time plus timeout is used. | [optional] 
**EntitySelector** | Pointer to **string** | The [entity selector](https://dt-url.net/apientityselector), defining a set of Dynatrace entities to be associated with the event.  Only entities that have been active within the last 24 hours can be selected.    If not set, the event is associated with the environment (&#x60;dt.entity.environment&#x60;) entity. | [optional] 
**EventType** | **string** | The type of the event. | 
**Properties** | Pointer to **map[string]string** | A map of event properties.  Keys with prefix &#x60;dt.*&#x60; are generally disallowed, with the exceptions of &#x60;dt.event.*&#x60;, &#x60;dt.davis.*&#x60; and &#x60;dt.entity.*&#x60;. These reserved keys may be used to set event properties with predefined semantics within the Dynatrace product. &#x60;dt.entity.*&#x60; keys may be used to provide additional information on an event, but will not lead to the event being tied to the specified entities. All other keys are interpreted as user-defined event properties.  Values of Dynatrace-reserved properties must fulfill the requirements of the respective property. | [optional] 
**StartTime** | Pointer to **int64** | The start time of the event, in UTC milliseconds.  If not set, the current timestamp is used.  Depending on the event type, the start time must not lie in the future more than 5 minutes for trigger events and 7 days for info events. | [optional] 
**Timeout** | Pointer to **int32** | The timeout of the event, in minutes.  If not set, 15 is used.  The timeout will automatically be capped to a maximum of 360 minutes (6 hours).  Problem-opening events can be refreshed and therefore kept open by sending the same payload again. | [optional] 
**Title** | **string** | The title of the event. | 

## Methods

### NewEventIngest

`func NewEventIngest(eventType string, title string, ) *EventIngest`

NewEventIngest instantiates a new EventIngest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventIngestWithDefaults

`func NewEventIngestWithDefaults() *EventIngest`

NewEventIngestWithDefaults instantiates a new EventIngest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *EventIngest) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *EventIngest) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *EventIngest) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *EventIngest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetEntitySelector

`func (o *EventIngest) GetEntitySelector() string`

GetEntitySelector returns the EntitySelector field if non-nil, zero value otherwise.

### GetEntitySelectorOk

`func (o *EventIngest) GetEntitySelectorOk() (*string, bool)`

GetEntitySelectorOk returns a tuple with the EntitySelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelector

`func (o *EventIngest) SetEntitySelector(v string)`

SetEntitySelector sets EntitySelector field to given value.

### HasEntitySelector

`func (o *EventIngest) HasEntitySelector() bool`

HasEntitySelector returns a boolean if a field has been set.

### GetEventType

`func (o *EventIngest) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventIngest) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventIngest) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetProperties

`func (o *EventIngest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *EventIngest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *EventIngest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *EventIngest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStartTime

`func (o *EventIngest) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *EventIngest) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *EventIngest) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *EventIngest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTimeout

`func (o *EventIngest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *EventIngest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *EventIngest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *EventIngest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTitle

`func (o *EventIngest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EventIngest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EventIngest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


