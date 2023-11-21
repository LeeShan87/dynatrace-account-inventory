# WebApplicationMetricDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The type of the web application metric. | 
**PropertyKey** | Pointer to **string** | The key of the user action property.    Only applicable for &#x60;DoubleProperty&#x60; and &#x60;LongProperty&#x60; metrics. | [optional] 

## Methods

### NewWebApplicationMetricDefinition

`func NewWebApplicationMetricDefinition(metric string, ) *WebApplicationMetricDefinition`

NewWebApplicationMetricDefinition instantiates a new WebApplicationMetricDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebApplicationMetricDefinitionWithDefaults

`func NewWebApplicationMetricDefinitionWithDefaults() *WebApplicationMetricDefinition`

NewWebApplicationMetricDefinitionWithDefaults instantiates a new WebApplicationMetricDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *WebApplicationMetricDefinition) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *WebApplicationMetricDefinition) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *WebApplicationMetricDefinition) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetPropertyKey

`func (o *WebApplicationMetricDefinition) GetPropertyKey() string`

GetPropertyKey returns the PropertyKey field if non-nil, zero value otherwise.

### GetPropertyKeyOk

`func (o *WebApplicationMetricDefinition) GetPropertyKeyOk() (*string, bool)`

GetPropertyKeyOk returns a tuple with the PropertyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyKey

`func (o *WebApplicationMetricDefinition) SetPropertyKey(v string)`

SetPropertyKey sets PropertyKey field to given value.

### HasPropertyKey

`func (o *WebApplicationMetricDefinition) HasPropertyKey() bool`

HasPropertyKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


