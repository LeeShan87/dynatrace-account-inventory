# SyntheticOnDemandExecution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchId** | **string** | The identifier of the batch. | 
**CustomizedScript** | Pointer to **map[string]interface{}** | Customized script properties for this on-demand batch execution. | [optional] 
**DataDeliveryTimestamp** | **int64** | The timestamp when whole data set has been collected on server, in UTC milliseconds. | 
**ExecutionId** | **string** | The identifier of the execution. | 
**ExecutionStage** | **string** | Execution stage. | 
**ExecutionTimestamp** | **int64** | The timestamp when execution was finished, in UTC milliseconds. | 
**FullResults** | Pointer to [**ExecutionFullResults**](ExecutionFullResults.md) |  | [optional] 
**LocationId** | **string** | The identifier of the location from where the monitor is to be executed. | 
**Metadata** | Pointer to **map[string]string** | Metadata map for the execution batch. | [optional] 
**MonitorId** | **string** | The identifier of the monitor. | 
**NextExecutionId** | Pointer to **int64** | Next execution id for sequential mode. | [optional] 
**ProcessingMode** | **string** | The processing mode of the execution. | 
**SchedulingTimestamp** | **int64** | The scheduling timestamp, in UTC milliseconds. | 
**SimpleResults** | Pointer to [**ExecutionSimpleResults**](ExecutionSimpleResults.md) |  | [optional] 
**Source** | **string** | The source of the triggering request. | 
**UserId** | **string** | The name of the user who triggered the on-demand execution. | 

## Methods

### NewSyntheticOnDemandExecution

`func NewSyntheticOnDemandExecution(batchId string, dataDeliveryTimestamp int64, executionId string, executionStage string, executionTimestamp int64, locationId string, monitorId string, processingMode string, schedulingTimestamp int64, source string, userId string, ) *SyntheticOnDemandExecution`

NewSyntheticOnDemandExecution instantiates a new SyntheticOnDemandExecution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticOnDemandExecutionWithDefaults

`func NewSyntheticOnDemandExecutionWithDefaults() *SyntheticOnDemandExecution`

NewSyntheticOnDemandExecutionWithDefaults instantiates a new SyntheticOnDemandExecution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchId

`func (o *SyntheticOnDemandExecution) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *SyntheticOnDemandExecution) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *SyntheticOnDemandExecution) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.


### GetCustomizedScript

`func (o *SyntheticOnDemandExecution) GetCustomizedScript() map[string]interface{}`

GetCustomizedScript returns the CustomizedScript field if non-nil, zero value otherwise.

### GetCustomizedScriptOk

`func (o *SyntheticOnDemandExecution) GetCustomizedScriptOk() (*map[string]interface{}, bool)`

GetCustomizedScriptOk returns a tuple with the CustomizedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedScript

`func (o *SyntheticOnDemandExecution) SetCustomizedScript(v map[string]interface{})`

SetCustomizedScript sets CustomizedScript field to given value.

### HasCustomizedScript

`func (o *SyntheticOnDemandExecution) HasCustomizedScript() bool`

HasCustomizedScript returns a boolean if a field has been set.

### GetDataDeliveryTimestamp

`func (o *SyntheticOnDemandExecution) GetDataDeliveryTimestamp() int64`

GetDataDeliveryTimestamp returns the DataDeliveryTimestamp field if non-nil, zero value otherwise.

### GetDataDeliveryTimestampOk

`func (o *SyntheticOnDemandExecution) GetDataDeliveryTimestampOk() (*int64, bool)`

GetDataDeliveryTimestampOk returns a tuple with the DataDeliveryTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDeliveryTimestamp

`func (o *SyntheticOnDemandExecution) SetDataDeliveryTimestamp(v int64)`

SetDataDeliveryTimestamp sets DataDeliveryTimestamp field to given value.


### GetExecutionId

`func (o *SyntheticOnDemandExecution) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *SyntheticOnDemandExecution) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *SyntheticOnDemandExecution) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetExecutionStage

`func (o *SyntheticOnDemandExecution) GetExecutionStage() string`

GetExecutionStage returns the ExecutionStage field if non-nil, zero value otherwise.

### GetExecutionStageOk

`func (o *SyntheticOnDemandExecution) GetExecutionStageOk() (*string, bool)`

GetExecutionStageOk returns a tuple with the ExecutionStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionStage

`func (o *SyntheticOnDemandExecution) SetExecutionStage(v string)`

SetExecutionStage sets ExecutionStage field to given value.


### GetExecutionTimestamp

`func (o *SyntheticOnDemandExecution) GetExecutionTimestamp() int64`

