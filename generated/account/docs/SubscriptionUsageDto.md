# SubscriptionUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CapabilityKey** | **string** | The capability key | 
**CapabilityName** | **string** | The name of the capability | 
**StartTime** | **string** | The start time of the capability usage in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**EndTime** | **string** | The end time of the capability usage in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**Value** | **float32** | The usage for the capability | 
**UnitMeasure** | **string** | The unit of measure for the capability usage | 

## Methods

### NewSubscriptionUsageDto

`func NewSubscriptionUsageDto(capabilityKey string, capabilityName string, startTime string, endTime string, value float32, unitMeasure string, ) *SubscriptionUsageDto`

NewSubscriptionUsageDto instantiates a new SubscriptionUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUsageDtoWithDefaults

`func NewSubscriptionUsageDtoWithDefaults() *SubscriptionUsageDto`

NewSubscriptionUsageDtoWithDefaults instantiates a new SubscriptionUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilityKey

`func (o *SubscriptionUsageDto) GetCapabilityKey() string`

GetCapabilityKey returns the CapabilityKey field if non-nil, zero value otherwise.

### GetCapabilityKeyOk

`func (o *SubscriptionUsageDto) GetCapabilityKeyOk() (*string, bool)`

GetCapabilityKeyOk returns a tuple with the CapabilityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityKey

`func (o *SubscriptionUsageDto) SetCapabilityKey(v string)`

SetCapabilityKey sets CapabilityKey field to given value.


### GetCapabilityName

`func (o *SubscriptionUsageDto) GetCapabilityName() string`

GetCapabilityName returns the CapabilityName field if non-nil, zero value otherwise.

### GetCapabilityNameOk

`func (o *SubscriptionUsageDto) GetCapabilityNameOk() (*string, bool)`

GetCapabilityNameOk returns a tuple with the CapabilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilityName

`func (o *SubscriptionUsageDto) SetCapabilityName(v string)`

SetCapabilityName sets CapabilityName field to given value.


### GetStartTime

`func (o *SubscriptionUsageDto) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SubscriptionUsageDto) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SubscriptionUsageDto) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SubscriptionUsageDto) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SubscriptionUsageDto) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SubscriptionUsageDto) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.


### GetValue

`func (o *SubscriptionUsageDto) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SubscriptionUsageDto) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SubscriptionUsageDto) SetValue(v float32)`

SetValue sets Value field to given value.


### GetUnitMeasure

`func (o *SubscriptionUsageDto) GetUnitMeasure() string`

GetUnitMeasure returns the UnitMeasure field if non-nil, zero value otherwise.

### GetUnitMeasureOk

`func (o *SubscriptionUsageDto) GetUnitMeasureOk() (*string, bool)`

GetUnitMeasureOk returns a tuple with the UnitMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitMeasure

`func (o *SubscriptionUsageDto) SetUnitMeasure(v string)`

SetUnitMeasure sets UnitMeasure field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


