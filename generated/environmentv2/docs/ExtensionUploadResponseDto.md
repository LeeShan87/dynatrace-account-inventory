# ExtensionUploadResponseDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetsInfo** | Pointer to [**[]AssetInfo**](AssetInfo.md) | Information about extension assets included | [optional] 
**Author** | Pointer to [**AuthorDto**](AuthorDto.md) |  | [optional] 
**DataSources** | Pointer to **[]string** | Data sources that extension uses to gather data | [optional] 
**ExtensionName** | Pointer to **string** | Extension name | [optional] 
**FeatureSets** | Pointer to **[]string** | Available feature sets | [optional] 
**FeatureSetsDetails** | Pointer to [**map[string]FeatureSetDetails**](FeatureSetDetails.md) | Details of feature sets | [optional] 
**FileHash** | Pointer to **string** | SHA-256 hash of uploaded Extension file | [optional] 
**MinDynatraceVersion** | Pointer to **string** | Minimal Dynatrace version that works with the extension | [optional] 
**MinEECVersion** | Pointer to **string** | Minimal Extension Execution Controller version that works with the extension | [optional] 
**Variables** | Pointer to **[]string** | Custom variables used in extension configuration | [optional] 
**Version** | Pointer to **string** | Extension version | [optional] 

## Methods

### NewExtensionUploadResponseDto

`func NewExtensionUploadResponseDto() *ExtensionUploadResponseDto`

NewExtensionUploadResponseDto instantiates a new ExtensionUploadResponseDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionUploadResponseDtoWithDefaults

`func NewExtensionUploadResponseDtoWithDefaults() *ExtensionUploadResponseDto`

NewExtensionUploadResponseDtoWithDefaults instantiates a new ExtensionUploadResponseDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetsInfo

`func (o *ExtensionUploadResponseDto) GetAssetsInfo() []AssetInfo`

GetAssetsInfo returns the AssetsInfo field if non-nil, zero value otherwise.

### GetAssetsInfoOk

`func (o *ExtensionUploadResponseDto) GetAssetsInfoOk() (*[]AssetInfo, bool)`

GetAssetsInfoOk returns a tuple with the AssetsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsInfo

`func (o *ExtensionUploadResponseDto) SetAssetsInfo(v []AssetInfo)`

SetAssetsInfo sets AssetsInfo field to given value.

### HasAssetsInfo

`func (o *ExtensionUploadResponseDto) HasAssetsInfo() bool`

HasAssetsInfo returns a boolean if a field has been set.

### GetAuthor

`func (o *ExtensionUploadResponseDto) GetAuthor() AuthorDto`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ExtensionUploadResponseDto) GetAuthorOk() (*AuthorDto, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ExtensionUploadResponseDto) SetAuthor(v AuthorDto)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ExtensionUploadResponseDto) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDataSources

`func (o *ExtensionUploadResponseDto) GetDataSources() []string`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *ExtensionUploadResponseDto) GetDataSourcesOk() (*[]string, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *ExtensionUploadResponseDto) SetDataSources(v []string)`

SetDataSources sets DataSources field to given value.

### HasDataSources

`func (o *ExtensionUploadResponseDto) HasDataSources() bool`

HasDataSources returns a boolean if a field has been set.

### GetExtensionName

`func (o *ExtensionUploadResponseDto) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *ExtensionUploadResponseDto) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *ExtensionUploadResponseDto) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *ExtensionUploadResponseDto) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### GetFeatureSets

`func (o *ExtensionUploadResponseDto) GetFeatureSets() []string`

GetFeatureSets returns the FeatureSets field if non-nil, zero value otherwise.

### GetFeatureSetsOk

`func (o *ExtensionUploadResponseDto) GetFeatureSetsOk() (*[]string, bool)`

GetFeatureSetsOk returns a tuple with the FeatureSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSets

`func (o *ExtensionUploadResponseDto) SetFeatureSets(v []string)`

SetFeatureSets sets FeatureSets field to given value.

### HasFeatureSets

`func (o *ExtensionUploadResponseDto) HasFeatureSets() bool`

HasFeatureSets returns a boolean if a field has been set.

### GetFeatureSetsDetails

`func (o *ExtensionUploadResponseDto) GetFeatureSetsDetails() map[string]FeatureSetDetails`

GetFeatureSetsDetails returns the FeatureSetsDetails field if non-nil, zero value otherwise.

### GetFeatureSetsDetailsOk

`func (o *ExtensionUploadResponseDto) GetFeatureSetsDetailsOk() (*map[string]FeatureSetDetails, bool)`

GetFeatureSetsDetailsOk returns a tuple with the FeatureSetsDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSetsDetails

`func (o *ExtensionUploadResponseDto) SetFeatureSetsDetails(v map[string]FeatureSetDetails)`

SetFeatureSetsDetails sets FeatureSetsDetails field to given value.

### HasFeatureSetsDetails

`func (o *ExtensionUploadResponseDto) HasFeatureSetsDetails() bool`

HasFeatureSetsDetails returns a boolean if a field has been set.

### GetFileHash

`func (o *ExtensionUploadResponseDto) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *ExtensionUploadResponseDto) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *ExtensionUploadResponseDto) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.

### HasFileHash

`func (o *ExtensionUploadResponseDto) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.

### GetMinDynatraceVersion

`func (o *ExtensionUploadResponseDto) GetMinDynatraceVersion() string`

GetMinDynatraceVersion returns the MinDynatraceVersion field if non-nil, zero value otherwise.

### GetMinDynatraceVersionOk

`func (o *ExtensionUploadResponseDto) GetMinDynatraceVersionOk() (*string, bool)`

GetMinDynatraceVersionOk returns a tuple with the MinDynatraceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinDynatraceVersion

`func (o *ExtensionUploadResponseDto) SetMinDynatraceVersion(v string)`

SetMinDynatraceVersion sets MinDynatraceVersion field to given value.

### HasMinDynatraceVersion

`func (o *ExtensionUploadResponseDto) HasMinDynatraceVersion() bool`

HasMinDynatraceVersion returns a boolean if a field has been set.

### GetMinEECVersion

`func (o *ExtensionUploadResponseDto) GetMinEECVersion() string`

GetMinEECVersion returns the MinEECVersion field if non-nil, zero value otherwise.

### GetMinEECVersionOk

`func (o *ExtensionUploadResponseDto) GetMinEECVersionOk() (*string, bool)`

GetMinEECVersionOk returns a tuple with the MinEECVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinEECVersion

`func (o *ExtensionUploadResponseDto) SetMinEECVersion(v string)`

SetMinEECVersion sets MinEECVersion field to given value.

### HasMinEECVersion

`func (o *ExtensionUploadResponseDto) HasMinEECVersion() bool`

HasMinEECVersion returns a boolean if a field has been set.

### GetVariables

`func (o *ExtensionUploadResponseDto) GetVariables() []string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ExtensionUploadResponseDto) GetVariablesOk() (*[]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ExtensionUploadResponseDto) SetVariables(v []string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ExtensionUploadResponseDto) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVersion

`func (o *ExtensionUploadResponseDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExtensionUploadResponseDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExtensionUploadResponseDto) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExtensionUploadResponseDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


