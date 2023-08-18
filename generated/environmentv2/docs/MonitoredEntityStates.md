# MonitoredEntityStates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | The Dynatrace entity ID of the process group instance. | [optional] 
**Params** | Pointer to [**[]MonitoredEntityStateParam**](MonitoredEntityStateParam.md) | Additional parameters of the monitoring state. | [optional] 
**Severity** | Pointer to **string** | The type of the monitoring state. | [optional] 
**State** | Pointer to **string** | The name of the monitoring state. | [optional] 

## Methods

### NewMonitoredEntityStates

`func NewMonitoredEntityStates() *MonitoredEntityStates`

NewMonitoredEntityStates instantiates a new MonitoredEntityStates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoredEntityStatesWithDefaults

`func NewMonitoredEntityStatesWithDefaults() *MonitoredEntityStates`

NewMonitoredEntityStatesWithDefaults instantiates a new MonitoredEntityStates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *MonitoredEntityStates) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *MonitoredEntityStates) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *MonitoredEntityStates) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *MonitoredEntityStates) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetParams

`func (o *MonitoredEntityStates) GetParams() []MonitoredEntityStateParam`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *MonitoredEntityStates) GetParamsOk() (*[]MonitoredEntityStateParam, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *MonitoredEntityStates) SetParams(v []MonitoredEntityStateParam)`

SetParams sets Params field to given value.

### HasParams

`func (o *MonitoredEntityStates) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSeverity

`func (o *MonitoredEntityStates) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *MonitoredEntityStates) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *MonitoredEntityStates) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *MonitoredEntityStates) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetState

`func (o *MonitoredEntityStates) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MonitoredEntityStates) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MonitoredEntityStates) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MonitoredEntityStates) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


