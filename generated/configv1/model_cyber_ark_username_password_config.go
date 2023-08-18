/*
Dynatrace Configuration API

Documentation of the Dynatrace Configuration API. To read about use-cases and examples, see [Dynatrace Documentation](https://dt-url.net/4u43kxw).  Notes about compatibility: * Operations marked as early adopter or preview may be changed in non-compatible ways, although we try to avoid this. * We may add new enum constants without incrementing the API version; thus, clients need to handle unknown enum constants gracefully.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package configv1

import (
	"encoding/json"
)

// checks if the CyberArkUsernamePasswordConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CyberArkUsernamePasswordConfig{}

// CyberArkUsernamePasswordConfig struct for CyberArkUsernamePasswordConfig
type CyberArkUsernamePasswordConfig struct {
	AccountName *string `json:"accountName,omitempty"`
	ApplicationId *string `json:"applicationId,omitempty"`
	Certificate *string `json:"certificate,omitempty"`
	FolderName *string `json:"folderName,omitempty"`
	SafeName *string `json:"safeName,omitempty"`
	UsernamePasswordForCPM *string `json:"usernamePasswordForCPM,omitempty"`
}

// NewCyberArkUsernamePasswordConfig instantiates a new CyberArkUsernamePasswordConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCyberArkUsernamePasswordConfig() *CyberArkUsernamePasswordConfig {
	this := CyberArkUsernamePasswordConfig{}
	return &this
}

// NewCyberArkUsernamePasswordConfigWithDefaults instantiates a new CyberArkUsernamePasswordConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCyberArkUsernamePasswordConfigWithDefaults() *CyberArkUsernamePasswordConfig {
	this := CyberArkUsernamePasswordConfig{}
	return &this
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *CyberArkUsernamePasswordConfig) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePasswordConfig) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *CyberArkUsernamePasswordConfig) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *CyberArkUsernamePasswordConfig) SetAccountName(v string) {
	o.AccountName = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *CyberArkUsernamePasswordConfig) GetApplicationId() string {
	if o == nil || IsNil(o.ApplicationId) {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePasswordConfig) GetApplicationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *CyberArkUsernamePasswordConfig) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *CyberArkUsernamePasswordConfig) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *CyberArkUsernamePasswordConfig) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePasswordConfig) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *CyberArkUsernamePasswordConfig) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *CyberArkUsernamePasswordConfig) SetCertificate(v string) {
	o.Certificate = &v
}

// GetFolderName returns the FolderName field value if set, zero value otherwise.
func (o *CyberArkUsernamePasswordConfig) GetFolderName() string {
	if o == nil || IsNil(o.FolderName) {
		var ret string
		return ret
	}
	return *o.FolderName
}

// GetFolderNameOk returns a tuple with the FolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePasswordConfig) GetFolderNameOk() (*string, bool) {
	if o == nil || IsNil(o.FolderName) {
		return nil, false
	}
	return o.FolderName, true
}

// HasFolderName returns a boolean if a field has been set.
func (o *CyberArkUsernamePasswordConfig) HasFolderName() bool {
	if o != nil && !IsNil(o.FolderName) {
		return true
	}

	return false
}

// SetFolderName gets a reference to the given string and assigns it to the FolderName field.
func (o *CyberArkUsernamePasswordConfig) SetFolderName(v string) {
	o.FolderName = &v
}

// GetSafeName returns the SafeName field value if set, zero value otherwise.
func (o *CyberArkUsernamePasswordConfig) GetSafeName() string {
	if o == nil || IsNil(o.SafeName) {
		var ret string
		return ret
	}
	return *o.SafeName
}

// GetSafeNameOk returns a tuple with the SafeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePasswordConfig) GetSafeNameOk() (*string, bool) {
	if o == nil || IsNil(o.SafeName) {
		return nil, false
	}
	return o.SafeName, true
}

// HasSafeName returns a boolean if a field has been set.
func (o *CyberArkUsernamePasswordConfig) HasSafeName() bool {
	if o != nil && !IsNil(o.SafeName) {
		return true
	}

	return false
}

// SetSafeName gets a reference to the given string and assigns it to the SafeName field.
func (o *CyberArkUsernamePasswordConfig) SetSafeName(v string) {
	o.SafeName = &v
}

// GetUsernamePasswordForCPM returns the UsernamePasswordForCPM field value if set, zero value otherwise.
func (o *CyberArkUsernamePasswordConfig) GetUsernamePasswordForCPM() string {
	if o == nil || IsNil(o.UsernamePasswordForCPM) {
		var ret string
		return ret
	}
	return *o.UsernamePasswordForCPM
}

// GetUsernamePasswordForCPMOk returns a tuple with the UsernamePasswordForCPM field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CyberArkUsernamePasswordConfig) GetUsernamePasswordForCPMOk() (*string, bool) {
	if o == nil || IsNil(o.UsernamePasswordForCPM) {
		return nil, false
	}
	return o.UsernamePasswordForCPM, true
}

// HasUsernamePasswordForCPM returns a boolean if a field has been set.
func (o *CyberArkUsernamePasswordConfig) HasUsernamePasswordForCPM() bool {
	if o != nil && !IsNil(o.UsernamePasswordForCPM) {
		return true
	}

	return false
}

// SetUsernamePasswordForCPM gets a reference to the given string and assigns it to the UsernamePasswordForCPM field.
func (o *CyberArkUsernamePasswordConfig) SetUsernamePasswordForCPM(v string) {
	o.UsernamePasswordForCPM = &v
}

func (o CyberArkUsernamePasswordConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CyberArkUsernamePasswordConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !IsNil(o.ApplicationId) {
		toSerialize["applicationId"] = o.ApplicationId
	}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.FolderName) {
		toSerialize["folderName"] = o.FolderName
	}
	if !IsNil(o.SafeName) {
		toSerialize["safeName"] = o.SafeName
	}
	if !IsNil(o.UsernamePasswordForCPM) {
		toSerialize["usernamePasswordForCPM"] = o.UsernamePasswordForCPM
	}
	return toSerialize, nil
}

type NullableCyberArkUsernamePasswordConfig struct {
	value *CyberArkUsernamePasswordConfig
	isSet bool
}

func (v NullableCyberArkUsernamePasswordConfig) Get() *CyberArkUsernamePasswordConfig {
	return v.value
}

func (v *NullableCyberArkUsernamePasswordConfig) Set(val *CyberArkUsernamePasswordConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCyberArkUsernamePasswordConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCyberArkUsernamePasswordConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCyberArkUsernamePasswordConfig(val *CyberArkUsernamePasswordConfig) *NullableCyberArkUsernamePasswordConfig {
	return &NullableCyberArkUsernamePasswordConfig{value: val, isSet: true}
}

func (v NullableCyberArkUsernamePasswordConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCyberArkUsernamePasswordConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


