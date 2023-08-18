# DatasourceDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterProperties** | **[]string** | The properties to filter the datasource options on. | 
**FullContext** | **bool** | Whether this datasource expects full setting payload as the context. | 
**Identifier** | **string** | The identifier of a custom data source of the property&#39;s value. | 
**ResetValue** | Pointer to **string** | When to reset datasource value in the UI on filter change. | [optional] 
**UseApiSearch** | **bool** | If true, the datasource should use the api to filter the results instead of client-side filtering. | 
**Validate** | **bool** | Whether to validate input to only allow values returned by the datasource. | 

## Methods

### NewDatasourceDefinition

`func NewDatasourceDefinition(filterProperties []string, fullContext bool, identifier string, useApiSearch bool, validate bool, ) *DatasourceDefinition`

NewDatasourceDefinition instantiates a new DatasourceDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasourceDefinitionWithDefaults

`func NewDatasourceDefinitionWithDefaults() *DatasourceDefinition`

NewDatasourceDefinitionWithDefaults instantiates a new DatasourceDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterProperties

`func (o *DatasourceDefinition) GetFilterProperties() []string`

GetFilterProperties returns the FilterProperties field if non-nil, zero value otherwise.

### GetFilterPropertiesOk

`func (o *DatasourceDefinition) GetFilterPropertiesOk() (*[]string, bool)`

GetFilterPropertiesOk returns a tuple with the FilterProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterProperties

`func (o *DatasourceDefinition) SetFilterProperties(v []string)`

SetFilterProperties sets FilterProperties field to given value.


### GetFullContext

`func (o *DatasourceDefinition) GetFullContext() bool`

GetFullContext returns the FullContext field if non-nil, zero value otherwise.

### GetFullContextOk

`func (o *DatasourceDefinition) GetFullContextOk() (*bool, bool)`

GetFullContextOk returns a tuple with the FullContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullContext

`func (o *DatasourceDefinition) SetFullContext(v bool)`

SetFullContext sets FullContext field to given value.


### GetIdentifier

`func (o *DatasourceDefinition) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *DatasourceDefinition) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *DatasourceDefinition) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetResetValue

`func (o *DatasourceDefinition) GetResetValue() string`

GetResetValue returns the ResetValue field if non-nil, zero value otherwise.

### GetResetValueOk

`func (o *DatasourceDefinition) GetResetValueOk() (*string, bool)`

GetResetValueOk returns a tuple with the ResetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetValue

`func (o *DatasourceDefinition) SetResetValue(v string)`

SetResetValue sets ResetValue field to given value.

### HasResetValue

`func (o *DatasourceDefinition) HasResetValue() bool`

HasResetValue returns a boolean if a field has been set.

### GetUseApiSearch

`func (o *DatasourceDefinition) GetUseApiSearch() bool`

GetUseApiSearch returns the UseApiSearch field if non-nil, zero value otherwise.

### GetUseApiSearchOk

`func (o *DatasourceDefinition) GetUseApiSearchOk() (*bool, bool)`

GetUseApiSearchOk returns a tuple with the UseApiSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseApiSearch

`func (o *DatasourceDefinition) SetUseApiSearch(v bool)`

SetUseApiSearch sets UseApiSearch field to given value.


### GetValidate

`func (o *DatasourceDefinition) GetValidate() bool`

GetValidate returns the Validate field if non-nil, zero value otherwise.

### GetValidateOk

`func (o *DatasourceDefinition) GetValidateOk() (*bool, bool)`

GetValidateOk returns a tuple with the Validate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidate

`func (o *DatasourceDefinition) SetValidate(v bool)`

SetValidate sets Validate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


