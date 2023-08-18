# SecurityProblemDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedEntities** | Pointer to **[]string** | A list of affected entities of the security problem.   An affected entity is an entity where a vulnerable component runs. | [optional] [readonly] 
**CodeLevelVulnerabilityDetails** | Pointer to [**CodeLevelVulnerabilityDetails**](CodeLevelVulnerabilityDetails.md) |  | [optional] 
**CveIds** | Pointer to **[]string** | A list of CVE IDs of the security problem. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the security problem. | [optional] [readonly] 
**DisplayId** | Pointer to **string** | The display ID of the security problem. | [optional] [readonly] 
**EntryPoints** | Pointer to [**EntryPoints**](EntryPoints.md) |  | [optional] 
**Events** | Pointer to [**[]SecurityProblemEvent**](SecurityProblemEvent.md) | An ordered (newest first) list of events of the security problem. | [optional] 
**ExposedEntities** | Pointer to **[]string** | A list of exposed entities of the security problem.   An exposed entity is an affected entity that is exposed to the internet. | [optional] [readonly] 
**ExternalVulnerabilityId** | Pointer to **string** | The external vulnerability ID of the security problem. | [optional] [readonly] 
**FilteredCounts** | Pointer to [**FilteredCountsDto**](FilteredCountsDto.md) |  | [optional] 
**FirstSeenTimestamp** | Pointer to **int64** | The timestamp of the first occurrence of the security problem. | [optional] [readonly] 
**GlobalCounts** | Pointer to [**GlobalCountsDto**](GlobalCountsDto.md) |  | [optional] 
**LastOpenedTimestamp** | Pointer to **int64** | The timestamp when the security problem was last opened. | [optional] [readonly] 
**LastResolvedTimestamp** | Pointer to **int64** | The timestamp when the security problem was last resolved. | [optional] [readonly] 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp of the most recent security problem change. | [optional] [readonly] 
**ManagementZones** | Pointer to [**[]ManagementZone**](ManagementZone.md) | A list of management zones which the affected entities belong to. | [optional] [readonly] 
**MuteStateChangeInProgress** | Pointer to **bool** | If &#x60;true&#x60; a change of the mute state is in progress. | [optional] [readonly] 
**Muted** | Pointer to **bool** | The security problem is (&#x60;true&#x60;) or is not (&#x60;false&#x60;) muted. | [optional] [readonly] 
**PackageName** | Pointer to **string** | The package name of the security problem. | [optional] [readonly] 
**ReachableDataAssets** | Pointer to **[]string** | A list of data assets reachable by affected entities of the security problem.   A data asset is a service that has database access. | [optional] [readonly] 
**RelatedAttacks** | Pointer to [**RelatedAttacksList**](RelatedAttacksList.md) |  | [optional] 
**RelatedContainerImages** | Pointer to [**[]SecurityProblemDetailsRelatedContainerImagesInner**](SecurityProblemDetailsRelatedContainerImagesInner.md) | A list of related container images of the security problem.   A related container image is a container image that contains at least one *affected entity*. | [optional] [readonly] 
**RelatedEntities** | Pointer to [**RelatedEntitiesList**](RelatedEntitiesList.md) |  | [optional] 
**RiskAssessment** | Pointer to [**RiskAssessment**](RiskAssessment.md) |  | [optional] 
**SecurityProblemId** | Pointer to **string** | The ID of the security problem. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the security problem. | [optional] [readonly] 
**Technology** | Pointer to **string** | The technology of the security problem. | [optional] [readonly] 
**Title** | Pointer to **string** | The title of the security problem. | [optional] [readonly] 
**Url** | Pointer to **string** | The URL to the security problem details page. | [optional] [readonly] 
**VulnerabilityType** | Pointer to **string** | The type of the vulnerability. | [optional] [readonly] 
**VulnerableComponents** | Pointer to [**[]VulnerableComponent**](VulnerableComponent.md) | A list of vulnerable components of the security problem.   A vulnerable component is what causes the security problem. | [optional] [readonly] 

## Methods

### NewSecurityProblemDetails

`func NewSecurityProblemDetails() *SecurityProblemDetails`

NewSecurityProblemDetails instantiates a new SecurityProblemDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityProblemDetailsWithDefaults

`func NewSecurityProblemDetailsWithDefaults() *SecurityProblemDetails`

NewSecurityProblemDetailsWithDefaults instantiates a new SecurityProblemDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedEntities

`func (o *SecurityProblemDetails) GetAffectedEntities() []string`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *SecurityProblemDetails) GetAffectedEntitiesOk() (*[]string, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *SecurityProblemDetails) SetAffectedEntities(v []string)`

SetAffectedEntities sets AffectedEntities field to given value.

### HasAffectedEntities

`func (o *SecurityProblemDetails) HasAffectedEntities() bool`

HasAffectedEntities returns a boolean if a field has been set.

### GetCodeLevelVulnerabilityDetails

`func (o *SecurityProblemDetails) GetCodeLevelVulnerabilityDetails() CodeLevelVulnerabilityDetails`

GetCodeLevelVulnerabilityDetails returns the CodeLevelVulnerabilityDetails field if non-nil, zero value otherwise.

### GetCodeLevelVulnerabilityDetailsOk

`func (o *SecurityProblemDetails) GetCodeLevelVulnerabilityDetailsOk() (*CodeLevelVulnerabilityDetails, bool)`

GetCodeLevelVulnerabilityDetailsOk returns a tuple with the CodeLevelVulnerabilityDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeLevelVulnerabilityDetails

`func (o *SecurityProblemDetails) SetCodeLevelVulnerabilityDetails(v CodeLevelVulnerabilityDetails)`

SetCodeLevelVulnerabilityDetails sets CodeLevelVulnerabilityDetails field to given value.

### HasCodeLevelVulnerabilityDetails

`func (o *SecurityProblemDetails) HasCodeLevelVulnerabilityDetails() bool`

HasCodeLevelVulnerabilityDetails returns a boolean if a field has been set.

### GetCveIds

`func (o *SecurityProblemDetails) GetCveIds() []string`

GetCveIds returns the CveIds field if non-nil, zero value otherwise.

### GetCveIdsOk

`func (o *SecurityProblemDetails) GetCveIdsOk() (*[]string, bool)`

GetCveIdsOk returns a tuple with the CveIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCveIds

`func (o *SecurityProblemDetails) SetCveIds(v []string)`

SetCveIds sets CveIds field to given value.

### HasCveIds

`func (o *SecurityProblemDetails) HasCveIds() bool`

HasCveIds returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityProblemDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityProblemDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityProblemDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityProblemDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayId

`func (o *SecurityProblemDetails) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *SecurityProblemDetails) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *SecurityProblemDetails) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *SecurityProblemDetails) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetEntryPoints

