# CreateLevelPolicyBindingsRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyBindings** | [**[]Binding**](Binding.md) | A list of policy bindings of the user group. | 

## Methods

### NewCreateLevelPolicyBindingsRequestDto

`func NewCreateLevelPolicyBindingsRequestDto(policyBindings []Binding, ) *CreateLevelPolicyBindingsRequestDto`

NewCreateLevelPolicyBindingsRequestDto instantiates a new CreateLevelPolicyBindingsRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLevelPolicyBindingsRequestDtoWithDefaults

`func NewCreateLevelPolicyBindingsRequestDtoWithDefaults() *CreateLevelPolicyBindingsRequestDto`

NewCreateLevelPolicyBindingsRequestDtoWithDefaults instantiates a new CreateLevelPolicyBindingsRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyBindings

`func (o *CreateLevelPolicyBindingsRequestDto) GetPolicyBindings() []Binding`

GetPolicyBindings returns the PolicyBindings field if non-nil, zero value otherwise.

### GetPolicyBindingsOk

`func (o *CreateLevelPolicyBindingsRequestDto) GetPolicyBindingsOk() (*[]Binding, bool)`

GetPolicyBindingsOk returns a tuple with the PolicyBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyBindings

`func (o *CreateLevelPolicyBindingsRequestDto) SetPolicyBindings(v []Binding)`

SetPolicyBindings sets PolicyBindings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


