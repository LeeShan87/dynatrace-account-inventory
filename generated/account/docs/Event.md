# Event

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentUuid** | **string** | The UUID of the environment that threw the event. | 
**Capability** | **string** | The subscription capability name. | 
**Date** | **time.Time** | The date the event was fired in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**CreatedAt** | **time.Time** | The date the notification was created in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**Severity** | **string** | The severity of the even. | 
**Message** | **string** | The message from the event. | 
**EventType** | **string** | The type of event: Forecast or usage. | 
**NotificationLevel** | **string** | The notification level of the event. | 

## Methods

### NewEvent

`func NewEvent(environmentUuid string, capability string, date time.Time, createdAt time.Time, severity string, message string, eventType string, notificationLevel string, ) *Event`

NewEvent instantiates a new Event object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventWithDefaults

`func NewEventWithDefaults() *Event`

NewEventWithDefaults instantiates a new Event object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentUuid

`func (o *Event) GetEnvironmentUuid() string`

GetEnvironmentUuid returns the EnvironmentUuid field if non-nil, zero value otherwise.

### GetEnvironmentUuidOk

`func (o *Event) GetEnvironmentUuidOk() (*string, bool)`

GetEnvironmentUuidOk returns a tuple with the EnvironmentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentUuid

`func (o *Event) SetEnvironmentUuid(v string)`

SetEnvironmentUuid sets EnvironmentUuid field to given value.


### GetCapability

`func (o *Event) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *Event) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *Event) SetCapability(v string)`

SetCapability sets Capability field to given value.


### GetDate

`func (o *Event) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Event) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Event) SetDate(v time.Time)`

SetDate sets Date field to given value.


### GetCreatedAt

`func (o *Event) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Event) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Event) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetSeverity

`func (o *Event) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Event) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Event) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetMessage

`func (o *Event) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Event) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Event) SetMessage(v string)`

SetMessage sets Message field to given value.


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


### GetNotificationLevel

`func (o *Event) GetNotificationLevel() string`

GetNotificationLevel returns the NotificationLevel field if non-nil, zero value otherwise.

### GetNotificationLevelOk

`func (o *Event) GetNotificationLevelOk() (*string, bool)`

GetNotificationLevelOk returns a tuple with the NotificationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationLevel

`func (o *Event) SetNotificationLevel(v string)`

SetNotificationLevel sets NotificationLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


