# ModificationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deletable** | Pointer to **bool** | If settings value can be deleted | [optional] 
**First** | Pointer to **bool** | If non-moveable settings value is in the first group of non-moveable settings, or in the last (start or end of list) | [optional] 
**Modifiable** | Pointer to **bool** | If settings value can be modified | [optional] 
**ModifiablePaths** | Pointer to **[]string** | Property paths which are modifiable, regardless of the state of &#x60;modifiable&#x60; | [optional] 
**Movable** | Pointer to **bool** | If settings value can be moved/reordered. Only applicable for ordered list schema | [optional] 
**NonModifiablePaths** | Pointer to **[]string** | Property paths which are not modifiable, when &#x60;modifiable&#x60; is true | [optional] 

## Methods

### NewModificationInfo

`func NewModificationInfo() *ModificationInfo`

NewModificationInfo instantiates a new ModificationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModificationInfoWithDefaults

`func NewModificationInfoWithDefaults() *ModificationInfo`

NewModificationInfoWithDefaults instantiates a new ModificationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletable

`func (o *ModificationInfo) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *ModificationInfo) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *ModificationInfo) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *ModificationInfo) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetFirst

`func (o *ModificationInfo) GetFirst() bool`

GetFirst returns the First field if non-nil, zero value otherwise.

### GetFirstOk

`func (o *ModificationInfo) GetFirstOk() (*bool, bool)`

GetFirstOk returns a tuple with the First field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirst

`func (o *ModificationInfo) SetFirst(v bool)`

SetFirst sets First field to given value.

### HasFirst

`func (o *ModificationInfo) HasFirst() bool`

HasFirst returns a boolean if a field has been set.

### GetModifiable

`func (o *ModificationInfo) GetModifiable() bool`

GetModifiable returns the Modifiable field if non-nil, zero value otherwise.

### GetModifiableOk

`func (o *ModificationInfo) GetModifiableOk() (*bool, bool)`

GetModifiableOk returns a tuple with the Modifiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiable

`func (o *ModificationInfo) SetModifiable(v bool)`

SetModifiable sets Modifiable field to given value.

### HasModifiable

`func (o *ModificationInfo) HasModifiable() bool`

HasModifiable returns a boolean if a field has been set.

### GetModifiablePaths

`func (o *ModificationInfo) GetModifiablePaths() []string`

GetModifiablePaths returns the ModifiablePaths field if non-nil, zero value otherwise.

### GetModifiablePathsOk

`func (o *ModificationInfo) GetModifiablePathsOk() (*[]string, bool)`

GetModifiablePathsOk returns a tuple with the ModifiablePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiablePaths

`func (o *ModificationInfo) SetModifiablePaths(v []string)`

SetModifiablePaths sets ModifiablePaths field to given value.

### HasModifiablePaths

`func (o *ModificationInfo) HasModifiablePaths() bool`

HasModifiablePaths returns a boolean if a field has been set.

### GetMovable

`func (o *ModificationInfo) GetMovable() bool`

GetMovable returns the Movable field if non-nil, zero value otherwise.

### GetMovableOk

`func (o *ModificationInfo) GetMovableOk() (*bool, bool)`

GetMovableOk returns a tuple with the Movable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMovable

`func (o *ModificationInfo) SetMovable(v bool)`

SetMovable sets Movable field to given value.

### HasMovable

`func (o *ModificationInfo) HasMovable() bool`

HasMovable returns a boolean if a field has been set.

### GetNonModifiablePaths

`func (o *ModificationInfo) GetNonModifiablePaths() []string`

GetNonModifiablePaths returns the NonModifiablePaths field if non-nil, zero value otherwise.

### GetNonModifiablePathsOk

`func (o *ModificationInfo) GetNonModifiablePathsOk() (*[]string, bool)`

GetNonModifiablePathsOk returns a tuple with the NonModifiablePaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonModifiablePaths

`func (o *ModificationInfo) SetNonModifiablePaths(v []string)`

SetNonModifiablePaths sets NonModifiablePaths field to given value.

### HasNonModifiablePaths

`func (o *ModificationInfo) HasNonModifiablePaths() bool`

HasNonModifiablePaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


