# ExecutionStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorType** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;BROWSER&#x60; -&gt; BMAction  * &#x60;HTTP&#x60; -&gt; MonitorRequestExecutionResult   | 

## Methods

### NewExecutionStep

`func NewExecutionStep(monitorType string, ) *ExecutionStep`

NewExecutionStep instantiates a new ExecutionStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionStepWithDefaults

`func NewExecutionStepWithDefaults() *ExecutionStep`

NewExecutionStepWithDefaults instantiates a new ExecutionStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorType

`func (o *ExecutionStep) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *ExecutionStep) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *ExecutionStep) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


