# AppendLevelPolicyBindingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | **[]string** | A list of user groups (specified by IDs) to which the policy applies. | 

## Methods

### NewAppendLevelPolicyBindingsRequestDto

`func NewAppendLevelPolicyBindingsRequestDto(groups []string, ) *AppendLevelPolicyBindingsRequestDto`

NewAppendLevelPolicyBindingsRequestDto instantiates a new AppendLevelPolicyBindingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppendLevelPolicyBindingsRequestDtoWithDefaults

`func NewAppendLevelPolicyBindingsRequestDtoWithDefaults() *AppendLevelPolicyBindingsRequestDto`

NewAppendLevelPolicyBindingsRequestDtoWithDefaults instantiates a new AppendLevelPolicyBindingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *AppendLevelPolicyBindingsRequestDto) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *AppendLevelPolicyBindingsRequestDto) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *AppendLevelPolicyBindingsRequestDto) SetGroups(v []string)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


