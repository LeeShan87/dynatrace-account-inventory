# UiTableColumnCustomization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltinColumnRef** | Pointer to **string** | The ui specific builtin column-implementation for this column. | [optional] 
**ColumnRef** | Pointer to **string** | The referenced column from the &#39;tableColumns&#39; property of the schema for this column. | [optional] 
**DisplayName** | Pointer to **string** | The display name for this column. | [optional] 
**Items** | Pointer to [**[]UiTableColumnItemCustomization**](UiTableColumnItemCustomization.md) | The possible items of this column. | [optional] 
**PropertyRef** | Pointer to **string** | The referenced property for this column. | [optional] 
**Type** | Pointer to **string** | The ui specific type for this column. | [optional] 

## Methods

### NewUiTableColumnCustomization

`func NewUiTableColumnCustomization() *UiTableColumnCustomization`

NewUiTableColumnCustomization instantiates a new UiTableColumnCustomization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUiTableColumnCustomizationWithDefaults

`func NewUiTableColumnCustomizationWithDefaults() *UiTableColumnCustomization`

NewUiTableColumnCustomizationWithDefaults instantiates a new UiTableColumnCustomization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltinColumnRef

`func (o *UiTableColumnCustomization) GetBuiltinColumnRef() string`

GetBuiltinColumnRef returns the BuiltinColumnRef field if non-nil, zero value otherwise.

### GetBuiltinColumnRefOk

`func (o *UiTableColumnCustomization) GetBuiltinColumnRefOk() (*string, bool)`

GetBuiltinColumnRefOk returns a tuple with the BuiltinColumnRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltinColumnRef

`func (o *UiTableColumnCustomization) SetBuiltinColumnRef(v string)`

SetBuiltinColumnRef sets BuiltinColumnRef field to given value.

### HasBuiltinColumnRef

`func (o *UiTableColumnCustomization) HasBuiltinColumnRef() bool`

HasBuiltinColumnRef returns a boolean if a field has been set.

### GetColumnRef

`func (o *UiTableColumnCustomization) GetColumnRef() string`

GetColumnRef returns the ColumnRef field if non-nil, zero value otherwise.

### GetColumnRefOk

`func (o *UiTableColumnCustomization) GetColumnRefOk() (*string, bool)`

GetColumnRefOk returns a tuple with the ColumnRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumnRef

`func (o *UiTableColumnCustomization) SetColumnRef(v string)`

SetColumnRef sets ColumnRef field to given value.

### HasColumnRef

`func (o *UiTableColumnCustomization) HasColumnRef() bool`

HasColumnRef returns a boolean if a field has been set.

### GetDisplayName

`func (o *UiTableColumnCustomization) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UiTableColumnCustomization) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UiTableColumnCustomization) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UiTableColumnCustomization) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetItems

`func (o *UiTableColumnCustomization) GetItems() []UiTableColumnItemCustomization`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UiTableColumnCustomization) GetItemsOk() (*[]UiTableColumnItemCustomization, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UiTableColumnCustomization) SetItems(v []UiTableColumnItemCustomization)`

SetItems sets Items field to given value.

### HasItems

`func (o *UiTableColumnCustomization) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPropertyRef

`func (o *UiTableColumnCustomization) GetPropertyRef() string`

GetPropertyRef returns the PropertyRef field if non-nil, zero value otherwise.

### GetPropertyRefOk

`func (o *UiTableColumnCustomization) GetPropertyRefOk() (*string, bool)`

GetPropertyRefOk returns a tuple with the PropertyRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyRef

`func (o *UiTableColumnCustomization) SetPropertyRef(v string)`

SetPropertyRef sets PropertyRef field to given value.

### HasPropertyRef

`func (o *UiTableColumnCustomization) HasPropertyRef() bool`

HasPropertyRef returns a boolean if a field has been set.

### GetType

`func (o *UiTableColumnCustomization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UiTableColumnCustomization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UiTableColumnCustomization) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UiTableColumnCustomization) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


