# LevelPolicyDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | The ID of the policy. | 
**Name** | **string** | The display name of the policy. | 
**Tags** | **[]string** | A list of tags. | 
**Description** | **string** | A short description of the policy. | 
**StatementQuery** | **string** | The [statement](https://dt-url.net/ht03ucb) of the policy. | 
**Statements** | [**[]Statement**](Statement.md) | The expanded form of the policy statement. | 

## Methods

### NewLevelPolicyDto

`func NewLevelPolicyDto(uuid string, name string, tags []string, description string, statementQuery string, statements []Statement, ) *LevelPolicyDto`

NewLevelPolicyDto instantiates a new LevelPolicyDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLevelPolicyDtoWithDefaults

`func NewLevelPolicyDtoWithDefaults() *LevelPolicyDto`

NewLevelPolicyDtoWithDefaults instantiates a new LevelPolicyDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *LevelPolicyDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *LevelPolicyDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *LevelPolicyDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *LevelPolicyDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LevelPolicyDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LevelPolicyDto) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *LevelPolicyDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LevelPolicyDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LevelPolicyDto) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetDescription

`func (o *LevelPolicyDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LevelPolicyDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LevelPolicyDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatementQuery

`func (o *LevelPolicyDto) GetStatementQuery() string`

GetStatementQuery returns the StatementQuery field if non-nil, zero value otherwise.

### GetStatementQueryOk

`func (o *LevelPolicyDto) GetStatementQueryOk() (*string, bool)`

GetStatementQueryOk returns a tuple with the StatementQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementQuery

`func (o *LevelPolicyDto) SetStatementQuery(v string)`

SetStatementQuery sets StatementQuery field to given value.


### GetStatements

`func (o *LevelPolicyDto) GetStatements() []Statement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *LevelPolicyDto) GetStatementsOk() (*[]Statement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *LevelPolicyDto) SetStatements(v []Statement)`

SetStatements sets Statements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


