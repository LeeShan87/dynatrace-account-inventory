# RelatedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**ExternalLink** | Pointer to **string** | External link (marketing/documentation) that can provide with additional information. | [optional] 
**HasClusterLink** | Pointer to **bool** | Indicates whether there is a page within the product to activate this item. | [optional] 
**IconUrl** | Pointer to **string** | The logo of the item. Can be a URL or Base64 encoded. Intended for &lt;image&gt; html tags | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Represents the type of item. It can be TECHNOLOGY, EXTENSION1 or EXTENSION2. | [optional] 

## Methods

### NewRelatedItem

`func NewRelatedItem() *RelatedItem`

NewRelatedItem instantiates a new RelatedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedItemWithDefaults

`func NewRelatedItemWithDefaults() *RelatedItem`

NewRelatedItemWithDefaults instantiates a new RelatedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RelatedItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RelatedItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RelatedItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RelatedItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalLink

`func (o *RelatedItem) GetExternalLink() string`

GetExternalLink returns the ExternalLink field if non-nil, zero value otherwise.

### GetExternalLinkOk

`func (o *RelatedItem) GetExternalLinkOk() (*string, bool)`

GetExternalLinkOk returns a tuple with the ExternalLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLink

`func (o *RelatedItem) SetExternalLink(v string)`

SetExternalLink sets ExternalLink field to given value.

### HasExternalLink

`func (o *RelatedItem) HasExternalLink() bool`

HasExternalLink returns a boolean if a field has been set.

### GetHasClusterLink

`func (o *RelatedItem) GetHasClusterLink() bool`

GetHasClusterLink returns the HasClusterLink field if non-nil, zero value otherwise.

### GetHasClusterLinkOk

`func (o *RelatedItem) GetHasClusterLinkOk() (*bool, bool)`

GetHasClusterLinkOk returns a tuple with the HasClusterLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasClusterLink

`func (o *RelatedItem) SetHasClusterLink(v bool)`

SetHasClusterLink sets HasClusterLink field to given value.

### HasHasClusterLink

`func (o *RelatedItem) HasHasClusterLink() bool`

HasHasClusterLink returns a boolean if a field has been set.

### GetIconUrl

`func (o *RelatedItem) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *RelatedItem) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *RelatedItem) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *RelatedItem) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetId

`func (o *RelatedItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RelatedItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RelatedItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RelatedItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RelatedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelatedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelatedItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RelatedItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *RelatedItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelatedItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelatedItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RelatedItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


