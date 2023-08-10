# EnvironmentResourceDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantResources** | [**[]TenantResourceDto**](TenantResourceDto.md) | A list of environments in the account. | 
**ManagementZoneResources** | [**[]ManagementZoneResourceDto**](ManagementZoneResourceDto.md) | A list of management zones in the account. | 

## Methods

### NewEnvironmentResourceDto

`func NewEnvironmentResourceDto(tenantResources []TenantResourceDto, managementZoneResources []ManagementZoneResourceDto, ) *EnvironmentResourceDto`

NewEnvironmentResourceDto instantiates a new EnvironmentResourceDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentResourceDtoWithDefaults

`func NewEnvironmentResourceDtoWithDefaults() *EnvironmentResourceDto`

NewEnvironmentResourceDtoWithDefaults instantiates a new EnvironmentResourceDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantResources

`func (o *EnvironmentResourceDto) GetTenantResources() []TenantResourceDto`

GetTenantResources returns the TenantResources field if non-nil, zero value otherwise.

### GetTenantResourcesOk

`func (o *EnvironmentResourceDto) GetTenantResourcesOk() (*[]TenantResourceDto, bool)`

GetTenantResourcesOk returns a tuple with the TenantResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantResources

`func (o *EnvironmentResourceDto) SetTenantResources(v []TenantResourceDto)`

SetTenantResources sets TenantResources field to given value.


### GetManagementZoneResources

`func (o *EnvironmentResourceDto) GetManagementZoneResources() []ManagementZoneResourceDto`

GetManagementZoneResources returns the ManagementZoneResources field if non-nil, zero value otherwise.

### GetManagementZoneResourcesOk

`func (o *EnvironmentResourceDto) GetManagementZoneResourcesOk() (*[]ManagementZoneResourceDto, bool)`

GetManagementZoneResourcesOk returns a tuple with the ManagementZoneResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZoneResources

`func (o *EnvironmentResourceDto) SetManagementZoneResources(v []ManagementZoneResourceDto)`

SetManagementZoneResources sets ManagementZoneResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


