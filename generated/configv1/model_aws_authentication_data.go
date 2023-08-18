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

// checks if the AwsAuthenticationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AwsAuthenticationData{}

// AwsAuthenticationData A credentials for the AWS authentication.
type AwsAuthenticationData struct {
	KeyBasedAuthentication *KeyBasedAuthentication `json:"keyBasedAuthentication,omitempty"`
	RoleBasedAuthentication *RoleBasedAuthentication `json:"roleBasedAuthentication,omitempty"`
	// The type of the authentication: role-based or key-based.
	Type string `json:"type"`
}

// NewAwsAuthenticationData instantiates a new AwsAuthenticationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAwsAuthenticationData(type_ string) *AwsAuthenticationData {
	this := AwsAuthenticationData{}
	this.Type = type_
	return &this
}

// NewAwsAuthenticationDataWithDefaults instantiates a new AwsAuthenticationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAwsAuthenticationDataWithDefaults() *AwsAuthenticationData {
	this := AwsAuthenticationData{}
	return &this
}

// GetKeyBasedAuthentication returns the KeyBasedAuthentication field value if set, zero value otherwise.
func (o *AwsAuthenticationData) GetKeyBasedAuthentication() KeyBasedAuthentication {
	if o == nil || IsNil(o.KeyBasedAuthentication) {
		var ret KeyBasedAuthentication
		return ret
	}
	return *o.KeyBasedAuthentication
}

// GetKeyBasedAuthenticationOk returns a tuple with the KeyBasedAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAuthenticationData) GetKeyBasedAuthenticationOk() (*KeyBasedAuthentication, bool) {
	if o == nil || IsNil(o.KeyBasedAuthentication) {
		return nil, false
	}
	return o.KeyBasedAuthentication, true
}

// HasKeyBasedAuthentication returns a boolean if a field has been set.
func (o *AwsAuthenticationData) HasKeyBasedAuthentication() bool {
	if o != nil && !IsNil(o.KeyBasedAuthentication) {
		return true
	}

	return false
}

// SetKeyBasedAuthentication gets a reference to the given KeyBasedAuthentication and assigns it to the KeyBasedAuthentication field.
func (o *AwsAuthenticationData) SetKeyBasedAuthentication(v KeyBasedAuthentication) {
	o.KeyBasedAuthentication = &v
}

// GetRoleBasedAuthentication returns the RoleBasedAuthentication field value if set, zero value otherwise.
func (o *AwsAuthenticationData) GetRoleBasedAuthentication() RoleBasedAuthentication {
	if o == nil || IsNil(o.RoleBasedAuthentication) {
		var ret RoleBasedAuthentication
		return ret
	}
	return *o.RoleBasedAuthentication
}

// GetRoleBasedAuthenticationOk returns a tuple with the RoleBasedAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AwsAuthenticationData) GetRoleBasedAuthenticationOk() (*RoleBasedAuthentication, bool) {
	if o == nil || IsNil(o.RoleBasedAuthentication) {
		return nil, false
	}
	return o.RoleBasedAuthentication, true
}

// HasRoleBasedAuthentication returns a boolean if a field has been set.
func (o *AwsAuthenticationData) HasRoleBasedAuthentication() bool {
	if o != nil && !IsNil(o.RoleBasedAuthentication) {
		return true
	}

	return false
}

// SetRoleBasedAuthentication gets a reference to the given RoleBasedAuthentication and assigns it to the RoleBasedAuthentication field.
func (o *AwsAuthenticationData) SetRoleBasedAuthentication(v RoleBasedAuthentication) {
	o.RoleBasedAuthentication = &v
}

// GetType returns the Type field value
func (o *AwsAuthenticationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AwsAuthenticationData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AwsAuthenticationData) SetType(v string) {
	o.Type = v
}

func (o AwsAuthenticationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AwsAuthenticationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeyBasedAuthentication) {
		toSerialize["keyBasedAuthentication"] = o.KeyBasedAuthentication
	}
	if !IsNil(o.RoleBasedAuthentication) {
		toSerialize["roleBasedAuthentication"] = o.RoleBasedAuthentication
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableAwsAuthenticationData struct {
	value *AwsAuthenticationData
	isSet bool
}

func (v NullableAwsAuthenticationData) Get() *AwsAuthenticationData {
	return v.value
}

func (v *NullableAwsAuthenticationData) Set(val *AwsAuthenticationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAwsAuthenticationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAwsAuthenticationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAwsAuthenticationData(val *AwsAuthenticationData) *NullableAwsAuthenticationData {
	return &NullableAwsAuthenticationData{value: val, isSet: true}
}

func (v NullableAwsAuthenticationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAwsAuthenticationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


