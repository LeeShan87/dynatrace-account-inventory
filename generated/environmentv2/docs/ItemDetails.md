# ItemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorLogo** | Pointer to **NullableString** | Url for the author&#39;s logo. | [optional] 
**AuthorName** | Pointer to **NullableString** | Name of the author of the item. | [optional] 
**ClusterCompatible** | Pointer to **NullableBool** | Checks if the item is compatible with the cluster version. | [optional] 
**ClusterMaxVersion** | Pointer to **NullableInt32** | The maximum supported cluster version for this item. | [optional] 
**ClusterMinVersion** | Pointer to **NullableInt32** | The minimum cluster version required to use this item. | [optional] 
**Description** | Pointer to **NullableString** | Description of the item. | [optional] 
**DescriptionBlocks** | Pointer to [**[]DescriptionBlock**](DescriptionBlock.md) |  | [optional] 
**DocumentationLink** | Pointer to **NullableString** | An absolute link to a documentation page explaining the item. | [optional] 
**Extension1Details** | Pointer to [**NullableExtension1Details**](Extension1Details.md) |  | [optional] 
**Extension2Details** | Pointer to [**NullableExtension2Details**](Extension2Details.md) |  | [optional] 
**ItemId** | Pointer to **string** | Unique Id of the item. | [optional] 
**Logo** | Pointer to **NullableString** | The logo of the item. Can be a URL or Base64 encoded. Intended for &lt;image&gt; html tags. | [optional] 
**MarketingLink** | Pointer to **NullableString** | An absolute link to a marketing page promoting how the item can be used with dynatrace. | [optional] 
**Name** | Pointer to **string** | Name of the item. | [optional] 
**NotCompatibleReason** | Pointer to **NullableString** | The reason why the item is not compatible with the cluster version. | [optional] 
**RelatedItems** | Pointer to [**[]RelatedItem**](RelatedItem.md) | Related items. | [optional] 
**Tags** | Pointer to **[]string** | Grouping of items with keywords. | [optional] 
**TechnologyDetails** | Pointer to [**NullableTechnologyDetails**](TechnologyDetails.md) |  | [optional] 
**Type** | Pointer to **string** | Represents the type of item. It can be TECHNOLOGY, EXTENSION1 or EXTENSION2. | [optional] 

## Methods

### NewItemDetails

`func NewItemDetails() *ItemDetails`

NewItemDetails instantiates a new ItemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemDetailsWithDefaults

`func NewItemDetailsWithDefaults() *ItemDetails`

NewItemDetailsWithDefaults instantiates a new ItemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorLogo

`func (o *ItemDetails) GetAuthorLogo() string`

GetAuthorLogo returns the AuthorLogo field if non-nil, zero value otherwise.

### GetAuthorLogoOk

`func (o *ItemDetails) GetAuthorLogoOk() (*string, bool)`

GetAuthorLogoOk returns a tuple with the AuthorLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorLogo

`func (o *ItemDetails) SetAuthorLogo(v string)`

SetAuthorLogo sets AuthorLogo field to given value.

### HasAuthorLogo

`func (o *ItemDetails) HasAuthorLogo() bool`

HasAuthorLogo returns a boolean if a field has been set.

### SetAuthorLogoNil

`func (o *ItemDetails) SetAuthorLogoNil(b bool)`

 SetAuthorLogoNil sets the value for AuthorLogo to be an explicit nil

### UnsetAuthorLogo
`func (o *ItemDetails) UnsetAuthorLogo()`

UnsetAuthorLogo ensures that no value is present for AuthorLogo, not even an explicit nil
### GetAuthorName

