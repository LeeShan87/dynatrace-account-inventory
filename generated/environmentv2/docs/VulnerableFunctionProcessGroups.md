# VulnerableFunctionProcessGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Function** | Pointer to [**VulnerableFunction**](VulnerableFunction.md) |  | [optional] 
**ProcessGroupsInUse** | Pointer to **[]string** | The process group identifiers, where this vulnerable function is in use. | [optional] [readonly] 
**ProcessGroupsNotAvailable** | Pointer to **[]string** | The process group identifiers, where information about the usage of this function not available. | [optional] [readonly] 
**ProcessGroupsNotInUse** | Pointer to **[]string** | The process group identifiers, where this vulnerable function is not in use. | [optional] [readonly] 
**Usage** | Pointer to **string** | The vulnerable function usage based on the given process groups: * IN_USE if at least one process group calls this vulnerable function. * NOT_IN_USE if all process groups do not call this vulnerable function. * NOT_AVAILABLE if vulnerable function usage could not be calculated for at least one process group and no process group calls this vulnerable function. | [optional] [readonly] 

## Methods

### NewVulnerableFunctionProcessGroups

`func NewVulnerableFunctionProcessGroups() *VulnerableFunctionProcessGroups`

NewVulnerableFunctionProcessGroups instantiates a new VulnerableFunctionProcessGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVulnerableFunctionProcessGroupsWithDefaults

`func NewVulnerableFunctionProcessGroupsWithDefaults() *VulnerableFunctionProcessGroups`

NewVulnerableFunctionProcessGroupsWithDefaults instantiates a new VulnerableFunctionProcessGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFunction

`func (o *VulnerableFunctionProcessGroups) GetFunction() VulnerableFunction`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *VulnerableFunctionProcessGroups) GetFunctionOk() (*VulnerableFunction, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *VulnerableFunctionProcessGroups) SetFunction(v VulnerableFunction)`

SetFunction sets Function field to given value.

### HasFunction

`func (o *VulnerableFunctionProcessGroups) HasFunction() bool`

HasFunction returns a boolean if a field has been set.

### GetProcessGroupsInUse

`func (o *VulnerableFunctionProcessGroups) GetProcessGroupsInUse() []string`

GetProcessGroupsInUse returns the ProcessGroupsInUse field if non-nil, zero value otherwise.

### GetProcessGroupsInUseOk

`func (o *VulnerableFunctionProcessGroups) GetProcessGroupsInUseOk() (*[]string, bool)`

GetProcessGroupsInUseOk returns a tuple with the ProcessGroupsInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupsInUse

`func (o *VulnerableFunctionProcessGroups) SetProcessGroupsInUse(v []string)`

SetProcessGroupsInUse sets ProcessGroupsInUse field to given value.

### HasProcessGroupsInUse

`func (o *VulnerableFunctionProcessGroups) HasProcessGroupsInUse() bool`

HasProcessGroupsInUse returns a boolean if a field has been set.

### GetProcessGroupsNotAvailable

`func (o *VulnerableFunctionProcessGroups) GetProcessGroupsNotAvailable() []string`

GetProcessGroupsNotAvailable returns the ProcessGroupsNotAvailable field if non-nil, zero value otherwise.

### GetProcessGroupsNotAvailableOk

`func (o *VulnerableFunctionProcessGroups) GetProcessGroupsNotAvailableOk() (*[]string, bool)`

GetProcessGroupsNotAvailableOk returns a tuple with the ProcessGroupsNotAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupsNotAvailable

`func (o *VulnerableFunctionProcessGroups) SetProcessGroupsNotAvailable(v []string)`

SetProcessGroupsNotAvailable sets ProcessGroupsNotAvailable field to given value.

### HasProcessGroupsNotAvailable

`func (o *VulnerableFunctionProcessGroups) HasProcessGroupsNotAvailable() bool`

HasProcessGroupsNotAvailable returns a boolean if a field has been set.

### GetProcessGroupsNotInUse

`func (o *VulnerableFunctionProcessGroups) GetProcessGroupsNotInUse() []string`

GetProcessGroupsNotInUse returns the ProcessGroupsNotInUse field if non-nil, zero value otherwise.

### GetProcessGroupsNotInUseOk

`func (o *VulnerableFunctionProcessGroups) GetProcessGroupsNotInUseOk() (*[]string, bool)`

GetProcessGroupsNotInUseOk returns a tuple with the ProcessGroupsNotInUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessGroupsNotInUse

`func (o *VulnerableFunctionProcessGroups) SetProcessGroupsNotInUse(v []string)`

SetProcessGroupsNotInUse sets ProcessGroupsNotInUse field to given value.

### HasProcessGroupsNotInUse

`func (o *VulnerableFunctionProcessGroups) HasProcessGroupsNotInUse() bool`

HasProcessGroupsNotInUse returns a boolean if a field has been set.

### GetUsage

`func (o *VulnerableFunctionProcessGroups) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *VulnerableFunctionProcessGroups) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *VulnerableFunctionProcessGroups) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *VulnerableFunctionProcessGroups) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


