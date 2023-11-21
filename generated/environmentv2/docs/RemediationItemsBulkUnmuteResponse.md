# RemediationItemsBulkUnmuteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | [**[]RemediationItemMutingSummary**](RemediationItemMutingSummary.md) | The summary of which remediation items were un-muted and which already were un-muted previously. | 

## Methods

### NewRemediationItemsBulkUnmuteResponse

`func NewRemediationItemsBulkUnmuteResponse(summary []RemediationItemMutingSummary, ) *RemediationItemsBulkUnmuteResponse`

NewRemediationItemsBulkUnmuteResponse instantiates a new RemediationItemsBulkUnmuteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationItemsBulkUnmuteResponseWithDefaults

`func NewRemediationItemsBulkUnmuteResponseWithDefaults() *RemediationItemsBulkUnmuteResponse`

NewRemediationItemsBulkUnmuteResponseWithDefaults instantiates a new RemediationItemsBulkUnmuteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *RemediationItemsBulkUnmuteResponse) GetSummary() []RemediationItemMutingSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *RemediationItemsBulkUnmuteResponse) GetSummaryOk() (*[]RemediationItemMutingSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *RemediationItemsBulkUnmuteResponse) SetSummary(v []RemediationItemMutingSummary)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


