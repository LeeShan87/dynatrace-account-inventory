# SecurityProblemUnmute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment about the un-muting reason. | [optional] 
**Reason** | **string** | The reason for un-muting a security problem. | 

## Methods

### NewSecurityProblemUnmute

`func NewSecurityProblemUnmute(reason string, ) *SecurityProblemUnmute`

NewSecurityProblemUnmute instantiates a new SecurityProblemUnmute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemUnmuteWithDefaults

`func NewSecurityProblemUnmuteWithDefaults() *SecurityProblemUnmute`

NewSecurityProblemUnmuteWithDefaults instantiates a new SecurityProblemUnmute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SecurityProblemUnmute) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SecurityProblemUnmute) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SecurityProblemUnmute) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SecurityProblemUnmute) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReason

`func (o *SecurityProblemUnmute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SecurityProblemUnmute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SecurityProblemUnmute) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


