# EventIngestResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIngestResults** | Pointer to [**[]EventIngestResult**](EventIngestResult.md) | The result of each created event report. | [optional] 
**ReportCount** | Pointer to **int32** | The number of created event reports. | [optional] 

## Methods

### NewEventIngestResults

`func NewEventIngestResults() *EventIngestResults`

NewEventIngestResults instantiates a new EventIngestResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventIngestResultsWithDefaults

`func NewEventIngestResultsWithDefaults() *EventIngestResults`

NewEventIngestResultsWithDefaults instantiates a new EventIngestResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIngestResults

`func (o *EventIngestResults) GetEventIngestResults() []EventIngestResult`

GetEventIngestResults returns the EventIngestResults field if non-nil, zero value otherwise.

### GetEventIngestResultsOk

`func (o *EventIngestResults) GetEventIngestResultsOk() (*[]EventIngestResult, bool)`

GetEventIngestResultsOk returns a tuple with the EventIngestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIngestResults

`func (o *EventIngestResults) SetEventIngestResults(v []EventIngestResult)`

SetEventIngestResults sets EventIngestResults field to given value.

### HasEventIngestResults

`func (o *EventIngestResults) HasEventIngestResults() bool`

HasEventIngestResults returns a boolean if a field has been set.

### GetReportCount

`func (o *EventIngestResults) GetReportCount() int32`

GetReportCount returns the ReportCount field if non-nil, zero value otherwise.

### GetReportCountOk

`func (o *EventIngestResults) GetReportCountOk() (*int32, bool)`

GetReportCountOk returns a tuple with the ReportCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportCount

`func (o *EventIngestResults) SetReportCount(v int32)`

SetReportCount sets ReportCount field to given value.

### HasReportCount

`func (o *EventIngestResults) HasReportCount() bool`

HasReportCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


