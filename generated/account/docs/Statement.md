# Statement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | The effect of the policy (for example, allow something). | 
**Service** | **string** | The service to which the policy applies. | 
**Permissions** | **[]string** | A list of granted permissions. | 
**Conditions** | [**[]Condition**](Condition.md) | A list of conditions limiting the granted permissions. | 

## Methods

### NewStatement

`func NewStatement(effect string, service string, permissions []string, conditions []Condition, ) *Statement`

NewStatement instantiates a new Statement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatementWithDefaults

`func NewStatementWithDefaults() *Statement`

NewStatementWithDefaults instantiates a new Statement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *Statement) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *Statement) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *Statement) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetService

`func (o *Statement) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *Statement) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *Statement) SetService(v string)`

SetService sets Service field to given value.


### GetPermissions

`func (o *Statement) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *Statement) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *Statement) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetConditions

`func (o *Statement) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Statement) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Statement) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


