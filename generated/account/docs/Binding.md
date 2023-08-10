# Binding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyUuid** | **string** | The ID of the policy. | 
**Groups** | **[]string** | A list of user groups to which the policy applies. | 

## Methods

### NewBinding

`func NewBinding(policyUuid string, groups []string, ) *Binding`

NewBinding instantiates a new Binding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBindingWithDefaults

`func NewBindingWithDefaults() *Binding`

NewBindingWithDefaults instantiates a new Binding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyUuid

`func (o *Binding) GetPolicyUuid() string`

GetPolicyUuid returns the PolicyUuid field if non-nil, zero value otherwise.

### GetPolicyUuidOk

`func (o *Binding) GetPolicyUuidOk() (*string, bool)`

GetPolicyUuidOk returns a tuple with the PolicyUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyUuid

`func (o *Binding) SetPolicyUuid(v string)`

SetPolicyUuid sets PolicyUuid field to given value.


### GetGroups

`func (o *Binding) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *Binding) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *Binding) SetGroups(v []string)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


