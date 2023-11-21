# ExtensionRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Represents whether this version is active version | [optional] 
**ArtifactSha256** | Pointer to **NullableString** | Sha256 hash for the distributed extension. | [optional] 
**AssetsInfo** | Pointer to [**[]AssetInfo**](AssetInfo.md) | Assets types and its count | [optional] 
**ConfiguredFeatureSets** | Pointer to **[]string** | Configured feature sets for an installed release | [optional] 
**DataSources** | Pointer to **[]string** | Available data sources for the given release | [optional] 
**Distributed** | Pointer to **bool** | Represents whether the release is distributed | [optional] 
**FeatureSets** | Pointer to [**map[string]FeatureSetDetails**](FeatureSetDetails.md) | Feature sets contained in the given release | [optional] 
**MinClusterVersion** | Pointer to **NullableInt32** | Minimum cluster version for the release | [optional] 
**Registered** | Pointer to **bool** | Represents whether extension is already registered | [optional] 
**ReleaseNotes** | Pointer to **NullableString** | Release notes for the extension. | [optional] 
**Unpublished** | Pointer to **bool** | Represents whether the extension is unpublished. | [optional] 
**UnpublishedDescription** | Pointer to **NullableString** | The description why the extension was unpublished. | [optional] 
**UnpublishedSeverity** | Pointer to **NullableInt32** | The severity of unpublished extension. 5 indicates an error state | [optional] 
**Version** | Pointer to **string** | Version number of the extension. | [optional] 

## Methods

### NewExtensionRelease

`func NewExtensionRelease() *ExtensionRelease`

NewExtensionRelease instantiates a new ExtensionRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionReleaseWithDefaults

`func NewExtensionReleaseWithDefaults() *ExtensionRelease`

NewExtensionReleaseWithDefaults instantiates a new ExtensionRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ExtensionRelease) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ExtensionRelease) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ExtensionRelease) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ExtensionRelease) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetArtifactSha256

`func (o *ExtensionRelease) GetArtifactSha256() string`

GetArtifactSha256 returns the ArtifactSha256 field if non-nil, zero value otherwise.

### GetArtifactSha256Ok

`func (o *ExtensionRelease) GetArtifactSha256Ok() (*string, bool)`

GetArtifactSha256Ok returns a tuple with the ArtifactSha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactSha256

`func (o *ExtensionRelease) SetArtifactSha256(v string)`

SetArtifactSha256 sets ArtifactSha256 field to given value.

### HasArtifactSha256

`func (o *ExtensionRelease) HasArtifactSha256() bool`

HasArtifactSha256 returns a boolean if a field has been set.

### SetArtifactSha256Nil

`func (o *ExtensionRelease) SetArtifactSha256Nil(b bool)`

 SetArtifactSha256Nil sets the value for ArtifactSha256 to be an explicit nil

### UnsetArtifactSha256
`func (o *ExtensionRelease) UnsetArtifactSha256()`

UnsetArtifactSha256 ensures that no value is present for ArtifactSha256, not even an explicit nil
### GetAssetsInfo

`func (o *ExtensionRelease) GetAssetsInfo() []AssetInfo`

GetAssetsInfo returns the AssetsInfo field if non-nil, zero value otherwise.

### GetAssetsInfoOk

`func (o *ExtensionRelease) GetAssetsInfoOk() (*[]AssetInfo, bool)`

GetAssetsInfoOk returns a tuple with the AssetsInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetsInfo

`func (o *ExtensionRelease) SetAssetsInfo(v []AssetInfo)`

SetAssetsInfo sets AssetsInfo field to given value.

### HasAssetsInfo

`func (o *ExtensionRelease) HasAssetsInfo() bool`

HasAssetsInfo returns a boolean if a field has been set.

### GetConfiguredFeatureSets

`func (o *ExtensionRelease) GetConfiguredFeatureSets() []string`

GetConfiguredFeatureSets returns the ConfiguredFeatureSets field if non-nil, zero value otherwise.

### GetConfiguredFeatureSetsOk

