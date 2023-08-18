# EntitySelectorBasedMzRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | The rule is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | [optional] 
**EntitySelector** | **string** | The entity selector string, by which the entities are selected. | 

## Methods

### NewEntitySelectorBasedMzRule

`func NewEntitySelectorBasedMzRule(entitySelector string, ) *EntitySelectorBasedMzRule`

NewEntitySelectorBasedMzRule instantiates a new EntitySelectorBasedMzRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitySelectorBasedMzRuleWithDefaults

`func NewEntitySelectorBasedMzRuleWithDefaults() *EntitySelectorBasedMzRule`

NewEntitySelectorBasedMzRuleWithDefaults instantiates a new EntitySelectorBasedMzRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *EntitySelectorBasedMzRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EntitySelectorBasedMzRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EntitySelectorBasedMzRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EntitySelectorBasedMzRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEntitySelector

`func (o *EntitySelectorBasedMzRule) GetEntitySelector() string`

GetEntitySelector returns the EntitySelector field if non-nil, zero value otherwise.

### GetEntitySelectorOk

`func (o *EntitySelectorBasedMzRule) GetEntitySelectorOk() (*string, bool)`

GetEntitySelectorOk returns a tuple with the EntitySelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitySelector

`func (o *EntitySelectorBasedMzRule) SetEntitySelector(v string)`

SetEntitySelector sets EntitySelector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


