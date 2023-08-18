# YAxis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultAxis** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Max** | **string** |  | 
**Min** | **string** |  | 
**Position** | **string** |  | 
**QueryIds** | **[]string** |  | 
**Visible** | Pointer to **bool** |  | [optional] 

## Methods

### NewYAxis

`func NewYAxis(max string, min string, position string, queryIds []string, ) *YAxis`

NewYAxis instantiates a new YAxis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYAxisWithDefaults

`func NewYAxisWithDefaults() *YAxis`

NewYAxisWithDefaults instantiates a new YAxis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultAxis

`func (o *YAxis) GetDefaultAxis() bool`

GetDefaultAxis returns the DefaultAxis field if non-nil, zero value otherwise.

### GetDefaultAxisOk

`func (o *YAxis) GetDefaultAxisOk() (*bool, bool)`

GetDefaultAxisOk returns a tuple with the DefaultAxis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAxis

`func (o *YAxis) SetDefaultAxis(v bool)`

SetDefaultAxis sets DefaultAxis field to given value.

### HasDefaultAxis

`func (o *YAxis) HasDefaultAxis() bool`

HasDefaultAxis returns a boolean if a field has been set.

### GetDisplayName

`func (o *YAxis) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *YAxis) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *YAxis) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *YAxis) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMax

`func (o *YAxis) GetMax() string`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *YAxis) GetMaxOk() (*string, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *YAxis) SetMax(v string)`

SetMax sets Max field to given value.


### GetMin

`func (o *YAxis) GetMin() string`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *YAxis) GetMinOk() (*string, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *YAxis) SetMin(v string)`

SetMin sets Min field to given value.


### GetPosition

`func (o *YAxis) GetPosition() string`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *YAxis) GetPositionOk() (*string, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *YAxis) SetPosition(v string)`

SetPosition sets Position field to given value.


### GetQueryIds

`func (o *YAxis) GetQueryIds() []string`

GetQueryIds returns the QueryIds field if non-nil, zero value otherwise.

### GetQueryIdsOk

`func (o *YAxis) GetQueryIdsOk() (*[]string, bool)`

GetQueryIdsOk returns a tuple with the QueryIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryIds

`func (o *YAxis) SetQueryIds(v []string)`

SetQueryIds sets QueryIds field to given value.


### GetVisible

`func (o *YAxis) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *YAxis) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *YAxis) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *YAxis) HasVisible() bool`

HasVisible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


