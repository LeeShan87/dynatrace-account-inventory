# MetricDescriptor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationTypes** | Pointer to **[]string** | The list of allowed aggregations for this metric. | [optional] 
**Billable** | Pointer to **bool** | If &#x60;true&#x60;the usage of metric is billable.   [Metric expressions](https://dt-url.net/metricExpression) don&#39;t return this field. | [optional] 
**Created** | Pointer to **int64** | The timestamp of metric creation.   Built-in metrics and metric expressions have the value of &#x60;null&#x60;. | [optional] 
**DduBillable** | Pointer to **bool** | If &#x60;true&#x60; the usage of metric consumes [Davis data units](https://dt-url.net/ddu). Deprecated and always &#x60;false&#x60; for Dynatrace Platform Subscription. Superseded by &#x60;isBillable&#x60;.   [Metric expressions](https://dt-url.net/metricExpression) don&#39;t return this field. | [optional] 
**DefaultAggregation** | Pointer to [**MetricDefaultAggregation**](MetricDefaultAggregation.md) |  | [optional] 
**Description** | Pointer to **string** | A short description of the metric. | [optional] 
**DimensionCardinalities** | Pointer to [**[]MetricDimensionCardinality**](MetricDimensionCardinality.md) | The cardinalities of MINT metric dimensions. | [optional] 
**DimensionDefinitions** | Pointer to [**[]MetricDimensionDefinition**](MetricDimensionDefinition.md) | The fine metric division (for example, process group and process ID for some process-related metric).   For [ingested metrics](https://dt-url.net/5d63ic1), dimensions that doesn&#39;t have have any data within the last 15 days are omitted.  | [optional] 
**DisplayName** | Pointer to **string** | The name of the metric in the user interface. | [optional] 
**EntityType** | Pointer to **[]string** | List of admissible primary entity types for this metric. Can be used for the &#x60;type&#x60; predicate in the &#x60;entitySelector&#x60;. | [optional] 
**ImpactRelevant** | Pointer to **bool** | The metric is (&#x60;true&#x60;) or is not (&#x60;false&#x60;) impact relevant.    An impact-relevant metric is highly dependent on other metrics and changes because an underlying root-cause metric has changed.   [Metric expressions](https://dt-url.net/metricExpression) don&#39;t return this field. | [optional] 
**LastWritten** | Pointer to **int64** | The timestamp when the metric was last written.   Has the value of &#x60;null&#x60; for metric expressions or if the data has never been written. | [optional] 
**Latency** | Pointer to **int64** | The latency of the metric, in minutes.    The latency is the expected reporting delay (for example, caused by constraints of cloud vendors or other third-party data sources) between the observation of a metric data point and its availability in Dynatrace.   The allowed value range is from 1 to 60 minutes.   [Metric expressions](https://dt-url.net/metricExpression) don&#39;t return this field. | [optional] 
**MaximumValue** | Pointer to **float64** | The maximum allowed value of the metric.   [Metric expressions](https://dt-url.net/metricExpression) don&#39;t return this field. | [optional] 
**MetricId** | **string** | The fully qualified key of the metric.   If a transformation has been used it is reflected in the metric key. | 
**MetricSelector** | Pointer to **string** | The metric selector that is used when querying a func: metric. | [optional] 
**MetricValueType** | Pointer to [**MetricValueType**](MetricValueType.md) |  | [optional] 
**MinimumValue** | Pointer to **float64** | The minimum allowed value of the metric.   [Metric expressions](https://dt-url.net/metricExpression) don&#39;t return this field. | [optional] 
**ResolutionInfSupported** | Pointer to **bool** | If &#39;true&#39;, resolution&#x3D;Inf can be applied to the metric query. | [optional] 
**RootCauseRelevant** | Pointer to **bool** | The metric is (&#x60;true&#x60;) or is not (&#x60;false&#x60;) root cause relevant.    A root-cause relevant metric represents a strong indicator for a faulty component.   [Metric expressions](https://dt-url.net/metricExpression) don&#39;t return this field. | [optional] 
**Scalar** | Pointer to **bool** | Indicates whether the metric expression resolves to a scalar (&#x60;true&#x60;) or to a series (&#x60;false&#x60;).  A scalar result always contains one data point. The amount of data points in a series result depends on the resolution you&#39;re using. | [optional] 
**Tags** | Pointer to **[]string** | The tags applied to the metric.    [Metric expressions](https://dt-url.net/metricExpression) don&#39;t return this field. | [optional] 
**Transformations** | Pointer to **[]string** | Transform operators that could be appended to the current transformation list. | [optional] 
**Unit** | Pointer to **string** | The unit of the metric. | [optional] 
**UnitDisplayFormat** | Pointer to **string** | The raw value is stored in bits or bytes. The user interface can display it in these numeral systems:  Binary: 1 MiB &#x3D; 1024 KiB &#x3D; 1,048,576 bytes  Decimal: 1 MB &#x3D; 1000 kB &#x3D; 1,000,000 bytes  If not set, the decimal system is used.   [Metric expressions](https://dt-url.net/metricExpression) don&#39;t return this field. | [optional] 
**Warnings** | Pointer to **[]string** | A list of potential warnings that affect this ID. For example deprecated feature usage etc. | [optional] 

## Methods

### NewMetricDescriptor

`func NewMetricDescriptor(metricId string, ) *MetricDescriptor`

NewMetricDescriptor instantiates a new MetricDescriptor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDescriptorWithDefaults

`func NewMetricDescriptorWithDefaults() *MetricDescriptor`

NewMetricDescriptorWithDefaults instantiates a new MetricDescriptor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationTypes

`func (o *MetricDescriptor) GetAggregationTypes() []string`

GetAggregationTypes returns the AggregationTypes field if non-nil, zero value otherwise.

### GetAggregationTypesOk

`func (o *MetricDescriptor) GetAggregationTypesOk() (*[]string, bool)`

GetAggregationTypesOk returns a tuple with the AggregationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationTypes

`func (o *MetricDescriptor) SetAggregationTypes(v []string)`

SetAggregationTypes sets AggregationTypes field to given value.

### HasAggregationTypes

`func (o *MetricDescriptor) HasAggregationTypes() bool`

HasAggregationTypes returns a boolean if a field has been set.

### GetBillable

`func (o *MetricDescriptor) GetBillable() bool`

GetBillable returns the Billable field if non-nil, zero value otherwise.

### GetBillableOk

`func (o *MetricDescriptor) GetBillableOk() (*bool, bool)`

GetBillableOk returns a tuple with the Billable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillable

`func (o *MetricDescriptor) SetBillable(v bool)`

SetBillable sets Billable field to given value.

### HasBillable

`func (o *MetricDescriptor) HasBillable() bool`

HasBillable returns a boolean if a field has been set.

### GetCreated

`func (o *MetricDescriptor) GetCreated() int64`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *MetricDescriptor) GetCreatedOk() (*int64, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *MetricDescriptor) SetCreated(v int64)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *MetricDescriptor) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDduBillable

`func (o *MetricDescriptor) GetDduBillable() bool`

GetDduBillable returns the DduBillable field if non-nil, zero value otherwise.

### GetDduBillableOk

`func (o *MetricDescriptor) GetDduBillableOk() (*bool, bool)`

GetDduBillableOk returns a tuple with the DduBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDduBillable

`func (o *MetricDescriptor) SetDduBillable(v bool)`

SetDduBillable sets DduBillable field to given value.

### HasDduBillable

`func (o *MetricDescriptor) HasDduBillable() bool`

HasDduBillable returns a boolean if a field has been set.

### GetDefaultAggregation

`func (o *MetricDescriptor) GetDefaultAggregation() MetricDefaultAggregation`

GetDefaultAggregation returns the DefaultAggregation field if non-nil, zero value otherwise.

### GetDefaultAggregationOk

`func (o *MetricDescriptor) GetDefaultAggregationOk() (*MetricDefaultAggregation, bool)`

GetDefaultAggregationOk returns a tuple with the DefaultAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAggregation

`func (o *MetricDescriptor) SetDefaultAggregation(v MetricDefaultAggregation)`

SetDefaultAggregation sets DefaultAggregation field to given value.

### HasDefaultAggregation

`func (o *MetricDescriptor) HasDefaultAggregation() bool`

HasDefaultAggregation returns a boolean if a field has been set.

### GetDescription

`func (o *MetricDescriptor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricDescriptor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricDescriptor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricDescriptor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDimensionCardinalities

`func (o *MetricDescriptor) GetDimensionCardinalities() []MetricDimensionCardinality`

GetDimensionCardinalities returns the DimensionCardinalities field if non-nil, zero value otherwise.

### GetDimensionCardinalitiesOk

`func (o *MetricDescriptor) GetDimensionCardinalitiesOk() (*[]MetricDimensionCardinality, bool)`

GetDimensionCardinalitiesOk returns a tuple with the DimensionCardinalities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionCardinalities

`func (o *MetricDescriptor) SetDimensionCardinalities(v []MetricDimensionCardinality)`

SetDimensionCardinalities sets DimensionCardinalities field to given value.

### HasDimensionCardinalities

`func (o *MetricDescriptor) HasDimensionCardinalities() bool`

HasDimensionCardinalities returns a boolean if a field has been set.

### GetDimensionDefinitions

`func (o *MetricDescriptor) GetDimensionDefinitions() []MetricDimensionDefinition`

GetDimensionDefinitions returns the DimensionDefinitions field if non-nil, zero value otherwise.

### GetDimensionDefinitionsOk

`func (o *MetricDescriptor) GetDimensionDefinitionsOk() (*[]MetricDimensionDefinition, bool)`

GetDimensionDefinitionsOk returns a tuple with the DimensionDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionDefinitions

`func (o *MetricDescriptor) SetDimensionDefinitions(v []MetricDimensionDefinition)`

SetDimensionDefinitions sets DimensionDefinitions field to given value.

### HasDimensionDefinitions

`func (o *MetricDescriptor) HasDimensionDefinitions() bool`

HasDimensionDefinitions returns a boolean if a field has been set.

### GetDisplayName

`func (o *MetricDescriptor) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MetricDescriptor) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MetricDescriptor) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MetricDescriptor) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEntityType

`func (o *MetricDescriptor) GetEntityType() []string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *MetricDescriptor) GetEntityTypeOk() (*[]string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *MetricDescriptor) SetEntityType(v []string)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *MetricDescriptor) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetImpactRelevant

`func (o *MetricDescriptor) GetImpactRelevant() bool`

GetImpactRelevant returns the ImpactRelevant field if non-nil, zero value otherwise.

### GetImpactRelevantOk

`func (o *MetricDescriptor) GetImpactRelevantOk() (*bool, bool)`

GetImpactRelevantOk returns a tuple with the ImpactRelevant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpactRelevant

`func (o *MetricDescriptor) SetImpactRelevant(v bool)`

SetImpactRelevant sets ImpactRelevant field to given value.

### HasImpactRelevant

`func (o *MetricDescriptor) HasImpactRelevant() bool`

HasImpactRelevant returns a boolean if a field has been set.

### GetLastWritten

`func (o *MetricDescriptor) GetLastWritten() int64`

GetLastWritten returns the LastWritten field if non-nil, zero value otherwise.

### GetLastWrittenOk

`func (o *MetricDescriptor) GetLastWrittenOk() (*int64, bool)`

GetLastWrittenOk returns a tuple with the LastWritten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastWritten

`func (o *MetricDescriptor) SetLastWritten(v int64)`

SetLastWritten sets LastWritten field to given value.

### HasLastWritten

`func (o *MetricDescriptor) HasLastWritten() bool`

HasLastWritten returns a boolean if a field has been set.

### GetLatency

`func (o *MetricDescriptor) GetLatency() int64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *MetricDescriptor) GetLatencyOk() (*int64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *MetricDescriptor) SetLatency(v int64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *MetricDescriptor) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetMaximumValue

`func (o *MetricDescriptor) GetMaximumValue() float64`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *MetricDescriptor) GetMaximumValueOk() (*float64, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *MetricDescriptor) SetMaximumValue(v float64)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *MetricDescriptor) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### GetMetricId

