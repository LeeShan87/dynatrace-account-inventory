# MzRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]EntityRuleEngineCondition**](EntityRuleEngineCondition.md) | A list of matching rules for the management zone.   The management zone applies only if **all** conditions are fulfilled. | 
**Enabled** | **bool** | The rule is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**PropagationTypes** | Pointer to **[]string** | How to apply the management zone to underlying entities:   * &#x60;SERVICE_TO_HOST_LIKE&#x60;: Apply to underlying hosts of matching services.  * &#x60;SERVICE_TO_PROCESS_GROUP_LIKE&#x60;: Apply to underlying process groups of matching services.  * &#x60;PROCESS_GROUP_TO_HOST&#x60;: Apply to underlying hosts of matching process groups.  * &#x60;PROCESS_GROUP_TO_SERVICE&#x60;: Apply to all services provided by matching process groups.  * &#x60;HOST_TO_PROCESS_GROUP_INSTANCE&#x60;: Apply to processes running on matching hosts.  * &#x60;CUSTOM_DEVICE_GROUP_TO_CUSTOM_DEVICE&#x60;: Apply to custom devices in matching custom device groups.  * &#x60;AZURE_TO_PG&#x60;: Apply to process groups connected to matching Azure entities.  * &#x60;AZURE_TO_SERVICE&#x60;: Apply to services provided by matching Azure entities. | [optional] 
**Type** | **string** | The type of Dynatrace entities the management zone can be applied to. | 

## Methods

### NewMzRule

`func NewMzRule(conditions []EntityRuleEngineCondition, enabled bool, type_ string, ) *MzRule`

NewMzRule instantiates a new MzRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMzRuleWithDefaults

`func NewMzRuleWithDefaults() *MzRule`

NewMzRuleWithDefaults instantiates a new MzRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *MzRule) GetConditions() []EntityRuleEngineCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *MzRule) GetConditionsOk() (*[]EntityRuleEngineCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *MzRule) SetConditions(v []EntityRuleEngineCondition)`

SetConditions sets Conditions field to given value.


### GetEnabled

`func (o *MzRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MzRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MzRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPropagationTypes

`func (o *MzRule) GetPropagationTypes() []string`

GetPropagationTypes returns the PropagationTypes field if non-nil, zero value otherwise.

### GetPropagationTypesOk

`func (o *MzRule) GetPropagationTypesOk() (*[]string, bool)`

GetPropagationTypesOk returns a tuple with the PropagationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagationTypes

`func (o *MzRule) SetPropagationTypes(v []string)`

SetPropagationTypes sets PropagationTypes field to given value.

### HasPropagationTypes

`func (o *MzRule) HasPropagationTypes() bool`

HasPropagationTypes returns a boolean if a field has been set.

### GetType

`func (o *MzRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MzRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MzRule) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


