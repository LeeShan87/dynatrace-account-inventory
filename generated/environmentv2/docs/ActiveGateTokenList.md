# ActiveGateTokenList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveGateTokens** | Pointer to  | A list of ActiveGate tokens. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 

## Methods

### NewActiveGateTokenList

`func NewActiveGateTokenList(totalCount int64, ) *ActiveGateTokenList`

NewActiveGateTokenList instantiates a new ActiveGateTokenList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGateTokenListWithDefaults

`func NewActiveGateTokenListWithDefaults() *ActiveGateTokenList`

NewActiveGateTokenListWithDefaults instantiates a new ActiveGateTokenList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveGateTokens

`func (o *ActiveGateTokenList) GetActiveGateTokens() []ActiveGateToken`

GetActiveGateTokens returns the ActiveGateTokens field if non-nil, zero value otherwise.

### GetActiveGateTokensOk

`func (o *ActiveGateTokenList) GetActiveGateTokensOk() (*[]ActiveGateToken, bool)`

GetActiveGateTokensOk returns a tuple with the ActiveGateTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveGateTokens

`func (o *ActiveGateTokenList) SetActiveGateTokens(v []ActiveGateToken)`

SetActiveGateTokens sets ActiveGateTokens field to given value.

### HasActiveGateTokens

`func (o *ActiveGateTokenList) HasActiveGateTokens() bool`

HasActiveGateTokens returns a boolean if a field has been set.

### GetNextPageKey

`func (o *ActiveGateTokenList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *ActiveGateTokenList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *ActiveGateTokenList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *ActiveGateTokenList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetPageSize

`func (o *ActiveGateTokenList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ActiveGateTokenList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ActiveGateTokenList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *ActiveGateTokenList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *ActiveGateTokenList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ActiveGateTokenList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ActiveGateTokenList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


