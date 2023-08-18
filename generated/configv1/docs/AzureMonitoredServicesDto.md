# AzureMonitoredServicesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Services** | [**[]AzureSupportingService**](AzureSupportingService.md) | A list of Azure services to be monitored. Available services are listed by [/azure/supportedServices](https://dt-url.net/wt42sdq) operation.  For each service, a list of metrics and dimensions can be specified. A list of supported metrics and dimensions for a given service can be checked in [documentation](https://dt-url.net/kx2351b).  List of metrics can be skipped (set to null), resulting in recommended (default) set of metrics and dimensions being chosen for monitoring. For built-in services, adjusting the list of metrics is not supported, therefore it needs to be null. | 

## Methods

### NewAzureMonitoredServicesDto

`func NewAzureMonitoredServicesDto(services []AzureSupportingService, ) *AzureMonitoredServicesDto`

NewAzureMonitoredServicesDto instantiates a new AzureMonitoredServicesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureMonitoredServicesDtoWithDefaults

`func NewAzureMonitoredServicesDtoWithDefaults() *AzureMonitoredServicesDto`

NewAzureMonitoredServicesDtoWithDefaults instantiates a new AzureMonitoredServicesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AzureMonitoredServicesDto) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AzureMonitoredServicesDto) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AzureMonitoredServicesDto) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AzureMonitoredServicesDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetServices

`func (o *AzureMonitoredServicesDto) GetServices() []AzureSupportingService`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AzureMonitoredServicesDto) GetServicesOk() (*[]AzureSupportingService, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AzureMonitoredServicesDto) SetServices(v []AzureSupportingService)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


