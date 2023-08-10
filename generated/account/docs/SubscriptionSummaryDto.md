# SubscriptionSummaryDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The UUID of the Dynatrace Platform Subscription. | 
**Type** | **string** | The type of the Dynatrace Platform Subscription. | 
**Name** | **string** | The name of the Dynatrace Platform Subscription. | 
**Status** | **string** | The status of the Dynatrace Platform Subscription. | 
**StartTime** | **string** | The start date of the subscription in &#x60;2021-05-01&#x60; format. | 
**EndTime** | **string** | The end date of the subscription in &#x60;2021-05-01&#x60; format. | 

## Methods

### NewSubscriptionSummaryDto

`func NewSubscriptionSummaryDto(uuid string, type_ string, name string, status string, startTime string, endTime string, ) *SubscriptionSummaryDto`

NewSubscriptionSummaryDto instantiates a new SubscriptionSummaryDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionSummaryDtoWithDefaults

`func NewSubscriptionSummaryDtoWithDefaults() *SubscriptionSummaryDto`

NewSubscriptionSummaryDtoWithDefaults instantiates a new SubscriptionSummaryDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *SubscriptionSummaryDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SubscriptionSummaryDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SubscriptionSummaryDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetType

`func (o *SubscriptionSummaryDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionSummaryDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionSummaryDto) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *SubscriptionSummaryDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionSummaryDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionSummaryDto) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *SubscriptionSummaryDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionSummaryDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionSummaryDto) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStartTime

`func (o *SubscriptionSummaryDto) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SubscriptionSummaryDto) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SubscriptionSummaryDto) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SubscriptionSummaryDto) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SubscriptionSummaryDto) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SubscriptionSummaryDto) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


