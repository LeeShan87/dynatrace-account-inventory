# ExtensionEventDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Content of the event | [optional] 
**DtActiveGateId** | Pointer to **string** | Hexadecimal ID of Active Gate that uses this monitoring configuration.  Example: &#x60;0x1a2b3c4d&#x60; | [optional] 
**DtEntityHost** | Pointer to **string** | Host that uses this monitoring configuration.  Example: &#x60;HOST-ABCDEF0123456789&#x60; | [optional] 
**DtExtensionDs** | Pointer to **string** | Data source that uses this monitoring configuration.  Example: &#x60;snmp&#x60; | [optional] 
**Severity** | Pointer to **string** | Severity of the event | [optional] 
**Status** | Pointer to **string** | Status of the event | [optional] 
**Timestamp** | Pointer to **string** | Timestamp of the event | [optional] 

## Methods

### NewExtensionEventDto

`func NewExtensionEventDto() *ExtensionEventDto`

NewExtensionEventDto instantiates a new ExtensionEventDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionEventDtoWithDefaults

`func NewExtensionEventDtoWithDefaults() *ExtensionEventDto`

NewExtensionEventDtoWithDefaults instantiates a new ExtensionEventDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ExtensionEventDto) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ExtensionEventDto) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ExtensionEventDto) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ExtensionEventDto) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetDtActiveGateId

`func (o *ExtensionEventDto) GetDtActiveGateId() string`

GetDtActiveGateId returns the DtActiveGateId field if non-nil, zero value otherwise.

### GetDtActiveGateIdOk

`func (o *ExtensionEventDto) GetDtActiveGateIdOk() (*string, bool)`

GetDtActiveGateIdOk returns a tuple with the DtActiveGateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtActiveGateId

`func (o *ExtensionEventDto) SetDtActiveGateId(v string)`

SetDtActiveGateId sets DtActiveGateId field to given value.

### HasDtActiveGateId

`func (o *ExtensionEventDto) HasDtActiveGateId() bool`

HasDtActiveGateId returns a boolean if a field has been set.

### GetDtEntityHost

`func (o *ExtensionEventDto) GetDtEntityHost() string`

GetDtEntityHost returns the DtEntityHost field if non-nil, zero value otherwise.

### GetDtEntityHostOk

`func (o *ExtensionEventDto) GetDtEntityHostOk() (*string, bool)`

GetDtEntityHostOk returns a tuple with the DtEntityHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtEntityHost

`func (o *ExtensionEventDto) SetDtEntityHost(v string)`

SetDtEntityHost sets DtEntityHost field to given value.

### HasDtEntityHost

`func (o *ExtensionEventDto) HasDtEntityHost() bool`

HasDtEntityHost returns a boolean if a field has been set.

### GetDtExtensionDs

`func (o *ExtensionEventDto) GetDtExtensionDs() string`

GetDtExtensionDs returns the DtExtensionDs field if non-nil, zero value otherwise.

### GetDtExtensionDsOk

`func (o *ExtensionEventDto) GetDtExtensionDsOk() (*string, bool)`

GetDtExtensionDsOk returns a tuple with the DtExtensionDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtExtensionDs

`func (o *ExtensionEventDto) SetDtExtensionDs(v string)`

SetDtExtensionDs sets DtExtensionDs field to given value.

### HasDtExtensionDs

`func (o *ExtensionEventDto) HasDtExtensionDs() bool`

HasDtExtensionDs returns a boolean if a field has been set.

### GetSeverity

`func (o *ExtensionEventDto) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ExtensionEventDto) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ExtensionEventDto) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ExtensionEventDto) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetStatus

`func (o *ExtensionEventDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExtensionEventDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExtensionEventDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExtensionEventDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimestamp

`func (o *ExtensionEventDto) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ExtensionEventDto) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ExtensionEventDto) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ExtensionEventDto) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


