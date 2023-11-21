# RemediationItemsBulkMuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | [**[]RemediationItemMutingSummary**](RemediationItemMutingSummary.md) | The summary of which remediation items were muted and which already were muted previously. | 

## Methods

### NewRemediationItemsBulkMuteResponse

`func NewRemediationItemsBulkMuteResponse(summary []RemediationItemMutingSummary, ) *RemediationItemsBulkMuteResponse`

NewRemediationItemsBulkMuteResponse instantiates a new RemediationItemsBulkMuteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationItemsBulkMuteResponseWithDefaults

`func NewRemediationItemsBulkMuteResponseWithDefaults() *RemediationItemsBulkMuteResponse`

NewRemediationItemsBulkMuteResponseWithDefaults instantiates a new RemediationItemsBulkMuteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *RemediationItemsBulkMuteResponse) GetSummary() []RemediationItemMutingSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RemediationItemsBulkMuteResponse) GetSummaryOk() (*[]RemediationItemMutingSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RemediationItemsBulkMuteResponse) SetSummary(v []RemediationItemMutingSummary)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


