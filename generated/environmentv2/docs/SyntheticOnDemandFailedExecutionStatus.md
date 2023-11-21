# SyntheticOnDemandFailedExecutionStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** | Error code. | 
**ExecutionId** | **string** | The identifier of the execution. | 
**ExecutionStage** | Pointer to **string** | Execution stage. | [optional] 
**ExecutionTimestamp** | Pointer to **int64** | The timestamp when execution was finished, in UTC milliseconds. | [optional] 
**FailureMessage** | Pointer to **string** | Failure message. | [optional] 
**LocationId** | **string** | The identifier of the location from where the monitor is to be executed. | 
**MonitorId** | **string** | The identifier of the monitor. | 

## Methods

### NewSyntheticOnDemandFailedExecutionStatus

`func NewSyntheticOnDemandFailedExecutionStatus(errorCode string, executionId string, locationId string, monitorId string, ) *SyntheticOnDemandFailedExecutionStatus`

NewSyntheticOnDemandFailedExecutionStatus instantiates a new SyntheticOnDemandFailedExecutionStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticOnDemandFailedExecutionStatusWithDefaults

`func NewSyntheticOnDemandFailedExecutionStatusWithDefaults() *SyntheticOnDemandFailedExecutionStatus`

NewSyntheticOnDemandFailedExecutionStatusWithDefaults instantiates a new SyntheticOnDemandFailedExecutionStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *SyntheticOnDemandFailedExecutionStatus) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *SyntheticOnDemandFailedExecutionStatus) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *SyntheticOnDemandFailedExecutionStatus) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetExecutionId

`func (o *SyntheticOnDemandFailedExecutionStatus) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *SyntheticOnDemandFailedExecutionStatus) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *SyntheticOnDemandFailedExecutionStatus) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetExecutionStage

`func (o *SyntheticOnDemandFailedExecutionStatus) GetExecutionStage() string`

GetExecutionStage returns the ExecutionStage field if non-nil, zero value otherwise.

### GetExecutionStageOk

`func (o *SyntheticOnDemandFailedExecutionStatus) GetExecutionStageOk() (*string, bool)`

GetExecutionStageOk returns a tuple with the ExecutionStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStage

`func (o *SyntheticOnDemandFailedExecutionStatus) SetExecutionStage(v string)`

SetExecutionStage sets ExecutionStage field to given value.

### HasExecutionStage

`func (o *SyntheticOnDemandFailedExecutionStatus) HasExecutionStage() bool`

HasExecutionStage returns a boolean if a field has been set.

### GetExecutionTimestamp

`func (o *SyntheticOnDemandFailedExecutionStatus) GetExecutionTimestamp() int64`

GetExecutionTimestamp returns the ExecutionTimestamp field if non-nil, zero value otherwise.

### GetExecutionTimestampOk

`func (o *SyntheticOnDemandFailedExecutionStatus) GetExecutionTimestampOk() (*int64, bool)`

GetExecutionTimestampOk returns a tuple with the ExecutionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimestamp

`func (o *SyntheticOnDemandFailedExecutionStatus) SetExecutionTimestamp(v int64)`

SetExecutionTimestamp sets ExecutionTimestamp field to given value.

### HasExecutionTimestamp

`func (o *SyntheticOnDemandFailedExecutionStatus) HasExecutionTimestamp() bool`

HasExecutionTimestamp returns a boolean if a field has been set.

### GetFailureMessage

`func (o *SyntheticOnDemandFailedExecutionStatus) GetFailureMessage() string`

GetFailureMessage returns the FailureMessage field if non-nil, zero value otherwise.

### GetFailureMessageOk

`func (o *SyntheticOnDemandFailedExecutionStatus) GetFailureMessageOk() (*string, bool)`

GetFailureMessageOk returns a tuple with the FailureMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMessage

`func (o *SyntheticOnDemandFailedExecutionStatus) SetFailureMessage(v string)`

SetFailureMessage sets FailureMessage field to given value.

### HasFailureMessage

`func (o *SyntheticOnDemandFailedExecutionStatus) HasFailureMessage() bool`

HasFailureMessage returns a boolean if a field has been set.

### GetLocationId

`func (o *SyntheticOnDemandFailedExecutionStatus) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *SyntheticOnDemandFailedExecutionStatus) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *SyntheticOnDemandFailedExecutionStatus) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetMonitorId

`func (o *SyntheticOnDemandFailedExecutionStatus) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *SyntheticOnDemandFailedExecutionStatus) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *SyntheticOnDemandFailedExecutionStatus) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


