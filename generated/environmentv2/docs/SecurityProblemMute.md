# SecurityProblemMute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment about the muting reason. | [optional] 
**Reason** | **string** | The reason for muting a security problem. | 

## Methods

### NewSecurityProblemMute

`func NewSecurityProblemMute(reason string, ) *SecurityProblemMute`

NewSecurityProblemMute instantiates a new SecurityProblemMute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemMuteWithDefaults

`func NewSecurityProblemMuteWithDefaults() *SecurityProblemMute`

NewSecurityProblemMuteWithDefaults instantiates a new SecurityProblemMute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SecurityProblemMute) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SecurityProblemMute) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SecurityProblemMute) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SecurityProblemMute) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReason

`func (o *SecurityProblemMute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SecurityProblemMute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SecurityProblemMute) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


