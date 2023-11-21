# AbstractCredentialsResponseElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialUsageSummary** | [**[]CredentialUsageHandler**](CredentialUsageHandler.md) | The list contains summary data related to the use of credentials. | 
**Description** | **string** | A short description of the credentials set. | 
**ExternalVault** | Pointer to [**ExternalVaultConfig**](ExternalVaultConfig.md) |  | [optional] 
**Id** | Pointer to **string** | The ID of the credentials set. | [optional] 
**Name** | **string** | The name of the credentials set. | 
**Owner** | **string** | The owner of the credential (user for which used API token was created). | 
**OwnerAccessOnly** | **bool** | Flag indicating that this credential is visible only to the owner. | 
**Scope** | Pointer to **string** | The scope of the credentials set. | [optional] 
**Scopes** | Pointer to **[]string** | The set of scopes of the credentials set. | [optional] 
**Type** | **string** | Defines the actual set of fields depending on the value. See one of the following objects:   * &#x60;USERNAME_PASSWORD&#x60; -&gt; CredentialsDetailsUsernamePasswordResponseElement  * &#x60;CERTIFICATE&#x60; -&gt; CredentialsDetailsCertificateResponseElement  * &#x60;TOKEN&#x60; -&gt; CredentialsDetailsTokenResponseElement  * &#x60;PUBLIC_CERTIFICATE&#x60; -&gt; CredentialsDetailsCertificateResponseElement   | 

## Methods

### NewAbstractCredentialsResponseElement

`func NewAbstractCredentialsResponseElement(credentialUsageSummary []CredentialUsageHandler, description string, name string, owner string, ownerAccessOnly bool, type_ string, ) *AbstractCredentialsResponseElement`

NewAbstractCredentialsResponseElement instantiates a new AbstractCredentialsResponseElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAbstractCredentialsResponseElementWithDefaults

`func NewAbstractCredentialsResponseElementWithDefaults() *AbstractCredentialsResponseElement`

NewAbstractCredentialsResponseElementWithDefaults instantiates a new AbstractCredentialsResponseElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialUsageSummary

`func (o *AbstractCredentialsResponseElement) GetCredentialUsageSummary() []CredentialUsageHandler`

GetCredentialUsageSummary returns the CredentialUsageSummary field if non-nil, zero value otherwise.

### GetCredentialUsageSummaryOk

`func (o *AbstractCredentialsResponseElement) GetCredentialUsageSummaryOk() (*[]CredentialUsageHandler, bool)`

GetCredentialUsageSummaryOk returns a tuple with the CredentialUsageSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialUsageSummary

`func (o *AbstractCredentialsResponseElement) SetCredentialUsageSummary(v []CredentialUsageHandler)`

SetCredentialUsageSummary sets CredentialUsageSummary field to given value.


### GetDescription

`func (o *AbstractCredentialsResponseElement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AbstractCredentialsResponseElement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AbstractCredentialsResponseElement) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExternalVault

`func (o *AbstractCredentialsResponseElement) GetExternalVault() ExternalVaultConfig`

GetExternalVault returns the ExternalVault field if non-nil, zero value otherwise.

### GetExternalVaultOk

`func (o *AbstractCredentialsResponseElement) GetExternalVaultOk() (*ExternalVaultConfig, bool)`

GetExternalVaultOk returns a tuple with the ExternalVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVault

`func (o *AbstractCredentialsResponseElement) SetExternalVault(v ExternalVaultConfig)`

SetExternalVault sets ExternalVault field to given value.

### HasExternalVault

`func (o *AbstractCredentialsResponseElement) HasExternalVault() bool`

HasExternalVault returns a boolean if a field has been set.

### GetId

`func (o *AbstractCredentialsResponseElement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AbstractCredentialsResponseElement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AbstractCredentialsResponseElement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AbstractCredentialsResponseElement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AbstractCredentialsResponseElement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AbstractCredentialsResponseElement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AbstractCredentialsResponseElement) SetName(v string)`

SetName sets Name field to given value.


### GetOwner

`func (o *AbstractCredentialsResponseElement) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AbstractCredentialsResponseElement) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AbstractCredentialsResponseElement) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetOwnerAccessOnly

`func (o *AbstractCredentialsResponseElement) GetOwnerAccessOnly() bool`

GetOwnerAccessOnly returns the OwnerAccessOnly field if non-nil, zero value otherwise.

### GetOwnerAccessOnlyOk

`func (o *AbstractCredentialsResponseElement) GetOwnerAccessOnlyOk() (*bool, bool)`

GetOwnerAccessOnlyOk returns a tuple with the OwnerAccessOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAccessOnly

`func (o *AbstractCredentialsResponseElement) SetOwnerAccessOnly(v bool)`

SetOwnerAccessOnly sets OwnerAccessOnly field to given value.


### GetScope

`func (o *AbstractCredentialsResponseElement) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AbstractCredentialsResponseElement) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AbstractCredentialsResponseElement) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AbstractCredentialsResponseElement) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetScopes

`func (o *AbstractCredentialsResponseElement) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AbstractCredentialsResponseElement) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AbstractCredentialsResponseElement) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AbstractCredentialsResponseElement) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetType

`func (o *AbstractCredentialsResponseElement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AbstractCredentialsResponseElement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AbstractCredentialsResponseElement) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


