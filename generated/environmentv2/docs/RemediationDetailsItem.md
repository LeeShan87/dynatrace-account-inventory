# RemediationDetailsItem

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
**VulnerableComponents** | Pointer to [**[]RemediationItemDetailsVulnerableComponent**](RemediationItemDetailsVulnerableComponent.md) | A list of vulnerable components of the remediation item.   A vulnerable component is what causes the security problem. | [optional] [readonly] 

## Methods

### NewRemediationDetailsItem

`func NewRemediationDetailsItem() *RemediationDetailsItem`

NewRemediationDetailsItem instantiates a new RemediationDetailsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationDetailsItemWithDefaults

`func NewRemediationDetailsItemWithDefaults() *RemediationDetailsItem`

NewRemediationDetailsItemWithDefaults instantiates a new RemediationDetailsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssessment

`func (o *RemediationDetailsItem) GetAssessment() RemediationAssessment`

GetAssessment returns the Assessment field if non-nil, zero value otherwise.

### GetAssessmentOk

`func (o *RemediationDetailsItem) GetAssessmentOk() (*RemediationAssessment, bool)`

GetAssessmentOk returns a tuple with the Assessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessment

`func (o *RemediationDetailsItem) SetAssessment(v RemediationAssessment)`

SetAssessment sets Assessment field to given value.

### HasAssessment

`func (o *RemediationDetailsItem) HasAssessment() bool`

HasAssessment returns a boolean if a field has been set.

### GetEntityIds

`func (o *RemediationDetailsItem) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *RemediationDetailsItem) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *RemediationDetailsItem) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *RemediationDetailsItem) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### GetFirstAffectedTimestamp

`func (o *RemediationDetailsItem) GetFirstAffectedTimestamp() int64`

GetFirstAffectedTimestamp returns the FirstAffectedTimestamp field if non-nil, zero value otherwise.

### GetFirstAffectedTimestampOk

`func (o *RemediationDetailsItem) GetFirstAffectedTimestampOk() (*int64, bool)`

GetFirstAffectedTimestampOk returns a tuple with the FirstAffectedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAffectedTimestamp

`func (o *RemediationDetailsItem) SetFirstAffectedTimestamp(v int64)`

SetFirstAffectedTimestamp sets FirstAffectedTimestamp field to given value.

### HasFirstAffectedTimestamp

`func (o *RemediationDetailsItem) HasFirstAffectedTimestamp() bool`

HasFirstAffectedTimestamp returns a boolean if a field has been set.

### GetId

`func (o *RemediationDetailsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemediationDetailsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemediationDetailsItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemediationDetailsItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMuteState

`func (o *RemediationDetailsItem) GetMuteState() RemediationItemMuteState`

GetMuteState returns the MuteState field if non-nil, zero value otherwise.

### GetMuteStateOk

`func (o *RemediationDetailsItem) GetMuteStateOk() (*RemediationItemMuteState, bool)`

GetMuteStateOk returns a tuple with the MuteState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteState

`func (o *RemediationDetailsItem) SetMuteState(v RemediationItemMuteState)`

SetMuteState sets MuteState field to given value.

### HasMuteState

`func (o *RemediationDetailsItem) HasMuteState() bool`

HasMuteState returns a boolean if a field has been set.

### GetName

`func (o *RemediationDetailsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemediationDetailsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemediationDetailsItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemediationDetailsItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRemediationProgress

`func (o *RemediationDetailsItem) GetRemediationProgress() RemediationProgress`

GetRemediationProgress returns the RemediationProgress field if non-nil, zero value otherwise.

### GetRemediationProgressOk

`func (o *RemediationDetailsItem) GetRemediationProgressOk() (*RemediationProgress, bool)`

GetRemediationProgressOk returns a tuple with the RemediationProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemediationProgress

`func (o *RemediationDetailsItem) SetRemediationProgress(v RemediationProgress)`

SetRemediationProgress sets RemediationProgress field to given value.

### HasRemediationProgress

`func (o *RemediationDetailsItem) HasRemediationProgress() bool`

HasRemediationProgress returns a boolean if a field has been set.

### GetResolvedTimestamp

`func (o *RemediationDetailsItem) GetResolvedTimestamp() int64`

GetResolvedTimestamp returns the ResolvedTimestamp field if non-nil, zero value otherwise.

### GetResolvedTimestampOk

`func (o *RemediationDetailsItem) GetResolvedTimestampOk() (*int64, bool)`

GetResolvedTimestampOk returns a tuple with the ResolvedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedTimestamp

`func (o *RemediationDetailsItem) SetResolvedTimestamp(v int64)`

SetResolvedTimestamp sets ResolvedTimestamp field to given value.

### HasResolvedTimestamp

`func (o *RemediationDetailsItem) HasResolvedTimestamp() bool`

HasResolvedTimestamp returns a boolean if a field has been set.

### GetVulnerabilityState

`func (o *RemediationDetailsItem) GetVulnerabilityState() string`

GetVulnerabilityState returns the VulnerabilityState field if non-nil, zero value otherwise.

### GetVulnerabilityStateOk

`func (o *RemediationDetailsItem) GetVulnerabilityStateOk() (*string, bool)`

GetVulnerabilityStateOk returns a tuple with the VulnerabilityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerabilityState

`func (o *RemediationDetailsItem) SetVulnerabilityState(v string)`

SetVulnerabilityState sets VulnerabilityState field to given value.

### HasVulnerabilityState

`func (o *RemediationDetailsItem) HasVulnerabilityState() bool`

HasVulnerabilityState returns a boolean if a field has been set.

### GetVulnerableComponents

`func (o *RemediationDetailsItem) GetVulnerableComponents() []RemediationItemDetailsVulnerableComponent`

GetVulnerableComponents returns the VulnerableComponents field if non-nil, zero value otherwise.

### GetVulnerableComponentsOk

`func (o *RemediationDetailsItem) GetVulnerableComponentsOk() (*[]RemediationItemDetailsVulnerableComponent, bool)`

GetVulnerableComponentsOk returns a tuple with the VulnerableComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableComponents

`func (o *RemediationDetailsItem) SetVulnerableComponents(v []RemediationItemDetailsVulnerableComponent)`

SetVulnerableComponents sets VulnerableComponents field to given value.

### HasVulnerableComponents

`func (o *RemediationDetailsItem) HasVulnerableComponents() bool`

HasVulnerableComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


