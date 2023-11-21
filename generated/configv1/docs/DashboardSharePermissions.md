# DashboardSharePermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the user or group to whom the permission is granted.  Not applicable if the **type** is set to &#x60;ALL&#x60;. | [optional] 
**Permission** | **string** | The level of the permission:    * &#x60;VIEW&#x60;: The dashboard is shared with read permission.  * &#x60;EDIT&#x60;: The dashboard is shared with edit permission.   | 
**Type** | **string** | The type of the permission:   * &#x60;USER&#x60;: The dashboard is shared with the specified user.  * &#x60;GROUP&#x60;: The dashboard is shared with all users of the specified group.  * &#x60;ALL&#x60;: The dashboard is shared via link. Any authenticated user with the link can view the dashboard. | 

## Methods

### NewDashboardSharePermissions

`func NewDashboardSharePermissions(permission string, type_ string, ) *DashboardSharePermissions`

NewDashboardSharePermissions instantiates a new DashboardSharePermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSharePermissionsWithDefaults

`func NewDashboardSharePermissionsWithDefaults() *DashboardSharePermissions`

NewDashboardSharePermissionsWithDefaults instantiates a new DashboardSharePermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DashboardSharePermissions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardSharePermissions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardSharePermissions) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DashboardSharePermissions) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPermission

`func (o *DashboardSharePermissions) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *DashboardSharePermissions) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *DashboardSharePermissions) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetType

`func (o *DashboardSharePermissions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardSharePermissions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardSharePermissions) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


