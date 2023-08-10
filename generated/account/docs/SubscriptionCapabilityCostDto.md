# SubscriptionCapabilityCostDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartTime** | **string** | The start time for the capability cost in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**EndTime** | **string** | The end time for the capability cost in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**Value** | **float32** | The total cost for all the capabilities combined | 
**CurrencyCode** | **string** | The currency code for the cost | 
**CapabilityKey** | **string** | The capability key | 
**CapabilityName** | **string** | The name of the capability | 

## Methods

### NewSubscriptionCapabilityCostDto

`func NewSubscriptionCapabilityCostDto(startTime string, endTime string, value float32, currencyCode string, capabilityKey string, capabilityName string, ) *SubscriptionCapabilityCostDto`

NewSubscriptionCapabilityCostDto instantiates a new SubscriptionCapabilityCostDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCapabilityCostDtoWithDefaults

`func NewSubscriptionCapabilityCostDtoWithDefaults() *SubscriptionCapabilityCostDto`

NewSubscriptionCapabilityCostDtoWithDefaults instantiates a new SubscriptionCapabilityCostDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartTime

`func (o *SubscriptionCapabilityCostDto) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SubscriptionCapabilityCostDto) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SubscriptionCapabilityCostDto) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SubscriptionCapabilityCostDto) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SubscriptionCapabilityCostDto) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SubscriptionCapabilityCostDto) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetValue

`func (o *SubscriptionCapabilityCostDto) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SubscriptionCapabilityCostDto) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SubscriptionCapabilityCostDto) SetValue(v float32)`

SetValue sets Value field to given value.


### GetCurrencyCode

`func (o *SubscriptionCapabilityCostDto) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SubscriptionCapabilityCostDto) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SubscriptionCapabilityCostDto) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetCapabilityKey

`func (o *SubscriptionCapabilityCostDto) GetCapabilityKey() string`

GetCapabilityKey returns the CapabilityKey field if non-nil, zero value otherwise.

### GetCapabilityKeyOk

`func (o *SubscriptionCapabilityCostDto) GetCapabilityKeyOk() (*string, bool)`

GetCapabilityKeyOk returns a tuple with the CapabilityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityKey

`func (o *SubscriptionCapabilityCostDto) SetCapabilityKey(v string)`

SetCapabilityKey sets CapabilityKey field to given value.


### GetCapabilityName

`func (o *SubscriptionCapabilityCostDto) GetCapabilityName() string`

GetCapabilityName returns the CapabilityName field if non-nil, zero value otherwise.

### GetCapabilityNameOk

`func (o *SubscriptionCapabilityCostDto) GetCapabilityNameOk() (*string, bool)`

GetCapabilityNameOk returns a tuple with the CapabilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityName

`func (o *SubscriptionCapabilityCostDto) SetCapabilityName(v string)`

SetCapabilityName sets CapabilityName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


