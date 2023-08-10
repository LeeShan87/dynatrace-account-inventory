# EffectivePermissionEffects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | Effect of policy | 
**Conditions** | [**[]Condition**](Condition.md) | Policy condition | 
**EffectivePolicies** | [**[]EffectivePolicyWithBinding**](EffectivePolicyWithBinding.md) | A list of effective policies. | 

## Methods

### NewEffectivePermissionEffects

`func NewEffectivePermissionEffects(effect string, conditions []Condition, effectivePolicies []EffectivePolicyWithBinding, ) *EffectivePermissionEffects`

NewEffectivePermissionEffects instantiates a new EffectivePermissionEffects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectivePermissionEffectsWithDefaults

`func NewEffectivePermissionEffectsWithDefaults() *EffectivePermissionEffects`

NewEffectivePermissionEffectsWithDefaults instantiates a new EffectivePermissionEffects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *EffectivePermissionEffects) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *EffectivePermissionEffects) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *EffectivePermissionEffects) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetConditions

`func (o *EffectivePermissionEffects) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *EffectivePermissionEffects) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *EffectivePermissionEffects) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.


### GetEffectivePolicies

`func (o *EffectivePermissionEffects) GetEffectivePolicies() []EffectivePolicyWithBinding`

GetEffectivePolicies returns the EffectivePolicies field if non-nil, zero value otherwise.

### GetEffectivePoliciesOk

`func (o *EffectivePermissionEffects) GetEffectivePoliciesOk() (*[]EffectivePolicyWithBinding, bool)`

GetEffectivePoliciesOk returns a tuple with the EffectivePolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePolicies

`func (o *EffectivePermissionEffects) SetEffectivePolicies(v []EffectivePolicyWithBinding)`

SetEffectivePolicies sets EffectivePolicies field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


