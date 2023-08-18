# UiTableCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]UiTableColumnCustomization**](UiTableColumnCustomization.md) | A list of columns for the UI table | [optional] 
**EmptyState** | Pointer to [**UiEmptyStateCustomization**](UiEmptyStateCustomization.md) |  | [optional] 

## Methods

### NewUiTableCustomization

`func NewUiTableCustomization() *UiTableCustomization`

NewUiTableCustomization instantiates a new UiTableCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiTableCustomizationWithDefaults

`func NewUiTableCustomizationWithDefaults() *UiTableCustomization`

NewUiTableCustomizationWithDefaults instantiates a new UiTableCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *UiTableCustomization) GetColumns() []UiTableColumnCustomization`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *UiTableCustomization) GetColumnsOk() (*[]UiTableColumnCustomization, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *UiTableCustomization) SetColumns(v []UiTableColumnCustomization)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *UiTableCustomization) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetEmptyState

`func (o *UiTableCustomization) GetEmptyState() UiEmptyStateCustomization`

GetEmptyState returns the EmptyState field if non-nil, zero value otherwise.

### GetEmptyStateOk

`func (o *UiTableCustomization) GetEmptyStateOk() (*UiEmptyStateCustomization, bool)`

GetEmptyStateOk returns a tuple with the EmptyState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyState

`func (o *UiTableCustomization) SetEmptyState(v UiEmptyStateCustomization)`

SetEmptyState sets EmptyState field to given value.

### HasEmptyState

`func (o *UiTableCustomization) HasEmptyState() bool`

HasEmptyState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


