# PropertyDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**[]Constraint**](Constraint.md) | A list of constraints limiting the values to be accepted. | [optional] 
**Datasource** | Pointer to [**DatasourceDefinition**](DatasourceDefinition.md) |  | [optional] 
**Default** | Pointer to **map[string]interface{}** | The default value to be used when no value is provided.   If a non-singleton has the value of &#x60;null&#x60;, it means an empty collection. | [optional] 
**Description** | Pointer to **string** | A short description of the property. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the property. | [optional] 
**Documentation** | Pointer to **string** | An extended description and/or links to documentation. | [optional] 
**ForceSecretResubmission** | Pointer to **bool** | Defines if value is allowed to be modified when secret properties are not | [optional] 
**Items** | Pointer to [**Item**](Item.md) |  | [optional] 
**MaxObjects** | **int32** | The maximum number of **objects** in a collection property.    Has the value of &#x60;1&#x60; for singletons. | 
**Metadata** | Pointer to **map[string]string** | Metadata of the property. | [optional] 
**MinObjects** | Pointer to **int32** | The minimum number of **objects** in a collection property. | [optional] 
**ModificationPolicy** | Pointer to **string** | Modification policy of the property. | [optional] 
**Nullable** | **bool** | The value can (&#x60;true&#x60;) or can&#39;t (&#x60;false&#x60;) be &#x60;null&#x60;. | 
**Precondition** | Pointer to [**Precondition**](Precondition.md) |  | [optional] 
**ReferencedType** | Pointer to **string** | The type referenced by the property value | [optional] 
**SubType** | Pointer to **string** | The subtype of the property&#39;s value. | [optional] 
**Type** | [**PropertyDefinitionType**](PropertyDefinitionType.md) |  | 
**UiCustomization** | Pointer to [**UiCustomization**](UiCustomization.md) |  | [optional] 

## Methods

### NewPropertyDefinition

`func NewPropertyDefinition(maxObjects int32, nullable bool, type_ PropertyDefinitionType, ) *PropertyDefinition`

NewPropertyDefinition instantiates a new PropertyDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyDefinitionWithDefaults

`func NewPropertyDefinitionWithDefaults() *PropertyDefinition`

NewPropertyDefinitionWithDefaults instantiates a new PropertyDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *PropertyDefinition) GetConstraints() []Constraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *PropertyDefinition) GetConstraintsOk() (*[]Constraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *PropertyDefinition) SetConstraints(v []Constraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *PropertyDefinition) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetDatasource

`func (o *PropertyDefinition) GetDatasource() DatasourceDefinition`

GetDatasource returns the Datasource field if non-nil, zero value otherwise.

### GetDatasourceOk

`func (o *PropertyDefinition) GetDatasourceOk() (*DatasourceDefinition, bool)`

GetDatasourceOk returns a tuple with the Datasource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasource

`func (o *PropertyDefinition) SetDatasource(v DatasourceDefinition)`

SetDatasource sets Datasource field to given value.

### HasDatasource

`func (o *PropertyDefinition) HasDatasource() bool`

HasDatasource returns a boolean if a field has been set.

### GetDefault

`func (o *PropertyDefinition) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PropertyDefinition) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PropertyDefinition) SetDefault(v map[string]interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PropertyDefinition) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *PropertyDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PropertyDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PropertyDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PropertyDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *PropertyDefinition) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PropertyDefinition) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PropertyDefinition) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PropertyDefinition) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDocumentation

`func (o *PropertyDefinition) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *PropertyDefinition) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *PropertyDefinition) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *PropertyDefinition) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetForceSecretResubmission

`func (o *PropertyDefinition) GetForceSecretResubmission() bool`

GetForceSecretResubmission returns the ForceSecretResubmission field if non-nil, zero value otherwise.

### GetForceSecretResubmissionOk

`func (o *PropertyDefinition) GetForceSecretResubmissionOk() (*bool, bool)`

GetForceSecretResubmissionOk returns a tuple with the ForceSecretResubmission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSecretResubmission

`func (o *PropertyDefinition) SetForceSecretResubmission(v bool)`

SetForceSecretResubmission sets ForceSecretResubmission field to given value.

### HasForceSecretResubmission

`func (o *PropertyDefinition) HasForceSecretResubmission() bool`

HasForceSecretResubmission returns a boolean if a field has been set.

### GetItems

`func (o *PropertyDefinition) GetItems() Item`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PropertyDefinition) GetItemsOk() (*Item, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PropertyDefinition) SetItems(v Item)`

SetItems sets Items field to given value.

### HasItems

`func (o *PropertyDefinition) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMaxObjects

