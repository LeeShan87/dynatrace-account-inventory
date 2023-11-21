# SloConfigItemDtoImpl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the SLO. | [optional] 
**Enabled** | Pointer to **bool** | The SLO is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;).   If not defined, the SLO is disabled by default. | [optional] 
**ErrorBudgetBurnRate** | Pointer to [**SloBurnRateConfig**](SloBurnRateConfig.md) |  | [optional] 
**EvaluationType** | **string** | The evaluation type of the SLO. | 
**Filter** | Pointer to **string** | The entity filter for the SLO evaluation. The total length of the entitySelector string in SLOs is limited to 1,000 characters. Use the [syntax of entity selector](https://dt-url.net/entityselector). | [optional] 
**MetricDenominator** | Pointer to **string** | The total count metric (the denominator in rate calculation).   Required when the **useRateMetric** is set to &#x60;false&#x60;. | [optional] 
**MetricExpression** | Pointer to **string** | The percentage-based metric expression for the calculation of the SLO. | [optional] 
**MetricName** | Pointer to **string** | The name that is used to create SLO func metrics keys. Once created, metric name cannot be changed. | [optional] 
**MetricNumerator** | Pointer to **string** | The metric for the count of successes (the numerator in rate calculation).   Required when the **useRateMetric** is set to &#x60;false&#x60;. | [optional] 
**MetricRate** | Pointer to **string** | The percentage-based metric for the calculation of the SLO.   Required when the **useRateMetric** is set to &#x60;true&#x60;. | [optional] 
**Name** | **string** | The name of the SLO. | 
**Target** | **float64** | The target value of the SLO. | 
**Timeframe** | **string** | The timeframe for the SLO evaluation. Use the syntax of the global timeframe selector. | 
**UseRateMetric** | Pointer to **NullableBool** | The type of the metric to use for SLO calculation:   * &#x60;true&#x60;: An existing percentage-based metric.  * &#x60;false&#x60;: A ratio of two metrics.   For a list of available metrics, see [Built-in metric page](https://dt-url.net/be03kow) or try the [GET metrics](https://dt-url.net/8e43kxf) API call. | [optional] 
**Warning** | **float64** | The warning value of the SLO.    At warning state the SLO is still fulfilled but is getting close to failure. | 

## Methods

### NewSloConfigItemDtoImpl

`func NewSloConfigItemDtoImpl(evaluationType string, name string, target float64, timeframe string, warning float64, ) *SloConfigItemDtoImpl`

NewSloConfigItemDtoImpl instantiates a new SloConfigItemDtoImpl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSloConfigItemDtoImplWithDefaults

`func NewSloConfigItemDtoImplWithDefaults() *SloConfigItemDtoImpl`

NewSloConfigItemDtoImplWithDefaults instantiates a new SloConfigItemDtoImpl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SloConfigItemDtoImpl) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SloConfigItemDtoImpl) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SloConfigItemDtoImpl) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SloConfigItemDtoImpl) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SloConfigItemDtoImpl) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SloConfigItemDtoImpl) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SloConfigItemDtoImpl) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SloConfigItemDtoImpl) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetErrorBudgetBurnRate

`func (o *SloConfigItemDtoImpl) GetErrorBudgetBurnRate() SloBurnRateConfig`

GetErrorBudgetBurnRate returns the ErrorBudgetBurnRate field if non-nil, zero value otherwise.

### GetErrorBudgetBurnRateOk

`func (o *SloConfigItemDtoImpl) GetErrorBudgetBurnRateOk() (*SloBurnRateConfig, bool)`

GetErrorBudgetBurnRateOk returns a tuple with the ErrorBudgetBurnRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorBudgetBurnRate

`func (o *SloConfigItemDtoImpl) SetErrorBudgetBurnRate(v SloBurnRateConfig)`

SetErrorBudgetBurnRate sets ErrorBudgetBurnRate field to given value.

### HasErrorBudgetBurnRate

`func (o *SloConfigItemDtoImpl) HasErrorBudgetBurnRate() bool`

HasErrorBudgetBurnRate returns a boolean if a field has been set.

### GetEvaluationType

`func (o *SloConfigItemDtoImpl) GetEvaluationType() string`

GetEvaluationType returns the EvaluationType field if non-nil, zero value otherwise.

### GetEvaluationTypeOk

`func (o *SloConfigItemDtoImpl) GetEvaluationTypeOk() (*string, bool)`

GetEvaluationTypeOk returns a tuple with the EvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationType

`func (o *SloConfigItemDtoImpl) SetEvaluationType(v string)`

SetEvaluationType sets EvaluationType field to given value.


### GetFilter

