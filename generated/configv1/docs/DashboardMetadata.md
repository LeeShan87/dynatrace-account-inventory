# DashboardMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardFilter** | Pointer to [**DashboardFilter**](DashboardFilter.md) |  | [optional] 
**DynamicFilters** | Pointer to [**DynamicFilters**](DynamicFilters.md) |  | [optional] 
**HasConsistentColors** | Pointer to **bool** | The tile uses consistent colors when rendering its content. | [optional] 
**Name** | **string** | The name of the dashboard. | 
**Owner** | **string** | The owner of the dashboard. | 
**Preset** | Pointer to **bool** | The dashboard is a preset (&#x60;true&#x60;) or a custom (&#x60;false&#x60;) dashboard. | [optional] 
**Shared** | Pointer to **bool** | The dashboard is shared (&#x60;true&#x60;) or private (&#x60;false&#x60;). | [optional] 
**Tags** | Pointer to **[]string** | A set of tags assigned to the dashboard. | [optional] 
**TilesNameSize** | Pointer to **string** | The general size of the tiles tile. Default value is medium | [optional] 

## Methods

### NewDashboardMetadata

`func NewDashboardMetadata(name string, owner string, ) *DashboardMetadata`

NewDashboardMetadata instantiates a new DashboardMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardMetadataWithDefaults

`func NewDashboardMetadataWithDefaults() *DashboardMetadata`

NewDashboardMetadataWithDefaults instantiates a new DashboardMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardFilter

`func (o *DashboardMetadata) GetDashboardFilter() DashboardFilter`

GetDashboardFilter returns the DashboardFilter field if non-nil, zero value otherwise.

### GetDashboardFilterOk

`func (o *DashboardMetadata) GetDashboardFilterOk() (*DashboardFilter, bool)`

GetDashboardFilterOk returns a tuple with the DashboardFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardFilter

`func (o *DashboardMetadata) SetDashboardFilter(v DashboardFilter)`

SetDashboardFilter sets DashboardFilter field to given value.

### HasDashboardFilter

`func (o *DashboardMetadata) HasDashboardFilter() bool`

HasDashboardFilter returns a boolean if a field has been set.

### GetDynamicFilters

`func (o *DashboardMetadata) GetDynamicFilters() DynamicFilters`

GetDynamicFilters returns the DynamicFilters field if non-nil, zero value otherwise.

### GetDynamicFiltersOk

`func (o *DashboardMetadata) GetDynamicFiltersOk() (*DynamicFilters, bool)`

GetDynamicFiltersOk returns a tuple with the DynamicFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicFilters

`func (o *DashboardMetadata) SetDynamicFilters(v DynamicFilters)`

SetDynamicFilters sets DynamicFilters field to given value.

### HasDynamicFilters

`func (o *DashboardMetadata) HasDynamicFilters() bool`

HasDynamicFilters returns a boolean if a field has been set.

### GetHasConsistentColors

`func (o *DashboardMetadata) GetHasConsistentColors() bool`

GetHasConsistentColors returns the HasConsistentColors field if non-nil, zero value otherwise.

### GetHasConsistentColorsOk

`func (o *DashboardMetadata) GetHasConsistentColorsOk() (*bool, bool)`

GetHasConsistentColorsOk returns a tuple with the HasConsistentColors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConsistentColors

`func (o *DashboardMetadata) SetHasConsistentColors(v bool)`

SetHasConsistentColors sets HasConsistentColors field to given value.

### HasHasConsistentColors

`func (o *DashboardMetadata) HasHasConsistentColors() bool`

HasHasConsistentColors returns a boolean if a field has been set.

### GetName

`func (o *DashboardMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DashboardMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DashboardMetadata) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *DashboardMetadata) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *DashboardMetadata) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *DashboardMetadata) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPreset

`func (o *DashboardMetadata) GetPreset() bool`

GetPreset returns the Preset field if non-nil, zero value otherwise.

### GetPresetOk

`func (o *DashboardMetadata) GetPresetOk() (*bool, bool)`

GetPresetOk returns a tuple with the Preset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreset

`func (o *DashboardMetadata) SetPreset(v bool)`

SetPreset sets Preset field to given value.

### HasPreset

`func (o *DashboardMetadata) HasPreset() bool`

HasPreset returns a boolean if a field has been set.

### GetShared

`func (o *DashboardMetadata) GetShared() bool`

GetShared returns the Shared field if non-nil, zero value otherwise.

### GetSharedOk

`func (o *DashboardMetadata) GetSharedOk() (*bool, bool)`

GetSharedOk returns a tuple with the Shared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShared

`func (o *DashboardMetadata) SetShared(v bool)`

SetShared sets Shared field to given value.

### HasShared

`func (o *DashboardMetadata) HasShared() bool`

HasShared returns a boolean if a field has been set.

### GetTags

`func (o *DashboardMetadata) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DashboardMetadata) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DashboardMetadata) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DashboardMetadata) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTilesNameSize

`func (o *DashboardMetadata) GetTilesNameSize() string`

GetTilesNameSize returns the TilesNameSize field if non-nil, zero value otherwise.

### GetTilesNameSizeOk

`func (o *DashboardMetadata) GetTilesNameSizeOk() (*string, bool)`

GetTilesNameSizeOk returns a tuple with the TilesNameSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTilesNameSize

`func (o *DashboardMetadata) SetTilesNameSize(v string)`

SetTilesNameSize sets TilesNameSize field to given value.

### HasTilesNameSize

`func (o *DashboardMetadata) HasTilesNameSize() bool`

HasTilesNameSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


