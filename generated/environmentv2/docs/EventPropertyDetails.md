# EventPropertyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A short description of the event property. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the event property. | [optional] 
**Key** | Pointer to **string** | The key of the event property. | [optional] 
**Writable** | Pointer to **bool** | The property can (&#x60;true&#x60;) or cannot (&#x60;false&#x60;) be set during event ingestion. | [optional] 

## Methods

### NewEventPropertyDetails

`func NewEventPropertyDetails() *EventPropertyDetails`

NewEventPropertyDetails instantiates a new EventPropertyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPropertyDetailsWithDefaults

`func NewEventPropertyDetailsWithDefaults() *EventPropertyDetails`

NewEventPropertyDetailsWithDefaults instantiates a new EventPropertyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EventPropertyDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EventPropertyDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EventPropertyDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EventPropertyDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *EventPropertyDetails) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EventPropertyDetails) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EventPropertyDetails) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EventPropertyDetails) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetKey

`func (o *EventPropertyDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EventPropertyDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EventPropertyDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EventPropertyDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetWritable

`func (o *EventPropertyDetails) GetWritable() bool`

GetWritable returns the Writable field if non-nil, zero value otherwise.

### GetWritableOk

`func (o *EventPropertyDetails) GetWritableOk() (*bool, bool)`

GetWritableOk returns a tuple with the Writable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWritable

`func (o *EventPropertyDetails) SetWritable(v bool)`

SetWritable sets Writable field to given value.

### HasWritable

`func (o *EventPropertyDetails) HasWritable() bool`

HasWritable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


