# RemediationProgressEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assessment** | Pointer to [**RemediationProgressEntityAssessment**](RemediationProgressEntityAssessment.md) |  | [optional] 
**FirstAffectedTimestamp** | Pointer to **int64** | The timestamp when the remediation progress entity has first been related to the vulnerability. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the remediation progress entity. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the remediation progress entity. | [optional] [readonly] 
**State** | Pointer to **string** | The current state of the remediation progress entity. | [optional] [readonly] 
**VulnerableComponents** | Pointer to [**[]RemediationProgressVulnerableComponent**](RemediationProgressVulnerableComponent.md) | A list of vulnerable components of the remediation item.   A vulnerable component is what causes the security problem. | [optional] [readonly] 

## Methods

### NewRemediationProgressEntity

`func NewRemediationProgressEntity() *RemediationProgressEntity`

NewRemediationProgressEntity instantiates a new RemediationProgressEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationProgressEntityWithDefaults

`func NewRemediationProgressEntityWithDefaults() *RemediationProgressEntity`

NewRemediationProgressEntityWithDefaults instantiates a new RemediationProgressEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssessment

`func (o *RemediationProgressEntity) GetAssessment() RemediationProgressEntityAssessment`

GetAssessment returns the Assessment field if non-nil, zero value otherwise.

### GetAssessmentOk

`func (o *RemediationProgressEntity) GetAssessmentOk() (*RemediationProgressEntityAssessment, bool)`

GetAssessmentOk returns a tuple with the Assessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessment

`func (o *RemediationProgressEntity) SetAssessment(v RemediationProgressEntityAssessment)`

SetAssessment sets Assessment field to given value.

### HasAssessment

`func (o *RemediationProgressEntity) HasAssessment() bool`

HasAssessment returns a boolean if a field has been set.

### GetFirstAffectedTimestamp

`func (o *RemediationProgressEntity) GetFirstAffectedTimestamp() int64`

GetFirstAffectedTimestamp returns the FirstAffectedTimestamp field if non-nil, zero value otherwise.

### GetFirstAffectedTimestampOk

`func (o *RemediationProgressEntity) GetFirstAffectedTimestampOk() (*int64, bool)`

GetFirstAffectedTimestampOk returns a tuple with the FirstAffectedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAffectedTimestamp

`func (o *RemediationProgressEntity) SetFirstAffectedTimestamp(v int64)`

SetFirstAffectedTimestamp sets FirstAffectedTimestamp field to given value.

### HasFirstAffectedTimestamp

`func (o *RemediationProgressEntity) HasFirstAffectedTimestamp() bool`

HasFirstAffectedTimestamp returns a boolean if a field has been set.

### GetId

`func (o *RemediationProgressEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemediationProgressEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemediationProgressEntity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RemediationProgressEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RemediationProgressEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RemediationProgressEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RemediationProgressEntity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RemediationProgressEntity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *RemediationProgressEntity) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RemediationProgressEntity) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RemediationProgressEntity) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RemediationProgressEntity) HasState() bool`

HasState returns a boolean if a field has been set.

### GetVulnerableComponents

`func (o *RemediationProgressEntity) GetVulnerableComponents() []RemediationProgressVulnerableComponent`

GetVulnerableComponents returns the VulnerableComponents field if non-nil, zero value otherwise.

### GetVulnerableComponentsOk

`func (o *RemediationProgressEntity) GetVulnerableComponentsOk() (*[]RemediationProgressVulnerableComponent, bool)`

GetVulnerableComponentsOk returns a tuple with the VulnerableComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVulnerableComponents

`func (o *RemediationProgressEntity) SetVulnerableComponents(v []RemediationProgressVulnerableComponent)`

SetVulnerableComponents sets VulnerableComponents field to given value.

### HasVulnerableComponents

`func (o *RemediationProgressEntity) HasVulnerableComponents() bool`

HasVulnerableComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


