# RemoteConfigurationManagementValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvalidEntities** | Pointer to [**[]RemoteConfigurationManagementEntityValidationError**](RemoteConfigurationManagementEntityValidationError.md) | A list of validation errors for entities. | [optional] 
**InvalidOperations** | Pointer to [**[]RemoteConfigurationManagementOperationValidationError**](RemoteConfigurationManagementOperationValidationError.md) | A list of validation errors for operations. | [optional] 

## Methods

### NewRemoteConfigurationManagementValidationResult

`func NewRemoteConfigurationManagementValidationResult() *RemoteConfigurationManagementValidationResult`

NewRemoteConfigurationManagementValidationResult instantiates a new RemoteConfigurationManagementValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConfigurationManagementValidationResultWithDefaults

`func NewRemoteConfigurationManagementValidationResultWithDefaults() *RemoteConfigurationManagementValidationResult`

NewRemoteConfigurationManagementValidationResultWithDefaults instantiates a new RemoteConfigurationManagementValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvalidEntities

`func (o *RemoteConfigurationManagementValidationResult) GetInvalidEntities() []RemoteConfigurationManagementEntityValidationError`

GetInvalidEntities returns the InvalidEntities field if non-nil, zero value otherwise.

### GetInvalidEntitiesOk

`func (o *RemoteConfigurationManagementValidationResult) GetInvalidEntitiesOk() (*[]RemoteConfigurationManagementEntityValidationError, bool)`

GetInvalidEntitiesOk returns a tuple with the InvalidEntities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidEntities

`func (o *RemoteConfigurationManagementValidationResult) SetInvalidEntities(v []RemoteConfigurationManagementEntityValidationError)`

SetInvalidEntities sets InvalidEntities field to given value.

### HasInvalidEntities

`func (o *RemoteConfigurationManagementValidationResult) HasInvalidEntities() bool`

HasInvalidEntities returns a boolean if a field has been set.

### GetInvalidOperations

`func (o *RemoteConfigurationManagementValidationResult) GetInvalidOperations() []RemoteConfigurationManagementOperationValidationError`

GetInvalidOperations returns the InvalidOperations field if non-nil, zero value otherwise.

### GetInvalidOperationsOk

`func (o *RemoteConfigurationManagementValidationResult) GetInvalidOperationsOk() (*[]RemoteConfigurationManagementOperationValidationError, bool)`

GetInvalidOperationsOk returns a tuple with the InvalidOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidOperations

`func (o *RemoteConfigurationManagementValidationResult) SetInvalidOperations(v []RemoteConfigurationManagementOperationValidationError)`

SetInvalidOperations sets InvalidOperations field to given value.

### HasInvalidOperations

`func (o *RemoteConfigurationManagementValidationResult) HasInvalidOperations() bool`

HasInvalidOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


