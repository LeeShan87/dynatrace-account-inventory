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

// checks if the CredentialsDetailsUsernamePasswordResponseElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsDetailsUsernamePasswordResponseElement{}

// CredentialsDetailsUsernamePasswordResponseElement Details of username and password credentials set.
type CredentialsDetailsUsernamePasswordResponseElement struct {
	// Plain text password value
	Password *string `json:"password,omitempty"`
	// Plain text username value
	Username *string `json:"username,omitempty"`
}

// NewCredentialsDetailsUsernamePasswordResponseElement instantiates a new CredentialsDetailsUsernamePasswordResponseElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsDetailsUsernamePasswordResponseElement(credentialUsageSummary []CredentialUsageHandler, description string, name string, owner string, ownerAccessOnly bool, type_ string) *CredentialsDetailsUsernamePasswordResponseElement {
	this := CredentialsDetailsUsernamePasswordResponseElement{}
	this.CredentialUsageSummary = credentialUsageSummary
	this.Description = description
	this.Name = name
	this.Owner = owner
	this.OwnerAccessOnly = ownerAccessOnly
	this.Type = type_
	return &this
}

// NewCredentialsDetailsUsernamePasswordResponseElementWithDefaults instantiates a new CredentialsDetailsUsernamePasswordResponseElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsDetailsUsernamePasswordResponseElementWithDefaults() *CredentialsDetailsUsernamePasswordResponseElement {
	this := CredentialsDetailsUsernamePasswordResponseElement{}
	return &this
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *CredentialsDetailsUsernamePasswordResponseElement) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsDetailsUsernamePasswordResponseElement) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *CredentialsDetailsUsernamePasswordResponseElement) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *CredentialsDetailsUsernamePasswordResponseElement) SetPassword(v string) {
	o.Password = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CredentialsDetailsUsernamePasswordResponseElement) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsDetailsUsernamePasswordResponseElement) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CredentialsDetailsUsernamePasswordResponseElement) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CredentialsDetailsUsernamePasswordResponseElement) SetUsername(v string) {
	o.Username = &v
}

func (o CredentialsDetailsUsernamePasswordResponseElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialsDetailsUsernamePasswordResponseElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableCredentialsDetailsUsernamePasswordResponseElement struct {
	value *CredentialsDetailsUsernamePasswordResponseElement
	isSet bool
}

func (v NullableCredentialsDetailsUsernamePasswordResponseElement) Get() *CredentialsDetailsUsernamePasswordResponseElement {
	return v.value
}

func (v *NullableCredentialsDetailsUsernamePasswordResponseElement) Set(val *CredentialsDetailsUsernamePasswordResponseElement) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsDetailsUsernamePasswordResponseElement) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsDetailsUsernamePasswordResponseElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsDetailsUsernamePasswordResponseElement(val *CredentialsDetailsUsernamePasswordResponseElement) *NullableCredentialsDetailsUsernamePasswordResponseElement {
	return &NullableCredentialsDetailsUsernamePasswordResponseElement{value: val, isSet: true}
}

func (v NullableCredentialsDetailsUsernamePasswordResponseElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsDetailsUsernamePasswordResponseElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


