# ActiveGateTokenEnforcement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoEnforced** | Pointer to **bool** | If &#x60;true&#x60;, ActiveGate tokens are enforced automatically. | [optional] 
**AutoEnforcementEstimation** | Pointer to [**Duration**](Duration.md) |  | [optional] 
**ManualEnforced** | Pointer to **bool** | If &#x60;true&#x60;, ActiveGate tokens are manually enforced by user. | [optional] 

## Methods

### NewActiveGateTokenEnforcement

`func NewActiveGateTokenEnforcement() *ActiveGateTokenEnforcement`

NewActiveGateTokenEnforcement instantiates a new ActiveGateTokenEnforcement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveGateTokenEnforcementWithDefaults

`func NewActiveGateTokenEnforcementWithDefaults() *ActiveGateTokenEnforcement`

NewActiveGateTokenEnforcementWithDefaults instantiates a new ActiveGateTokenEnforcement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoEnforced

`func (o *ActiveGateTokenEnforcement) GetAutoEnforced() bool`

GetAutoEnforced returns the AutoEnforced field if non-nil, zero value otherwise.

### GetAutoEnforcedOk

`func (o *ActiveGateTokenEnforcement) GetAutoEnforcedOk() (*bool, bool)`

GetAutoEnforcedOk returns a tuple with the AutoEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnforced

`func (o *ActiveGateTokenEnforcement) SetAutoEnforced(v bool)`

SetAutoEnforced sets AutoEnforced field to given value.

### HasAutoEnforced

`func (o *ActiveGateTokenEnforcement) HasAutoEnforced() bool`

HasAutoEnforced returns a boolean if a field has been set.

### GetAutoEnforcementEstimation

`func (o *ActiveGateTokenEnforcement) GetAutoEnforcementEstimation() Duration`

GetAutoEnforcementEstimation returns the AutoEnforcementEstimation field if non-nil, zero value otherwise.

### GetAutoEnforcementEstimationOk

`func (o *ActiveGateTokenEnforcement) GetAutoEnforcementEstimationOk() (*Duration, bool)`

GetAutoEnforcementEstimationOk returns a tuple with the AutoEnforcementEstimation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnforcementEstimation

`func (o *ActiveGateTokenEnforcement) SetAutoEnforcementEstimation(v Duration)`

SetAutoEnforcementEstimation sets AutoEnforcementEstimation field to given value.

### HasAutoEnforcementEstimation

`func (o *ActiveGateTokenEnforcement) HasAutoEnforcementEstimation() bool`

HasAutoEnforcementEstimation returns a boolean if a field has been set.

### GetManualEnforced

`func (o *ActiveGateTokenEnforcement) GetManualEnforced() bool`

GetManualEnforced returns the ManualEnforced field if non-nil, zero value otherwise.

### GetManualEnforcedOk

`func (o *ActiveGateTokenEnforcement) GetManualEnforcedOk() (*bool, bool)`

GetManualEnforcedOk returns a tuple with the ManualEnforced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualEnforced

`func (o *ActiveGateTokenEnforcement) SetManualEnforced(v bool)`

SetManualEnforced sets ManualEnforced field to given value.

### HasManualEnforced

`func (o *ActiveGateTokenEnforcement) HasManualEnforced() bool`

HasManualEnforced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


