# SecurityProblemsBulkMuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | [**[]SecurityProblemBulkMutingSummary**](SecurityProblemBulkMutingSummary.md) | The summary of which security problems were muted and which already were muted previously. | [readonly] 

## Methods

### NewSecurityProblemsBulkMuteResponse

`func NewSecurityProblemsBulkMuteResponse(summary []SecurityProblemBulkMutingSummary, ) *SecurityProblemsBulkMuteResponse`

NewSecurityProblemsBulkMuteResponse instantiates a new SecurityProblemsBulkMuteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemsBulkMuteResponseWithDefaults

`func NewSecurityProblemsBulkMuteResponseWithDefaults() *SecurityProblemsBulkMuteResponse`

NewSecurityProblemsBulkMuteResponseWithDefaults instantiates a new SecurityProblemsBulkMuteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *SecurityProblemsBulkMuteResponse) GetSummary() []SecurityProblemBulkMutingSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SecurityProblemsBulkMuteResponse) GetSummaryOk() (*[]SecurityProblemBulkMutingSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SecurityProblemsBulkMuteResponse) SetSummary(v []SecurityProblemBulkMutingSummary)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


