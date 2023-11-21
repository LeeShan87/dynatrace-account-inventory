# AttackEntrypoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CodeLocation** | Pointer to [**CodeLocation**](CodeLocation.md) |  | [optional] 
**EntrypointFunction** | Pointer to [**FunctionDefinition**](FunctionDefinition.md) |  | [optional] 
**Payload** | Pointer to [**[]AttackEntrypointPayloadInner**](AttackEntrypointPayloadInner.md) | All relevant payload data that has been sent during the attack. | [optional] [readonly] 

## Methods

### NewAttackEntrypoint

`func NewAttackEntrypoint() *AttackEntrypoint`

NewAttackEntrypoint instantiates a new AttackEntrypoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackEntrypointWithDefaults

`func NewAttackEntrypointWithDefaults() *AttackEntrypoint`

NewAttackEntrypointWithDefaults instantiates a new AttackEntrypoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCodeLocation

`func (o *AttackEntrypoint) GetCodeLocation() CodeLocation`

GetCodeLocation returns the CodeLocation field if non-nil, zero value otherwise.

### GetCodeLocationOk

`func (o *AttackEntrypoint) GetCodeLocationOk() (*CodeLocation, bool)`

GetCodeLocationOk returns a tuple with the CodeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLocation

`func (o *AttackEntrypoint) SetCodeLocation(v CodeLocation)`

SetCodeLocation sets CodeLocation field to given value.

### HasCodeLocation

`func (o *AttackEntrypoint) HasCodeLocation() bool`

HasCodeLocation returns a boolean if a field has been set.

### GetEntrypointFunction

`func (o *AttackEntrypoint) GetEntrypointFunction() FunctionDefinition`

GetEntrypointFunction returns the EntrypointFunction field if non-nil, zero value otherwise.

### GetEntrypointFunctionOk

`func (o *AttackEntrypoint) GetEntrypointFunctionOk() (*FunctionDefinition, bool)`

GetEntrypointFunctionOk returns a tuple with the EntrypointFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypointFunction

`func (o *AttackEntrypoint) SetEntrypointFunction(v FunctionDefinition)`

SetEntrypointFunction sets EntrypointFunction field to given value.

### HasEntrypointFunction

`func (o *AttackEntrypoint) HasEntrypointFunction() bool`

HasEntrypointFunction returns a boolean if a field has been set.

### GetPayload

`func (o *AttackEntrypoint) GetPayload() []AttackEntrypointPayloadInner`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AttackEntrypoint) GetPayloadOk() (*[]AttackEntrypointPayloadInner, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AttackEntrypoint) SetPayload(v []AttackEntrypointPayloadInner)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AttackEntrypoint) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


