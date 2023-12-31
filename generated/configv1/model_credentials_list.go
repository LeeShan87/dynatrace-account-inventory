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

// checks if the CredentialsList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsList{}

// CredentialsList A list of credentials sets for Synthetic monitors.
type CredentialsList struct {
	// A list of credentials sets for Synthetic monitors.
	Credentials []CredentialsResponseElement `json:"credentials"`
	NextPageKey *string `json:"nextPageKey,omitempty"`
	PageSize *int32 `json:"pageSize,omitempty"`
	TotalCount *int64 `json:"totalCount,omitempty"`
}

// NewCredentialsList instantiates a new CredentialsList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsList(credentials []CredentialsResponseElement) *CredentialsList {
	this := CredentialsList{}
	this.Credentials = credentials
	return &this
}

// NewCredentialsListWithDefaults instantiates a new CredentialsList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsListWithDefaults() *CredentialsList {
	this := CredentialsList{}
	return &this
}

// GetCredentials returns the Credentials field value
func (o *CredentialsList) GetCredentials() []CredentialsResponseElement {
	if o == nil {
		var ret []CredentialsResponseElement
		return ret
	}

	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value
// and a boolean to check if the value has been set.
func (o *CredentialsList) GetCredentialsOk() ([]CredentialsResponseElement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Credentials, true
}

// SetCredentials sets field value
func (o *CredentialsList) SetCredentials(v []CredentialsResponseElement) {
	o.Credentials = v
}

// GetNextPageKey returns the NextPageKey field value if set, zero value otherwise.
func (o *CredentialsList) GetNextPageKey() string {
	if o == nil || IsNil(o.NextPageKey) {
		var ret string
		return ret
	}
	return *o.NextPageKey
}

// GetNextPageKeyOk returns a tuple with the NextPageKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsList) GetNextPageKeyOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageKey) {
		return nil, false
	}
	return o.NextPageKey, true
}

// HasNextPageKey returns a boolean if a field has been set.
func (o *CredentialsList) HasNextPageKey() bool {
	if o != nil && !IsNil(o.NextPageKey) {
		return true
	}

	return false
}

// SetNextPageKey gets a reference to the given string and assigns it to the NextPageKey field.
func (o *CredentialsList) SetNextPageKey(v string) {
	o.NextPageKey = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *CredentialsList) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsList) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *CredentialsList) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *CredentialsList) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *CredentialsList) GetTotalCount() int64 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int64
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialsList) GetTotalCountOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *CredentialsList) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int64 and assigns it to the TotalCount field.
func (o *CredentialsList) SetTotalCount(v int64) {
	o.TotalCount = &v
}

func (o CredentialsList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialsList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["credentials"] = o.Credentials
	if !IsNil(o.NextPageKey) {
		toSerialize["nextPageKey"] = o.NextPageKey
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableCredentialsList struct {
	value *CredentialsList
	isSet bool
}

func (v NullableCredentialsList) Get() *CredentialsList {
	return v.value
}

func (v *NullableCredentialsList) Set(val *CredentialsList) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsList) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsList(val *CredentialsList) *NullableCredentialsList {
	return &NullableCredentialsList{value: val, isSet: true}
}

func (v NullableCredentialsList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


