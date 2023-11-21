# RemoteConfigurationManagementOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attribute** | **string** | The attribute which is affected by the operation. | 
**Operation** | **string** | The operation performed on given attribute. | 
**Value** | Pointer to **string** | The value which should be assigned to given attribute. | [optional] 

## Methods

### NewRemoteConfigurationManagementOperation

`func NewRemoteConfigurationManagementOperation(attribute string, operation string, ) *RemoteConfigurationManagementOperation`

NewRemoteConfigurationManagementOperation instantiates a new RemoteConfigurationManagementOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConfigurationManagementOperationWithDefaults

`func NewRemoteConfigurationManagementOperationWithDefaults() *RemoteConfigurationManagementOperation`

NewRemoteConfigurationManagementOperationWithDefaults instantiates a new RemoteConfigurationManagementOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttribute

`func (o *RemoteConfigurationManagementOperation) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *RemoteConfigurationManagementOperation) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *RemoteConfigurationManagementOperation) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.


### GetOperation

`func (o *RemoteConfigurationManagementOperation) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RemoteConfigurationManagementOperation) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RemoteConfigurationManagementOperation) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetValue

`func (o *RemoteConfigurationManagementOperation) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RemoteConfigurationManagementOperation) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RemoteConfigurationManagementOperation) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RemoteConfigurationManagementOperation) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


