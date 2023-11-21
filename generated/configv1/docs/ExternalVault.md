# ExternalVault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationForSynchronizationId** | Pointer to **string** | Id of a location used by the synchronizing monitor | [optional] 
**PasswordSecretName** | Pointer to **string** | The name of the secret saved in external vault where password is stored. | [optional] 
**SourceAuthMethod** | Pointer to **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;HASHICORP_VAULT_APPROLE&#x60; -&gt; HashicorpApprole  * &#x60;HASHICORP_VAULT_CERTIFICATE&#x60; -&gt; HashicorpCertificate  * &#x60;AZURE_KEY_VAULT_CLIENT_SECRET&#x60; -&gt; AzureClientSecret  * &#x60;CYBERARK_VAULT_USERNAME_PASSWORD&#x60; -&gt; CyberArkUsernamePassword  * &#x60;CYBERARK_VAULT_ALLOWED_LOCATION&#x60; -&gt; CyberArkAllowedLocationDto   | [optional] 
**TokenSecretName** | Pointer to **string** | The name of the secret saved in external vault where token is stored. | [optional] 
**UsernameSecretName** | Pointer to **string** | The name of the secret saved in external vault where username is stored. | [optional] 
**VaultUrl** | Pointer to **string** | External vault URL. | [optional] 

## Methods

### NewExternalVault

`func NewExternalVault() *ExternalVault`

NewExternalVault instantiates a new ExternalVault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalVaultWithDefaults

`func NewExternalVaultWithDefaults() *ExternalVault`

NewExternalVaultWithDefaults instantiates a new ExternalVault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationForSynchronizationId

`func (o *ExternalVault) GetLocationForSynchronizationId() string`

GetLocationForSynchronizationId returns the LocationForSynchronizationId field if non-nil, zero value otherwise.

### GetLocationForSynchronizationIdOk

`func (o *ExternalVault) GetLocationForSynchronizationIdOk() (*string, bool)`

GetLocationForSynchronizationIdOk returns a tuple with the LocationForSynchronizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationForSynchronizationId

`func (o *ExternalVault) SetLocationForSynchronizationId(v string)`

SetLocationForSynchronizationId sets LocationForSynchronizationId field to given value.

### HasLocationForSynchronizationId

`func (o *ExternalVault) HasLocationForSynchronizationId() bool`

HasLocationForSynchronizationId returns a boolean if a field has been set.

### GetPasswordSecretName

`func (o *ExternalVault) GetPasswordSecretName() string`

GetPasswordSecretName returns the PasswordSecretName field if non-nil, zero value otherwise.

### GetPasswordSecretNameOk

`func (o *ExternalVault) GetPasswordSecretNameOk() (*string, bool)`

GetPasswordSecretNameOk returns a tuple with the PasswordSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSecretName

`func (o *ExternalVault) SetPasswordSecretName(v string)`

SetPasswordSecretName sets PasswordSecretName field to given value.

### HasPasswordSecretName

`func (o *ExternalVault) HasPasswordSecretName() bool`

HasPasswordSecretName returns a boolean if a field has been set.

### GetSourceAuthMethod

`func (o *ExternalVault) GetSourceAuthMethod() string`

GetSourceAuthMethod returns the SourceAuthMethod field if non-nil, zero value otherwise.

### GetSourceAuthMethodOk

`func (o *ExternalVault) GetSourceAuthMethodOk() (*string, bool)`

GetSourceAuthMethodOk returns a tuple with the SourceAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAuthMethod

`func (o *ExternalVault) SetSourceAuthMethod(v string)`

SetSourceAuthMethod sets SourceAuthMethod field to given value.

### HasSourceAuthMethod

`func (o *ExternalVault) HasSourceAuthMethod() bool`

HasSourceAuthMethod returns a boolean if a field has been set.

### GetTokenSecretName

`func (o *ExternalVault) GetTokenSecretName() string`

GetTokenSecretName returns the TokenSecretName field if non-nil, zero value otherwise.

### GetTokenSecretNameOk

`func (o *ExternalVault) GetTokenSecretNameOk() (*string, bool)`

GetTokenSecretNameOk returns a tuple with the TokenSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSecretName

`func (o *ExternalVault) SetTokenSecretName(v string)`

SetTokenSecretName sets TokenSecretName field to given value.

### HasTokenSecretName

`func (o *ExternalVault) HasTokenSecretName() bool`

HasTokenSecretName returns a boolean if a field has been set.

### GetUsernameSecretName

`func (o *ExternalVault) GetUsernameSecretName() string`

GetUsernameSecretName returns the UsernameSecretName field if non-nil, zero value otherwise.

### GetUsernameSecretNameOk

`func (o *ExternalVault) GetUsernameSecretNameOk() (*string, bool)`

GetUsernameSecretNameOk returns a tuple with the UsernameSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameSecretName

`func (o *ExternalVault) SetUsernameSecretName(v string)`

SetUsernameSecretName sets UsernameSecretName field to given value.

### HasUsernameSecretName

`func (o *ExternalVault) HasUsernameSecretName() bool`

HasUsernameSecretName returns a boolean if a field has been set.

### GetVaultUrl

`func (o *ExternalVault) GetVaultUrl() string`

GetVaultUrl returns the VaultUrl field if non-nil, zero value otherwise.

### GetVaultUrlOk

`func (o *ExternalVault) GetVaultUrlOk() (*string, bool)`

GetVaultUrlOk returns a tuple with the VaultUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultUrl

`func (o *ExternalVault) SetVaultUrl(v string)`

SetVaultUrl sets VaultUrl field to given value.

### HasVaultUrl

`func (o *ExternalVault) HasVaultUrl() bool`

HasVaultUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


