# EventTypeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventTypeInfos** | Pointer to [**[]EventType**](EventType.md) | A list of event types. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 

## Methods

### NewEventTypeList

`func NewEventTypeList(totalCount int64, ) *EventTypeList`

NewEventTypeList instantiates a new EventTypeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventTypeListWithDefaults

`func NewEventTypeListWithDefaults() *EventTypeList`

NewEventTypeListWithDefaults instantiates a new EventTypeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventTypeInfos

`func (o *EventTypeList) GetEventTypeInfos() []EventType`

GetEventTypeInfos returns the EventTypeInfos field if non-nil, zero value otherwise.

### GetEventTypeInfosOk

`func (o *EventTypeList) GetEventTypeInfosOk() (*[]EventType, bool)`

GetEventTypeInfosOk returns a tuple with the EventTypeInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTypeInfos

`func (o *EventTypeList) SetEventTypeInfos(v []EventType)`

SetEventTypeInfos sets EventTypeInfos field to given value.

### HasEventTypeInfos

`func (o *EventTypeList) HasEventTypeInfos() bool`

HasEventTypeInfos returns a boolean if a field has been set.

### GetNextPageKey

`func (o *EventTypeList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *EventTypeList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *EventTypeList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *EventTypeList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetPageSize

`func (o *EventTypeList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EventTypeList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EventTypeList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *EventTypeList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *EventTypeList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EventTypeList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EventTypeList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


