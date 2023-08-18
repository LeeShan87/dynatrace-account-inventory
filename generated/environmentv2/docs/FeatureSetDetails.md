# FeatureSetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**[]MetricDto**](MetricDto.md) | Feature set metrics | [optional] 

## Methods

### NewFeatureSetDetails

`func NewFeatureSetDetails() *FeatureSetDetails`

NewFeatureSetDetails instantiates a new FeatureSetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureSetDetailsWithDefaults

`func NewFeatureSetDetailsWithDefaults() *FeatureSetDetails`

NewFeatureSetDetailsWithDefaults instantiates a new FeatureSetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *FeatureSetDetails) GetMetrics() []MetricDto`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *FeatureSetDetails) GetMetricsOk() (*[]MetricDto, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *FeatureSetDetails) SetMetrics(v []MetricDto)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *FeatureSetDetails) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


