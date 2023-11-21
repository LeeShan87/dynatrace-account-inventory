# SyntheticOnDemandBatchStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | **string** | The identifier of the batch. | 
**BatchStatus** | **string** | The status of the batch. | 
**ExecutedCount** | **int32** | The number of triggered executions with the result SUCCESS or FAILED. | 
**FailedCount** | **int32** | The number of triggered executions with the result FAILED. | 
**FailedExecutions** | Pointer to [**[]SyntheticOnDemandFailedExecutionStatus**](SyntheticOnDemandFailedExecutionStatus.md) |  | [optional] 
**FailedToExecute** | Pointer to [**[]SyntheticOnDemandFailedExecutionStatus**](SyntheticOnDemandFailedExecutionStatus.md) |  | [optional] 
**FailedToExecuteCount** | **int32** | The number of executions that were triggered and timed out because of a problem with the Synthetic engine. | 
**Metadata** | Pointer to **map[string]string** | String to string map of metadata properties for batch | [optional] 
**TriggeredCount** | **int32** | The number of triggered executions within the batch. | 
**TriggeringProblems** | Pointer to [**[]SyntheticOnDemandTriggeringProblemDetails**](SyntheticOnDemandTriggeringProblemDetails.md) |  | [optional] 
**TriggeringProblemsCount** | Pointer to **int32** | The number of executions that were not triggered due to some problems. | [optional] 
**UserId** | **string** | The name of the user who triggered execution of the batch. | 

## Methods

### NewSyntheticOnDemandBatchStatus

`func NewSyntheticOnDemandBatchStatus(batchId string, batchStatus string, executedCount int32, failedCount int32, failedToExecuteCount int32, triggeredCount int32, userId string, ) *SyntheticOnDemandBatchStatus`

NewSyntheticOnDemandBatchStatus instantiates a new SyntheticOnDemandBatchStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticOnDemandBatchStatusWithDefaults

`func NewSyntheticOnDemandBatchStatusWithDefaults() *SyntheticOnDemandBatchStatus`

NewSyntheticOnDemandBatchStatusWithDefaults instantiates a new SyntheticOnDemandBatchStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *SyntheticOnDemandBatchStatus) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *SyntheticOnDemandBatchStatus) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *SyntheticOnDemandBatchStatus) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.


### GetBatchStatus

`func (o *SyntheticOnDemandBatchStatus) GetBatchStatus() string`

GetBatchStatus returns the BatchStatus field if non-nil, zero value otherwise.

### GetBatchStatusOk

`func (o *SyntheticOnDemandBatchStatus) GetBatchStatusOk() (*string, bool)`

GetBatchStatusOk returns a tuple with the BatchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchStatus

`func (o *SyntheticOnDemandBatchStatus) SetBatchStatus(v string)`

SetBatchStatus sets BatchStatus field to given value.


### GetExecutedCount

`func (o *SyntheticOnDemandBatchStatus) GetExecutedCount() int32`

GetExecutedCount returns the ExecutedCount field if non-nil, zero value otherwise.

### GetExecutedCountOk

`func (o *SyntheticOnDemandBatchStatus) GetExecutedCountOk() (*int32, bool)`

GetExecutedCountOk returns a tuple with the ExecutedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedCount

`func (o *SyntheticOnDemandBatchStatus) SetExecutedCount(v int32)`

SetExecutedCount sets ExecutedCount field to given value.


### GetFailedCount

`func (o *SyntheticOnDemandBatchStatus) GetFailedCount() int32`

GetFailedCount returns the FailedCount field if non-nil, zero value otherwise.

### GetFailedCountOk

`func (o *SyntheticOnDemandBatchStatus) GetFailedCountOk() (*int32, bool)`

GetFailedCountOk returns a tuple with the FailedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedCount

`func (o *SyntheticOnDemandBatchStatus) SetFailedCount(v int32)`

SetFailedCount sets FailedCount field to given value.


### GetFailedExecutions

`func (o *SyntheticOnDemandBatchStatus) GetFailedExecutions() []SyntheticOnDemandFailedExecutionStatus`

GetFailedExecutions returns the FailedExecutions field if non-nil, zero value otherwise.

### GetFailedExecutionsOk

`func (o *SyntheticOnDemandBatchStatus) GetFailedExecutionsOk() (*[]SyntheticOnDemandFailedExecutionStatus, bool)`

GetFailedExecutionsOk returns a tuple with the FailedExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedExecutions

`func (o *SyntheticOnDemandBatchStatus) SetFailedExecutions(v []SyntheticOnDemandFailedExecutionStatus)`

