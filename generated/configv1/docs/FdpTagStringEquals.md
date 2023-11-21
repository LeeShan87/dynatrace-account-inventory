# FdpTagStringEquals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreCase** | Pointer to **bool** |  | [optional] 
**Negated** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **string** | The reference value. The condition is fulfilled when the tag (which is a string) equals this value. | [optional] 

## Methods

### NewFdpTagStringEquals

`func NewFdpTagStringEquals() *FdpTagStringEquals`

NewFdpTagStringEquals instantiates a new FdpTagStringEquals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdpTagStringEqualsWithDefaults

`func NewFdpTagStringEqualsWithDefaults() *FdpTagStringEquals`

NewFdpTagStringEqualsWithDefaults instantiates a new FdpTagStringEquals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreCase

`func (o *FdpTagStringEquals) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *FdpTagStringEquals) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *FdpTagStringEquals) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.

### HasIgnoreCase

`func (o *FdpTagStringEquals) HasIgnoreCase() bool`

HasIgnoreCase returns a boolean if a field has been set.

### GetNegated

`func (o *FdpTagStringEquals) GetNegated() bool`

GetNegated returns the Negated field if non-nil, zero value otherwise.

### GetNegatedOk

`func (o *FdpTagStringEquals) GetNegatedOk() (*bool, bool)`

GetNegatedOk returns a tuple with the Negated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegated

`func (o *FdpTagStringEquals) SetNegated(v bool)`

SetNegated sets Negated field to given value.

### HasNegated

`func (o *FdpTagStringEquals) HasNegated() bool`

HasNegated returns a boolean if a field has been set.

### GetValue

`func (o *FdpTagStringEquals) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FdpTagStringEquals) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FdpTagStringEquals) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FdpTagStringEquals) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


