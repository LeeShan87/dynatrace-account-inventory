# RemoteConfigurationManagementEntityValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to **string** | The ID of the entity for which validation failed. | [optional] 
**Reasons** | Pointer to **[]string** | The reason of entity validation failure. | [optional] 

## Methods

### NewRemoteConfigurationManagementEntityValidationError

`func NewRemoteConfigurationManagementEntityValidationError() *RemoteConfigurationManagementEntityValidationError`

NewRemoteConfigurationManagementEntityValidationError instantiates a new RemoteConfigurationManagementEntityValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConfigurationManagementEntityValidationErrorWithDefaults

`func NewRemoteConfigurationManagementEntityValidationErrorWithDefaults() *RemoteConfigurationManagementEntityValidationError`

NewRemoteConfigurationManagementEntityValidationErrorWithDefaults instantiates a new RemoteConfigurationManagementEntityValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *RemoteConfigurationManagementEntityValidationError) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *RemoteConfigurationManagementEntityValidationError) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *RemoteConfigurationManagementEntityValidationError) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *RemoteConfigurationManagementEntityValidationError) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetReasons

`func (o *RemoteConfigurationManagementEntityValidationError) GetReasons() []string`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *RemoteConfigurationManagementEntityValidationError) GetReasonsOk() (*[]string, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *RemoteConfigurationManagementEntityValidationError) SetReasons(v []string)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *RemoteConfigurationManagementEntityValidationError) HasReasons() bool`

HasReasons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


