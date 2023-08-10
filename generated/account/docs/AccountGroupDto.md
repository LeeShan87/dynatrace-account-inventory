# AccountGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupName** | **string** | The name of the user group. | 
**Uuid** | **string** | The UUID of the user group. | 
**Owner** | **string** | The identity provider from which the group originates. | 
**AccountUUID** | **string** | The UUID of the Dynatrace account. | 
**AccountName** | **string** | The name of the Dynatrace account. | 
**Description** | **string** | A short description of the group. | 
**CreatedAt** | **string** | The date and time of the group creation in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**UpdatedAt** | **string** | The date and time of the most recent modification to the group in &#x60;2021-05-01T15:11:00Z&#x60; format. | 

## Methods

### NewAccountGroupDto

`func NewAccountGroupDto(groupName string, uuid string, owner string, accountUUID string, accountName string, description string, createdAt string, updatedAt string, ) *AccountGroupDto`

NewAccountGroupDto instantiates a new AccountGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountGroupDtoWithDefaults

`func NewAccountGroupDtoWithDefaults() *AccountGroupDto`

NewAccountGroupDtoWithDefaults instantiates a new AccountGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupName

`func (o *AccountGroupDto) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *AccountGroupDto) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *AccountGroupDto) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.


### GetUuid

`func (o *AccountGroupDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AccountGroupDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AccountGroupDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetOwner

`func (o *AccountGroupDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AccountGroupDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AccountGroupDto) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetAccountUUID

`func (o *AccountGroupDto) GetAccountUUID() string`

GetAccountUUID returns the AccountUUID field if non-nil, zero value otherwise.

### GetAccountUUIDOk

`func (o *AccountGroupDto) GetAccountUUIDOk() (*string, bool)`

GetAccountUUIDOk returns a tuple with the AccountUUID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUUID

`func (o *AccountGroupDto) SetAccountUUID(v string)`

SetAccountUUID sets AccountUUID field to given value.


### GetAccountName

`func (o *AccountGroupDto) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AccountGroupDto) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AccountGroupDto) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetDescription

`func (o *AccountGroupDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountGroupDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountGroupDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *AccountGroupDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccountGroupDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccountGroupDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AccountGroupDto) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AccountGroupDto) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AccountGroupDto) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


