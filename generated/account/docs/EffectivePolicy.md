# EffectivePolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The ID of the policy. | 
**Name** | **string** | The display name of the policy. | 
**StatementQuery** | **string** | The the statement query of the policy. | 
**LevelType** | **string** | The type of the level to which the policy applies. | 
**LevelId** | **string** | The ID of the level to which the policy applies. | 

## Methods

### NewEffectivePolicy

`func NewEffectivePolicy(uuid string, name string, statementQuery string, levelType string, levelId string, ) *EffectivePolicy`

NewEffectivePolicy instantiates a new EffectivePolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectivePolicyWithDefaults

`func NewEffectivePolicyWithDefaults() *EffectivePolicy`

NewEffectivePolicyWithDefaults instantiates a new EffectivePolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *EffectivePolicy) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EffectivePolicy) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EffectivePolicy) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *EffectivePolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EffectivePolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EffectivePolicy) SetName(v string)`

SetName sets Name field to given value.


### GetStatementQuery

`func (o *EffectivePolicy) GetStatementQuery() string`

GetStatementQuery returns the StatementQuery field if non-nil, zero value otherwise.

### GetStatementQueryOk

`func (o *EffectivePolicy) GetStatementQueryOk() (*string, bool)`

GetStatementQueryOk returns a tuple with the StatementQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementQuery

`func (o *EffectivePolicy) SetStatementQuery(v string)`

SetStatementQuery sets StatementQuery field to given value.


### GetLevelType

`func (o *EffectivePolicy) GetLevelType() string`

GetLevelType returns the LevelType field if non-nil, zero value otherwise.

### GetLevelTypeOk

`func (o *EffectivePolicy) GetLevelTypeOk() (*string, bool)`

GetLevelTypeOk returns a tuple with the LevelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelType

`func (o *EffectivePolicy) SetLevelType(v string)`

SetLevelType sets LevelType field to given value.


### GetLevelId

`func (o *EffectivePolicy) GetLevelId() string`

GetLevelId returns the LevelId field if non-nil, zero value otherwise.

### GetLevelIdOk

`func (o *EffectivePolicy) GetLevelIdOk() (*string, bool)`

GetLevelIdOk returns a tuple with the LevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelId

`func (o *EffectivePolicy) SetLevelId(v string)`

SetLevelId sets LevelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


