# AttackTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | Pointer to **string** | The monitored entity ID of the targeted host/database. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the targeted host/database. | [optional] [readonly] 

## Methods

### NewAttackTarget

`func NewAttackTarget() *AttackTarget`

NewAttackTarget instantiates a new AttackTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttackTargetWithDefaults

`func NewAttackTargetWithDefaults() *AttackTarget`

NewAttackTargetWithDefaults instantiates a new AttackTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *AttackTarget) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *AttackTarget) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *AttackTarget) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *AttackTarget) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### GetName

`func (o *AttackTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AttackTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AttackTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AttackTarget) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


