# HashicorpApprole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PathToCredentials** | **string** | Path to folder where credentials in HashiCorp Vault are stored. | 
**RoleId** | **string** | Role ID is similar to username when you want to authenticate in HashiCorp Vault using AppRole. | 
**SecretId** | **string** | Secret ID is similar to password when you want to authenticate in HashiCorp Vault using AppRole. ID of token representing secret ID saved in Dynatrace CV. | 
**VaultNamespace** | **string** | Vault namespace in HashiCorp Vault. It is an information you set as environmental variable VAULT_NAMESPACE if you are accessing HashiCorp Vault from command line. | 

## Methods

### NewHashicorpApprole

`func NewHashicorpApprole(pathToCredentials string, roleId string, secretId string, vaultNamespace string, ) *HashicorpApprole`

NewHashicorpApprole instantiates a new HashicorpApprole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHashicorpApproleWithDefaults

`func NewHashicorpApproleWithDefaults() *HashicorpApprole`

NewHashicorpApproleWithDefaults instantiates a new HashicorpApprole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPathToCredentials

`func (o *HashicorpApprole) GetPathToCredentials() string`

GetPathToCredentials returns the PathToCredentials field if non-nil, zero value otherwise.

### GetPathToCredentialsOk

`func (o *HashicorpApprole) GetPathToCredentialsOk() (*string, bool)`

GetPathToCredentialsOk returns a tuple with the PathToCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathToCredentials

`func (o *HashicorpApprole) SetPathToCredentials(v string)`

SetPathToCredentials sets PathToCredentials field to given value.


### GetRoleId

`func (o *HashicorpApprole) GetRoleId() string`

GetRoleId returns the RoleId field if non-nil, zero value otherwise.

### GetRoleIdOk

`func (o *HashicorpApprole) GetRoleIdOk() (*string, bool)`

GetRoleIdOk returns a tuple with the RoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleId

`func (o *HashicorpApprole) SetRoleId(v string)`

SetRoleId sets RoleId field to given value.


### GetSecretId

`func (o *HashicorpApprole) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *HashicorpApprole) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *HashicorpApprole) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.


### GetVaultNamespace

`func (o *HashicorpApprole) GetVaultNamespace() string`

GetVaultNamespace returns the VaultNamespace field if non-nil, zero value otherwise.

### GetVaultNamespaceOk

`func (o *HashicorpApprole) GetVaultNamespaceOk() (*string, bool)`

GetVaultNamespaceOk returns a tuple with the VaultNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultNamespace

`func (o *HashicorpApprole) SetVaultNamespace(v string)`

SetVaultNamespace sets VaultNamespace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