`func (o *ItemDetails) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *ItemDetails) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *ItemDetails) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *ItemDetails) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### SetAuthorNameNil

`func (o *ItemDetails) SetAuthorNameNil(b bool)`

 SetAuthorNameNil sets the value for AuthorName to be an explicit nil

### UnsetAuthorName
`func (o *ItemDetails) UnsetAuthorName()`

UnsetAuthorName ensures that no value is present for AuthorName, not even an explicit nil
### GetClusterCompatible

`func (o *ItemDetails) GetClusterCompatible() bool`

GetClusterCompatible returns the ClusterCompatible field if non-nil, zero value otherwise.

### GetClusterCompatibleOk

`func (o *ItemDetails) GetClusterCompatibleOk() (*bool, bool)`

GetClusterCompatibleOk returns a tuple with the ClusterCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCompatible

`func (o *ItemDetails) SetClusterCompatible(v bool)`

SetClusterCompatible sets ClusterCompatible field to given value.

### HasClusterCompatible

`func (o *ItemDetails) HasClusterCompatible() bool`

HasClusterCompatible returns a boolean if a field has been set.

### SetClusterCompatibleNil

`func (o *ItemDetails) SetClusterCompatibleNil(b bool)`

 SetClusterCompatibleNil sets the value for ClusterCompatible to be an explicit nil

### UnsetClusterCompatible
`func (o *ItemDetails) UnsetClusterCompatible()`

UnsetClusterCompatible ensures that no value is present for ClusterCompatible, not even an explicit nil
### GetClusterMaxVersion

`func (o *ItemDetails) GetClusterMaxVersion() int32`

GetClusterMaxVersion returns the ClusterMaxVersion field if non-nil, zero value otherwise.

### GetClusterMaxVersionOk

`func (o *ItemDetails) GetClusterMaxVersionOk() (*int32, bool)`

GetClusterMaxVersionOk returns a tuple with the ClusterMaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMaxVersion

`func (o *ItemDetails) SetClusterMaxVersion(v int32)`

SetClusterMaxVersion sets ClusterMaxVersion field to given value.

### HasClusterMaxVersion

`func (o *ItemDetails) HasClusterMaxVersion() bool`

HasClusterMaxVersion returns a boolean if a field has been set.

### SetClusterMaxVersionNil

`func (o *ItemDetails) SetClusterMaxVersionNil(b bool)`

 SetClusterMaxVersionNil sets the value for ClusterMaxVersion to be an explicit nil

### UnsetClusterMaxVersion
`func (o *ItemDetails) UnsetClusterMaxVersion()`

UnsetClusterMaxVersion ensures that no value is present for ClusterMaxVersion, not even an explicit nil
### GetClusterMinVersion

`func (o *ItemDetails) GetClusterMinVersion() int32`

GetClusterMinVersion returns the ClusterMinVersion field if non-nil, zero value otherwise.

### GetClusterMinVersionOk

`func (o *ItemDetails) GetClusterMinVersionOk() (*int32, bool)`

GetClusterMinVersionOk returns a tuple with the ClusterMinVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterMinVersion

`func (o *ItemDetails) SetClusterMinVersion(v int32)`

SetClusterMinVersion sets ClusterMinVersion field to given value.

### HasClusterMinVersion

`func (o *ItemDetails) HasClusterMinVersion() bool`

HasClusterMinVersion returns a boolean if a field has been set.

### SetClusterMinVersionNil

`func (o *ItemDetails) SetClusterMinVersionNil(b bool)`

 SetClusterMinVersionNil sets the value for ClusterMinVersion to be an explicit nil

### UnsetClusterMinVersion
`func (o *ItemDetails) UnsetClusterMinVersion()`

UnsetClusterMinVersion ensures that no value is present for ClusterMinVersion, not even an explicit nil
### GetDescription

`func (o *ItemDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ItemDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ItemDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDescriptionBlocks

`func (o *ItemDetails) GetDescriptionBlocks() []DescriptionBlock`

GetDescriptionBlocks returns the DescriptionBlocks field if non-nil, zero value otherwise.

### GetDescriptionBlocksOk

`func (o *ItemDetails) GetDescriptionBlocksOk() (*[]DescriptionBlock, bool)`

GetDescriptionBlocksOk returns a tuple with the DescriptionBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionBlocks

`func (o *ItemDetails) SetDescriptionBlocks(v []DescriptionBlock)`

SetDescriptionBlocks sets DescriptionBlocks field to given value.

### HasDescriptionBlocks

`func (o *ItemDetails) HasDescriptionBlocks() bool`

HasDescriptionBlocks returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *ItemDetails) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *ItemDetails) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *ItemDetails) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *ItemDetails) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### SetDocumentationLinkNil

