# NetworkErrorsDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomThresholds** | Pointer to [**NetworkErrorsThresholds**](NetworkErrorsThresholds.md) |  | [optional] 
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 

## Methods

### NewNetworkErrorsDetectionConfig

`func NewNetworkErrorsDetectionConfig(enabled bool, ) *NetworkErrorsDetectionConfig`

NewNetworkErrorsDetectionConfig instantiates a new NetworkErrorsDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkErrorsDetectionConfigWithDefaults

`func NewNetworkErrorsDetectionConfigWithDefaults() *NetworkErrorsDetectionConfig`

NewNetworkErrorsDetectionConfigWithDefaults instantiates a new NetworkErrorsDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomThresholds

`func (o *NetworkErrorsDetectionConfig) GetCustomThresholds() NetworkErrorsThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *NetworkErrorsDetectionConfig) GetCustomThresholdsOk() (*NetworkErrorsThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *NetworkErrorsDetectionConfig) SetCustomThresholds(v NetworkErrorsThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *NetworkErrorsDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkErrorsDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkErrorsDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkErrorsDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


