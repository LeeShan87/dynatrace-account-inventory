# EntryPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]EntryPoint**](EntryPoint.md) | A list of entry points. | [optional] [readonly] 
**Truncated** | Pointer to **bool** | Indicates whether the list of entry points was truncated or not. | [optional] [readonly] 

## Methods

### NewEntryPoints

`func NewEntryPoints() *EntryPoints`

NewEntryPoints instantiates a new EntryPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntryPointsWithDefaults

`func NewEntryPointsWithDefaults() *EntryPoints`

NewEntryPointsWithDefaults instantiates a new EntryPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EntryPoints) GetItems() []EntryPoint`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EntryPoints) GetItemsOk() (*[]EntryPoint, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EntryPoints) SetItems(v []EntryPoint)`

SetItems sets Items field to given value.

### HasItems

`func (o *EntryPoints) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTruncated

`func (o *EntryPoints) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *EntryPoints) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *EntryPoints) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *EntryPoints) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


