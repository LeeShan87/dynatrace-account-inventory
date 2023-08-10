# PermissionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the permission. | 
**Description** | **string** | The display name of the permission. | 

## Methods

### NewPermissionDto

`func NewPermissionDto(id string, description string, ) *PermissionDto`

NewPermissionDto instantiates a new PermissionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionDtoWithDefaults

`func NewPermissionDtoWithDefaults() *PermissionDto`

NewPermissionDtoWithDefaults instantiates a new PermissionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PermissionDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionDto) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *PermissionDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PermissionDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PermissionDto) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


