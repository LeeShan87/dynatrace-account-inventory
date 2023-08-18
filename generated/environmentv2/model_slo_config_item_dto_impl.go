/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
)

// checks if the SloConfigItemDtoImpl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SloConfigItemDtoImpl{}

// SloConfigItemDtoImpl struct for SloConfigItemDtoImpl
type SloConfigItemDtoImpl struct {
	// The description of the SLO.
	Description *string `json:"description,omitempty"`
	// The SLO is enabled (`true`) or disabled (`false`).   If not defined, the SLO is disabled by default.
	Enabled *bool `json:"enabled,omitempty"`
	ErrorBudgetBurnRate *SloBurnRateConfig `json:"errorBudgetBurnRate,omitempty"`
	// The evaluation type of the SLO.
	EvaluationType string `json:"evaluationType"`
	// The entity filter for the SLO evaluation. The total length of the entitySelector string in SLOs is limited to 1,000 characters. Use the [syntax of entity selector](https://dt-url.net/entityselector).
	Filter *string `json:"filter,omitempty"`
	// The total count metric (the denominator in rate calculation).   Required when the **useRateMetric** is set to `false`.
	// Deprecated
	MetricDenominator *string `json:"metricDenominator,omitempty"`
	// The percentage-based metric expression for the calculation of the SLO.
	MetricExpression *string `json:"metricExpression,omitempty"`
	// The name that is used to create SLO func metrics keys. Once created, metric name cannot be changed.
	MetricName *string `json:"metricName,omitempty"`
	// The metric for the count of successes (the numerator in rate calculation).   Required when the **useRateMetric** is set to `false`.
	// Deprecated
	MetricNumerator *string `json:"metricNumerator,omitempty"`
	// The percentage-based metric for the calculation of the SLO.   Required when the **useRateMetric** is set to `true`.
	// Deprecated
	MetricRate *string `json:"metricRate,omitempty"`
	// The name of the SLO.
	Name string `json:"name"`
	// The target value of the SLO.
	Target float64 `json:"target"`
	// The timeframe for the SLO evaluation. Use the syntax of the global timeframe selector.
	Timeframe string `json:"timeframe"`
	// The type of the metric to use for SLO calculation:   * `true`: An existing percentage-based metric.  * `false`: A ratio of two metrics.   For a list of available metrics, see [Built-in metric page](https://dt-url.net/be03kow) or try the [GET metrics](https://dt-url.net/8e43kxf) API call.
	// Deprecated
	UseRateMetric NullableBool `json:"useRateMetric,omitempty"`
	// The warning value of the SLO.    At warning state the SLO is still fulfilled but is getting close to failure.
	Warning float64 `json:"warning"`
}

// NewSloConfigItemDtoImpl instantiates a new SloConfigItemDtoImpl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSloConfigItemDtoImpl(evaluationType string, name string, target float64, timeframe string, warning float64) *SloConfigItemDtoImpl {
	this := SloConfigItemDtoImpl{}
	this.EvaluationType = evaluationType
	this.Name = name
	this.Target = target
	this.Timeframe = timeframe
	this.Warning = warning
	return &this
}

// NewSloConfigItemDtoImplWithDefaults instantiates a new SloConfigItemDtoImpl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSloConfigItemDtoImplWithDefaults() *SloConfigItemDtoImpl {
	this := SloConfigItemDtoImpl{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SloConfigItemDtoImpl) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SloConfigItemDtoImpl) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SloConfigItemDtoImpl) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SloConfigItemDtoImpl) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SloConfigItemDtoImpl) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SloConfigItemDtoImpl) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetErrorBudgetBurnRate returns the ErrorBudgetBurnRate field value if set, zero value otherwise.
func (o *SloConfigItemDtoImpl) GetErrorBudgetBurnRate() SloBurnRateConfig {
	if o == nil || IsNil(o.ErrorBudgetBurnRate) {
		var ret SloBurnRateConfig
		return ret
	}
	return *o.ErrorBudgetBurnRate
}

// GetErrorBudgetBurnRateOk returns a tuple with the ErrorBudgetBurnRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetErrorBudgetBurnRateOk() (*SloBurnRateConfig, bool) {
	if o == nil || IsNil(o.ErrorBudgetBurnRate) {
		return nil, false
	}
	return o.ErrorBudgetBurnRate, true
}

// HasErrorBudgetBurnRate returns a boolean if a field has been set.
func (o *SloConfigItemDtoImpl) HasErrorBudgetBurnRate() bool {
	if o != nil && !IsNil(o.ErrorBudgetBurnRate) {
		return true
	}

	return false
}

// SetErrorBudgetBurnRate gets a reference to the given SloBurnRateConfig and assigns it to the ErrorBudgetBurnRate field.
func (o *SloConfigItemDtoImpl) SetErrorBudgetBurnRate(v SloBurnRateConfig) {
	o.ErrorBudgetBurnRate = &v
}

// GetEvaluationType returns the EvaluationType field value
func (o *SloConfigItemDtoImpl) GetEvaluationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvaluationType
}