`func (o *PropertyDefinition) GetMaxObjects() int32`

GetMaxObjects returns the MaxObjects field if non-nil, zero value otherwise.

### GetMaxObjectsOk

`func (o *PropertyDefinition) GetMaxObjectsOk() (*int32, bool)`

GetMaxObjectsOk returns a tuple with the MaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxObjects

`func (o *PropertyDefinition) SetMaxObjects(v int32)`

SetMaxObjects sets MaxObjects field to given value.


### GetMetadata

`func (o *PropertyDefinition) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PropertyDefinition) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PropertyDefinition) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PropertyDefinition) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMinObjects

`func (o *PropertyDefinition) GetMinObjects() int32`

GetMinObjects returns the MinObjects field if non-nil, zero value otherwise.

### GetMinObjectsOk

`func (o *PropertyDefinition) GetMinObjectsOk() (*int32, bool)`

GetMinObjectsOk returns a tuple with the MinObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinObjects

`func (o *PropertyDefinition) SetMinObjects(v int32)`

SetMinObjects sets MinObjects field to given value.

### HasMinObjects

`func (o *PropertyDefinition) HasMinObjects() bool`

HasMinObjects returns a boolean if a field has been set.

### GetModificationPolicy

`func (o *PropertyDefinition) GetModificationPolicy() string`

GetModificationPolicy returns the ModificationPolicy field if non-nil, zero value otherwise.

### GetModificationPolicyOk

`func (o *PropertyDefinition) GetModificationPolicyOk() (*string, bool)`

GetModificationPolicyOk returns a tuple with the ModificationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModificationPolicy

`func (o *PropertyDefinition) SetModificationPolicy(v string)`

SetModificationPolicy sets ModificationPolicy field to given value.

### HasModificationPolicy

`func (o *PropertyDefinition) HasModificationPolicy() bool`

HasModificationPolicy returns a boolean if a field has been set.

### GetNullable

`func (o *PropertyDefinition) GetNullable() bool`

GetNullable returns the Nullable field if non-nil, zero value otherwise.

### GetNullableOk

`func (o *PropertyDefinition) GetNullableOk() (*bool, bool)`

GetNullableOk returns a tuple with the Nullable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNullable

`func (o *PropertyDefinition) SetNullable(v bool)`

SetNullable sets Nullable field to given value.


### GetPrecondition

`func (o *PropertyDefinition) GetPrecondition() Precondition`

GetPrecondition returns the Precondition field if non-nil, zero value otherwise.

### GetPreconditionOk

`func (o *PropertyDefinition) GetPreconditionOk() (*Precondition, bool)`

GetPreconditionOk returns a tuple with the Precondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecondition

`func (o *PropertyDefinition) SetPrecondition(v Precondition)`

SetPrecondition sets Precondition field to given value.

### HasPrecondition

`func (o *PropertyDefinition) HasPrecondition() bool`

HasPrecondition returns a boolean if a field has been set.

### GetReferencedType

`func (o *PropertyDefinition) GetReferencedType() string`

GetReferencedType returns the ReferencedType field if non-nil, zero value otherwise.

### GetReferencedTypeOk

`func (o *PropertyDefinition) GetReferencedTypeOk() (*string, bool)`

GetReferencedTypeOk returns a tuple with the ReferencedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedType

`func (o *PropertyDefinition) SetReferencedType(v string)`

SetReferencedType sets ReferencedType field to given value.

### HasReferencedType

`func (o *PropertyDefinition) HasReferencedType() bool`

HasReferencedType returns a boolean if a field has been set.

### GetSubType

`func (o *PropertyDefinition) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *PropertyDefinition) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *PropertyDefinition) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *PropertyDefinition) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetType

`func (o *PropertyDefinition) GetType() PropertyDefinitionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PropertyDefinition) GetTypeOk() (*PropertyDefinitionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PropertyDefinition) SetType(v PropertyDefinitionType)`

SetType sets Type field to given value.


### GetUiCustomization

`func (o *PropertyDefinition) GetUiCustomization() UiCustomization`

GetUiCustomization returns the UiCustomization field if non-nil, zero value otherwise.

### GetUiCustomizationOk

`func (o *PropertyDefinition) GetUiCustomizationOk() (*UiCustomization, bool)`

GetUiCustomizationOk returns a tuple with the UiCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiCustomization

`func (o *PropertyDefinition) SetUiCustomization(v UiCustomization)`

SetUiCustomization sets UiCustomization field to given value.

### HasUiCustomization

`func (o *PropertyDefinition) HasUiCustomization() bool`

HasUiCustomization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