`func (o *ItemDetails) SetDocumentationLinkNil(b bool)`

 SetDocumentationLinkNil sets the value for DocumentationLink to be an explicit nil

### UnsetDocumentationLink
`func (o *ItemDetails) UnsetDocumentationLink()`

UnsetDocumentationLink ensures that no value is present for DocumentationLink, not even an explicit nil
### GetExtension1Details

`func (o *ItemDetails) GetExtension1Details() Extension1Details`

GetExtension1Details returns the Extension1Details field if non-nil, zero value otherwise.

### GetExtension1DetailsOk

`func (o *ItemDetails) GetExtension1DetailsOk() (*Extension1Details, bool)`

GetExtension1DetailsOk returns a tuple with the Extension1Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension1Details

`func (o *ItemDetails) SetExtension1Details(v Extension1Details)`

SetExtension1Details sets Extension1Details field to given value.

### HasExtension1Details

`func (o *ItemDetails) HasExtension1Details() bool`

HasExtension1Details returns a boolean if a field has been set.

### SetExtension1DetailsNil

`func (o *ItemDetails) SetExtension1DetailsNil(b bool)`

 SetExtension1DetailsNil sets the value for Extension1Details to be an explicit nil

### UnsetExtension1Details
`func (o *ItemDetails) UnsetExtension1Details()`

UnsetExtension1Details ensures that no value is present for Extension1Details, not even an explicit nil
### GetExtension2Details

`func (o *ItemDetails) GetExtension2Details() Extension2Details`

GetExtension2Details returns the Extension2Details field if non-nil, zero value otherwise.

### GetExtension2DetailsOk

`func (o *ItemDetails) GetExtension2DetailsOk() (*Extension2Details, bool)`

GetExtension2DetailsOk returns a tuple with the Extension2Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension2Details

`func (o *ItemDetails) SetExtension2Details(v Extension2Details)`

SetExtension2Details sets Extension2Details field to given value.

### HasExtension2Details

`func (o *ItemDetails) HasExtension2Details() bool`

HasExtension2Details returns a boolean if a field has been set.

### SetExtension2DetailsNil

`func (o *ItemDetails) SetExtension2DetailsNil(b bool)`

 SetExtension2DetailsNil sets the value for Extension2Details to be an explicit nil

### UnsetExtension2Details
`func (o *ItemDetails) UnsetExtension2Details()`

UnsetExtension2Details ensures that no value is present for Extension2Details, not even an explicit nil
### GetItemId

`func (o *ItemDetails) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemDetails) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemDetails) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemDetails) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetLogo

`func (o *ItemDetails) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ItemDetails) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ItemDetails) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ItemDetails) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### SetLogoNil

`func (o *ItemDetails) SetLogoNil(b bool)`

 SetLogoNil sets the value for Logo to be an explicit nil

### UnsetLogo
`func (o *ItemDetails) UnsetLogo()`

UnsetLogo ensures that no value is present for Logo, not even an explicit nil
### GetMarketingLink

`func (o *ItemDetails) GetMarketingLink() string`

GetMarketingLink returns the MarketingLink field if non-nil, zero value otherwise.

### GetMarketingLinkOk

`func (o *ItemDetails) GetMarketingLinkOk() (*string, bool)`

GetMarketingLinkOk returns a tuple with the MarketingLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingLink

