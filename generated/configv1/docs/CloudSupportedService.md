# CloudSupportedService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloudProviderServiceType** | **string** | Name of service used by cloud provider. | 
**DisplayName** | **string** | Display name for service on Dynatrace UI | 
**EntityType** | **string** | Entity type monitored by this service | 
**Name** | **string** | Service unique name used by Dynatrace.  | 

## Methods

### NewCloudSupportedService

`func NewCloudSupportedService(cloudProviderServiceType string, displayName string, entityType string, name string, ) *CloudSupportedService`

NewCloudSupportedService instantiates a new CloudSupportedService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSupportedServiceWithDefaults

`func NewCloudSupportedServiceWithDefaults() *CloudSupportedService`

NewCloudSupportedServiceWithDefaults instantiates a new CloudSupportedService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloudProviderServiceType

`func (o *CloudSupportedService) GetCloudProviderServiceType() string`

GetCloudProviderServiceType returns the CloudProviderServiceType field if non-nil, zero value otherwise.

### GetCloudProviderServiceTypeOk

`func (o *CloudSupportedService) GetCloudProviderServiceTypeOk() (*string, bool)`

GetCloudProviderServiceTypeOk returns a tuple with the CloudProviderServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderServiceType

`func (o *CloudSupportedService) SetCloudProviderServiceType(v string)`

SetCloudProviderServiceType sets CloudProviderServiceType field to given value.


### GetDisplayName

`func (o *CloudSupportedService) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudSupportedService) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudSupportedService) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEntityType

`func (o *CloudSupportedService) GetEntityType() string`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CloudSupportedService) GetEntityTypeOk() (*string, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CloudSupportedService) SetEntityType(v string)`

SetEntityType sets EntityType field to given value.


### GetName

`func (o *CloudSupportedService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudSupportedService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudSupportedService) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


