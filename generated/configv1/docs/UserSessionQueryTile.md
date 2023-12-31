# UserSessionQueryTile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomName** | **string** | The name of the tile, set by user. | 
**Limit** | Pointer to **int32** | The limit of the results, if not set will use the default value of the system | [optional] 
**Query** | **string** | A [user session query](https://dt-url.net/dtusql) executed by the tile. | 
**TimeFrameShift** | Pointer to **string** | The comparison timeframe of the query.    If specified, you additionally get the results of the same query with the specified time shift. | [optional] 
**Type** | **string** | The visualization of the tile. | 
**VisualizationConfig** | Pointer to [**UserSessionQueryTileConfiguration**](UserSessionQueryTileConfiguration.md) |  | [optional] 

## Methods

### NewUserSessionQueryTile

`func NewUserSessionQueryTile(customName string, query string, type_ string, ) *UserSessionQueryTile`

NewUserSessionQueryTile instantiates a new UserSessionQueryTile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSessionQueryTileWithDefaults

`func NewUserSessionQueryTileWithDefaults() *UserSessionQueryTile`

NewUserSessionQueryTileWithDefaults instantiates a new UserSessionQueryTile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomName

`func (o *UserSessionQueryTile) GetCustomName() string`

GetCustomName returns the CustomName field if non-nil, zero value otherwise.

### GetCustomNameOk

`func (o *UserSessionQueryTile) GetCustomNameOk() (*string, bool)`

GetCustomNameOk returns a tuple with the CustomName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomName

`func (o *UserSessionQueryTile) SetCustomName(v string)`

SetCustomName sets CustomName field to given value.


### GetLimit

`func (o *UserSessionQueryTile) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *UserSessionQueryTile) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *UserSessionQueryTile) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *UserSessionQueryTile) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetQuery

`func (o *UserSessionQueryTile) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *UserSessionQueryTile) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *UserSessionQueryTile) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTimeFrameShift

`func (o *UserSessionQueryTile) GetTimeFrameShift() string`

GetTimeFrameShift returns the TimeFrameShift field if non-nil, zero value otherwise.

### GetTimeFrameShiftOk

`func (o *UserSessionQueryTile) GetTimeFrameShiftOk() (*string, bool)`

GetTimeFrameShiftOk returns a tuple with the TimeFrameShift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFrameShift

`func (o *UserSessionQueryTile) SetTimeFrameShift(v string)`

SetTimeFrameShift sets TimeFrameShift field to given value.

### HasTimeFrameShift

`func (o *UserSessionQueryTile) HasTimeFrameShift() bool`

HasTimeFrameShift returns a boolean if a field has been set.

### GetType

`func (o *UserSessionQueryTile) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserSessionQueryTile) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserSessionQueryTile) SetType(v string)`

SetType sets Type field to given value.


### GetVisualizationConfig

`func (o *UserSessionQueryTile) GetVisualizationConfig() UserSessionQueryTileConfiguration`

GetVisualizationConfig returns the VisualizationConfig field if non-nil, zero value otherwise.

### GetVisualizationConfigOk

`func (o *UserSessionQueryTile) GetVisualizationConfigOk() (*UserSessionQueryTileConfiguration, bool)`

GetVisualizationConfigOk returns a tuple with the VisualizationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizationConfig

`func (o *UserSessionQueryTile) SetVisualizationConfig(v UserSessionQueryTileConfiguration)`

SetVisualizationConfig sets VisualizationConfig field to given value.

### HasVisualizationConfig

`func (o *UserSessionQueryTile) HasVisualizationConfig() bool`

HasVisualizationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


