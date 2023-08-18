# MzDimensionalRuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionType** | **string** | The type of the condition. | 
**Key** | **string** | The reference value for comparison.   For conditions of the &#x60;DIMENSION&#x60; type, specify the key here. | 
**RuleMatcher** | **string** | How we compare the values. | 
**Value** | Pointer to **string** | The value of the dimension.   Only applicable when the **conditionType** is set to &#x60;DIMENSION&#x60;. | [optional] 

## Methods

### NewMzDimensionalRuleCondition

`func NewMzDimensionalRuleCondition(conditionType string, key string, ruleMatcher string, ) *MzDimensionalRuleCondition`

NewMzDimensionalRuleCondition instantiates a new MzDimensionalRuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMzDimensionalRuleConditionWithDefaults

`func NewMzDimensionalRuleConditionWithDefaults() *MzDimensionalRuleCondition`

NewMzDimensionalRuleConditionWithDefaults instantiates a new MzDimensionalRuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionType

`func (o *MzDimensionalRuleCondition) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *MzDimensionalRuleCondition) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *MzDimensionalRuleCondition) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.


### GetKey

`func (o *MzDimensionalRuleCondition) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MzDimensionalRuleCondition) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MzDimensionalRuleCondition) SetKey(v string)`

SetKey sets Key field to given value.


### GetRuleMatcher

`func (o *MzDimensionalRuleCondition) GetRuleMatcher() string`

GetRuleMatcher returns the RuleMatcher field if non-nil, zero value otherwise.

### GetRuleMatcherOk

`func (o *MzDimensionalRuleCondition) GetRuleMatcherOk() (*string, bool)`

GetRuleMatcherOk returns a tuple with the RuleMatcher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleMatcher

`func (o *MzDimensionalRuleCondition) SetRuleMatcher(v string)`

SetRuleMatcher sets RuleMatcher field to given value.


### GetValue

`func (o *MzDimensionalRuleCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MzDimensionalRuleCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MzDimensionalRuleCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *MzDimensionalRuleCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


