# EntitySelectorBasedAutoTagRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | The rule is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**EntitySelector** | **string** | The entity selector string, by which the entities are selected. | 
**Normalization** | Pointer to **NullableString** | Changes applied to the value after applying the value format. Default is LEAVE_TEXT_AS_IS. | [optional] 
**ValueFormat** | Pointer to **string** | The value of the entity-selector-based auto-tag. If specified, the tag is used in the &#x60;name:valueFormat&#x60; format.   For example, you can extend the &#x60;Infrastructure&#x60; tag to &#x60;Infrastructure:Windows&#x60; and &#x60;Infrastructure:Linux&#x60;.    | [optional] 

## Methods

### NewEntitySelectorBasedAutoTagRule

`func NewEntitySelectorBasedAutoTagRule(entitySelector string, ) *EntitySelectorBasedAutoTagRule`

NewEntitySelectorBasedAutoTagRule instantiates a new EntitySelectorBasedAutoTagRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySelectorBasedAutoTagRuleWithDefaults

`func NewEntitySelectorBasedAutoTagRuleWithDefaults() *EntitySelectorBasedAutoTagRule`

NewEntitySelectorBasedAutoTagRuleWithDefaults instantiates a new EntitySelectorBasedAutoTagRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *EntitySelectorBasedAutoTagRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EntitySelectorBasedAutoTagRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EntitySelectorBasedAutoTagRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EntitySelectorBasedAutoTagRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntitySelector

`func (o *EntitySelectorBasedAutoTagRule) GetEntitySelector() string`

GetEntitySelector returns the EntitySelector field if non-nil, zero value otherwise.

### GetEntitySelectorOk

`func (o *EntitySelectorBasedAutoTagRule) GetEntitySelectorOk() (*string, bool)`

GetEntitySelectorOk returns a tuple with the EntitySelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelector

`func (o *EntitySelectorBasedAutoTagRule) SetEntitySelector(v string)`

SetEntitySelector sets EntitySelector field to given value.


### GetNormalization

`func (o *EntitySelectorBasedAutoTagRule) GetNormalization() string`

GetNormalization returns the Normalization field if non-nil, zero value otherwise.

### GetNormalizationOk

`func (o *EntitySelectorBasedAutoTagRule) GetNormalizationOk() (*string, bool)`

GetNormalizationOk returns a tuple with the Normalization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalization

`func (o *EntitySelectorBasedAutoTagRule) SetNormalization(v string)`

SetNormalization sets Normalization field to given value.

### HasNormalization

`func (o *EntitySelectorBasedAutoTagRule) HasNormalization() bool`

HasNormalization returns a boolean if a field has been set.

### SetNormalizationNil

`func (o *EntitySelectorBasedAutoTagRule) SetNormalizationNil(b bool)`

 SetNormalizationNil sets the value for Normalization to be an explicit nil

### UnsetNormalization
`func (o *EntitySelectorBasedAutoTagRule) UnsetNormalization()`

UnsetNormalization ensures that no value is present for Normalization, not even an explicit nil
### GetValueFormat

`func (o *EntitySelectorBasedAutoTagRule) GetValueFormat() string`

GetValueFormat returns the ValueFormat field if non-nil, zero value otherwise.

### GetValueFormatOk

`func (o *EntitySelectorBasedAutoTagRule) GetValueFormatOk() (*string, bool)`

GetValueFormatOk returns a tuple with the ValueFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueFormat

`func (o *EntitySelectorBasedAutoTagRule) SetValueFormat(v string)`

SetValueFormat sets ValueFormat field to given value.

### HasValueFormat

`func (o *EntitySelectorBasedAutoTagRule) HasValueFormat() bool`

HasValueFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


