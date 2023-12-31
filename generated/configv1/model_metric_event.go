/*
Dynatrace Configuration API

Documentation of the Dynatrace Configuration API. To read about use-cases and examples, see [Dynatrace Documentation](https://dt-url.net/4u43kxw).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configv1

import (
	"encoding/json"
)

// checks if the MetricEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricEvent{}

// MetricEvent The configuration of the metric event.
type MetricEvent struct {
	// How the metric data points are aggregated for the evaluation.    The timeseries must support this aggregation.
	AggregationType *string `json:"aggregationType,omitempty"`
	// Defines the scope of the metric event. Only one filter is allowed per filter type, except for tags, where up to 3 are allowed. The filters are combined by conjunction.
	AlertingScope []MetricEventAlertingScope `json:"alertingScope,omitempty"`
	// The description of the metric event.
	Description string `json:"description"`
	// The reason of automatic disabling.   The `NONE` means config was not disabled automatically.
	DisabledReason *string `json:"disabledReason,omitempty"`
	// The metric event is enabled (`true`) or disabled (`false`).
	Enabled bool `json:"enabled"`
	// The ID of the metric event.
	Id *string `json:"id,omitempty"`
	Metadata *ConfigurationMetadata `json:"metadata,omitempty"`
	// Defines the dimensions of the metric to alert on. The filters are combined by conjunction.
	MetricDimensions []MetricEventDimensions `json:"metricDimensions,omitempty"`
	// The ID of the metric evaluated by the metric event.
	MetricId *string `json:"metricId,omitempty"`
	// The metric selector that should be executed.
	MetricSelector *string `json:"metricSelector,omitempty"`
	MonitoringStrategy MetricEventMonitoringStrategy `json:"monitoringStrategy"`
	// The name of the metric event displayed in the UI.
	Name string `json:"name"`
	// Defines which dimension key should be used for the **alertingScope**.
	PrimaryDimensionKey *string `json:"primaryDimensionKey,omitempty"`
	// Defines the query offset to adapt the evaluation timeframe for known metric latency.
	QueryOffset *int64 `json:"queryOffset,omitempty"`
	// The type of the event to trigger on the threshold violation.   The `CUSTOM_ALERT` type is not correlated with other alerts. The `INFO` type does not open a problem.
	Severity *string `json:"severity,omitempty"`
	// The reason of a warning set on the config.   The `NONE` means config has no warnings.
	WarningReason *string `json:"warningReason,omitempty"`
}

// NewMetricEvent instantiates a new MetricEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricEvent(description string, enabled bool, monitoringStrategy MetricEventMonitoringStrategy, name string) *MetricEvent {
	this := MetricEvent{}
	this.Description = description
	this.Enabled = enabled
	this.MonitoringStrategy = monitoringStrategy
	this.Name = name
	return &this
}

// NewMetricEventWithDefaults instantiates a new MetricEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricEventWithDefaults() *MetricEvent {
	this := MetricEvent{}
	return &this
}

// GetAggregationType returns the AggregationType field value if set, zero value otherwise.
func (o *MetricEvent) GetAggregationType() string {
	if o == nil || IsNil(o.AggregationType) {
		var ret string
		return ret
	}
	return *o.AggregationType
}

// GetAggregationTypeOk returns a tuple with the AggregationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetAggregationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AggregationType) {
		return nil, false
	}
	return o.AggregationType, true
}

// HasAggregationType returns a boolean if a field has been set.
func (o *MetricEvent) HasAggregationType() bool {
	if o != nil && !IsNil(o.AggregationType) {
		return true
	}

	return false
}

// SetAggregationType gets a reference to the given string and assigns it to the AggregationType field.
func (o *MetricEvent) SetAggregationType(v string) {
	o.AggregationType = &v
}

// GetAlertingScope returns the AlertingScope field value if set, zero value otherwise.
func (o *MetricEvent) GetAlertingScope() []MetricEventAlertingScope {
	if o == nil || IsNil(o.AlertingScope) {
		var ret []MetricEventAlertingScope
		return ret
	}
	return o.AlertingScope
}

// GetAlertingScopeOk returns a tuple with the AlertingScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetAlertingScopeOk() ([]MetricEventAlertingScope, bool) {
	if o == nil || IsNil(o.AlertingScope) {
		return nil, false
	}
	return o.AlertingScope, true
}

// HasAlertingScope returns a boolean if a field has been set.
func (o *MetricEvent) HasAlertingScope() bool {
	if o != nil && !IsNil(o.AlertingScope) {
		return true
	}

	return false
}

// SetAlertingScope gets a reference to the given []MetricEventAlertingScope and assigns it to the AlertingScope field.
func (o *MetricEvent) SetAlertingScope(v []MetricEventAlertingScope) {
	o.AlertingScope = v
}

// GetDescription returns the Description field value
func (o *MetricEvent) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *MetricEvent) SetDescription(v string) {
	o.Description = v
}

// GetDisabledReason returns the DisabledReason field value if set, zero value otherwise.
func (o *MetricEvent) GetDisabledReason() string {
	if o == nil || IsNil(o.DisabledReason) {
		var ret string
		return ret
	}
	return *o.DisabledReason
}

// GetDisabledReasonOk returns a tuple with the DisabledReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetDisabledReasonOk() (*string, bool) {
	if o == nil || IsNil(o.DisabledReason) {
		return nil, false
	}
	return o.DisabledReason, true
}

// HasDisabledReason returns a boolean if a field has been set.
func (o *MetricEvent) HasDisabledReason() bool {
	if o != nil && !IsNil(o.DisabledReason) {
		return true
	}

	return false
}

// SetDisabledReason gets a reference to the given string and assigns it to the DisabledReason field.
func (o *MetricEvent) SetDisabledReason(v string) {
	o.DisabledReason = &v
}

// GetEnabled returns the Enabled field value
func (o *MetricEvent) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *MetricEvent) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MetricEvent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MetricEvent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MetricEvent) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *MetricEvent) GetMetadata() ConfigurationMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ConfigurationMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetMetadataOk() (*ConfigurationMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *MetricEvent) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ConfigurationMetadata and assigns it to the Metadata field.
func (o *MetricEvent) SetMetadata(v ConfigurationMetadata) {
	o.Metadata = &v
}

// GetMetricDimensions returns the MetricDimensions field value if set, zero value otherwise.
func (o *MetricEvent) GetMetricDimensions() []MetricEventDimensions {
	if o == nil || IsNil(o.MetricDimensions) {
		var ret []MetricEventDimensions
		return ret
	}
	return o.MetricDimensions
}

// GetMetricDimensionsOk returns a tuple with the MetricDimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetMetricDimensionsOk() ([]MetricEventDimensions, bool) {
	if o == nil || IsNil(o.MetricDimensions) {
		return nil, false
	}
	return o.MetricDimensions, true
}

// HasMetricDimensions returns a boolean if a field has been set.
func (o *MetricEvent) HasMetricDimensions() bool {
	if o != nil && !IsNil(o.MetricDimensions) {
		return true
	}

	return false
}

// SetMetricDimensions gets a reference to the given []MetricEventDimensions and assigns it to the MetricDimensions field.
func (o *MetricEvent) SetMetricDimensions(v []MetricEventDimensions) {
	o.MetricDimensions = v
}

// GetMetricId returns the MetricId field value if set, zero value otherwise.
func (o *MetricEvent) GetMetricId() string {
	if o == nil || IsNil(o.MetricId) {
		var ret string
		return ret
	}
	return *o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetMetricIdOk() (*string, bool) {
	if o == nil || IsNil(o.MetricId) {
		return nil, false
	}
	return o.MetricId, true
}

// HasMetricId returns a boolean if a field has been set.
func (o *MetricEvent) HasMetricId() bool {
	if o != nil && !IsNil(o.MetricId) {
		return true
	}

	return false
}

// SetMetricId gets a reference to the given string and assigns it to the MetricId field.
func (o *MetricEvent) SetMetricId(v string) {
	o.MetricId = &v
}

// GetMetricSelector returns the MetricSelector field value if set, zero value otherwise.
func (o *MetricEvent) GetMetricSelector() string {
	if o == nil || IsNil(o.MetricSelector) {
		var ret string
		return ret
	}
	return *o.MetricSelector
}

// GetMetricSelectorOk returns a tuple with the MetricSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetMetricSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.MetricSelector) {
		return nil, false
	}
	return o.MetricSelector, true
}

// HasMetricSelector returns a boolean if a field has been set.
func (o *MetricEvent) HasMetricSelector() bool {
	if o != nil && !IsNil(o.MetricSelector) {
		return true
	}

	return false
}

// SetMetricSelector gets a reference to the given string and assigns it to the MetricSelector field.
func (o *MetricEvent) SetMetricSelector(v string) {
	o.MetricSelector = &v
}

// GetMonitoringStrategy returns the MonitoringStrategy field value
func (o *MetricEvent) GetMonitoringStrategy() MetricEventMonitoringStrategy {
	if o == nil {
		var ret MetricEventMonitoringStrategy
		return ret
	}

	return o.MonitoringStrategy
}

// GetMonitoringStrategyOk returns a tuple with the MonitoringStrategy field value
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetMonitoringStrategyOk() (*MetricEventMonitoringStrategy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MonitoringStrategy, true
}

// SetMonitoringStrategy sets field value
func (o *MetricEvent) SetMonitoringStrategy(v MetricEventMonitoringStrategy) {
	o.MonitoringStrategy = v
}

// GetName returns the Name field value
func (o *MetricEvent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetricEvent) SetName(v string) {
	o.Name = v
}

// GetPrimaryDimensionKey returns the PrimaryDimensionKey field value if set, zero value otherwise.
func (o *MetricEvent) GetPrimaryDimensionKey() string {
	if o == nil || IsNil(o.PrimaryDimensionKey) {
		var ret string
		return ret
	}
	return *o.PrimaryDimensionKey
}

// GetPrimaryDimensionKeyOk returns a tuple with the PrimaryDimensionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetPrimaryDimensionKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryDimensionKey) {
		return nil, false
	}
	return o.PrimaryDimensionKey, true
}

// HasPrimaryDimensionKey returns a boolean if a field has been set.
func (o *MetricEvent) HasPrimaryDimensionKey() bool {
	if o != nil && !IsNil(o.PrimaryDimensionKey) {
		return true
	}

	return false
}

// SetPrimaryDimensionKey gets a reference to the given string and assigns it to the PrimaryDimensionKey field.
func (o *MetricEvent) SetPrimaryDimensionKey(v string) {
	o.PrimaryDimensionKey = &v
}

// GetQueryOffset returns the QueryOffset field value if set, zero value otherwise.
func (o *MetricEvent) GetQueryOffset() int64 {
	if o == nil || IsNil(o.QueryOffset) {
		var ret int64
		return ret
	}
	return *o.QueryOffset
}

// GetQueryOffsetOk returns a tuple with the QueryOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetQueryOffsetOk() (*int64, bool) {
	if o == nil || IsNil(o.QueryOffset) {
		return nil, false
	}
	return o.QueryOffset, true
}

// HasQueryOffset returns a boolean if a field has been set.
func (o *MetricEvent) HasQueryOffset() bool {
	if o != nil && !IsNil(o.QueryOffset) {
		return true
	}

	return false
}

// SetQueryOffset gets a reference to the given int64 and assigns it to the QueryOffset field.
func (o *MetricEvent) SetQueryOffset(v int64) {
	o.QueryOffset = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *MetricEvent) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *MetricEvent) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *MetricEvent) SetSeverity(v string) {
	o.Severity = &v
}

// GetWarningReason returns the WarningReason field value if set, zero value otherwise.
func (o *MetricEvent) GetWarningReason() string {
	if o == nil || IsNil(o.WarningReason) {
		var ret string
		return ret
	}
	return *o.WarningReason
}

// GetWarningReasonOk returns a tuple with the WarningReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricEvent) GetWarningReasonOk() (*string, bool) {
	if o == nil || IsNil(o.WarningReason) {
		return nil, false
	}
	return o.WarningReason, true
}

// HasWarningReason returns a boolean if a field has been set.
func (o *MetricEvent) HasWarningReason() bool {
	if o != nil && !IsNil(o.WarningReason) {
		return true
	}

	return false
}

// SetWarningReason gets a reference to the given string and assigns it to the WarningReason field.
func (o *MetricEvent) SetWarningReason(v string) {
	o.WarningReason = &v
}

func (o MetricEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationType) {
		toSerialize["aggregationType"] = o.AggregationType
	}
	if !IsNil(o.AlertingScope) {
		toSerialize["alertingScope"] = o.AlertingScope
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.DisabledReason) {
		toSerialize["disabledReason"] = o.DisabledReason
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MetricDimensions) {
		toSerialize["metricDimensions"] = o.MetricDimensions
	}
	if !IsNil(o.MetricId) {
		toSerialize["metricId"] = o.MetricId
	}
	if !IsNil(o.MetricSelector) {
		toSerialize["metricSelector"] = o.MetricSelector
	}
	toSerialize["monitoringStrategy"] = o.MonitoringStrategy
	toSerialize["name"] = o.Name
	if !IsNil(o.PrimaryDimensionKey) {
		toSerialize["primaryDimensionKey"] = o.PrimaryDimensionKey
	}
	if !IsNil(o.QueryOffset) {
		toSerialize["queryOffset"] = o.QueryOffset
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.WarningReason) {
		toSerialize["warningReason"] = o.WarningReason
	}
	return toSerialize, nil
}

type NullableMetricEvent struct {
	value *MetricEvent
	isSet bool
}

func (v NullableMetricEvent) Get() *MetricEvent {
	return v.value
}

func (v *NullableMetricEvent) Set(val *MetricEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricEvent(val *MetricEvent) *NullableMetricEvent {
	return &NullableMetricEvent{value: val, isSet: true}
}

func (v NullableMetricEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


