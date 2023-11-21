# CreateMonitoringConfiguration207ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | The HTTP Status code | [optional] 
**ObjectId** | Pointer to **string** | The identifier of the new configuration | [optional] 
**Error** | Pointer to [**Error**](Error.md) |  | [optional] 

## Methods

### NewCreateMonitoringConfiguration207ResponseInner

`func NewCreateMonitoringConfiguration207ResponseInner() *CreateMonitoringConfiguration207ResponseInner`

NewCreateMonitoringConfiguration207ResponseInner instantiates a new CreateMonitoringConfiguration207ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMonitoringConfiguration207ResponseInnerWithDefaults

`func NewCreateMonitoringConfiguration207ResponseInnerWithDefaults() *CreateMonitoringConfiguration207ResponseInner`

NewCreateMonitoringConfiguration207ResponseInnerWithDefaults instantiates a new CreateMonitoringConfiguration207ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CreateMonitoringConfiguration207ResponseInner) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateMonitoringConfiguration207ResponseInner) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateMonitoringConfiguration207ResponseInner) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateMonitoringConfiguration207ResponseInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetObjectId

`func (o *CreateMonitoringConfiguration207ResponseInner) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *CreateMonitoringConfiguration207ResponseInner) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *CreateMonitoringConfiguration207ResponseInner) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *CreateMonitoringConfiguration207ResponseInner) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetError

`func (o *CreateMonitoringConfiguration207ResponseInner) GetError() Error`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateMonitoringConfiguration207ResponseInner) GetErrorOk() (*Error, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateMonitoringConfiguration207ResponseInner) SetError(v Error)`

SetError sets Error field to given value.

### HasError

`func (o *CreateMonitoringConfiguration207ResponseInner) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


