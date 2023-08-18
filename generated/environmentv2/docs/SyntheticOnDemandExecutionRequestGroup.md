# SyntheticOnDemandExecutionRequestGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to **[]string** | List of application identifiers. Only monitors with all applications assigned will be executed. | [optional] 
**Locations** | Pointer to **[]string** | The locations from where monitors are to be executed. | [optional] 
**Services** | Pointer to **[]string** | List of service identifiers. Only monitors with all services assigned will be executed. | [optional] 
**Tags** | Pointer to **[]string** | List of tags. Only monitors with all tags assigned will be executed. | [optional] 

## Methods

### NewSyntheticOnDemandExecutionRequestGroup

`func NewSyntheticOnDemandExecutionRequestGroup() *SyntheticOnDemandExecutionRequestGroup`

NewSyntheticOnDemandExecutionRequestGroup instantiates a new SyntheticOnDemandExecutionRequestGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticOnDemandExecutionRequestGroupWithDefaults

`func NewSyntheticOnDemandExecutionRequestGroupWithDefaults() *SyntheticOnDemandExecutionRequestGroup`

NewSyntheticOnDemandExecutionRequestGroupWithDefaults instantiates a new SyntheticOnDemandExecutionRequestGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *SyntheticOnDemandExecutionRequestGroup) GetApplications() []string`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *SyntheticOnDemandExecutionRequestGroup) GetApplicationsOk() (*[]string, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *SyntheticOnDemandExecutionRequestGroup) SetApplications(v []string)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *SyntheticOnDemandExecutionRequestGroup) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetLocations

`func (o *SyntheticOnDemandExecutionRequestGroup) GetLocations() []string`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *SyntheticOnDemandExecutionRequestGroup) GetLocationsOk() (*[]string, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *SyntheticOnDemandExecutionRequestGroup) SetLocations(v []string)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *SyntheticOnDemandExecutionRequestGroup) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetServices

`func (o *SyntheticOnDemandExecutionRequestGroup) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *SyntheticOnDemandExecutionRequestGroup) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *SyntheticOnDemandExecutionRequestGroup) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *SyntheticOnDemandExecutionRequestGroup) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetTags

`func (o *SyntheticOnDemandExecutionRequestGroup) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SyntheticOnDemandExecutionRequestGroup) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SyntheticOnDemandExecutionRequestGroup) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SyntheticOnDemandExecutionRequestGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


