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

// checks if the AzureClientSecretConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AzureClientSecretConfig{}

// AzureClientSecretConfig struct for AzureClientSecretConfig
type AzureClientSecretConfig struct {
	ClientId *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
}

// NewAzureClientSecretConfig instantiates a new AzureClientSecretConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAzureClientSecretConfig() *AzureClientSecretConfig {
	this := AzureClientSecretConfig{}
	return &this
}

// NewAzureClientSecretConfigWithDefaults instantiates a new AzureClientSecretConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAzureClientSecretConfigWithDefaults() *AzureClientSecretConfig {
	this := AzureClientSecretConfig{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AzureClientSecretConfig) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureClientSecretConfig) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AzureClientSecretConfig) HasClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AzureClientSecretConfig) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AzureClientSecretConfig) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureClientSecretConfig) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *AzureClientSecretConfig) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AzureClientSecretConfig) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *AzureClientSecretConfig) GetTenantId() string {
	if o == nil || IsNil(o.TenantId) {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AzureClientSecretConfig) GetTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *AzureClientSecretConfig) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *AzureClientSecretConfig) SetTenantId(v string) {
	o.TenantId = &v
}

func (o AzureClientSecretConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AzureClientSecretConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	return toSerialize, nil
}

type NullableAzureClientSecretConfig struct {
	value *AzureClientSecretConfig
	isSet bool
}

func (v NullableAzureClientSecretConfig) Get() *AzureClientSecretConfig {
	return v.value
}

func (v *NullableAzureClientSecretConfig) Set(val *AzureClientSecretConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAzureClientSecretConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAzureClientSecretConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAzureClientSecretConfig(val *AzureClientSecretConfig) *NullableAzureClientSecretConfig {
	return &NullableAzureClientSecretConfig{value: val, isSet: true}
}

func (v NullableAzureClientSecretConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAzureClientSecretConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


