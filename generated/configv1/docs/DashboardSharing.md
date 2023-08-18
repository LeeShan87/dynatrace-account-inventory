# DashboardSharing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | The dashboard is shared (&#x60;true&#x60;) or private (&#x60;false&#x60;). | [optional] 
**Id** | **string** | The Dynatrace entity ID of the dashboard. | 
**Permissions** | [**[]DashboardSharePermissions**](DashboardSharePermissions.md) | A list of permissions to access the dashboard. | 
**Preset** | Pointer to **bool** | If &#x60;true&#x60; the dashboard will be marked as preset. | [optional] 
**PublicAccess** | [**DashboardAnonymousAccess**](DashboardAnonymousAccess.md) |  | 

## Methods

### NewDashboardSharing

`func NewDashboardSharing(id string, permissions []DashboardSharePermissions, publicAccess DashboardAnonymousAccess, ) *DashboardSharing`

NewDashboardSharing instantiates a new DashboardSharing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSharingWithDefaults

`func NewDashboardSharingWithDefaults() *DashboardSharing`

NewDashboardSharingWithDefaults instantiates a new DashboardSharing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DashboardSharing) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DashboardSharing) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DashboardSharing) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DashboardSharing) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *DashboardSharing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DashboardSharing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DashboardSharing) SetId(v string)`

SetId sets Id field to given value.


### GetPermissions

`func (o *DashboardSharing) GetPermissions() []DashboardSharePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DashboardSharing) GetPermissionsOk() (*[]DashboardSharePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DashboardSharing) SetPermissions(v []DashboardSharePermissions)`

SetPermissions sets Permissions field to given value.


### GetPreset

`func (o *DashboardSharing) GetPreset() bool`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *DashboardSharing) GetPresetOk() (*bool, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *DashboardSharing) SetPreset(v bool)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *DashboardSharing) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetPublicAccess

`func (o *DashboardSharing) GetPublicAccess() DashboardAnonymousAccess`

GetPublicAccess returns the PublicAccess field if non-nil, zero value otherwise.

### GetPublicAccessOk

`func (o *DashboardSharing) GetPublicAccessOk() (*DashboardAnonymousAccess, bool)`

GetPublicAccessOk returns a tuple with the PublicAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccess

`func (o *DashboardSharing) SetPublicAccess(v DashboardAnonymousAccess)`

SetPublicAccess sets PublicAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


