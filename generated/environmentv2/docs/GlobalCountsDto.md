# GlobalCountsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedNodes** | Pointer to **int32** | Number of affected nodes | [optional] [readonly] 
**AffectedProcessGroupInstances** | Pointer to **int32** | Number of affected process group instances | [optional] [readonly] 
**AffectedProcessGroups** | Pointer to **int32** | Number of affected process groups | [optional] [readonly] 
**ExposedProcessGroups** | Pointer to **int32** | Number of exposed process groups | [optional] [readonly] 
**ReachableDataAssets** | Pointer to **int32** | Number of reachable data assets exposed | [optional] [readonly] 
**RelatedApplications** | Pointer to **int32** | Number of related applications | [optional] [readonly] 
**RelatedAttacks** | Pointer to **int32** | Number of attacks on the exposed security problem | [optional] [readonly] 
**RelatedHosts** | Pointer to **int32** | Number of related hosts | [optional] [readonly] 
**RelatedKubernetesClusters** | Pointer to **int32** | Number of related kubernetes cluster | [optional] [readonly] 
**RelatedKubernetesWorkloads** | Pointer to **int32** | Number of related kubernetes workloads | [optional] [readonly] 
**RelatedServices** | Pointer to **int32** | Number of related services | [optional] [readonly] 
**VulnerableComponents** | Pointer to **int32** | Number of vulnerable components | [optional] [readonly] 

## Methods

### NewGlobalCountsDto

`func NewGlobalCountsDto() *GlobalCountsDto`

NewGlobalCountsDto instantiates a new GlobalCountsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalCountsDtoWithDefaults

`func NewGlobalCountsDtoWithDefaults() *GlobalCountsDto`

NewGlobalCountsDtoWithDefaults instantiates a new GlobalCountsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedNodes

`func (o *GlobalCountsDto) GetAffectedNodes() int32`

GetAffectedNodes returns the AffectedNodes field if non-nil, zero value otherwise.

### GetAffectedNodesOk

`func (o *GlobalCountsDto) GetAffectedNodesOk() (*int32, bool)`

GetAffectedNodesOk returns a tuple with the AffectedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedNodes

`func (o *GlobalCountsDto) SetAffectedNodes(v int32)`

SetAffectedNodes sets AffectedNodes field to given value.

### HasAffectedNodes

`func (o *GlobalCountsDto) HasAffectedNodes() bool`

HasAffectedNodes returns a boolean if a field has been set.

### GetAffectedProcessGroupInstances

`func (o *GlobalCountsDto) GetAffectedProcessGroupInstances() int32`

GetAffectedProcessGroupInstances returns the AffectedProcessGroupInstances field if non-nil, zero value otherwise.

### GetAffectedProcessGroupInstancesOk

`func (o *GlobalCountsDto) GetAffectedProcessGroupInstancesOk() (*int32, bool)`

GetAffectedProcessGroupInstancesOk returns a tuple with the AffectedProcessGroupInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProcessGroupInstances

`func (o *GlobalCountsDto) SetAffectedProcessGroupInstances(v int32)`

SetAffectedProcessGroupInstances sets AffectedProcessGroupInstances field to given value.

### HasAffectedProcessGroupInstances

`func (o *GlobalCountsDto) HasAffectedProcessGroupInstances() bool`

HasAffectedProcessGroupInstances returns a boolean if a field has been set.

### GetAffectedProcessGroups

`func (o *GlobalCountsDto) GetAffectedProcessGroups() int32`

GetAffectedProcessGroups returns the AffectedProcessGroups field if non-nil, zero value otherwise.

### GetAffectedProcessGroupsOk

`func (o *GlobalCountsDto) GetAffectedProcessGroupsOk() (*int32, bool)`

GetAffectedProcessGroupsOk returns a tuple with the AffectedProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedProcessGroups

`func (o *GlobalCountsDto) SetAffectedProcessGroups(v int32)`

SetAffectedProcessGroups sets AffectedProcessGroups field to given value.

### HasAffectedProcessGroups

`func (o *GlobalCountsDto) HasAffectedProcessGroups() bool`

HasAffectedProcessGroups returns a boolean if a field has been set.

### GetExposedProcessGroups

`func (o *GlobalCountsDto) GetExposedProcessGroups() int32`

GetExposedProcessGroups returns the ExposedProcessGroups field if non-nil, zero value otherwise.

### GetExposedProcessGroupsOk

`func (o *GlobalCountsDto) GetExposedProcessGroupsOk() (*int32, bool)`

GetExposedProcessGroupsOk returns a tuple with the ExposedProcessGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedProcessGroups

`func (o *GlobalCountsDto) SetExposedProcessGroups(v int32)`

SetExposedProcessGroups sets ExposedProcessGroups field to given value.

### HasExposedProcessGroups

`func (o *GlobalCountsDto) HasExposedProcessGroups() bool`

HasExposedProcessGroups returns a boolean if a field has been set.

### GetReachableDataAssets

`func (o *GlobalCountsDto) GetReachableDataAssets() int32`

GetReachableDataAssets returns the ReachableDataAssets field if non-nil, zero value otherwise.

### GetReachableDataAssetsOk

`func (o *GlobalCountsDto) GetReachableDataAssetsOk() (*int32, bool)`

GetReachableDataAssetsOk returns a tuple with the ReachableDataAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableDataAssets

`func (o *GlobalCountsDto) SetReachableDataAssets(v int32)`

SetReachableDataAssets sets ReachableDataAssets field to given value.