`func (o *ExtensionRelease) GetConfiguredFeatureSetsOk() (*[]string, bool)`

GetConfiguredFeatureSetsOk returns a tuple with the ConfiguredFeatureSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredFeatureSets

`func (o *ExtensionRelease) SetConfiguredFeatureSets(v []string)`

SetConfiguredFeatureSets sets ConfiguredFeatureSets field to given value.

### HasConfiguredFeatureSets

`func (o *ExtensionRelease) HasConfiguredFeatureSets() bool`

HasConfiguredFeatureSets returns a boolean if a field has been set.

### GetDataSources

`func (o *ExtensionRelease) GetDataSources() []string`

GetDataSources returns the DataSources field if non-nil, zero value otherwise.

### GetDataSourcesOk

`func (o *ExtensionRelease) GetDataSourcesOk() (*[]string, bool)`

GetDataSourcesOk returns a tuple with the DataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSources

`func (o *ExtensionRelease) SetDataSources(v []string)`

SetDataSources sets DataSources field to given value.

### HasDataSources

`func (o *ExtensionRelease) HasDataSources() bool`

HasDataSources returns a boolean if a field has been set.

### GetDistributed

`func (o *ExtensionRelease) GetDistributed() bool`

GetDistributed returns the Distributed field if non-nil, zero value otherwise.

### GetDistributedOk

`func (o *ExtensionRelease) GetDistributedOk() (*bool, bool)`

GetDistributedOk returns a tuple with the Distributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributed

`func (o *ExtensionRelease) SetDistributed(v bool)`

SetDistributed sets Distributed field to given value.

### HasDistributed

`func (o *ExtensionRelease) HasDistributed() bool`

HasDistributed returns a boolean if a field has been set.

### GetFeatureSets

`func (o *ExtensionRelease) GetFeatureSets() map[string]FeatureSetDetails`

GetFeatureSets returns the FeatureSets field if non-nil, zero value otherwise.

### GetFeatureSetsOk

`func (o *ExtensionRelease) GetFeatureSetsOk() (*map[string]FeatureSetDetails, bool)`

GetFeatureSetsOk returns a tuple with the FeatureSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureSets

`func (o *ExtensionRelease) SetFeatureSets(v map[string]FeatureSetDetails)`

SetFeatureSets sets FeatureSets field to given value.

### HasFeatureSets

`func (o *ExtensionRelease) HasFeatureSets() bool`

HasFeatureSets returns a boolean if a field has been set.

### GetMinClusterVersion

`func (o *ExtensionRelease) GetMinClusterVersion() int32`

GetMinClusterVersion returns the MinClusterVersion field if non-nil, zero value otherwise.

### GetMinClusterVersionOk

`func (o *ExtensionRelease) GetMinClusterVersionOk() (*int32, bool)`

GetMinClusterVersionOk returns a tuple with the MinClusterVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinClusterVersion

`func (o *ExtensionRelease) SetMinClusterVersion(v int32)`

SetMinClusterVersion sets MinClusterVersion field to given value.

### HasMinClusterVersion

`func (o *ExtensionRelease) HasMinClusterVersion() bool`

HasMinClusterVersion returns a boolean if a field has been set.

### SetMinClusterVersionNil

`func (o *ExtensionRelease) SetMinClusterVersionNil(b bool)`

 SetMinClusterVersionNil sets the value for MinClusterVersion to be an explicit nil

### UnsetMinClusterVersion
`func (o *ExtensionRelease) UnsetMinClusterVersion()`

UnsetMinClusterVersion ensures that no value is present for MinClusterVersion, not even an explicit nil
### GetRegistered

