# Item

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**[]Constraint**](Constraint.md) | A list of constraints limiting the values to be accepted. | [optional] 
**Datasource** | Pointer to [**DatasourceDefinition**](DatasourceDefinition.md) |  | [optional] 
**Description** | Pointer to **string** | A short description of the item. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the item. | [optional] 
**Documentation** | Pointer to **string** | An extended description and/or links to documentation. | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata of the items. | [optional] 
**ReferencedType** | Pointer to **string** | The type referenced by the item&#39;s value. | [optional] 
**SubType** | Pointer to **string** | The subtype of the item&#39;s value. | [optional] 
**Type** | [**ItemType**](ItemType.md) |  | 
**UiCustomization** | Pointer to [**UiCustomization**](UiCustomization.md) |  | [optional] 

## Methods

### NewItem

`func NewItem(type_ ItemType, ) *Item`

NewItem instantiates a new Item object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemWithDefaults

`func NewItemWithDefaults() *Item`

NewItemWithDefaults instantiates a new Item object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *Item) GetConstraints() []Constraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *Item) GetConstraintsOk() (*[]Constraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *Item) SetConstraints(v []Constraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *Item) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetDatasource

`func (o *Item) GetDatasource() DatasourceDefinition`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *Item) GetDatasourceOk() (*DatasourceDefinition, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *Item) SetDatasource(v DatasourceDefinition)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *Item) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetDescription

`func (o *Item) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Item) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Item) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Item) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *Item) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Item) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Item) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Item) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDocumentation

`func (o *Item) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *Item) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *Item) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *Item) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetMetadata

`func (o *Item) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Item) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Item) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Item) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReferencedType

`func (o *Item) GetReferencedType() string`

GetReferencedType returns the ReferencedType field if non-nil, zero value otherwise.

### GetReferencedTypeOk

`func (o *Item) GetReferencedTypeOk() (*string, bool)`

GetReferencedTypeOk returns a tuple with the ReferencedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedType

`func (o *Item) SetReferencedType(v string)`

SetReferencedType sets ReferencedType field to given value.

### HasReferencedType

`func (o *Item) HasReferencedType() bool`

HasReferencedType returns a boolean if a field has been set.

### GetSubType

`func (o *Item) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Item) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Item) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Item) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetType

`func (o *Item) GetType() ItemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Item) GetTypeOk() (*ItemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Item) SetType(v ItemType)`

SetType sets Type field to given value.


### GetUiCustomization

`func (o *Item) GetUiCustomization() UiCustomization`

GetUiCustomization returns the UiCustomization field if non-nil, zero value otherwise.

### GetUiCustomizationOk

`func (o *Item) GetUiCustomizationOk() (*UiCustomization, bool)`

GetUiCustomizationOk returns a tuple with the UiCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiCustomization

`func (o *Item) SetUiCustomization(v UiCustomization)`

SetUiCustomization sets UiCustomization field to given value.

### HasUiCustomization

`func (o *Item) HasUiCustomization() bool`

HasUiCustomization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


