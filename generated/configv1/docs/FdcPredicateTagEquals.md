# FdcPredicateTagEquals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]string** | A list of reference tag keys.   The condition is fulfilled when the tag matches *any* key:value pair, consisting of a key and a value that are at the **same position** in the respective list. | [optional] 
**Values** | Pointer to **[]string** | A list of reference tag values.   The condition is fulfilled when the tag matches *any* key:value pair, consisting of a key and a value that are at the **same position** in the respective list. | [optional] 

## Methods

### NewFdcPredicateTagEquals

`func NewFdcPredicateTagEquals() *FdcPredicateTagEquals`

NewFdcPredicateTagEquals instantiates a new FdcPredicateTagEquals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdcPredicateTagEqualsWithDefaults

`func NewFdcPredicateTagEqualsWithDefaults() *FdcPredicateTagEquals`

NewFdcPredicateTagEqualsWithDefaults instantiates a new FdcPredicateTagEquals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *FdcPredicateTagEquals) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *FdcPredicateTagEquals) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *FdcPredicateTagEquals) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *FdcPredicateTagEquals) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetValues

`func (o *FdcPredicateTagEquals) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FdcPredicateTagEquals) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FdcPredicateTagEquals) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *FdcPredicateTagEquals) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


