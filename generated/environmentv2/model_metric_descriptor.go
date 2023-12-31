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

// checks if the MetricDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricDescriptor{}

// MetricDescriptor The descriptor of a metric.
type MetricDescriptor struct {
	// The list of allowed aggregations for this metric.
	AggregationTypes []string `json:"aggregationTypes,omitempty"`
	// If `true`the usage of metric is billable.   [Metric expressions](https://dt-url.net/metricExpression) don't return this field.
	Billable *bool `json:"billable,omitempty"`
	// The timestamp of metric creation.   Built-in metrics and metric expressions have the value of `null`.
	Created *int64 `json:"created,omitempty"`
	// If `true` the usage of metric consumes [Davis data units](https://dt-url.net/ddu). Deprecated and always `false` for Dynatrace Platform Subscription. Superseded by `isBillable`.   [Metric expressions](https://dt-url.net/metricExpression) don't return this field.
	DduBillable *bool `json:"dduBillable,omitempty"`
	DefaultAggregation *MetricDefaultAggregation `json:"defaultAggregation,omitempty"`
	// A short description of the metric.
	Description *string `json:"description,omitempty"`
	// The cardinalities of MINT metric dimensions.
	DimensionCardinalities []MetricDimensionCardinality `json:"dimensionCardinalities,omitempty"`
	// The fine metric division (for example, process group and process ID for some process-related metric).   For [ingested metrics](https://dt-url.net/5d63ic1), dimensions that doesn't have have any data within the last 15 days are omitted. 
	DimensionDefinitions []MetricDimensionDefinition `json:"dimensionDefinitions,omitempty"`
	// The name of the metric in the user interface.
	DisplayName *string `json:"displayName,omitempty"`
	// List of admissible primary entity types for this metric. Can be used for the `type` predicate in the `entitySelector`.
	EntityType []string `json:"entityType,omitempty"`
	// The metric is (`true`) or is not (`false`) impact relevant.    An impact-relevant metric is highly dependent on other metrics and changes because an underlying root-cause metric has changed.   [Metric expressions](https://dt-url.net/metricExpression) don't return this field.
	ImpactRelevant *bool `json:"impactRelevant,omitempty"`
	// The timestamp when the metric was last written.   Has the value of `null` for metric expressions or if the data has never been written.
	LastWritten *int64 `json:"lastWritten,omitempty"`
	// The latency of the metric, in minutes.    The latency is the expected reporting delay (for example, caused by constraints of cloud vendors or other third-party data sources) between the observation of a metric data point and its availability in Dynatrace.   The allowed value range is from 1 to 60 minutes.   [Metric expressions](https://dt-url.net/metricExpression) don't return this field.
	Latency *int64 `json:"latency,omitempty"`
	// The maximum allowed value of the metric.   [Metric expressions](https://dt-url.net/metricExpression) don't return this field.
	MaximumValue *float64 `json:"maximumValue,omitempty"`
	// The fully qualified key of the metric.   If a transformation has been used it is reflected in the metric key.
	MetricId string `json:"metricId"`
	// The metric selector that is used when querying a func: metric.
	MetricSelector *string `json:"metricSelector,omitempty"`
	MetricValueType *MetricValueType `json:"metricValueType,omitempty"`
	// The minimum allowed value of the metric.   [Metric expressions](https://dt-url.net/metricExpression) don't return this field.
	MinimumValue *float64 `json:"minimumValue,omitempty"`
	// If 'true', resolution=Inf can be applied to the metric query.
	ResolutionInfSupported *bool `json:"resolutionInfSupported,omitempty"`
	// The metric is (`true`) or is not (`false`) root cause relevant.    A root-cause relevant metric represents a strong indicator for a faulty component.   [Metric expressions](https://dt-url.net/metricExpression) don't return this field.
	RootCauseRelevant *bool `json:"rootCauseRelevant,omitempty"`
	// Indicates whether the metric expression resolves to a scalar (`true`) or to a series (`false`).  A scalar result always contains one data point. The amount of data points in a series result depends on the resolution you're using.
	Scalar *bool `json:"scalar,omitempty"`
	// The tags applied to the metric.    [Metric expressions](https://dt-url.net/metricExpression) don't return this field.
	Tags []string `json:"tags,omitempty"`
	// Transform operators that could be appended to the current transformation list.
	Transformations []string `json:"transformations,omitempty"`
	// The unit of the metric.
	Unit *string `json:"unit,omitempty"`
	// The raw value is stored in bits or bytes. The user interface can display it in these numeral systems:  Binary: 1 MiB = 1024 KiB = 1,048,576 bytes  Decimal: 1 MB = 1000 kB = 1,000,000 bytes  If not set, the decimal system is used.   [Metric expressions](https://dt-url.net/metricExpression) don't return this field.
	UnitDisplayFormat *string `json:"unitDisplayFormat,omitempty"`
	// A list of potential warnings that affect this ID. For example deprecated feature usage etc.
	Warnings []string `json:"warnings,omitempty"`
}

