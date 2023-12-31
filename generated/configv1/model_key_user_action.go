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

// checks if the KeyUserAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeyUserAction{}

// KeyUserAction Configuration of the key user action.
type KeyUserAction struct {
	// The type of the action.
	ActionType string `json:"actionType"`
	// The domain where the action is performed.
	Domain *string `json:"domain,omitempty"`
	// The Dynatrace entity ID of the action.
	MeIdentifier *string `json:"meIdentifier,omitempty"`
	// The name of the action.
	Name string `json:"name"`
}

// NewKeyUserAction instantiates a new KeyUserAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeyUserAction(actionType string, name string) *KeyUserAction {
	this := KeyUserAction{}
	this.ActionType = actionType
	this.Name = name
	return &this
}

// NewKeyUserActionWithDefaults instantiates a new KeyUserAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyUserActionWithDefaults() *KeyUserAction {
	this := KeyUserAction{}
	return &this
}

// GetActionType returns the ActionType field value
func (o *KeyUserAction) GetActionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *KeyUserAction) GetActionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *KeyUserAction) SetActionType(v string) {
	o.ActionType = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *KeyUserAction) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyUserAction) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *KeyUserAction) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *KeyUserAction) SetDomain(v string) {
	o.Domain = &v
}

// GetMeIdentifier returns the MeIdentifier field value if set, zero value otherwise.
func (o *KeyUserAction) GetMeIdentifier() string {
	if o == nil || IsNil(o.MeIdentifier) {
		var ret string
		return ret
	}
	return *o.MeIdentifier
}

// GetMeIdentifierOk returns a tuple with the MeIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeyUserAction) GetMeIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.MeIdentifier) {
		return nil, false
	}
	return o.MeIdentifier, true
}

// HasMeIdentifier returns a boolean if a field has been set.
func (o *KeyUserAction) HasMeIdentifier() bool {
	if o != nil && !IsNil(o.MeIdentifier) {
		return true
	}

	return false
}

// SetMeIdentifier gets a reference to the given string and assigns it to the MeIdentifier field.
func (o *KeyUserAction) SetMeIdentifier(v string) {
	o.MeIdentifier = &v
}

// GetName returns the Name field value
func (o *KeyUserAction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KeyUserAction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KeyUserAction) SetName(v string) {
	o.Name = v
}

func (o KeyUserAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KeyUserAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actionType"] = o.ActionType
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.MeIdentifier) {
		toSerialize["meIdentifier"] = o.MeIdentifier
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableKeyUserAction struct {
	value *KeyUserAction
	isSet bool
}

func (v NullableKeyUserAction) Get() *KeyUserAction {
	return v.value
}

func (v *NullableKeyUserAction) Set(val *KeyUserAction) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyUserAction) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyUserAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyUserAction(val *KeyUserAction) *NullableKeyUserAction {
	return &NullableKeyUserAction{value: val, isSet: true}
}

func (v NullableKeyUserAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyUserAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


