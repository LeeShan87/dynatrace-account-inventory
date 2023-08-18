# DataExplorerFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to  |  | [optional] 
**EntityAttribute** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to **string** |  | [optional] 
**FilterOperator** | Pointer to **string** |  | [optional] 
**FilterType** | Pointer to **string** |  | [optional] 
**GlobalEntity** | Pointer to **string** |  | [optional] 
**NestedFilters** | Pointer to  |  | [optional] 
**Relationship** | Pointer to [**DexpFilterRelationship**](DexpFilterRelationship.md) |  | [optional] 

## Methods

### NewDataExplorerFilter

`func NewDataExplorerFilter() *DataExplorerFilter`

NewDataExplorerFilter instantiates a new DataExplorerFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataExplorerFilterWithDefaults

`func NewDataExplorerFilterWithDefaults() *DataExplorerFilter`

NewDataExplorerFilterWithDefaults instantiates a new DataExplorerFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *DataExplorerFilter) GetCriteria() []DexpFilterCriterion`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *DataExplorerFilter) GetCriteriaOk() (*[]DexpFilterCriterion, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *DataExplorerFilter) SetCriteria(v []DexpFilterCriterion)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *DataExplorerFilter) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetEntityAttribute

`func (o *DataExplorerFilter) GetEntityAttribute() string`

GetEntityAttribute returns the EntityAttribute field if non-nil, zero value otherwise.

### GetEntityAttributeOk

`func (o *DataExplorerFilter) GetEntityAttributeOk() (*string, bool)`

GetEntityAttributeOk returns a tuple with the EntityAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityAttribute

`func (o *DataExplorerFilter) SetEntityAttribute(v string)`

SetEntityAttribute sets EntityAttribute field to given value.

### HasEntityAttribute

`func (o *DataExplorerFilter) HasEntityAttribute() bool`

HasEntityAttribute returns a boolean if a field has been set.

### GetFilter

`func (o *DataExplorerFilter) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DataExplorerFilter) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DataExplorerFilter) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *DataExplorerFilter) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFilterOperator

`func (o *DataExplorerFilter) GetFilterOperator() string`

GetFilterOperator returns the FilterOperator field if non-nil, zero value otherwise.

### GetFilterOperatorOk

`func (o *DataExplorerFilter) GetFilterOperatorOk() (*string, bool)`

GetFilterOperatorOk returns a tuple with the FilterOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterOperator

`func (o *DataExplorerFilter) SetFilterOperator(v string)`

SetFilterOperator sets FilterOperator field to given value.

### HasFilterOperator

`func (o *DataExplorerFilter) HasFilterOperator() bool`

HasFilterOperator returns a boolean if a field has been set.

### GetFilterType

`func (o *DataExplorerFilter) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *DataExplorerFilter) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *DataExplorerFilter) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.

### HasFilterType

`func (o *DataExplorerFilter) HasFilterType() bool`

HasFilterType returns a boolean if a field has been set.

### GetGlobalEntity

`func (o *DataExplorerFilter) GetGlobalEntity() string`

GetGlobalEntity returns the GlobalEntity field if non-nil, zero value otherwise.

### GetGlobalEntityOk

`func (o *DataExplorerFilter) GetGlobalEntityOk() (*string, bool)`

GetGlobalEntityOk returns a tuple with the GlobalEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalEntity

`func (o *DataExplorerFilter) SetGlobalEntity(v string)`

SetGlobalEntity sets GlobalEntity field to given value.

### HasGlobalEntity

`func (o *DataExplorerFilter) HasGlobalEntity() bool`

HasGlobalEntity returns a boolean if a field has been set.

### GetNestedFilters

`func (o *DataExplorerFilter) GetNestedFilters() []DataExplorerFilter`

GetNestedFilters returns the NestedFilters field if non-nil, zero value otherwise.

### GetNestedFiltersOk

`func (o *DataExplorerFilter) GetNestedFiltersOk() (*[]DataExplorerFilter, bool)`

GetNestedFiltersOk returns a tuple with the NestedFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedFilters

`func (o *DataExplorerFilter) SetNestedFilters(v []DataExplorerFilter)`

SetNestedFilters sets NestedFilters field to given value.

### HasNestedFilters

`func (o *DataExplorerFilter) HasNestedFilters() bool`

HasNestedFilters returns a boolean if a field has been set.

### GetRelationship

`func (o *DataExplorerFilter) GetRelationship() DexpFilterRelationship`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *DataExplorerFilter) GetRelationshipOk() (*DexpFilterRelationship, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *DataExplorerFilter) SetRelationship(v DexpFilterRelationship)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *DataExplorerFilter) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


