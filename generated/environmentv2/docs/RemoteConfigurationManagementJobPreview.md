# RemoteConfigurationManagementJobPreview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyConfiguredEntitiesCount** | Pointer to **int32** | The number of entities that are currently configured as defined by remote configuration management operation. | [optional] 
**Attribute** | Pointer to **string** | The attribute which is affected by the operation. | [optional] 
**Operation** | Pointer to **string** | The operation performed on given attribute. | [optional] 
**TargetEntitiesCount** | Pointer to **int32** | The number of entities that will be configured as defined by remote configuration management after it is completed. | [optional] 
**Value** | Pointer to **string** | The value which should be assigned to given attribute. | [optional] 

## Methods

### NewRemoteConfigurationManagementJobPreview

`func NewRemoteConfigurationManagementJobPreview() *RemoteConfigurationManagementJobPreview`

NewRemoteConfigurationManagementJobPreview instantiates a new RemoteConfigurationManagementJobPreview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteConfigurationManagementJobPreviewWithDefaults

`func NewRemoteConfigurationManagementJobPreviewWithDefaults() *RemoteConfigurationManagementJobPreview`

NewRemoteConfigurationManagementJobPreviewWithDefaults instantiates a new RemoteConfigurationManagementJobPreview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyConfiguredEntitiesCount

`func (o *RemoteConfigurationManagementJobPreview) GetAlreadyConfiguredEntitiesCount() int32`

GetAlreadyConfiguredEntitiesCount returns the AlreadyConfiguredEntitiesCount field if non-nil, zero value otherwise.

### GetAlreadyConfiguredEntitiesCountOk

`func (o *RemoteConfigurationManagementJobPreview) GetAlreadyConfiguredEntitiesCountOk() (*int32, bool)`

GetAlreadyConfiguredEntitiesCountOk returns a tuple with the AlreadyConfiguredEntitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyConfiguredEntitiesCount

`func (o *RemoteConfigurationManagementJobPreview) SetAlreadyConfiguredEntitiesCount(v int32)`

SetAlreadyConfiguredEntitiesCount sets AlreadyConfiguredEntitiesCount field to given value.

### HasAlreadyConfiguredEntitiesCount

`func (o *RemoteConfigurationManagementJobPreview) HasAlreadyConfiguredEntitiesCount() bool`

HasAlreadyConfiguredEntitiesCount returns a boolean if a field has been set.

### GetAttribute

`func (o *RemoteConfigurationManagementJobPreview) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *RemoteConfigurationManagementJobPreview) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *RemoteConfigurationManagementJobPreview) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *RemoteConfigurationManagementJobPreview) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetOperation

`func (o *RemoteConfigurationManagementJobPreview) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RemoteConfigurationManagementJobPreview) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RemoteConfigurationManagementJobPreview) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *RemoteConfigurationManagementJobPreview) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetTargetEntitiesCount

`func (o *RemoteConfigurationManagementJobPreview) GetTargetEntitiesCount() int32`

GetTargetEntitiesCount returns the TargetEntitiesCount field if non-nil, zero value otherwise.

### GetTargetEntitiesCountOk

`func (o *RemoteConfigurationManagementJobPreview) GetTargetEntitiesCountOk() (*int32, bool)`

GetTargetEntitiesCountOk returns a tuple with the TargetEntitiesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEntitiesCount

`func (o *RemoteConfigurationManagementJobPreview) SetTargetEntitiesCount(v int32)`

SetTargetEntitiesCount sets TargetEntitiesCount field to given value.

### HasTargetEntitiesCount

`func (o *RemoteConfigurationManagementJobPreview) HasTargetEntitiesCount() bool`

HasTargetEntitiesCount returns a boolean if a field has been set.

### GetValue

`func (o *RemoteConfigurationManagementJobPreview) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RemoteConfigurationManagementJobPreview) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RemoteConfigurationManagementJobPreview) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RemoteConfigurationManagementJobPreview) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


