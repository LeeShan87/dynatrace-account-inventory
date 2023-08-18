# SecurityProblemsBulkUnmute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | A comment about the un-muting reason. | [optional] 
**Reason** | **string** | The reason for un-muting the security problems. | 
**SecurityProblemIds** | **[]string** | The ids of the security problems to be un-muted. | 

## Methods

### NewSecurityProblemsBulkUnmute

`func NewSecurityProblemsBulkUnmute(reason string, securityProblemIds []string, ) *SecurityProblemsBulkUnmute`

NewSecurityProblemsBulkUnmute instantiates a new SecurityProblemsBulkUnmute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemsBulkUnmuteWithDefaults

`func NewSecurityProblemsBulkUnmuteWithDefaults() *SecurityProblemsBulkUnmute`

NewSecurityProblemsBulkUnmuteWithDefaults instantiates a new SecurityProblemsBulkUnmute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *SecurityProblemsBulkUnmute) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SecurityProblemsBulkUnmute) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SecurityProblemsBulkUnmute) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SecurityProblemsBulkUnmute) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReason

`func (o *SecurityProblemsBulkUnmute) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SecurityProblemsBulkUnmute) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SecurityProblemsBulkUnmute) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSecurityProblemIds

`func (o *SecurityProblemsBulkUnmute) GetSecurityProblemIds() []string`

GetSecurityProblemIds returns the SecurityProblemIds field if non-nil, zero value otherwise.

### GetSecurityProblemIdsOk

`func (o *SecurityProblemsBulkUnmute) GetSecurityProblemIdsOk() (*[]string, bool)`

GetSecurityProblemIdsOk returns a tuple with the SecurityProblemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProblemIds

`func (o *SecurityProblemsBulkUnmute) SetSecurityProblemIds(v []string)`

SetSecurityProblemIds sets SecurityProblemIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


