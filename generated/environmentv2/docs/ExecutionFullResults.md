# ExecutionFullResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | Error code. | [optional] 
**ExecutionStepCount** | Pointer to **int32** | Number executed steps. | [optional] 
**ExecutionSteps** | Pointer to [**[]ExecutionStep**](ExecutionStep.md) | Details about the monitor&#39;s step execution. | [optional] 
**FailedStepName** | Pointer to **string** | Failed step name. | [optional] 
**FailedStepSequenceId** | Pointer to **int32** | Failed step sequence id. | [optional] 
**FailureMessage** | Pointer to **string** | Failure message. | [optional] 
**Status** | Pointer to **string** | Execution status. | [optional] 

## Methods

### NewExecutionFullResults

`func NewExecutionFullResults() *ExecutionFullResults`

NewExecutionFullResults instantiates a new ExecutionFullResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionFullResultsWithDefaults

`func NewExecutionFullResultsWithDefaults() *ExecutionFullResults`

NewExecutionFullResultsWithDefaults instantiates a new ExecutionFullResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ExecutionFullResults) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ExecutionFullResults) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ExecutionFullResults) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ExecutionFullResults) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetExecutionStepCount

`func (o *ExecutionFullResults) GetExecutionStepCount() int32`

GetExecutionStepCount returns the ExecutionStepCount field if non-nil, zero value otherwise.

### GetExecutionStepCountOk

`func (o *ExecutionFullResults) GetExecutionStepCountOk() (*int32, bool)`

GetExecutionStepCountOk returns a tuple with the ExecutionStepCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStepCount

`func (o *ExecutionFullResults) SetExecutionStepCount(v int32)`

SetExecutionStepCount sets ExecutionStepCount field to given value.

### HasExecutionStepCount

`func (o *ExecutionFullResults) HasExecutionStepCount() bool`

HasExecutionStepCount returns a boolean if a field has been set.

### GetExecutionSteps

`func (o *ExecutionFullResults) GetExecutionSteps() []ExecutionStep`

GetExecutionSteps returns the ExecutionSteps field if non-nil, zero value otherwise.

### GetExecutionStepsOk

`func (o *ExecutionFullResults) GetExecutionStepsOk() (*[]ExecutionStep, bool)`

GetExecutionStepsOk returns a tuple with the ExecutionSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionSteps

`func (o *ExecutionFullResults) SetExecutionSteps(v []ExecutionStep)`

SetExecutionSteps sets ExecutionSteps field to given value.

### HasExecutionSteps

`func (o *ExecutionFullResults) HasExecutionSteps() bool`

HasExecutionSteps returns a boolean if a field has been set.

### GetFailedStepName

`func (o *ExecutionFullResults) GetFailedStepName() string`

GetFailedStepName returns the FailedStepName field if non-nil, zero value otherwise.

### GetFailedStepNameOk

`func (o *ExecutionFullResults) GetFailedStepNameOk() (*string, bool)`

GetFailedStepNameOk returns a tuple with the FailedStepName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedStepName

`func (o *ExecutionFullResults) SetFailedStepName(v string)`

SetFailedStepName sets FailedStepName field to given value.

### HasFailedStepName

`func (o *ExecutionFullResults) HasFailedStepName() bool`

HasFailedStepName returns a boolean if a field has been set.

### GetFailedStepSequenceId

`func (o *ExecutionFullResults) GetFailedStepSequenceId() int32`

GetFailedStepSequenceId returns the FailedStepSequenceId field if non-nil, zero value otherwise.

### GetFailedStepSequenceIdOk

`func (o *ExecutionFullResults) GetFailedStepSequenceIdOk() (*int32, bool)`

GetFailedStepSequenceIdOk returns a tuple with the FailedStepSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedStepSequenceId

`func (o *ExecutionFullResults) SetFailedStepSequenceId(v int32)`

SetFailedStepSequenceId sets FailedStepSequenceId field to given value.

### HasFailedStepSequenceId

`func (o *ExecutionFullResults) HasFailedStepSequenceId() bool`

HasFailedStepSequenceId returns a boolean if a field has been set.

### GetFailureMessage

`func (o *ExecutionFullResults) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *ExecutionFullResults) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *ExecutionFullResults) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *ExecutionFullResults) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionFullResults) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionFullResults) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionFullResults) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionFullResults) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


