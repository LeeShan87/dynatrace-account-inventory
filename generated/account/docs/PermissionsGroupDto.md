# PermissionsGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID of the user group. | [optional] 
**Name** | **string** | The name of the user group. | 
**Description** | Pointer to **string** | A short description of the user group. | [optional] 
**FederatedAttributeValues** | Pointer to **[]string** | A list of values associating this group with the corresponding claim from an identity provider. | [optional] 
**Owner** | **string** | The identity provider from which the group originates. | 
**CreatedAt** | **string** | The date and time of the group creation in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**UpdatedAt** | **string** | The date and time of the most recent group modification in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**Permissions** | [**[]PermissionsDto**](PermissionsDto.md) | A list of permissions assigned to the group. | 

## Methods

### NewPermissionsGroupDto

`func NewPermissionsGroupDto(name string, owner string, createdAt string, updatedAt string, permissions []PermissionsDto, ) *PermissionsGroupDto`

NewPermissionsGroupDto instantiates a new PermissionsGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsGroupDtoWithDefaults

`func NewPermissionsGroupDtoWithDefaults() *PermissionsGroupDto`

NewPermissionsGroupDtoWithDefaults instantiates a new PermissionsGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PermissionsGroupDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PermissionsGroupDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PermissionsGroupDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PermissionsGroupDto) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *PermissionsGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionsGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionsGroupDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PermissionsGroupDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PermissionsGroupDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PermissionsGroupDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PermissionsGroupDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFederatedAttributeValues

`func (o *PermissionsGroupDto) GetFederatedAttributeValues() []string`

GetFederatedAttributeValues returns the FederatedAttributeValues field if non-nil, zero value otherwise.

### GetFederatedAttributeValuesOk

`func (o *PermissionsGroupDto) GetFederatedAttributeValuesOk() (*[]string, bool)`

GetFederatedAttributeValuesOk returns a tuple with the FederatedAttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedAttributeValues

`func (o *PermissionsGroupDto) SetFederatedAttributeValues(v []string)`

SetFederatedAttributeValues sets FederatedAttributeValues field to given value.

### HasFederatedAttributeValues

`func (o *PermissionsGroupDto) HasFederatedAttributeValues() bool`

HasFederatedAttributeValues returns a boolean if a field has been set.

### GetOwner

`func (o *PermissionsGroupDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PermissionsGroupDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PermissionsGroupDto) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetCreatedAt

`func (o *PermissionsGroupDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PermissionsGroupDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PermissionsGroupDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PermissionsGroupDto) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PermissionsGroupDto) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PermissionsGroupDto) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPermissions

`func (o *PermissionsGroupDto) GetPermissions() []PermissionsDto`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionsGroupDto) GetPermissionsOk() (*[]PermissionsDto, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionsGroupDto) SetPermissions(v []PermissionsDto)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


