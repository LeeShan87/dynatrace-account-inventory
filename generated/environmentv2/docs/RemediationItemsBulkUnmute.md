# RemediationItemsBulkUnmute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment about the un-muting reason. | [optional] 
**Reason** | **string** | The reason for un-muting the remediation items. | 
**RemediationItemIds** | **[]string** | The ids of the remediation items to be un-muted. | 

## Methods

### NewRemediationItemsBulkUnmute

`func NewRemediationItemsBulkUnmute(reason string, remediationItemIds []string, ) *RemediationItemsBulkUnmute`

NewRemediationItemsBulkUnmute instantiates a new RemediationItemsBulkUnmute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationItemsBulkUnmuteWithDefaults

`func NewRemediationItemsBulkUnmuteWithDefaults() *RemediationItemsBulkUnmute`

NewRemediationItemsBulkUnmuteWithDefaults instantiates a new RemediationItemsBulkUnmute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *RemediationItemsBulkUnmute) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *RemediationItemsBulkUnmute) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *RemediationItemsBulkUnmute) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *RemediationItemsBulkUnmute) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReason

`func (o *RemediationItemsBulkUnmute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RemediationItemsBulkUnmute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RemediationItemsBulkUnmute) SetReason(v string)`

SetReason sets Reason field to given value.


### GetRemediationItemIds

`func (o *RemediationItemsBulkUnmute) GetRemediationItemIds() []string`

GetRemediationItemIds returns the RemediationItemIds field if non-nil, zero value otherwise.

### GetRemediationItemIdsOk

`func (o *RemediationItemsBulkUnmute) GetRemediationItemIdsOk() (*[]string, bool)`

GetRemediationItemIdsOk returns a tuple with the RemediationItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationItemIds

`func (o *RemediationItemsBulkUnmute) SetRemediationItemIds(v []string)`

SetRemediationItemIds sets RemediationItemIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