### HasReachableDataAssets

`func (o *GlobalCountsDto) HasReachableDataAssets() bool`

HasReachableDataAssets returns a boolean if a field has been set.

### GetRelatedApplications

`func (o *GlobalCountsDto) GetRelatedApplications() int32`

GetRelatedApplications returns the RelatedApplications field if non-nil, zero value otherwise.

### GetRelatedApplicationsOk

`func (o *GlobalCountsDto) GetRelatedApplicationsOk() (*int32, bool)`

GetRelatedApplicationsOk returns a tuple with the RelatedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedApplications

`func (o *GlobalCountsDto) SetRelatedApplications(v int32)`

SetRelatedApplications sets RelatedApplications field to given value.

### HasRelatedApplications

`func (o *GlobalCountsDto) HasRelatedApplications() bool`

HasRelatedApplications returns a boolean if a field has been set.

### GetRelatedAttacks

`func (o *GlobalCountsDto) GetRelatedAttacks() int32`

GetRelatedAttacks returns the RelatedAttacks field if non-nil, zero value otherwise.

### GetRelatedAttacksOk

`func (o *GlobalCountsDto) GetRelatedAttacksOk() (*int32, bool)`

GetRelatedAttacksOk returns a tuple with the RelatedAttacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedAttacks

`func (o *GlobalCountsDto) SetRelatedAttacks(v int32)`

SetRelatedAttacks sets RelatedAttacks field to given value.

### HasRelatedAttacks

`func (o *GlobalCountsDto) HasRelatedAttacks() bool`

HasRelatedAttacks returns a boolean if a field has been set.

### GetRelatedHosts

`func (o *GlobalCountsDto) GetRelatedHosts() int32`

GetRelatedHosts returns the RelatedHosts field if non-nil, zero value otherwise.

### GetRelatedHostsOk

`func (o *GlobalCountsDto) GetRelatedHostsOk() (*int32, bool)`

GetRelatedHostsOk returns a tuple with the RelatedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedHosts

`func (o *GlobalCountsDto) SetRelatedHosts(v int32)`

SetRelatedHosts sets RelatedHosts field to given value.

### HasRelatedHosts

`func (o *GlobalCountsDto) HasRelatedHosts() bool`

HasRelatedHosts returns a boolean if a field has been set.

### GetRelatedKubernetesClusters

`func (o *GlobalCountsDto) GetRelatedKubernetesClusters() int32`

GetRelatedKubernetesClusters returns the RelatedKubernetesClusters field if non-nil, zero value otherwise.

### GetRelatedKubernetesClustersOk

`func (o *GlobalCountsDto) GetRelatedKubernetesClustersOk() (*int32, bool)`

GetRelatedKubernetesClustersOk returns a tuple with the RelatedKubernetesClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedKubernetesClusters

`func (o *GlobalCountsDto) SetRelatedKubernetesClusters(v int32)`

SetRelatedKubernetesClusters sets RelatedKubernetesClusters field to given value.

### HasRelatedKubernetesClusters

`func (o *GlobalCountsDto) HasRelatedKubernetesClusters() bool`

HasRelatedKubernetesClusters returns a boolean if a field has been set.

### GetRelatedKubernetesWorkloads

`func (o *GlobalCountsDto) GetRelatedKubernetesWorkloads() int32`

GetRelatedKubernetesWorkloads returns the RelatedKubernetesWorkloads field if non-nil, zero value otherwise.

### GetRelatedKubernetesWorkloadsOk

`func (o *GlobalCountsDto) GetRelatedKubernetesWorkloadsOk() (*int32, bool)`

GetRelatedKubernetesWorkloadsOk returns a tuple with the RelatedKubernetesWorkloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedKubernetesWorkloads

`func (o *GlobalCountsDto) SetRelatedKubernetesWorkloads(v int32)`

SetRelatedKubernetesWorkloads sets RelatedKubernetesWorkloads field to given value.

### HasRelatedKubernetesWorkloads

`func (o *GlobalCountsDto) HasRelatedKubernetesWorkloads() bool`

HasRelatedKubernetesWorkloads returns a boolean if a field has been set.

### GetRelatedServices

`func (o *GlobalCountsDto) GetRelatedServices() int32`

GetRelatedServices returns the RelatedServices field if non-nil, zero value otherwise.

### GetRelatedServicesOk

`func (o *GlobalCountsDto) GetRelatedServicesOk() (*int32, bool)`

GetRelatedServicesOk returns a tuple with the RelatedServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedServices

`func (o *GlobalCountsDto) SetRelatedServices(v int32)`

SetRelatedServices sets RelatedServices field to given value.

### HasRelatedServices

`func (o *GlobalCountsDto) HasRelatedServices() bool`

HasRelatedServices returns a boolean if a field has been set.

### GetVulnerableComponents

`func (o *GlobalCountsDto) GetVulnerableComponents() int32`

GetVulnerableComponents returns the VulnerableComponents field if non-nil, zero value otherwise.

### GetVulnerableComponentsOk

`func (o *GlobalCountsDto) GetVulnerableComponentsOk() (*int32, bool)`

GetVulnerableComponentsOk returns a tuple with the VulnerableComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableComponents

`func (o *GlobalCountsDto) SetVulnerableComponents(v int32)`

SetVulnerableComponents sets VulnerableComponents field to given value.

### HasVulnerableComponents

`func (o *GlobalCountsDto) HasVulnerableComponents() bool`

HasVulnerableComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