// NewMetricDescriptor instantiates a new MetricDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricDescriptor(metricId string) *MetricDescriptor {
	this := MetricDescriptor{}
	this.MetricId = metricId
	return &this
}

// NewMetricDescriptorWithDefaults instantiates a new MetricDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricDescriptorWithDefaults() *MetricDescriptor {
	this := MetricDescriptor{}
	return &this
}

// GetAggregationTypes returns the AggregationTypes field value if set, zero value otherwise.
func (o *MetricDescriptor) GetAggregationTypes() []string {
	if o == nil || IsNil(o.AggregationTypes) {
		var ret []string
		return ret
	}
	return o.AggregationTypes
}

// GetAggregationTypesOk returns a tuple with the AggregationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetAggregationTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AggregationTypes) {
		return nil, false
	}
	return o.AggregationTypes, true
}

// HasAggregationTypes returns a boolean if a field has been set.
func (o *MetricDescriptor) HasAggregationTypes() bool {
	if o != nil && !IsNil(o.AggregationTypes) {
		return true
	}

	return false
}

// SetAggregationTypes gets a reference to the given []string and assigns it to the AggregationTypes field.
func (o *MetricDescriptor) SetAggregationTypes(v []string) {
	o.AggregationTypes = v
}

// GetBillable returns the Billable field value if set, zero value otherwise.
func (o *MetricDescriptor) GetBillable() bool {
	if o == nil || IsNil(o.Billable) {
		var ret bool
		return ret
	}
	return *o.Billable
}

// GetBillableOk returns a tuple with the Billable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetBillableOk() (*bool, bool) {
	if o == nil || IsNil(o.Billable) {
		return nil, false
	}
	return o.Billable, true
}

// HasBillable returns a boolean if a field has been set.
func (o *MetricDescriptor) HasBillable() bool {
	if o != nil && !IsNil(o.Billable) {
		return true
	}

	return false
}

// SetBillable gets a reference to the given bool and assigns it to the Billable field.
func (o *MetricDescriptor) SetBillable(v bool) {
	o.Billable = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *MetricDescriptor) GetCreated() int64 {
	if o == nil || IsNil(o.Created) {
		var ret int64
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetCreatedOk() (*int64, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *MetricDescriptor) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given int64 and assigns it to the Created field.
func (o *MetricDescriptor) SetCreated(v int64) {
	o.Created = &v
}

// GetDduBillable returns the DduBillable field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDduBillable() bool {
	if o == nil || IsNil(o.DduBillable) {
		var ret bool
		return ret
	}
	return *o.DduBillable
}

// GetDduBillableOk returns a tuple with the DduBillable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDduBillableOk() (*bool, bool) {
	if o == nil || IsNil(o.DduBillable) {
		return nil, false
	}
	return o.DduBillable, true
}

// HasDduBillable returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDduBillable() bool {
	if o != nil && !IsNil(o.DduBillable) {
		return true
	}

	return false
}

// SetDduBillable gets a reference to the given bool and assigns it to the DduBillable field.
func (o *MetricDescriptor) SetDduBillable(v bool) {
	o.DduBillable = &v
}

// GetDefaultAggregation returns the DefaultAggregation field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDefaultAggregation() MetricDefaultAggregation {
	if o == nil || IsNil(o.DefaultAggregation) {
		var ret MetricDefaultAggregation
		return ret
	}
	return *o.DefaultAggregation
}

// GetDefaultAggregationOk returns a tuple with the DefaultAggregation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDefaultAggregationOk() (*MetricDefaultAggregation, bool) {
	if o == nil || IsNil(o.DefaultAggregation) {
		return nil, false
	}
	return o.DefaultAggregation, true
}

