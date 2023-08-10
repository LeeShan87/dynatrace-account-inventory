# SubscriptionCostDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **string** | The start time for the capability cost in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**EndTime** | **string** | The end time for the capability cost in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**Value** | **float32** | The total cost for all the capabilities combined | 
**CurrencyCode** | **string** | The currency code for the cost | 

## Methods

### NewSubscriptionCostDto

`func NewSubscriptionCostDto(startTime string, endTime string, value float32, currencyCode string, ) *SubscriptionCostDto`

NewSubscriptionCostDto instantiates a new SubscriptionCostDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCostDtoWithDefaults

`func NewSubscriptionCostDtoWithDefaults() *SubscriptionCostDto`

NewSubscriptionCostDtoWithDefaults instantiates a new SubscriptionCostDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *SubscriptionCostDto) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SubscriptionCostDto) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SubscriptionCostDto) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SubscriptionCostDto) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SubscriptionCostDto) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SubscriptionCostDto) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetValue

`func (o *SubscriptionCostDto) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SubscriptionCostDto) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SubscriptionCostDto) SetValue(v float32)`

SetValue sets Value field to given value.


### GetCurrencyCode

`func (o *SubscriptionCostDto) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SubscriptionCostDto) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SubscriptionCostDto) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


