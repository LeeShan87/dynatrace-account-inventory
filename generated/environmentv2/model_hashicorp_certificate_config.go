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

// checks if the HashicorpCertificateConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HashicorpCertificateConfig{}

// HashicorpCertificateConfig struct for HashicorpCertificateConfig
type HashicorpCertificateConfig struct {
	Certificate *string `json:"certificate,omitempty"`
	PathToCredentials *string `json:"pathToCredentials,omitempty"`
}

// NewHashicorpCertificateConfig instantiates a new HashicorpCertificateConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashicorpCertificateConfig() *HashicorpCertificateConfig {
	this := HashicorpCertificateConfig{}
	return &this
}

// NewHashicorpCertificateConfigWithDefaults instantiates a new HashicorpCertificateConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashicorpCertificateConfigWithDefaults() *HashicorpCertificateConfig {
	this := HashicorpCertificateConfig{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *HashicorpCertificateConfig) GetCertificate() string {
	if o == nil || IsNil(o.Certificate) {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashicorpCertificateConfig) GetCertificateOk() (*string, bool) {
	if o == nil || IsNil(o.Certificate) {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *HashicorpCertificateConfig) HasCertificate() bool {
	if o != nil && !IsNil(o.Certificate) {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *HashicorpCertificateConfig) SetCertificate(v string) {
	o.Certificate = &v
}

// GetPathToCredentials returns the PathToCredentials field value if set, zero value otherwise.
func (o *HashicorpCertificateConfig) GetPathToCredentials() string {
	if o == nil || IsNil(o.PathToCredentials) {
		var ret string
		return ret
	}
	return *o.PathToCredentials
}

// GetPathToCredentialsOk returns a tuple with the PathToCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashicorpCertificateConfig) GetPathToCredentialsOk() (*string, bool) {
	if o == nil || IsNil(o.PathToCredentials) {
		return nil, false
	}
	return o.PathToCredentials, true
}

// HasPathToCredentials returns a boolean if a field has been set.
func (o *HashicorpCertificateConfig) HasPathToCredentials() bool {
	if o != nil && !IsNil(o.PathToCredentials) {
		return true
	}

	return false
}

// SetPathToCredentials gets a reference to the given string and assigns it to the PathToCredentials field.
func (o *HashicorpCertificateConfig) SetPathToCredentials(v string) {
	o.PathToCredentials = &v
}

func (o HashicorpCertificateConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HashicorpCertificateConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Certificate) {
		toSerialize["certificate"] = o.Certificate
	}
	if !IsNil(o.PathToCredentials) {
		toSerialize["pathToCredentials"] = o.PathToCredentials
	}
	return toSerialize, nil
}

type NullableHashicorpCertificateConfig struct {
	value *HashicorpCertificateConfig
	isSet bool
}

func (v NullableHashicorpCertificateConfig) Get() *HashicorpCertificateConfig {
	return v.value
}

func (v *NullableHashicorpCertificateConfig) Set(val *HashicorpCertificateConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableHashicorpCertificateConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableHashicorpCertificateConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashicorpCertificateConfig(val *HashicorpCertificateConfig) *NullableHashicorpCertificateConfig {
	return &NullableHashicorpCertificateConfig{value: val, isSet: true}
}

func (v NullableHashicorpCertificateConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashicorpCertificateConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


