# FailureDetectionCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **string** | The attribute to be checked. | [optional] 
**Predicate** | Pointer to [**FdcPredicate**](FdcPredicate.md) |  | [optional] 

## Methods

### NewFailureDetectionCondition

`func NewFailureDetectionCondition() *FailureDetectionCondition`

NewFailureDetectionCondition instantiates a new FailureDetectionCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureDetectionConditionWithDefaults

`func NewFailureDetectionConditionWithDefaults() *FailureDetectionCondition`

NewFailureDetectionConditionWithDefaults instantiates a new FailureDetectionCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *FailureDetectionCondition) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *FailureDetectionCondition) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *FailureDetectionCondition) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *FailureDetectionCondition) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetPredicate

`func (o *FailureDetectionCondition) GetPredicate() FdcPredicate`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *FailureDetectionCondition) GetPredicateOk() (*FdcPredicate, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *FailureDetectionCondition) SetPredicate(v FdcPredicate)`

SetPredicate sets Predicate field to given value.

### HasPredicate

`func (o *FailureDetectionCondition) HasPredicate() bool`

HasPredicate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


