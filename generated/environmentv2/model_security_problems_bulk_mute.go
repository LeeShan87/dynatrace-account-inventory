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

// checks if the SecurityProblemsBulkMute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecurityProblemsBulkMute{}

// SecurityProblemsBulkMute Information on muting several security problems.
type SecurityProblemsBulkMute struct {
	// A comment about the muting reason.
	Comment *string `json:"comment,omitempty"`
	// The reason for muting the security problems.
	Reason string `json:"reason"`
	// The ids of the security problems to be muted.
	SecurityProblemIds []string `json:"securityProblemIds"`
}

// NewSecurityProblemsBulkMute instantiates a new SecurityProblemsBulkMute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityProblemsBulkMute(reason string, securityProblemIds []string) *SecurityProblemsBulkMute {
	this := SecurityProblemsBulkMute{}
	this.Reason = reason
	this.SecurityProblemIds = securityProblemIds
	return &this
}

// NewSecurityProblemsBulkMuteWithDefaults instantiates a new SecurityProblemsBulkMute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityProblemsBulkMuteWithDefaults() *SecurityProblemsBulkMute {
	this := SecurityProblemsBulkMute{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *SecurityProblemsBulkMute) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityProblemsBulkMute) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *SecurityProblemsBulkMute) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *SecurityProblemsBulkMute) SetComment(v string) {
	o.Comment = &v
}

// GetReason returns the Reason field value
func (o *SecurityProblemsBulkMute) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *SecurityProblemsBulkMute) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *SecurityProblemsBulkMute) SetReason(v string) {
	o.Reason = v
}

// GetSecurityProblemIds returns the SecurityProblemIds field value
func (o *SecurityProblemsBulkMute) GetSecurityProblemIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SecurityProblemIds
}

// GetSecurityProblemIdsOk returns a tuple with the SecurityProblemIds field value
// and a boolean to check if the value has been set.
func (o *SecurityProblemsBulkMute) GetSecurityProblemIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecurityProblemIds, true
}

// SetSecurityProblemIds sets field value
func (o *SecurityProblemsBulkMute) SetSecurityProblemIds(v []string) {
	o.SecurityProblemIds = v
}

func (o SecurityProblemsBulkMute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecurityProblemsBulkMute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	toSerialize["reason"] = o.Reason
	toSerialize["securityProblemIds"] = o.SecurityProblemIds
	return toSerialize, nil
}

type NullableSecurityProblemsBulkMute struct {
	value *SecurityProblemsBulkMute
	isSet bool
}

func (v NullableSecurityProblemsBulkMute) Get() *SecurityProblemsBulkMute {
	return v.value
}

func (v *NullableSecurityProblemsBulkMute) Set(val *SecurityProblemsBulkMute) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityProblemsBulkMute) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityProblemsBulkMute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityProblemsBulkMute(val *SecurityProblemsBulkMute) *NullableSecurityProblemsBulkMute {
	return &NullableSecurityProblemsBulkMute{value: val, isSet: true}
}

func (v NullableSecurityProblemsBulkMute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityProblemsBulkMute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

