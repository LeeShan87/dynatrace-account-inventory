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

// checks if the UserPasswordCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPasswordCredentials{}

// UserPasswordCredentials A credentials set of the `USERNAME_PASSWORD` type.
type UserPasswordCredentials struct {
	ExternalVault *ExternalVault `json:"externalVault,omitempty"`
	// The password of the credential.
	Password *string `json:"password,omitempty"`
	// The username of the credentials set.
	User *string `json:"user,omitempty"`
}

// NewUserPasswordCredentials instantiates a new UserPasswordCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPasswordCredentials(name string, scope string, scopes []string) *UserPasswordCredentials {
	this := UserPasswordCredentials{}
	this.Name = name
	this.Scope = scope
	this.Scopes = scopes
	return &this
}

// NewUserPasswordCredentialsWithDefaults instantiates a new UserPasswordCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPasswordCredentialsWithDefaults() *UserPasswordCredentials {
	this := UserPasswordCredentials{}
	return &this
}

// GetExternalVault returns the ExternalVault field value if set, zero value otherwise.
func (o *UserPasswordCredentials) GetExternalVault() ExternalVault {
	if o == nil || IsNil(o.ExternalVault) {
		var ret ExternalVault
		return ret
	}
	return *o.ExternalVault
}

// GetExternalVaultOk returns a tuple with the ExternalVault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPasswordCredentials) GetExternalVaultOk() (*ExternalVault, bool) {
	if o == nil || IsNil(o.ExternalVault) {
		return nil, false
	}
	return o.ExternalVault, true
}

// HasExternalVault returns a boolean if a field has been set.
func (o *UserPasswordCredentials) HasExternalVault() bool {
	if o != nil && !IsNil(o.ExternalVault) {
		return true
	}

	return false
}

// SetExternalVault gets a reference to the given ExternalVault and assigns it to the ExternalVault field.
func (o *UserPasswordCredentials) SetExternalVault(v ExternalVault) {
	o.ExternalVault = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserPasswordCredentials) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPasswordCredentials) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserPasswordCredentials) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserPasswordCredentials) SetPassword(v string) {
	o.Password = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *UserPasswordCredentials) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPasswordCredentials) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *UserPasswordCredentials) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *UserPasswordCredentials) SetUser(v string) {
	o.User = &v
}

func (o UserPasswordCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPasswordCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalVault) {
		toSerialize["externalVault"] = o.ExternalVault
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableUserPasswordCredentials struct {
	value *UserPasswordCredentials
	isSet bool
}

func (v NullableUserPasswordCredentials) Get() *UserPasswordCredentials {
	return v.value
}

func (v *NullableUserPasswordCredentials) Set(val *UserPasswordCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPasswordCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPasswordCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPasswordCredentials(val *UserPasswordCredentials) *NullableUserPasswordCredentials {
	return &NullableUserPasswordCredentials{value: val, isSet: true}
}

func (v NullableUserPasswordCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPasswordCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

