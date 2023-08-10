# UserLoginMetaDataDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessfulLoginCounter** | **float32** | The number of successful sign-ins. | 
**FailedLoginCounter** | **float32** | The number of failed sign-ins. | 
**LastSuccessfulLogin** | **string** | The date and time of the most recent successful sign-in in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**LastFailedLogin** | **string** | The date and time of the most recent failed sign-in in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**CreatedAt** | **string** | The date and time of user creation in &#x60;2021-05-01T15:11:00Z&#x60; format. | 
**UpdatedAt** | **string** | The date and time of the most recent modification to the user in &#x60;2021-05-01T15:11:00Z&#x60; format. | 

## Methods

### NewUserLoginMetaDataDto

`func NewUserLoginMetaDataDto(successfulLoginCounter float32, failedLoginCounter float32, lastSuccessfulLogin string, lastFailedLogin string, createdAt string, updatedAt string, ) *UserLoginMetaDataDto`

NewUserLoginMetaDataDto instantiates a new UserLoginMetaDataDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserLoginMetaDataDtoWithDefaults

`func NewUserLoginMetaDataDtoWithDefaults() *UserLoginMetaDataDto`

NewUserLoginMetaDataDtoWithDefaults instantiates a new UserLoginMetaDataDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessfulLoginCounter

`func (o *UserLoginMetaDataDto) GetSuccessfulLoginCounter() float32`

GetSuccessfulLoginCounter returns the SuccessfulLoginCounter field if non-nil, zero value otherwise.

### GetSuccessfulLoginCounterOk

`func (o *UserLoginMetaDataDto) GetSuccessfulLoginCounterOk() (*float32, bool)`

GetSuccessfulLoginCounterOk returns a tuple with the SuccessfulLoginCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulLoginCounter

`func (o *UserLoginMetaDataDto) SetSuccessfulLoginCounter(v float32)`

SetSuccessfulLoginCounter sets SuccessfulLoginCounter field to given value.


### GetFailedLoginCounter

`func (o *UserLoginMetaDataDto) GetFailedLoginCounter() float32`

GetFailedLoginCounter returns the FailedLoginCounter field if non-nil, zero value otherwise.

### GetFailedLoginCounterOk

`func (o *UserLoginMetaDataDto) GetFailedLoginCounterOk() (*float32, bool)`

GetFailedLoginCounterOk returns a tuple with the FailedLoginCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginCounter

`func (o *UserLoginMetaDataDto) SetFailedLoginCounter(v float32)`

SetFailedLoginCounter sets FailedLoginCounter field to given value.


### GetLastSuccessfulLogin

`func (o *UserLoginMetaDataDto) GetLastSuccessfulLogin() string`

GetLastSuccessfulLogin returns the LastSuccessfulLogin field if non-nil, zero value otherwise.

### GetLastSuccessfulLoginOk

`func (o *UserLoginMetaDataDto) GetLastSuccessfulLoginOk() (*string, bool)`

GetLastSuccessfulLoginOk returns a tuple with the LastSuccessfulLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulLogin

`func (o *UserLoginMetaDataDto) SetLastSuccessfulLogin(v string)`

SetLastSuccessfulLogin sets LastSuccessfulLogin field to given value.


### GetLastFailedLogin

`func (o *UserLoginMetaDataDto) GetLastFailedLogin() string`

GetLastFailedLogin returns the LastFailedLogin field if non-nil, zero value otherwise.

### GetLastFailedLoginOk

`func (o *UserLoginMetaDataDto) GetLastFailedLoginOk() (*string, bool)`

GetLastFailedLoginOk returns a tuple with the LastFailedLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFailedLogin

`func (o *UserLoginMetaDataDto) SetLastFailedLogin(v string)`

SetLastFailedLogin sets LastFailedLogin field to given value.


### GetCreatedAt

`func (o *UserLoginMetaDataDto) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserLoginMetaDataDto) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserLoginMetaDataDto) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UserLoginMetaDataDto) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UserLoginMetaDataDto) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UserLoginMetaDataDto) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


