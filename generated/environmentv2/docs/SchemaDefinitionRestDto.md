# SchemaDefinitionRestDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedScopes** | **[]string** | A list of scopes where the schema can be used. | 
**Constraints** | Pointer to [**[]ComplexConstraint**](ComplexConstraint.md) | A list of constrains limiting the values to be accepted by the schema. | [optional] 
**DeletionConstraints** | Pointer to [**[]DeletionConstraint**](DeletionConstraint.md) | Constraints limiting the values to be deleted. | [optional] 
**Description** | **string** | A short description of the schema. | 
**DisplayName** | **string** | The display name of the schema. | 
**Documentation** | Pointer to **string** | An extended description of the schema and/or links to documentation. | [optional] 
**Dynatrace** | **string** | The version of the data format. | 
**Enums** | [**map[string]EnumType**](EnumType.md) | A list of definitions of enum properties. | 
**KeyProperty** | Pointer to **string** | Name of the key property in this schema. | [optional] 
**MaxObjects** | **int32** | The maximum amount of objects per scope.   Only applicable when **multiObject** is set to &#x60;true&#x60;. | 
**Metadata** | Pointer to **map[string]string** | Metadata of the setting. | [optional] 
**MultiObject** | **bool** | Multiple (&#x60;true&#x60;) objects per scope are permitted or a single (&#x60;false&#x60;) object per scope is permitted. | 
**Ordered** | Pointer to **bool** | If &#x60;true&#x60; the order of objects has semantic significance.   Only applicable when **multiObject** is set to &#x60;true&#x60;. | [optional] 
**Properties** | [**map[string]PropertyDefinition**](PropertyDefinition.md) | A list of schema&#39;s properties. | 
**SchemaConstraints** | Pointer to [**[]SchemaConstraintRestDto**](SchemaConstraintRestDto.md) | Constraints limiting the values as a whole to be accepted in this configuration element. | [optional] 
**SchemaGroups** | Pointer to **[]string** | Names of the groups, which the schema belongs to. | [optional] 
**SchemaId** | **string** | The ID of the schema. | 
**TableColumns** | Pointer to [**map[string]TableColumn**](TableColumn.md) | Table column definitions for use in the ui. | [optional] 
**Types** | [**map[string]SchemaType**](SchemaType.md) | A list of definitions of types.    A type is a complex property that contains its own set of subproperties. | 
**UiCustomization** | Pointer to [**UiCustomization**](UiCustomization.md) |  | [optional] 
**Version** | **string** | The version of the schema. | 

## Methods

### NewSchemaDefinitionRestDto

`func NewSchemaDefinitionRestDto(allowedScopes []string, description string, displayName string, dynatrace string, enums map[string]EnumType, maxObjects int32, multiObject bool, properties map[string]PropertyDefinition, schemaId string, types map[string]SchemaType, version string, ) *SchemaDefinitionRestDto`

NewSchemaDefinitionRestDto instantiates a new SchemaDefinitionRestDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaDefinitionRestDtoWithDefaults

`func NewSchemaDefinitionRestDtoWithDefaults() *SchemaDefinitionRestDto`

NewSchemaDefinitionRestDtoWithDefaults instantiates a new SchemaDefinitionRestDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedScopes

`func (o *SchemaDefinitionRestDto) GetAllowedScopes() []string`

GetAllowedScopes returns the AllowedScopes field if non-nil, zero value otherwise.

### GetAllowedScopesOk

`func (o *SchemaDefinitionRestDto) GetAllowedScopesOk() (*[]string, bool)`

GetAllowedScopesOk returns a tuple with the AllowedScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedScopes

`func (o *SchemaDefinitionRestDto) SetAllowedScopes(v []string)`

SetAllowedScopes sets AllowedScopes field to given value.


### GetConstraints

`func (o *SchemaDefinitionRestDto) GetConstraints() []ComplexConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *SchemaDefinitionRestDto) GetConstraintsOk() (*[]ComplexConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *SchemaDefinitionRestDto) SetConstraints(v []ComplexConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *SchemaDefinitionRestDto) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetDeletionConstraints

`func (o *SchemaDefinitionRestDto) GetDeletionConstraints() []DeletionConstraint`

GetDeletionConstraints returns the DeletionConstraints field if non-nil, zero value otherwise.

### GetDeletionConstraintsOk

`func (o *SchemaDefinitionRestDto) GetDeletionConstraintsOk() (*[]DeletionConstraint, bool)`

GetDeletionConstraintsOk returns a tuple with the DeletionConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionConstraints

`func (o *SchemaDefinitionRestDto) SetDeletionConstraints(v []DeletionConstraint)`

SetDeletionConstraints sets DeletionConstraints field to given value.

### HasDeletionConstraints

`func (o *SchemaDefinitionRestDto) HasDeletionConstraints() bool`

HasDeletionConstraints returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaDefinitionRestDto) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaDefinitionRestDto) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaDefinitionRestDto) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *SchemaDefinitionRestDto) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SchemaDefinitionRestDto) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SchemaDefinitionRestDto) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDocumentation

`func (o *SchemaDefinitionRestDto) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *SchemaDefinitionRestDto) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *SchemaDefinitionRestDto) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.

