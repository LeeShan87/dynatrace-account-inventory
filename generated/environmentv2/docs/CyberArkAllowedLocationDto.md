# CyberArkAllowedLocationDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | **string** | Account name that stores the username and password to retrieve and synchronize with the Dynatrace Credential Vault: This is NOT the name of the account logged into the CyberArk Central Credential Provider. | 
**ApplicationId** | **string** | Application ID connected to CyberArk Vault. | 
**Certificate** | Pointer to **string** | [Optional] Certificate used for authentication to CyberArk application. ID of certificate credential saved in Dynatrace CV. | [optional] 
**FolderName** | Pointer to **string** | [Optional] Folder name where credentials in CyberArk Vault are stored. Default folder name is &#39;Root&#39;. | [optional] 
**SafeName** | **string** | Safe name connected to CyberArk Vault. | 

## Methods

### NewCyberArkAllowedLocationDto

`func NewCyberArkAllowedLocationDto(accountName string, applicationId string, safeName string, ) *CyberArkAllowedLocationDto`

NewCyberArkAllowedLocationDto instantiates a new CyberArkAllowedLocationDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCyberArkAllowedLocationDtoWithDefaults

`func NewCyberArkAllowedLocationDtoWithDefaults() *CyberArkAllowedLocationDto`

NewCyberArkAllowedLocationDtoWithDefaults instantiates a new CyberArkAllowedLocationDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *CyberArkAllowedLocationDto) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *CyberArkAllowedLocationDto) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *CyberArkAllowedLocationDto) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### GetApplicationId

`func (o *CyberArkAllowedLocationDto) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CyberArkAllowedLocationDto) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CyberArkAllowedLocationDto) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetCertificate

`func (o *CyberArkAllowedLocationDto) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CyberArkAllowedLocationDto) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CyberArkAllowedLocationDto) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CyberArkAllowedLocationDto) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetFolderName

`func (o *CyberArkAllowedLocationDto) GetFolderName() string`

GetFolderName returns the FolderName field if non-nil, zero value otherwise.

### GetFolderNameOk

`func (o *CyberArkAllowedLocationDto) GetFolderNameOk() (*string, bool)`

GetFolderNameOk returns a tuple with the FolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderName

`func (o *CyberArkAllowedLocationDto) SetFolderName(v string)`

SetFolderName sets FolderName field to given value.

### HasFolderName

`func (o *CyberArkAllowedLocationDto) HasFolderName() bool`

HasFolderName returns a boolean if a field has been set.

### GetSafeName

`func (o *CyberArkAllowedLocationDto) GetSafeName() string`

GetSafeName returns the SafeName field if non-nil, zero value otherwise.

### GetSafeNameOk

`func (o *CyberArkAllowedLocationDto) GetSafeNameOk() (*string, bool)`

GetSafeNameOk returns a tuple with the SafeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSafeName

`func (o *CyberArkAllowedLocationDto) SetSafeName(v string)`

SetSafeName sets SafeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