// GetEvaluationTypeOk returns a tuple with the EvaluationType field value
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetEvaluationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvaluationType, true
}

// SetEvaluationType sets field value
func (o *SloConfigItemDtoImpl) SetEvaluationType(v string) {
	o.EvaluationType = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *SloConfigItemDtoImpl) GetFilter() string {
	if o == nil || IsNil(o.Filter) {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetFilterOk() (*string, bool) {
	if o == nil || IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *SloConfigItemDtoImpl) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *SloConfigItemDtoImpl) SetFilter(v string) {
	o.Filter = &v
}

// GetMetricDenominator returns the MetricDenominator field value if set, zero value otherwise.
// Deprecated
func (o *SloConfigItemDtoImpl) GetMetricDenominator() string {
	if o == nil || IsNil(o.MetricDenominator) {
		var ret string
		return ret
	}
	return *o.MetricDenominator
}

// GetMetricDenominatorOk returns a tuple with the MetricDenominator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SloConfigItemDtoImpl) GetMetricDenominatorOk() (*string, bool) {
	if o == nil || IsNil(o.MetricDenominator) {
		return nil, false
	}
	return o.MetricDenominator, true
}

// HasMetricDenominator returns a boolean if a field has been set.
func (o *SloConfigItemDtoImpl) HasMetricDenominator() bool {
	if o != nil && !IsNil(o.MetricDenominator) {
		return true
	}

	return false
}

// SetMetricDenominator gets a reference to the given string and assigns it to the MetricDenominator field.
// Deprecated
func (o *SloConfigItemDtoImpl) SetMetricDenominator(v string) {
	o.MetricDenominator = &v
}

// GetMetricExpression returns the MetricExpression field value if set, zero value otherwise.
func (o *SloConfigItemDtoImpl) GetMetricExpression() string {
	if o == nil || IsNil(o.MetricExpression) {
		var ret string
		return ret
	}
	return *o.MetricExpression
}

// GetMetricExpressionOk returns a tuple with the MetricExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetMetricExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.MetricExpression) {
		return nil, false
	}
	return o.MetricExpression, true
}

// HasMetricExpression returns a boolean if a field has been set.
func (o *SloConfigItemDtoImpl) HasMetricExpression() bool {
	if o != nil && !IsNil(o.MetricExpression) {
		return true
	}

	return false
}

// SetMetricExpression gets a reference to the given string and assigns it to the MetricExpression field.
func (o *SloConfigItemDtoImpl) SetMetricExpression(v string) {
	o.MetricExpression = &v
}

// GetMetricName returns the MetricName field value if set, zero value otherwise.
func (o *SloConfigItemDtoImpl) GetMetricName() string {
	if o == nil || IsNil(o.MetricName) {
		var ret string
		return ret
	}
	return *o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetMetricNameOk() (*string, bool) {
	if o == nil || IsNil(o.MetricName) {
		return nil, false
	}
	return o.MetricName, true
}

// HasMetricName returns a boolean if a field has been set.
func (o *SloConfigItemDtoImpl) HasMetricName() bool {
	if o != nil && !IsNil(o.MetricName) {
		return true
	}

	return false
}

// SetMetricName gets a reference to the given string and assigns it to the MetricName field.
func (o *SloConfigItemDtoImpl) SetMetricName(v string) {
	o.MetricName = &v
}

// GetMetricNumerator returns the MetricNumerator field value if set, zero value otherwise.
// Deprecated
func (o *SloConfigItemDtoImpl) GetMetricNumerator() string {
	if o == nil || IsNil(o.MetricNumerator) {
		var ret string
		return ret
	}
	return *o.MetricNumerator
}

// GetMetricNumeratorOk returns a tuple with the MetricNumerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SloConfigItemDtoImpl) GetMetricNumeratorOk() (*string, bool) {
	if o == nil || IsNil(o.MetricNumerator) {
		return nil, false
	}
	return o.MetricNumerator, true
}

// HasMetricNumerator returns a boolean if a field has been set.
func (o *SloConfigItemDtoImpl) HasMetricNumerator() bool {
	if o != nil && !IsNil(o.MetricNumerator) {
		return true
	}

	return false
}

// SetMetricNumerator gets a reference to the given string and assigns it to the MetricNumerator field.
// Deprecated
func (o *SloConfigItemDtoImpl) SetMetricNumerator(v string) {
	o.MetricNumerator = &v
}

// GetMetricRate returns the MetricRate field value if set, zero value otherwise.
// Deprecated
func (o *SloConfigItemDtoImpl) GetMetricRate() string {
	if o == nil || IsNil(o.MetricRate) {
		var ret string
		return ret
	}
	return *o.MetricRate
}

// GetMetricRateOk returns a tuple with the MetricRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *SloConfigItemDtoImpl) GetMetricRateOk() (*string, bool) {
	if o == nil || IsNil(o.MetricRate) {
		return nil, false
	}
	return o.MetricRate, true
}

