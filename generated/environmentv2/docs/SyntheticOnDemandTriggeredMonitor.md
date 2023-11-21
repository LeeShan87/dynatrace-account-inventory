# SyntheticOnDemandTriggeredMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Executions** | [**[]SyntheticOnDemandTriggeredExecutionDetails**](SyntheticOnDemandTriggeredExecutionDetails.md) | The list of triggered executions. | 
**MonitorId** | **string** | The monitor identifier. | 

## Methods

### NewSyntheticOnDemandTriggeredMonitor

`func NewSyntheticOnDemandTriggeredMonitor(executions []SyntheticOnDemandTriggeredExecutionDetails, monitorId string, ) *SyntheticOnDemandTriggeredMonitor`

NewSyntheticOnDemandTriggeredMonitor instantiates a new SyntheticOnDemandTriggeredMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticOnDemandTriggeredMonitorWithDefaults

`func NewSyntheticOnDemandTriggeredMonitorWithDefaults() *SyntheticOnDemandTriggeredMonitor`

NewSyntheticOnDemandTriggeredMonitorWithDefaults instantiates a new SyntheticOnDemandTriggeredMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutions

`func (o *SyntheticOnDemandTriggeredMonitor) GetExecutions() []SyntheticOnDemandTriggeredExecutionDetails`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *SyntheticOnDemandTriggeredMonitor) GetExecutionsOk() (*[]SyntheticOnDemandTriggeredExecutionDetails, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *SyntheticOnDemandTriggeredMonitor) SetExecutions(v []SyntheticOnDemandTriggeredExecutionDetails)`

SetExecutions sets Executions field to given value.


### GetMonitorId

`func (o *SyntheticOnDemandTriggeredMonitor) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *SyntheticOnDemandTriggeredMonitor) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *SyntheticOnDemandTriggeredMonitor) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


