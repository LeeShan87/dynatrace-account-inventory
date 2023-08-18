# FdpTagCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Predicate** | [**FdpTagPredicate**](FdpTagPredicate.md) |  | 
**TagKey** | **string** | The key of the tag to be checked. | 

## Methods

### NewFdpTagCondition

`func NewFdpTagCondition(predicate FdpTagPredicate, tagKey string, ) *FdpTagCondition`

NewFdpTagCondition instantiates a new FdpTagCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFdpTagConditionWithDefaults

`func NewFdpTagConditionWithDefaults() *FdpTagCondition`

NewFdpTagConditionWithDefaults instantiates a new FdpTagCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredicate

`func (o *FdpTagCondition) GetPredicate() FdpTagPredicate`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *FdpTagCondition) GetPredicateOk() (*FdpTagPredicate, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *FdpTagCondition) SetPredicate(v FdpTagPredicate)`

SetPredicate sets Predicate field to given value.


### GetTagKey

`func (o *FdpTagCondition) GetTagKey() string`

GetTagKey returns the TagKey field if non-nil, zero value otherwise.

### GetTagKeyOk

`func (o *FdpTagCondition) GetTagKeyOk() (*string, bool)`

GetTagKeyOk returns a tuple with the TagKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagKey

`func (o *FdpTagCondition) SetTagKey(v string)`

SetTagKey sets TagKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


