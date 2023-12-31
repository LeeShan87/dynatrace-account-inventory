# RdsHighMemoryDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomThresholds** | Pointer to [**RdsHighMemoryThresholds**](RdsHighMemoryThresholds.md) |  | [optional] 
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 

## Methods

### NewRdsHighMemoryDetectionConfig

`func NewRdsHighMemoryDetectionConfig(enabled bool, ) *RdsHighMemoryDetectionConfig`

NewRdsHighMemoryDetectionConfig instantiates a new RdsHighMemoryDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRdsHighMemoryDetectionConfigWithDefaults

`func NewRdsHighMemoryDetectionConfigWithDefaults() *RdsHighMemoryDetectionConfig`

NewRdsHighMemoryDetectionConfigWithDefaults instantiates a new RdsHighMemoryDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomThresholds

`func (o *RdsHighMemoryDetectionConfig) GetCustomThresholds() RdsHighMemoryThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *RdsHighMemoryDetectionConfig) GetCustomThresholdsOk() (*RdsHighMemoryThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *RdsHighMemoryDetectionConfig) SetCustomThresholds(v RdsHighMemoryThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *RdsHighMemoryDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.

### GetEnabled

`func (o *RdsHighMemoryDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RdsHighMemoryDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RdsHighMemoryDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


