# FailureDetectionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]FailureDetectionCondition**](FailureDetectionCondition.md) | A list of conditions of the rule.   The rule applies when **all** conditions are fulfilled. | 
**Description** | Pointer to **string** | A short description of the rule. | [optional] 
**Enabled** | **bool** | The rule is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**FdpId** | **string** | The failure detection parameter (FDP) set of the rule.   Specify the ID of the set here. The FDP set must exist at the time of rule creation. | 
**Id** | Pointer to **string** | The ID of the rule. | [optional] 
**Name** | Pointer to **string** | The display name of the rule.   The length of the name is limited to 150 characters. | [optional] 

## Methods

### NewFailureDetectionRule

`func NewFailureDetectionRule(conditions []FailureDetectionCondition, enabled bool, fdpId string, ) *FailureDetectionRule`

NewFailureDetectionRule instantiates a new FailureDetectionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureDetectionRuleWithDefaults

`func NewFailureDetectionRuleWithDefaults() *FailureDetectionRule`

NewFailureDetectionRuleWithDefaults instantiates a new FailureDetectionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *FailureDetectionRule) GetConditions() []FailureDetectionCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *FailureDetectionRule) GetConditionsOk() (*[]FailureDetectionCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *FailureDetectionRule) SetConditions(v []FailureDetectionCondition)`

SetConditions sets Conditions field to given value.


### GetDescription

`func (o *FailureDetectionRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FailureDetectionRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FailureDetectionRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FailureDetectionRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *FailureDetectionRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FailureDetectionRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FailureDetectionRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFdpId

`func (o *FailureDetectionRule) GetFdpId() string`

GetFdpId returns the FdpId field if non-nil, zero value otherwise.

### GetFdpIdOk

`func (o *FailureDetectionRule) GetFdpIdOk() (*string, bool)`

GetFdpIdOk returns a tuple with the FdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdpId

`func (o *FailureDetectionRule) SetFdpId(v string)`

SetFdpId sets FdpId field to given value.


### GetId

`func (o *FailureDetectionRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FailureDetectionRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FailureDetectionRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FailureDetectionRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FailureDetectionRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FailureDetectionRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FailureDetectionRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FailureDetectionRule) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


