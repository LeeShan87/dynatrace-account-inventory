# VulnerableFunctionsContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VulnerableFunctions** | Pointer to [**[]VulnerableFunctionProcessGroups**](VulnerableFunctionProcessGroups.md) | A list of vulnerable functions, their security problem wide usages and their usages per process group. | [optional] [readonly] 
**VulnerableFunctionsByProcessGroup** | Pointer to [**[]ProcessGroupVulnerableFunctions**](ProcessGroupVulnerableFunctions.md) | A list of vulnerable function usages per process group for a security problem. The result is sorted based on the following criteria:  * the number of vulnerable functions in use (descending). * the number of vulnerable functions not in use (descending). * the number of vulnerable functions not available (descending). * the process group identifier (ascending) | [optional] [readonly] 

## Methods

### NewVulnerableFunctionsContainer

`func NewVulnerableFunctionsContainer() *VulnerableFunctionsContainer`

NewVulnerableFunctionsContainer instantiates a new VulnerableFunctionsContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnerableFunctionsContainerWithDefaults

`func NewVulnerableFunctionsContainerWithDefaults() *VulnerableFunctionsContainer`

NewVulnerableFunctionsContainerWithDefaults instantiates a new VulnerableFunctionsContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVulnerableFunctions

`func (o *VulnerableFunctionsContainer) GetVulnerableFunctions() []VulnerableFunctionProcessGroups`

GetVulnerableFunctions returns the VulnerableFunctions field if non-nil, zero value otherwise.

### GetVulnerableFunctionsOk

`func (o *VulnerableFunctionsContainer) GetVulnerableFunctionsOk() (*[]VulnerableFunctionProcessGroups, bool)`

GetVulnerableFunctionsOk returns a tuple with the VulnerableFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctions

`func (o *VulnerableFunctionsContainer) SetVulnerableFunctions(v []VulnerableFunctionProcessGroups)`

SetVulnerableFunctions sets VulnerableFunctions field to given value.

### HasVulnerableFunctions

`func (o *VulnerableFunctionsContainer) HasVulnerableFunctions() bool`

HasVulnerableFunctions returns a boolean if a field has been set.

### GetVulnerableFunctionsByProcessGroup

`func (o *VulnerableFunctionsContainer) GetVulnerableFunctionsByProcessGroup() []ProcessGroupVulnerableFunctions`

GetVulnerableFunctionsByProcessGroup returns the VulnerableFunctionsByProcessGroup field if non-nil, zero value otherwise.

### GetVulnerableFunctionsByProcessGroupOk

`func (o *VulnerableFunctionsContainer) GetVulnerableFunctionsByProcessGroupOk() (*[]ProcessGroupVulnerableFunctions, bool)`

GetVulnerableFunctionsByProcessGroupOk returns a tuple with the VulnerableFunctionsByProcessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctionsByProcessGroup

`func (o *VulnerableFunctionsContainer) SetVulnerableFunctionsByProcessGroup(v []ProcessGroupVulnerableFunctions)`

SetVulnerableFunctionsByProcessGroup sets VulnerableFunctionsByProcessGroup field to given value.

### HasVulnerableFunctionsByProcessGroup

`func (o *VulnerableFunctionsContainer) HasVulnerableFunctionsByProcessGroup() bool`

HasVulnerableFunctionsByProcessGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


