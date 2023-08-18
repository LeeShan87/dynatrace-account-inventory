# RiskAssessmentSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exposure** | Pointer to **string** | The level of exposure of affected entities. | [optional] [readonly] 
**NumberOfAffectedEntities** | Pointer to **int32** | The number of currently affected entities. | [optional] [readonly] 
**NumberOfReachableDataAssets** | Pointer to **int32** | The number of data assets that are currently reachable by affected entities. | [optional] [readonly] 
**PublicExploit** | Pointer to **string** | The availability status of public exploits. | [optional] [readonly] 
**RiskLevel** | Pointer to **string** | The Davis risk level.    It is calculated by Dynatrace on the basis of CVSS score. | [optional] [readonly] 
**RiskScore** | Pointer to **float32** | The Davis risk score (1-10).    It is calculated by Dynatrace on the basis of CVSS score. | [optional] [readonly] 
**VulnerableFunctionUsage** | Pointer to **string** | The state of vulnerable code execution. | [optional] [readonly] 

## Methods

### NewRiskAssessmentSnapshot

`func NewRiskAssessmentSnapshot() *RiskAssessmentSnapshot`

NewRiskAssessmentSnapshot instantiates a new RiskAssessmentSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAssessmentSnapshotWithDefaults

`func NewRiskAssessmentSnapshotWithDefaults() *RiskAssessmentSnapshot`

NewRiskAssessmentSnapshotWithDefaults instantiates a new RiskAssessmentSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExposure

`func (o *RiskAssessmentSnapshot) GetExposure() string`

GetExposure returns the Exposure field if non-nil, zero value otherwise.

### GetExposureOk

`func (o *RiskAssessmentSnapshot) GetExposureOk() (*string, bool)`

GetExposureOk returns a tuple with the Exposure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposure

`func (o *RiskAssessmentSnapshot) SetExposure(v string)`

SetExposure sets Exposure field to given value.

### HasExposure

`func (o *RiskAssessmentSnapshot) HasExposure() bool`

HasExposure returns a boolean if a field has been set.

### GetNumberOfAffectedEntities

`func (o *RiskAssessmentSnapshot) GetNumberOfAffectedEntities() int32`

GetNumberOfAffectedEntities returns the NumberOfAffectedEntities field if non-nil, zero value otherwise.

### GetNumberOfAffectedEntitiesOk

`func (o *RiskAssessmentSnapshot) GetNumberOfAffectedEntitiesOk() (*int32, bool)`

GetNumberOfAffectedEntitiesOk returns a tuple with the NumberOfAffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAffectedEntities

`func (o *RiskAssessmentSnapshot) SetNumberOfAffectedEntities(v int32)`

SetNumberOfAffectedEntities sets NumberOfAffectedEntities field to given value.

### HasNumberOfAffectedEntities

`func (o *RiskAssessmentSnapshot) HasNumberOfAffectedEntities() bool`

HasNumberOfAffectedEntities returns a boolean if a field has been set.

### GetNumberOfReachableDataAssets

`func (o *RiskAssessmentSnapshot) GetNumberOfReachableDataAssets() int32`

GetNumberOfReachableDataAssets returns the NumberOfReachableDataAssets field if non-nil, zero value otherwise.

### GetNumberOfReachableDataAssetsOk

`func (o *RiskAssessmentSnapshot) GetNumberOfReachableDataAssetsOk() (*int32, bool)`

GetNumberOfReachableDataAssetsOk returns a tuple with the NumberOfReachableDataAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfReachableDataAssets

`func (o *RiskAssessmentSnapshot) SetNumberOfReachableDataAssets(v int32)`

SetNumberOfReachableDataAssets sets NumberOfReachableDataAssets field to given value.

### HasNumberOfReachableDataAssets

`func (o *RiskAssessmentSnapshot) HasNumberOfReachableDataAssets() bool`

HasNumberOfReachableDataAssets returns a boolean if a field has been set.

### GetPublicExploit

`func (o *RiskAssessmentSnapshot) GetPublicExploit() string`

GetPublicExploit returns the PublicExploit field if non-nil, zero value otherwise.

### GetPublicExploitOk

`func (o *RiskAssessmentSnapshot) GetPublicExploitOk() (*string, bool)`

GetPublicExploitOk returns a tuple with the PublicExploit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicExploit

`func (o *RiskAssessmentSnapshot) SetPublicExploit(v string)`

SetPublicExploit sets PublicExploit field to given value.

### HasPublicExploit

`func (o *RiskAssessmentSnapshot) HasPublicExploit() bool`

HasPublicExploit returns a boolean if a field has been set.

### GetRiskLevel

`func (o *RiskAssessmentSnapshot) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *RiskAssessmentSnapshot) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *RiskAssessmentSnapshot) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *RiskAssessmentSnapshot) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### GetRiskScore

`func (o *RiskAssessmentSnapshot) GetRiskScore() float32`

GetRiskScore returns the RiskScore field if non-nil, zero value otherwise.

### GetRiskScoreOk

`func (o *RiskAssessmentSnapshot) GetRiskScoreOk() (*float32, bool)`

GetRiskScoreOk returns a tuple with the RiskScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskScore

`func (o *RiskAssessmentSnapshot) SetRiskScore(v float32)`

SetRiskScore sets RiskScore field to given value.

### HasRiskScore

`func (o *RiskAssessmentSnapshot) HasRiskScore() bool`

HasRiskScore returns a boolean if a field has been set.

### GetVulnerableFunctionUsage

`func (o *RiskAssessmentSnapshot) GetVulnerableFunctionUsage() string`

GetVulnerableFunctionUsage returns the VulnerableFunctionUsage field if non-nil, zero value otherwise.

### GetVulnerableFunctionUsageOk

`func (o *RiskAssessmentSnapshot) GetVulnerableFunctionUsageOk() (*string, bool)`

GetVulnerableFunctionUsageOk returns a tuple with the VulnerableFunctionUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableFunctionUsage

`func (o *RiskAssessmentSnapshot) SetVulnerableFunctionUsage(v string)`

SetVulnerableFunctionUsage sets VulnerableFunctionUsage field to given value.

### HasVulnerableFunctionUsage

`func (o *RiskAssessmentSnapshot) HasVulnerableFunctionUsage() bool`

HasVulnerableFunctionUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


