# SLO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BurnRateMetricKey** | **string** | The key for the SLO&#39;s error budget burn rate func metric. | 
**DenominatorValue** | Pointer to **float64** | The denominator value used to evaluate the SLO when **useRateMetric** is set to &#x60;false&#x60;. | [optional] 
**Description** | Pointer to **string** | A short description of the SLO. | [optional] 
**Enabled** | **bool** | The SLO is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**Error** | **string** | The error of the SLO calculation.   If the value differs from &#x60;NONE&#x60;, there is something wrong with the SLO calculation. | 
**ErrorBudget** | **float64** | The error budget of the calculated SLO.   The error budget is the difference between the calculated and target values. A positive number means all is good; a negative number means trouble.  Has the value of the evaluated error budget or the value of &#x60;-1&#x60;:   * If there is an error with the SLO calculation; in that case check the value of the **error** property.  * If the evaluate parameter has not been set to &#x60;true&#x60;; in that case the **error** property will contain no error. | 
**ErrorBudgetBurnRate** | [**SloBurnRate**](SloBurnRate.md) |  | 
**ErrorBudgetMetricKey** | **string** | The key for the SLO&#39;s error budget func metric. | 
**EvaluatedPercentage** | **float64** | The calculated status value of the SLO. Has the value of the evaluated SLO status or the value of &#x60;-1&#x60;:   * If there is an error with the SLO calculation; in that case check the value of the **error** property.  * If the evaluate parameter has not been set to &#x60;true&#x60;; in that case the **error** property will contain no error. | 
**EvaluationType** | **string** | The evaluation type of the SLO. | 
**Filter** | **string** | The entity filter for the SLO evaluation. The total length of the entitySelector string in SLOs is limited to 1,000 characters. Use the [syntax of entity selector](https://dt-url.net/entityselector). | 
**Id** | **string** | The ID of the SLO | 
**MetricDenominator** | Pointer to **string** | The total count metric (the denominator in rate calculation).   Required when the **useRateMetric** is set to &#x60;false&#x60;. | [optional] 
**MetricExpression** | **string** | The percentage-based metric expression for the calculation of the SLO. | 
**MetricKey** | **string** | The key for the SLO&#39;s status func metric. | 
**MetricNumerator** | Pointer to **string** | The metric for the count of successes (the numerator in rate calculation).   Required when the **useRateMetric** is set to &#x60;false&#x60;. | [optional] 
**MetricRate** | Pointer to **string** | The percentage-based metric for the calculation of the SLO.   Required when the **useRateMetric** is set to &#x60;true&#x60;. | [optional] 
**Name** | **string** | The name of the SLO. | 
**NormalizedErrorBudgetMetricKey** | **string** | The key for the SLO&#39;s normalized error budget func metric. | 
**NumeratorValue** | Pointer to **float64** | The numerator value used to evaluate the SLO when **useRateMetric** is set to &#x60;false&#x60;. | [optional] 
**ProblemFilters** | Pointer to **[]string** | The entity filter for fetching the number of problems related to an SLO. Auto-generated in case no filter has been added to the SLO. | [optional] 
**RelatedOpenProblems** | Pointer to **int32** | Number of open problems related to the SLO.   Has the value of &#x60;-1&#x60; if there&#39;s an error with fetching SLO related problems. | [optional] 
**RelatedTotalProblems** | Pointer to **int32** | Total number of problems related to the SLO.   Has the value of &#x60;-1&#x60; if there&#39;s an error with fetching SLO related problems. | [optional] 
**Status** | **string** | The status of the calculated SLO. | 
**Target** | **float64** | The target value of the SLO. | 
**Timeframe** | **string** | The timeframe for the SLO evaluation. Use the syntax of the global timeframe selector. | 
**UseRateMetric** | Pointer to **bool** | The type of the metric to use for SLO calculation:   * &#x60;true&#x60;: An existing percentage-based metric.  * &#x60;false&#x60;: A ratio of two metrics.   For a list of available metrics, see [Built-in metric page](https://dt-url.net/be03kow) or try the [GET metrics](https://dt-url.net/8e43kxf) API call. | [optional] 
**Warning** | **float64** | The warning value of the SLO.    At warning state the SLO is still fulfilled but is getting close to failure. | 

## Methods

### NewSLO

`func NewSLO(burnRateMetricKey string, enabled bool, error_ string, errorBudget float64, errorBudgetBurnRate SloBurnRate, errorBudgetMetricKey string, evaluatedPercentage float64, evaluationType string, filter string, id string, metricExpression string, metricKey string, name string, normalizedErrorBudgetMetricKey string, status string, target float64, timeframe string, warning float64, ) *SLO`

NewSLO instantiates a new SLO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSLOWithDefaults

`func NewSLOWithDefaults() *SLO`

NewSLOWithDefaults instantiates a new SLO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBurnRateMetricKey

`func (o *SLO) GetBurnRateMetricKey() string`

GetBurnRateMetricKey returns the BurnRateMetricKey field if non-nil, zero value otherwise.

### GetBurnRateMetricKeyOk

`func (o *SLO) GetBurnRateMetricKeyOk() (*string, bool)`

GetBurnRateMetricKeyOk returns a tuple with the BurnRateMetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurnRateMetricKey

`func (o *SLO) SetBurnRateMetricKey(v string)`

SetBurnRateMetricKey sets BurnRateMetricKey field to given value.


### GetDenominatorValue

`func (o *SLO) GetDenominatorValue() float64`

GetDenominatorValue returns the DenominatorValue field if non-nil, zero value otherwise.

### GetDenominatorValueOk

`func (o *SLO) GetDenominatorValueOk() (*float64, bool)`

GetDenominatorValueOk returns a tuple with the DenominatorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenominatorValue

`func (o *SLO) SetDenominatorValue(v float64)`

SetDenominatorValue sets DenominatorValue field to given value.

### HasDenominatorValue

`func (o *SLO) HasDenominatorValue() bool`

HasDenominatorValue returns a boolean if a field has been set.

### GetDescription

`func (o *SLO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SLO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SLO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SLO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SLO) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SLO) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SLO) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetError

`func (o *SLO) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SLO) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SLO) SetError(v string)`

SetError sets Error field to given value.


### GetErrorBudget

`func (o *SLO) GetErrorBudget() float64`

GetErrorBudget returns the ErrorBudget field if non-nil, zero value otherwise.

### GetErrorBudgetOk

`func (o *SLO) GetErrorBudgetOk() (*float64, bool)`

GetErrorBudgetOk returns a tuple with the ErrorBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorBudget

`func (o *SLO) SetErrorBudget(v float64)`

SetErrorBudget sets ErrorBudget field to given value.


### GetErrorBudgetBurnRate

`func (o *SLO) GetErrorBudgetBurnRate() SloBurnRate`

GetErrorBudgetBurnRate returns the ErrorBudgetBurnRate field if non-nil, zero value otherwise.

### GetErrorBudgetBurnRateOk

`func (o *SLO) GetErrorBudgetBurnRateOk() (*SloBurnRate, bool)`

GetErrorBudgetBurnRateOk returns a tuple with the ErrorBudgetBurnRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorBudgetBurnRate

`func (o *SLO) SetErrorBudgetBurnRate(v SloBurnRate)`

SetErrorBudgetBurnRate sets ErrorBudgetBurnRate field to given value.


### GetErrorBudgetMetricKey

`func (o *SLO) GetErrorBudgetMetricKey() string`

GetErrorBudgetMetricKey returns the ErrorBudgetMetricKey field if non-nil, zero value otherwise.

### GetErrorBudgetMetricKeyOk

`func (o *SLO) GetErrorBudgetMetricKeyOk() (*string, bool)`

GetErrorBudgetMetricKeyOk returns a tuple with the ErrorBudgetMetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorBudgetMetricKey

`func (o *SLO) SetErrorBudgetMetricKey(v string)`

SetErrorBudgetMetricKey sets ErrorBudgetMetricKey field to given value.


### GetEvaluatedPercentage

`func (o *SLO) GetEvaluatedPercentage() float64`

GetEvaluatedPercentage returns the EvaluatedPercentage field if non-nil, zero value otherwise.

### GetEvaluatedPercentageOk

`func (o *SLO) GetEvaluatedPercentageOk() (*float64, bool)`

GetEvaluatedPercentageOk returns a tuple with the EvaluatedPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluatedPercentage

`func (o *SLO) SetEvaluatedPercentage(v float64)`

SetEvaluatedPercentage sets EvaluatedPercentage field to given value.


### GetEvaluationType

`func (o *SLO) GetEvaluationType() string`

GetEvaluationType returns the EvaluationType field if non-nil, zero value otherwise.

### GetEvaluationTypeOk

`func (o *SLO) GetEvaluationTypeOk() (*string, bool)`

GetEvaluationTypeOk returns a tuple with the EvaluationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationType

`func (o *SLO) SetEvaluationType(v string)`

SetEvaluationType sets EvaluationType field to given value.


### GetFilter

`func (o *SLO) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *SLO) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *SLO) SetFilter(v string)`

SetFilter sets Filter field to given value.


### GetId

`func (o *SLO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SLO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SLO) SetId(v string)`

SetId sets Id field to given value.


### GetMetricDenominator

`func (o *SLO) GetMetricDenominator() string`

GetMetricDenominator returns the MetricDenominator field if non-nil, zero value otherwise.

### GetMetricDenominatorOk

`func (o *SLO) GetMetricDenominatorOk() (*string, bool)`

GetMetricDenominatorOk returns a tuple with the MetricDenominator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDenominator

`func (o *SLO) SetMetricDenominator(v string)`

SetMetricDenominator sets MetricDenominator field to given value.

### HasMetricDenominator

`func (o *SLO) HasMetricDenominator() bool`

HasMetricDenominator returns a boolean if a field has been set.

### GetMetricExpression

`func (o *SLO) GetMetricExpression() string`

GetMetricExpression returns the MetricExpression field if non-nil, zero value otherwise.

### GetMetricExpressionOk

`func (o *SLO) GetMetricExpressionOk() (*string, bool)`

GetMetricExpressionOk returns a tuple with the MetricExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricExpression

`func (o *SLO) SetMetricExpression(v string)`

SetMetricExpression sets MetricExpression field to given value.


### GetMetricKey

`func (o *SLO) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *SLO) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *SLO) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetMetricNumerator

