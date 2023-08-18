# FilteredCountsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedNodes** | Pointer to **int32** | Number of affected nodes | [optional] [readonly] 
**AffectedProcessGroupInstances** | Pointer to **int32** | Number of affected processes | [optional] [readonly] 
**AffectedProcessGroups** | Pointer to **int32** | Number of affected process groups | [optional] [readonly] 
**ExposedProcessGroups** | Pointer to **int32** | Number of exposed process groups | [optional] [readonly] 
**ReachableDataAssets** | Pointer to **int32** | Number of reachable data assets | [optional] [readonly] 
**RelatedApplications** | Pointer to **int32** | Number of related applications | [optional] [readonly] 
**RelatedAttacks** | Pointer to **int32** | Number of related attacks | [optional] [readonly] 
**RelatedDatabases** | Pointer to **int32** | Number of related databases | [optional] [readonly] 
**RelatedHosts** | Pointer to **int32** | Number of related hosts | [optional] [readonly] 
**RelatedKubernetesClusters** | Pointer to **int32** | Number of related Kubernetes clusters | [optional] [readonly] 
**RelatedKubernetesWorkloads** | Pointer to **int32** | Number of related Kubernetes workloads | [optional] [readonly] 
**RelatedServices** | Pointer to **int32** | Number of related services | [optional] [readonly] 
**VulnerableComponents** | Pointer to **int32** | Number of vulnerable components | [optional] [readonly] 

## Methods

### NewFilteredCountsDto

`func NewFilteredCountsDto() *FilteredCountsDto`

NewFilteredCountsDto instantiates a new FilteredCountsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilteredCountsDtoWithDefaults

`func NewFilteredCountsDtoWithDefaults() *FilteredCountsDto`

NewFilteredCountsDtoWithDefaults instantiates a new FilteredCountsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedNodes

`func (o *FilteredCountsDto) GetAffectedNodes() int32`

GetAffectedNodes returns the AffectedNodes field if non-nil, zero value otherwise.

### GetAffectedNodesOk

`func (o *FilteredCountsDto) GetAffectedNodesOk() (*int32, bool)`

GetAffectedNodesOk returns a tuple with the AffectedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedNodes

`func (o *FilteredCountsDto) SetAffectedNodes(v int32)`

SetAffectedNodes sets AffectedNodes field to given value.

### HasAffectedNodes

`func (o *FilteredCountsDto) HasAffectedNodes() bool`

HasAffectedNodes returns a boolean if a field has been set.

### GetAffectedProcessGroupInstances

`func (o *FilteredCountsDto) GetAffectedProcessGroupInstances() int32`

GetAffectedProcessGroupInstances returns the AffectedProcessGroupInstances field if non-nil, zero value otherwise.

### GetAffectedProcessGroupInstancesOk

`func (o *FilteredCountsDto) GetAffectedProcessGroupInstancesOk() (*int32, bool)`

GetAffectedProcessGroupInstancesOk returns a tuple with the AffectedProcessGroupInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProcessGroupInstances

`func (o *FilteredCountsDto) SetAffectedProcessGroupInstances(v int32)`

SetAffectedProcessGroupInstances sets AffectedProcessGroupInstances field to given value.

### HasAffectedProcessGroupInstances

`func (o *FilteredCountsDto) HasAffectedProcessGroupInstances() bool`

HasAffectedProcessGroupInstances returns a boolean if a field has been set.

### GetAffectedProcessGroups

`func (o *FilteredCountsDto) GetAffectedProcessGroups() int32`

GetAffectedProcessGroups returns the AffectedProcessGroups field if non-nil, zero value otherwise.

### GetAffectedProcessGroupsOk

`func (o *FilteredCountsDto) GetAffectedProcessGroupsOk() (*int32, bool)`

GetAffectedProcessGroupsOk returns a tuple with the AffectedProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProcessGroups

`func (o *FilteredCountsDto) SetAffectedProcessGroups(v int32)`

