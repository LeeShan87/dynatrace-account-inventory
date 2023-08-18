# TokenCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalVault** | Pointer to [**ExternalVault**](ExternalVault.md) |  | [optional] 
**Token** | Pointer to **string** | Token in the string format. | [optional] 

## Methods

### NewTokenCredentials

`func NewTokenCredentials() *TokenCredentials`

NewTokenCredentials instantiates a new TokenCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenCredentialsWithDefaults

`func NewTokenCredentialsWithDefaults() *TokenCredentials`

NewTokenCredentialsWithDefaults instantiates a new TokenCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalVault

`func (o *TokenCredentials) GetExternalVault() ExternalVault`

GetExternalVault returns the ExternalVault field if non-nil, zero value otherwise.

### GetExternalVaultOk

`func (o *TokenCredentials) GetExternalVaultOk() (*ExternalVault, bool)`

GetExternalVaultOk returns a tuple with the ExternalVault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalVault

`func (o *TokenCredentials) SetExternalVault(v ExternalVault)`

SetExternalVault sets ExternalVault field to given value.

### HasExternalVault

`func (o *TokenCredentials) HasExternalVault() bool`

HasExternalVault returns a boolean if a field has been set.

### GetToken

`func (o *TokenCredentials) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenCredentials) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenCredentials) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenCredentials) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


