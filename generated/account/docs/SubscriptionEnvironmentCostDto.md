# SubscriptionEnvironmentCostDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentId** | **string** | The id of the environment | 
**Cost** | [**[]SubscriptionCapabilityCostDto**](SubscriptionCapabilityCostDto.md) | A list of subscription cost for the environment. | 

## Methods

### NewSubscriptionEnvironmentCostDto

`func NewSubscriptionEnvironmentCostDto(environmentId string, cost []SubscriptionCapabilityCostDto, ) *SubscriptionEnvironmentCostDto`

NewSubscriptionEnvironmentCostDto instantiates a new SubscriptionEnvironmentCostDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionEnvironmentCostDtoWithDefaults

`func NewSubscriptionEnvironmentCostDtoWithDefaults() *SubscriptionEnvironmentCostDto`

NewSubscriptionEnvironmentCostDtoWithDefaults instantiates a new SubscriptionEnvironmentCostDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentId

`func (o *SubscriptionEnvironmentCostDto) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *SubscriptionEnvironmentCostDto) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *SubscriptionEnvironmentCostDto) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetCost

`func (o *SubscriptionEnvironmentCostDto) GetCost() []SubscriptionCapabilityCostDto`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SubscriptionEnvironmentCostDto) GetCostOk() (*[]SubscriptionCapabilityCostDto, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SubscriptionEnvironmentCostDto) SetCost(v []SubscriptionCapabilityCostDto)`

SetCost sets Cost field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


