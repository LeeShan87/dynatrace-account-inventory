# AssetInfoDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetSchemaDetails** | Pointer to [**AssetSchemaDetailsDto**](AssetSchemaDetailsDto.md) |  | [optional] 
**DisplayName** | Pointer to **string** | User-friendly name of the asset. | [optional] 
**Id** | Pointer to **string** | ID of the asset. Identifies the asset in REST API and/or UI (where applicable). | [optional] 
**Type** | Pointer to **string** | The type of the asset. | [optional] 

## Methods

### NewAssetInfoDto

`func NewAssetInfoDto() *AssetInfoDto`

NewAssetInfoDto instantiates a new AssetInfoDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetInfoDtoWithDefaults

`func NewAssetInfoDtoWithDefaults() *AssetInfoDto`

NewAssetInfoDtoWithDefaults instantiates a new AssetInfoDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetSchemaDetails

`func (o *AssetInfoDto) GetAssetSchemaDetails() AssetSchemaDetailsDto`

GetAssetSchemaDetails returns the AssetSchemaDetails field if non-nil, zero value otherwise.

### GetAssetSchemaDetailsOk

`func (o *AssetInfoDto) GetAssetSchemaDetailsOk() (*AssetSchemaDetailsDto, bool)`

GetAssetSchemaDetailsOk returns a tuple with the AssetSchemaDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetSchemaDetails

`func (o *AssetInfoDto) SetAssetSchemaDetails(v AssetSchemaDetailsDto)`

SetAssetSchemaDetails sets AssetSchemaDetails field to given value.

### HasAssetSchemaDetails

`func (o *AssetInfoDto) HasAssetSchemaDetails() bool`

HasAssetSchemaDetails returns a boolean if a field has been set.

### GetDisplayName

`func (o *AssetInfoDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AssetInfoDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AssetInfoDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AssetInfoDto) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *AssetInfoDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetInfoDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetInfoDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetInfoDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AssetInfoDto) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssetInfoDto) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssetInfoDto) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AssetInfoDto) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