`func (o *SecurityProblemDetails) GetEntryPoints() EntryPoints`

GetEntryPoints returns the EntryPoints field if non-nil, zero value otherwise.

### GetEntryPointsOk

`func (o *SecurityProblemDetails) GetEntryPointsOk() (*EntryPoints, bool)`

GetEntryPointsOk returns a tuple with the EntryPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryPoints

`func (o *SecurityProblemDetails) SetEntryPoints(v EntryPoints)`

SetEntryPoints sets EntryPoints field to given value.

### HasEntryPoints

`func (o *SecurityProblemDetails) HasEntryPoints() bool`

HasEntryPoints returns a boolean if a field has been set.

### GetEvents

`func (o *SecurityProblemDetails) GetEvents() []SecurityProblemEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SecurityProblemDetails) GetEventsOk() (*[]SecurityProblemEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SecurityProblemDetails) SetEvents(v []SecurityProblemEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *SecurityProblemDetails) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetExposedEntities

`func (o *SecurityProblemDetails) GetExposedEntities() []string`

GetExposedEntities returns the ExposedEntities field if non-nil, zero value otherwise.

### GetExposedEntitiesOk

`func (o *SecurityProblemDetails) GetExposedEntitiesOk() (*[]string, bool)`

GetExposedEntitiesOk returns a tuple with the ExposedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedEntities

`func (o *SecurityProblemDetails) SetExposedEntities(v []string)`

SetExposedEntities sets ExposedEntities field to given value.

### HasExposedEntities

`func (o *SecurityProblemDetails) HasExposedEntities() bool`

HasExposedEntities returns a boolean if a field has been set.

### GetExternalVulnerabilityId

`func (o *SecurityProblemDetails) GetExternalVulnerabilityId() string`

GetExternalVulnerabilityId returns the ExternalVulnerabilityId field if non-nil, zero value otherwise.

### GetExternalVulnerabilityIdOk

`func (o *SecurityProblemDetails) GetExternalVulnerabilityIdOk() (*string, bool)`

GetExternalVulnerabilityIdOk returns a tuple with the ExternalVulnerabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVulnerabilityId

`func (o *SecurityProblemDetails) SetExternalVulnerabilityId(v string)`

SetExternalVulnerabilityId sets ExternalVulnerabilityId field to given value.

### HasExternalVulnerabilityId

`func (o *SecurityProblemDetails) HasExternalVulnerabilityId() bool`

HasExternalVulnerabilityId returns a boolean if a field has been set.

### GetFilteredCounts

`func (o *SecurityProblemDetails) GetFilteredCounts() FilteredCountsDto`

GetFilteredCounts returns the FilteredCounts field if non-nil, zero value otherwise.

### GetFilteredCountsOk

`func (o *SecurityProblemDetails) GetFilteredCountsOk() (*FilteredCountsDto, bool)`

GetFilteredCountsOk returns a tuple with the FilteredCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilteredCounts

`func (o *SecurityProblemDetails) SetFilteredCounts(v FilteredCountsDto)`

SetFilteredCounts sets FilteredCounts field to given value.

### HasFilteredCounts

`func (o *SecurityProblemDetails) HasFilteredCounts() bool`

HasFilteredCounts returns a boolean if a field has been set.

### GetFirstSeenTimestamp

`func (o *SecurityProblemDetails) GetFirstSeenTimestamp() int64`

GetFirstSeenTimestamp returns the FirstSeenTimestamp field if non-nil, zero value otherwise.

### GetFirstSeenTimestampOk

`func (o *SecurityProblemDetails) GetFirstSeenTimestampOk() (*int64, bool)`

GetFirstSeenTimestampOk returns a tuple with the FirstSeenTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeenTimestamp

`func (o *SecurityProblemDetails) SetFirstSeenTimestamp(v int64)`

SetFirstSeenTimestamp sets FirstSeenTimestamp field to given value.

### HasFirstSeenTimestamp

`func (o *SecurityProblemDetails) HasFirstSeenTimestamp() bool`

HasFirstSeenTimestamp returns a boolean if a field has been set.

### GetGlobalCounts

`func (o *SecurityProblemDetails) GetGlobalCounts() GlobalCountsDto`

GetGlobalCounts returns the GlobalCounts field if non-nil, zero value otherwise.

### GetGlobalCountsOk

`func (o *SecurityProblemDetails) GetGlobalCountsOk() (*GlobalCountsDto, bool)`

GetGlobalCountsOk returns a tuple with the GlobalCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalCounts

`func (o *SecurityProblemDetails) SetGlobalCounts(v GlobalCountsDto)`

SetGlobalCounts sets GlobalCounts field to given value.

### HasGlobalCounts

`func (o *SecurityProblemDetails) HasGlobalCounts() bool`

HasGlobalCounts returns a boolean if a field has been set.

### GetLastOpenedTimestamp

`func (o *SecurityProblemDetails) GetLastOpenedTimestamp() int64`

GetLastOpenedTimestamp returns the LastOpenedTimestamp field if non-nil, zero value otherwise.

### GetLastOpenedTimestampOk

`func (o *SecurityProblemDetails) GetLastOpenedTimestampOk() (*int64, bool)`

GetLastOpenedTimestampOk returns a tuple with the LastOpenedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastOpenedTimestamp

`func (o *SecurityProblemDetails) SetLastOpenedTimestamp(v int64)`

SetLastOpenedTimestamp sets LastOpenedTimestamp field to given value.

### HasLastOpenedTimestamp

`func (o *SecurityProblemDetails) HasLastOpenedTimestamp() bool`

HasLastOpenedTimestamp returns a boolean if a field has been set.

### GetLastResolvedTimestamp

`func (o *SecurityProblemDetails) GetLastResolvedTimestamp() int64`

GetLastResolvedTimestamp returns the LastResolvedTimestamp field if non-nil, zero value otherwise.

### GetLastResolvedTimestampOk

`func (o *SecurityProblemDetails) GetLastResolvedTimestampOk() (*int64, bool)`

GetLastResolvedTimestampOk returns a tuple with the LastResolvedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResolvedTimestamp

`func (o *SecurityProblemDetails) SetLastResolvedTimestamp(v int64)`

SetLastResolvedTimestamp sets LastResolvedTimestamp field to given value.

### HasLastResolvedTimestamp

`func (o *SecurityProblemDetails) HasLastResolvedTimestamp() bool`

HasLastResolvedTimestamp returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *SecurityProblemDetails) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *SecurityProblemDetails) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *SecurityProblemDetails) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *SecurityProblemDetails) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetManagementZones

`func (o *SecurityProblemDetails) GetManagementZones() []ManagementZone`

GetManagementZones returns the ManagementZones field if non-nil, zero value otherwise.

### GetManagementZonesOk

`func (o *SecurityProblemDetails) GetManagementZonesOk() (*[]ManagementZone, bool)`

GetManagementZonesOk returns a tuple with the ManagementZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementZones

`func (o *SecurityProblemDetails) SetManagementZones(v []ManagementZone)`

SetManagementZones sets ManagementZones field to given value.

### HasManagementZones

`func (o *SecurityProblemDetails) HasManagementZones() bool`

HasManagementZones returns a boolean if a field has been set.

### GetMuteStateChangeInProgress

`func (o *SecurityProblemDetails) GetMuteStateChangeInProgress() bool`

GetMuteStateChangeInProgress returns the MuteStateChangeInProgress field if non-nil, zero value otherwise.

### GetMuteStateChangeInProgressOk

`func (o *SecurityProblemDetails) GetMuteStateChangeInProgressOk() (*bool, bool)`

GetMuteStateChangeInProgressOk returns a tuple with the MuteStateChangeInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteStateChangeInProgress

`func (o *SecurityProblemDetails) SetMuteStateChangeInProgress(v bool)`

SetMuteStateChangeInProgress sets MuteStateChangeInProgress field to given value.

### HasMuteStateChangeInProgress

`func (o *SecurityProblemDetails) HasMuteStateChangeInProgress() bool`

HasMuteStateChangeInProgress returns a boolean if a field has been set.

### GetMuted

`func (o *SecurityProblemDetails) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *SecurityProblemDetails) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *SecurityProblemDetails) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *SecurityProblemDetails) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetPackageName

