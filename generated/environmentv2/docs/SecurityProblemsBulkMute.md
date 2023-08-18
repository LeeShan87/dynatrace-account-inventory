# SecurityProblemsBulkMute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment about the muting reason. | [optional] 
**Reason** | **string** | The reason for muting the security problems. | 
**SecurityProblemIds** | **[]string** | The ids of the security problems to be muted. | 

## Methods

### NewSecurityProblemsBulkMute

`func NewSecurityProblemsBulkMute(reason string, securityProblemIds []string, ) *SecurityProblemsBulkMute`

NewSecurityProblemsBulkMute instantiates a new SecurityProblemsBulkMute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemsBulkMuteWithDefaults

`func NewSecurityProblemsBulkMuteWithDefaults() *SecurityProblemsBulkMute`

NewSecurityProblemsBulkMuteWithDefaults instantiates a new SecurityProblemsBulkMute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SecurityProblemsBulkMute) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SecurityProblemsBulkMute) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SecurityProblemsBulkMute) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SecurityProblemsBulkMute) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReason

`func (o *SecurityProblemsBulkMute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SecurityProblemsBulkMute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SecurityProblemsBulkMute) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSecurityProblemIds

`func (o *SecurityProblemsBulkMute) GetSecurityProblemIds() []string`

GetSecurityProblemIds returns the SecurityProblemIds field if non-nil, zero value otherwise.

### GetSecurityProblemIdsOk

`func (o *SecurityProblemsBulkMute) GetSecurityProblemIdsOk() (*[]string, bool)`

GetSecurityProblemIdsOk returns a tuple with the SecurityProblemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProblemIds

`func (o *SecurityProblemsBulkMute) SetSecurityProblemIds(v []string)`

SetSecurityProblemIds sets SecurityProblemIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


