# ExternalVaultConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialsUsedForExternalSynchronization** | Pointer to **[]string** |  | [optional] 
**PasswordSecretName** | Pointer to **string** |  | [optional] 
**SourceAuthMethod** | Pointer to **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;HASHICORP_VAULT_APPROLE&#x60; -&gt; HashicorpApproleConfig  * &#x60;HASHICORP_VAULT_CERTIFICATE&#x60; -&gt; HashicorpCertificateConfig  * &#x60;AZURE_KEY_VAULT_CLIENT_SECRET&#x60; -&gt; AzureClientSecretConfig  * &#x60;CYBERARK_VAULT_USERNAME_PASSWORD&#x60; -&gt; CyberArkUsernamePasswordConfig  * &#x60;CYBERARK_VAULT_ALLOWED_LOCATION&#x60; -&gt; CyberArkAllowedLocationConfig   | [optional] 
**TokenSecretName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UsernameSecretName** | Pointer to **string** |  | [optional] 
**VaultUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewExternalVaultConfig

`func NewExternalVaultConfig() *ExternalVaultConfig`

NewExternalVaultConfig instantiates a new ExternalVaultConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalVaultConfigWithDefaults

`func NewExternalVaultConfigWithDefaults() *ExternalVaultConfig`

NewExternalVaultConfigWithDefaults instantiates a new ExternalVaultConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialsUsedForExternalSynchronization

`func (o *ExternalVaultConfig) GetCredentialsUsedForExternalSynchronization() []string`

GetCredentialsUsedForExternalSynchronization returns the CredentialsUsedForExternalSynchronization field if non-nil, zero value otherwise.

### GetCredentialsUsedForExternalSynchronizationOk

`func (o *ExternalVaultConfig) GetCredentialsUsedForExternalSynchronizationOk() (*[]string, bool)`

GetCredentialsUsedForExternalSynchronizationOk returns a tuple with the CredentialsUsedForExternalSynchronization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsUsedForExternalSynchronization

`func (o *ExternalVaultConfig) SetCredentialsUsedForExternalSynchronization(v []string)`

SetCredentialsUsedForExternalSynchronization sets CredentialsUsedForExternalSynchronization field to given value.

### HasCredentialsUsedForExternalSynchronization

`func (o *ExternalVaultConfig) HasCredentialsUsedForExternalSynchronization() bool`

HasCredentialsUsedForExternalSynchronization returns a boolean if a field has been set.

### GetPasswordSecretName

`func (o *ExternalVaultConfig) GetPasswordSecretName() string`

GetPasswordSecretName returns the PasswordSecretName field if non-nil, zero value otherwise.

### GetPasswordSecretNameOk

`func (o *ExternalVaultConfig) GetPasswordSecretNameOk() (*string, bool)`

GetPasswordSecretNameOk returns a tuple with the PasswordSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSecretName

`func (o *ExternalVaultConfig) SetPasswordSecretName(v string)`

SetPasswordSecretName sets PasswordSecretName field to given value.

### HasPasswordSecretName

`func (o *ExternalVaultConfig) HasPasswordSecretName() bool`

HasPasswordSecretName returns a boolean if a field has been set.

### GetSourceAuthMethod

`func (o *ExternalVaultConfig) GetSourceAuthMethod() string`

GetSourceAuthMethod returns the SourceAuthMethod field if non-nil, zero value otherwise.

### GetSourceAuthMethodOk

`func (o *ExternalVaultConfig) GetSourceAuthMethodOk() (*string, bool)`

GetSourceAuthMethodOk returns a tuple with the SourceAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAuthMethod

`func (o *ExternalVaultConfig) SetSourceAuthMethod(v string)`

SetSourceAuthMethod sets SourceAuthMethod field to given value.

### HasSourceAuthMethod

`func (o *ExternalVaultConfig) HasSourceAuthMethod() bool`

HasSourceAuthMethod returns a boolean if a field has been set.

### GetTokenSecretName

`func (o *ExternalVaultConfig) GetTokenSecretName() string`

GetTokenSecretName returns the TokenSecretName field if non-nil, zero value otherwise.

### GetTokenSecretNameOk

`func (o *ExternalVaultConfig) GetTokenSecretNameOk() (*string, bool)`

GetTokenSecretNameOk returns a tuple with the TokenSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSecretName

`func (o *ExternalVaultConfig) SetTokenSecretName(v string)`

SetTokenSecretName sets TokenSecretName field to given value.

### HasTokenSecretName

`func (o *ExternalVaultConfig) HasTokenSecretName() bool`

HasTokenSecretName returns a boolean if a field has been set.

### GetType

`func (o *ExternalVaultConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalVaultConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalVaultConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExternalVaultConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsernameSecretName

`func (o *ExternalVaultConfig) GetUsernameSecretName() string`

GetUsernameSecretName returns the UsernameSecretName field if non-nil, zero value otherwise.

### GetUsernameSecretNameOk

`func (o *ExternalVaultConfig) GetUsernameSecretNameOk() (*string, bool)`

GetUsernameSecretNameOk returns a tuple with the UsernameSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameSecretName

`func (o *ExternalVaultConfig) SetUsernameSecretName(v string)`

SetUsernameSecretName sets UsernameSecretName field to given value.

### HasUsernameSecretName

`func (o *ExternalVaultConfig) HasUsernameSecretName() bool`

HasUsernameSecretName returns a boolean if a field has been set.

### GetVaultUrl

`func (o *ExternalVaultConfig) GetVaultUrl() string`

GetVaultUrl returns the VaultUrl field if non-nil, zero value otherwise.

### GetVaultUrlOk

`func (o *ExternalVaultConfig) GetVaultUrlOk() (*string, bool)`

GetVaultUrlOk returns a tuple with the VaultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultUrl

`func (o *ExternalVaultConfig) SetVaultUrl(v string)`

SetVaultUrl sets VaultUrl field to given value.

### HasVaultUrl

`func (o *ExternalVaultConfig) HasVaultUrl() bool`

HasVaultUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


