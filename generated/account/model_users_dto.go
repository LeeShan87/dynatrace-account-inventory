/*
Dynatrace Account Management API

The enterprise management API for Dynatrace SaaS enables automation of operational tasks related to user access and environment lifecycle management.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package account

import (
	"encoding/json"
)

// checks if the UsersDto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsersDto{}

// UsersDto struct for UsersDto
type UsersDto struct {
	// The UUID of the user.
	Uid string `json:"uid"`
	// The email address of the user.
	Email string `json:"email"`
	// The first name of the user.
	Name *string `json:"name,omitempty"`
	// The last name of the user.
	Surname *string `json:"surname,omitempty"`
	// The status of this user in Dynatrace:   * `ACTIVE`: The user is active. * `INACTIVE`: The user is deactivated and cannot sign in to Dynatrace.  * `PENDING`: The user received an invitation, but hasn't completed sign-up yet.  * `DELETED`: The user was deleted and cannot sign in to Dynatrace anymore.  * `ECUSTOMS_MANUALLY_BLOCKED`: The user is blocked due to to a trade and export compliance violation.  
	UserStatus *string `json:"userStatus,omitempty"`
	// The user is (`true`) or is not (`false`) an emergency contact for the account.
	EmergencyContact *bool `json:"emergencyContact,omitempty"`
	UserLoginMetadata *UserLoginMetaDataDto `json:"userLoginMetadata,omitempty"`
}

// NewUsersDto instantiates a new UsersDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsersDto(uid string, email string) *UsersDto {
	this := UsersDto{}
	this.Uid = uid
	this.Email = email
	return &this
}

// NewUsersDtoWithDefaults instantiates a new UsersDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsersDtoWithDefaults() *UsersDto {
	this := UsersDto{}
	return &this
}

// GetUid returns the Uid field value
func (o *UsersDto) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *UsersDto) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *UsersDto) SetUid(v string) {
	o.Uid = v
}

// GetEmail returns the Email field value
func (o *UsersDto) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *UsersDto) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *UsersDto) SetEmail(v string) {
	o.Email = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UsersDto) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersDto) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UsersDto) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UsersDto) SetName(v string) {
	o.Name = &v
}

// GetSurname returns the Surname field value if set, zero value otherwise.
func (o *UsersDto) GetSurname() string {
	if o == nil || IsNil(o.Surname) {
		var ret string
		return ret
	}
	return *o.Surname
}

// GetSurnameOk returns a tuple with the Surname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersDto) GetSurnameOk() (*string, bool) {
	if o == nil || IsNil(o.Surname) {
		return nil, false
	}
	return o.Surname, true
}

// HasSurname returns a boolean if a field has been set.
func (o *UsersDto) HasSurname() bool {
	if o != nil && !IsNil(o.Surname) {
		return true
	}

	return false
}

// SetSurname gets a reference to the given string and assigns it to the Surname field.
func (o *UsersDto) SetSurname(v string) {
	o.Surname = &v
}

// GetUserStatus returns the UserStatus field value if set, zero value otherwise.
func (o *UsersDto) GetUserStatus() string {
	if o == nil || IsNil(o.UserStatus) {
		var ret string
		return ret
	}
	return *o.UserStatus
}

// GetUserStatusOk returns a tuple with the UserStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersDto) GetUserStatusOk() (*string, bool) {
	if o == nil || IsNil(o.UserStatus) {
		return nil, false
	}
	return o.UserStatus, true
}

// HasUserStatus returns a boolean if a field has been set.
func (o *UsersDto) HasUserStatus() bool {
	if o != nil && !IsNil(o.UserStatus) {
		return true
	}

	return false
}

// SetUserStatus gets a reference to the given string and assigns it to the UserStatus field.
func (o *UsersDto) SetUserStatus(v string) {
	o.UserStatus = &v
}

// GetEmergencyContact returns the EmergencyContact field value if set, zero value otherwise.
func (o *UsersDto) GetEmergencyContact() bool {
	if o == nil || IsNil(o.EmergencyContact) {
		var ret bool
		return ret
	}
	return *o.EmergencyContact
}

// GetEmergencyContactOk returns a tuple with the EmergencyContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersDto) GetEmergencyContactOk() (*bool, bool) {
	if o == nil || IsNil(o.EmergencyContact) {
		return nil, false
	}
	return o.EmergencyContact, true
}

// HasEmergencyContact returns a boolean if a field has been set.
func (o *UsersDto) HasEmergencyContact() bool {
	if o != nil && !IsNil(o.EmergencyContact) {
		return true
	}

	return false
}

// SetEmergencyContact gets a reference to the given bool and assigns it to the EmergencyContact field.
func (o *UsersDto) SetEmergencyContact(v bool) {
	o.EmergencyContact = &v
}

// GetUserLoginMetadata returns the UserLoginMetadata field value if set, zero value otherwise.
func (o *UsersDto) GetUserLoginMetadata() UserLoginMetaDataDto {
	if o == nil || IsNil(o.UserLoginMetadata) {
		var ret UserLoginMetaDataDto
		return ret
	}
	return *o.UserLoginMetadata
}

// GetUserLoginMetadataOk returns a tuple with the UserLoginMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsersDto) GetUserLoginMetadataOk() (*UserLoginMetaDataDto, bool) {
	if o == nil || IsNil(o.UserLoginMetadata) {
		return nil, false
	}
	return o.UserLoginMetadata, true
}

// HasUserLoginMetadata returns a boolean if a field has been set.
func (o *UsersDto) HasUserLoginMetadata() bool {
	if o != nil && !IsNil(o.UserLoginMetadata) {
		return true
	}

	return false
}

// SetUserLoginMetadata gets a reference to the given UserLoginMetaDataDto and assigns it to the UserLoginMetadata field.
func (o *UsersDto) SetUserLoginMetadata(v UserLoginMetaDataDto) {
	o.UserLoginMetadata = &v
}

func (o UsersDto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsersDto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uid"] = o.Uid
	toSerialize["email"] = o.Email
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Surname) {
		toSerialize["surname"] = o.Surname
	}
	if !IsNil(o.UserStatus) {
		toSerialize["userStatus"] = o.UserStatus
	}
	if !IsNil(o.EmergencyContact) {
		toSerialize["emergencyContact"] = o.EmergencyContact
	}
	if !IsNil(o.UserLoginMetadata) {
		toSerialize["userLoginMetadata"] = o.UserLoginMetadata
	}
	return toSerialize, nil
}

type NullableUsersDto struct {
	value *UsersDto
	isSet bool
}

func (v NullableUsersDto) Get() *UsersDto {
	return v.value
}

func (v *NullableUsersDto) Set(val *UsersDto) {
	v.value = val
	v.isSet = true
}

func (v NullableUsersDto) IsSet() bool {
	return v.isSet
}

func (v *NullableUsersDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsersDto(val *UsersDto) *NullableUsersDto {
	return &NullableUsersDto{value: val, isSet: true}
}

func (v NullableUsersDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsersDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


