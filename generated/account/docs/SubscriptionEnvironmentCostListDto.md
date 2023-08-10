# SubscriptionEnvironmentCostListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SubscriptionEnvironmentCostDto**](SubscriptionEnvironmentCostDto.md) | Subscription cost data | 
**LastModifiedTime** | **string** | The time the subscription data was last modified in &#x60;2021-05-01T15:11:00Z&#x60; format. | 

## Methods

### NewSubscriptionEnvironmentCostListDto

`func NewSubscriptionEnvironmentCostListDto(data []SubscriptionEnvironmentCostDto, lastModifiedTime string, ) *SubscriptionEnvironmentCostListDto`

NewSubscriptionEnvironmentCostListDto instantiates a new SubscriptionEnvironmentCostListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionEnvironmentCostListDtoWithDefaults

`func NewSubscriptionEnvironmentCostListDtoWithDefaults() *SubscriptionEnvironmentCostListDto`

NewSubscriptionEnvironmentCostListDtoWithDefaults instantiates a new SubscriptionEnvironmentCostListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionEnvironmentCostListDto) GetData() []SubscriptionEnvironmentCostDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionEnvironmentCostListDto) GetDataOk() (*[]SubscriptionEnvironmentCostDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionEnvironmentCostListDto) SetData(v []SubscriptionEnvironmentCostDto)`

SetData sets Data field to given value.


### GetLastModifiedTime

`func (o *SubscriptionEnvironmentCostListDto) GetLastModifiedTime() string`

GetLastModifiedTime returns the LastModifiedTime field if non-nil, zero value otherwise.

### GetLastModifiedTimeOk

`func (o *SubscriptionEnvironmentCostListDto) GetLastModifiedTimeOk() (*string, bool)`

GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedTime

`func (o *SubscriptionEnvironmentCostListDto) SetLastModifiedTime(v string)`

SetLastModifiedTime sets LastModifiedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


