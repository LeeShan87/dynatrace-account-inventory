# SubscriptionPeriodDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **string** | The subscription period start time in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**EndTime** | **string** | The subscription period end time in &#x60;2021-05-01T15:11:00Z&#x60; format. | 

## Methods

### NewSubscriptionPeriodDto

`func NewSubscriptionPeriodDto(startTime string, endTime string, ) *SubscriptionPeriodDto`

NewSubscriptionPeriodDto instantiates a new SubscriptionPeriodDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPeriodDtoWithDefaults

`func NewSubscriptionPeriodDtoWithDefaults() *SubscriptionPeriodDto`

NewSubscriptionPeriodDtoWithDefaults instantiates a new SubscriptionPeriodDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *SubscriptionPeriodDto) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SubscriptionPeriodDto) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SubscriptionPeriodDto) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SubscriptionPeriodDto) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SubscriptionPeriodDto) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SubscriptionPeriodDto) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


