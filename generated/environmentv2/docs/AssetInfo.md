# AssetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetType** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewAssetInfo

`func NewAssetInfo() *AssetInfo`

NewAssetInfo instantiates a new AssetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetInfoWithDefaults

`func NewAssetInfoWithDefaults() *AssetInfo`

NewAssetInfoWithDefaults instantiates a new AssetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *AssetInfo) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AssetInfo) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AssetInfo) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *AssetInfo) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetCount

`func (o *AssetInfo) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *AssetInfo) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *AssetInfo) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *AssetInfo) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


