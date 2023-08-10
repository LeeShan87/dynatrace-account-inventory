# EffectivePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permission** | **string** | One of a effective permissions | 
**Effects** | [**[]EffectivePermissionEffects**](EffectivePermissionEffects.md) | A list of policies. | 

## Methods

### NewEffectivePermission

`func NewEffectivePermission(permission string, effects []EffectivePermissionEffects, ) *EffectivePermission`

NewEffectivePermission instantiates a new EffectivePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectivePermissionWithDefaults

`func NewEffectivePermissionWithDefaults() *EffectivePermission`

NewEffectivePermissionWithDefaults instantiates a new EffectivePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermission

`func (o *EffectivePermission) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *EffectivePermission) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *EffectivePermission) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetEffects

`func (o *EffectivePermission) GetEffects() []EffectivePermissionEffects`

GetEffects returns the Effects field if non-nil, zero value otherwise.

### GetEffectsOk

`func (o *EffectivePermission) GetEffectsOk() (*[]EffectivePermissionEffects, bool)`

GetEffectsOk returns a tuple with the Effects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffects

`func (o *EffectivePermission) SetEffects(v []EffectivePermissionEffects)`

SetEffects sets Effects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


