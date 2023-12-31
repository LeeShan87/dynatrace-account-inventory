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

// checks if the UserActionDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserActionDetails{}

// UserActionDetails Configuration of a user action-based conversion goal.
type UserActionDetails struct {
	// Type of the action to which the rule applies.
	ActionType *string `json:"actionType,omitempty"`
	// The match is case-sensitive (`true`) or (`false`).
	CaseSensitive *bool `json:"caseSensitive,omitempty"`
	// The type of the entity to which the rule applies.
	MatchEntity *string `json:"matchEntity,omitempty"`
	// The operator of the match.
	MatchType *string `json:"matchType,omitempty"`
	// The value to be matched to hit the conversion goal.
	Value *string `json:"value,omitempty"`
}

// NewUserActionDetails instantiates a new UserActionDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActionDetails() *UserActionDetails {
	this := UserActionDetails{}
	return &this
}

// NewUserActionDetailsWithDefaults instantiates a new UserActionDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActionDetailsWithDefaults() *UserActionDetails {
	this := UserActionDetails{}
	return &this
}

// GetActionType returns the ActionType field value if set, zero value otherwise.
func (o *UserActionDetails) GetActionType() string {
	if o == nil || IsNil(o.ActionType) {
		var ret string
		return ret
	}
	return *o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionDetails) GetActionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActionType) {
		return nil, false
	}
	return o.ActionType, true
}

// HasActionType returns a boolean if a field has been set.
func (o *UserActionDetails) HasActionType() bool {
	if o != nil && !IsNil(o.ActionType) {
		return true
	}

	return false
}

// SetActionType gets a reference to the given string and assigns it to the ActionType field.
func (o *UserActionDetails) SetActionType(v string) {
	o.ActionType = &v
}

// GetCaseSensitive returns the CaseSensitive field value if set, zero value otherwise.
func (o *UserActionDetails) GetCaseSensitive() bool {
	if o == nil || IsNil(o.CaseSensitive) {
		var ret bool
		return ret
	}
	return *o.CaseSensitive
}

// GetCaseSensitiveOk returns a tuple with the CaseSensitive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionDetails) GetCaseSensitiveOk() (*bool, bool) {
	if o == nil || IsNil(o.CaseSensitive) {
		return nil, false
	}
	return o.CaseSensitive, true
}

// HasCaseSensitive returns a boolean if a field has been set.
func (o *UserActionDetails) HasCaseSensitive() bool {
	if o != nil && !IsNil(o.CaseSensitive) {
		return true
	}

	return false
}

// SetCaseSensitive gets a reference to the given bool and assigns it to the CaseSensitive field.
func (o *UserActionDetails) SetCaseSensitive(v bool) {
	o.CaseSensitive = &v
}

// GetMatchEntity returns the MatchEntity field value if set, zero value otherwise.
func (o *UserActionDetails) GetMatchEntity() string {
	if o == nil || IsNil(o.MatchEntity) {
		var ret string
		return ret
	}
	return *o.MatchEntity
}

// GetMatchEntityOk returns a tuple with the MatchEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionDetails) GetMatchEntityOk() (*string, bool) {
	if o == nil || IsNil(o.MatchEntity) {
		return nil, false
	}
	return o.MatchEntity, true
}

// HasMatchEntity returns a boolean if a field has been set.
func (o *UserActionDetails) HasMatchEntity() bool {
	if o != nil && !IsNil(o.MatchEntity) {
		return true
	}

	return false
}

// SetMatchEntity gets a reference to the given string and assigns it to the MatchEntity field.
func (o *UserActionDetails) SetMatchEntity(v string) {
	o.MatchEntity = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *UserActionDetails) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionDetails) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *UserActionDetails) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *UserActionDetails) SetMatchType(v string) {
	o.MatchType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UserActionDetails) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActionDetails) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UserActionDetails) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *UserActionDetails) SetValue(v string) {
	o.Value = &v
}

func (o UserActionDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserActionDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActionType) {
		toSerialize["actionType"] = o.ActionType
	}
	if !IsNil(o.CaseSensitive) {
		toSerialize["caseSensitive"] = o.CaseSensitive
	}
	if !IsNil(o.MatchEntity) {
		toSerialize["matchEntity"] = o.MatchEntity
	}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableUserActionDetails struct {
	value *UserActionDetails
	isSet bool
}

func (v NullableUserActionDetails) Get() *UserActionDetails {
	return v.value
}

func (v *NullableUserActionDetails) Set(val *UserActionDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActionDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActionDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActionDetails(val *UserActionDetails) *NullableUserActionDetails {
	return &NullableUserActionDetails{value: val, isSet: true}
}

func (v NullableUserActionDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActionDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