`func (o *ExtensionRelease) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *ExtensionRelease) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *ExtensionRelease) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *ExtensionRelease) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.

### GetReleaseNotes

`func (o *ExtensionRelease) GetReleaseNotes() string`

GetReleaseNotes returns the ReleaseNotes field if non-nil, zero value otherwise.

### GetReleaseNotesOk

`func (o *ExtensionRelease) GetReleaseNotesOk() (*string, bool)`

GetReleaseNotesOk returns a tuple with the ReleaseNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseNotes

`func (o *ExtensionRelease) SetReleaseNotes(v string)`

SetReleaseNotes sets ReleaseNotes field to given value.

### HasReleaseNotes

`func (o *ExtensionRelease) HasReleaseNotes() bool`

HasReleaseNotes returns a boolean if a field has been set.

### SetReleaseNotesNil

`func (o *ExtensionRelease) SetReleaseNotesNil(b bool)`

 SetReleaseNotesNil sets the value for ReleaseNotes to be an explicit nil

### UnsetReleaseNotes
`func (o *ExtensionRelease) UnsetReleaseNotes()`

UnsetReleaseNotes ensures that no value is present for ReleaseNotes, not even an explicit nil
### GetUnpublished

`func (o *ExtensionRelease) GetUnpublished() bool`

GetUnpublished returns the Unpublished field if non-nil, zero value otherwise.

### GetUnpublishedOk

`func (o *ExtensionRelease) GetUnpublishedOk() (*bool, bool)`

GetUnpublishedOk returns a tuple with the Unpublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublished

`func (o *ExtensionRelease) SetUnpublished(v bool)`

SetUnpublished sets Unpublished field to given value.

### HasUnpublished

`func (o *ExtensionRelease) HasUnpublished() bool`

HasUnpublished returns a boolean if a field has been set.

### GetUnpublishedDescription

`func (o *ExtensionRelease) GetUnpublishedDescription() string`

GetUnpublishedDescription returns the UnpublishedDescription field if non-nil, zero value otherwise.

### GetUnpublishedDescriptionOk

`func (o *ExtensionRelease) GetUnpublishedDescriptionOk() (*string, bool)`

GetUnpublishedDescriptionOk returns a tuple with the UnpublishedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublishedDescription

`func (o *ExtensionRelease) SetUnpublishedDescription(v string)`

SetUnpublishedDescription sets UnpublishedDescription field to given value.

### HasUnpublishedDescription

`func (o *ExtensionRelease) HasUnpublishedDescription() bool`

HasUnpublishedDescription returns a boolean if a field has been set.

### SetUnpublishedDescriptionNil

`func (o *ExtensionRelease) SetUnpublishedDescriptionNil(b bool)`

 SetUnpublishedDescriptionNil sets the value for UnpublishedDescription to be an explicit nil

### UnsetUnpublishedDescription
`func (o *ExtensionRelease) UnsetUnpublishedDescription()`

UnsetUnpublishedDescription ensures that no value is present for UnpublishedDescription, not even an explicit nil
### GetUnpublishedSeverity

`func (o *ExtensionRelease) GetUnpublishedSeverity() int32`

GetUnpublishedSeverity returns the UnpublishedSeverity field if non-nil, zero value otherwise.

### GetUnpublishedSeverityOk

`func (o *ExtensionRelease) GetUnpublishedSeverityOk() (*int32, bool)`

GetUnpublishedSeverityOk returns a tuple with the UnpublishedSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnpublishedSeverity

`func (o *ExtensionRelease) SetUnpublishedSeverity(v int32)`

SetUnpublishedSeverity sets UnpublishedSeverity field to given value.

### HasUnpublishedSeverity

`func (o *ExtensionRelease) HasUnpublishedSeverity() bool`

HasUnpublishedSeverity returns a boolean if a field has been set.

### SetUnpublishedSeverityNil

`func (o *ExtensionRelease) SetUnpublishedSeverityNil(b bool)`

 SetUnpublishedSeverityNil sets the value for UnpublishedSeverity to be an explicit nil

### UnsetUnpublishedSeverity
`func (o *ExtensionRelease) UnsetUnpublishedSeverity()`

UnsetUnpublishedSeverity ensures that no value is present for UnpublishedSeverity, not even an explicit nil
### GetVersion

`func (o *ExtensionRelease) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ExtensionRelease) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ExtensionRelease) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ExtensionRelease) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


