/*
Dynatrace Environment API

 Documentation of the Dynatrace Environment API v2. Resources here generally supersede those in v1. Migration of resources from v1 is in progress.   If you miss a resource, consider using the Dynatrace Environment API v1. To read about use cases and examples, see [Dynatrace Documentation](https://dt-url.net/2u23k1k) .  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package environmentv2

import (
	"encoding/json"
)

// checks if the CyberArkUsernamePassword type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CyberArkUsernamePassword{}

// CyberArkUsernamePassword Synchronization credentials with CyberArk Vault using username password authentication method.
type CyberArkUsernamePassword struct {
	// Account name that stores the username and password to retrieve and synchronize with the Dynatrace Credential Vault: This is NOT the name of the account logged into the CyberArk Central Credential Provider.
	AccountName string `json:"accountName"`
	// Application ID connected to CyberArk Vault.
	ApplicationId string `json:"applicationId"`
	// [Optional] Certificate used for authentication to CyberArk application. ID of certificate credential saved in Dynatrace CV.
	Certificate *string `json:"certificate,omitempty"`
	// [Optional] Folder name where credentials in CyberArk Vault are stored. Default folder name is 'Root'.
	FolderName *string `json:"folderName,omitempty"`
	// Safe name connected to CyberArk Vault.
	SafeName string `json:"safeName"`
	// Dynatrace credential ID of the username-password pair used for authentication to the CyberArk Central Credential Provider
	UsernamePasswordForCPM string `json:"usernamePasswordForCPM"`
}

// NewCyberArkUsernamePassword instantiates a new CyberArkUsernamePassword object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCyberArkUsernamePassword(accountName string, applicationId string, safeName string, usernamePasswordForCPM string) *CyberArkUsernamePassword {
	this := CyberArkUsernamePassword{}
	this.AccountName = accountName
	this.ApplicationId = applicationId
	this.SafeName = safeName
	this.UsernamePasswordForCPM = usernamePasswordForCPM
	return &this
}

// NewCyberArkUsernamePasswordWithDefaults instantiates a new CyberArkUsernamePassword object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCyberArkUsernamePasswordWithDefaults() *CyberArkUsernamePassword {
	this := CyberArkUsernamePassword{}
	return &this
}

// GetAccountName returns the AccountName field value
func (o *CyberArkUsernamePassword) GetAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePassword) GetAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountName, true
}

// SetAccountName sets field value
func (o *CyberArkUsernamePassword) SetAccountName(v string) {
	o.AccountName = v
}

// GetApplicationId returns the ApplicationId field value
func (o *CyberArkUsernamePassword) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePassword) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *CyberArkUsernamePassword) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CyberArkUsernamePassword) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePassword) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CyberArkUsernamePassword) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *CyberArkUsernamePassword) SetCertificate(v string) {
	o.Certificate = &v
}

// GetFolderName returns the FolderName field value if set, zero value otherwise.
func (o *CyberArkUsernamePassword) GetFolderName() string {
	if o == nil || IsNil(o.FolderName) {
		var ret string
		return ret
	}
	return *o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePassword) GetFolderNameOk() (*string, bool) {
	if o == nil || IsNil(o.FolderName) {
		return nil, false
	}
	return o.FolderName, true
}

// HasFolderName returns a boolean if a field has been set.
func (o *CyberArkUsernamePassword) HasFolderName() bool {
	if o != nil && !IsNil(o.FolderName) {
		return true
	}

	return false
}

// SetFolderName gets a reference to the given string and assigns it to the FolderName field.
func (o *CyberArkUsernamePassword) SetFolderName(v string) {
	o.FolderName = &v
}

// GetSafeName returns the SafeName field value
func (o *CyberArkUsernamePassword) GetSafeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SafeName
}

// GetSafeNameOk returns a tuple with the SafeName field value
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePassword) GetSafeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SafeName, true
}

// SetSafeName sets field value
func (o *CyberArkUsernamePassword) SetSafeName(v string) {
	o.SafeName = v
}

// GetUsernamePasswordForCPM returns the UsernamePasswordForCPM field value
func (o *CyberArkUsernamePassword) GetUsernamePasswordForCPM() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsernamePasswordForCPM
}

// GetUsernamePasswordForCPMOk returns a tuple with the UsernamePasswordForCPM field value
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePassword) GetUsernamePasswordForCPMOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsernamePasswordForCPM, true
}

// SetUsernamePasswordForCPM sets field value
func (o *CyberArkUsernamePassword) SetUsernamePasswordForCPM(v string) {
	o.UsernamePasswordForCPM = v
}

func (o CyberArkUsernamePassword) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CyberArkUsernamePassword) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountName"] = o.AccountName
	toSerialize["applicationId"] = o.ApplicationId
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.FolderName) {
		toSerialize["folderName"] = o.FolderName
	}
	toSerialize["safeName"] = o.SafeName
	toSerialize["usernamePasswordForCPM"] = o.UsernamePasswordForCPM
	return toSerialize, nil
}

type NullableCyberArkUsernamePassword struct {
	value *CyberArkUsernamePassword
	isSet bool
}

func (v NullableCyberArkUsernamePassword) Get() *CyberArkUsernamePassword {
	return v.value
}

func (v *NullableCyberArkUsernamePassword) Set(val *CyberArkUsernamePassword) {
	v.value = val
	v.isSet = true
}

func (v NullableCyberArkUsernamePassword) IsSet() bool {
	return v.isSet
}

func (v *NullableCyberArkUsernamePassword) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCyberArkUsernamePassword(val *CyberArkUsernamePassword) *NullableCyberArkUsernamePassword {
	return &NullableCyberArkUsernamePassword{value: val, isSet: true}
}

func (v NullableCyberArkUsernamePassword) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCyberArkUsernamePassword) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

