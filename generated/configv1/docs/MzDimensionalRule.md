# MzDimensionalRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppliesTo** | **string** | The target of the rule. | 
**Conditions** | [**[]MzDimensionalRuleCondition**](MzDimensionalRuleCondition.md) | A list of conditions for the management zone.    The management zone applies only if **all** conditions are fulfilled. | 
**Enabled** | **bool** | The rule is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 

## Methods

### NewMzDimensionalRule

`func NewMzDimensionalRule(appliesTo string, conditions []MzDimensionalRuleCondition, enabled bool, ) *MzDimensionalRule`

NewMzDimensionalRule instantiates a new MzDimensionalRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMzDimensionalRuleWithDefaults

`func NewMzDimensionalRuleWithDefaults() *MzDimensionalRule`

NewMzDimensionalRuleWithDefaults instantiates a new MzDimensionalRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppliesTo

`func (o *MzDimensionalRule) GetAppliesTo() string`

GetAppliesTo returns the AppliesTo field if non-nil, zero value otherwise.

### GetAppliesToOk

`func (o *MzDimensionalRule) GetAppliesToOk() (*string, bool)`

GetAppliesToOk returns a tuple with the AppliesTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppliesTo

`func (o *MzDimensionalRule) SetAppliesTo(v string)`

SetAppliesTo sets AppliesTo field to given value.


### GetConditions

`func (o *MzDimensionalRule) GetConditions() []MzDimensionalRuleCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MzDimensionalRule) GetConditionsOk() (*[]MzDimensionalRuleCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MzDimensionalRule) SetConditions(v []MzDimensionalRuleCondition)`

SetConditions sets Conditions field to given value.


### GetEnabled

`func (o *MzDimensionalRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MzDimensionalRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MzDimensionalRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


