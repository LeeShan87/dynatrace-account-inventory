# HashicorpCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | ID of certificate saved in Dynatrace CV. Using this certificate you can authenticate to your HashiCorp Vault. | [optional] 
**PathToCredentials** | Pointer to **string** | Path to folder where credentials in HashiCorp Vault are stored. | [optional] 

## Methods

### NewHashicorpCertificate

`func NewHashicorpCertificate() *HashicorpCertificate`

NewHashicorpCertificate instantiates a new HashicorpCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHashicorpCertificateWithDefaults

`func NewHashicorpCertificateWithDefaults() *HashicorpCertificate`

NewHashicorpCertificateWithDefaults instantiates a new HashicorpCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *HashicorpCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *HashicorpCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *HashicorpCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *HashicorpCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetPathToCredentials

`func (o *HashicorpCertificate) GetPathToCredentials() string`

GetPathToCredentials returns the PathToCredentials field if non-nil, zero value otherwise.

### GetPathToCredentialsOk

`func (o *HashicorpCertificate) GetPathToCredentialsOk() (*string, bool)`

GetPathToCredentialsOk returns a tuple with the PathToCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathToCredentials

`func (o *HashicorpCertificate) SetPathToCredentials(v string)`

SetPathToCredentials sets PathToCredentials field to given value.

### HasPathToCredentials

`func (o *HashicorpCertificate) HasPathToCredentials() bool`

HasPathToCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


