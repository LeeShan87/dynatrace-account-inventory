# Condition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the condition.    It indicates which part of the **services** is checked by the condition. | 
**Operator** | **string** | The operator of the condition. | 
**Values** | **[]string** | A list of reference values of the condition. | 

## Methods

### NewCondition

`func NewCondition(name string, operator string, values []string, ) *Condition`

NewCondition instantiates a new Condition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionWithDefaults

`func NewConditionWithDefaults() *Condition`

NewConditionWithDefaults instantiates a new Condition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Condition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Condition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Condition) SetName(v string)`

SetName sets Name field to given value.


### GetOperator

`func (o *Condition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Condition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Condition) SetOperator(v string)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *Condition) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Condition) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Condition) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


