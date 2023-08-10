# PutGroupDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID of the user group. | [optional] 
**Name** | **string** | The name of the user group. | 
**Description** | Pointer to **string** | A short description of the user group. | [optional] 
**FederatedAttributeValues** | Pointer to **[]string** | A list of values associating this group with the corresponding claim from an identity provider. If present, sets &#x60;owner&#x60; to SAML, otherwise &#x60;owner&#x60; will be LOCAL | [optional] 
**Owner** | Pointer to **map[string]interface{}** | The owner type of the group. | [optional] 

## Methods

### NewPutGroupDto

`func NewPutGroupDto(name string, ) *PutGroupDto`

NewPutGroupDto instantiates a new PutGroupDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutGroupDtoWithDefaults

`func NewPutGroupDtoWithDefaults() *PutGroupDto`

NewPutGroupDtoWithDefaults instantiates a new PutGroupDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *PutGroupDto) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *PutGroupDto) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *PutGroupDto) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *PutGroupDto) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *PutGroupDto) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutGroupDto) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutGroupDto) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PutGroupDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutGroupDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutGroupDto) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutGroupDto) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFederatedAttributeValues

`func (o *PutGroupDto) GetFederatedAttributeValues() []string`

GetFederatedAttributeValues returns the FederatedAttributeValues field if non-nil, zero value otherwise.

### GetFederatedAttributeValuesOk

`func (o *PutGroupDto) GetFederatedAttributeValuesOk() (*[]string, bool)`

GetFederatedAttributeValuesOk returns a tuple with the FederatedAttributeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedAttributeValues

`func (o *PutGroupDto) SetFederatedAttributeValues(v []string)`

SetFederatedAttributeValues sets FederatedAttributeValues field to given value.

### HasFederatedAttributeValues

`func (o *PutGroupDto) HasFederatedAttributeValues() bool`

HasFederatedAttributeValues returns a boolean if a field has been set.

### GetOwner

`func (o *PutGroupDto) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PutGroupDto) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PutGroupDto) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PutGroupDto) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


