# PermissionsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionName** | **string** | The name of the permission. | 
**Scope** | **string** | The scope of the permission. Depending on the scope type, it is defined by:   * &#x60;account&#x60;: The UUID of the account.  * &#x60;tenant&#x60;: The ID of the environment.  * &#x60;management-zone&#x60;: The ID of the management zone from an environment in &#x60;{environment-id}:{management-zone-id}&#x60; format. | 
**ScopeType** | **string** | The type of the permission scope. | 
**CreatedAt** | Pointer to **string** | The date and time of the permission creation in &#x60;2021-05-01T15:11:00Z&#x60; format. | [optional] 
**UpdatedAt** | Pointer to **string** | The date and time of the most recent permission modification in &#x60;2021-05-01T15:11:00Z&#x60; format. | [optional] 

## Methods

### NewPermissionsDto

`func NewPermissionsDto(permissionName string, scope string, scopeType string, ) *PermissionsDto`

NewPermissionsDto instantiates a new PermissionsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionsDtoWithDefaults

`func NewPermissionsDtoWithDefaults() *PermissionsDto`

NewPermissionsDtoWithDefaults instantiates a new PermissionsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionName

`func (o *PermissionsDto) GetPermissionName() string`

GetPermissionName returns the PermissionName field if non-nil, zero value otherwise.

### GetPermissionNameOk

`func (o *PermissionsDto) GetPermissionNameOk() (*string, bool)`

GetPermissionNameOk returns a tuple with the PermissionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionName

`func (o *PermissionsDto) SetPermissionName(v string)`

SetPermissionName sets PermissionName field to given value.


### GetScope

`func (o *PermissionsDto) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PermissionsDto) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PermissionsDto) SetScope(v string)`

SetScope sets Scope field to given value.


### GetScopeType

`func (o *PermissionsDto) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *PermissionsDto) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *PermissionsDto) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.


### GetCreatedAt

`func (o *PermissionsDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PermissionsDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PermissionsDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PermissionsDto) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PermissionsDto) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PermissionsDto) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PermissionsDto) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PermissionsDto) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


