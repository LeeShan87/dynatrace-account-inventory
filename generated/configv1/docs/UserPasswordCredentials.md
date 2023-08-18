# UserPasswordCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalVault** | Pointer to [**ExternalVault**](ExternalVault.md) |  | [optional] 
**Password** | Pointer to **string** | The password of the credential. | [optional] 
**User** | Pointer to **string** | The username of the credentials set. | [optional] 

## Methods

### NewUserPasswordCredentials

`func NewUserPasswordCredentials() *UserPasswordCredentials`

NewUserPasswordCredentials instantiates a new UserPasswordCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordCredentialsWithDefaults

`func NewUserPasswordCredentialsWithDefaults() *UserPasswordCredentials`

NewUserPasswordCredentialsWithDefaults instantiates a new UserPasswordCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalVault

`func (o *UserPasswordCredentials) GetExternalVault() ExternalVault`

GetExternalVault returns the ExternalVault field if non-nil, zero value otherwise.

### GetExternalVaultOk

`func (o *UserPasswordCredentials) GetExternalVaultOk() (*ExternalVault, bool)`

GetExternalVaultOk returns a tuple with the ExternalVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVault

`func (o *UserPasswordCredentials) SetExternalVault(v ExternalVault)`

SetExternalVault sets ExternalVault field to given value.

### HasExternalVault

`func (o *UserPasswordCredentials) HasExternalVault() bool`

HasExternalVault returns a boolean if a field has been set.

### GetPassword

`func (o *UserPasswordCredentials) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserPasswordCredentials) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserPasswordCredentials) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserPasswordCredentials) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUser

`func (o *UserPasswordCredentials) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserPasswordCredentials) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserPasswordCredentials) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *UserPasswordCredentials) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