### HasDocumentation

`func (o *SchemaDefinitionRestDto) HasDocumentation() bool`

HasDocumentation returns a boolean if a field has been set.

### GetDynatrace

`func (o *SchemaDefinitionRestDto) GetDynatrace() string`

GetDynatrace returns the Dynatrace field if non-nil, zero value otherwise.

### GetDynatraceOk

`func (o *SchemaDefinitionRestDto) GetDynatraceOk() (*string, bool)`

GetDynatraceOk returns a tuple with the Dynatrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynatrace

`func (o *SchemaDefinitionRestDto) SetDynatrace(v string)`

SetDynatrace sets Dynatrace field to given value.


### GetEnums

`func (o *SchemaDefinitionRestDto) GetEnums() map[string]EnumType`

GetEnums returns the Enums field if non-nil, zero value otherwise.

### GetEnumsOk

`func (o *SchemaDefinitionRestDto) GetEnumsOk() (*map[string]EnumType, bool)`

GetEnumsOk returns a tuple with the Enums field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnums

`func (o *SchemaDefinitionRestDto) SetEnums(v map[string]EnumType)`

SetEnums sets Enums field to given value.


### GetKeyProperty

`func (o *SchemaDefinitionRestDto) GetKeyProperty() string`

GetKeyProperty returns the KeyProperty field if non-nil, zero value otherwise.

### GetKeyPropertyOk

`func (o *SchemaDefinitionRestDto) GetKeyPropertyOk() (*string, bool)`

GetKeyPropertyOk returns a tuple with the KeyProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyProperty

`func (o *SchemaDefinitionRestDto) SetKeyProperty(v string)`

SetKeyProperty sets KeyProperty field to given value.

### HasKeyProperty

`func (o *SchemaDefinitionRestDto) HasKeyProperty() bool`

HasKeyProperty returns a boolean if a field has been set.

### GetMaxObjects

`func (o *SchemaDefinitionRestDto) GetMaxObjects() int32`

GetMaxObjects returns the MaxObjects field if non-nil, zero value otherwise.

### GetMaxObjectsOk

`func (o *SchemaDefinitionRestDto) GetMaxObjectsOk() (*int32, bool)`

GetMaxObjectsOk returns a tuple with the MaxObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxObjects

`func (o *SchemaDefinitionRestDto) SetMaxObjects(v int32)`

SetMaxObjects sets MaxObjects field to given value.


### GetMetadata

`func (o *SchemaDefinitionRestDto) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SchemaDefinitionRestDto) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SchemaDefinitionRestDto) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SchemaDefinitionRestDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMultiObject

`func (o *SchemaDefinitionRestDto) GetMultiObject() bool`

GetMultiObject returns the MultiObject field if non-nil, zero value otherwise.

### GetMultiObjectOk

`func (o *SchemaDefinitionRestDto) GetMultiObjectOk() (*bool, bool)`

GetMultiObjectOk returns a tuple with the MultiObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiObject

`func (o *SchemaDefinitionRestDto) SetMultiObject(v bool)`

SetMultiObject sets MultiObject field to given value.


### GetOrdered

`func (o *SchemaDefinitionRestDto) GetOrdered() bool`

GetOrdered returns the Ordered field if non-nil, zero value otherwise.

### GetOrderedOk

`func (o *SchemaDefinitionRestDto) GetOrderedOk() (*bool, bool)`

GetOrderedOk returns a tuple with the Ordered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdered

`func (o *SchemaDefinitionRestDto) SetOrdered(v bool)`

SetOrdered sets Ordered field to given value.

### HasOrdered

`func (o *SchemaDefinitionRestDto) HasOrdered() bool`

HasOrdered returns a boolean if a field has been set.

### GetProperties

`func (o *SchemaDefinitionRestDto) GetProperties() map[string]PropertyDefinition`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SchemaDefinitionRestDto) GetPropertiesOk() (*map[string]PropertyDefinition, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SchemaDefinitionRestDto) SetProperties(v map[string]PropertyDefinition)`

