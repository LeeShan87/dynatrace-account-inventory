# OutOfThreadsDetectionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomThresholds** | Pointer to [**OutOfThreadsThresholds**](OutOfThreadsThresholds.md) |  | [optional] 
**Enabled** | **bool** | The detection is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 

## Methods

### NewOutOfThreadsDetectionConfig

`func NewOutOfThreadsDetectionConfig(enabled bool, ) *OutOfThreadsDetectionConfig`

NewOutOfThreadsDetectionConfig instantiates a new OutOfThreadsDetectionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutOfThreadsDetectionConfigWithDefaults

`func NewOutOfThreadsDetectionConfigWithDefaults() *OutOfThreadsDetectionConfig`

NewOutOfThreadsDetectionConfigWithDefaults instantiates a new OutOfThreadsDetectionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomThresholds

`func (o *OutOfThreadsDetectionConfig) GetCustomThresholds() OutOfThreadsThresholds`

GetCustomThresholds returns the CustomThresholds field if non-nil, zero value otherwise.

### GetCustomThresholdsOk

`func (o *OutOfThreadsDetectionConfig) GetCustomThresholdsOk() (*OutOfThreadsThresholds, bool)`

GetCustomThresholdsOk returns a tuple with the CustomThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomThresholds

`func (o *OutOfThreadsDetectionConfig) SetCustomThresholds(v OutOfThreadsThresholds)`

SetCustomThresholds sets CustomThresholds field to given value.

### HasCustomThresholds

`func (o *OutOfThreadsDetectionConfig) HasCustomThresholds() bool`

HasCustomThresholds returns a boolean if a field has been set.

### GetEnabled

`func (o *OutOfThreadsDetectionConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OutOfThreadsDetectionConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OutOfThreadsDetectionConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


