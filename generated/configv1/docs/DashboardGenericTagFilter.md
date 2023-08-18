# DashboardGenericTagFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The display name used to identify this generic filter. | [optional] 
**EntityTypes** | **[]string** | Entity types affected by tag. | 
**SuggestionsFromEntityType** | Pointer to **string** | The entity type for which the suggestions should be provided. | [optional] 
**TagKey** | Pointer to **string** | The tag key for this filter. | [optional] 

## Methods

### NewDashboardGenericTagFilter

`func NewDashboardGenericTagFilter(entityTypes []string, ) *DashboardGenericTagFilter`

NewDashboardGenericTagFilter instantiates a new DashboardGenericTagFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardGenericTagFilterWithDefaults

`func NewDashboardGenericTagFilterWithDefaults() *DashboardGenericTagFilter`

NewDashboardGenericTagFilterWithDefaults instantiates a new DashboardGenericTagFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *DashboardGenericTagFilter) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DashboardGenericTagFilter) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DashboardGenericTagFilter) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DashboardGenericTagFilter) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEntityTypes

`func (o *DashboardGenericTagFilter) GetEntityTypes() []string`

GetEntityTypes returns the EntityTypes field if non-nil, zero value otherwise.

### GetEntityTypesOk

`func (o *DashboardGenericTagFilter) GetEntityTypesOk() (*[]string, bool)`

GetEntityTypesOk returns a tuple with the EntityTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityTypes

`func (o *DashboardGenericTagFilter) SetEntityTypes(v []string)`

SetEntityTypes sets EntityTypes field to given value.


### GetSuggestionsFromEntityType

`func (o *DashboardGenericTagFilter) GetSuggestionsFromEntityType() string`

GetSuggestionsFromEntityType returns the SuggestionsFromEntityType field if non-nil, zero value otherwise.

### GetSuggestionsFromEntityTypeOk

`func (o *DashboardGenericTagFilter) GetSuggestionsFromEntityTypeOk() (*string, bool)`

GetSuggestionsFromEntityTypeOk returns a tuple with the SuggestionsFromEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionsFromEntityType

`func (o *DashboardGenericTagFilter) SetSuggestionsFromEntityType(v string)`

SetSuggestionsFromEntityType sets SuggestionsFromEntityType field to given value.

### HasSuggestionsFromEntityType

`func (o *DashboardGenericTagFilter) HasSuggestionsFromEntityType() bool`

HasSuggestionsFromEntityType returns a boolean if a field has been set.

### GetTagKey

`func (o *DashboardGenericTagFilter) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *DashboardGenericTagFilter) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *DashboardGenericTagFilter) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.

### HasTagKey

`func (o *DashboardGenericTagFilter) HasTagKey() bool`

HasTagKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