// HasMetricRate returns a boolean if a field has been set.
func (o *SloConfigItemDtoImpl) HasMetricRate() bool {
	if o != nil && !IsNil(o.MetricRate) {
		return true
	}

	return false
}

// SetMetricRate gets a reference to the given string and assigns it to the MetricRate field.
// Deprecated
func (o *SloConfigItemDtoImpl) SetMetricRate(v string) {
	o.MetricRate = &v
}

// GetName returns the Name field value
func (o *SloConfigItemDtoImpl) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SloConfigItemDtoImpl) SetName(v string) {
	o.Name = v
}

// GetTarget returns the Target field value
func (o *SloConfigItemDtoImpl) GetTarget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetTargetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *SloConfigItemDtoImpl) SetTarget(v float64) {
	o.Target = v
}

// GetTimeframe returns the Timeframe field value
func (o *SloConfigItemDtoImpl) GetTimeframe() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetTimeframeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframe, true
}

// SetTimeframe sets field value
func (o *SloConfigItemDtoImpl) SetTimeframe(v string) {
	o.Timeframe = v
}

// GetUseRateMetric returns the UseRateMetric field value if set, zero value otherwise (both if not set or set to explicit null).
// Deprecated
func (o *SloConfigItemDtoImpl) GetUseRateMetric() bool {
	if o == nil || IsNil(o.UseRateMetric.Get()) {
		var ret bool
		return ret
	}
	return *o.UseRateMetric.Get()
}

// GetUseRateMetricOk returns a tuple with the UseRateMetric field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *SloConfigItemDtoImpl) GetUseRateMetricOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseRateMetric.Get(), o.UseRateMetric.IsSet()
}

// HasUseRateMetric returns a boolean if a field has been set.
func (o *SloConfigItemDtoImpl) HasUseRateMetric() bool {
	if o != nil && o.UseRateMetric.IsSet() {
		return true
	}

	return false
}

// SetUseRateMetric gets a reference to the given NullableBool and assigns it to the UseRateMetric field.
// Deprecated
func (o *SloConfigItemDtoImpl) SetUseRateMetric(v bool) {
	o.UseRateMetric.Set(&v)
}
// SetUseRateMetricNil sets the value for UseRateMetric to be an explicit nil
func (o *SloConfigItemDtoImpl) SetUseRateMetricNil() {
	o.UseRateMetric.Set(nil)
}

// UnsetUseRateMetric ensures that no value is present for UseRateMetric, not even an explicit nil
func (o *SloConfigItemDtoImpl) UnsetUseRateMetric() {
	o.UseRateMetric.Unset()
}

// GetWarning returns the Warning field value
func (o *SloConfigItemDtoImpl) GetWarning() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Warning
}

// GetWarningOk returns a tuple with the Warning field value
// and a boolean to check if the value has been set.
func (o *SloConfigItemDtoImpl) GetWarningOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Warning, true
}

// SetWarning sets field value
func (o *SloConfigItemDtoImpl) SetWarning(v float64) {
	o.Warning = v
}

func (o SloConfigItemDtoImpl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SloConfigItemDtoImpl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ErrorBudgetBurnRate) {
		toSerialize["errorBudgetBurnRate"] = o.ErrorBudgetBurnRate
	}
	toSerialize["evaluationType"] = o.EvaluationType
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.MetricDenominator) {
		toSerialize["metricDenominator"] = o.MetricDenominator
	}
	if !IsNil(o.MetricExpression) {
		toSerialize["metricExpression"] = o.MetricExpression
	}
	if !IsNil(o.MetricName) {
		toSerialize["metricName"] = o.MetricName
	}
	if !IsNil(o.MetricNumerator) {
		toSerialize["metricNumerator"] = o.MetricNumerator
	}
	if !IsNil(o.MetricRate) {
		toSerialize["metricRate"] = o.MetricRate
	}
	toSerialize["name"] = o.Name
	toSerialize["target"] = o.Target
	toSerialize["timeframe"] = o.Timeframe
	if o.UseRateMetric.IsSet() {
		toSerialize["useRateMetric"] = o.UseRateMetric.Get()
	}
	toSerialize["warning"] = o.Warning
	return toSerialize, nil
}

type NullableSloConfigItemDtoImpl struct {
	value *SloConfigItemDtoImpl
	isSet bool
}

func (v NullableSloConfigItemDtoImpl) Get() *SloConfigItemDtoImpl {
	return v.value
}

func (v *NullableSloConfigItemDtoImpl) Set(val *SloConfigItemDtoImpl) {
	v.value = val
	v.isSet = true
}

func (v NullableSloConfigItemDtoImpl) IsSet() bool {
	return v.isSet
}

func (v *NullableSloConfigItemDtoImpl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSloConfigItemDtoImpl(val *SloConfigItemDtoImpl) *NullableSloConfigItemDtoImpl {
	return &NullableSloConfigItemDtoImpl{value: val, isSet: true}
}

func (v NullableSloConfigItemDtoImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSloConfigItemDtoImpl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


