# GetGroupDto

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

## Methods

### NewGetGroupDto

`func NewGetGroupDto(name string, owner string, createdAt string, updatedAt string, ) *GetGroupDto`

NewGetGroupDto instantiates a new GetGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupDtoWithDefaults

`func NewGetGroupDtoWithDefaults() *GetGroupDto`

NewGetGroupDtoWithDefaults instantiates a new GetGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *GetGroupDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGroupDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGroupDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGroupDto) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *GetGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGroupDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GetGroupDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGroupDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGroupDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGroupDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFederatedAttributeValues

`func (o *GetGroupDto) GetFederatedAttributeValues() []string`

GetFederatedAttributeValues returns the FederatedAttributeValues field if non-nil, zero value otherwise.

### GetFederatedAttributeValuesOk

`func (o *GetGroupDto) GetFederatedAttributeValuesOk() (*[]string, bool)`

GetFederatedAttributeValuesOk returns a tuple with the FederatedAttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedAttributeValues

`func (o *GetGroupDto) SetFederatedAttributeValues(v []string)`

SetFederatedAttributeValues sets FederatedAttributeValues field to given value.

### HasFederatedAttributeValues

`func (o *GetGroupDto) HasFederatedAttributeValues() bool`

HasFederatedAttributeValues returns a boolean if a field has been set.

### GetOwner

`func (o *GetGroupDto) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetGroupDto) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetGroupDto) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetCreatedAt

`func (o *GetGroupDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetGroupDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetGroupDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetGroupDto) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetGroupDto) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetGroupDto) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


