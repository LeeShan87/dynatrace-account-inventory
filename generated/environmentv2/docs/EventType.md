# EventType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A short description of the event type. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the event type. | [optional] 
**SeverityLevel** | Pointer to **string** | The severity level associated with the event type. | [optional] 
**Type** | Pointer to **string** | The event type. | [optional] 

## Methods

### NewEventType

`func NewEventType() *EventType`

NewEventType instantiates a new EventType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTypeWithDefaults

`func NewEventTypeWithDefaults() *EventType`

NewEventTypeWithDefaults instantiates a new EventType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EventType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *EventType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EventType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EventType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EventType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSeverityLevel

`func (o *EventType) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *EventType) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *EventType) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.

### HasSeverityLevel

`func (o *EventType) HasSeverityLevel() bool`

HasSeverityLevel returns a boolean if a field has been set.

### GetType

`func (o *EventType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EventType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EventType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EventType) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