`func (o *SLO) GetMetricNumerator() string`

GetMetricNumerator returns the MetricNumerator field if non-nil, zero value otherwise.

### GetMetricNumeratorOk

`func (o *SLO) GetMetricNumeratorOk() (*string, bool)`

GetMetricNumeratorOk returns a tuple with the MetricNumerator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricNumerator

`func (o *SLO) SetMetricNumerator(v string)`

SetMetricNumerator sets MetricNumerator field to given value.

### HasMetricNumerator

`func (o *SLO) HasMetricNumerator() bool`

HasMetricNumerator returns a boolean if a field has been set.

### GetMetricRate

`func (o *SLO) GetMetricRate() string`

GetMetricRate returns the MetricRate field if non-nil, zero value otherwise.

### GetMetricRateOk

`func (o *SLO) GetMetricRateOk() (*string, bool)`

GetMetricRateOk returns a tuple with the MetricRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricRate

`func (o *SLO) SetMetricRate(v string)`

SetMetricRate sets MetricRate field to given value.

### HasMetricRate

`func (o *SLO) HasMetricRate() bool`

HasMetricRate returns a boolean if a field has been set.

### GetName

`func (o *SLO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SLO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SLO) SetName(v string)`

SetName sets Name field to given value.


### GetNormalizedErrorBudgetMetricKey