`func (o *SecurityProblemDetails) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *SecurityProblemDetails) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *SecurityProblemDetails) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *SecurityProblemDetails) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetReachableDataAssets

`func (o *SecurityProblemDetails) GetReachableDataAssets() []string`

GetReachableDataAssets returns the ReachableDataAssets field if non-nil, zero value otherwise.

### GetReachableDataAssetsOk

`func (o *SecurityProblemDetails) GetReachableDataAssetsOk() (*[]string, bool)`

GetReachableDataAssetsOk returns a tuple with the ReachableDataAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReachableDataAssets

`func (o *SecurityProblemDetails) SetReachableDataAssets(v []string)`

SetReachableDataAssets sets ReachableDataAssets field to given value.

### HasReachableDataAssets

`func (o *SecurityProblemDetails) HasReachableDataAssets() bool`

HasReachableDataAssets returns a boolean if a field has been set.

### GetRelatedAttacks

`func (o *SecurityProblemDetails) GetRelatedAttacks() RelatedAttacksList`

GetRelatedAttacks returns the RelatedAttacks field if non-nil, zero value otherwise.

### GetRelatedAttacksOk

`func (o *SecurityProblemDetails) GetRelatedAttacksOk() (*RelatedAttacksList, bool)`

GetRelatedAttacksOk returns a tuple with the RelatedAttacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedAttacks

`func (o *SecurityProblemDetails) SetRelatedAttacks(v RelatedAttacksList)`

SetRelatedAttacks sets RelatedAttacks field to given value.

### HasRelatedAttacks

`func (o *SecurityProblemDetails) HasRelatedAttacks() bool`

HasRelatedAttacks returns a boolean if a field has been set.

### GetRelatedContainerImages

`func (o *SecurityProblemDetails) GetRelatedContainerImages() []SecurityProblemDetailsRelatedContainerImagesInner`

GetRelatedContainerImages returns the RelatedContainerImages field if non-nil, zero value otherwise.

### GetRelatedContainerImagesOk

`func (o *SecurityProblemDetails) GetRelatedContainerImagesOk() (*[]SecurityProblemDetailsRelatedContainerImagesInner, bool)`

GetRelatedContainerImagesOk returns a tuple with the RelatedContainerImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedContainerImages

`func (o *SecurityProblemDetails) SetRelatedContainerImages(v []SecurityProblemDetailsRelatedContainerImagesInner)`

SetRelatedContainerImages sets RelatedContainerImages field to given value.

### HasRelatedContainerImages

`func (o *SecurityProblemDetails) HasRelatedContainerImages() bool`

HasRelatedContainerImages returns a boolean if a field has been set.

### GetRelatedEntities

`func (o *SecurityProblemDetails) GetRelatedEntities() RelatedEntitiesList`

GetRelatedEntities returns the RelatedEntities field if non-nil, zero value otherwise.

### GetRelatedEntitiesOk

`func (o *SecurityProblemDetails) GetRelatedEntitiesOk() (*RelatedEntitiesList, bool)`

GetRelatedEntitiesOk returns a tuple with the RelatedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedEntities

`func (o *SecurityProblemDetails) SetRelatedEntities(v RelatedEntitiesList)`

SetRelatedEntities sets RelatedEntities field to given value.

### HasRelatedEntities

`func (o *SecurityProblemDetails) HasRelatedEntities() bool`

HasRelatedEntities returns a boolean if a field has been set.

### GetRiskAssessment

`func (o *SecurityProblemDetails) GetRiskAssessment() RiskAssessment`

GetRiskAssessment returns the RiskAssessment field if non-nil, zero value otherwise.

### GetRiskAssessmentOk

`func (o *SecurityProblemDetails) GetRiskAssessmentOk() (*RiskAssessment, bool)`

GetRiskAssessmentOk returns a tuple with the RiskAssessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAssessment

`func (o *SecurityProblemDetails) SetRiskAssessment(v RiskAssessment)`

SetRiskAssessment sets RiskAssessment field to given value.

### HasRiskAssessment

`func (o *SecurityProblemDetails) HasRiskAssessment() bool`

HasRiskAssessment returns a boolean if a field has been set.

### GetSecurityProblemId

`func (o *SecurityProblemDetails) GetSecurityProblemId() string`

GetSecurityProblemId returns the SecurityProblemId field if non-nil, zero value otherwise.

### GetSecurityProblemIdOk

`func (o *SecurityProblemDetails) GetSecurityProblemIdOk() (*string, bool)`

GetSecurityProblemIdOk returns a tuple with the SecurityProblemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityProblemId

`func (o *SecurityProblemDetails) SetSecurityProblemId(v string)`

SetSecurityProblemId sets SecurityProblemId field to given value.

### HasSecurityProblemId

`func (o *SecurityProblemDetails) HasSecurityProblemId() bool`

HasSecurityProblemId returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityProblemDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityProblemDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityProblemDetails) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecurityProblemDetails) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTechnology

`func (o *SecurityProblemDetails) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *SecurityProblemDetails) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *SecurityProblemDetails) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *SecurityProblemDetails) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

