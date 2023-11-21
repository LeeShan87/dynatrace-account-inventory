# SecurityProblemBulkMutingSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MuteStateChangeTriggered** | **bool** | Whether a mute state change for the given security problem was triggered by this request. | [readonly] 
**Reason** | Pointer to **string** | Contains a reason, in case the requested operation was not executed. | [optional] [readonly] 
**SecurityProblemId** | **string** | The id of the security problem that was (un-)muted. | [readonly] 

## Methods

### NewSecurityProblemBulkMutingSummary

`func NewSecurityProblemBulkMutingSummary(muteStateChangeTriggered bool, securityProblemId string, ) *SecurityProblemBulkMutingSummary`

NewSecurityProblemBulkMutingSummary instantiates a new SecurityProblemBulkMutingSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemBulkMutingSummaryWithDefaults

`func NewSecurityProblemBulkMutingSummaryWithDefaults() *SecurityProblemBulkMutingSummary`

NewSecurityProblemBulkMutingSummaryWithDefaults instantiates a new SecurityProblemBulkMutingSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMuteStateChangeTriggered

`func (o *SecurityProblemBulkMutingSummary) GetMuteStateChangeTriggered() bool`

GetMuteStateChangeTriggered returns the MuteStateChangeTriggered field if non-nil, zero value otherwise.

### GetMuteStateChangeTriggeredOk

`func (o *SecurityProblemBulkMutingSummary) GetMuteStateChangeTriggeredOk() (*bool, bool)`

GetMuteStateChangeTriggeredOk returns a tuple with the MuteStateChangeTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteStateChangeTriggered

`func (o *SecurityProblemBulkMutingSummary) SetMuteStateChangeTriggered(v bool)`

SetMuteStateChangeTriggered sets MuteStateChangeTriggered field to given value.


### GetReason

`func (o *SecurityProblemBulkMutingSummary) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SecurityProblemBulkMutingSummary) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SecurityProblemBulkMutingSummary) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SecurityProblemBulkMutingSummary) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSecurityProblemId

`func (o *SecurityProblemBulkMutingSummary) GetSecurityProblemId() string`

GetSecurityProblemId returns the SecurityProblemId field if non-nil, zero value otherwise.

### GetSecurityProblemIdOk

`func (o *SecurityProblemBulkMutingSummary) GetSecurityProblemIdOk() (*string, bool)`

GetSecurityProblemIdOk returns a tuple with the SecurityProblemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProblemId

`func (o *SecurityProblemBulkMutingSummary) SetSecurityProblemId(v string)`

SetSecurityProblemId sets SecurityProblemId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


