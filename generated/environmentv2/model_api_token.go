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

// checks if the ApiToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiToken{}

// ApiToken Metadata of an API token.
type ApiToken struct {
	// Contains additional properties for specific kinds of token. Examples:  * A `dashboardId` property for dashboard sharing tokens. * A `reportId` property for report sharing tokens
	AdditionalMetadata map[string]interface{} `json:"additionalMetadata,omitempty"`
	// Token creation date in ISO 8601 format (`yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`)
	CreationDate *string `json:"creationDate,omitempty"`
	// The token is enabled (`true`) or disabled (`false`).
	Enabled *bool `json:"enabled,omitempty"`
	// Token expiration date in ISO 8601 format (`yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`).    If not set, the token never expires.
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// The ID of the token, consisting of prefix and public part of the token.
	Id *string `json:"id,omitempty"`
	// Token last used date in ISO 8601 format (`yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`)
	LastUsedDate *string `json:"lastUsedDate,omitempty"`
	// Token last used IP address.
	LastUsedIpAddress *string `json:"lastUsedIpAddress,omitempty"`
	// Token last modified date in ISO 8601 format (`yyyy-MM-dd'T'HH:mm:ss.SSS'Z'`). Updating scopes or name counts as modification, enabling or disabling a token does not.
	ModifiedDate *string `json:"modifiedDate,omitempty"`
	// The name of the token.
	Name *string `json:"name,omitempty"`
	// The owner of the token.
	Owner *string `json:"owner,omitempty"`
	// The token is a [personal access token](https://dt-url.net/wm03sop) (`true`) or an API token (`false`).
	PersonalAccessToken *bool `json:"personalAccessToken,omitempty"`
	// A list of scopes assigned to the token.
	Scopes []string `json:"scopes,omitempty"`
}

// NewApiToken instantiates a new ApiToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiToken() *ApiToken {
	this := ApiToken{}
	return &this
}

// NewApiTokenWithDefaults instantiates a new ApiToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTokenWithDefaults() *ApiToken {
	this := ApiToken{}
	return &this
}

// GetAdditionalMetadata returns the AdditionalMetadata field value if set, zero value otherwise.
func (o *ApiToken) GetAdditionalMetadata() map[string]interface{} {
	if o == nil || IsNil(o.AdditionalMetadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.AdditionalMetadata
}

// GetAdditionalMetadataOk returns a tuple with the AdditionalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetAdditionalMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AdditionalMetadata) {
		return map[string]interface{}{}, false
	}
	return o.AdditionalMetadata, true
}

// HasAdditionalMetadata returns a boolean if a field has been set.
func (o *ApiToken) HasAdditionalMetadata() bool {
	if o != nil && !IsNil(o.AdditionalMetadata) {
		return true
	}

	return false
}

// SetAdditionalMetadata gets a reference to the given map[string]interface{} and assigns it to the AdditionalMetadata field.
func (o *ApiToken) SetAdditionalMetadata(v map[string]interface{}) {
	o.AdditionalMetadata = v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *ApiToken) GetCreationDate() string {
	if o == nil || IsNil(o.CreationDate) {
		var ret string
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetCreationDateOk() (*string, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *ApiToken) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given string and assigns it to the CreationDate field.
func (o *ApiToken) SetCreationDate(v string) {
	o.CreationDate = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApiToken) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApiToken) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApiToken) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ApiToken) GetExpirationDate() string {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetExpirationDateOk() (*string, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ApiToken) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *ApiToken) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiToken) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiToken) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiToken) SetId(v string) {
	o.Id = &v
}

// GetLastUsedDate returns the LastUsedDate field value if set, zero value otherwise.
func (o *ApiToken) GetLastUsedDate() string {
	if o == nil || IsNil(o.LastUsedDate) {
		var ret string
		return ret
	}
	return *o.LastUsedDate
}

// GetLastUsedDateOk returns a tuple with the LastUsedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetLastUsedDateOk() (*string, bool) {
	if o == nil || IsNil(o.LastUsedDate) {
		return nil, false
	}
	return o.LastUsedDate, true
}

