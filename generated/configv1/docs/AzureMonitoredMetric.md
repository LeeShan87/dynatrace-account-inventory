# AzureMonitoredMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | **[]string** | A list of metric&#39;s dimensions names. It must include all the recommended dimensions. | 
**Name** | **string** | The name of the metric of the service. | 

## Methods

### NewAzureMonitoredMetric

`func NewAzureMonitoredMetric(dimensions []string, name string, ) *AzureMonitoredMetric`

NewAzureMonitoredMetric instantiates a new AzureMonitoredMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMonitoredMetricWithDefaults

`func NewAzureMonitoredMetricWithDefaults() *AzureMonitoredMetric`

NewAzureMonitoredMetricWithDefaults instantiates a new AzureMonitoredMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *AzureMonitoredMetric) GetDimensions() []string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *AzureMonitoredMetric) GetDimensionsOk() (*[]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *AzureMonitoredMetric) SetDimensions(v []string)`

SetDimensions sets Dimensions field to given value.


### GetName

`func (o *AzureMonitoredMetric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureMonitoredMetric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureMonitoredMetric) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


