# RemediationItemMutingSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MuteStateChangeTriggered** | **bool** | Whether a mute state change for the given remediation item was triggered by this request. | [readonly] 
**Reason** | Pointer to **string** | Contains a reason, in case the requested operation was not executed. | [optional] [readonly] 
**RemediationItemId** | **string** | The id of the remediation item that will be (un-)muted. | [readonly] 

## Methods

### NewRemediationItemMutingSummary

`func NewRemediationItemMutingSummary(muteStateChangeTriggered bool, remediationItemId string, ) *RemediationItemMutingSummary`

NewRemediationItemMutingSummary instantiates a new RemediationItemMutingSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationItemMutingSummaryWithDefaults

`func NewRemediationItemMutingSummaryWithDefaults() *RemediationItemMutingSummary`

NewRemediationItemMutingSummaryWithDefaults instantiates a new RemediationItemMutingSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMuteStateChangeTriggered

`func (o *RemediationItemMutingSummary) GetMuteStateChangeTriggered() bool`

GetMuteStateChangeTriggered returns the MuteStateChangeTriggered field if non-nil, zero value otherwise.

### GetMuteStateChangeTriggeredOk

`func (o *RemediationItemMutingSummary) GetMuteStateChangeTriggeredOk() (*bool, bool)`

GetMuteStateChangeTriggeredOk returns a tuple with the MuteStateChangeTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteStateChangeTriggered

`func (o *RemediationItemMutingSummary) SetMuteStateChangeTriggered(v bool)`

SetMuteStateChangeTriggered sets MuteStateChangeTriggered field to given value.


### GetReason

`func (o *RemediationItemMutingSummary) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RemediationItemMutingSummary) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RemediationItemMutingSummary) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RemediationItemMutingSummary) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRemediationItemId

`func (o *RemediationItemMutingSummary) GetRemediationItemId() string`

GetRemediationItemId returns the RemediationItemId field if non-nil, zero value otherwise.

### GetRemediationItemIdOk

`func (o *RemediationItemMutingSummary) GetRemediationItemIdOk() (*string, bool)`

GetRemediationItemIdOk returns a tuple with the RemediationItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationItemId

`func (o *RemediationItemMutingSummary) SetRemediationItemId(v string)`

SetRemediationItemId sets RemediationItemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