GetExecutionTimestamp returns the ExecutionTimestamp field if non-nil, zero value otherwise.

### GetExecutionTimestampOk

`func (o *SyntheticOnDemandExecution) GetExecutionTimestampOk() (*int64, bool)`

GetExecutionTimestampOk returns a tuple with the ExecutionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTimestamp

`func (o *SyntheticOnDemandExecution) SetExecutionTimestamp(v int64)`

SetExecutionTimestamp sets ExecutionTimestamp field to given value.


### GetFullResults

`func (o *SyntheticOnDemandExecution) GetFullResults() ExecutionFullResults`

GetFullResults returns the FullResults field if non-nil, zero value otherwise.

### GetFullResultsOk

`func (o *SyntheticOnDemandExecution) GetFullResultsOk() (*ExecutionFullResults, bool)`

GetFullResultsOk returns a tuple with the FullResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullResults

`func (o *SyntheticOnDemandExecution) SetFullResults(v ExecutionFullResults)`

SetFullResults sets FullResults field to given value.

### HasFullResults

`func (o *SyntheticOnDemandExecution) HasFullResults() bool`

HasFullResults returns a boolean if a field has been set.

### GetLocationId

`func (o *SyntheticOnDemandExecution) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *SyntheticOnDemandExecution) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *SyntheticOnDemandExecution) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetMetadata

`func (o *SyntheticOnDemandExecution) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SyntheticOnDemandExecution) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SyntheticOnDemandExecution) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SyntheticOnDemandExecution) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMonitorId

`func (o *SyntheticOnDemandExecution) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *SyntheticOnDemandExecution) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *SyntheticOnDemandExecution) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.


### GetNextExecutionId

`func (o *SyntheticOnDemandExecution) GetNextExecutionId() int64`

GetNextExecutionId returns the NextExecutionId field if non-nil, zero value otherwise.

### GetNextExecutionIdOk

`func (o *SyntheticOnDemandExecution) GetNextExecutionIdOk() (*int64, bool)`

GetNextExecutionIdOk returns a tuple with the NextExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecutionId

`func (o *SyntheticOnDemandExecution) SetNextExecutionId(v int64)`

SetNextExecutionId sets NextExecutionId field to given value.

### HasNextExecutionId

`func (o *SyntheticOnDemandExecution) HasNextExecutionId() bool`

HasNextExecutionId returns a boolean if a field has been set.

### GetProcessingMode

`func (o *SyntheticOnDemandExecution) GetProcessingMode() string`

GetProcessingMode returns the ProcessingMode field if non-nil, zero value otherwise.

### GetProcessingModeOk

`func (o *SyntheticOnDemandExecution) GetProcessingModeOk() (*string, bool)`

GetProcessingModeOk returns a tuple with the ProcessingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessingMode

`func (o *SyntheticOnDemandExecution) SetProcessingMode(v string)`

SetProcessingMode sets ProcessingMode field to given value.


### GetSchedulingTimestamp

`func (o *SyntheticOnDemandExecution) GetSchedulingTimestamp() int64`

GetSchedulingTimestamp returns the SchedulingTimestamp field if non-nil, zero value otherwise.

### GetSchedulingTimestampOk

`func (o *SyntheticOnDemandExecution) GetSchedulingTimestampOk() (*int64, bool)`

GetSchedulingTimestampOk returns a tuple with the SchedulingTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedulingTimestamp

`func (o *SyntheticOnDemandExecution) SetSchedulingTimestamp(v int64)`

SetSchedulingTimestamp sets SchedulingTimestamp field to given value.


### GetSimpleResults

`func (o *SyntheticOnDemandExecution) GetSimpleResults() ExecutionSimpleResults`

GetSimpleResults returns the SimpleResults field if non-nil, zero value otherwise.

### GetSimpleResultsOk

`func (o *SyntheticOnDemandExecution) GetSimpleResultsOk() (*ExecutionSimpleResults, bool)`

GetSimpleResultsOk returns a tuple with the SimpleResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleResults

`func (o *SyntheticOnDemandExecution) SetSimpleResults(v ExecutionSimpleResults)`

SetSimpleResults sets SimpleResults field to given value.

### HasSimpleResults

`func (o *SyntheticOnDemandExecution) HasSimpleResults() bool`

HasSimpleResults returns a boolean if a field has been set.

### GetSource

`func (o *SyntheticOnDemandExecution) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *SyntheticOnDemandExecution) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *SyntheticOnDemandExecution) SetSource(v string)`

SetSource sets Source field to given value.


### GetUserId

`func (o *SyntheticOnDemandExecution) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SyntheticOnDemandExecution) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SyntheticOnDemandExecution) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


