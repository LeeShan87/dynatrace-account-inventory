# RemediationAssessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataAssets** | Pointer to **string** | The reachability of related data assets by affected entities. | [optional] [readonly] 
**Exposure** | Pointer to **string** | The level of exposure of affected entities. | [optional] [readonly] 
**NumberOfDataAssets** | Pointer to **int64** | The number of related data assets. | [optional] [readonly] 
**VulnerableFunctionUsage** | Pointer to **string** | The usage of vulnerable functions | [optional] [readonly] 
**VulnerableFunctionsInUse** | Pointer to [**[]VulnerableFunction**](VulnerableFunction.md) | A list of vulnerable functions that are in use. | [optional] [readonly] 
**VulnerableFunctionsNotAvailable** | Pointer to [**[]VulnerableFunction**](VulnerableFunction.md) | A list of vulnerable functions that are not available. | [optional] [readonly] 
**VulnerableFunctionsNotInUse** | Pointer to [**[]VulnerableFunction**](VulnerableFunction.md) | A list of vulnerable functions that are not in use. | [optional] [readonly] 

## Methods

### NewRemediationAssessment

`func NewRemediationAssessment() *RemediationAssessment`

NewRemediationAssessment instantiates a new RemediationAssessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationAssessmentWithDefaults

`func NewRemediationAssessmentWithDefaults() *RemediationAssessment`

NewRemediationAssessmentWithDefaults instantiates a new RemediationAssessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataAssets

`func (o *RemediationAssessment) GetDataAssets() string`

GetDataAssets returns the DataAssets field if non-nil, zero value otherwise.

### GetDataAssetsOk

`func (o *RemediationAssessment) GetDataAssetsOk() (*string, bool)`

GetDataAssetsOk returns a tuple with the DataAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataAssets

`func (o *RemediationAssessment) SetDataAssets(v string)`

SetDataAssets sets DataAssets field to given value.

### HasDataAssets

`func (o *RemediationAssessment) HasDataAssets() bool`

HasDataAssets returns a boolean if a field has been set.

### GetExposure

`func (o *RemediationAssessment) GetExposure() string`

GetExposure returns the Exposure field if non-nil, zero value otherwise.

### GetExposureOk

`func (o *RemediationAssessment) GetExposureOk() (*string, bool)`

GetExposureOk returns a tuple with the Exposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposure

`func (o *RemediationAssessment) SetExposure(v string)`

SetExposure sets Exposure field to given value.

### HasExposure

`func (o *RemediationAssessment) HasExposure() bool`

HasExposure returns a boolean if a field has been set.

### GetNumberOfDataAssets

`func (o *RemediationAssessment) GetNumberOfDataAssets() int64`

GetNumberOfDataAssets returns the NumberOfDataAssets field if non-nil, zero value otherwise.

### GetNumberOfDataAssetsOk

`func (o *RemediationAssessment) GetNumberOfDataAssetsOk() (*int64, bool)`

GetNumberOfDataAssetsOk returns a tuple with the NumberOfDataAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDataAssets

`func (o *RemediationAssessment) SetNumberOfDataAssets(v int64)`

SetNumberOfDataAssets sets NumberOfDataAssets field to given value.

### HasNumberOfDataAssets

`func (o *RemediationAssessment) HasNumberOfDataAssets() bool`

HasNumberOfDataAssets returns a boolean if a field has been set.

### GetVulnerableFunctionUsage

`func (o *RemediationAssessment) GetVulnerableFunctionUsage() string`

GetVulnerableFunctionUsage returns the VulnerableFunctionUsage field if non-nil, zero value otherwise.

### GetVulnerableFunctionUsageOk

`func (o *RemediationAssessment) GetVulnerableFunctionUsageOk() (*string, bool)`

GetVulnerableFunctionUsageOk returns a tuple with the VulnerableFunctionUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctionUsage

`func (o *RemediationAssessment) SetVulnerableFunctionUsage(v string)`

SetVulnerableFunctionUsage sets VulnerableFunctionUsage field to given value.

### HasVulnerableFunctionUsage

`func (o *RemediationAssessment) HasVulnerableFunctionUsage() bool`

HasVulnerableFunctionUsage returns a boolean if a field has been set.

### GetVulnerableFunctionsInUse

`func (o *RemediationAssessment) GetVulnerableFunctionsInUse() []VulnerableFunction`

GetVulnerableFunctionsInUse returns the VulnerableFunctionsInUse field if non-nil, zero value otherwise.

### GetVulnerableFunctionsInUseOk

`func (o *RemediationAssessment) GetVulnerableFunctionsInUseOk() (*[]VulnerableFunction, bool)`

GetVulnerableFunctionsInUseOk returns a tuple with the VulnerableFunctionsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctionsInUse

`func (o *RemediationAssessment) SetVulnerableFunctionsInUse(v []VulnerableFunction)`

SetVulnerableFunctionsInUse sets VulnerableFunctionsInUse field to given value.

### HasVulnerableFunctionsInUse

`func (o *RemediationAssessment) HasVulnerableFunctionsInUse() bool`

HasVulnerableFunctionsInUse returns a boolean if a field has been set.

### GetVulnerableFunctionsNotAvailable

`func (o *RemediationAssessment) GetVulnerableFunctionsNotAvailable() []VulnerableFunction`

GetVulnerableFunctionsNotAvailable returns the VulnerableFunctionsNotAvailable field if non-nil, zero value otherwise.

### GetVulnerableFunctionsNotAvailableOk

`func (o *RemediationAssessment) GetVulnerableFunctionsNotAvailableOk() (*[]VulnerableFunction, bool)`

GetVulnerableFunctionsNotAvailableOk returns a tuple with the VulnerableFunctionsNotAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctionsNotAvailable

`func (o *RemediationAssessment) SetVulnerableFunctionsNotAvailable(v []VulnerableFunction)`

SetVulnerableFunctionsNotAvailable sets VulnerableFunctionsNotAvailable field to given value.

### HasVulnerableFunctionsNotAvailable

`func (o *RemediationAssessment) HasVulnerableFunctionsNotAvailable() bool`

HasVulnerableFunctionsNotAvailable returns a boolean if a field has been set.

### GetVulnerableFunctionsNotInUse

`func (o *RemediationAssessment) GetVulnerableFunctionsNotInUse() []VulnerableFunction`

GetVulnerableFunctionsNotInUse returns the VulnerableFunctionsNotInUse field if non-nil, zero value otherwise.

### GetVulnerableFunctionsNotInUseOk

`func (o *RemediationAssessment) GetVulnerableFunctionsNotInUseOk() (*[]VulnerableFunction, bool)`

GetVulnerableFunctionsNotInUseOk returns a tuple with the VulnerableFunctionsNotInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctionsNotInUse

`func (o *RemediationAssessment) SetVulnerableFunctionsNotInUse(v []VulnerableFunction)`

SetVulnerableFunctionsNotInUse sets VulnerableFunctionsNotInUse field to given value.

### HasVulnerableFunctionsNotInUse

`func (o *RemediationAssessment) HasVulnerableFunctionsNotInUse() bool`

HasVulnerableFunctionsNotInUse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


