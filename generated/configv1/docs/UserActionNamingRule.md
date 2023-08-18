# UserActionNamingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]UserActionNamingRuleCondition**](UserActionNamingRuleCondition.md) | Defines the conditions when the naming rule should apply. | [optional] 
**Template** | **string** | Naming pattern. Use Curly brackets &#x60;{}&#x60; to select placeholders. | 
**UseOrConditions** | Pointer to **bool** | If set to &#x60;true&#x60; the conditions will be connected by logical OR instead of logical AND. | [optional] 

## Methods

### NewUserActionNamingRule

`func NewUserActionNamingRule(template string, ) *UserActionNamingRule`

NewUserActionNamingRule instantiates a new UserActionNamingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActionNamingRuleWithDefaults

`func NewUserActionNamingRuleWithDefaults() *UserActionNamingRule`

NewUserActionNamingRuleWithDefaults instantiates a new UserActionNamingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *UserActionNamingRule) GetConditions() []UserActionNamingRuleCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *UserActionNamingRule) GetConditionsOk() (*[]UserActionNamingRuleCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *UserActionNamingRule) SetConditions(v []UserActionNamingRuleCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *UserActionNamingRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetTemplate

`func (o *UserActionNamingRule) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UserActionNamingRule) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UserActionNamingRule) SetTemplate(v string)`

SetTemplate sets Template field to given value.


### GetUseOrConditions

`func (o *UserActionNamingRule) GetUseOrConditions() bool`

GetUseOrConditions returns the UseOrConditions field if non-nil, zero value otherwise.

### GetUseOrConditionsOk

`func (o *UserActionNamingRule) GetUseOrConditionsOk() (*bool, bool)`

GetUseOrConditionsOk returns a tuple with the UseOrConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOrConditions

`func (o *UserActionNamingRule) SetUseOrConditions(v bool)`

SetUseOrConditions sets UseOrConditions field to given value.

### HasUseOrConditions

`func (o *UserActionNamingRule) HasUseOrConditions() bool`

HasUseOrConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


