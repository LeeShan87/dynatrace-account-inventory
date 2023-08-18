# RemediationItemMuteStateChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | **string** | A comment about the mute state change reason. | 
**Muted** | **bool** | The desired mute state of the remediation item. | 
**Reason** | **string** | The reason for the mute state change. | 

## Methods

### NewRemediationItemMuteStateChange

`func NewRemediationItemMuteStateChange(comment string, muted bool, reason string, ) *RemediationItemMuteStateChange`

NewRemediationItemMuteStateChange instantiates a new RemediationItemMuteStateChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationItemMuteStateChangeWithDefaults

`func NewRemediationItemMuteStateChangeWithDefaults() *RemediationItemMuteStateChange`

NewRemediationItemMuteStateChangeWithDefaults instantiates a new RemediationItemMuteStateChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *RemediationItemMuteStateChange) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RemediationItemMuteStateChange) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RemediationItemMuteStateChange) SetComment(v string)`

SetComment sets Comment field to given value.


### GetMuted

`func (o *RemediationItemMuteStateChange) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *RemediationItemMuteStateChange) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *RemediationItemMuteStateChange) SetMuted(v bool)`

SetMuted sets Muted field to given value.


### GetReason

`func (o *RemediationItemMuteStateChange) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RemediationItemMuteStateChange) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RemediationItemMuteStateChange) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


