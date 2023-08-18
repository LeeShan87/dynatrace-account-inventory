# SecurityProblemEventsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | Pointer to [**[]SecurityProblemEvent**](SecurityProblemEvent.md) | A list of events for a security problem. | [optional] [readonly] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 

## Methods

### NewSecurityProblemEventsList

`func NewSecurityProblemEventsList(totalCount int64, ) *SecurityProblemEventsList`

NewSecurityProblemEventsList instantiates a new SecurityProblemEventsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemEventsListWithDefaults

`func NewSecurityProblemEventsListWithDefaults() *SecurityProblemEventsList`

NewSecurityProblemEventsListWithDefaults instantiates a new SecurityProblemEventsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *SecurityProblemEventsList) GetEvents() []SecurityProblemEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SecurityProblemEventsList) GetEventsOk() (*[]SecurityProblemEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SecurityProblemEventsList) SetEvents(v []SecurityProblemEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SecurityProblemEventsList) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetNextPageKey

`func (o *SecurityProblemEventsList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *SecurityProblemEventsList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *SecurityProblemEventsList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *SecurityProblemEventsList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetPageSize

`func (o *SecurityProblemEventsList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *SecurityProblemEventsList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *SecurityProblemEventsList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *SecurityProblemEventsList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *SecurityProblemEventsList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *SecurityProblemEventsList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *SecurityProblemEventsList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