SetProperties sets Properties field to given value.


### GetSchemaConstraints

`func (o *SchemaDefinitionRestDto) GetSchemaConstraints() []SchemaConstraintRestDto`

GetSchemaConstraints returns the SchemaConstraints field if non-nil, zero value otherwise.

### GetSchemaConstraintsOk

`func (o *SchemaDefinitionRestDto) GetSchemaConstraintsOk() (*[]SchemaConstraintRestDto, bool)`

GetSchemaConstraintsOk returns a tuple with the SchemaConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaConstraints

`func (o *SchemaDefinitionRestDto) SetSchemaConstraints(v []SchemaConstraintRestDto)`

SetSchemaConstraints sets SchemaConstraints field to given value.

### HasSchemaConstraints

`func (o *SchemaDefinitionRestDto) HasSchemaConstraints() bool`

HasSchemaConstraints returns a boolean if a field has been set.

### GetSchemaGroups

`func (o *SchemaDefinitionRestDto) GetSchemaGroups() []string`

GetSchemaGroups returns the SchemaGroups field if non-nil, zero value otherwise.

### GetSchemaGroupsOk

`func (o *SchemaDefinitionRestDto) GetSchemaGroupsOk() (*[]string, bool)`

GetSchemaGroupsOk returns a tuple with the SchemaGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaGroups

`func (o *SchemaDefinitionRestDto) SetSchemaGroups(v []string)`

SetSchemaGroups sets SchemaGroups field to given value.

### HasSchemaGroups

`func (o *SchemaDefinitionRestDto) HasSchemaGroups() bool`

HasSchemaGroups returns a boolean if a field has been set.

### GetSchemaId

`func (o *SchemaDefinitionRestDto) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *SchemaDefinitionRestDto) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *SchemaDefinitionRestDto) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.


### GetTableColumns

`func (o *SchemaDefinitionRestDto) GetTableColumns() map[string]TableColumn`

GetTableColumns returns the TableColumns field if non-nil, zero value otherwise.

### GetTableColumnsOk

`func (o *SchemaDefinitionRestDto) GetTableColumnsOk() (*map[string]TableColumn, bool)`

GetTableColumnsOk returns a tuple with the TableColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableColumns

`func (o *SchemaDefinitionRestDto) SetTableColumns(v map[string]TableColumn)`

SetTableColumns sets TableColumns field to given value.

### HasTableColumns

`func (o *SchemaDefinitionRestDto) HasTableColumns() bool`

HasTableColumns returns a boolean if a field has been set.

### GetTypes

`func (o *SchemaDefinitionRestDto) GetTypes() map[string]SchemaType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *SchemaDefinitionRestDto) GetTypesOk() (*map[string]SchemaType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *SchemaDefinitionRestDto) SetTypes(v map[string]SchemaType)`

SetTypes sets Types field to given value.


### GetUiCustomization

`func (o *SchemaDefinitionRestDto) GetUiCustomization() UiCustomization`

GetUiCustomization returns the UiCustomization field if non-nil, zero value otherwise.

### GetUiCustomizationOk

`func (o *SchemaDefinitionRestDto) GetUiCustomizationOk() (*UiCustomization, bool)`

GetUiCustomizationOk returns a tuple with the UiCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUiCustomization

`func (o *SchemaDefinitionRestDto) SetUiCustomization(v UiCustomization)`

SetUiCustomization sets UiCustomization field to given value.

### HasUiCustomization

`func (o *SchemaDefinitionRestDto) HasUiCustomization() bool`

HasUiCustomization returns a boolean if a field has been set.

### GetVersion

`func (o *SchemaDefinitionRestDto) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaDefinitionRestDto) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaDefinitionRestDto) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


