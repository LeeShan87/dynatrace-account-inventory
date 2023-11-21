# AzureClientSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Client (application) ID of Azure application in Azure Active Directory which has permission to access secrets in Azure Key Vault. | [optional] 
**ClientSecret** | Pointer to **string** | Client secret generated for Azure application in Azure Active Directory used for proving identity when requesting a token used later for accessing secrets in Azure Key Vault. | [optional] 
**TenantId** | Pointer to **string** | Tenant (directory) ID of Azure application in Azure Active Directory which has permission to access secrets in Azure Key Vault. | [optional] 

## Methods

### NewAzureClientSecret

`func NewAzureClientSecret() *AzureClientSecret`

NewAzureClientSecret instantiates a new AzureClientSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureClientSecretWithDefaults

`func NewAzureClientSecretWithDefaults() *AzureClientSecret`

NewAzureClientSecretWithDefaults instantiates a new AzureClientSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *AzureClientSecret) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AzureClientSecret) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AzureClientSecret) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AzureClientSecret) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *AzureClientSecret) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AzureClientSecret) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AzureClientSecret) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AzureClientSecret) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetTenantId

`func (o *AzureClientSecret) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureClientSecret) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureClientSecret) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureClientSecret) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


