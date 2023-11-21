# AttackList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attacks** | Pointer to [**[]Attack**](Attack.md) | A list of attacks. | [optional] [readonly] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 

## Methods

### NewAttackList

`func NewAttackList(totalCount int64, ) *AttackList`

NewAttackList instantiates a new AttackList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackListWithDefaults

`func NewAttackListWithDefaults() *AttackList`

NewAttackListWithDefaults instantiates a new AttackList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttacks

`func (o *AttackList) GetAttacks() []Attack`

GetAttacks returns the Attacks field if non-nil, zero value otherwise.

### GetAttacksOk

`func (o *AttackList) GetAttacksOk() (*[]Attack, bool)`

GetAttacksOk returns a tuple with the Attacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttacks

`func (o *AttackList) SetAttacks(v []Attack)`

SetAttacks sets Attacks field to given value.

### HasAttacks

`func (o *AttackList) HasAttacks() bool`

HasAttacks returns a boolean if a field has been set.

### GetNextPageKey

`func (o *AttackList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *AttackList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *AttackList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *AttackList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetPageSize

`func (o *AttackList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *AttackList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *AttackList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *AttackList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *AttackList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *AttackList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *AttackList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


