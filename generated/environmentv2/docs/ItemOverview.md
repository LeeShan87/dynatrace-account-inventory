# ItemOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationLink** | Pointer to **string** | The activation link for a technology | [optional] 
**ArtifactId** | Pointer to **string** | The unique ID used by the artifacts contained in releases. | [optional] 
**AuthorLogo** | Pointer to **string** | Url for the author&#39;s logo. | [optional] 
**AuthorName** | Pointer to **string** | Name of the author of the item. | [optional] 
**ClusterCompatible** | Pointer to **bool** | Checks if the item is compatible with the cluster version. | [optional] 
**ComingSoon** | Pointer to **bool** | Whether or not the item is planned to be released soon | [optional] 
**Description** | Pointer to **string** | Description of the item. | [optional] 
**DocumentationLink** | Pointer to **string** | An absolute link to the documentation page of this item. | [optional] 
**HasDescriptionBlocks** | Pointer to **bool** | Whether or not the details call will contain description blocks. | [optional] 
**ItemId** | Pointer to **string** | Unique Id of the item. | [optional] 
**Logo** | Pointer to **string** | The logo of the item. Can be a URL or Base64 encoded. Intended for &lt;image&gt; html tags | [optional] 
**MarketingLink** | Pointer to **string** | An absolute link to the marketing page of this item. | [optional] 
**Name** | Pointer to **string** | Name of the item. | [optional] 
**NotCompatibleReason** | Pointer to **string** | The reason why the item is not compatible with the cluster version. | [optional] 
**Tags** | Pointer to **[]string** | Grouping of items with keywords. | [optional] 
**Type** | Pointer to **string** | Represents the type of item. It can be TECHNOLOGY, EXTENSION1 or EXTENSION2. | [optional] 

## Methods

### NewItemOverview

`func NewItemOverview() *ItemOverview`

NewItemOverview instantiates a new ItemOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemOverviewWithDefaults

`func NewItemOverviewWithDefaults() *ItemOverview`

NewItemOverviewWithDefaults instantiates a new ItemOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationLink

`func (o *ItemOverview) GetActivationLink() string`

GetActivationLink returns the ActivationLink field if non-nil, zero value otherwise.

### GetActivationLinkOk

`func (o *ItemOverview) GetActivationLinkOk() (*string, bool)`

GetActivationLinkOk returns a tuple with the ActivationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationLink

`func (o *ItemOverview) SetActivationLink(v string)`

SetActivationLink sets ActivationLink field to given value.

### HasActivationLink

`func (o *ItemOverview) HasActivationLink() bool`

HasActivationLink returns a boolean if a field has been set.

### GetArtifactId

`func (o *ItemOverview) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *ItemOverview) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *ItemOverview) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *ItemOverview) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetAuthorLogo

`func (o *ItemOverview) GetAuthorLogo() string`

GetAuthorLogo returns the AuthorLogo field if non-nil, zero value otherwise.

### GetAuthorLogoOk

`func (o *ItemOverview) GetAuthorLogoOk() (*string, bool)`

GetAuthorLogoOk returns a tuple with the AuthorLogo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorLogo

`func (o *ItemOverview) SetAuthorLogo(v string)`

SetAuthorLogo sets AuthorLogo field to given value.

### HasAuthorLogo

`func (o *ItemOverview) HasAuthorLogo() bool`

HasAuthorLogo returns a boolean if a field has been set.

### GetAuthorName

`func (o *ItemOverview) GetAuthorName() string`

GetAuthorName returns the AuthorName field if non-nil, zero value otherwise.

### GetAuthorNameOk

`func (o *ItemOverview) GetAuthorNameOk() (*string, bool)`

GetAuthorNameOk returns a tuple with the AuthorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorName

`func (o *ItemOverview) SetAuthorName(v string)`

SetAuthorName sets AuthorName field to given value.

### HasAuthorName

`func (o *ItemOverview) HasAuthorName() bool`

HasAuthorName returns a boolean if a field has been set.

### GetClusterCompatible

`func (o *ItemOverview) GetClusterCompatible() bool`

GetClusterCompatible returns the ClusterCompatible field if non-nil, zero value otherwise.

### GetClusterCompatibleOk

`func (o *ItemOverview) GetClusterCompatibleOk() (*bool, bool)`

GetClusterCompatibleOk returns a tuple with the ClusterCompatible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterCompatible

`func (o *ItemOverview) SetClusterCompatible(v bool)`

SetClusterCompatible sets ClusterCompatible field to given value.

### HasClusterCompatible

`func (o *ItemOverview) HasClusterCompatible() bool`

HasClusterCompatible returns a boolean if a field has been set.

### GetComingSoon

`func (o *ItemOverview) GetComingSoon() bool`

GetComingSoon returns the ComingSoon field if non-nil, zero value otherwise.

### GetComingSoonOk

`func (o *ItemOverview) GetComingSoonOk() (*bool, bool)`

GetComingSoonOk returns a tuple with the ComingSoon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComingSoon

`func (o *ItemOverview) SetComingSoon(v bool)`

SetComingSoon sets ComingSoon field to given value.

### HasComingSoon

`func (o *ItemOverview) HasComingSoon() bool`

HasComingSoon returns a boolean if a field has been set.

### GetDescription

`func (o *ItemOverview) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ItemOverview) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ItemOverview) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ItemOverview) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDocumentationLink

