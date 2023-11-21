# RemediationItemsBulkMute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment about the muting reason. | [optional] 
**Reason** | **string** | The reason for muting the remediation items. | 
**RemediationItemIds** | **[]string** | The ids of the remediation items to be muted. | 

## Methods

### NewRemediationItemsBulkMute

`func NewRemediationItemsBulkMute(reason string, remediationItemIds []string, ) *RemediationItemsBulkMute`

NewRemediationItemsBulkMute instantiates a new RemediationItemsBulkMute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationItemsBulkMuteWithDefaults

`func NewRemediationItemsBulkMuteWithDefaults() *RemediationItemsBulkMute`

NewRemediationItemsBulkMuteWithDefaults instantiates a new RemediationItemsBulkMute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *RemediationItemsBulkMute) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RemediationItemsBulkMute) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RemediationItemsBulkMute) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RemediationItemsBulkMute) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReason

`func (o *RemediationItemsBulkMute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RemediationItemsBulkMute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RemediationItemsBulkMute) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRemediationItemIds

`func (o *RemediationItemsBulkMute) GetRemediationItemIds() []string`

GetRemediationItemIds returns the RemediationItemIds field if non-nil, zero value otherwise.

### GetRemediationItemIdsOk

`func (o *RemediationItemsBulkMute) GetRemediationItemIdsOk() (*[]string, bool)`

GetRemediationItemIdsOk returns a tuple with the RemediationItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationItemIds

`func (o *RemediationItemsBulkMute) SetRemediationItemIds(v []string)`

SetRemediationItemIds sets RemediationItemIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


