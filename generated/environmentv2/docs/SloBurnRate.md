# SloBurnRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurnRateType** | Pointer to **string** | The calculated burn rate type.   Has a value of &#39;FAST&#39;, &#39;SLOW&#39; or &#39;NONE&#39;. | [optional] 
**BurnRateValue** | Pointer to **float64** | The burn rate of the SLO, calculated for the last hour. | [optional] 
**BurnRateVisualizationEnabled** | **bool** | The error budget burn rate calculation is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).   In case of &#x60;false&#x60;, no calculated values will be present here. | 
**EstimatedTimeToConsumeErrorBudget** | Pointer to **float64** | The estimated time left to consume the error budget in hours. | [optional] 
**FastBurnThreshold** | Pointer to **float64** | The threshold between a slow and a fast burn rate. | [optional] 
**SloValue** | Pointer to **float64** | The calculated value of the SLO for the timeframe chosen for the burn rate calculation. | [optional] 

## Methods

### NewSloBurnRate

`func NewSloBurnRate(burnRateVisualizationEnabled bool, ) *SloBurnRate`

NewSloBurnRate instantiates a new SloBurnRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloBurnRateWithDefaults

`func NewSloBurnRateWithDefaults() *SloBurnRate`

NewSloBurnRateWithDefaults instantiates a new SloBurnRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurnRateType

`func (o *SloBurnRate) GetBurnRateType() string`

GetBurnRateType returns the BurnRateType field if non-nil, zero value otherwise.

### GetBurnRateTypeOk

`func (o *SloBurnRate) GetBurnRateTypeOk() (*string, bool)`

GetBurnRateTypeOk returns a tuple with the BurnRateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnRateType

`func (o *SloBurnRate) SetBurnRateType(v string)`

SetBurnRateType sets BurnRateType field to given value.

### HasBurnRateType

`func (o *SloBurnRate) HasBurnRateType() bool`

HasBurnRateType returns a boolean if a field has been set.

### GetBurnRateValue

`func (o *SloBurnRate) GetBurnRateValue() float64`

GetBurnRateValue returns the BurnRateValue field if non-nil, zero value otherwise.

### GetBurnRateValueOk

`func (o *SloBurnRate) GetBurnRateValueOk() (*float64, bool)`

GetBurnRateValueOk returns a tuple with the BurnRateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnRateValue

`func (o *SloBurnRate) SetBurnRateValue(v float64)`

SetBurnRateValue sets BurnRateValue field to given value.

### HasBurnRateValue

`func (o *SloBurnRate) HasBurnRateValue() bool`

HasBurnRateValue returns a boolean if a field has been set.

### GetBurnRateVisualizationEnabled

`func (o *SloBurnRate) GetBurnRateVisualizationEnabled() bool`

GetBurnRateVisualizationEnabled returns the BurnRateVisualizationEnabled field if non-nil, zero value otherwise.

### GetBurnRateVisualizationEnabledOk

`func (o *SloBurnRate) GetBurnRateVisualizationEnabledOk() (*bool, bool)`

GetBurnRateVisualizationEnabledOk returns a tuple with the BurnRateVisualizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnRateVisualizationEnabled

`func (o *SloBurnRate) SetBurnRateVisualizationEnabled(v bool)`

SetBurnRateVisualizationEnabled sets BurnRateVisualizationEnabled field to given value.


### GetEstimatedTimeToConsumeErrorBudget

`func (o *SloBurnRate) GetEstimatedTimeToConsumeErrorBudget() float64`

GetEstimatedTimeToConsumeErrorBudget returns the EstimatedTimeToConsumeErrorBudget field if non-nil, zero value otherwise.

### GetEstimatedTimeToConsumeErrorBudgetOk

`func (o *SloBurnRate) GetEstimatedTimeToConsumeErrorBudgetOk() (*float64, bool)`

GetEstimatedTimeToConsumeErrorBudgetOk returns a tuple with the EstimatedTimeToConsumeErrorBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTimeToConsumeErrorBudget

`func (o *SloBurnRate) SetEstimatedTimeToConsumeErrorBudget(v float64)`

SetEstimatedTimeToConsumeErrorBudget sets EstimatedTimeToConsumeErrorBudget field to given value.

### HasEstimatedTimeToConsumeErrorBudget

`func (o *SloBurnRate) HasEstimatedTimeToConsumeErrorBudget() bool`

HasEstimatedTimeToConsumeErrorBudget returns a boolean if a field has been set.

### GetFastBurnThreshold

`func (o *SloBurnRate) GetFastBurnThreshold() float64`

GetFastBurnThreshold returns the FastBurnThreshold field if non-nil, zero value otherwise.

### GetFastBurnThresholdOk

`func (o *SloBurnRate) GetFastBurnThresholdOk() (*float64, bool)`

GetFastBurnThresholdOk returns a tuple with the FastBurnThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastBurnThreshold

`func (o *SloBurnRate) SetFastBurnThreshold(v float64)`

SetFastBurnThreshold sets FastBurnThreshold field to given value.

### HasFastBurnThreshold

`func (o *SloBurnRate) HasFastBurnThreshold() bool`

HasFastBurnThreshold returns a boolean if a field has been set.

### GetSloValue

`func (o *SloBurnRate) GetSloValue() float64`

GetSloValue returns the SloValue field if non-nil, zero value otherwise.

### GetSloValueOk

`func (o *SloBurnRate) GetSloValueOk() (*float64, bool)`

GetSloValueOk returns a tuple with the SloValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloValue

`func (o *SloBurnRate) SetSloValue(v float64)`

SetSloValue sets SloValue field to given value.

### HasSloValue

`func (o *SloBurnRate) HasSloValue() bool`

HasSloValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


