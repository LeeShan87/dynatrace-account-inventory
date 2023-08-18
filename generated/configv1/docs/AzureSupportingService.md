# AzureSupportingService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitoredMetrics** | Pointer to [**[]AzureMonitoredMetric**](AzureMonitoredMetric.md) | A list of metrics to be monitored for this service. It must include all the recommended metrics. If the list is null then recommended list of metrics for this service will be monitored. | [optional] 
**Name** | **string** | The name of the service. Valid supported service names can be discovered using /azure/supportedServices restAPI | 

## Methods

### NewAzureSupportingService

`func NewAzureSupportingService(name string, ) *AzureSupportingService`

NewAzureSupportingService instantiates a new AzureSupportingService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSupportingServiceWithDefaults

`func NewAzureSupportingServiceWithDefaults() *AzureSupportingService`

NewAzureSupportingServiceWithDefaults instantiates a new AzureSupportingService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitoredMetrics

`func (o *AzureSupportingService) GetMonitoredMetrics() []AzureMonitoredMetric`

GetMonitoredMetrics returns the MonitoredMetrics field if non-nil, zero value otherwise.

### GetMonitoredMetricsOk

`func (o *AzureSupportingService) GetMonitoredMetricsOk() (*[]AzureMonitoredMetric, bool)`

GetMonitoredMetricsOk returns a tuple with the MonitoredMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoredMetrics

`func (o *AzureSupportingService) SetMonitoredMetrics(v []AzureMonitoredMetric)`

SetMonitoredMetrics sets MonitoredMetrics field to given value.

### HasMonitoredMetrics

`func (o *AzureSupportingService) HasMonitoredMetrics() bool`

HasMonitoredMetrics returns a boolean if a field has been set.

### GetName

`func (o *AzureSupportingService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureSupportingService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureSupportingService) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