// HasDefaultAggregation returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDefaultAggregation() bool {
	if o != nil && !IsNil(o.DefaultAggregation) {
		return true
	}

	return false
}

// SetDefaultAggregation gets a reference to the given MetricDefaultAggregation and assigns it to the DefaultAggregation field.
func (o *MetricDescriptor) SetDefaultAggregation(v MetricDefaultAggregation) {
	o.DefaultAggregation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MetricDescriptor) SetDescription(v string) {
	o.Description = &v
}

// GetDimensionCardinalities returns the DimensionCardinalities field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDimensionCardinalities() []MetricDimensionCardinality {
	if o == nil || IsNil(o.DimensionCardinalities) {
		var ret []MetricDimensionCardinality
		return ret
	}
	return o.DimensionCardinalities
}

// GetDimensionCardinalitiesOk returns a tuple with the DimensionCardinalities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDimensionCardinalitiesOk() ([]MetricDimensionCardinality, bool) {
	if o == nil || IsNil(o.DimensionCardinalities) {
		return nil, false
	}
	return o.DimensionCardinalities, true
}

// HasDimensionCardinalities returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDimensionCardinalities() bool {
	if o != nil && !IsNil(o.DimensionCardinalities) {
		return true
	}

	return false
}

// SetDimensionCardinalities gets a reference to the given []MetricDimensionCardinality and assigns it to the DimensionCardinalities field.
func (o *MetricDescriptor) SetDimensionCardinalities(v []MetricDimensionCardinality) {
	o.DimensionCardinalities = v
}

// GetDimensionDefinitions returns the DimensionDefinitions field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDimensionDefinitions() []MetricDimensionDefinition {
	if o == nil || IsNil(o.DimensionDefinitions) {
		var ret []MetricDimensionDefinition
		return ret
	}
	return o.DimensionDefinitions
}

// GetDimensionDefinitionsOk returns a tuple with the DimensionDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDimensionDefinitionsOk() ([]MetricDimensionDefinition, bool) {
	if o == nil || IsNil(o.DimensionDefinitions) {
		return nil, false
	}
	return o.DimensionDefinitions, true
}

// HasDimensionDefinitions returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDimensionDefinitions() bool {
	if o != nil && !IsNil(o.DimensionDefinitions) {
		return true
	}

	return false
}

