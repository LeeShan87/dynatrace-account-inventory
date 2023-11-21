# RemediationProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AffectedEntities** | Pointer to **[]string** | A list of related entities that are affected by the security problem. | [optional] [readonly] 
**UnaffectedEntities** | Pointer to **[]string** | A list of related entities that are affected by the security problem. | [optional] [readonly] 

## Methods

### NewRemediationProgress

`func NewRemediationProgress() *RemediationProgress`

NewRemediationProgress instantiates a new RemediationProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediationProgressWithDefaults

`func NewRemediationProgressWithDefaults() *RemediationProgress`

NewRemediationProgressWithDefaults instantiates a new RemediationProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffectedEntities

`func (o *RemediationProgress) GetAffectedEntities() []string`

GetAffectedEntities returns the AffectedEntities field if non-nil, zero value otherwise.

### GetAffectedEntitiesOk

`func (o *RemediationProgress) GetAffectedEntitiesOk() (*[]string, bool)`

GetAffectedEntitiesOk returns a tuple with the AffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedEntities

`func (o *RemediationProgress) SetAffectedEntities(v []string)`

SetAffectedEntities sets AffectedEntities field to given value.

### HasAffectedEntities

`func (o *RemediationProgress) HasAffectedEntities() bool`

HasAffectedEntities returns a boolean if a field has been set.

### GetUnaffectedEntities

`func (o *RemediationProgress) GetUnaffectedEntities() []string`

GetUnaffectedEntities returns the UnaffectedEntities field if non-nil, zero value otherwise.

### GetUnaffectedEntitiesOk

`func (o *RemediationProgress) GetUnaffectedEntitiesOk() (*[]string, bool)`

GetUnaffectedEntitiesOk returns a tuple with the UnaffectedEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnaffectedEntities

`func (o *RemediationProgress) SetUnaffectedEntities(v []string)`

SetUnaffectedEntities sets UnaffectedEntities field to given value.

### HasUnaffectedEntities

`func (o *RemediationProgress) HasUnaffectedEntities() bool`

HasUnaffectedEntities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


