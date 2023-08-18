# CyberArkUsernamePassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** | Account name that stores the username and password to retrieve and synchronize with the Dynatrace Credential Vault: This is NOT the name of the account logged into the CyberArk Central Credential Provider. | 
**ApplicationId** | **string** | Application ID connected to CyberArk Vault. | 
**Certificate** | Pointer to **string** | [Optional] Certificate used for authentication to CyberArk application. ID of certificate credential saved in Dynatrace CV. | [optional] 
**FolderName** | Pointer to **string** | [Optional] Folder name where credentials in CyberArk Vault are stored. Default folder name is &#39;Root&#39;. | [optional] 
**SafeName** | **string** | Safe name connected to CyberArk Vault. | 
**UsernamePasswordForCPM** | **string** | Dynatrace credential ID of the username-password pair used for authentication to the CyberArk Central Credential Provider | 

## Methods

### NewCyberArkUsernamePassword

`func NewCyberArkUsernamePassword(accountName string, applicationId string, safeName string, usernamePasswordForCPM string, ) *CyberArkUsernamePassword`

NewCyberArkUsernamePassword instantiates a new CyberArkUsernamePassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCyberArkUsernamePasswordWithDefaults

`func NewCyberArkUsernamePasswordWithDefaults() *CyberArkUsernamePassword`

NewCyberArkUsernamePasswordWithDefaults instantiates a new CyberArkUsernamePassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *CyberArkUsernamePassword) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CyberArkUsernamePassword) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CyberArkUsernamePassword) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetApplicationId

`func (o *CyberArkUsernamePassword) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CyberArkUsernamePassword) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CyberArkUsernamePassword) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetCertificate

`func (o *CyberArkUsernamePassword) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CyberArkUsernamePassword) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CyberArkUsernamePassword) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CyberArkUsernamePassword) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetFolderName

`func (o *CyberArkUsernamePassword) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *CyberArkUsernamePassword) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *CyberArkUsernamePassword) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *CyberArkUsernamePassword) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### GetSafeName

`func (o *CyberArkUsernamePassword) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *CyberArkUsernamePassword) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *CyberArkUsernamePassword) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.


### GetUsernamePasswordForCPM

`func (o *CyberArkUsernamePassword) GetUsernamePasswordForCPM() string`

GetUsernamePasswordForCPM returns the UsernamePasswordForCPM field if non-nil, zero value otherwise.

### GetUsernamePasswordForCPMOk

`func (o *CyberArkUsernamePassword) GetUsernamePasswordForCPMOk() (*string, bool)`

GetUsernamePasswordForCPMOk returns a tuple with the UsernamePasswordForCPM field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernamePasswordForCPM

`func (o *CyberArkUsernamePassword) SetUsernamePasswordForCPM(v string)`

SetUsernamePasswordForCPM sets UsernamePasswordForCPM field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


