# RemoteConfigurationManagementOperationActiveGateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entities** | **[]string** | A list of entities IDs for which remote configuration management is to be executed. | 
**Operations** | [**[]RemoteConfigurationManagementOperation**](RemoteConfigurationManagementOperation.md) | A list of remote configuration management operations to be executed. | 

## Methods

### NewRemoteConfigurationManagementOperationActiveGateRequest

`func NewRemoteConfigurationManagementOperationActiveGateRequest(entities []string, operations []RemoteConfigurationManagementOperation, ) *RemoteConfigurationManagementOperationActiveGateRequest`

NewRemoteConfigurationManagementOperationActiveGateRequest instantiates a new RemoteConfigurationManagementOperationActiveGateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConfigurationManagementOperationActiveGateRequestWithDefaults

`func NewRemoteConfigurationManagementOperationActiveGateRequestWithDefaults() *RemoteConfigurationManagementOperationActiveGateRequest`

NewRemoteConfigurationManagementOperationActiveGateRequestWithDefaults instantiates a new RemoteConfigurationManagementOperationActiveGateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntities

`func (o *RemoteConfigurationManagementOperationActiveGateRequest) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *RemoteConfigurationManagementOperationActiveGateRequest) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *RemoteConfigurationManagementOperationActiveGateRequest) SetEntities(v []string)`

SetEntities sets Entities field to given value.


### GetOperations

`func (o *RemoteConfigurationManagementOperationActiveGateRequest) GetOperations() []RemoteConfigurationManagementOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *RemoteConfigurationManagementOperationActiveGateRequest) GetOperationsOk() (*[]RemoteConfigurationManagementOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *RemoteConfigurationManagementOperationActiveGateRequest) SetOperations(v []RemoteConfigurationManagementOperation)`

SetOperations sets Operations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