`func (o *MetricDescriptor) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricDescriptor) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricDescriptor) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.


### GetMetricSelector

`func (o *MetricDescriptor) GetMetricSelector() string`

GetMetricSelector returns the MetricSelector field if non-nil, zero value otherwise.

### GetMetricSelectorOk

`func (o *MetricDescriptor) GetMetricSelectorOk() (*string, bool)`

GetMetricSelectorOk returns a tuple with the MetricSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricSelector

`func (o *MetricDescriptor) SetMetricSelector(v string)`

SetMetricSelector sets MetricSelector field to given value.

### HasMetricSelector

`func (o *MetricDescriptor) HasMetricSelector() bool`

HasMetricSelector returns a boolean if a field has been set.

### GetMetricValueType

`func (o *MetricDescriptor) GetMetricValueType() MetricValueType`

GetMetricValueType returns the MetricValueType field if non-nil, zero value otherwise.

### GetMetricValueTypeOk

`func (o *MetricDescriptor) GetMetricValueTypeOk() (*MetricValueType, bool)`

GetMetricValueTypeOk returns a tuple with the MetricValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricValueType

`func (o *MetricDescriptor) SetMetricValueType(v MetricValueType)`

SetMetricValueType sets MetricValueType field to given value.

### HasMetricValueType

`func (o *MetricDescriptor) HasMetricValueType() bool`

HasMetricValueType returns a boolean if a field has been set.

### GetMinimumValue

`func (o *MetricDescriptor) GetMinimumValue() float64`

GetMinimumValue returns the MinimumValue field if non-nil, zero value otherwise.

### GetMinimumValueOk

`func (o *MetricDescriptor) GetMinimumValueOk() (*float64, bool)`

GetMinimumValueOk returns a tuple with the MinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumValue

`func (o *MetricDescriptor) SetMinimumValue(v float64)`

SetMinimumValue sets MinimumValue field to given value.

### HasMinimumValue

`func (o *MetricDescriptor) HasMinimumValue() bool`

HasMinimumValue returns a boolean if a field has been set.

### GetResolutionInfSupported

`func (o *MetricDescriptor) GetResolutionInfSupported() bool`

GetResolutionInfSupported returns the ResolutionInfSupported field if non-nil, zero value otherwise.

### GetResolutionInfSupportedOk

`func (o *MetricDescriptor) GetResolutionInfSupportedOk() (*bool, bool)`

GetResolutionInfSupportedOk returns a tuple with the ResolutionInfSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolutionInfSupported

`func (o *MetricDescriptor) SetResolutionInfSupported(v bool)`

SetResolutionInfSupported sets ResolutionInfSupported field to given value.

### HasResolutionInfSupported

`func (o *MetricDescriptor) HasResolutionInfSupported() bool`

HasResolutionInfSupported returns a boolean if a field has been set.

### GetRootCauseRelevant

`func (o *MetricDescriptor) GetRootCauseRelevant() bool`

GetRootCauseRelevant returns the RootCauseRelevant field if non-nil, zero value otherwise.

### GetRootCauseRelevantOk

`func (o *MetricDescriptor) GetRootCauseRelevantOk() (*bool, bool)`

GetRootCauseRelevantOk returns a tuple with the RootCauseRelevant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCauseRelevant

`func (o *MetricDescriptor) SetRootCauseRelevant(v bool)`

SetRootCauseRelevant sets RootCauseRelevant field to given value.

### HasRootCauseRelevant

`func (o *MetricDescriptor) HasRootCauseRelevant() bool`

HasRootCauseRelevant returns a boolean if a field has been set.

### GetScalar

`func (o *MetricDescriptor) GetScalar() bool`

GetScalar returns the Scalar field if non-nil, zero value otherwise.

### GetScalarOk

`func (o *MetricDescriptor) GetScalarOk() (*bool, bool)`

GetScalarOk returns a tuple with the Scalar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalar

`func (o *MetricDescriptor) SetScalar(v bool)`

SetScalar sets Scalar field to given value.

### HasScalar

`func (o *MetricDescriptor) HasScalar() bool`

HasScalar returns a boolean if a field has been set.

### GetTags

`func (o *MetricDescriptor) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *MetricDescriptor) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *MetricDescriptor) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *MetricDescriptor) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTransformations

