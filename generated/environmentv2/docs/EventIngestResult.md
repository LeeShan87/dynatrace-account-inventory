# EventIngestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | The correlation ID of the created event. | [optional] 
**Status** | Pointer to **string** | The status of the ingestion. | [optional] 

## Methods

### NewEventIngestResult

`func NewEventIngestResult() *EventIngestResult`

NewEventIngestResult instantiates a new EventIngestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventIngestResultWithDefaults

`func NewEventIngestResultWithDefaults() *EventIngestResult`

NewEventIngestResultWithDefaults instantiates a new EventIngestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *EventIngestResult) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *EventIngestResult) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *EventIngestResult) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *EventIngestResult) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetStatus

`func (o *EventIngestResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EventIngestResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EventIngestResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EventIngestResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


