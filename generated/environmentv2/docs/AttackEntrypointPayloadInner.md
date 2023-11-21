# AttackEntrypointPayloadInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TruncationInfo** | Pointer to [**TruncationInfo**](TruncationInfo.md) |  | [optional] 
**Values** | Pointer to [**[]EntrypointPayload**](EntrypointPayload.md) | Values of the list. | [optional] [readonly] 

## Methods

### NewAttackEntrypointPayloadInner

`func NewAttackEntrypointPayloadInner() *AttackEntrypointPayloadInner`

NewAttackEntrypointPayloadInner instantiates a new AttackEntrypointPayloadInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackEntrypointPayloadInnerWithDefaults

`func NewAttackEntrypointPayloadInnerWithDefaults() *AttackEntrypointPayloadInner`

NewAttackEntrypointPayloadInnerWithDefaults instantiates a new AttackEntrypointPayloadInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTruncationInfo

`func (o *AttackEntrypointPayloadInner) GetTruncationInfo() TruncationInfo`

GetTruncationInfo returns the TruncationInfo field if non-nil, zero value otherwise.

### GetTruncationInfoOk

`func (o *AttackEntrypointPayloadInner) GetTruncationInfoOk() (*TruncationInfo, bool)`

GetTruncationInfoOk returns a tuple with the TruncationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncationInfo

`func (o *AttackEntrypointPayloadInner) SetTruncationInfo(v TruncationInfo)`

SetTruncationInfo sets TruncationInfo field to given value.

### HasTruncationInfo

`func (o *AttackEntrypointPayloadInner) HasTruncationInfo() bool`

HasTruncationInfo returns a boolean if a field has been set.

### GetValues

`func (o *AttackEntrypointPayloadInner) GetValues() []EntrypointPayload`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *AttackEntrypointPayloadInner) GetValuesOk() (*[]EntrypointPayload, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *AttackEntrypointPayloadInner) SetValues(v []EntrypointPayload)`

SetValues sets Values field to given value.

### HasValues

`func (o *AttackEntrypointPayloadInner) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


