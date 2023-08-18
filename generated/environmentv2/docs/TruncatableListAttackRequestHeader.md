# TruncatableListAttackRequestHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TruncationInfo** | Pointer to [**TruncationInfo**](TruncationInfo.md) |  | [optional] 
**Values** | Pointer to [**[]AttackRequestHeader**](AttackRequestHeader.md) | Values of the list. | [optional] [readonly] 

## Methods

### NewTruncatableListAttackRequestHeader

`func NewTruncatableListAttackRequestHeader() *TruncatableListAttackRequestHeader`

NewTruncatableListAttackRequestHeader instantiates a new TruncatableListAttackRequestHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTruncatableListAttackRequestHeaderWithDefaults

`func NewTruncatableListAttackRequestHeaderWithDefaults() *TruncatableListAttackRequestHeader`

NewTruncatableListAttackRequestHeaderWithDefaults instantiates a new TruncatableListAttackRequestHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTruncationInfo

`func (o *TruncatableListAttackRequestHeader) GetTruncationInfo() TruncationInfo`

GetTruncationInfo returns the TruncationInfo field if non-nil, zero value otherwise.

### GetTruncationInfoOk

`func (o *TruncatableListAttackRequestHeader) GetTruncationInfoOk() (*TruncationInfo, bool)`

GetTruncationInfoOk returns a tuple with the TruncationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncationInfo

`func (o *TruncatableListAttackRequestHeader) SetTruncationInfo(v TruncationInfo)`

SetTruncationInfo sets TruncationInfo field to given value.

### HasTruncationInfo

`func (o *TruncatableListAttackRequestHeader) HasTruncationInfo() bool`

HasTruncationInfo returns a boolean if a field has been set.

### GetValues

`func (o *TruncatableListAttackRequestHeader) GetValues() []AttackRequestHeader`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *TruncatableListAttackRequestHeader) GetValuesOk() (*[]AttackRequestHeader, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *TruncatableListAttackRequestHeader) SetValues(v []AttackRequestHeader)`

SetValues sets Values field to given value.

### HasValues

`func (o *TruncatableListAttackRequestHeader) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


