# SubscriptionListDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]SubscriptionSummaryDto**](SubscriptionSummaryDto.md) | A list of subscriptions of the account. | 

## Methods

### NewSubscriptionListDto

`func NewSubscriptionListDto(data []SubscriptionSummaryDto, ) *SubscriptionListDto`

NewSubscriptionListDto instantiates a new SubscriptionListDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionListDtoWithDefaults

`func NewSubscriptionListDtoWithDefaults() *SubscriptionListDto`

NewSubscriptionListDtoWithDefaults instantiates a new SubscriptionListDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SubscriptionListDto) GetData() []SubscriptionSummaryDto`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscriptionListDto) GetDataOk() (*[]SubscriptionSummaryDto, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscriptionListDto) SetData(v []SubscriptionSummaryDto)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


