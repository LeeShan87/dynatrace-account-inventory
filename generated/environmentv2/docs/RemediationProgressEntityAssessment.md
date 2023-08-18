# RemediationProgressEntityAssessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VulnerableFunctionUsage** | Pointer to **string** | The usage of vulnerable functions | [optional] [readonly] 
**VulnerableFunctionsInUse** | Pointer to [**[]VulnerableFunction**](VulnerableFunction.md) | A list of vulnerable functions that are in use. | [optional] [readonly] 
**VulnerableFunctionsNotAvailable** | Pointer to [**[]VulnerableFunction**](VulnerableFunction.md) | A list of vulnerable functions that are not available. | [optional] [readonly] 
**VulnerableFunctionsNotInUse** | Pointer to [**[]VulnerableFunction**](VulnerableFunction.md) | A list of vulnerable functions that are not in use. | [optional] [readonly] 

## Methods

### NewRemediationProgressEntityAssessment

`func NewRemediationProgressEntityAssessment() *RemediationProgressEntityAssessment`

NewRemediationProgressEntityAssessment instantiates a new RemediationProgressEntityAssessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationProgressEntityAssessmentWithDefaults

`func NewRemediationProgressEntityAssessmentWithDefaults() *RemediationProgressEntityAssessment`

NewRemediationProgressEntityAssessmentWithDefaults instantiates a new RemediationProgressEntityAssessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVulnerableFunctionUsage

`func (o *RemediationProgressEntityAssessment) GetVulnerableFunctionUsage() string`

GetVulnerableFunctionUsage returns the VulnerableFunctionUsage field if non-nil, zero value otherwise.

### GetVulnerableFunctionUsageOk

`func (o *RemediationProgressEntityAssessment) GetVulnerableFunctionUsageOk() (*string, bool)`

GetVulnerableFunctionUsageOk returns a tuple with the VulnerableFunctionUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctionUsage

`func (o *RemediationProgressEntityAssessment) SetVulnerableFunctionUsage(v string)`

SetVulnerableFunctionUsage sets VulnerableFunctionUsage field to given value.

### HasVulnerableFunctionUsage

`func (o *RemediationProgressEntityAssessment) HasVulnerableFunctionUsage() bool`

HasVulnerableFunctionUsage returns a boolean if a field has been set.

### GetVulnerableFunctionsInUse

`func (o *RemediationProgressEntityAssessment) GetVulnerableFunctionsInUse() []VulnerableFunction`

GetVulnerableFunctionsInUse returns the VulnerableFunctionsInUse field if non-nil, zero value otherwise.

### GetVulnerableFunctionsInUseOk

`func (o *RemediationProgressEntityAssessment) GetVulnerableFunctionsInUseOk() (*[]VulnerableFunction, bool)`

GetVulnerableFunctionsInUseOk returns a tuple with the VulnerableFunctionsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctionsInUse

`func (o *RemediationProgressEntityAssessment) SetVulnerableFunctionsInUse(v []VulnerableFunction)`

SetVulnerableFunctionsInUse sets VulnerableFunctionsInUse field to given value.

### HasVulnerableFunctionsInUse

`func (o *RemediationProgressEntityAssessment) HasVulnerableFunctionsInUse() bool`

HasVulnerableFunctionsInUse returns a boolean if a field has been set.

### GetVulnerableFunctionsNotAvailable

`func (o *RemediationProgressEntityAssessment) GetVulnerableFunctionsNotAvailable() []VulnerableFunction`

GetVulnerableFunctionsNotAvailable returns the VulnerableFunctionsNotAvailable field if non-nil, zero value otherwise.

### GetVulnerableFunctionsNotAvailableOk

`func (o *RemediationProgressEntityAssessment) GetVulnerableFunctionsNotAvailableOk() (*[]VulnerableFunction, bool)`

GetVulnerableFunctionsNotAvailableOk returns a tuple with the VulnerableFunctionsNotAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctionsNotAvailable

`func (o *RemediationProgressEntityAssessment) SetVulnerableFunctionsNotAvailable(v []VulnerableFunction)`

SetVulnerableFunctionsNotAvailable sets VulnerableFunctionsNotAvailable field to given value.

### HasVulnerableFunctionsNotAvailable

`func (o *RemediationProgressEntityAssessment) HasVulnerableFunctionsNotAvailable() bool`

HasVulnerableFunctionsNotAvailable returns a boolean if a field has been set.

### GetVulnerableFunctionsNotInUse

`func (o *RemediationProgressEntityAssessment) GetVulnerableFunctionsNotInUse() []VulnerableFunction`

GetVulnerableFunctionsNotInUse returns the VulnerableFunctionsNotInUse field if non-nil, zero value otherwise.

### GetVulnerableFunctionsNotInUseOk

`func (o *RemediationProgressEntityAssessment) GetVulnerableFunctionsNotInUseOk() (*[]VulnerableFunction, bool)`

GetVulnerableFunctionsNotInUseOk returns a tuple with the VulnerableFunctionsNotInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctionsNotInUse

`func (o *RemediationProgressEntityAssessment) SetVulnerableFunctionsNotInUse(v []VulnerableFunction)`

SetVulnerableFunctionsNotInUse sets VulnerableFunctionsNotInUse field to given value.

### HasVulnerableFunctionsNotInUse

`func (o *RemediationProgressEntityAssessment) HasVulnerableFunctionsNotInUse() bool`

HasVulnerableFunctionsNotInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


