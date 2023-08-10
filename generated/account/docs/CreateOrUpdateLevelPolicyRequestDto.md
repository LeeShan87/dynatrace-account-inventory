# CreateOrUpdateLevelPolicyRequestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The display name of the policy. | 
**Description** | **string** | A short description of the policy. | 
**Tags** | **[]string** | A list of tags. | 
**StatementQuery** | **string** | The [statement](https://dt-url.net/ht03ucb) of the policy. | 

## Methods

### NewCreateOrUpdateLevelPolicyRequestDto

`func NewCreateOrUpdateLevelPolicyRequestDto(name string, description string, tags []string, statementQuery string, ) *CreateOrUpdateLevelPolicyRequestDto`

NewCreateOrUpdateLevelPolicyRequestDto instantiates a new CreateOrUpdateLevelPolicyRequestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrUpdateLevelPolicyRequestDtoWithDefaults

`func NewCreateOrUpdateLevelPolicyRequestDtoWithDefaults() *CreateOrUpdateLevelPolicyRequestDto`

NewCreateOrUpdateLevelPolicyRequestDtoWithDefaults instantiates a new CreateOrUpdateLevelPolicyRequestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateOrUpdateLevelPolicyRequestDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateOrUpdateLevelPolicyRequestDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateOrUpdateLevelPolicyRequestDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateOrUpdateLevelPolicyRequestDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateOrUpdateLevelPolicyRequestDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateOrUpdateLevelPolicyRequestDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTags

`func (o *CreateOrUpdateLevelPolicyRequestDto) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateOrUpdateLevelPolicyRequestDto) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateOrUpdateLevelPolicyRequestDto) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetStatementQuery

`func (o *CreateOrUpdateLevelPolicyRequestDto) GetStatementQuery() string`

GetStatementQuery returns the StatementQuery field if non-nil, zero value otherwise.

### GetStatementQueryOk

`func (o *CreateOrUpdateLevelPolicyRequestDto) GetStatementQueryOk() (*string, bool)`

GetStatementQueryOk returns a tuple with the StatementQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementQuery

`func (o *CreateOrUpdateLevelPolicyRequestDto) SetStatementQuery(v string)`

SetStatementQuery sets StatementQuery field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


