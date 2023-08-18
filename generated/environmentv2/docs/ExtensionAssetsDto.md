# ExtensionAssetsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assets** | Pointer to [**[]AssetInfoDto**](AssetInfoDto.md) | The list of the imported assets. | [optional] 
**Errors** | Pointer to **[]string** | List of errors during asset import | [optional] 
**Status** | Pointer to **string** | The status of the assets list. | [optional] 
**Version** | Pointer to **string** | Extension version | [optional] 

## Methods

### NewExtensionAssetsDto

`func NewExtensionAssetsDto() *ExtensionAssetsDto`

NewExtensionAssetsDto instantiates a new ExtensionAssetsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionAssetsDtoWithDefaults

`func NewExtensionAssetsDtoWithDefaults() *ExtensionAssetsDto`

NewExtensionAssetsDtoWithDefaults instantiates a new ExtensionAssetsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssets

`func (o *ExtensionAssetsDto) GetAssets() []AssetInfoDto`

GetAssets returns the Assets field if non-nil, zero value otherwise.

### GetAssetsOk

`func (o *ExtensionAssetsDto) GetAssetsOk() (*[]AssetInfoDto, bool)`

GetAssetsOk returns a tuple with the Assets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssets

`func (o *ExtensionAssetsDto) SetAssets(v []AssetInfoDto)`

SetAssets sets Assets field to given value.

### HasAssets

`func (o *ExtensionAssetsDto) HasAssets() bool`

HasAssets returns a boolean if a field has been set.

### GetErrors

`func (o *ExtensionAssetsDto) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ExtensionAssetsDto) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ExtensionAssetsDto) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ExtensionAssetsDto) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *ExtensionAssetsDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExtensionAssetsDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExtensionAssetsDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExtensionAssetsDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersion

`func (o *ExtensionAssetsDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExtensionAssetsDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExtensionAssetsDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExtensionAssetsDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


