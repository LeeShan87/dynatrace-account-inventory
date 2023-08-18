# Extension2Details

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Distributed** | Pointer to **bool** | Whether this extension is available in the central hub catalog. | [optional] 
**ExtensionName** | Pointer to **string** | Fully qualified name of the extension. | [optional] 
**RecommendedCatalogVersion** | Pointer to **NullableString** | Recommended version of this extension to use. This is the latest compatible published release. | [optional] 
**Releases** | Pointer to [**[]ExtensionRelease**](ExtensionRelease.md) | Releases for the extension. | [optional] 

## Methods

### NewExtension2Details

`func NewExtension2Details() *Extension2Details`

NewExtension2Details instantiates a new Extension2Details object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtension2DetailsWithDefaults

`func NewExtension2DetailsWithDefaults() *Extension2Details`

NewExtension2DetailsWithDefaults instantiates a new Extension2Details object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistributed

`func (o *Extension2Details) GetDistributed() bool`

GetDistributed returns the Distributed field if non-nil, zero value otherwise.

### GetDistributedOk

`func (o *Extension2Details) GetDistributedOk() (*bool, bool)`

GetDistributedOk returns a tuple with the Distributed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributed

`func (o *Extension2Details) SetDistributed(v bool)`

SetDistributed sets Distributed field to given value.

### HasDistributed

`func (o *Extension2Details) HasDistributed() bool`

HasDistributed returns a boolean if a field has been set.

### GetExtensionName

`func (o *Extension2Details) GetExtensionName() string`

GetExtensionName returns the ExtensionName field if non-nil, zero value otherwise.

### GetExtensionNameOk

`func (o *Extension2Details) GetExtensionNameOk() (*string, bool)`

GetExtensionNameOk returns a tuple with the ExtensionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionName

`func (o *Extension2Details) SetExtensionName(v string)`

SetExtensionName sets ExtensionName field to given value.

### HasExtensionName

`func (o *Extension2Details) HasExtensionName() bool`

HasExtensionName returns a boolean if a field has been set.

### GetRecommendedCatalogVersion

`func (o *Extension2Details) GetRecommendedCatalogVersion() string`

GetRecommendedCatalogVersion returns the RecommendedCatalogVersion field if non-nil, zero value otherwise.

### GetRecommendedCatalogVersionOk

`func (o *Extension2Details) GetRecommendedCatalogVersionOk() (*string, bool)`

GetRecommendedCatalogVersionOk returns a tuple with the RecommendedCatalogVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendedCatalogVersion

`func (o *Extension2Details) SetRecommendedCatalogVersion(v string)`

SetRecommendedCatalogVersion sets RecommendedCatalogVersion field to given value.

### HasRecommendedCatalogVersion

`func (o *Extension2Details) HasRecommendedCatalogVersion() bool`

HasRecommendedCatalogVersion returns a boolean if a field has been set.

### SetRecommendedCatalogVersionNil

`func (o *Extension2Details) SetRecommendedCatalogVersionNil(b bool)`

 SetRecommendedCatalogVersionNil sets the value for RecommendedCatalogVersion to be an explicit nil

### UnsetRecommendedCatalogVersion
`func (o *Extension2Details) UnsetRecommendedCatalogVersion()`

UnsetRecommendedCatalogVersion ensures that no value is present for RecommendedCatalogVersion, not even an explicit nil
### GetReleases

`func (o *Extension2Details) GetReleases() []ExtensionRelease`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *Extension2Details) GetReleasesOk() (*[]ExtensionRelease, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *Extension2Details) SetReleases(v []ExtensionRelease)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *Extension2Details) HasReleases() bool`

HasReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