// SetDimensionDefinitions gets a reference to the given []MetricDimensionDefinition and assigns it to the DimensionDefinitions field.
func (o *MetricDescriptor) SetDimensionDefinitions(v []MetricDimensionDefinition) {
	o.DimensionDefinitions = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MetricDescriptor) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MetricDescriptor) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MetricDescriptor) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *MetricDescriptor) GetEntityType() []string {
	if o == nil || IsNil(o.EntityType) {
		var ret []string
		return ret
	}
	return o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetEntityTypeOk() ([]string, bool) {
	if o == nil || IsNil(o.EntityType) {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *MetricDescriptor) HasEntityType() bool {
	if o != nil && !IsNil(o.EntityType) {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given []string and assigns it to the EntityType field.
func (o *MetricDescriptor) SetEntityType(v []string) {
	o.EntityType = v
}

// GetImpactRelevant returns the ImpactRelevant field value if set, zero value otherwise.
func (o *MetricDescriptor) GetImpactRelevant() bool {
	if o == nil || IsNil(o.ImpactRelevant) {
		var ret bool
		return ret
	}
	return *o.ImpactRelevant
}

// GetImpactRelevantOk returns a tuple with the ImpactRelevant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetImpactRelevantOk() (*bool, bool) {
	if o == nil || IsNil(o.ImpactRelevant) {
		return nil, false
	}
	return o.ImpactRelevant, true
}

// HasImpactRelevant returns a boolean if a field has been set.
func (o *MetricDescriptor) HasImpactRelevant() bool {
	if o != nil && !IsNil(o.ImpactRelevant) {
		return true
	}

	return false
}

// SetImpactRelevant gets a reference to the given bool and assigns it to the ImpactRelevant field.
func (o *MetricDescriptor) SetImpactRelevant(v bool) {
	o.ImpactRelevant = &v
}

// GetLastWritten returns the LastWritten field value if set, zero value otherwise.
func (o *MetricDescriptor) GetLastWritten() int64 {
	if o == nil || IsNil(o.LastWritten) {
		var ret int64
		return ret
	}
	return *o.LastWritten
}

// GetLastWrittenOk returns a tuple with the LastWritten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetLastWrittenOk() (*int64, bool) {
	if o == nil || IsNil(o.LastWritten) {
		return nil, false
	}
	return o.LastWritten, true
}

// HasLastWritten returns a boolean if a field has been set.
func (o *MetricDescriptor) HasLastWritten() bool {
	if o != nil && !IsNil(o.LastWritten) {
		return true
	}

	return false
}

// SetLastWritten gets a reference to the given int64 and assigns it to the LastWritten field.
func (o *MetricDescriptor) SetLastWritten(v int64) {
	o.LastWritten = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *MetricDescriptor) GetLatency() int64 {
	if o == nil || IsNil(o.Latency) {
		var ret int64
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetLatencyOk() (*int64, bool) {
	if o == nil || IsNil(o.Latency) {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *MetricDescriptor) HasLatency() bool {
	if o != nil && !IsNil(o.Latency) {
		return true
	}

	return false
}

// SetLatency gets a reference to the given int64 and assigns it to the Latency field.
func (o *MetricDescriptor) SetLatency(v int64) {
	o.Latency = &v
}

// GetMaximumValue returns the MaximumValue field value if set, zero value otherwise.
func (o *MetricDescriptor) GetMaximumValue() float64 {
	if o == nil || IsNil(o.MaximumValue) {
		var ret float64
		return ret
	}
	return *o.MaximumValue
}

// GetMaximumValueOk returns a tuple with the MaximumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetMaximumValueOk() (*float64, bool) {
	if o == nil || IsNil(o.MaximumValue) {
		return nil, false
	}
	return o.MaximumValue, true
}

// HasMaximumValue returns a boolean if a field has been set.
func (o *MetricDescriptor) HasMaximumValue() bool {
	if o != nil && !IsNil(o.MaximumValue) {
		return true
	}

	return false
}

// SetMaximumValue gets a reference to the given float64 and assigns it to the MaximumValue field.
func (o *MetricDescriptor) SetMaximumValue(v float64) {
	o.MaximumValue = &v
}

// GetMetricId returns the MetricId field value
func (o *MetricDescriptor) GetMetricId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricId
}

// GetMetricIdOk returns a tuple with the MetricId field value
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetMetricIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricId, true
}

// SetMetricId sets field value
func (o *MetricDescriptor) SetMetricId(v string) {
	o.MetricId = v
}

// GetMetricSelector returns the MetricSelector field value if set, zero value otherwise.
func (o *MetricDescriptor) GetMetricSelector() string {
	if o == nil || IsNil(o.MetricSelector) {
		var ret string
		return ret
	}
	return *o.MetricSelector
}

// GetMetricSelectorOk returns a tuple with the MetricSelector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetMetricSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.MetricSelector) {
		return nil, false
	}
	return o.MetricSelector, true
}

// HasMetricSelector returns a boolean if a field has been set.
func (o *MetricDescriptor) HasMetricSelector() bool {
	if o != nil && !IsNil(o.MetricSelector) {
		return true
	}

	return false
}

// SetMetricSelector gets a reference to the given string and assigns it to the MetricSelector field.
func (o *MetricDescriptor) SetMetricSelector(v string) {
	o.MetricSelector = &v
}

// GetMetricValueType returns the MetricValueType field value if set, zero value otherwise.
func (o *MetricDescriptor) GetMetricValueType() MetricValueType {
	if o == nil || IsNil(o.MetricValueType) {
		var ret MetricValueType
		return ret
	}
	return *o.MetricValueType
}

// GetMetricValueTypeOk returns a tuple with the MetricValueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetMetricValueTypeOk() (*MetricValueType, bool) {
	if o == nil || IsNil(o.MetricValueType) {
		return nil, false
	}
	return o.MetricValueType, true
}

// HasMetricValueType returns a boolean if a field has been set.
func (o *MetricDescriptor) HasMetricValueType() bool {
	if o != nil && !IsNil(o.MetricValueType) {
		return true
	}

	return false
}

// SetMetricValueType gets a reference to the given MetricValueType and assigns it to the MetricValueType field.
func (o *MetricDescriptor) SetMetricValueType(v MetricValueType) {
	o.MetricValueType = &v
}

