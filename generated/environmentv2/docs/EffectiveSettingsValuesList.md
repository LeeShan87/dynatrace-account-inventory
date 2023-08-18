# EffectiveSettingsValuesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]EffectiveSettingsValue**](EffectiveSettingsValue.md) | A list of effective settings values. | 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**PageSize** | **int32** | The number of entries per page. | 
**TotalCount** | **int64** | The total number of entries in the result. | 

## Methods

### NewEffectiveSettingsValuesList

`func NewEffectiveSettingsValuesList(items []EffectiveSettingsValue, pageSize int32, totalCount int64, ) *EffectiveSettingsValuesList`

NewEffectiveSettingsValuesList instantiates a new EffectiveSettingsValuesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectiveSettingsValuesListWithDefaults

`func NewEffectiveSettingsValuesListWithDefaults() *EffectiveSettingsValuesList`

NewEffectiveSettingsValuesListWithDefaults instantiates a new EffectiveSettingsValuesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EffectiveSettingsValuesList) GetItems() []EffectiveSettingsValue`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EffectiveSettingsValuesList) GetItemsOk() (*[]EffectiveSettingsValue, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EffectiveSettingsValuesList) SetItems(v []EffectiveSettingsValue)`

SetItems sets Items field to given value.


### GetNextPageKey

`func (o *EffectiveSettingsValuesList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *EffectiveSettingsValuesList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *EffectiveSettingsValuesList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *EffectiveSettingsValuesList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetPageSize

`func (o *EffectiveSettingsValuesList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EffectiveSettingsValuesList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EffectiveSettingsValuesList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetTotalCount

`func (o *EffectiveSettingsValuesList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EffectiveSettingsValuesList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EffectiveSettingsValuesList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


