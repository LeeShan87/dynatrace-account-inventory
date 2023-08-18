# MaskingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaskingRuleType** | **string** | The type of the masking rule. | 
**Selector** | **string** | The selector for the element or the attribute to be masked.   Specify a CSS expression for an element or a [regular expression](https://dt-url.net/k9e0iaq) for an attribute. | 
**UserInteractionHidden** | **bool** | Interactions with the element are (&#x60;true&#x60;) or are not (&#x60;false) masked. | 

## Methods

### NewMaskingRule

`func NewMaskingRule(maskingRuleType string, selector string, userInteractionHidden bool, ) *MaskingRule`

NewMaskingRule instantiates a new MaskingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaskingRuleWithDefaults

`func NewMaskingRuleWithDefaults() *MaskingRule`

NewMaskingRuleWithDefaults instantiates a new MaskingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaskingRuleType

`func (o *MaskingRule) GetMaskingRuleType() string`

GetMaskingRuleType returns the MaskingRuleType field if non-nil, zero value otherwise.

### GetMaskingRuleTypeOk

`func (o *MaskingRule) GetMaskingRuleTypeOk() (*string, bool)`

GetMaskingRuleTypeOk returns a tuple with the MaskingRuleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskingRuleType

`func (o *MaskingRule) SetMaskingRuleType(v string)`

SetMaskingRuleType sets MaskingRuleType field to given value.


### GetSelector

`func (o *MaskingRule) GetSelector() string`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *MaskingRule) GetSelectorOk() (*string, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *MaskingRule) SetSelector(v string)`

SetSelector sets Selector field to given value.


### GetUserInteractionHidden

`func (o *MaskingRule) GetUserInteractionHidden() bool`

GetUserInteractionHidden returns the UserInteractionHidden field if non-nil, zero value otherwise.

### GetUserInteractionHiddenOk

`func (o *MaskingRule) GetUserInteractionHiddenOk() (*bool, bool)`

GetUserInteractionHiddenOk returns a tuple with the UserInteractionHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInteractionHidden

`func (o *MaskingRule) SetUserInteractionHidden(v bool)`

SetUserInteractionHidden sets UserInteractionHidden field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