// GetMinimumValue returns the MinimumValue field value if set, zero value otherwise.
func (o *MetricDescriptor) GetMinimumValue() float64 {
	if o == nil || IsNil(o.MinimumValue) {
		var ret float64
		return ret
	}
	return *o.MinimumValue
}

// GetMinimumValueOk returns a tuple with the MinimumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetMinimumValueOk() (*float64, bool) {
	if o == nil || IsNil(o.MinimumValue) {
		return nil, false
	}
	return o.MinimumValue, true
}

// HasMinimumValue returns a boolean if a field has been set.
func (o *MetricDescriptor) HasMinimumValue() bool {
	if o != nil && !IsNil(o.MinimumValue) {
		return true
	}

	return false
}

// SetMinimumValue gets a reference to the given float64 and assigns it to the MinimumValue field.
func (o *MetricDescriptor) SetMinimumValue(v float64) {
	o.MinimumValue = &v
}

// GetResolutionInfSupported returns the ResolutionInfSupported field value if set, zero value otherwise.
func (o *MetricDescriptor) GetResolutionInfSupported() bool {
	if o == nil || IsNil(o.ResolutionInfSupported) {
		var ret bool
		return ret
	}
	return *o.ResolutionInfSupported
}

// GetResolutionInfSupportedOk returns a tuple with the ResolutionInfSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetResolutionInfSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.ResolutionInfSupported) {
		return nil, false
	}
	return o.ResolutionInfSupported, true
}

// HasResolutionInfSupported returns a boolean if a field has been set.
func (o *MetricDescriptor) HasResolutionInfSupported() bool {
	if o != nil && !IsNil(o.ResolutionInfSupported) {
		return true
	}

	return false
}

// SetResolutionInfSupported gets a reference to the given bool and assigns it to the ResolutionInfSupported field.
func (o *MetricDescriptor) SetResolutionInfSupported(v bool) {
	o.ResolutionInfSupported = &v
}

// GetRootCauseRelevant returns the RootCauseRelevant field value if set, zero value otherwise.
func (o *MetricDescriptor) GetRootCauseRelevant() bool {
	if o == nil || IsNil(o.RootCauseRelevant) {
		var ret bool
		return ret
	}
	return *o.RootCauseRelevant
}

// GetRootCauseRelevantOk returns a tuple with the RootCauseRelevant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetRootCauseRelevantOk() (*bool, bool) {
	if o == nil || IsNil(o.RootCauseRelevant) {
		return nil, false
	}
	return o.RootCauseRelevant, true
}

// HasRootCauseRelevant returns a boolean if a field has been set.
func (o *MetricDescriptor) HasRootCauseRelevant() bool {
	if o != nil && !IsNil(o.RootCauseRelevant) {
		return true
	}

	return false
}

// SetRootCauseRelevant gets a reference to the given bool and assigns it to the RootCauseRelevant field.
func (o *MetricDescriptor) SetRootCauseRelevant(v bool) {
	o.RootCauseRelevant = &v
}

// GetScalar returns the Scalar field value if set, zero value otherwise.
func (o *MetricDescriptor) GetScalar() bool {
	if o == nil || IsNil(o.Scalar) {
		var ret bool
		return ret
	}
	return *o.Scalar
}

// GetScalarOk returns a tuple with the Scalar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetScalarOk() (*bool, bool) {
	if o == nil || IsNil(o.Scalar) {
		return nil, false
	}
	return o.Scalar, true
}

// HasScalar returns a boolean if a field has been set.
func (o *MetricDescriptor) HasScalar() bool {
	if o != nil && !IsNil(o.Scalar) {
		return true
	}

	return false
}

// SetScalar gets a reference to the given bool and assigns it to the Scalar field.
func (o *MetricDescriptor) SetScalar(v bool) {
	o.Scalar = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MetricDescriptor) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MetricDescriptor) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MetricDescriptor) SetTags(v []string) {
	o.Tags = v
}

// GetTransformations returns the Transformations field value if set, zero value otherwise.
func (o *MetricDescriptor) GetTransformations() []string {
	if o == nil || IsNil(o.Transformations) {
		var ret []string
		return ret
	}
	return o.Transformations
}

// GetTransformationsOk returns a tuple with the Transformations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetTransformationsOk() ([]string, bool) {
	if o == nil || IsNil(o.Transformations) {
		return nil, false
	}
	return o.Transformations, true
}

