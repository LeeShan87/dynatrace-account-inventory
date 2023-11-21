# SyntheticOnDemandExecutionRequestMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomizedScript** | Pointer to **map[string]interface{}** | Customized script properties for this on-demand batch execution. | [optional] 
**ExecutionCount** | Pointer to **int32** | The number of times the monitor is to be executed per location; if not set, the monitor will be executed once. | [optional] [default to 1]
**Locations** | Pointer to **[]string** | The locations from where the monitor is to be executed. | [optional] 
**MonitorId** | **string** | The monitor identifier. | 
**RepeatMode** | Pointer to **string** | Execution repeat mode. If not set, the mode is SEQUENTIAL. | [optional] [default to "SEQUENTIAL"]

## Methods

### NewSyntheticOnDemandExecutionRequestMonitor

`func NewSyntheticOnDemandExecutionRequestMonitor(monitorId string, ) *SyntheticOnDemandExecutionRequestMonitor`

NewSyntheticOnDemandExecutionRequestMonitor instantiates a new SyntheticOnDemandExecutionRequestMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticOnDemandExecutionRequestMonitorWithDefaults

`func NewSyntheticOnDemandExecutionRequestMonitorWithDefaults() *SyntheticOnDemandExecutionRequestMonitor`

NewSyntheticOnDemandExecutionRequestMonitorWithDefaults instantiates a new SyntheticOnDemandExecutionRequestMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomizedScript

`func (o *SyntheticOnDemandExecutionRequestMonitor) GetCustomizedScript() map[string]interface{}`

GetCustomizedScript returns the CustomizedScript field if non-nil, zero value otherwise.

### GetCustomizedScriptOk

`func (o *SyntheticOnDemandExecutionRequestMonitor) GetCustomizedScriptOk() (*map[string]interface{}, bool)`

GetCustomizedScriptOk returns a tuple with the CustomizedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedScript

`func (o *SyntheticOnDemandExecutionRequestMonitor) SetCustomizedScript(v map[string]interface{})`

SetCustomizedScript sets CustomizedScript field to given value.

### HasCustomizedScript

`func (o *SyntheticOnDemandExecutionRequestMonitor) HasCustomizedScript() bool`

HasCustomizedScript returns a boolean if a field has been set.

### GetExecutionCount

`func (o *SyntheticOnDemandExecutionRequestMonitor) GetExecutionCount() int32`

GetExecutionCount returns the ExecutionCount field if non-nil, zero value otherwise.

### GetExecutionCountOk

`func (o *SyntheticOnDemandExecutionRequestMonitor) GetExecutionCountOk() (*int32, bool)`

GetExecutionCountOk returns a tuple with the ExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionCount

`func (o *SyntheticOnDemandExecutionRequestMonitor) SetExecutionCount(v int32)`

SetExecutionCount sets ExecutionCount field to given value.

### HasExecutionCount

`func (o *SyntheticOnDemandExecutionRequestMonitor) HasExecutionCount() bool`

HasExecutionCount returns a boolean if a field has been set.

### GetLocations

`func (o *SyntheticOnDemandExecutionRequestMonitor) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticOnDemandExecutionRequestMonitor) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticOnDemandExecutionRequestMonitor) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SyntheticOnDemandExecutionRequestMonitor) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetMonitorId

`func (o *SyntheticOnDemandExecutionRequestMonitor) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *SyntheticOnDemandExecutionRequestMonitor) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *SyntheticOnDemandExecutionRequestMonitor) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.


### GetRepeatMode

`func (o *SyntheticOnDemandExecutionRequestMonitor) GetRepeatMode() string`

GetRepeatMode returns the RepeatMode field if non-nil, zero value otherwise.

### GetRepeatModeOk

`func (o *SyntheticOnDemandExecutionRequestMonitor) GetRepeatModeOk() (*string, bool)`

GetRepeatModeOk returns a tuple with the RepeatMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatMode

`func (o *SyntheticOnDemandExecutionRequestMonitor) SetRepeatMode(v string)`

SetRepeatMode sets RepeatMode field to given value.

### HasRepeatMode

`func (o *SyntheticOnDemandExecutionRequestMonitor) HasRepeatMode() bool`

HasRepeatMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