`func (o *SLO) GetNormalizedErrorBudgetMetricKey() string`

GetNormalizedErrorBudgetMetricKey returns the NormalizedErrorBudgetMetricKey field if non-nil, zero value otherwise.

### GetNormalizedErrorBudgetMetricKeyOk

`func (o *SLO) GetNormalizedErrorBudgetMetricKeyOk() (*string, bool)`

GetNormalizedErrorBudgetMetricKeyOk returns a tuple with the NormalizedErrorBudgetMetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedErrorBudgetMetricKey

`func (o *SLO) SetNormalizedErrorBudgetMetricKey(v string)`

SetNormalizedErrorBudgetMetricKey sets NormalizedErrorBudgetMetricKey field to given value.


### GetNumeratorValue

`func (o *SLO) GetNumeratorValue() float64`

GetNumeratorValue returns the NumeratorValue field if non-nil, zero value otherwise.

### GetNumeratorValueOk

`func (o *SLO) GetNumeratorValueOk() (*float64, bool)`

GetNumeratorValueOk returns a tuple with the NumeratorValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumeratorValue

`func (o *SLO) SetNumeratorValue(v float64)`

SetNumeratorValue sets NumeratorValue field to given value.

### HasNumeratorValue

`func (o *SLO) HasNumeratorValue() bool`