// HasTransformations returns a boolean if a field has been set.
func (o *MetricDescriptor) HasTransformations() bool {
	if o != nil && !IsNil(o.Transformations) {
		return true
	}

	return false
}

// SetTransformations gets a reference to the given []string and assigns it to the Transformations field.
func (o *MetricDescriptor) SetTransformations(v []string) {
	o.Transformations = v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *MetricDescriptor) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *MetricDescriptor) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *MetricDescriptor) SetUnit(v string) {
	o.Unit = &v
}

// GetUnitDisplayFormat returns the UnitDisplayFormat field value if set, zero value otherwise.
func (o *MetricDescriptor) GetUnitDisplayFormat() string {
	if o == nil || IsNil(o.UnitDisplayFormat) {
		var ret string
		return ret
	}
	return *o.UnitDisplayFormat
}

// GetUnitDisplayFormatOk returns a tuple with the UnitDisplayFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetUnitDisplayFormatOk() (*string, bool) {
	if o == nil || IsNil(o.UnitDisplayFormat) {
		return nil, false
	}
	return o.UnitDisplayFormat, true
}

// HasUnitDisplayFormat returns a boolean if a field has been set.
func (o *MetricDescriptor) HasUnitDisplayFormat() bool {
	if o != nil && !IsNil(o.UnitDisplayFormat) {
		return true
	}

	return false
}

// SetUnitDisplayFormat gets a reference to the given string and assigns it to the UnitDisplayFormat field.
func (o *MetricDescriptor) SetUnitDisplayFormat(v string) {
	o.UnitDisplayFormat = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *MetricDescriptor) GetWarnings() []string {
	if o == nil || IsNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricDescriptor) GetWarningsOk() ([]string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *MetricDescriptor) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *MetricDescriptor) SetWarnings(v []string) {
	o.Warnings = v
}

func (o MetricDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AggregationTypes) {
		toSerialize["aggregationTypes"] = o.AggregationTypes
	}
	if !IsNil(o.Billable) {
		toSerialize["billable"] = o.Billable
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.DduBillable) {
		toSerialize["dduBillable"] = o.DduBillable
	}
	if !IsNil(o.DefaultAggregation) {
		toSerialize["defaultAggregation"] = o.DefaultAggregation
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DimensionCardinalities) {
		toSerialize["dimensionCardinalities"] = o.DimensionCardinalities
	}
	if !IsNil(o.DimensionDefinitions) {
		toSerialize["dimensionDefinitions"] = o.DimensionDefinitions
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.EntityType) {
		toSerialize["entityType"] = o.EntityType
	}
	if !IsNil(o.ImpactRelevant) {
		toSerialize["impactRelevant"] = o.ImpactRelevant
	}
	if !IsNil(o.LastWritten) {
		toSerialize["lastWritten"] = o.LastWritten
	}
	if !IsNil(o.Latency) {
		toSerialize["latency"] = o.Latency
	}
	if !IsNil(o.MaximumValue) {
		toSerialize["maximumValue"] = o.MaximumValue
	}
	toSerialize["metricId"] = o.MetricId
	if !IsNil(o.MetricSelector) {
		toSerialize["metricSelector"] = o.MetricSelector
	}
	if !IsNil(o.MetricValueType) {
		toSerialize["metricValueType"] = o.MetricValueType
	}
	if !IsNil(o.MinimumValue) {
		toSerialize["minimumValue"] = o.MinimumValue
	}
	if !IsNil(o.ResolutionInfSupported) {
		toSerialize["resolutionInfSupported"] = o.ResolutionInfSupported
	}
	if !IsNil(o.RootCauseRelevant) {
		toSerialize["rootCauseRelevant"] = o.RootCauseRelevant
	}
	if !IsNil(o.Scalar) {
		toSerialize["scalar"] = o.Scalar
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Transformations) {
		toSerialize["transformations"] = o.Transformations
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !IsNil(o.UnitDisplayFormat) {
		toSerialize["unitDisplayFormat"] = o.UnitDisplayFormat
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableMetricDescriptor struct {
	value *MetricDescriptor
	isSet bool
}

func (v NullableMetricDescriptor) Get() *MetricDescriptor {
	return v.value
}

func (v *NullableMetricDescriptor) Set(val *MetricDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricDescriptor(val *MetricDescriptor) *NullableMetricDescriptor {
	return &NullableMetricDescriptor{value: val, isSet: true}
}

func (v NullableMetricDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


