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

// checks if the ActiveGateTokenCreated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActiveGateTokenCreated{}

// ActiveGateTokenCreated The newly created ActiveGate token.
type ActiveGateTokenCreated struct {
	// The token expiration date in ISO 8601 format (`yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`).
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// The ActiveGate token identifier, consisting of [prefix and public part](https://dt-url.net/rn00tjg) of the token.
	Id string `json:"id"`
	// The secret of the token.
	Token string `json:"token"`
}

// NewActiveGateTokenCreated instantiates a new ActiveGateTokenCreated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActiveGateTokenCreated(id string, token string) *ActiveGateTokenCreated {
	this := ActiveGateTokenCreated{}
	this.Id = id
	this.Token = token
	return &this
}

// NewActiveGateTokenCreatedWithDefaults instantiates a new ActiveGateTokenCreated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActiveGateTokenCreatedWithDefaults() *ActiveGateTokenCreated {
	this := ActiveGateTokenCreated{}
	return &this
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ActiveGateTokenCreated) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenCreated) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ActiveGateTokenCreated) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *ActiveGateTokenCreated) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetId returns the Id field value
func (o *ActiveGateTokenCreated) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenCreated) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ActiveGateTokenCreated) SetId(v string) {
	o.Id = v
}

// GetToken returns the Token field value
func (o *ActiveGateTokenCreated) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ActiveGateTokenCreated) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ActiveGateTokenCreated) SetToken(v string) {
	o.Token = v
}

func (o ActiveGateTokenCreated) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActiveGateTokenCreated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	toSerialize["id"] = o.Id
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

type NullableActiveGateTokenCreated struct {
	value *ActiveGateTokenCreated
	isSet bool
}

func (v NullableActiveGateTokenCreated) Get() *ActiveGateTokenCreated {
	return v.value
}

func (v *NullableActiveGateTokenCreated) Set(val *ActiveGateTokenCreated) {
	v.value = val
	v.isSet = true
}

func (v NullableActiveGateTokenCreated) IsSet() bool {
	return v.isSet
}

func (v *NullableActiveGateTokenCreated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActiveGateTokenCreated(val *ActiveGateTokenCreated) *NullableActiveGateTokenCreated {
	return &NullableActiveGateTokenCreated{value: val, isSet: true}
}

func (v NullableActiveGateTokenCreated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActiveGateTokenCreated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


