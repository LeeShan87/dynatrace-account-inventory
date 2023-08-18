# SloBurnRateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurnRateVisualizationEnabled** | Pointer to **bool** | The error budget burn rate calculation is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).   In case of &#x60;false&#x60;, no calculated values will be present here.   If not defined, the error budget burn rate calculation is disabled by default. | [optional] 
**FastBurnThreshold** | Pointer to **float64** | The threshold between a slow and a fast burn rate. | [optional] 

## Methods

### NewSloBurnRateConfig

`func NewSloBurnRateConfig() *SloBurnRateConfig`

NewSloBurnRateConfig instantiates a new SloBurnRateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloBurnRateConfigWithDefaults

`func NewSloBurnRateConfigWithDefaults() *SloBurnRateConfig`

NewSloBurnRateConfigWithDefaults instantiates a new SloBurnRateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurnRateVisualizationEnabled

`func (o *SloBurnRateConfig) GetBurnRateVisualizationEnabled() bool`

GetBurnRateVisualizationEnabled returns the BurnRateVisualizationEnabled field if non-nil, zero value otherwise.

### GetBurnRateVisualizationEnabledOk

`func (o *SloBurnRateConfig) GetBurnRateVisualizationEnabledOk() (*bool, bool)`

GetBurnRateVisualizationEnabledOk returns a tuple with the BurnRateVisualizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnRateVisualizationEnabled

`func (o *SloBurnRateConfig) SetBurnRateVisualizationEnabled(v bool)`

SetBurnRateVisualizationEnabled sets BurnRateVisualizationEnabled field to given value.

### HasBurnRateVisualizationEnabled

`func (o *SloBurnRateConfig) HasBurnRateVisualizationEnabled() bool`

HasBurnRateVisualizationEnabled returns a boolean if a field has been set.

### GetFastBurnThreshold

`func (o *SloBurnRateConfig) GetFastBurnThreshold() float64`

GetFastBurnThreshold returns the FastBurnThreshold field if non-nil, zero value otherwise.

### GetFastBurnThresholdOk

`func (o *SloBurnRateConfig) GetFastBurnThresholdOk() (*float64, bool)`

GetFastBurnThresholdOk returns a tuple with the FastBurnThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastBurnThreshold

`func (o *SloBurnRateConfig) SetFastBurnThreshold(v float64)`

SetFastBurnThreshold sets FastBurnThreshold field to given value.

### HasFastBurnThreshold

`func (o *SloBurnRateConfig) HasFastBurnThreshold() bool`

HasFastBurnThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