SetAffectedProcessGroups sets AffectedProcessGroups field to given value.

### HasAffectedProcessGroups

`func (o *FilteredCountsDto) HasAffectedProcessGroups() bool`

HasAffectedProcessGroups returns a boolean if a field has been set.

### GetExposedProcessGroups

`func (o *FilteredCountsDto) GetExposedProcessGroups() int32`

GetExposedProcessGroups returns the ExposedProcessGroups field if non-nil, zero value otherwise.

### GetExposedProcessGroupsOk

`func (o *FilteredCountsDto) GetExposedProcessGroupsOk() (*int32, bool)`

GetExposedProcessGroupsOk returns a tuple with the ExposedProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedProcessGroups

`func (o *FilteredCountsDto) SetExposedProcessGroups(v int32)`

SetExposedProcessGroups sets ExposedProcessGroups field to given value.

### HasExposedProcessGroups

`func (o *FilteredCountsDto) HasExposedProcessGroups() bool`

HasExposedProcessGroups returns a boolean if a field has been set.

### GetReachableDataAssets

`func (o *FilteredCountsDto) GetReachableDataAssets() int32`

GetReachableDataAssets returns the ReachableDataAssets field if non-nil, zero value otherwise.

### GetReachableDataAssetsOk

`func (o *FilteredCountsDto) GetReachableDataAssetsOk() (*int32, bool)`

GetReachableDataAssetsOk returns a tuple with the ReachableDataAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableDataAssets

`func (o *FilteredCountsDto) SetReachableDataAssets(v int32)`

SetReachableDataAssets sets ReachableDataAssets field to given value.

### HasReachableDataAssets

`func (o *FilteredCountsDto) HasReachableDataAssets() bool`

HasReachableDataAssets returns a boolean if a field has been set.

### GetRelatedApplications

`func (o *FilteredCountsDto) GetRelatedApplications() int32`

GetRelatedApplications returns the RelatedApplications field if non-nil, zero value otherwise.

### GetRelatedApplicationsOk

`func (o *FilteredCountsDto) GetRelatedApplicationsOk() (*int32, bool)`

GetRelatedApplicationsOk returns a tuple with the RelatedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedApplications

`func (o *FilteredCountsDto) SetRelatedApplications(v int32)`

SetRelatedApplications sets RelatedApplications field to given value.

### HasRelatedApplications

`func (o *FilteredCountsDto) HasRelatedApplications() bool`

HasRelatedApplications returns a boolean if a field has been set.

### GetRelatedAttacks

`func (o *FilteredCountsDto) GetRelatedAttacks() int32`

GetRelatedAttacks returns the RelatedAttacks field if non-nil, zero value otherwise.

### GetRelatedAttacksOk

`func (o *FilteredCountsDto) GetRelatedAttacksOk() (*int32, bool)`

GetRelatedAttacksOk returns a tuple with the RelatedAttacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedAttacks

`func (o *FilteredCountsDto) SetRelatedAttacks(v int32)`

SetRelatedAttacks sets RelatedAttacks field to given value.

### HasRelatedAttacks

`func (o *FilteredCountsDto) HasRelatedAttacks() bool`

HasRelatedAttacks returns a boolean if a field has been set.

### GetRelatedDatabases

`func (o *FilteredCountsDto) GetRelatedDatabases() int32`

GetRelatedDatabases returns the RelatedDatabases field if non-nil, zero value otherwise.

### GetRelatedDatabasesOk

`func (o *FilteredCountsDto) GetRelatedDatabasesOk() (*int32, bool)`

GetRelatedDatabasesOk returns a tuple with the RelatedDatabases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedDatabases

`func (o *FilteredCountsDto) SetRelatedDatabases(v int32)`

SetRelatedDatabases sets RelatedDatabases field to given value.

### HasRelatedDatabases

`func (o *FilteredCountsDto) HasRelatedDatabases() bool`

