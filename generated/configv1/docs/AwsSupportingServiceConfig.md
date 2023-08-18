# AwsSupportingServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitoredMetrics** | Pointer to [**[]AwsSupportingServiceMetric**](AwsSupportingServiceMetric.md) | A list of metrics to be monitored for this service. If the list is null then recommended list of metrics for this service will be monitored. | [optional] 
**Name** | **string** | The name of the service. Valid supported service names can be discovered using /aws/supportedServices restAPI | 

## Methods

### NewAwsSupportingServiceConfig

`func NewAwsSupportingServiceConfig(name string, ) *AwsSupportingServiceConfig`

NewAwsSupportingServiceConfig instantiates a new AwsSupportingServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsSupportingServiceConfigWithDefaults

`func NewAwsSupportingServiceConfigWithDefaults() *AwsSupportingServiceConfig`

NewAwsSupportingServiceConfigWithDefaults instantiates a new AwsSupportingServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitoredMetrics

`func (o *AwsSupportingServiceConfig) GetMonitoredMetrics() []AwsSupportingServiceMetric`

GetMonitoredMetrics returns the MonitoredMetrics field if non-nil, zero value otherwise.

### GetMonitoredMetricsOk

`func (o *AwsSupportingServiceConfig) GetMonitoredMetricsOk() (*[]AwsSupportingServiceMetric, bool)`

GetMonitoredMetricsOk returns a tuple with the MonitoredMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredMetrics

`func (o *AwsSupportingServiceConfig) SetMonitoredMetrics(v []AwsSupportingServiceMetric)`

SetMonitoredMetrics sets MonitoredMetrics field to given value.

### HasMonitoredMetrics

`func (o *AwsSupportingServiceConfig) HasMonitoredMetrics() bool`

HasMonitoredMetrics returns a boolean if a field has been set.

### GetName

`func (o *AwsSupportingServiceConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AwsSupportingServiceConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AwsSupportingServiceConfig) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


