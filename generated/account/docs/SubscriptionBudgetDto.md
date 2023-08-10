# SubscriptionBudgetDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **float32** | The total budget for the subscription | 
**Used** | **float32** | The total budget used for the subscription | 
**CurrencyCode** | **string** | The currency code for the subscription | 

## Methods

### NewSubscriptionBudgetDto

`func NewSubscriptionBudgetDto(total float32, used float32, currencyCode string, ) *SubscriptionBudgetDto`

NewSubscriptionBudgetDto instantiates a new SubscriptionBudgetDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionBudgetDtoWithDefaults

`func NewSubscriptionBudgetDtoWithDefaults() *SubscriptionBudgetDto`

NewSubscriptionBudgetDtoWithDefaults instantiates a new SubscriptionBudgetDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *SubscriptionBudgetDto) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SubscriptionBudgetDto) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SubscriptionBudgetDto) SetTotal(v float32)`

SetTotal sets Total field to given value.


### GetUsed

`func (o *SubscriptionBudgetDto) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *SubscriptionBudgetDto) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *SubscriptionBudgetDto) SetUsed(v float32)`

SetUsed sets Used field to given value.


### GetCurrencyCode

`func (o *SubscriptionBudgetDto) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *SubscriptionBudgetDto) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *SubscriptionBudgetDto) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


