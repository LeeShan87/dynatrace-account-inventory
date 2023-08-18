# AwsMonitoredServicesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ConfigurationMetadata**](ConfigurationMetadata.md) |  | [optional] 
**Services** | [**[]AwsSupportingServiceConfig**](AwsSupportingServiceConfig.md) | A list of AWS services to be monitored. Available services are listed by [/aws/supportedServices](https://dt-url.net/me02sh2) operation.  For each service, a list of metrics and dimensions can be specified. A list of supported metrics and dimensions for a given service can be checked in [documentation](https://dt-url.net/r12v0pkl).  List of metrics can be skipped (set to null), resulting in recommended (default) set of metrics and dimensions being chosen for monitoring. For built-in services, adjusting the list of metrics is not supported, therefore it needs to be null. | 

## Methods

### NewAwsMonitoredServicesDto

`func NewAwsMonitoredServicesDto(services []AwsSupportingServiceConfig, ) *AwsMonitoredServicesDto`

NewAwsMonitoredServicesDto instantiates a new AwsMonitoredServicesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsMonitoredServicesDtoWithDefaults

`func NewAwsMonitoredServicesDtoWithDefaults() *AwsMonitoredServicesDto`

NewAwsMonitoredServicesDtoWithDefaults instantiates a new AwsMonitoredServicesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *AwsMonitoredServicesDto) GetMetadata() ConfigurationMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AwsMonitoredServicesDto) GetMetadataOk() (*ConfigurationMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AwsMonitoredServicesDto) SetMetadata(v ConfigurationMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AwsMonitoredServicesDto) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetServices

`func (o *AwsMonitoredServicesDto) GetServices() []AwsSupportingServiceConfig`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *AwsMonitoredServicesDto) GetServicesOk() (*[]AwsSupportingServiceConfig, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *AwsMonitoredServicesDto) SetServices(v []AwsSupportingServiceConfig)`

SetServices sets Services field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


