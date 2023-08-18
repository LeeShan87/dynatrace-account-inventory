# MonitoredStates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitoringStates** | Pointer to [**[]MonitoredEntityStates**](MonitoredEntityStates.md) | A list of process group instances and their monitoring states. | [optional] 
**TotalCount** | Pointer to **int64** | The total number of entities in the response. | [optional] 

## Methods

### NewMonitoredStates

`func NewMonitoredStates() *MonitoredStates`

NewMonitoredStates instantiates a new MonitoredStates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoredStatesWithDefaults

`func NewMonitoredStatesWithDefaults() *MonitoredStates`

NewMonitoredStatesWithDefaults instantiates a new MonitoredStates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitoringStates

`func (o *MonitoredStates) GetMonitoringStates() []MonitoredEntityStates`

GetMonitoringStates returns the MonitoringStates field if non-nil, zero value otherwise.

### GetMonitoringStatesOk

`func (o *MonitoredStates) GetMonitoringStatesOk() (*[]MonitoredEntityStates, bool)`

GetMonitoringStatesOk returns a tuple with the MonitoringStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringStates

`func (o *MonitoredStates) SetMonitoringStates(v []MonitoredEntityStates)`

SetMonitoringStates sets MonitoringStates field to given value.

### HasMonitoringStates

`func (o *MonitoredStates) HasMonitoringStates() bool`

HasMonitoringStates returns a boolean if a field has been set.

### GetTotalCount

`func (o *MonitoredStates) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *MonitoredStates) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *MonitoredStates) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *MonitoredStates) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


