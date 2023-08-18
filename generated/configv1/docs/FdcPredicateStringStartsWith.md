# FdcPredicateStringStartsWith

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreCase** | **bool** | The condition is case sensitive (&#x60;false&#x60;) or case insensitive (&#x60;true&#x60;).   If not set, then &#x60;false&#x60; is used, making the condition case sensitive. | 
**Values** | **[]string** | A list of reference values. The condition is fulfilled when the attribute (which is a string) start with *any* of these. | 

## Methods

### NewFdcPredicateStringStartsWith

`func NewFdcPredicateStringStartsWith(ignoreCase bool, values []string, ) *FdcPredicateStringStartsWith`

NewFdcPredicateStringStartsWith instantiates a new FdcPredicateStringStartsWith object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdcPredicateStringStartsWithWithDefaults

`func NewFdcPredicateStringStartsWithWithDefaults() *FdcPredicateStringStartsWith`

NewFdcPredicateStringStartsWithWithDefaults instantiates a new FdcPredicateStringStartsWith object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreCase

`func (o *FdcPredicateStringStartsWith) GetIgnoreCase() bool`

GetIgnoreCase returns the IgnoreCase field if non-nil, zero value otherwise.

### GetIgnoreCaseOk

`func (o *FdcPredicateStringStartsWith) GetIgnoreCaseOk() (*bool, bool)`

GetIgnoreCaseOk returns a tuple with the IgnoreCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCase

`func (o *FdcPredicateStringStartsWith) SetIgnoreCase(v bool)`

SetIgnoreCase sets IgnoreCase field to given value.


### GetValues

`func (o *FdcPredicateStringStartsWith) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FdcPredicateStringStartsWith) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FdcPredicateStringStartsWith) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


