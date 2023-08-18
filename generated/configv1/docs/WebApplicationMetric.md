# WebApplicationMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationIdentifier** | **string** | The Dynatrace entity ID of the application to which the metric belongs. | 
**Dimensions** | Pointer to [**[]WebApplicationDimensionDefinition**](WebApplicationDimensionDefinition.md) | A list of metric dimensions. | [optional] 
**Enabled** | **bool** | The metric is enabled (&#x60;true&#x60;) or disabled (&#x60;false&#x60;). | 
**MetricDefinition** | [**WebApplicationMetricDefinition**](WebApplicationMetricDefinition.md) |  | 
**MetricKey** | **string** | The unique key of the metric.    The key must have the &#x60;calc:apps&#x60; prefix. | 
**Name** | **string** | The displayed name of the metric. | 
**UserActionFilter** | Pointer to [**UserActionFilter**](UserActionFilter.md) |  | [optional] 

## Methods

### NewWebApplicationMetric

`func NewWebApplicationMetric(applicationIdentifier string, enabled bool, metricDefinition WebApplicationMetricDefinition, metricKey string, name string, ) *WebApplicationMetric`

NewWebApplicationMetric instantiates a new WebApplicationMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebApplicationMetricWithDefaults

`func NewWebApplicationMetricWithDefaults() *WebApplicationMetric`

NewWebApplicationMetricWithDefaults instantiates a new WebApplicationMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationIdentifier

`func (o *WebApplicationMetric) GetApplicationIdentifier() string`

GetApplicationIdentifier returns the ApplicationIdentifier field if non-nil, zero value otherwise.

### GetApplicationIdentifierOk

`func (o *WebApplicationMetric) GetApplicationIdentifierOk() (*string, bool)`

GetApplicationIdentifierOk returns a tuple with the ApplicationIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIdentifier

`func (o *WebApplicationMetric) SetApplicationIdentifier(v string)`

SetApplicationIdentifier sets ApplicationIdentifier field to given value.


### GetDimensions

`func (o *WebApplicationMetric) GetDimensions() []WebApplicationDimensionDefinition`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *WebApplicationMetric) GetDimensionsOk() (*[]WebApplicationDimensionDefinition, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *WebApplicationMetric) SetDimensions(v []WebApplicationDimensionDefinition)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *WebApplicationMetric) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetEnabled

`func (o *WebApplicationMetric) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WebApplicationMetric) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WebApplicationMetric) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMetricDefinition

`func (o *WebApplicationMetric) GetMetricDefinition() WebApplicationMetricDefinition`

GetMetricDefinition returns the MetricDefinition field if non-nil, zero value otherwise.

### GetMetricDefinitionOk

`func (o *WebApplicationMetric) GetMetricDefinitionOk() (*WebApplicationMetricDefinition, bool)`

GetMetricDefinitionOk returns a tuple with the MetricDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDefinition

`func (o *WebApplicationMetric) SetMetricDefinition(v WebApplicationMetricDefinition)`

SetMetricDefinition sets MetricDefinition field to given value.


### GetMetricKey

`func (o *WebApplicationMetric) GetMetricKey() string`

GetMetricKey returns the MetricKey field if non-nil, zero value otherwise.

### GetMetricKeyOk

`func (o *WebApplicationMetric) GetMetricKeyOk() (*string, bool)`

GetMetricKeyOk returns a tuple with the MetricKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricKey

`func (o *WebApplicationMetric) SetMetricKey(v string)`

SetMetricKey sets MetricKey field to given value.


### GetName

`func (o *WebApplicationMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebApplicationMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebApplicationMetric) SetName(v string)`

SetName sets Name field to given value.


### GetUserActionFilter

`func (o *WebApplicationMetric) GetUserActionFilter() UserActionFilter`

GetUserActionFilter returns the UserActionFilter field if non-nil, zero value otherwise.

### GetUserActionFilterOk

`func (o *WebApplicationMetric) GetUserActionFilterOk() (*UserActionFilter, bool)`

GetUserActionFilterOk returns a tuple with the UserActionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionFilter

`func (o *WebApplicationMetric) SetUserActionFilter(v UserActionFilter)`

SetUserActionFilter sets UserActionFilter field to given value.

### HasUserActionFilter

`func (o *WebApplicationMetric) HasUserActionFilter() bool`

HasUserActionFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