### GetTitle

`func (o *SecurityProblemDetails) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SecurityProblemDetails) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SecurityProblemDetails) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SecurityProblemDetails) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *SecurityProblemDetails) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SecurityProblemDetails) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SecurityProblemDetails) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SecurityProblemDetails) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVulnerabilityType

`func (o *SecurityProblemDetails) GetVulnerabilityType() string`

GetVulnerabilityType returns the VulnerabilityType field if non-nil, zero value otherwise.

### GetVulnerabilityTypeOk

`func (o *SecurityProblemDetails) GetVulnerabilityTypeOk() (*string, bool)`

GetVulnerabilityTypeOk returns a tuple with the VulnerabilityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityType

`func (o *SecurityProblemDetails) SetVulnerabilityType(v string)`

SetVulnerabilityType sets VulnerabilityType field to given value.

### HasVulnerabilityType

`func (o *SecurityProblemDetails) HasVulnerabilityType() bool`

HasVulnerabilityType returns a boolean if a field has been set.

### GetVulnerableComponents

`func (o *SecurityProblemDetails) GetVulnerableComponents() []VulnerableComponent`

GetVulnerableComponents returns the VulnerableComponents field if non-nil, zero value otherwise.

### GetVulnerableComponentsOk

`func (o *SecurityProblemDetails) GetVulnerableComponentsOk() (*[]VulnerableComponent, bool)`

GetVulnerableComponentsOk returns a tuple with the VulnerableComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableComponents

`func (o *SecurityProblemDetails) SetVulnerableComponents(v []VulnerableComponent)`

SetVulnerableComponents sets VulnerableComponents field to given value.

### HasVulnerableComponents

`func (o *SecurityProblemDetails) HasVulnerableComponents() bool`

HasVulnerableComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


