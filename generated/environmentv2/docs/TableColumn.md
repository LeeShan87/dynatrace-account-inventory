# TableColumn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | **string** | Pattern with references to properties to create a single value for the column. | 

## Methods

### NewTableColumn

`func NewTableColumn(pattern string, ) *TableColumn`

NewTableColumn instantiates a new TableColumn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTableColumnWithDefaults

`func NewTableColumnWithDefaults() *TableColumn`

NewTableColumnWithDefaults instantiates a new TableColumn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *TableColumn) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *TableColumn) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *TableColumn) SetPattern(v string)`

SetPattern sets Pattern field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


