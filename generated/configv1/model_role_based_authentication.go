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

// checks if the RoleBasedAuthentication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleBasedAuthentication{}

// RoleBasedAuthentication The credentials for the role-based authentication.
type RoleBasedAuthentication struct {
	// The ID of the Amazon account.
	AccountId string `json:"accountId"`
	// The external ID token for setting an IAM role.    You can obtain it with the `GET /aws/iamExternalId` request.
	ExternalId *string `json:"externalId,omitempty"`
	// The IAM role to be used by Dynatrace to get monitoring data.
	IamRole string `json:"iamRole"`
}

// NewRoleBasedAuthentication instantiates a new RoleBasedAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleBasedAuthentication(accountId string, iamRole string) *RoleBasedAuthentication {
	this := RoleBasedAuthentication{}
	this.AccountId = accountId
	this.IamRole = iamRole
	return &this
}

// NewRoleBasedAuthenticationWithDefaults instantiates a new RoleBasedAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleBasedAuthenticationWithDefaults() *RoleBasedAuthentication {
	this := RoleBasedAuthentication{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *RoleBasedAuthentication) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *RoleBasedAuthentication) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *RoleBasedAuthentication) SetAccountId(v string) {
	o.AccountId = v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *RoleBasedAuthentication) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBasedAuthentication) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *RoleBasedAuthentication) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *RoleBasedAuthentication) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetIamRole returns the IamRole field value
func (o *RoleBasedAuthentication) GetIamRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IamRole
}

// GetIamRoleOk returns a tuple with the IamRole field value
// and a boolean to check if the value has been set.
func (o *RoleBasedAuthentication) GetIamRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IamRole, true
}

// SetIamRole sets field value
func (o *RoleBasedAuthentication) SetIamRole(v string) {
	o.IamRole = v
}

func (o RoleBasedAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleBasedAuthentication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	toSerialize["iamRole"] = o.IamRole
	return toSerialize, nil
}

type NullableRoleBasedAuthentication struct {
	value *RoleBasedAuthentication
	isSet bool
}

func (v NullableRoleBasedAuthentication) Get() *RoleBasedAuthentication {
	return v.value
}

func (v *NullableRoleBasedAuthentication) Set(val *RoleBasedAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleBasedAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleBasedAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleBasedAuthentication(val *RoleBasedAuthentication) *NullableRoleBasedAuthentication {
	return &NullableRoleBasedAuthentication{value: val, isSet: true}
}

func (v NullableRoleBasedAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleBasedAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

