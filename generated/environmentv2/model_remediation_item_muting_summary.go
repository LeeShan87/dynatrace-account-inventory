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

// checks if the RemediationItemMutingSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemediationItemMutingSummary{}

// RemediationItemMutingSummary Summary of (un-)muting a remediation item.
type RemediationItemMutingSummary struct {
	// Whether a mute state change for the given remediation item was triggered by this request.
	MuteStateChangeTriggered bool `json:"muteStateChangeTriggered"`
	// Contains a reason, in case the requested operation was not executed.
	Reason *string `json:"reason,omitempty"`
	// The id of the remediation item that will be (un-)muted.
	RemediationItemId string `json:"remediationItemId"`
}

// NewRemediationItemMutingSummary instantiates a new RemediationItemMutingSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemediationItemMutingSummary(muteStateChangeTriggered bool, remediationItemId string) *RemediationItemMutingSummary {
	this := RemediationItemMutingSummary{}
	this.MuteStateChangeTriggered = muteStateChangeTriggered
	this.RemediationItemId = remediationItemId
	return &this
}

// NewRemediationItemMutingSummaryWithDefaults instantiates a new RemediationItemMutingSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemediationItemMutingSummaryWithDefaults() *RemediationItemMutingSummary {
	this := RemediationItemMutingSummary{}
	return &this
}

// GetMuteStateChangeTriggered returns the MuteStateChangeTriggered field value
func (o *RemediationItemMutingSummary) GetMuteStateChangeTriggered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MuteStateChangeTriggered
}

// GetMuteStateChangeTriggeredOk returns a tuple with the MuteStateChangeTriggered field value
// and a boolean to check if the value has been set.
func (o *RemediationItemMutingSummary) GetMuteStateChangeTriggeredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MuteStateChangeTriggered, true
}

// SetMuteStateChangeTriggered sets field value
func (o *RemediationItemMutingSummary) SetMuteStateChangeTriggered(v bool) {
	o.MuteStateChangeTriggered = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *RemediationItemMutingSummary) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RemediationItemMutingSummary) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *RemediationItemMutingSummary) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *RemediationItemMutingSummary) SetReason(v string) {
	o.Reason = &v
}

// GetRemediationItemId returns the RemediationItemId field value
func (o *RemediationItemMutingSummary) GetRemediationItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemediationItemId
}

// GetRemediationItemIdOk returns a tuple with the RemediationItemId field value
// and a boolean to check if the value has been set.
func (o *RemediationItemMutingSummary) GetRemediationItemIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemediationItemId, true
}

// SetRemediationItemId sets field value
func (o *RemediationItemMutingSummary) SetRemediationItemId(v string) {
	o.RemediationItemId = v
}

func (o RemediationItemMutingSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemediationItemMutingSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["muteStateChangeTriggered"] = o.MuteStateChangeTriggered
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	toSerialize["remediationItemId"] = o.RemediationItemId
	return toSerialize, nil
}

type NullableRemediationItemMutingSummary struct {
	value *RemediationItemMutingSummary
	isSet bool
}

func (v NullableRemediationItemMutingSummary) Get() *RemediationItemMutingSummary {
	return v.value
}

func (v *NullableRemediationItemMutingSummary) Set(val *RemediationItemMutingSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableRemediationItemMutingSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableRemediationItemMutingSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemediationItemMutingSummary(val *RemediationItemMutingSummary) *NullableRemediationItemMutingSummary {
	return &NullableRemediationItemMutingSummary{value: val, isSet: true}
}

func (v NullableRemediationItemMutingSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemediationItemMutingSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


