# SchemaType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Constraints** | Pointer to [**[]ComplexConstraint**](ComplexConstraint.md) | A list of constraints limiting the values to be accepted. | [optional] 
**Description** | **string** | A short description of the property. | 
**DisplayName** | Pointer to **string** | The display name of the property. | [optional] 
**Documentation** | **string** | An extended description and/or links to documentation. | 
**Properties** | [**map[string]PropertyDefinition**](PropertyDefinition.md) | Definition of properties that can be persisted. | 
**SearchPattern** | Pointer to **string** | The pattern for the summary search(for example, \&quot;Alert after *X* minutes.\&quot;) of the configuration in the UI. | [optional] 
**SummaryPattern** | **string** | The pattern for the summary (for example, \&quot;Alert after *X* minutes.\&quot;) of the configuration in the UI. | 
**Type** | **string** | Type of the reference type. | 
**Version** | **string** | The version of the type. | 
**VersionInfo** | Pointer to **string** | A short description of the version. | [optional] 

## Methods

### NewSchemaType

`func NewSchemaType(description string, documentation string, properties map[string]PropertyDefinition, summaryPattern string, type_ string, version string, ) *SchemaType`

NewSchemaType instantiates a new SchemaType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaTypeWithDefaults

`func NewSchemaTypeWithDefaults() *SchemaType`

NewSchemaTypeWithDefaults instantiates a new SchemaType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConstraints

`func (o *SchemaType) GetConstraints() []ComplexConstraint`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *SchemaType) GetConstraintsOk() (*[]ComplexConstraint, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *SchemaType) SetConstraints(v []ComplexConstraint)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *SchemaType) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaType) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDisplayName

`func (o *SchemaType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SchemaType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SchemaType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SchemaType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDocumentation

`func (o *SchemaType) GetDocumentation() string`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *SchemaType) GetDocumentationOk() (*string, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *SchemaType) SetDocumentation(v string)`

SetDocumentation sets Documentation field to given value.


### GetProperties

`func (o *SchemaType) GetProperties() map[string]PropertyDefinition`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SchemaType) GetPropertiesOk() (*map[string]PropertyDefinition, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SchemaType) SetProperties(v map[string]PropertyDefinition)`

SetProperties sets Properties field to given value.


### GetSearchPattern

`func (o *SchemaType) GetSearchPattern() string`

GetSearchPattern returns the SearchPattern field if non-nil, zero value otherwise.

### GetSearchPatternOk

`func (o *SchemaType) GetSearchPatternOk() (*string, bool)`

GetSearchPatternOk returns a tuple with the SearchPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPattern

`func (o *SchemaType) SetSearchPattern(v string)`

SetSearchPattern sets SearchPattern field to given value.

### HasSearchPattern

`func (o *SchemaType) HasSearchPattern() bool`

HasSearchPattern returns a boolean if a field has been set.

### GetSummaryPattern

`func (o *SchemaType) GetSummaryPattern() string`

GetSummaryPattern returns the SummaryPattern field if non-nil, zero value otherwise.

### GetSummaryPatternOk

`func (o *SchemaType) GetSummaryPatternOk() (*string, bool)`

GetSummaryPatternOk returns a tuple with the SummaryPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryPattern

`func (o *SchemaType) SetSummaryPattern(v string)`

SetSummaryPattern sets SummaryPattern field to given value.


### GetType

`func (o *SchemaType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaType) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *SchemaType) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaType) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaType) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersionInfo

`func (o *SchemaType) GetVersionInfo() string`

GetVersionInfo returns the VersionInfo field if non-nil, zero value otherwise.

### GetVersionInfoOk

`func (o *SchemaType) GetVersionInfoOk() (*string, bool)`

GetVersionInfoOk returns a tuple with the VersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionInfo

`func (o *SchemaType) SetVersionInfo(v string)`

SetVersionInfo sets VersionInfo field to given value.

### HasVersionInfo

`func (o *SchemaType) HasVersionInfo() bool`

HasVersionInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


