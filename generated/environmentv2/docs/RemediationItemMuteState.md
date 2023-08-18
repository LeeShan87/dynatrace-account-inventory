# RemediationItemMuteState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A short comment about the most recent mute state change. | [optional] [readonly] 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp (UTC milliseconds) of the last update of the mute state. | [optional] [readonly] 
**Muted** | Pointer to **bool** | The remediation is (&#x60;true&#x60;) or is not (&#x60;false&#x60;) muted. | [optional] [readonly] 
**Reason** | Pointer to **string** | The reason for the most recent mute state change. | [optional] [readonly] 
**User** | Pointer to **string** | The user who last changed the mute state. | [optional] [readonly] 

## Methods

### NewRemediationItemMuteState

`func NewRemediationItemMuteState() *RemediationItemMuteState`

NewRemediationItemMuteState instantiates a new RemediationItemMuteState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationItemMuteStateWithDefaults

`func NewRemediationItemMuteStateWithDefaults() *RemediationItemMuteState`

NewRemediationItemMuteStateWithDefaults instantiates a new RemediationItemMuteState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *RemediationItemMuteState) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RemediationItemMuteState) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RemediationItemMuteState) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RemediationItemMuteState) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *RemediationItemMuteState) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *RemediationItemMuteState) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *RemediationItemMuteState) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *RemediationItemMuteState) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetMuted

`func (o *RemediationItemMuteState) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *RemediationItemMuteState) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *RemediationItemMuteState) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *RemediationItemMuteState) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetReason

`func (o *RemediationItemMuteState) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RemediationItemMuteState) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RemediationItemMuteState) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RemediationItemMuteState) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetUser

`func (o *RemediationItemMuteState) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RemediationItemMuteState) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RemediationItemMuteState) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RemediationItemMuteState) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


