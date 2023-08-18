# MetricEventStaticThresholdMonitoringStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertCondition** | **string** | The condition for the **threshold** value check: above or below. | 
**AlertingOnMissingData** | Pointer to **bool** | If true, also one-minute samples without data are counted as violating samples. | [optional] 
**DealertingSamples** | **int32** | The number of one-minute samples within the evaluation window that must go back to normal to close the event. | 
**Samples** | **int32** | The number of one-minute samples that form the sliding evaluation window. | 
**Threshold** | **float64** | The value of the static threshold based on the specified unit. | 
**Unit** | Pointer to **string** | The unit of the threshold, matching the metric definition. | [optional] 
**ViolatingSamples** | **int32** | The number of one-minute samples within the evaluation window that must violate the threshold to trigger an event. | 

## Methods

### NewMetricEventStaticThresholdMonitoringStrategy

`func NewMetricEventStaticThresholdMonitoringStrategy(alertCondition string, dealertingSamples int32, samples int32, threshold float64, violatingSamples int32, ) *MetricEventStaticThresholdMonitoringStrategy`

NewMetricEventStaticThresholdMonitoringStrategy instantiates a new MetricEventStaticThresholdMonitoringStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricEventStaticThresholdMonitoringStrategyWithDefaults

`func NewMetricEventStaticThresholdMonitoringStrategyWithDefaults() *MetricEventStaticThresholdMonitoringStrategy`

NewMetricEventStaticThresholdMonitoringStrategyWithDefaults instantiates a new MetricEventStaticThresholdMonitoringStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertCondition

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetAlertCondition() string`

GetAlertCondition returns the AlertCondition field if non-nil, zero value otherwise.

### GetAlertConditionOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetAlertConditionOk() (*string, bool)`

GetAlertConditionOk returns a tuple with the AlertCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCondition

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetAlertCondition(v string)`

SetAlertCondition sets AlertCondition field to given value.


### GetAlertingOnMissingData

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetAlertingOnMissingData() bool`

GetAlertingOnMissingData returns the AlertingOnMissingData field if non-nil, zero value otherwise.

### GetAlertingOnMissingDataOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetAlertingOnMissingDataOk() (*bool, bool)`

GetAlertingOnMissingDataOk returns a tuple with the AlertingOnMissingData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertingOnMissingData

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetAlertingOnMissingData(v bool)`

SetAlertingOnMissingData sets AlertingOnMissingData field to given value.

### HasAlertingOnMissingData

`func (o *MetricEventStaticThresholdMonitoringStrategy) HasAlertingOnMissingData() bool`

HasAlertingOnMissingData returns a boolean if a field has been set.

### GetDealertingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetDealertingSamples() int32`

GetDealertingSamples returns the DealertingSamples field if non-nil, zero value otherwise.

### GetDealertingSamplesOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetDealertingSamplesOk() (*int32, bool)`

GetDealertingSamplesOk returns a tuple with the DealertingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDealertingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetDealertingSamples(v int32)`

SetDealertingSamples sets DealertingSamples field to given value.


### GetSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetSamples() int32`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetSamplesOk() (*int32, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetSamples(v int32)`

SetSamples sets Samples field to given value.


### GetThreshold

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.


### GetUnit

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricEventStaticThresholdMonitoringStrategy) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetViolatingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetViolatingSamples() int32`

GetViolatingSamples returns the ViolatingSamples field if non-nil, zero value otherwise.

### GetViolatingSamplesOk

`func (o *MetricEventStaticThresholdMonitoringStrategy) GetViolatingSamplesOk() (*int32, bool)`

GetViolatingSamplesOk returns a tuple with the ViolatingSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolatingSamples

`func (o *MetricEventStaticThresholdMonitoringStrategy) SetViolatingSamples(v int32)`

SetViolatingSamples sets ViolatingSamples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


