# Tile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bounds** | [**TileBounds**](TileBounds.md) |  | 
**Configured** | Pointer to **bool** | The tile is configured and ready to use (&#x60;true&#x60;) or just placed on the dashboard (&#x60;false&#x60;). | [optional] 
**IsAutoRefreshDisabled** | Pointer to **bool** | The tile auto refresh is disabled. Only works for certain tile types. | [optional] 
**Name** | **string** | The name of the tile. | 
**NameSize** | Pointer to **string** | The size of the tile name. Default value is null. | [optional] 
**TileFilter** | Pointer to [**TileFilter**](TileFilter.md) |  | [optional] 
**TileType** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;DATA_EXPLORER&#x60; -&gt; DataExplorerTile  * &#x60;CUSTOM_CHARTING&#x60; -&gt; CustomChartingTile  * &#x60;DTAQL&#x60; -&gt; UserSessionQueryTile  * &#x60;MARKDOWN&#x60; -&gt; MarkdownTile  * &#x60;IMAGE&#x60; -&gt; ImageTile  * &#x60;HOSTS&#x60; -&gt; FilterableEntityTile  * &#x60;APPLICATIONS&#x60; -&gt; FilterableEntityTile  * &#x60;SERVICES&#x60; -&gt; FilterableEntityTile  * &#x60;DATABASES_OVERVIEW&#x60; -&gt; FilterableEntityTile  * &#x60;SYNTHETIC_TESTS&#x60; -&gt; FilterableEntityTile  * &#x60;APPLICATION_WORLDMAP&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;RESOURCES&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;THIRD_PARTY_MOST_ACTIVE&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;UEM_CONVERSIONS_PER_GOAL&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;HOST&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;PROCESS_GROUPS_ONE&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;SYNTHETIC_SINGLE_WEBCHECK&#x60; -&gt; SyntheticSingleWebcheckTile  * &#x60;APPLICATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;VIRTUALIZATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;AWS&#x60; -&gt; AssignedEntitiesTile  * &#x60;SERVICE_VERSATILE&#x60; -&gt; AssignedEntitiesTile  * &#x60;SESSION_METRICS&#x60; -&gt; AssignedEntitiesTile  * &#x60;USERS&#x60; -&gt; AssignedEntitiesTile  * &#x60;UEM_KEY_USER_ACTIONS&#x60; -&gt; AssignedEntitiesTile  * &#x60;BOUNCE_RATE&#x60; -&gt; AssignedEntitiesTile  * &#x60;UEM_CONVERSIONS_OVERALL&#x60; -&gt; AssignedEntitiesTile  * &#x60;UEM_JSERRORS_OVERALL&#x60; -&gt; AssignedEntitiesTile  * &#x60;MOBILE_APPLICATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;SYNTHETIC_SINGLE_EXT_TEST&#x60; -&gt; AssignedEntitiesTile  * &#x60;SYNTHETIC_HTTP_MONITOR&#x60; -&gt; AssignedEntitiesTile  * &#x60;DATABASE&#x60; -&gt; AssignedEntitiesTile  * &#x60;CUSTOM_APPLICATION&#x60; -&gt; AssignedEntitiesTile  * &#x60;APPLICATION_METHOD&#x60; -&gt; AssignedEntitiesTile  * &#x60;LOG_ANALYTICS&#x60; -&gt; AssignedEntitiesTile  * &#x60;OPENSTACK&#x60; -&gt; AssignedEntitiesTile  * &#x60;OPENSTACK_PROJECT&#x60; -&gt; AssignedEntitiesTile  * &#x60;OPENSTACK_AV_ZONE&#x60; -&gt; AssignedEntitiesTile  * &#x60;DEVICE_APPLICATION_METHOD&#x60; -&gt; AssignedEntitiesTile  * &#x60;DEM_KEY_USER_ACTION&#x60; -&gt; AssignedEntitiesTile  * &#x60;SLO&#x60; -&gt; AssignedEntitiesWithMetricTile  * &#x60;SCALABLE_LIST&#x60; -&gt; ScalableListTile  * &#x60;HEADER&#x60; -&gt; Tile  * &#x60;OPEN_PROBLEMS&#x60; -&gt; ProblemTile  * &#x60;PURE_MODEL&#x60; -&gt; Tile  * &#x60;DOCKER&#x60; -&gt; Tile  * &#x60;NETWORK_MEDIUM&#x60; -&gt; Tile  * &#x60;APPLICATIONS_MOST_ACTIVE&#x60; -&gt; Tile  * &#x60;NETWORK&#x60; -&gt; Tile  * &#x60;UEM_ACTIVE_SESSIONS&#x60; -&gt; Tile  * &#x60;DCRUM_SERVICES&#x60; -&gt; Tile   | 

## Methods

### NewTile

`func NewTile(bounds TileBounds, name string, tileType string, ) *Tile`

NewTile instantiates a new Tile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTileWithDefaults

`func NewTileWithDefaults() *Tile`

NewTileWithDefaults instantiates a new Tile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBounds

`func (o *Tile) GetBounds() TileBounds`

GetBounds returns the Bounds field if non-nil, zero value otherwise.

### GetBoundsOk

`func (o *Tile) GetBoundsOk() (*TileBounds, bool)`

GetBoundsOk returns a tuple with the Bounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBounds

`func (o *Tile) SetBounds(v TileBounds)`

SetBounds sets Bounds field to given value.


### GetConfigured

`func (o *Tile) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *Tile) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *Tile) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *Tile) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetIsAutoRefreshDisabled

`func (o *Tile) GetIsAutoRefreshDisabled() bool`

GetIsAutoRefreshDisabled returns the IsAutoRefreshDisabled field if non-nil, zero value otherwise.

### GetIsAutoRefreshDisabledOk

`func (o *Tile) GetIsAutoRefreshDisabledOk() (*bool, bool)`

GetIsAutoRefreshDisabledOk returns a tuple with the IsAutoRefreshDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoRefreshDisabled

`func (o *Tile) SetIsAutoRefreshDisabled(v bool)`

SetIsAutoRefreshDisabled sets IsAutoRefreshDisabled field to given value.

### HasIsAutoRefreshDisabled

`func (o *Tile) HasIsAutoRefreshDisabled() bool`

HasIsAutoRefreshDisabled returns a boolean if a field has been set.

### GetName

`func (o *Tile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tile) SetName(v string)`

SetName sets Name field to given value.


### GetNameSize

`func (o *Tile) GetNameSize() string`

GetNameSize returns the NameSize field if non-nil, zero value otherwise.

### GetNameSizeOk

`func (o *Tile) GetNameSizeOk() (*string, bool)`

GetNameSizeOk returns a tuple with the NameSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSize

`func (o *Tile) SetNameSize(v string)`

SetNameSize sets NameSize field to given value.

### HasNameSize

`func (o *Tile) HasNameSize() bool`

HasNameSize returns a boolean if a field has been set.

### GetTileFilter

`func (o *Tile) GetTileFilter() TileFilter`

GetTileFilter returns the TileFilter field if non-nil, zero value otherwise.

### GetTileFilterOk

`func (o *Tile) GetTileFilterOk() (*TileFilter, bool)`

GetTileFilterOk returns a tuple with the TileFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTileFilter

`func (o *Tile) SetTileFilter(v TileFilter)`

SetTileFilter sets TileFilter field to given value.

### HasTileFilter

`func (o *Tile) HasTileFilter() bool`

HasTileFilter returns a boolean if a field has been set.

### GetTileType

`func (o *Tile) GetTileType() string`

GetTileType returns the TileType field if non-nil, zero value otherwise.

### GetTileTypeOk

`func (o *Tile) GetTileTypeOk() (*string, bool)`

GetTileTypeOk returns a tuple with the TileType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTileType

`func (o *Tile) SetTileType(v string)`

SetTileType sets TileType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


