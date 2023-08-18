# RemoteConfigurationManagementOperationValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | Pointer to **string** | The attribute which is affected by the operation. | [optional] 
**Operation** | Pointer to **string** | The operation performed on given attribute. | [optional] 
**Reason** | Pointer to **string** | The reason of validation failure. | [optional] 
**Value** | Pointer to **string** | The value which should be assigned to given attribute. | [optional] 

## Methods

### NewRemoteConfigurationManagementOperationValidationError

`func NewRemoteConfigurationManagementOperationValidationError() *RemoteConfigurationManagementOperationValidationError`

NewRemoteConfigurationManagementOperationValidationError instantiates a new RemoteConfigurationManagementOperationValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConfigurationManagementOperationValidationErrorWithDefaults

`func NewRemoteConfigurationManagementOperationValidationErrorWithDefaults() *RemoteConfigurationManagementOperationValidationError`

NewRemoteConfigurationManagementOperationValidationErrorWithDefaults instantiates a new RemoteConfigurationManagementOperationValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *RemoteConfigurationManagementOperationValidationError) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *RemoteConfigurationManagementOperationValidationError) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *RemoteConfigurationManagementOperationValidationError) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *RemoteConfigurationManagementOperationValidationError) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetOperation

`func (o *RemoteConfigurationManagementOperationValidationError) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RemoteConfigurationManagementOperationValidationError) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RemoteConfigurationManagementOperationValidationError) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *RemoteConfigurationManagementOperationValidationError) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetReason

`func (o *RemoteConfigurationManagementOperationValidationError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *RemoteConfigurationManagementOperationValidationError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *RemoteConfigurationManagementOperationValidationError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *RemoteConfigurationManagementOperationValidationError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetValue

`func (o *RemoteConfigurationManagementOperationValidationError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RemoteConfigurationManagementOperationValidationError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RemoteConfigurationManagementOperationValidationError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RemoteConfigurationManagementOperationValidationError) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