`func (o *ItemDetails) SetMarketingLink(v string)`

SetMarketingLink sets MarketingLink field to given value.

### HasMarketingLink

`func (o *ItemDetails) HasMarketingLink() bool`

HasMarketingLink returns a boolean if a field has been set.

### SetMarketingLinkNil

`func (o *ItemDetails) SetMarketingLinkNil(b bool)`

 SetMarketingLinkNil sets the value for MarketingLink to be an explicit nil

### UnsetMarketingLink
`func (o *ItemDetails) UnsetMarketingLink()`

UnsetMarketingLink ensures that no value is present for MarketingLink, not even an explicit nil
### GetName

`func (o *ItemDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ItemDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotCompatibleReason

`func (o *ItemDetails) GetNotCompatibleReason() string`

GetNotCompatibleReason returns the NotCompatibleReason field if non-nil, zero value otherwise.

### GetNotCompatibleReasonOk

`func (o *ItemDetails) GetNotCompatibleReasonOk() (*string, bool)`

GetNotCompatibleReasonOk returns a tuple with the NotCompatibleReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCompatibleReason

`func (o *ItemDetails) SetNotCompatibleReason(v string)`

SetNotCompatibleReason sets NotCompatibleReason field to given value.

### HasNotCompatibleReason

`func (o *ItemDetails) HasNotCompatibleReason() bool`

HasNotCompatibleReason returns a boolean if a field has been set.

### SetNotCompatibleReasonNil

`func (o *ItemDetails) SetNotCompatibleReasonNil(b bool)`

 SetNotCompatibleReasonNil sets the value for NotCompatibleReason to be an explicit nil

### UnsetNotCompatibleReason
`func (o *ItemDetails) UnsetNotCompatibleReason()`

UnsetNotCompatibleReason ensures that no value is present for NotCompatibleReason, not even an explicit nil
### GetRelatedItems

`func (o *ItemDetails) GetRelatedItems() []RelatedItem`

GetRelatedItems returns the RelatedItems field if non-nil, zero value otherwise.

### GetRelatedItemsOk

`func (o *ItemDetails) GetRelatedItemsOk() (*[]RelatedItem, bool)`

GetRelatedItemsOk returns a tuple with the RelatedItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedItems

`func (o *ItemDetails) SetRelatedItems(v []RelatedItem)`

SetRelatedItems sets RelatedItems field to given value.

### HasRelatedItems

`func (o *ItemDetails) HasRelatedItems() bool`

HasRelatedItems returns a boolean if a field has been set.

### GetTags

`func (o *ItemDetails) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ItemDetails) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ItemDetails) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ItemDetails) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTechnologyDetails

`func (o *ItemDetails) GetTechnologyDetails() TechnologyDetails`

GetTechnologyDetails returns the TechnologyDetails field if non-nil, zero value otherwise.

### GetTechnologyDetailsOk

`func (o *ItemDetails) GetTechnologyDetailsOk() (*TechnologyDetails, bool)`

GetTechnologyDetailsOk returns a tuple with the TechnologyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnologyDetails

`func (o *ItemDetails) SetTechnologyDetails(v TechnologyDetails)`

SetTechnologyDetails sets TechnologyDetails field to given value.

### HasTechnologyDetails

`func (o *ItemDetails) HasTechnologyDetails() bool`

HasTechnologyDetails returns a boolean if a field has been set.

### SetTechnologyDetailsNil

`func (o *ItemDetails) SetTechnologyDetailsNil(b bool)`

 SetTechnologyDetailsNil sets the value for TechnologyDetails to be an explicit nil

### UnsetTechnologyDetails
`func (o *ItemDetails) UnsetTechnologyDetails()`

UnsetTechnologyDetails ensures that no value is present for TechnologyDetails, not even an explicit nil
### GetType

`func (o *ItemDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemDetails) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ItemDetails) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


