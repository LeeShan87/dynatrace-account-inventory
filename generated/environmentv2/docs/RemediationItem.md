# RemediationItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assessment** | Pointer to [**RemediationAssessment**](RemediationAssessment.md) |  | [optional] 
**EntityIds** | Pointer to **[]string** |  | [optional] 
**FirstAffectedTimestamp** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MuteState** | Pointer to [**RemediationItemMuteState**](RemediationItemMuteState.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RemediationProgress** | Pointer to [**RemediationProgress**](RemediationProgress.md) |  | [optional] 
**ResolvedTimestamp** | Pointer to **int64** |  | [optional] 
**VulnerabilityState** | Pointer to **string** |  | [optional] 
**VulnerableComponents** | Pointer to [**[]VulnerableComponent**](VulnerableComponent.md) | A list of vulnerable components of the remediation item.   A vulnerable component is what causes the security problem. | [optional] [readonly] 

## Methods

### NewRemediationItem

`func NewRemediationItem() *RemediationItem`

NewRemediationItem instantiates a new RemediationItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationItemWithDefaults

`func NewRemediationItemWithDefaults() *RemediationItem`

NewRemediationItemWithDefaults instantiates a new RemediationItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssessment

`func (o *RemediationItem) GetAssessment() RemediationAssessment`

GetAssessment returns the Assessment field if non-nil, zero value otherwise.

### GetAssessmentOk

`func (o *RemediationItem) GetAssessmentOk() (*RemediationAssessment, bool)`

GetAssessmentOk returns a tuple with the Assessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessment

`func (o *RemediationItem) SetAssessment(v RemediationAssessment)`

SetAssessment sets Assessment field to given value.

### HasAssessment

`func (o *RemediationItem) HasAssessment() bool`

HasAssessment returns a boolean if a field has been set.

### GetEntityIds

`func (o *RemediationItem) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *RemediationItem) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *RemediationItem) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *RemediationItem) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### GetFirstAffectedTimestamp

`func (o *RemediationItem) GetFirstAffectedTimestamp() int64`

GetFirstAffectedTimestamp returns the FirstAffectedTimestamp field if non-nil, zero value otherwise.

### GetFirstAffectedTimestampOk

`func (o *RemediationItem) GetFirstAffectedTimestampOk() (*int64, bool)`

GetFirstAffectedTimestampOk returns a tuple with the FirstAffectedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAffectedTimestamp

`func (o *RemediationItem) SetFirstAffectedTimestamp(v int64)`

SetFirstAffectedTimestamp sets FirstAffectedTimestamp field to given value.

### HasFirstAffectedTimestamp

`func (o *RemediationItem) HasFirstAffectedTimestamp() bool`

HasFirstAffectedTimestamp returns a boolean if a field has been set.

### GetId

`func (o *RemediationItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemediationItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemediationItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemediationItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMuteState

`func (o *RemediationItem) GetMuteState() RemediationItemMuteState`

GetMuteState returns the MuteState field if non-nil, zero value otherwise.

### GetMuteStateOk

`func (o *RemediationItem) GetMuteStateOk() (*RemediationItemMuteState, bool)`

GetMuteStateOk returns a tuple with the MuteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteState

`func (o *RemediationItem) SetMuteState(v RemediationItemMuteState)`

SetMuteState sets MuteState field to given value.

### HasMuteState

`func (o *RemediationItem) HasMuteState() bool`

HasMuteState returns a boolean if a field has been set.

### GetName

`func (o *RemediationItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemediationItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemediationItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemediationItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemediationProgress

`func (o *RemediationItem) GetRemediationProgress() RemediationProgress`

GetRemediationProgress returns the RemediationProgress field if non-nil, zero value otherwise.

### GetRemediationProgressOk

`func (o *RemediationItem) GetRemediationProgressOk() (*RemediationProgress, bool)`

GetRemediationProgressOk returns a tuple with the RemediationProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationProgress

`func (o *RemediationItem) SetRemediationProgress(v RemediationProgress)`

SetRemediationProgress sets RemediationProgress field to given value.

### HasRemediationProgress

`func (o *RemediationItem) HasRemediationProgress() bool`

HasRemediationProgress returns a boolean if a field has been set.

### GetResolvedTimestamp

`func (o *RemediationItem) GetResolvedTimestamp() int64`

GetResolvedTimestamp returns the ResolvedTimestamp field if non-nil, zero value otherwise.

### GetResolvedTimestampOk

`func (o *RemediationItem) GetResolvedTimestampOk() (*int64, bool)`

GetResolvedTimestampOk returns a tuple with the ResolvedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedTimestamp

`func (o *RemediationItem) SetResolvedTimestamp(v int64)`

SetResolvedTimestamp sets ResolvedTimestamp field to given value.

### HasResolvedTimestamp

`func (o *RemediationItem) HasResolvedTimestamp() bool`

HasResolvedTimestamp returns a boolean if a field has been set.

### GetVulnerabilityState

`func (o *RemediationItem) GetVulnerabilityState() string`

GetVulnerabilityState returns the VulnerabilityState field if non-nil, zero value otherwise.

### GetVulnerabilityStateOk

`func (o *RemediationItem) GetVulnerabilityStateOk() (*string, bool)`

GetVulnerabilityStateOk returns a tuple with the VulnerabilityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityState

`func (o *RemediationItem) SetVulnerabilityState(v string)`

SetVulnerabilityState sets VulnerabilityState field to given value.

### HasVulnerabilityState

`func (o *RemediationItem) HasVulnerabilityState() bool`

HasVulnerabilityState returns a boolean if a field has been set.

### GetVulnerableComponents

`func (o *RemediationItem) GetVulnerableComponents() []VulnerableComponent`

GetVulnerableComponents returns the VulnerableComponents field if non-nil, zero value otherwise.

### GetVulnerableComponentsOk

`func (o *RemediationItem) GetVulnerableComponentsOk() (*[]VulnerableComponent, bool)`

GetVulnerableComponentsOk returns a tuple with the VulnerableComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableComponents

`func (o *RemediationItem) SetVulnerableComponents(v []VulnerableComponent)`

SetVulnerableComponents sets VulnerableComponents field to given value.

### HasVulnerableComponents

`func (o *RemediationItem) HasVulnerableComponents() bool`

HasVulnerableComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


