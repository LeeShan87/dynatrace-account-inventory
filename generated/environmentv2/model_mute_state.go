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

// checks if the MuteState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MuteState{}

// MuteState Metadata of the muted state of a security problem in relation to an event.
type MuteState struct {
	// A user's comment.
	Comment *string `json:"comment,omitempty"`
	// The reason for the mute state change.
	Reason *string `json:"reason,omitempty"`
	// The user who has muted or unmuted the problem.
	User *string `json:"user,omitempty"`
}

// NewMuteState instantiates a new MuteState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMuteState() *MuteState {
	this := MuteState{}
	return &this
}

// NewMuteStateWithDefaults instantiates a new MuteState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMuteStateWithDefaults() *MuteState {
	this := MuteState{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *MuteState) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteState) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *MuteState) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *MuteState) SetComment(v string) {
	o.Comment = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *MuteState) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteState) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *MuteState) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *MuteState) SetReason(v string) {
	o.Reason = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MuteState) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MuteState) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MuteState) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *MuteState) SetUser(v string) {
	o.User = &v
}

func (o MuteState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MuteState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableMuteState struct {
	value *MuteState
	isSet bool
}

func (v NullableMuteState) Get() *MuteState {
	return v.value
}

func (v *NullableMuteState) Set(val *MuteState) {
	v.value = val
	v.isSet = true
}

func (v NullableMuteState) IsSet() bool {
	return v.isSet
}

func (v *NullableMuteState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMuteState(val *MuteState) *NullableMuteState {
	return &NullableMuteState{value: val, isSet: true}
}

func (v NullableMuteState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMuteState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

