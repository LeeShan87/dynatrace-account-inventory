# EventPropertiesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventProperties** | Pointer to [**[]EventPropertyDetails**](EventPropertyDetails.md) | A list of event properties. | [optional] 
**NextPageKey** | Pointer to **string** | The cursor for the next page of results. Has the value of &#x60;null&#x60; on the last page.   Use it in the **nextPageKey** query parameter to obtain subsequent pages of the result. | [optional] 
**PageSize** | Pointer to **int32** | The number of entries per page. | [optional] 
**TotalCount** | **int64** | The total number of entries in the result. | 

## Methods

### NewEventPropertiesList

`func NewEventPropertiesList(totalCount int64, ) *EventPropertiesList`

NewEventPropertiesList instantiates a new EventPropertiesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventPropertiesListWithDefaults

`func NewEventPropertiesListWithDefaults() *EventPropertiesList`

NewEventPropertiesListWithDefaults instantiates a new EventPropertiesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventProperties

`func (o *EventPropertiesList) GetEventProperties() []EventPropertyDetails`

GetEventProperties returns the EventProperties field if non-nil, zero value otherwise.

### GetEventPropertiesOk

`func (o *EventPropertiesList) GetEventPropertiesOk() (*[]EventPropertyDetails, bool)`

GetEventPropertiesOk returns a tuple with the EventProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventProperties

`func (o *EventPropertiesList) SetEventProperties(v []EventPropertyDetails)`

SetEventProperties sets EventProperties field to given value.

### HasEventProperties

`func (o *EventPropertiesList) HasEventProperties() bool`

HasEventProperties returns a boolean if a field has been set.

### GetNextPageKey

`func (o *EventPropertiesList) GetNextPageKey() string`

GetNextPageKey returns the NextPageKey field if non-nil, zero value otherwise.

### GetNextPageKeyOk

`func (o *EventPropertiesList) GetNextPageKeyOk() (*string, bool)`

GetNextPageKeyOk returns a tuple with the NextPageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageKey

`func (o *EventPropertiesList) SetNextPageKey(v string)`

SetNextPageKey sets NextPageKey field to given value.

### HasNextPageKey

`func (o *EventPropertiesList) HasNextPageKey() bool`

HasNextPageKey returns a boolean if a field has been set.

### GetPageSize

`func (o *EventPropertiesList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *EventPropertiesList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *EventPropertiesList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *EventPropertiesList) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetTotalCount

`func (o *EventPropertiesList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EventPropertiesList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EventPropertiesList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