`func (o *SloConfigItemDtoImpl) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SloConfigItemDtoImpl) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SloConfigItemDtoImpl) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *SloConfigItemDtoImpl) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetMetricDenominator

`func (o *SloConfigItemDtoImpl) GetMetricDenominator() string`

GetMetricDenominator returns the MetricDenominator field if non-nil, zero value otherwise.

### GetMetricDenominatorOk

`func (o *SloConfigItemDtoImpl) GetMetricDenominatorOk() (*string, bool)`

GetMetricDenominatorOk returns a tuple with the MetricDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDenominator

`func (o *SloConfigItemDtoImpl) SetMetricDenominator(v string)`

SetMetricDenominator sets MetricDenominator field to given value.

### HasMetricDenominator

`func (o *SloConfigItemDtoImpl) HasMetricDenominator() bool`

HasMetricDenominator returns a boolean if a field has been set.

### GetMetricExpression

`func (o *SloConfigItemDtoImpl) GetMetricExpression() string`

GetMetricExpression returns the MetricExpression field if non-nil, zero value otherwise.

### GetMetricExpressionOk

`func (o *SloConfigItemDtoImpl) GetMetricExpressionOk() (*string, bool)`

GetMetricExpressionOk returns a tuple with the MetricExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricExpression

`func (o *SloConfigItemDtoImpl) SetMetricExpression(v string)`

SetMetricExpression sets MetricExpression field to given value.

### HasMetricExpression

`func (o *SloConfigItemDtoImpl) HasMetricExpression() bool`

HasMetricExpression returns a boolean if a field has been set.

### GetMetricName

`func (o *SloConfigItemDtoImpl) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *SloConfigItemDtoImpl) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *SloConfigItemDtoImpl) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.

### HasMetricName

`func (o *SloConfigItemDtoImpl) HasMetricName() bool`

HasMetricName returns a boolean if a field has been set.

### GetMetricNumerator

`func (o *SloConfigItemDtoImpl) GetMetricNumerator() string`

GetMetricNumerator returns the MetricNumerator field if non-nil, zero value otherwise.

### GetMetricNumeratorOk

`func (o *SloConfigItemDtoImpl) GetMetricNumeratorOk() (*string, bool)`

GetMetricNumeratorOk returns a tuple with the MetricNumerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricNumerator

`func (o *SloConfigItemDtoImpl) SetMetricNumerator(v string)`

SetMetricNumerator sets MetricNumerator field to given value.

### HasMetricNumerator

`func (o *SloConfigItemDtoImpl) HasMetricNumerator() bool`

HasMetricNumerator returns a boolean if a field has been set.

### GetMetricRate

`func (o *SloConfigItemDtoImpl) GetMetricRate() string`

GetMetricRate returns the MetricRate field if non-nil, zero value otherwise.

### GetMetricRateOk

`func (o *SloConfigItemDtoImpl) GetMetricRateOk() (*string, bool)`

GetMetricRateOk returns a tuple with the MetricRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricRate

`func (o *SloConfigItemDtoImpl) SetMetricRate(v string)`

SetMetricRate sets MetricRate field to given value.

### HasMetricRate

`func (o *SloConfigItemDtoImpl) HasMetricRate() bool`

HasMetricRate returns a boolean if a field has been set.

### GetName

`func (o *SloConfigItemDtoImpl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SloConfigItemDtoImpl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SloConfigItemDtoImpl) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *SloConfigItemDtoImpl) GetTarget() float64`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SloConfigItemDtoImpl) GetTargetOk() (*float64, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SloConfigItemDtoImpl) SetTarget(v float64)`

SetTarget sets Target field to given value.


### GetTimeframe

`func (o *SloConfigItemDtoImpl) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *SloConfigItemDtoImpl) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *SloConfigItemDtoImpl) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.


### GetUseRateMetric

`func (o *SloConfigItemDtoImpl) GetUseRateMetric() bool`

GetUseRateMetric returns the UseRateMetric field if non-nil, zero value otherwise.

### GetUseRateMetricOk

`func (o *SloConfigItemDtoImpl) GetUseRateMetricOk() (*bool, bool)`

GetUseRateMetricOk returns a tuple with the UseRateMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRateMetric

`func (o *SloConfigItemDtoImpl) SetUseRateMetric(v bool)`

SetUseRateMetric sets UseRateMetric field to given value.

### HasUseRateMetric

`func (o *SloConfigItemDtoImpl) HasUseRateMetric() bool`

HasUseRateMetric returns a boolean if a field has been set.

### SetUseRateMetricNil

`func (o *SloConfigItemDtoImpl) SetUseRateMetricNil(b bool)`

 SetUseRateMetricNil sets the value for UseRateMetric to be an explicit nil

### UnsetUseRateMetric
`func (o *SloConfigItemDtoImpl) UnsetUseRateMetric()`

UnsetUseRateMetric ensures that no value is present for UseRateMetric, not even an explicit nil
### GetWarning

`func (o *SloConfigItemDtoImpl) GetWarning() float64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *SloConfigItemDtoImpl) GetWarningOk() (*float64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *SloConfigItemDtoImpl) SetWarning(v float64)`

SetWarning sets Warning field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