`func (o *MetricDescriptor) GetTransformations() []string`

GetTransformations returns the Transformations field if non-nil, zero value otherwise.

### GetTransformationsOk

`func (o *MetricDescriptor) GetTransformationsOk() (*[]string, bool)`

GetTransformationsOk returns a tuple with the Transformations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransformations

`func (o *MetricDescriptor) SetTransformations(v []string)`

SetTransformations sets Transformations field to given value.

### HasTransformations

`func (o *MetricDescriptor) HasTransformations() bool`

HasTransformations returns a boolean if a field has been set.

### GetUnit

`func (o *MetricDescriptor) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricDescriptor) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricDescriptor) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricDescriptor) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetUnitDisplayFormat

`func (o *MetricDescriptor) GetUnitDisplayFormat() string`

GetUnitDisplayFormat returns the UnitDisplayFormat field if non-nil, zero value otherwise.

### GetUnitDisplayFormatOk

`func (o *MetricDescriptor) GetUnitDisplayFormatOk() (*string, bool)`

GetUnitDisplayFormatOk returns a tuple with the UnitDisplayFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitDisplayFormat

`func (o *MetricDescriptor) SetUnitDisplayFormat(v string)`

SetUnitDisplayFormat sets UnitDisplayFormat field to given value.

### HasUnitDisplayFormat

`func (o *MetricDescriptor) HasUnitDisplayFormat() bool`

HasUnitDisplayFormat returns a boolean if a field has been set.

### GetWarnings

`func (o *MetricDescriptor) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *MetricDescriptor) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *MetricDescriptor) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *MetricDescriptor) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