`func (o *ItemOverview) GetDocumentationLink() string`

GetDocumentationLink returns the DocumentationLink field if non-nil, zero value otherwise.

### GetDocumentationLinkOk

`func (o *ItemOverview) GetDocumentationLinkOk() (*string, bool)`

GetDocumentationLinkOk returns a tuple with the DocumentationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLink

`func (o *ItemOverview) SetDocumentationLink(v string)`

SetDocumentationLink sets DocumentationLink field to given value.

### HasDocumentationLink

`func (o *ItemOverview) HasDocumentationLink() bool`

HasDocumentationLink returns a boolean if a field has been set.

### GetHasDescriptionBlocks

`func (o *ItemOverview) GetHasDescriptionBlocks() bool`

GetHasDescriptionBlocks returns the HasDescriptionBlocks field if non-nil, zero value otherwise.

### GetHasDescriptionBlocksOk

`func (o *ItemOverview) GetHasDescriptionBlocksOk() (*bool, bool)`

GetHasDescriptionBlocksOk returns a tuple with the HasDescriptionBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDescriptionBlocks

`func (o *ItemOverview) SetHasDescriptionBlocks(v bool)`

SetHasDescriptionBlocks sets HasDescriptionBlocks field to given value.

### HasHasDescriptionBlocks

`func (o *ItemOverview) HasHasDescriptionBlocks() bool`

HasHasDescriptionBlocks returns a boolean if a field has been set.

### GetItemId

`func (o *ItemOverview) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ItemOverview) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ItemOverview) SetItemId(v string)`

SetItemId sets ItemId field to given value.

### HasItemId

`func (o *ItemOverview) HasItemId() bool`

HasItemId returns a boolean if a field has been set.

### GetLogo

`func (o *ItemOverview) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ItemOverview) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ItemOverview) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ItemOverview) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetMarketingLink

`func (o *ItemOverview) GetMarketingLink() string`

GetMarketingLink returns the MarketingLink field if non-nil, zero value otherwise.

### GetMarketingLinkOk

`func (o *ItemOverview) GetMarketingLinkOk() (*string, bool)`

GetMarketingLinkOk returns a tuple with the MarketingLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketingLink

`func (o *ItemOverview) SetMarketingLink(v string)`

SetMarketingLink sets MarketingLink field to given value.

### HasMarketingLink

`func (o *ItemOverview) HasMarketingLink() bool`

HasMarketingLink returns a boolean if a field has been set.

### GetName

`func (o *ItemOverview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ItemOverview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ItemOverview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ItemOverview) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotCompatibleReason

`func (o *ItemOverview) GetNotCompatibleReason() string`

GetNotCompatibleReason returns the NotCompatibleReason field if non-nil, zero value otherwise.

### GetNotCompatibleReasonOk

`func (o *ItemOverview) GetNotCompatibleReasonOk() (*string, bool)`

GetNotCompatibleReasonOk returns a tuple with the NotCompatibleReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotCompatibleReason

`func (o *ItemOverview) SetNotCompatibleReason(v string)`

SetNotCompatibleReason sets NotCompatibleReason field to given value.

### HasNotCompatibleReason

`func (o *ItemOverview) HasNotCompatibleReason() bool`

HasNotCompatibleReason returns a boolean if a field has been set.

### GetTags

`func (o *ItemOverview) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ItemOverview) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ItemOverview) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ItemOverview) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *ItemOverview) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ItemOverview) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ItemOverview) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ItemOverview) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


