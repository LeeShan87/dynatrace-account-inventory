# SubscriptionCostListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SubscriptionCostDto**](SubscriptionCostDto.md) | Subscription cost data | 
**LastModifiedTime** | **string** | The time the subscription data was last modified in &#x60;2021-05-01T15:11:00Z&#x60; format. | 

## Methods

### NewSubscriptionCostListDto

`func NewSubscriptionCostListDto(data []SubscriptionCostDto, lastModifiedTime string, ) *SubscriptionCostListDto`

NewSubscriptionCostListDto instantiates a new SubscriptionCostListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCostListDtoWithDefaults

`func NewSubscriptionCostListDtoWithDefaults() *SubscriptionCostListDto`

NewSubscriptionCostListDtoWithDefaults instantiates a new SubscriptionCostListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionCostListDto) GetData() []SubscriptionCostDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionCostListDto) GetDataOk() (*[]SubscriptionCostDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionCostListDto) SetData(v []SubscriptionCostDto)`

SetData sets Data field to given value.


### GetLastModifiedTime

`func (o *SubscriptionCostListDto) GetLastModifiedTime() string`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *SubscriptionCostListDto) GetLastModifiedTimeOk() (*string, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *SubscriptionCostListDto) SetLastModifiedTime(v string)`

SetLastModifiedTime sets LastModifiedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


