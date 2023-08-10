# SubscriptionEnvironmentUsageDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | The id of the environment | 
**Usage** | [**[]SubscriptionUsageDto**](SubscriptionUsageDto.md) | A list of subscription usage for the environment. | 

## Methods

### NewSubscriptionEnvironmentUsageDto

`func NewSubscriptionEnvironmentUsageDto(environmentId string, usage []SubscriptionUsageDto, ) *SubscriptionEnvironmentUsageDto`

NewSubscriptionEnvironmentUsageDto instantiates a new SubscriptionEnvironmentUsageDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionEnvironmentUsageDtoWithDefaults

`func NewSubscriptionEnvironmentUsageDtoWithDefaults() *SubscriptionEnvironmentUsageDto`

NewSubscriptionEnvironmentUsageDtoWithDefaults instantiates a new SubscriptionEnvironmentUsageDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *SubscriptionEnvironmentUsageDto) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SubscriptionEnvironmentUsageDto) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SubscriptionEnvironmentUsageDto) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetUsage

`func (o *SubscriptionEnvironmentUsageDto) GetUsage() []SubscriptionUsageDto`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SubscriptionEnvironmentUsageDto) GetUsageOk() (*[]SubscriptionUsageDto, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SubscriptionEnvironmentUsageDto) SetUsage(v []SubscriptionUsageDto)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


