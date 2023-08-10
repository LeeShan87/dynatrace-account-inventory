# LevelPolicyBindingDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LevelType** | **string** | The type of the policy level. | 
**LevelId** | **string** | The ID of the policy level. | 
**PolicyBindings** | [**[]Binding**](Binding.md) |  | 

## Methods

### NewLevelPolicyBindingDto

`func NewLevelPolicyBindingDto(levelType string, levelId string, policyBindings []Binding, ) *LevelPolicyBindingDto`

NewLevelPolicyBindingDto instantiates a new LevelPolicyBindingDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLevelPolicyBindingDtoWithDefaults

`func NewLevelPolicyBindingDtoWithDefaults() *LevelPolicyBindingDto`

NewLevelPolicyBindingDtoWithDefaults instantiates a new LevelPolicyBindingDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevelType

`func (o *LevelPolicyBindingDto) GetLevelType() string`

GetLevelType returns the LevelType field if non-nil, zero value otherwise.

### GetLevelTypeOk

`func (o *LevelPolicyBindingDto) GetLevelTypeOk() (*string, bool)`

GetLevelTypeOk returns a tuple with the LevelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelType

`func (o *LevelPolicyBindingDto) SetLevelType(v string)`

SetLevelType sets LevelType field to given value.


### GetLevelId

`func (o *LevelPolicyBindingDto) GetLevelId() string`

GetLevelId returns the LevelId field if non-nil, zero value otherwise.

### GetLevelIdOk

`func (o *LevelPolicyBindingDto) GetLevelIdOk() (*string, bool)`

GetLevelIdOk returns a tuple with the LevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelId

`func (o *LevelPolicyBindingDto) SetLevelId(v string)`

SetLevelId sets LevelId field to given value.


### GetPolicyBindings

`func (o *LevelPolicyBindingDto) GetPolicyBindings() []Binding`

GetPolicyBindings returns the PolicyBindings field if non-nil, zero value otherwise.

### GetPolicyBindingsOk

`func (o *LevelPolicyBindingDto) GetPolicyBindingsOk() (*[]Binding, bool)`

GetPolicyBindingsOk returns a tuple with the PolicyBindings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyBindings

`func (o *LevelPolicyBindingDto) SetPolicyBindings(v []Binding)`

SetPolicyBindings sets PolicyBindings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


