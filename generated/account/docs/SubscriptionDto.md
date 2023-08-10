# SubscriptionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The UUID of the Dynatrace Platform Subscription. | 
**Type** | **string** | The type of the Dynatrace Platform Subscription. | 
**Name** | **string** | The name of the Dynatrace Platform Subscription. | 
**Status** | **string** | The status of the Dynatrace Platform Subscription. | 
**StartTime** | **string** | The start date of the subscription in &#x60;2021-05-01&#x60; format. | 
**EndTime** | **string** | The end date of the subscription in &#x60;2021-05-01&#x60; format. | 
**Account** | [**SubscriptionAccountDto**](SubscriptionAccountDto.md) |  | 
**Budget** | [**SubscriptionBudgetDto**](SubscriptionBudgetDto.md) |  | 
**CurrentPeriod** | [**SubscriptionCurrentPeriodDto**](SubscriptionCurrentPeriodDto.md) |  | 
**Periods** | [**[]SubscriptionPeriodDto**](SubscriptionPeriodDto.md) | A list of period data of the subscription. | 
**Capabilities** | [**[]SubscriptionCapabilityDto**](SubscriptionCapabilityDto.md) | A list of subscription capabilities. | 

## Methods

### NewSubscriptionDto

`func NewSubscriptionDto(uuid string, type_ string, name string, status string, startTime string, endTime string, account SubscriptionAccountDto, budget SubscriptionBudgetDto, currentPeriod SubscriptionCurrentPeriodDto, periods []SubscriptionPeriodDto, capabilities []SubscriptionCapabilityDto, ) *SubscriptionDto`

NewSubscriptionDto instantiates a new SubscriptionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionDtoWithDefaults

`func NewSubscriptionDtoWithDefaults() *SubscriptionDto`

NewSubscriptionDtoWithDefaults instantiates a new SubscriptionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SubscriptionDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SubscriptionDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SubscriptionDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *SubscriptionDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionDto) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *SubscriptionDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionDto) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *SubscriptionDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartTime

`func (o *SubscriptionDto) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SubscriptionDto) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SubscriptionDto) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SubscriptionDto) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SubscriptionDto) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SubscriptionDto) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetAccount

`func (o *SubscriptionDto) GetAccount() SubscriptionAccountDto`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SubscriptionDto) GetAccountOk() (*SubscriptionAccountDto, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SubscriptionDto) SetAccount(v SubscriptionAccountDto)`

SetAccount sets Account field to given value.


### GetBudget

`func (o *SubscriptionDto) GetBudget() SubscriptionBudgetDto`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *SubscriptionDto) GetBudgetOk() (*SubscriptionBudgetDto, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *SubscriptionDto) SetBudget(v SubscriptionBudgetDto)`

SetBudget sets Budget field to given value.


### GetCurrentPeriod

`func (o *SubscriptionDto) GetCurrentPeriod() SubscriptionCurrentPeriodDto`

GetCurrentPeriod returns the CurrentPeriod field if non-nil, zero value otherwise.

### GetCurrentPeriodOk

`func (o *SubscriptionDto) GetCurrentPeriodOk() (*SubscriptionCurrentPeriodDto, bool)`

GetCurrentPeriodOk returns a tuple with the CurrentPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriod

`func (o *SubscriptionDto) SetCurrentPeriod(v SubscriptionCurrentPeriodDto)`

SetCurrentPeriod sets CurrentPeriod field to given value.


### GetPeriods

`func (o *SubscriptionDto) GetPeriods() []SubscriptionPeriodDto`

GetPeriods returns the Periods field if non-nil, zero value otherwise.

### GetPeriodsOk

`func (o *SubscriptionDto) GetPeriodsOk() (*[]SubscriptionPeriodDto, bool)`

GetPeriodsOk returns a tuple with the Periods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriods

`func (o *SubscriptionDto) SetPeriods(v []SubscriptionPeriodDto)`

SetPeriods sets Periods field to given value.


### GetCapabilities

`func (o *SubscriptionDto) GetCapabilities() []SubscriptionCapabilityDto`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *SubscriptionDto) GetCapabilitiesOk() (*[]SubscriptionCapabilityDto, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *SubscriptionDto) SetCapabilities(v []SubscriptionCapabilityDto)`

SetCapabilities sets Capabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