// HasLastUsedDate returns a boolean if a field has been set.
func (o *ApiToken) HasLastUsedDate() bool {
	if o != nil && !IsNil(o.LastUsedDate) {
		return true
	}

	return false
}

// SetLastUsedDate gets a reference to the given string and assigns it to the LastUsedDate field.
func (o *ApiToken) SetLastUsedDate(v string) {
	o.LastUsedDate = &v
}

// GetLastUsedIpAddress returns the LastUsedIpAddress field value if set, zero value otherwise.
func (o *ApiToken) GetLastUsedIpAddress() string {
	if o == nil || IsNil(o.LastUsedIpAddress) {
		var ret string
		return ret
	}
	return *o.LastUsedIpAddress
}

// GetLastUsedIpAddressOk returns a tuple with the LastUsedIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetLastUsedIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.LastUsedIpAddress) {
		return nil, false
	}
	return o.LastUsedIpAddress, true
}

// HasLastUsedIpAddress returns a boolean if a field has been set.
func (o *ApiToken) HasLastUsedIpAddress() bool {
	if o != nil && !IsNil(o.LastUsedIpAddress) {
		return true
	}

	return false
}

// SetLastUsedIpAddress gets a reference to the given string and assigns it to the LastUsedIpAddress field.
func (o *ApiToken) SetLastUsedIpAddress(v string) {
	o.LastUsedIpAddress = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *ApiToken) GetModifiedDate() string {
	if o == nil || IsNil(o.ModifiedDate) {
		var ret string
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetModifiedDateOk() (*string, bool) {
	if o == nil || IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *ApiToken) HasModifiedDate() bool {
	if o != nil && !IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given string and assigns it to the ModifiedDate field.
func (o *ApiToken) SetModifiedDate(v string) {
	o.ModifiedDate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiToken) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiToken) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiToken) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ApiToken) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ApiToken) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ApiToken) SetOwner(v string) {
	o.Owner = &v
}

// GetPersonalAccessToken returns the PersonalAccessToken field value if set, zero value otherwise.
func (o *ApiToken) GetPersonalAccessToken() bool {
	if o == nil || IsNil(o.PersonalAccessToken) {
		var ret bool
		return ret
	}
	return *o.PersonalAccessToken
}

// GetPersonalAccessTokenOk returns a tuple with the PersonalAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetPersonalAccessTokenOk() (*bool, bool) {
	if o == nil || IsNil(o.PersonalAccessToken) {
		return nil, false
	}
	return o.PersonalAccessToken, true
}

// HasPersonalAccessToken returns a boolean if a field has been set.
func (o *ApiToken) HasPersonalAccessToken() bool {
	if o != nil && !IsNil(o.PersonalAccessToken) {
		return true
	}

	return false
}

// SetPersonalAccessToken gets a reference to the given bool and assigns it to the PersonalAccessToken field.
func (o *ApiToken) SetPersonalAccessToken(v bool) {
	o.PersonalAccessToken = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *ApiToken) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiToken) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *ApiToken) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *ApiToken) SetScopes(v []string) {
	o.Scopes = v
}

func (o ApiToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalMetadata) {
		toSerialize["additionalMetadata"] = o.AdditionalMetadata
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creationDate"] = o.CreationDate
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastUsedDate) {
		toSerialize["lastUsedDate"] = o.LastUsedDate
	}
	if !IsNil(o.LastUsedIpAddress) {
		toSerialize["lastUsedIpAddress"] = o.LastUsedIpAddress
	}
	if !IsNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.PersonalAccessToken) {
		toSerialize["personalAccessToken"] = o.PersonalAccessToken
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableApiToken struct {
	value *ApiToken
	isSet bool
}

func (v NullableApiToken) Get() *ApiToken {
	return v.value
}

func (v *NullableApiToken) Set(val *ApiToken) {
	v.value = val
	v.isSet = true
}

func (v NullableApiToken) IsSet() bool {
	return v.isSet
}

func (v *NullableApiToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiToken(val *ApiToken) *NullableApiToken {
	return &NullableApiToken{value: val, isSet: true}
}

func (v NullableApiToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