SetFailedExecutions sets FailedExecutions field to given value.

### HasFailedExecutions

`func (o *SyntheticOnDemandBatchStatus) HasFailedExecutions() bool`

HasFailedExecutions returns a boolean if a field has been set.

### GetFailedToExecute

`func (o *SyntheticOnDemandBatchStatus) GetFailedToExecute() []SyntheticOnDemandFailedExecutionStatus`

GetFailedToExecute returns the FailedToExecute field if non-nil, zero value otherwise.

### GetFailedToExecuteOk

`func (o *SyntheticOnDemandBatchStatus) GetFailedToExecuteOk() (*[]SyntheticOnDemandFailedExecutionStatus, bool)`

GetFailedToExecuteOk returns a tuple with the FailedToExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToExecute

`func (o *SyntheticOnDemandBatchStatus) SetFailedToExecute(v []SyntheticOnDemandFailedExecutionStatus)`

SetFailedToExecute sets FailedToExecute field to given value.

### HasFailedToExecute

`func (o *SyntheticOnDemandBatchStatus) HasFailedToExecute() bool`

HasFailedToExecute returns a boolean if a field has been set.

### GetFailedToExecuteCount

`func (o *SyntheticOnDemandBatchStatus) GetFailedToExecuteCount() int32`

GetFailedToExecuteCount returns the FailedToExecuteCount field if non-nil, zero value otherwise.

### GetFailedToExecuteCountOk

`func (o *SyntheticOnDemandBatchStatus) GetFailedToExecuteCountOk() (*int32, bool)`

GetFailedToExecuteCountOk returns a tuple with the FailedToExecuteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedToExecuteCount

`func (o *SyntheticOnDemandBatchStatus) SetFailedToExecuteCount(v int32)`

SetFailedToExecuteCount sets FailedToExecuteCount field to given value.


### GetMetadata

`func (o *SyntheticOnDemandBatchStatus) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SyntheticOnDemandBatchStatus) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SyntheticOnDemandBatchStatus) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SyntheticOnDemandBatchStatus) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTriggeredCount

`func (o *SyntheticOnDemandBatchStatus) GetTriggeredCount() int32`

GetTriggeredCount returns the TriggeredCount field if non-nil, zero value otherwise.

### GetTriggeredCountOk

`func (o *SyntheticOnDemandBatchStatus) GetTriggeredCountOk() (*int32, bool)`

GetTriggeredCountOk returns a tuple with the TriggeredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredCount

`func (o *SyntheticOnDemandBatchStatus) SetTriggeredCount(v int32)`

SetTriggeredCount sets TriggeredCount field to given value.


### GetTriggeringProblems

`func (o *SyntheticOnDemandBatchStatus) GetTriggeringProblems() []SyntheticOnDemandTriggeringProblemDetails`

GetTriggeringProblems returns the TriggeringProblems field if non-nil, zero value otherwise.

### GetTriggeringProblemsOk

`func (o *SyntheticOnDemandBatchStatus) GetTriggeringProblemsOk() (*[]SyntheticOnDemandTriggeringProblemDetails, bool)`

GetTriggeringProblemsOk returns a tuple with the TriggeringProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeringProblems

`func (o *SyntheticOnDemandBatchStatus) SetTriggeringProblems(v []SyntheticOnDemandTriggeringProblemDetails)`

SetTriggeringProblems sets TriggeringProblems field to given value.

### HasTriggeringProblems

`func (o *SyntheticOnDemandBatchStatus) HasTriggeringProblems() bool`

HasTriggeringProblems returns a boolean if a field has been set.

### GetTriggeringProblemsCount

`func (o *SyntheticOnDemandBatchStatus) GetTriggeringProblemsCount() int32`

GetTriggeringProblemsCount returns the TriggeringProblemsCount field if non-nil, zero value otherwise.

### GetTriggeringProblemsCountOk

`func (o *SyntheticOnDemandBatchStatus) GetTriggeringProblemsCountOk() (*int32, bool)`

GetTriggeringProblemsCountOk returns a tuple with the TriggeringProblemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeringProblemsCount

`func (o *SyntheticOnDemandBatchStatus) SetTriggeringProblemsCount(v int32)`

SetTriggeringProblemsCount sets TriggeringProblemsCount field to given value.

### HasTriggeringProblemsCount

`func (o *SyntheticOnDemandBatchStatus) HasTriggeringProblemsCount() bool`

HasTriggeringProblemsCount returns a boolean if a field has been set.

### GetUserId

`func (o *SyntheticOnDemandBatchStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SyntheticOnDemandBatchStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SyntheticOnDemandBatchStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