HasRelatedDatabases returns a boolean if a field has been set.

### GetRelatedHosts

`func (o *FilteredCountsDto) GetRelatedHosts() int32`

GetRelatedHosts returns the RelatedHosts field if non-nil, zero value otherwise.

### GetRelatedHostsOk

`func (o *FilteredCountsDto) GetRelatedHostsOk() (*int32, bool)`

GetRelatedHostsOk returns a tuple with the RelatedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedHosts

`func (o *FilteredCountsDto) SetRelatedHosts(v int32)`

SetRelatedHosts sets RelatedHosts field to given value.

### HasRelatedHosts

`func (o *FilteredCountsDto) HasRelatedHosts() bool`

HasRelatedHosts returns a boolean if a field has been set.

### GetRelatedKubernetesClusters

`func (o *FilteredCountsDto) GetRelatedKubernetesClusters() int32`

GetRelatedKubernetesClusters returns the RelatedKubernetesClusters field if non-nil, zero value otherwise.

### GetRelatedKubernetesClustersOk

`func (o *FilteredCountsDto) GetRelatedKubernetesClustersOk() (*int32, bool)`

GetRelatedKubernetesClustersOk returns a tuple with the RelatedKubernetesClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedKubernetesClusters

`func (o *FilteredCountsDto) SetRelatedKubernetesClusters(v int32)`

SetRelatedKubernetesClusters sets RelatedKubernetesClusters field to given value.

### HasRelatedKubernetesClusters

`func (o *FilteredCountsDto) HasRelatedKubernetesClusters() bool`

HasRelatedKubernetesClusters returns a boolean if a field has been set.

### GetRelatedKubernetesWorkloads

`func (o *FilteredCountsDto) GetRelatedKubernetesWorkloads() int32`

GetRelatedKubernetesWorkloads returns the RelatedKubernetesWorkloads field if non-nil, zero value otherwise.

### GetRelatedKubernetesWorkloadsOk

`func (o *FilteredCountsDto) GetRelatedKubernetesWorkloadsOk() (*int32, bool)`

GetRelatedKubernetesWorkloadsOk returns a tuple with the RelatedKubernetesWorkloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedKubernetesWorkloads

`func (o *FilteredCountsDto) SetRelatedKubernetesWorkloads(v int32)`

SetRelatedKubernetesWorkloads sets RelatedKubernetesWorkloads field to given value.

### HasRelatedKubernetesWorkloads

`func (o *FilteredCountsDto) HasRelatedKubernetesWorkloads() bool`

HasRelatedKubernetesWorkloads returns a boolean if a field has been set.

### GetRelatedServices

`func (o *FilteredCountsDto) GetRelatedServices() int32`

GetRelatedServices returns the RelatedServices field if non-nil, zero value otherwise.

### GetRelatedServicesOk

`func (o *FilteredCountsDto) GetRelatedServicesOk() (*int32, bool)`

GetRelatedServicesOk returns a tuple with the RelatedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedServices

`func (o *FilteredCountsDto) SetRelatedServices(v int32)`

SetRelatedServices sets RelatedServices field to given value.

### HasRelatedServices

`func (o *FilteredCountsDto) HasRelatedServices() bool`

HasRelatedServices returns a boolean if a field has been set.

### GetVulnerableComponents

`func (o *FilteredCountsDto) GetVulnerableComponents() int32`

GetVulnerableComponents returns the VulnerableComponents field if non-nil, zero value otherwise.

### GetVulnerableComponentsOk

`func (o *FilteredCountsDto) GetVulnerableComponentsOk() (*int32, bool)`

GetVulnerableComponentsOk returns a tuple with the VulnerableComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableComponents

`func (o *FilteredCountsDto) SetVulnerableComponents(v int32)`

SetVulnerableComponents sets VulnerableComponents field to given value.

### HasVulnerableComponents

`func (o *FilteredCountsDto) HasVulnerableComponents() bool`

HasVulnerableComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


