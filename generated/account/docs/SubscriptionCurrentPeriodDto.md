# SubscriptionCurrentPeriodDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **string** | The current period start date in &#x60;2021-05-01&#x60; format. | 
**EndTime** | **string** | The current period end date in &#x60;2021-05-01&#x60; format. | 
**DaysRemaining** | **float32** | Days remaining in the current period | 

## Methods

### NewSubscriptionCurrentPeriodDto

`func NewSubscriptionCurrentPeriodDto(startTime string, endTime string, daysRemaining float32, ) *SubscriptionCurrentPeriodDto`

NewSubscriptionCurrentPeriodDto instantiates a new SubscriptionCurrentPeriodDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCurrentPeriodDtoWithDefaults

`func NewSubscriptionCurrentPeriodDtoWithDefaults() *SubscriptionCurrentPeriodDto`

NewSubscriptionCurrentPeriodDtoWithDefaults instantiates a new SubscriptionCurrentPeriodDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *SubscriptionCurrentPeriodDto) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SubscriptionCurrentPeriodDto) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SubscriptionCurrentPeriodDto) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SubscriptionCurrentPeriodDto) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SubscriptionCurrentPeriodDto) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SubscriptionCurrentPeriodDto) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetDaysRemaining

`func (o *SubscriptionCurrentPeriodDto) GetDaysRemaining() float32`

GetDaysRemaining returns the DaysRemaining field if non-nil, zero value otherwise.

### GetDaysRemainingOk

`func (o *SubscriptionCurrentPeriodDto) GetDaysRemainingOk() (*float32, bool)`

GetDaysRemainingOk returns a tuple with the DaysRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysRemaining

`func (o *SubscriptionCurrentPeriodDto) SetDaysRemaining(v float32)`

SetDaysRemaining sets DaysRemaining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


