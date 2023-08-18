# ProcessGroupVulnerableFunctions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FunctionsInUse** | Pointer to [**[]VulnerableFunction**](VulnerableFunction.md) | A list of vulnerable functions in use. | [optional] [readonly] 
**FunctionsNotAvailable** | Pointer to [**[]VulnerableFunction**](VulnerableFunction.md) | A list of vulnerable functions with unknown state. | [optional] [readonly] 
**FunctionsNotInUse** | Pointer to [**[]VulnerableFunction**](VulnerableFunction.md) | A list of vulnerable functions not in use. | [optional] [readonly] 
**ProcessGroup** | Pointer to **string** | The process group identifier. | [optional] [readonly] 

## Methods

### NewProcessGroupVulnerableFunctions

`func NewProcessGroupVulnerableFunctions() *ProcessGroupVulnerableFunctions`

NewProcessGroupVulnerableFunctions instantiates a new ProcessGroupVulnerableFunctions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessGroupVulnerableFunctionsWithDefaults

`func NewProcessGroupVulnerableFunctionsWithDefaults() *ProcessGroupVulnerableFunctions`

NewProcessGroupVulnerableFunctionsWithDefaults instantiates a new ProcessGroupVulnerableFunctions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunctionsInUse

`func (o *ProcessGroupVulnerableFunctions) GetFunctionsInUse() []VulnerableFunction`

GetFunctionsInUse returns the FunctionsInUse field if non-nil, zero value otherwise.

### GetFunctionsInUseOk

`func (o *ProcessGroupVulnerableFunctions) GetFunctionsInUseOk() (*[]VulnerableFunction, bool)`

GetFunctionsInUseOk returns a tuple with the FunctionsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsInUse

`func (o *ProcessGroupVulnerableFunctions) SetFunctionsInUse(v []VulnerableFunction)`

SetFunctionsInUse sets FunctionsInUse field to given value.

### HasFunctionsInUse

`func (o *ProcessGroupVulnerableFunctions) HasFunctionsInUse() bool`

HasFunctionsInUse returns a boolean if a field has been set.

### GetFunctionsNotAvailable

`func (o *ProcessGroupVulnerableFunctions) GetFunctionsNotAvailable() []VulnerableFunction`

GetFunctionsNotAvailable returns the FunctionsNotAvailable field if non-nil, zero value otherwise.

### GetFunctionsNotAvailableOk

`func (o *ProcessGroupVulnerableFunctions) GetFunctionsNotAvailableOk() (*[]VulnerableFunction, bool)`

GetFunctionsNotAvailableOk returns a tuple with the FunctionsNotAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsNotAvailable

`func (o *ProcessGroupVulnerableFunctions) SetFunctionsNotAvailable(v []VulnerableFunction)`

SetFunctionsNotAvailable sets FunctionsNotAvailable field to given value.

### HasFunctionsNotAvailable

`func (o *ProcessGroupVulnerableFunctions) HasFunctionsNotAvailable() bool`

HasFunctionsNotAvailable returns a boolean if a field has been set.

### GetFunctionsNotInUse

`func (o *ProcessGroupVulnerableFunctions) GetFunctionsNotInUse() []VulnerableFunction`

GetFunctionsNotInUse returns the FunctionsNotInUse field if non-nil, zero value otherwise.

### GetFunctionsNotInUseOk

`func (o *ProcessGroupVulnerableFunctions) GetFunctionsNotInUseOk() (*[]VulnerableFunction, bool)`

GetFunctionsNotInUseOk returns a tuple with the FunctionsNotInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunctionsNotInUse

`func (o *ProcessGroupVulnerableFunctions) SetFunctionsNotInUse(v []VulnerableFunction)`

SetFunctionsNotInUse sets FunctionsNotInUse field to given value.

### HasFunctionsNotInUse

`func (o *ProcessGroupVulnerableFunctions) HasFunctionsNotInUse() bool`

HasFunctionsNotInUse returns a boolean if a field has been set.

### GetProcessGroup

`func (o *ProcessGroupVulnerableFunctions) GetProcessGroup() string`

GetProcessGroup returns the ProcessGroup field if non-nil, zero value otherwise.

### GetProcessGroupOk

`func (o *ProcessGroupVulnerableFunctions) GetProcessGroupOk() (*string, bool)`

GetProcessGroupOk returns a tuple with the ProcessGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroup

`func (o *ProcessGroupVulnerableFunctions) SetProcessGroup(v string)`

SetProcessGroup sets ProcessGroup field to given value.

### HasProcessGroup

`func (o *ProcessGroupVulnerableFunctions) HasProcessGroup() bool`

HasProcessGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


