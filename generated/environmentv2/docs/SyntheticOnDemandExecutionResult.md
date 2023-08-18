# SyntheticOnDemandExecutionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | **string** | The batch identifier of the triggered executions. | 
**Triggered** | Pointer to [**[]SyntheticOnDemandTriggeredMonitor**](SyntheticOnDemandTriggeredMonitor.md) | Monitors for which on-demand executions were triggered. | [optional] 
**TriggeredCount** | **int32** | The total number of the triggered executions within the batch. | 
**TriggeringProblemsCount** | **int32** | The total number of problems within the batch. | 
**TriggeringProblemsDetails** | Pointer to [**[]SyntheticOnDemandTriggeringProblemDetails**](SyntheticOnDemandTriggeringProblemDetails.md) | List with the entities for which triggering problems occurred. | [optional] 

## Methods

### NewSyntheticOnDemandExecutionResult

`func NewSyntheticOnDemandExecutionResult(batchId string, triggeredCount int32, triggeringProblemsCount int32, ) *SyntheticOnDemandExecutionResult`

NewSyntheticOnDemandExecutionResult instantiates a new SyntheticOnDemandExecutionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticOnDemandExecutionResultWithDefaults

`func NewSyntheticOnDemandExecutionResultWithDefaults() *SyntheticOnDemandExecutionResult`

NewSyntheticOnDemandExecutionResultWithDefaults instantiates a new SyntheticOnDemandExecutionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *SyntheticOnDemandExecutionResult) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *SyntheticOnDemandExecutionResult) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *SyntheticOnDemandExecutionResult) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.


### GetTriggered

`func (o *SyntheticOnDemandExecutionResult) GetTriggered() []SyntheticOnDemandTriggeredMonitor`

GetTriggered returns the Triggered field if non-nil, zero value otherwise.

### GetTriggeredOk

`func (o *SyntheticOnDemandExecutionResult) GetTriggeredOk() (*[]SyntheticOnDemandTriggeredMonitor, bool)`

GetTriggeredOk returns a tuple with the Triggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggered

`func (o *SyntheticOnDemandExecutionResult) SetTriggered(v []SyntheticOnDemandTriggeredMonitor)`

SetTriggered sets Triggered field to given value.

### HasTriggered

`func (o *SyntheticOnDemandExecutionResult) HasTriggered() bool`

HasTriggered returns a boolean if a field has been set.

### GetTriggeredCount

`func (o *SyntheticOnDemandExecutionResult) GetTriggeredCount() int32`

GetTriggeredCount returns the TriggeredCount field if non-nil, zero value otherwise.

### GetTriggeredCountOk

`func (o *SyntheticOnDemandExecutionResult) GetTriggeredCountOk() (*int32, bool)`

GetTriggeredCountOk returns a tuple with the TriggeredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredCount

`func (o *SyntheticOnDemandExecutionResult) SetTriggeredCount(v int32)`

SetTriggeredCount sets TriggeredCount field to given value.


### GetTriggeringProblemsCount

`func (o *SyntheticOnDemandExecutionResult) GetTriggeringProblemsCount() int32`

GetTriggeringProblemsCount returns the TriggeringProblemsCount field if non-nil, zero value otherwise.

### GetTriggeringProblemsCountOk

`func (o *SyntheticOnDemandExecutionResult) GetTriggeringProblemsCountOk() (*int32, bool)`

GetTriggeringProblemsCountOk returns a tuple with the TriggeringProblemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeringProblemsCount

`func (o *SyntheticOnDemandExecutionResult) SetTriggeringProblemsCount(v int32)`

SetTriggeringProblemsCount sets TriggeringProblemsCount field to given value.


### GetTriggeringProblemsDetails

`func (o *SyntheticOnDemandExecutionResult) GetTriggeringProblemsDetails() []SyntheticOnDemandTriggeringProblemDetails`

GetTriggeringProblemsDetails returns the TriggeringProblemsDetails field if non-nil, zero value otherwise.

### GetTriggeringProblemsDetailsOk

`func (o *SyntheticOnDemandExecutionResult) GetTriggeringProblemsDetailsOk() (*[]SyntheticOnDemandTriggeringProblemDetails, bool)`

GetTriggeringProblemsDetailsOk returns a tuple with the TriggeringProblemsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeringProblemsDetails

`func (o *SyntheticOnDemandExecutionResult) SetTriggeringProblemsDetails(v []SyntheticOnDemandTriggeringProblemDetails)`

SetTriggeringProblemsDetails sets TriggeringProblemsDetails field to given value.

### HasTriggeringProblemsDetails

`func (o *SyntheticOnDemandExecutionResult) HasTriggeringProblemsDetails() bool`

HasTriggeringProblemsDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