HasNumeratorValue returns a boolean if a field has been set.

### GetProblemFilters

`func (o *SLO) GetProblemFilters() []string`

GetProblemFilters returns the ProblemFilters field if non-nil, zero value otherwise.

### GetProblemFiltersOk

`func (o *SLO) GetProblemFiltersOk() (*[]string, bool)`

GetProblemFiltersOk returns a tuple with the ProblemFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemFilters

`func (o *SLO) SetProblemFilters(v []string)`

SetProblemFilters sets ProblemFilters field to given value.

### HasProblemFilters

`func (o *SLO) HasProblemFilters() bool`

HasProblemFilters returns a boolean if a field has been set.

### GetRelatedOpenProblems

`func (o *SLO) GetRelatedOpenProblems() int32`

GetRelatedOpenProblems returns the RelatedOpenProblems field if non-nil, zero value otherwise.

### GetRelatedOpenProblemsOk

`func (o *SLO) GetRelatedOpenProblemsOk() (*int32, bool)`

GetRelatedOpenProblemsOk returns a tuple with the RelatedOpenProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedOpenProblems

`func (o *SLO) SetRelatedOpenProblems(v int32)`

SetRelatedOpenProblems sets RelatedOpenProblems field to given value.

### HasRelatedOpenProblems

`func (o *SLO) HasRelatedOpenProblems() bool`

HasRelatedOpenProblems returns a boolean if a field has been set.

### GetRelatedTotalProblems

`func (o *SLO) GetRelatedTotalProblems() int32`

GetRelatedTotalProblems returns the RelatedTotalProblems field if non-nil, zero value otherwise.

### GetRelatedTotalProblemsOk

`func (o *SLO) GetRelatedTotalProblemsOk() (*int32, bool)`

GetRelatedTotalProblemsOk returns a tuple with the RelatedTotalProblems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedTotalProblems

`func (o *SLO) SetRelatedTotalProblems(v int32)`

SetRelatedTotalProblems sets RelatedTotalProblems field to given value.

### HasRelatedTotalProblems

`func (o *SLO) HasRelatedTotalProblems() bool`

HasRelatedTotalProblems returns a boolean if a field has been set.

### GetStatus

`func (o *SLO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SLO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SLO) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTarget

`func (o *SLO) GetTarget() float64`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SLO) GetTargetOk() (*float64, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SLO) SetTarget(v float64)`

SetTarget sets Target field to given value.


### GetTimeframe

`func (o *SLO) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *SLO) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *SLO) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.


### GetUseRateMetric

`func (o *SLO) GetUseRateMetric() bool`

GetUseRateMetric returns the UseRateMetric field if non-nil, zero value otherwise.

### GetUseRateMetricOk

`func (o *SLO) GetUseRateMetricOk() (*bool, bool)`

GetUseRateMetricOk returns a tuple with the UseRateMetric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRateMetric

`func (o *SLO) SetUseRateMetric(v bool)`

SetUseRateMetric sets UseRateMetric field to given value.

### HasUseRateMetric

`func (o *SLO) HasUseRateMetric() bool`

HasUseRateMetric returns a boolean if a field has been set.

### GetWarning

`func (o *SLO) GetWarning() float64`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *SLO) GetWarningOk() (*float64, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *SLO) SetWarning(v float64)`

SetWarning sets Warning field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


