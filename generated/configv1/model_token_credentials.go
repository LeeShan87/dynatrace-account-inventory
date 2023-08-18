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

// checks if the TokenCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenCredentials{}

// TokenCredentials A credentials set of the `TOKEN` type.
type TokenCredentials struct {
	ExternalVault *ExternalVault `json:"externalVault,omitempty"`
	// Token in the string format.
	Token *string `json:"token,omitempty"`
}

// NewTokenCredentials instantiates a new TokenCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenCredentials(name string, scope string, scopes []string) *TokenCredentials {
	this := TokenCredentials{}
	this.Name = name
	this.Scope = scope
	this.Scopes = scopes
	return &this
}

// NewTokenCredentialsWithDefaults instantiates a new TokenCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenCredentialsWithDefaults() *TokenCredentials {
	this := TokenCredentials{}
	return &this
}

// GetExternalVault returns the ExternalVault field value if set, zero value otherwise.
func (o *TokenCredentials) GetExternalVault() ExternalVault {
	if o == nil || IsNil(o.ExternalVault) {
		var ret ExternalVault
		return ret
	}
	return *o.ExternalVault
}

// GetExternalVaultOk returns a tuple with the ExternalVault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCredentials) GetExternalVaultOk() (*ExternalVault, bool) {
	if o == nil || IsNil(o.ExternalVault) {
		return nil, false
	}
	return o.ExternalVault, true
}

// HasExternalVault returns a boolean if a field has been set.
func (o *TokenCredentials) HasExternalVault() bool {
	if o != nil && !IsNil(o.ExternalVault) {
		return true
	}

	return false
}

// SetExternalVault gets a reference to the given ExternalVault and assigns it to the ExternalVault field.
func (o *TokenCredentials) SetExternalVault(v ExternalVault) {
	o.ExternalVault = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TokenCredentials) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenCredentials) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TokenCredentials) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *TokenCredentials) SetToken(v string) {
	o.Token = &v
}

func (o TokenCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExternalVault) {
		toSerialize["externalVault"] = o.ExternalVault
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

type NullableTokenCredentials struct {
	value *TokenCredentials
	isSet bool
}

func (v NullableTokenCredentials) Get() *TokenCredentials {
	return v.value
}

func (v *NullableTokenCredentials) Set(val *TokenCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenCredentials(val *TokenCredentials) *NullableTokenCredentials {
	return &NullableTokenCredentials{value: val, isSet: true}
}

func (v NullableTokenCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


