# ElbHighConnectionErrorsDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomThresholds** | Pointer to [**ElbHighConnectionErrorsThresholds**](ElbHighConnectionErrorsThresholds.md) |  | [optional] 
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 

## Methods

### NewElbHighConnectionErrorsDetectionConfig

`func NewElbHighConnectionErrorsDetectionConfig(enabled bool, ) *ElbHighConnectionErrorsDetectionConfig`

NewElbHighConnectionErrorsDetectionConfig instantiates a new ElbHighConnectionErrorsDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElbHighConnectionErrorsDetectionConfigWithDefaults

`func NewElbHighConnectionErrorsDetectionConfigWithDefaults() *ElbHighConnectionErrorsDetectionConfig`

NewElbHighConnectionErrorsDetectionConfigWithDefaults instantiates a new ElbHighConnectionErrorsDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomThresholds

`func (o *ElbHighConnectionErrorsDetectionConfig) GetCustomThresholds() ElbHighConnectionErrorsThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *ElbHighConnectionErrorsDetectionConfig) GetCustomThresholdsOk() (*ElbHighConnectionErrorsThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *ElbHighConnectionErrorsDetectionConfig) SetCustomThresholds(v ElbHighConnectionErrorsThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *ElbHighConnectionErrorsDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.

### GetEnabled

`func (o *ElbHighConnectionErrorsDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ElbHighConnectionErrorsDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ElbHighConnectionErrorsDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


